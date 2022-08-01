package out
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\ia32-doc\out\ia32_defines_only.h.back

const(
CR0_PROTECTION_ENABLE =                                        0x01 //col:1
CR0_MONITOR_COPROCESSOR =                                      0x02 //col:2
CR0_EMULATE_FPU =                                              0x04 //col:3
CR0_TASK_SWITCHED =                                            0x08 //col:4
CR0_EXTENSION_TYPE =                                           0x10 //col:5
CR0_NUMERIC_ERROR =                                            0x20 //col:6
CR0_WRITE_PROTECT =                                            0x10000 //col:7
CR0_ALIGNMENT_MASK =                                           0x40000 //col:8
CR0_NOT_WRITE_THROUGH =                                        0x20000000 //col:9
CR0_CACHE_DISABLE =                                            0x40000000 //col:10
CR0_PAGING_ENABLE =                                            0x80000000 //col:11
CR3_PAGE_LEVEL_WRITE_THROUGH =                                 0x08 //col:12
CR3_PAGE_LEVEL_CACHE_DISABLE =                                 0x10 //col:13
CR3_ADDRESS_OF_PAGE_DIRECTORY =                                0xFFFFFFFFF000 //col:14
CR4_VIRTUAL_MODE_EXTENSIONS =                                  0x01 //col:15
CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS =                        0x02 //col:16
CR4_TIMESTAMP_DISABLE =                                        0x04 //col:17
CR4_DEBUGGING_EXTENSIONS =                                     0x08 //col:18
CR4_PAGE_SIZE_EXTENSIONS =                                     0x10 //col:19
CR4_PHYSICAL_ADDRESS_EXTENSION =                               0x20 //col:20
CR4_MACHINE_CHECK_ENABLE =                                     0x40 //col:21
CR4_PAGE_GLOBAL_ENABLE =                                       0x80 //col:22
CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE =                    0x100 //col:23
CR4_OS_FXSAVE_FXRSTOR_SUPPORT =                                0x200 //col:24
CR4_OS_XMM_EXCEPTION_SUPPORT =                                 0x400 //col:25
CR4_USERMODE_INSTRUCTION_PREVENTION =                          0x800 //col:26
CR4_LINEAR_ADDRESSES_57_BIT =                                  0x1000 //col:27
CR4_VMX_ENABLE =                                               0x2000 //col:28
CR4_SMX_ENABLE =                                               0x4000 //col:29
CR4_FSGSBASE_ENABLE =                                          0x10000 //col:30
CR4_PCID_ENABLE =                                              0x20000 //col:31
CR4_OS_XSAVE =                                                 0x40000 //col:32
CR4_KEY_LOCKER_ENABLE =                                        0x80000 //col:33
CR4_SMEP_ENABLE =                                              0x100000 //col:34
CR4_SMAP_ENABLE =                                              0x200000 //col:35
CR4_PROTECTION_KEY_ENABLE =                                    0x400000 //col:36
CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE =                          0x800000 //col:37
CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE =                0x1000000 //col:38
CR8_TASK_PRIORITY_LEVEL =                                      0x0F //col:39
CR8_RESERVED =                                                 0xFFFFFFFFFFFFFFF0 //col:40
DR6_BREAKPOINT_CONDITION =                                     0x0F //col:41
DR6_DEBUG_REGISTER_ACCESS_DETECTED =                           0x2000 //col:42
DR6_SINGLE_INSTRUCTION =                                       0x4000 //col:43
DR6_TASK_SWITCH =                                              0x8000 //col:44
DR6_RESTRICTED_TRANSACTIONAL_MEMORY =                          0x10000 //col:45
DR7_LOCAL_BREAKPOINT_0 =                                       0x01 //col:46
DR7_GLOBAL_BREAKPOINT_0 =                                      0x02 //col:47
DR7_LOCAL_BREAKPOINT_1 =                                       0x04 //col:48
DR7_GLOBAL_BREAKPOINT_1 =                                      0x08 //col:49
DR7_LOCAL_BREAKPOINT_2 =                                       0x10 //col:50
DR7_GLOBAL_BREAKPOINT_2 =                                      0x20 //col:51
DR7_LOCAL_BREAKPOINT_3 =                                       0x40 //col:52
DR7_GLOBAL_BREAKPOINT_3 =                                      0x80 //col:53
DR7_LOCAL_EXACT_BREAKPOINT =                                   0x100 //col:54
DR7_GLOBAL_EXACT_BREAKPOINT =                                  0x200 //col:55
DR7_RESTRICTED_TRANSACTIONAL_MEMORY =                          0x800 //col:56
DR7_GENERAL_DETECT =                                           0x2000 //col:57
DR7_READ_WRITE_0 =                                             0x30000 //col:58
DR7_LENGTH_0 =                                                 0xC0000 //col:59
DR7_READ_WRITE_1 =                                             0x300000 //col:60
DR7_LENGTH_1 =                                                 0xC00000 //col:61
DR7_READ_WRITE_2 =                                             0x3000000 //col:62
DR7_LENGTH_2 =                                                 0xC000000 //col:63
DR7_READ_WRITE_3 =                                             0x30000000 //col:64
DR7_LENGTH_3 =                                                 0xC0000000 //col:65
CPUID_SIGNATURE =                                              0x00000000 //col:66
CPUID_VERSION_INFO =                                           0x00000001 //col:67
CPUID_EAX_STEPPING_ID =                                        0x0F //col:68
CPUID_EAX_MODEL =                                              0xF0 //col:69
CPUID_EAX_FAMILY_ID =                                          0xF00 //col:70
CPUID_EAX_PROCESSOR_TYPE =                                     0x3000 //col:71
CPUID_EAX_EXTENDED_MODEL_ID =                                  0xF0000 //col:72
CPUID_EAX_EXTENDED_FAMILY_ID =                                 0xFF00000 //col:73
CPUID_EBX_BRAND_INDEX =                                        0xFF //col:74
CPUID_EBX_CLFLUSH_LINE_SIZE =                                  0xFF00 //col:75
CPUID_EBX_MAX_ADDRESSABLE_IDS =                                0xFF0000 //col:76
CPUID_EBX_INITIAL_APIC_ID =                                    0xFF000000 //col:77
CPUID_ECX_STREAMING_SIMD_EXTENSIONS_3 =                        0x01 //col:78
CPUID_ECX_PCLMULQDQ_INSTRUCTION =                              0x02 //col:79
CPUID_ECX_DS_AREA_64BIT_LAYOUT =                               0x04 //col:80
CPUID_ECX_MONITOR_MWAIT_INSTRUCTION =                          0x08 //col:81
CPUID_ECX_CPL_QUALIFIED_DEBUG_STORE =                          0x10 //col:82
CPUID_ECX_VIRTUAL_MACHINE_EXTENSIONS =                         0x20 //col:83
CPUID_ECX_SAFER_MODE_EXTENSIONS =                              0x40 //col:84
CPUID_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY =                0x80 //col:85
CPUID_ECX_THERMAL_MONITOR_2 =                                  0x100 //col:86
CPUID_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3 =           0x200 //col:87
CPUID_ECX_L1_CONTEXT_ID =                                      0x400 //col:88
CPUID_ECX_SILICON_DEBUG =                                      0x800 //col:89
CPUID_ECX_FMA_EXTENSIONS =                                     0x1000 //col:90
CPUID_ECX_CMPXCHG16B_INSTRUCTION =                             0x2000 //col:91
CPUID_ECX_XTPR_UPDATE_CONTROL =                                0x4000 //col:92
CPUID_ECX_PERFMON_AND_DEBUG_CAPABILITY =                       0x8000 //col:93
CPUID_ECX_PROCESS_CONTEXT_IDENTIFIERS =                        0x20000 //col:94
CPUID_ECX_DIRECT_CACHE_ACCESS =                                0x40000 //col:95
CPUID_ECX_SSE41_SUPPORT =                                      0x80000 //col:96
CPUID_ECX_SSE42_SUPPORT =                                      0x100000 //col:97
CPUID_ECX_X2APIC_SUPPORT =                                     0x200000 //col:98
CPUID_ECX_MOVBE_INSTRUCTION =                                  0x400000 //col:99
CPUID_ECX_POPCNT_INSTRUCTION =                                 0x800000 //col:100
CPUID_ECX_TSC_DEADLINE =                                       0x1000000 //col:101
CPUID_ECX_AESNI_INSTRUCTION_EXTENSIONS =                       0x2000000 //col:102
CPUID_ECX_XSAVE_XRSTOR_INSTRUCTION =                           0x4000000 //col:103
CPUID_ECX_OSX_SAVE =                                           0x8000000 //col:104
CPUID_ECX_AVX_SUPPORT =                                        0x10000000 //col:105
CPUID_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS =             0x20000000 //col:106
CPUID_ECX_RDRAND_INSTRUCTION =                                 0x40000000 //col:107
CPUID_EDX_FLOATING_POINT_UNIT_ON_CHIP =                        0x01 //col:108
CPUID_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS =                     0x02 //col:109
CPUID_EDX_DEBUGGING_EXTENSIONS =                               0x04 //col:110
CPUID_EDX_PAGE_SIZE_EXTENSION =                                0x08 //col:111
CPUID_EDX_TIMESTAMP_COUNTER =                                  0x10 //col:112
CPUID_EDX_RDMSR_WRMSR_INSTRUCTIONS =                           0x20 //col:113
CPUID_EDX_PHYSICAL_ADDRESS_EXTENSION =                         0x40 //col:114
CPUID_EDX_MACHINE_CHECK_EXCEPTION =                            0x80 //col:115
CPUID_EDX_CMPXCHG8B_INSTRUCTION =                              0x100 //col:116
CPUID_EDX_APIC_ON_CHIP =                                       0x200 //col:117
CPUID_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS =                      0x800 //col:118
CPUID_EDX_MEMORY_TYPE_RANGE_REGISTERS =                        0x1000 //col:119
CPUID_EDX_PAGE_GLOBAL_BIT =                                    0x2000 //col:120
CPUID_EDX_MACHINE_CHECK_ARCHITECTURE =                         0x4000 //col:121
CPUID_EDX_CONDITIONAL_MOVE_INSTRUCTIONS =                      0x8000 //col:122
CPUID_EDX_PAGE_ATTRIBUTE_TABLE =                               0x10000 //col:123
CPUID_EDX_PAGE_SIZE_EXTENSION_36BIT =                          0x20000 //col:124
CPUID_EDX_PROCESSOR_SERIAL_NUMBER =                            0x40000 //col:125
CPUID_EDX_CLFLUSH_INSTRUCTION =                                0x80000 //col:126
CPUID_EDX_DEBUG_STORE =                                        0x200000 //col:127
CPUID_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI =                      0x400000 //col:128
CPUID_EDX_MMX_SUPPORT =                                        0x800000 //col:129
CPUID_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS =                        0x1000000 //col:130
CPUID_EDX_SSE_SUPPORT =                                        0x2000000 //col:131
CPUID_EDX_SSE2_SUPPORT =                                       0x4000000 //col:132
CPUID_EDX_SELF_SNOOP =                                         0x8000000 //col:133
CPUID_EDX_HYPER_THREADING_TECHNOLOGY =                         0x10000000 //col:134
CPUID_EDX_THERMAL_MONITOR =                                    0x20000000 //col:135
CPUID_EDX_PENDING_BREAK_ENABLE =                               0x80000000 //col:136
CPUID_CACHE_PARAMS =                                           0x00000004 //col:137
CPUID_EAX_CACHE_TYPE_FIELD =                                   0x1F //col:138
CPUID_EAX_CACHE_LEVEL =                                        0xE0 //col:139
CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL =                      0x100 //col:140
CPUID_EAX_FULLY_ASSOCIATIVE_CACHE =                            0x200 //col:141
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE = 0x3FFC000 //col:142
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE = 0xFC000000 //col:143
CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE =                         0xFFF //col:144
CPUID_EBX_PHYSICAL_LINE_PARTITIONS =                           0x3FF000 //col:145
CPUID_EBX_WAYS_OF_ASSOCIATIVITY =                              0xFFC00000 //col:146
CPUID_ECX_NUMBER_OF_SETS =                                     0xFFFFFFFF //col:147
CPUID_EDX_WRITE_BACK_INVALIDATE =                              0x01 //col:148
CPUID_EDX_CACHE_INCLUSIVENESS =                                0x02 //col:149
CPUID_EDX_COMPLEX_CACHE_INDEXING =                             0x04 //col:150
CPUID_MONITOR_MWAIT =                                          0x00000005 //col:151
CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE =                         0xFFFF //col:152
CPUID_EBX_LARGEST_MONITOR_LINE_SIZE =                          0xFFFF //col:153
CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS =            0x01 //col:154
CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT = 0x02 //col:155
CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES =                          0x0F //col:156
CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES =                          0xF0 //col:157
CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES =                          0xF00 //col:158
CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES =                          0xF000 //col:159
CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES =                          0xF0000 //col:160
CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES =                          0xF00000 //col:161
CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES =                          0xF000000 //col:162
CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES =                          0xF0000000 //col:163
CPUID_THERMAL_POWER_MANAGEMENT =                               0x00000006 //col:164
CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED =                       0x01 //col:165
CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE =             0x02 //col:166
CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING =                          0x04 //col:167
CPUID_EAX_POWER_LIMIT_NOTIFICATION =                           0x10 //col:168
CPUID_EAX_CLOCK_MODULATION_DUTY =                              0x20 //col:169
CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT =                         0x40 //col:170
CPUID_EAX_HWP_BASE_REGISTERS =                                 0x80 //col:171
CPUID_EAX_HWP_NOTIFICATION =                                   0x100 //col:172
CPUID_EAX_HWP_ACTIVITY_WINDOW =                                0x200 //col:173
CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE =                  0x400 //col:174
CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST =                          0x800 //col:175
CPUID_EAX_HDC =                                                0x2000 //col:176
CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE =       0x4000 //col:177
CPUID_EAX_HWP_CAPABILITIES =                                   0x8000 //col:178
CPUID_EAX_HWP_PECI_OVERRIDE =                                  0x10000 //col:179
CPUID_EAX_FLEXIBLE_HWP =                                       0x20000 //col:180
CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR =               0x40000 //col:181
CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST =        0x100000 //col:182
CPUID_EAX_INTEL_THREAD_DIRECTOR =                              0x800000 //col:183
CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR =   0x0F //col:184
CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY =          0x01 //col:185
CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES =            0x08 //col:186
CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE =                 0xFF00 //col:187
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:188
CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS =                      0x00000007 //col:189
CPUID_EAX_NUMBER_OF_SUB_LEAVES =                               0xFFFFFFFF //col:190
CPUID_EBX_FSGSBASE =                                           0x01 //col:191
CPUID_EBX_IA32_TSC_ADJUST_MSR =                                0x02 //col:192
CPUID_EBX_SGX =                                                0x04 //col:193
CPUID_EBX_BMI1 =                                               0x08 //col:194
CPUID_EBX_HLE =                                                0x10 //col:195
CPUID_EBX_AVX2 =                                               0x20 //col:196
CPUID_EBX_FDP_EXCPTN_ONLY =                                    0x40 //col:197
CPUID_EBX_SMEP =                                               0x80 //col:198
CPUID_EBX_BMI2 =                                               0x100 //col:199
CPUID_EBX_ENHANCED_REP_MOVSB_STOSB =                           0x200 //col:200
CPUID_EBX_INVPCID =                                            0x400 //col:201
CPUID_EBX_RTM =                                                0x800 //col:202
CPUID_EBX_RDT_M =                                              0x1000 //col:203
CPUID_EBX_DEPRECATES =                                         0x2000 //col:204
CPUID_EBX_MPX =                                                0x4000 //col:205
CPUID_EBX_RDT =                                                0x8000 //col:206
CPUID_EBX_AVX512F =                                            0x10000 //col:207
CPUID_EBX_AVX512DQ =                                           0x20000 //col:208
CPUID_EBX_RDSEED =                                             0x40000 //col:209
CPUID_EBX_ADX =                                                0x80000 //col:210
CPUID_EBX_SMAP =                                               0x100000 //col:211
CPUID_EBX_AVX512_IFMA =                                        0x200000 //col:212
CPUID_EBX_CLFLUSHOPT =                                         0x800000 //col:213
CPUID_EBX_CLWB =                                               0x1000000 //col:214
CPUID_EBX_INTEL =                                              0x2000000 //col:215
CPUID_EBX_AVX512PF =                                           0x4000000 //col:216
CPUID_EBX_AVX512ER =                                           0x8000000 //col:217
CPUID_EBX_AVX512CD =                                           0x10000000 //col:218
CPUID_EBX_SHA =                                                0x20000000 //col:219
CPUID_EBX_AVX512BW =                                           0x40000000 //col:220
CPUID_EBX_AVX512VL =                                           0x80000000 //col:221
CPUID_ECX_PREFETCHWT1 =                                        0x01 //col:222
CPUID_ECX_AVX512_VBMI =                                        0x02 //col:223
CPUID_ECX_UMIP =                                               0x04 //col:224
CPUID_ECX_PKU =                                                0x08 //col:225
CPUID_ECX_OSPKE =                                              0x10 //col:226
CPUID_ECX_WAITPKG =                                            0x20 //col:227
CPUID_ECX_AVX512_VBMI2 =                                       0x40 //col:228
CPUID_ECX_CET_SS =                                             0x80 //col:229
CPUID_ECX_GFNI =                                               0x100 //col:230
CPUID_ECX_VAES =                                               0x200 //col:231
CPUID_ECX_VPCLMULQDQ =                                         0x400 //col:232
CPUID_ECX_AVX512_VNNI =                                        0x800 //col:233
CPUID_ECX_AVX512_BITALG =                                      0x1000 //col:234
CPUID_ECX_TME_EN =                                             0x2000 //col:235
CPUID_ECX_AVX512_VPOPCNTDQ =                                   0x4000 //col:236
CPUID_ECX_LA57 =                                               0x10000 //col:237
CPUID_ECX_MAWAU =                                              0x3E0000 //col:238
CPUID_ECX_RDPID =                                              0x400000 //col:239
CPUID_ECX_KL =                                                 0x800000 //col:240
CPUID_ECX_CLDEMOTE =                                           0x2000000 //col:241
CPUID_ECX_MOVDIRI =                                            0x8000000 //col:242
CPUID_ECX_MOVDIR64B =                                          0x10000000 //col:243
CPUID_ECX_SGX_LC =                                             0x40000000 //col:244
CPUID_ECX_PKS =                                                0x80000000 //col:245
CPUID_EDX_AVX512_4VNNIW =                                      0x04 //col:246
CPUID_EDX_AVX512_4FMAPS =                                      0x08 //col:247
CPUID_EDX_FAST_SHORT_REP_MOV =                                 0x10 //col:248
CPUID_EDX_AVX512_VP2INTERSECT =                                0x100 //col:249
CPUID_EDX_MD_CLEAR =                                           0x400 //col:250
CPUID_EDX_SERIALIZE =                                          0x4000 //col:251
CPUID_EDX_HYBRID =                                             0x8000 //col:252
CPUID_EDX_PCONFIG =                                            0x40000 //col:253
CPUID_EDX_CET_IBT =                                            0x100000 //col:254
CPUID_EDX_IBRS_IBPB =                                          0x4000000 //col:255
CPUID_EDX_STIBP =                                              0x8000000 //col:256
CPUID_EDX_L1D_FLUSH =                                          0x10000000 //col:257
CPUID_EDX_IA32_ARCH_CAPABILITIES =                             0x20000000 //col:258
CPUID_EDX_IA32_CORE_CAPABILITIES =                             0x40000000 //col:259
CPUID_EDX_SSBD =                                               0x80000000 //col:260
CPUID_DIRECT_CACHE_ACCESS_INFO =                               0x00000009 //col:261
CPUID_EAX_IA32_PLATFORM_DCA_CAP =                              0xFFFFFFFF //col:262
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:263
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:264
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:265
CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING =                   0x0000000A //col:266
CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING = 0xFF //col:267
CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR = 0xFF00 //col:268
CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER =        0xFF0000 //col:269
CPUID_EAX_EBX_BIT_VECTOR_LENGTH =                              0xFF000000 //col:270
CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE =                     0x01 //col:271
CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE =            0x02 //col:272
CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE =               0x04 //col:273
CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE =     0x08 //col:274
CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE =        0x10 //col:275
CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE =     0x20 //col:276
CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE =      0x40 //col:277
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:278
CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS =      0x1F //col:279
CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS =   0x1FE0 //col:280
CPUID_EDX_ANY_THREAD_DEPRECATION =                             0x8000 //col:281
CPUID_EXTENDED_TOPOLOGY =                                      0x0000000B //col:282
CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT =              0x1F //col:283
CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE =    0xFFFF //col:284
CPUID_ECX_LEVEL_NUMBER =                                       0xFF //col:285
CPUID_ECX_LEVEL_TYPE =                                         0xFF00 //col:286
CPUID_EDX_X2APIC_ID =                                          0xFFFFFFFF //col:287
CPUID_EXTENDED_STATE =                                         0x0000000D //col:288
CPUID_EAX_X87_STATE =                                          0x01 //col:289
CPUID_EAX_SSE_STATE =                                          0x02 //col:290
CPUID_EAX_AVX_STATE =                                          0x04 //col:291
CPUID_EAX_MPX_STATE =                                          0x18 //col:292
CPUID_EAX_AVX_512_STATE =                                      0xE0 //col:293
CPUID_EAX_USED_FOR_IA32_XSS_1 =                                0x100 //col:294
CPUID_EAX_PKRU_STATE =                                         0x200 //col:295
CPUID_EAX_USED_FOR_IA32_XSS_2 =                                0x2000 //col:296
CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0 =      0xFFFFFFFF //col:297
CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA =                 0xFFFFFFFF //col:298
CPUID_EDX_XCR0_SUPPORTED_BITS =                                0xFFFFFFFF //col:299
CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR =               0x02 //col:300
CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1 =                         0x04 //col:301
CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS =                 0x08 //col:302
CPUID_EBX_SIZE_OF_XSAVE_AREAD =                                0xFFFFFFFF //col:303
CPUID_ECX_USED_FOR_XCR0_1 =                                    0xFF //col:304
CPUID_ECX_PT_STATE =                                           0x100 //col:305
CPUID_ECX_USED_FOR_XCR0_2 =                                    0x200 //col:306
CPUID_ECX_CET_USER_STATE =                                     0x800 //col:307
CPUID_ECX_CET_SUPERVISOR_STATE =                               0x1000 //col:308
CPUID_ECX_HDC_STATE =                                          0x2000 //col:309
CPUID_ECX_LBR_STATE =                                          0x8000 //col:310
CPUID_ECX_HWP_STATE =                                          0x10000 //col:311
CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS =                      0xFFFFFFFF //col:312
CPUID_EAX_IA32_PLATFORM_DCA_CAP =                              0xFFFFFFFF //col:313
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:314
CPUID_ECX_ECX_2 =                                              0x01 //col:315
CPUID_ECX_ECX_1 =                                              0x02 //col:316
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:317
CPUID_INTEL_RDT_MONITORING =                                   0x0000000F //col:318
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:319
CPUID_EBX_RMID_MAX_RANGE =                                     0xFFFFFFFF //col:320
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:321
CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING =             0x02 //col:322
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:323
CPUID_EBX_CONVERSION_FACTOR =                                  0xFFFFFFFF //col:324
CPUID_ECX_RMID_MAX_RANGE =                                     0xFFFFFFFF //col:325
CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING =                   0x01 //col:326
CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING =             0x02 //col:327
CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING =             0x04 //col:328
CPUID_INTEL_RDT_ALLOCATION =                                   0x00000010 //col:329
CPUID_EAX_IA32_PLATFORM_DCA_CAP =                              0xFFFFFFFF //col:330
CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY =            0x02 //col:331
CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY =            0x04 //col:332
CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION =               0x08 //col:333
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:334
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:335
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK =                        0x1F //col:336
CPUID_EBX_EBX_0 =                                              0xFFFFFFFF //col:337
CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED =    0x04 //col:338
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED =                       0xFFFF //col:339
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK =                        0x1F //col:340
CPUID_EBX_EBX_0 =                                              0xFFFFFFFF //col:341
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:342
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED =                       0xFFFF //col:343
CPUID_EAX_MAX_MBA_THROTTLING_VALUE =                           0xFFF //col:344
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:345
CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR =                        0x04 //col:346
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED =                       0xFFFF //col:347
CPUID_INTEL_SGX =                                              0x00000012 //col:348
CPUID_EAX_SGX1 =                                               0x01 //col:349
CPUID_EAX_SGX2 =                                               0x02 //col:350
CPUID_EAX_SGX_ENCLV_ADVANCED =                                 0x20 //col:351
CPUID_EAX_SGX_ENCLS_ADVANCED =                                 0x40 //col:352
CPUID_EBX_MISCSELECT =                                         0xFFFFFFFF //col:353
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:354
CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64 =                             0xFF //col:355
CPUID_EDX_MAX_ENCLAVE_SIZE_64 =                                0xFF00 //col:356
CPUID_EAX_VALID_SECS_ATTRIBUTES_0 =                            0xFFFFFFFF //col:357
CPUID_EBX_VALID_SECS_ATTRIBUTES_1 =                            0xFFFFFFFF //col:358
CPUID_ECX_VALID_SECS_ATTRIBUTES_2 =                            0xFFFFFFFF //col:359
CPUID_EDX_VALID_SECS_ATTRIBUTES_3 =                            0xFFFFFFFF //col:360
CPUID_EAX_SUB_LEAF_TYPE =                                      0x0F //col:361
CPUID_EBX_ZERO =                                               0xFFFFFFFF //col:362
CPUID_ECX_ZERO =                                               0xFFFFFFFF //col:363
CPUID_EDX_ZERO =                                               0xFFFFFFFF //col:364
CPUID_EAX_SUB_LEAF_TYPE =                                      0x0F //col:365
CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1 =                        0xFFFFF000 //col:366
CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2 =                        0xFFFFF //col:367
CPUID_ECX_EPC_SECTION_PROPERTY =                               0x0F //col:368
CPUID_ECX_EPC_SIZE_1 =                                         0xFFFFF000 //col:369
CPUID_EDX_EPC_SIZE_2 =                                         0xFFFFF //col:370
CPUID_INTEL_PROCESSOR_TRACE =                                  0x00000014 //col:371
CPUID_EAX_MAX_SUB_LEAF =                                       0xFFFFFFFF //col:372
CPUID_EBX_FLAG0 =                                              0x01 //col:373
CPUID_EBX_FLAG1 =                                              0x02 //col:374
CPUID_EBX_FLAG2 =                                              0x04 //col:375
CPUID_EBX_FLAG3 =                                              0x08 //col:376
CPUID_EBX_FLAG4 =                                              0x10 //col:377
CPUID_EBX_FLAG5 =                                              0x20 //col:378
CPUID_EBX_FLAG6 =                                              0x40 //col:379
CPUID_EBX_FLAG7 =                                              0x80 //col:380
CPUID_EBX_FLAG8 =                                              0x100 //col:381
CPUID_ECX_FLAG0 =                                              0x01 //col:382
CPUID_ECX_FLAG1 =                                              0x02 //col:383
CPUID_ECX_FLAG2 =                                              0x04 //col:384
CPUID_ECX_FLAG3 =                                              0x08 //col:385
CPUID_ECX_FLAG31 =                                             0x80000000 //col:386
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:387
CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING = 0x07 //col:388
CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS =           0xFFFF0000 //col:389
CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS = 0xFFFF //col:390
CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS = 0xFFFF0000 //col:391
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:392
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:393
CPUID_TIME_STAMP_COUNTER =                                     0x00000015 //col:394
CPUID_EAX_DENOMINATOR =                                        0xFFFFFFFF //col:395
CPUID_EBX_NUMERATOR =                                          0xFFFFFFFF //col:396
CPUID_ECX_NOMINAL_FREQUENCY =                                  0xFFFFFFFF //col:397
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:398
CPUID_PROCESSOR_FREQUENCY =                                    0x00000016 //col:399
CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ =                        0xFFFF //col:400
CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ =                    0xFFFF //col:401
CPUID_ECX_BUS_FREQUENCY_MHZ =                                  0xFFFF //col:402
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:403
CPUID_SOC_VENDOR =                                             0x00000017 //col:404
CPUID_EAX_MAX_SOC_ID_INDEX =                                   0xFFFFFFFF //col:405
CPUID_EBX_SOC_VENDOR_ID =                                      0xFFFF //col:406
CPUID_EBX_IS_VENDOR_SCHEME =                                   0x10000 //col:407
CPUID_ECX_PROJECT_ID =                                         0xFFFFFFFF //col:408
CPUID_EDX_STEPPING_ID =                                        0xFFFFFFFF //col:409
CPUID_EAX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:410
CPUID_EBX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:411
CPUID_ECX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:412
CPUID_EDX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:413
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:414
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:415
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:416
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:417
CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS =           0x00000018 //col:418
CPUID_EAX_MAX_SUB_LEAF =                                       0xFFFFFFFF //col:419
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED =                         0x01 //col:420
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED =                         0x02 //col:421
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED =                         0x04 //col:422
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED =                         0x08 //col:423
CPUID_EBX_PARTITIONING =                                       0x700 //col:424
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00 =                           0xFFFF0000 //col:425
CPUID_ECX_NUMBER_OF_SETS =                                     0xFFFFFFFF //col:426
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD =                       0x1F //col:427
CPUID_EDX_TRANSLATION_CACHE_LEVEL =                            0xE0 //col:428
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE =                        0x100 //col:429
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS =         0x3FFC000 //col:430
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:431
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED =                         0x01 //col:432
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED =                         0x02 //col:433
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED =                         0x04 //col:434
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED =                         0x08 //col:435
CPUID_EBX_PARTITIONING =                                       0x700 //col:436
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01 =                           0xFFFF0000 //col:437
CPUID_ECX_NUMBER_OF_SETS =                                     0xFFFFFFFF //col:438
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD =                       0x1F //col:439
CPUID_EDX_TRANSLATION_CACHE_LEVEL =                            0xE0 //col:440
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE =                        0x100 //col:441
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS =         0x3FFC000 //col:442
CPUID_EXTENDED_FUNCTION =                                      0x80000000 //col:443
CPUID_EAX_MAX_EXTENDED_FUNCTIONS =                             0xFFFFFFFF //col:444
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:445
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:446
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:447
CPUID_EXTENDED_CPU_SIGNATURE =                                 0x80000001 //col:448
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:449
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:450
CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE =                 0x01 //col:451
CPUID_ECX_LZCNT =                                              0x20 //col:452
CPUID_ECX_PREFETCHW =                                          0x100 //col:453
CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE =            0x800 //col:454
CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE =                      0x100000 //col:455
CPUID_EDX_PAGES_1GB_AVAILABLE =                                0x4000000 //col:456
CPUID_EDX_RDTSCP_AVAILABLE =                                   0x8000000 //col:457
CPUID_EDX_IA64_AVAILABLE =                                     0x20000000 //col:458
CPUID_BRAND_STRING1 =                                          0x80000002 //col:459
CPUID_BRAND_STRING2 =                                          0x80000003 //col:460
CPUID_BRAND_STRING3 =                                          0x80000004 //col:461
CPUID_EAX_PROCESSOR_BRAND_STRING_1 =                           0xFFFFFFFF //col:462
CPUID_EBX_PROCESSOR_BRAND_STRING_2 =                           0xFFFFFFFF //col:463
CPUID_ECX_PROCESSOR_BRAND_STRING_3 =                           0xFFFFFFFF //col:464
CPUID_EDX_PROCESSOR_BRAND_STRING_4 =                           0xFFFFFFFF //col:465
CPUID_EAX_PROCESSOR_BRAND_STRING_5 =                           0xFFFFFFFF //col:466
CPUID_EBX_PROCESSOR_BRAND_STRING_6 =                           0xFFFFFFFF //col:467
CPUID_ECX_PROCESSOR_BRAND_STRING_7 =                           0xFFFFFFFF //col:468
CPUID_EDX_PROCESSOR_BRAND_STRING_8 =                           0xFFFFFFFF //col:469
CPUID_EAX_PROCESSOR_BRAND_STRING_9 =                           0xFFFFFFFF //col:470
CPUID_EBX_PROCESSOR_BRAND_STRING_10 =                          0xFFFFFFFF //col:471
CPUID_ECX_PROCESSOR_BRAND_STRING_11 =                          0xFFFFFFFF //col:472
CPUID_EDX_PROCESSOR_BRAND_STRING_12 =                          0xFFFFFFFF //col:473
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:474
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:475
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:476
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:477
CPUID_EXTENDED_CACHE_INFO =                                    0x80000006 //col:478
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:479
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:480
CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES =                           0xFF //col:481
CPUID_ECX_L2_ASSOCIATIVITY_FIELD =                             0xF000 //col:482
CPUID_ECX_CACHE_SIZE_IN_1K_UNITS =                             0xFFFF0000 //col:483
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:484
CPUID_EXTENDED_TIME_STAMP_COUNTER =                            0x80000007 //col:485
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:486
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:487
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:488
CPUID_EDX_INVARIANT_TSC_AVAILABLE =                            0x100 //col:489
CPUID_EXTENDED_VIRT_PHYS_ADDRESS_SIZE =                        0x80000008 //col:490
CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS =                    0xFF //col:491
CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS =                      0xFF00 //col:492
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:493
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:494
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:495
IA32_P5_MC_ADDR =                                              0x00000000 //col:496
IA32_P5_MC_TYPE =                                              0x00000001 //col:497
IA32_MONITOR_FILTER_SIZE =                                     0x00000006 //col:498
IA32_TIME_STAMP_COUNTER =                                      0x00000010 //col:499
IA32_PLATFORM_ID =                                             0x00000017 //col:500
IA32_PLATFORM_ID_PLATFORM_ID =                                 0x1C000000000000 //col:501
IA32_APIC_BASE =                                               0x0000001B //col:502
IA32_APIC_BASE_BSP_FLAG =                                      0x100 //col:503
IA32_APIC_BASE_ENABLE_X2APIC_MODE =                            0x400 //col:504
IA32_APIC_BASE_APIC_GLOBAL_ENABLE =                            0x800 //col:505
IA32_APIC_BASE_APIC_BASE =                                     0xFFFFFFFFF000 //col:506
IA32_FEATURE_CONTROL =                                         0x0000003A //col:507
IA32_FEATURE_CONTROL_LOCK_BIT =                                0x01 //col:508
IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX =                   0x02 //col:509
IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX =                  0x04 //col:510
IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES =           0x7F00 //col:511
IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE =                    0x8000 //col:512
IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE =               0x20000 //col:513
IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE =                       0x40000 //col:514
IA32_FEATURE_CONTROL_LMCE_ON =                                 0x100000 //col:515
IA32_TSC_ADJUST =                                              0x0000003B //col:516
IA32_SPEC_CTRL =                                               0x00000048 //col:517
IA32_SPEC_CTRL_IBRS =                                          0x01 //col:518
IA32_SPEC_CTRL_STIBP =                                         0x02 //col:519
IA32_SPEC_CTRL_SSBD =                                          0x04 //col:520
IA32_PRED_CMD =                                                0x00000049 //col:521
IA32_PRED_CMD_IBPB =                                           0x01 //col:522
IA32_BIOS_UPDT_TRIG =                                          0x00000079 //col:523
IA32_BIOS_SIGN_ID =                                            0x0000008B //col:524
IA32_BIOS_SIGN_ID_RESERVED =                                   0xFFFFFFFF //col:525
IA32_BIOS_SIGN_ID_MICROCODE_UPDATE_SIGNATURE =                 0xFFFFFFFF00000000 //col:526
IA32_SGXLEPUBKEYHASH0 =                                        0x0000008C //col:527
IA32_SGXLEPUBKEYHASH1 =                                        0x0000008D //col:528
IA32_SGXLEPUBKEYHASH2 =                                        0x0000008E //col:529
IA32_SGXLEPUBKEYHASH3 =                                        0x0000008F //col:530
IA32_SMM_MONITOR_CTL =                                         0x0000009B //col:531
IA32_SMM_MONITOR_CTL_VALID =                                   0x01 //col:532
IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF =                0x04 //col:533
IA32_SMM_MONITOR_CTL_MSEG_BASE =                               0xFFFFF000 //col:534
IA32_STM_FEATURES_IA32E =                                      0x00000001 //col:535
IA32_SMBASE =                                                  0x0000009E //col:536
IA32_PMC0 =                                                    0x000000C1 //col:537
IA32_PMC1 =                                                    0x000000C2 //col:538
IA32_PMC2 =                                                    0x000000C3 //col:539
IA32_PMC3 =                                                    0x000000C4 //col:540
IA32_PMC4 =                                                    0x000000C5 //col:541
IA32_PMC5 =                                                    0x000000C6 //col:542
IA32_PMC6 =                                                    0x000000C7 //col:543
IA32_PMC7 =                                                    0x000000C8 //col:544
IA32_MPERF =                                                   0x000000E7 //col:545
IA32_APERF =                                                   0x000000E8 //col:546
IA32_MTRRCAP =                                                 0x000000FE //col:547
IA32_MTRRCAP_VARIABLE_RANGE_REGISTERS_COUNT =                  0xFF //col:548
IA32_MTRRCAP_FIXED_RANGE_REGISTERS_SUPPORTED =                 0x100 //col:549
IA32_MTRRCAP_WRITE_COMBINING =                                 0x400 //col:550
IA32_MTRRCAP_SYSTEM_MANAGEMENT_RANGE_REGISTER =                0x800 //col:551
IA32_ARCH_CAPABILITIES =                                       0x0000010A //col:552
IA32_ARCH_CAPABILITIES_RDCL_NO =                               0x01 //col:553
IA32_ARCH_CAPABILITIES_IBRS_ALL =                              0x02 //col:554
IA32_ARCH_CAPABILITIES_RSBA =                                  0x04 //col:555
IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY =                    0x08 //col:556
IA32_ARCH_CAPABILITIES_SSB_NO =                                0x10 //col:557
IA32_ARCH_CAPABILITIES_MDS_NO =                                0x20 //col:558
IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO =                     0x40 //col:559
IA32_ARCH_CAPABILITIES_TSX_CTRL =                              0x80 //col:560
IA32_ARCH_CAPABILITIES_TAA_NO =                                0x100 //col:561
IA32_FLUSH_CMD =                                               0x0000010B //col:562
IA32_FLUSH_CMD_L1D_FLUSH =                                     0x01 //col:563
IA32_TSX_CTRL =                                                0x00000122 //col:564
IA32_TSX_CTRL_RTM_DISABLE =                                    0x01 //col:565
IA32_TSX_CTRL_TSX_CPUID_CLEAR =                                0x02 //col:566
IA32_SYSENTER_CS =                                             0x00000174 //col:567
IA32_SYSENTER_CS_CS_SELECTOR =                                 0xFFFF //col:568
IA32_SYSENTER_CS_NOT_USED_1 =                                  0xFFFF0000 //col:569
IA32_SYSENTER_CS_NOT_USED_2 =                                  0xFFFFFFFF00000000 //col:570
IA32_SYSENTER_ESP =                                            0x00000175 //col:571
IA32_SYSENTER_EIP =                                            0x00000176 //col:572
IA32_MCG_CAP =                                                 0x00000179 //col:573
IA32_MCG_CAP_COUNT =                                           0xFF //col:574
IA32_MCG_CAP_MCG_CTL_P =                                       0x100 //col:575
IA32_MCG_CAP_MCG_EXT_P =                                       0x200 //col:576
IA32_MCG_CAP_MCP_CMCI_P =                                      0x400 //col:577
IA32_MCG_CAP_MCG_TES_P =                                       0x800 //col:578
IA32_MCG_CAP_MCG_EXT_CNT =                                     0xFF0000 //col:579
IA32_MCG_CAP_MCG_SER_P =                                       0x1000000 //col:580
IA32_MCG_CAP_MCG_ELOG_P =                                      0x4000000 //col:581
IA32_MCG_CAP_MCG_LMCE_P =                                      0x8000000 //col:582
IA32_MCG_STATUS =                                              0x0000017A //col:583
IA32_MCG_STATUS_RIPV =                                         0x01 //col:584
IA32_MCG_STATUS_EIPV =                                         0x02 //col:585
IA32_MCG_STATUS_MCIP =                                         0x04 //col:586
IA32_MCG_STATUS_LMCE_S =                                       0x08 //col:587
IA32_MCG_CTL =                                                 0x0000017B //col:588
IA32_PERFEVTSEL0 =                                             0x00000186 //col:589
IA32_PERFEVTSEL1 =                                             0x00000187 //col:590
IA32_PERFEVTSEL2 =                                             0x00000188 //col:591
IA32_PERFEVTSEL3 =                                             0x00000189 //col:592
IA32_PERFEVTSEL_EVENT_SELECT =                                 0xFF //col:593
IA32_PERFEVTSEL_U_MASK =                                       0xFF00 //col:594
IA32_PERFEVTSEL_USR =                                          0x10000 //col:595
IA32_PERFEVTSEL_OS =                                           0x20000 //col:596
IA32_PERFEVTSEL_EDGE =                                         0x40000 //col:597
IA32_PERFEVTSEL_PC =                                           0x80000 //col:598
IA32_PERFEVTSEL_INTR =                                         0x100000 //col:599
IA32_PERFEVTSEL_ANY_THREAD =                                   0x200000 //col:600
IA32_PERFEVTSEL_EN =                                           0x400000 //col:601
IA32_PERFEVTSEL_INV =                                          0x800000 //col:602
IA32_PERFEVTSEL_CMASK =                                        0xFF000000 //col:603
IA32_PERF_STATUS =                                             0x00000198 //col:604
IA32_PERF_STATUS_CURRENT_PERFORMANCE_STATE_VALUE =             0xFFFF //col:605
IA32_PERF_CTL =                                                0x00000199 //col:606
IA32_PERF_CTL_TARGET_PERFORMANCE_STATE_VALUE =                 0xFFFF //col:607
IA32_PERF_CTL_IDA_ENGAGE =                                     0x100000000 //col:608
IA32_CLOCK_MODULATION =                                        0x0000019A //col:609
IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE = 0x01 //col:610
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE =  0x0E //col:611
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE =      0x10 //col:612
IA32_THERM_INTERRUPT =                                         0x0000019B //col:613
IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE =       0x01 //col:614
IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE =        0x02 //col:615
IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE =                0x04 //col:616
IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE =                0x08 //col:617
IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE =   0x10 //col:618
IA32_THERM_INTERRUPT_THRESHOLD1_VALUE =                        0x7F00 //col:619
IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE =             0x8000 //col:620
IA32_THERM_INTERRUPT_THRESHOLD2_VALUE =                        0x7F0000 //col:621
IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE =             0x800000 //col:622
IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE =         0x1000000 //col:623
IA32_THERM_STATUS =                                            0x0000019C //col:624
IA32_THERM_STATUS_THERMAL_STATUS =                             0x01 //col:625
IA32_THERM_STATUS_THERMAL_STATUS_LOG =                         0x02 //col:626
IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT =                      0x04 //col:627
IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG =                        0x08 //col:628
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS =                0x10 //col:629
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG =            0x20 //col:630
IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS =                  0x40 //col:631
IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG =                     0x80 //col:632
IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS =                  0x100 //col:633
IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG =                     0x200 //col:634
IA32_THERM_STATUS_POWER_LIMITATION_STATUS =                    0x400 //col:635
IA32_THERM_STATUS_POWER_LIMITATION_LOG =                       0x800 //col:636
IA32_THERM_STATUS_CURRENT_LIMIT_STATUS =                       0x1000 //col:637
IA32_THERM_STATUS_CURRENT_LIMIT_LOG =                          0x2000 //col:638
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS =                  0x4000 //col:639
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG =                     0x8000 //col:640
IA32_THERM_STATUS_DIGITAL_READOUT =                            0x7F0000 //col:641
IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS =              0x78000000 //col:642
IA32_THERM_STATUS_READING_VALID =                              0x80000000 //col:643
IA32_MISC_ENABLE =                                             0x000001A0 //col:644
IA32_MISC_ENABLE_FAST_STRINGS_ENABLE =                         0x01 //col:645
IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE =    0x08 //col:646
IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE =            0x80 //col:647
IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE =            0x800 //col:648
IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE =  0x1000 //col:649
IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE =  0x10000 //col:650
IA32_MISC_ENABLE_ENABLE_MONITOR_FSM =                          0x40000 //col:651
IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL =                          0x400000 //col:652
IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE =                        0x800000 //col:653
IA32_MISC_ENABLE_XD_BIT_DISABLE =                              0x400000000 //col:654
IA32_ENERGY_PERF_BIAS =                                        0x000001B0 //col:655
IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE =                0x0F //col:656
IA32_PACKAGE_THERM_STATUS =                                    0x000001B1 //col:657
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS =                     0x01 //col:658
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG =                 0x02 //col:659
IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT =                      0x04 //col:660
IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG =                        0x08 //col:661
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS =        0x10 //col:662
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG =    0x20 //col:663
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS =          0x40 //col:664
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG =             0x80 //col:665
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS =          0x100 //col:666
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG =             0x200 //col:667
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS =            0x400 //col:668
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG =               0x800 //col:669
IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT =                    0x7F0000 //col:670
IA32_PACKAGE_THERM_INTERRUPT =                                 0x000001B2 //col:671
IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE = 0x01 //col:672
IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE = 0x02 //col:673
IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE =        0x04 //col:674
IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE =       0x10 //col:675
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE =                0x7F00 //col:676
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE =     0x8000 //col:677
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE =                0x7F0000 //col:678
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE =     0x800000 //col:679
IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE = 0x1000000 //col:680
IA32_DEBUGCTL =                                                0x000001D9 //col:681
IA32_DEBUGCTL_LBR =                                            0x01 //col:682
IA32_DEBUGCTL_BTF =                                            0x02 //col:683
IA32_DEBUGCTL_TR =                                             0x40 //col:684
IA32_DEBUGCTL_BTS =                                            0x80 //col:685
IA32_DEBUGCTL_BTINT =                                          0x100 //col:686
IA32_DEBUGCTL_BTS_OFF_OS =                                     0x200 //col:687
IA32_DEBUGCTL_BTS_OFF_USR =                                    0x400 //col:688
IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI =                             0x800 //col:689
IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI =                          0x1000 //col:690
IA32_DEBUGCTL_ENABLE_UNCORE_PMI =                              0x2000 //col:691
IA32_DEBUGCTL_FREEZE_WHILE_SMM =                               0x4000 //col:692
IA32_DEBUGCTL_RTM_DEBUG =                                      0x8000 //col:693
IA32_SMRR_PHYSBASE =                                           0x000001F2 //col:694
IA32_SMRR_PHYSBASE_TYPE =                                      0xFF //col:695
IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS =                0xFFFFF000 //col:696
IA32_SMRR_PHYSMASK =                                           0x000001F3 //col:697
IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK =                         0x800 //col:698
IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK =                   0xFFFFF000 //col:699
IA32_PLATFORM_DCA_CAP =                                        0x000001F8 //col:700
IA32_CPU_DCA_CAP =                                             0x000001F9 //col:701
IA32_DCA_0_CAP =                                               0x000001FA //col:702
IA32_DCA_0_CAP_DCA_ACTIVE =                                    0x01 //col:703
IA32_DCA_0_CAP_TRANSACTION =                                   0x06 //col:704
IA32_DCA_0_CAP_DCA_TYPE =                                      0x78 //col:705
IA32_DCA_0_CAP_DCA_QUEUE_SIZE =                                0x780 //col:706
IA32_DCA_0_CAP_DCA_DELAY =                                     0x1E000 //col:707
IA32_DCA_0_CAP_SW_BLOCK =                                      0x1000000 //col:708
IA32_DCA_0_CAP_HW_BLOCK =                                      0x4000000 //col:709
IA32_MTRR_PHYSBASE_TYPE =                                      0xFF //col:710
IA32_MTRR_PHYSBASE_PHYSICAL_ADDRES_BASE =                      0xFFFFFFFFF000 //col:711
IA32_MTRR_PHYSBASE0 =                                          0x00000200 //col:712
IA32_MTRR_PHYSBASE1 =                                          0x00000202 //col:713
IA32_MTRR_PHYSBASE2 =                                          0x00000204 //col:714
IA32_MTRR_PHYSBASE3 =                                          0x00000206 //col:715
IA32_MTRR_PHYSBASE4 =                                          0x00000208 //col:716
IA32_MTRR_PHYSBASE5 =                                          0x0000020A //col:717
IA32_MTRR_PHYSBASE6 =                                          0x0000020C //col:718
IA32_MTRR_PHYSBASE7 =                                          0x0000020E //col:719
IA32_MTRR_PHYSBASE8 =                                          0x00000210 //col:720
IA32_MTRR_PHYSBASE9 =                                          0x00000212 //col:721
IA32_MTRR_PHYSMASK_VALID =                                     0x800 //col:722
IA32_MTRR_PHYSMASK_PHYSICAL_ADDRES_MASK =                      0xFFFFFFFFF000 //col:723
IA32_MTRR_PHYSMASK0 =                                          0x00000201 //col:724
IA32_MTRR_PHYSMASK1 =                                          0x00000203 //col:725
IA32_MTRR_PHYSMASK2 =                                          0x00000205 //col:726
IA32_MTRR_PHYSMASK3 =                                          0x00000207 //col:727
IA32_MTRR_PHYSMASK4 =                                          0x00000209 //col:728
IA32_MTRR_PHYSMASK5 =                                          0x0000020B //col:729
IA32_MTRR_PHYSMASK6 =                                          0x0000020D //col:730
IA32_MTRR_PHYSMASK7 =                                          0x0000020F //col:731
IA32_MTRR_PHYSMASK8 =                                          0x00000211 //col:732
IA32_MTRR_PHYSMASK9 =                                          0x00000213 //col:733
IA32_MTRR_FIX64K_BASE =                                        0x00000000 //col:734
IA32_MTRR_FIX64K_SIZE =                                        0x00010000 //col:735
IA32_MTRR_FIX64K_00000 =                                       0x00000250 //col:736
IA32_MTRR_FIX16K_BASE =                                        0x00080000 //col:737
IA32_MTRR_FIX16K_SIZE =                                        0x00004000 //col:738
IA32_MTRR_FIX16K_80000 =                                       0x00000258 //col:739
IA32_MTRR_FIX16K_A0000 =                                       0x00000259 //col:740
IA32_MTRR_FIX4K_BASE =                                         0x000C0000 //col:741
IA32_MTRR_FIX4K_SIZE =                                         0x00001000 //col:742
IA32_MTRR_FIX4K_C0000 =                                        0x00000268 //col:743
IA32_MTRR_FIX4K_C8000 =                                        0x00000269 //col:744
IA32_MTRR_FIX4K_D0000 =                                        0x0000026A //col:745
IA32_MTRR_FIX4K_D8000 =                                        0x0000026B //col:746
IA32_MTRR_FIX4K_E0000 =                                        0x0000026C //col:747
IA32_MTRR_FIX4K_E8000 =                                        0x0000026D //col:748
IA32_MTRR_FIX4K_F0000 =                                        0x0000026E //col:749
IA32_MTRR_FIX4K_F8000 =                                        0x0000026F //col:750
IA32_MTRR_FIX_COUNT =                                          ((1 + 2 + 8) * 8) //col:751
IA32_MTRR_VARIABLE_COUNT =                                     0x0000000A //col:752
IA32_MTRR_COUNT =                                              (IA32_MTRR_FIX_COUNT + IA32_MTRR_VARIABLE_COUNT) //col:753
IA32_PAT =                                                     0x00000277 //col:754
IA32_PAT_PA0 =                                                 0x07 //col:755
IA32_PAT_PA1 =                                                 0x700 //col:756
IA32_PAT_PA2 =                                                 0x70000 //col:757
IA32_PAT_PA3 =                                                 0x7000000 //col:758
IA32_PAT_PA4 =                                                 0x700000000 //col:759
IA32_PAT_PA5 =                                                 0x70000000000 //col:760
IA32_PAT_PA6 =                                                 0x7000000000000 //col:761
IA32_PAT_PA7 =                                                 0x700000000000000 //col:762
IA32_MC0_CTL2 =                                                0x00000280 //col:763
IA32_MC1_CTL2 =                                                0x00000281 //col:764
IA32_MC2_CTL2 =                                                0x00000282 //col:765
IA32_MC3_CTL2 =                                                0x00000283 //col:766
IA32_MC4_CTL2 =                                                0x00000284 //col:767
IA32_MC5_CTL2 =                                                0x00000285 //col:768
IA32_MC6_CTL2 =                                                0x00000286 //col:769
IA32_MC7_CTL2 =                                                0x00000287 //col:770
IA32_MC8_CTL2 =                                                0x00000288 //col:771
IA32_MC9_CTL2 =                                                0x00000289 //col:772
IA32_MC10_CTL2 =                                               0x0000028A //col:773
IA32_MC11_CTL2 =                                               0x0000028B //col:774
IA32_MC12_CTL2 =                                               0x0000028C //col:775
IA32_MC13_CTL2 =                                               0x0000028D //col:776
IA32_MC14_CTL2 =                                               0x0000028E //col:777
IA32_MC15_CTL2 =                                               0x0000028F //col:778
IA32_MC16_CTL2 =                                               0x00000290 //col:779
IA32_MC17_CTL2 =                                               0x00000291 //col:780
IA32_MC18_CTL2 =                                               0x00000292 //col:781
IA32_MC19_CTL2 =                                               0x00000293 //col:782
IA32_MC20_CTL2 =                                               0x00000294 //col:783
IA32_MC21_CTL2 =                                               0x00000295 //col:784
IA32_MC22_CTL2 =                                               0x00000296 //col:785
IA32_MC23_CTL2 =                                               0x00000297 //col:786
IA32_MC24_CTL2 =                                               0x00000298 //col:787
IA32_MC25_CTL2 =                                               0x00000299 //col:788
IA32_MC26_CTL2 =                                               0x0000029A //col:789
IA32_MC27_CTL2 =                                               0x0000029B //col:790
IA32_MC28_CTL2 =                                               0x0000029C //col:791
IA32_MC29_CTL2 =                                               0x0000029D //col:792
IA32_MC30_CTL2 =                                               0x0000029E //col:793
IA32_MC31_CTL2 =                                               0x0000029F //col:794
IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD =                 0x7FFF //col:795
IA32_MC_CTL2_CMCI_EN =                                         0x40000000 //col:796
IA32_MTRR_DEF_TYPE =                                           0x000002FF //col:797
IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE =                       0x07 //col:798
IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE =                   0x400 //col:799
IA32_MTRR_DEF_TYPE_MTRR_ENABLE =                               0x800 //col:800
IA32_FIXED_CTR0 =                                              0x00000309 //col:801
IA32_FIXED_CTR1 =                                              0x0000030A //col:802
IA32_FIXED_CTR2 =                                              0x0000030B //col:803
IA32_PERF_CAPABILITIES =                                       0x00000345 //col:804
IA32_PERF_CAPABILITIES_LBR_FORMAT =                            0x3F //col:805
IA32_PERF_CAPABILITIES_PEBS_TRAP =                             0x40 //col:806
IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS =                   0x80 //col:807
IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT =                    0xF00 //col:808
IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED =         0x1000 //col:809
IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE =              0x2000 //col:810
IA32_FIXED_CTR_CTRL =                                          0x0000038D //col:811
IA32_FIXED_CTR_CTRL_EN0_OS =                                   0x01 //col:812
IA32_FIXED_CTR_CTRL_EN0_USR =                                  0x02 //col:813
IA32_FIXED_CTR_CTRL_ANY_THREAD0 =                              0x04 //col:814
IA32_FIXED_CTR_CTRL_EN0_PMI =                                  0x08 //col:815
IA32_FIXED_CTR_CTRL_EN1_OS =                                   0x10 //col:816
IA32_FIXED_CTR_CTRL_EN1_USR =                                  0x20 //col:817
IA32_FIXED_CTR_CTRL_ANY_THREAD1 =                              0x40 //col:818
IA32_FIXED_CTR_CTRL_EN1_PMI =                                  0x80 //col:819
IA32_FIXED_CTR_CTRL_EN2_OS =                                   0x100 //col:820
IA32_FIXED_CTR_CTRL_EN2_USR =                                  0x200 //col:821
IA32_FIXED_CTR_CTRL_ANY_THREAD2 =                              0x400 //col:822
IA32_FIXED_CTR_CTRL_EN2_PMI =                                  0x800 //col:823
IA32_PERF_GLOBAL_STATUS =                                      0x0000038E //col:824
IA32_PERF_GLOBAL_STATUS_OVF_PMC0 =                             0x01 //col:825
IA32_PERF_GLOBAL_STATUS_OVF_PMC1 =                             0x02 //col:826
IA32_PERF_GLOBAL_STATUS_OVF_PMC2 =                             0x04 //col:827
IA32_PERF_GLOBAL_STATUS_OVF_PMC3 =                             0x08 //col:828
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0 =                        0x100000000 //col:829
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1 =                        0x200000000 //col:830
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2 =                        0x400000000 //col:831
IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI =                       0x80000000000000 //col:832
IA32_PERF_GLOBAL_STATUS_LBR_FRZ =                              0x400000000000000 //col:833
IA32_PERF_GLOBAL_STATUS_CTR_FRZ =                              0x800000000000000 //col:834
IA32_PERF_GLOBAL_STATUS_ASCI =                                 0x1000000000000000 //col:835
IA32_PERF_GLOBAL_STATUS_OVF_UNCORE =                           0x2000000000000000 //col:836
IA32_PERF_GLOBAL_STATUS_OVF_BUF =                              0x4000000000000000 //col:837
IA32_PERF_GLOBAL_STATUS_COND_CHGD =                            0x8000000000000000 //col:838
IA32_PERF_GLOBAL_CTRL =                                        0x0000038F //col:839
IA32_PERF_GLOBAL_CTRL_EN_PMCN =                                0xFFFFFFFF //col:840
IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN =                          0xFFFFFFFF00000000 //col:841
IA32_PERF_GLOBAL_STATUS_RESET =                                0x00000390 //col:842
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN =                 0xFFFFFFFF //col:843
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN =           0x700000000 //col:844
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI =           0x80000000000000 //col:845
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ =                  0x400000000000000 //col:846
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ =                  0x800000000000000 //col:847
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI =                     0x1000000000000000 //col:848
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE =               0x2000000000000000 //col:849
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF =                  0x4000000000000000 //col:850
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD =                0x8000000000000000 //col:851
IA32_PERF_GLOBAL_STATUS_SET =                                  0x00000391 //col:852
IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN =                         0xFFFFFFFF //col:853
IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN =                   0x700000000 //col:854
IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI =                   0x80000000000000 //col:855
IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ =                          0x400000000000000 //col:856
IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ =                          0x800000000000000 //col:857
IA32_PERF_GLOBAL_STATUS_SET_ASCI =                             0x1000000000000000 //col:858
IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE =                       0x2000000000000000 //col:859
IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF =                          0x4000000000000000 //col:860
IA32_PERF_GLOBAL_INUSE =                                       0x00000392 //col:861
IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE =               0xFFFFFFFF //col:862
IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE =                0x700000000 //col:863
IA32_PERF_GLOBAL_INUSE_PMI_IN_USE =                            0x8000000000000000 //col:864
IA32_PEBS_ENABLE =                                             0x000003F1 //col:865
IA32_PEBS_ENABLE_ENABLE_PEBS =                                 0x01 //col:866
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1 =                    0x0E //col:867
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2 =                    0xF00000000 //col:868
IA32_MC0_CTL =                                                 0x00000400 //col:869
IA32_MC1_CTL =                                                 0x00000404 //col:870
IA32_MC2_CTL =                                                 0x00000408 //col:871
IA32_MC3_CTL =                                                 0x0000040C //col:872
IA32_MC4_CTL =                                                 0x00000410 //col:873
IA32_MC5_CTL =                                                 0x00000414 //col:874
IA32_MC6_CTL =                                                 0x00000418 //col:875
IA32_MC7_CTL =                                                 0x0000041C //col:876
IA32_MC8_CTL =                                                 0x00000420 //col:877
IA32_MC9_CTL =                                                 0x00000424 //col:878
IA32_MC10_CTL =                                                0x00000428 //col:879
IA32_MC11_CTL =                                                0x0000042C //col:880
IA32_MC12_CTL =                                                0x00000430 //col:881
IA32_MC13_CTL =                                                0x00000434 //col:882
IA32_MC14_CTL =                                                0x00000438 //col:883
IA32_MC15_CTL =                                                0x0000043C //col:884
IA32_MC16_CTL =                                                0x00000440 //col:885
IA32_MC17_CTL =                                                0x00000444 //col:886
IA32_MC18_CTL =                                                0x00000448 //col:887
IA32_MC19_CTL =                                                0x0000044C //col:888
IA32_MC20_CTL =                                                0x00000450 //col:889
IA32_MC21_CTL =                                                0x00000454 //col:890
IA32_MC22_CTL =                                                0x00000458 //col:891
IA32_MC23_CTL =                                                0x0000045C //col:892
IA32_MC24_CTL =                                                0x00000460 //col:893
IA32_MC25_CTL =                                                0x00000464 //col:894
IA32_MC26_CTL =                                                0x00000468 //col:895
IA32_MC27_CTL =                                                0x0000046C //col:896
IA32_MC28_CTL =                                                0x00000470 //col:897
IA32_MC0_STATUS =                                              0x00000401 //col:898
IA32_MC1_STATUS =                                              0x00000405 //col:899
IA32_MC2_STATUS =                                              0x00000409 //col:900
IA32_MC3_STATUS =                                              0x0000040D //col:901
IA32_MC4_STATUS =                                              0x00000411 //col:902
IA32_MC5_STATUS =                                              0x00000415 //col:903
IA32_MC6_STATUS =                                              0x00000419 //col:904
IA32_MC7_STATUS =                                              0x0000041D //col:905
IA32_MC8_STATUS =                                              0x00000421 //col:906
IA32_MC9_STATUS =                                              0x00000425 //col:907
IA32_MC10_STATUS =                                             0x00000429 //col:908
IA32_MC11_STATUS =                                             0x0000042D //col:909
IA32_MC12_STATUS =                                             0x00000431 //col:910
IA32_MC13_STATUS =                                             0x00000435 //col:911
IA32_MC14_STATUS =                                             0x00000439 //col:912
IA32_MC15_STATUS =                                             0x0000043D //col:913
IA32_MC16_STATUS =                                             0x00000441 //col:914
IA32_MC17_STATUS =                                             0x00000445 //col:915
IA32_MC18_STATUS =                                             0x00000449 //col:916
IA32_MC19_STATUS =                                             0x0000044D //col:917
IA32_MC20_STATUS =                                             0x00000451 //col:918
IA32_MC21_STATUS =                                             0x00000455 //col:919
IA32_MC22_STATUS =                                             0x00000459 //col:920
IA32_MC23_STATUS =                                             0x0000045D //col:921
IA32_MC24_STATUS =                                             0x00000461 //col:922
IA32_MC25_STATUS =                                             0x00000465 //col:923
IA32_MC26_STATUS =                                             0x00000469 //col:924
IA32_MC27_STATUS =                                             0x0000046D //col:925
IA32_MC28_STATUS =                                             0x00000471 //col:926
IA32_MC0_ADDR =                                                0x00000402 //col:927
IA32_MC1_ADDR =                                                0x00000406 //col:928
IA32_MC2_ADDR =                                                0x0000040A //col:929
IA32_MC3_ADDR =                                                0x0000040E //col:930
IA32_MC4_ADDR =                                                0x00000412 //col:931
IA32_MC5_ADDR =                                                0x00000416 //col:932
IA32_MC6_ADDR =                                                0x0000041A //col:933
IA32_MC7_ADDR =                                                0x0000041E //col:934
IA32_MC8_ADDR =                                                0x00000422 //col:935
IA32_MC9_ADDR =                                                0x00000426 //col:936
IA32_MC10_ADDR =                                               0x0000042A //col:937
IA32_MC11_ADDR =                                               0x0000042E //col:938
IA32_MC12_ADDR =                                               0x00000432 //col:939
IA32_MC13_ADDR =                                               0x00000436 //col:940
IA32_MC14_ADDR =                                               0x0000043A //col:941
IA32_MC15_ADDR =                                               0x0000043E //col:942
IA32_MC16_ADDR =                                               0x00000442 //col:943
IA32_MC17_ADDR =                                               0x00000446 //col:944
IA32_MC18_ADDR =                                               0x0000044A //col:945
IA32_MC19_ADDR =                                               0x0000044E //col:946
IA32_MC20_ADDR =                                               0x00000452 //col:947
IA32_MC21_ADDR =                                               0x00000456 //col:948
IA32_MC22_ADDR =                                               0x0000045A //col:949
IA32_MC23_ADDR =                                               0x0000045E //col:950
IA32_MC24_ADDR =                                               0x00000462 //col:951
IA32_MC25_ADDR =                                               0x00000466 //col:952
IA32_MC26_ADDR =                                               0x0000046A //col:953
IA32_MC27_ADDR =                                               0x0000046E //col:954
IA32_MC28_ADDR =                                               0x00000472 //col:955
IA32_MC0_MISC =                                                0x00000403 //col:956
IA32_MC1_MISC =                                                0x00000407 //col:957
IA32_MC2_MISC =                                                0x0000040B //col:958
IA32_MC3_MISC =                                                0x0000040F //col:959
IA32_MC4_MISC =                                                0x00000413 //col:960
IA32_MC5_MISC =                                                0x00000417 //col:961
IA32_MC6_MISC =                                                0x0000041B //col:962
IA32_MC7_MISC =                                                0x0000041F //col:963
IA32_MC8_MISC =                                                0x00000423 //col:964
IA32_MC9_MISC =                                                0x00000427 //col:965
IA32_MC10_MISC =                                               0x0000042B //col:966
IA32_MC11_MISC =                                               0x0000042F //col:967
IA32_MC12_MISC =                                               0x00000433 //col:968
IA32_MC13_MISC =                                               0x00000437 //col:969
IA32_MC14_MISC =                                               0x0000043B //col:970
IA32_MC15_MISC =                                               0x0000043F //col:971
IA32_MC16_MISC =                                               0x00000443 //col:972
IA32_MC17_MISC =                                               0x00000447 //col:973
IA32_MC18_MISC =                                               0x0000044B //col:974
IA32_MC19_MISC =                                               0x0000044F //col:975
IA32_MC20_MISC =                                               0x00000453 //col:976
IA32_MC21_MISC =                                               0x00000457 //col:977
IA32_MC22_MISC =                                               0x0000045B //col:978
IA32_MC23_MISC =                                               0x0000045F //col:979
IA32_MC24_MISC =                                               0x00000463 //col:980
IA32_MC25_MISC =                                               0x00000467 //col:981
IA32_MC26_MISC =                                               0x0000046B //col:982
IA32_MC27_MISC =                                               0x0000046F //col:983
IA32_MC28_MISC =                                               0x00000473 //col:984
IA32_VMX_BASIC =                                               0x00000480 //col:985
IA32_VMX_BASIC_VMCS_REVISION_ID =                              0x7FFFFFFF //col:986
IA32_VMX_BASIC_MUST_BE_ZERO =                                  0x80000000 //col:987
IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES =                            0x1FFF00000000 //col:988
IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH =                   0x1000000000000 //col:989
IA32_VMX_BASIC_DUAL_MONITOR =                                  0x2000000000000 //col:990
IA32_VMX_BASIC_MEMORY_TYPE =                                   0x3C000000000000 //col:991
IA32_VMX_BASIC_INS_OUTS_VMEXIT_INFORMATION =                   0x40000000000000 //col:992
IA32_VMX_BASIC_TRUE_CONTROLS =                                 0x80000000000000 //col:993
IA32_VMX_PINBASED_CTLS =                                       0x00000481 //col:994
IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING =            0x01 //col:995
IA32_VMX_PINBASED_CTLS_NMI_EXITING =                           0x08 //col:996
IA32_VMX_PINBASED_CTLS_VIRTUAL_NMIS =                          0x20 //col:997
IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER =         0x40 //col:998
IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS =             0x80 //col:999
IA32_VMX_PROCBASED_CTLS =                                      0x00000482 //col:1000
IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING =             0x04 //col:1001
IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING =                   0x08 //col:1002
IA32_VMX_PROCBASED_CTLS_HLT_EXITING =                          0x80 //col:1003
IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING =                       0x200 //col:1004
IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING =                        0x400 //col:1005
IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING =                        0x800 //col:1006
IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING =                        0x1000 //col:1007
IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING =                     0x8000 //col:1008
IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING =                    0x10000 //col:1009
IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS =           0x20000 //col:1010
IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING =                     0x80000 //col:1011
IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING =                    0x100000 //col:1012
IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW =                       0x200000 //col:1013
IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING =                   0x400000 //col:1014
IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING =                       0x800000 //col:1015
IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING =             0x1000000 //col:1016
IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS =                       0x2000000 //col:1017
IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG =                    0x8000000 //col:1018
IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS =                      0x10000000 //col:1019
IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING =                      0x20000000 //col:1020
IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING =                        0x40000000 //col:1021
IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS =          0x80000000 //col:1022
IA32_VMX_EXIT_CTLS =                                           0x00000483 //col:1023
IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS =                       0x04 //col:1024
IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE =                   0x200 //col:1025
IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL =                0x1000 //col:1026
IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT =             0x8000 //col:1027
IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT =                             0x40000 //col:1028
IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT =                             0x80000 //col:1029
IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER =                            0x100000 //col:1030
IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER =                            0x200000 //col:1031
IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE =           0x400000 //col:1032
IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS =                        0x800000 //col:1033
IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT =                       0x1000000 //col:1034
IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL =                       0x2000000 //col:1035
IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL =                        0x4000000 //col:1036
IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE =                       0x10000000 //col:1037
IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS =                            0x20000000 //col:1038
IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS =               0x80000000 //col:1039
IA32_VMX_ENTRY_CTLS =                                          0x00000484 //col:1040
IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS =                      0x04 //col:1041
IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST =                         0x200 //col:1042
IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM =                             0x400 //col:1043
IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT =        0x800 //col:1044
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL =               0x2000 //col:1045
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT =                            0x4000 //col:1046
IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER =                           0x8000 //col:1047
IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS =                        0x10000 //col:1048
IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT =                      0x20000 //col:1049
IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL =                       0x40000 //col:1050
IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE =                           0x100000 //col:1051
IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL =                        0x200000 //col:1052
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS =                           0x400000 //col:1053
IA32_VMX_MISC =                                                0x00000485 //col:1054
IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP =              0x1F //col:1055
IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT =                       0x20 //col:1056
IA32_VMX_MISC_ACTIVITY_STATES =                                0x1C0 //col:1057
IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX =                      0x4000 //col:1058
IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM =          0x8000 //col:1059
IA32_VMX_MISC_CR3_TARGET_COUNT =                               0x1FF0000 //col:1060
IA32_VMX_MISC_MAX_NUMBER_OF_MSR =                              0xE000000 //col:1061
IA32_VMX_MISC_SMM_MONITOR_CTL_B2 =                             0x10000000 //col:1062
IA32_VMX_MISC_VMWRITE_VMEXIT_INFO =                            0x20000000 //col:1063
IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION =      0x40000000 //col:1064
IA32_VMX_MISC_MSEG_ID =                                        0xFFFFFFFF00000000 //col:1065
IA32_VMX_CR0_FIXED0 =                                          0x00000486 //col:1066
IA32_VMX_CR0_FIXED1 =                                          0x00000487 //col:1067
IA32_VMX_CR4_FIXED0 =                                          0x00000488 //col:1068
IA32_VMX_CR4_FIXED1 =                                          0x00000489 //col:1069
IA32_VMX_VMCS_ENUM =                                           0x0000048A //col:1070
IA32_VMX_VMCS_ENUM_ACCESS_TYPE =                               0x01 //col:1071
IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE =                       0x3FE //col:1072
IA32_VMX_VMCS_ENUM_FIELD_TYPE =                                0xC00 //col:1073
IA32_VMX_VMCS_ENUM_FIELD_WIDTH =                               0x6000 //col:1074
IA32_VMX_PROCBASED_CTLS2 =                                     0x0000048B //col:1075
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES =            0x01 //col:1076
IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT =                          0x02 //col:1077
IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING =            0x04 //col:1078
IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP =                       0x08 //col:1079
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE =              0x10 //col:1080
IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID =                         0x20 //col:1081
IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING =                      0x40 //col:1082
IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST =                  0x80 //col:1083
IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION =        0x100 //col:1084
IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY =          0x200 //col:1085
IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING =                  0x400 //col:1086
IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING =                      0x800 //col:1087
IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID =                      0x1000 //col:1088
IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS =                 0x2000 //col:1089
IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING =                      0x4000 //col:1090
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING =                0x8000 //col:1091
IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING =                      0x10000 //col:1092
IA32_VMX_PROCBASED_CTLS2_ENABLE_PML =                          0x20000 //col:1093
IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION =                       0x40000 //col:1094
IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT =                 0x80000 //col:1095
IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES =                       0x100000 //col:1096
IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT =  0x400000 //col:1097
IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT =  0x800000 //col:1098
IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES =    0x1000000 //col:1099
IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING =                     0x2000000 //col:1100
IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE =              0x4000000 //col:1101
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING =                0x10000000 //col:1102
IA32_VMX_EPT_VPID_CAP =                                        0x0000048C //col:1103
IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES =                     0x01 //col:1104
IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4 =                     0x40 //col:1105
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE =                0x100 //col:1106
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK =                 0x4000 //col:1107
IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES =                          0x10000 //col:1108
IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES =                        0x20000 //col:1109
IA32_VMX_EPT_VPID_CAP_INVEPT =                                 0x100000 //col:1110
IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS =           0x200000 //col:1111
IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION = 0x400000 //col:1112
IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK =                0x800000 //col:1113
IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT =                  0x2000000 //col:1114
IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS =                    0x4000000 //col:1115
IA32_VMX_EPT_VPID_CAP_INVVPID =                                0x100000000 //col:1116
IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS =             0x10000000000 //col:1117
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT =                 0x20000000000 //col:1118
IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS =                   0x40000000000 //col:1119
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS =  0x80000000000 //col:1120
IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE =                   0x3F000000000000 //col:1121
IA32_VMX_TRUE_PINBASED_CTLS =                                  0x0000048D //col:1122
IA32_VMX_TRUE_PROCBASED_CTLS =                                 0x0000048E //col:1123
IA32_VMX_TRUE_EXIT_CTLS =                                      0x0000048F //col:1124
IA32_VMX_TRUE_ENTRY_CTLS =                                     0x00000490 //col:1125
IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS =                        0xFFFFFFFF //col:1126
IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS =                        0xFFFFFFFF00000000 //col:1127
IA32_VMX_VMFUNC =                                              0x00000491 //col:1128
IA32_VMX_VMFUNC_EPTP_SWITCHING =                               0x01 //col:1129
IA32_VMX_PROCBASED_CTLS3 =                                     0x00000492 //col:1130
IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING =                   0x01 //col:1131
IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT =                         0x02 //col:1132
IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE =                    0x04 //col:1133
IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING =                        0x08 //col:1134
IA32_VMX_EXIT_CTLS2 =                                          0x00000493 //col:1135
IA32_VMX_EXIT_CTLS2_RESERVED =                                 0xFFFFFFFFFFFFFFFF //col:1136
IA32_A_PMC0 =                                                  0x000004C1 //col:1137
IA32_A_PMC1 =                                                  0x000004C2 //col:1138
IA32_A_PMC2 =                                                  0x000004C3 //col:1139
IA32_A_PMC3 =                                                  0x000004C4 //col:1140
IA32_A_PMC4 =                                                  0x000004C5 //col:1141
IA32_A_PMC5 =                                                  0x000004C6 //col:1142
IA32_A_PMC6 =                                                  0x000004C7 //col:1143
IA32_A_PMC7 =                                                  0x000004C8 //col:1144
IA32_MCG_EXT_CTL =                                             0x000004D0 //col:1145
IA32_MCG_EXT_CTL_LMCE_EN =                                     0x01 //col:1146
IA32_SGX_SVN_STATUS =                                          0x00000500 //col:1147
IA32_SGX_SVN_STATUS_LOCK =                                     0x01 //col:1148
IA32_SGX_SVN_STATUS_SGX_SVN_SINIT =                            0xFF0000 //col:1149
IA32_RTIT_OUTPUT_BASE =                                        0x00000560 //col:1150
IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS =                  0xFFFFFFFFFF80 //col:1151
IA32_RTIT_OUTPUT_MASK_PTRS =                                   0x00000561 //col:1152
IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK =                        0x7F //col:1153
IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET =              0xFFFFFF80 //col:1154
IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET =                     0xFFFFFFFF00000000 //col:1155
IA32_RTIT_CTL =                                                0x00000570 //col:1156
IA32_RTIT_CTL_TRACE_EN =                                       0x01 //col:1157
IA32_RTIT_CTL_CYC_EN =                                         0x02 //col:1158
IA32_RTIT_CTL_OS =                                             0x04 //col:1159
IA32_RTIT_CTL_USER =                                           0x08 //col:1160
IA32_RTIT_CTL_PWR_EVT_EN =                                     0x10 //col:1161
IA32_RTIT_CTL_FUP_ON_PTW =                                     0x20 //col:1162
IA32_RTIT_CTL_FABRIC_EN =                                      0x40 //col:1163
IA32_RTIT_CTL_CR3_FILTER =                                     0x80 //col:1164
IA32_RTIT_CTL_TOPA =                                           0x100 //col:1165
IA32_RTIT_CTL_MTC_EN =                                         0x200 //col:1166
IA32_RTIT_CTL_TSC_EN =                                         0x400 //col:1167
IA32_RTIT_CTL_DIS_RETC =                                       0x800 //col:1168
IA32_RTIT_CTL_PTW_EN =                                         0x1000 //col:1169
IA32_RTIT_CTL_BRANCH_EN =                                      0x2000 //col:1170
IA32_RTIT_CTL_MTC_FREQ =                                       0x3C000 //col:1171
IA32_RTIT_CTL_CYC_THRESH =                                     0x780000 //col:1172
IA32_RTIT_CTL_PSB_FREQ =                                       0xF000000 //col:1173
IA32_RTIT_CTL_ADDR0_CFG =                                      0xF00000000 //col:1174
IA32_RTIT_CTL_ADDR1_CFG =                                      0xF000000000 //col:1175
IA32_RTIT_CTL_ADDR2_CFG =                                      0xF0000000000 //col:1176
IA32_RTIT_CTL_ADDR3_CFG =                                      0xF00000000000 //col:1177
IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE =                       0x100000000000000 //col:1178
IA32_RTIT_STATUS =                                             0x00000571 //col:1179
IA32_RTIT_STATUS_FILTER_EN =                                   0x01 //col:1180
IA32_RTIT_STATUS_CONTEX_EN =                                   0x02 //col:1181
IA32_RTIT_STATUS_TRIGGER_EN =                                  0x04 //col:1182
IA32_RTIT_STATUS_ERROR =                                       0x10 //col:1183
IA32_RTIT_STATUS_STOPPED =                                     0x20 //col:1184
IA32_RTIT_STATUS_PEND_PSB =                                    0x40 //col:1185
IA32_RTIT_STATUS_PEND_TOPA_PMI =                               0x80 //col:1186
IA32_RTIT_STATUS_PACKET_BYTE_CNT =                             0x1FFFF00000000 //col:1187
IA32_RTIT_CR3_MATCH =                                          0x00000572 //col:1188
IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH =                       0xFFFFFFFFFFFFFFE0 //col:1189
IA32_RTIT_ADDR0_A =                                            0x00000580 //col:1190
IA32_RTIT_ADDR1_A =                                            0x00000582 //col:1191
IA32_RTIT_ADDR2_A =                                            0x00000584 //col:1192
IA32_RTIT_ADDR3_A =                                            0x00000586 //col:1193
IA32_RTIT_ADDR0_B =                                            0x00000581 //col:1194
IA32_RTIT_ADDR1_B =                                            0x00000583 //col:1195
IA32_RTIT_ADDR2_B =                                            0x00000585 //col:1196
IA32_RTIT_ADDR3_B =                                            0x00000587 //col:1197
IA32_RTIT_ADDR_VIRTUAL_ADDRESS =                               0xFFFFFFFFFFFF //col:1198
IA32_RTIT_ADDR_SIGN_EXT_VA =                                   0xFFFF000000000000 //col:1199
IA32_DS_AREA =                                                 0x00000600 //col:1200
IA32_U_CET =                                                   0x000006A0 //col:1201
IA32_U_CET_SH_STK_EN =                                         0x01 //col:1202
IA32_U_CET_WR_SHSTK_EN =                                       0x02 //col:1203
IA32_U_CET_ENDBR_EN =                                          0x04 //col:1204
IA32_U_CET_LEG_IW_EN =                                         0x08 //col:1205
IA32_U_CET_NO_TRACK_EN =                                       0x10 //col:1206
IA32_U_CET_SUPPRESS_DIS =                                      0x20 //col:1207
IA32_U_CET_SUPPRESS =                                          0x400 //col:1208
IA32_U_CET_TRACKER =                                           0x800 //col:1209
IA32_U_CET_EB_LEG_BITMAP_BASE =                                0xFFFFFFFFFFFFF000 //col:1210
IA32_S_CET =                                                   0x000006A2 //col:1211
IA32_S_CET_SH_STK_EN =                                         0x01 //col:1212
IA32_S_CET_WR_SHSTK_EN =                                       0x02 //col:1213
IA32_S_CET_ENDBR_EN =                                          0x04 //col:1214
IA32_S_CET_LEG_IW_EN =                                         0x08 //col:1215
IA32_S_CET_NO_TRACK_EN =                                       0x10 //col:1216
IA32_S_CET_SUPPRESS_DIS =                                      0x20 //col:1217
IA32_S_CET_SUPPRESS =                                          0x400 //col:1218
IA32_S_CET_TRACKER =                                           0x800 //col:1219
IA32_S_CET_EB_LEG_BITMAP_BASE =                                0xFFFFFFFFFFFFF000 //col:1220
IA32_PL0_SSP =                                                 0x000006A4 //col:1221
IA32_PL1_SSP =                                                 0x000006A5 //col:1222
IA32_PL2_SSP =                                                 0x000006A6 //col:1223
IA32_PL3_SSP =                                                 0x000006A7 //col:1224
IA32_INTERRUPT_SSP_TABLE_ADDR =                                0x000006A8 //col:1225
IA32_TSC_DEADLINE =                                            0x000006E0 //col:1226
IA32_PM_ENABLE =                                               0x00000770 //col:1227
IA32_PM_ENABLE_HWP_ENABLE =                                    0x01 //col:1228
IA32_HWP_CAPABILITIES =                                        0x00000771 //col:1229
IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE =                    0xFF //col:1230
IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE =                 0xFF00 //col:1231
IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE =             0xFF0000 //col:1232
IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE =                     0xFF000000 //col:1233
IA32_HWP_REQUEST_PKG =                                         0x00000772 //col:1234
IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE =                     0xFF //col:1235
IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE =                     0xFF00 //col:1236
IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE =                     0xFF0000 //col:1237
IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE =           0xFF000000 //col:1238
IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW =                         0x3FF00000000 //col:1239
IA32_HWP_INTERRUPT =                                           0x00000773 //col:1240
IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE =          0x01 //col:1241
IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM =                      0x02 //col:1242
IA32_HWP_REQUEST =                                             0x00000774 //col:1243
IA32_HWP_REQUEST_MINIMUM_PERFORMANCE =                         0xFF //col:1244
IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE =                         0xFF00 //col:1245
IA32_HWP_REQUEST_DESIRED_PERFORMANCE =                         0xFF0000 //col:1246
IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE =               0xFF000000 //col:1247
IA32_HWP_REQUEST_ACTIVITY_WINDOW =                             0x3FF00000000 //col:1248
IA32_HWP_REQUEST_PACKAGE_CONTROL =                             0x40000000000 //col:1249
IA32_HWP_STATUS =                                              0x00000777 //col:1250
IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE =                0x01 //col:1251
IA32_HWP_STATUS_EXCURSION_TO_MINIMUM =                         0x04 //col:1252
IA32_X2APIC_APICID =                                           0x00000802 //col:1253
IA32_X2APIC_VERSION =                                          0x00000803 //col:1254
IA32_X2APIC_TPR =                                              0x00000808 //col:1255
IA32_X2APIC_PPR =                                              0x0000080A //col:1256
IA32_X2APIC_EOI =                                              0x0000080B //col:1257
IA32_X2APIC_LDR =                                              0x0000080D //col:1258
IA32_X2APIC_SIVR =                                             0x0000080F //col:1259
IA32_X2APIC_ISR0 =                                             0x00000810 //col:1260
IA32_X2APIC_ISR1 =                                             0x00000811 //col:1261
IA32_X2APIC_ISR2 =                                             0x00000812 //col:1262
IA32_X2APIC_ISR3 =                                             0x00000813 //col:1263
IA32_X2APIC_ISR4 =                                             0x00000814 //col:1264
IA32_X2APIC_ISR5 =                                             0x00000815 //col:1265
IA32_X2APIC_ISR6 =                                             0x00000816 //col:1266
IA32_X2APIC_ISR7 =                                             0x00000817 //col:1267
IA32_X2APIC_TMR0 =                                             0x00000818 //col:1268
IA32_X2APIC_TMR1 =                                             0x00000819 //col:1269
IA32_X2APIC_TMR2 =                                             0x0000081A //col:1270
IA32_X2APIC_TMR3 =                                             0x0000081B //col:1271
IA32_X2APIC_TMR4 =                                             0x0000081C //col:1272
IA32_X2APIC_TMR5 =                                             0x0000081D //col:1273
IA32_X2APIC_TMR6 =                                             0x0000081E //col:1274
IA32_X2APIC_TMR7 =                                             0x0000081F //col:1275
IA32_X2APIC_IRR0 =                                             0x00000820 //col:1276
IA32_X2APIC_IRR1 =                                             0x00000821 //col:1277
IA32_X2APIC_IRR2 =                                             0x00000822 //col:1278
IA32_X2APIC_IRR3 =                                             0x00000823 //col:1279
IA32_X2APIC_IRR4 =                                             0x00000824 //col:1280
IA32_X2APIC_IRR5 =                                             0x00000825 //col:1281
IA32_X2APIC_IRR6 =                                             0x00000826 //col:1282
IA32_X2APIC_IRR7 =                                             0x00000827 //col:1283
IA32_X2APIC_ESR =                                              0x00000828 //col:1284
IA32_X2APIC_LVT_CMCI =                                         0x0000082F //col:1285
IA32_X2APIC_ICR =                                              0x00000830 //col:1286
IA32_X2APIC_LVT_TIMER =                                        0x00000832 //col:1287
IA32_X2APIC_LVT_THERMAL =                                      0x00000833 //col:1288
IA32_X2APIC_LVT_PMI =                                          0x00000834 //col:1289
IA32_X2APIC_LVT_LINT0 =                                        0x00000835 //col:1290
IA32_X2APIC_LVT_LINT1 =                                        0x00000836 //col:1291
IA32_X2APIC_LVT_ERROR =                                        0x00000837 //col:1292
IA32_X2APIC_INIT_COUNT =                                       0x00000838 //col:1293
IA32_X2APIC_CUR_COUNT =                                        0x00000839 //col:1294
IA32_X2APIC_DIV_CONF =                                         0x0000083E //col:1295
IA32_X2APIC_SELF_IPI =                                         0x0000083F //col:1296
IA32_DEBUG_INTERFACE =                                         0x00000C80 //col:1297
IA32_DEBUG_INTERFACE_ENABLE =                                  0x01 //col:1298
IA32_DEBUG_INTERFACE_LOCK =                                    0x40000000 //col:1299
IA32_DEBUG_INTERFACE_DEBUG_OCCURRED =                          0x80000000 //col:1300
IA32_L3_QOS_CFG =                                              0x00000C81 //col:1301
IA32_L3_QOS_CFG_ENABLE =                                       0x01 //col:1302
IA32_L2_QOS_CFG =                                              0x00000C82 //col:1303
IA32_L2_QOS_CFG_ENABLE =                                       0x01 //col:1304
IA32_QM_EVTSEL =                                               0x00000C8D //col:1305
IA32_QM_EVTSEL_EVENT_ID =                                      0xFF //col:1306
IA32_QM_EVTSEL_RESOURCE_MONITORING_ID =                        0xFFFFFFFF00000000 //col:1307
IA32_QM_CTR =                                                  0x00000C8E //col:1308
IA32_QM_CTR_RESOURCE_MONITORED_DATA =                          0x3FFFFFFFFFFFFFFF //col:1309
IA32_QM_CTR_UNAVAILABLE =                                      0x4000000000000000 //col:1310
IA32_QM_CTR_ERROR =                                            0x8000000000000000 //col:1311
IA32_PQR_ASSOC =                                               0x00000C8F //col:1312
IA32_PQR_ASSOC_RESOURCE_MONITORING_ID =                        0xFFFFFFFF //col:1313
IA32_PQR_ASSOC_COS =                                           0xFFFFFFFF00000000 //col:1314
IA32_BNDCFGS =                                                 0x00000D90 //col:1315
IA32_BNDCFGS_ENABLE =                                          0x01 //col:1316
IA32_BNDCFGS_BND_PRESERVE =                                    0x02 //col:1317
IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS =                    0xFFFFFFFFFFFFF000 //col:1318
IA32_XSS =                                                     0x00000DA0 //col:1319
IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE =                    0x100 //col:1320
IA32_PKG_HDC_CTL =                                             0x00000DB0 //col:1321
IA32_PKG_HDC_CTL_HDC_PKG_ENABLE =                              0x01 //col:1322
IA32_PM_CTL1 =                                                 0x00000DB1 //col:1323
IA32_PM_CTL1_HDC_ALLOW_BLOCK =                                 0x01 //col:1324
IA32_THREAD_STALL =                                            0x00000DB2 //col:1325
IA32_EFER =                                                    0xC0000080 //col:1326
IA32_EFER_SYSCALL_ENABLE =                                     0x01 //col:1327
IA32_EFER_IA32E_MODE_ENABLE =                                  0x100 //col:1328
IA32_EFER_IA32E_MODE_ACTIVE =                                  0x400 //col:1329
IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE =                         0x800 //col:1330
IA32_STAR =                                                    0xC0000081 //col:1331
IA32_LSTAR =                                                   0xC0000082 //col:1332
IA32_CSTAR =                                                   0xC0000083 //col:1333
IA32_FMASK =                                                   0xC0000084 //col:1334
IA32_FS_BASE =                                                 0xC0000100 //col:1335
IA32_GS_BASE =                                                 0xC0000101 //col:1336
IA32_KERNEL_GS_BASE =                                          0xC0000102 //col:1337
IA32_TSC_AUX =                                                 0xC0000103 //col:1338
IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE =                         0xFFFFFFFF //col:1339
PDE_4MB_32_PRESENT =                                           0x01 //col:1340
PDE_4MB_32_WRITE =                                             0x02 //col:1341
PDE_4MB_32_SUPERVISOR =                                        0x04 //col:1342
PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH =                          0x08 //col:1343
PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE =                          0x10 //col:1344
PDE_4MB_32_ACCESSED =                                          0x20 //col:1345
PDE_4MB_32_DIRTY =                                             0x40 //col:1346
PDE_4MB_32_LARGE_PAGE =                                        0x80 //col:1347
PDE_4MB_32_GLOBAL =                                            0x100 //col:1348
PDE_4MB_32_IGNORED_1 =                                         0xE00 //col:1349
PDE_4MB_32_PAT =                                               0x1000 //col:1350
PDE_4MB_32_PAGE_FRAME_NUMBER_LOW =                             0x1FE000 //col:1351
PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH =                            0xFFC00000 //col:1352
PDE_32_PRESENT =                                               0x01 //col:1353
PDE_32_WRITE =                                                 0x02 //col:1354
PDE_32_SUPERVISOR =                                            0x04 //col:1355
PDE_32_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:1356
PDE_32_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:1357
PDE_32_ACCESSED =                                              0x20 //col:1358
PDE_32_IGNORED_1 =                                             0x40 //col:1359
PDE_32_LARGE_PAGE =                                            0x80 //col:1360
PDE_32_IGNORED_2 =                                             0xF00 //col:1361
PDE_32_PAGE_FRAME_NUMBER =                                     0xFFFFF000 //col:1362
PTE_32_PRESENT =                                               0x01 //col:1363
PTE_32_WRITE =                                                 0x02 //col:1364
PTE_32_SUPERVISOR =                                            0x04 //col:1365
PTE_32_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:1366
PTE_32_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:1367
PTE_32_ACCESSED =                                              0x20 //col:1368
PTE_32_DIRTY =                                                 0x40 //col:1369
PTE_32_PAT =                                                   0x80 //col:1370
PTE_32_GLOBAL =                                                0x100 //col:1371
PTE_32_IGNORED_1 =                                             0xE00 //col:1372
PTE_32_PAGE_FRAME_NUMBER =                                     0xFFFFF000 //col:1373
PT_ENTRY_32_PRESENT =                                          0x01 //col:1374
PT_ENTRY_32_WRITE =                                            0x02 //col:1375
PT_ENTRY_32_SUPERVISOR =                                       0x04 //col:1376
PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH =                         0x08 //col:1377
PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE =                         0x10 //col:1378
PT_ENTRY_32_ACCESSED =                                         0x20 //col:1379
PT_ENTRY_32_DIRTY =                                            0x40 //col:1380
PT_ENTRY_32_LARGE_PAGE =                                       0x80 //col:1381
PT_ENTRY_32_GLOBAL =                                           0x100 //col:1382
PT_ENTRY_32_IGNORED_1 =                                        0xE00 //col:1383
PT_ENTRY_32_PAGE_FRAME_NUMBER =                                0xFFFFF000 //col:1384
PDE_ENTRY_COUNT_32 =                                           0x00000400 //col:1385
PTE_ENTRY_COUNT_32 =                                           0x00000400 //col:1386
PML4E_64_PRESENT =                                             0x01 //col:1387
PML4E_64_WRITE =                                               0x02 //col:1388
PML4E_64_SUPERVISOR =                                          0x04 //col:1389
PML4E_64_PAGE_LEVEL_WRITE_THROUGH =                            0x08 //col:1390
PML4E_64_PAGE_LEVEL_CACHE_DISABLE =                            0x10 //col:1391
PML4E_64_ACCESSED =                                            0x20 //col:1392
PML4E_64_MUST_BE_ZERO =                                        0x80 //col:1393
PML4E_64_IGNORED_1 =                                           0x700 //col:1394
PML4E_64_RESTART =                                             0x800 //col:1395
PML4E_64_PAGE_FRAME_NUMBER =                                   0xFFFFFFFFF000 //col:1396
PML4E_64_IGNORED_2 =                                           0x7FF0000000000000 //col:1397
PML4E_64_EXECUTE_DISABLE =                                     0x8000000000000000 //col:1398
PDPTE_1GB_64_PRESENT =                                         0x01 //col:1399
PDPTE_1GB_64_WRITE =                                           0x02 //col:1400
PDPTE_1GB_64_SUPERVISOR =                                      0x04 //col:1401
PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH =                        0x08 //col:1402
PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE =                        0x10 //col:1403
PDPTE_1GB_64_ACCESSED =                                        0x20 //col:1404
PDPTE_1GB_64_DIRTY =                                           0x40 //col:1405
PDPTE_1GB_64_LARGE_PAGE =                                      0x80 //col:1406
PDPTE_1GB_64_GLOBAL =                                          0x100 //col:1407
PDPTE_1GB_64_IGNORED_1 =                                       0x600 //col:1408
PDPTE_1GB_64_RESTART =                                         0x800 //col:1409
PDPTE_1GB_64_PAT =                                             0x1000 //col:1410
PDPTE_1GB_64_PAGE_FRAME_NUMBER =                               0xFFFFC0000000 //col:1411
PDPTE_1GB_64_IGNORED_2 =                                       0x7F0000000000000 //col:1412
PDPTE_1GB_64_PROTECTION_KEY =                                  0x7800000000000000 //col:1413
PDPTE_1GB_64_EXECUTE_DISABLE =                                 0x8000000000000000 //col:1414
PDPTE_64_PRESENT =                                             0x01 //col:1415
PDPTE_64_WRITE =                                               0x02 //col:1416
PDPTE_64_SUPERVISOR =                                          0x04 //col:1417
PDPTE_64_PAGE_LEVEL_WRITE_THROUGH =                            0x08 //col:1418
PDPTE_64_PAGE_LEVEL_CACHE_DISABLE =                            0x10 //col:1419
PDPTE_64_ACCESSED =                                            0x20 //col:1420
PDPTE_64_LARGE_PAGE =                                          0x80 //col:1421
PDPTE_64_IGNORED_1 =                                           0x700 //col:1422
PDPTE_64_RESTART =                                             0x800 //col:1423
PDPTE_64_PAGE_FRAME_NUMBER =                                   0xFFFFFFFFF000 //col:1424
PDPTE_64_IGNORED_2 =                                           0x7FF0000000000000 //col:1425
PDPTE_64_EXECUTE_DISABLE =                                     0x8000000000000000 //col:1426
PDE_2MB_64_PRESENT =                                           0x01 //col:1427
PDE_2MB_64_WRITE =                                             0x02 //col:1428
PDE_2MB_64_SUPERVISOR =                                        0x04 //col:1429
PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH =                          0x08 //col:1430
PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE =                          0x10 //col:1431
PDE_2MB_64_ACCESSED =                                          0x20 //col:1432
PDE_2MB_64_DIRTY =                                             0x40 //col:1433
PDE_2MB_64_LARGE_PAGE =                                        0x80 //col:1434
PDE_2MB_64_GLOBAL =                                            0x100 //col:1435
PDE_2MB_64_IGNORED_1 =                                         0x600 //col:1436
PDE_2MB_64_RESTART =                                           0x800 //col:1437
PDE_2MB_64_PAT =                                               0x1000 //col:1438
PDE_2MB_64_PAGE_FRAME_NUMBER =                                 0xFFFFFFE00000 //col:1439
PDE_2MB_64_IGNORED_2 =                                         0x7F0000000000000 //col:1440
PDE_2MB_64_PROTECTION_KEY =                                    0x7800000000000000 //col:1441
PDE_2MB_64_EXECUTE_DISABLE =                                   0x8000000000000000 //col:1442
PDE_64_PRESENT =                                               0x01 //col:1443
PDE_64_WRITE =                                                 0x02 //col:1444
PDE_64_SUPERVISOR =                                            0x04 //col:1445
PDE_64_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:1446
PDE_64_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:1447
PDE_64_ACCESSED =                                              0x20 //col:1448
PDE_64_LARGE_PAGE =                                            0x80 //col:1449
PDE_64_IGNORED_1 =                                             0x700 //col:1450
PDE_64_RESTART =                                               0x800 //col:1451
PDE_64_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:1452
PDE_64_IGNORED_2 =                                             0x7FF0000000000000 //col:1453
PDE_64_EXECUTE_DISABLE =                                       0x8000000000000000 //col:1454
PTE_64_PRESENT =                                               0x01 //col:1455
PTE_64_WRITE =                                                 0x02 //col:1456
PTE_64_SUPERVISOR =                                            0x04 //col:1457
PTE_64_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:1458
PTE_64_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:1459
PTE_64_ACCESSED =                                              0x20 //col:1460
PTE_64_DIRTY =                                                 0x40 //col:1461
PTE_64_PAT =                                                   0x80 //col:1462
PTE_64_GLOBAL =                                                0x100 //col:1463
PTE_64_IGNORED_1 =                                             0x600 //col:1464
PTE_64_RESTART =                                               0x800 //col:1465
PTE_64_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:1466
PTE_64_IGNORED_2 =                                             0x7F0000000000000 //col:1467
PTE_64_PROTECTION_KEY =                                        0x7800000000000000 //col:1468
PTE_64_EXECUTE_DISABLE =                                       0x8000000000000000 //col:1469
PT_ENTRY_64_PRESENT =                                          0x01 //col:1470
PT_ENTRY_64_WRITE =                                            0x02 //col:1471
PT_ENTRY_64_SUPERVISOR =                                       0x04 //col:1472
PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH =                         0x08 //col:1473
PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE =                         0x10 //col:1474
PT_ENTRY_64_ACCESSED =                                         0x20 //col:1475
PT_ENTRY_64_DIRTY =                                            0x40 //col:1476
PT_ENTRY_64_LARGE_PAGE =                                       0x80 //col:1477
PT_ENTRY_64_GLOBAL =                                           0x100 //col:1478
PT_ENTRY_64_IGNORED_1 =                                        0x600 //col:1479
PT_ENTRY_64_RESTART =                                          0x800 //col:1480
PT_ENTRY_64_PAGE_FRAME_NUMBER =                                0xFFFFFFFFF000 //col:1481
PT_ENTRY_64_IGNORED_2 =                                        0x7F0000000000000 //col:1482
PT_ENTRY_64_PROTECTION_KEY =                                   0x7800000000000000 //col:1483
PT_ENTRY_64_EXECUTE_DISABLE =                                  0x8000000000000000 //col:1484
PML4_ENTRY_COUNT_64 =                                          0x00000200 //col:1485
PDPTE_ENTRY_COUNT_64 =                                         0x00000200 //col:1486
PDE_ENTRY_COUNT_64 =                                           0x00000200 //col:1487
PTE_ENTRY_COUNT_64 =                                           0x00000200 //col:1488
INVPCID_INDIVIDUAL_ADDRESS =                                   0x00000000 //col:1489
INVPCID_SINGLE_CONTEXT =                                       0x00000001 //col:1490
INVPCID_ALL_CONTEXT_WITH_GLOBALS =                             0x00000002 //col:1491
INVPCID_ALL_CONTEXT =                                          0x00000003 //col:1492
INVPCID_DESCRIPTOR_PCID =                                      0xFFF //col:1493
INVPCID_DESCRIPTOR_RESERVED1 =                                 0xFFFFFFFFFFFFF000 //col:1494
INVPCID_DESCRIPTOR_LINEAR_ADDRESS =                            0xFFFFFFFFFFFFFFFF0000000000000000 //col:1495
SEGMENT_ACCESS_RIGHTS_TYPE =                                   0xF00 //col:1496
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE =                        0x1000 //col:1497
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL =             0x6000 //col:1498
SEGMENT_ACCESS_RIGHTS_PRESENT =                                0x8000 //col:1499
SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT =                          0x100000 //col:1500
SEGMENT_ACCESS_RIGHTS_LONG_MODE =                              0x200000 //col:1501
SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG =                            0x400000 //col:1502
SEGMENT_ACCESS_RIGHTS_GRANULARITY =                            0x800000 //col:1503
SEGMENT__BASE_ADDRESS_MIDDLE =                                 0xFF //col:1504
SEGMENT__TYPE =                                                0xF00 //col:1505
SEGMENT__DESCRIPTOR_TYPE =                                     0x1000 //col:1506
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL =                          0x6000 //col:1507
SEGMENT__PRESENT =                                             0x8000 //col:1508
SEGMENT__SEGMENT_LIMIT_HIGH =                                  0xF0000 //col:1509
SEGMENT__AVAILABLE_BIT =                                       0x100000 //col:1510
SEGMENT__LONG_MODE =                                           0x200000 //col:1511
SEGMENT__DEFAULT_BIG =                                         0x400000 //col:1512
SEGMENT__GRANULARITY =                                         0x800000 //col:1513
SEGMENT__BASE_ADDRESS_HIGH =                                   0xFF000000 //col:1514
SEGMENT__BASE_ADDRESS_MIDDLE =                                 0xFF //col:1515
SEGMENT__TYPE =                                                0xF00 //col:1516
SEGMENT__DESCRIPTOR_TYPE =                                     0x1000 //col:1517
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL =                          0x6000 //col:1518
SEGMENT__PRESENT =                                             0x8000 //col:1519
SEGMENT__SEGMENT_LIMIT_HIGH =                                  0xF0000 //col:1520
SEGMENT__AVAILABLE_BIT =                                       0x100000 //col:1521
SEGMENT__LONG_MODE =                                           0x200000 //col:1522
SEGMENT__DEFAULT_BIG =                                         0x400000 //col:1523
SEGMENT__GRANULARITY =                                         0x800000 //col:1524
SEGMENT__BASE_ADDRESS_HIGH =                                   0xFF000000 //col:1525
SEGMENT__INTERRUPT_STACK_TABLE =                               0x07 //col:1526
SEGMENT__MUST_BE_ZERO_0 =                                      0xF8 //col:1527
SEGMENT__TYPE =                                                0xF00 //col:1528
SEGMENT__MUST_BE_ZERO_1 =                                      0x1000 //col:1529
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL =                          0x6000 //col:1530
SEGMENT__PRESENT =                                             0x8000 //col:1531
SEGMENT__OFFSET_MIDDLE =                                       0xFFFF0000 //col:1532
SEGMENT_DESCRIPTOR_TYPE_SYSTEM =                               0x00000000 //col:1533
SEGMENT_DESCRIPTOR_TYPE_CODE_OR_DATA =                         0x00000001 //col:1534
SEGMENT_DESCRIPTOR_TYPE_DATA_R =                               0x00000000 //col:1535
SEGMENT_DESCRIPTOR_TYPE_DATA_RA =                              0x00000001 //col:1536
SEGMENT_DESCRIPTOR_TYPE_DATA_RW =                              0x00000002 //col:1537
SEGMENT_DESCRIPTOR_TYPE_DATA_RWA =                             0x00000003 //col:1538
SEGMENT_DESCRIPTOR_TYPE_DATA_RE =                              0x00000004 //col:1539
SEGMENT_DESCRIPTOR_TYPE_DATA_REA =                             0x00000005 //col:1540
SEGMENT_DESCRIPTOR_TYPE_DATA_RWE =                             0x00000006 //col:1541
SEGMENT_DESCRIPTOR_TYPE_DATA_RWEA =                            0x00000007 //col:1542
SEGMENT_DESCRIPTOR_TYPE_CODE_E =                               0x00000008 //col:1543
SEGMENT_DESCRIPTOR_TYPE_CODE_EA =                              0x00000009 //col:1544
SEGMENT_DESCRIPTOR_TYPE_CODE_ER =                              0x0000000A //col:1545
SEGMENT_DESCRIPTOR_TYPE_CODE_ERA =                             0x0000000B //col:1546
SEGMENT_DESCRIPTOR_TYPE_CODE_EC =                              0x0000000C //col:1547
SEGMENT_DESCRIPTOR_TYPE_CODE_ECA =                             0x0000000D //col:1548
SEGMENT_DESCRIPTOR_TYPE_CODE_ERC =                             0x0000000E //col:1549
SEGMENT_DESCRIPTOR_TYPE_CODE_ERCA =                            0x0000000F //col:1550
SEGMENT_DESCRIPTOR_TYPE_RESERVED_1 =                           0x00000000 //col:1551
SEGMENT_DESCRIPTOR_TYPE_TSS_16_AVAILABLE =                     0x00000001 //col:1552
SEGMENT_DESCRIPTOR_TYPE_LDT =                                  0x00000002 //col:1553
SEGMENT_DESCRIPTOR_TYPE_TSS_16_BUSY =                          0x00000003 //col:1554
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE_16 =                         0x00000004 //col:1555
SEGMENT_DESCRIPTOR_TYPE_TASK_GATE =                            0x00000005 //col:1556
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE_16 =                    0x00000006 //col:1557
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE_16 =                         0x00000007 //col:1558
SEGMENT_DESCRIPTOR_TYPE_RESERVED_2 =                           0x00000008 //col:1559
SEGMENT_DESCRIPTOR_TYPE_TSS_AVAILABLE =                        0x00000009 //col:1560
SEGMENT_DESCRIPTOR_TYPE_RESERVED_3 =                           0x0000000A //col:1561
SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY =                             0x0000000B //col:1562
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE =                            0x0000000C //col:1563
SEGMENT_DESCRIPTOR_TYPE_RESERVED_4 =                           0x0000000D //col:1564
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE =                       0x0000000E //col:1565
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE =                            0x0000000F //col:1566
SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL =                     0x03 //col:1567
SEGMENT_SELECTOR_TABLE_INDICATOR =                             0x04 //col:1568
SEGMENT_SELECTOR_INDEX =                                       0xFFF8 //col:1569
VMX_EXIT_REASON_XCPT_OR_NMI =                                  0x00000000 //col:1570
VMX_EXIT_REASON_EXT_INT =                                      0x00000001 //col:1571
VMX_EXIT_REASON_TRIPLE_FAULT =                                 0x00000002 //col:1572
VMX_EXIT_REASON_INIT_SIGNAL =                                  0x00000003 //col:1573
VMX_EXIT_REASON_SIPI =                                         0x00000004 //col:1574
VMX_EXIT_REASON_IO_SMI =                                       0x00000005 //col:1575
VMX_EXIT_REASON_SMI =                                          0x00000006 //col:1576
VMX_EXIT_REASON_INT_WINDOW =                                   0x00000007 //col:1577
VMX_EXIT_REASON_NMI_WINDOW =                                   0x00000008 //col:1578
VMX_EXIT_REASON_TASK_SWITCH =                                  0x00000009 //col:1579
VMX_EXIT_REASON_CPUID =                                        0x0000000A //col:1580
VMX_EXIT_REASON_GETSEC =                                       0x0000000B //col:1581
VMX_EXIT_REASON_HLT =                                          0x0000000C //col:1582
VMX_EXIT_REASON_INVD =                                         0x0000000D //col:1583
VMX_EXIT_REASON_INVLPG =                                       0x0000000E //col:1584
VMX_EXIT_REASON_RDPMC =                                        0x0000000F //col:1585
VMX_EXIT_REASON_RDTSC =                                        0x00000010 //col:1586
VMX_EXIT_REASON_RSM =                                          0x00000011 //col:1587
VMX_EXIT_REASON_VMCALL =                                       0x00000012 //col:1588
VMX_EXIT_REASON_VMCLEAR =                                      0x00000013 //col:1589
VMX_EXIT_REASON_VMLAUNCH =                                     0x00000014 //col:1590
VMX_EXIT_REASON_VMPTRLD =                                      0x00000015 //col:1591
VMX_EXIT_REASON_VMPTRST =                                      0x00000016 //col:1592
VMX_EXIT_REASON_VMREAD =                                       0x00000017 //col:1593
VMX_EXIT_REASON_VMRESUME =                                     0x00000018 //col:1594
VMX_EXIT_REASON_VMWRITE =                                      0x00000019 //col:1595
VMX_EXIT_REASON_VMXOFF =                                       0x0000001A //col:1596
VMX_EXIT_REASON_VMXON =                                        0x0000001B //col:1597
VMX_EXIT_REASON_MOV_CRX =                                      0x0000001C //col:1598
VMX_EXIT_REASON_MOV_DRX =                                      0x0000001D //col:1599
VMX_EXIT_REASON_IO_INSTR =                                     0x0000001E //col:1600
VMX_EXIT_REASON_RDMSR =                                        0x0000001F //col:1601
VMX_EXIT_REASON_WRMSR =                                        0x00000020 //col:1602
VMX_EXIT_REASON_ERR_INVALID_GUEST_STATE =                      0x00000021 //col:1603
VMX_EXIT_REASON_ERR_MSR_LOAD =                                 0x00000022 //col:1604
VMX_EXIT_REASON_MWAIT =                                        0x00000024 //col:1605
VMX_EXIT_REASON_MTF =                                          0x00000025 //col:1606
VMX_EXIT_REASON_MONITOR =                                      0x00000027 //col:1607
VMX_EXIT_REASON_PAUSE =                                        0x00000028 //col:1608
VMX_EXIT_REASON_ERR_MACHINE_CHECK =                            0x00000029 //col:1609
VMX_EXIT_REASON_TPR_BELOW_THRESHOLD =                          0x0000002B //col:1610
VMX_EXIT_REASON_APIC_ACCESS =                                  0x0000002C //col:1611
VMX_EXIT_REASON_VIRTUALIZED_EOI =                              0x0000002D //col:1612
VMX_EXIT_REASON_XDTR_ACCESS =                                  0x0000002E //col:1613
VMX_EXIT_REASON_TR_ACCESS =                                    0x0000002F //col:1614
VMX_EXIT_REASON_EPT_VIOLATION =                                0x00000030 //col:1615
VMX_EXIT_REASON_EPT_MISCONFIG =                                0x00000031 //col:1616
VMX_EXIT_REASON_INVEPT =                                       0x00000032 //col:1617
VMX_EXIT_REASON_RDTSCP =                                       0x00000033 //col:1618
VMX_EXIT_REASON_PREEMPT_TIMER =                                0x00000034 //col:1619
VMX_EXIT_REASON_INVVPID =                                      0x00000035 //col:1620
VMX_EXIT_REASON_WBINVD =                                       0x00000036 //col:1621
VMX_EXIT_REASON_XSETBV =                                       0x00000037 //col:1622
VMX_EXIT_REASON_APIC_WRITE =                                   0x00000038 //col:1623
VMX_EXIT_REASON_RDRAND =                                       0x00000039 //col:1624
VMX_EXIT_REASON_INVPCID =                                      0x0000003A //col:1625
VMX_EXIT_REASON_VMFUNC =                                       0x0000003B //col:1626
VMX_EXIT_REASON_ENCLS =                                        0x0000003C //col:1627
VMX_EXIT_REASON_RDSEED =                                       0x0000003D //col:1628
VMX_EXIT_REASON_PML_FULL =                                     0x0000003E //col:1629
VMX_EXIT_REASON_XSAVES =                                       0x0000003F //col:1630
VMX_EXIT_REASON_XRSTORS =                                      0x00000040 //col:1631
VMX_ERROR_VMCALL =                                             0x00000001 //col:1632
VMX_ERROR_VMCLEAR_INVALID_PHYS_ADDR =                          0x00000002 //col:1633
VMX_ERROR_VMCLEAR_INVALID_VMXON_PTR =                          0x00000003 //col:1634
VMX_ERROR_VMLAUCH_NON_CLEAR_VMCS =                             0x00000004 //col:1635
VMX_ERROR_VMRESUME_NON_LAUNCHED_VMCS =                         0x00000005 //col:1636
VMX_ERROR_VMRESUME_CORRUPTED_VMCS =                            0x00000006 //col:1637
VMX_ERROR_VMENTRY_INVALID_CONTROL_FIELDS =                     0x00000007 //col:1638
VMX_ERROR_VMENTRY_INVALID_HOST_STATE =                         0x00000008 //col:1639
VMX_ERROR_VMPTRLD_INVALID_PHYS_ADDR =                          0x00000009 //col:1640
VMX_ERROR_VMPTRLD_VMXON_PTR =                                  0x0000000A //col:1641
VMX_ERROR_VMPTRLD_WRONG_VMCS_REVISION =                        0x0000000B //col:1642
VMX_ERROR_VMREAD_VMWRITE_INVALID_COMPONENT =                   0x0000000C //col:1643
VMX_ERROR_VMWRITE_READONLY_COMPONENT =                         0x0000000D //col:1644
VMX_ERROR_VMXON_IN_VMX_ROOT_OP =                               0x0000000F //col:1645
VMX_ERROR_VMENTRY_INVALID_VMCS_EXEC_PTR =                      0x00000010 //col:1646
VMX_ERROR_VMENTRY_NON_LAUNCHED_EXEC_VMCS =                     0x00000011 //col:1647
VMX_ERROR_VMENTRY_EXEC_VMCS_PTR =                              0x00000012 //col:1648
VMX_ERROR_VMCALL_NON_CLEAR_VMCS =                              0x00000013 //col:1649
VMX_ERROR_VMCALL_INVALID_VMEXIT_FIELDS =                       0x00000014 //col:1650
VMX_ERROR_VMCALL_INVALID_MSEG_REVISION =                       0x00000016 //col:1651
VMX_ERROR_VMXOFF_DUAL_MONITOR =                                0x00000017 //col:1652
VMX_ERROR_VMCALL_INVALID_SMM_MONITOR =                         0x00000018 //col:1653
VMX_ERROR_VMENTRY_INVALID_VM_EXEC_CTRL =                       0x00000019 //col:1654
VMX_ERROR_VMENTRY_MOV_SS =                                     0x0000001A //col:1655
VMX_ERROR_INVEPTVPID_INVALID_OPERAND =                         0x0000001C //col:1656
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION =  0x0F //col:1657
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED = 0x2000 //col:1658
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION =    0x4000 //col:1659
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR =                  0xFFFF //col:1660
VMX_EXIT_QUALIFICATION_TASK_SWITCH_TYPE =                      0xC0000000 //col:1661
VMX_EXIT_QUALIFICATION_TYPE_CALL =                             0x00000000 //col:1662
VMX_EXIT_QUALIFICATION_TYPE_IRET =                             0x00000001 //col:1663
VMX_EXIT_QUALIFICATION_TYPE_JMP =                              0x00000002 //col:1664
VMX_EXIT_QUALIFICATION_TYPE_IDT =                              0x00000003 //col:1665
VMX_EXIT_QUALIFICATION_CR_ACCESS_CR_NUMBER =                   0x0F //col:1666
VMX_EXIT_QUALIFICATION_REGISTER_CR0 =                          0x00000000 //col:1667
VMX_EXIT_QUALIFICATION_REGISTER_CR2 =                          0x00000002 //col:1668
VMX_EXIT_QUALIFICATION_REGISTER_CR3 =                          0x00000003 //col:1669
VMX_EXIT_QUALIFICATION_REGISTER_CR4 =                          0x00000004 //col:1670
VMX_EXIT_QUALIFICATION_REGISTER_CR8 =                          0x00000008 //col:1671
VMX_EXIT_QUALIFICATION_CR_ACCESS_ACCESS_TYPE =                 0x30 //col:1672
VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR =                      0x00000000 //col:1673
VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR =                    0x00000001 //col:1674
VMX_EXIT_QUALIFICATION_ACCESS_CLTS =                           0x00000002 //col:1675
VMX_EXIT_QUALIFICATION_ACCESS_LMSW =                           0x00000003 //col:1676
VMX_EXIT_QUALIFICATION_CR_ACCESS_LMSW_OPERAND_TYPE =           0x40 //col:1677
VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER =                      0x00000000 //col:1678
VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY =                        0x00000001 //col:1679
VMX_EXIT_QUALIFICATION_CR_ACCESS_GP_REGISTER =                 0xF00 //col:1680
VMX_EXIT_QUALIFICATION_GENREG_RAX =                            0x00000000 //col:1681
VMX_EXIT_QUALIFICATION_GENREG_RCX =                            0x00000001 //col:1682
VMX_EXIT_QUALIFICATION_GENREG_RDX =                            0x00000002 //col:1683
VMX_EXIT_QUALIFICATION_GENREG_RBX =                            0x00000003 //col:1684
VMX_EXIT_QUALIFICATION_GENREG_RSP =                            0x00000004 //col:1685
VMX_EXIT_QUALIFICATION_GENREG_RBP =                            0x00000005 //col:1686
VMX_EXIT_QUALIFICATION_GENREG_RSI =                            0x00000006 //col:1687
VMX_EXIT_QUALIFICATION_GENREG_RDI =                            0x00000007 //col:1688
VMX_EXIT_QUALIFICATION_GENREG_R8 =                             0x00000008 //col:1689
VMX_EXIT_QUALIFICATION_GENREG_R9 =                             0x00000009 //col:1690
VMX_EXIT_QUALIFICATION_GENREG_R10 =                            0x0000000A //col:1691
VMX_EXIT_QUALIFICATION_GENREG_R11 =                            0x0000000B //col:1692
VMX_EXIT_QUALIFICATION_GENREG_R12 =                            0x0000000C //col:1693
VMX_EXIT_QUALIFICATION_GENREG_R13 =                            0x0000000D //col:1694
VMX_EXIT_QUALIFICATION_GENREG_R14 =                            0x0000000E //col:1695
VMX_EXIT_QUALIFICATION_GENREG_R15 =                            0x0000000F //col:1696
VMX_EXIT_QUALIFICATION_CR_ACCESS_LMSW_SOURCE_DATA =            0xFFFF0000 //col:1697
VMX_EXIT_QUALIFICATION_DR_ACCESS_DR_NUMBER =                   0x07 //col:1698
VMX_EXIT_QUALIFICATION_REGISTER_DR0 =                          0x00000000 //col:1699
VMX_EXIT_QUALIFICATION_REGISTER_DR1 =                          0x00000001 //col:1700
VMX_EXIT_QUALIFICATION_REGISTER_DR2 =                          0x00000002 //col:1701
VMX_EXIT_QUALIFICATION_REGISTER_DR3 =                          0x00000003 //col:1702
VMX_EXIT_QUALIFICATION_REGISTER_DR6 =                          0x00000006 //col:1703
VMX_EXIT_QUALIFICATION_REGISTER_DR7 =                          0x00000007 //col:1704
VMX_EXIT_QUALIFICATION_DR_ACCESS_DIRECTION_OF_ACCESS =         0x10 //col:1705
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR =                   0x00000000 //col:1706
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR =                 0x00000001 //col:1707
VMX_EXIT_QUALIFICATION_DR_ACCESS_GP_REGISTER =                 0xF00 //col:1708
VMX_EXIT_QUALIFICATION_IO_INST_SIZE_OF_ACCESS =                0x07 //col:1709
VMX_EXIT_QUALIFICATION_WIDTH_1B =                              0x00000000 //col:1710
VMX_EXIT_QUALIFICATION_WIDTH_2B =                              0x00000001 //col:1711
VMX_EXIT_QUALIFICATION_WIDTH_4B =                              0x00000003 //col:1712
VMX_EXIT_QUALIFICATION_IO_INST_DIRECTION_OF_ACCESS =           0x08 //col:1713
VMX_EXIT_QUALIFICATION_DIRECTION_OUT =                         0x00000000 //col:1714
VMX_EXIT_QUALIFICATION_DIRECTION_IN =                          0x00000001 //col:1715
VMX_EXIT_QUALIFICATION_IO_INST_STRING_INSTRUCTION =            0x10 //col:1716
VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING =                  0x00000000 //col:1717
VMX_EXIT_QUALIFICATION_IS_STRING_STRING =                      0x00000001 //col:1718
VMX_EXIT_QUALIFICATION_IO_INST_REP_PREFIXED =                  0x20 //col:1719
VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP =                        0x00000000 //col:1720
VMX_EXIT_QUALIFICATION_IS_REP_REP =                            0x00000001 //col:1721
VMX_EXIT_QUALIFICATION_IO_INST_OPERAND_ENCODING =              0x40 //col:1722
VMX_EXIT_QUALIFICATION_ENCODING_DX =                           0x00000000 //col:1723
VMX_EXIT_QUALIFICATION_ENCODING_IMM =                          0x00000001 //col:1724
VMX_EXIT_QUALIFICATION_IO_INST_PORT_NUMBER =                   0xFFFF0000 //col:1725
VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET =               0xFFF //col:1726
VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE =               0xF000 //col:1727
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_READ =                      0x00000000 //col:1728
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_WRITE =                     0x00000001 //col:1729
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_INSTR_FETCH =               0x00000002 //col:1730
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_EVENT_DELIVERY =            0x00000003 //col:1731
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_EVENT_DELIVERY =          0x0000000A //col:1732
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_INSTR =                   0x0000000F //col:1733
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_DATA_READ =               0x01 //col:1734
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_DATA_WRITE =              0x02 //col:1735
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_INSTRUCTION_FETCH =       0x04 //col:1736
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_PRESENT =           0x08 //col:1737
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_WRITE =             0x10 //col:1738
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_EXECUTE =           0x20 //col:1739
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_EXECUTE_FOR_USER_MODE = 0x40 //col:1740
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS = 0x80 //col:1741
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_TRANSLATED_ACCESS =   0x100 //col:1742
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS = 0x200 //col:1743
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE =  0x400 //col:1744
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE =    0x800 //col:1745
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING =          0x1000 //col:1746
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS =     0x2000 //col:1747
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK = 0x4000 //col:1748
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION = 0x8000 //col:1749
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION = 0x10000 //col:1750
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE =            0x380 //col:1751
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER =        0x38000 //col:1752
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING =               0x03 //col:1753
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE =          0x380 //col:1754
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER =      0x38000 //col:1755
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GP_REGISTER =           0x3C0000 //col:1756
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GP_REGISTER_INVALID =   0x400000 //col:1757
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER =         0x7800000 //col:1758
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID = 0x8000000 //col:1759
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2 =            0xF0000000 //col:1760
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING =         0x03 //col:1761
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE =    0x380 //col:1762
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE =    0x800 //col:1763
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER = 0x38000 //col:1764
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GP_REGISTER =     0x3C0000 //col:1765
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GP_REGISTER_INVALID = 0x400000 //col:1766
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER =   0x7800000 //col:1767
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID = 0x8000000 //col:1768
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_IDENTITY = 0x30000000 //col:1769
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING =           0x03 //col:1770
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1 =             0x78 //col:1771
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE =      0x380 //col:1772
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER =   0x400 //col:1773
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER =  0x38000 //col:1774
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GP_REGISTER =       0x3C0000 //col:1775
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GP_REGISTER_INVALID = 0x400000 //col:1776
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER =     0x7800000 //col:1777
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID = 0x8000000 //col:1778
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_IDENTITY = 0x30000000 //col:1779
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER = 0x78 //col:1780
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE =       0x1800 //col:1781
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING =           0x03 //col:1782
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE =      0x380 //col:1783
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER =  0x38000 //col:1784
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GP_REGISTER =       0x3C0000 //col:1785
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GP_REGISTER_INVALID = 0x400000 //col:1786
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER =     0x7800000 //col:1787
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID = 0x8000000 //col:1788
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING =           0x03 //col:1789
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1 =        0x78 //col:1790
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE =      0x380 //col:1791
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER =   0x400 //col:1792
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER =  0x38000 //col:1793
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GP_REGISTER =       0x3C0000 //col:1794
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GP_REGISTER_INVALID = 0x400000 //col:1795
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER =     0x7800000 //col:1796
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID = 0x8000000 //col:1797
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2 =        0xF0000000 //col:1798
VMX_SEGMENT_ACCESS_RIGHTS_TYPE =                               0x0F //col:1799
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE =                    0x10 //col:1800
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL =         0x60 //col:1801
VMX_SEGMENT_ACCESS_RIGHTS_PRESENT =                            0x80 //col:1802
VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT =                      0x1000 //col:1803
VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE =                          0x2000 //col:1804
VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG =                        0x4000 //col:1805
VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY =                        0x8000 //col:1806
VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE =                           0x10000 //col:1807
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI =                   0x01 //col:1808
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS =                0x02 //col:1809
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI =                   0x04 //col:1810
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI =                   0x08 //col:1811
VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION =              0x10 //col:1812
VMX_ACTIVE =                                                   0x00000000 //col:1813
VMX_HLT =                                                      0x00000001 //col:1814
VMX_SHUTDOWN =                                                 0x00000002 //col:1815
VMX_WAIT_FOR_SIPI =                                            0x00000003 //col:1816
VMX_PENDING_DEBUG_EXCEPTIONS_B0 =                              0x01 //col:1817
VMX_PENDING_DEBUG_EXCEPTIONS_B1 =                              0x02 //col:1818
VMX_PENDING_DEBUG_EXCEPTIONS_B2 =                              0x04 //col:1819
VMX_PENDING_DEBUG_EXCEPTIONS_B3 =                              0x08 //col:1820
VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT =              0x1000 //col:1821
VMX_PENDING_DEBUG_EXCEPTIONS_BS =                              0x4000 //col:1822
VMX_PENDING_DEBUG_EXCEPTIONS_RTM =                             0x10000 //col:1823
VMX_VMEXIT_REASON_BASIC_EXIT_REASON =                          0xFFFF //col:1824
VMX_VMEXIT_REASON_ALWAYS0 =                                    0x10000 //col:1825
VMX_VMEXIT_REASON_RESERVED1 =                                  0x7FE0000 //col:1826
VMX_VMEXIT_REASON_ENCLAVE_MODE =                               0x8000000 //col:1827
VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT =                        0x10000000 //col:1828
VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT =                      0x20000000 //col:1829
VMX_VMEXIT_REASON_RESERVED2 =                                  0x40000000 //col:1830
VMX_VMEXIT_REASON_VM_ENTRY_FAILURE =                           0x80000000 //col:1831
IO_BITMAP_A_MIN =                                              0x00000000 //col:1832
IO_BITMAP_A_MAX =                                              0x00007FFF //col:1833
IO_BITMAP_B_MIN =                                              0x00008000 //col:1834
IO_BITMAP_B_MAX =                                              0x0000FFFF //col:1835
MSR_ID_LOW_MIN =                                               0x00000000 //col:1836
MSR_ID_LOW_MAX =                                               0x00001FFF //col:1837
MSR_ID_HIGH_MIN =                                              0xC0000000 //col:1838
MSR_ID_HIGH_MAX =                                              0xC0001FFF //col:1839
EPTP_MEMORY_TYPE =                                             0x07 //col:1840
EPTP_PAGE_WALK_LENGTH =                                        0x38 //col:1841
EPT_PAGE_WALK_LENGTH_4 =                                       0x00000003 //col:1842
EPTP_ENABLE_ACCESS_AND_DIRTY_FLAGS =                           0x40 //col:1843
EPTP_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES =                    0x80 //col:1844
EPTP_PAGE_FRAME_NUMBER =                                       0xFFFFFFFFF000 //col:1845
EPML4E_READ_ACCESS =                                           0x01 //col:1846
EPML4E_WRITE_ACCESS =                                          0x02 //col:1847
EPML4E_EXECUTE_ACCESS =                                        0x04 //col:1848
EPML4E_ACCESSED =                                              0x100 //col:1849
EPML4E_USER_MODE_EXECUTE =                                     0x400 //col:1850
EPML4E_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:1851
EPDPTE_1GB_READ_ACCESS =                                       0x01 //col:1852
EPDPTE_1GB_WRITE_ACCESS =                                      0x02 //col:1853
EPDPTE_1GB_EXECUTE_ACCESS =                                    0x04 //col:1854
EPDPTE_1GB_MEMORY_TYPE =                                       0x38 //col:1855
EPDPTE_1GB_IGNORE_PAT =                                        0x40 //col:1856
EPDPTE_1GB_LARGE_PAGE =                                        0x80 //col:1857
EPDPTE_1GB_ACCESSED =                                          0x100 //col:1858
EPDPTE_1GB_DIRTY =                                             0x200 //col:1859
EPDPTE_1GB_USER_MODE_EXECUTE =                                 0x400 //col:1860
EPDPTE_1GB_PAGE_FRAME_NUMBER =                                 0xFFFFC0000000 //col:1861
EPDPTE_1GB_VERIFY_GUEST_PAGING =                               0x200000000000000 //col:1862
EPDPTE_1GB_PAGING_WRITE_ACCESS =                               0x400000000000000 //col:1863
EPDPTE_1GB_SUPERVISOR_SHADOW_STACK =                           0x1000000000000000 //col:1864
EPDPTE_1GB_SUPPRESS_VE =                                       0x8000000000000000 //col:1865
EPDPTE_READ_ACCESS =                                           0x01 //col:1866
EPDPTE_WRITE_ACCESS =                                          0x02 //col:1867
EPDPTE_EXECUTE_ACCESS =                                        0x04 //col:1868
EPDPTE_ACCESSED =                                              0x100 //col:1869
EPDPTE_USER_MODE_EXECUTE =                                     0x400 //col:1870
EPDPTE_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:1871
EPDE_2MB_READ_ACCESS =                                         0x01 //col:1872
EPDE_2MB_WRITE_ACCESS =                                        0x02 //col:1873
EPDE_2MB_EXECUTE_ACCESS =                                      0x04 //col:1874
EPDE_2MB_MEMORY_TYPE =                                         0x38 //col:1875
EPDE_2MB_IGNORE_PAT =                                          0x40 //col:1876
EPDE_2MB_LARGE_PAGE =                                          0x80 //col:1877
EPDE_2MB_ACCESSED =                                            0x100 //col:1878
EPDE_2MB_DIRTY =                                               0x200 //col:1879
EPDE_2MB_USER_MODE_EXECUTE =                                   0x400 //col:1880
EPDE_2MB_PAGE_FRAME_NUMBER =                                   0xFFFFFFE00000 //col:1881
EPDE_2MB_VERIFY_GUEST_PAGING =                                 0x200000000000000 //col:1882
EPDE_2MB_PAGING_WRITE_ACCESS =                                 0x400000000000000 //col:1883
EPDE_2MB_SUPERVISOR_SHADOW_STACK =                             0x1000000000000000 //col:1884
EPDE_2MB_SUPPRESS_VE =                                         0x8000000000000000 //col:1885
EPDE_READ_ACCESS =                                             0x01 //col:1886
EPDE_WRITE_ACCESS =                                            0x02 //col:1887
EPDE_EXECUTE_ACCESS =                                          0x04 //col:1888
EPDE_ACCESSED =                                                0x100 //col:1889
EPDE_USER_MODE_EXECUTE =                                       0x400 //col:1890
EPDE_PAGE_FRAME_NUMBER =                                       0xFFFFFFFFF000 //col:1891
EPTE_READ_ACCESS =                                             0x01 //col:1892
EPTE_WRITE_ACCESS =                                            0x02 //col:1893
EPTE_EXECUTE_ACCESS =                                          0x04 //col:1894
EPTE_MEMORY_TYPE =                                             0x38 //col:1895
EPTE_IGNORE_PAT =                                              0x40 //col:1896
EPTE_ACCESSED =                                                0x100 //col:1897
EPTE_DIRTY =                                                   0x200 //col:1898
EPTE_USER_MODE_EXECUTE =                                       0x400 //col:1899
EPTE_PAGE_FRAME_NUMBER =                                       0xFFFFFFFFF000 //col:1900
EPTE_VERIFY_GUEST_PAGING =                                     0x200000000000000 //col:1901
EPTE_PAGING_WRITE_ACCESS =                                     0x400000000000000 //col:1902
EPTE_SUPERVISOR_SHADOW_STACK =                                 0x1000000000000000 //col:1903
EPTE_SUB_PAGE_WRITE_PERMISSIONS =                              0x2000000000000000 //col:1904
EPTE_SUPPRESS_VE =                                             0x8000000000000000 //col:1905
EPT_ENTRY_READ_ACCESS =                                        0x01 //col:1906
EPT_ENTRY_WRITE_ACCESS =                                       0x02 //col:1907
EPT_ENTRY_EXECUTE_ACCESS =                                     0x04 //col:1908
EPT_ENTRY_MEMORY_TYPE =                                        0x38 //col:1909
EPT_ENTRY_IGNORE_PAT =                                         0x40 //col:1910
EPT_ENTRY_LARGE_PAGE =                                         0x80 //col:1911
EPT_ENTRY_ACCESSED =                                           0x100 //col:1912
EPT_ENTRY_DIRTY =                                              0x200 //col:1913
EPT_ENTRY_USER_MODE_EXECUTE =                                  0x400 //col:1914
EPT_ENTRY_PAGE_FRAME_NUMBER =                                  0xFFFFFFFFF000 //col:1915
EPT_ENTRY_SUPPRESS_VE =                                        0x8000000000000000 //col:1916
EPT_LEVEL_PML4E =                                              0x00000003 //col:1917
EPT_LEVEL_PDPTE =                                              0x00000002 //col:1918
EPT_LEVEL_PDE =                                                0x00000001 //col:1919
EPT_LEVEL_PTE =                                                0x00000000 //col:1920
EPML4_ENTRY_COUNT =                                            0x00000200 //col:1921
EPDPTE_ENTRY_COUNT =                                           0x00000200 //col:1922
EPDE_ENTRY_COUNT =                                             0x00000200 //col:1923
EPTE_ENTRY_COUNT =                                             0x00000200 //col:1924
INVEPT_SINGLE_CONTEXT =                                        0x00000001 //col:1925
INVEPT_ALL_CONTEXT =                                           0x00000002 //col:1926
INVVPID_INDIVIDUAL_ADDRESS =                                   0x00000000 //col:1927
INVVPID_SINGLE_CONTEXT =                                       0x00000001 //col:1928
INVVPID_ALL_CONTEXT =                                          0x00000002 //col:1929
INVVPID_SINGLE_CONTEXT_RETAINING_GLOBALS =                     0x00000003 //col:1930
HLATP_PAGE_LEVEL_WRITE_THROUGH =                               0x08 //col:1931
HLATP_PAGE_LEVEL_CACHE_DISABLE =                               0x10 //col:1932
HLATP_PAGE_FRAME_NUMBER =                                      0xFFFFFFFFF000 //col:1933
VMCS_COMPONENT_ENCODING_ACCESS_TYPE =                          0x01 //col:1934
VMCS_COMPONENT_ENCODING_INDEX =                                0x3FE //col:1935
VMCS_COMPONENT_ENCODING_TYPE =                                 0xC00 //col:1936
VMCS_COMPONENT_ENCODING_MUST_BE_ZERO =                         0x1000 //col:1937
VMCS_COMPONENT_ENCODING_WIDTH =                                0x6000 //col:1938
VMCS_CTRL_VPID =                                               0x00000000 //col:1939
VMCS_CTRL_POSTED_INTR_NOTIFY_VECTOR =                          0x00000002 //col:1940
VMCS_CTRL_EPTP_INDEX =                                         0x00000004 //col:1941
VMCS_CTRL_HLAT_PREFIX_SIZE =                                   0x00000006 //col:1942
VMCS_GUEST_ES_SEL =                                            0x00000800 //col:1943
VMCS_GUEST_CS_SEL =                                            0x00000802 //col:1944
VMCS_GUEST_SS_SEL =                                            0x00000804 //col:1945
VMCS_GUEST_DS_SEL =                                            0x00000806 //col:1946
VMCS_GUEST_FS_SEL =                                            0x00000808 //col:1947
VMCS_GUEST_GS_SEL =                                            0x0000080A //col:1948
VMCS_GUEST_LDTR_SEL =                                          0x0000080C //col:1949
VMCS_GUEST_TR_SEL =                                            0x0000080E //col:1950
VMCS_GUEST_INTR_STATUS =                                       0x00000810 //col:1951
VMCS_GUEST_PML_INDEX =                                         0x00000812 //col:1952
VMCS_HOST_ES_SEL =                                             0x00000C00 //col:1953
VMCS_HOST_CS_SEL =                                             0x00000C02 //col:1954
VMCS_HOST_SS_SEL =                                             0x00000C04 //col:1955
VMCS_HOST_DS_SEL =                                             0x00000C06 //col:1956
VMCS_HOST_FS_SEL =                                             0x00000C08 //col:1957
VMCS_HOST_GS_SEL =                                             0x00000C0A //col:1958
VMCS_HOST_TR_SEL =                                             0x00000C0C //col:1959
VMCS_CTRL_IO_BITMAP_A =                                        0x00002000 //col:1960
VMCS_CTRL_IO_BITMAP_B =                                        0x00002002 //col:1961
VMCS_CTRL_MSR_BITMAP =                                         0x00002004 //col:1962
VMCS_CTRL_VMEXIT_MSR_STORE =                                   0x00002006 //col:1963
VMCS_CTRL_VMEXIT_MSR_LOAD =                                    0x00002008 //col:1964
VMCS_CTRL_VMENTRY_MSR_LOAD =                                   0x0000200A //col:1965
VMCS_CTRL_EXEC_VMCS_PTR =                                      0x0000200C //col:1966
VMCS_CTRL_PML_ADDR =                                           0x0000200E //col:1967
VMCS_CTRL_TSC_OFFSET =                                         0x00002010 //col:1968
VMCS_CTRL_VAPIC_PAGEADDR =                                     0x00002012 //col:1969
VMCS_CTRL_APIC_ACCESSADDR =                                    0x00002014 //col:1970
VMCS_CTRL_POSTED_INTR_DESC =                                   0x00002016 //col:1971
VMCS_CTRL_VMFUNC_CTRLS =                                       0x00002018 //col:1972
VMCS_CTRL_EPTP =                                               0x0000201A //col:1973
VMCS_CTRL_EOI_BITMAP_0 =                                       0x0000201C //col:1974
VMCS_CTRL_EOI_BITMAP_1 =                                       0x0000201E //col:1975
VMCS_CTRL_EOI_BITMAP_2 =                                       0x00002020 //col:1976
VMCS_CTRL_EOI_BITMAP_3 =                                       0x00002022 //col:1977
VMCS_CTRL_EPTP_LIST =                                          0x00002024 //col:1978
VMCS_CTRL_VMREAD_BITMAP =                                      0x00002026 //col:1979
VMCS_CTRL_VMWRITE_BITMAP =                                     0x00002028 //col:1980
VMCS_CTRL_VIRTXCPT_INFO_ADDR =                                 0x0000202A //col:1981
VMCS_CTRL_XSS_EXITING_BITMAP =                                 0x0000202C //col:1982
VMCS_CTRL_ENCLS_EXITING_BITMAP =                               0x0000202E //col:1983
VMCS_CTRL_SPP_TABLE_POINTER =                                  0x00002030 //col:1984
VMCS_CTRL_TSC_MULTIPLIER =                                     0x00002032 //col:1985
VMCS_CTRL_PROC_EXEC3 =                                         0x00002034 //col:1986
VMCS_CTRL_ENCLV_EXITING_BITMAP =                               0x00002036 //col:1987
VMCS_CTRL_HLATP =                                              0x00002040 //col:1988
VMCS_CTRL_SECONDARY_EXIT =                                     0x00002044 //col:1989
VMCS_GUEST_PHYS_ADDR =                                         0x00002400 //col:1990
VMCS_GUEST_VMCS_LINK_PTR =                                     0x00002800 //col:1991
VMCS_GUEST_DEBUGCTL =                                          0x00002802 //col:1992
VMCS_GUEST_PAT =                                               0x00002804 //col:1993
VMCS_GUEST_EFER =                                              0x00002806 //col:1994
VMCS_GUEST_PERF_GLOBAL_CTRL =                                  0x00002808 //col:1995
VMCS_GUEST_PDPTE0 =                                            0x0000280A //col:1996
VMCS_GUEST_PDPTE1 =                                            0x0000280C //col:1997
VMCS_GUEST_PDPTE2 =                                            0x0000280E //col:1998
VMCS_GUEST_PDPTE3 =                                            0x00002810 //col:1999
VMCS_GUEST_BNDCFGS =                                           0x00002812 //col:2000
VMCS_GUEST_RTIT_CTL =                                          0x00002814 //col:2001
VMCS_GUEST_LBR_CTL =                                           0x00002816 //col:2002
VMCS_GUEST_PKRS =                                              0x00002818 //col:2003
VMCS_HOST_PAT =                                                0x00002C00 //col:2004
VMCS_HOST_EFER =                                               0x00002C02 //col:2005
VMCS_HOST_PERF_GLOBAL_CTRL =                                   0x00002C04 //col:2006
VMCS_HOST_PKRS =                                               0x00002C06 //col:2007
VMCS_CTRL_PIN_EXEC =                                           0x00004000 //col:2008
VMCS_CTRL_PROC_EXEC =                                          0x00004002 //col:2009
VMCS_CTRL_EXCEPTION_BITMAP =                                   0x00004004 //col:2010
VMCS_CTRL_PAGEFAULT_ERROR_MASK =                               0x00004006 //col:2011
VMCS_CTRL_PAGEFAULT_ERROR_MATCH =                              0x00004008 //col:2012
VMCS_CTRL_CR3_TARGET_COUNT =                                   0x0000400A //col:2013
VMCS_CTRL_PRIMARY_EXIT =                                       0x0000400C //col:2014
VMCS_CTRL_EXIT_MSR_STORE_COUNT =                               0x0000400E //col:2015
VMCS_CTRL_EXIT_MSR_LOAD_COUNT =                                0x00004010 //col:2016
VMCS_CTRL_ENTRY =                                              0x00004012 //col:2017
VMCS_CTRL_ENTRY_MSR_LOAD_COUNT =                               0x00004014 //col:2018
VMCS_CTRL_ENTRY_INTERRUPTION_INFO =                            0x00004016 //col:2019
VMCS_CTRL_ENTRY_EXCEPTION_ERRCODE =                            0x00004018 //col:2020
VMCS_CTRL_ENTRY_INSTR_LENGTH =                                 0x0000401A //col:2021
VMCS_CTRL_TPR_THRESHOLD =                                      0x0000401C //col:2022
VMCS_CTRL_PROC_EXEC2 =                                         0x0000401E //col:2023
VMCS_CTRL_PLE_GAP =                                            0x00004020 //col:2024
VMCS_CTRL_PLE_WINDOW =                                         0x00004022 //col:2025
VMCS_VM_INSTR_ERROR =                                          0x00004400 //col:2026
VMCS_EXIT_REASON =                                             0x00004402 //col:2027
VMCS_EXIT_INTERRUPTION_INFO =                                  0x00004404 //col:2028
VMCS_EXIT_INTERRUPTION_ERROR_CODE =                            0x00004406 //col:2029
VMCS_IDT_VECTORING_INFO =                                      0x00004408 //col:2030
VMCS_IDT_VECTORING_ERROR_CODE =                                0x0000440A //col:2031
VMCS_EXIT_INSTR_LENGTH =                                       0x0000440C //col:2032
VMCS_EXIT_INSTR_INFO =                                         0x0000440E //col:2033
VMCS_GUEST_ES_LIMIT =                                          0x00004800 //col:2034
VMCS_GUEST_CS_LIMIT =                                          0x00004802 //col:2035
VMCS_GUEST_SS_LIMIT =                                          0x00004804 //col:2036
VMCS_GUEST_DS_LIMIT =                                          0x00004806 //col:2037
VMCS_GUEST_FS_LIMIT =                                          0x00004808 //col:2038
VMCS_GUEST_GS_LIMIT =                                          0x0000480A //col:2039
VMCS_GUEST_LDTR_LIMIT =                                        0x0000480C //col:2040
VMCS_GUEST_TR_LIMIT =                                          0x0000480E //col:2041
VMCS_GUEST_GDTR_LIMIT =                                        0x00004810 //col:2042
VMCS_GUEST_IDTR_LIMIT =                                        0x00004812 //col:2043
VMCS_GUEST_ES_ACCESS_RIGHTS =                                  0x00004814 //col:2044
VMCS_GUEST_CS_ACCESS_RIGHTS =                                  0x00004816 //col:2045
VMCS_GUEST_SS_ACCESS_RIGHTS =                                  0x00004818 //col:2046
VMCS_GUEST_DS_ACCESS_RIGHTS =                                  0x0000481A //col:2047
VMCS_GUEST_FS_ACCESS_RIGHTS =                                  0x0000481C //col:2048
VMCS_GUEST_GS_ACCESS_RIGHTS =                                  0x0000481E //col:2049
VMCS_GUEST_LDTR_ACCESS_RIGHTS =                                0x00004820 //col:2050
VMCS_GUEST_TR_ACCESS_RIGHTS =                                  0x00004822 //col:2051
VMCS_GUEST_INTERRUPTIBILITY_STATE =                            0x00004824 //col:2052
VMCS_GUEST_ACTIVITY_STATE =                                    0x00004826 //col:2053
VMCS_GUEST_SMBASE =                                            0x00004828 //col:2054
VMCS_GUEST_SYSENTER_CS =                                       0x0000482A //col:2055
VMCS_GUEST_PREEMPT_TIMER_VALUE =                               0x0000482E //col:2056
VMCS_HOST_SYSENTER_CS =                                        0x00004C00 //col:2057
VMCS_CTRL_CR0_MASK =                                           0x00006000 //col:2058
VMCS_CTRL_CR4_MASK =                                           0x00006002 //col:2059
VMCS_CTRL_CR0_READ_SHADOW =                                    0x00006004 //col:2060
VMCS_CTRL_CR4_READ_SHADOW =                                    0x00006006 //col:2061
VMCS_CTRL_CR3_TARGET_VAL0 =                                    0x00006008 //col:2062
VMCS_CTRL_CR3_TARGET_VAL1 =                                    0x0000600A //col:2063
VMCS_CTRL_CR3_TARGET_VAL2 =                                    0x0000600C //col:2064
VMCS_CTRL_CR3_TARGET_VAL3 =                                    0x0000600E //col:2065
VMCS_EXIT_QUALIFICATION =                                      0x00006400 //col:2066
VMCS_IO_RCX =                                                  0x00006402 //col:2067
VMCS_IO_RSI =                                                  0x00006404 //col:2068
VMCS_IO_RDI =                                                  0x00006406 //col:2069
VMCS_IO_RIP =                                                  0x00006408 //col:2070
VMCS_EXIT_GUEST_LINEAR_ADDR =                                  0x0000640A //col:2071
VMCS_GUEST_CR0 =                                               0x00006800 //col:2072
VMCS_GUEST_CR3 =                                               0x00006802 //col:2073
VMCS_GUEST_CR4 =                                               0x00006804 //col:2074
VMCS_GUEST_ES_BASE =                                           0x00006806 //col:2075
VMCS_GUEST_CS_BASE =                                           0x00006808 //col:2076
VMCS_GUEST_SS_BASE =                                           0x0000680A //col:2077
VMCS_GUEST_DS_BASE =                                           0x0000680C //col:2078
VMCS_GUEST_FS_BASE =                                           0x0000680E //col:2079
VMCS_GUEST_GS_BASE =                                           0x00006810 //col:2080
VMCS_GUEST_LDTR_BASE =                                         0x00006812 //col:2081
VMCS_GUEST_TR_BASE =                                           0x00006814 //col:2082
VMCS_GUEST_GDTR_BASE =                                         0x00006816 //col:2083
VMCS_GUEST_IDTR_BASE =                                         0x00006818 //col:2084
VMCS_GUEST_DR7 =                                               0x0000681A //col:2085
VMCS_GUEST_RSP =                                               0x0000681C //col:2086
VMCS_GUEST_RIP =                                               0x0000681E //col:2087
VMCS_GUEST_RFLAGS =                                            0x00006820 //col:2088
VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS =                          0x00006822 //col:2089
VMCS_GUEST_SYSENTER_ESP =                                      0x00006824 //col:2090
VMCS_GUEST_SYSENTER_EIP =                                      0x00006826 //col:2091
VMCS_GUEST_S_CET =                                             0x00006C28 //col:2092
VMCS_GUEST_SSP =                                               0x00006C2A //col:2093
VMCS_GUEST_INTERRUPT_SSP_TABLE_ADDR =                          0x00006C2C //col:2094
VMCS_HOST_CR0 =                                                0x00006C00 //col:2095
VMCS_HOST_CR3 =                                                0x00006C02 //col:2096
VMCS_HOST_CR4 =                                                0x00006C04 //col:2097
VMCS_HOST_FS_BASE =                                            0x00006C06 //col:2098
VMCS_HOST_GS_BASE =                                            0x00006C08 //col:2099
VMCS_HOST_TR_BASE =                                            0x00006C0A //col:2100
VMCS_HOST_GDTR_BASE =                                          0x00006C0C //col:2101
VMCS_HOST_IDTR_BASE =                                          0x00006C0E //col:2102
VMCS_HOST_SYSENTER_ESP =                                       0x00006C10 //col:2103
VMCS_HOST_SYSENTER_EIP =                                       0x00006C12 //col:2104
VMCS_HOST_RSP =                                                0x00006C14 //col:2105
VMCS_HOST_RIP =                                                0x00006C16 //col:2106
VMCS_HOST_S_CET =                                              0x00006C18 //col:2107
VMCS_HOST_SSP =                                                0x00006C1A //col:2108
VMCS_HOST_INTERRUPT_SSP_TABLE_ADDR =                           0x00006C1C //col:2109
EXTERNAL_INTERRUPT =                                           0x00000000 //col:2110
NON_MASKABLE_INTERRUPT =                                       0x00000002 //col:2111
HARDWARE_EXCEPTION =                                           0x00000003 //col:2112
SOFTWARE_INTERRUPT =                                           0x00000004 //col:2113
PRIVILEGED_SOFTWARE_EXCEPTION =                                0x00000005 //col:2114
SOFTWARE_EXCEPTION =                                           0x00000006 //col:2115
OTHER_EVENT =                                                  0x00000007 //col:2116
VMENTRY_INTERRUPT_INFO_VECTOR =                                0xFF //col:2117
VMENTRY_INTERRUPT_INFO_INTERRUPTION_TYPE =                     0x700 //col:2118
VMENTRY_INTERRUPT_INFO_DELIVER_ERROR_CODE =                    0x800 //col:2119
VMENTRY_INTERRUPT_INFO_VALID =                                 0x80000000 //col:2120
VMEXIT_INTERRUPT_INFO_VECTOR =                                 0xFF //col:2121
VMEXIT_INTERRUPT_INFO_INTERRUPTION_TYPE =                      0x700 //col:2122
VMEXIT_INTERRUPT_INFO_ERROR_CODE_VALID =                       0x800 //col:2123
VMEXIT_INTERRUPT_INFO_NMI_UNBLOCKING =                         0x1000 //col:2124
VMEXIT_INTERRUPT_INFO_VALID =                                  0x80000000 //col:2125
APIC_BASE =                                                    0xFEE00000 //col:2126
APIC_ID =                                                      0x00000020 //col:2127
APIC_VERSION =                                                 0x00000030 //col:2128
APIC_TPR =                                                     0x00000080 //col:2129
APIC_APR =                                                     0x00000090 //col:2130
APIC_PPR =                                                     0x000000A0 //col:2131
APIC_EOI =                                                     0x000000B0 //col:2132
APIC_REMOTE_READ =                                             0x000000C0 //col:2133
APIC_LOGICAL_DESTINATION =                                     0x000000D0 //col:2134
APIC_DESTINATION_FORMAT =                                      0x000000E0 //col:2135
APIC_SIV =                                                     0x000000F0 //col:2136
APIC_ISR_31_0 =                                                0x00000100 //col:2137
APIC_ISR_63_32 =                                               0x00000110 //col:2138
APIC_ISR_95_64 =                                               0x00000120 //col:2139
APIC_ISR_127_96 =                                              0x00000130 //col:2140
APIC_ISR_159_128 =                                             0x00000140 //col:2141
APIC_ISR_191_160 =                                             0x00000150 //col:2142
APIC_ISR_223_192 =                                             0x00000160 //col:2143
APIC_ISR_255_224 =                                             0x00000170 //col:2144
APIC_TMR_31_0 =                                                0x00000180 //col:2145
APIC_TMR_63_32 =                                               0x00000190 //col:2146
APIC_TMR_95_64 =                                               0x000001A0 //col:2147
APIC_TMR_127_96 =                                              0x000001B0 //col:2148
APIC_TMR_159_128 =                                             0x000001C0 //col:2149
APIC_TMR_191_160 =                                             0x000001D0 //col:2150
APIC_TMR_223_192 =                                             0x000001E0 //col:2151
APIC_TMR_255_224 =                                             0x000001F0 //col:2152
APIC_IRR_31_0 =                                                0x00000200 //col:2153
APIC_IRR_63_32 =                                               0x00000210 //col:2154
APIC_IRR_95_64 =                                               0x00000220 //col:2155
APIC_IRR_127_96 =                                              0x00000230 //col:2156
APIC_IRR_159_128 =                                             0x00000240 //col:2157
APIC_IRR_191_160 =                                             0x00000250 //col:2158
APIC_IRR_223_192 =                                             0x00000260 //col:2159
APIC_IRR_255_224 =                                             0x00000270 //col:2160
APIC_ERROR_STATUS =                                            0x00000280 //col:2161
APIC_CMCI =                                                    0x000002F0 //col:2162
APIC_ICR_0_31 =                                                0x00000300 //col:2163
APIC_ICR_32_63 =                                               0x00000310 //col:2164
APIC_LVT_TIMER =                                               0x00000320 //col:2165
APIC_LVT_THERMAL_SENSOR =                                      0x00000330 //col:2166
APIC_LVT_PERFORMANCE_MONITORING_COUNTERS =                     0x00000340 //col:2167
APIC_LVT_LINT0 =                                               0x00000350 //col:2168
APIC_LVT_LINT1 =                                               0x00000360 //col:2169
APIC_LVT_ERROR =                                               0x00000370 //col:2170
APIC_INITIAL_COUNT =                                           0x00000380 //col:2171
APIC_CURRENT_COUNT =                                           0x00000390 //col:2172
APIC_DIVIDE_CONFIGURATION =                                    0x000003E0 //col:2173
EFL_CARRY_FLAG =                                               0x01 //col:2174
EFL_READ_AS_1 =                                                0x02 //col:2175
EFL_PARITY_FLAG =                                              0x04 //col:2176
EFL_AUXILIARY_CARRY_FLAG =                                     0x10 //col:2177
EFL_ZERO_FLAG =                                                0x40 //col:2178
EFL_SIGN_FLAG =                                                0x80 //col:2179
EFL_TRAP_FLAG =                                                0x100 //col:2180
EFL_INTERRUPT_ENABLE_FLAG =                                    0x200 //col:2181
EFL_DIRECTION_FLAG =                                           0x400 //col:2182
EFL_OVERFLOW_FLAG =                                            0x800 //col:2183
EFL_IO_PRIVILEGE_LEVEL =                                       0x3000 //col:2184
EFL_NESTED_TASK_FLAG =                                         0x4000 //col:2185
EFL_RESUME_FLAG =                                              0x10000 //col:2186
EFL_VIRTUAL_8086_MODE_FLAG =                                   0x20000 //col:2187
EFL_ALIGNMENT_CHECK_FLAG =                                     0x40000 //col:2188
EFL_VIRTUAL_INTERRUPT_FLAG =                                   0x80000 //col:2189
EFL_VIRTUAL_INTERRUPT_PENDING_FLAG =                           0x100000 //col:2190
EFL_IDENTIFICATION_FLAG =                                      0x200000 //col:2191
RFL_CARRY_FLAG =                                               0x01 //col:2192
RFL_READ_AS_1 =                                                0x02 //col:2193
RFL_PARITY_FLAG =                                              0x04 //col:2194
RFL_AUXILIARY_CARRY_FLAG =                                     0x10 //col:2195
RFL_ZERO_FLAG =                                                0x40 //col:2196
RFL_SIGN_FLAG =                                                0x80 //col:2197
RFL_TRAP_FLAG =                                                0x100 //col:2198
RFL_INTERRUPT_ENABLE_FLAG =                                    0x200 //col:2199
RFL_DIRECTION_FLAG =                                           0x400 //col:2200
RFL_OVERFLOW_FLAG =                                            0x800 //col:2201
RFL_IO_PRIVILEGE_LEVEL =                                       0x3000 //col:2202
RFL_NESTED_TASK_FLAG =                                         0x4000 //col:2203
RFL_RESUME_FLAG =                                              0x10000 //col:2204
RFL_VIRTUAL_8086_MODE_FLAG =                                   0x20000 //col:2205
RFL_ALIGNMENT_CHECK_FLAG =                                     0x40000 //col:2206
RFL_VIRTUAL_INTERRUPT_FLAG =                                   0x80000 //col:2207
RFL_VIRTUAL_INTERRUPT_PENDING_FLAG =                           0x100000 //col:2208
RFL_IDENTIFICATION_FLAG =                                      0x200000 //col:2209
CONTROL_PROTECTION_EXCEPTION_CPEC =                            0x7FFF //col:2210
CONTROL_PROTECTION_EXCEPTION_ENCL =                            0x8000 //col:2211
DIVIDE_ERROR =                                                 0x00000000 //col:2212
DEBUG =                                                        0x00000001 //col:2213
NMI =                                                          0x00000002 //col:2214
BREAKPOINT =                                                   0x00000003 //col:2215
OVERFLOW =                                                     0x00000004 //col:2216
BOUND_RANGE_EXCEEDED =                                         0x00000005 //col:2217
INVALID_OPCODE =                                               0x00000006 //col:2218
DEVICE_NOT_AVAILABLE =                                         0x00000007 //col:2219
DOUBLE_FAULT =                                                 0x00000008 //col:2220
COPROCESSOR_SEGMENT_OVERRUN =                                  0x00000009 //col:2221
INVALID_TSS =                                                  0x0000000A //col:2222
SEGMENT_NOT_PRESENT =                                          0x0000000B //col:2223
STACK_SEGMENT_FAULT =                                          0x0000000C //col:2224
GENERAL_PROTECTION =                                           0x0000000D //col:2225
PAGE_FAULT =                                                   0x0000000E //col:2226
X87_FLOATING_POINT_ERROR =                                     0x00000010 //col:2227
ALIGNMENT_CHECK =                                              0x00000011 //col:2228
MACHINE_CHECK =                                                0x00000012 //col:2229
SIMD_FLOATING_POINT_ERROR =                                    0x00000013 //col:2230
VIRTUALIZATION_EXCEPTION =                                     0x00000014 //col:2231
CONTROL_PROTECTION =                                           0x00000015 //col:2232
EXCEPTION_ERROR_CODE_EXTERNAL_EVENT =                          0x01 //col:2233
EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION =                     0x02 //col:2234
EXCEPTION_ERROR_CODE_GDT_LDT =                                 0x04 //col:2235
EXCEPTION_ERROR_CODE_INDEX =                                   0xFFF8 //col:2236
PAGE_FAULT_EXCEPTION_PRESENT =                                 0x01 //col:2237
PAGE_FAULT_EXCEPTION_WRITE =                                   0x02 //col:2238
PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS =                        0x04 //col:2239
PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION =                  0x08 //col:2240
PAGE_FAULT_EXCEPTION_EXECUTE =                                 0x10 //col:2241
PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION =                0x20 //col:2242
PAGE_FAULT_EXCEPTION_SHADOW_STACK =                            0x40 //col:2243
PAGE_FAULT_EXCEPTION_HLAT =                                    0x80 //col:2244
PAGE_FAULT_EXCEPTION_SGX =                                     0x8000 //col:2245
MEMORY_TYPE_UC =                                               0x00000000 //col:2246
MEMORY_TYPE_WC =                                               0x00000001 //col:2247
MEMORY_TYPE_WT =                                               0x00000004 //col:2248
MEMORY_TYPE_WP =                                               0x00000005 //col:2249
MEMORY_TYPE_WB =                                               0x00000006 //col:2250
MEMORY_TYPE_UC_MINUS =                                         0x00000007 //col:2251
MEMORY_TYPE_INVALID =                                          0x000000FF //col:2252
VTD_Lower64_PRESENT =                                          0x01 //col:2253
VTD_Lower64_CONTEXT_TABLE_POINTER =                            0xFFFFFFFFFFFFF000 //col:2254
VTD_Upper64_RESERVED =                                         0xFFFFFFFFFFFFFFFF //col:2255
VTD_Lower64_PRESENT =                                          0x01 //col:2256
VTD_Lower64_FAULT_PROCESSING_DISABLE =                         0x02 //col:2257
VTD_Lower64_TRANSLATION_TYPE =                                 0x0C //col:2258
VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER =            0xFFFFFFFFFFFFF000 //col:2259
VTD_Upper64_ADDRESS_WIDTH =                                    0x07 //col:2260
VTD_Upper64_IGNORED =                                          0x78 //col:2261
VTD_Upper64_DOMAIN_IDENTIFIER =                                0x3FF00 //col:2262
VTD_ROOT_ENTRY_COUNT =                                         0x00000100 //col:2263
VTD_CONTEXT_ENTRY_COUNT =                                      0x00000100 //col:2264
VTD_VERSION =                                                  0x00000000 //col:2265
VTD_VERSION_MINOR =                                            0x0F //col:2266
VTD_VERSION_MAJOR =                                            0xF0 //col:2267
VTD_CAPABILITY =                                               0x00000008 //col:2268
VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED =                   0x07 //col:2269
VTD_CAPABILITY_ADVANCED_FAULT_LOGGING =                        0x08 //col:2270
VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING =                0x10 //col:2271
VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION =                   0x20 //col:2272
VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION =                  0x40 //col:2273
VTD_CAPABILITY_CACHING_MODE =                                  0x80 //col:2274
VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS =       0x1F00 //col:2275
VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH =                   0x3F0000 //col:2276
VTD_CAPABILITY_ZERO_LENGTH_READ =                              0x400000 //col:2277
VTD_CAPABILITY_DEPRECATED =                                    0x800000 //col:2278
VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET =               0x3FF000000 //col:2279
VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT =               0x3C00000000 //col:2280
VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION =                   0x8000000000 //col:2281
VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS =           0xFF0000000000 //col:2282
VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE =                    0x3F000000000000 //col:2283
VTD_CAPABILITY_WRITE_DRAINING =                                0x40000000000000 //col:2284
VTD_CAPABILITY_READ_DRAINING =                                 0x80000000000000 //col:2285
VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT =               0x100000000000000 //col:2286
VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT =                     0x800000000000000 //col:2287
VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT =             0x1000000000000000 //col:2288
VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT = 0x4000000000000000 //col:2289
VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT =       0x8000000000000000 //col:2290
VTD_EXTENDED_CAPABILITY =                                      0x00000010 //col:2291
VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY =                  0x01 //col:2292
VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT =          0x02 //col:2293
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT =                   0x04 //col:2294
VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT =          0x08 //col:2295
VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE =              0x10 //col:2296
VTD_EXTENDED_CAPABILITY_DEPRECATED1 =                          0x20 //col:2297
VTD_EXTENDED_CAPABILITY_PASS_THROUGH =                         0x40 //col:2298
VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL =                        0x80 //col:2299
VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET =                0x3FF00 //col:2300
VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE =            0xF00000 //col:2301
VTD_EXTENDED_CAPABILITY_DEPRECATED2 =                          0x1000000 //col:2302
VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT =                  0x2000000 //col:2303
VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT =           0x4000000 //col:2304
VTD_EXTENDED_CAPABILITY_DEPRECATED3 =                          0x10000000 //col:2305
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT =                 0x20000000 //col:2306
VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT =              0x40000000 //col:2307
VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT =                0x200000000 //col:2308
VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT =       0x400000000 //col:2309
VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED =                 0xF800000000 //col:2310
VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT =     0x10000000000 //col:2311
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE =     0x20000000000 //col:2312
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT =           0x40000000000 //col:2313
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT =    0x80000000000 //col:2314
VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT =              0x100000000000 //col:2315
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT =  0x200000000000 //col:2316
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT =     0x400000000000 //col:2317
VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT =      0x800000000000 //col:2318
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY =    0x1000000000000 //col:2319
VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT =                    0x2000000000000 //col:2320
VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT =               0x10000000000000 //col:2321
VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT =                     0x20000000000000 //col:2322
VTD_GLOBAL_COMMAND =                                           0x00000018 //col:2323
VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT =            0x800000 //col:2324
VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER =         0x1000000 //col:2325
VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE =                0x2000000 //col:2326
VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE =                0x4000000 //col:2327
VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH =                        0x8000000 //col:2328
VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING =             0x10000000 //col:2329
VTD_GLOBAL_COMMAND_SET_FAULT_LOG =                             0x20000000 //col:2330
VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER =                    0x40000000 //col:2331
VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE =                        0x80000000 //col:2332
VTD_GLOBAL_STATUS =                                            0x0000001C //col:2333
VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS =      0x800000 //col:2334
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS =   0x1000000 //col:2335
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS =          0x2000000 //col:2336
VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS =          0x4000000 //col:2337
VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS =                  0x8000000 //col:2338
VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS =              0x10000000 //col:2339
VTD_GLOBAL_STATUS_FAULT_LOG_STATUS =                           0x20000000 //col:2340
VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS =                  0x40000000 //col:2341
VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS =                  0x80000000 //col:2342
VTD_ROOT_TABLE_ADDRESS =                                       0x00000020 //col:2343
VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE =                0xC00 //col:2344
VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS =                    0xFFFFFFFFFFFFF000 //col:2345
VTD_CONTEXT_COMMAND =                                          0x00000028 //col:2346
VTD_CONTEXT_COMMAND_DOMAIN_ID =                                0xFFFF //col:2347
VTD_CONTEXT_COMMAND_SOURCE_ID =                                0xFFFF0000 //col:2348
VTD_CONTEXT_COMMAND_FUNCTION_MASK =                            0x300000000 //col:2349
VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY =  0x1800000000000000 //col:2350
VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY = 0x6000000000000000 //col:2351
VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE =                 0x8000000000000000 //col:2352
VTD_INVALIDATE_ADDRESS =                                       0x00000000 //col:2353
VTD_INVALIDATE_ADDRESS_ADDRESS_MASK =                          0x3F //col:2354
VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT =                     0x40 //col:2355
VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS =                          0xFFFFFFFFFFFFF000 //col:2356
VTD_IOTLB_INVALIDATE =                                         0x00000008 //col:2357
VTD_IOTLB_INVALIDATE_DOMAIN_ID =                               0xFFFF00000000 //col:2358
VTD_IOTLB_INVALIDATE_DRAIN_WRITES =                            0x1000000000000 //col:2359
VTD_IOTLB_INVALIDATE_DRAIN_READS =                             0x2000000000000 //col:2360
VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY =   0x600000000000000 //col:2361
VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY =  0x3000000000000000 //col:2362
VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB =                        0x8000000000000000 //col:2363
XCR0_X87 =                                                     0x01 //col:2364
XCR0_SSE =                                                     0x02 //col:2365
XCR0_AVX =                                                     0x04 //col:2366
XCR0_BNDREG =                                                  0x08 //col:2367
XCR0_BNDCSR =                                                  0x10 //col:2368
XCR0_OPMASK =                                                  0x20 //col:2369
XCR0_ZMM_HI256 =                                               0x40 //col:2370
XCR0_ZMM_HI16 =                                                0x80 //col:2371
XCR0_PKRU =                                                    0x200 //col:2372
)

type typedef struct { struct{
max_cpuid_input_value uint32_t //col:2
ebx_value_genu uint32_t //col:3
ecx_value_ntel uint32_t //col:4
edx_value_inei uint32_t //col:5
}


type typedef struct { struct{
stepping_id uint32_t //col:10
model uint32_t //col:11
family_id uint32_t //col:12
processor_type uint32_t //col:13
reserved_1 uint32_t //col:14
extended_model_id uint32_t //col:15
extended_family_id uint32_t //col:16
reserved_2 uint32_t //col:17
}


type typedef struct { struct{
cache_type_field uint32_t //col:108
cache_level uint32_t //col:109
self_initializing_cache_level uint32_t //col:110
fully_associative_cache uint32_t //col:111
reserved_1 uint32_t //col:112
max_addressable_ids_for_logical_processors_sharing_this_cache: uint32_t //col:113
max_addressable_ids_for_processor_cores_in_physical_package: uint32_t //col:114
}


type typedef struct { struct{
smallest_monitor_line_size uint32_t //col:145
reserved_1 uint32_t //col:146
}


type typedef struct { struct{
temperature_sensor_supported uint32_t //col:182
intel_turbo_boost_technology_available uint32_t //col:183
apic_timer_always_running uint32_t //col:184
reserved_1 uint32_t //col:185
power_limit_notification uint32_t //col:186
clock_modulation_duty uint32_t //col:187
package_thermal_management uint32_t //col:188
hwp_base_registers uint32_t //col:189
hwp_notification uint32_t //col:190
hwp_activity_window uint32_t //col:191
hwp_energy_performance_preference uint32_t //col:192
hwp_package_level_request uint32_t //col:193
reserved_2 uint32_t //col:194
hdc uint32_t //col:195
intel_turbo_boost_max_technology_3_available uint32_t //col:196
hwp_capabilities uint32_t //col:197
hwp_peci_override uint32_t //col:198
flexible_hwp uint32_t //col:199
fast_access_mode_for_hwp_request_msr uint32_t //col:200
reserved_3 uint32_t //col:201
ignoring_idle_logical_processor_hwp_request uint32_t //col:202
reserved_4 uint32_t //col:203
intel_thread_director uint32_t //col:204
reserved_5 uint32_t //col:205
}


type typedef struct { struct{
number_of_sub_leaves uint32_t //col:237
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t //col:342
}


type typedef struct { struct{
version_id_of_architectural_performance_monitoring uint32_t //col:368
number_of_performance_monitoring_counter_per_logical_processor: uint32_t //col:369
bit_width_of_performance_monitoring_counter uint32_t //col:370
ebx_bit_vector_length uint32_t //col:371
}


type typedef struct { struct{
x2apic_id_to_unique_topology_id_shift uint32_t //col:408
reserved_1 uint32_t //col:409
}


type typedef struct { struct{
x87_state uint32_t //col:438
sse_state uint32_t //col:439
avx_state uint32_t //col:440
mpx_state uint32_t //col:441
avx_512_state uint32_t //col:442
used_for_ia32_xss_1 uint32_t //col:443
pkru_state uint32_t //col:444
reserved_1 uint32_t //col:445
used_for_ia32_xss_2 uint32_t //col:446
reserved_2 uint32_t //col:447
}


type typedef struct { struct{
reserved_1 uint32_t //col:473
supports_xsavec_and_compacted_xrstor uint32_t //col:474
supports_xgetbv_with_ecx_1 uint32_t //col:475
supports_xsave_xrstor_and_ia32_xss uint32_t //col:476
reserved_2 uint32_t //col:477
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t //col:513
}


type typedef struct { struct{
reserved uint32_t //col:541
}


type typedef struct { struct{
reserved uint32_t //col:569
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t //col:598
}


type typedef struct { struct{
length_of_capacity_bit_mask uint32_t //col:628
reserved_1 uint32_t //col:629
}


type typedef struct { struct{
length_of_capacity_bit_mask uint32_t //col:658
reserved_1 uint32_t //col:659
}


type typedef struct { struct{
max_mba_throttling_value uint32_t //col:686
reserved_1 uint32_t //col:687
}


type typedef struct { struct{
sgx1 uint32_t //col:716
sgx2 uint32_t //col:717
reserved_1 uint32_t //col:718
sgx_enclv_advanced uint32_t //col:719
sgx_encls_advanced uint32_t //col:720
reserved_2 uint32_t //col:721
}


type typedef struct { struct{
valid_secs_attributes_0 uint32_t //col:749
}


type typedef struct { struct{
sub_leaf_type uint32_t //col:775
reserved_1 uint32_t //col:776
}


type typedef struct { struct{
sub_leaf_type uint32_t //col:802
reserved_1 uint32_t //col:803
epc_base_physical_address_1 uint32_t //col:804
}


type typedef struct { struct{
max_sub_leaf uint32_t //col:834
}


type typedef struct { struct{
number_of_configurable_address_ranges_for_filtering uint32_t //col:874
reserved_1 uint32_t //col:875
bitmap_of_supported_mtc_period_encodings uint32_t //col:876
}


type typedef struct { struct{
denominator uint32_t //col:903
}


type typedef struct { struct{
procesor_base_frequency_mhz uint32_t //col:929
reserved_1 uint32_t //col:930
}


type typedef struct { struct{
max_soc_id_index uint32_t //col:958
}


type typedef struct { struct{
soc_vendor_brand_string uint32_t //col:986
}


type typedef struct { struct{
reserved uint32_t //col:1012
}


type typedef struct { struct{
max_sub_leaf uint32_t //col:1038
}


type typedef struct { struct{
reserved uint32_t //col:1076
}


type typedef struct { struct{
max_extended_functions uint32_t //col:1114
}


type typedef struct { struct{
reserved uint32_t //col:1140
}


type typedef struct { struct{
processor_brand_string_1 uint32_t //col:1180
}


type typedef struct { struct{
processor_brand_string_5 uint32_t //col:1206
}


type typedef struct { struct{
processor_brand_string_9 uint32_t //col:1232
}


type typedef struct { struct{
reserved uint32_t //col:1258
}


type typedef struct { struct{
reserved uint32_t //col:1284
}


type typedef struct { struct{
reserved uint32_t //col:1313
}


type typedef struct { struct{
number_of_physical_address_bits uint32_t //col:1341
number_of_linear_address_bits uint32_t //col:1342
reserved_1 uint32_t //col:1343
}


type typedef struct { struct{
thread_adjust uint64_t //col:1367
}


type typedef struct { struct{
mseg_header_revision uint32_t //col:1370
monitor_features uint32_t //col:1371
gdtr_limit uint32_t //col:1372
gdtr_base_offset uint32_t //col:1373
cs_selector uint32_t //col:1374
eip_offset uint32_t //col:1375
esp_offset uint32_t //col:1376
cr3_offset uint32_t //col:1377
}


type typedef struct { struct{
c0_mcnt uint64_t //col:1380
}


type typedef struct { struct{
c0_acnt uint64_t //col:1383
}


type typedef struct { struct{
stall_cycle_cnt uint64_t //col:1386
}


type typedef struct { struct{
limit uint16_t //col:1389
base_address uint32_t //col:1390
}


type typedef struct { struct{
limit uint16_t //col:1393
base_address uint64_t //col:1394
}


type typedef struct { struct{
segment_limit_low uint16_t //col:1397
base_address_low uint16_t //col:1398
base_address_middle uint32_t //col:1401
type uint32_t //col:1402
descriptor_type uint32_t //col:1403
descriptor_privilege_level uint32_t //col:1404
present uint32_t //col:1405
segment_limit_high uint32_t //col:1406
available_bit uint32_t //col:1407
long_mode uint32_t //col:1408
default_big uint32_t //col:1409
granularity uint32_t //col:1410
base_address_high uint32_t //col:1411
}


type typedef struct { struct{
segment_limit_low uint16_t //col:1417
base_address_low uint16_t //col:1418
base_address_middle uint32_t //col:1421
type uint32_t //col:1422
descriptor_type uint32_t //col:1423
descriptor_privilege_level uint32_t //col:1424
present uint32_t //col:1425
segment_limit_high uint32_t //col:1426
available_bit uint32_t //col:1427
long_mode uint32_t //col:1428
default_big uint32_t //col:1429
granularity uint32_t //col:1430
base_address_high uint32_t //col:1431
}


type typedef struct { struct{
offset_low uint16_t //col:1439
segment_selector uint16_t //col:1440
interrupt_stack_table uint32_t //col:1443
must_be_zero_0 uint32_t //col:1444
type uint32_t //col:1445
must_be_zero_1 uint32_t //col:1446
descriptor_privilege_level uint32_t //col:1447
present uint32_t //col:1448
offset_middle uint32_t //col:1449
}


type typedef struct { struct{
reserved_0 uint32_t //col:1457
rsp0 uint64_t //col:1458
rsp1 uint64_t //col:1459
rsp2 uint64_t //col:1460
reserved_1 uint64_t //col:1461
ist1 uint64_t //col:1462
ist2 uint64_t //col:1463
ist3 uint64_t //col:1464
ist4 uint64_t //col:1465
ist5 uint64_t //col:1466
ist6 uint64_t //col:1467
ist7 uint64_t //col:1468
reserved_2 uint64_t //col:1469
reserved_3 uint16_t //col:1470
io_map_base uint16_t //col:1471
}


type typedef struct { struct{
reason uint32_t //col:1474
exception_mask uint32_t //col:1475
exit uint64_t //col:1476
guest_linear_address uint64_t //col:1477
guest_physical_address uint64_t //col:1478
current_eptp_index uint16_t //col:1479
}


type typedef struct { struct{
io_a[4096] uint8_t //col:1482
io_b[4096] uint8_t //col:1483
}


type typedef struct { struct{
rdmsr_low[1024] uint8_t //col:1486
rdmsr_high[1024] uint8_t //col:1487
wrmsr_low[1024] uint8_t //col:1488
wrmsr_high[1024] uint8_t //col:1489
}


type typedef struct { struct{
ept_pointer uint64_t //col:1492
reserved uint64_t //col:1493
}


type typedef struct { struct{
vpid uint16_t //col:1496
reserved1 uint16_t //col:1497
reserved2 uint32_t //col:1498
linear_address uint64_t //col:1499
}


type typedef struct { struct{
revision_id uint32_t //col:1503
shadow_vmcs_indicator uint32_t //col:1504
}


type typedef struct { struct{
revision_id uint32_t //col:1511
must_be_zero uint32_t //col:1512
}


type typedef struct { struct{
present uint64_t //col:1519
reserved_1 uint64_t //col:1520
context_table_pointer uint64_t //col:1521
}


type typedef struct { struct{
present uint64_t //col:1535
fault_processing_disable uint64_t //col:1536
translation_type uint64_t //col:1537
reserved_1 uint64_t //col:1538
second_level_page_translation_pointer uint64_t //col:1539
}




