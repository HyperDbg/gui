package misc

const MaximumCallInstrSize = 8

type CallstackDisplayMethod int

const (
	CallstackDisplayMethodWithParams CallstackDisplayMethod = iota
	CallstackDisplayMethodWithoutParams
)

type CallstackFrame struct {
	Value                 uint64
	IsValidAddress        bool
	IsExecutable          bool
	InstructionBytesOnRip [MaximumCallInstrSize]byte
}

type FunctionDescription struct {
	BaseAddress uint64
	Size        uint64
	Name        string
}

func CallstackReturnAddressToCallingAddress(returnAddress []byte) (bool, uint32) {
	if len(returnAddress) < 7 {
		return false, 0
	}

	if returnAddress[len(returnAddress)-7] == 0x9A {
		return true, 7
	} else if len(returnAddress) >= 5 && returnAddress[len(returnAddress)-5] == 0xE8 {
		return true, 5
	} else if returnAddress[len(returnAddress)-2] == 0xFF {
		rmMask := uint8(0xF8)
		returnAddrLen := len(returnAddress)

		if returnAddrLen >= 7 &&
			returnAddress[returnAddrLen-7] == 0xFF &&
			(returnAddress[returnAddrLen-6] == 0x94 || returnAddress[returnAddrLen-6] == 0x9C) {
			return true, 7
		}

		if returnAddrLen >= 6 &&
			returnAddress[returnAddrLen-6] == 0xFF &&
			((returnAddress[returnAddrLen-5]&rmMask == 0x90 || returnAddress[returnAddrLen-5]&rmMask == 0x98) &&
				(returnAddress[returnAddrLen-5] != 0x94 && returnAddress[returnAddrLen-5] != 0x9C)) {
			return true, 6
		}

		if returnAddrLen >= 6 &&
			returnAddress[returnAddrLen-6] == 0xFF &&
			(returnAddress[returnAddrLen-5] == 0x15 || returnAddress[returnAddrLen-5] == 0x1D) {
			return true, 6
		}

		if returnAddrLen >= 4 &&
			returnAddress[returnAddrLen-4] == 0xFF &&
			(returnAddress[returnAddrLen-3] == 0x54 || returnAddress[returnAddrLen-3] == 0x5C) {
			return true, 4
		}

		if returnAddrLen >= 3 &&
			returnAddress[returnAddrLen-3] == 0xFF &&
			((returnAddress[returnAddrLen-2]&rmMask == 0x50 || returnAddress[returnAddrLen-2]&rmMask == 0x58) &&
				(returnAddress[returnAddrLen-2] != 0x54 && returnAddress[returnAddrLen-2] != 0x5C)) {
			return true, 3
		}

		if returnAddrLen >= 3 &&
			returnAddress[returnAddrLen-3] == 0xFF &&
			(returnAddress[returnAddrLen-2] == 0x14 || returnAddress[returnAddrLen-2] == 0x1C) {
			return true, 3
		}

		if returnAddrLen >= 2 &&
			returnAddress[returnAddrLen-2] == 0xFF &&
			(returnAddress[returnAddrLen-1]&rmMask == 0xD0 || returnAddress[returnAddrLen-1]&rmMask == 0xD8) {
			return true, 2
		}

		if returnAddrLen >= 2 &&
			returnAddress[returnAddrLen-2] == 0xFF &&
			((returnAddress[returnAddrLen-1]&rmMask == 0x10 || returnAddress[returnAddrLen-1]&rmMask == 0x18) &&
				(returnAddress[returnAddrLen-1] != 0x14 && returnAddress[returnAddrLen-1] != 0x15 &&
					returnAddress[returnAddrLen-1] != 0x1C && returnAddress[returnAddrLen-1] != 0x1D)) {
			return true, 2
		}

		return false, 0
	}

	return false, 0
}

func CallstackShowFrames(frames []CallstackFrame, frameCount uint32, displayMethod CallstackDisplayMethod, is32Bit bool) {
	for i := range frameCount {
		isCall := false
		var targetAddress uint64

		if frames[i].IsValidAddress {
			var callLength uint32
			instrBytes := frames[i].InstructionBytesOnRip[:]

			found, length := CallstackReturnAddressToCallingAddress(instrBytes)
			if frames[i].IsExecutable && found {
				if found {
					callLength = length
					targetAddress = frames[i].Value - uint64(callLength)
					isCall = true
				}
			} else {
				if displayMethod == CallstackDisplayMethodWithoutParams {
					continue
				}
				isCall = false
				targetAddress = frames[i].Value
			}

			offset := i * 4
			if !is32Bit {
				offset = i * 8
			}

			if isCall {
				if is32Bit {
					printf("[$+%03x]  %08x    (from ", offset, targetAddress)
				} else {
					printf("[$+%03x]  %016x    (from ", offset, targetAddress)
				}
			} else {
				if is32Bit {
					printf("[$+%03x]     %08x (addr ", offset, targetAddress)
				} else {
					printf("[$+%03x]     %016x (addr ", offset, targetAddress)
				}
			}

			if is32Bit {
				printf("<%08x>)\n", targetAddress)
			} else {
				printf("<%016x>)\n", targetAddress)
			}
		}
	}
}

func printf(format string, args ...any) {
}
