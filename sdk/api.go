package sdk

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"unsafe"
)

type API interface {
	init()
	DetectVmxSupport() Boolean
	ReadVendorString(arg0 *Char)
	LoadVmm() Int
	UnloadVmm() Int
	InstallVmmDriver() Int
	UninstallVmmDriver() Int
	StopVmmDriver() Int
	TestCommandParser(command *Char, number_of_tokens Uint32, tokens_list **Char, failed_token_num *Uint32, failed_token_position *Uint32) Boolean
	TestCommandParserShowTokens(command *Char)
	SetupPathForFilename(filename *Char, file_location *Char, buffer_len Uint32, check_file_existence Boolean) Boolean
	RunCommand(command *Char) Int
	ShowSignature()
	SetTextMessageCallback(handler unsafe.Pointer)
	SetTextMessageCallbackUsingSharedBuffer(handler unsafe.Pointer) unsafe.Pointer
	UnsetTextMessageCallback()
	ScriptReadFileAndExecuteCommandline(argc Int, argv **Char) Int
	ContinuePreviousCommand() Boolean
	CheckMultilineCommand(current_command *Char, reset Boolean) Boolean
	SetCustomDriverPath(driver_file_path *Char, driver_name *Char) Boolean
	UseDefaultDriverPath()
	ConnectLocalDebugger()
	ConnectRemoteDebugger(ip *Char, port *Char) Boolean
	ConnectRemoteDebuggerUsingComPort(port_name *Char, baudrate Dword, pause_after_connection Boolean) Boolean
	ConnectRemoteDebuggerUsingNamedPipe(named_pipe *Char, pause_after_connection Boolean) Boolean
	ConnectCurrentDebuggerUsingComPort(port_name *Char, baudrate Dword) Boolean
	DebugCloseRemoteDebugger() Boolean
	GetKernelBase() Uint64
	ReadMemory(target_address Uint64, memory_type DebuggerReadMemoryType, reading_Type DebuggerReadReadingType, pid Uint32, size Uint32, get_address_mode Boolean, address_mode *DebuggerReadMemoryAddressMode, target_buffer_to_store *Byte, return_length *Uint32) Boolean
	ShowMemoryOrDisassemble(style DebuggerShowMemoryStyle, address Uint64, memory_type DebuggerReadMemoryType, reading_type DebuggerReadReadingType, pid Uint32, size Uint32, dt_details PdebuggerDtCommandOptions)
	WriteMemory(destination_address unsafe.Pointer, memory_type DebuggerEditMemoryType, process_id Uint32, source_address unsafe.Pointer, number_of_bytes Uint32) Boolean
	ReadAllRegisters(guest_registers unsafe.Pointer, extra_registers unsafe.Pointer) Boolean
	ReadTargetRegister(register_id RegsEnum, target_register *Uint64) Boolean
	WriteTargetRegister(register_id RegsEnum, value Uint64) Boolean
	ShowAllRegisters() Boolean
	ShowTargetRegister(register_id RegsEnum) Boolean
	ContinueDebuggee()
	PauseDebuggee()
	SetBreakpoint(address Uint64, pid Uint32, tid Uint32, core_numer Uint32)
	SteppingInstrumentationStepIn() Boolean
	SteppingRegularStepIn() Boolean
	SteppingStepOver() Boolean
	SteppingInstrumentationStepInForTracking() Boolean
	SteppingStepOverForGu(last_instruction Boolean) Boolean
	StartProcess(path *Wchar) Boolean
	StartProcessWithArgs(path *Wchar, arguments *Wchar) Boolean
	CommandGetLocalApic(local_apic PlapicPage, is_using_x2apic *Boolean) Boolean
	AssembleGetLength(assembly_code *Char, start_address Uint64, length *Uint32) Boolean
	Assemble(assembly_code *Char, start_address Uint64, buffer_to_store_assembled_data unsafe.Pointer, buffer_size Uint32) Boolean
	HwdbgScriptRunScript(script *Char, instance_filepath_to_read *Char, hardware_script_file_path_to_save *Char, initial_bram_buffer_size Uint32) Boolean
	HwdbgScriptEngineWrapperTestParser(Expr *Char)
	HardwareScriptInterpreterShowScriptCapabilities(InstanceInfo *HwdbgInstanceInformation)
	HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(InstanceInfo *HwdbgInstanceInformation, ScriptBuffer unsafe.Pointer, CountOfScriptSymbolChunks Uint32, NumberOfStages *Uint32, NumberOfOperands *Uint32, NumberOfOperandsImplemented *Uint32) Boolean
	HardwareScriptInterpreterCompressBuffer(Buffer *Uint64, BufferLength int32, ScriptVariableLength Uint32, BramDataWidth Uint32, NewBufferSize *int32, NumberOfBytesPerChunk *int32) Boolean
	HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(InstanceInfo *HwdbgInstanceInformation, SymbolBuffer unsafe.Pointer, SymbolBufferLength int32, NumberOfStages Uint32, NewShortSymbolBuffer *unsafe.Pointer, NewBufferSize *int32) Boolean
	HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(NewShortSymbolBuffer unsafe.Pointer)
	ScriptEngineParse(str *byte) unsafe.Pointer
	ScriptEngineSetHwdbgInstanceInfo(InstancInfo *HwdbgInstanceInformation) Boolean
	PrintSymbolBuffer(SymbolBuffer unsafe.Pointer)
	RemoveSymbolBuffer(SymbolBuffer unsafe.Pointer)
	PrintSymbol(Symbol unsafe.Pointer)
	ScriptEngineConvertNameToAddress(FunctionOrVariableName *byte, WasFound Pboolean) Uint64
	ScriptEngineLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) Uint32
	ScriptEngineUnloadAllSymbols() Uint32
	ScriptEngineUnloadModuleSymbol(ModuleName *byte) Uint32
	ScriptEngineSearchSymbolForMask(SearchMask *byte) Uint32
	ScriptEngineGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *Uint32) Boolean
	ScriptEngineGetDataTypeSize(TypeName *Char, TypeSize *Uint64) Boolean
	ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) Boolean
	ScriptEngineConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) Boolean
	ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) Boolean
	ScriptEngineSymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength Uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) Boolean
	ScriptEngineShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) Boolean
	ScriptEngineSymbolAbortLoading()
	ScriptEngineSetTextMessageCallback(Handler unsafe.Pointer)
	SymSetTextMessageCallback(Handler unsafe.Pointer)
	SymbolAbortLoading()
	SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound Pboolean) Uint64
	SymLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) Uint32
	SymUnloadAllSymbols() Uint32
	SymUnloadModuleSymbol(ModuleName *byte) Uint32
	SymSearchSymbolForMask(SearchMask *byte) Uint32
	SymGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *Uint32) Boolean
	SymGetDataTypeSize(TypeName *Char, TypeSize *Uint64) Boolean
	SymCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) Boolean
	SymConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) Boolean
	SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) Boolean
	SymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength Uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) Boolean
	SymShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) Boolean
	SymQuerySizeof(StructNameOrTypeName *byte, SizeOfField *Uint32) Boolean
	SymCastingQueryForFiledsAndTypes(StructName *byte, FiledOfStructName *byte, IsStructNamePointerOrNot Pboolean, IsFiledOfStructNamePointerOrNot Pboolean, NewStructOrTypeName **byte, OffsetOfFieldFromTop *Uint32, SizeOfField *Uint32) Boolean
}

type hyperDbg struct {
}

func New() API {
	return &hyperDbg{}
}

func SetCustomDriverPathEx(DriverFilePath string) bool {
	b := SetCustomDriverPath(StringToBytePointer(DriverFilePath), StringToBytePointer(stream.BaseName(DriverFilePath)))
	return Boolean2Bool(b)
}

func RunCommandEx(command string) (status string) {
	mylog.Info("command", command)
	code := RunCommand(StringToBytePointer(command))
	return string(code)
}

func ConnectRemoteDebuggerEx(ip string, port string) bool {
	return ConnectRemoteDebugger(StringToBytePointer(ip), StringToBytePointer(port)) == 0
}

func (h *hyperDbg) init() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) DetectVmxSupport() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadVendorString(arg0 *Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) LoadVmm() Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UnloadVmm() Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) InstallVmmDriver() Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UninstallVmmDriver() Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StopVmmDriver() Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) TestCommandParser(command *Char, number_of_tokens Uint32, tokens_list **Char, failed_token_num *Uint32, failed_token_position *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) TestCommandParserShowTokens(command *Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetupPathForFilename(filename *Char, file_location *Char, buffer_len Uint32, check_file_existence Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) RunCommand(command *Char) Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowSignature() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetTextMessageCallback(handler unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetTextMessageCallbackUsingSharedBuffer(handler unsafe.Pointer) unsafe.Pointer {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UnsetTextMessageCallback() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptReadFileAndExecuteCommandline(argc Int, argv **Char) Int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ContinuePreviousCommand() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) CheckMultilineCommand(current_command *Char, reset Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetCustomDriverPath(driver_file_path *Char, driver_name *Char) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UseDefaultDriverPath() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectLocalDebugger() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebugger(ip *Char, port *Char) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebuggerUsingComPort(port_name *Char, baudrate Dword, pause_after_connection Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebuggerUsingNamedPipe(named_pipe *Char, pause_after_connection Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectCurrentDebuggerUsingComPort(port_name *Char, baudrate Dword) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) DebugCloseRemoteDebugger() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) GetKernelBase() Uint64 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadMemory(target_address Uint64, memory_type DebuggerReadMemoryType, reading_Type DebuggerReadReadingType, pid Uint32, size Uint32, get_address_mode Boolean, address_mode *DebuggerReadMemoryAddressMode, target_buffer_to_store *Byte, return_length *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowMemoryOrDisassemble(style DebuggerShowMemoryStyle, address Uint64, memory_type DebuggerReadMemoryType, reading_type DebuggerReadReadingType, pid Uint32, size Uint32, dt_details PdebuggerDtCommandOptions) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) WriteMemory(destination_address unsafe.Pointer, memory_type DebuggerEditMemoryType, process_id Uint32, source_address unsafe.Pointer, number_of_bytes Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadAllRegisters(guest_registers unsafe.Pointer, extra_registers unsafe.Pointer) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadTargetRegister(register_id RegsEnum, target_register *Uint64) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) WriteTargetRegister(register_id RegsEnum, value Uint64) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowAllRegisters() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowTargetRegister(register_id RegsEnum) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ContinueDebuggee() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) PauseDebuggee() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetBreakpoint(address Uint64, pid Uint32, tid Uint32, core_numer Uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingInstrumentationStepIn() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingRegularStepIn() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingStepOver() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingInstrumentationStepInForTracking() Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingStepOverForGu(last_instruction Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StartProcess(path *Wchar) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StartProcessWithArgs(path *Wchar, arguments *Wchar) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) CommandGetLocalApic(local_apic PlapicPage, is_using_x2apic *Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) AssembleGetLength(assembly_code *Char, start_address Uint64, length *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) Assemble(assembly_code *Char, start_address Uint64, buffer_to_store_assembled_data unsafe.Pointer, buffer_size Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HwdbgScriptRunScript(script *Char, instance_filepath_to_read *Char, hardware_script_file_path_to_save *Char, initial_bram_buffer_size Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HwdbgScriptEngineWrapperTestParser(Expr *Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterShowScriptCapabilities(InstanceInfo *HwdbgInstanceInformation) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(InstanceInfo *HwdbgInstanceInformation, ScriptBuffer unsafe.Pointer, CountOfScriptSymbolChunks Uint32, NumberOfStages *Uint32, NumberOfOperands *Uint32, NumberOfOperandsImplemented *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterCompressBuffer(Buffer *Uint64, BufferLength int32, ScriptVariableLength Uint32, BramDataWidth Uint32, NewBufferSize *int32, NumberOfBytesPerChunk *int32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(InstanceInfo *HwdbgInstanceInformation, SymbolBuffer unsafe.Pointer, SymbolBufferLength int32, NumberOfStages Uint32, NewShortSymbolBuffer *unsafe.Pointer, NewBufferSize *int32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(NewShortSymbolBuffer unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineParse(str *byte) unsafe.Pointer {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSetHwdbgInstanceInfo(InstancInfo *HwdbgInstanceInformation) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) PrintSymbolBuffer(SymbolBuffer unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) RemoveSymbolBuffer(SymbolBuffer unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) PrintSymbol(Symbol unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertNameToAddress(FunctionOrVariableName *byte, WasFound Pboolean) Uint64 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineUnloadAllSymbols() Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineUnloadModuleSymbol(ModuleName *byte) Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSearchSymbolForMask(SearchMask *byte) Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineGetDataTypeSize(TypeName *Char, TypeSize *Uint64) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength Uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSymbolAbortLoading() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSetTextMessageCallback(Handler unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymSetTextMessageCallback(Handler unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymbolAbortLoading() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound Pboolean) Uint64 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymUnloadAllSymbols() Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymUnloadModuleSymbol(ModuleName *byte) Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymSearchSymbolForMask(SearchMask *byte) Uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymGetDataTypeSize(TypeName *Char, TypeSize *Uint64) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength Uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymQuerySizeof(StructNameOrTypeName *byte, SizeOfField *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymCastingQueryForFiledsAndTypes(StructName *byte, FiledOfStructName *byte, IsStructNamePointerOrNot Pboolean, IsFiledOfStructNamePointerOrNot Pboolean, NewStructOrTypeName **byte, OffsetOfFieldFromTop *Uint32, SizeOfField *Uint32) Boolean {
	//TODO implement me
	panic("implement me")
}
