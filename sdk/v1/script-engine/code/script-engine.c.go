package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-engine\code\script-engine.c.back

type (
	ScriptEngine interface {
		ScriptEngineConvertNameToAddress() (ok bool)    //col:813
		ScriptEngineBooleanExpresssionParse() (ok bool) //col:937
		RemoveSymbol() (ok bool)                        //col:1031
		PseudoRegToInt() (ok bool)                      //col:1042
		SemanticRuleToInt() (ok bool)                   //col:1053
		HandleError() (ok bool)                         //col:1103
		GetLocalIdentifierVal() (ok bool)               //col:1116
		NewGlobalIdentifier() (ok bool)                 //col:1137
		LalrIsOperandType() (ok bool)                   //col:1185
	}
	scriptEngine struct{}
)

func NewScriptEngine() ScriptEngine { return &scriptEngine{} }

func (s *scriptEngine) ScriptEngineConvertNameToAddress() (ok bool) { //col:813
	/*
	   ScriptEngineConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound)

	   	{
	   	    return SymConvertNameToAddress(FunctionOrVariableName, WasFound);

	   ScriptEngineLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName)

	   	{
	   	    return SymLoadFileSymbol(BaseAddress, PdbFileName);

	   ScriptEngineSetTextMessageCallback(PVOID Handler)

	   	{
	   	    SymSetTextMessageCallback(Handler);

	   ScriptEngineUnloadAllSymbols()

	   	{
	   	    return SymUnloadAllSymbols();

	   ScriptEngineUnloadModuleSymbol(char * ModuleName)

	   	{
	   	    return SymUnloadModuleSymbol(ModuleName);

	   ScriptEngineSearchSymbolForMask(const char * SearchMask)

	   	{
	   	    return SymSearchSymbolForMask(SearchMask);

	   ScriptEngineGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)

	   	{
	   	    return SymGetFieldOffset(TypeName, FieldName, FieldOffset);

	   ScriptEngineGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize)

	   	{
	   	    return SymGetDataTypeSize(TypeName, TypeSize);

	   ScriptEngineCreateSymbolTableForDisassembler(void * CallbackFunction)

	   	{
	   	    return SymCreateSymbolTableForDisassembler(CallbackFunction);

	   ScriptEngineConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath)

	   	{
	   	    return SymConvertFileToPdbPath(LocalFilePath, ResultPath);

	   ScriptEngineSymbolInitLoad(PVOID        BufferToStoreDetails,

	   	UINT32       StoredLength,
	   	BOOLEAN      DownloadIfAvailable,
	   	const char * SymbolPath,
	   	BOOLEAN      IsSilentLoad)

	   	{
	   	    return SymbolInitLoad(BufferToStoreDetails, StoredLength, DownloadIfAvailable, SymbolPath, IsSilentLoad);

	   ScriptEngineShowDataBasedOnSymbolTypes(const char * TypeName,

	   	UINT64       Address,
	   	BOOLEAN      IsStruct,
	   	PVOID        BufferAddress,
	   	const char * AdditionalParameters)

	   	{
	   	    return SymShowDataBasedOnSymbolTypes(TypeName, Address, IsStruct, BufferAddress, AdditionalParameters);

	   ScriptEngineSymbolAbortLoading()

	   	{
	   	    return SymbolAbortLoading();

	   ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath, char * PdbFilePath, char * GuidAndAgeDetails)

	   	{
	   	    return SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails);

	   ScriptEngineParse(char * str)

	   	{
	   	    PTOKEN_LIST Stack = NewTokenList();
	   	    PTOKEN_LIST    MatchedStack = NewTokenList();
	   	    PSYMBOL_BUFFER CodeBuffer   = NewSymbolBuffer();
	   	    if (FirstCall)
	   	    {
	   	        IdTable   = NewTokenList();
	   	    PTOKEN TopToken = NewUnknownToken();
	   	    PTOKEN EndToken = NewToken(END_OF_STACK, "$");
	   	    PTOKEN StartToken = NewToken(NON_TERMINAL, START_VARIABLE);
	   	    Push(Stack, EndToken);
	   	    Push(Stack, StartToken);
	   	    c = sgetc(str);
	   	    PTOKEN CurrentIn = Scan(str, &c);
	   	    if (CurrentIn->Type == UNKNOWN)
	   	    {
	   	        Error               = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	        ErrorMessage        = HandleError(&Error, str);
	   	        RemoveTokenList(Stack);
	   	        RemoveTokenList(MatchedStack);
	   	        RemoveToken(&CurrentIn);
	   	        RemoveToken(&TopToken);
	   	        TopToken = Pop(Stack);
	   	        printf("\nTop Token :\n");
	   	        PrintToken(TopToken);
	   	        printf("\nCurrent Input :\n");
	   	        PrintToken(CurrentIn);
	   	        printf("\n");
	   	        if (TopToken->Type == NON_TERMINAL)
	   	        {
	   	            if (!strcmp(TopToken->Value, "BOOLEAN_EXPRESSION"))
	   	            {
	   	                UINT64 BooleanExpressionSize = BooleanExpressionExtractEnd(str, &WaitForWaitStatementBooleanExpression, CurrentIn);
	   	                ScriptEngineBooleanExpresssionParse(BooleanExpressionSize, CurrentIn, MatchedStack, CodeBuffer, str, &c, &Error);
	   	                if (Error != SCRIPT_ENGINE_ERROR_FREE)
	   	                {
	   	                    break;
	   	                }
	   	                RemoveToken(&CurrentIn);
	   	                CurrentIn = Scan(str, &c);
	   	                if (CurrentIn->Type == UNKNOWN)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN;
	   	                    break;
	   	                }
	   	                RemoveToken(&CurrentIn);
	   	                CurrentIn = Scan(str, &c);
	   	                if (CurrentIn->Type == UNKNOWN)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN;
	   	                    break;
	   	                }
	   	                RemoveToken(&TopToken);
	   	                TopToken = Pop(Stack);
	   	                NonTerminalId = GetNonTerminalId(TopToken);
	   	                if (NonTerminalId == INVALID)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    break;
	   	                }
	   	                TerminalId = GetTerminalId(CurrentIn);
	   	                if (TerminalId == INVALID)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    break;
	   	                }
	   	                RuleId = ParseTable[NonTerminalId][TerminalId];
	   	                if (RuleId == INVALID)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    break;
	   	                }
	   	                for (int i = RhsSize[RuleId] - 1; i >= 0; i--)
	   	                {
	   	                    PTOKEN Token = &Rhs[RuleId][i];
	   	                    if (Token->Type == EPSILON)
	   	                        break;
	   	                    PTOKEN DuplicatedToken = CopyToken(Token);
	   	                    Push(Stack, DuplicatedToken);
	   	        else if (TopToken->Type == SEMANTIC_RULE)
	   	        {
	   	            {
	   	                RemoveToken(&TopToken);
	   	                TopToken = Pop(Stack);
	   	                Push(MatchedStack, CurrentIn);
	   	                CurrentIn = Scan(str, &c);
	   	                if (CurrentIn->Type == UNKNOWN)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    break;
	   	                }
	   	            }
	   	            else
	   	            {
	   	                {
	   	                    WaitForWaitStatementBooleanExpression = TRUE;
	   	                }
	   	                CodeGen(MatchedStack, CodeBuffer, TopToken, &Error);
	   	                if (Error != SCRIPT_ENGINE_ERROR_FREE)
	   	                {
	   	                    break;
	   	                }
	   	            }
	   	        }
	   	        else
	   	        {
	   	            if (!IsEqual(TopToken, CurrentIn))
	   	            {
	   	                Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                break;
	   	            }
	   	            else
	   	            {
	   	                RemoveToken(&CurrentIn);
	   	                CurrentIn = Scan(str, &c);
	   	                if (CurrentIn->Type == UNKNOWN)
	   	                {
	   	                    Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    break;
	   	                }
	   	            }
	   	        }

	   #ifdef _SCRIPT_ENGINE_LL1_DBG_EN

	   	    printf("Stack: \n");
	   	    PrintTokenList(Stack);
	   	    printf("\n");
	   	} while (TopToken->Type != END_OF_STACK);
	   	if (Error != SCRIPT_ENGINE_ERROR_FREE)
	   	{
	   	    ErrorMessage = HandleError(&Error, str);
	   	    CleanTempList();
	   	if (Stack)
	   	    RemoveTokenList(Stack);
	   	if (MatchedStack)
	   	    RemoveTokenList(MatchedStack);
	   	if (CurrentIn)
	   	    RemoveToken(&CurrentIn);
	   	if (TopToken)
	   	    RemoveToken(&TopToken);

	   CodeGen(PTOKEN_LIST MatchedStack, PSYMBOL_BUFFER CodeBuffer, PTOKEN Operator, PSCRIPT_ENGINE_ERROR_TYPE Error)

	   	{
	   	    PTOKEN Op0  = NULL;
	   	    PTOKEN Op1  = NULL;
	   	    PTOKEN Op2  = NULL;
	   	    PTOKEN Temp = NULL;
	   	    PSYMBOL OperatorSymbol = NULL;
	   	    PSYMBOL Op0Symbol      = NULL;
	   	    PSYMBOL Op1Symbol      = NULL;
	   	    PSYMBOL Op2Symbol      = NULL;
	   	    PSYMBOL TempSymbol     = NULL;
	   	    OperatorSymbol = ToSymbol(Operator, Error);
	   	    printf("Operator :\n");
	   	    PrintToken(Operator);
	   	    printf("\n");
	   	    printf("Semantic Stack:\n");
	   	    PrintTokenList(MatchedStack);
	   	    printf("\n");
	   	    printf("Code Buffer:\n");
	   	    PrintSymbolBuffer(CodeBuffer);
	   	    printf(".\n.\n.\n\n");
	   	    while (TRUE)
	   	    {
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            Op1 = Pop(MatchedStack);
	   	            if (Op1->Type == GLOBAL_UNRESOLVED_ID)
	   	            {
	   	                Op1Symbol = NewSymbol();
	   	                free(Op1Symbol->Value);
	   	                Op1Symbol->Value = NewGlobalIdentifier(Op1);
	   	                SetType(&Op1Symbol->Type, SYMBOL_GLOBAL_ID_TYPE);
	   	            else if (Op1->Type == LOCAL_UNRESOLVED_ID)
	   	            {
	   	                Op1Symbol = NewSymbol();
	   	                free(Op1Symbol->Value);
	   	                Op1Symbol->Value = NewLocalIdentifier(Op1);
	   	                SetType(&Op1Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
	   	                Op1Symbol = ToSymbol(Op1, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, Op1Symbol);
	   	            FreeTemp(Op0);
	   	            FreeTemp(Op1);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        else if (IsType2Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        else if (IsType1Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            Temp = NewTemp(Error);
	   	            Push(MatchedStack, Temp);
	   	            TempSymbol = ToSymbol(Temp, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, TempSymbol);
	   	            FreeTemp(Op0);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        else if (IsType4Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            PSYMBOL_BUFFER TempStack    = NewSymbolBuffer();
	   	                if (Op1)
	   	                {
	   	                    RemoveToken(&Op1);
	   	                Op1 = Pop(MatchedStack);
	   	                if (Op1->Type != SEMANTIC_RULE)
	   	                {
	   	                    Op1Symbol = ToSymbol(Op1, Error);
	   	                    FreeTemp(Op1);
	   	                    PushSymbol(TempStack, Op1Symbol);
	   	                    RemoveSymbol(&Op1Symbol);
	   	                    if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	                    {
	   	                        RemoveSymbolBuffer(TempStack);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            FreeTemp(Op0);
	   	            PSYMBOL OperandCountSymbol = NewSymbol();
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, OperandCountSymbol);
	   	            RemoveSymbol(&OperandCountSymbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                RemoveSymbolBuffer(TempStack);
	   	            PSYMBOL FirstArg = (PSYMBOL)((unsigned long long)CodeBuffer->Head +
	   	                                         (unsigned long long)(CodeBuffer->Pointer * sizeof(SYMBOL)));
	   	            for (int i = TempStack->Pointer - 1; i >= 0; i--)
	   	            {
	   	                Symbol = TempStack->Head + i;
	   	                PushSymbol(CodeBuffer, Symbol);
	   	            RemoveSymbolBuffer(TempStack);
	   	                if (*Str == '%')
	   	                {
	   	                    CHAR Temp = *(Str + 1);
	   	                    if (Temp == 'd' || Temp == 'i' || Temp == 'u' || Temp == 'o' ||
	   	                        Temp == 'x' || Temp == 'c' || Temp == 'p' || Temp == 's' ||
	   	                        !strncmp(Str, "%ws", 3) || !strncmp(Str, "%ls", 3) ||
	   	                        !strncmp(Str, "%ld", 3) || !strncmp(Str, "%li", 3) ||
	   	                        !strncmp(Str, "%lu", 3) || !strncmp(Str, "%lo", 3) ||
	   	                        !strncmp(Str, "%lx", 3) ||
	   	                        !strncmp(Str, "%hd", 3) || !strncmp(Str, "%hi", 3) ||
	   	                        !strncmp(Str, "%hu", 3) || !strncmp(Str, "%ho", 3) ||
	   	                        !strncmp(Str, "%hx", 3) ||
	   	                        !strncmp(Str, "%lld", 4) || !strncmp(Str, "%lli", 4) ||
	   	                        !strncmp(Str, "%llu", 4) || !strncmp(Str, "%llo", 4) ||
	   	                        !strncmp(Str, "%llx", 4)
	   	                    )
	   	                    {
	   	                        if (i < ArgCount)
	   	                        {
	   	                            Symbol = FirstArg + i;
	   	                        }
	   	                        else
	   	                        {
	   	                            *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                            break;
	   	                        }
	   	                        Symbol->Type &= 0xffffffff;
	   	                        Symbol->Type |= (UINT64)(Str - Format - 1) << 32;
	   	                        i++;
	   	                    }
	   	                }
	   	                Str++;
	   	            } while (*Str);
	   	            if (i != ArgCount)
	   	            {
	   	                *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	            }
	   	            if (*Error == SCRIPT_ENGINE_ERROR_SYNTAX)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        else if (IsType5Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	        {
	   	            Op0 = Pop(MatchedStack);
	   	        else if (IsType6Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            Op1       = Pop(MatchedStack);
	   	            Op1Symbol = ToSymbol(Op1, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, Op1Symbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            Temp = NewTemp(Error);
	   	            Push(MatchedStack, Temp);
	   	            TempSymbol = ToSymbol(Temp, Error);
	   	            PushSymbol(CodeBuffer, TempSymbol);
	   	            FreeTemp(Op0);
	   	            FreeTemp(Op1);
	   	        else if (IsType7Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            Op1       = Pop(MatchedStack);
	   	            Op1Symbol = ToSymbol(Op1, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, Op1Symbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            FreeTemp(Op0);
	   	            FreeTemp(Op1);
	   	        else if (IsType8Func(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            Op1       = Pop(MatchedStack);
	   	            Op1Symbol = ToSymbol(Op1, Error);
	   	            PTOKEN  Op2       = Pop(MatchedStack);
	   	            PSYMBOL Op2Symbol = ToSymbol(Op2, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, Op1Symbol);
	   	            PushSymbol(CodeBuffer, Op2Symbol);
	   	            Temp = NewTemp(Error);
	   	            Push(MatchedStack, Temp);
	   	            TempSymbol = ToSymbol(Temp, Error);
	   	            PushSymbol(CodeBuffer, TempSymbol);
	   	            FreeTemp(Op2);
	   	            FreeTemp(Op0);
	   	            FreeTemp(Op1);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        else if (IsTwoOperandOperator(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            Op1       = Pop(MatchedStack);
	   	            Op1Symbol = ToSymbol(Op1, Error);
	   	            Temp = NewTemp(Error);
	   	            Push(MatchedStack, Temp);
	   	            TempSymbol = ToSymbol(Temp, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            PushSymbol(CodeBuffer, Op1Symbol);
	   	            PushSymbol(CodeBuffer, TempSymbol);
	   	            FreeTemp(Op0);
	   	            FreeTemp(Op1);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        else if (IsOneOperandOperator(Operator))
	   	        {
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            FreeTemp(Op0);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        {
	   	            PTOKEN OperatorCopy = CopyToken(Operator);
	   	            Push(MatchedStack, OperatorCopy);
	   	        {
	   	            PTOKEN OperatorCopy = CopyToken(Operator);
	   	            Push(MatchedStack, OperatorCopy);
	   	        {
	   	            UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            PSYMBOL JumpAddressSymbol = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            sprintf(str, "%llu", CodeBuffer->Pointer);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	            FreeTemp(Op0);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        {
	   	            UINT64  CurrentPointer           = CodeBuffer->Pointer;
	   	            PTOKEN  JumpSemanticAddressToken = Pop(MatchedStack);
	   	            UINT64  JumpSemanticAddress      = DecimalToInt(JumpSemanticAddressToken->Value);
	   	            PSYMBOL JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpSemanticAddress - 2);
	   	            RemoveToken(&JumpSemanticAddressToken);
	   	            PSYMBOL JumpInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpInstruction);
	   	            RemoveSymbol(&JumpInstruction);
	   	            JumpAddressSymbol        = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            sprintf(str, "%llu", CurrentPointer);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	        {
	   	            UINT64  CurrentPointer           = CodeBuffer->Pointer;
	   	            PTOKEN  JumpSemanticAddressToken = Pop(MatchedStack);
	   	            {
	   	                UINT64 JumpSemanticAddress = DecimalToInt(JumpSemanticAddressToken->Value);
	   	                JumpAddressSymbol          = (PSYMBOL)(CodeBuffer->Head + JumpSemanticAddress + 1);
	   	                RemoveToken(&JumpSemanticAddressToken);
	   	                JumpSemanticAddressToken = Pop(MatchedStack);
	   	            RemoveToken(&JumpSemanticAddressToken);
	   	        {
	   	            PTOKEN OperatorCopy = CopyToken(Operator);
	   	            Push(MatchedStack, OperatorCopy);
	   	            sprintf(str, "%llu", CodeBuffer->Pointer);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	        {
	   	            UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	            RemoveSymbol(&OperatorSymbol);
	   	            OperatorSymbol = ToSymbol(JzToken, Error);
	   	            RemoveToken(&JzToken);
	   	            PSYMBOL JumpAddressSymbol = NewSymbol();
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            PTOKEN StartOfWhileToken = Pop(MatchedStack);
	   	            sprintf(str, "%llu", CurrentPointer + 1);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	            Push(MatchedStack, StartOfWhileToken);
	   	            PushSymbol(CodeBuffer, OperatorSymbol);
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            FreeTemp(Op0);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	        }
	   	        {
	   	            PSYMBOL JumpInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpInstruction);
	   	            RemoveSymbol(&JumpInstruction);
	   	            PTOKEN  JumpAddressToken  = Pop(MatchedStack);
	   	            UINT64  JumpAddress       = DecimalToInt(JumpAddressToken->Value);
	   	            PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	            do
	   	            {
	   	                RemoveToken(&JumpAddressToken);
	   	                JumpAddressToken = Pop(MatchedStack);
	   	                {
	   	                    break;
	   	                }
	   	                JumpAddress              = DecimalToInt(JumpAddressToken->Value);
	   	                JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
	   	            } while (TRUE);
	   	            RemoveToken(&JumpAddressToken);
	   	        {
	   	            PTOKEN OperatorCopy = CopyToken(Operator);
	   	            Push(MatchedStack, OperatorCopy);
	   	            sprintf(str, "%llu", CodeBuffer->Pointer);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	        {
	   	            PSYMBOL JumpInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpInstruction);
	   	            RemoveSymbol(&JumpInstruction);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            PTOKEN JumpAddressToken = Pop(MatchedStack);
	   	            UINT64 JumpAddress      = DecimalToInt(JumpAddressToken->Value);
	   	            PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            FreeTemp(Op0);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	            do
	   	            {
	   	                RemoveToken(&JumpAddressToken);
	   	                JumpAddressToken = Pop(MatchedStack);
	   	                {
	   	                    break;
	   	                }
	   	                JumpAddress = DecimalToInt(JumpAddressToken->Value);
	   	                printf("Jz Jump Address = %d\n", JumpAddress);
	   	                JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
	   	            } while (TRUE);
	   	            RemoveToken(&JumpAddressToken);
	   	        {
	   	            PTOKEN OperatorCopy = CopyToken(Operator);
	   	            Push(MatchedStack, OperatorCopy);
	   	            sprintf(str, "%llu", CodeBuffer->Pointer);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	        {
	   	            PSYMBOL JnzInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JnzInstruction);
	   	            RemoveSymbol(&JnzInstruction);
	   	            PSYMBOL JnzAddressSymbol = NewSymbol();
	   	            PushSymbol(CodeBuffer, JnzAddressSymbol);
	   	            RemoveSymbol(&JnzAddressSymbol);
	   	            Op0       = Pop(MatchedStack);
	   	            Op0Symbol = ToSymbol(Op0, Error);
	   	            PushSymbol(CodeBuffer, Op0Symbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            PSYMBOL JumpInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpInstruction);
	   	            RemoveSymbol(&JumpInstruction);
	   	            PSYMBOL JumpAddressSymbol = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            PTOKEN StartOfForAddressToken = Pop(MatchedStack);
	   	            sprintf(str, "%llu", CodeBuffer->Pointer);
	   	            PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, CurrentAddressToken);
	   	            Push(MatchedStack, StartOfForAddressToken);
	   	        {
	   	            PSYMBOL JumpInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpInstruction);
	   	            RemoveSymbol(&JumpInstruction);
	   	            PTOKEN JumpAddressToken = Pop(MatchedStack);
	   	            UINT64 JumpAddress      = DecimalToInt(JumpAddressToken->Value);
	   	            PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            RemoveToken(&JumpAddressToken);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	            JumpAddressToken      = Pop(MatchedStack);
	   	            JumpAddress           = DecimalToInt(JumpAddressToken->Value);
	   	            JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpAddress - 1);
	   	            sprintf(str, "%llu", JumpAddress - 4);
	   	            PTOKEN JzAddressToken = NewToken(DECIMAL, str);
	   	            Push(MatchedStack, JzAddressToken);
	   	            Push(MatchedStack, IncDecToken);
	   	            Push(MatchedStack, JumpAddressToken);
	   	        {
	   	            PSYMBOL JumpInstruction = NewSymbol();
	   	            PushSymbol(CodeBuffer, JumpInstruction);
	   	            RemoveSymbol(&JumpInstruction);
	   	            PTOKEN JumpAddressToken = Pop(MatchedStack);
	   	            UINT64 JumpAddress      = DecimalToInt(JumpAddressToken->Value);
	   	            PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
	   	            PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	            RemoveSymbol(&JumpAddressSymbol);
	   	            RemoveToken(&JumpAddressToken);
	   	            if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	            {
	   	                break;
	   	            }
	   	            JumpAddressToken = Pop(MatchedStack);
	   	                RemoveToken(&JumpAddressToken);
	   	                JumpAddressToken = Pop(MatchedStack);
	   	                {
	   	                    break;
	   	                }
	   	                else
	   	                {
	   	                    JumpAddress = DecimalToInt(JumpAddressToken->Value);
	   	                    JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
	   	            } while (TRUE);
	   	            RemoveToken(&JumpAddressToken);
	   	        {
	   	            PTOKEN_LIST TempStack = NewTokenList();
	   	                TempToken = Pop(MatchedStack);
	   	                {
	   	                    Push(MatchedStack, TempToken);
	   	                    sprintf(str, "%llu", CurrentPointer);
	   	                    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	                    Push(MatchedStack, CurrentAddressToken);
	   	                    PSYMBOL JumpInstruction = NewSymbol();
	   	                    PushSymbol(CodeBuffer, JumpInstruction);
	   	                    RemoveSymbol(&JumpInstruction);
	   	                    PSYMBOL JumpAddressSymbol = NewSymbol();
	   	                    PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	                    RemoveSymbol(&JumpAddressSymbol);
	   	                        TempToken = Pop(TempStack);
	   	                        Push(MatchedStack, TempToken);
	   	                    } while (TempStack->Pointer != 0);
	   	                else if (MatchedStack->Pointer == 0)
	   	                {
	   	                    *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    RemoveToken(&TempToken);
	   	                    Push(TempStack, TempToken);
	   	            } while (TRUE);
	   	            RemoveTokenList(TempStack);
	   	        {
	   	            PTOKEN_LIST TempStack = NewTokenList();
	   	                TempToken = Pop(MatchedStack);
	   	                {
	   	                    Push(MatchedStack, TempToken);
	   	                    PSYMBOL JumpInstruction = NewSymbol();
	   	                    PushSymbol(CodeBuffer, JumpInstruction);
	   	                    RemoveSymbol(&JumpInstruction);
	   	                    TempToken = Pop(TempStack);
	   	                    Push(MatchedStack, TempToken);
	   	                    PSYMBOL JumpAddressSymbol = NewSymbol();
	   	                    JumpAddressSymbol->Value  = DecimalToInt(TempToken->Value);
	   	                    PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	                    RemoveSymbol(&JumpAddressSymbol);
	   	                        TempToken = Pop(TempStack);
	   	                        Push(MatchedStack, TempToken);
	   	                    } while (TempStack->Pointer != 0);
	   	                else if (MatchedStack->Pointer == 0)
	   	                {
	   	                    *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                    break;
	   	                }
	   	                else
	   	                {
	   	                    Push(TempStack, TempToken);
	   	            } while (TRUE);
	   	            RemoveTokenList(TempStack);
	   	    printf("Semantic Stack:\n");
	   	    PrintTokenList(MatchedStack);
	   	    printf("\n");
	   	    printf("Code Buffer:\n");
	   	    PrintSymbolBuffer(CodeBuffer);
	   	    printf("------------------------------------------\n\n");
	   	    if (Op0)
	   	        RemoveToken(&Op0);
	   	    if (Op1)
	   	        RemoveToken(&Op1);
	   	    if (Op2)
	   	        RemoveToken(&Op2);
	   	    RemoveSymbol(&OperatorSymbol);
	   	    if (Op0Symbol)
	   	        RemoveSymbol(&Op0Symbol);
	   	    if (Op1Symbol)
	   	        RemoveSymbol(&Op1Symbol);
	   	    if (Op2Symbol)
	   	        RemoveSymbol(&Op2Symbol);
	   	    if (TempSymbol)
	   	        RemoveSymbol(&TempSymbol);

	   BooleanExpressionExtractEnd(char * str, BOOL * WaitForWaitStatementBooleanExpression, PTOKEN CurrentIn)

	   	{
	   	    UINT64 BooleanExpressionSize = 0;
	   	    if (*WaitForWaitStatementBooleanExpression)
	   	    {
	   	        while (str[InputIdx + BooleanExpressionSize - 1] != ';')
	   	        {
	   	            BooleanExpressionSize += 1;
	   	        }
	   	        *WaitForWaitStatementBooleanExpression = FALSE;
	   	        return InputIdx + BooleanExpressionSize - 1;
	   	    }
	   	    else
	   	    {
	   	        int OpenParanthesesCount = 1;
	   	        if (!strcmp(CurrentIn->Value, "("))
	   	        {
	   	            OpenParanthesesCount++;
	   	        }
	   	        while (str[InputIdx + BooleanExpressionSize - 1] != '\0')
	   	        {
	   	            if (str[InputIdx + BooleanExpressionSize - 1] == ')')
	   	            {
	   	                OpenParanthesesCount--;
	   	                if (OpenParanthesesCount == 0)
	   	                {
	   	                    return InputIdx + BooleanExpressionSize - 1;
	   	                }
	   	            }
	   	            else if (str[InputIdx + BooleanExpressionSize - 1] == '(')
	   	            {
	   	                OpenParanthesesCount++;
	   	            }
	   	            BooleanExpressionSize++;
	   	        }
	   	    }
	   	    return -1;
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineBooleanExpresssionParse() (ok bool) { //col:937
	/*
	   ScriptEngineBooleanExpresssionParse(

	   	UINT64                    BooleanExpressionSize,
	   	PTOKEN                    FirstToken,
	   	PTOKEN_LIST               MatchedStack,
	   	PSYMBOL_BUFFER            CodeBuffer,
	   	char *                    str,
	   	char *                    c,
	   	PSCRIPT_ENGINE_ERROR_TYPE Error)

	   	{
	   	    PTOKEN_LIST Stack = NewTokenList();
	   	    PTOKEN State = NewToken(STATE_ID, "0");
	   	    Push(Stack, State);
	   	    printf("Boolean Expression: ");
	   	    printf("%s", FirstToken->Value);
	   	    for (int i = InputIdx - 1; i < BooleanExpressionSize; i++)
	   	    {
	   	        printf("%c", str[i]);
	   	    printf("\n\n");
	   	    PTOKEN EndToken = NewToken(END_OF_STACK, "$");
	   	    PTOKEN CurrentIn = CopyToken(FirstToken);
	   	    while (1)
	   	    {
	   	        TopToken       = Top(Stack);
	   	        int TerminalId = LalrGetTerminalId(CurrentIn);
	   	        StateId        = DecimalToSignedInt(TopToken->Value);
	   	        if (StateId == INVALID || TerminalId < 0)
	   	        {
	   	            *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	            break;
	   	        }
	   	        Action = LalrActionTable[StateId][TerminalId];

	   #ifdef _SCRIPT_ENGINE_LALR_DBG_EN

	   	    printf("Stack :\n");
	   	    PrintTokenList(Stack);
	   	    printf("Action : %d\n\n", Action);
	   	    if (Action == LALR_ACCEPT)
	   	    {
	   	        *Error = SCRIPT_ENGINE_ERROR_FREE;
	   	        break;
	   	    }
	   	    if (Action == INVALID)
	   	    {
	   	        *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	        break;
	   	    }
	   	    if (Action == 0)
	   	    {
	   	        *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	        break;
	   	    }
	   	    else if (Action >= 0)
	   	    {
	   	        StateId = Action;
	   	        Push(Stack, CurrentIn);
	   	        sprintf(buffer, "%d", StateId);
	   	        State = NewToken(STATE_ID, buffer);
	   	        Push(Stack, State);
	   	        CurrentIn = Scan(str, c);
	   	        if (InputIdx - 1 > BooleanExpressionSize)
	   	        {
	   	            InputIdx = InputIdxTemp;
	   	            *c       = Ctemp;
	   	            RemoveToken(&CurrentIn);
	   	            CurrentIn = CopyToken(EndToken);
	   	    else if (Action < 0)
	   	    {
	   	        StateId      = -Action;
	   	        Lhs          = &LalrLhs[StateId - 1];
	   	        RhsSize      = LalrGetRhsSize(StateId - 1);
	   	        for (int i = 0; i < 2 * RhsSize; i++)
	   	        {
	   	            Temp = Pop(Stack);
	   	            {
	   	                if (LalrIsOperandType(Temp))
	   	                {
	   	                    Operand = Temp;
	   	                    Push(MatchedStack, Operand);
	   	                    RemoveToken(&Temp);
	   	                RemoveToken(&Temp);
	   	        if (SemanticRule->Type == SEMANTIC_RULE)
	   	        {
	   	            {
	   	            }
	   	            else
	   	            {
	   	                CodeGen(MatchedStack, CodeBuffer, SemanticRule, Error);
	   	                if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	                {
	   	                    break;
	   	                }
	   	            }
	   	        }
	   	        Temp    = Top(Stack);
	   	        StateId = DecimalToSignedInt(Temp->Value);
	   	        Goto = LalrGotoTable[StateId][LalrGetNonTerminalId(Lhs)];
	   	        PTOKEN LhsCopy = CopyToken(Lhs);
	   	        sprintf(buffer, "%d", Goto);
	   	        State = NewToken(STATE_ID, buffer);
	   	        Push(Stack, LhsCopy);
	   	        Push(Stack, State);
	   	if (EndToken)
	   	    RemoveToken(&EndToken);
	   	if (Stack)
	   	    RemoveTokenList(Stack);
	   	if (CurrentIn)
	   	    RemoveToken(&CurrentIn);

	   NewSymbol(void)

	   	{
	   	    PSYMBOL Symbol;
	   	    Symbol        = (PSYMBOL)malloc(sizeof(SYMBOL));

	   NewStringSymbol(char * value)

	   	{
	   	    PSYMBOL Symbol;
	   	    int     BufferSize = (sizeof(unsigned long long) + (strlen(value))) / sizeof(SYMBOL) + 1;
	   	    Symbol             = (unsigned long long)malloc(BufferSize * sizeof(SYMBOL));
	   	    strcpy(&Symbol->Value, value);
	   	    SetType(&Symbol->Type, SYMBOL_STRING_TYPE);

	   GetStringSymbolSize(PSYMBOL Symbol)

	   	{
	   	    int Temp = (sizeof(unsigned long long) + (strlen(&Symbol->Value))) / sizeof(SYMBOL) + 1;
	   	    return Temp;
	   	}
	*/
	return true
}

func (s *scriptEngine) RemoveSymbol() (ok bool) { //col:1031
	/*
	   RemoveSymbol(PSYMBOL * Symbol)

	   	{
	   	    free(*Symbol);

	   PrintSymbol(PSYMBOL Symbol)

	   	{
	   	    if (Symbol->Type == SYMBOL_STRING_TYPE)
	   	    {
	   	        printf("Type:%llx, Value:%s\n", Symbol->Type, &Symbol->Value);
	   	        printf("Type:%llx, Value:0x%llx\n", Symbol->Type, Symbol->Value);

	   ToSymbol(PTOKEN Token, PSCRIPT_ENGINE_ERROR_TYPE Error)

	   	{
	   	    PSYMBOL Symbol = NewSymbol();
	   	    switch (Token->Type)
	   	    {
	   	    case GLOBAL_ID:
	   	        Symbol->Value = GetGlobalIdentifierVal(Token);
	   	        SetType(&Symbol->Type, SYMBOL_GLOBAL_ID_TYPE);
	   	        Symbol->Value = GetLocalIdentifierVal(Token);
	   	        SetType(&Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
	   	        Symbol->Value = DecimalToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        Symbol->Value = HexToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        Symbol->Value = OctalToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        Symbol->Value = BinaryToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        Symbol->Value = RegisterToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_REGISTER_TYPE);
	   	        Symbol->Value = PseudoRegToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_PSEUDO_REG_TYPE);
	   	        Symbol->Value = SemanticRuleToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_SEMANTIC_RULE_TYPE);
	   	        Symbol->Value = DecimalToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_TEMP_TYPE);
	   	        RemoveSymbol(&Symbol);
	   	        return NewStringSymbol(Token->Value);

	   NewSymbolBuffer(void)

	   	{
	   	    PSYMBOL_BUFFER SymbolBuffer;
	   	    SymbolBuffer          = (PSYMBOL_BUFFER)malloc(sizeof(*SymbolBuffer));
	   	    SymbolBuffer->Head    = (PSYMBOL)malloc(SymbolBuffer->Size * sizeof(SYMBOL));

	   RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer)

	   	{
	   	    free(SymbolBuffer->Message);
	   	    free(SymbolBuffer->Head);
	   	    free(SymbolBuffer);

	   PushSymbol(PSYMBOL_BUFFER SymbolBuffer, const PSYMBOL Symbol)

	   	{
	   	    uintptr_t Head      = (uintptr_t)SymbolBuffer->Head;
	   	    uintptr_t Pointer   = (uintptr_t)SymbolBuffer->Pointer;
	   	    PSYMBOL   WriteAddr = (PSYMBOL)(Head + Pointer * sizeof(SYMBOL));
	   	    if (Symbol->Type == SYMBOL_STRING_TYPE)
	   	    {
	   	        SymbolBuffer->Pointer += GetStringSymbolSize(Symbol);
	   	        if (SymbolBuffer->Pointer >= SymbolBuffer->Size - 1)
	   	        {
	   	            int NewSize = SymbolBuffer->Size;
	   	            do
	   	            {
	   	                NewSize *= 2;
	   	            } while (NewSize <= SymbolBuffer->Pointer);
	   	            PSYMBOL NewHead = (PSYMBOL)malloc(NewSize * sizeof(SYMBOL));
	   	            memcpy(NewHead, SymbolBuffer->Head, SymbolBuffer->Size * sizeof(SYMBOL));
	   	            free(SymbolBuffer->Head);
	   	        WriteAddr       = (PSYMBOL)((uintptr_t)SymbolBuffer->Head + (uintptr_t)Pointer * (uintptr_t)sizeof(SYMBOL));
	   	        strcpy((char *)&WriteAddr->Value, (char *)&Symbol->Value);
	   	        if (Pointer == SymbolBuffer->Size - 1)
	   	        {
	   	            PSYMBOL NewHead = (PSYMBOL)malloc(2 * SymbolBuffer->Size * sizeof(SYMBOL));
	   	            memcpy(NewHead, SymbolBuffer->Head, SymbolBuffer->Size * sizeof(SYMBOL));
	   	            free(SymbolBuffer->Head);

	   PrintSymbolBuffer(const PSYMBOL_BUFFER SymbolBuffer)

	   	{
	   	    PSYMBOL Symbol;
	   	    for (int i = 0; i < SymbolBuffer->Pointer;)
	   	    {
	   	        Symbol = SymbolBuffer->Head + i;
	   	        printf("%8x:", i);
	   	        PrintSymbol(Symbol);
	   	        if (Symbol->Type == SYMBOL_STRING_TYPE)
	   	        {
	   	            int temp = GetStringSymbolSize(Symbol);

	   RegisterToInt(char * str)

	   	{
	   	    for (int i = 0; i < REGISTER_MAP_LIST_LENGTH; i++)
	   	    {
	   	        if (!strcmp(str, RegisterMapList[i].Name))
	   	        {
	   	            return RegisterMapList[i].Type;
	   	        }
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (s *scriptEngine) PseudoRegToInt() (ok bool) { //col:1042
	/*
	   PseudoRegToInt(char * str)

	   	{
	   	    for (int i = 0; i < PSEUDO_REGISTER_MAP_LIST_LENGTH; i++)
	   	    {
	   	        if (!strcmp(str, PseudoRegisterMapList[i].Name))
	   	        {
	   	            return PseudoRegisterMapList[i].Type;
	   	        }
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (s *scriptEngine) SemanticRuleToInt() (ok bool) { //col:1053
	/*
	   SemanticRuleToInt(char * str)

	   	{
	   	    for (int i = 0; i < SEMANTIC_RULES_MAP_LIST_LENGTH; i++)
	   	    {
	   	        if (!strcmp(str, SemanticRulesMapList[i].Name))
	   	        {
	   	            return SemanticRulesMapList[i].Type;
	   	        }
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (s *scriptEngine) HandleError() (ok bool) { //col:1103
	/*
	   HandleError(PSCRIPT_ENGINE_ERROR_TYPE Error, char * str)

	   	{
	   	    unsigned int LineEnd;
	   	    for (int i = InputIdx;; i++)
	   	    {
	   	        if (str[i] == '\n' || str[i] == '\0')
	   	        {
	   	            LineEnd = i;
	   	            break;
	   	        }
	   	    }
	   	    int    MessageSize = 16 + 100 + (CurrentTokenIdx - CurrentLineIdx) + (LineEnd - CurrentLineIdx);
	   	    char * Message     = (char *)malloc(MessageSize);
	   	    strcpy(Message, "Line ");
	   	    sprintf(Line, "%d:\n", CurrentLine);
	   	    strcat(Message, Line);
	   	    strncat(Message, (str + CurrentLineIdx), LineEnd - CurrentLineIdx);
	   	    strcat(Message, "\n");
	   	    int  n     = (CurrentTokenIdx - CurrentLineIdx);
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        strncat(Message, &Space, 1);
	   	    strcat(Message, "^\n");
	   	    switch (*Error)
	   	    {
	   	    case SCRIPT_ENGINE_ERROR_SYNTAX:
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Invalid Syntax");
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Unknown Token");
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Unresolved Variable");
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Unhandled Semantic Rule");
	   	        strcat(Message, "Internal Error: ");
	   	        strcat(Message, "Please split the expression to many smaller expressions.");
	   	        strcat(Message, "Unkown Error: ");

	   GetGlobalIdentifierVal(PTOKEN Token)

	   	{
	   	    PTOKEN CurrentToken;
	   	    for (uintptr_t i = 0; i < IdTable->Pointer; i++)
	   	    {
	   	        CurrentToken = *(IdTable->Head + i);
	   	        if (!strcmp(Token->Value, CurrentToken->Value))
	   	        {
	   	            return (int)i;
	   	        }
	   	    }
	   	    return -1;
	   	}
	*/
	return true
}

func (s *scriptEngine) GetLocalIdentifierVal() (ok bool) { //col:1116
	/*
	   GetLocalIdentifierVal(PTOKEN Token)

	   	{
	   	    PTOKEN CurrentToken;
	   	    for (uintptr_t i = 0; i < IdTable->Pointer; i++)
	   	    {
	   	        CurrentToken = *(IdTable->Head + i);
	   	        if (!strcmp(Token->Value, CurrentToken->Value))
	   	        {
	   	            return (int)i;
	   	        }
	   	    }
	   	    return -1;
	   	}
	*/
	return true
}

func (s *scriptEngine) NewGlobalIdentifier() (ok bool) { //col:1137
	/*
	   NewGlobalIdentifier(PTOKEN Token)

	   	{
	   	    PTOKEN CopiedToken = CopyToken(Token);
	   	    IdTable            = Push(IdTable, CopiedToken);

	   NewLocalIdentifier(PTOKEN Token)

	   	{
	   	    PTOKEN CopiedToken = CopyToken(Token);
	   	    IdTable            = Push(IdTable, CopiedToken);

	   LalrGetRhsSize(int RuleId)

	   	{
	   	    int Counter = 0;
	   	    int N       = LalrRhsSize[RuleId];
	   	    for (int i = 0; i < N; i++)
	   	    {
	   	        if (LalrRhs[RuleId][i].Type != EPSILON && LalrRhs[RuleId][i].Type != SEMANTIC_RULE)
	   	        {
	   	            Counter++;
	   	        }
	   	    }
	   	    return Counter;
	   	}
	*/
	return true
}

func (s *scriptEngine) LalrIsOperandType() (ok bool) { //col:1185
	/*
	   LalrIsOperandType(PTOKEN Token)

	   	{
	   	    if (Token->Type == GLOBAL_ID)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == GLOBAL_UNRESOLVED_ID)
	   	    {
	   	        return TRUE;
	   	    }
	   	    if (Token->Type == LOCAL_ID)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == LOCAL_UNRESOLVED_ID)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == DECIMAL)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == HEX)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == OCTAL)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == BINARY)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == REGISTER)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == PSEUDO_REGISTER)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else if (Token->Type == TEMP)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

