package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-engine\code\common.c.back

type (
	Common interface {
		NewUnknownToken() (ok bool)      //col:111
		IsDecimal() (ok bool)            //col:118
		IsLetter() (ok bool)             //col:127
		IsBinary() (ok bool)             //col:136
		IsOctal() (ok bool)              //col:143
		NewTemp() (ok bool)              //col:171
		CleanTempList() (ok bool)        //col:178
		IsType1Func() (ok bool)          //col:190
		IsType2Func() (ok bool)          //col:202
		IsTwoOperandOperator() (ok bool) //col:214
		IsOneOperandOperator() (ok bool) //col:226
		IsType4Func() (ok bool)          //col:238
		IsType5Func() (ok bool)          //col:250
		IsType6Func() (ok bool)          //col:262
		IsType7Func() (ok bool)          //col:274
		IsType8Func() (ok bool)          //col:286
		IsNoneTerminal() (ok bool)       //col:293
		IsSemanticRule() (ok bool)       //col:300
		GetNonTerminalId() (ok bool)     //col:309
		GetTerminalId() (ok bool)        //col:382
		LalrGetNonTerminalId() (ok bool) //col:391
		LalrGetTerminalId() (ok bool)    //col:464
		IsEqual() (ok bool)              //col:498
		SetType() (ok bool)              //col:502
		DecimalToInt() (ok bool)         //col:550
	}
	common struct{}
)

func NewCommon() Common { return &common{} }

func (c *common) NewUnknownToken() (ok bool) { //col:111
	/*
	   NewUnknownToken()

	   	{
	   	    PTOKEN Token;
	   	    Token        = (PTOKEN)malloc(sizeof(TOKEN));
	   	    Token->Value = (char *)calloc(TOKEN_VALUE_MAX_LEN + 1, sizeof(char));
	   	    strcpy(Token->Value, "");

	   NewToken(TOKEN_TYPE Type, char * Value)

	   	{
	   	    PTOKEN Token = (PTOKEN)malloc(sizeof(TOKEN));
	   	    unsigned int Len = strlen(Value);
	   	    Token->Value     = (char *)calloc(Token->MaxLen + 1, sizeof(char));
	   	    strcpy(Token->Value, Value);

	   RemoveToken(PTOKEN * Token)

	   	{
	   	    free((*Token)->Value);
	   	    free(*Token);

	   PrintToken(PTOKEN Token)

	   	{
	   	    if (Token->Type == WHITE_SPACE)
	   	    {
	   	        printf("< :");
	   	        printf("<'%s' : ", Token->Value);
	   	    switch (Token->Type)
	   	    {
	   	    case GLOBAL_ID:
	   	        printf(" GLOBAL_ID>\n");
	   	        printf(" GLOBAL_UNRESOLVED_ID>\n");
	   	        printf(" LOCAL_ID>\n");
	   	        printf(" LOCAL_UNRESOLVED_ID>\n");
	   	        printf(" STATE_ID>\n");
	   	        printf(" DECIMAL>\n");
	   	        printf(" HEX>\n");
	   	        printf(" OCTAL>\n");
	   	        printf(" BINARY>\n");
	   	        printf(" SPECIAL_TOKEN>\n");
	   	        printf(" KEYWORD>\n");
	   	        printf(" WHITE_SPACE>\n");
	   	        printf(" COMMENT>\n");
	   	        printf(" REGISTER>\n");
	   	        printf(" PSEUDO_REGISTER>\n");
	   	        printf(" SEMANTIC_RULE>\n");
	   	        printf(" NON_TERMINAL>\n");
	   	        printf(" END_OF_STACK>\n");
	   	        printf(" STRING>\n");
	   	        printf(" TEMP>\n");
	   	        printf(" UNKNOWN>\n");
	   	        printf(" ERROR>\n");

	   Append(PTOKEN Token, char c)

	   	{
	   	    if (Token->Len >= Token->MaxLen - 1)
	   	    {
	   	        Token->MaxLen *= 2;
	   	        char * NewValue = (char *)calloc(Token->MaxLen + 1, sizeof(char));
	   	        strncpy(NewValue, Token->Value, Token->Len);
	   	        free(Token->Value);
	   	    strncat(Token->Value, &c, 1);

	   CopyToken(PTOKEN Token)

	   	{
	   	    PTOKEN TokenCopy  = (PTOKEN)malloc(sizeof(TOKEN));
	   	    TokenCopy->Value  = (char *)calloc(strlen(Token->Value) + 1, sizeof(char));
	   	    strcpy(TokenCopy->Value, Token->Value);

	   NewTokenList(void)

	   	{
	   	    PTOKEN_LIST TokenList;
	   	    TokenList = (PTOKEN_LIST)malloc(sizeof(*TokenList));
	   	    TokenList->Head = (PTOKEN *)malloc(TokenList->Size * sizeof(PTOKEN));

	   RemoveTokenList(PTOKEN_LIST TokenList)

	   	{
	   	    PTOKEN Token;
	   	    for (uintptr_t i = 0; i < TokenList->Pointer; i++)
	   	    {
	   	        Token = *(TokenList->Head + i);
	   	        RemoveToken(&Token);
	   	    free(TokenList->Head);
	   	    free(TokenList);

	   PrintTokenList(PTOKEN_LIST TokenList)

	   	{
	   	    PTOKEN Token;
	   	    for (uintptr_t i = 0; i < TokenList->Pointer; i++)
	   	    {
	   	        Token = *(TokenList->Head + i);
	   	        PrintToken(Token);

	   Push(PTOKEN_LIST TokenList, PTOKEN Token)

	   	{
	   	    uintptr_t Head      = (uintptr_t)TokenList->Head;
	   	    uintptr_t Pointer   = (uintptr_t)TokenList->Pointer;
	   	    PTOKEN *  WriteAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
	   	    if (Pointer == TokenList->Size - 1)
	   	    {
	   	        PTOKEN * NewHead = (PTOKEN *)malloc(2 * TokenList->Size * sizeof(PTOKEN));
	   	        memcpy(NewHead, TokenList->Head, TokenList->Size * sizeof(PTOKEN));
	   	        free(TokenList->Head);

	   Pop(PTOKEN_LIST TokenList)

	   	{
	   	    if (TokenList->Pointer > 0)
	   	        TokenList->Pointer--;
	   	    uintptr_t Head     = (uintptr_t)TokenList->Head;
	   	    uintptr_t Pointer  = (uintptr_t)TokenList->Pointer;
	   	    PTOKEN *  ReadAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));

	   Top(PTOKEN_LIST TokenList)

	   	{
	   	    uintptr_t Head     = (uintptr_t)TokenList->Head;
	   	    uintptr_t Pointer  = (uintptr_t)TokenList->Pointer - 1;
	   	    PTOKEN *  ReadAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));

	   IsHex(char c)

	   	{
	   	    if ((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F'))
	   	        return 1;
	   	    else
	   	        return 0;
	   	}
	*/
	return true
}

func (c *common) IsDecimal() (ok bool) { //col:118
	/*
	   IsDecimal(char c)

	   	{
	   	    if (c >= '0' && c <= '9')
	   	        return 1;
	   	    else
	   	        return 0;
	   	}
	*/
	return true
}

func (c *common) IsLetter() (ok bool) { //col:127
	/*
	   IsLetter(char c)

	   	{
	   	    if ((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'))
	   	        return 1;
	   	    else
	   	    {
	   	        return 0;
	   	    }
	   	}
	*/
	return true
}

func (c *common) IsBinary() (ok bool) { //col:136
	/*
	   IsBinary(char c)

	   	{
	   	    if (c == '0' || c == '1')
	   	        return 1;
	   	    else
	   	    {
	   	        return 0;
	   	    }
	   	}
	*/
	return true
}

func (c *common) IsOctal() (ok bool) { //col:143
	/*
	   IsOctal(char c)

	   	{
	   	    if (c >= '0' && c <= '7')
	   	        return 1;
	   	    else
	   	        return 0;
	   	}
	*/
	return true
}

func (c *common) NewTemp() (ok bool) { //col:171
	/*
	   NewTemp(PSCRIPT_ENGINE_ERROR_TYPE Error)

	   	{
	   	    static unsigned int TempID = 0;
	   	    int                 i;
	   	    for (i = 0; i < MAX_TEMP_COUNT; i++)
	   	    {
	   	        if (TempMap[i] == 0)
	   	        {
	   	            TempID     = i;
	   	            TempMap[i] = 1;
	   	            break;
	   	        }
	   	    }
	   	    if (i == MAX_TEMP_COUNT)
	   	    {
	   	        *Error = SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL;
	   	    }
	   	    PTOKEN Temp = NewUnknownToken();
	   	    sprintf(TempValue, "%d", TempID);
	   	    strcpy(Temp->Value, TempValue);

	   FreeTemp(PTOKEN Temp)

	   	{
	   	    int id = DecimalToInt(Temp->Value);
	   	    if (Temp->Type == TEMP)
	   	    {
	   	        TempMap[id] = 0;
	   	    }
	   	}
	*/
	return true
}

func (c *common) CleanTempList() (ok bool) { //col:178
	/*
	   CleanTempList(void)

	   	{
	   	    for (int i = 0; i < MAX_TEMP_COUNT; i++)
	   	    {
	   	        TempMap[i] = 0;
	   	    }
	   	}
	*/
	return true
}

func (c *common) IsType1Func() (ok bool) { //col:190
	/*
	   IsType1Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = ONEOPFUNC1_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, OneOpFunc1[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsType2Func() (ok bool) { //col:202
	/*
	   IsType2Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = ONEOPFUNC2_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, OneOpFunc2[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsTwoOperandOperator() (ok bool) { //col:214
	/*
	   IsTwoOperandOperator(PTOKEN Operator)

	   	{
	   	    unsigned int n = OPERATORS_TWO_OPERAND_LIST_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, OperatorsTwoOperandList[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsOneOperandOperator() (ok bool) { //col:226
	/*
	   IsOneOperandOperator(PTOKEN Operator)

	   	{
	   	    unsigned int n = OPERATORS_ONE_OPERAND_LIST_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, OperatorsOneOperandList[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsType4Func() (ok bool) { //col:238
	/*
	   IsType4Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = VARARGFUNC1_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, VarArgFunc1[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsType5Func() (ok bool) { //col:250
	/*
	   IsType5Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = ZEROOPFUNC1_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, ZeroOpFunc1[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsType6Func() (ok bool) { //col:262
	/*
	   IsType6Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = TWOOPFUNC1_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, TwoOpFunc1[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsType7Func() (ok bool) { //col:274
	/*
	   IsType7Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = TWOOPFUNC2_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, TwoOpFunc2[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsType8Func() (ok bool) { //col:286
	/*
	   IsType8Func(PTOKEN Operator)

	   	{
	   	    unsigned int n = THREEOPFUNC1_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(Operator->Value, ThreeOpFunc1[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) IsNoneTerminal() (ok bool) { //col:293
	/*
	   IsNoneTerminal(PTOKEN Token)

	   	{
	   	    if (Token->Value[0] >= 'A' && Token->Value[0] <= 'Z')
	   	        return 1;
	   	    else
	   	        return 0;
	   	}
	*/
	return true
}

func (c *common) IsSemanticRule() (ok bool) { //col:300
	/*
	   IsSemanticRule(PTOKEN Token)

	   	{
	   	        return 1;
	   	    else
	   	        return 0;
	   	}
	*/
	return true
}

func (c *common) GetNonTerminalId() (ok bool) { //col:309
	/*
	   GetNonTerminalId(PTOKEN Token)

	   	{
	   	    for (int i = 0; i < NONETERMINAL_COUNT; i++)
	   	    {
	   	        if (!strcmp(Token->Value, NoneTerminalMap[i]))
	   	            return i;
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (c *common) GetTerminalId() (ok bool) { //col:382
	/*
	   GetTerminalId(PTOKEN Token)

	   	{
	   	    for (int i = 0; i < TERMINAL_COUNT; i++)
	   	    {
	   	        if (Token->Type == HEX)
	   	        {
	   	            if (!strcmp("_hex", TerminalMap[i]))
	   	                return i;
	   	        }
	   	        else if (Token->Type == GLOBAL_ID || Token->Type == GLOBAL_UNRESOLVED_ID)
	   	        {
	   	            if (!strcmp("_global_id", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == LOCAL_ID || Token->Type == LOCAL_UNRESOLVED_ID)
	   	        {
	   	            if (!strcmp("_local_id", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == REGISTER)
	   	        {
	   	            if (!strcmp("_register", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == PSEUDO_REGISTER)
	   	        {
	   	            if (!strcmp("_pseudo_register", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == DECIMAL)
	   	        {
	   	            if (!strcmp("_decimal", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == BINARY)
	   	        {
	   	            if (!strcmp("_binary", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == OCTAL)
	   	        {
	   	            if (!strcmp("_octal", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == STRING)
	   	        {
	   	            if (!strcmp("_string", TerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            if (!strcmp(Token->Value, TerminalMap[i]))
	   	                return i;
	   	        }
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (c *common) LalrGetNonTerminalId() (ok bool) { //col:391
	/*
	   LalrGetNonTerminalId(PTOKEN Token)

	   	{
	   	    for (int i = 0; i < LALR_NONTERMINAL_COUNT; i++)
	   	    {
	   	        if (!strcmp(Token->Value, LalrNoneTerminalMap[i]))
	   	            return i;
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (c *common) LalrGetTerminalId() (ok bool) { //col:464
	/*
	   LalrGetTerminalId(PTOKEN Token)

	   	{
	   	    for (int i = 0; i < LALR_TERMINAL_COUNT; i++)
	   	    {
	   	        if (Token->Type == HEX)
	   	        {
	   	            if (!strcmp("_hex", LalrTerminalMap[i]))
	   	                return i;
	   	        }
	   	        else if (Token->Type == GLOBAL_ID || Token->Type == GLOBAL_UNRESOLVED_ID)
	   	        {
	   	            if (!strcmp("_global_id", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == LOCAL_ID || Token->Type == LOCAL_UNRESOLVED_ID)
	   	        {
	   	            if (!strcmp("_local_id", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == REGISTER)
	   	        {
	   	            if (!strcmp("_register", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == PSEUDO_REGISTER)
	   	        {
	   	            if (!strcmp("_pseudo_register", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == DECIMAL)
	   	        {
	   	            if (!strcmp("_decimal", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == BINARY)
	   	        {
	   	            if (!strcmp("_binary", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == OCTAL)
	   	        {
	   	            if (!strcmp("_octal", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else if (Token->Type == STRING)
	   	        {
	   	            if (!strcmp("_string", LalrTerminalMap[i]))
	   	            {
	   	                return i;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            if (!strcmp(Token->Value, LalrTerminalMap[i]))
	   	                return i;
	   	        }
	   	    }
	   	    return INVALID;
	   	}
	*/
	return true
}

func (c *common) IsEqual() (ok bool) { //col:498
	/*
	   IsEqual(const PTOKEN Token1, const PTOKEN Token2)

	   	{
	   	    if (Token1->Type == Token2->Type)
	   	    {
	   	        if (Token1->Type == SPECIAL_TOKEN)
	   	        {
	   	            if (!strcmp(Token1->Value, Token2->Value))
	   	            {
	   	                return 1;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    if (Token1->Type == GLOBAL_ID && Token2->Type == GLOBAL_UNRESOLVED_ID)
	   	    {
	   	        return 1;
	   	    }
	   	    if (Token1->Type == GLOBAL_UNRESOLVED_ID && Token2->Type == GLOBAL_ID)
	   	    {
	   	        return 1;
	   	    }
	   	    if (Token1->Type == LOCAL_ID && Token2->Type == LOCAL_UNRESOLVED_ID)
	   	    {
	   	        return 1;
	   	    }
	   	    if (Token1->Type == LOCAL_UNRESOLVED_ID && Token2->Type == LOCAL_ID)
	   	    {
	   	        return 1;
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (c *common) SetType() (ok bool) { //col:502
	/*
	   SetType(unsigned long long * Val, unsigned char Type)

	   	{
	   	    *Val = (unsigned long long int)Type;
	   	}
	*/
	return true
}

func (c *common) DecimalToInt() (ok bool) { //col:550
	/*
	   DecimalToInt(char * str)

	   	{
	   	    unsigned long long int Acc = 0;
	   	    size_t                 Len;
	   	    Len = strlen(str);
	   	    for (int i = 0; i < Len; i++)
	   	    {
	   	        Acc *= 10;
	   	        Acc += (str[i] - '0');

	   DecimalToSignedInt(char * str)

	   	{
	   	    long long int Acc = 0;
	   	    size_t        Len;
	   	    if (str[0] == '-')
	   	    {
	   	        Len = strlen(str);
	   	        for (int i = 1; i < Len; i++)
	   	        {
	   	            Acc *= 10;
	   	            Acc += (str[i] - '0');
	   	        Len = strlen(str);
	   	        for (int i = 0; i < Len; i++)
	   	        {
	   	            Acc *= 10;
	   	            Acc += (str[i] - '0');

	   HexToInt(char * str)

	   	{
	   	    char                   temp;
	   	    size_t                 len = strlen(str);
	   	    for (int i = 0; i < len; i++)
	   	    {
	   	        Acc <<= 4;
	   	        if (str[i] >= '0' && str[i] <= '9')
	   	        {
	   	            temp = str[i] - '0';
	   	        }
	   	        else if (str[i] >= 'a' && str[i] <= 'f')
	   	        {
	   	            temp = str[i] - 'a' + 10;
	   	        }
	   	        else
	   	        {
	   	            temp = str[i] - 'A' + 10;
	   	        }
	   	        Acc += temp;
	   	    }
	   	    return Acc;
	   	}
	*/
	return true
}

