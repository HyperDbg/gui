package script-engine-wrapper
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\script-engine-wrapper\script-engine-wrapper.cpp.back

type ALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING struct{
* int8 //col:3
* int8 //col:4
* int8 //col:5
* int8 //col:6
* int8 //col:7
* int8 //col:8
}


type     typedef struct _UNICODE_STRING struct{
Length uint16 //col:12
MaximumLength uint16 //col:13
Buffer PWSTR //col:14
}


type     typedef struct _STUPID_STRUCT1 struct{
Flag32 uint32 //col:18
Flag64 uint64 //col:19
Context PVOID //col:20
StringValue PUNICODE_STRING //col:21
}


type     typedef struct _STUPID_STRUCT2 struct{
Sina32 uint32 //col:25
Sina64 uint64 //col:26
AghaaSina PVOID //col:27
UnicodeStr PUNICODE_STRING //col:28
StupidStruct1 PSTUPID_STRUCT1 //col:29
}


type     typedef struct _TEST_STRUCT struct{
Var1 uint64 //col:68
Var2 uint64 //col:69
Var3 uint64 //col:70
Var4 uint64 //col:71
}



type (
ScriptEngineWrapper interface{
ScriptEngineConvertNameToAddressWrapper()(ok bool)//col:4
ScriptEngineLoadFileSymbolWrapper()(ok bool)//col:8
ScriptEngineSetTextMessageCallbackWrapper()(ok bool)//col:12
ScriptEngineUnloadAllSymbolsWrapper()(ok bool)//col:16
ScriptEngineUnloadModuleSymbolWrapper()(ok bool)//col:20
ScriptEngineSearchSymbolForMaskWrapper()(ok bool)//col:24
ScriptEngineGetFieldOffsetWrapper()(ok bool)//col:28
ScriptEngineGetDataTypeSizeWrapper()(ok bool)//col:32
ScriptEngineCreateSymbolTableForDisassemblerWrapper()(ok bool)//col:36
ScriptEngineConvertFileToPdbPathWrapper()(ok bool)//col:40
ScriptEngineSymbolInitLoadWrapper()(ok bool)//col:48
ScriptEngineShowDataBasedOnSymbolTypesWrapper()(ok bool)//col:57
ScriptEngineSymbolAbortLoadingWrapper()(ok bool)//col:61
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper()(ok bool)//col:67
ScriptEngineParseWrapper()(ok bool)//col:85
PrintSymbolBufferWrapper()(ok bool)//col:89
ScriptEngineEvalWrapper()(ok bool)//col:146
ScriptAutomaticStatementsTestWrapper()(ok bool)//col:160
AllocateStructForCasting()(ok bool)//col:174
ScriptEngineWrapperGetHead()(ok bool)//col:178
ScriptEngineWrapperGetSize()(ok bool)//col:184
ScriptEngineWrapperGetPointer()(ok bool)//col:188
ScriptEngineWrapperRemoveSymbolBuffer()(ok bool)//col:192
}
scriptEngineWrapper struct{}
)

func NewScriptEngineWrapper()ScriptEngineWrapper{ return & scriptEngineWrapper{} }

func (s *scriptEngineWrapper)ScriptEngineConvertNameToAddressWrapper()(ok bool){//col:4
/*ScriptEngineConvertNameToAddressWrapper(const char * FunctionOrVariableName, PBOOLEAN WasFound)
{
    return ScriptEngineConvertNameToAddress(FunctionOrVariableName, WasFound);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineLoadFileSymbolWrapper()(ok bool){//col:8
/*ScriptEngineLoadFileSymbolWrapper(UINT64 BaseAddress, const char * PdbFileName)
{
    return ScriptEngineLoadFileSymbol(BaseAddress, PdbFileName);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSetTextMessageCallbackWrapper()(ok bool){//col:12
/*ScriptEngineSetTextMessageCallbackWrapper(PVOID Handler)
{
    return ScriptEngineSetTextMessageCallback(Handler);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineUnloadAllSymbolsWrapper()(ok bool){//col:16
/*ScriptEngineUnloadAllSymbolsWrapper()
{
    return ScriptEngineUnloadAllSymbols();
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineUnloadModuleSymbolWrapper()(ok bool){//col:20
/*ScriptEngineUnloadModuleSymbolWrapper(char * ModuleName)
{
    return ScriptEngineUnloadModuleSymbol(ModuleName);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSearchSymbolForMaskWrapper()(ok bool){//col:24
/*ScriptEngineSearchSymbolForMaskWrapper(const char * SearchMask)
{
    return ScriptEngineSearchSymbolForMask(SearchMask);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineGetFieldOffsetWrapper()(ok bool){//col:28
/*ScriptEngineGetFieldOffsetWrapper(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)
{
    return ScriptEngineGetFieldOffset(TypeName, FieldName, FieldOffset);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineGetDataTypeSizeWrapper()(ok bool){//col:32
/*ScriptEngineGetDataTypeSizeWrapper(CHAR * TypeName, UINT64 * TypeSize)
{
    return ScriptEngineGetDataTypeSize(TypeName, TypeSize);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineCreateSymbolTableForDisassemblerWrapper()(ok bool){//col:36
/*ScriptEngineCreateSymbolTableForDisassemblerWrapper(void * CallbackFunction)
{
    return ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineConvertFileToPdbPathWrapper()(ok bool){//col:40
/*ScriptEngineConvertFileToPdbPathWrapper(const char * LocalFilePath, char * ResultPath)
{
    return ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSymbolInitLoadWrapper()(ok bool){//col:48
/*ScriptEngineSymbolInitLoadWrapper(PMODULE_SYMBOL_DETAIL BufferToStoreDetails,
                                  UINT32                StoredLength,
                                  BOOLEAN               DownloadIfAvailable,
                                  const char *          SymbolPath,
                                  BOOLEAN               IsSilentLoad)
{
    return ScriptEngineSymbolInitLoad(BufferToStoreDetails, StoredLength, DownloadIfAvailable, SymbolPath, IsSilentLoad);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineShowDataBasedOnSymbolTypesWrapper()(ok bool){//col:57
/*ScriptEngineShowDataBasedOnSymbolTypesWrapper(
    const char * TypeName,
    UINT64       Address,
    BOOLEAN      IsStruct,
    PVOID        BufferAddress,
    const char * AdditionalParameters)
{
    return ScriptEngineShowDataBasedOnSymbolTypes(TypeName, Address, IsStruct, BufferAddress, AdditionalParameters);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSymbolAbortLoadingWrapper()(ok bool){//col:61
/*ScriptEngineSymbolAbortLoadingWrapper()
{
    return ScriptEngineSymbolAbortLoading();
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper()(ok bool){//col:67
/*ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(const char * LocalFilePath,
                                                            char *       PdbFilePath,
                                                            char *       GuidAndAgeDetails)
{
    return ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineParseWrapper()(ok bool){//col:85
/*ScriptEngineParseWrapper(char * Expr, BOOLEAN ShowErrorMessageIfAny)
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
        }
        ScriptEngineWrapperRemoveSymbolBuffer(SymbolBuffer);
        return NULL;
    }
}*/
return true
}

func (s *scriptEngineWrapper)PrintSymbolBufferWrapper()(ok bool){//col:89
/*PrintSymbolBufferWrapper(PVOID SymbolBuffer)
{
    PrintSymbolBuffer((PSYMBOL_BUFFER)SymbolBuffer);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineEvalWrapper()(ok bool){//col:146
/*ScriptEngineEvalWrapper(PGUEST_REGS GuestRegs,
                        string      Expr)
{
    SCRIPT_ENGINE_VARIABLES_LIST VariablesList = {0};
    if (!g_ScriptGlobalVariables)
    {
        g_ScriptGlobalVariables = (UINT64 *)malloc(MAX_VAR_COUNT * sizeof(UINT64));
        RtlZeroMemory(g_ScriptGlobalVariables, MAX_VAR_COUNT * sizeof(UINT64));
    }
    if (!g_ScriptLocalVariables)
    {
        g_ScriptLocalVariables = (UINT64 *)malloc(MAX_VAR_COUNT * sizeof(UINT64));
        RtlZeroMemory(g_ScriptLocalVariables, MAX_VAR_COUNT * sizeof(UINT64));
    }
    if (!g_ScriptTempVariables)
    {
        g_ScriptTempVariables = (UINT64 *)malloc(MAX_TEMP_COUNT * sizeof(UINT64));
        RtlZeroMemory(g_ScriptTempVariables, MAX_TEMP_COUNT * sizeof(UINT64));
    }
    PSYMBOL_BUFFER CodeBuffer = ScriptEngineParse((char *)Expr.c_str());
    ACTION_BUFFER ActionBuffer = {0};
    SYMBOL        ErrorSymbol  = {0};
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
                g_CurrentExprEvalResultHasError = TRUE;
                g_CurrentExprEvalResult         = NULL;
                break;
            }
        }
    }
    else
    {
        ShowMessages("%s\n", CodeBuffer->Message);
    }
    RemoveSymbolBuffer(CodeBuffer);
    return;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptAutomaticStatementsTestWrapper()(ok bool){//col:160
/*ScriptAutomaticStatementsTestWrapper(const string & Expr, UINT64 ExpectationValue, BOOLEAN ExceptError)
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

func (s *scriptEngineWrapper)AllocateStructForCasting()(ok bool){//col:174
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
    *HasError = g_CurrentExprEvalResultHasError;
    return g_CurrentExprEvalResult;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetHead()(ok bool){//col:178
/*ScriptEngineWrapperGetHead(PVOID SymbolBuffer)
{
    return (UINT64)((PSYMBOL_BUFFER)SymbolBuffer)->Head;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetSize()(ok bool){//col:184
/*ScriptEngineWrapperGetSize(PVOID SymbolBuffer)
{
    UINT32 Size =
        (UINT32)((PSYMBOL_BUFFER)SymbolBuffer)->Pointer * sizeof(SYMBOL);
    return Size;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetPointer()(ok bool){//col:188
/*ScriptEngineWrapperGetPointer(PVOID SymbolBuffer)
{
    return (UINT32)((PSYMBOL_BUFFER)SymbolBuffer)->Pointer;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperRemoveSymbolBuffer()(ok bool){//col:192
/*ScriptEngineWrapperRemoveSymbolBuffer(PVOID SymbolBuffer)
{
    RemoveSymbolBuffer((PSYMBOL_BUFFER)SymbolBuffer);
}*/
return true
}



