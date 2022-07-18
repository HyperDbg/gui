#pragma once
#ifndef SCANNER_H
#    define SCANNER_H
PTOKEN_LIST  IdTable;
unsigned int InputIdx;
unsigned int CurrentLine;
unsigned int CurrentLineIdx;
unsigned int CurrentTokenIdx;
PTOKEN
GetToken(char * c, char * str);
PTOKEN
Scan(char * str, char * c);
char
sgetc(char * str);
char
IsKeyword(char * str);
char
IsId(char * str);
char
IsRegister(char * str);
#endif // !SCANNER_H
