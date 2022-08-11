#pragma once
#include <Windows.h>
#include <iomanip>
#include <sstream>
#include <string>
#include <vector>
#define _NO_CVCONST_H 
#include <DbgHelp.h>
#define USE_LIB_IA32
#if defined(USE_LIB_IA32)
#pragma warning(push, 0)
#include <ia32-doc/out/ia32.h>
#pragma warning(pop)
typedef RFLAGS *PRFLAGS;
#endif 
#include "..\symbol-parser\header\common-utils.h"
#include "..\symbol-parser\header\symbol-parser.h"
#include "Definition.h"
#include "SDK/HyperDbgSdk.h"
#include "SDK/Imports/HyperDbgCtrlImports.h"
using namespace std;
#pragma comment(lib, "dbghelp.lib")
#pragma comment(lib, "Urlmon.lib")
