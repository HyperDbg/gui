package pe

import (
	"fmt"
	"log"
)

// Parse the export directory.
//
// Given the RVA of the export directory, it will process all
// its entries.
//
// The exports will be made available as a list of ExportData
// instances in the ExportDescriptors PE attribute.
func (p *PEFile) parseExportDirectory(rva, size uint32) (err error) {
	exportDir := NewExportDirectory(p.getOffsetFromRva(rva))
	start, _ := p.getDataBounds(rva, 0)
	if err = p.parseInterface(&exportDir.ImageExportDirectory, start, exportDir.size); err != nil {
		return err
	}
	p.ExportDirectory = exportDir

	//log.Println(exportDir)
	startAddrOfNames, _ := p.getDataBounds(exportDir.AddressOfNames, 0)
	startAddrOfOrdinals, _ := p.getDataBounds(exportDir.AddressOfNameOrdinals, 0)
	startAddrOfFuncs, _ := p.getDataBounds(exportDir.AddressOfFunctions, 0)

	errMsg := "RVA %s in the export directory points to an invalid address: %x"
	//maxErrors := 10

	section := p.getSectionByRva(exportDir.AddressOfNames)
	if section == nil {
		log.Printf(errMsg, "AddressOfNames", exportDir.AddressOfNames)
		return fmt.Errorf(errMsg, "AddressOfNames", exportDir.AddressOfNames)
	}

	safetyBoundary := section.VirtualAddress + section.SizeOfRawData - exportDir.AddressOfNames
	numberOfNames := MinUInt32(safetyBoundary/4, exportDir.NumberOfNames)

	// A hash set for tracking seen ordinals
	ordMap := make(map[uint16]bool)

	for i := uint32(0); i < numberOfNames; i++ {
		sym := new(ExportData)

		// Name and name offset
		var symNameAddr uint32
		sym.NameOffset = startAddrOfNames + (int(i) * 4)
		if err = p.parseInterface(&symNameAddr, sym.NameOffset, 4); err != nil {
			return err
		}
		sym.Name = p.getStringAtRva(symNameAddr)
		//log.Printf("%s\n", sym.Name)
		if !validFuncName(sym.Name) {
			break
		}
		sym.NameOffset = p.getOffsetFromRva(symNameAddr)

		// Ordinal
		sym.OrdinalOffset = startAddrOfOrdinals + (int(i) * 2)
		if err = p.parseInterface(&sym.Ordinal, sym.OrdinalOffset, 2); err != nil {
			return err
		}

		// Address
		sym.AddressOffset = startAddrOfFuncs + (int(sym.Ordinal) * 4)
		if err = p.parseInterface(&sym.Address, sym.AddressOffset, 4); err != nil {
			return err
		}
		if sym.Address == 0 {
			continue
		}

		// Forwarder if applicable
		if sym.Address >= rva && sym.Address < rva+size {
			sym.Forwarder = p.getStringAtRva(sym.Address)
			sym.ForwarderOffset = p.getOffsetFromRva(sym.Address)
		}

		sym.Ordinal += uint16(exportDir.Base)

		ordMap[sym.Ordinal] = true
		exportDir.Exports = append(exportDir.Exports, sym)
	}

	// Check for any missing function symbols
	section = p.getSectionByRva(exportDir.AddressOfFunctions)
	if section == nil {
		log.Printf(errMsg, "AddressOfFunctions", exportDir.AddressOfFunctions)
		return fmt.Errorf(errMsg, "AddressOfFunctions", exportDir.AddressOfFunctions)
	}

	safetyBoundary = section.VirtualAddress + section.SizeOfRawData - exportDir.AddressOfFunctions
	numberOfNames = MinUInt32(safetyBoundary/4, exportDir.NumberOfFunctions)

	for i := uint32(0); i < numberOfNames; i++ {
		if _, ok := ordMap[uint16(i+exportDir.Base)]; ok {
			continue
		}

		sym := new(ExportData)

		// Address
		sym.AddressOffset = startAddrOfFuncs + (int(sym.Ordinal) * 4)
		if err = p.parseInterface(&sym.Address, sym.AddressOffset, 4); err != nil {
			return err
		}
		if sym.Address == 0 {
			continue
		}

		// Forwarder if applicable
		if sym.Address >= rva && sym.Address < rva+size {
			sym.Forwarder = p.getStringAtRva(sym.Address)
			sym.ForwarderOffset = p.getOffsetFromRva(sym.Address)
		}

		sym.Ordinal = uint16(exportDir.Base + i)

		exportDir.Exports = append(exportDir.Exports, sym)
	}

	return nil
}
