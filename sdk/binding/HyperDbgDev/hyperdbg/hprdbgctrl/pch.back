










#pragma once









#define WIN32_LEAN_AND_MEAN





#define USE_LIB_IA32
#if defined(USE_LIB_IA32)
#    pragma warning(push, 0)

#    include <ia32-doc/out/ia32.h>
#    pragma warning(pop)
typedef RFLAGS * PRFLAGS;
#endif 




#define USE__NATIVE_PHNT_HEADERS

#if defined(USE__NATIVE_PHNT_HEADERS)




typedef const wchar_t *LPCWCHAR, *PCWCHAR;

#    define PHNT_MODE    PHNT_MODE_USER
#    define PHNT_VERSION PHNT_WIN11 

#    define PHNT_PATCH_FOR_HYPERDBG TRUE

#    include <phnt/phnt_windows.h>
#    include <phnt/phnt.h>

#elif defined(USE_NATIVE_SDK_HEADERS)

#    include <winternl.h>
#    include <Windows.h>
#    include <winioctl.h>

#endif

#include <winsock2.h>
#include <ws2tcpip.h>
#include <strsafe.h>
#include <shlobj.h>
#include <tchar.h>
#include <tlhelp32.h>
#include <shlwapi.h>
#include <VersionHelpers.h>
#include <tchar.h>
#include <psapi.h>
#include <time.h>

#include <conio.h>
#include <intrin.h>
#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>




#include <algorithm>
#include <string>
#include <vector>
#include <array>
#include <bitset>
#include <iomanip>
#include <iostream>
#include <iterator>
#include <sstream>
#include <fstream>
#include <map>
#include <numeric>
#include <list>
#include <locale>
#include <memory>
#include <cctype>
#include <cstring>




#define SCRIPT_ENGINE_USER_MODE
#define HYPERDBG_USER_MODE




#include "SDK/HyperDbgSdk.h"
#include "SDK/Imports/HyperDbgCtrlImports.h"
#include "Configuration.h"
#include "Definition.h"




#include "..\script-eval\header\ScriptEngineCommonDefinitions.h"
#include "..\script-eval\header\ScriptEngineHeader.h"




#include "SDK/Imports/HyperDbgScriptImports.h"




#include "header/inipp.h"
#include "header/commands.h"
#include "header/common.h"
#include "header/symbol.h"
#include "header/debugger.h"
#include "header/script-engine.h"
#include "header/exports.h"
#include "header/help.h"
#include "header/install.h"
#include "header/list.h"
#include "header/tests.h"
#include "header/transparency.h"
#include "header/communication.h"
#include "header/namedpipe.h"
#include "header/forwarding.h"
#include "header/kd.h"
#include "header/pe-parser.h"
#include "header/ud.h"
#include "header/objects.h"

#pragma comment(lib, "ntdll.lib")




#pragma comment(lib, "Shlwapi.lib")





#pragma comment(lib, "Ws2_32.lib")
#pragma comment(lib, "Mswsock.lib")
#pragma comment(lib, "AdvApi32.lib")






#pragma comment(lib, "Psapi.lib")
#pragma comment(lib, "Kernel32.lib")
