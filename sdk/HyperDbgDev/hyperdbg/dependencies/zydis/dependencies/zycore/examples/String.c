/***************************************************************************************************

  Zyan Core Library (Zycore-C)

  Original Author : Florian Bernd

 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.

***************************************************************************************************/

/**
 * @file
 * @brief   Demonstrates the `String` implementation.
 */

#include <Zycore/Allocator.h>
#include <Zycore/Defines.h>
#include <Zycore/LibC.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#include <stdio.h>

/* ==============================================================================================
 */
/* Enums and types */
/* ==============================================================================================
 */

/* ==============================================================================================
 */
/* Helper functions */
/* ==============================================================================================
 */

/* ==============================================================================================
 */
/* Tests */
/* ==============================================================================================
 */

/* ----------------------------------------------------------------------------------------------
 */
/* Basic tests */
/* ----------------------------------------------------------------------------------------------
 */

/**
 * @brief   Performs some basic test on the given `ZyanString` instance.
 *
 * @param   string  A pointer to the `ZyanString` instance.
 *
 * @return  A zyan status code.
 */
static ZyanStatus PerformBasicTests(ZyanString *string) {
  ZYAN_ASSERT(string);
  ZYAN_UNUSED(string);

  return ZYAN_STATUS_SUCCESS;
}

/**
 * @brief   Performs basic tests on a string that dynamically manages memory.
 *
 * @return  A zyan status code.
 */
static ZyanStatus TestDynamic(void) {
  PerformBasicTests(ZYAN_NULL);
  return ZYAN_STATUS_SUCCESS;
}

/**
 * @brief   Performs basic tests on a string that uses a static buffer.
 *
 * @return  A zyan status code.
 */
static ZyanStatus TestStatic(void) {
  PerformBasicTests(ZYAN_NULL);
  return ZYAN_STATUS_SUCCESS;
}

/* ----------------------------------------------------------------------------------------------
 */
/* Custom allocator */
/* ----------------------------------------------------------------------------------------------
 */

// static ZyanStatus AllocatorAllocate(ZyanAllocator* allocator, void** p,
// ZyanUSize element_size,
//     ZyanUSize n)
//{
//     ZYAN_ASSERT(allocator);
//     ZYAN_ASSERT(p);
//     ZYAN_ASSERT(element_size);
//     ZYAN_ASSERT(n);
//
//     ZYAN_UNUSED(allocator);
//
//     *p = ZYAN_MALLOC(element_size * n);
//     if (!*p)
//     {
//         return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
//     }
//
//     return ZYAN_STATUS_SUCCESS;
// }
//
// static ZyanStatus AllocatorReallocate(ZyanAllocator* allocator, void** p,
// ZyanUSize element_size,
//     ZyanUSize n)
//{
//     ZYAN_ASSERT(allocator);
//     ZYAN_ASSERT(p);
//     ZYAN_ASSERT(element_size);
//     ZYAN_ASSERT(n);
//
//     ZYAN_UNUSED(allocator);
//
//     void* const x = ZYAN_REALLOC(*p, element_size * n);
//     if (!x)
//     {
//         return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
//     }
//     *p = x;
//
//     return ZYAN_STATUS_SUCCESS;
// }
//
// static ZyanStatus AllocatorDeallocate(ZyanAllocator* allocator, void* p,
// ZyanUSize element_size,
//     ZyanUSize n)
//{
//     ZYAN_ASSERT(allocator);
//     ZYAN_ASSERT(p);
//     ZYAN_ASSERT(element_size);
//     ZYAN_ASSERT(n);
//
//     ZYAN_UNUSED(allocator);
//     ZYAN_UNUSED(element_size);
//     ZYAN_UNUSED(n);
//
//     ZYAN_FREE(p);
//
//     return ZYAN_STATUS_SUCCESS;
// }

/* ----------------------------------------------------------------------------------------------
 */

/**
 * @brief   Performs basic tests on a vector that dynamically manages memory
 * using a custom allocator and modified growth-factor/shrink-threshold.
 *
 * @return  A zyan status code.
 */
static ZyanStatus TestAllocator(void) { return ZYAN_STATUS_SUCCESS; }

/* ----------------------------------------------------------------------------------------------
 */

/* ==============================================================================================
 */
/* Entry point */
/* ==============================================================================================
 */

int main() {
  if (!ZYAN_SUCCESS(TestDynamic())) {
    return EXIT_FAILURE;
  }
  if (!ZYAN_SUCCESS(TestStatic())) {
    return EXIT_FAILURE;
  }
  if (!ZYAN_SUCCESS(TestAllocator())) {
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}

/* ==============================================================================================
 */

// ===========================
// TranslationUnit
//  `+Files = 
//   `-File
//    |+Name = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c
//    |+Decl = 
//    | |-VarDecl
//    | | |+Type = <nil>
//    | | |+Name = ZyanStatus
//    | | `+Init = <nil>
//    | |-VarDecl
//    | | |+Type = <nil>
//    | | |+Name = ZyanStatus
//    | | `+Init = <nil>
//    | |-VarDecl
//    | | |+Type =  void
//    | | |+Name = <nil>
//    | | `+Init = <nil>
//    | |-VarDecl
//    | | |+Type = <nil>
//    | | |+Name = ZyanStatus
//    | | `+Init = <nil>
//    | |-VarDecl
//    | | |+Type =  void
//    | | |+Name = <nil>
//    | | `+Init = <nil>
//    | |-VarDecl
//    | | |+Type = <nil>
//    | | |+Name = ZyanStatus
//    | | `+Init = <nil>
//    | |-VarDecl
//    | | |+Type =  void
//    | | |+Name = <nil>
//    | | `+Init = <nil>
//    | `-FuncDecl
//    |  |+Name = main
//    |  |+Type =  int ()
//    |  |+Decl = 
//    |  `+Body = CompoundStmt
//    |   |+Lbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:183:12
//    |   |+Stmts = 
//    |   | |-IfStmt
//    |   | | |+If = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:3
//    |   | | |+X = UnaryExpr
//    |   | | | |+Op = "!"<PUNCTUATOR@D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:7>
//    |   | | | `+X = CallExpr
//    |   | | |  |+Func = ZYAN_SUCCESS
//    |   | | |  |+Lparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:20
//    |   | | |  |+Args = 
//    |   | | |  | `-CallExpr
//    |   | | |  |  |+Func = TestDynamic
//    |   | | |  |  |+Lparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:32
//    |   | | |  |  |+Args = 
//    |   | | |  |  `+Rparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:33
//    |   | | |  `+Rparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:34
//    |   | | |+Then = CompoundStmt
//    |   | | | |+Lbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:37
//    |   | | | |+Stmts = 
//    |   | | | | `-ReturnStmt
//    |   | | | |  |+Return = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:185:5
//    |   | | | |  |+X = EXIT_FAILURE
//    |   | | | |  `+Semicolon = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:185:24
//    |   | | | `+Rbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:186:3
//    |   | | `+Else = <nil>
//    |   | |-IfStmt
//    |   | | |+If = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:3
//    |   | | |+X = UnaryExpr
//    |   | | | |+Op = "!"<PUNCTUATOR@D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:7>
//    |   | | | `+X = CallExpr
//    |   | | |  |+Func = ZYAN_SUCCESS
//    |   | | |  |+Lparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:20
//    |   | | |  |+Args = 
//    |   | | |  | `-CallExpr
//    |   | | |  |  |+Func = TestStatic
//    |   | | |  |  |+Lparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:31
//    |   | | |  |  |+Args = 
//    |   | | |  |  `+Rparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:32
//    |   | | |  `+Rparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:33
//    |   | | |+Then = CompoundStmt
//    |   | | | |+Lbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:36
//    |   | | | |+Stmts = 
//    |   | | | | `-ReturnStmt
//    |   | | | |  |+Return = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:188:5
//    |   | | | |  |+X = EXIT_FAILURE
//    |   | | | |  `+Semicolon = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:188:24
//    |   | | | `+Rbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:189:3
//    |   | | `+Else = <nil>
//    |   | |-IfStmt
//    |   | | |+If = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:3
//    |   | | |+X = UnaryExpr
//    |   | | | |+Op = "!"<PUNCTUATOR@D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:7>
//    |   | | | `+X = CallExpr
//    |   | | |  |+Func = ZYAN_SUCCESS
//    |   | | |  |+Lparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:20
//    |   | | |  |+Args = 
//    |   | | |  | `-CallExpr
//    |   | | |  |  |+Func = TestAllocator
//    |   | | |  |  |+Lparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:34
//    |   | | |  |  |+Args = 
//    |   | | |  |  `+Rparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:35
//    |   | | |  `+Rparen = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:36
//    |   | | |+Then = CompoundStmt
//    |   | | | |+Lbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:39
//    |   | | | |+Stmts = 
//    |   | | | | `-ReturnStmt
//    |   | | | |  |+Return = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:191:5
//    |   | | | |  |+X = EXIT_FAILURE
//    |   | | | |  `+Semicolon = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:191:24
//    |   | | | `+Rbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:192:3
//    |   | | `+Else = <nil>
//    |   | `-ReturnStmt
//    |   |  |+Return = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:194:3
//    |   |  |+X = EXIT_SUCCESS
//    |   |  `+Semicolon = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:194:22
//    |   `+Rbrace = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:195:1
//    `+Unresolved = 
// ===========================
//
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行19列: 这里应该是一个 ; ，不应该出现 PerformBasicTests
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行19列: 这里应该是一个 定义语句 ，不应该出现 PerformBasicTests
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:36
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行36列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:37
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行37列: 这里应该是一个 定义语句 ，不应该出现 ZyanString
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:48
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行48列: 这里应该是一个 定义语句 ，不应该出现 *
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:49
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行49列: 这里应该是一个 定义语句 ，不应该出现 string
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:55
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行55列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:57
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第71行57列: 这里应该是一个 定义语句 ，不应该出现 {
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:72:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第72行3列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_ASSERT
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:72:14
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第72行14列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:72:15
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第72行15列: 这里应该是一个 定义语句 ，不应该出现 string
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:72:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第72行21列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:72:22
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第72行22列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:73:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第73行3列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_UNUSED
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:73:14
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第73行14列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:73:15
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第73行15列: 这里应该是一个 定义语句 ，不应该出现 string
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:73:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第73行21列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:73:22
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第73行22列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:75:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第75行3列: 这里应该是一个 定义语句 ，不应该出现 return
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:75:10
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第75行10列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_STATUS_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:75:29
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第75行29列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:76:1
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第76行1列: 这里应该是一个 定义语句 ，不应该出现 }
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行8列: 重复声明的变量名 ZyanStatus，上次声明的位置 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:8
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行19列: 这里应该是一个 ; ，不应该出现 TestDynamic
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行19列: 这里应该是一个 定义语句 ，不应该出现 TestDynamic
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:30
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行30列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:35
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行35列: 这里应该是一个 ; ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:35
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行35列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:83:37
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第83行37列: 这里应该是一个 定义语句 ，不应该出现 {
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:84:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第84行3列: 这里应该是一个 定义语句 ，不应该出现 PerformBasicTests
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:84:20
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第84行20列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:84:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第84行21列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_NULL
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:84:30
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第84行30列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:84:31
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第84行31列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:85:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第85行3列: 这里应该是一个 定义语句 ，不应该出现 return
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:85:10
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第85行10列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_STATUS_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:85:29
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第85行29列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:86:1
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第86行1列: 这里应该是一个 定义语句 ，不应该出现 }
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行8列: 重复声明的变量名 ZyanStatus，上次声明的位置 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:8
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行19列: 这里应该是一个 ; ，不应该出现 TestStatic
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行19列: 这里应该是一个 定义语句 ，不应该出现 TestStatic
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:29
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行29列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:34
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行34列: 这里应该是一个 ; ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:34
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行34列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:93:36
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第93行36列: 这里应该是一个 定义语句 ，不应该出现 {
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:94:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第94行3列: 这里应该是一个 定义语句 ，不应该出现 PerformBasicTests
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:94:20
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第94行20列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:94:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第94行21列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_NULL
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:94:30
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第94行30列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:94:31
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第94行31列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:95:3
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第95行3列: 这里应该是一个 定义语句 ，不应该出现 return
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:95:10
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第95行10列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_STATUS_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:95:29
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第95行29列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:96:1
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第96行1列: 这里应该是一个 定义语句 ，不应该出现 }
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行8列: 重复声明的变量名 ZyanStatus，上次声明的位置 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:71:8
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行19列: 这里应该是一个 ; ，不应该出现 TestAllocator
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:19
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行19列: 这里应该是一个 定义语句 ，不应该出现 TestAllocator
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:32
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行32列: 这里应该是一个 定义语句 ，不应该出现 (
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:37
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行37列: 这里应该是一个 ; ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:37
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行37列: 这里应该是一个 定义语句 ，不应该出现 )
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:39
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行39列: 这里应该是一个 定义语句 ，不应该出现 {
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:41
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行41列: 这里应该是一个 定义语句 ，不应该出现 return
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:48
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行48列: 这里应该是一个 定义语句 ，不应该出现 ZYAN_STATUS_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:67
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行67列: 这里应该是一个 定义语句 ，不应该出现 ;
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:172:69
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第172行69列: 这里应该是一个 定义语句 ，不应该出现 }
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第184行8列: 未定义的标识符 ZYAN_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:184:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第184行21列: 未定义的标识符 TestDynamic
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:185:12
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第185行12列: 未定义的标识符 EXIT_FAILURE
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第187行8列: 未定义的标识符 ZYAN_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:187:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第187行21列: 未定义的标识符 TestStatic
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:188:12
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第188行12列: 未定义的标识符 EXIT_FAILURE
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:8
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第190行8列: 未定义的标识符 ZYAN_SUCCESS
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:190:21
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第190行21列: 未定义的标识符 TestAllocator
// |-Error
// | |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:191:12
// | |+Typ = 0
// | `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第191行12列: 未定义的标识符 EXIT_FAILURE
// `-Error
//  |+Pos = D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c:194:10
//  |+Typ = 0
//  `+Msg = 在 D:\codespace\workspace\src\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c 文件的第194行10列: 未定义的标识符 EXIT_SUCCESS
// ===========================
