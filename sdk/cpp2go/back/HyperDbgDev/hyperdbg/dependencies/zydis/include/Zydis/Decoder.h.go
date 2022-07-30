package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Decoder.h.back

const(
ZYDIS_DECODER_H =  //col:33
)

type     /** uint32
const(
    /** typedef enum ZydisDecoderMode_ = 1  //col:57
     * Enables minimal instruction decoding without semantic analysis. typedef enum ZydisDecoderMode_ = 2  //col:58
     * typedef enum ZydisDecoderMode_ = 3  //col:59
     * This mode provides access to the mnemonic the instruction-length the effective typedef enum ZydisDecoderMode_ = 4  //col:60
     * operand-size the effective address-width some attributes (e.g. `ZYDIS_ATTRIB_IS_RELATIVE`) typedef enum ZydisDecoderMode_ = 5  //col:61
     * and all of the information in the `raw` field of the `ZydisDecodedInstruction` struct. typedef enum ZydisDecoderMode_ = 6  //col:62
     * typedef enum ZydisDecoderMode_ = 7  //col:63
     * Operands most attributes and other specific information (like `AVX` info) are not typedef enum ZydisDecoderMode_ = 8  //col:64
     * accessible in this mode. typedef enum ZydisDecoderMode_ = 9  //col:65
     * typedef enum ZydisDecoderMode_ = 10  //col:66
     * This mode is NOT enabled by default. typedef enum ZydisDecoderMode_ = 11  //col:67
     */ typedef enum ZydisDecoderMode_ = 12  //col:68
    ZYDIS_DECODER_MODE_MINIMAL typedef enum ZydisDecoderMode_ = 13  //col:69
    /** typedef enum ZydisDecoderMode_ = 14  //col:70
     * Enables the `AMD`-branch mode. typedef enum ZydisDecoderMode_ = 15  //col:71
     * typedef enum ZydisDecoderMode_ = 16  //col:72
     * Intel ignores the operand-size override-prefix (`0x66`) for all branches with 32-bit typedef enum ZydisDecoderMode_ = 17  //col:73
     * immediates and forces the operand-size of the instruction to 64-bit in 64-bit mode. typedef enum ZydisDecoderMode_ = 18  //col:74
     * In `AMD`-branch mode `0x66` is not ignored and changes the operand-size and the size of the typedef enum ZydisDecoderMode_ = 19  //col:75
     * immediate to 16-bit. typedef enum ZydisDecoderMode_ = 20  //col:76
     * typedef enum ZydisDecoderMode_ = 21  //col:77
     * This mode is NOT enabled by default. typedef enum ZydisDecoderMode_ = 22  //col:78
     */ typedef enum ZydisDecoderMode_ = 23  //col:79
    ZYDIS_DECODER_MODE_AMD_BRANCHES typedef enum ZydisDecoderMode_ = 24  //col:80
    /** typedef enum ZydisDecoderMode_ = 25  //col:81
     * Enables `KNC` compatibility-mode. typedef enum ZydisDecoderMode_ = 26  //col:82
     * typedef enum ZydisDecoderMode_ = 27  //col:83
     * `KNC` and `KNL+` chips are sharing opcodes and encodings for some mask-related instructions. typedef enum ZydisDecoderMode_ = 28  //col:84
     * Enable this mode to use the old `KNC` specifications (different mnemonics operands ..). typedef enum ZydisDecoderMode_ = 29  //col:85
     * typedef enum ZydisDecoderMode_ = 30  //col:86
     * This mode is NOT enabled by default. typedef enum ZydisDecoderMode_ = 31  //col:87
     */ typedef enum ZydisDecoderMode_ = 32  //col:88
    ZYDIS_DECODER_MODE_KNC typedef enum ZydisDecoderMode_ = 33  //col:89
    /** typedef enum ZydisDecoderMode_ = 34  //col:90
     * Enables the `MPX` mode. typedef enum ZydisDecoderMode_ = 35  //col:91
     * typedef enum ZydisDecoderMode_ = 36  //col:92
     * The `MPX` isa-extension reuses (overrides) some of the widenop instruction opcodes. typedef enum ZydisDecoderMode_ = 37  //col:93
     * typedef enum ZydisDecoderMode_ = 38  //col:94
     * This mode is enabled by default. typedef enum ZydisDecoderMode_ = 39  //col:95
     */ typedef enum ZydisDecoderMode_ = 40  //col:96
    ZYDIS_DECODER_MODE_MPX typedef enum ZydisDecoderMode_ = 41  //col:97
    /** typedef enum ZydisDecoderMode_ = 42  //col:98
     * Enables the `CET` mode. typedef enum ZydisDecoderMode_ = 43  //col:99
     * typedef enum ZydisDecoderMode_ = 44  //col:100
     * The `CET` isa-extension reuses (overrides) some of the widenop instruction opcodes. typedef enum ZydisDecoderMode_ = 45  //col:101
     * typedef enum ZydisDecoderMode_ = 46  //col:102
     * This mode is enabled by default. typedef enum ZydisDecoderMode_ = 47  //col:103
     */ typedef enum ZydisDecoderMode_ = 48  //col:104
    ZYDIS_DECODER_MODE_CET typedef enum ZydisDecoderMode_ = 49  //col:105
    /** typedef enum ZydisDecoderMode_ = 50  //col:106
     * Enables the `LZCNT` mode. typedef enum ZydisDecoderMode_ = 51  //col:107
     * typedef enum ZydisDecoderMode_ = 52  //col:108
     * The `LZCNT` isa-extension reuses (overrides) some of the widenop instruction opcodes. typedef enum ZydisDecoderMode_ = 53  //col:109
     * typedef enum ZydisDecoderMode_ = 54  //col:110
     * This mode is enabled by default. typedef enum ZydisDecoderMode_ = 55  //col:111
     */ typedef enum ZydisDecoderMode_ = 56  //col:112
    ZYDIS_DECODER_MODE_LZCNT typedef enum ZydisDecoderMode_ = 57  //col:113
    /** typedef enum ZydisDecoderMode_ = 58  //col:114
     * Enables the `TZCNT` mode. typedef enum ZydisDecoderMode_ = 59  //col:115
     * typedef enum ZydisDecoderMode_ = 60  //col:116
     * The `TZCNT` isa-extension reuses (overrides) some of the widenop instruction opcodes. typedef enum ZydisDecoderMode_ = 61  //col:117
     * typedef enum ZydisDecoderMode_ = 62  //col:118
     * This mode is enabled by default. typedef enum ZydisDecoderMode_ = 63  //col:119
     */ typedef enum ZydisDecoderMode_ = 64  //col:120
    ZYDIS_DECODER_MODE_TZCNT typedef enum ZydisDecoderMode_ = 65  //col:121
    /** typedef enum ZydisDecoderMode_ = 66  //col:122
     * Enables the `WBNOINVD` mode. typedef enum ZydisDecoderMode_ = 67  //col:123
     * typedef enum ZydisDecoderMode_ = 68  //col:124
     * The `WBINVD` instruction is interpreted as `WBNOINVD` on ICL chips if a `F3` prefix is typedef enum ZydisDecoderMode_ = 69  //col:125
     * used. typedef enum ZydisDecoderMode_ = 70  //col:126
     * typedef enum ZydisDecoderMode_ = 71  //col:127
     * This mode is disabled by default. typedef enum ZydisDecoderMode_ = 72  //col:128
     */ typedef enum ZydisDecoderMode_ = 73  //col:129
    ZYDIS_DECODER_MODE_WBNOINVD typedef enum ZydisDecoderMode_ = 74  //col:130
     /** typedef enum ZydisDecoderMode_ = 75  //col:131
     * Enables the `CLDEMOTE` mode. typedef enum ZydisDecoderMode_ = 76  //col:132
     * typedef enum ZydisDecoderMode_ = 77  //col:133
     * The `CLDEMOTE` isa-extension reuses (overrides) some of the widenop instruction opcodes. typedef enum ZydisDecoderMode_ = 78  //col:134
     * typedef enum ZydisDecoderMode_ = 79  //col:135
     * This mode is enabled by default. typedef enum ZydisDecoderMode_ = 80  //col:136
     */ typedef enum ZydisDecoderMode_ = 81  //col:137
    ZYDIS_DECODER_MODE_CLDEMOTE typedef enum ZydisDecoderMode_ = 82  //col:138
    /** typedef enum ZydisDecoderMode_ = 83  //col:140
     * Maximum value of this enum. typedef enum ZydisDecoderMode_ = 84  //col:141
     */ typedef enum ZydisDecoderMode_ = 85  //col:142
    ZYDIS_DECODER_MODE_MAX_VALUE  typedef enum ZydisDecoderMode_ =  ZYDIS_DECODER_MODE_CLDEMOTE  //col:143
    /** typedef enum ZydisDecoderMode_ = 87  //col:144
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisDecoderMode_ = 88  //col:145
     */ typedef enum ZydisDecoderMode_ = 89  //col:146
    ZYDIS_DECODER_MODE_REQUIRED_BITS  typedef enum ZydisDecoderMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_DECODER_MODE_MAX_VALUE)  //col:147
)



type (
Decoder interface{
  Zyan Disassembler Library ()(ok bool)//col:148
ZYDIS_EXPORT ZyanStatus ZydisDecoderInit()(ok bool)//col:231
}
)

func NewDecoder() { return & decoder{} }

func (d *decoder)  Zyan Disassembler Library ()(ok bool){//col:148
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
 * Functions for decoding instructions.
#ifndef ZYDIS_DECODER_H
#define ZYDIS_DECODER_H
#include <Zycore/Types.h>
#include <Zycore/Defines.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
 * Defines the `ZydisDecoderMode` enum.
typedef enum ZydisDecoderMode_
{
     * Enables minimal instruction decoding without semantic analysis.
     *
     * This mode provides access to the mnemonic, the instruction-length, the effective
     * operand-size, the effective address-width, some attributes (e.g. `ZYDIS_ATTRIB_IS_RELATIVE`)
     * and all of the information in the `raw` field of the `ZydisDecodedInstruction` struct.
     *
     * Operands, most attributes and other specific information (like `AVX` info) are not
     * accessible in this mode.
     *
     * This mode is NOT enabled by default.
    ZYDIS_DECODER_MODE_MINIMAL,
     * Enables the `AMD`-branch mode.
     *
     * Intel ignores the operand-size override-prefix (`0x66`) for all branches with 32-bit
     * immediates and forces the operand-size of the instruction to 64-bit in 64-bit mode.
     * In `AMD`-branch mode `0x66` is not ignored and changes the operand-size and the size of the
     * immediate to 16-bit.
     *
     * This mode is NOT enabled by default.
    ZYDIS_DECODER_MODE_AMD_BRANCHES,
     * Enables `KNC` compatibility-mode.
     *
     * `KNC` and `KNL+` chips are sharing opcodes and encodings for some mask-related instructions.
     * Enable this mode to use the old `KNC` specifications (different mnemonics, operands, ..).
     *
     * This mode is NOT enabled by default.
    ZYDIS_DECODER_MODE_KNC,
     * Enables the `MPX` mode.
     *
     * The `MPX` isa-extension reuses (overrides) some of the widenop instruction opcodes.
     *
     * This mode is enabled by default.
    ZYDIS_DECODER_MODE_MPX,
     * Enables the `CET` mode.
     *
     * The `CET` isa-extension reuses (overrides) some of the widenop instruction opcodes.
     *
     * This mode is enabled by default.
    ZYDIS_DECODER_MODE_CET,
     * Enables the `LZCNT` mode.
     *
     * The `LZCNT` isa-extension reuses (overrides) some of the widenop instruction opcodes.
     *
     * This mode is enabled by default.
    ZYDIS_DECODER_MODE_LZCNT,
     * Enables the `TZCNT` mode.
     *
     * The `TZCNT` isa-extension reuses (overrides) some of the widenop instruction opcodes.
     *
     * This mode is enabled by default.
    ZYDIS_DECODER_MODE_TZCNT,
     * Enables the `WBNOINVD` mode.
     *
     * The `WBINVD` instruction is interpreted as `WBNOINVD` on ICL chips, if a `F3` prefix is
     * used.
     *
     * This mode is disabled by default.
    ZYDIS_DECODER_MODE_WBNOINVD,
     * Enables the `CLDEMOTE` mode.
     *
     * The `CLDEMOTE` isa-extension reuses (overrides) some of the widenop instruction opcodes.
     *
     * This mode is enabled by default.
    ZYDIS_DECODER_MODE_CLDEMOTE,
     * Maximum value of this enum.
    ZYDIS_DECODER_MODE_MAX_VALUE = ZYDIS_DECODER_MODE_CLDEMOTE,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_DECODER_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_DECODER_MODE_MAX_VALUE)
} ZydisDecoderMode;*/
return true
}

func (d *decoder)ZYDIS_EXPORT ZyanStatus ZydisDecoderInit()(ok bool){//col:231
/*ZYDIS_EXPORT ZyanStatus ZydisDecoderInit(ZydisDecoder* decoder, ZydisMachineMode machine_mode,
    ZydisAddressWidth address_width);
 * Enables or disables the specified decoder-mode.
 *
 *
ZYDIS_EXPORT ZyanStatus ZydisDecoderEnableMode(ZydisDecoder* decoder, ZydisDecoderMode mode,
    ZyanBool enabled);
 * Decodes the instruction in the given input `buffer`.
 *
 *                      details about the decoded instruction.
 *
ZYDIS_EXPORT ZyanStatus ZydisDecoderDecodeBuffer(const ZydisDecoder* decoder,
    const void* buffer, ZyanUSize length, ZydisDecodedInstruction* instruction);
#ifdef __cplusplus
}*/
return true
}



