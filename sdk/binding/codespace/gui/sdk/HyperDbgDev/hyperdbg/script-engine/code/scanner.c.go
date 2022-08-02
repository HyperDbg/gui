package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-engine\code\scanner.c.back

type (
	Scanner interface {
		GetToken() (ok bool)   //col:510
		IsKeyword() (ok bool)  //col:530
		IsRegister() (ok bool) //col:536
		IsId() (ok bool)       //col:540
	}
	scanner struct{}
)

func NewScanner() Scanner { return &scanner{} }

func (s *scanner) GetToken() (ok bool) { //col:510
	/*
	   GetToken(char * c, char * str)

	   	{
	   	    PTOKEN Token = NewUnknownToken();
	   	    switch (*c)
	   	    {
	   	    case '"':
	   	        do
	   	        {
	   	            *c = sgetc(str);
	   	            if (*c == '\\')
	   	            {
	   	                *c = sgetc(str);
	   	                if (*c == 'n')
	   	                {
	   	                    Append(Token, '\n');
	   	                if (*c == '\\')
	   	                {
	   	                    Append(Token, '\\');
	   	                else if (*c == 't')
	   	                {
	   	                    Append(Token, '\t');
	   	                else if (*c == '"')
	   	                {
	   	                    Append(Token, '"');
	   	                    *c          = sgetc(str);
	   	            else if (*c == '"')
	   	            {
	   	                break;
	   	            }
	   	            else
	   	            {
	   	                Append(Token, *c);
	   	        } while (1);
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, "~");
	   	        *c          = sgetc(str);
	   	        *c = sgetc(str);
	   	        if (*c == '+')
	   	        {
	   	            strcpy(Token->Value, "++");
	   	            *c          = sgetc(str);
	   	        else if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "+=");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "+");
	   	        *c = sgetc(str);
	   	        if (*c == '-')
	   	        {
	   	            strcpy(Token->Value, "--");
	   	            *c          = sgetc(str);
	   	        else if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "-=");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "-");
	   	        *c = sgetc(str);
	   	        if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "*=");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "*");
	   	        *c = sgetc(str);
	   	        if (*c == '>')
	   	        {
	   	            strcpy(Token->Value, ">>");
	   	            *c          = sgetc(str);
	   	        else if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, ">=");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, ">");
	   	        *c = sgetc(str);
	   	        if (*c == '<')
	   	        {
	   	            strcpy(Token->Value, "<<");
	   	            *c          = sgetc(str);
	   	        else if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "<=");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "<");
	   	        *c = sgetc(str);
	   	        if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "/=");
	   	            *c          = sgetc(str);
	   	        else if (*c == '/')
	   	        {
	   	            do
	   	            {
	   	                *c = sgetc(str);
	   	            } while (*c != '\n' && (int)*c != EOF);
	   	            *c          = sgetc(str);
	   	        else if (*c == '*')
	   	        {
	   	            do
	   	            {
	   	                *c = sgetc(str);
	   	                if (*c == '*')
	   	                {
	   	                    *c = sgetc(str);
	   	                    if (*c == '/')
	   	                    {
	   	                        Token->Type = COMMENT;
	   	                        *c          = sgetc(str);
	   	                if ((int)*c == EOF)
	   	                    break;
	   	            } while (1);
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "/");
	   	        *c = sgetc(str);
	   	        if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "==");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "=");
	   	        *c = sgetc(str);
	   	        if (*c == '=')
	   	        {
	   	            strcpy(Token->Value, "!=");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "!");
	   	        strcpy(Token->Value, "%");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, ",");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, ";");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, ":");
	   	        *c          = sgetc(str);
	   	    case '(':
	   	        strcpy(Token->Value, "(");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, ")");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, "{");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, "}");
	   	        *c          = sgetc(str);
	   	        *c = sgetc(str);
	   	        if (*c == '|')
	   	        {
	   	            strcpy(Token->Value, "||");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "|");
	   	        *c = sgetc(str);
	   	        if (*c == '&')
	   	        {
	   	            strcpy(Token->Value, "&&");
	   	            *c          = sgetc(str);
	   	            strcpy(Token->Value, "&");
	   	        strcpy(Token->Value, "^");
	   	        *c          = sgetc(str);
	   	        *c = sgetc(str);
	   	        if (IsLetter(*c))
	   	        {
	   	            while (IsLetter(*c) || IsDecimal(*c))
	   	            {
	   	                Append(Token, *c);
	   	                *c = sgetc(str);
	   	            if (RegisterToInt(Token->Value) != INVALID)
	   	            {
	   	                Token->Type = REGISTER;
	   	            }
	   	            else
	   	            {
	   	                Token->Type = UNKNOWN;
	   	            }
	   	            return Token;
	   	        }
	   	    case '$':
	   	        *c = sgetc(str);
	   	        if (IsLetter(*c))
	   	        {
	   	            while (IsLetter(*c) || IsDecimal(*c))
	   	            {
	   	                Append(Token, *c);
	   	                *c = sgetc(str);
	   	            if (PseudoRegToInt(Token->Value) != INVALID)
	   	            {
	   	                Token->Type = PSEUDO_REGISTER;
	   	            }
	   	            else
	   	            {
	   	                Token->Type = UNKNOWN;
	   	            }
	   	            return Token;
	   	        }
	   	    case '.':
	   	        Append(Token, *c);
	   	        *c = sgetc(str);
	   	        if (IsLetter(*c) || IsHex(*c) || (*c == '_') || (*c == '!'))
	   	        {
	   	            do
	   	            {
	   	                Append(Token, *c);
	   	                *c = sgetc(str);
	   	            } while (IsLetter(*c) || IsHex(*c) || (*c == '_') || (*c == '!'));
	   	            BOOLEAN HasBang  = strstr(Token->Value, "!") != 0;
	   	            UINT64  Address  = 0;
	   	            if (HasBang)
	   	            {
	   	                Address = ScriptEngineConvertNameToAddress(Token->Value, &WasFound);
	   	            if (WasFound)
	   	            {
	   	                RemoveToken(&Token);
	   	                sprintf(str, "%llx", Address);
	   	                Token = NewToken(HEX, str);
	   	                if (HasBang)
	   	                {
	   	                    Token->Type = UNKNOWN;
	   	                    return Token;
	   	                }
	   	                else
	   	                {
	   	                    if (GetGlobalIdentifierVal(Token) != -1)
	   	                    {
	   	                        Token->Type = GLOBAL_ID;
	   	                    }
	   	                    else
	   	                    {
	   	                        Token->Type = GLOBAL_UNRESOLVED_ID;
	   	                    }
	   	                }
	   	            }
	   	        }
	   	        else
	   	        {
	   	            Token->Type = UNKNOWN;
	   	            return Token;
	   	        }
	   	        return Token;
	   	    case ' ':
	   	    case '\t':
	   	        strcpy(Token->Value, "");
	   	        *c          = sgetc(str);
	   	        strcpy(Token->Value, "\n");
	   	        *c          = sgetc(str);
	   	        *c = sgetc(str);
	   	        if (*c == 'x')
	   	        {
	   	            *c = sgetc(str);
	   	            while (IsHex(*c) || *c == '`')
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	        else if (*c == 'o')
	   	        {
	   	            *c = sgetc(str);
	   	            while (IsOctal(*c) || *c == '`')
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	        else if (*c == 'n')
	   	        {
	   	            *c = sgetc(str);
	   	            while (IsDecimal(*c) || *c == '`')
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	        else if (*c == 'y')
	   	        {
	   	            *c = sgetc(str);
	   	            while (IsBinary(*c) || *c == '`')
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	        else if (IsHex(*c))
	   	        {
	   	            do
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	            } while (IsHex(*c) || *c == '`');
	   	            strcpy(Token->Value, "0");
	   	        if (*c >= '0' && *c <= '9')
	   	        {
	   	            do
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	            } while (IsHex(*c) || *c == '`');
	   	        else if ((*c >= 'a' && *c <= 'f') || (*c >= 'A' && *c <= 'F') || (*c == '_') || (*c == '!'))
	   	        {
	   	            uint8_t NotHex = 0;
	   	            do
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	                if (IsHex(*c) || *c == '`')
	   	                {
	   	                }
	   	                else if ((*c >= 'G' && *c <= 'Z') || (*c >= 'g' && *c <= 'z'))
	   	                {
	   	                    NotHex = 1;
	   	                    break;
	   	                }
	   	                else
	   	                {
	   	                    break;
	   	                }
	   	            } while (1);
	   	            if (NotHex)
	   	            {
	   	                do
	   	                {
	   	                    if (*c != '`')
	   	                        Append(Token, *c);
	   	                    *c = sgetc(str);
	   	                } while (IsLetter(*c) || IsHex(*c) || (*c == '_') || (*c == '!'));
	   	                if (IsKeyword(Token->Value))
	   	                {
	   	                    Token->Type = KEYWORD;
	   	                }
	   	                else if (IsRegister(Token->Value))
	   	                {
	   	                    Token->Type = REGISTER;
	   	                }
	   	                else
	   	                {
	   	                    BOOLEAN WasFound = FALSE;
	   	                    BOOLEAN HasBang  = strstr(Token->Value, "!") != 0;
	   	                    UINT64  Address  = 0;
	   	                    if (HasBang)
	   	                    {
	   	                        Address = ScriptEngineConvertNameToAddress(Token->Value, &WasFound);
	   	                    if (WasFound)
	   	                    {
	   	                        RemoveToken(&Token);
	   	                        sprintf(str, "%llx", Address);
	   	                        Token = NewToken(HEX, str);
	   	                        if (HasBang)
	   	                        {
	   	                            Token->Type = UNKNOWN;
	   	                            return Token;
	   	                        }
	   	                        else
	   	                        {
	   	                            if (GetLocalIdentifierVal(Token) != -1)
	   	                            {
	   	                                Token->Type = LOCAL_ID;
	   	                            }
	   	                            else
	   	                            {
	   	                                Token->Type = LOCAL_UNRESOLVED_ID;
	   	                            }
	   	                        }
	   	                    }
	   	                }
	   	                return Token;
	   	            }
	   	            else
	   	            {
	   	                if (IsKeyword(Token->Value))
	   	                {
	   	                    Token->Type = KEYWORD;
	   	                }
	   	                else if (IsRegister(Token->Value))
	   	                {
	   	                    Token->Type = REGISTER;
	   	                }
	   	                else if (IsId(Token->Value))
	   	                {
	   	                    BOOLEAN WasFound = FALSE;
	   	                    BOOLEAN HasBang  = strstr(Token->Value, "!") != 0;
	   	                    UINT64  Address  = 0;
	   	                    if (HasBang)
	   	                    {
	   	                        Address = ScriptEngineConvertNameToAddress(Token->Value, &WasFound);
	   	                    if (WasFound)
	   	                    {
	   	                        RemoveToken(&Token);
	   	                        sprintf(str, "%llx", Address);
	   	                        Token = NewToken(HEX, str);
	   	                        if (HasBang)
	   	                        {
	   	                            Token->Type = UNKNOWN;
	   	                            return Token;
	   	                        }
	   	                        else
	   	                        {
	   	                            if (GetLocalIdentifierVal(Token) != -1)
	   	                            {
	   	                                Token->Type = LOCAL_ID;
	   	                            }
	   	                            else
	   	                            {
	   	                                Token->Type = LOCAL_UNRESOLVED_ID;
	   	                            }
	   	                        }
	   	                    }
	   	                }
	   	                else
	   	                {
	   	                    Token->Type = HEX;
	   	                }
	   	                return Token;
	   	            }
	   	        }
	   	        else if ((*c >= 'G' && *c <= 'Z') || (*c >= 'g' && *c <= 'z') || (*c == '_') || (*c == '!'))
	   	        {
	   	            do
	   	            {
	   	                if (*c != '`')
	   	                    Append(Token, *c);
	   	                *c = sgetc(str);
	   	            } while (IsLetter(*c) || IsHex(*c) || (*c == '_') || (*c == '!'));
	   	            if (IsKeyword(Token->Value))
	   	            {
	   	                Token->Type = KEYWORD;
	   	            }
	   	            else if (IsRegister(Token->Value))
	   	            {
	   	                Token->Type = REGISTER;
	   	            }
	   	            else
	   	            {
	   	                BOOLEAN WasFound = FALSE;
	   	                BOOLEAN HasBang  = strstr(Token->Value, "!") != 0;
	   	                UINT64  Address  = 0;
	   	                if (HasBang)
	   	                {
	   	                    Address = ScriptEngineConvertNameToAddress(Token->Value, &WasFound);
	   	                if (WasFound)
	   	                {
	   	                    RemoveToken(&Token);
	   	                    sprintf(str, "%llx", Address);
	   	                    Token = NewToken(HEX, str);
	   	                    if (HasBang)
	   	                    {
	   	                        Token->Type = UNKNOWN;
	   	                        return Token;
	   	                    }
	   	                    else
	   	                    {
	   	                        if (GetLocalIdentifierVal(Token) != -1)
	   	                        {
	   	                            Token->Type = LOCAL_ID;
	   	                        }
	   	                        else
	   	                        {
	   	                            Token->Type = LOCAL_UNRESOLVED_ID;
	   	                        }
	   	                    }
	   	                }
	   	            }
	   	            return Token;
	   	        }
	   	        Token->Type = UNKNOWN;
	   	        *c          = sgetc(str);

	   Scan(char * str, char * c)

	   	{
	   	    static BOOLEAN ReturnEndOfString;
	   	    PTOKEN         Token;
	   	    if (InputIdx <= 1)
	   	    {
	   	        ReturnEndOfString = FALSE;
	   	    }
	   	    if (ReturnEndOfString)
	   	    {
	   	        Token = NewToken(END_OF_STACK, "$");
	   	    if (str[InputIdx - 1] == '\0')
	   	    {
	   	    }
	   	    while (1)
	   	    {
	   	        CurrentTokenIdx = InputIdx - 1;
	   	        Token = GetToken(c, str);
	   	        if ((char)*c == EOF)
	   	        {
	   	            ReturnEndOfString = TRUE;
	   	        }
	   	        if (Token->Type == WHITE_SPACE)
	   	        {
	   	            if (!strcmp(Token->Value, "\n"))
	   	            {
	   	                CurrentLine++;
	   	                CurrentLineIdx = InputIdx;
	   	            }
	   	            RemoveToken(&Token);
	   	            if (ReturnEndOfString)
	   	            {
	   	                Token = NewToken(END_OF_STACK, "$");
	   	        else if (Token->Type == COMMENT)
	   	        {
	   	            RemoveToken(&Token);
	   	            if (ReturnEndOfString)
	   	            {
	   	                Token = NewToken(END_OF_STACK, "$");

	   sgetc(char * str)

	   	{
	   	    char c = str[InputIdx];
	   	    if (c)
	   	    {
	   	        InputIdx++;
	   	        return c;
	   	    }
	   	    else
	   	    {
	   	        return EOF;
	   	    }
	   	}
	*/
	return true
}

func (s *scanner) IsKeyword() (ok bool) { //col:530
	/*
	   IsKeyword(char * str)

	   	{
	   	    int n = KEYWORD_LIST_LENGTH;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(str, KeywordList[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    n = TERMINAL_COUNT;
	   	    for (int i = 0; i < n; i++)
	   	    {
	   	        if (!strcmp(str, TerminalMap[i]))
	   	        {
	   	            return 1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (s *scanner) IsRegister() (ok bool) { //col:536
	/*
	   IsRegister(char * str)

	   	{
	   	    if (RegisterToInt(str) == INVALID)
	   	        return 0;
	   	    return 1;
	   	}
	*/
	return true
}

func (s *scanner) IsId() (ok bool) { //col:540
	/*
	   IsId(char * str)

	   	{
	   	    return 0;
	   	}
	*/
	return true
}

