package pe

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// Parse the import directory.
//
// Given the RVA of the export directory, it will process all
// its entries.
//
// The exports will be made available as a list of ImportData
// instances in the ImportDescriptors PE attribute.
func (p *PEFile) parseImportDirectory(rva, size uint32) (err error) {
	p.ImportDescriptors = make([]*ImportDescriptor, 0)

	for {
		fileOffset := p.getOffsetFromRva(rva)
		importDesc := NewImportDescriptor(fileOffset)

		if err = p.parseInterface(&importDesc.ImageImportDescriptor, fileOffset, importDesc.size); err != nil {
			return err
		}

		//log.Printf("0x%x == %s", importDesc.Name, p.getStringAtRva(importDesc.Name))
		if EmptyStruct(importDesc.ImageImportDescriptor) {
			break
		}

		rva += uint32(importDesc.size)

		importDesc.Module = p.getStringAtRva(importDesc.Name)
		if !validDosFilename(importDesc.Module) {
			importDesc.Module = invalidImportName
		}

		if p.OptionalHeader64 != nil {
			if err := p.parseImports64(importDesc); err != nil {
				return err
			}
			// Give pretty names to well known dll files
			for _, imp := range importDesc.Imports64 {
				if imp.ImportByOrdinal {
					if funcname := OrdLookup(string(importDesc.Module), imp.Ordinal, false); len(funcname) > 0 {
						imp.Name = []byte(funcname)
					}
				}
			}
		} else {
			if err := p.parseImports(importDesc); err != nil {
				return err
			}
			// Give pretty names to well known dll files
			for _, imp := range importDesc.Imports {
				if imp.ImportByOrdinal {
					if funcname := OrdLookup(string(importDesc.Module), uint64(imp.Ordinal), false); len(funcname) > 0 {
						imp.Name = []byte(funcname)
					}
				}
			}
		}

		p.ImportDescriptors = append(p.ImportDescriptors, importDesc)
	}
	return nil
}

// Parse the imported symbols.
//
// It will fill a list, which will be available as the dictionary
// attribute "imports". Its keys will be the DLL names and the values
// all the symbols imported from that object.
func (p *PEFile) parseImports(importDesc *ImportDescriptor) (err error) {
	var table []*ThunkData32

	ilt, err := p.getImportTable(importDesc.Characteristics, importDesc)
	if err != nil {
		return err
	}

	iat, err := p.getImportTable(importDesc.FirstThunk, importDesc)
	if err != nil {
		return err
	}

	if len(iat) == 0 && len(ilt) == 0 {
		return errors.New("invalid import table information - both ILT and IAT appear to be broken")
	}

	impOffset := uint32(0x4)
	addressMask := uint32(0x7fffffff)
	ordinalFlag := IMAGE_ORDINAL_FLAG

	numInvalid := uint32(0)

	if len(ilt) > 0 {
		table = ilt
	} else {
		table = iat
	}

	for idx := uint32(0); idx < uint32(len(table)); idx++ {
		imp := new(ImportData32)
		imp.StructTable = table[idx]
		imp.OrdinalOffset = table[idx].fileOffset

		if table[idx].AddressOfData > 0 {

			// If imported by ordinal, we will append the ordinal numberx
			if table[idx].AddressOfData&ordinalFlag > 0 {
				imp.ImportByOrdinal = true
				imp.Ordinal = table[idx].AddressOfData & uint32(0xffff)
			} else {
				imp.ImportByOrdinal = false
				imp.HintNameTableRva = table[idx].AddressOfData & addressMask

				if err := p.parseInterface(&imp.Hint, int(imp.HintNameTableRva), 2); err != nil {
					return err
				}

				imp.Name = p.getStringAtRva(table[idx].AddressOfData + 2)

				if !validFuncName(imp.Name) {
					imp.Name = invalidImportName
				}
				imp.NameOffset = p.getOffsetFromRva(table[idx].AddressOfData + 2)
			}
			imp.ThunkOffset = table[idx].fileOffset
			imp.ThunkRva = p.getRvaFromOffset(imp.ThunkOffset)
		}

		imp.Address = importDesc.FirstThunk + p.OptionalHeader.ImageBase + (idx * impOffset)

		if len(iat) > 0 && len(ilt) > 0 && ilt[idx].AddressOfData != iat[idx].AddressOfData {
			imp.Bound = iat[idx].AddressOfData
			imp.StructIat = iat[idx]
		}

		hasName := len(imp.Name) > 0

		// The file with hashe:
		// SHA256: 3d22f8b001423cb460811ab4f4789f277b35838d45c62ec0454c877e7c82c7f5
		// has an invalid table built in a way that it's parseable but contains
		// invalid entries
		if imp.Ordinal == 0 && !hasName {
			return errors.New("must have either an ordinal or a name in an import")
		}

		// Some PEs appear to interleave valid and invalid imports. Instead of
		// aborting the parsing altogether, we will simply skip the invalid entries.
		// Although if we see 1000 invalid entries and no legitimate ones, we abort.
		if reflect.DeepEqual(imp.Name, invalidImportName) {
			if numInvalid > 1000 && numInvalid == idx {
				return errors.New("too many invalid names, aborting parsing")
			}
			numInvalid += 1
			continue
		}

		if imp.Ordinal > 0 || hasName {
			importDesc.Imports = append(importDesc.Imports, imp)
		}
	}

	return nil
}

func (p *PEFile) getImportTable(rva uint32, importDesc *ImportDescriptor) ([]*ThunkData32, error) {
	// Setup variables
	thunkTable := make(map[uint32]*ThunkData32)
	data := make([]*ThunkData32, 0)

	MaxAddressSpread := uint32(134217728) // 128 MB
	MaxRepeatedAddresses := uint32(16)

	ordinalFlag := IMAGE_ORDINAL_FLAG
	repeatedAddresses := uint32(0)
	startRva := rva

	minAddressOfData := ^uint32(0)
	maxAddressOfData := uint32(0)

	maxLen := uint32(p.dataLen - importDesc.fileOffset)
	if rva > importDesc.Characteristics || rva > importDesc.FirstThunk {
		maxLen = MaxUInt32(rva-importDesc.Characteristics, rva-importDesc.FirstThunk)
	}
	lastAddr := rva + maxLen

	// logic start
	for {
		if rva >= lastAddr {
			log.Println("Error parsing the import table. Entries go beyond bounds.")
			break
		}
		// if we see too many times the same entry we assume it could be
		// a table containing bogus data (with malicious intent or otherwise)
		if repeatedAddresses >= MaxRepeatedAddresses {
			return []*ThunkData32{}, errors.New("bogus data found in imports")
		}

		// if the addresses point somewhere but the difference between the highest
		// and lowest address is larger than MAX_ADDRESS_SPREAD we assume a bogus
		// table as the addresses should be contained within a module
		if maxAddressOfData-minAddressOfData > MaxAddressSpread {
			return []*ThunkData32{}, errors.New("data addresses too spread out")
		}

		thunk := NewThunkData32(p.getOffsetFromRva(rva))
		if err := p.parseInterface(&thunk.ImageThunkData32, thunk.fileOffset, thunk.size); err != nil {
			msg := fmt.Sprintf("Error Parsing the import table.\nInvalid data at RVA: 0x%x", rva)
			log.Println(msg)
			return []*ThunkData32{}, errors.New(msg)
		}
		if EmptyStruct(thunk.ImageThunkData32) {
			break
		}

		// Check if the AddressOfData lies within the range of RVAs that it's
		// being scanned, abort if that is the case, as it is very unlikely
		// to be legitimate data.
		// Seen in PE with SHA256:
		// 5945bb6f0ac879ddf61b1c284f3b8d20c06b228e75ae4f571fa87f5b9512902c
		if thunk.AddressOfData >= startRva && thunk.AddressOfData <= rva {
			log.Printf("Error parsing the import table. "+
				"AddressOfData overlaps with THUNK_DATA for THUNK at:\n  "+
				"RVA 0x%x", rva)
			break
		}

		if thunk.AddressOfData > 0 {
			// If the entry looks like could be an ordinal...
			if thunk.AddressOfData&ordinalFlag > 0 {
				// but its value is beyond 2^16, we will assume it's a
				// corrupted and ignore it altogether
				if thunk.AddressOfData&uint32(0x7fffffff) > uint32(0xffff) {
					msg := fmt.Sprintf("Corruption detected in thunk data at 0x%x", rva)
					log.Printf(msg)
					return []*ThunkData32{}, errors.New(msg)
				}
				// and if it looks like it should be an RVA
			} else {
				// keep track of the RVAs seen and store them to study their
				// properties. When certain non-standard features are detected
				// the parsing will be aborted
				if _, ok := thunkTable[rva]; ok {
					repeatedAddresses += 1
				}
				if thunk.AddressOfData > maxAddressOfData {
					maxAddressOfData = thunk.AddressOfData
				}
				if thunk.AddressOfData < minAddressOfData {
					minAddressOfData = thunk.AddressOfData
				}
			}
		}

		thunkTable[rva] = thunk
		data = append(data, thunk)
		rva += uint32(thunk.size)
	}

	return data, nil
}

func (p *PEFile) parseImports64(importDesc *ImportDescriptor) (err error) {
	// TODO: not implemented yet
	return nil
}

func (p *PEFile) getImportTable64(rva uint64) []*ThunkData64 {
	// TODO: not implemeted yet
	return []*ThunkData64{}
}
