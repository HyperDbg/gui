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
	TestCommandParser(command *Char, number_of_tokens uint32, tokens_list **Char, failed_token_num *uint32, failed_token_position *uint32) Boolean
	TestCommandParserShowTokens(command *Char)
	SetupPathForFilename(filename *Char, file_location *Char, buffer_len uint32, check_file_existence Boolean) Boolean
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
	ReadMemory(target_address Uint64, memory_type DebuggerReadMemoryType, reading_Type DebuggerReadReadingType, pid uint32, size uint32, get_address_mode Boolean, address_mode *DebuggerReadMemoryAddressMode, target_buffer_to_store *Byte, return_length *uint32) Boolean
	ShowMemoryOrDisassemble(style DebuggerShowMemoryStyle, address Uint64, memory_type DebuggerReadMemoryType, reading_type DebuggerReadReadingType, pid uint32, size uint32, dt_details PdebuggerDtCommandOptions)
	WriteMemory(destination_address unsafe.Pointer, memory_type DebuggerEditMemoryType, process_id uint32, source_address unsafe.Pointer, number_of_bytes uint32) Boolean
	ReadAllRegisters(guest_registers unsafe.Pointer, extra_registers unsafe.Pointer) Boolean
	ReadTargetRegister(register_id RegsEnum, target_register *Uint64) Boolean
	WriteTargetRegister(register_id RegsEnum, value Uint64) Boolean
	ShowAllRegisters() Boolean
	ShowTargetRegister(register_id RegsEnum) Boolean
	ContinueDebuggee()
	PauseDebuggee()
	SetBreakpoint(address Uint64, pid uint32, tid uint32, core_numer uint32)
	SteppingInstrumentationStepIn() Boolean
	SteppingRegularStepIn() Boolean
	SteppingStepOver() Boolean
	SteppingInstrumentationStepInForTracking() Boolean
	SteppingStepOverForGu(last_instruction Boolean) Boolean
	StartProcess(path *Wchar) Boolean
	StartProcessWithArgs(path *Wchar, arguments *Wchar) Boolean
	CommandGetLocalApic(local_apic PlapicPage, is_using_x2apic *Boolean) Boolean
	AssembleGetLength(assembly_code *Char, start_address Uint64, length *uint32) Boolean
	Assemble(assembly_code *Char, start_address Uint64, buffer_to_store_assembled_data unsafe.Pointer, buffer_size uint32) Boolean
	HwdbgScriptRunScript(script *Char, instance_filepath_to_read *Char, hardware_script_file_path_to_save *Char, initial_bram_buffer_size uint32) Boolean
	HwdbgScriptEngineWrapperTestParser(Expr *Char)
	HardwareScriptInterpreterShowScriptCapabilities(InstanceInfo *HwdbgInstanceInformation)
	HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(InstanceInfo *HwdbgInstanceInformation, ScriptBuffer unsafe.Pointer, CountOfScriptSymbolChunks uint32, NumberOfStages *uint32, NumberOfOperands *uint32, NumberOfOperandsImplemented *uint32) Boolean
	HardwareScriptInterpreterCompressBuffer(Buffer *Uint64, BufferLength int32, ScriptVariableLength uint32, BramDataWidth uint32, NewBufferSize *int32, NumberOfBytesPerChunk *int32) Boolean
	HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(InstanceInfo *HwdbgInstanceInformation, SymbolBuffer unsafe.Pointer, SymbolBufferLength int32, NumberOfStages uint32, NewShortSymbolBuffer *unsafe.Pointer, NewBufferSize *int32) Boolean
	HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(NewShortSymbolBuffer unsafe.Pointer)
	ScriptEngineParse(str *byte) unsafe.Pointer
	ScriptEngineSetHwdbgInstanceInfo(InstancInfo *HwdbgInstanceInformation) Boolean
	PrintSymbolBuffer(SymbolBuffer unsafe.Pointer)
	RemoveSymbolBuffer(SymbolBuffer unsafe.Pointer)
	PrintSymbol(Symbol unsafe.Pointer)
	ScriptEngineConvertNameToAddress(FunctionOrVariableName *byte, WasFound Pboolean) Uint64
	ScriptEngineLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) uint32
	ScriptEngineUnloadAllSymbols() uint32
	ScriptEngineUnloadModuleSymbol(ModuleName *byte) uint32
	ScriptEngineSearchSymbolForMask(SearchMask *byte) uint32
	ScriptEngineGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *uint32) Boolean
	ScriptEngineGetDataTypeSize(TypeName *Char, TypeSize *Uint64) Boolean
	ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) Boolean
	ScriptEngineConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) Boolean
	ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) Boolean
	ScriptEngineSymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) Boolean
	ScriptEngineShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) Boolean
	ScriptEngineSymbolAbortLoading()
	ScriptEngineSetTextMessageCallback(Handler unsafe.Pointer)
	SymSetTextMessageCallback(Handler unsafe.Pointer)
	SymbolAbortLoading()
	SymConvertNameToAddress(FunctionOrVariableName *byte, WasFound Pboolean) Uint64
	SymLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) uint32
	SymUnloadAllSymbols() uint32
	SymUnloadModuleSymbol(ModuleName *byte) uint32
	SymSearchSymbolForMask(SearchMask *byte) uint32
	SymGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *uint32) Boolean
	SymGetDataTypeSize(TypeName *Char, TypeSize *Uint64) Boolean
	SymCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) Boolean
	SymConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) Boolean
	SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) Boolean
	SymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) Boolean
	SymShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) Boolean
	SymQuerySizeof(StructNameOrTypeName *byte, SizeOfField *uint32) Boolean
	SymCastingQueryForFiledsAndTypes(StructName *byte, FiledOfStructName *byte, IsStructNamePointerOrNot Pboolean, IsFiledOfStructNamePointerOrNot Pboolean, NewStructOrTypeName **byte, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) Boolean
}

type hyperDbg struct {
}

func New() *hyperDbg {
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

func (h *hyperDbg) DetectVmxSupport() bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadVendorString(arg0 string) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) LoadVmm() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UnloadVmm() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) InstallVmmDriver() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UninstallVmmDriver() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StopVmmDriver() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) TestCommandParser(command *Char, number_of_tokens uint32, tokens_list **Char, failed_token_num *uint32, failed_token_position *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) TestCommandParserShowTokens(command *Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetupPathForFilename(filename *Char, file_location *Char, buffer_len uint32, check_file_existence Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) RunCommand(command *Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowSignature() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetLogCallback(handler func()) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UnsetLogCallback() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptReadFileAndExecuteCommandline(argc Int, argv **Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ContinuePreviousCommand() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) CheckMultilineCommand(current_command *Char, reset Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetCustomDriverPath(driver_file_path *Char, driver_name *Char) {
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

func (h *hyperDbg) ConnectRemoteDebugger(ip *Char, port *Char) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebuggerUsingComPort(port_name *Char, baudrate Dword, pause_after_connection Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebuggerUsingNamedPipe(named_pipe *Char, pause_after_connection Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectCurrentDebuggerUsingComPort(port_name *Char, baudrate Dword) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) DebugCloseRemoteDebugger() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) GetKernelBase() uint64 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadMemory(target_address Uint64, memory_type DebuggerReadMemoryType, reading_Type DebuggerReadReadingType, pid uint32, size uint32, get_address_mode Boolean, address_mode *DebuggerReadMemoryAddressMode, target_buffer_to_store *Byte, return_length *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowMemoryOrDisassemble(style DebuggerShowMemoryStyle, address Uint64, memory_type DebuggerReadMemoryType, reading_type DebuggerReadReadingType, pid uint32, size uint32, dt_details PdebuggerDtCommandOptions) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) WriteMemory(destination_address unsafe.Pointer, memory_type DebuggerEditMemoryType, process_id uint32, source_address unsafe.Pointer, number_of_bytes uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadAllRegisters(guest_registers unsafe.Pointer, extra_registers unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadTargetRegister(register_id RegsEnum, target_register *Uint64) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) WriteTargetRegister(register_id RegsEnum, value Uint64) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowAllRegisters() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowTargetRegister(register_id RegsEnum) {
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

func (h *hyperDbg) SetBreakpoint(address Uint64, pid uint32, tid uint32, core_numer uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingInstrumentationStepIn() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingRegularStepIn() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingStepOver() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingInstrumentationStepInForTracking() {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingStepOverForGu(last_instruction Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StartProcess(path *Wchar) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StartProcessWithArgs(path *Wchar, arguments *Wchar) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) CommandGetLocalApic(local_apic PlapicPage, is_using_x2apic *Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) AssembleGetLength(assembly_code *Char, start_address Uint64, length *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) Assemble(assembly_code *Char, start_address Uint64, buffer_to_store_assembled_data unsafe.Pointer, buffer_size uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HwdbgScriptRunScript(script *Char, instance_filepath_to_read *Char, hardware_script_file_path_to_save *Char, initial_bram_buffer_size uint32) {
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

func (h *hyperDbg) HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(InstanceInfo *HwdbgInstanceInformation, ScriptBuffer unsafe.Pointer, CountOfScriptSymbolChunks uint32, NumberOfStages *uint32, NumberOfOperands *uint32, NumberOfOperandsImplemented *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterCompressBuffer(Buffer *Uint64, BufferLength int32, ScriptVariableLength uint32, BramDataWidth uint32, NewBufferSize *int32, NumberOfBytesPerChunk *int32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(InstanceInfo *HwdbgInstanceInformation, SymbolBuffer unsafe.Pointer, SymbolBufferLength int32, NumberOfStages uint32, NewShortSymbolBuffer *unsafe.Pointer, NewBufferSize *int32) {
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

func (h *hyperDbg) ScriptEngineSetHwdbgInstanceInfo(InstancInfo *HwdbgInstanceInformation) {
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

func (h *hyperDbg) ScriptEngineLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineUnloadAllSymbols() uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineUnloadModuleSymbol(ModuleName *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSearchSymbolForMask(SearchMask *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineGetDataTypeSize(TypeName *Char, TypeSize *Uint64) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) {
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

func (h *hyperDbg) SymLoadFileSymbol(BaseAddress Uint64, PdbFileName *byte, CustomModuleName *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymUnloadAllSymbols() uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymUnloadModuleSymbol(ModuleName *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymSearchSymbolForMask(SearchMask *byte) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymGetFieldOffset(TypeName *Char, FieldName *Char, FieldOffset *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymGetDataTypeSize(TypeName *Char, TypeSize *Uint64) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymCreateSymbolTableForDisassembler(CallbackFunction unsafe.Pointer) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertFileToPdbPath(LocalFilePath *byte, ResultPath *byte) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath *byte, PdbFilePath *byte, GuidAndAgeDetails *byte, Is32BitModule Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymbolInitLoad(BufferToStoreDetails unsafe.Pointer, StoredLength uint32, DownloadIfAvailable Boolean, SymbolPath *byte, IsSilentLoad Boolean) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymShowDataBasedOnSymbolTypes(TypeName *byte, Address Uint64, IsStruct Boolean, BufferAddress unsafe.Pointer, AdditionalParameters *byte) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymQuerySizeof(StructNameOrTypeName *byte, SizeOfField *uint32) {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymCastingQueryForFiledsAndTypes(StructName *byte, FiledOfStructName *byte, IsStructNamePointerOrNot Pboolean, IsFiledOfStructNamePointerOrNot Pboolean, NewStructOrTypeName **byte, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) {
	//TODO implement me
	panic("implement me")
}
