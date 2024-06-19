#include <stdint.h>

__declspec(dllexport) int __stdcall UnpackPdataSection(char *MSNameOfProtected, char *MSNameOfDumped, char *MSWarning);

__declspec(dllexport) int __stdcall GetNameFileOptimized(char *MSFileNameOrig, char *MSFileNameOptimized);

__declspec(dllexport) int
__stdcall RebuildSectionsFromArmadillo(char *MSNameOfProtected, char *MSNameOfDumped, char *MSWarning);

__declspec(dllexport) int
__stdcall TryGetImportedFunction(uint32_t IRProcessId, uint32_t IRVAddress, uint32_t IROrdinal, uint32_t *IRHint,
                                 char *IRFunctionName, char *IRModule);

__declspec(dllexport) int
__stdcall SearchAndRebuildImportsNoNewSection(uint32_t IRProcessId, char *IRNameOfDumped, uint32_t IROEP,
                                              uint32_t IRSaveOEPToFile, uint32_t *IRIATRVA, uint32_t *IRIATSize,
                                              char *IRWarning);

__declspec(dllexport) int
__stdcall SearchAndRebuildImportsIATOptimized(uint32_t IRProcessId, char *IRNameOfDumped, uint32_t IROEP,
                                              uint32_t IRSaveOEPToFile, uint32_t *IRIATRVA, uint32_t *IRIATSize,
                                              char *IRWarning);

__declspec(dllexport) int
__stdcall SearchAndRebuildImports(uint32_t IRProcessId, char *IRNameOfDumped, uint32_t IROEP, uint32_t IRSaveOEPToFile,
                                  uint32_t *IRIATRVA, uint32_t *IRIATSize, char *IRWarning);

__declspec(dllexport) int
__stdcall GetProcNameAndOrdinal(uint32_t IRHModule, uint32_t IRAddress, uint32_t *IROrdinal, uint32_t *IRHint,
                                char *IRProcName);

__declspec(dllexport) uint32_t __stdcall GetProcOrdinal(uint32_t IRHModule, uint32_t IRAddress);

__declspec(dllexport) int
__stdcall GetProcName(uint32_t IRHModule, uint32_t IRAddress, uint32_t *IRHInt, char *IRProcName);

__declspec(dllexport) void GetAllVAddressesOfImports();//todo







 