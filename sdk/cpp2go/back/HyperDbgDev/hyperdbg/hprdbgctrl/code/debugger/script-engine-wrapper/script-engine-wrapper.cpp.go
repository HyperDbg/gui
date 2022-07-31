package script-engine-wrapper
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\script-engine-wrapper\script-engine-wrapper.cpp.back

type (
ScriptEngineWrapper interface{
ScriptEngineConvertNameToAddressWrapper()(ok bool)//col:54
ScriptEngineLoadFileSymbolWrapper()(ok bool)//col:69
ScriptEngineSetTextMessageCallbackWrapper()(ok bool)//col:82
ScriptEngineUnloadAllSymbolsWrapper()(ok bool)//col:93
ScriptEngineUnloadModuleSymbolWrapper()(ok bool)//col:105
ScriptEngineSearchSymbolForMaskWrapper()(ok bool)//col:118
ScriptEngineGetFieldOffsetWrapper()(ok bool)//col:134
ScriptEngineGetDataTypeSizeWrapper()(ok bool)//col:148
ScriptEngineCreateSymbolTableForDisassemblerWrapper()(ok bool)//col:161
ScriptEngineConvertFileToPdbPathWrapper()(ok bool)//col:176
ScriptEngineSymbolInitLoadWrapper()(ok bool)//col:197
ScriptEngineShowDataBasedOnSymbolTypesWrapper()(ok bool)//col:219
ScriptEngineSymbolAbortLoadingWrapper()(ok bool)//col:231
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper()(ok bool)//col:249
ScriptEngineParseWrapper()(ok bool)//col:288
PrintSymbolBufferWrapper()(ok bool)//col:301
ScriptEngineEvalWrapper()(ok bool)//col:409
ScriptAutomaticStatementsTestWrapper()(ok bool)//col:446
AllocateStructForCasting()(ok bool)//col:534
ScriptEngineWrapperTestParser()(ok bool)//col:597
ScriptEngineEvalUInt64StyleExpressionWrapper()(ok bool)//col:622
ScriptEngineWrapperGetHead()(ok bool)//col:634
ScriptEngineWrapperGetSize()(ok bool)//col:648
ScriptEngineWrapperGetPointer()(ok bool)//col:660
ScriptEngineWrapperRemoveSymbolBuffer()(ok bool)//col:672
}
)

func NewScriptEngineWrapper() { return & scriptEngineWrapper{} }

func (s *scriptEngineWrapper)ScriptEngineConvertNameToAddressWrapper()(ok bool){//col:54
/*ScriptEngineConvertNameToAddressWrapper(const char * FunctionOrVariableName, PBOOLEAN WasFound)
{
    return ScriptEngineConvertNameToAddress(FunctionOrVariableName, WasFound);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineLoadFileSymbolWrapper()(ok bool){//col:69
/*ScriptEngineLoadFileSymbolWrapper(UINT64 BaseAddress, const char * PdbFileName)
{
    return ScriptEngineLoadFileSymbol(BaseAddress, PdbFileName);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSetTextMessageCallbackWrapper()(ok bool){//col:82
/*ScriptEngineSetTextMessageCallbackWrapper(PVOID Handler)
{
    return ScriptEngineSetTextMessageCallback(Handler);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineUnloadAllSymbolsWrapper()(ok bool){//col:93
/*ScriptEngineUnloadAllSymbolsWrapper()
{
    return ScriptEngineUnloadAllSymbols();
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineUnloadModuleSymbolWrapper()(ok bool){//col:105
/*ScriptEngineUnloadModuleSymbolWrapper(char * ModuleName)
{
    return ScriptEngineUnloadModuleSymbol(ModuleName);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSearchSymbolForMaskWrapper()(ok bool){//col:118
/*ScriptEngineSearchSymbolForMaskWrapper(const char * SearchMask)
{
    return ScriptEngineSearchSymbolForMask(SearchMask);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineGetFieldOffsetWrapper()(ok bool){//col:134
/*ScriptEngineGetFieldOffsetWrapper(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)
{
    return ScriptEngineGetFieldOffset(TypeName, FieldName, FieldOffset);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineGetDataTypeSizeWrapper()(ok bool){//col:148
/*ScriptEngineGetDataTypeSizeWrapper(CHAR * TypeName, UINT64 * TypeSize)
{
    return ScriptEngineGetDataTypeSize(TypeName, TypeSize);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineCreateSymbolTableForDisassemblerWrapper()(ok bool){//col:161
/*ScriptEngineCreateSymbolTableForDisassemblerWrapper(void * CallbackFunction)
{
    return ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineConvertFileToPdbPathWrapper()(ok bool){//col:176
/*ScriptEngineConvertFileToPdbPathWrapper(const char * LocalFilePath, char * ResultPath)
{
    return ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineSymbolInitLoadWrapper()(ok bool){//col:197
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

func (s *scriptEngineWrapper)ScriptEngineShowDataBasedOnSymbolTypesWrapper()(ok bool){//col:219
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

func (s *scriptEngineWrapper)ScriptEngineSymbolAbortLoadingWrapper()(ok bool){//col:231
/*ScriptEngineSymbolAbortLoadingWrapper()
{
    return ScriptEngineSymbolAbortLoading();
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper()(ok bool){//col:249
/*ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(const char * LocalFilePath,
                                                            char *       PdbFilePath,
                                                            char *       GuidAndAgeDetails)
{
    return ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineParseWrapper()(ok bool){//col:288
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

func (s *scriptEngineWrapper)PrintSymbolBufferWrapper()(ok bool){//col:301
/*PrintSymbolBufferWrapper(PVOID SymbolBuffer)
{
    PrintSymbolBuffer((PSYMBOL_BUFFER)SymbolBuffer);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineEvalWrapper()(ok bool){//col:409
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

func (s *scriptEngineWrapper)ScriptAutomaticStatementsTestWrapper()(ok bool){//col:446
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

func (s *scriptEngineWrapper)AllocateStructForCasting()(ok bool){//col:534
/*AllocateStructForCasting(PALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING AllocationsForCastings)
{
    typedef struct _UNICODE_STRING
    {
    } UNICODE_STRING, *PUNICODE_STRING;
    typedef struct _STUPID_STRUCT1
    {
    } STUPID_STRUCT1, *PSTUPID_STRUCT1;
    typedef struct _STUPID_STRUCT2
    {
    } STUPID_STRUCT2, *PSTUPID_STRUCT2;
    WCHAR           MyString1[40]   = L"Hi come from stupid struct 1!";
    UINT32          SizeOfMyString1 = wcslen(MyString1) * sizeof(WCHAR) + 2;
    PUNICODE_STRING UnicodeStr1     = (PUNICODE_STRING)malloc(sizeof(UNICODE_STRING));
    AllocationsForCastings->Buff1   = (CHAR *)UnicodeStr1;
    WCHAR * Buff1                   = (WCHAR *)malloc(SizeOfMyString1);
    AllocationsForCastings->Buff2   = (CHAR *)Buff1;
    RtlZeroMemory(Buff1, SizeOfMyString1);
    UnicodeStr1->Buffer = Buff1;
    UnicodeStr1->Length = UnicodeStr1->MaximumLength = SizeOfMyString1;
    memcpy(UnicodeStr1->Buffer, MyString1, SizeOfMyString1);
    WCHAR           MyString2[40]   = L"Goodbye I'm at stupid struct 2!";
    UINT32          SizeOfMyString2 = wcslen(MyString2) * sizeof(WCHAR) + 2;
    PUNICODE_STRING UnicodeStr2     = (PUNICODE_STRING)malloc(sizeof(UNICODE_STRING));
    AllocationsForCastings->Buff3   = (CHAR *)UnicodeStr2;
    WCHAR * Buff2                   = (WCHAR *)malloc(SizeOfMyString2);
    AllocationsForCastings->Buff4   = (CHAR *)Buff2;
    RtlZeroMemory(Buff2, SizeOfMyString2);
    UnicodeStr2->Buffer = Buff2;
    UnicodeStr2->Length = UnicodeStr2->MaximumLength = SizeOfMyString2;
    memcpy(UnicodeStr2->Buffer, MyString2, SizeOfMyString2);
    PSTUPID_STRUCT1 StupidStruct1 = (PSTUPID_STRUCT1)malloc(sizeof(STUPID_STRUCT1));
    AllocationsForCastings->Buff5 = (CHAR *)StupidStruct1;
    StupidStruct1->Flag32         = 0x3232;
    StupidStruct1->Flag64         = 0x6464;
    StupidStruct1->Context        = (PVOID)0x85;
    StupidStruct1->StringValue    = UnicodeStr1;
    PSTUPID_STRUCT2 StupidStruct2 = (PSTUPID_STRUCT2)malloc(sizeof(STUPID_STRUCT2));
    AllocationsForCastings->Buff6 = (CHAR *)StupidStruct2;
    StupidStruct2->Sina32        = 0x32;
    StupidStruct2->Sina64        = 0x64;
    StupidStruct2->AghaaSina     = (PVOID)0x55;
    StupidStruct2->UnicodeStr    = UnicodeStr2;
    StupidStruct2->StupidStruct1 = StupidStruct1;
    return StupidStruct2;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperTestParser()(ok bool){//col:597
/*ScriptEngineWrapperTestParser(const string & Expr)
{
    ALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING AllocationsForCastings = {0};
    typedef struct _TEST_STRUCT
    {
        UINT64 Var1;
        UINT64 Var2;
        UINT64 Var3;
        UINT64 Var4;
    } TEST_STRUCT, *PTEST_STRUCT;
    PTEST_STRUCT TestStruct = (PTEST_STRUCT)malloc(sizeof(TEST_STRUCT));
    RtlZeroMemory(TestStruct, sizeof(TEST_STRUCT));
    TestStruct->Var1 = 0x41414141;
    TestStruct->Var3 = 0x4242424242424242;
    GUEST_REGS GuestRegs = {0};
    char    test[] = "Hello world !";
    wchar_t testw[] =
        L"A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 0 1 2 3 4 5 6 7 8 "
        L"9 a b c d e f g h i j k l m n o p q r s t u v w x y z";
    char * RspReg = (char *)malloc(0x100);
    memcpy(RspReg, testw, 0x100);
    GuestRegs.rax = 0x1;
    GuestRegs.rdx = 0x3;
    GuestRegs.rbx = 0x4;
    GuestRegs.rsp = (UINT64)RspReg + 0x50;
    GuestRegs.rbp = 0x6;
    GuestRegs.rsi = 0x7;
    GuestRegs.rdi = 0x8;
    GuestRegs.r8  = 0x9;
    GuestRegs.r9  = 0xa;
    GuestRegs.r10 = 0xb;
    GuestRegs.r11 = 0xc;
    GuestRegs.r12 = 0xd;
    GuestRegs.r13 = 0xe;
    GuestRegs.r14 = (UINT64)testw;
    GuestRegs.r15 = (UINT64)test;
    ScriptEngineEvalWrapper(&GuestRegs, Expr);
    free(RspReg);
    free(TestStruct);
    free(AllocationsForCastings.Buff1);
    free(AllocationsForCastings.Buff2);
    free(AllocationsForCastings.Buff3);
    free(AllocationsForCastings.Buff4);
    free(AllocationsForCastings.Buff5);
    free(AllocationsForCastings.Buff6);
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineEvalUInt64StyleExpressionWrapper()(ok bool){//col:622
/*ScriptEngineEvalUInt64StyleExpressionWrapper(const string & Expr, PBOOLEAN HasError)
{
    GUEST_REGS GuestRegs = {0};
    ScriptEngineEvalWrapper(&GuestRegs, Expr);
    *HasError = g_CurrentExprEvalResultHasError;
    return g_CurrentExprEvalResult;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetHead()(ok bool){//col:634
/*ScriptEngineWrapperGetHead(PVOID SymbolBuffer)
{
    return (UINT64)((PSYMBOL_BUFFER)SymbolBuffer)->Head;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetSize()(ok bool){//col:648
/*ScriptEngineWrapperGetSize(PVOID SymbolBuffer)
{
    UINT32 Size =
        (UINT32)((PSYMBOL_BUFFER)SymbolBuffer)->Pointer * sizeof(SYMBOL);
    return Size;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperGetPointer()(ok bool){//col:660
/*ScriptEngineWrapperGetPointer(PVOID SymbolBuffer)
{
    return (UINT32)((PSYMBOL_BUFFER)SymbolBuffer)->Pointer;
}*/
return true
}

func (s *scriptEngineWrapper)ScriptEngineWrapperRemoveSymbolBuffer()(ok bool){//col:672
/*ScriptEngineWrapperRemoveSymbolBuffer(PVOID SymbolBuffer)
{
    RemoveSymbolBuffer((PSYMBOL_BUFFER)SymbolBuffer);
}*/
return true
}



