package sdk

import (
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

type API interface {
	init()
	DetectVmxSupport() bool
	ReadVendorString(arg0 string)
	LoadVmm()
	UnloadVmm()
	InstallVmmDriver()
	UninstallVmmDriver()
	StopVmmDriver()
	TestCommandParser(command string, number_of_tokens uint32, tokens_list *string, failed_token_num *uint32, failed_token_position *uint32) bool
	TestCommandParserShowTokens(command string)
	SetupPathForFilename(filename string, file_location string, buffer_len uint32, check_file_existence bool) bool
	RunCommand(command string)
	ShowSignature()
	SetTextMessageCallback(handler uint64)
	SetTextMessageCallbackUsingSharedBuffer(handler uint64) uint64
	UnsetTextMessageCallback()
	ScriptReadFileAndExecuteCommandline(argc Int, argv *string)
	ContinuePreviousCommand() bool
	CheckMultilineCommand(current_command string, reset bool) bool
	SetCustomDriverPath(driver_file_path string, driver_name string) bool
	UseDefaultDriverPath()
	ConnectLocalDebugger()
	ConnectRemoteDebugger(ip string, port string) bool
	ConnectRemoteDebuggerUsingComPort(port_name string, baudrate Dword, pause_after_connection bool) bool
	ConnectRemoteDebuggerUsingNamedPipe(named_pipe string, pause_after_connection bool) bool
	ConnectCurrentDebuggerUsingComPort(port_name string, baudrate Dword) bool
	DebugCloseRemoteDebugger() bool
	GetKernelBase() uint64
	ReadMemory(target_address uint64, memory_type DebuggerReadMemoryType, reading_Type DebuggerReadReadingType, pid uint32, size uint32, get_address_mode bool, address_mode *DebuggerReadMemoryAddressMode, target_buffer_to_store []byte, return_length *uint32) bool
	ShowMemoryOrDisassemble(style DebuggerShowMemoryStyle, address uint64, memory_type DebuggerReadMemoryType, reading_type DebuggerReadReadingType, pid uint32, size uint32, dt_details PdebuggerDtCommandOptions)
	WriteMemory(destination_address uint64, memory_type DebuggerEditMemoryType, process_id uint32, source_address uint64, number_of_bytes uint32) bool
	ReadAllRegisters(guest_registers uint64, extra_registers uint64) bool
	ReadTargetRegister(register_id RegsEnum, target_register *uint64) bool
	WriteTargetRegister(register_id RegsEnum, value uint64) bool
	ShowAllRegisters() bool
	ShowTargetRegister(register_id RegsEnum) bool
	ContinueDebuggee()
	PauseDebuggee()
	SetBreakpoint(address uint64, pid uint32, tid uint32, core_numer uint32)
	SteppingInstrumentationStepIn() bool
	SteppingRegularStepIn() bool
	SteppingStepOver() bool
	SteppingInstrumentationStepInForTracking() bool
	SteppingStepOverForGu(last_instruction bool) bool
	StartProcess(path *Wchar) bool
	StartProcessWithArgs(path *Wchar, arguments *Wchar) bool
	CommandGetLocalApic(local_apic PlapicPage, is_using_x2apic *bool) bool
	AssembleGetLength(assembly_code string, start_address uint64, length *uint32) bool
	Assemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data uint64, buffer_size uint32) bool
	HwdbgScriptRunScript(script string, instance_filepath_to_read string, hardware_script_file_path_to_save string, initial_bram_buffer_size uint32) bool
	HwdbgScriptEngineWrapperTestParser(Expr string)
	HardwareScriptInterpreterShowScriptCapabilities(InstanceInfo *HwdbgInstanceInformation)
	HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(InstanceInfo *HwdbgInstanceInformation, ScriptBuffer uint64, CountOfScriptSymbolChunks uint32, NumberOfStages *uint32, NumberOfOperands *uint32, NumberOfOperandsImplemented *uint32) bool
	HardwareScriptInterpreterCompressBuffer(Buffer *uint64, BufferLength int32, ScriptVariableLength uint32, BramDataWidth uint32, NewBufferSize *int32, NumberOfBytesPerChunk *int32) bool
	HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(InstanceInfo *HwdbgInstanceInformation, SymbolBuffer uint64, SymbolBufferLength int32, NumberOfStages uint32, NewShortSymbolBuffer *uint64, NewBufferSize *int32) bool
	HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(NewShortSymbolBuffer uint64)
	ScriptEngineParse(str []byte) uint64
	ScriptEngineSetHwdbgInstanceInfo(InstancInfo *HwdbgInstanceInformation) bool
	PrintSymbolBuffer(SymbolBuffer uint64)
	RemoveSymbolBuffer(SymbolBuffer uint64)
	PrintSymbol(Symbol uint64)
	ScriptEngineConvertNameToAddress(FunctionOrVariableName []byte, WasFound bool) uint64
	ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName []byte, CustomModuleName []byte) uint32
	ScriptEngineUnloadAllSymbols() uint32
	ScriptEngineUnloadModuleSymbol(ModuleName []byte) uint32
	ScriptEngineSearchSymbolForMask(SearchMask []byte) uint32
	ScriptEngineGetFieldOffset(TypeName string, FieldName string, FieldOffset *uint32) bool
	ScriptEngineGetDataTypeSize(TypeName string, TypeSize *uint64) bool
	ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction uint64) bool
	ScriptEngineConvertFileToPdbPath(LocalFilePath []byte, ResultPath []byte) bool
	ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath []byte, PdbFilePath []byte, GuidAndAgeDetails []byte, Is32BitModule bool) bool
	ScriptEngineSymbolInitLoad(BufferToStoreDetails uint64, StoredLength uint32, DownloadIfAvailable bool, SymbolPath []byte, IsSilentLoad bool) bool
	ScriptEngineShowDataBasedOnSymbolTypes(TypeName []byte, Address uint64, IsStruct bool, BufferAddress uint64, AdditionalParameters []byte) bool
	ScriptEngineSymbolAbortLoading()
	ScriptEngineSetTextMessageCallback(Handler uint64)
	SymSetTextMessageCallback(Handler uint64)
	SymbolAbortLoading()
	SymConvertNameToAddress(FunctionOrVariableName []byte, WasFound bool) uint64
	SymLoadFileSymbol(BaseAddress uint64, PdbFileName []byte, CustomModuleName []byte) uint32
	SymUnloadAllSymbols() uint32
	SymUnloadModuleSymbol(ModuleName []byte) uint32
	SymSearchSymbolForMask(SearchMask []byte) uint32
	SymGetFieldOffset(TypeName string, FieldName string, FieldOffset *uint32) bool
	SymGetDataTypeSize(TypeName string, TypeSize *uint64) bool
	SymCreateSymbolTableForDisassembler(CallbackFunction uint64) bool
	SymConvertFileToPdbPath(LocalFilePath []byte, ResultPath []byte) bool
	SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath []byte, PdbFilePath []byte, GuidAndAgeDetails []byte, Is32BitModule bool) bool
	SymbolInitLoad(BufferToStoreDetails uint64, StoredLength uint32, DownloadIfAvailable bool, SymbolPath []byte, IsSilentLoad bool) bool
	SymShowDataBasedOnSymbolTypes(TypeName []byte, Address uint64, IsStruct bool, BufferAddress uint64, AdditionalParameters []byte) bool
	SymQuerySizeof(StructNameOrTypeName []byte, SizeOfField *uint32) bool
	SymCastingQueryForFiledsAndTypes(StructName []byte, FiledOfStructName []byte, IsStructNamePointerOrNot bool, IsFiledOfStructNamePointerOrNot bool, NewStructOrTypeName *[]byte, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) bool
}

type hyperDbg struct{}

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
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) DetectVmxSupport() bool {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadVendorString(arg0 string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) LoadVmm() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UnloadVmm() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) InstallVmmDriver() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UninstallVmmDriver() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StopVmmDriver() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) TestCommandParser(command string, number_of_tokens uint32, tokens_list *string, failed_token_num *uint32, failed_token_position *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) TestCommandParserShowTokens(command string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetupPathForFilename(filename string, file_location string, buffer_len uint32, check_file_existence bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) RunCommand(command string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowSignature() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetLogCallback(handler func()) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UnsetLogCallback() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptReadFileAndExecuteCommandline(argc Int, argv *string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ContinuePreviousCommand() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) CheckMultilineCommand(current_command string, reset bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetCustomDriverPath(driver_file_path string, driver_name string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) UseDefaultDriverPath() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectLocalDebugger() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebugger(ip string, port string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebuggerUsingComPort(port_name string, baudrate Dword, pause_after_connection bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectRemoteDebuggerUsingNamedPipe(named_pipe string, pause_after_connection bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ConnectCurrentDebuggerUsingComPort(port_name string, baudrate Dword) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) DebugCloseRemoteDebugger() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) GetKernelBase() uint64 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadMemory(target_address uint64, memory_type DebuggerReadMemoryType, reading_Type DebuggerReadReadingType, pid uint32, size uint32, get_address_mode bool, address_mode *DebuggerReadMemoryAddressMode, target_buffer_to_store []byte, return_length *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowMemoryOrDisassemble(style DebuggerShowMemoryStyle, address uint64, memory_type DebuggerReadMemoryType, reading_type DebuggerReadReadingType, pid uint32, size uint32, dt_details PdebuggerDtCommandOptions) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) WriteMemory(destination_address uint64, memory_type DebuggerEditMemoryType, process_id uint32, source_address uint64, number_of_bytes uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadAllRegisters(guest_registers uint64, extra_registers uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ReadTargetRegister(register_id RegsEnum, target_register *uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) WriteTargetRegister(register_id RegsEnum, value uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowAllRegisters() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ShowTargetRegister(register_id RegsEnum) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ContinueDebuggee() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) PauseDebuggee() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SetBreakpoint(address uint64, pid uint32, tid uint32, core_numer uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingInstrumentationStepIn() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingRegularStepIn() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingStepOver() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingInstrumentationStepInForTracking() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SteppingStepOverForGu(last_instruction bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StartProcess(path *Wchar) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) StartProcessWithArgs(path *Wchar, arguments *Wchar) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) CommandGetLocalApic(local_apic PlapicPage, is_using_x2apic *bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) AssembleGetLength(assembly_code string, start_address uint64, length *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) Assemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data uint64, buffer_size uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HwdbgScriptRunScript(script string, instance_filepath_to_read string, hardware_script_file_path_to_save string, initial_bram_buffer_size uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HwdbgScriptEngineWrapperTestParser(Expr string) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterShowScriptCapabilities(InstanceInfo *HwdbgInstanceInformation) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(InstanceInfo *HwdbgInstanceInformation, ScriptBuffer uint64, CountOfScriptSymbolChunks uint32, NumberOfStages *uint32, NumberOfOperands *uint32, NumberOfOperandsImplemented *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterCompressBuffer(Buffer *uint64, BufferLength int32, ScriptVariableLength uint32, BramDataWidth uint32, NewBufferSize *int32, NumberOfBytesPerChunk *int32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(InstanceInfo *HwdbgInstanceInformation, SymbolBuffer uint64, SymbolBufferLength int32, NumberOfStages uint32, NewShortSymbolBuffer *uint64, NewBufferSize *int32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(NewShortSymbolBuffer uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineParse(str []byte) uint64 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSetHwdbgInstanceInfo(InstancInfo *HwdbgInstanceInformation) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) PrintSymbolBuffer(SymbolBuffer uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) RemoveSymbolBuffer(SymbolBuffer uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) PrintSymbol(Symbol uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertNameToAddress(FunctionOrVariableName []byte, WasFound bool) uint64 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName []byte, CustomModuleName []byte) uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineUnloadAllSymbols() uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineUnloadModuleSymbol(ModuleName []byte) uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSearchSymbolForMask(SearchMask []byte) uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineGetFieldOffset(TypeName string, FieldName string, FieldOffset *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineGetDataTypeSize(TypeName string, TypeSize *uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertFileToPdbPath(LocalFilePath []byte, ResultPath []byte) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath []byte, PdbFilePath []byte, GuidAndAgeDetails []byte, Is32BitModule bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSymbolInitLoad(BufferToStoreDetails uint64, StoredLength uint32, DownloadIfAvailable bool, SymbolPath []byte, IsSilentLoad bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineShowDataBasedOnSymbolTypes(TypeName []byte, Address uint64, IsStruct bool, BufferAddress uint64, AdditionalParameters []byte) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSymbolAbortLoading() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) ScriptEngineSetTextMessageCallback(Handler uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymSetTextMessageCallback(Handler uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymbolAbortLoading() {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertNameToAddress(FunctionOrVariableName []byte, WasFound bool) uint64 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymLoadFileSymbol(BaseAddress uint64, PdbFileName []byte, CustomModuleName []byte) uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymUnloadAllSymbols() uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymUnloadModuleSymbol(ModuleName []byte) uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymSearchSymbolForMask(SearchMask []byte) uint32 {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymGetFieldOffset(TypeName string, FieldName string, FieldOffset *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymGetDataTypeSize(TypeName string, TypeSize *uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymCreateSymbolTableForDisassembler(CallbackFunction uint64) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertFileToPdbPath(LocalFilePath []byte, ResultPath []byte) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath []byte, PdbFilePath []byte, GuidAndAgeDetails []byte, Is32BitModule bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymbolInitLoad(BufferToStoreDetails uint64, StoredLength uint32, DownloadIfAvailable bool, SymbolPath []byte, IsSilentLoad bool) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymShowDataBasedOnSymbolTypes(TypeName []byte, Address uint64, IsStruct bool, BufferAddress uint64, AdditionalParameters []byte) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymQuerySizeof(StructNameOrTypeName []byte, SizeOfField *uint32) {
	// TODO implement me
	panic("implement me")
}

func (h *hyperDbg) SymCastingQueryForFiledsAndTypes(StructName []byte, FiledOfStructName []byte, IsStructNamePointerOrNot bool, IsFiledOfStructNamePointerOrNot bool, NewStructOrTypeName *[]byte, OffsetOfFieldFromTop *uint32, SizeOfField *uint32) {
	// TODO implement me
	panic("implement me")
}
