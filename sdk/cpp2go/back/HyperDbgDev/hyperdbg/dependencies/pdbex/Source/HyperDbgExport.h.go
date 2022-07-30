package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\HyperDbgExport.h.back

const(
HYPERDBG_CODES =  //col:11
)

type (
HyperDbgExport interface{
typedef int ()(ok bool)//col:16
}
)

func NewHyperDbgExport() { return & hyperDbgExport{} }

func (h *hyperDbgExport)typedef int ()(ok bool){//col:16
/*typedef int (*Callback)(const char * Text);
#define HYPERDBG_CODES
extern "C" {
__declspec(dllexport) int pdbex_export(int argc, char ** argv, bool is_struct, void * buffer_address);
__declspec(dllexport) void pdbex_set_logging_method_export(PVOID handler);
}*/
return true
}



