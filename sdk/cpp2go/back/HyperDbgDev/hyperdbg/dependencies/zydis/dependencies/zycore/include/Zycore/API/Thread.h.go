package API
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\API\Thread.h.back

const(
ZYCORE_THREAD_H =  //col:33
ZYAN_THREAD_DECLARE_TLS_CALLBACK(name, param_type, param_name) = void name(param_type* param_name) //col:88
ZYAN_THREAD_DECLARE_TLS_CALLBACK(name, param_type, param_name) = VOID NTAPI name(param_type* param_name) //col:132
)

type (
Thread interface{
  Zyan Core Library ()(ok bool)//col:239
}
)

func NewThread() { return & thread{} }

func (t *thread)  Zyan Core Library ()(ok bool){//col:239
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
#ifndef ZYCORE_THREAD_H
#define ZYCORE_THREAD_H
#ifndef ZYAN_NO_LIBC
#include <ZycoreExportConfig.h>
#include <Zycore/Defines.h>
#include <Zycore/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
#if   defined(ZYAN_POSIX)
#include <pthread.h>
typedef pthread_t ZyanThread;
typedef ZyanU64 ZyanThreadId;
typedef pthread_key_t ZyanThreadTlsIndex;
typedef void(*ZyanThreadTlsCallback)(void* data);
 *
#define ZYAN_THREAD_DECLARE_TLS_CALLBACK(name, param_type, param_name) \
    void name(param_type* param_name)
#elif defined(ZYAN_WINDOWS)
#include <windows.h>
typedef HANDLE ZyanThread;
typedef DWORD ZyanThreadId;
typedef DWORD ZyanThreadTlsIndex;
typedef PFLS_CALLBACK_FUNCTION ZyanThreadTlsCallback;
 *
#define ZYAN_THREAD_DECLARE_TLS_CALLBACK(name, param_type, param_name) \
    VOID NTAPI name(param_type* param_name)
#else
#   error "Unsupported platform detected"
#endif
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id);
 *
 *                      in the TLS slot or `ZYAN_NULL`, if not needed.
 *
 * The maximum available number of TLS slots is implementation specific and different on each
 * platform:
 * - Windows
 *   - A total amount of 128 slots per process are guaranteed
 * - POSIX
 *   - A total amount of 128 slots per process are guaranteed
 *   - Some systems guarantee larger amounts like e.g. 1024 slots per process
 *
 * Note that the invocation rules for the destructor callback are implementation specific and
 * different on each platform:
 * - Windows
 *   - The callback is invoked when a thread exits
 *   - The callback is invoked when the process exits
 *   - The callback is invoked when the TLS slot is released
 * - POSIX
 *   - The callback is invoked when a thread exits and the stored value is not null
 *   - The callback is NOT invoked when the process exits
 *   - The callback is NOT invoked when the TLS slot is released
 *
ZYCORE_EXPORT ZyanStatus ZyanThreadTlsAlloc(ZyanThreadTlsIndex* index,
    ZyanThreadTlsCallback destructor);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index);
 *          thread.
 *
 *                  calling thread.
 *
ZYCORE_EXPORT ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data);
 *
 *                  calling thread
 *
ZYCORE_EXPORT ZyanStatus ZyanThreadTlsSetValue(ZyanThreadTlsIndex index, void* data);
#ifdef __cplusplus
}*/
return true
}



