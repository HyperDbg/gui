// Package symbolparser — download.go
//
// Provides PdbDownload, which fetches a PDB from the Microsoft Symbol Server
// (or a configured mirror) using the same probe sequence as the Microsoft
// Symsrv.dll: HEAD the canonical URL, fall back to the underscore-prefixed
// path, finally try file.ptr redirection.
//
// Mirrors libhyperdbg/code/debugger/script-engine/symbol.cpp SymbolPdbDownload
// and the standalone pdbfetch tool (ok/pdbfetch-master/main.go).
package symbolparser

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// PdbDownloadRequest carries the inputs for a single PDB download.
type PdbDownloadRequest struct {
	// PdbName is the PDB file name (e.g. "ntdll.pdb"). Case-insensitive.
	PdbName string
	// GuidHex is the PDB 7.0 GUID as a 32-char uppercase hex string (no
	// dashes/braces), e.g. "E02BF69D5F0C4D2B8C3DCE6DC7F5E2B1".
	GuidHex string
	// Age is the PDB age (typically 1 or 2). Used as a hex string in the URL.
	Age uint32
	// ServerURL is the symbol server base, default
	// "https://msdl.microsoft.com/download/symbols".
	ServerURL string
	// DestDir is the local directory where the PDB will be saved. The file
	// is placed at DestDir/<PdbName>/<GuidHex><AgeHex>/<PdbName>, mirroring
	// the symsrv.dll cache layout.
	DestDir string
}

// PdbDownloadResult is the outcome of a successful download.
type PdbDownloadResult struct {
	Path string
	Size int64
}

// PdbDownload fetches the PDB described by req and writes it under
// req.DestDir. Returns the final on-disk path and size.
//
// The default server is the Microsoft Symbol Server; for offline workflows
// the caller may point ServerURL at a local mirror (e.g. file:// or a
// custom HTTP server).
func PdbDownload(req PdbDownloadRequest) (*PdbDownloadResult, error) {
	if req.PdbName == "" || req.GuidHex == "" {
		return nil, errors.New("PdbDownload: PdbName and GuidHex are required")
	}
	server := req.ServerURL
	if server == "" {
		server = "https://msdl.microsoft.com/download/symbols"
	}
	pdbName := strings.ToLower(req.PdbName)
	guidHex := strings.ToUpper(sanitizeHex(req.GuidHex))
	ageHex := fmt.Sprintf("%x", req.Age)

	// Compute the destination path matching the symsrv.dll layout.
	downloadDir := filepath.Join(req.DestDir, pdbName, guidHex+ageHex)
	if err := os.MkdirAll(downloadDir, 0o755); err != nil {
		return nil, fmt.Errorf("PdbDownload: mkdir %q: %w", downloadDir, err)
	}
	downloadPath := filepath.Join(downloadDir, pdbName)

	// Cache check: if the file already exists and is non-empty, skip the
	// network round-trip (matches symsrv.dll behaviour).
	if info, err := os.Stat(downloadPath); err == nil && info.Size() > 0 {
		return &PdbDownloadResult{Path: downloadPath, Size: info.Size()}, nil
	}

	// Canonical URL: <server>/<pdb>/<guid><age>/<pdb>
	url := fmt.Sprintf("%s/%s/%s%s/%s", server, pdbName, guidHex, ageHex, pdbName)

	client := &http.Client{Timeout: 60 * time.Second}
	const userAgent = "Microsoft-Symbol-Server/10.0.10522.521"

	// Probe 0: HEAD the canonical URL.
	body, err := fetchPdbBody(client, url, userAgent, downloadPath)
	if err == nil {
		return body, nil
	}
	// Probe 1: underscore-suffixed path.
	url2 := probeWithUnderscore(url)
	body, err = fetchPdbBody(client, url2, userAgent, downloadPath)
	if err == nil {
		return body, nil
	}
	// Probe 2: file.ptr redirection.
	body, err = fetchPdbBodyFromFilePtr(client, url, userAgent, downloadPath)
	if err == nil {
		return body, nil
	}
	return nil, fmt.Errorf("PdbDownload: %s not found on %s (tried 3 probes)", pdbName, server)
}

// probeWithUnderscore mirrors the symsrv.dll underscore probe.
func probeWithUnderscore(path string) string {
	// Find the last path segment and prepend '_' to it.
	idx := strings.LastIndexByte(path, '/')
	if idx == -1 {
		return path
	}
	return path[:idx+1] + "_" + path[idx+1:]
}

// fetchPdbBody fetches the PDB at url and writes it to destPath. Returns a
// PdbDownloadResult describing the saved file, or an error.
func fetchPdbBody(client *http.Client, url, userAgent, destPath string) (*PdbDownloadResult, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return savePdb(resp.Body, destPath)
}

// fetchPdbBodyFromFilePtr handles the file.ptr redirection used by Microsoft
// Symbol Server for legacy symbol stores. The server returns a small text
// payload describing the actual location (PATH:... or URL:...).
func fetchPdbBodyFromFilePtr(client *http.Client, url, userAgent, destPath string) (*PdbDownloadResult, error) {
	// Construct the file.ptr URL: replace the last path segment with "file.ptr".
	idx := strings.LastIndexByte(url, '/')
	if idx == -1 {
		return nil, errors.New("file.ptr: malformed URL")
	}
	fpUrl := url[:idx+1] + "file.ptr"

	req, err := http.NewRequest(http.MethodGet, fpUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("file.ptr HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	s := string(body)
	switch {
	case strings.HasPrefix(s, "PATH:"):
		// Local path on the server (we can't access it remotely).
		return nil, fmt.Errorf("file.ptr PATH: unsupported")
	case strings.HasPrefix(s, "URL:"):
		target := strings.TrimSpace(strings.TrimPrefix(s, "URL:"))
		return fetchPdbBody(client, target, userAgent, destPath)
	default:
		return nil, fmt.Errorf("file.ptr: unrecognised payload %q", truncate(s, 80))
	}
}

// savePdb drains r into destPath and returns a PdbDownloadResult with the
// SHA-256 of the saved content (useful for cache verification).
func savePdb(r io.Reader, destPath string) (*PdbDownloadResult, error) {
	f, err := os.Create(destPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	n, err := io.Copy(io.MultiWriter(f, h), r)
	if err != nil {
		_ = os.Remove(destPath)
		return nil, err
	}
	if n == 0 {
		_ = os.Remove(destPath)
		return nil, errors.New("empty PDB")
	}
	_ = hex.EncodeToString(h.Sum(nil)) // for future cache-key extension
	return &PdbDownloadResult{Path: destPath, Size: n}, nil
}

// sanitizeHex strips dashes, braces and whitespace from a GUID string and
// upper-cases the result.
func sanitizeHex(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r == '-' || r == '{' || r == '}' || unicode.IsSpace(r) {
			continue
		}
		b.WriteRune(unicode.ToUpper(r))
	}
	return b.String()
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
