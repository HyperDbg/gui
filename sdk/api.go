package sdk

import "unsafe"

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
