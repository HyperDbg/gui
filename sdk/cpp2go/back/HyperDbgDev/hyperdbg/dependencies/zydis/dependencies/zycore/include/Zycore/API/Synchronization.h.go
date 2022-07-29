package API
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\API\Synchronization.h.back

const(
ZYCORE_SYNCHRONIZATION_H =  //col:33
)

type (
Synchronization interface{
  Zyan Core Library ()(ok bool)//col:128
}
)

func NewSynchronization() { return & synchronization{} }

func (s *synchronization)  Zyan Core Library ()(ok bool){//col:128
/*  Zyan Core Library (Zycore-C)
  Original Author : Florian Bernd
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
#ifndef ZYCORE_SYNCHRONIZATION_H
#define ZYCORE_SYNCHRONIZATION_H
#ifndef ZYAN_NO_LIBC
#include <ZycoreExportConfig.h>
#include <Zycore/Defines.h>
#include <Zycore/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
#if   defined(ZYAN_POSIX)
#include <pthread.h>
typedef pthread_mutex_t ZyanCriticalSection;
#elif defined(ZYAN_WINDOWS)
#include <windows.h>
typedef CRITICAL_SECTION ZyanCriticalSection;
#else
#   error "Unsupported platform detected"
#endif
 *
ZYCORE_EXPORT ZyanStatus ZyanCriticalSectionInitialize(ZyanCriticalSection* critical_section);
 *
ZYCORE_EXPORT ZyanStatus ZyanCriticalSectionEnter(ZyanCriticalSection* critical_section);
 *
 *
 *          if not.
ZYCORE_EXPORT ZyanBool ZyanCriticalSectionTryEnter(ZyanCriticalSection* critical_section);
 *
ZYCORE_EXPORT ZyanStatus ZyanCriticalSectionLeave(ZyanCriticalSection* critical_section);
 *
ZYCORE_EXPORT ZyanStatus ZyanCriticalSectionDelete(ZyanCriticalSection* critical_section);
#ifdef __cplusplus
}*/
return true
}



