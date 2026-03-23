package kerneldump

import (
	"encoding/binary"
	"fmt"
)

type PhysicalMemoryRun struct {
	BasePage  uint64
	PageCount uint64
}

type PhysicalMemoryBlock struct {
	NumberOfRuns uint64
	Runs        []PhysicalMemoryRun
}

func (kd *KernelDump) GetPhysicalMemoryBlock() (*PhysicalMemoryBlock, error) {
	if kd.Header.PhysicalMemoryBlock == 0 || kd.Header.PhysicalMemoryBlock >= uint64(len(kd.Data)) {
		return nil, fmt.Errorf("invalid physical memory block RVA")
	}

	offset := int(kd.Header.PhysicalMemoryBlock)
	if offset+16 > len(kd.Data) {
		return nil, fmt.Errorf("physical memory block exceeds file size")
	}

	pmb := &PhysicalMemoryBlock{
		NumberOfRuns: binary.LittleEndian.Uint64(kd.Data[offset : offset+8]),
	}

	for i := uint64(0); i < pmb.NumberOfRuns; i++ {
		runOffset := offset + 16 + int(i*24)
		if runOffset+24 > len(kd.Data) {
			break
		}

		run := PhysicalMemoryRun{
			BasePage:  binary.LittleEndian.Uint64(kd.Data[runOffset : runOffset+8]),
			PageCount: binary.LittleEndian.Uint64(kd.Data[runOffset+8 : runOffset+16]),
		}
		pmb.Runs = append(pmb.Runs, run)
	}

	return pmb, nil
}

func (kd *KernelDump) ReadMemory(physicalAddr uint64, size uint64) ([]byte, error) {
	pmb, err := kd.GetPhysicalMemoryBlock()
	if err != nil {
		return nil, err
	}

	pageNumber := physicalAddr / 4096
	offsetInPage := physicalAddr % 4096

	for _, run := range pmb.Runs {
		if pageNumber >= run.BasePage && pageNumber < run.BasePage+run.PageCount {
			runOffset := pageNumber - run.BasePage
			fileOffset := int(run.BasePage*4096 + runOffset*4096 + offsetInPage)

			if fileOffset+int(size) > len(kd.Data) {
				size = uint64(len(kd.Data) - fileOffset)
			}

			return kd.Data[fileOffset : fileOffset+int(size)], nil
		}
	}

	return nil, fmt.Errorf("physical address 0x%016X not found in dump", physicalAddr)
}
