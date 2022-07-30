package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\FormatterIntel.h.back

const(
ZYDIS_FORMATTER_INTEL_H =  //col:33
)

type (
FormatterIntel interface{
  Zyan Disassembler Library ()(ok bool)//col:170
}
)

func NewFormatterIntel() { return & formatterIntel{} }

func (f *formatterIntel)  Zyan Disassembler Library ()(ok bool){//col:170
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
 * Implements the `INTEL` style instruction-formatter.
#ifndef ZYDIS_FORMATTER_INTEL_H
#define ZYDIS_FORMATTER_INTEL_H
#include <Zydis/Formatter.h>
#include <Zydis/Internal/FormatterBase.h>
#include <Zydis/Internal/String.h>
#ifdef __cplusplus
extern "C" {
#endif
ZyanStatus ZydisFormatterIntelFormatInstruction(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterIntelFormatOperandMEM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterIntelPrintMnemonic(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterIntelPrintRegister(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg);
ZyanStatus ZydisFormatterIntelPrintDISP(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterIntelPrintTypecast(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterIntelFormatInstructionMASM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterIntelPrintAddressMASM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
 * The default formatter configuration for `INTEL` style disassembly.
static const ZydisFormatter FORMATTER_INTEL =
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



