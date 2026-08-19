package pe

import (
	"errors"
	"unsafe"
)

// Parse the debug directory.
func (p *PEFile) parseDebugDirectory(rva, size uint32) (err error) {
	if size == 0 {
		// Simply no debug directory exists.
		return nil
	}

	// Parse all debug directories.
	var offset uint32 = 0
	for offset < size {
		debugDir := NewDebugDirectory(p.getOffsetFromRva(rva + offset))
		start, _ := p.getDataBounds(rva+offset, 0)
		if err = p.parseInterface(&debugDir.ImageDebugDirectory, start, debugDir.size); err != nil {
			return err
		}

		if debugDir.SizeOfData != 0 {
			dataSize := int(debugDir.SizeOfData)
			dataOffset := int(debugDir.PointerToRawData)

			debugDir.RawData = make([]byte, debugDir.SizeOfData)
			copy(debugDir.RawData, p.data[dataOffset:dataOffset+dataSize])

			if debugDir.Type == IMAGE_DEBUG_TYPE_CODEVIEW {
				var omfSignature OMFSignature
				const omfSignatureSize = int(unsafe.Sizeof(omfSignature))

				// Check data length
				if len(debugDir.RawData) < omfSignatureSize {
					return errors.New("corrupt codeview data")
				}

				// Get OMF signature
				if err = p.parseInterface(&omfSignature, dataOffset, omfSignatureSize); err != nil {
					return err
				}

				if omfSignature.Signature == CV_PDB_70_SIGNATUE {
					// Signature is "RSDS" - PDB 7.0
					var cvInfoPdb CvInfoPdb70
					const cvInfoPdb70Size = int(unsafe.Sizeof(cvInfoPdb))

					// Check data length
					if dataSize < cvInfoPdb70Size {
						return errors.New("corrupt PDB 7.0 data")
					}

					// Get CV info
					if err = p.parseInterface(&cvInfoPdb, dataOffset, cvInfoPdb70Size); err != nil {
						return err
					}
					debugDir.InfoPdb70 = &cvInfoPdb

					// Get the symbol file name.
					debugDir.SymbolName = p.getStringFromData(dataOffset + cvInfoPdb70Size)

				} else if omfSignature.Signature == CV_PDB_20_SIGNATUE {
					// Signature is "NB10" - PDB 2.0
					var cvInfoPdb CvInfoPdb20
					const cvInfoPdb20Size = int(unsafe.Sizeof(cvInfoPdb))

					// Check data length
					if dataSize < cvInfoPdb20Size {
						return errors.New("corrupt PDB 2.0 data")
					}

					// Get CV info
					if err = p.parseInterface(&cvInfoPdb, dataOffset, cvInfoPdb20Size); err != nil {
						return err
					}
					debugDir.InfoPdb20 = &cvInfoPdb

					// Get the symbol file name.
					debugDir.SymbolName = p.getStringFromData(dataOffset + cvInfoPdb20Size)
				}
			}
		}

		// Is this safe?
		p.DebugDirectories = append(p.DebugDirectories, debugDir)

		offset += uint32(debugDir.size)
	}

	return nil
}
