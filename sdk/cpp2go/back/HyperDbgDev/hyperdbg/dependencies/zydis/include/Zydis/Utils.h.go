package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Utils.h.back

const(
ZYDIS_UTILS_H =  //col:33
ZYDIS_MAX_INSTRUCTION_SEGMENT_COUNT = 9 //col:51
)

type     ZYDIS_INSTR_SEGMENT_NONE uint32
const(
    ZYDIS_INSTR_SEGMENT_NONE typedef enum ZydisInstructionSegment_ = 1  //col:64
    /** typedef enum ZydisInstructionSegment_ = 2  //col:65
     * The legacy prefixes (including ignored `REX` prefixes). typedef enum ZydisInstructionSegment_ = 3  //col:66
     */ typedef enum ZydisInstructionSegment_ = 4  //col:67
    ZYDIS_INSTR_SEGMENT_PREFIXES typedef enum ZydisInstructionSegment_ = 5  //col:68
    /** typedef enum ZydisInstructionSegment_ = 6  //col:69
     * The effective `REX` prefix byte. typedef enum ZydisInstructionSegment_ = 7  //col:70
     */ typedef enum ZydisInstructionSegment_ = 8  //col:71
    ZYDIS_INSTR_SEGMENT_REX typedef enum ZydisInstructionSegment_ = 9  //col:72
    /** typedef enum ZydisInstructionSegment_ = 10  //col:73
     * The `XOP` prefix bytes. typedef enum ZydisInstructionSegment_ = 11  //col:74
     */ typedef enum ZydisInstructionSegment_ = 12  //col:75
    ZYDIS_INSTR_SEGMENT_XOP typedef enum ZydisInstructionSegment_ = 13  //col:76
    /** typedef enum ZydisInstructionSegment_ = 14  //col:77
     * The `VEX` prefix bytes. typedef enum ZydisInstructionSegment_ = 15  //col:78
     */ typedef enum ZydisInstructionSegment_ = 16  //col:79
    ZYDIS_INSTR_SEGMENT_VEX typedef enum ZydisInstructionSegment_ = 17  //col:80
    /** typedef enum ZydisInstructionSegment_ = 18  //col:81
     * The `EVEX` prefix bytes. typedef enum ZydisInstructionSegment_ = 19  //col:82
     */ typedef enum ZydisInstructionSegment_ = 20  //col:83
    ZYDIS_INSTR_SEGMENT_EVEX typedef enum ZydisInstructionSegment_ = 21  //col:84
    /** typedef enum ZydisInstructionSegment_ = 22  //col:85
     * The `MVEX` prefix bytes. typedef enum ZydisInstructionSegment_ = 23  //col:86
     */ typedef enum ZydisInstructionSegment_ = 24  //col:87
    ZYDIS_INSTR_SEGMENT_MVEX typedef enum ZydisInstructionSegment_ = 25  //col:88
    /** typedef enum ZydisInstructionSegment_ = 26  //col:89
     * The opcode bytes. typedef enum ZydisInstructionSegment_ = 27  //col:90
     */ typedef enum ZydisInstructionSegment_ = 28  //col:91
    ZYDIS_INSTR_SEGMENT_OPCODE typedef enum ZydisInstructionSegment_ = 29  //col:92
    /** typedef enum ZydisInstructionSegment_ = 30  //col:93
     * The `ModRM` byte. typedef enum ZydisInstructionSegment_ = 31  //col:94
     */ typedef enum ZydisInstructionSegment_ = 32  //col:95
    ZYDIS_INSTR_SEGMENT_MODRM typedef enum ZydisInstructionSegment_ = 33  //col:96
    /** typedef enum ZydisInstructionSegment_ = 34  //col:97
     * The `SIB` byte. typedef enum ZydisInstructionSegment_ = 35  //col:98
     */ typedef enum ZydisInstructionSegment_ = 36  //col:99
    ZYDIS_INSTR_SEGMENT_SIB typedef enum ZydisInstructionSegment_ = 37  //col:100
    /** typedef enum ZydisInstructionSegment_ = 38  //col:101
     * The displacement bytes. typedef enum ZydisInstructionSegment_ = 39  //col:102
     */ typedef enum ZydisInstructionSegment_ = 40  //col:103
    ZYDIS_INSTR_SEGMENT_DISPLACEMENT typedef enum ZydisInstructionSegment_ = 41  //col:104
    /** typedef enum ZydisInstructionSegment_ = 42  //col:105
     * The immediate bytes. typedef enum ZydisInstructionSegment_ = 43  //col:106
     */ typedef enum ZydisInstructionSegment_ = 44  //col:107
    ZYDIS_INSTR_SEGMENT_IMMEDIATE typedef enum ZydisInstructionSegment_ = 45  //col:108
    /** typedef enum ZydisInstructionSegment_ = 46  //col:110
     * Maximum value of this enum. typedef enum ZydisInstructionSegment_ = 47  //col:111
     */ typedef enum ZydisInstructionSegment_ = 48  //col:112
    ZYDIS_INSTR_SEGMENT_MAX_VALUE  typedef enum ZydisInstructionSegment_ =  ZYDIS_INSTR_SEGMENT_IMMEDIATE  //col:113
    /** typedef enum ZydisInstructionSegment_ = 50  //col:114
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisInstructionSegment_ = 51  //col:115
     */ typedef enum ZydisInstructionSegment_ = 52  //col:116
    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS  typedef enum ZydisInstructionSegment_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTR_SEGMENT_MAX_VALUE)  //col:117
)



type (
Utils interface{
  Zyan Disassembler Library ()(ok bool)//col:118
         * The offset of the segment relative to the start of the instruction ()(ok bool)//col:144
 * - `IMM` operands with relative address ()(ok bool)//col:266
}
)

func NewUtils() { return & utils{} }

func (u *utils)  Zyan Disassembler Library ()(ok bool){//col:118
/*  Zyan Disassembler Library (Zydis)
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
 * Other utility functions.
#ifndef ZYDIS_UTILS_H
#define ZYDIS_UTILS_H
#include <Zycore/Defines.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYDIS_MAX_INSTRUCTION_SEGMENT_COUNT 9
 * Defines the `ZydisInstructionSegment` struct.
typedef enum ZydisInstructionSegment_
{
    ZYDIS_INSTR_SEGMENT_NONE,
     * The legacy prefixes (including ignored `REX` prefixes).
    ZYDIS_INSTR_SEGMENT_PREFIXES,
     * The effective `REX` prefix byte.
    ZYDIS_INSTR_SEGMENT_REX,
     * The `XOP` prefix bytes.
    ZYDIS_INSTR_SEGMENT_XOP,
     * The `VEX` prefix bytes.
    ZYDIS_INSTR_SEGMENT_VEX,
     * The `EVEX` prefix bytes.
    ZYDIS_INSTR_SEGMENT_EVEX,
     * The `MVEX` prefix bytes.
    ZYDIS_INSTR_SEGMENT_MVEX,
     * The opcode bytes.
    ZYDIS_INSTR_SEGMENT_OPCODE,
     * The `ModRM` byte.
    ZYDIS_INSTR_SEGMENT_MODRM,
     * The `SIB` byte.
    ZYDIS_INSTR_SEGMENT_SIB,
     * The displacement bytes.
    ZYDIS_INSTR_SEGMENT_DISPLACEMENT,
     * The immediate bytes.
    ZYDIS_INSTR_SEGMENT_IMMEDIATE,
     * Maximum value of this enum.
    ZYDIS_INSTR_SEGMENT_MAX_VALUE = ZYDIS_INSTR_SEGMENT_IMMEDIATE,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTR_SEGMENT_MAX_VALUE)
} ZydisInstructionSegment;*/
return true
}

func (u *utils)         * The offset of the segment relative to the start of the instruction ()(ok bool){//col:144
/*         * The offset of the segment relative to the start of the instruction (in bytes).
        ZyanU8 offset;
         * The size of the segment, in bytes.
        ZyanU8 size;
    } segments[ZYDIS_MAX_INSTRUCTION_SEGMENT_COUNT];
} ZydisInstructionSegments;*/
return true
}

func (u *utils) * - `IMM` operands with relative address ()(ok bool){//col:266
/* * - `IMM` operands with relative address (e.g. `JMP`, `CALL`, ...)
 * - `MEM` operands with `RIP`/`EIP`-relative address (e.g. `MOV RAX, [RIP+0x12345678]`)
 * - `MEM` operands with absolute address (e.g. `MOV RAX, [0x12345678]`)
 *   - The displacement needs to get truncated and zero extended
ZYDIS_EXPORT ZyanStatus ZydisCalcAbsoluteAddress(const ZydisDecodedInstruction* instruction,
    const ZydisDecodedOperand* operand, ZyanU64 runtime_address, ZyanU64* result_address);
 * Calculates the absolute address value for the given instruction operand.
 *
 *
 *
 * This function behaves like `ZydisCalcAbsoluteAddress` but takes an additional register-context
 * argument to allow calculation of addresses depending on runtime register values.
 *
 * Note that `IP/EIP/RIP` from the register-context will be ignored in favor of the passed
 * runtime-address.
ZYDIS_EXPORT ZyanStatus ZydisCalcAbsoluteAddressEx(const ZydisDecodedInstruction* instruction,
    const ZydisDecodedOperand* operand, ZyanU64 runtime_address,
    const ZydisRegisterContext* register_context, ZyanU64* result_address);
 * Returns a mask of accessed CPU-flags matching the given `action`.
 *
 *
ZYDIS_EXPORT ZyanStatus ZydisGetAccessedFlagsByAction(const ZydisDecodedInstruction* instruction,
    ZydisCPUFlagAction action, ZydisCPUFlags* flags);
 * Returns a mask of accessed CPU-flags that are read (tested) by the current instruction.
 *
 *
ZYDIS_EXPORT ZyanStatus ZydisGetAccessedFlagsRead(const ZydisDecodedInstruction* instruction,
    ZydisCPUFlags* flags);
 * Returns a mask of accessed CPU-flags that are written (modified, undefined)
 * by the current instruction.
 *
 *
ZYDIS_EXPORT ZyanStatus ZydisGetAccessedFlagsWritten(const ZydisDecodedInstruction* instruction,
    ZydisCPUFlags* flags);
 * Returns offsets and sizes of all logical instruction segments (e.g. `OPCODE`,
 * `MODRM`, ...).
 *
 *
ZYDIS_EXPORT ZyanStatus ZydisGetInstructionSegments(const ZydisDecodedInstruction* instruction,
    ZydisInstructionSegments* segments);
#ifdef __cplusplus
}*/
return true
}



