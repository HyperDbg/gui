package vmx
type (
	MsrHandlers interface {
		MsrHandleRdmsrVmexit() (ok bool)    //col:129
		MsrHandleWrmsrVmexit() (ok bool)    //col:239
		MsrHandleSetMsrBitmap() (ok bool)   //col:290
		MsrHandleUnSetMsrBitmap() (ok bool) //col:341
		*not valid or should be ignored ()
(ok bool) //col:363
* not valid or should be ignored ()(ok bool) //col:392
MsrHandlePerformMsrBitmapReadChange()(ok bool) //col:425
MsrHandlePerformMsrBitmapReadReset()(ok bool) //col:442
MsrHandlePerformMsrBitmapWriteChange()(ok bool) //col:474
MsrHandlePerformMsrBitmapWriteReset()(ok bool)  //col:491
}

)
func NewMsrHandlers() { return &msrHandlers{} }
func (m *msrHandlers) MsrHandleRdmsrVmexit() (ok bool) { //col:129
	return true
}

func (m *msrHandlers) MsrHandleWrmsrVmexit() (ok bool) { //col:239
	return true
}

func (m *msrHandlers) MsrHandleSetMsrBitmap() (ok bool) { //col:290
	return true
}

func (m *msrHandlers) MsrHandleUnSetMsrBitmap() (ok bool) { //col:341
	return true
}

func (m *msrHandlers) * not valid or should be ignored ()(ok bool) { //col:363
	return true
}

func (m *msrHandlers) * not valid or should be ignored ()(ok bool) { //col:392
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapReadChange() (ok bool) { //col:425
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapReadReset() (ok bool) { //col:442
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapWriteChange() (ok bool) { //col:474
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapWriteReset() (ok bool) { //col:491
	return true
}

