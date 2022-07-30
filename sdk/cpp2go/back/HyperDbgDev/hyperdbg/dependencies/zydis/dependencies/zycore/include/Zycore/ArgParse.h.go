package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\ArgParse.h.back

const(
ZYCORE_ARGPARSE_H =  //col:33
)

type (
ArgParse interface{
  Zyan Core Library ()(ok bool)//col:68
ZYCORE_EXPORT ZyanStatus ZyanArgParse()(ok bool)//col:170
}
)

func NewArgParse() { return & argParse{} }

func (a *argParse)  Zyan Core Library ()(ok bool){//col:68
/*  Zyan Core Library (Zycore-C)
  Original Author : Joel Hoener
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#ifndef ZYCORE_ARGPARSE_H
#define ZYCORE_ARGPARSE_H
#include <Zycore/Types.h>
#include <Zycore/Status.h>
#include <Zycore/Vector.h>
#include <Zycore/String.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef struct ZyanArgParseDefinition_
{
     *
     * Must start with either one or two dashes. Single dash arguments must consist of a single
     * character, (e.g. `-n`), double-dash arguments can be of arbitrary length.
    const char* name;
    ZyanBool boolean;
    ZyanBool required;
} ZyanArgParseDefinition;*/
return true
}

func (a *argParse)ZYCORE_EXPORT ZyanStatus ZyanArgParse()(ok bool){//col:170
/*ZYCORE_EXPORT ZyanStatus ZyanArgParse(const ZyanArgParseConfig *cfg, ZyanVector* parsed,
    const char** error_token);
#endif
 *
 * This version allows specification of a custom memory allocator and thus supports no-libc.
 *
 *                      transferred to the user. Input is expected to be uninitialized. On error,
 *                      the vector remains uninitialized.
 *                      error. Optional, may be `ZYAN_NULL`. The pointer borrows into the `cfg`
 *                      struct and doesn't have to be freed by the user.
 *
ZYCORE_EXPORT ZyanStatus ZyanArgParseEx(const ZyanArgParseConfig *cfg, ZyanVector* parsed,
    const char** error_token, ZyanAllocator* allocator);
#ifdef __cplusplus
}*/
return true
}



