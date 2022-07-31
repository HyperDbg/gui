package Source


const(
IMAGE_FILE_MACHINE_CHPE_X86 = 0x3A64 //col:12
PDBEX_VERSION_MAJOR = 0 //col:15
PDBEX_VERSION_MINOR = 18 //col:16
PDBEX_VERSION_STRING = "0.18" //col:18
)

type (
PdbExtractor interface{
		int Run()(ok bool)//col:98
}






)

func NewPdbExtractor() { return & pdbExtractor{} }

func (p *pdbExtractor)		int Run()(ok bool){//col:98









































return true
}









