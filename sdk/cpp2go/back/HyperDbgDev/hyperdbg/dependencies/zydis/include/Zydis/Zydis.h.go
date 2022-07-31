package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Zydis.h.back

const(
ZYDIS_H =  //col:33
ZYDIS_VERSION = (ZyanU64)0x0003000100000000 //col:69
ZYDIS_VERSION_MAJOR(version) = (ZyanU16)(((version) & 0xFFFF000000000000) >> 48) //col:80
ZYDIS_VERSION_MINOR(version) = (ZyanU16)(((version) & 0x0000FFFF00000000) >> 32) //col:87
ZYDIS_VERSION_PATCH(version) = (ZyanU16)(((version) & 0x00000000FFFF0000) >> 16) //col:94
ZYDIS_VERSION_BUILD(version) = (ZyanU16)((version) & 0x000000000000FFFF) //col:101
)

type     ZYDIS_FEATURE_DECODER uint32
const(
    ZYDIS_FEATURE_DECODER typedef enum ZydisFeature_ = 1  //col:114
    ZYDIS_FEATURE_FORMATTER typedef enum ZydisFeature_ = 2  //col:115
    ZYDIS_FEATURE_AVX512 typedef enum ZydisFeature_ = 3  //col:116
    ZYDIS_FEATURE_KNC typedef enum ZydisFeature_ = 4  //col:117
    /** typedef enum ZydisFeature_ = 5  //col:119
     * Maximum value of this enum. typedef enum ZydisFeature_ = 6  //col:120
     */ typedef enum ZydisFeature_ = 7  //col:121
    ZYDIS_FEATURE_MAX_VALUE  typedef enum ZydisFeature_ =  ZYDIS_FEATURE_KNC  //col:122
    /** typedef enum ZydisFeature_ = 9  //col:123
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisFeature_ = 10  //col:124
     */ typedef enum ZydisFeature_ = 11  //col:125
    ZYDIS_FEATURE_REQUIRED_BITS  typedef enum ZydisFeature_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FEATURE_MAX_VALUE)  //col:126
)



type (
Zydis interface{
  Zyan Disassembler Library ()(ok bool)//col:127
ZYDIS_EXPORT ZyanU64 ZydisGetVersion()(ok bool)//col:166
}
)

func NewZydis() { return & zydis{} }

func (z *zydis)  Zyan Disassembler Library ()(ok bool){//col:127
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
 * Master include file, including everything else.
#ifndef ZYDIS_H
#define ZYDIS_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#ifndef ZYDIS_DISABLE_DECODER
#   include <Zydis/Decoder.h>
#   include <Zydis/DecoderTypes.h>
#endif
#ifndef ZYDIS_DISABLE_FORMATTER
#   include <Zydis/Formatter.h>
#endif
#include <Zydis/MetaInfo.h>
#include <Zydis/Mnemonic.h>
#include <Zydis/Register.h>
#include <Zydis/SharedTypes.h>
#include <Zydis/Status.h>
#include <Zydis/Utils.h>
#ifdef __cplusplus
extern "C" {
#endif
 * A macro that defines the zydis version.
#define ZYDIS_VERSION (ZyanU64)0x0003000100000000
 * Extracts the major-part of the zydis version.
 *
#define ZYDIS_VERSION_MAJOR(version) (ZyanU16)(((version) & 0xFFFF000000000000) >> 48)
 * Extracts the minor-part of the zydis version.
 *
#define ZYDIS_VERSION_MINOR(version) (ZyanU16)(((version) & 0x0000FFFF00000000) >> 32)
 * Extracts the patch-part of the zydis version.
 *
#define ZYDIS_VERSION_PATCH(version) (ZyanU16)(((version) & 0x00000000FFFF0000) >> 16)
 * Extracts the build-part of the zydis version.
 *
#define ZYDIS_VERSION_BUILD(version) (ZyanU16)((version) & 0x000000000000FFFF)
 * Defines the `ZydisFeature` enum.
typedef enum ZydisFeature_
{
    ZYDIS_FEATURE_DECODER,
    ZYDIS_FEATURE_FORMATTER,
    ZYDIS_FEATURE_AVX512,
    ZYDIS_FEATURE_KNC,
     * Maximum value of this enum.
    ZYDIS_FEATURE_MAX_VALUE = ZYDIS_FEATURE_KNC,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_FEATURE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FEATURE_MAX_VALUE)
} ZydisFeature;*/
return true
}

func (z *zydis)ZYDIS_EXPORT ZyanU64 ZydisGetVersion()(ok bool){//col:166
/*ZYDIS_EXPORT ZyanU64 ZydisGetVersion(void);
 * Checks, if the specified feature is enabled in the current zydis library instance.
 *
 *
 *          zyan status code, if an error occured.
ZYDIS_EXPORT ZyanStatus ZydisIsFeatureEnabled(ZydisFeature feature);
#ifdef __cplusplus
}*/
return true
}



