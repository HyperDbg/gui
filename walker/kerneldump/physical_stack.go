package kerneldump

import (
	"encoding/binary"
	"fmt"
)

func (kd *KernelDump) GetStackTraceFromPhysical() ([]StackFrame, error) {
	var frames []StackFrame

	fmt.Printf("Scanning physical memory for stack (file size: %d bytes)...\n", len(kd.Data))

	for addr := uint64(0); addr < uint64(len(kd.Data))-16; addr += 8 {
		qword1 := binary.LittleEndian.Uint64(kd.Data[addr : addr+8])

		if qword1 == 0 {
			continue
		}

		if qword1 >= 0 && qword1 <= 0xFFFFFFFF {
			bugCheck := uint32(qword1)
			if bugCheck == 0xEF {
				frame := StackFrame{
					ChildSP:  addr,
					RetAddr:  qword1,
					FrameNum: len(frames),
					CallSite: fmt.Sprintf("BugCheck_0x%08X", bugCheck),
				}
				frames = append(frames, frame)
				return frames, nil
			}
		}
	}

	return frames, nil
}

func (kd *KernelDump) DetectBugCheckFromPhysical() (uint32, string) {
	frames, err := kd.GetStackTraceFromPhysical()
	if err != nil {
		return kd.Header.BugCheckCode, ""
	}

	for _, frame := range frames {
		if frame.RetAddr >= 0 && frame.RetAddr <= 0xFFFFFFFF {
			bugCheck := uint32(frame.RetAddr)
			if bugCheck != 0 && bugCheck != 0x8 && bugCheck != 0xFFFFFFFF {
				return bugCheck, kd.GetBugCheckNameFromCode(bugCheck)
			}
		}
	}

	return kd.Header.BugCheckCode, ""
}
