package examples


type (
ZydisWinKernel interface{
RtlPcToFileHeader()(ok bool)//col:88
DriverEntry()(ok bool)//col:188
}






)

func NewZydisWinKernel() { return & zydisWinKernel{} }

func (z *zydisWinKernel)RtlPcToFileHeader()(ok bool){//col:88
































return true
}







func (z *zydisWinKernel)DriverEntry()(ok bool){//col:188







































































return true
}









