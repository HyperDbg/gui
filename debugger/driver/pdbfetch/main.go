/**
 * pdbfetch
 * Copyright (c) 2019, Aidan Khoury. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @file main.go
 * @author Aidan Khoury (ajkhoury)
 * @date 11/13/2019
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ddkwork/golibrary/std/mylog"

	"github.com/ddkwork/HyperDbg/debugger/driver/pdbfetch/pkg/pe"
)

func probeWithUnderscore(path string) string {
	index := len(path) - 1
	return path[:index] + "_" + path[index+1:]
}

func probeWithFileptr(path string) string {
	index := strings.LastIndexByte(path, byte('/'))
	return path[:index] + "/file.ptr" + path[index+1:]
}

func retry(url string, head bool, probeMethod int) (*http.Response, error) {
	client := &http.Client{Timeout: 30 * time.Second}

	var method string
	if !head {
		method = http.MethodGet
	} else {
		method = http.MethodHead
	}

	var urlStr string
	switch probeMethod {
	case 1:
		urlStr = probeWithUnderscore(url)
	case 2:
		urlStr = probeWithFileptr(url)
	default:
		urlStr = url
	}

	req := mylog.Check2(http.NewRequest(method, urlStr, nil))

	const symbolServerUserAgent = "Microsoft-Symbol-Server/10.0.10522.521"
	req.Header.Set("User-Agent", symbolServerUserAgent)

	return client.Do(req)
}

func processFileptrResponse(resp *http.Response) (int64, string) {
	fileptrBytes := mylog.Check2(io.ReadAll(resp.Body))

	fileptrString := string(fileptrBytes)
	log.Println(fileptrString)

	if strings.Contains(fileptrString, "PATH:") {
		fileptrString = fileptrString[6:] // Removing "PATH:" from the output
		fileptrInfo := mylog.Check2(os.Stat(fileptrString))

		return fileptrInfo.Size(), fileptrString
	}

	return 0, ""
}

// downloadSymbolFile will download the symbol at the given url to a local filepath.
// TODO: Handle compressed symbols
func downloadSymbolFile(filepath string, url string) {
	var totalSize int64
	var probeMethod int

	// Probe 0
	probeMethod = 0
	resp := mylog.Check2(retry(url, true, probeMethod))

	defer resp.Body.Close()

	// Probe 1
	if resp.StatusCode == http.StatusNotFound {
		probeMethod = 1
		resp := mylog.Check2(retry(url, true, probeMethod))

		defer resp.Body.Close()
	}

	if resp.StatusCode == http.StatusOK {
		totalSize += resp.ContentLength
	}

	// Probe 2
	if resp.StatusCode == http.StatusNotFound {
		probeMethod = 2
		resp := mylog.Check2(retry(url, false, probeMethod))

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fileptrLength, fileptrPath := processFileptrResponse(resp)
			if fileptrLength > 0 {
				totalSize += fileptrLength
				url = fileptrPath
			}
		}
	}

	if totalSize == 0 {
		mylog.Check("symbol not found on server")
	}

	log.Println("PDB size:", totalSize/1024, "KiB")

	var reader io.Reader

	if probeMethod == 2 {
		f := mylog.Check2(os.Open(url))

		info := mylog.Check2(f.Stat())

		byteData := make([]byte, 0, info.Size())
		mylog.Check2(f.Read(byteData))

		reader = bytes.NewReader(byteData)

	} else {

		resp := mylog.Check2(retry(url, false, probeMethod))

		defer resp.Body.Close()

		reader = resp.Body
	}

	log.Println("Saving PDB to", filepath)

	// Create the file
	out := mylog.Check2(os.Create(filepath))

	defer out.Close()

	// Write the body to file
	mylog.Check2(io.Copy(out, reader))
}

func help() {
	fmt.Println("Fetches PDB symbol files directly from Microsoft's symbol servers.")
	fmt.Println("A PE file is supplied by the `pefile` argument.")
	fmt.Println("An optional save directory can be supplied by the `directory` argument.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("   ", filepath.Base(os.Args[0]), "pefile [directory]")
	fmt.Println("Example:")
	switch runtime.GOOS {
	case "windows":
		fmt.Println("   ", filepath.Base(os.Args[0]), "ExamplePE.exe \"C:\\symbols\"")
	case "linux":
		fmt.Println("   ", filepath.Base(os.Args[0]), "ExamplePE.exe \"/usr/share/symbols\"")
	}
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("   ", "-help", "\t", "display help information")
}

func main() {
	// Make sure the executable file path is given
	nArgs := len(os.Args)
	if nArgs < 2 {
		log.Println("executable file path not supplied")
		help()
		os.Exit(1)
	}
	// Check given args
	if os.Args[1] == "-help" || os.Args[1] == "--help" {
		help()
		os.Exit(0)
	}

	// PE path argument.
	pePath := os.Args[1]

	// Determine the symbol save directory.
	var saveDirectory string
	if nArgs > 2 {
		saveDirectory = os.Args[2]
	} else {
		absPath := mylog.Check2(filepath.Abs(os.Args[0]))

		saveDirectory = filepath.Dir(absPath)
	}

	// Check if the given PE path exists.
	mylog.Check2(os.Stat(pePath))

	// Parse the PE file of the target.
	pefile := mylog.Check2(pe.PE(pePath))

	// Enumerate the parsed debug directories in the PE.
	for _, debugDir := range pefile.DebugDirectories {
		// Only handle CV type debug directories.
		if debugDir.Type != pe.IMAGE_DEBUG_TYPE_CODEVIEW {
			continue
		}

		// Only PDB 7 files are supported.
		if debugDir.InfoPdb70 != nil {
			// Get the PDB GUID string.
			pdbGuid := pe.GuidFromWindowsArray(debugDir.InfoPdb70.Signature)
			pdbGuidStr := mylog.Check2(pdbGuid.ToString("N"))

			// Get the PDB filename.
			fullPdbPath := string(debugDir.SymbolName)
			pdbName := filepath.Base(fullPdbPath)

			// Get the PDB age string.
			pdbAgeStr := strconv.FormatUint(uint64(debugDir.InfoPdb70.Age), 16)

			// Compile the download URL for the symbol file.
			downloadURL := "http://msdl.microsoft.com/download/symbols/" + pdbName + "/" + strings.ToUpper(pdbGuidStr) + strings.ToUpper(pdbAgeStr) + "/" + pdbName

			log.Println(string(debugDir.SymbolName)+":", pdbGuid, pdbAgeStr, downloadURL)

			// Create directory for saving symbol file.
			downloadDir := filepath.Join(saveDirectory, pdbName, strings.ToUpper(pdbGuidStr)+strings.ToUpper(pdbAgeStr))
			mylog.Check(os.MkdirAll(downloadDir, 0o755))

			// Download the symbol file. FilePath concatenates guid and age
			downloadPath := filepath.Join(downloadDir, pdbName)
			(downloadSymbolFile(downloadPath, downloadURL))

		}
	}
}
