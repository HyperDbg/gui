package API
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\API\Memory.h.back

const(
ZYCORE_MEMORY_H =  //col:33
)

type #if   defined(ZYAN_WINDOWS) uint32
const(
#if   defined(ZYAN_WINDOWS) typedef enum ZyanMemoryPageProtection_ = 1  //col:57
    ZYAN_PAGE_READONLY           typedef enum ZyanMemoryPageProtection_ =  PAGE_READONLY             //col:59
    ZYAN_PAGE_READWRITE          typedef enum ZyanMemoryPageProtection_ =  PAGE_READWRITE  //col:60
    ZYAN_PAGE_EXECUTE            typedef enum ZyanMemoryPageProtection_ =  PAGE_EXECUTE               //col:61
    ZYAN_PAGE_EXECUTE_READ       typedef enum ZyanMemoryPageProtection_ =  PAGE_EXECUTE_READ  //col:62
    ZYAN_PAGE_EXECUTE_READWRITE  typedef enum ZyanMemoryPageProtection_ =  PAGE_EXECUTE_READWRITE  //col:63
#elif defined(ZYAN_POSIX) typedef enum ZyanMemoryPageProtection_ = 7  //col:65
    ZYAN_PAGE_READONLY           typedef enum ZyanMemoryPageProtection_ =  PROT_READ             //col:67
    ZYAN_PAGE_READWRITE          typedef enum ZyanMemoryPageProtection_ =  PROT_READ | PROT_WRITE  //col:68
    ZYAN_PAGE_EXECUTE            typedef enum ZyanMemoryPageProtection_ =  PROT_EXEC               //col:69
    ZYAN_PAGE_EXECUTE_READ       typedef enum ZyanMemoryPageProtection_ =  PROT_EXEC | PROT_READ  //col:70
    ZYAN_PAGE_EXECUTE_READWRITE  typedef enum ZyanMemoryPageProtection_ =  PROT_EXEC | PROT_READ | PROT_WRITE  //col:71
#endif typedef enum ZyanMemoryPageProtection_ = 13  //col:73
)



type (
Memory interface{
  Zyan Core Library ()(ok bool)//col:74
}
)

func NewMemory() { return & memory{} }

func (m *memory)  Zyan Core Library ()(ok bool){//col:74
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
#ifndef ZYCORE_MEMORY_H
#define ZYCORE_MEMORY_H
#include <ZycoreExportConfig.h>
#include <Zycore/Defines.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#if   defined(ZYAN_WINDOWS)
#   include <windows.h>
#elif defined(ZYAN_POSIX)
#   include <sys/mman.h>
#else
#   error "Unsupported platform detected"
#endif
typedef enum ZyanMemoryPageProtection_
{
#if   defined(ZYAN_WINDOWS)
    ZYAN_PAGE_READONLY          = PAGE_READONLY,           
    ZYAN_PAGE_READWRITE         = PAGE_READWRITE,
    ZYAN_PAGE_EXECUTE           = PAGE_EXECUTE,             
    ZYAN_PAGE_EXECUTE_READ      = PAGE_EXECUTE_READ,
    ZYAN_PAGE_EXECUTE_READWRITE = PAGE_EXECUTE_READWRITE
#elif defined(ZYAN_POSIX)
    ZYAN_PAGE_READONLY          = PROT_READ,           
    ZYAN_PAGE_READWRITE         = PROT_READ | PROT_WRITE,
    ZYAN_PAGE_EXECUTE           = PROT_EXEC,             
    ZYAN_PAGE_EXECUTE_READ      = PROT_EXEC | PROT_READ,
    ZYAN_PAGE_EXECUTE_READWRITE = PROT_EXEC | PROT_READ | PROT_WRITE
#endif
} ZyanMemoryPageProtection;*/
return true
}



