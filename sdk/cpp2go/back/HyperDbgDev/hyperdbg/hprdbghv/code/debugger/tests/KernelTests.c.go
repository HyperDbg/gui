package tests
type (
	KernelTests interface {
		TestKernelConfigureTagsAndCallTargetFunction() (ok bool) //col:72
		TestKernelPerformTests() (ok bool)                       //col:109
		TestKernelGetInformation() (ok bool)                     //col:164
	}
)
func NewKernelTests() { return &kernelTests{} }
func (k *kernelTests) TestKernelConfigureTagsAndCallTargetFunction() (ok bool) { //col:72
	return true
}

func (k *kernelTests) TestKernelPerformTests() (ok bool) { //col:109
	return true
}

func (k *kernelTests) TestKernelGetInformation() (ok bool) { //col:164
	return true
}

