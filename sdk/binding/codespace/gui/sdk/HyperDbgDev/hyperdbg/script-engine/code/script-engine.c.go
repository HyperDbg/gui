package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-engine\code\script-engine.c.back

type (
	ScriptEngine interface {
		ScriptEngineConvertNameToAddress() (ok bool)                     //col:16
		ScriptEngineLoadFileSymbol() (ok bool)                           //col:20
		ScriptEngineSetTextMessageCallback() (ok bool)                   //col:24
		ScriptEngineUnloadAllSymbols() (ok bool)                         //col:28
		ScriptEngineUnloadModuleSymbol() (ok bool)                       //col:32
		ScriptEngineSearchSymbolForMask() (ok bool)                      //col:36
		ScriptEngineGetFieldOffset() (ok bool)                           //col:40
		ScriptEngineGetDataTypeSize() (ok bool)                          //col:44
		ScriptEngineCreateSymbolTableForDisassembler() (ok bool)         //col:48
		ScriptEngineConvertFileToPdbPath() (ok bool)                     //col:52
		ScriptEngineSymbolInitLoad() (ok bool)                           //col:60
		ScriptEngineShowDataBasedOnSymbolTypes() (ok bool)               //col:68
		ScriptEngineSymbolAbortLoading() (ok bool)                       //col:72
		ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails() (ok bool) //col:76
		ScriptEngineParse() (ok bool)                                    //col:253
		CodeGen() (ok bool)                                              //col:954
		BooleanExpressionExtractEnd() (ok bool)                          //col:992
		ScriptEngineBooleanExpresssionParse() (ok bool)                  //col:1136
		NewSymbol() (ok bool)                                            //col:1144
		NewStringSymbol() (ok bool)                                      //col:1153
		GetStringSymbolSize() (ok bool)                                  //col:1158
		RemoveSymbol() (ok bool)                                         //col:1164
		PrintSymbol() (ok bool)                                          //col:1175
		ToSymbol() (ok bool)                                             //col:1230
		NewSymbolBuffer() (ok bool)                                      //col:1240
		RemoveSymbolBuffer() (ok bool)                                   //col:1246
		PushSymbol() (ok bool)                                           //col:1286
		PrintSymbolBuffer() (ok bool)                                    //col:1305
		RegisterToInt() (ok bool)                                        //col:1316
		PseudoRegToInt() (ok bool)                                       //col:1327
		SemanticRuleToInt() (ok bool)                                    //col:1338
		HandleError() (ok bool)                                          //col:1391
		GetGlobalIdentifierVal() (ok bool)                               //col:1404
		GetLocalIdentifierVal() (ok bool)                                //col:1417
		NewGlobalIdentifier() (ok bool)                                  //col:1423
		NewLocalIdentifier() (ok bool)                                   //col:1429
		LalrGetRhsSize() (ok bool)                                       //col:1442
		LalrIsOperandType() (ok bool)                                    //col:1490
	}
	scriptEngine struct{}
)

func NewScriptEngine() ScriptEngine { return &scriptEngine{} }

func (s *scriptEngine) ScriptEngineConvertNameToAddress() (ok bool) { //col:16
	/*
	   ScriptEngineConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound)

	   	{
	   	    return SymConvertNameToAddress(FunctionOrVariableName, WasFound);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineLoadFileSymbol() (ok bool) { //col:20
	/*
	   ScriptEngineLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName)

	   	{
	   	    return SymLoadFileSymbol(BaseAddress, PdbFileName);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineSetTextMessageCallback() (ok bool) { //col:24
	/*
	   ScriptEngineSetTextMessageCallback(PVOID Handler)

	   	{
	   	    SymSetTextMessageCallback(Handler);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineUnloadAllSymbols() (ok bool) { //col:28
	/*
	   ScriptEngineUnloadAllSymbols()

	   	{
	   	    return SymUnloadAllSymbols();
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineUnloadModuleSymbol() (ok bool) { //col:32
	/*
	   ScriptEngineUnloadModuleSymbol(char * ModuleName)

	   	{
	   	    return SymUnloadModuleSymbol(ModuleName);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineSearchSymbolForMask() (ok bool) { //col:36
	/*
	   ScriptEngineSearchSymbolForMask(const char * SearchMask)

	   	{
	   	    return SymSearchSymbolForMask(SearchMask);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineGetFieldOffset() (ok bool) { //col:40
	/*
	   ScriptEngineGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)

	   	{
	   	    return SymGetFieldOffset(TypeName, FieldName, FieldOffset);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineGetDataTypeSize() (ok bool) { //col:44
	/*
	   ScriptEngineGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize)

	   	{
	   	    return SymGetDataTypeSize(TypeName, TypeSize);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineCreateSymbolTableForDisassembler() (ok bool) { //col:48
	/*
	   ScriptEngineCreateSymbolTableForDisassembler(void * CallbackFunction)

	   	{
	   	    return SymCreateSymbolTableForDisassembler(CallbackFunction);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineConvertFileToPdbPath() (ok bool) { //col:52
	/*
	   ScriptEngineConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath)

	   	{
	   	    return SymConvertFileToPdbPath(LocalFilePath, ResultPath);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineSymbolInitLoad() (ok bool) { //col:60
	/*
	   ScriptEngineSymbolInitLoad(PVOID        BufferToStoreDetails,

	   	UINT32       StoredLength,
	   	BOOLEAN      DownloadIfAvailable,
	   	const char * SymbolPath,
	   	BOOLEAN      IsSilentLoad)

	   	{
	   	    return SymbolInitLoad(BufferToStoreDetails, StoredLength, DownloadIfAvailable, SymbolPath, IsSilentLoad);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineShowDataBasedOnSymbolTypes() (ok bool) { //col:68
	/*
	   ScriptEngineShowDataBasedOnSymbolTypes(const char * TypeName,

	   	UINT64       Address,
	   	BOOLEAN      IsStruct,
	   	PVOID        BufferAddress,
	   	const char * AdditionalParameters)

	   	{
	   	    return SymShowDataBasedOnSymbolTypes(TypeName, Address, IsStruct, BufferAddress, AdditionalParameters);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineSymbolAbortLoading() (ok bool) { //col:72
	/*
	   ScriptEngineSymbolAbortLoading()

	   	{
	   	    return SymbolAbortLoading();
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails() (ok bool) { //col:76
	/*
	   ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath, char * PdbFilePath, char * GuidAndAgeDetails)

	   	{
	   	    return SymConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails);
	   	}
	*/
	return true
}

func (s *scriptEngine) ScriptEngineParse() (ok bool) { //col:253
	/*
	   ScriptEngineParse(char * str)

	   	{
	   	    PTOKEN_LIST Stack = NewTokenList();
	   	    PTOKEN_LIST    MatchedStack = NewTokenList();
	   	    PSYMBOL_BUFFER CodeBuffer   = NewSymbolBuffer();
	   	    SCRIPT_ENGINE_ERROR_TYPE Error        = SCRIPT_ENGINE_ERROR_FREE;
	   	    char *                   ErrorMessage = NULL;
	   	    static FirstCall = 1;
	   	    if (FirstCall)
	   	    {
	   	        IdTable   = NewTokenList();
	   	        FirstCall = 0;
	   	    }
	   	    PTOKEN TopToken = NewUnknownToken();
	   	    int  NonTerminalId;
	   	    int  TerminalId;
	   	    int  RuleId;
	   	    char c;
	   	    BOOL WaitForWaitStatementBooleanExpression = FALSE;
	   	    InputIdx       = 0;
	   	    CurrentLine    = 0;
	   	    CurrentLineIdx = 0;
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
	   	        CodeBuffer->Message = ErrorMessage;
	   	        RemoveTokenList(Stack);
	   	        RemoveTokenList(MatchedStack);
	   	        RemoveToken(&CurrentIn);
	   	        return CodeBuffer;
	   	    }
	   	    do
	   	    {
	   	        RemoveToken(&TopToken);
	   	        TopToken = Pop(Stack);

	   #ifdef _SCRIPT_ENGINE_LL1_DBG_EN

	   	printf("\nTop Token :\n");
	   	PrintToken(TopToken);
	   	printf("\nCurrent Input :\n");
	   	PrintToken(CurrentIn);
	   	printf("\n");

	   #endif

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
	   	    }
	   	    else
	   	    {
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
	   	        }
	   	    }
	   	}
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

	   #endif

	   	    } while (TopToken->Type != END_OF_STACK);
	   	    if (Error != SCRIPT_ENGINE_ERROR_FREE)
	   	    {
	   	        ErrorMessage = HandleError(&Error, str);
	   	        CleanTempList();
	   	    }
	   	    else
	   	    {
	   	        ErrorMessage = NULL;
	   	    }
	   	    CodeBuffer->Message = ErrorMessage;
	   	    if (Stack)
	   	        RemoveTokenList(Stack);
	   	    if (MatchedStack)
	   	        RemoveTokenList(MatchedStack);
	   	    if (CurrentIn)
	   	        RemoveToken(&CurrentIn);
	   	    if (TopToken)
	   	        RemoveToken(&TopToken);
	   	    return CodeBuffer;
	   	}
	*/
	return true
}

func (s *scriptEngine) CodeGen() (ok bool) { //col:954
	/*
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

	   #ifdef _SCRIPT_ENGINE_CODEGEN_DBG_EN

	   	printf("Operator :\n");
	   	PrintToken(Operator);
	   	printf("\n");
	   	printf("Semantic Stack:\n");
	   	PrintTokenList(MatchedStack);
	   	printf("\n");
	   	printf("Code Buffer:\n");
	   	PrintSymbolBuffer(CodeBuffer);
	   	printf(".\n.\n.\n\n");

	   #endif

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
	   	        }
	   	        else if (Op1->Type == LOCAL_UNRESOLVED_ID)
	   	        {
	   	            Op1Symbol = NewSymbol();
	   	            free(Op1Symbol->Value);
	   	            Op1Symbol->Value = NewLocalIdentifier(Op1);
	   	            SetType(&Op1Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
	   	        }
	   	        else
	   	        {
	   	            Op1Symbol = ToSymbol(Op1, Error);
	   	        }
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
	   	        UINT32         OperandCount = 0;
	   	        do
	   	        {
	   	            if (Op1)
	   	            {
	   	                RemoveToken(&Op1);
	   	            }
	   	            Op1 = Pop(MatchedStack);
	   	            if (Op1->Type != SEMANTIC_RULE)
	   	            {
	   	                Op1Symbol = ToSymbol(Op1, Error);
	   	                FreeTemp(Op1);
	   	                PushSymbol(TempStack, Op1Symbol);
	   	                RemoveSymbol(&Op1Symbol);
	   	                OperandCount++;
	   	                if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	                {
	   	                    RemoveSymbolBuffer(TempStack);
	   	                    break;
	   	                }
	   	            }
	   	        if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	        {
	   	            break;
	   	        }
	   	        Op0       = Pop(MatchedStack);
	   	        Op0Symbol = ToSymbol(Op0, Error);
	   	        FreeTemp(Op0);
	   	        char * Format = Op0->Value;
	   	        PSYMBOL OperandCountSymbol = NewSymbol();
	   	        OperandCountSymbol->Type   = SYMBOL_VARIABLE_COUNT_TYPE;
	   	        OperandCountSymbol->Value  = OperandCount;
	   	        PushSymbol(CodeBuffer, Op0Symbol);
	   	        PushSymbol(CodeBuffer, OperandCountSymbol);
	   	        RemoveSymbol(&OperandCountSymbol);
	   	        if (*Error != SCRIPT_ENGINE_ERROR_FREE)
	   	        {
	   	            RemoveSymbolBuffer(TempStack);
	   	            break;
	   	        }
	   	        PSYMBOL FirstArg = (PSYMBOL)((unsigned long long)CodeBuffer->Head +
	   	                                     (unsigned long long)(CodeBuffer->Pointer * sizeof(SYMBOL)));
	   	        PSYMBOL Symbol;
	   	        int     ArgCount = TempStack->Pointer;
	   	        for (int i = TempStack->Pointer - 1; i >= 0; i--)
	   	        {
	   	            Symbol = TempStack->Head + i;
	   	            PushSymbol(CodeBuffer, Symbol);
	   	        }
	   	        RemoveSymbolBuffer(TempStack);
	   	        UINT32 i   = 0;
	   	        char * Str = Format;
	   	        do
	   	        {
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
	   	    }
	   	    {
	   	        Op0 = Pop(MatchedStack);
	   	    }
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
	   	    }
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
	   	    }
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
	   	    }
	   	    {
	   	        PTOKEN OperatorCopy = CopyToken(Operator);
	   	        Push(MatchedStack, OperatorCopy);
	   	    }
	   	    {
	   	        UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	        PushSymbol(CodeBuffer, OperatorSymbol);
	   	        PSYMBOL JumpAddressSymbol = NewSymbol();
	   	        JumpAddressSymbol->Type   = SYMBOL_NUM_TYPE;
	   	        JumpAddressSymbol->Value  = 0xffffffffffffffff;
	   	        PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	        RemoveSymbol(&JumpAddressSymbol);
	   	        Op0       = Pop(MatchedStack);
	   	        Op0Symbol = ToSymbol(Op0, Error);
	   	        PushSymbol(CodeBuffer, Op0Symbol);
	   	        char str[20] = {0};
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
	   	        JumpAddressSymbol->Value         = CurrentPointer + 2;
	   	        RemoveToken(&JumpSemanticAddressToken);
	   	        PSYMBOL JumpInstruction = NewSymbol();
	   	        JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JumpInstruction->Value  = FUNC_JMP;
	   	        PushSymbol(CodeBuffer, JumpInstruction);
	   	        RemoveSymbol(&JumpInstruction);
	   	        JumpAddressSymbol        = NewSymbol();
	   	        JumpAddressSymbol->Type  = SYMBOL_NUM_TYPE;
	   	        JumpAddressSymbol->Value = 0xffffffffffffffff;
	   	        PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	        RemoveSymbol(&JumpAddressSymbol);
	   	        char str[20] = {0};
	   	        sprintf(str, "%llu", CurrentPointer);
	   	        PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	        Push(MatchedStack, CurrentAddressToken);
	   	    }
	   	    {
	   	        UINT64  CurrentPointer           = CodeBuffer->Pointer;
	   	        PTOKEN  JumpSemanticAddressToken = Pop(MatchedStack);
	   	        PSYMBOL JumpAddressSymbol;
	   	        {
	   	            UINT64 JumpSemanticAddress = DecimalToInt(JumpSemanticAddressToken->Value);
	   	            JumpAddressSymbol          = (PSYMBOL)(CodeBuffer->Head + JumpSemanticAddress + 1);
	   	            JumpAddressSymbol->Value   = CurrentPointer;
	   	            RemoveToken(&JumpSemanticAddressToken);
	   	            JumpSemanticAddressToken = Pop(MatchedStack);
	   	        }
	   	        RemoveToken(&JumpSemanticAddressToken);
	   	    }
	   	    {
	   	        PTOKEN OperatorCopy = CopyToken(Operator);
	   	        Push(MatchedStack, OperatorCopy);
	   	        char str[20] = {0};
	   	        sprintf(str, "%llu", CodeBuffer->Pointer);
	   	        PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	        Push(MatchedStack, CurrentAddressToken);
	   	    }
	   	    {
	   	        UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	        RemoveSymbol(&OperatorSymbol);
	   	        OperatorSymbol = ToSymbol(JzToken, Error);
	   	        RemoveToken(&JzToken);
	   	        PSYMBOL JumpAddressSymbol = NewSymbol();
	   	        JumpAddressSymbol->Type   = SYMBOL_NUM_TYPE;
	   	        JumpAddressSymbol->Value  = 0xffffffffffffffff;
	   	        Op0       = Pop(MatchedStack);
	   	        Op0Symbol = ToSymbol(Op0, Error);
	   	        PTOKEN StartOfWhileToken = Pop(MatchedStack);
	   	        char str[20];
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
	   	        JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JumpInstruction->Value  = FUNC_JMP;
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
	   	            JumpAddressSymbol->Value = CurrentPointer;
	   	        } while (TRUE);
	   	        RemoveToken(&JumpAddressToken);
	   	    }
	   	    {
	   	        PTOKEN OperatorCopy = CopyToken(Operator);
	   	        Push(MatchedStack, OperatorCopy);
	   	        char str[20];
	   	        sprintf(str, "%llu", CodeBuffer->Pointer);
	   	        PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	        Push(MatchedStack, CurrentAddressToken);
	   	    }
	   	    {
	   	        PSYMBOL JumpInstruction = NewSymbol();
	   	        JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JumpInstruction->Value  = FUNC_JNZ;
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

	   #ifdef _SCRIPT_ENGINE_LL1_DBG_EN

	   	printf("Jz Jump Address = %d\n", JumpAddress);

	   #endif

	   	            JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
	   	            JumpAddressSymbol->Value = CurrentPointer;
	   	        } while (TRUE);
	   	        RemoveToken(&JumpAddressToken);
	   	    }
	   	    {
	   	        PTOKEN OperatorCopy = CopyToken(Operator);
	   	        Push(MatchedStack, OperatorCopy);
	   	        char str[20] = {0};
	   	        sprintf(str, "%llu", CodeBuffer->Pointer);
	   	        PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	        Push(MatchedStack, CurrentAddressToken);
	   	    }
	   	    {
	   	        PSYMBOL JnzInstruction = NewSymbol();
	   	        JnzInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JnzInstruction->Value  = FUNC_JZ;
	   	        PushSymbol(CodeBuffer, JnzInstruction);
	   	        RemoveSymbol(&JnzInstruction);
	   	        PSYMBOL JnzAddressSymbol = NewSymbol();
	   	        JnzAddressSymbol->Type   = SYMBOL_NUM_TYPE;
	   	        JnzAddressSymbol->Value  = 0xffffffffffffffff;
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
	   	        JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JumpInstruction->Value  = FUNC_JMP;
	   	        PushSymbol(CodeBuffer, JumpInstruction);
	   	        RemoveSymbol(&JumpInstruction);
	   	        PSYMBOL JumpAddressSymbol = NewSymbol();
	   	        JumpAddressSymbol->Type   = SYMBOL_NUM_TYPE;
	   	        JumpAddressSymbol->Value  = 0xffffffffffffffff;
	   	        PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	        RemoveSymbol(&JumpAddressSymbol);
	   	        PTOKEN StartOfForAddressToken = Pop(MatchedStack);
	   	        char str[20] = {0};
	   	        sprintf(str, "%llu", CodeBuffer->Pointer);
	   	        PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	        Push(MatchedStack, CurrentAddressToken);
	   	        Push(MatchedStack, StartOfForAddressToken);
	   	    }
	   	    {
	   	        PSYMBOL JumpInstruction = NewSymbol();
	   	        JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JumpInstruction->Value  = FUNC_JMP;
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
	   	        JumpAddressSymbol->Value = CurrentPointer;
	   	        char str[20] = {0};
	   	        sprintf(str, "%llu", JumpAddress - 4);
	   	        PTOKEN JzAddressToken = NewToken(DECIMAL, str);
	   	        Push(MatchedStack, JzAddressToken);
	   	        Push(MatchedStack, IncDecToken);
	   	        Push(MatchedStack, JumpAddressToken);
	   	    }
	   	    {
	   	        PSYMBOL JumpInstruction = NewSymbol();
	   	        JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	        JumpInstruction->Value  = FUNC_JMP;
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
	   	        UINT64 CurrentPointer = CodeBuffer->Pointer;
	   	        do
	   	        {
	   	            RemoveToken(&JumpAddressToken);
	   	            JumpAddressToken = Pop(MatchedStack);
	   	            {
	   	                break;
	   	            }
	   	            else
	   	            {
	   	                JumpAddress = DecimalToInt(JumpAddressToken->Value);
	   	                JumpAddressSymbol        = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
	   	                JumpAddressSymbol->Value = CurrentPointer;
	   	            }
	   	        } while (TRUE);
	   	        RemoveToken(&JumpAddressToken);
	   	    }
	   	    {
	   	        PTOKEN_LIST TempStack = NewTokenList();
	   	        PTOKEN      TempToken;
	   	        do
	   	        {
	   	            TempToken = Pop(MatchedStack);
	   	            {
	   	                Push(MatchedStack, TempToken);
	   	                UINT64 CurrentPointer = CodeBuffer->Pointer + 1;
	   	                char   str[20];
	   	                sprintf(str, "%llu", CurrentPointer);
	   	                PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
	   	                Push(MatchedStack, CurrentAddressToken);
	   	                PSYMBOL JumpInstruction = NewSymbol();
	   	                JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	                JumpInstruction->Value  = FUNC_JMP;
	   	                PushSymbol(CodeBuffer, JumpInstruction);
	   	                RemoveSymbol(&JumpInstruction);
	   	                PSYMBOL JumpAddressSymbol = NewSymbol();
	   	                JumpAddressSymbol->Type   = SYMBOL_NUM_TYPE;
	   	                JumpAddressSymbol->Value  = 0xffffffffffffffff;
	   	                PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	                RemoveSymbol(&JumpAddressSymbol);
	   	                do
	   	                {
	   	                    TempToken = Pop(TempStack);
	   	                    Push(MatchedStack, TempToken);
	   	                } while (TempStack->Pointer != 0);
	   	                break;
	   	            }
	   	            else if (MatchedStack->Pointer == 0)
	   	            {
	   	                *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                RemoveToken(&TempToken);
	   	                break;
	   	            }
	   	            else
	   	            {
	   	                Push(TempStack, TempToken);
	   	            }
	   	        } while (TRUE);
	   	        RemoveTokenList(TempStack);
	   	    }
	   	    {
	   	        PTOKEN_LIST TempStack = NewTokenList();
	   	        PTOKEN      TempToken;
	   	        do
	   	        {
	   	            TempToken = Pop(MatchedStack);
	   	            {
	   	                Push(MatchedStack, TempToken);
	   	                PSYMBOL JumpInstruction = NewSymbol();
	   	                JumpInstruction->Type   = SYMBOL_SEMANTIC_RULE_TYPE;
	   	                JumpInstruction->Value  = FUNC_JMP;
	   	                PushSymbol(CodeBuffer, JumpInstruction);
	   	                RemoveSymbol(&JumpInstruction);
	   	                TempToken = Pop(TempStack);
	   	                Push(MatchedStack, TempToken);
	   	                PSYMBOL JumpAddressSymbol = NewSymbol();
	   	                JumpAddressSymbol->Type   = SYMBOL_NUM_TYPE;
	   	                JumpAddressSymbol->Value  = DecimalToInt(TempToken->Value);
	   	                PushSymbol(CodeBuffer, JumpAddressSymbol);
	   	                RemoveSymbol(&JumpAddressSymbol);
	   	                do
	   	                {
	   	                    TempToken = Pop(TempStack);
	   	                    Push(MatchedStack, TempToken);
	   	                } while (TempStack->Pointer != 0);
	   	                break;
	   	            }
	   	            else if (MatchedStack->Pointer == 0)
	   	            {
	   	                *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
	   	                break;
	   	            }
	   	            else
	   	            {
	   	                Push(TempStack, TempToken);
	   	            }
	   	        } while (TRUE);
	   	        RemoveTokenList(TempStack);
	   	    }
	   	    else
	   	    {
	   	        *Error = SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE;
	   	    }
	   	    break;
	   	}

	   #ifdef _SCRIPT_ENGINE_CODEGEN_DBG_EN

	   	printf("Semantic Stack:\n");
	   	PrintTokenList(MatchedStack);
	   	printf("\n");
	   	printf("Code Buffer:\n");
	   	PrintSymbolBuffer(CodeBuffer);
	   	printf("------------------------------------------\n\n");

	   #endif

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
	   	    return;
	   	}
	*/
	return true
}

func (s *scriptEngine) BooleanExpressionExtractEnd() (ok bool) { //col:992
	/*
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

func (s *scriptEngine) ScriptEngineBooleanExpresssionParse() (ok bool) { //col:1136
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

	   #ifdef _SCRIPT_ENGINE_LALR_DBG_EN

	   	printf("Boolean Expression: ");
	   	printf("%s", FirstToken->Value);
	   	for (int i = InputIdx - 1; i < BooleanExpressionSize; i++)
	   	{
	   	    printf("%c", str[i]);
	   	}
	   	printf("\n\n");

	   #endif

	   	PTOKEN EndToken = NewToken(END_OF_STACK, "$");
	   	PTOKEN CurrentIn = CopyToken(FirstToken);
	   	PTOKEN TopToken     = NULL;
	   	PTOKEN Lhs          = NULL;
	   	PTOKEN Temp         = NULL;
	   	PTOKEN Operand      = NULL;
	   	PTOKEN SemanticRule = NULL;
	   	int          Action       = INVALID;
	   	int          StateId      = 0;
	   	int          Goto         = 0;
	   	int          InputPointer = 0;
	   	int          RhsSize      = 0;
	   	unsigned int InputIdxTemp;
	   	char         Ctemp;
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

	   #endif

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
	   	            char buffer[20] = {0};
	   	            sprintf(buffer, "%d", StateId);
	   	            State = NewToken(STATE_ID, buffer);
	   	            Push(Stack, State);
	   	            InputIdxTemp = InputIdx;
	   	            Ctemp        = *c;
	   	            CurrentIn = Scan(str, c);
	   	            if (InputIdx - 1 > BooleanExpressionSize)
	   	            {
	   	                InputIdx = InputIdxTemp;
	   	                *c       = Ctemp;
	   	                RemoveToken(&CurrentIn);
	   	                CurrentIn = CopyToken(EndToken);
	   	            }
	   	        }
	   	        else if (Action < 0)
	   	        {
	   	            StateId      = -Action;
	   	            Lhs          = &LalrLhs[StateId - 1];
	   	            RhsSize      = LalrGetRhsSize(StateId - 1);
	   	            SemanticRule = &LalrSemanticRules[StateId - 1];
	   	            for (int i = 0; i < 2 * RhsSize; i++)
	   	            {
	   	                Temp = Pop(Stack);
	   	                {
	   	                    if (LalrIsOperandType(Temp))
	   	                    {
	   	                        Operand = Temp;
	   	                        Push(MatchedStack, Operand);
	   	                    }
	   	                    else
	   	                    {
	   	                        RemoveToken(&Temp);
	   	                    }
	   	                }
	   	                else
	   	                {
	   	                    RemoveToken(&Temp);
	   	                }
	   	            }
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
	   	            char buffer[20] = {0};
	   	            sprintf(buffer, "%d", Goto);
	   	            State = NewToken(STATE_ID, buffer);
	   	            Push(Stack, LhsCopy);
	   	            Push(Stack, State);
	   	        }
	   	    }
	   	    if (EndToken)
	   	        RemoveToken(&EndToken);
	   	    if (Stack)
	   	        RemoveTokenList(Stack);
	   	    if (CurrentIn)
	   	        RemoveToken(&CurrentIn);
	   	    return;
	   	}
	*/
	return true
}

func (s *scriptEngine) NewSymbol() (ok bool) { //col:1144
	/*
	   NewSymbol(void)

	   	{
	   	    PSYMBOL Symbol;
	   	    Symbol        = (PSYMBOL)malloc(sizeof(SYMBOL));
	   	    Symbol->Value = 0;
	   	    Symbol->Type  = 0;
	   	    return Symbol;
	   	}
	*/
	return true
}

func (s *scriptEngine) NewStringSymbol() (ok bool) { //col:1153
	/*
	   NewStringSymbol(char * value)

	   	{
	   	    PSYMBOL Symbol;
	   	    int     BufferSize = (sizeof(unsigned long long) + (strlen(value))) / sizeof(SYMBOL) + 1;
	   	    Symbol             = (unsigned long long)malloc(BufferSize * sizeof(SYMBOL));
	   	    strcpy(&Symbol->Value, value);
	   	    SetType(&Symbol->Type, SYMBOL_STRING_TYPE);
	   	    return Symbol;
	   	}
	*/
	return true
}

func (s *scriptEngine) GetStringSymbolSize() (ok bool) { //col:1158
	/*
	   GetStringSymbolSize(PSYMBOL Symbol)

	   	{
	   	    int Temp = (sizeof(unsigned long long) + (strlen(&Symbol->Value))) / sizeof(SYMBOL) + 1;
	   	    return Temp;
	   	}
	*/
	return true
}

func (s *scriptEngine) RemoveSymbol() (ok bool) { //col:1164
	/*
	   RemoveSymbol(PSYMBOL * Symbol)

	   	{
	   	    free(*Symbol);
	   	    *Symbol = NULL;
	   	    return;
	   	}
	*/
	return true
}

func (s *scriptEngine) PrintSymbol() (ok bool) { //col:1175
	/*
	   PrintSymbol(PSYMBOL Symbol)

	   	{
	   	    if (Symbol->Type == SYMBOL_STRING_TYPE)
	   	    {
	   	        printf("Type:%llx, Value:%s\n", Symbol->Type, &Symbol->Value);
	   	    }
	   	    else
	   	    {
	   	        printf("Type:%llx, Value:0x%llx\n", Symbol->Type, Symbol->Value);
	   	    }
	   	}
	*/
	return true
}

func (s *scriptEngine) ToSymbol() (ok bool) { //col:1230
	/*
	   ToSymbol(PTOKEN Token, PSCRIPT_ENGINE_ERROR_TYPE Error)

	   	{
	   	    PSYMBOL Symbol = NewSymbol();
	   	    switch (Token->Type)
	   	    {
	   	    case GLOBAL_ID:
	   	        Symbol->Value = GetGlobalIdentifierVal(Token);
	   	        SetType(&Symbol->Type, SYMBOL_GLOBAL_ID_TYPE);
	   	        return Symbol;
	   	    case LOCAL_ID:
	   	        Symbol->Value = GetLocalIdentifierVal(Token);
	   	        SetType(&Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
	   	        return Symbol;
	   	    case DECIMAL:
	   	        Symbol->Value = DecimalToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        return Symbol;
	   	    case HEX:
	   	        Symbol->Value = HexToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        return Symbol;
	   	    case OCTAL:
	   	        Symbol->Value = OctalToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        return Symbol;
	   	    case BINARY:
	   	        Symbol->Value = BinaryToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
	   	        return Symbol;
	   	    case REGISTER:
	   	        Symbol->Value = RegisterToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_REGISTER_TYPE);
	   	        return Symbol;
	   	    case PSEUDO_REGISTER:
	   	        Symbol->Value = PseudoRegToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_PSEUDO_REG_TYPE);
	   	        return Symbol;
	   	    case SEMANTIC_RULE:
	   	        Symbol->Value = SemanticRuleToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_SEMANTIC_RULE_TYPE);
	   	        return Symbol;
	   	    case TEMP:
	   	        Symbol->Value = DecimalToInt(Token->Value);
	   	        SetType(&Symbol->Type, SYMBOL_TEMP_TYPE);
	   	        return Symbol;
	   	    case STRING:
	   	        RemoveSymbol(&Symbol);
	   	        return NewStringSymbol(Token->Value);
	   	    default:
	   	        *Error        = SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE;
	   	        Symbol->Type  = INVALID;
	   	        Symbol->Value = INVALID;
	   	        return Symbol;
	   	    }
	   	}
	*/
	return true
}

func (s *scriptEngine) NewSymbolBuffer() (ok bool) { //col:1240
	/*
	   NewSymbolBuffer(void)

	   	{
	   	    PSYMBOL_BUFFER SymbolBuffer;
	   	    SymbolBuffer          = (PSYMBOL_BUFFER)malloc(sizeof(*SymbolBuffer));
	   	    SymbolBuffer->Pointer = 0;
	   	    SymbolBuffer->Size    = SYMBOL_BUFFER_INIT_SIZE;
	   	    SymbolBuffer->Head    = (PSYMBOL)malloc(SymbolBuffer->Size * sizeof(SYMBOL));
	   	    SymbolBuffer->Message = NULL;
	   	    return SymbolBuffer;
	   	}
	*/
	return true
}

func (s *scriptEngine) RemoveSymbolBuffer() (ok bool) { //col:1246
	/*
	   RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer)

	   	{
	   	    free(SymbolBuffer->Message);
	   	    free(SymbolBuffer->Head);
	   	    free(SymbolBuffer);
	   	}
	*/
	return true
}

func (s *scriptEngine) PushSymbol() (ok bool) { //col:1286
	/*
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
	   	            SymbolBuffer->Size = NewSize;
	   	            SymbolBuffer->Head = NewHead;
	   	        }
	   	        WriteAddr       = (PSYMBOL)((uintptr_t)SymbolBuffer->Head + (uintptr_t)Pointer * (uintptr_t)sizeof(SYMBOL));
	   	        WriteAddr->Type = Symbol->Type;
	   	        strcpy((char *)&WriteAddr->Value, (char *)&Symbol->Value);
	   	    }
	   	    else
	   	    {
	   	        *WriteAddr = *Symbol;
	   	        SymbolBuffer->Pointer++;
	   	        if (Pointer == SymbolBuffer->Size - 1)
	   	        {
	   	            PSYMBOL NewHead = (PSYMBOL)malloc(2 * SymbolBuffer->Size * sizeof(SYMBOL));
	   	            memcpy(NewHead, SymbolBuffer->Head, SymbolBuffer->Size * sizeof(SYMBOL));
	   	            free(SymbolBuffer->Head);
	   	            SymbolBuffer->Size *= 2;
	   	            SymbolBuffer->Head = NewHead;
	   	        }
	   	    }
	   	    return SymbolBuffer;
	   	}
	*/
	return true
}

func (s *scriptEngine) PrintSymbolBuffer() (ok bool) { //col:1305
	/*
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
	   	            i += temp;
	   	        }
	   	        else
	   	        {
	   	            i++;
	   	        }
	   	    }
	   	}
	*/
	return true
}

func (s *scriptEngine) RegisterToInt() (ok bool) { //col:1316
	/*
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

func (s *scriptEngine) PseudoRegToInt() (ok bool) { //col:1327
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

func (s *scriptEngine) SemanticRuleToInt() (ok bool) { //col:1338
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

func (s *scriptEngine) HandleError() (ok bool) { //col:1391
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
	   	    char Line[16] = {0};
	   	    sprintf(Line, "%d:\n", CurrentLine);
	   	    strcat(Message, Line);
	   	    strncat(Message, (str + CurrentLineIdx), LineEnd - CurrentLineIdx);
	   	    strcat(Message, "\n");
	   	    char Space = ' ';
	   	    int  n     = (CurrentTokenIdx - CurrentLineIdx);
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        strncat(Message, &Space, 1);
	   	    }
	   	    strcat(Message, "^\n");
	   	    switch (*Error)
	   	    {
	   	    case SCRIPT_ENGINE_ERROR_SYNTAX:
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Invalid Syntax");
	   	        return Message;
	   	    case SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN:
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Unknown Token");
	   	        return Message;
	   	    case SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE:
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Unresolved Variable");
	   	        return Message;
	   	    case SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE:
	   	        strcat(Message, "Syntax Error: ");
	   	        strcat(Message, "Unhandled Semantic Rule");
	   	        return Message;
	   	    case SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL:
	   	        strcat(Message, "Internal Error: ");
	   	        strcat(Message, "Please split the expression to many smaller expressions.");
	   	        return Message;
	   	    default:
	   	        strcat(Message, "Unkown Error: ");
	   	        return Message;
	   	    }
	   	}
	*/
	return true
}

func (s *scriptEngine) GetGlobalIdentifierVal() (ok bool) { //col:1404
	/*
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

func (s *scriptEngine) GetLocalIdentifierVal() (ok bool) { //col:1417
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

func (s *scriptEngine) NewGlobalIdentifier() (ok bool) { //col:1423
	/*
	   NewGlobalIdentifier(PTOKEN Token)

	   	{
	   	    PTOKEN CopiedToken = CopyToken(Token);
	   	    IdTable            = Push(IdTable, CopiedToken);
	   	    return IdTable->Pointer - 1;
	   	}
	*/
	return true
}

func (s *scriptEngine) NewLocalIdentifier() (ok bool) { //col:1429
	/*
	   NewLocalIdentifier(PTOKEN Token)

	   	{
	   	    PTOKEN CopiedToken = CopyToken(Token);
	   	    IdTable            = Push(IdTable, CopiedToken);
	   	    return IdTable->Pointer - 1;
	   	}
	*/
	return true
}

func (s *scriptEngine) LalrGetRhsSize() (ok bool) { //col:1442
	/*
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

func (s *scriptEngine) LalrIsOperandType() (ok bool) { //col:1490
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

