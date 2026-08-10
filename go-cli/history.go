// history.go implements the in-memory and on-disk command history for the
// Go CLI REPL. It mirrors the per-user persistent history that shells like
// bash keep under ~/.bash_history; here the file lives at
// ~/.hyperdbg/history.
//
// The History type is goroutine-safe (the REPL reads/writes it from the
// main loop while signal handlers may also touch it), so every method
// guards its fields with a mutex — the same pattern used by
// go-libhyperdbg/debugger/userlevel/ud.go.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// defaultHistoryMax is the upper bound on the number of entries kept in
// memory and persisted to disk. It matches the order of magnitude used by
// common shells (bash defaults to 500/2000; we pick 1000).
const defaultHistoryMax = 1000

// History owns the REPL command history. All fields are guarded by mu.
type History struct {
	mu      sync.Mutex
	entries []string
	maxSize int
}

// NewHistory returns a History capped at maxSize entries. A maxSize <= 0
// means unbounded growth (use with care).
func NewHistory(maxSize int) *History {
	return &History{maxSize: maxSize}
}

// Add appends a line to the history. Empty lines and exact duplicates of
// the most recent entry are ignored (matches typical shell behaviour so
// pressing Enter repeatedly does not flood the history). When the buffer
// exceeds maxSize the oldest entries are dropped.
func (h *History) Add(line string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if line == "" {
		return
	}
	if n := len(h.entries); n > 0 && h.entries[n-1] == line {
		return
	}
	h.entries = append(h.entries, line)
	if h.maxSize > 0 && len(h.entries) > h.maxSize {
		h.entries = h.entries[len(h.entries)-h.maxSize:]
	}
}

// Get returns the nth entry using 1-based indexing (so Get(1) is the
// oldest command and Get(len) is the most recent). This matches the
// "!n" expansion syntax the REPL exposes. ok is false when n is out of
// range.
func (h *History) Get(n int) (string, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if n < 1 || n > len(h.entries) {
		return "", false
	}
	return h.entries[n-1], true
}

// List returns a copy of the history entries in chronological order
// (oldest first). Callers may iterate the result without holding the lock.
func (h *History) List() []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	out := make([]string, len(h.entries))
	copy(out, h.entries)
	return out
}

// Len returns the number of entries currently held.
func (h *History) Len() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return len(h.entries)
}

// Save writes the history to path, one entry per line. The parent
// directory is created with mode 0o755 if it does not already exist so
// that ~/.hyperdbg/history works on a fresh machine. A missing file is
// not an error — the call simply creates it.
func (h *History) Save(path string) error {
	if path == "" {
		return nil
	}
	h.mu.Lock()
	entries := make([]string, len(h.entries))
	copy(entries, h.entries)
	h.mu.Unlock()

	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("history save: %w", err)
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("history save: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, e := range entries {
		if _, err := fmt.Fprintln(w, e); err != nil {
			return fmt.Errorf("history save: %w", err)
		}
	}
	return w.Flush()
}

// Load replaces the in-memory history with the contents of path. A
// missing file is a silent no-op (first run); any other I/O error is
// returned wrapped. Entries past maxSize are trimmed.
func (h *History) Load(path string) error {
	if path == "" {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("history load: %w", err)
	}
	defer f.Close()

	var entries []string
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 4096), 1<<20)
	for sc.Scan() {
		entries = append(entries, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("history load: %w", err)
	}

	h.mu.Lock()
	h.entries = entries
	if h.maxSize > 0 && len(h.entries) > h.maxSize {
		h.entries = h.entries[len(h.entries)-h.maxSize:]
	}
	h.mu.Unlock()
	return nil
}
