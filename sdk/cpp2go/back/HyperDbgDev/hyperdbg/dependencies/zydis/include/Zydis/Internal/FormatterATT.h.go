package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\FormatterATT.h.back

const(
ZYDIS_FORMATTER_ATT_H =  //col:33
)

type (
FormatterAtt interface{
  Zyan Disassembler Library ()(ok bool)//col:168
}
)

func NewFormatterAtt() { return & formatterAtt{} }

func (f *formatterAtt)  Zyan Disassembler Library ()(ok bool){//col:168
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd, Joel Hoener
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
 * Implements the `AT&T` style instruction-formatter.
#ifndef ZYDIS_FORMATTER_ATT_H
#define ZYDIS_FORMATTER_ATT_H
#include <Zydis/Formatter.h>
#include <Zydis/Internal/FormatterBase.h>
#include <Zydis/Internal/String.h>
#ifdef __cplusplus
extern "C" {
#endif
ZyanStatus ZydisFormatterATTFormatInstruction(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTFormatOperandMEM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTPrintMnemonic(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTPrintRegister(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg);
ZyanStatus ZydisFormatterATTPrintDISP(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTPrintIMM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
 * The default formatter configuration for `AT&T` style disassembly.
static const ZydisFormatter FORMATTER_ATT =
{
    {
        {
            {
            },
            {
            }
        },
        {
            {
                                    ZYDIS_NUMERIC_BASE_HEX][0].string_data,
            },
            {
            }
        }
    },
};*/
return true
}



