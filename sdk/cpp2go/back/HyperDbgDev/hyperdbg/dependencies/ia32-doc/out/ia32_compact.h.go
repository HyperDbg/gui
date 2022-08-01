package out
//back\HyperDbgDev\hyperdbg\dependencies\ia32-doc\out\ia32_compact.h.back

const(
CPUID_SIGNATURE =                                              0x00000000 //col:1
CPUID_VERSION_INFO =                                           0x00000001 //col:2
CPUID_CACHE_PARAMS =                                           0x00000004 //col:3
CPUID_MONITOR_MWAIT =                                          0x00000005 //col:4
CPUID_THERMAL_POWER_MANAGEMENT =                               0x00000006 //col:5
CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS =                      0x00000007 //col:6
CPUID_DIRECT_CACHE_ACCESS_INFO =                               0x00000009 //col:7
CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING =                   0x0000000A //col:8
CPUID_EXTENDED_TOPOLOGY =                                      0x0000000B //col:9
CPUID_EXTENDED_STATE =                                         0x0000000D //col:10
CPUID_INTEL_RDT_MONITORING =                                   0x0000000F //col:11
CPUID_INTEL_RDT_ALLOCATION =                                   0x00000010 //col:12
CPUID_INTEL_SGX =                                              0x00000012 //col:13
CPUID_INTEL_PROCESSOR_TRACE =                                  0x00000014 //col:14
CPUID_TIME_STAMP_COUNTER =                                     0x00000015 //col:15
CPUID_PROCESSOR_FREQUENCY =                                    0x00000016 //col:16
CPUID_SOC_VENDOR =                                             0x00000017 //col:17
CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS =           0x00000018 //col:18
CPUID_EXTENDED_FUNCTION =                                      0x80000000 //col:19
CPUID_EXTENDED_CPU_SIGNATURE =                                 0x80000001 //col:20
CPUID_BRAND_STRING1 =                                          0x80000002 //col:21
CPUID_BRAND_STRING2 =                                          0x80000003 //col:22
CPUID_BRAND_STRING3 =                                          0x80000004 //col:23
CPUID_EXTENDED_CACHE_INFO =                                    0x80000006 //col:24
CPUID_EXTENDED_TIME_STAMP_COUNTER =                            0x80000007 //col:25
CPUID_EXTENDED_VIRT_PHYS_ADDRESS_SIZE =                        0x80000008 //col:26
IA32_P5_MC_ADDR =                                              0x00000000 //col:27
IA32_P5_MC_TYPE =                                              0x00000001 //col:28
IA32_MONITOR_FILTER_SIZE =                                     0x00000006 //col:29
IA32_TIME_STAMP_COUNTER =                                      0x00000010 //col:30
IA32_PLATFORM_ID =                                             0x00000017 //col:31
IA32_APIC_BASE =                                               0x0000001B //col:32
IA32_FEATURE_CONTROL =                                         0x0000003A //col:33
IA32_TSC_ADJUST =                                              0x0000003B //col:34
IA32_SPEC_CTRL =                                               0x00000048 //col:35
IA32_PRED_CMD =                                                0x00000049 //col:36
IA32_BIOS_UPDT_TRIG =                                          0x00000079 //col:37
IA32_BIOS_SIGN_ID =                                            0x0000008B //col:38
IA32_SGXLEPUBKEYHASH0 =                                        0x0000008C //col:39
IA32_SGXLEPUBKEYHASH1 =                                        0x0000008D //col:40
IA32_SGXLEPUBKEYHASH2 =                                        0x0000008E //col:41
IA32_SGXLEPUBKEYHASH3 =                                        0x0000008F //col:42
IA32_SMM_MONITOR_CTL =                                         0x0000009B //col:43
IA32_STM_FEATURES_IA32E =                                      0x00000001 //col:44
IA32_SMBASE =                                                  0x0000009E //col:45
IA32_PMC0 =                                                    0x000000C1 //col:46
IA32_PMC1 =                                                    0x000000C2 //col:47
IA32_PMC2 =                                                    0x000000C3 //col:48
IA32_PMC3 =                                                    0x000000C4 //col:49
IA32_PMC4 =                                                    0x000000C5 //col:50
IA32_PMC5 =                                                    0x000000C6 //col:51
IA32_PMC6 =                                                    0x000000C7 //col:52
IA32_PMC7 =                                                    0x000000C8 //col:53
IA32_MPERF =                                                   0x000000E7 //col:54
IA32_APERF =                                                   0x000000E8 //col:55
IA32_MTRRCAP =                                                 0x000000FE //col:56
IA32_ARCH_CAPABILITIES =                                       0x0000010A //col:57
IA32_FLUSH_CMD =                                               0x0000010B //col:58
IA32_TSX_CTRL =                                                0x00000122 //col:59
IA32_SYSENTER_CS =                                             0x00000174 //col:60
IA32_SYSENTER_ESP =                                            0x00000175 //col:61
IA32_SYSENTER_EIP =                                            0x00000176 //col:62
IA32_MCG_CAP =                                                 0x00000179 //col:63
IA32_MCG_STATUS =                                              0x0000017A //col:64
IA32_MCG_CTL =                                                 0x0000017B //col:65
IA32_PERFEVTSEL0 =                                             0x00000186 //col:66
IA32_PERFEVTSEL1 =                                             0x00000187 //col:67
IA32_PERFEVTSEL2 =                                             0x00000188 //col:68
IA32_PERFEVTSEL3 =                                             0x00000189 //col:69
IA32_PERF_STATUS =                                             0x00000198 //col:70
IA32_PERF_CTL =                                                0x00000199 //col:71
IA32_CLOCK_MODULATION =                                        0x0000019A //col:72
IA32_THERM_INTERRUPT =                                         0x0000019B //col:73
IA32_THERM_STATUS =                                            0x0000019C //col:74
IA32_MISC_ENABLE =                                             0x000001A0 //col:75
IA32_ENERGY_PERF_BIAS =                                        0x000001B0 //col:76
IA32_PACKAGE_THERM_STATUS =                                    0x000001B1 //col:77
IA32_PACKAGE_THERM_INTERRUPT =                                 0x000001B2 //col:78
IA32_DEBUGCTL =                                                0x000001D9 //col:79
IA32_SMRR_PHYSBASE =                                           0x000001F2 //col:80
IA32_SMRR_PHYSMASK =                                           0x000001F3 //col:81
IA32_PLATFORM_DCA_CAP =                                        0x000001F8 //col:82
IA32_CPU_DCA_CAP =                                             0x000001F9 //col:83
IA32_DCA_0_CAP =                                               0x000001FA //col:84
IA32_MTRR_PHYSBASE0 =                                          0x00000200 //col:85
IA32_MTRR_PHYSBASE1 =                                          0x00000202 //col:86
IA32_MTRR_PHYSBASE2 =                                          0x00000204 //col:87
IA32_MTRR_PHYSBASE3 =                                          0x00000206 //col:88
IA32_MTRR_PHYSBASE4 =                                          0x00000208 //col:89
IA32_MTRR_PHYSBASE5 =                                          0x0000020A //col:90
IA32_MTRR_PHYSBASE6 =                                          0x0000020C //col:91
IA32_MTRR_PHYSBASE7 =                                          0x0000020E //col:92
IA32_MTRR_PHYSBASE8 =                                          0x00000210 //col:93
IA32_MTRR_PHYSBASE9 =                                          0x00000212 //col:94
IA32_MTRR_PHYSMASK0 =                                          0x00000201 //col:95
IA32_MTRR_PHYSMASK1 =                                          0x00000203 //col:96
IA32_MTRR_PHYSMASK2 =                                          0x00000205 //col:97
IA32_MTRR_PHYSMASK3 =                                          0x00000207 //col:98
IA32_MTRR_PHYSMASK4 =                                          0x00000209 //col:99
IA32_MTRR_PHYSMASK5 =                                          0x0000020B //col:100
IA32_MTRR_PHYSMASK6 =                                          0x0000020D //col:101
IA32_MTRR_PHYSMASK7 =                                          0x0000020F //col:102
IA32_MTRR_PHYSMASK8 =                                          0x00000211 //col:103
IA32_MTRR_PHYSMASK9 =                                          0x00000213 //col:104
IA32_MTRR_FIX64K_BASE =                                        0x00000000 //col:105
IA32_MTRR_FIX64K_SIZE =                                        0x00010000 //col:106
IA32_MTRR_FIX64K_00000 =                                       0x00000250 //col:107
IA32_MTRR_FIX16K_BASE =                                        0x00080000 //col:108
IA32_MTRR_FIX16K_SIZE =                                        0x00004000 //col:109
IA32_MTRR_FIX16K_80000 =                                       0x00000258 //col:110
IA32_MTRR_FIX16K_A0000 =                                       0x00000259 //col:111
IA32_MTRR_FIX4K_BASE =                                         0x000C0000 //col:112
IA32_MTRR_FIX4K_SIZE =                                         0x00001000 //col:113
IA32_MTRR_FIX4K_C0000 =                                        0x00000268 //col:114
IA32_MTRR_FIX4K_C8000 =                                        0x00000269 //col:115
IA32_MTRR_FIX4K_D0000 =                                        0x0000026A //col:116
IA32_MTRR_FIX4K_D8000 =                                        0x0000026B //col:117
IA32_MTRR_FIX4K_E0000 =                                        0x0000026C //col:118
IA32_MTRR_FIX4K_E8000 =                                        0x0000026D //col:119
IA32_MTRR_FIX4K_F0000 =                                        0x0000026E //col:120
IA32_MTRR_FIX4K_F8000 =                                        0x0000026F //col:121
IA32_MTRR_FIX_COUNT =                                          ((1 + 2 + 8) * 8) //col:122
IA32_MTRR_VARIABLE_COUNT =                                     0x0000000A //col:123
IA32_MTRR_COUNT =                                              (IA32_MTRR_FIX_COUNT + IA32_MTRR_VARIABLE_COUNT) //col:124
IA32_PAT =                                                     0x00000277 //col:125
IA32_MC0_CTL2 =                                                0x00000280 //col:126
IA32_MC1_CTL2 =                                                0x00000281 //col:127
IA32_MC2_CTL2 =                                                0x00000282 //col:128
IA32_MC3_CTL2 =                                                0x00000283 //col:129
IA32_MC4_CTL2 =                                                0x00000284 //col:130
IA32_MC5_CTL2 =                                                0x00000285 //col:131
IA32_MC6_CTL2 =                                                0x00000286 //col:132
IA32_MC7_CTL2 =                                                0x00000287 //col:133
IA32_MC8_CTL2 =                                                0x00000288 //col:134
IA32_MC9_CTL2 =                                                0x00000289 //col:135
IA32_MC10_CTL2 =                                               0x0000028A //col:136
IA32_MC11_CTL2 =                                               0x0000028B //col:137
IA32_MC12_CTL2 =                                               0x0000028C //col:138
IA32_MC13_CTL2 =                                               0x0000028D //col:139
IA32_MC14_CTL2 =                                               0x0000028E //col:140
IA32_MC15_CTL2 =                                               0x0000028F //col:141
IA32_MC16_CTL2 =                                               0x00000290 //col:142
IA32_MC17_CTL2 =                                               0x00000291 //col:143
IA32_MC18_CTL2 =                                               0x00000292 //col:144
IA32_MC19_CTL2 =                                               0x00000293 //col:145
IA32_MC20_CTL2 =                                               0x00000294 //col:146
IA32_MC21_CTL2 =                                               0x00000295 //col:147
IA32_MC22_CTL2 =                                               0x00000296 //col:148
IA32_MC23_CTL2 =                                               0x00000297 //col:149
IA32_MC24_CTL2 =                                               0x00000298 //col:150
IA32_MC25_CTL2 =                                               0x00000299 //col:151
IA32_MC26_CTL2 =                                               0x0000029A //col:152
IA32_MC27_CTL2 =                                               0x0000029B //col:153
IA32_MC28_CTL2 =                                               0x0000029C //col:154
IA32_MC29_CTL2 =                                               0x0000029D //col:155
IA32_MC30_CTL2 =                                               0x0000029E //col:156
IA32_MC31_CTL2 =                                               0x0000029F //col:157
IA32_MTRR_DEF_TYPE =                                           0x000002FF //col:158
IA32_FIXED_CTR0 =                                              0x00000309 //col:159
IA32_FIXED_CTR1 =                                              0x0000030A //col:160
IA32_FIXED_CTR2 =                                              0x0000030B //col:161
IA32_PERF_CAPABILITIES =                                       0x00000345 //col:162
IA32_FIXED_CTR_CTRL =                                          0x0000038D //col:163
IA32_PERF_GLOBAL_STATUS =                                      0x0000038E //col:164
IA32_PERF_GLOBAL_CTRL =                                        0x0000038F //col:165
IA32_PERF_GLOBAL_STATUS_RESET =                                0x00000390 //col:166
IA32_PERF_GLOBAL_STATUS_SET =                                  0x00000391 //col:167
IA32_PERF_GLOBAL_INUSE =                                       0x00000392 //col:168
IA32_PEBS_ENABLE =                                             0x000003F1 //col:169
IA32_MC0_CTL =                                                 0x00000400 //col:170
IA32_MC1_CTL =                                                 0x00000404 //col:171
IA32_MC2_CTL =                                                 0x00000408 //col:172
IA32_MC3_CTL =                                                 0x0000040C //col:173
IA32_MC4_CTL =                                                 0x00000410 //col:174
IA32_MC5_CTL =                                                 0x00000414 //col:175
IA32_MC6_CTL =                                                 0x00000418 //col:176
IA32_MC7_CTL =                                                 0x0000041C //col:177
IA32_MC8_CTL =                                                 0x00000420 //col:178
IA32_MC9_CTL =                                                 0x00000424 //col:179
IA32_MC10_CTL =                                                0x00000428 //col:180
IA32_MC11_CTL =                                                0x0000042C //col:181
IA32_MC12_CTL =                                                0x00000430 //col:182
IA32_MC13_CTL =                                                0x00000434 //col:183
IA32_MC14_CTL =                                                0x00000438 //col:184
IA32_MC15_CTL =                                                0x0000043C //col:185
IA32_MC16_CTL =                                                0x00000440 //col:186
IA32_MC17_CTL =                                                0x00000444 //col:187
IA32_MC18_CTL =                                                0x00000448 //col:188
IA32_MC19_CTL =                                                0x0000044C //col:189
IA32_MC20_CTL =                                                0x00000450 //col:190
IA32_MC21_CTL =                                                0x00000454 //col:191
IA32_MC22_CTL =                                                0x00000458 //col:192
IA32_MC23_CTL =                                                0x0000045C //col:193
IA32_MC24_CTL =                                                0x00000460 //col:194
IA32_MC25_CTL =                                                0x00000464 //col:195
IA32_MC26_CTL =                                                0x00000468 //col:196
IA32_MC27_CTL =                                                0x0000046C //col:197
IA32_MC28_CTL =                                                0x00000470 //col:198
IA32_MC0_STATUS =                                              0x00000401 //col:199
IA32_MC1_STATUS =                                              0x00000405 //col:200
IA32_MC2_STATUS =                                              0x00000409 //col:201
IA32_MC3_STATUS =                                              0x0000040D //col:202
IA32_MC4_STATUS =                                              0x00000411 //col:203
IA32_MC5_STATUS =                                              0x00000415 //col:204
IA32_MC6_STATUS =                                              0x00000419 //col:205
IA32_MC7_STATUS =                                              0x0000041D //col:206
IA32_MC8_STATUS =                                              0x00000421 //col:207
IA32_MC9_STATUS =                                              0x00000425 //col:208
IA32_MC10_STATUS =                                             0x00000429 //col:209
IA32_MC11_STATUS =                                             0x0000042D //col:210
IA32_MC12_STATUS =                                             0x00000431 //col:211
IA32_MC13_STATUS =                                             0x00000435 //col:212
IA32_MC14_STATUS =                                             0x00000439 //col:213
IA32_MC15_STATUS =                                             0x0000043D //col:214
IA32_MC16_STATUS =                                             0x00000441 //col:215
IA32_MC17_STATUS =                                             0x00000445 //col:216
IA32_MC18_STATUS =                                             0x00000449 //col:217
IA32_MC19_STATUS =                                             0x0000044D //col:218
IA32_MC20_STATUS =                                             0x00000451 //col:219
IA32_MC21_STATUS =                                             0x00000455 //col:220
IA32_MC22_STATUS =                                             0x00000459 //col:221
IA32_MC23_STATUS =                                             0x0000045D //col:222
IA32_MC24_STATUS =                                             0x00000461 //col:223
IA32_MC25_STATUS =                                             0x00000465 //col:224
IA32_MC26_STATUS =                                             0x00000469 //col:225
IA32_MC27_STATUS =                                             0x0000046D //col:226
IA32_MC28_STATUS =                                             0x00000471 //col:227
IA32_MC0_ADDR =                                                0x00000402 //col:228
IA32_MC1_ADDR =                                                0x00000406 //col:229
IA32_MC2_ADDR =                                                0x0000040A //col:230
IA32_MC3_ADDR =                                                0x0000040E //col:231
IA32_MC4_ADDR =                                                0x00000412 //col:232
IA32_MC5_ADDR =                                                0x00000416 //col:233
IA32_MC6_ADDR =                                                0x0000041A //col:234
IA32_MC7_ADDR =                                                0x0000041E //col:235
IA32_MC8_ADDR =                                                0x00000422 //col:236
IA32_MC9_ADDR =                                                0x00000426 //col:237
IA32_MC10_ADDR =                                               0x0000042A //col:238
IA32_MC11_ADDR =                                               0x0000042E //col:239
IA32_MC12_ADDR =                                               0x00000432 //col:240
IA32_MC13_ADDR =                                               0x00000436 //col:241
IA32_MC14_ADDR =                                               0x0000043A //col:242
IA32_MC15_ADDR =                                               0x0000043E //col:243
IA32_MC16_ADDR =                                               0x00000442 //col:244
IA32_MC17_ADDR =                                               0x00000446 //col:245
IA32_MC18_ADDR =                                               0x0000044A //col:246
IA32_MC19_ADDR =                                               0x0000044E //col:247
IA32_MC20_ADDR =                                               0x00000452 //col:248
IA32_MC21_ADDR =                                               0x00000456 //col:249
IA32_MC22_ADDR =                                               0x0000045A //col:250
IA32_MC23_ADDR =                                               0x0000045E //col:251
IA32_MC24_ADDR =                                               0x00000462 //col:252
IA32_MC25_ADDR =                                               0x00000466 //col:253
IA32_MC26_ADDR =                                               0x0000046A //col:254
IA32_MC27_ADDR =                                               0x0000046E //col:255
IA32_MC28_ADDR =                                               0x00000472 //col:256
IA32_MC0_MISC =                                                0x00000403 //col:257
IA32_MC1_MISC =                                                0x00000407 //col:258
IA32_MC2_MISC =                                                0x0000040B //col:259
IA32_MC3_MISC =                                                0x0000040F //col:260
IA32_MC4_MISC =                                                0x00000413 //col:261
IA32_MC5_MISC =                                                0x00000417 //col:262
IA32_MC6_MISC =                                                0x0000041B //col:263
IA32_MC7_MISC =                                                0x0000041F //col:264
IA32_MC8_MISC =                                                0x00000423 //col:265
IA32_MC9_MISC =                                                0x00000427 //col:266
IA32_MC10_MISC =                                               0x0000042B //col:267
IA32_MC11_MISC =                                               0x0000042F //col:268
IA32_MC12_MISC =                                               0x00000433 //col:269
IA32_MC13_MISC =                                               0x00000437 //col:270
IA32_MC14_MISC =                                               0x0000043B //col:271
IA32_MC15_MISC =                                               0x0000043F //col:272
IA32_MC16_MISC =                                               0x00000443 //col:273
IA32_MC17_MISC =                                               0x00000447 //col:274
IA32_MC18_MISC =                                               0x0000044B //col:275
IA32_MC19_MISC =                                               0x0000044F //col:276
IA32_MC20_MISC =                                               0x00000453 //col:277
IA32_MC21_MISC =                                               0x00000457 //col:278
IA32_MC22_MISC =                                               0x0000045B //col:279
IA32_MC23_MISC =                                               0x0000045F //col:280
IA32_MC24_MISC =                                               0x00000463 //col:281
IA32_MC25_MISC =                                               0x00000467 //col:282
IA32_MC26_MISC =                                               0x0000046B //col:283
IA32_MC27_MISC =                                               0x0000046F //col:284
IA32_MC28_MISC =                                               0x00000473 //col:285
IA32_VMX_BASIC =                                               0x00000480 //col:286
IA32_VMX_PINBASED_CTLS =                                       0x00000481 //col:287
IA32_VMX_PROCBASED_CTLS =                                      0x00000482 //col:288
IA32_VMX_EXIT_CTLS =                                           0x00000483 //col:289
IA32_VMX_ENTRY_CTLS =                                          0x00000484 //col:290
IA32_VMX_MISC =                                                0x00000485 //col:291
IA32_VMX_CR0_FIXED0 =                                          0x00000486 //col:292
IA32_VMX_CR0_FIXED1 =                                          0x00000487 //col:293
IA32_VMX_CR4_FIXED0 =                                          0x00000488 //col:294
IA32_VMX_CR4_FIXED1 =                                          0x00000489 //col:295
IA32_VMX_VMCS_ENUM =                                           0x0000048A //col:296
IA32_VMX_PROCBASED_CTLS2 =                                     0x0000048B //col:297
IA32_VMX_EPT_VPID_CAP =                                        0x0000048C //col:298
IA32_VMX_TRUE_PINBASED_CTLS =                                  0x0000048D //col:299
IA32_VMX_TRUE_PROCBASED_CTLS =                                 0x0000048E //col:300
IA32_VMX_TRUE_EXIT_CTLS =                                      0x0000048F //col:301
IA32_VMX_TRUE_ENTRY_CTLS =                                     0x00000490 //col:302
IA32_VMX_VMFUNC =                                              0x00000491 //col:303
IA32_VMX_PROCBASED_CTLS3 =                                     0x00000492 //col:304
IA32_VMX_EXIT_CTLS2 =                                          0x00000493 //col:305
IA32_A_PMC0 =                                                  0x000004C1 //col:306
IA32_A_PMC1 =                                                  0x000004C2 //col:307
IA32_A_PMC2 =                                                  0x000004C3 //col:308
IA32_A_PMC3 =                                                  0x000004C4 //col:309
IA32_A_PMC4 =                                                  0x000004C5 //col:310
IA32_A_PMC5 =                                                  0x000004C6 //col:311
IA32_A_PMC6 =                                                  0x000004C7 //col:312
IA32_A_PMC7 =                                                  0x000004C8 //col:313
IA32_MCG_EXT_CTL =                                             0x000004D0 //col:314
IA32_SGX_SVN_STATUS =                                          0x00000500 //col:315
IA32_RTIT_OUTPUT_BASE =                                        0x00000560 //col:316
IA32_RTIT_OUTPUT_MASK_PTRS =                                   0x00000561 //col:317
IA32_RTIT_CTL =                                                0x00000570 //col:318
IA32_RTIT_STATUS =                                             0x00000571 //col:319
IA32_RTIT_CR3_MATCH =                                          0x00000572 //col:320
IA32_RTIT_ADDR0_A =                                            0x00000580 //col:321
IA32_RTIT_ADDR1_A =                                            0x00000582 //col:322
IA32_RTIT_ADDR2_A =                                            0x00000584 //col:323
IA32_RTIT_ADDR3_A =                                            0x00000586 //col:324
IA32_RTIT_ADDR0_B =                                            0x00000581 //col:325
IA32_RTIT_ADDR1_B =                                            0x00000583 //col:326
IA32_RTIT_ADDR2_B =                                            0x00000585 //col:327
IA32_RTIT_ADDR3_B =                                            0x00000587 //col:328
IA32_DS_AREA =                                                 0x00000600 //col:329
IA32_U_CET =                                                   0x000006A0 //col:330
IA32_S_CET =                                                   0x000006A2 //col:331
IA32_PL0_SSP =                                                 0x000006A4 //col:332
IA32_PL1_SSP =                                                 0x000006A5 //col:333
IA32_PL2_SSP =                                                 0x000006A6 //col:334
IA32_PL3_SSP =                                                 0x000006A7 //col:335
IA32_INTERRUPT_SSP_TABLE_ADDR =                                0x000006A8 //col:336
IA32_TSC_DEADLINE =                                            0x000006E0 //col:337
IA32_PM_ENABLE =                                               0x00000770 //col:338
IA32_HWP_CAPABILITIES =                                        0x00000771 //col:339
IA32_HWP_REQUEST_PKG =                                         0x00000772 //col:340
IA32_HWP_INTERRUPT =                                           0x00000773 //col:341
IA32_HWP_REQUEST =                                             0x00000774 //col:342
IA32_HWP_STATUS =                                              0x00000777 //col:343
IA32_X2APIC_APICID =                                           0x00000802 //col:344
IA32_X2APIC_VERSION =                                          0x00000803 //col:345
IA32_X2APIC_TPR =                                              0x00000808 //col:346
IA32_X2APIC_PPR =                                              0x0000080A //col:347
IA32_X2APIC_EOI =                                              0x0000080B //col:348
IA32_X2APIC_LDR =                                              0x0000080D //col:349
IA32_X2APIC_SIVR =                                             0x0000080F //col:350
IA32_X2APIC_ISR0 =                                             0x00000810 //col:351
IA32_X2APIC_ISR1 =                                             0x00000811 //col:352
IA32_X2APIC_ISR2 =                                             0x00000812 //col:353
IA32_X2APIC_ISR3 =                                             0x00000813 //col:354
IA32_X2APIC_ISR4 =                                             0x00000814 //col:355
IA32_X2APIC_ISR5 =                                             0x00000815 //col:356
IA32_X2APIC_ISR6 =                                             0x00000816 //col:357
IA32_X2APIC_ISR7 =                                             0x00000817 //col:358
IA32_X2APIC_TMR0 =                                             0x00000818 //col:359
IA32_X2APIC_TMR1 =                                             0x00000819 //col:360
IA32_X2APIC_TMR2 =                                             0x0000081A //col:361
IA32_X2APIC_TMR3 =                                             0x0000081B //col:362
IA32_X2APIC_TMR4 =                                             0x0000081C //col:363
IA32_X2APIC_TMR5 =                                             0x0000081D //col:364
IA32_X2APIC_TMR6 =                                             0x0000081E //col:365
IA32_X2APIC_TMR7 =                                             0x0000081F //col:366
IA32_X2APIC_IRR0 =                                             0x00000820 //col:367
IA32_X2APIC_IRR1 =                                             0x00000821 //col:368
IA32_X2APIC_IRR2 =                                             0x00000822 //col:369
IA32_X2APIC_IRR3 =                                             0x00000823 //col:370
IA32_X2APIC_IRR4 =                                             0x00000824 //col:371
IA32_X2APIC_IRR5 =                                             0x00000825 //col:372
IA32_X2APIC_IRR6 =                                             0x00000826 //col:373
IA32_X2APIC_IRR7 =                                             0x00000827 //col:374
IA32_X2APIC_ESR =                                              0x00000828 //col:375
IA32_X2APIC_LVT_CMCI =                                         0x0000082F //col:376
IA32_X2APIC_ICR =                                              0x00000830 //col:377
IA32_X2APIC_LVT_TIMER =                                        0x00000832 //col:378
IA32_X2APIC_LVT_THERMAL =                                      0x00000833 //col:379
IA32_X2APIC_LVT_PMI =                                          0x00000834 //col:380
IA32_X2APIC_LVT_LINT0 =                                        0x00000835 //col:381
IA32_X2APIC_LVT_LINT1 =                                        0x00000836 //col:382
IA32_X2APIC_LVT_ERROR =                                        0x00000837 //col:383
IA32_X2APIC_INIT_COUNT =                                       0x00000838 //col:384
IA32_X2APIC_CUR_COUNT =                                        0x00000839 //col:385
IA32_X2APIC_DIV_CONF =                                         0x0000083E //col:386
IA32_X2APIC_SELF_IPI =                                         0x0000083F //col:387
IA32_DEBUG_INTERFACE =                                         0x00000C80 //col:388
IA32_L3_QOS_CFG =                                              0x00000C81 //col:389
IA32_L2_QOS_CFG =                                              0x00000C82 //col:390
IA32_QM_EVTSEL =                                               0x00000C8D //col:391
IA32_QM_CTR =                                                  0x00000C8E //col:392
IA32_PQR_ASSOC =                                               0x00000C8F //col:393
IA32_BNDCFGS =                                                 0x00000D90 //col:394
IA32_XSS =                                                     0x00000DA0 //col:395
IA32_PKG_HDC_CTL =                                             0x00000DB0 //col:396
IA32_PM_CTL1 =                                                 0x00000DB1 //col:397
IA32_THREAD_STALL =                                            0x00000DB2 //col:398
IA32_EFER =                                                    0xC0000080 //col:399
IA32_STAR =                                                    0xC0000081 //col:400
IA32_LSTAR =                                                   0xC0000082 //col:401
IA32_CSTAR =                                                   0xC0000083 //col:402
IA32_FMASK =                                                   0xC0000084 //col:403
IA32_FS_BASE =                                                 0xC0000100 //col:404
IA32_GS_BASE =                                                 0xC0000101 //col:405
IA32_KERNEL_GS_BASE =                                          0xC0000102 //col:406
IA32_TSC_AUX =                                                 0xC0000103 //col:407
PDE_ENTRY_COUNT_32 =                                           0x00000400 //col:408
PTE_ENTRY_COUNT_32 =                                           0x00000400 //col:409
PML4_ENTRY_COUNT_64 =                                          0x00000200 //col:410
PDPTE_ENTRY_COUNT_64 =                                         0x00000200 //col:411
PDE_ENTRY_COUNT_64 =                                           0x00000200 //col:412
PTE_ENTRY_COUNT_64 =                                           0x00000200 //col:413
SEGMENT_DESCRIPTOR_TYPE_SYSTEM =                               0x00000000 //col:414
SEGMENT_DESCRIPTOR_TYPE_CODE_OR_DATA =                         0x00000001 //col:415
SEGMENT_DESCRIPTOR_TYPE_DATA_R =                               0x00000000 //col:416
SEGMENT_DESCRIPTOR_TYPE_DATA_RA =                              0x00000001 //col:417
SEGMENT_DESCRIPTOR_TYPE_DATA_RW =                              0x00000002 //col:418
SEGMENT_DESCRIPTOR_TYPE_DATA_RWA =                             0x00000003 //col:419
SEGMENT_DESCRIPTOR_TYPE_DATA_RE =                              0x00000004 //col:420
SEGMENT_DESCRIPTOR_TYPE_DATA_REA =                             0x00000005 //col:421
SEGMENT_DESCRIPTOR_TYPE_DATA_RWE =                             0x00000006 //col:422
SEGMENT_DESCRIPTOR_TYPE_DATA_RWEA =                            0x00000007 //col:423
SEGMENT_DESCRIPTOR_TYPE_CODE_E =                               0x00000008 //col:424
SEGMENT_DESCRIPTOR_TYPE_CODE_EA =                              0x00000009 //col:425
SEGMENT_DESCRIPTOR_TYPE_CODE_ER =                              0x0000000A //col:426
SEGMENT_DESCRIPTOR_TYPE_CODE_ERA =                             0x0000000B //col:427
SEGMENT_DESCRIPTOR_TYPE_CODE_EC =                              0x0000000C //col:428
SEGMENT_DESCRIPTOR_TYPE_CODE_ECA =                             0x0000000D //col:429
SEGMENT_DESCRIPTOR_TYPE_CODE_ERC =                             0x0000000E //col:430
SEGMENT_DESCRIPTOR_TYPE_CODE_ERCA =                            0x0000000F //col:431
SEGMENT_DESCRIPTOR_TYPE_RESERVED_1 =                           0x00000000 //col:432
SEGMENT_DESCRIPTOR_TYPE_TSS_16_AVAILABLE =                     0x00000001 //col:433
SEGMENT_DESCRIPTOR_TYPE_LDT =                                  0x00000002 //col:434
SEGMENT_DESCRIPTOR_TYPE_TSS_16_BUSY =                          0x00000003 //col:435
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE_16 =                         0x00000004 //col:436
SEGMENT_DESCRIPTOR_TYPE_TASK_GATE =                            0x00000005 //col:437
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE_16 =                    0x00000006 //col:438
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE_16 =                         0x00000007 //col:439
SEGMENT_DESCRIPTOR_TYPE_RESERVED_2 =                           0x00000008 //col:440
SEGMENT_DESCRIPTOR_TYPE_TSS_AVAILABLE =                        0x00000009 //col:441
SEGMENT_DESCRIPTOR_TYPE_RESERVED_3 =                           0x0000000A //col:442
SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY =                             0x0000000B //col:443
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE =                            0x0000000C //col:444
SEGMENT_DESCRIPTOR_TYPE_RESERVED_4 =                           0x0000000D //col:445
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE =                       0x0000000E //col:446
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE =                            0x0000000F //col:447
VMX_EXIT_REASON_XCPT_OR_NMI =                                  0x00000000 //col:448
VMX_EXIT_REASON_EXT_INT =                                      0x00000001 //col:449
VMX_EXIT_REASON_TRIPLE_FAULT =                                 0x00000002 //col:450
VMX_EXIT_REASON_INIT_SIGNAL =                                  0x00000003 //col:451
VMX_EXIT_REASON_SIPI =                                         0x00000004 //col:452
VMX_EXIT_REASON_IO_SMI =                                       0x00000005 //col:453
VMX_EXIT_REASON_SMI =                                          0x00000006 //col:454
VMX_EXIT_REASON_INT_WINDOW =                                   0x00000007 //col:455
VMX_EXIT_REASON_NMI_WINDOW =                                   0x00000008 //col:456
VMX_EXIT_REASON_TASK_SWITCH =                                  0x00000009 //col:457
VMX_EXIT_REASON_CPUID =                                        0x0000000A //col:458
VMX_EXIT_REASON_GETSEC =                                       0x0000000B //col:459
VMX_EXIT_REASON_HLT =                                          0x0000000C //col:460
VMX_EXIT_REASON_INVD =                                         0x0000000D //col:461
VMX_EXIT_REASON_INVLPG =                                       0x0000000E //col:462
VMX_EXIT_REASON_RDPMC =                                        0x0000000F //col:463
VMX_EXIT_REASON_RDTSC =                                        0x00000010 //col:464
VMX_EXIT_REASON_RSM =                                          0x00000011 //col:465
VMX_EXIT_REASON_VMCALL =                                       0x00000012 //col:466
VMX_EXIT_REASON_VMCLEAR =                                      0x00000013 //col:467
VMX_EXIT_REASON_VMLAUNCH =                                     0x00000014 //col:468
VMX_EXIT_REASON_VMPTRLD =                                      0x00000015 //col:469
VMX_EXIT_REASON_VMPTRST =                                      0x00000016 //col:470
VMX_EXIT_REASON_VMREAD =                                       0x00000017 //col:471
VMX_EXIT_REASON_VMRESUME =                                     0x00000018 //col:472
VMX_EXIT_REASON_VMWRITE =                                      0x00000019 //col:473
VMX_EXIT_REASON_VMXOFF =                                       0x0000001A //col:474
VMX_EXIT_REASON_VMXON =                                        0x0000001B //col:475
VMX_EXIT_REASON_MOV_CRX =                                      0x0000001C //col:476
VMX_EXIT_REASON_MOV_DRX =                                      0x0000001D //col:477
VMX_EXIT_REASON_IO_INSTR =                                     0x0000001E //col:478
VMX_EXIT_REASON_RDMSR =                                        0x0000001F //col:479
VMX_EXIT_REASON_WRMSR =                                        0x00000020 //col:480
VMX_EXIT_REASON_ERR_INVALID_GUEST_STATE =                      0x00000021 //col:481
VMX_EXIT_REASON_ERR_MSR_LOAD =                                 0x00000022 //col:482
VMX_EXIT_REASON_MWAIT =                                        0x00000024 //col:483
VMX_EXIT_REASON_MTF =                                          0x00000025 //col:484
VMX_EXIT_REASON_MONITOR =                                      0x00000027 //col:485
VMX_EXIT_REASON_PAUSE =                                        0x00000028 //col:486
VMX_EXIT_REASON_ERR_MACHINE_CHECK =                            0x00000029 //col:487
VMX_EXIT_REASON_TPR_BELOW_THRESHOLD =                          0x0000002B //col:488
VMX_EXIT_REASON_APIC_ACCESS =                                  0x0000002C //col:489
VMX_EXIT_REASON_VIRTUALIZED_EOI =                              0x0000002D //col:490
VMX_EXIT_REASON_XDTR_ACCESS =                                  0x0000002E //col:491
VMX_EXIT_REASON_TR_ACCESS =                                    0x0000002F //col:492
VMX_EXIT_REASON_EPT_VIOLATION =                                0x00000030 //col:493
VMX_EXIT_REASON_EPT_MISCONFIG =                                0x00000031 //col:494
VMX_EXIT_REASON_INVEPT =                                       0x00000032 //col:495
VMX_EXIT_REASON_RDTSCP =                                       0x00000033 //col:496
VMX_EXIT_REASON_PREEMPT_TIMER =                                0x00000034 //col:497
VMX_EXIT_REASON_INVVPID =                                      0x00000035 //col:498
VMX_EXIT_REASON_WBINVD =                                       0x00000036 //col:499
VMX_EXIT_REASON_XSETBV =                                       0x00000037 //col:500
VMX_EXIT_REASON_APIC_WRITE =                                   0x00000038 //col:501
VMX_EXIT_REASON_RDRAND =                                       0x00000039 //col:502
VMX_EXIT_REASON_INVPCID =                                      0x0000003A //col:503
VMX_EXIT_REASON_VMFUNC =                                       0x0000003B //col:504
VMX_EXIT_REASON_ENCLS =                                        0x0000003C //col:505
VMX_EXIT_REASON_RDSEED =                                       0x0000003D //col:506
VMX_EXIT_REASON_PML_FULL =                                     0x0000003E //col:507
VMX_EXIT_REASON_XSAVES =                                       0x0000003F //col:508
VMX_EXIT_REASON_XRSTORS =                                      0x00000040 //col:509
VMX_ERROR_VMCALL =                                             0x00000001 //col:510
VMX_ERROR_VMCLEAR_INVALID_PHYS_ADDR =                          0x00000002 //col:511
VMX_ERROR_VMCLEAR_INVALID_VMXON_PTR =                          0x00000003 //col:512
VMX_ERROR_VMLAUCH_NON_CLEAR_VMCS =                             0x00000004 //col:513
VMX_ERROR_VMRESUME_NON_LAUNCHED_VMCS =                         0x00000005 //col:514
VMX_ERROR_VMRESUME_CORRUPTED_VMCS =                            0x00000006 //col:515
VMX_ERROR_VMENTRY_INVALID_CONTROL_FIELDS =                     0x00000007 //col:516
VMX_ERROR_VMENTRY_INVALID_HOST_STATE =                         0x00000008 //col:517
VMX_ERROR_VMPTRLD_INVALID_PHYS_ADDR =                          0x00000009 //col:518
VMX_ERROR_VMPTRLD_VMXON_PTR =                                  0x0000000A //col:519
VMX_ERROR_VMPTRLD_WRONG_VMCS_REVISION =                        0x0000000B //col:520
VMX_ERROR_VMREAD_VMWRITE_INVALID_COMPONENT =                   0x0000000C //col:521
VMX_ERROR_VMWRITE_READONLY_COMPONENT =                         0x0000000D //col:522
VMX_ERROR_VMXON_IN_VMX_ROOT_OP =                               0x0000000F //col:523
VMX_ERROR_VMENTRY_INVALID_VMCS_EXEC_PTR =                      0x00000010 //col:524
VMX_ERROR_VMENTRY_NON_LAUNCHED_EXEC_VMCS =                     0x00000011 //col:525
VMX_ERROR_VMENTRY_EXEC_VMCS_PTR =                              0x00000012 //col:526
VMX_ERROR_VMCALL_NON_CLEAR_VMCS =                              0x00000013 //col:527
VMX_ERROR_VMCALL_INVALID_VMEXIT_FIELDS =                       0x00000014 //col:528
VMX_ERROR_VMCALL_INVALID_MSEG_REVISION =                       0x00000016 //col:529
VMX_ERROR_VMXOFF_DUAL_MONITOR =                                0x00000017 //col:530
VMX_ERROR_VMCALL_INVALID_SMM_MONITOR =                         0x00000018 //col:531
VMX_ERROR_VMENTRY_INVALID_VM_EXEC_CTRL =                       0x00000019 //col:532
VMX_ERROR_VMENTRY_MOV_SS =                                     0x0000001A //col:533
VMX_ERROR_INVEPTVPID_INVALID_OPERAND =                         0x0000001C //col:534
VMX_EXIT_QUALIFICATION_TYPE_CALL =                             0x00000000 //col:535
VMX_EXIT_QUALIFICATION_TYPE_IRET =                             0x00000001 //col:536
VMX_EXIT_QUALIFICATION_TYPE_JMP =                              0x00000002 //col:537
VMX_EXIT_QUALIFICATION_TYPE_IDT =                              0x00000003 //col:538
VMX_EXIT_QUALIFICATION_REGISTER_CR0 =                          0x00000000 //col:539
VMX_EXIT_QUALIFICATION_REGISTER_CR2 =                          0x00000002 //col:540
VMX_EXIT_QUALIFICATION_REGISTER_CR3 =                          0x00000003 //col:541
VMX_EXIT_QUALIFICATION_REGISTER_CR4 =                          0x00000004 //col:542
VMX_EXIT_QUALIFICATION_REGISTER_CR8 =                          0x00000008 //col:543
VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR =                      0x00000000 //col:544
VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR =                    0x00000001 //col:545
VMX_EXIT_QUALIFICATION_ACCESS_CLTS =                           0x00000002 //col:546
VMX_EXIT_QUALIFICATION_ACCESS_LMSW =                           0x00000003 //col:547
VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER =                      0x00000000 //col:548
VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY =                        0x00000001 //col:549
VMX_EXIT_QUALIFICATION_GENREG_RAX =                            0x00000000 //col:550
VMX_EXIT_QUALIFICATION_GENREG_RCX =                            0x00000001 //col:551
VMX_EXIT_QUALIFICATION_GENREG_RDX =                            0x00000002 //col:552
VMX_EXIT_QUALIFICATION_GENREG_RBX =                            0x00000003 //col:553
VMX_EXIT_QUALIFICATION_GENREG_RSP =                            0x00000004 //col:554
VMX_EXIT_QUALIFICATION_GENREG_RBP =                            0x00000005 //col:555
VMX_EXIT_QUALIFICATION_GENREG_RSI =                            0x00000006 //col:556
VMX_EXIT_QUALIFICATION_GENREG_RDI =                            0x00000007 //col:557
VMX_EXIT_QUALIFICATION_GENREG_R8 =                             0x00000008 //col:558
VMX_EXIT_QUALIFICATION_GENREG_R9 =                             0x00000009 //col:559
VMX_EXIT_QUALIFICATION_GENREG_R10 =                            0x0000000A //col:560
VMX_EXIT_QUALIFICATION_GENREG_R11 =                            0x0000000B //col:561
VMX_EXIT_QUALIFICATION_GENREG_R12 =                            0x0000000C //col:562
VMX_EXIT_QUALIFICATION_GENREG_R13 =                            0x0000000D //col:563
VMX_EXIT_QUALIFICATION_GENREG_R14 =                            0x0000000E //col:564
VMX_EXIT_QUALIFICATION_GENREG_R15 =                            0x0000000F //col:565
VMX_EXIT_QUALIFICATION_REGISTER_DR0 =                          0x00000000 //col:566
VMX_EXIT_QUALIFICATION_REGISTER_DR1 =                          0x00000001 //col:567
VMX_EXIT_QUALIFICATION_REGISTER_DR2 =                          0x00000002 //col:568
VMX_EXIT_QUALIFICATION_REGISTER_DR3 =                          0x00000003 //col:569
VMX_EXIT_QUALIFICATION_REGISTER_DR6 =                          0x00000006 //col:570
VMX_EXIT_QUALIFICATION_REGISTER_DR7 =                          0x00000007 //col:571
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR =                   0x00000000 //col:572
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR =                 0x00000001 //col:573
VMX_EXIT_QUALIFICATION_WIDTH_1B =                              0x00000000 //col:574
VMX_EXIT_QUALIFICATION_WIDTH_2B =                              0x00000001 //col:575
VMX_EXIT_QUALIFICATION_WIDTH_4B =                              0x00000003 //col:576
VMX_EXIT_QUALIFICATION_DIRECTION_OUT =                         0x00000000 //col:577
VMX_EXIT_QUALIFICATION_DIRECTION_IN =                          0x00000001 //col:578
VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING =                  0x00000000 //col:579
VMX_EXIT_QUALIFICATION_IS_STRING_STRING =                      0x00000001 //col:580
VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP =                        0x00000000 //col:581
VMX_EXIT_QUALIFICATION_IS_REP_REP =                            0x00000001 //col:582
VMX_EXIT_QUALIFICATION_ENCODING_DX =                           0x00000000 //col:583
VMX_EXIT_QUALIFICATION_ENCODING_IMM =                          0x00000001 //col:584
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_READ =                      0x00000000 //col:585
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_WRITE =                     0x00000001 //col:586
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_INSTR_FETCH =               0x00000002 //col:587
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_EVENT_DELIVERY =            0x00000003 //col:588
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_EVENT_DELIVERY =          0x0000000A //col:589
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_INSTR =                   0x0000000F //col:590
IO_BITMAP_A_MIN =                                              0x00000000 //col:591
IO_BITMAP_A_MAX =                                              0x00007FFF //col:592
IO_BITMAP_B_MIN =                                              0x00008000 //col:593
IO_BITMAP_B_MAX =                                              0x0000FFFF //col:594
MSR_ID_LOW_MIN =                                               0x00000000 //col:595
MSR_ID_LOW_MAX =                                               0x00001FFF //col:596
MSR_ID_HIGH_MIN =                                              0xC0000000 //col:597
MSR_ID_HIGH_MAX =                                              0xC0001FFF //col:598
EPT_PAGE_WALK_LENGTH_4 =                                       0x00000003 //col:599
EPT_LEVEL_PML4E =                                              0x00000003 //col:600
EPT_LEVEL_PDPTE =                                              0x00000002 //col:601
EPT_LEVEL_PDE =                                                0x00000001 //col:602
EPT_LEVEL_PTE =                                                0x00000000 //col:603
EPML4_ENTRY_COUNT =                                            0x00000200 //col:604
EPDPTE_ENTRY_COUNT =                                           0x00000200 //col:605
EPDE_ENTRY_COUNT =                                             0x00000200 //col:606
EPTE_ENTRY_COUNT =                                             0x00000200 //col:607
VMCS_CTRL_VPID =                                               0x00000000 //col:608
VMCS_CTRL_POSTED_INTR_NOTIFY_VECTOR =                          0x00000002 //col:609
VMCS_CTRL_EPTP_INDEX =                                         0x00000004 //col:610
VMCS_CTRL_HLAT_PREFIX_SIZE =                                   0x00000006 //col:611
VMCS_GUEST_ES_SEL =                                            0x00000800 //col:612
VMCS_GUEST_CS_SEL =                                            0x00000802 //col:613
VMCS_GUEST_SS_SEL =                                            0x00000804 //col:614
VMCS_GUEST_DS_SEL =                                            0x00000806 //col:615
VMCS_GUEST_FS_SEL =                                            0x00000808 //col:616
VMCS_GUEST_GS_SEL =                                            0x0000080A //col:617
VMCS_GUEST_LDTR_SEL =                                          0x0000080C //col:618
VMCS_GUEST_TR_SEL =                                            0x0000080E //col:619
VMCS_GUEST_INTR_STATUS =                                       0x00000810 //col:620
VMCS_GUEST_PML_INDEX =                                         0x00000812 //col:621
VMCS_HOST_ES_SEL =                                             0x00000C00 //col:622
VMCS_HOST_CS_SEL =                                             0x00000C02 //col:623
VMCS_HOST_SS_SEL =                                             0x00000C04 //col:624
VMCS_HOST_DS_SEL =                                             0x00000C06 //col:625
VMCS_HOST_FS_SEL =                                             0x00000C08 //col:626
VMCS_HOST_GS_SEL =                                             0x00000C0A //col:627
VMCS_HOST_TR_SEL =                                             0x00000C0C //col:628
VMCS_CTRL_IO_BITMAP_A =                                        0x00002000 //col:629
VMCS_CTRL_IO_BITMAP_B =                                        0x00002002 //col:630
VMCS_CTRL_MSR_BITMAP =                                         0x00002004 //col:631
VMCS_CTRL_VMEXIT_MSR_STORE =                                   0x00002006 //col:632
VMCS_CTRL_VMEXIT_MSR_LOAD =                                    0x00002008 //col:633
VMCS_CTRL_VMENTRY_MSR_LOAD =                                   0x0000200A //col:634
VMCS_CTRL_EXEC_VMCS_PTR =                                      0x0000200C //col:635
VMCS_CTRL_PML_ADDR =                                           0x0000200E //col:636
VMCS_CTRL_TSC_OFFSET =                                         0x00002010 //col:637
VMCS_CTRL_VAPIC_PAGEADDR =                                     0x00002012 //col:638
VMCS_CTRL_APIC_ACCESSADDR =                                    0x00002014 //col:639
VMCS_CTRL_POSTED_INTR_DESC =                                   0x00002016 //col:640
VMCS_CTRL_VMFUNC_CTRLS =                                       0x00002018 //col:641
VMCS_CTRL_EPTP =                                               0x0000201A //col:642
VMCS_CTRL_EOI_BITMAP_0 =                                       0x0000201C //col:643
VMCS_CTRL_EOI_BITMAP_1 =                                       0x0000201E //col:644
VMCS_CTRL_EOI_BITMAP_2 =                                       0x00002020 //col:645
VMCS_CTRL_EOI_BITMAP_3 =                                       0x00002022 //col:646
VMCS_CTRL_EPTP_LIST =                                          0x00002024 //col:647
VMCS_CTRL_VMREAD_BITMAP =                                      0x00002026 //col:648
VMCS_CTRL_VMWRITE_BITMAP =                                     0x00002028 //col:649
VMCS_CTRL_VIRTXCPT_INFO_ADDR =                                 0x0000202A //col:650
VMCS_CTRL_XSS_EXITING_BITMAP =                                 0x0000202C //col:651
VMCS_CTRL_ENCLS_EXITING_BITMAP =                               0x0000202E //col:652
VMCS_CTRL_SPP_TABLE_POINTER =                                  0x00002030 //col:653
VMCS_CTRL_TSC_MULTIPLIER =                                     0x00002032 //col:654
VMCS_CTRL_PROC_EXEC3 =                                         0x00002034 //col:655
VMCS_CTRL_ENCLV_EXITING_BITMAP =                               0x00002036 //col:656
VMCS_CTRL_HLATP =                                              0x00002040 //col:657
VMCS_CTRL_SECONDARY_EXIT =                                     0x00002044 //col:658
VMCS_GUEST_PHYS_ADDR =                                         0x00002400 //col:659
VMCS_GUEST_VMCS_LINK_PTR =                                     0x00002800 //col:660
VMCS_GUEST_DEBUGCTL =                                          0x00002802 //col:661
VMCS_GUEST_PAT =                                               0x00002804 //col:662
VMCS_GUEST_EFER =                                              0x00002806 //col:663
VMCS_GUEST_PERF_GLOBAL_CTRL =                                  0x00002808 //col:664
VMCS_GUEST_PDPTE0 =                                            0x0000280A //col:665
VMCS_GUEST_PDPTE1 =                                            0x0000280C //col:666
VMCS_GUEST_PDPTE2 =                                            0x0000280E //col:667
VMCS_GUEST_PDPTE3 =                                            0x00002810 //col:668
VMCS_GUEST_BNDCFGS =                                           0x00002812 //col:669
VMCS_GUEST_RTIT_CTL =                                          0x00002814 //col:670
VMCS_GUEST_LBR_CTL =                                           0x00002816 //col:671
VMCS_GUEST_PKRS =                                              0x00002818 //col:672
VMCS_HOST_PAT =                                                0x00002C00 //col:673
VMCS_HOST_EFER =                                               0x00002C02 //col:674
VMCS_HOST_PERF_GLOBAL_CTRL =                                   0x00002C04 //col:675
VMCS_HOST_PKRS =                                               0x00002C06 //col:676
VMCS_CTRL_PIN_EXEC =                                           0x00004000 //col:677
VMCS_CTRL_PROC_EXEC =                                          0x00004002 //col:678
VMCS_CTRL_EXCEPTION_BITMAP =                                   0x00004004 //col:679
VMCS_CTRL_PAGEFAULT_ERROR_MASK =                               0x00004006 //col:680
VMCS_CTRL_PAGEFAULT_ERROR_MATCH =                              0x00004008 //col:681
VMCS_CTRL_CR3_TARGET_COUNT =                                   0x0000400A //col:682
VMCS_CTRL_PRIMARY_EXIT =                                       0x0000400C //col:683
VMCS_CTRL_EXIT_MSR_STORE_COUNT =                               0x0000400E //col:684
VMCS_CTRL_EXIT_MSR_LOAD_COUNT =                                0x00004010 //col:685
VMCS_CTRL_ENTRY =                                              0x00004012 //col:686
VMCS_CTRL_ENTRY_MSR_LOAD_COUNT =                               0x00004014 //col:687
VMCS_CTRL_ENTRY_INTERRUPTION_INFO =                            0x00004016 //col:688
VMCS_CTRL_ENTRY_EXCEPTION_ERRCODE =                            0x00004018 //col:689
VMCS_CTRL_ENTRY_INSTR_LENGTH =                                 0x0000401A //col:690
VMCS_CTRL_TPR_THRESHOLD =                                      0x0000401C //col:691
VMCS_CTRL_PROC_EXEC2 =                                         0x0000401E //col:692
VMCS_CTRL_PLE_GAP =                                            0x00004020 //col:693
VMCS_CTRL_PLE_WINDOW =                                         0x00004022 //col:694
VMCS_VM_INSTR_ERROR =                                          0x00004400 //col:695
VMCS_EXIT_REASON =                                             0x00004402 //col:696
VMCS_EXIT_INTERRUPTION_INFO =                                  0x00004404 //col:697
VMCS_EXIT_INTERRUPTION_ERROR_CODE =                            0x00004406 //col:698
VMCS_IDT_VECTORING_INFO =                                      0x00004408 //col:699
VMCS_IDT_VECTORING_ERROR_CODE =                                0x0000440A //col:700
VMCS_EXIT_INSTR_LENGTH =                                       0x0000440C //col:701
VMCS_EXIT_INSTR_INFO =                                         0x0000440E //col:702
VMCS_GUEST_ES_LIMIT =                                          0x00004800 //col:703
VMCS_GUEST_CS_LIMIT =                                          0x00004802 //col:704
VMCS_GUEST_SS_LIMIT =                                          0x00004804 //col:705
VMCS_GUEST_DS_LIMIT =                                          0x00004806 //col:706
VMCS_GUEST_FS_LIMIT =                                          0x00004808 //col:707
VMCS_GUEST_GS_LIMIT =                                          0x0000480A //col:708
VMCS_GUEST_LDTR_LIMIT =                                        0x0000480C //col:709
VMCS_GUEST_TR_LIMIT =                                          0x0000480E //col:710
VMCS_GUEST_GDTR_LIMIT =                                        0x00004810 //col:711
VMCS_GUEST_IDTR_LIMIT =                                        0x00004812 //col:712
VMCS_GUEST_ES_ACCESS_RIGHTS =                                  0x00004814 //col:713
VMCS_GUEST_CS_ACCESS_RIGHTS =                                  0x00004816 //col:714
VMCS_GUEST_SS_ACCESS_RIGHTS =                                  0x00004818 //col:715
VMCS_GUEST_DS_ACCESS_RIGHTS =                                  0x0000481A //col:716
VMCS_GUEST_FS_ACCESS_RIGHTS =                                  0x0000481C //col:717
VMCS_GUEST_GS_ACCESS_RIGHTS =                                  0x0000481E //col:718
VMCS_GUEST_LDTR_ACCESS_RIGHTS =                                0x00004820 //col:719
VMCS_GUEST_TR_ACCESS_RIGHTS =                                  0x00004822 //col:720
VMCS_GUEST_INTERRUPTIBILITY_STATE =                            0x00004824 //col:721
VMCS_GUEST_ACTIVITY_STATE =                                    0x00004826 //col:722
VMCS_GUEST_SMBASE =                                            0x00004828 //col:723
VMCS_GUEST_SYSENTER_CS =                                       0x0000482A //col:724
VMCS_GUEST_PREEMPT_TIMER_VALUE =                               0x0000482E //col:725
VMCS_HOST_SYSENTER_CS =                                        0x00004C00 //col:726
VMCS_CTRL_CR0_MASK =                                           0x00006000 //col:727
VMCS_CTRL_CR4_MASK =                                           0x00006002 //col:728
VMCS_CTRL_CR0_READ_SHADOW =                                    0x00006004 //col:729
VMCS_CTRL_CR4_READ_SHADOW =                                    0x00006006 //col:730
VMCS_CTRL_CR3_TARGET_VAL0 =                                    0x00006008 //col:731
VMCS_CTRL_CR3_TARGET_VAL1 =                                    0x0000600A //col:732
VMCS_CTRL_CR3_TARGET_VAL2 =                                    0x0000600C //col:733
VMCS_CTRL_CR3_TARGET_VAL3 =                                    0x0000600E //col:734
VMCS_EXIT_QUALIFICATION =                                      0x00006400 //col:735
VMCS_IO_RCX =                                                  0x00006402 //col:736
VMCS_IO_RSI =                                                  0x00006404 //col:737
VMCS_IO_RDI =                                                  0x00006406 //col:738
VMCS_IO_RIP =                                                  0x00006408 //col:739
VMCS_EXIT_GUEST_LINEAR_ADDR =                                  0x0000640A //col:740
VMCS_GUEST_CR0 =                                               0x00006800 //col:741
VMCS_GUEST_CR3 =                                               0x00006802 //col:742
VMCS_GUEST_CR4 =                                               0x00006804 //col:743
VMCS_GUEST_ES_BASE =                                           0x00006806 //col:744
VMCS_GUEST_CS_BASE =                                           0x00006808 //col:745
VMCS_GUEST_SS_BASE =                                           0x0000680A //col:746
VMCS_GUEST_DS_BASE =                                           0x0000680C //col:747
VMCS_GUEST_FS_BASE =                                           0x0000680E //col:748
VMCS_GUEST_GS_BASE =                                           0x00006810 //col:749
VMCS_GUEST_LDTR_BASE =                                         0x00006812 //col:750
VMCS_GUEST_TR_BASE =                                           0x00006814 //col:751
VMCS_GUEST_GDTR_BASE =                                         0x00006816 //col:752
VMCS_GUEST_IDTR_BASE =                                         0x00006818 //col:753
VMCS_GUEST_DR7 =                                               0x0000681A //col:754
VMCS_GUEST_RSP =                                               0x0000681C //col:755
VMCS_GUEST_RIP =                                               0x0000681E //col:756
VMCS_GUEST_RFLAGS =                                            0x00006820 //col:757
VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS =                          0x00006822 //col:758
VMCS_GUEST_SYSENTER_ESP =                                      0x00006824 //col:759
VMCS_GUEST_SYSENTER_EIP =                                      0x00006826 //col:760
VMCS_GUEST_S_CET =                                             0x00006C28 //col:761
VMCS_GUEST_SSP =                                               0x00006C2A //col:762
VMCS_GUEST_INTERRUPT_SSP_TABLE_ADDR =                          0x00006C2C //col:763
VMCS_HOST_CR0 =                                                0x00006C00 //col:764
VMCS_HOST_CR3 =                                                0x00006C02 //col:765
VMCS_HOST_CR4 =                                                0x00006C04 //col:766
VMCS_HOST_FS_BASE =                                            0x00006C06 //col:767
VMCS_HOST_GS_BASE =                                            0x00006C08 //col:768
VMCS_HOST_TR_BASE =                                            0x00006C0A //col:769
VMCS_HOST_GDTR_BASE =                                          0x00006C0C //col:770
VMCS_HOST_IDTR_BASE =                                          0x00006C0E //col:771
VMCS_HOST_SYSENTER_ESP =                                       0x00006C10 //col:772
VMCS_HOST_SYSENTER_EIP =                                       0x00006C12 //col:773
VMCS_HOST_RSP =                                                0x00006C14 //col:774
VMCS_HOST_RIP =                                                0x00006C16 //col:775
VMCS_HOST_S_CET =                                              0x00006C18 //col:776
VMCS_HOST_SSP =                                                0x00006C1A //col:777
VMCS_HOST_INTERRUPT_SSP_TABLE_ADDR =                           0x00006C1C //col:778
APIC_BASE =                                                    0xFEE00000 //col:779
APIC_ID =                                                      0x00000020 //col:780
APIC_VERSION =                                                 0x00000030 //col:781
APIC_TPR =                                                     0x00000080 //col:782
APIC_APR =                                                     0x00000090 //col:783
APIC_PPR =                                                     0x000000A0 //col:784
APIC_EOI =                                                     0x000000B0 //col:785
APIC_REMOTE_READ =                                             0x000000C0 //col:786
APIC_LOGICAL_DESTINATION =                                     0x000000D0 //col:787
APIC_DESTINATION_FORMAT =                                      0x000000E0 //col:788
APIC_SIV =                                                     0x000000F0 //col:789
APIC_ISR_31_0 =                                                0x00000100 //col:790
APIC_ISR_63_32 =                                               0x00000110 //col:791
APIC_ISR_95_64 =                                               0x00000120 //col:792
APIC_ISR_127_96 =                                              0x00000130 //col:793
APIC_ISR_159_128 =                                             0x00000140 //col:794
APIC_ISR_191_160 =                                             0x00000150 //col:795
APIC_ISR_223_192 =                                             0x00000160 //col:796
APIC_ISR_255_224 =                                             0x00000170 //col:797
APIC_TMR_31_0 =                                                0x00000180 //col:798
APIC_TMR_63_32 =                                               0x00000190 //col:799
APIC_TMR_95_64 =                                               0x000001A0 //col:800
APIC_TMR_127_96 =                                              0x000001B0 //col:801
APIC_TMR_159_128 =                                             0x000001C0 //col:802
APIC_TMR_191_160 =                                             0x000001D0 //col:803
APIC_TMR_223_192 =                                             0x000001E0 //col:804
APIC_TMR_255_224 =                                             0x000001F0 //col:805
APIC_IRR_31_0 =                                                0x00000200 //col:806
APIC_IRR_63_32 =                                               0x00000210 //col:807
APIC_IRR_95_64 =                                               0x00000220 //col:808
APIC_IRR_127_96 =                                              0x00000230 //col:809
APIC_IRR_159_128 =                                             0x00000240 //col:810
APIC_IRR_191_160 =                                             0x00000250 //col:811
APIC_IRR_223_192 =                                             0x00000260 //col:812
APIC_IRR_255_224 =                                             0x00000270 //col:813
APIC_ERROR_STATUS =                                            0x00000280 //col:814
APIC_CMCI =                                                    0x000002F0 //col:815
APIC_ICR_0_31 =                                                0x00000300 //col:816
APIC_ICR_32_63 =                                               0x00000310 //col:817
APIC_LVT_TIMER =                                               0x00000320 //col:818
APIC_LVT_THERMAL_SENSOR =                                      0x00000330 //col:819
APIC_LVT_PERFORMANCE_MONITORING_COUNTERS =                     0x00000340 //col:820
APIC_LVT_LINT0 =                                               0x00000350 //col:821
APIC_LVT_LINT1 =                                               0x00000360 //col:822
APIC_LVT_ERROR =                                               0x00000370 //col:823
APIC_INITIAL_COUNT =                                           0x00000380 //col:824
APIC_CURRENT_COUNT =                                           0x00000390 //col:825
APIC_DIVIDE_CONFIGURATION =                                    0x000003E0 //col:826
MEMORY_TYPE_UC =                                               0x00000000 //col:827
MEMORY_TYPE_WC =                                               0x00000001 //col:828
MEMORY_TYPE_WT =                                               0x00000004 //col:829
MEMORY_TYPE_WP =                                               0x00000005 //col:830
MEMORY_TYPE_WB =                                               0x00000006 //col:831
MEMORY_TYPE_UC_MINUS =                                         0x00000007 //col:832
MEMORY_TYPE_INVALID =                                          0x000000FF //col:833
VTD_ROOT_ENTRY_COUNT =                                         0x00000100 //col:834
VTD_CONTEXT_ENTRY_COUNT =                                      0x00000100 //col:835
VTD_VERSION =                                                  0x00000000 //col:836
VTD_CAPABILITY =                                               0x00000008 //col:837
VTD_EXTENDED_CAPABILITY =                                      0x00000010 //col:838
VTD_GLOBAL_COMMAND =                                           0x00000018 //col:839
VTD_GLOBAL_STATUS =                                            0x0000001C //col:840
VTD_ROOT_TABLE_ADDRESS =                                       0x00000020 //col:841
VTD_CONTEXT_COMMAND =                                          0x00000028 //col:842
VTD_INVALIDATE_ADDRESS =                                       0x00000000 //col:843
VTD_IOTLB_INVALIDATE =                                         0x00000008 //col:844
)

const(
  invpcid_individual_address                                    =  0x00000000  //col:2
  invpcid_single_context                                        =  0x00000001  //col:3
  invpcid_all_context_with_globals                              =  0x00000002  //col:4
  invpcid_all_context                                           =  0x00000003  //col:5
)


const(
  vmx_active                                                    =  0x00000000  //col:8
  vmx_hlt                                                       =  0x00000001  //col:9
  vmx_shutdown                                                  =  0x00000002  //col:10
  vmx_wait_for_sipi                                             =  0x00000003  //col:11
)


const(
  invept_single_context                                         =  0x00000001  //col:14
  invept_all_context                                            =  0x00000002  //col:15
)


const(
  invvpid_individual_address                                    =  0x00000000  //col:18
  invvpid_single_context                                        =  0x00000001  //col:19
  invvpid_all_context                                           =  0x00000002  //col:20
  invvpid_single_context_retaining_globals                      =  0x00000003  //col:21
)


const(
  non_maskable_interrupt                                        =  0x00000002  //col:24
  hardware_exception                                            =  0x00000003  //col:25
  software_interrupt                                            =  0x00000004  //col:26
  privileged_software_exception                                 =  0x00000005  //col:27
  software_exception                                            =  0x00000006  //col:28
  other_event                                                   =  0x00000007  //col:29
)


const(
  divide_error                                                  =  0x00000000  //col:32
  debug                                                         =  0x00000001  //col:33
  nmi                                                           =  0x00000002  //col:34
  breakpoint                                                    =  0x00000003  //col:35
  overflow                                                      =  0x00000004  //col:36
  bound_range_exceeded                                          =  0x00000005  //col:37
  invalid_opcode                                                =  0x00000006  //col:38
  device_not_available                                          =  0x00000007  //col:39
  double_fault                                                  =  0x00000008  //col:40
  coprocessor_segment_overrun                                   =  0x00000009  //col:41
  invalid_tss                                                   =  0x0000000A  //col:42
  segment_not_present                                           =  0x0000000B  //col:43
  stack_segment_fault                                           =  0x0000000C  //col:44
  general_protection                                            =  0x0000000D  //col:45
  page_fault                                                    =  0x0000000E  //col:46
  x87_floating_point_error                                      =  0x00000010  //col:47
  alignment_check                                               =  0x00000011  //col:48
  machine_check                                                 =  0x00000012  //col:49
  simd_floating_point_error                                     =  0x00000013  //col:50
  virtualization_exception                                      =  0x00000014  //col:51
  control_protection                                            =  0x00000015  //col:52
)



type typedef struct { struct{
max_cpuid_input_value uint32_t
ebx_value_genu uint32_t
ecx_value_ntel uint32_t
edx_value_inei uint32_t
}


type typedef struct { struct{
stepping_id uint32_t
model uint32_t
family_id uint32_t
processor_type uint32_t
reserved_1 uint32_t
extended_model_id uint32_t
extended_family_id uint32_t
}


type typedef struct { struct{
cache_type_field uint32_t
cache_level uint32_t
self_initializing_cache_level uint32_t
fully_associative_cache uint32_t
reserved_1 uint32_t
max_addressable_ids_for_logical_processors_sharing_this_cache: uint32_t
max_addressable_ids_for_processor_cores_in_physical_package: uint32_t
}


type typedef struct { struct{
smallest_monitor_line_size uint32_t
}


type typedef struct { struct{
temperature_sensor_supported uint32_t
intel_turbo_boost_technology_available uint32_t
apic_timer_always_running uint32_t
reserved_1 uint32_t
power_limit_notification uint32_t
clock_modulation_duty uint32_t
package_thermal_management uint32_t
hwp_base_registers uint32_t
hwp_notification uint32_t
hwp_activity_window uint32_t
hwp_energy_performance_preference uint32_t
hwp_package_level_request uint32_t
reserved_2 uint32_t
hdc uint32_t
intel_turbo_boost_max_technology_3_available uint32_t
hwp_capabilities uint32_t
hwp_peci_override uint32_t
flexible_hwp uint32_t
fast_access_mode_for_hwp_request_msr uint32_t
reserved_3 uint32_t
ignoring_idle_logical_processor_hwp_request uint32_t
reserved_4 uint32_t
intel_thread_director uint32_t
}


type typedef struct { struct{
number_of_sub_leaves uint32_t
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t
}


type typedef struct { struct{
version_id_of_architectural_performance_monitoring uint32_t
number_of_performance_monitoring_counter_per_logical_processor: uint32_t
bit_width_of_performance_monitoring_counter uint32_t
ebx_bit_vector_length uint32_t
}


type typedef struct { struct{
x2apic_id_to_unique_topology_id_shift uint32_t
}


type typedef struct { struct{
x87_state uint32_t
sse_state uint32_t
avx_state uint32_t
mpx_state uint32_t
avx_512_state uint32_t
used_for_ia32_xss_1 uint32_t
pkru_state uint32_t
reserved_1 uint32_t
used_for_ia32_xss_2 uint32_t
}


type typedef struct { struct{
reserved_1 uint32_t
supports_xsavec_and_compacted_xrstor uint32_t
supports_xgetbv_with_ecx_1 uint32_t
supports_xsave_xrstor_and_ia32_xss uint32_t
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t
}


type typedef struct { struct{
length_of_capacity_bit_mask uint32_t
}


type typedef struct { struct{
length_of_capacity_bit_mask uint32_t
}


type typedef struct { struct{
max_mba_throttling_value uint32_t
}


type typedef struct { struct{
sgx1 uint32_t
sgx2 uint32_t
reserved_1 uint32_t
sgx_enclv_advanced uint32_t
sgx_encls_advanced uint32_t
}


type typedef struct { struct{
valid_secs_attributes_0 uint32_t
}


type typedef struct { struct{
sub_leaf_type uint32_t
}


type typedef struct { struct{
sub_leaf_type uint32_t
reserved_1 uint32_t
epc_base_physical_address_1 uint32_t
}


type typedef struct { struct{
max_sub_leaf uint32_t
}


type typedef struct { struct{
number_of_configurable_address_ranges_for_filtering uint32_t
reserved_1 uint32_t
bitmap_of_supported_mtc_period_encodings uint32_t
}


type typedef struct { struct{
denominator uint32_t
}


type typedef struct { struct{
procesor_base_frequency_mhz uint32_t
}


type typedef struct { struct{
max_soc_id_index uint32_t
}


type typedef struct { struct{
soc_vendor_brand_string uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
max_sub_leaf uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
max_extended_functions uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
processor_brand_string_1 uint32_t
}


type typedef struct { struct{
processor_brand_string_5 uint32_t
}


type typedef struct { struct{
processor_brand_string_9 uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
reserved uint32_t
}


type typedef struct { struct{
number_of_physical_address_bits uint32_t
number_of_linear_address_bits uint32_t
}


type typedef struct { struct{
thread_adjust uint64_t
}


type typedef struct { struct{
mseg_header_revision uint32_t
monitor_features uint32_t
gdtr_limit uint32_t
gdtr_base_offset uint32_t
cs_selector uint32_t
eip_offset uint32_t
esp_offset uint32_t
cr3_offset uint32_t
}


type typedef struct { struct{
c0_mcnt uint64_t
}


type typedef struct { struct{
c0_acnt uint64_t
}


type typedef struct { struct{
stall_cycle_cnt uint64_t
}


type typedef struct { struct{
limit uint16_t
base_address uint32_t
}


type typedef struct { struct{
limit uint16_t
base_address uint64_t
}


type typedef struct { struct{
segment_limit_low uint16_t
base_address_low uint16_t
base_address_middle uint32_t
type uint32_t
descriptor_type uint32_t
descriptor_privilege_level uint32_t
present uint32_t
segment_limit_high uint32_t
available_bit uint32_t
long_mode uint32_t
default_big uint32_t
granularity uint32_t
base_address_high uint32_t
}


type typedef struct { struct{
segment_limit_low uint16_t
base_address_low uint16_t
base_address_middle uint32_t
type uint32_t
descriptor_type uint32_t
descriptor_privilege_level uint32_t
present uint32_t
segment_limit_high uint32_t
available_bit uint32_t
long_mode uint32_t
default_big uint32_t
granularity uint32_t
base_address_high uint32_t
}


type typedef struct { struct{
offset_low uint16_t
segment_selector uint16_t
interrupt_stack_table uint32_t
must_be_zero_0 uint32_t
type uint32_t
must_be_zero_1 uint32_t
descriptor_privilege_level uint32_t
present uint32_t
offset_middle uint32_t
}


type typedef struct { struct{
reserved_0 uint32_t
rsp0 uint64_t
rsp1 uint64_t
rsp2 uint64_t
reserved_1 uint64_t
ist1 uint64_t
ist2 uint64_t
ist3 uint64_t
ist4 uint64_t
ist5 uint64_t
ist6 uint64_t
ist7 uint64_t
reserved_2 uint64_t
reserved_3 uint16_t
io_map_base uint16_t
}


type typedef struct { struct{
reason uint32_t
exception_mask uint32_t
exit uint64_t
guest_linear_address uint64_t
guest_physical_address uint64_t
current_eptp_index uint16_t
}


type typedef struct { struct{
io_a[4096] uint8_t
io_b[4096] uint8_t
}


type typedef struct { struct{
rdmsr_low[1024] uint8_t
rdmsr_high[1024] uint8_t
wrmsr_low[1024] uint8_t
wrmsr_high[1024] uint8_t
}


type typedef struct { struct{
ept_pointer uint64_t
reserved uint64_t
}


type typedef struct { struct{
vpid uint16_t
reserved1 uint16_t
reserved2 uint32_t
linear_address uint64_t
}


type typedef struct { struct{
revision_id uint32_t
shadow_vmcs_indicator uint32_t
}


type typedef struct { struct{
revision_id uint32_t
must_be_zero uint32_t
}


type typedef struct { struct{
present uint64_t
reserved_1 uint64_t
context_table_pointer uint64_t
}


type typedef struct { struct{
present uint64_t
fault_processing_disable uint64_t
translation_type uint64_t
reserved_1 uint64_t
second_level_page_translation_pointer uint64_t
}




