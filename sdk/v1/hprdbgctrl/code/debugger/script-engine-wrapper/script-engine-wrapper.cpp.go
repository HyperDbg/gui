package script_engine_wrapper
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\script-engine-wrapper\script-engine-wrapper.cpp.back

type  _ALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING struct{
Buff1 *int8 //col:11
Buff2 *int8 //col:12
Buff3 *int8 //col:13
Buff4 *int8 //col:14
Buff5 *int8 //col:15
Buff6 *int8 //col:16
}


type     typedef struct _UNICODE_STRING struct{
Length uint16 //col:17
MaximumLength uint16 //col:18
Buffer *uint32 //col:19
}


type     typedef struct _STUPID_STRUCT1 struct{
Flag32 uint32 //col:24
Flag64 uint64 //col:25
Context uintptr //col:26
StringValue *uint32 //col:27
}


type     typedef struct _STUPID_STRUCT2 struct{
Sina32 uint32 //col:32
Sina64 uint64 //col:33
AghaaSina uintptr //col:34
UnicodeStr *uint32 //col:35
StupidStruct1 PSTUPID_STRUCT1 //col:36
}


type     typedef struct _TEST_STRUCT struct{
Var1 uint64 //col:74
Var2 uint64 //col:75
Var3 uint64 //col:76
Var4 uint64 //col:77
}



type (
ScriptEngineWrapper interface{
ScriptEngineConvertNameToAddressWrapper()(ok bool)//col:125
AllocateStructForCasting()(ok bool)//col:140
ScriptEngineWrapperGetSize()(ok bool)//col:148
}
scriptEngineWrapper struct{}
)

func NewScriptEngineWrapper()ScriptEngineWrapper{ return & scriptEngineWrapper{} }

func (s *scriptEngineWrapper)ScriptEngineConvertNameToAddressWrapper()(ok bool){//col:125
/*ScriptEngineConvertNameToAddressWrapper(const char * FunctionOrVariableName, PBOOLEAN WasFound)
{
    return ScriptEngineConvertNameToAddress(FunctionOrVariableName, WasFound);
ScriptEngineLoadFileSymbolWrapper(UINT64 BaseAddress, const char * PdbFileName)
{
    return ScriptEngineLoadFileSymbol(BaseAddress, PdbFileName);
ScriptEngineSetTextMessageCallbackWrapper(PVOID Handler)
{
    return ScriptEngineSetTextMessageCallback(Handler);
ScriptEngineUnloadAllSymbolsWrapper()
{
    return ScriptEngineUnloadAllSymbols();
ScriptEngineUnloadModuleSymbolWrapper(char * ModuleName)
{
    return ScriptEngineUnloadModuleSymbol(ModuleName);
ScriptEngineSearchSymbolForMaskWrapper(const char * SearchMask)
{
    return ScriptEngineSearchSymbolForMask(SearchMask);
ScriptEngineGetFieldOffsetWrapper(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)
{
    return ScriptEngineGetFieldOffset(TypeName, FieldName, FieldOffset);
ScriptEngineGetDataTypeSizeWrapper(CHAR * TypeName, UINT64 * TypeSize)
{
    return ScriptEngineGetDataTypeSize(TypeName, TypeSize);
ScriptEngineCreateSymbolTableForDisassemblerWrapper(void * CallbackFunction)
{
    return ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction);
ScriptEngineConvertFileToPdbPathWrapper(const char * LocalFilePath, char * ResultPath)
{
    return ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath);
ScriptEngineSymbolInitLoadWrapper(PMODULE_SYMBOL_DETAIL BufferToStoreDetails,
                                  UINT32                StoredLength,
                                  BOOLEAN               DownloadIfAvailable,
                                  const char *          SymbolPath,
                                  BOOLEAN               IsSilentLoad)
{
    return ScriptEngineSymbolInitLoad(BufferToStoreDetails, StoredLength, DownloadIfAvailable, SymbolPath, IsSilentLoad);
ScriptEngineShowDataBasedOnSymbolTypesWrapper(
    const char * TypeName,
    UINT64       Address,
    BOOLEAN      IsStruct,
    PVOID        BufferAddress,
    const char * AdditionalParameters)
{
    return ScriptEngineShowDataBasedOnSymbolTypes(TypeName, Address, IsStruct, BufferAddress, AdditionalParameters);
ScriptEngineSymbolAbortLoadingWrapper()
{
    return ScriptEngineSymbolAbortLoading();
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(const char * LocalFilePath,
                                                            char *       PdbFilePath,
                                                            char *       GuidAndAgeDetails)
{
    return ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails);
ScriptEngineParseWrapper(char * Expr, BOOLEAN ShowErrorMessageIfAny)
{
    PSYMBOL_BUFFER SymbolBuffer;
    SymbolBuffer = ScriptEngineParse(Expr);
    if (SymbolBuffer->Message == NULL)
    {
        return SymbolBuffer;
    }
    else
    {
        if (ShowErrorMessageIfAny)
        {
            ShowMessages("%s\n", SymbolBuffer->Message);
        ScriptEngineWrapperRemoveSymbolBuffer(SymbolBuffer);
PrintSymbolBufferWrapper(PVOID SymbolBuffer)
{
    PrintSymbolBuffer((PSYMBOL_BUFFER)SymbolBuffer);
ScriptEngineEvalWrapper(PGUEST_REGS GuestRegs,
                        string      Expr)
{
    SCRIPT_ENGINE_VARIABLES_LIST VariablesList = {0};
    if (!g_ScriptGlobalVariables)
    {
        g_ScriptGlobalVariables = (UINT64 *)malloc(MAX_VAR_COUNT * sizeof(UINT64));
        RtlZeroMemory(g_ScriptGlobalVariables, MAX_VAR_COUNT * sizeof(UINT64));
    if (!g_ScriptLocalVariables)
    {
        g_ScriptLocalVariables = (UINT64 *)malloc(MAX_VAR_COUNT * sizeof(UINT64));
        RtlZeroMemory(g_ScriptLocalVariables, MAX_VAR_COUNT * sizeof(UINT64));
    if (!g_ScriptTempVariables)
    {
        g_ScriptTempVariables = (UINT64 *)malloc(MAX_TEMP_COUNT * sizeof(UINT64));
        RtlZeroMemory(g_ScriptTempVariables, MAX_TEMP_COUNT * sizeof(UINT64));
    PSYMBOL_BUFFER CodeBuffer = ScriptEngineParse((char *)Expr.c_str());
    if (CodeBuffer->Message == NULL)
    {
        for (int i = 0; i < CodeBuffer->Pointer;)
        {
            ActionBuffer.Context                   = NULL;
            ActionBuffer.CurrentAction             = NULL;
            ActionBuffer.ImmediatelySendTheResults = FALSE;
            ActionBuffer.Tag                       = NULL;
            VariablesList.TempList            = g_ScriptTempVariables;
            VariablesList.GlobalVariablesList = g_ScriptGlobalVariables;
            VariablesList.LocalVariablesList  = g_ScriptLocalVariables;
            if (ScriptEngineExecute(GuestRegs,
                                    &ActionBuffer,
                                    &VariablesList,
                                    CodeBuffer,
                                    &i,
                                    &ErrorSymbol) == TRUE)
            {
                CHAR NameOfOperator[MAX_FUNCTION_NAME_LENGTH] = {0};
                ScriptEngineGetOperatorName(&ErrorSymbol, NameOfOperator);
                ShowMessages("invalid returning address for operator: %s",
                             NameOfOperator);
        ShowMessages("%s\n", CodeBuffer->Message);
    RemoveSymbolBuffer(CodeBuffer);
ScriptAutomaticStatementsTestWrapper(const string & Expr, UINT64 ExpectationValue, BOOLEAN ExceptError)
{
    g_CurrentExprEvalResult = 0;
    ScriptEngineWrapperTestParser(Expr);
    if (g_CurrentExprEvalResultHasError && ExceptError)
    {
        return TRUE;
    }
    else if (ExpectationValue == g_CurrentExprEvalResult)
    {
        return TRUE;
    }
    return FALSE;
}*/
return true
}

func (s *scriptEngineWrapper)AllocateStructForCasting()(ok bool){//col:140
/*AllocateStructForCasting(PALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING AllocationsForCastings)
{
VOID
ScriptEngineWrapperTestParser(const string & Expr)
{
    ALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING AllocationsForCastings = {0};
UINT64
ScriptEngineEvalUInt64StyleExpressionWrapper(const string & Expr, PBOOLEAN HasError)
{
    GUEST_REGS GuestRegs = {0};
    ScriptEngineEvalWrapper(&GuestRegs, Expr);
ScriptEngineWrapperGetHead(PVOID SymbolBuffer)
{
    return (UINT64)((PSYMBOL_BUFFER)SymbolBuffer)->Head;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetSize()(ok bool){//col:148
/*ScriptEngineWrapperGetSize(PVOID SymbolBuffer)
{
    UINT32 Size =
        (UINT32)((PSYMBOL_BUFFER)SymbolBuffer)->Pointer * sizeof(SYMBOL);
ScriptEngineWrapperGetPointer(PVOID SymbolBuffer)
{
    return (UINT32)((PSYMBOL_BUFFER)SymbolBuffer)->Pointer;
}*/
return true
}



