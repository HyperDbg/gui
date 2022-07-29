package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Formatter.h.back

const(
ZYDIS_FORMATTER_H =  //col:33
ZYDIS_RUNTIME_ADDRESS_NONE = (ZyanU64)(-1) //col:53
)

type     /** uint32
const(
    /** typedef enum ZydisFormatterStyle_ = 1  //col:68
     * Generates `AT&T`-style disassembly. typedef enum ZydisFormatterStyle_ = 2  //col:69
     */ typedef enum ZydisFormatterStyle_ = 3  //col:70
    ZYDIS_FORMATTER_STYLE_ATT typedef enum ZydisFormatterStyle_ = 4  //col:71
    /** typedef enum ZydisFormatterStyle_ = 5  //col:72
     * Generates `Intel`-style disassembly. typedef enum ZydisFormatterStyle_ = 6  //col:73
     */ typedef enum ZydisFormatterStyle_ = 7  //col:74
    ZYDIS_FORMATTER_STYLE_INTEL typedef enum ZydisFormatterStyle_ = 8  //col:75
    /** typedef enum ZydisFormatterStyle_ = 9  //col:76
     * Generates `MASM`-style disassembly that is directly accepted as input for typedef enum ZydisFormatterStyle_ = 10  //col:77
     * the `MASM` assembler. typedef enum ZydisFormatterStyle_ = 11  //col:78
     * typedef enum ZydisFormatterStyle_ = 12  //col:79
     * The runtime-address is ignored in this mode. typedef enum ZydisFormatterStyle_ = 13  //col:80
     */ typedef enum ZydisFormatterStyle_ = 14  //col:81
    ZYDIS_FORMATTER_STYLE_INTEL_MASM typedef enum ZydisFormatterStyle_ = 15  //col:82
    /** typedef enum ZydisFormatterStyle_ = 16  //col:84
     * Maximum value of this enum. typedef enum ZydisFormatterStyle_ = 17  //col:85
     */ typedef enum ZydisFormatterStyle_ = 18  //col:86
    ZYDIS_FORMATTER_STYLE_MAX_VALUE  typedef enum ZydisFormatterStyle_ =  ZYDIS_FORMATTER_STYLE_INTEL_MASM  //col:87
    /** typedef enum ZydisFormatterStyle_ = 20  //col:88
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisFormatterStyle_ = 21  //col:89
     */ typedef enum ZydisFormatterStyle_ = 22  //col:90
    ZYDIS_FORMATTER_STYLE_REQUIRED_BITS  typedef enum ZydisFormatterStyle_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_STYLE_MAX_VALUE)  //col:91
)


type     /* ---------------------------------------------------------------------------------------- */ uint32
const(
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 1  //col:103
    /* General                                                                                  */ typedef enum ZydisFormatterProperty_ = 2  //col:104
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 3  //col:105
    /** typedef enum ZydisFormatterProperty_ = 4  //col:107
     * Controls the printing of effective operand-size suffixes (`AT&T`) or operand-sizes typedef enum ZydisFormatterProperty_ = 5  //col:108
     * of memory operands (`INTEL`). typedef enum ZydisFormatterProperty_ = 6  //col:109
     * typedef enum ZydisFormatterProperty_ = 7  //col:110
     * Pass `ZYAN_TRUE` as value to force the formatter to always print the size or `ZYAN_FALSE` typedef enum ZydisFormatterProperty_ = 8  //col:111
     * to only print it if needed. typedef enum ZydisFormatterProperty_ = 9  //col:112
     */ typedef enum ZydisFormatterProperty_ = 10  //col:113
    ZYDIS_FORMATTER_PROP_FORCE_SIZE typedef enum ZydisFormatterProperty_ = 11  //col:114
    /** typedef enum ZydisFormatterProperty_ = 12  //col:115
     * Controls the printing of segment prefixes. typedef enum ZydisFormatterProperty_ = 13  //col:116
     * typedef enum ZydisFormatterProperty_ = 14  //col:117
     * Pass `ZYAN_TRUE` as value to force the formatter to always print the segment register of typedef enum ZydisFormatterProperty_ = 15  //col:118
     * memory-operands or `ZYAN_FALSE` to omit implicit `DS`/`SS` segments. typedef enum ZydisFormatterProperty_ = 16  //col:119
     */ typedef enum ZydisFormatterProperty_ = 17  //col:120
    ZYDIS_FORMATTER_PROP_FORCE_SEGMENT typedef enum ZydisFormatterProperty_ = 18  //col:121
    /** typedef enum ZydisFormatterProperty_ = 19  //col:122
     * Controls the printing of branch addresses. typedef enum ZydisFormatterProperty_ = 20  //col:123
     * typedef enum ZydisFormatterProperty_ = 21  //col:124
     * Pass `ZYAN_TRUE` as value to force the formatter to always print relative branch addresses typedef enum ZydisFormatterProperty_ = 22  //col:125
     * or `ZYAN_FALSE` to use absolute addresses if a runtime-address different to typedef enum ZydisFormatterProperty_ = 23  //col:126
     * `ZYDIS_RUNTIME_ADDRESS_NONE` was passed. typedef enum ZydisFormatterProperty_ = 24  //col:127
     */ typedef enum ZydisFormatterProperty_ = 25  //col:128
    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_BRANCHES typedef enum ZydisFormatterProperty_ = 26  //col:129
    /** typedef enum ZydisFormatterProperty_ = 27  //col:130
     * Controls the printing of `EIP`/`RIP`-relative addresses. typedef enum ZydisFormatterProperty_ = 28  //col:131
     * typedef enum ZydisFormatterProperty_ = 29  //col:132
     * Pass `ZYAN_TRUE` as value to force the formatter to always print relative addresses for typedef enum ZydisFormatterProperty_ = 30  //col:133
     * `EIP`/`RIP`-relative operands or `ZYAN_FALSE` to use absolute addresses if a runtime- typedef enum ZydisFormatterProperty_ = 31  //col:134
     * address different to `ZYDIS_RUNTIME_ADDRESS_NONE` was passed. typedef enum ZydisFormatterProperty_ = 32  //col:135
     */ typedef enum ZydisFormatterProperty_ = 33  //col:136
    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_RIPREL typedef enum ZydisFormatterProperty_ = 34  //col:137
    /** typedef enum ZydisFormatterProperty_ = 35  //col:138
     * Controls the printing of branch-instructions sizes. typedef enum ZydisFormatterProperty_ = 36  //col:139
     * typedef enum ZydisFormatterProperty_ = 37  //col:140
     * Pass `ZYAN_TRUE` as value to print the size (`short` `near`) of branch typedef enum ZydisFormatterProperty_ = 38  //col:141
     * instructions or `ZYAN_FALSE` to hide it. typedef enum ZydisFormatterProperty_ = 39  //col:142
     * typedef enum ZydisFormatterProperty_ = 40  //col:143
     * Note that the `far`/`l` modifier is always printed. typedef enum ZydisFormatterProperty_ = 41  //col:144
     */ typedef enum ZydisFormatterProperty_ = 42  //col:145
    ZYDIS_FORMATTER_PROP_PRINT_BRANCH_SIZE typedef enum ZydisFormatterProperty_ = 43  //col:146
    /** typedef enum ZydisFormatterProperty_ = 44  //col:148
     * Controls the printing of instruction prefixes. typedef enum ZydisFormatterProperty_ = 45  //col:149
     * typedef enum ZydisFormatterProperty_ = 46  //col:150
     * Pass `ZYAN_TRUE` as value to print all instruction-prefixes (even ignored or duplicate typedef enum ZydisFormatterProperty_ = 47  //col:151
     * ones) or `ZYAN_FALSE` to only print prefixes that are effectively used by the instruction. typedef enum ZydisFormatterProperty_ = 48  //col:152
     */ typedef enum ZydisFormatterProperty_ = 49  //col:153
    ZYDIS_FORMATTER_PROP_DETAILED_PREFIXES typedef enum ZydisFormatterProperty_ = 50  //col:154
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 51  //col:156
    /* Numeric values                                                                           */ typedef enum ZydisFormatterProperty_ = 52  //col:157
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 53  //col:158
    /** typedef enum ZydisFormatterProperty_ = 54  //col:160
     * Controls the base of address values. typedef enum ZydisFormatterProperty_ = 55  //col:161
     */ typedef enum ZydisFormatterProperty_ = 56  //col:162
    ZYDIS_FORMATTER_PROP_ADDR_BASE typedef enum ZydisFormatterProperty_ = 57  //col:163
    /** typedef enum ZydisFormatterProperty_ = 58  //col:164
     * Controls the signedness of relative addresses. Absolute addresses are typedef enum ZydisFormatterProperty_ = 59  //col:165
     * always unsigned. typedef enum ZydisFormatterProperty_ = 60  //col:166
     */ typedef enum ZydisFormatterProperty_ = 61  //col:167
    ZYDIS_FORMATTER_PROP_ADDR_SIGNEDNESS typedef enum ZydisFormatterProperty_ = 62  //col:168
    /** typedef enum ZydisFormatterProperty_ = 63  //col:169
     * Controls the padding of absolute address values. typedef enum ZydisFormatterProperty_ = 64  //col:170
     * typedef enum ZydisFormatterProperty_ = 65  //col:171
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding `ZYDIS_PADDING_AUTO` to padd all typedef enum ZydisFormatterProperty_ = 66  //col:172
     * addresses to the current stack width (hexadecimal only) or any other integer value for typedef enum ZydisFormatterProperty_ = 67  //col:173
     * custom padding. typedef enum ZydisFormatterProperty_ = 68  //col:174
     */ typedef enum ZydisFormatterProperty_ = 69  //col:175
    ZYDIS_FORMATTER_PROP_ADDR_PADDING_ABSOLUTE typedef enum ZydisFormatterProperty_ = 70  //col:176
    /** typedef enum ZydisFormatterProperty_ = 71  //col:177
     * Controls the padding of relative address values. typedef enum ZydisFormatterProperty_ = 72  //col:178
     * typedef enum ZydisFormatterProperty_ = 73  //col:179
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding `ZYDIS_PADDING_AUTO` to padd all typedef enum ZydisFormatterProperty_ = 74  //col:180
     * addresses to the current stack width (hexadecimal only) or any other integer value for typedef enum ZydisFormatterProperty_ = 75  //col:181
     * custom padding. typedef enum ZydisFormatterProperty_ = 76  //col:182
     */ typedef enum ZydisFormatterProperty_ = 77  //col:183
    ZYDIS_FORMATTER_PROP_ADDR_PADDING_RELATIVE typedef enum ZydisFormatterProperty_ = 78  //col:184
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 79  //col:186
    /** typedef enum ZydisFormatterProperty_ = 80  //col:188
     * Controls the base of displacement values. typedef enum ZydisFormatterProperty_ = 81  //col:189
     */ typedef enum ZydisFormatterProperty_ = 82  //col:190
    ZYDIS_FORMATTER_PROP_DISP_BASE typedef enum ZydisFormatterProperty_ = 83  //col:191
    /** typedef enum ZydisFormatterProperty_ = 84  //col:192
     * Controls the signedness of displacement values. typedef enum ZydisFormatterProperty_ = 85  //col:193
     */ typedef enum ZydisFormatterProperty_ = 86  //col:194
    ZYDIS_FORMATTER_PROP_DISP_SIGNEDNESS typedef enum ZydisFormatterProperty_ = 87  //col:195
    /** typedef enum ZydisFormatterProperty_ = 88  //col:196
     * Controls the padding of displacement values. typedef enum ZydisFormatterProperty_ = 89  //col:197
     * typedef enum ZydisFormatterProperty_ = 90  //col:198
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding or any other integer value for custom typedef enum ZydisFormatterProperty_ = 91  //col:199
     * padding. typedef enum ZydisFormatterProperty_ = 92  //col:200
     */ typedef enum ZydisFormatterProperty_ = 93  //col:201
    ZYDIS_FORMATTER_PROP_DISP_PADDING typedef enum ZydisFormatterProperty_ = 94  //col:202
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 95  //col:204
    /** typedef enum ZydisFormatterProperty_ = 96  //col:206
     * Controls the base of immediate values. typedef enum ZydisFormatterProperty_ = 97  //col:207
     */ typedef enum ZydisFormatterProperty_ = 98  //col:208
    ZYDIS_FORMATTER_PROP_IMM_BASE typedef enum ZydisFormatterProperty_ = 99  //col:209
    /** typedef enum ZydisFormatterProperty_ = 100  //col:210
     * Controls the signedness of immediate values. typedef enum ZydisFormatterProperty_ = 101  //col:211
     * typedef enum ZydisFormatterProperty_ = 102  //col:212
     * Pass `ZYDIS_SIGNEDNESS_AUTO` to automatically choose the most suitable mode based on the typedef enum ZydisFormatterProperty_ = 103  //col:213
     * operands `ZydisDecodedOperand.imm.is_signed` attribute. typedef enum ZydisFormatterProperty_ = 104  //col:214
     */ typedef enum ZydisFormatterProperty_ = 105  //col:215
    ZYDIS_FORMATTER_PROP_IMM_SIGNEDNESS typedef enum ZydisFormatterProperty_ = 106  //col:216
    /** typedef enum ZydisFormatterProperty_ = 107  //col:217
     * Controls the padding of immediate values. typedef enum ZydisFormatterProperty_ = 108  //col:218
     * typedef enum ZydisFormatterProperty_ = 109  //col:219
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding `ZYDIS_PADDING_AUTO` to padd all typedef enum ZydisFormatterProperty_ = 110  //col:220
     * immediates to the operand-width (hexadecimal only) or any other integer value for custom typedef enum ZydisFormatterProperty_ = 111  //col:221
     * padding. typedef enum ZydisFormatterProperty_ = 112  //col:222
     */ typedef enum ZydisFormatterProperty_ = 113  //col:223
    ZYDIS_FORMATTER_PROP_IMM_PADDING typedef enum ZydisFormatterProperty_ = 114  //col:224
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 115  //col:226
    /* Text formatting                                                                          */ typedef enum ZydisFormatterProperty_ = 116  //col:227
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 117  //col:228
    /** typedef enum ZydisFormatterProperty_ = 118  //col:230
     * Controls the letter-case for prefixes. typedef enum ZydisFormatterProperty_ = 119  //col:231
     * typedef enum ZydisFormatterProperty_ = 120  //col:232
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase. typedef enum ZydisFormatterProperty_ = 121  //col:233
     */ typedef enum ZydisFormatterProperty_ = 122  //col:234
    ZYDIS_FORMATTER_PROP_UPPERCASE_PREFIXES typedef enum ZydisFormatterProperty_ = 123  //col:235
    /** typedef enum ZydisFormatterProperty_ = 124  //col:236
     * Controls the letter-case for the mnemonic. typedef enum ZydisFormatterProperty_ = 125  //col:237
     * typedef enum ZydisFormatterProperty_ = 126  //col:238
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase. typedef enum ZydisFormatterProperty_ = 127  //col:239
     */ typedef enum ZydisFormatterProperty_ = 128  //col:240
    ZYDIS_FORMATTER_PROP_UPPERCASE_MNEMONIC typedef enum ZydisFormatterProperty_ = 129  //col:241
    /** typedef enum ZydisFormatterProperty_ = 130  //col:242
     * Controls the letter-case for registers. typedef enum ZydisFormatterProperty_ = 131  //col:243
     * typedef enum ZydisFormatterProperty_ = 132  //col:244
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase. typedef enum ZydisFormatterProperty_ = 133  //col:245
     */ typedef enum ZydisFormatterProperty_ = 134  //col:246
    ZYDIS_FORMATTER_PROP_UPPERCASE_REGISTERS typedef enum ZydisFormatterProperty_ = 135  //col:247
    /** typedef enum ZydisFormatterProperty_ = 136  //col:248
     * Controls the letter-case for typecasts. typedef enum ZydisFormatterProperty_ = 137  //col:249
     * typedef enum ZydisFormatterProperty_ = 138  //col:250
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase. typedef enum ZydisFormatterProperty_ = 139  //col:251
     */ typedef enum ZydisFormatterProperty_ = 140  //col:252
    ZYDIS_FORMATTER_PROP_UPPERCASE_TYPECASTS typedef enum ZydisFormatterProperty_ = 141  //col:253
    /** typedef enum ZydisFormatterProperty_ = 142  //col:254
     * Controls the letter-case for decorators. typedef enum ZydisFormatterProperty_ = 143  //col:255
     * typedef enum ZydisFormatterProperty_ = 144  //col:256
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase. typedef enum ZydisFormatterProperty_ = 145  //col:257
     */ typedef enum ZydisFormatterProperty_ = 146  //col:258
    ZYDIS_FORMATTER_PROP_UPPERCASE_DECORATORS typedef enum ZydisFormatterProperty_ = 147  //col:259
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 148  //col:261
    /* Number formatting                                                                        */ typedef enum ZydisFormatterProperty_ = 149  //col:262
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 150  //col:263
    /** typedef enum ZydisFormatterProperty_ = 151  //col:265
     * Controls the prefix for decimal values. typedef enum ZydisFormatterProperty_ = 152  //col:266
     * typedef enum ZydisFormatterProperty_ = 153  //col:267
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters typedef enum ZydisFormatterProperty_ = 154  //col:268
     * to set a custom prefix or `ZYAN_NULL` to disable it. typedef enum ZydisFormatterProperty_ = 155  //col:269
     * typedef enum ZydisFormatterProperty_ = 156  //col:270
     * The string is deep-copied into an internal buffer. typedef enum ZydisFormatterProperty_ = 157  //col:271
     */ typedef enum ZydisFormatterProperty_ = 158  //col:272
    ZYDIS_FORMATTER_PROP_DEC_PREFIX typedef enum ZydisFormatterProperty_ = 159  //col:273
    /** typedef enum ZydisFormatterProperty_ = 160  //col:274
     * Controls the suffix for decimal values. typedef enum ZydisFormatterProperty_ = 161  //col:275
     * typedef enum ZydisFormatterProperty_ = 162  //col:276
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters typedef enum ZydisFormatterProperty_ = 163  //col:277
     * to set a custom suffix or `ZYAN_NULL` to disable it. typedef enum ZydisFormatterProperty_ = 164  //col:278
     * typedef enum ZydisFormatterProperty_ = 165  //col:279
     * The string is deep-copied into an internal buffer. typedef enum ZydisFormatterProperty_ = 166  //col:280
     */ typedef enum ZydisFormatterProperty_ = 167  //col:281
    ZYDIS_FORMATTER_PROP_DEC_SUFFIX typedef enum ZydisFormatterProperty_ = 168  //col:282
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 169  //col:284
    /** typedef enum ZydisFormatterProperty_ = 170  //col:286
     * Controls the letter-case of hexadecimal values. typedef enum ZydisFormatterProperty_ = 171  //col:287
     * typedef enum ZydisFormatterProperty_ = 172  //col:288
     * Pass `ZYAN_TRUE` as value to format in uppercase and `ZYAN_FALSE` to format in lowercase. typedef enum ZydisFormatterProperty_ = 173  //col:289
     * typedef enum ZydisFormatterProperty_ = 174  //col:290
     * The default value is `ZYAN_TRUE`. typedef enum ZydisFormatterProperty_ = 175  //col:291
     */ typedef enum ZydisFormatterProperty_ = 176  //col:292
    ZYDIS_FORMATTER_PROP_HEX_UPPERCASE typedef enum ZydisFormatterProperty_ = 177  //col:293
    /** typedef enum ZydisFormatterProperty_ = 178  //col:294
     * Controls the prefix for hexadecimal values. typedef enum ZydisFormatterProperty_ = 179  //col:295
     * typedef enum ZydisFormatterProperty_ = 180  //col:296
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters typedef enum ZydisFormatterProperty_ = 181  //col:297
     * to set a custom prefix or `ZYAN_NULL` to disable it. typedef enum ZydisFormatterProperty_ = 182  //col:298
     * typedef enum ZydisFormatterProperty_ = 183  //col:299
     * The string is deep-copied into an internal buffer. typedef enum ZydisFormatterProperty_ = 184  //col:300
     */ typedef enum ZydisFormatterProperty_ = 185  //col:301
    ZYDIS_FORMATTER_PROP_HEX_PREFIX typedef enum ZydisFormatterProperty_ = 186  //col:302
    /** typedef enum ZydisFormatterProperty_ = 187  //col:303
     * Controls the suffix for hexadecimal values. typedef enum ZydisFormatterProperty_ = 188  //col:304
     * typedef enum ZydisFormatterProperty_ = 189  //col:305
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters typedef enum ZydisFormatterProperty_ = 190  //col:306
     * to set a custom suffix or `ZYAN_NULL` to disable it. typedef enum ZydisFormatterProperty_ = 191  //col:307
     * typedef enum ZydisFormatterProperty_ = 192  //col:308
     * The string is deep-copied into an internal buffer. typedef enum ZydisFormatterProperty_ = 193  //col:309
     */ typedef enum ZydisFormatterProperty_ = 194  //col:310
    ZYDIS_FORMATTER_PROP_HEX_SUFFIX typedef enum ZydisFormatterProperty_ = 195  //col:311
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterProperty_ = 196  //col:313
    /** typedef enum ZydisFormatterProperty_ = 197  //col:315
     * Maximum value of this enum. typedef enum ZydisFormatterProperty_ = 198  //col:316
     */ typedef enum ZydisFormatterProperty_ = 199  //col:317
    ZYDIS_FORMATTER_PROP_MAX_VALUE  typedef enum ZydisFormatterProperty_ =  ZYDIS_FORMATTER_PROP_HEX_SUFFIX  //col:318
    /** typedef enum ZydisFormatterProperty_ = 201  //col:319
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisFormatterProperty_ = 202  //col:320
     */ typedef enum ZydisFormatterProperty_ = 203  //col:321
    ZYDIS_FORMATTER_PROP_REQUIRED_BITS  typedef enum ZydisFormatterProperty_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_PROP_MAX_VALUE)  //col:322
)


type     /** uint32
const(
    /** typedef enum ZydisNumericBase_ = 1  //col:332
     * Decimal system. typedef enum ZydisNumericBase_ = 2  //col:333
     */ typedef enum ZydisNumericBase_ = 3  //col:334
    ZYDIS_NUMERIC_BASE_DEC typedef enum ZydisNumericBase_ = 4  //col:335
    /** typedef enum ZydisNumericBase_ = 5  //col:336
     * Hexadecimal system. typedef enum ZydisNumericBase_ = 6  //col:337
     */ typedef enum ZydisNumericBase_ = 7  //col:338
    ZYDIS_NUMERIC_BASE_HEX typedef enum ZydisNumericBase_ = 8  //col:339
    /** typedef enum ZydisNumericBase_ = 9  //col:341
     * Maximum value of this enum. typedef enum ZydisNumericBase_ = 10  //col:342
     */ typedef enum ZydisNumericBase_ = 11  //col:343
    ZYDIS_NUMERIC_BASE_MAX_VALUE  typedef enum ZydisNumericBase_ =  ZYDIS_NUMERIC_BASE_HEX  //col:344
    /** typedef enum ZydisNumericBase_ = 13  //col:345
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisNumericBase_ = 14  //col:346
     */ typedef enum ZydisNumericBase_ = 15  //col:347
    ZYDIS_NUMERIC_BASE_REQUIRED_BITS  typedef enum ZydisNumericBase_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_NUMERIC_BASE_MAX_VALUE)  //col:348
)


type     /** uint32
const(
    /** typedef enum ZydisSignedness_ = 1  //col:358
     * Automatically choose the most suitable mode based on the operands typedef enum ZydisSignedness_ = 2  //col:359
     * ZydisDecodedOperand.imm.is_signed` attribute. typedef enum ZydisSignedness_ = 3  //col:360
     */ typedef enum ZydisSignedness_ = 4  //col:361
    ZYDIS_SIGNEDNESS_AUTO typedef enum ZydisSignedness_ = 5  //col:362
    /** typedef enum ZydisSignedness_ = 6  //col:363
     * Force signed values. typedef enum ZydisSignedness_ = 7  //col:364
     */ typedef enum ZydisSignedness_ = 8  //col:365
    ZYDIS_SIGNEDNESS_SIGNED typedef enum ZydisSignedness_ = 9  //col:366
    /** typedef enum ZydisSignedness_ = 10  //col:367
     * Force unsigned values. typedef enum ZydisSignedness_ = 11  //col:368
     */ typedef enum ZydisSignedness_ = 12  //col:369
    ZYDIS_SIGNEDNESS_UNSIGNED typedef enum ZydisSignedness_ = 13  //col:370
    /** typedef enum ZydisSignedness_ = 14  //col:372
     * Maximum value of this enum. typedef enum ZydisSignedness_ = 15  //col:373
     */ typedef enum ZydisSignedness_ = 16  //col:374
    ZYDIS_SIGNEDNESS_MAX_VALUE  typedef enum ZydisSignedness_ =  ZYDIS_SIGNEDNESS_UNSIGNED  //col:375
    /** typedef enum ZydisSignedness_ = 18  //col:376
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisSignedness_ = 19  //col:377
     */ typedef enum ZydisSignedness_ = 20  //col:378
    ZYDIS_SIGNEDNESS_REQUIRED_BITS  typedef enum ZydisSignedness_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_SIGNEDNESS_MAX_VALUE)  //col:379
)


type     /** uint32
const(
    /** typedef enum ZydisPadding_ = 1  //col:389
     * Disables padding. typedef enum ZydisPadding_ = 2  //col:390
     */ typedef enum ZydisPadding_ = 3  //col:391
    ZYDIS_PADDING_DISABLED  typedef enum ZydisPadding_ =  0  //col:392
    /** typedef enum ZydisPadding_ = 5  //col:393
     * Padds the value to the current stack-width for addresses or to the typedef enum ZydisPadding_ = 6  //col:394
     * operand-width for immediate values (hexadecimal only). typedef enum ZydisPadding_ = 7  //col:395
     */ typedef enum ZydisPadding_ = 8  //col:396
    ZYDIS_PADDING_AUTO      typedef enum ZydisPadding_ =  (-1)  //col:397
    /** typedef enum ZydisPadding_ = 10  //col:399
     * Maximum value of this enum. typedef enum ZydisPadding_ = 11  //col:400
     */ typedef enum ZydisPadding_ = 12  //col:401
    ZYDIS_PADDING_MAX_VALUE  typedef enum ZydisPadding_ =  ZYDIS_PADDING_AUTO  //col:402
    /** typedef enum ZydisPadding_ = 14  //col:403
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisPadding_ = 15  //col:404
     */ typedef enum ZydisPadding_ = 16  //col:405
    ZYDIS_PADDING_REQUIRED_BITS  typedef enum ZydisPadding_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_PADDING_MAX_VALUE)  //col:406
)


type     /* ---------------------------------------------------------------------------------------- */ uint32
const(
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 1  //col:421
    /* Instruction                                                                              */ typedef enum ZydisFormatterFunction_ = 2  //col:422
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 3  //col:423
    /** typedef enum ZydisFormatterFunction_ = 4  //col:425
     * This function is invoked before the formatter formats an instruction. typedef enum ZydisFormatterFunction_ = 5  //col:426
     */ typedef enum ZydisFormatterFunction_ = 6  //col:427
    ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION typedef enum ZydisFormatterFunction_ = 7  //col:428
    /** typedef enum ZydisFormatterFunction_ = 8  //col:429
     * This function is invoked after the formatter formatted an instruction. typedef enum ZydisFormatterFunction_ = 9  //col:430
     */ typedef enum ZydisFormatterFunction_ = 10  //col:431
    ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION typedef enum ZydisFormatterFunction_ = 11  //col:432
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 12  //col:434
    /** typedef enum ZydisFormatterFunction_ = 13  //col:436
     * This function refers to the main formatting function. typedef enum ZydisFormatterFunction_ = 14  //col:437
     * typedef enum ZydisFormatterFunction_ = 15  //col:438
     * Replacing this function allows for complete custom formatting but indirectly disables all typedef enum ZydisFormatterFunction_ = 16  //col:439
     * other hooks except for `ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION` and typedef enum ZydisFormatterFunction_ = 17  //col:440
     * `ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION`. typedef enum ZydisFormatterFunction_ = 18  //col:441
     */ typedef enum ZydisFormatterFunction_ = 19  //col:442
    ZYDIS_FORMATTER_FUNC_FORMAT_INSTRUCTION typedef enum ZydisFormatterFunction_ = 20  //col:443
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 21  //col:445
    /* Operands                                                                                 */ typedef enum ZydisFormatterFunction_ = 22  //col:446
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 23  //col:447
    /** typedef enum ZydisFormatterFunction_ = 24  //col:449
     * This function is invoked before the formatter formats an operand. typedef enum ZydisFormatterFunction_ = 25  //col:450
     */ typedef enum ZydisFormatterFunction_ = 26  //col:451
    ZYDIS_FORMATTER_FUNC_PRE_OPERAND typedef enum ZydisFormatterFunction_ = 27  //col:452
    /** typedef enum ZydisFormatterFunction_ = 28  //col:453
     * This function is invoked after the formatter formatted an operand. typedef enum ZydisFormatterFunction_ = 29  //col:454
     */ typedef enum ZydisFormatterFunction_ = 30  //col:455
    ZYDIS_FORMATTER_FUNC_POST_OPERAND typedef enum ZydisFormatterFunction_ = 31  //col:456
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 32  //col:458
    /** typedef enum ZydisFormatterFunction_ = 33  //col:460
     * This function is invoked to format a register operand. typedef enum ZydisFormatterFunction_ = 34  //col:461
     */ typedef enum ZydisFormatterFunction_ = 35  //col:462
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG typedef enum ZydisFormatterFunction_ = 36  //col:463
    /** typedef enum ZydisFormatterFunction_ = 37  //col:464
     * This function is invoked to format a memory operand. typedef enum ZydisFormatterFunction_ = 38  //col:465
     * typedef enum ZydisFormatterFunction_ = 39  //col:466
     * Replacing this function might indirectly disable some specific calls to the typedef enum ZydisFormatterFunction_ = 40  //col:467
     * `ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST` `ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT` typedef enum ZydisFormatterFunction_ = 41  //col:468
     * `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` and `ZYDIS_FORMATTER_FUNC_PRINT_DISP` functions. typedef enum ZydisFormatterFunction_ = 42  //col:469
     */ typedef enum ZydisFormatterFunction_ = 43  //col:470
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM typedef enum ZydisFormatterFunction_ = 44  //col:471
    /** typedef enum ZydisFormatterFunction_ = 45  //col:472
     * This function is invoked to format a pointer operand. typedef enum ZydisFormatterFunction_ = 46  //col:473
     */ typedef enum ZydisFormatterFunction_ = 47  //col:474
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR typedef enum ZydisFormatterFunction_ = 48  //col:475
    /** typedef enum ZydisFormatterFunction_ = 49  //col:476
     * This function is invoked to format an immediate operand. typedef enum ZydisFormatterFunction_ = 50  //col:477
     * typedef enum ZydisFormatterFunction_ = 51  //col:478
     * Replacing this function might indirectly disable some specific calls to the typedef enum ZydisFormatterFunction_ = 52  //col:479
     * `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL` and typedef enum ZydisFormatterFunction_ = 53  //col:480
     * `ZYDIS_FORMATTER_FUNC_PRINT_IMM` functions. typedef enum ZydisFormatterFunction_ = 54  //col:481
     */ typedef enum ZydisFormatterFunction_ = 55  //col:482
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM typedef enum ZydisFormatterFunction_ = 56  //col:483
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 57  //col:485
    /* Elemental tokens                                                                         */ typedef enum ZydisFormatterFunction_ = 58  //col:486
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 59  //col:487
    /** typedef enum ZydisFormatterFunction_ = 60  //col:489
     * This function is invoked to print the instruction mnemonic. typedef enum ZydisFormatterFunction_ = 61  //col:490
     */ typedef enum ZydisFormatterFunction_ = 62  //col:491
    ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC typedef enum ZydisFormatterFunction_ = 63  //col:492
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 64  //col:494
    /** typedef enum ZydisFormatterFunction_ = 65  //col:496
     * This function is invoked to print a register. typedef enum ZydisFormatterFunction_ = 66  //col:497
     */ typedef enum ZydisFormatterFunction_ = 67  //col:498
    ZYDIS_FORMATTER_FUNC_PRINT_REGISTER typedef enum ZydisFormatterFunction_ = 68  //col:499
    /** typedef enum ZydisFormatterFunction_ = 69  //col:500
     * This function is invoked to print absolute addresses. typedef enum ZydisFormatterFunction_ = 70  //col:501
     * typedef enum ZydisFormatterFunction_ = 71  //col:502
     * Conditionally invoked if a runtime-address different to `ZYDIS_RUNTIME_ADDRESS_NONE` was typedef enum ZydisFormatterFunction_ = 72  //col:503
     * passed: typedef enum ZydisFormatterFunction_ = 73  //col:504
     * - `IMM` operands with relative address (e.g. `JMP` `CALL` ...) typedef enum ZydisFormatterFunction_ = 74  //col:505
     * - `MEM` operands with `EIP`/`RIP`-relative address (e.g. `MOV RAX [RIP+0x12345678]`) typedef enum ZydisFormatterFunction_ = 75  //col:506
     * typedef enum ZydisFormatterFunction_ = 76  //col:507
     * Always invoked for: typedef enum ZydisFormatterFunction_ = 77  //col:508
     * - `MEM` operands with absolute address (e.g. `MOV RAX [0x12345678]`) typedef enum ZydisFormatterFunction_ = 78  //col:509
     */ typedef enum ZydisFormatterFunction_ = 79  //col:510
    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS typedef enum ZydisFormatterFunction_ = 80  //col:511
    /** typedef enum ZydisFormatterFunction_ = 81  //col:512
     * This function is invoked to print relative addresses. typedef enum ZydisFormatterFunction_ = 82  //col:513
     * typedef enum ZydisFormatterFunction_ = 83  //col:514
     * Conditionally invoked if `ZYDIS_RUNTIME_ADDRESS_NONE` was passed as runtime-address: typedef enum ZydisFormatterFunction_ = 84  //col:515
     * - `IMM` operands with relative address (e.g. `JMP` `CALL` ...) typedef enum ZydisFormatterFunction_ = 85  //col:516
     */ typedef enum ZydisFormatterFunction_ = 86  //col:517
    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL typedef enum ZydisFormatterFunction_ = 87  //col:518
    /** typedef enum ZydisFormatterFunction_ = 88  //col:519
     * This function is invoked to print a memory displacement value. typedef enum ZydisFormatterFunction_ = 89  //col:520
     * typedef enum ZydisFormatterFunction_ = 90  //col:521
     * If the memory displacement contains an address and a runtime-address different to typedef enum ZydisFormatterFunction_ = 91  //col:522
     * `ZYDIS_RUNTIME_ADDRESS_NONE` was passed `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` is called typedef enum ZydisFormatterFunction_ = 92  //col:523
     * instead. typedef enum ZydisFormatterFunction_ = 93  //col:524
     */ typedef enum ZydisFormatterFunction_ = 94  //col:525
    ZYDIS_FORMATTER_FUNC_PRINT_DISP typedef enum ZydisFormatterFunction_ = 95  //col:526
    /** typedef enum ZydisFormatterFunction_ = 96  //col:527
     * This function is invoked to print an immediate value. typedef enum ZydisFormatterFunction_ = 97  //col:528
     * typedef enum ZydisFormatterFunction_ = 98  //col:529
     * If the immediate contains an address and a runtime-address different to typedef enum ZydisFormatterFunction_ = 99  //col:530
     * `ZYDIS_RUNTIME_ADDRESS_NONE` was passed `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` is called typedef enum ZydisFormatterFunction_ = 100  //col:531
     * instead. typedef enum ZydisFormatterFunction_ = 101  //col:532
     * typedef enum ZydisFormatterFunction_ = 102  //col:533
     * If the immediate contains an address and `ZYDIS_RUNTIME_ADDRESS_NONE` was passed as typedef enum ZydisFormatterFunction_ = 103  //col:534
     * runtime-address `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL` is called instead. typedef enum ZydisFormatterFunction_ = 104  //col:535
     */ typedef enum ZydisFormatterFunction_ = 105  //col:536
    ZYDIS_FORMATTER_FUNC_PRINT_IMM typedef enum ZydisFormatterFunction_ = 106  //col:537
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 107  //col:539
    /* Optional tokens                                                                          */ typedef enum ZydisFormatterFunction_ = 108  //col:540
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 109  //col:541
    /** typedef enum ZydisFormatterFunction_ = 110  //col:543
     * This function is invoked to print the size of a memory operand (`INTEL` only). typedef enum ZydisFormatterFunction_ = 111  //col:544
     */ typedef enum ZydisFormatterFunction_ = 112  //col:545
    ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST typedef enum ZydisFormatterFunction_ = 113  //col:546
    /** typedef enum ZydisFormatterFunction_ = 114  //col:547
     * This function is invoked to print the segment-register of a memory operand. typedef enum ZydisFormatterFunction_ = 115  //col:548
     */ typedef enum ZydisFormatterFunction_ = 116  //col:549
    ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT typedef enum ZydisFormatterFunction_ = 117  //col:550
    /** typedef enum ZydisFormatterFunction_ = 118  //col:551
     * This function is invoked to print the instruction prefixes. typedef enum ZydisFormatterFunction_ = 119  //col:552
     */ typedef enum ZydisFormatterFunction_ = 120  //col:553
    ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES typedef enum ZydisFormatterFunction_ = 121  //col:554
    /** typedef enum ZydisFormatterFunction_ = 122  //col:555
     * This function is invoked after formatting an operand to print a `EVEX`/`MVEX` typedef enum ZydisFormatterFunction_ = 123  //col:556
     * decorator. typedef enum ZydisFormatterFunction_ = 124  //col:557
     */ typedef enum ZydisFormatterFunction_ = 125  //col:558
    ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR typedef enum ZydisFormatterFunction_ = 126  //col:559
    /* ---------------------------------------------------------------------------------------- */ typedef enum ZydisFormatterFunction_ = 127  //col:561
    /** typedef enum ZydisFormatterFunction_ = 128  //col:563
     * Maximum value of this enum. typedef enum ZydisFormatterFunction_ = 129  //col:564
     */ typedef enum ZydisFormatterFunction_ = 130  //col:565
    ZYDIS_FORMATTER_FUNC_MAX_VALUE  typedef enum ZydisFormatterFunction_ =  ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR  //col:566
    /** typedef enum ZydisFormatterFunction_ = 132  //col:567
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisFormatterFunction_ = 133  //col:568
     */ typedef enum ZydisFormatterFunction_ = 134  //col:569
    ZYDIS_FORMATTER_FUNC_REQUIRED_BITS  typedef enum ZydisFormatterFunction_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_FUNC_MAX_VALUE)  //col:570
)


type     ZYDIS_DECORATOR_INVALID uint32
const(
    ZYDIS_DECORATOR_INVALID typedef enum ZydisDecorator_ = 1  //col:582
    /** typedef enum ZydisDecorator_ = 2  //col:583
     * The embedded-mask decorator. typedef enum ZydisDecorator_ = 3  //col:584
     */ typedef enum ZydisDecorator_ = 4  //col:585
    ZYDIS_DECORATOR_MASK typedef enum ZydisDecorator_ = 5  //col:586
    /** typedef enum ZydisDecorator_ = 6  //col:587
     * The broadcast decorator. typedef enum ZydisDecorator_ = 7  //col:588
     */ typedef enum ZydisDecorator_ = 8  //col:589
    ZYDIS_DECORATOR_BC typedef enum ZydisDecorator_ = 9  //col:590
    /** typedef enum ZydisDecorator_ = 10  //col:591
     * The rounding-control decorator. typedef enum ZydisDecorator_ = 11  //col:592
     */ typedef enum ZydisDecorator_ = 12  //col:593
    ZYDIS_DECORATOR_RC typedef enum ZydisDecorator_ = 13  //col:594
    /** typedef enum ZydisDecorator_ = 14  //col:595
     * The suppress-all-exceptions decorator. typedef enum ZydisDecorator_ = 15  //col:596
     */ typedef enum ZydisDecorator_ = 16  //col:597
    ZYDIS_DECORATOR_SAE typedef enum ZydisDecorator_ = 17  //col:598
    /** typedef enum ZydisDecorator_ = 18  //col:599
     * The register-swizzle decorator. typedef enum ZydisDecorator_ = 19  //col:600
     */ typedef enum ZydisDecorator_ = 20  //col:601
    ZYDIS_DECORATOR_SWIZZLE typedef enum ZydisDecorator_ = 21  //col:602
    /** typedef enum ZydisDecorator_ = 22  //col:603
     * The conversion decorator. typedef enum ZydisDecorator_ = 23  //col:604
     */ typedef enum ZydisDecorator_ = 24  //col:605
    ZYDIS_DECORATOR_CONVERSION typedef enum ZydisDecorator_ = 25  //col:606
    /** typedef enum ZydisDecorator_ = 26  //col:607
     * The eviction-hint decorator. typedef enum ZydisDecorator_ = 27  //col:608
     */ typedef enum ZydisDecorator_ = 28  //col:609
    ZYDIS_DECORATOR_EH typedef enum ZydisDecorator_ = 29  //col:610
    /** typedef enum ZydisDecorator_ = 30  //col:612
     * Maximum value of this enum. typedef enum ZydisDecorator_ = 31  //col:613
     */ typedef enum ZydisDecorator_ = 32  //col:614
    ZYDIS_DECORATOR_MAX_VALUE  typedef enum ZydisDecorator_ =  ZYDIS_DECORATOR_EH  //col:615
    /** typedef enum ZydisDecorator_ = 34  //col:616
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisDecorator_ = 35  //col:617
     */ typedef enum ZydisDecorator_ = 36  //col:618
    ZYDIS_DECORATOR_REQUIRED_BITS  typedef enum ZydisDecorator_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_DECORATOR_MAX_VALUE)  //col:619
)



type (
Formatter interface{
  Zyan Disassembler Library ()(ok bool)//col:92
     * Controls the printing of effective operand-size suffixes ()(ok bool)//col:323
    ZYDIS_NUMERIC_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:349
    ZYDIS_SIGNEDNESS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:380
     * operand-width for immediate values ()(ok bool)//col:407
     * - `IMM` operands with relative address ()(ok bool)//col:571
    ZYDIS_DECORATOR_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:620
 * process to fail ()(ok bool)//col:936
ZYDIS_EXPORT ZyanStatus ZydisFormatterInit()(ok bool)//col:1176
}
)

func NewFormatter() { return & formatter{} }

func (f *formatter)  Zyan Disassembler Library ()(ok bool){//col:92
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
 * Functions for formatting instructions to human-readable text.
#ifndef ZYDIS_FORMATTER_H
#define ZYDIS_FORMATTER_H
#include <Zycore/Defines.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/FormatterBuffer.h>
#ifdef __cplusplus
extern "C" {
#endif
 * Use this constant as value for `runtime_address` in `ZydisFormatterFormatInstruction(Ex)`
 * or `ZydisFormatterFormatOperand(Ex)` to print relative values for all addresses.
#define ZYDIS_RUNTIME_ADDRESS_NONE (ZyanU64)(-1)
 * Defines the `ZydisFormatterStyle` enum.
typedef enum ZydisFormatterStyle_
{
     * Generates `AT&T`-style disassembly.
    ZYDIS_FORMATTER_STYLE_ATT,
     * Generates `Intel`-style disassembly.
    ZYDIS_FORMATTER_STYLE_INTEL,
     * Generates `MASM`-style disassembly that is directly accepted as input for
     * the `MASM` assembler.
     *
     * The runtime-address is ignored in this mode.
    ZYDIS_FORMATTER_STYLE_INTEL_MASM,
     * Maximum value of this enum.
    ZYDIS_FORMATTER_STYLE_MAX_VALUE = ZYDIS_FORMATTER_STYLE_INTEL_MASM,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_FORMATTER_STYLE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_STYLE_MAX_VALUE)
} ZydisFormatterStyle;*/
return true
}

func (f *formatter)     * Controls the printing of effective operand-size suffixes ()(ok bool){//col:323
/*     * Controls the printing of effective operand-size suffixes (`AT&T`) or operand-sizes
     * of memory operands (`INTEL`).
     *
     * Pass `ZYAN_TRUE` as value to force the formatter to always print the size, or `ZYAN_FALSE`
     * to only print it if needed.
    ZYDIS_FORMATTER_PROP_FORCE_SIZE,
     * Controls the printing of segment prefixes.
     *
     * Pass `ZYAN_TRUE` as value to force the formatter to always print the segment register of
     * memory-operands or `ZYAN_FALSE` to omit implicit `DS`/`SS` segments.
    ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
     * Controls the printing of branch addresses.
     *
     * Pass `ZYAN_TRUE` as value to force the formatter to always print relative branch addresses
     * or `ZYAN_FALSE` to use absolute addresses, if a runtime-address different to
     * `ZYDIS_RUNTIME_ADDRESS_NONE` was passed.
    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_BRANCHES,
     * Controls the printing of `EIP`/`RIP`-relative addresses.
     *
     * Pass `ZYAN_TRUE` as value to force the formatter to always print relative addresses for
     * `EIP`/`RIP`-relative operands or `ZYAN_FALSE` to use absolute addresses, if a runtime-
     * address different to `ZYDIS_RUNTIME_ADDRESS_NONE` was passed.
    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_RIPREL,
     * Controls the printing of branch-instructions sizes.
     *
     * Pass `ZYAN_TRUE` as value to print the size (`short`, `near`) of branch
     * instructions or `ZYAN_FALSE` to hide it.
     *
     * Note that the `far`/`l` modifier is always printed.
    ZYDIS_FORMATTER_PROP_PRINT_BRANCH_SIZE,
     * Controls the printing of instruction prefixes.
     *
     * Pass `ZYAN_TRUE` as value to print all instruction-prefixes (even ignored or duplicate
     * ones) or `ZYAN_FALSE` to only print prefixes that are effectively used by the instruction.
    ZYDIS_FORMATTER_PROP_DETAILED_PREFIXES,
     * Controls the base of address values.
    ZYDIS_FORMATTER_PROP_ADDR_BASE,
     * Controls the signedness of relative addresses. Absolute addresses are
     * always unsigned.
    ZYDIS_FORMATTER_PROP_ADDR_SIGNEDNESS,
     * Controls the padding of absolute address values.
     *
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding, `ZYDIS_PADDING_AUTO` to padd all
     * addresses to the current stack width (hexadecimal only), or any other integer value for
     * custom padding.
    ZYDIS_FORMATTER_PROP_ADDR_PADDING_ABSOLUTE,
     * Controls the padding of relative address values.
     *
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding, `ZYDIS_PADDING_AUTO` to padd all
     * addresses to the current stack width (hexadecimal only), or any other integer value for
     * custom padding.
    ZYDIS_FORMATTER_PROP_ADDR_PADDING_RELATIVE,
     * Controls the base of displacement values.
    ZYDIS_FORMATTER_PROP_DISP_BASE,
     * Controls the signedness of displacement values.
    ZYDIS_FORMATTER_PROP_DISP_SIGNEDNESS,
     * Controls the padding of displacement values.
     *
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding, or any other integer value for custom
     * padding.
    ZYDIS_FORMATTER_PROP_DISP_PADDING,
     * Controls the base of immediate values.
    ZYDIS_FORMATTER_PROP_IMM_BASE,
     * Controls the signedness of immediate values.
     *
     * Pass `ZYDIS_SIGNEDNESS_AUTO` to automatically choose the most suitable mode based on the
     * operands `ZydisDecodedOperand.imm.is_signed` attribute.
    ZYDIS_FORMATTER_PROP_IMM_SIGNEDNESS,
     * Controls the padding of immediate values.
     *
     * Pass `ZYDIS_PADDING_DISABLED` to disable padding, `ZYDIS_PADDING_AUTO` to padd all
     * immediates to the operand-width (hexadecimal only), or any other integer value for custom
     * padding.
    ZYDIS_FORMATTER_PROP_IMM_PADDING,
     * Controls the letter-case for prefixes.
     *
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase.
    ZYDIS_FORMATTER_PROP_UPPERCASE_PREFIXES,
     * Controls the letter-case for the mnemonic.
     *
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase.
    ZYDIS_FORMATTER_PROP_UPPERCASE_MNEMONIC,
     * Controls the letter-case for registers.
     *
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase.
    ZYDIS_FORMATTER_PROP_UPPERCASE_REGISTERS,
     * Controls the letter-case for typecasts.
     *
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase.
    ZYDIS_FORMATTER_PROP_UPPERCASE_TYPECASTS,
     * Controls the letter-case for decorators.
     *
     * Pass `ZYAN_TRUE` as value to format in uppercase or `ZYAN_FALSE` to format in lowercase.
    ZYDIS_FORMATTER_PROP_UPPERCASE_DECORATORS,
     * Controls the prefix for decimal values.
     *
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters
     * to set a custom prefix, or `ZYAN_NULL` to disable it.
     *
     * The string is deep-copied into an internal buffer.
    ZYDIS_FORMATTER_PROP_DEC_PREFIX,
     * Controls the suffix for decimal values.
     *
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters
     * to set a custom suffix, or `ZYAN_NULL` to disable it.
     *
     * The string is deep-copied into an internal buffer.
    ZYDIS_FORMATTER_PROP_DEC_SUFFIX,
     * Controls the letter-case of hexadecimal values.
     *
     * Pass `ZYAN_TRUE` as value to format in uppercase and `ZYAN_FALSE` to format in lowercase.
     *
     * The default value is `ZYAN_TRUE`.
    ZYDIS_FORMATTER_PROP_HEX_UPPERCASE,
     * Controls the prefix for hexadecimal values.
     *
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters
     * to set a custom prefix, or `ZYAN_NULL` to disable it.
     *
     * The string is deep-copied into an internal buffer.
    ZYDIS_FORMATTER_PROP_HEX_PREFIX,
     * Controls the suffix for hexadecimal values.
     *
     * Pass a pointer to a null-terminated C-style string with a maximum length of 10 characters
     * to set a custom suffix, or `ZYAN_NULL` to disable it.
     *
     * The string is deep-copied into an internal buffer.
    ZYDIS_FORMATTER_PROP_HEX_SUFFIX,
     * Maximum value of this enum.
    ZYDIS_FORMATTER_PROP_MAX_VALUE = ZYDIS_FORMATTER_PROP_HEX_SUFFIX,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_FORMATTER_PROP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_PROP_MAX_VALUE)
} ZydisFormatterProperty;*/
return true
}

func (f *formatter)    ZYDIS_NUMERIC_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:349
/*    ZYDIS_NUMERIC_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_NUMERIC_BASE_MAX_VALUE)
} ZydisNumericBase;*/
return true
}

func (f *formatter)    ZYDIS_SIGNEDNESS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:380
/*    ZYDIS_SIGNEDNESS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_SIGNEDNESS_MAX_VALUE)
} ZydisSignedness;*/
return true
}

func (f *formatter)     * operand-width for immediate values ()(ok bool){//col:407
/*     * operand-width for immediate values (hexadecimal only).
    ZYDIS_PADDING_AUTO     = (-1),
     * Maximum value of this enum.
    ZYDIS_PADDING_MAX_VALUE = ZYDIS_PADDING_AUTO,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_PADDING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_PADDING_MAX_VALUE)
} ZydisPadding;*/
return true
}

func (f *formatter)     * - `IMM` operands with relative address ()(ok bool){//col:571
/*     * - `IMM` operands with relative address (e.g. `JMP`, `CALL`, ...)
     * - `MEM` operands with `EIP`/`RIP`-relative address (e.g. `MOV RAX, [RIP+0x12345678]`)
     *
     * Always invoked for:
     * - `MEM` operands with absolute address (e.g. `MOV RAX, [0x12345678]`)
    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
     * This function is invoked to print relative addresses.
     *
     * Conditionally invoked, if `ZYDIS_RUNTIME_ADDRESS_NONE` was passed as runtime-address:
     * - `IMM` operands with relative address (e.g. `JMP`, `CALL`, ...)
    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL,
     * This function is invoked to print a memory displacement value.
     *
     * If the memory displacement contains an address and a runtime-address different to
     * `ZYDIS_RUNTIME_ADDRESS_NONE` was passed, `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` is called
     * instead.
    ZYDIS_FORMATTER_FUNC_PRINT_DISP,
     * This function is invoked to print an immediate value.
     *
     * If the immediate contains an address and a runtime-address different to
     * `ZYDIS_RUNTIME_ADDRESS_NONE` was passed, `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` is called
     * instead.
     *
     * If the immediate contains an address and `ZYDIS_RUNTIME_ADDRESS_NONE` was passed as
     * runtime-address, `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL` is called instead.
    ZYDIS_FORMATTER_FUNC_PRINT_IMM,
     * This function is invoked to print the size of a memory operand (`INTEL` only).
    ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST,
     * This function is invoked to print the segment-register of a memory operand.
    ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT,
     * This function is invoked to print the instruction prefixes.
    ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES,
     * This function is invoked after formatting an operand to print a `EVEX`/`MVEX`
     * decorator.
    ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR,
     * Maximum value of this enum.
    ZYDIS_FORMATTER_FUNC_MAX_VALUE = ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_FORMATTER_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_FUNC_MAX_VALUE)
} ZydisFormatterFunction;*/
return true
}

func (f *formatter)    ZYDIS_DECORATOR_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:620
/*    ZYDIS_DECORATOR_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_DECORATOR_MAX_VALUE)
} ZydisDecorator;*/
return true
}

func (f *formatter) * process to fail ()(ok bool){//col:936
/* * process to fail (see exceptions below).
 *
 * Returning `ZYDIS_STATUS_SKIP_TOKEN` is valid for functions of the following types and will
 * instruct the formatter to omit the whole operand:
 * - `ZYDIS_FORMATTER_FUNC_PRE_OPERAND`
 * - `ZYDIS_FORMATTER_FUNC_POST_OPERAND`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM`
 *
 * This function prototype is used by functions of the following types:
 * - `ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION`
 * - `ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION`
 * - `ZYDIS_FORMATTER_FUNC_PRE_OPERAND`
 * - `ZYDIS_FORMATTER_FUNC_POST_OPERAND`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_INSTRUCTION`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR`
 * - `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_DISP`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_IMM`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST`
 * - `ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT`
typedef ZyanStatus (*ZydisFormatterFunc)(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
 * Defines the `ZydisFormatterRegisterFunc` function prototype.
 *
 *
 *          formatting process to fail.
 *
 * This function prototype is used by functions of the following types:
 * - `ZYDIS_FORMATTER_FUNC_PRINT_REGISTER`.
typedef ZyanStatus (*ZydisFormatterRegisterFunc)(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg);
 * Defines the `ZydisFormatterDecoratorFunc` function prototype.
 *
 *
 *          formatting process to fail.
 *
 * This function type is used for:
 * - `ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR`
typedef ZyanStatus (*ZydisFormatterDecoratorFunc)(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisDecorator decorator);
 * Defines the `ZydisFormatter` struct.
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
 *
 * Do NOT change the order of the function fields or the values of the `ZydisFormatterFunction`
 * enum.
struct ZydisFormatter_
{
     * The formatter style.
    ZydisFormatterStyle style;
     * The `ZYDIS_FORMATTER_PROP_FORCE_SIZE` property.
    ZyanBool force_memory_size;
     * The `ZYDIS_FORMATTER_PROP_FORCE_SEGMENT` property.
    ZyanBool force_memory_segment;
     * The `ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_BRANCHES` property.
    ZyanBool force_relative_branches;
     * The `ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_RIPREL` property.
    ZyanBool force_relative_riprel;
     * The `ZYDIS_FORMATTER_PROP_PRINT_BRANCH_SIZE` property.
    ZyanBool print_branch_size;
     * The `ZYDIS_FORMATTER_DETAILED_PREFIXES` property.
    ZyanBool detailed_prefixes;
     * The `ZYDIS_FORMATTER_ADDR_BASE` property.
    ZydisNumericBase addr_base;
     * The `ZYDIS_FORMATTER_ADDR_SIGNEDNESS` property.
    ZydisSignedness addr_signedness;
     * The `ZYDIS_FORMATTER_ADDR_PADDING_ABSOLUTE` property.
    ZydisPadding addr_padding_absolute;
     * The `ZYDIS_FORMATTER_ADDR_PADDING_RELATIVE` property.
    ZydisPadding addr_padding_relative;
     * The `ZYDIS_FORMATTER_DISP_BASE` property.
    ZydisNumericBase disp_base;
     * The `ZYDIS_FORMATTER_DISP_SIGNEDNESS` property.
    ZydisSignedness disp_signedness;
     * The `ZYDIS_FORMATTER_DISP_PADDING` property.
    ZydisPadding disp_padding;
     * The `ZYDIS_FORMATTER_IMM_BASE` property.
    ZydisNumericBase imm_base;
     * The `ZYDIS_FORMATTER_IMM_SIGNEDNESS` property.
    ZydisSignedness imm_signedness;
     * The `ZYDIS_FORMATTER_IMM_PADDING` property.
    ZydisPadding imm_padding;
     * The `ZYDIS_FORMATTER_UPPERCASE_PREFIXES` property.
    ZyanI32 case_prefixes;
     * The `ZYDIS_FORMATTER_UPPERCASE_MNEMONIC` property.
    ZyanI32 case_mnemonic;
     * The `ZYDIS_FORMATTER_UPPERCASE_REGISTERS` property.
    ZyanI32 case_registers;
     * The `ZYDIS_FORMATTER_UPPERCASE_TYPECASTS` property.
    ZyanI32 case_typecasts;
     * The `ZYDIS_FORMATTER_UPPERCASE_DECORATORS` property.
    ZyanI32 case_decorators;
     * The `ZYDIS_FORMATTER_HEX_UPPERCASE` property.
    ZyanBool hex_uppercase;
     * The number formats for all numeric bases.
     *
     * Index 0 = prefix
     * Index 1 = suffix
    struct
    {
         * A pointer to the `ZyanStringView` to use as prefix/suffix.
        const ZyanStringView* string;
         * The `ZyanStringView` to use as prefix/suffix
        ZyanStringView string_data;
         * The actual string data.
        char buffer[11];
    } number_format[ZYDIS_NUMERIC_BASE_MAX_VALUE + 1][2];
     * The `ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION` function.
    ZydisFormatterFunc func_pre_instruction;
     * The `ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION` function.
    ZydisFormatterFunc func_post_instruction;
     * The `ZYDIS_FORMATTER_FUNC_FORMAT_INSTRUCTION` function.
    ZydisFormatterFunc func_format_instruction;
     * The `ZYDIS_FORMATTER_FUNC_PRE_OPERAND` function.
    ZydisFormatterFunc func_pre_operand;
     * The `ZYDIS_FORMATTER_FUNC_POST_OPERAND` function.
    ZydisFormatterFunc func_post_operand;
     * The `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG` function.
    ZydisFormatterFunc func_format_operand_reg;
     * The `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM` function.
    ZydisFormatterFunc func_format_operand_mem;
     * The `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR` function.
    ZydisFormatterFunc func_format_operand_ptr;
     * The `ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM` function.
    ZydisFormatterFunc func_format_operand_imm;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC function.
    ZydisFormatterFunc func_print_mnemonic;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_REGISTER` function.
    ZydisFormatterRegisterFunc func_print_register;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS` function.
    ZydisFormatterFunc func_print_address_abs;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL` function.
    ZydisFormatterFunc func_print_address_rel;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_DISP` function.
    ZydisFormatterFunc func_print_disp;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_IMM` function.
    ZydisFormatterFunc func_print_imm;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST` function.
    ZydisFormatterFunc func_print_typecast;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT` function.
    ZydisFormatterFunc func_print_segment;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES` function.
    ZydisFormatterFunc func_print_prefixes;
     * The `ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR` function.
    ZydisFormatterDecoratorFunc func_print_decorator;
};*/
return true
}

func (f *formatter)ZYDIS_EXPORT ZyanStatus ZydisFormatterInit()(ok bool){//col:1176
/*ZYDIS_EXPORT ZyanStatus ZydisFormatterInit(ZydisFormatter* formatter, ZydisFormatterStyle style);
 * Changes the value of the specified formatter `property`.
 *
 *
 *
 * This function returns `ZYAN_STATUS_INVALID_OPERATION` if a property can't be changed for the
 * current formatter-style.
ZYDIS_EXPORT ZyanStatus ZydisFormatterSetProperty(ZydisFormatter* formatter,
    ZydisFormatterProperty property, ZyanUPointer value);
 * Replaces a formatter function with a custom callback and/or retrieves the currently
 * used function.
 *
 *                      and receives the pointer of the currently used function.
 *
 *
 * Call this function with `callback` pointing to a `ZYAN_NULL` value to retrieve the currently
 * used function without replacing it.
 *
 * This function returns `ZYAN_STATUS_INVALID_OPERATION` if a function can't be replaced for the
 * current formatter-style.
ZYDIS_EXPORT ZyanStatus ZydisFormatterSetHook(ZydisFormatter* formatter,
    ZydisFormatterFunction type, const void** callback);
 * Formats the given instruction and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatInstruction(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address);
 * Formats the given instruction and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *                          callbacks.
 *
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatInstructionEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address, void* user_data);
 * Formats the given operand and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *
 *
 * Use `ZydisFormatterFormatInstruction` or `ZydisFormatterFormatInstructionEx` to format a
 * complete instruction.
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatOperand(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address);
 * Formats the given operand and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *                          callbacks.
 *
 *
 * Use `ZydisFormatterFormatInstruction` or `ZydisFormatterFormatInstructionEx` to format a
 * complete instruction.
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatOperandEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address, void* user_data);
 * Tokenizes the given instruction and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeInstruction(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token);
 * Tokenizes the given instruction and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *                          callbacks.
 *
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeInstructionEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token, void* user_data);
 * Tokenizes the given operand and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *
 *
 * Use `ZydisFormatterTokenizeInstruction` or `ZydisFormatterTokenizeInstructionEx` to tokenize a
 * complete instruction.
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeOperand(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token);
 * Tokenizes the given operand and writes it into the output buffer.
 *
 *                          to print relative addresses.
 *                          callbacks.
 *
 *
 * Use `ZydisFormatterTokenizeInstruction` or `ZydisFormatterTokenizeInstructionEx` to tokenize a
 * complete instruction.
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeOperandEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token, void* user_data);
#ifdef __cplusplus
}*/
return true
}



