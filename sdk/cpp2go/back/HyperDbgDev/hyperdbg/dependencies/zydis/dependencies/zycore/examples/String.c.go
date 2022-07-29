package examples
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c.back

type (
String interface{
  Zyan Core Library ()(ok bool)//col:71
static ZyanStatus TestDynamic()(ok bool)//col:82
static ZyanStatus TestStatic()(ok bool)//col:93
static ZyanStatus TestAllocator()(ok bool)//col:166
int main()(ok bool)//col:190
}
)

func NewString() { return & string{} }

func (s *string)  Zyan Core Library ()(ok bool){//col:71
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
#include <stdio.h>
#include <Zycore/Allocator.h>
#include <Zycore/Defines.h>
#include <Zycore/LibC.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
 *
 *
static ZyanStatus PerformBasicTests(ZyanString* string)
{
    ZYAN_ASSERT(string);
    ZYAN_UNUSED(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)static ZyanStatus TestDynamic()(ok bool){//col:82
/*static ZyanStatus TestDynamic(void)
{
    PerformBasicTests(ZYAN_NULL);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)static ZyanStatus TestStatic()(ok bool){//col:93
/*static ZyanStatus TestStatic(void)
{
    PerformBasicTests(ZYAN_NULL);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)static ZyanStatus TestAllocator()(ok bool){//col:166
/*static ZyanStatus TestAllocator(void)
{
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)int main()(ok bool){//col:190
/*int main()
{
    if (!ZYAN_SUCCESS(TestDynamic()))
    {
        return EXIT_FAILURE;
    }
    if (!ZYAN_SUCCESS(TestStatic()))
    {
        return EXIT_FAILURE;
    }
    if (!ZYAN_SUCCESS(TestAllocator()))
    {
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}*/
return true
}



