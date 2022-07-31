package vmx


type (
MsrHandlers interface{
MsrHandleRdmsrVmexit()(ok bool)//col:129
MsrHandleWrmsrVmexit()(ok bool)//col:240
MsrHandleSetMsrBitmap()(ok bool)//col:292
MsrHandleUnSetMsrBitmap()(ok bool)//col:344
MsrHandleFilterMsrReadBitmap()(ok bool)//col:367
MsrHandleFilterMsrWriteBitmap()(ok bool)//col:397
MsrHandlePerformMsrBitmapReadChange()(ok bool)//col:431
MsrHandlePerformMsrBitmapReadReset()(ok bool)//col:449
MsrHandlePerformMsrBitmapWriteChange()(ok bool)//col:482
MsrHandlePerformMsrBitmapWriteReset()(ok bool)//col:500
}

)

func NewMsrHandlers() { return & msrHandlers{} }

func (m *msrHandlers)MsrHandleRdmsrVmexit()(ok bool){//col:129



















































return true
}


func (m *msrHandlers)MsrHandleWrmsrVmexit()(ok bool){//col:240























































return true
}


func (m *msrHandlers)MsrHandleSetMsrBitmap()(ok bool){//col:292



































return true
}


func (m *msrHandlers)MsrHandleUnSetMsrBitmap()(ok bool){//col:344



































return true
}


func (m *msrHandlers)MsrHandleFilterMsrReadBitmap()(ok bool){//col:367







return true
}


func (m *msrHandlers)MsrHandleFilterMsrWriteBitmap()(ok bool){//col:397









return true
}


func (m *msrHandlers)MsrHandlePerformMsrBitmapReadChange()(ok bool){//col:431














return true
}


func (m *msrHandlers)MsrHandlePerformMsrBitmapReadReset()(ok bool){//col:449






return true
}


func (m *msrHandlers)MsrHandlePerformMsrBitmapWriteChange()(ok bool){//col:482














return true
}


func (m *msrHandlers)MsrHandlePerformMsrBitmapWriteReset()(ok bool){//col:500






return true
}




