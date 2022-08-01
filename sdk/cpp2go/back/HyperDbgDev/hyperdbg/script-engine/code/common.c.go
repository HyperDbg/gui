package code
//back\HyperDbgDev\hyperdbg\script-engine\code\common.c.back

type (
Common interface{
NewUnknownToken()(ok bool)//col:11
NewToken()(ok bool)//col:22
RemoveToken()(ok bool)//col:29
PrintToken()(ok bool)//col:109
Append()(ok bool)//col:122
CopyToken()(ok bool)//col:132
NewTokenList()(ok bool)//col:141
RemoveTokenList()(ok bool)//col:153
PrintTokenList()(ok bool)//col:162
Push()(ok bool)//col:179
Pop()(ok bool)//col:188
Top()(ok bool)//col:195
IsHex()(ok bool)//col:202
IsDecimal()(ok bool)//col:209
IsLetter()(ok bool)//col:218
IsBinary()(ok bool)//col:227
IsOctal()(ok bool)//col:234
NewTemp()(ok bool)//col:258
FreeTemp()(ok bool)//col:266
CleanTempList()(ok bool)//col:273
IsType1Func()(ok bool)//col:285
IsType2Func()(ok bool)//col:297
IsTwoOperandOperator()(ok bool)//col:309
IsOneOperandOperator()(ok bool)//col:321
IsType4Func()(ok bool)//col:333
IsType5Func()(ok bool)//col:345
IsType6Func()(ok bool)//col:357
IsType7Func()(ok bool)//col:369
IsType8Func()(ok bool)//col:381
IsNoneTerminal()(ok bool)//col:388
IsSemanticRule()(ok bool)//col:395
GetNonTerminalId()(ok bool)//col:404
GetTerminalId()(ok bool)//col:477
LalrGetNonTerminalId()(ok bool)//col:486
LalrGetTerminalId()(ok bool)//col:559
IsEqual()(ok bool)//col:593
SetType()(ok bool)//col:597
DecimalToInt()(ok bool)//col:609
DecimalToSignedInt()(ok bool)//col:634
HexToInt()(ok bool)//col:658
OctalToInt()(ok bool)//col:670
BinaryToInt()(ok bool)//col:682
}
)

func NewCommon() { return & common{} }

func (c *common)NewUnknownToken()(ok bool){//col:11
/*NewUnknownToken()
{
    PTOKEN Token;
    Token        = (PTOKEN)malloc(sizeof(TOKEN));
    Token->Value = (char *)calloc(TOKEN_VALUE_MAX_LEN + 1, sizeof(char));
    strcpy(Token->Value, "");
    Token->Type   = UNKNOWN;
    Token->Len    = 0;
    Token->MaxLen = TOKEN_VALUE_MAX_LEN;
    return Token;
}*/
return true
}

func (c *common)NewToken()(ok bool){//col:22
/*NewToken(TOKEN_TYPE Type, char * Value)
{
    PTOKEN Token = (PTOKEN)malloc(sizeof(TOKEN));
    unsigned int Len = strlen(Value);
    Token->Type      = Type;
    Token->Len       = Len;
    Token->MaxLen    = Len;
    Token->Value     = (char *)calloc(Token->MaxLen + 1, sizeof(char));
    strcpy(Token->Value, Value);
    return Token;
}*/
return true
}

func (c *common)RemoveToken()(ok bool){//col:29
/*RemoveToken(PTOKEN * Token)
{
    free((*Token)->Value);
    free(*Token);
    *Token = NULL;
    return;
}*/
return true
}

func (c *common)PrintToken()(ok bool){//col:109
/*PrintToken(PTOKEN Token)
{
    if (Token->Type == WHITE_SPACE)
    {
        printf("< :");
    }
    else
    {
        printf("<'%s' : ", Token->Value);
    }
    switch (Token->Type)
    {
    case GLOBAL_ID:
        printf(" GLOBAL_ID>\n");
        break;
    case GLOBAL_UNRESOLVED_ID:
        printf(" GLOBAL_UNRESOLVED_ID>\n");
        break;
    case LOCAL_ID:
        printf(" LOCAL_ID>\n");
        break;
    case LOCAL_UNRESOLVED_ID:
        printf(" LOCAL_UNRESOLVED_ID>\n");
        break;
    case STATE_ID:
        printf(" STATE_ID>\n");
        break;
    case DECIMAL:
        printf(" DECIMAL>\n");
        break;
    case HEX:
        printf(" HEX>\n");
        break;
    case OCTAL:
        printf(" OCTAL>\n");
        break;
    case BINARY:
        printf(" BINARY>\n");
        break;
    case SPECIAL_TOKEN:
        printf(" SPECIAL_TOKEN>\n");
        break;
    case KEYWORD:
        printf(" KEYWORD>\n");
        break;
    case WHITE_SPACE:
        printf(" WHITE_SPACE>\n");
        break;
    case COMMENT:
        printf(" COMMENT>\n");
        break;
    case REGISTER:
        printf(" REGISTER>\n");
        break;
    case PSEUDO_REGISTER:
        printf(" PSEUDO_REGISTER>\n");
        break;
    case SEMANTIC_RULE:
        printf(" SEMANTIC_RULE>\n");
        break;
    case NON_TERMINAL:
        printf(" NON_TERMINAL>\n");
        break;
    case END_OF_STACK:
        printf(" END_OF_STACK>\n");
        break;
    case STRING:
        printf(" STRING>\n");
        break;
    case TEMP:
        printf(" TEMP>\n");
        break;
    case UNKNOWN:
        printf(" UNKNOWN>\n");
        break;
    default:
        printf(" ERROR>\n");
        break;
    }
}*/
return true
}

func (c *common)Append()(ok bool){//col:122
/*Append(PTOKEN Token, char c)
{
    if (Token->Len >= Token->MaxLen - 1)
    {
        Token->MaxLen *= 2;
        char * NewValue = (char *)calloc(Token->MaxLen + 1, sizeof(char));
        strncpy(NewValue, Token->Value, Token->Len);
        free(Token->Value);
        Token->Value = NewValue;
    }
    strncat(Token->Value, &c, 1);
    Token->Len++;
}*/
return true
}

func (c *common)CopyToken()(ok bool){//col:132
/*CopyToken(PTOKEN Token)
{
    PTOKEN TokenCopy  = (PTOKEN)malloc(sizeof(TOKEN));
    TokenCopy->Type   = Token->Type;
    TokenCopy->MaxLen = Token->MaxLen;
    TokenCopy->Len    = Token->Len;
    TokenCopy->Value  = (char *)calloc(strlen(Token->Value) + 1, sizeof(char));
    strcpy(TokenCopy->Value, Token->Value);
    return TokenCopy;
}*/
return true
}

func (c *common)NewTokenList()(ok bool){//col:141
/*NewTokenList(void)
{
    PTOKEN_LIST TokenList;
    TokenList = (PTOKEN_LIST)malloc(sizeof(*TokenList));
    TokenList->Pointer = 0;
    TokenList->Size    = TOKEN_LIST_INIT_SIZE;
    TokenList->Head = (PTOKEN *)malloc(TokenList->Size * sizeof(PTOKEN));
    return TokenList;
}*/
return true
}

func (c *common)RemoveTokenList()(ok bool){//col:153
/*RemoveTokenList(PTOKEN_LIST TokenList)
{
    PTOKEN Token;
    for (uintptr_t i = 0; i < TokenList->Pointer; i++)
    {
        Token = *(TokenList->Head + i);
        RemoveToken(&Token);
    }
    free(TokenList->Head);
    free(TokenList);
    return;
}*/
return true
}

func (c *common)PrintTokenList()(ok bool){//col:162
/*PrintTokenList(PTOKEN_LIST TokenList)
{
    PTOKEN Token;
    for (uintptr_t i = 0; i < TokenList->Pointer; i++)
    {
        Token = *(TokenList->Head + i);
        PrintToken(Token);
    }
}*/
return true
}

func (c *common)Push()(ok bool){//col:179
/*Push(PTOKEN_LIST TokenList, PTOKEN Token)
{
    uintptr_t Head      = (uintptr_t)TokenList->Head;
    uintptr_t Pointer   = (uintptr_t)TokenList->Pointer;
    PTOKEN *  WriteAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
    *WriteAddr = Token;
    TokenList->Pointer++;
    if (Pointer == TokenList->Size - 1)
    {
        PTOKEN * NewHead = (PTOKEN *)malloc(2 * TokenList->Size * sizeof(PTOKEN));
        memcpy(NewHead, TokenList->Head, TokenList->Size * sizeof(PTOKEN));
        free(TokenList->Head);
        TokenList->Size = TokenList->Size * 2;
        TokenList->Head = NewHead;
    }
    return TokenList;
}*/
return true
}

func (c *common)Pop()(ok bool){//col:188
/*Pop(PTOKEN_LIST TokenList)
{
    if (TokenList->Pointer > 0)
        TokenList->Pointer--;
    uintptr_t Head     = (uintptr_t)TokenList->Head;
    uintptr_t Pointer  = (uintptr_t)TokenList->Pointer;
    PTOKEN *  ReadAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
    return *ReadAddr;
}*/
return true
}

func (c *common)Top()(ok bool){//col:195
/*Top(PTOKEN_LIST TokenList)
{
    uintptr_t Head     = (uintptr_t)TokenList->Head;
    uintptr_t Pointer  = (uintptr_t)TokenList->Pointer - 1;
    PTOKEN *  ReadAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
    return *ReadAddr;
}*/
return true
}

func (c *common)IsHex()(ok bool){//col:202
/*IsHex(char c)
{
    if ((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F'))
        return 1;
    else
        return 0;
}*/
return true
}

func (c *common)IsDecimal()(ok bool){//col:209
/*IsDecimal(char c)
{
    if (c >= '0' && c <= '9')
        return 1;
    else
        return 0;
}*/
return true
}

func (c *common)IsLetter()(ok bool){//col:218
/*IsLetter(char c)
{
    if ((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'))
        return 1;
    else
    {
        return 0;
    }
}*/
return true
}

func (c *common)IsBinary()(ok bool){//col:227
/*IsBinary(char c)
{
    if (c == '0' || c == '1')
        return 1;
    else
    {
        return 0;
    }
}*/
return true
}

func (c *common)IsOctal()(ok bool){//col:234
/*IsOctal(char c)
{
    if (c >= '0' && c <= '7')
        return 1;
    else
        return 0;
}*/
return true
}

func (c *common)NewTemp()(ok bool){//col:258
/*NewTemp(PSCRIPT_ENGINE_ERROR_TYPE Error)
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
    char   TempValue[8];
    sprintf(TempValue, "%d", TempID);
    strcpy(Temp->Value, TempValue);
    Temp->Type = TEMP;
    return Temp;
}*/
return true
}

func (c *common)FreeTemp()(ok bool){//col:266
/*FreeTemp(PTOKEN Temp)
{
    int id = DecimalToInt(Temp->Value);
    if (Temp->Type == TEMP)
    {
        TempMap[id] = 0;
    }
}*/
return true
}

func (c *common)CleanTempList()(ok bool){//col:273
/*CleanTempList(void)
{
    for (int i = 0; i < MAX_TEMP_COUNT; i++)
    {
        TempMap[i] = 0;
    }
}*/
return true
}

func (c *common)IsType1Func()(ok bool){//col:285
/*IsType1Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsType2Func()(ok bool){//col:297
/*IsType2Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsTwoOperandOperator()(ok bool){//col:309
/*IsTwoOperandOperator(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsOneOperandOperator()(ok bool){//col:321
/*IsOneOperandOperator(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsType4Func()(ok bool){//col:333
/*IsType4Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsType5Func()(ok bool){//col:345
/*IsType5Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsType6Func()(ok bool){//col:357
/*IsType6Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsType7Func()(ok bool){//col:369
/*IsType7Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsType8Func()(ok bool){//col:381
/*IsType8Func(PTOKEN Operator)
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
}*/
return true
}

func (c *common)IsNoneTerminal()(ok bool){//col:388
/*IsNoneTerminal(PTOKEN Token)
{
    if (Token->Value[0] >= 'A' && Token->Value[0] <= 'Z')
        return 1;
    else
        return 0;
}*/
return true
}

func (c *common)IsSemanticRule()(ok bool){//col:395
/*IsSemanticRule(PTOKEN Token)
{
        return 1;
    else
        return 0;
}*/
return true
}

func (c *common)GetNonTerminalId()(ok bool){//col:404
/*GetNonTerminalId(PTOKEN Token)
{
    for (int i = 0; i < NONETERMINAL_COUNT; i++)
    {
        if (!strcmp(Token->Value, NoneTerminalMap[i]))
            return i;
    }
    return INVALID;
}*/
return true
}

func (c *common)GetTerminalId()(ok bool){//col:477
/*GetTerminalId(PTOKEN Token)
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
}*/
return true
}

func (c *common)LalrGetNonTerminalId()(ok bool){//col:486
/*LalrGetNonTerminalId(PTOKEN Token)
{
    for (int i = 0; i < LALR_NONTERMINAL_COUNT; i++)
    {
        if (!strcmp(Token->Value, LalrNoneTerminalMap[i]))
            return i;
    }
    return INVALID;
}*/
return true
}

func (c *common)LalrGetTerminalId()(ok bool){//col:559
/*LalrGetTerminalId(PTOKEN Token)
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
}*/
return true
}

func (c *common)IsEqual()(ok bool){//col:593
/*IsEqual(const PTOKEN Token1, const PTOKEN Token2)
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
}*/
return true
}

func (c *common)SetType()(ok bool){//col:597
/*SetType(unsigned long long * Val, unsigned char Type)
{
    *Val = (unsigned long long int)Type;
}*/
return true
}

func (c *common)DecimalToInt()(ok bool){//col:609
/*DecimalToInt(char * str)
{
    unsigned long long int Acc = 0;
    size_t                 Len;
    Len = strlen(str);
    for (int i = 0; i < Len; i++)
    {
        Acc *= 10;
        Acc += (str[i] - '0');
    }
    return Acc;
}*/
return true
}

func (c *common)DecimalToSignedInt()(ok bool){//col:634
/*DecimalToSignedInt(char * str)
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
        }
        return -Acc;
    }
    else
    {
        Len = strlen(str);
        for (int i = 0; i < Len; i++)
        {
            Acc *= 10;
            Acc += (str[i] - '0');
        }
        return Acc;
    }
}*/
return true
}

func (c *common)HexToInt()(ok bool){//col:658
/*HexToInt(char * str)
{
    char                   temp;
    size_t                 len = strlen(str);
    unsigned long long int Acc = 0;
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
}*/
return true
}

func (c *common)OctalToInt()(ok bool){//col:670
/*OctalToInt(char * str)
{
    size_t                 Len;
    unsigned long long int Acc = 0;
    Len = strlen(str);
    for (int i = 0; i < Len; i++)
    {
        Acc <<= 3;
        Acc += (str[i] - '0');
    }
    return Acc;
}*/
return true
}

func (c *common)BinaryToInt()(ok bool){//col:682
/*BinaryToInt(char * str)
{
    size_t                 Len;
    unsigned long long int Acc = 0;
    Len = strlen(str);
    for (int i = 0; i < Len; i++)
    {
        Acc <<= 1;
        Acc += (str[i] - '0');
    }
    return Acc;
}*/
return true
}



