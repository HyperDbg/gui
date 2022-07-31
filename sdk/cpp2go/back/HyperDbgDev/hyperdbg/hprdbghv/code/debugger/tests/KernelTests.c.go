package tests


type (
KernelTests interface{
TestKernelConfigureTagsAndCallTargetFunction()(ok bool)//col:72
TestKernelPerformTests()(ok bool)//col:110
TestKernelGetInformation()(ok bool)//col:166
}






)

func NewKernelTests() { return & kernelTests{} }

func (k *kernelTests)TestKernelConfigureTagsAndCallTargetFunction()(ok bool){//col:72


























return true
}







func (k *kernelTests)TestKernelPerformTests()(ok bool){//col:110












return true
}







func (k *kernelTests)TestKernelGetInformation()(ok bool){//col:166




















return true
}









