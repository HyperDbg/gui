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
	"errors"
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

	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	const symbolServerUserAgent = "Microsoft-Symbol-Server/10.0.10522.521"
	req.Header.Set("User-Agent", symbolServerUserAgent)

	return client.Do(req)
}

func processFileptrResponse(resp *http.Response) (int64, string) {
	fileptrBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fileptrString := string(fileptrBytes)
	log.Println(fileptrString)

	if strings.Contains(fileptrString, "PATH:") {
		fileptrString = fileptrString[6:] // Removing "PATH:" from the output
		fileptrInfo, err := os.Stat(fileptrString)
		if err != nil {
			return 0, ""
		}

		return fileptrInfo.Size(), fileptrString
	}

	return 0, ""
}

// downloadSymbolFile will download the symbol at the given url to a local filepath.
// TODO: Handle compressed symbols
func downloadSymbolFile(filepath string, url string) error {
	var totalSize int64
	var probeMethod int

	// Probe 0
	probeMethod = 0
	resp, err := retry(url, true, probeMethod)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Probe 1
	if resp.StatusCode == http.StatusNotFound {
		probeMethod = 1
		resp, err := retry(url, true, probeMethod)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode == http.StatusOK {
		totalSize += resp.ContentLength
	}

	// Probe 2
	if resp.StatusCode == http.StatusNotFound {
		probeMethod = 2
		resp, err := retry(url, false, probeMethod)
		if err != nil {
			return err
		}
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
		return errors.New("symbol not found on server")
	}

	log.Println("PDB size:", totalSize/1024, "KiB")

	var reader io.Reader

	if probeMethod == 2 {
		f, err := os.Open(url)
		if err != nil {
			return err
		}
		info, err := f.Stat()
		if err != nil {
			return err
		}

		byteData := make([]byte, 0, info.Size())
		_, err = f.Read(byteData)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(byteData)

	} else {

		resp, err := retry(url, false, probeMethod)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		reader = resp.Body
	}

	log.Println("Saving PDB to", filepath)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, reader)
	return err
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
		absPath, err := filepath.Abs(os.Args[0])
		if err != nil {
			log.Fatal(err)
		}
		saveDirectory = filepath.Dir(absPath)
	}

	// Check if the given PE path exists.
	if _, err := os.Stat(pePath); os.IsNotExist(err) {
		log.Fatal("pe file does not exist")
	}

	// Parse the PE file of the target.
	pefile, err := pe.PE(pePath)
	if err != nil {
		log.Fatal(err)
	}

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
			pdbGuidStr, err := pdbGuid.ToString("N")
			if err != nil {
				log.Fatal(err)
			}

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
			err = os.MkdirAll(downloadDir, 0o755)
			if err != nil {
				log.Fatal(err)
			}

			// Download the symbol file. FilePath concatenates guid and age
			downloadPath := filepath.Join(downloadDir, pdbName)
			err = downloadSymbolFile(downloadPath, downloadURL)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
