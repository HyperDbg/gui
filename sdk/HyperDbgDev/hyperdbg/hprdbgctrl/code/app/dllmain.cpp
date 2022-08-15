/**
 * @file dllmain.cpp
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Defines the entry point for the DLL application
 * @details
 * @version 0.1
 * @date 2020-04-11
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#include "pch.h"

/**
 * @brief Dll Main Entry
 *
 * @param hModule
 * @param ul_reason_for_call
 * @param lpReserved
 * @return BOOL DllMain
 */
BOOL APIENTRY DllMain(HMODULE hModule, DWORD ul_reason_for_call,
                      LPVOID lpReserved) {
  switch (ul_reason_for_call) {
  case DLL_PROCESS_ATTACH:
  case DLL_THREAD_ATTACH:
  case DLL_THREAD_DETACH:
  case DLL_PROCESS_DETACH:
    break;
  }
  return TRUE;
}

// ===========================
// TranslationUnit
//  `+Files = 
//   `-File
//    |+Name = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp
//    |+Decl = 
//    `+Unresolved = 
// ===========================
//
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:1
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行1列: 这里应该是一个 定义语句 ，不应该出现 BOOL
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:6
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行6列: 这里应该是一个 定义语句 ，不应该出现 APIENTRY
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:15
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行15列: 这里应该是一个 定义语句 ，不应该出现 DllMain
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:22
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行22列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:23
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行23列: 这里应该是一个 定义语句 ，不应该出现 HMODULE
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:31
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行31列: 这里应该是一个 定义语句 ，不应该出现 hModule
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:38
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行38列: 这里应该是一个 定义语句 ，不应该出现 ,
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:40
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行40列: 这里应该是一个 定义语句 ，不应该出现 DWORD
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:46
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行46列: 这里应该是一个 定义语句 ，不应该出现 ul_reason_for_call
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:22:64
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第22行64列: 这里应该是一个 定义语句 ，不应该出现 ,
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:23:23
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第23行23列: 这里应该是一个 定义语句 ，不应该出现 LPVOID
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:23:30
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第23行30列: 这里应该是一个 定义语句 ，不应该出现 lpReserved
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:23:40
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第23行40列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:23:42
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第23行42列: 这里应该是一个 定义语句 ，不应该出现 {
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:24:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第24行3列: 这里应该是一个 定义语句 ，不应该出现 switch
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:24:10
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第24行10列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:24:11
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第24行11列: 这里应该是一个 定义语句 ，不应该出现 ul_reason_for_call
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:24:29
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第24行29列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:24:31
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第24行31列: 这里应该是一个 定义语句 ，不应该出现 {
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:25:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第25行3列: 这里应该是一个 定义语句 ，不应该出现 case
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:25:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第25行8列: 这里应该是一个 定义语句 ，不应该出现 DLL_PROCESS_ATTACH
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:25:26
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第25行26列: 这里应该是一个 定义语句 ，不应该出现 :
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:26:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第26行3列: 这里应该是一个 定义语句 ，不应该出现 case
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:26:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第26行8列: 这里应该是一个 定义语句 ，不应该出现 DLL_THREAD_ATTACH
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:26:25
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第26行25列: 这里应该是一个 定义语句 ，不应该出现 :
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:27:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第27行3列: 这里应该是一个 定义语句 ，不应该出现 case
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:27:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第27行8列: 这里应该是一个 定义语句 ，不应该出现 DLL_THREAD_DETACH
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:27:25
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第27行25列: 这里应该是一个 定义语句 ，不应该出现 :
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:28:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第28行3列: 这里应该是一个 定义语句 ，不应该出现 case
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:28:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第28行8列: 这里应该是一个 定义语句 ，不应该出现 DLL_PROCESS_DETACH
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:28:26
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第28行26列: 这里应该是一个 定义语句 ，不应该出现 :
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:29:5
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第29行5列: 这里应该是一个 定义语句 ，不应该出现 break
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:29:10
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第29行10列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:30:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第30行3列: 这里应该是一个 定义语句 ，不应该出现 }
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:31:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第31行3列: 这里应该是一个 定义语句 ，不应该出现 return
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:31:10
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第31行10列: 这里应该是一个 定义语句 ，不应该出现 TRUE
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:31:14
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第31行14列: 这里应该是一个 定义语句 ，不应该出现 ;
// `-Error
//  |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp:32:1
//  |+Typ = 0
//  `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp 文件的第32行1列: 这里应该是一个 定义语句 ，不应该出现 }
// ===========================
