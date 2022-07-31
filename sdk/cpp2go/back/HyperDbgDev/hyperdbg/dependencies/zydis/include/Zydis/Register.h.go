package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Register.h.back

const(
ZYDIS_REGISTER_H =  //col:33
)

type     ZYDIS_REGCLASS_INVALID uint32
const(
    ZYDIS_REGCLASS_INVALID typedef enum ZydisRegisterClass_ = 1  //col:70
    /** typedef enum ZydisRegisterClass_ = 2  //col:71
     * 8-bit general-purpose registers. typedef enum ZydisRegisterClass_ = 3  //col:72
     */ typedef enum ZydisRegisterClass_ = 4  //col:73
    ZYDIS_REGCLASS_GPR8 typedef enum ZydisRegisterClass_ = 5  //col:74
    /** typedef enum ZydisRegisterClass_ = 6  //col:75
     * 16-bit general-purpose registers. typedef enum ZydisRegisterClass_ = 7  //col:76
     */ typedef enum ZydisRegisterClass_ = 8  //col:77
    ZYDIS_REGCLASS_GPR16 typedef enum ZydisRegisterClass_ = 9  //col:78
    /** typedef enum ZydisRegisterClass_ = 10  //col:79
     * 32-bit general-purpose registers. typedef enum ZydisRegisterClass_ = 11  //col:80
     */ typedef enum ZydisRegisterClass_ = 12  //col:81
    ZYDIS_REGCLASS_GPR32 typedef enum ZydisRegisterClass_ = 13  //col:82
    /** typedef enum ZydisRegisterClass_ = 14  //col:83
     * 64-bit general-purpose registers. typedef enum ZydisRegisterClass_ = 15  //col:84
     */ typedef enum ZydisRegisterClass_ = 16  //col:85
    ZYDIS_REGCLASS_GPR64 typedef enum ZydisRegisterClass_ = 17  //col:86
    /** typedef enum ZydisRegisterClass_ = 18  //col:87
     * Floating point legacy registers. typedef enum ZydisRegisterClass_ = 19  //col:88
     */ typedef enum ZydisRegisterClass_ = 20  //col:89
    ZYDIS_REGCLASS_X87 typedef enum ZydisRegisterClass_ = 21  //col:90
    /** typedef enum ZydisRegisterClass_ = 22  //col:91
     * Floating point multimedia registers. typedef enum ZydisRegisterClass_ = 23  //col:92
     */ typedef enum ZydisRegisterClass_ = 24  //col:93
    ZYDIS_REGCLASS_MMX typedef enum ZydisRegisterClass_ = 25  //col:94
    /** typedef enum ZydisRegisterClass_ = 26  //col:95
     * 128-bit vector registers. typedef enum ZydisRegisterClass_ = 27  //col:96
     */ typedef enum ZydisRegisterClass_ = 28  //col:97
    ZYDIS_REGCLASS_XMM typedef enum ZydisRegisterClass_ = 29  //col:98
    /** typedef enum ZydisRegisterClass_ = 30  //col:99
     * 256-bit vector registers. typedef enum ZydisRegisterClass_ = 31  //col:100
     */ typedef enum ZydisRegisterClass_ = 32  //col:101
    ZYDIS_REGCLASS_YMM typedef enum ZydisRegisterClass_ = 33  //col:102
    /** typedef enum ZydisRegisterClass_ = 34  //col:103
     * 512-bit vector registers. typedef enum ZydisRegisterClass_ = 35  //col:104
     */ typedef enum ZydisRegisterClass_ = 36  //col:105
    ZYDIS_REGCLASS_ZMM typedef enum ZydisRegisterClass_ = 37  //col:106
    /** typedef enum ZydisRegisterClass_ = 38  //col:107
     * Flags registers. typedef enum ZydisRegisterClass_ = 39  //col:108
     */ typedef enum ZydisRegisterClass_ = 40  //col:109
    ZYDIS_REGCLASS_FLAGS typedef enum ZydisRegisterClass_ = 41  //col:110
    /** typedef enum ZydisRegisterClass_ = 42  //col:111
     * Instruction-pointer registers. typedef enum ZydisRegisterClass_ = 43  //col:112
     */ typedef enum ZydisRegisterClass_ = 44  //col:113
    ZYDIS_REGCLASS_IP typedef enum ZydisRegisterClass_ = 45  //col:114
    /** typedef enum ZydisRegisterClass_ = 46  //col:115
     * Segment registers. typedef enum ZydisRegisterClass_ = 47  //col:116
     */ typedef enum ZydisRegisterClass_ = 48  //col:117
    ZYDIS_REGCLASS_SEGMENT typedef enum ZydisRegisterClass_ = 49  //col:118
    /** typedef enum ZydisRegisterClass_ = 50  //col:119
     * Test registers. typedef enum ZydisRegisterClass_ = 51  //col:120
     */ typedef enum ZydisRegisterClass_ = 52  //col:121
    ZYDIS_REGCLASS_TEST typedef enum ZydisRegisterClass_ = 53  //col:122
    /** typedef enum ZydisRegisterClass_ = 54  //col:123
     * Control registers. typedef enum ZydisRegisterClass_ = 55  //col:124
     */ typedef enum ZydisRegisterClass_ = 56  //col:125
    ZYDIS_REGCLASS_CONTROL typedef enum ZydisRegisterClass_ = 57  //col:126
    /** typedef enum ZydisRegisterClass_ = 58  //col:127
     * Debug registers. typedef enum ZydisRegisterClass_ = 59  //col:128
     */ typedef enum ZydisRegisterClass_ = 60  //col:129
    ZYDIS_REGCLASS_DEBUG typedef enum ZydisRegisterClass_ = 61  //col:130
    /** typedef enum ZydisRegisterClass_ = 62  //col:131
     * Mask registers. typedef enum ZydisRegisterClass_ = 63  //col:132
     */ typedef enum ZydisRegisterClass_ = 64  //col:133
    ZYDIS_REGCLASS_MASK typedef enum ZydisRegisterClass_ = 65  //col:134
    /** typedef enum ZydisRegisterClass_ = 66  //col:135
     * Bound registers. typedef enum ZydisRegisterClass_ = 67  //col:136
     */ typedef enum ZydisRegisterClass_ = 68  //col:137
    ZYDIS_REGCLASS_BOUND typedef enum ZydisRegisterClass_ = 69  //col:138
    /** typedef enum ZydisRegisterClass_ = 70  //col:140
     * Maximum value of this enum. typedef enum ZydisRegisterClass_ = 71  //col:141
     */ typedef enum ZydisRegisterClass_ = 72  //col:142
    ZYDIS_REGCLASS_MAX_VALUE  typedef enum ZydisRegisterClass_ =  ZYDIS_REGCLASS_BOUND  //col:143
    /** typedef enum ZydisRegisterClass_ = 74  //col:144
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisRegisterClass_ = 75  //col:145
     */ typedef enum ZydisRegisterClass_ = 76  //col:146
    ZYDIS_REGCLASS_REQUIRED_BITS  typedef enum ZydisRegisterClass_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_REGCLASS_MAX_VALUE)  //col:147
)



type (
Register interface{
  Zyan Disassembler Library ()(ok bool)//col:148
ZYDIS_EXPORT ZydisRegister ZydisRegisterEncode()(ok bool)//col:286
}
)

func NewRegister() { return & register{} }

func (r *register)  Zyan Disassembler Library ()(ok bool){//col:148
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
 * Utility functions and constants for registers.
#ifndef ZYDIS_REGISTER_H
#define ZYDIS_REGISTER_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#include <Zydis/SharedTypes.h>
#include <Zydis/ShortString.h>
#ifdef __cplusplus
extern "C" {
#endif
#include <Zydis/Generated/EnumRegister.h>
 * Defines the `ZydisRegisterClass` enum.
 *
 * Please note that this enum does not contain a matching entry for all values of the
 * `ZydisRegister` enum, but only for those registers where it makes sense to logically group them
 * for decoding/encoding purposes.
 *
 * These are mainly the registers that can be identified by an id within their corresponding
 * register-class. The `IP` and `FLAGS` values are exceptions to this rule.
typedef enum ZydisRegisterClass_
{
    ZYDIS_REGCLASS_INVALID,
     * 8-bit general-purpose registers.
    ZYDIS_REGCLASS_GPR8,
     * 16-bit general-purpose registers.
    ZYDIS_REGCLASS_GPR16,
     * 32-bit general-purpose registers.
    ZYDIS_REGCLASS_GPR32,
     * 64-bit general-purpose registers.
    ZYDIS_REGCLASS_GPR64,
     * Floating point legacy registers.
    ZYDIS_REGCLASS_X87,
     * Floating point multimedia registers.
    ZYDIS_REGCLASS_MMX,
     * 128-bit vector registers.
    ZYDIS_REGCLASS_XMM,
     * 256-bit vector registers.
    ZYDIS_REGCLASS_YMM,
     * 512-bit vector registers.
    ZYDIS_REGCLASS_ZMM,
     * Flags registers.
    ZYDIS_REGCLASS_FLAGS,
     * Instruction-pointer registers.
    ZYDIS_REGCLASS_IP,
     * Segment registers.
    ZYDIS_REGCLASS_SEGMENT,
     * Test registers.
    ZYDIS_REGCLASS_TEST,
     * Control registers.
    ZYDIS_REGCLASS_CONTROL,
     * Debug registers.
    ZYDIS_REGCLASS_DEBUG,
     * Mask registers.
    ZYDIS_REGCLASS_MASK,
     * Bound registers.
    ZYDIS_REGCLASS_BOUND,
     * Maximum value of this enum.
    ZYDIS_REGCLASS_MAX_VALUE = ZYDIS_REGCLASS_BOUND,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_REGCLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_REGCLASS_MAX_VALUE)
} ZydisRegisterClass;*/
return true
}

func (r *register)ZYDIS_EXPORT ZydisRegister ZydisRegisterEncode()(ok bool){//col:286
/*ZYDIS_EXPORT ZydisRegister ZydisRegisterEncode(ZydisRegisterClass register_class, ZyanU8 id);
 * Returns the id of the specified register.
 *
 *
ZYDIS_EXPORT ZyanI8 ZydisRegisterGetId(ZydisRegister reg);
 * Returns the register-class of the specified register.
 *
 *
ZYDIS_EXPORT ZydisRegisterClass ZydisRegisterGetClass(ZydisRegister reg);
 * Returns the width of the specified register.
 *
 *
 *          invalid for the active machine-mode.
ZYDIS_EXPORT ZydisRegisterWidth ZydisRegisterGetWidth(ZydisMachineMode mode, ZydisRegister reg);
 * Returns the largest enclosing register of the given register.
 *
 *
 *          register is invalid for the active machine-mode or does not have an enclosing-register.
ZYDIS_EXPORT ZydisRegister ZydisRegisterGetLargestEnclosing(ZydisMachineMode mode,
    ZydisRegister reg);
 * Returns the specified register string.
 *
 *
ZYDIS_EXPORT const char* ZydisRegisterGetString(ZydisRegister reg);
 * Returns the specified register string as `ZydisShortString`.
 *
 *
 *
 * The `buffer` of the returned struct is guaranteed to be zero-terminated in this special case.
ZYDIS_EXPORT const ZydisShortString* ZydisRegisterGetStringWrapped(ZydisRegister reg);
 * Returns the width of the specified register-class.
 *
 *
ZYDIS_EXPORT ZydisRegisterWidth ZydisRegisterClassGetWidth(ZydisMachineMode mode,
    ZydisRegisterClass register_class);
#ifdef __cplusplus
}*/
return true
}



