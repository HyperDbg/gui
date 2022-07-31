package API
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Thread.c.back

type (
Thread interface{
  Zyan Core Library ()(ok bool)//col:58
ZYAN_STATIC_ASSERT()(ok bool)//col:69
ZyanStatus ZyanThreadTlsAlloc()(ok bool)//col:94
ZyanStatus ZyanThreadTlsFree()(ok bool)//col:99
ZyanStatus ZyanThreadTlsGetValue()(ok bool)//col:106
ZyanStatus ZyanThreadTlsSetValue()(ok bool)//col:121
#elif defined()(ok bool)//col:136
ZyanStatus ZyanThreadGetCurrentThreadId()(ok bool)//col:143
ZyanStatus ZyanThreadTlsAlloc()(ok bool)//col:159
ZyanStatus ZyanThreadTlsFree()(ok bool)//col:164
ZyanStatus ZyanThreadTlsGetValue()(ok bool)//col:171
ZyanStatus ZyanThreadTlsSetValue()(ok bool)//col:186
}
)

func NewThread() { return & thread{} }

func (t *thread)  Zyan Core Library ()(ok bool){//col:58
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
#include <Zycore/API/Thread.h>
#if   defined(ZYAN_POSIX)
#include <errno.h>
ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)
{
    *thread = pthread_self();
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZYAN_STATIC_ASSERT()(ok bool){//col:69
/*ZYAN_STATIC_ASSERT(sizeof(ZyanThreadId) <= sizeof(ZyanU64));
ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)
{
    pthread_t ptid = pthread_self();
    *thread_id = *(ZyanThreadId*)ptid;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsAlloc()(ok bool){//col:94
/*ZyanStatus ZyanThreadTlsAlloc(ZyanThreadTlsIndex* index, ZyanThreadTlsCallback destructor)
{
    ZyanThreadTlsIndex value;
    const int error = pthread_key_create(&value, destructor);
    if (error != 0)
    {
        if (error == EAGAIN)
        {
            return ZYAN_STATUS_OUT_OF_RESOURCES;
        }
        if (error == ENOMEM)
        {
            return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
        }
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
    *index = value;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsFree()(ok bool){//col:99
/*ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)
{
    return !pthread_key_delete(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsGetValue()(ok bool){//col:106
/*ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)
{
    *data = pthread_getspecific(index);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsSetValue()(ok bool){//col:121
/*ZyanStatus ZyanThreadTlsSetValue(ZyanThreadTlsIndex index, void* data)
{
    const int error = pthread_setspecific(index, data);
    if (error != 0)
    {
        if (error == EINVAL)
        {
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)#elif defined()(ok bool){//col:136
/*#elif defined(ZYAN_WINDOWS)
ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)
{
    *thread = GetCurrentThread();
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadGetCurrentThreadId()(ok bool){//col:143
/*ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)
{
    *thread_id = GetCurrentThreadId();
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsAlloc()(ok bool){//col:159
/*ZyanStatus ZyanThreadTlsAlloc(ZyanThreadTlsIndex* index, ZyanThreadTlsCallback destructor)
{
    const ZyanThreadTlsIndex value = FlsAlloc(destructor);
    if (value == FLS_OUT_OF_INDEXES)
    {
        return ZYAN_STATUS_OUT_OF_RESOURCES;
    }
    *index = value;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsFree()(ok bool){//col:164
/*ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)
{
    return FlsFree(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsGetValue()(ok bool){//col:171
/*ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)
{
    *data = FlsGetValue(index);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *thread)ZyanStatus ZyanThreadTlsSetValue()(ok bool){//col:186
/*ZyanStatus ZyanThreadTlsSetValue(ZyanThreadTlsIndex index, void* data)
{
    if (!FlsSetValue(index, data))
    {
        const DWORD error = GetLastError();
        if (error == ERROR_INVALID_PARAMETER)
        {
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}



