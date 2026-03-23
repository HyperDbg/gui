package kerneldump

type StackException struct {
	ExceptionCode    uint32
	ExceptionAddress uint64
	Found            bool
}

func (kd *KernelDump) ExtractExceptionFromStack() (*StackException, error) {
	frames, err := kd.GetStackTrace()
	if err != nil {
		return nil, err
	}

	for _, frame := range frames {
		if frame.RetAddr >= 0 && frame.RetAddr <= 0xFFFFFFFF {
			exceptCode := uint32(frame.RetAddr)

			if exceptCode >= 0x80000000 && exceptCode <= 0x800000FF {
				return &StackException{
					ExceptionCode:    exceptCode,
					ExceptionAddress: 0,
					Found:            true,
				}, nil
			}
		}
	}

	return &StackException{
		Found: false,
	}, nil
}

func (kd *KernelDump) GetExceptionFromStack() (uint32, uint64, bool) {
	except, err := kd.ExtractExceptionFromStack()
	if err != nil {
		return 0, 0, false
	}

	if except.Found {
		return except.ExceptionCode, except.ExceptionAddress, true
	}

	return 0, 0, false
}
