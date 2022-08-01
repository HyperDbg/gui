package out


const(
CR0_PROTECTION_ENABLE =                                        0x01 //col:22
CR0_MONITOR_COPROCESSOR =                                      0x02 //col:24
CR0_EMULATE_FPU =                                              0x04 //col:26
CR0_TASK_SWITCHED =                                            0x08 //col:28
CR0_EXTENSION_TYPE =                                           0x10 //col:30
CR0_NUMERIC_ERROR =                                            0x20 //col:32
CR0_WRITE_PROTECT =                                            0x10000 //col:35
CR0_ALIGNMENT_MASK =                                           0x40000 //col:38
CR0_NOT_WRITE_THROUGH =                                        0x20000000 //col:41
CR0_CACHE_DISABLE =                                            0x40000000 //col:43
CR0_PAGING_ENABLE =                                            0x80000000 //col:45
CR3_PAGE_LEVEL_WRITE_THROUGH =                                 0x08 //col:56
CR3_PAGE_LEVEL_CACHE_DISABLE =                                 0x10 //col:58
CR3_ADDRESS_OF_PAGE_DIRECTORY =                                0xFFFFFFFFF000 //col:61
CR4_VIRTUAL_MODE_EXTENSIONS =                                  0x01 //col:71
CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS =                        0x02 //col:73
CR4_TIMESTAMP_DISABLE =                                        0x04 //col:75
CR4_DEBUGGING_EXTENSIONS =                                     0x08 //col:77
CR4_PAGE_SIZE_EXTENSIONS =                                     0x10 //col:79
CR4_PHYSICAL_ADDRESS_EXTENSION =                               0x20 //col:81
CR4_MACHINE_CHECK_ENABLE =                                     0x40 //col:83
CR4_PAGE_GLOBAL_ENABLE =                                       0x80 //col:85
CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE =                    0x100 //col:87
CR4_OS_FXSAVE_FXRSTOR_SUPPORT =                                0x200 //col:89
CR4_OS_XMM_EXCEPTION_SUPPORT =                                 0x400 //col:91
CR4_USERMODE_INSTRUCTION_PREVENTION =                          0x800 //col:93
CR4_LINEAR_ADDRESSES_57_BIT =                                  0x1000 //col:95
CR4_VMX_ENABLE =                                               0x2000 //col:97
CR4_SMX_ENABLE =                                               0x4000 //col:99
CR4_FSGSBASE_ENABLE =                                          0x10000 //col:102
CR4_PCID_ENABLE =                                              0x20000 //col:104
CR4_OS_XSAVE =                                                 0x40000 //col:106
CR4_KEY_LOCKER_ENABLE =                                        0x80000 //col:108
CR4_SMEP_ENABLE =                                              0x100000 //col:110
CR4_SMAP_ENABLE =                                              0x200000 //col:112
CR4_PROTECTION_KEY_ENABLE =                                    0x400000 //col:114
CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE =                          0x800000 //col:116
CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE =                0x1000000 //col:118
CR8_TASK_PRIORITY_LEVEL =                                      0x0F //col:128
CR8_RESERVED =                                                 0xFFFFFFFFFFFFFFF0 //col:130
DR6_BREAKPOINT_CONDITION =                                     0x0F //col:148
DR6_DEBUG_REGISTER_ACCESS_DETECTED =                           0x2000 //col:151
DR6_SINGLE_INSTRUCTION =                                       0x4000 //col:153
DR6_TASK_SWITCH =                                              0x8000 //col:155
DR6_RESTRICTED_TRANSACTIONAL_MEMORY =                          0x10000 //col:157
DR7_LOCAL_BREAKPOINT_0 =                                       0x01 //col:167
DR7_GLOBAL_BREAKPOINT_0 =                                      0x02 //col:169
DR7_LOCAL_BREAKPOINT_1 =                                       0x04 //col:171
DR7_GLOBAL_BREAKPOINT_1 =                                      0x08 //col:173
DR7_LOCAL_BREAKPOINT_2 =                                       0x10 //col:175
DR7_GLOBAL_BREAKPOINT_2 =                                      0x20 //col:177
DR7_LOCAL_BREAKPOINT_3 =                                       0x40 //col:179
DR7_GLOBAL_BREAKPOINT_3 =                                      0x80 //col:181
DR7_LOCAL_EXACT_BREAKPOINT =                                   0x100 //col:183
DR7_GLOBAL_EXACT_BREAKPOINT =                                  0x200 //col:185
DR7_RESTRICTED_TRANSACTIONAL_MEMORY =                          0x800 //col:188
DR7_GENERAL_DETECT =                                           0x2000 //col:191
DR7_READ_WRITE_0 =                                             0x30000 //col:194
DR7_LENGTH_0 =                                                 0xC0000 //col:196
DR7_READ_WRITE_1 =                                             0x300000 //col:198
DR7_LENGTH_1 =                                                 0xC00000 //col:200
DR7_READ_WRITE_2 =                                             0x3000000 //col:202
DR7_LENGTH_2 =                                                 0xC000000 //col:204
DR7_READ_WRITE_3 =                                             0x30000000 //col:206
DR7_LENGTH_3 =                                                 0xC0000000 //col:208
CPUID_SIGNATURE =                                              0x00000000 //col:224
CPUID_VERSION_INFO =                                           0x00000001 //col:232
CPUID_EAX_STEPPING_ID =                                        0x0F //col:237
CPUID_EAX_MODEL =                                              0xF0 //col:239
CPUID_EAX_FAMILY_ID =                                          0xF00 //col:241
CPUID_EAX_PROCESSOR_TYPE =                                     0x3000 //col:243
CPUID_EAX_EXTENDED_MODEL_ID =                                  0xF0000 //col:246
CPUID_EAX_EXTENDED_FAMILY_ID =                                 0xFF00000 //col:248
CPUID_EBX_BRAND_INDEX =                                        0xFF //col:258
CPUID_EBX_CLFLUSH_LINE_SIZE =                                  0xFF00 //col:260
CPUID_EBX_MAX_ADDRESSABLE_IDS =                                0xFF0000 //col:262
CPUID_EBX_INITIAL_APIC_ID =                                    0xFF000000 //col:264
CPUID_ECX_STREAMING_SIMD_EXTENSIONS_3 =                        0x01 //col:273
CPUID_ECX_PCLMULQDQ_INSTRUCTION =                              0x02 //col:275
CPUID_ECX_DS_AREA_64BIT_LAYOUT =                               0x04 //col:277
CPUID_ECX_MONITOR_MWAIT_INSTRUCTION =                          0x08 //col:279
CPUID_ECX_CPL_QUALIFIED_DEBUG_STORE =                          0x10 //col:281
CPUID_ECX_VIRTUAL_MACHINE_EXTENSIONS =                         0x20 //col:283
CPUID_ECX_SAFER_MODE_EXTENSIONS =                              0x40 //col:285
CPUID_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY =                0x80 //col:287
CPUID_ECX_THERMAL_MONITOR_2 =                                  0x100 //col:289
CPUID_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3 =           0x200 //col:291
CPUID_ECX_L1_CONTEXT_ID =                                      0x400 //col:293
CPUID_ECX_SILICON_DEBUG =                                      0x800 //col:295
CPUID_ECX_FMA_EXTENSIONS =                                     0x1000 //col:297
CPUID_ECX_CMPXCHG16B_INSTRUCTION =                             0x2000 //col:299
CPUID_ECX_XTPR_UPDATE_CONTROL =                                0x4000 //col:301
CPUID_ECX_PERFMON_AND_DEBUG_CAPABILITY =                       0x8000 //col:303
CPUID_ECX_PROCESS_CONTEXT_IDENTIFIERS =                        0x20000 //col:306
CPUID_ECX_DIRECT_CACHE_ACCESS =                                0x40000 //col:308
CPUID_ECX_SSE41_SUPPORT =                                      0x80000 //col:310
CPUID_ECX_SSE42_SUPPORT =                                      0x100000 //col:312
CPUID_ECX_X2APIC_SUPPORT =                                     0x200000 //col:314
CPUID_ECX_MOVBE_INSTRUCTION =                                  0x400000 //col:316
CPUID_ECX_POPCNT_INSTRUCTION =                                 0x800000 //col:318
CPUID_ECX_TSC_DEADLINE =                                       0x1000000 //col:320
CPUID_ECX_AESNI_INSTRUCTION_EXTENSIONS =                       0x2000000 //col:322
CPUID_ECX_XSAVE_XRSTOR_INSTRUCTION =                           0x4000000 //col:324
CPUID_ECX_OSX_SAVE =                                           0x8000000 //col:326
CPUID_ECX_AVX_SUPPORT =                                        0x10000000 //col:328
CPUID_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS =             0x20000000 //col:330
CPUID_ECX_RDRAND_INSTRUCTION =                                 0x40000000 //col:332
CPUID_EDX_FLOATING_POINT_UNIT_ON_CHIP =                        0x01 //col:342
CPUID_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS =                     0x02 //col:344
CPUID_EDX_DEBUGGING_EXTENSIONS =                               0x04 //col:346
CPUID_EDX_PAGE_SIZE_EXTENSION =                                0x08 //col:348
CPUID_EDX_TIMESTAMP_COUNTER =                                  0x10 //col:350
CPUID_EDX_RDMSR_WRMSR_INSTRUCTIONS =                           0x20 //col:352
CPUID_EDX_PHYSICAL_ADDRESS_EXTENSION =                         0x40 //col:354
CPUID_EDX_MACHINE_CHECK_EXCEPTION =                            0x80 //col:356
CPUID_EDX_CMPXCHG8B_INSTRUCTION =                              0x100 //col:358
CPUID_EDX_APIC_ON_CHIP =                                       0x200 //col:360
CPUID_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS =                      0x800 //col:363
CPUID_EDX_MEMORY_TYPE_RANGE_REGISTERS =                        0x1000 //col:365
CPUID_EDX_PAGE_GLOBAL_BIT =                                    0x2000 //col:367
CPUID_EDX_MACHINE_CHECK_ARCHITECTURE =                         0x4000 //col:369
CPUID_EDX_CONDITIONAL_MOVE_INSTRUCTIONS =                      0x8000 //col:371
CPUID_EDX_PAGE_ATTRIBUTE_TABLE =                               0x10000 //col:373
CPUID_EDX_PAGE_SIZE_EXTENSION_36BIT =                          0x20000 //col:375
CPUID_EDX_PROCESSOR_SERIAL_NUMBER =                            0x40000 //col:377
CPUID_EDX_CLFLUSH_INSTRUCTION =                                0x80000 //col:379
CPUID_EDX_DEBUG_STORE =                                        0x200000 //col:382
CPUID_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI =                      0x400000 //col:384
CPUID_EDX_MMX_SUPPORT =                                        0x800000 //col:386
CPUID_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS =                        0x1000000 //col:388
CPUID_EDX_SSE_SUPPORT =                                        0x2000000 //col:390
CPUID_EDX_SSE2_SUPPORT =                                       0x4000000 //col:392
CPUID_EDX_SELF_SNOOP =                                         0x8000000 //col:394
CPUID_EDX_HYPER_THREADING_TECHNOLOGY =                         0x10000000 //col:396
CPUID_EDX_THERMAL_MONITOR =                                    0x20000000 //col:398
CPUID_EDX_PENDING_BREAK_ENABLE =                               0x80000000 //col:401
CPUID_CACHE_PARAMS =                                           0x00000004 //col:409
CPUID_EAX_CACHE_TYPE_FIELD =                                   0x1F //col:414
CPUID_EAX_CACHE_LEVEL =                                        0xE0 //col:416
CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL =                      0x100 //col:418
CPUID_EAX_FULLY_ASSOCIATIVE_CACHE =                            0x200 //col:420
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE = 0x3FFC000 //col:423
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE = 0xFC000000 //col:425
CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE =                         0xFFF //col:434
CPUID_EBX_PHYSICAL_LINE_PARTITIONS =                           0x3FF000 //col:436
CPUID_EBX_WAYS_OF_ASSOCIATIVITY =                              0xFFC00000 //col:438
CPUID_ECX_NUMBER_OF_SETS =                                     0xFFFFFFFF //col:447
CPUID_EDX_WRITE_BACK_INVALIDATE =                              0x01 //col:456
CPUID_EDX_CACHE_INCLUSIVENESS =                                0x02 //col:458
CPUID_EDX_COMPLEX_CACHE_INDEXING =                             0x04 //col:460
CPUID_MONITOR_MWAIT =                                          0x00000005 //col:469
CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE =                         0xFFFF //col:474
CPUID_EBX_LARGEST_MONITOR_LINE_SIZE =                          0xFFFF //col:484
CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS =            0x01 //col:494
CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT = 0x02 //col:496
CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES =                          0x0F //col:506
CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES =                          0xF0 //col:508
CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES =                          0xF00 //col:510
CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES =                          0xF000 //col:512
CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES =                          0xF0000 //col:514
CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES =                          0xF00000 //col:516
CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES =                          0xF000000 //col:518
CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES =                          0xF0000000 //col:520
CPUID_THERMAL_POWER_MANAGEMENT =                               0x00000006 //col:528
CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED =                       0x01 //col:533
CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE =             0x02 //col:535
CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING =                          0x04 //col:537
CPUID_EAX_POWER_LIMIT_NOTIFICATION =                           0x10 //col:540
CPUID_EAX_CLOCK_MODULATION_DUTY =                              0x20 //col:542
CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT =                         0x40 //col:544
CPUID_EAX_HWP_BASE_REGISTERS =                                 0x80 //col:546
CPUID_EAX_HWP_NOTIFICATION =                                   0x100 //col:548
CPUID_EAX_HWP_ACTIVITY_WINDOW =                                0x200 //col:550
CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE =                  0x400 //col:552
CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST =                          0x800 //col:554
CPUID_EAX_HDC =                                                0x2000 //col:557
CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE =       0x4000 //col:559
CPUID_EAX_HWP_CAPABILITIES =                                   0x8000 //col:561
CPUID_EAX_HWP_PECI_OVERRIDE =                                  0x10000 //col:563
CPUID_EAX_FLEXIBLE_HWP =                                       0x20000 //col:565
CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR =               0x40000 //col:567
CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST =        0x100000 //col:570
CPUID_EAX_INTEL_THREAD_DIRECTOR =                              0x800000 //col:573
CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR =   0x0F //col:583
CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY =          0x01 //col:593
CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES =            0x08 //col:596
CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE =                 0xFF00 //col:599
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:609
CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS =                      0x00000007 //col:617
CPUID_EAX_NUMBER_OF_SUB_LEAVES =                               0xFFFFFFFF //col:622
CPUID_EBX_FSGSBASE =                                           0x01 //col:631
CPUID_EBX_IA32_TSC_ADJUST_MSR =                                0x02 //col:633
CPUID_EBX_SGX =                                                0x04 //col:635
CPUID_EBX_BMI1 =                                               0x08 //col:637
CPUID_EBX_HLE =                                                0x10 //col:639
CPUID_EBX_AVX2 =                                               0x20 //col:641
CPUID_EBX_FDP_EXCPTN_ONLY =                                    0x40 //col:643
CPUID_EBX_SMEP =                                               0x80 //col:645
CPUID_EBX_BMI2 =                                               0x100 //col:647
CPUID_EBX_ENHANCED_REP_MOVSB_STOSB =                           0x200 //col:649
CPUID_EBX_INVPCID =                                            0x400 //col:651
CPUID_EBX_RTM =                                                0x800 //col:653
CPUID_EBX_RDT_M =                                              0x1000 //col:655
CPUID_EBX_DEPRECATES =                                         0x2000 //col:657
CPUID_EBX_MPX =                                                0x4000 //col:659
CPUID_EBX_RDT =                                                0x8000 //col:661
CPUID_EBX_AVX512F =                                            0x10000 //col:663
CPUID_EBX_AVX512DQ =                                           0x20000 //col:665
CPUID_EBX_RDSEED =                                             0x40000 //col:667
CPUID_EBX_ADX =                                                0x80000 //col:669
CPUID_EBX_SMAP =                                               0x100000 //col:671
CPUID_EBX_AVX512_IFMA =                                        0x200000 //col:673
CPUID_EBX_CLFLUSHOPT =                                         0x800000 //col:676
CPUID_EBX_CLWB =                                               0x1000000 //col:678
CPUID_EBX_INTEL =                                              0x2000000 //col:680
CPUID_EBX_AVX512PF =                                           0x4000000 //col:682
CPUID_EBX_AVX512ER =                                           0x8000000 //col:684
CPUID_EBX_AVX512CD =                                           0x10000000 //col:686
CPUID_EBX_SHA =                                                0x20000000 //col:688
CPUID_EBX_AVX512BW =                                           0x40000000 //col:690
CPUID_EBX_AVX512VL =                                           0x80000000 //col:692
CPUID_ECX_PREFETCHWT1 =                                        0x01 //col:701
CPUID_ECX_AVX512_VBMI =                                        0x02 //col:703
CPUID_ECX_UMIP =                                               0x04 //col:705
CPUID_ECX_PKU =                                                0x08 //col:707
CPUID_ECX_OSPKE =                                              0x10 //col:709
CPUID_ECX_WAITPKG =                                            0x20 //col:711
CPUID_ECX_AVX512_VBMI2 =                                       0x40 //col:713
CPUID_ECX_CET_SS =                                             0x80 //col:715
CPUID_ECX_GFNI =                                               0x100 //col:717
CPUID_ECX_VAES =                                               0x200 //col:719
CPUID_ECX_VPCLMULQDQ =                                         0x400 //col:721
CPUID_ECX_AVX512_VNNI =                                        0x800 //col:723
CPUID_ECX_AVX512_BITALG =                                      0x1000 //col:725
CPUID_ECX_TME_EN =                                             0x2000 //col:727
CPUID_ECX_AVX512_VPOPCNTDQ =                                   0x4000 //col:729
CPUID_ECX_LA57 =                                               0x10000 //col:732
CPUID_ECX_MAWAU =                                              0x3E0000 //col:734
CPUID_ECX_RDPID =                                              0x400000 //col:736
CPUID_ECX_KL =                                                 0x800000 //col:738
CPUID_ECX_CLDEMOTE =                                           0x2000000 //col:741
CPUID_ECX_MOVDIRI =                                            0x8000000 //col:744
CPUID_ECX_MOVDIR64B =                                          0x10000000 //col:746
CPUID_ECX_SGX_LC =                                             0x40000000 //col:749
CPUID_ECX_PKS =                                                0x80000000 //col:751
CPUID_EDX_AVX512_4VNNIW =                                      0x04 //col:761
CPUID_EDX_AVX512_4FMAPS =                                      0x08 //col:763
CPUID_EDX_FAST_SHORT_REP_MOV =                                 0x10 //col:765
CPUID_EDX_AVX512_VP2INTERSECT =                                0x100 //col:768
CPUID_EDX_MD_CLEAR =                                           0x400 //col:771
CPUID_EDX_SERIALIZE =                                          0x4000 //col:774
CPUID_EDX_HYBRID =                                             0x8000 //col:776
CPUID_EDX_PCONFIG =                                            0x40000 //col:779
CPUID_EDX_CET_IBT =                                            0x100000 //col:782
CPUID_EDX_IBRS_IBPB =                                          0x4000000 //col:785
CPUID_EDX_STIBP =                                              0x8000000 //col:787
CPUID_EDX_L1D_FLUSH =                                          0x10000000 //col:789
CPUID_EDX_IA32_ARCH_CAPABILITIES =                             0x20000000 //col:791
CPUID_EDX_IA32_CORE_CAPABILITIES =                             0x40000000 //col:793
CPUID_EDX_SSBD =                                               0x80000000 //col:795
CPUID_DIRECT_CACHE_ACCESS_INFO =                               0x00000009 //col:803
CPUID_EAX_IA32_PLATFORM_DCA_CAP =                              0xFFFFFFFF //col:808
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:817
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:826
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:835
CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING =                   0x0000000A //col:843
CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING = 0xFF //col:848
CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR = 0xFF00 //col:850
CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER =        0xFF0000 //col:852
CPUID_EAX_EBX_BIT_VECTOR_LENGTH =                              0xFF000000 //col:854
CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE =                     0x01 //col:863
CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE =            0x02 //col:865
CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE =               0x04 //col:867
CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE =     0x08 //col:869
CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE =        0x10 //col:871
CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE =     0x20 //col:873
CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE =      0x40 //col:875
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:885
CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS =      0x1F //col:894
CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS =   0x1FE0 //col:896
CPUID_EDX_ANY_THREAD_DEPRECATION =                             0x8000 //col:899
CPUID_EXTENDED_TOPOLOGY =                                      0x0000000B //col:908
CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT =              0x1F //col:913
CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE =    0xFFFF //col:923
CPUID_ECX_LEVEL_NUMBER =                                       0xFF //col:933
CPUID_ECX_LEVEL_TYPE =                                         0xFF00 //col:935
CPUID_EDX_X2APIC_ID =                                          0xFFFFFFFF //col:945
CPUID_EXTENDED_STATE =                                         0x0000000D //col:958
CPUID_EAX_X87_STATE =                                          0x01 //col:963
CPUID_EAX_SSE_STATE =                                          0x02 //col:965
CPUID_EAX_AVX_STATE =                                          0x04 //col:967
CPUID_EAX_MPX_STATE =                                          0x18 //col:969
CPUID_EAX_AVX_512_STATE =                                      0xE0 //col:971
CPUID_EAX_USED_FOR_IA32_XSS_1 =                                0x100 //col:973
CPUID_EAX_PKRU_STATE =                                         0x200 //col:975
CPUID_EAX_USED_FOR_IA32_XSS_2 =                                0x2000 //col:978
CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0 =      0xFFFFFFFF //col:988
CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA =                 0xFFFFFFFF //col:997
CPUID_EDX_XCR0_SUPPORTED_BITS =                                0xFFFFFFFF //col:1006
CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR =               0x02 //col:1019
CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1 =                         0x04 //col:1021
CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS =                 0x08 //col:1023
CPUID_EBX_SIZE_OF_XSAVE_AREAD =                                0xFFFFFFFF //col:1033
CPUID_ECX_USED_FOR_XCR0_1 =                                    0xFF //col:1042
CPUID_ECX_PT_STATE =                                           0x100 //col:1044
CPUID_ECX_USED_FOR_XCR0_2 =                                    0x200 //col:1046
CPUID_ECX_CET_USER_STATE =                                     0x800 //col:1049
CPUID_ECX_CET_SUPERVISOR_STATE =                               0x1000 //col:1051
CPUID_ECX_HDC_STATE =                                          0x2000 //col:1053
CPUID_ECX_LBR_STATE =                                          0x8000 //col:1056
CPUID_ECX_HWP_STATE =                                          0x10000 //col:1058
CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS =                      0xFFFFFFFF //col:1068
CPUID_EAX_IA32_PLATFORM_DCA_CAP =                              0xFFFFFFFF //col:1080
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:1089
CPUID_ECX_ECX_2 =                                              0x01 //col:1098
CPUID_ECX_ECX_1 =                                              0x02 //col:1100
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1110
CPUID_INTEL_RDT_MONITORING =                                   0x0000000F //col:1127
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:1132
CPUID_EBX_RMID_MAX_RANGE =                                     0xFFFFFFFF //col:1141
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:1150
CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING =             0x02 //col:1160
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:1173
CPUID_EBX_CONVERSION_FACTOR =                                  0xFFFFFFFF //col:1182
CPUID_ECX_RMID_MAX_RANGE =                                     0xFFFFFFFF //col:1191
CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING =                   0x01 //col:1200
CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING =             0x02 //col:1202
CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING =             0x04 //col:1204
CPUID_INTEL_RDT_ALLOCATION =                                   0x00000010 //col:1222
CPUID_EAX_IA32_PLATFORM_DCA_CAP =                              0xFFFFFFFF //col:1227
CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY =            0x02 //col:1237
CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY =            0x04 //col:1239
CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION =               0x08 //col:1241
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:1251
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1260
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK =                        0x1F //col:1272
CPUID_EBX_EBX_0 =                                              0xFFFFFFFF //col:1282
CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED =    0x04 //col:1292
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED =                       0xFFFF //col:1302
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK =                        0x1F //col:1315
CPUID_EBX_EBX_0 =                                              0xFFFFFFFF //col:1325
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:1334
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED =                       0xFFFF //col:1343
CPUID_EAX_MAX_MBA_THROTTLING_VALUE =                           0xFFF //col:1356
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:1366
CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR =                        0x04 //col:1376
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED =                       0xFFFF //col:1386
CPUID_INTEL_SGX =                                              0x00000012 //col:1404
CPUID_EAX_SGX1 =                                               0x01 //col:1409
CPUID_EAX_SGX2 =                                               0x02 //col:1411
CPUID_EAX_SGX_ENCLV_ADVANCED =                                 0x20 //col:1414
CPUID_EAX_SGX_ENCLS_ADVANCED =                                 0x40 //col:1416
CPUID_EBX_MISCSELECT =                                         0xFFFFFFFF //col:1426
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:1435
CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64 =                             0xFF //col:1444
CPUID_EDX_MAX_ENCLAVE_SIZE_64 =                                0xFF00 //col:1446
CPUID_EAX_VALID_SECS_ATTRIBUTES_0 =                            0xFFFFFFFF //col:1459
CPUID_EBX_VALID_SECS_ATTRIBUTES_1 =                            0xFFFFFFFF //col:1468
CPUID_ECX_VALID_SECS_ATTRIBUTES_2 =                            0xFFFFFFFF //col:1477
CPUID_EDX_VALID_SECS_ATTRIBUTES_3 =                            0xFFFFFFFF //col:1486
CPUID_EAX_SUB_LEAF_TYPE =                                      0x0F //col:1498
CPUID_EBX_ZERO =                                               0xFFFFFFFF //col:1508
CPUID_ECX_ZERO =                                               0xFFFFFFFF //col:1517
CPUID_EDX_ZERO =                                               0xFFFFFFFF //col:1526
CPUID_EAX_SUB_LEAF_TYPE =                                      0x0F //col:1538
CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1 =                        0xFFFFF000 //col:1541
CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2 =                        0xFFFFF //col:1550
CPUID_ECX_EPC_SECTION_PROPERTY =                               0x0F //col:1560
CPUID_ECX_EPC_SIZE_1 =                                         0xFFFFF000 //col:1563
CPUID_EDX_EPC_SIZE_2 =                                         0xFFFFF //col:1572
CPUID_INTEL_PROCESSOR_TRACE =                                  0x00000014 //col:1590
CPUID_EAX_MAX_SUB_LEAF =                                       0xFFFFFFFF //col:1595
CPUID_EBX_FLAG0 =                                              0x01 //col:1604
CPUID_EBX_FLAG1 =                                              0x02 //col:1606
CPUID_EBX_FLAG2 =                                              0x04 //col:1608
CPUID_EBX_FLAG3 =                                              0x08 //col:1610
CPUID_EBX_FLAG4 =                                              0x10 //col:1612
CPUID_EBX_FLAG5 =                                              0x20 //col:1614
CPUID_EBX_FLAG6 =                                              0x40 //col:1616
CPUID_EBX_FLAG7 =                                              0x80 //col:1618
CPUID_EBX_FLAG8 =                                              0x100 //col:1620
CPUID_ECX_FLAG0 =                                              0x01 //col:1630
CPUID_ECX_FLAG1 =                                              0x02 //col:1632
CPUID_ECX_FLAG2 =                                              0x04 //col:1634
CPUID_ECX_FLAG3 =                                              0x08 //col:1636
CPUID_ECX_FLAG31 =                                             0x80000000 //col:1639
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1648
CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING = 0x07 //col:1660
CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS =           0xFFFF0000 //col:1663
CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS = 0xFFFF //col:1672
CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS = 0xFFFF0000 //col:1674
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:1683
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1692
CPUID_TIME_STAMP_COUNTER =                                     0x00000015 //col:1704
CPUID_EAX_DENOMINATOR =                                        0xFFFFFFFF //col:1709
CPUID_EBX_NUMERATOR =                                          0xFFFFFFFF //col:1718
CPUID_ECX_NOMINAL_FREQUENCY =                                  0xFFFFFFFF //col:1727
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1736
CPUID_PROCESSOR_FREQUENCY =                                    0x00000016 //col:1744
CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ =                        0xFFFF //col:1749
CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ =                    0xFFFF //col:1759
CPUID_ECX_BUS_FREQUENCY_MHZ =                                  0xFFFF //col:1769
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1779
CPUID_SOC_VENDOR =                                             0x00000017 //col:1792
CPUID_EAX_MAX_SOC_ID_INDEX =                                   0xFFFFFFFF //col:1797
CPUID_EBX_SOC_VENDOR_ID =                                      0xFFFF //col:1806
CPUID_EBX_IS_VENDOR_SCHEME =                                   0x10000 //col:1808
CPUID_ECX_PROJECT_ID =                                         0xFFFFFFFF //col:1818
CPUID_EDX_STEPPING_ID =                                        0xFFFFFFFF //col:1827
CPUID_EAX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:1839
CPUID_EBX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:1848
CPUID_ECX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:1857
CPUID_EDX_SOC_VENDOR_BRAND_STRING =                            0xFFFFFFFF //col:1866
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:1878
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:1887
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:1896
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:1905
CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS =           0x00000018 //col:1922
CPUID_EAX_MAX_SUB_LEAF =                                       0xFFFFFFFF //col:1927
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED =                         0x01 //col:1936
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED =                         0x02 //col:1938
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED =                         0x04 //col:1940
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED =                         0x08 //col:1942
CPUID_EBX_PARTITIONING =                                       0x700 //col:1945
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00 =                           0xFFFF0000 //col:1948
CPUID_ECX_NUMBER_OF_SETS =                                     0xFFFFFFFF //col:1957
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD =                       0x1F //col:1966
CPUID_EDX_TRANSLATION_CACHE_LEVEL =                            0xE0 //col:1968
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE =                        0x100 //col:1970
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS =         0x3FFC000 //col:1973
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:1986
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED =                         0x01 //col:1995
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED =                         0x02 //col:1997
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED =                         0x04 //col:1999
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED =                         0x08 //col:2001
CPUID_EBX_PARTITIONING =                                       0x700 //col:2004
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01 =                           0xFFFF0000 //col:2007
CPUID_ECX_NUMBER_OF_SETS =                                     0xFFFFFFFF //col:2016
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD =                       0x1F //col:2025
CPUID_EDX_TRANSLATION_CACHE_LEVEL =                            0xE0 //col:2027
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE =                        0x100 //col:2029
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS =         0x3FFC000 //col:2032
CPUID_EXTENDED_FUNCTION =                                      0x80000000 //col:2045
CPUID_EAX_MAX_EXTENDED_FUNCTIONS =                             0xFFFFFFFF //col:2050
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:2059
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:2068
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:2077
CPUID_EXTENDED_CPU_SIGNATURE =                                 0x80000001 //col:2085
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:2090
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:2099
CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE =                 0x01 //col:2108
CPUID_ECX_LZCNT =                                              0x20 //col:2111
CPUID_ECX_PREFETCHW =                                          0x100 //col:2114
CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE =            0x800 //col:2125
CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE =                      0x100000 //col:2128
CPUID_EDX_PAGES_1GB_AVAILABLE =                                0x4000000 //col:2131
CPUID_EDX_RDTSCP_AVAILABLE =                                   0x8000000 //col:2133
CPUID_EDX_IA64_AVAILABLE =                                     0x20000000 //col:2136
CPUID_BRAND_STRING1 =                                          0x80000002 //col:2145
CPUID_BRAND_STRING2 =                                          0x80000003 //col:2146
CPUID_BRAND_STRING3 =                                          0x80000004 //col:2147
CPUID_EAX_PROCESSOR_BRAND_STRING_1 =                           0xFFFFFFFF //col:2152
CPUID_EBX_PROCESSOR_BRAND_STRING_2 =                           0xFFFFFFFF //col:2161
CPUID_ECX_PROCESSOR_BRAND_STRING_3 =                           0xFFFFFFFF //col:2170
CPUID_EDX_PROCESSOR_BRAND_STRING_4 =                           0xFFFFFFFF //col:2179
CPUID_EAX_PROCESSOR_BRAND_STRING_5 =                           0xFFFFFFFF //col:2191
CPUID_EBX_PROCESSOR_BRAND_STRING_6 =                           0xFFFFFFFF //col:2200
CPUID_ECX_PROCESSOR_BRAND_STRING_7 =                           0xFFFFFFFF //col:2209
CPUID_EDX_PROCESSOR_BRAND_STRING_8 =                           0xFFFFFFFF //col:2218
CPUID_EAX_PROCESSOR_BRAND_STRING_9 =                           0xFFFFFFFF //col:2230
CPUID_EBX_PROCESSOR_BRAND_STRING_10 =                          0xFFFFFFFF //col:2239
CPUID_ECX_PROCESSOR_BRAND_STRING_11 =                          0xFFFFFFFF //col:2248
CPUID_EDX_PROCESSOR_BRAND_STRING_12 =                          0xFFFFFFFF //col:2257
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:2269
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:2278
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:2287
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:2296
CPUID_EXTENDED_CACHE_INFO =                                    0x80000006 //col:2304
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:2309
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:2318
CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES =                           0xFF //col:2327
CPUID_ECX_L2_ASSOCIATIVITY_FIELD =                             0xF000 //col:2330
CPUID_ECX_CACHE_SIZE_IN_1K_UNITS =                             0xFFFF0000 //col:2332
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:2341
CPUID_EXTENDED_TIME_STAMP_COUNTER =                            0x80000007 //col:2349
CPUID_EAX_RESERVED =                                           0xFFFFFFFF //col:2354
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:2363
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:2372
CPUID_EDX_INVARIANT_TSC_AVAILABLE =                            0x100 //col:2382
CPUID_EXTENDED_VIRT_PHYS_ADDRESS_SIZE =                        0x80000008 //col:2391
CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS =                    0xFF //col:2396
CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS =                      0xFF00 //col:2398
CPUID_EBX_RESERVED =                                           0xFFFFFFFF //col:2408
CPUID_ECX_RESERVED =                                           0xFFFFFFFF //col:2417
CPUID_EDX_RESERVED =                                           0xFFFFFFFF //col:2426
IA32_P5_MC_ADDR =                                              0x00000000 //col:2448
IA32_P5_MC_TYPE =                                              0x00000001 //col:2449
IA32_MONITOR_FILTER_SIZE =                                     0x00000006 //col:2454
IA32_TIME_STAMP_COUNTER =                                      0x00000010 //col:2455
IA32_PLATFORM_ID =                                             0x00000017 //col:2456
IA32_PLATFORM_ID_PLATFORM_ID =                                 0x1C000000000000 //col:2461
IA32_APIC_BASE =                                               0x0000001B //col:2468
IA32_APIC_BASE_BSP_FLAG =                                      0x100 //col:2473
IA32_APIC_BASE_ENABLE_X2APIC_MODE =                            0x400 //col:2476
IA32_APIC_BASE_APIC_GLOBAL_ENABLE =                            0x800 //col:2478
IA32_APIC_BASE_APIC_BASE =                                     0xFFFFFFFFF000 //col:2480
IA32_FEATURE_CONTROL =                                         0x0000003A //col:2487
IA32_FEATURE_CONTROL_LOCK_BIT =                                0x01 //col:2491
IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX =                   0x02 //col:2493
IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX =                  0x04 //col:2495
IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES =           0x7F00 //col:2498
IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE =                    0x8000 //col:2500
IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE =               0x20000 //col:2503
IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE =                       0x40000 //col:2505
IA32_FEATURE_CONTROL_LMCE_ON =                                 0x100000 //col:2508
IA32_TSC_ADJUST =                                              0x0000003B //col:2515
IA32_SPEC_CTRL =                                               0x00000048 //col:2520
IA32_SPEC_CTRL_IBRS =                                          0x01 //col:2524
IA32_SPEC_CTRL_STIBP =                                         0x02 //col:2526
IA32_SPEC_CTRL_SSBD =                                          0x04 //col:2528
IA32_PRED_CMD =                                                0x00000049 //col:2535
IA32_PRED_CMD_IBPB =                                           0x01 //col:2539
IA32_BIOS_UPDT_TRIG =                                          0x00000079 //col:2546
IA32_BIOS_SIGN_ID =                                            0x0000008B //col:2547
IA32_BIOS_SIGN_ID_RESERVED =                                   0xFFFFFFFF //col:2551
IA32_BIOS_SIGN_ID_MICROCODE_UPDATE_SIGNATURE =                 0xFFFFFFFF00000000 //col:2553
IA32_SGXLEPUBKEYHASH0 =                                        0x0000008C //col:2564
IA32_SGXLEPUBKEYHASH1 =                                        0x0000008D //col:2565
IA32_SGXLEPUBKEYHASH2 =                                        0x0000008E //col:2566
IA32_SGXLEPUBKEYHASH3 =                                        0x0000008F //col:2567
IA32_SMM_MONITOR_CTL =                                         0x0000009B //col:2572
IA32_SMM_MONITOR_CTL_VALID =                                   0x01 //col:2576
IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF =                0x04 //col:2579
IA32_SMM_MONITOR_CTL_MSEG_BASE =                               0xFFFFF000 //col:2582
IA32_STM_FEATURES_IA32E =                                      0x00000001 //col:2593
IA32_SMBASE =                                                  0x0000009E //col:2602
IA32_PMC0 =                                                    0x000000C1 //col:2608
IA32_PMC1 =                                                    0x000000C2 //col:2609
IA32_PMC2 =                                                    0x000000C3 //col:2610
IA32_PMC3 =                                                    0x000000C4 //col:2611
IA32_PMC4 =                                                    0x000000C5 //col:2612
IA32_PMC5 =                                                    0x000000C6 //col:2613
IA32_PMC6 =                                                    0x000000C7 //col:2614
IA32_PMC7 =                                                    0x000000C8 //col:2615
IA32_MPERF =                                                   0x000000E7 //col:2620
IA32_APERF =                                                   0x000000E8 //col:2625
IA32_MTRRCAP =                                                 0x000000FE //col:2630
IA32_MTRRCAP_VARIABLE_RANGE_REGISTERS_COUNT =                  0xFF //col:2634
IA32_MTRRCAP_FIXED_RANGE_REGISTERS_SUPPORTED =                 0x100 //col:2636
IA32_MTRRCAP_WRITE_COMBINING =                                 0x400 //col:2639
IA32_MTRRCAP_SYSTEM_MANAGEMENT_RANGE_REGISTER =                0x800 //col:2641
IA32_ARCH_CAPABILITIES =                                       0x0000010A //col:2648
IA32_ARCH_CAPABILITIES_RDCL_NO =                               0x01 //col:2652
IA32_ARCH_CAPABILITIES_IBRS_ALL =                              0x02 //col:2654
IA32_ARCH_CAPABILITIES_RSBA =                                  0x04 //col:2656
IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY =                    0x08 //col:2658
IA32_ARCH_CAPABILITIES_SSB_NO =                                0x10 //col:2660
IA32_ARCH_CAPABILITIES_MDS_NO =                                0x20 //col:2662
IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO =                     0x40 //col:2664
IA32_ARCH_CAPABILITIES_TSX_CTRL =                              0x80 //col:2666
IA32_ARCH_CAPABILITIES_TAA_NO =                                0x100 //col:2668
IA32_FLUSH_CMD =                                               0x0000010B //col:2675
IA32_FLUSH_CMD_L1D_FLUSH =                                     0x01 //col:2679
IA32_TSX_CTRL =                                                0x00000122 //col:2686
IA32_TSX_CTRL_RTM_DISABLE =                                    0x01 //col:2690
IA32_TSX_CTRL_TSX_CPUID_CLEAR =                                0x02 //col:2692
IA32_SYSENTER_CS =                                             0x00000174 //col:2699
IA32_SYSENTER_CS_CS_SELECTOR =                                 0xFFFF //col:2703
IA32_SYSENTER_CS_NOT_USED_1 =                                  0xFFFF0000 //col:2705
IA32_SYSENTER_CS_NOT_USED_2 =                                  0xFFFFFFFF00000000 //col:2707
IA32_SYSENTER_ESP =                                            0x00000175 //col:2713
IA32_SYSENTER_EIP =                                            0x00000176 //col:2714
IA32_MCG_CAP =                                                 0x00000179 //col:2715
IA32_MCG_CAP_COUNT =                                           0xFF //col:2719
IA32_MCG_CAP_MCG_CTL_P =                                       0x100 //col:2721
IA32_MCG_CAP_MCG_EXT_P =                                       0x200 //col:2723
IA32_MCG_CAP_MCP_CMCI_P =                                      0x400 //col:2725
IA32_MCG_CAP_MCG_TES_P =                                       0x800 //col:2727
IA32_MCG_CAP_MCG_EXT_CNT =                                     0xFF0000 //col:2730
IA32_MCG_CAP_MCG_SER_P =                                       0x1000000 //col:2732
IA32_MCG_CAP_MCG_ELOG_P =                                      0x4000000 //col:2735
IA32_MCG_CAP_MCG_LMCE_P =                                      0x8000000 //col:2737
IA32_MCG_STATUS =                                              0x0000017A //col:2744
IA32_MCG_STATUS_RIPV =                                         0x01 //col:2748
IA32_MCG_STATUS_EIPV =                                         0x02 //col:2750
IA32_MCG_STATUS_MCIP =                                         0x04 //col:2752
IA32_MCG_STATUS_LMCE_S =                                       0x08 //col:2754
IA32_MCG_CTL =                                                 0x0000017B //col:2761
IA32_PERFEVTSEL0 =                                             0x00000186 //col:2767
IA32_PERFEVTSEL1 =                                             0x00000187 //col:2768
IA32_PERFEVTSEL2 =                                             0x00000188 //col:2769
IA32_PERFEVTSEL3 =                                             0x00000189 //col:2770
IA32_PERFEVTSEL_EVENT_SELECT =                                 0xFF //col:2774
IA32_PERFEVTSEL_U_MASK =                                       0xFF00 //col:2776
IA32_PERFEVTSEL_USR =                                          0x10000 //col:2778
IA32_PERFEVTSEL_OS =                                           0x20000 //col:2780
IA32_PERFEVTSEL_EDGE =                                         0x40000 //col:2782
IA32_PERFEVTSEL_PC =                                           0x80000 //col:2784
IA32_PERFEVTSEL_INTR =                                         0x100000 //col:2786
IA32_PERFEVTSEL_ANY_THREAD =                                   0x200000 //col:2788
IA32_PERFEVTSEL_EN =                                           0x400000 //col:2790
IA32_PERFEVTSEL_INV =                                          0x800000 //col:2792
IA32_PERFEVTSEL_CMASK =                                        0xFF000000 //col:2794
IA32_PERF_STATUS =                                             0x00000198 //col:2805
IA32_PERF_STATUS_CURRENT_PERFORMANCE_STATE_VALUE =             0xFFFF //col:2809
IA32_PERF_CTL =                                                0x00000199 //col:2816
IA32_PERF_CTL_TARGET_PERFORMANCE_STATE_VALUE =                 0xFFFF //col:2820
IA32_PERF_CTL_IDA_ENGAGE =                                     0x100000000 //col:2823
IA32_CLOCK_MODULATION =                                        0x0000019A //col:2830
IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE = 0x01 //col:2834
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE =  0x0E //col:2836
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE =      0x10 //col:2838
IA32_THERM_INTERRUPT =                                         0x0000019B //col:2845
IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE =       0x01 //col:2849
IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE =        0x02 //col:2851
IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE =                0x04 //col:2853
IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE =                0x08 //col:2855
IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE =   0x10 //col:2857
IA32_THERM_INTERRUPT_THRESHOLD1_VALUE =                        0x7F00 //col:2860
IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE =             0x8000 //col:2862
IA32_THERM_INTERRUPT_THRESHOLD2_VALUE =                        0x7F0000 //col:2864
IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE =             0x800000 //col:2866
IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE =         0x1000000 //col:2868
IA32_THERM_STATUS =                                            0x0000019C //col:2875
IA32_THERM_STATUS_THERMAL_STATUS =                             0x01 //col:2879
IA32_THERM_STATUS_THERMAL_STATUS_LOG =                         0x02 //col:2881
IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT =                      0x04 //col:2883
IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG =                        0x08 //col:2885
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS =                0x10 //col:2887
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG =            0x20 //col:2889
IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS =                  0x40 //col:2891
IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG =                     0x80 //col:2893
IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS =                  0x100 //col:2895
IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG =                     0x200 //col:2897
IA32_THERM_STATUS_POWER_LIMITATION_STATUS =                    0x400 //col:2899
IA32_THERM_STATUS_POWER_LIMITATION_LOG =                       0x800 //col:2901
IA32_THERM_STATUS_CURRENT_LIMIT_STATUS =                       0x1000 //col:2903
IA32_THERM_STATUS_CURRENT_LIMIT_LOG =                          0x2000 //col:2905
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS =                  0x4000 //col:2907
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG =                     0x8000 //col:2909
IA32_THERM_STATUS_DIGITAL_READOUT =                            0x7F0000 //col:2911
IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS =              0x78000000 //col:2914
IA32_THERM_STATUS_READING_VALID =                              0x80000000 //col:2916
IA32_MISC_ENABLE =                                             0x000001A0 //col:2923
IA32_MISC_ENABLE_FAST_STRINGS_ENABLE =                         0x01 //col:2927
IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE =    0x08 //col:2930
IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE =            0x80 //col:2933
IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE =            0x800 //col:2936
IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE =  0x1000 //col:2938
IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE =  0x10000 //col:2941
IA32_MISC_ENABLE_ENABLE_MONITOR_FSM =                          0x40000 //col:2944
IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL =                          0x400000 //col:2947
IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE =                        0x800000 //col:2949
IA32_MISC_ENABLE_XD_BIT_DISABLE =                              0x400000000 //col:2952
IA32_ENERGY_PERF_BIAS =                                        0x000001B0 //col:2959
IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE =                0x0F //col:2963
IA32_PACKAGE_THERM_STATUS =                                    0x000001B1 //col:2970
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS =                     0x01 //col:2974
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG =                 0x02 //col:2976
IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT =                      0x04 //col:2978
IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG =                        0x08 //col:2980
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS =        0x10 //col:2982
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG =    0x20 //col:2984
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS =          0x40 //col:2986
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG =             0x80 //col:2988
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS =          0x100 //col:2990
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG =             0x200 //col:2992
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS =            0x400 //col:2994
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG =               0x800 //col:2996
IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT =                    0x7F0000 //col:2999
IA32_PACKAGE_THERM_INTERRUPT =                                 0x000001B2 //col:3006
IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE = 0x01 //col:3010
IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE = 0x02 //col:3012
IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE =        0x04 //col:3014
IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE =       0x10 //col:3017
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE =                0x7F00 //col:3020
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE =     0x8000 //col:3022
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE =                0x7F0000 //col:3024
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE =     0x800000 //col:3026
IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE = 0x1000000 //col:3028
IA32_DEBUGCTL =                                                0x000001D9 //col:3035
IA32_DEBUGCTL_LBR =                                            0x01 //col:3039
IA32_DEBUGCTL_BTF =                                            0x02 //col:3041
IA32_DEBUGCTL_TR =                                             0x40 //col:3044
IA32_DEBUGCTL_BTS =                                            0x80 //col:3046
IA32_DEBUGCTL_BTINT =                                          0x100 //col:3048
IA32_DEBUGCTL_BTS_OFF_OS =                                     0x200 //col:3050
IA32_DEBUGCTL_BTS_OFF_USR =                                    0x400 //col:3052
IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI =                             0x800 //col:3054
IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI =                          0x1000 //col:3056
IA32_DEBUGCTL_ENABLE_UNCORE_PMI =                              0x2000 //col:3058
IA32_DEBUGCTL_FREEZE_WHILE_SMM =                               0x4000 //col:3060
IA32_DEBUGCTL_RTM_DEBUG =                                      0x8000 //col:3062
IA32_SMRR_PHYSBASE =                                           0x000001F2 //col:3069
IA32_SMRR_PHYSBASE_TYPE =                                      0xFF //col:3073
IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS =                0xFFFFF000 //col:3076
IA32_SMRR_PHYSMASK =                                           0x000001F3 //col:3083
IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK =                         0x800 //col:3088
IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK =                   0xFFFFF000 //col:3090
IA32_PLATFORM_DCA_CAP =                                        0x000001F8 //col:3097
IA32_CPU_DCA_CAP =                                             0x000001F9 //col:3098
IA32_DCA_0_CAP =                                               0x000001FA //col:3099
IA32_DCA_0_CAP_DCA_ACTIVE =                                    0x01 //col:3103
IA32_DCA_0_CAP_TRANSACTION =                                   0x06 //col:3105
IA32_DCA_0_CAP_DCA_TYPE =                                      0x78 //col:3107
IA32_DCA_0_CAP_DCA_QUEUE_SIZE =                                0x780 //col:3109
IA32_DCA_0_CAP_DCA_DELAY =                                     0x1E000 //col:3112
IA32_DCA_0_CAP_SW_BLOCK =                                      0x1000000 //col:3115
IA32_DCA_0_CAP_HW_BLOCK =                                      0x4000000 //col:3118
IA32_MTRR_PHYSBASE_TYPE =                                      0xFF //col:3133
IA32_MTRR_PHYSBASE_PHYSICAL_ADDRES_BASE =                      0xFFFFFFFFF000 //col:3136
IA32_MTRR_PHYSBASE0 =                                          0x00000200 //col:3143
IA32_MTRR_PHYSBASE1 =                                          0x00000202 //col:3144
IA32_MTRR_PHYSBASE2 =                                          0x00000204 //col:3145
IA32_MTRR_PHYSBASE3 =                                          0x00000206 //col:3146
IA32_MTRR_PHYSBASE4 =                                          0x00000208 //col:3147
IA32_MTRR_PHYSBASE5 =                                          0x0000020A //col:3148
IA32_MTRR_PHYSBASE6 =                                          0x0000020C //col:3149
IA32_MTRR_PHYSBASE7 =                                          0x0000020E //col:3150
IA32_MTRR_PHYSBASE8 =                                          0x00000210 //col:3151
IA32_MTRR_PHYSBASE9 =                                          0x00000212 //col:3152
IA32_MTRR_PHYSMASK_VALID =                                     0x800 //col:3166
IA32_MTRR_PHYSMASK_PHYSICAL_ADDRES_MASK =                      0xFFFFFFFFF000 //col:3168
IA32_MTRR_PHYSMASK0 =                                          0x00000201 //col:3175
IA32_MTRR_PHYSMASK1 =                                          0x00000203 //col:3176
IA32_MTRR_PHYSMASK2 =                                          0x00000205 //col:3177
IA32_MTRR_PHYSMASK3 =                                          0x00000207 //col:3178
IA32_MTRR_PHYSMASK4 =                                          0x00000209 //col:3179
IA32_MTRR_PHYSMASK5 =                                          0x0000020B //col:3180
IA32_MTRR_PHYSMASK6 =                                          0x0000020D //col:3181
IA32_MTRR_PHYSMASK7 =                                          0x0000020F //col:3182
IA32_MTRR_PHYSMASK8 =                                          0x00000211 //col:3183
IA32_MTRR_PHYSMASK9 =                                          0x00000213 //col:3184
IA32_MTRR_FIX64K_BASE =                                        0x00000000 //col:3199
IA32_MTRR_FIX64K_SIZE =                                        0x00010000 //col:3200
IA32_MTRR_FIX64K_00000 =                                       0x00000250 //col:3201
IA32_MTRR_FIX16K_BASE =                                        0x00080000 //col:3211
IA32_MTRR_FIX16K_SIZE =                                        0x00004000 //col:3212
IA32_MTRR_FIX16K_80000 =                                       0x00000258 //col:3213
IA32_MTRR_FIX16K_A0000 =                                       0x00000259 //col:3214
IA32_MTRR_FIX4K_BASE =                                         0x000C0000 //col:3224
IA32_MTRR_FIX4K_SIZE =                                         0x00001000 //col:3225
IA32_MTRR_FIX4K_C0000 =                                        0x00000268 //col:3226
IA32_MTRR_FIX4K_C8000 =                                        0x00000269 //col:3227
IA32_MTRR_FIX4K_D0000 =                                        0x0000026A //col:3228
IA32_MTRR_FIX4K_D8000 =                                        0x0000026B //col:3229
IA32_MTRR_FIX4K_E0000 =                                        0x0000026C //col:3230
IA32_MTRR_FIX4K_E8000 =                                        0x0000026D //col:3231
IA32_MTRR_FIX4K_F0000 =                                        0x0000026E //col:3232
IA32_MTRR_FIX4K_F8000 =                                        0x0000026F //col:3233
IA32_MTRR_FIX_COUNT =                                          ((1 + 2 + 8) * 8) //col:3238
IA32_MTRR_VARIABLE_COUNT =                                     0x0000000A //col:3239
IA32_MTRR_COUNT =                                              (IA32_MTRR_FIX_COUNT + IA32_MTRR_VARIABLE_COUNT) //col:3240
IA32_PAT =                                                     0x00000277 //col:3245
IA32_PAT_PA0 =                                                 0x07 //col:3249
IA32_PAT_PA1 =                                                 0x700 //col:3252
IA32_PAT_PA2 =                                                 0x70000 //col:3255
IA32_PAT_PA3 =                                                 0x7000000 //col:3258
IA32_PAT_PA4 =                                                 0x700000000 //col:3261
IA32_PAT_PA5 =                                                 0x70000000000 //col:3264
IA32_PAT_PA6 =                                                 0x7000000000000 //col:3267
IA32_PAT_PA7 =                                                 0x700000000000000 //col:3270
IA32_MC0_CTL2 =                                                0x00000280 //col:3282
IA32_MC1_CTL2 =                                                0x00000281 //col:3283
IA32_MC2_CTL2 =                                                0x00000282 //col:3284
IA32_MC3_CTL2 =                                                0x00000283 //col:3285
IA32_MC4_CTL2 =                                                0x00000284 //col:3286
IA32_MC5_CTL2 =                                                0x00000285 //col:3287
IA32_MC6_CTL2 =                                                0x00000286 //col:3288
IA32_MC7_CTL2 =                                                0x00000287 //col:3289
IA32_MC8_CTL2 =                                                0x00000288 //col:3290
IA32_MC9_CTL2 =                                                0x00000289 //col:3291
IA32_MC10_CTL2 =                                               0x0000028A //col:3292
IA32_MC11_CTL2 =                                               0x0000028B //col:3293
IA32_MC12_CTL2 =                                               0x0000028C //col:3294
IA32_MC13_CTL2 =                                               0x0000028D //col:3295
IA32_MC14_CTL2 =                                               0x0000028E //col:3296
IA32_MC15_CTL2 =                                               0x0000028F //col:3297
IA32_MC16_CTL2 =                                               0x00000290 //col:3298
IA32_MC17_CTL2 =                                               0x00000291 //col:3299
IA32_MC18_CTL2 =                                               0x00000292 //col:3300
IA32_MC19_CTL2 =                                               0x00000293 //col:3301
IA32_MC20_CTL2 =                                               0x00000294 //col:3302
IA32_MC21_CTL2 =                                               0x00000295 //col:3303
IA32_MC22_CTL2 =                                               0x00000296 //col:3304
IA32_MC23_CTL2 =                                               0x00000297 //col:3305
IA32_MC24_CTL2 =                                               0x00000298 //col:3306
IA32_MC25_CTL2 =                                               0x00000299 //col:3307
IA32_MC26_CTL2 =                                               0x0000029A //col:3308
IA32_MC27_CTL2 =                                               0x0000029B //col:3309
IA32_MC28_CTL2 =                                               0x0000029C //col:3310
IA32_MC29_CTL2 =                                               0x0000029D //col:3311
IA32_MC30_CTL2 =                                               0x0000029E //col:3312
IA32_MC31_CTL2 =                                               0x0000029F //col:3313
IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD =                 0x7FFF //col:3317
IA32_MC_CTL2_CMCI_EN =                                         0x40000000 //col:3320
IA32_MTRR_DEF_TYPE =                                           0x000002FF //col:3331
IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE =                       0x07 //col:3335
IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE =                   0x400 //col:3338
IA32_MTRR_DEF_TYPE_MTRR_ENABLE =                               0x800 //col:3340
IA32_FIXED_CTR0 =                                              0x00000309 //col:3352
IA32_FIXED_CTR1 =                                              0x0000030A //col:3353
IA32_FIXED_CTR2 =                                              0x0000030B //col:3354
IA32_PERF_CAPABILITIES =                                       0x00000345 //col:3359
IA32_PERF_CAPABILITIES_LBR_FORMAT =                            0x3F //col:3363
IA32_PERF_CAPABILITIES_PEBS_TRAP =                             0x40 //col:3365
IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS =                   0x80 //col:3367
IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT =                    0xF00 //col:3369
IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED =         0x1000 //col:3371
IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE =              0x2000 //col:3373
IA32_FIXED_CTR_CTRL =                                          0x0000038D //col:3380
IA32_FIXED_CTR_CTRL_EN0_OS =                                   0x01 //col:3384
IA32_FIXED_CTR_CTRL_EN0_USR =                                  0x02 //col:3386
IA32_FIXED_CTR_CTRL_ANY_THREAD0 =                              0x04 //col:3388
IA32_FIXED_CTR_CTRL_EN0_PMI =                                  0x08 //col:3390
IA32_FIXED_CTR_CTRL_EN1_OS =                                   0x10 //col:3392
IA32_FIXED_CTR_CTRL_EN1_USR =                                  0x20 //col:3394
IA32_FIXED_CTR_CTRL_ANY_THREAD1 =                              0x40 //col:3396
IA32_FIXED_CTR_CTRL_EN1_PMI =                                  0x80 //col:3398
IA32_FIXED_CTR_CTRL_EN2_OS =                                   0x100 //col:3400
IA32_FIXED_CTR_CTRL_EN2_USR =                                  0x200 //col:3402
IA32_FIXED_CTR_CTRL_ANY_THREAD2 =                              0x400 //col:3404
IA32_FIXED_CTR_CTRL_EN2_PMI =                                  0x800 //col:3406
IA32_PERF_GLOBAL_STATUS =                                      0x0000038E //col:3413
IA32_PERF_GLOBAL_STATUS_OVF_PMC0 =                             0x01 //col:3417
IA32_PERF_GLOBAL_STATUS_OVF_PMC1 =                             0x02 //col:3419
IA32_PERF_GLOBAL_STATUS_OVF_PMC2 =                             0x04 //col:3421
IA32_PERF_GLOBAL_STATUS_OVF_PMC3 =                             0x08 //col:3423
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0 =                        0x100000000 //col:3426
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1 =                        0x200000000 //col:3428
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2 =                        0x400000000 //col:3430
IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI =                       0x80000000000000 //col:3433
IA32_PERF_GLOBAL_STATUS_LBR_FRZ =                              0x400000000000000 //col:3436
IA32_PERF_GLOBAL_STATUS_CTR_FRZ =                              0x800000000000000 //col:3438
IA32_PERF_GLOBAL_STATUS_ASCI =                                 0x1000000000000000 //col:3440
IA32_PERF_GLOBAL_STATUS_OVF_UNCORE =                           0x2000000000000000 //col:3442
IA32_PERF_GLOBAL_STATUS_OVF_BUF =                              0x4000000000000000 //col:3444
IA32_PERF_GLOBAL_STATUS_COND_CHGD =                            0x8000000000000000 //col:3446
IA32_PERF_GLOBAL_CTRL =                                        0x0000038F //col:3452
IA32_PERF_GLOBAL_CTRL_EN_PMCN =                                0xFFFFFFFF //col:3456
IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN =                          0xFFFFFFFF00000000 //col:3458
IA32_PERF_GLOBAL_STATUS_RESET =                                0x00000390 //col:3464
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN =                 0xFFFFFFFF //col:3468
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN =           0x700000000 //col:3470
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI =           0x80000000000000 //col:3473
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ =                  0x400000000000000 //col:3476
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ =                  0x800000000000000 //col:3478
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI =                     0x1000000000000000 //col:3480
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE =               0x2000000000000000 //col:3482
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF =                  0x4000000000000000 //col:3484
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD =                0x8000000000000000 //col:3486
IA32_PERF_GLOBAL_STATUS_SET =                                  0x00000391 //col:3492
IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN =                         0xFFFFFFFF //col:3496
IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN =                   0x700000000 //col:3498
IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI =                   0x80000000000000 //col:3501
IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ =                          0x400000000000000 //col:3504
IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ =                          0x800000000000000 //col:3506
IA32_PERF_GLOBAL_STATUS_SET_ASCI =                             0x1000000000000000 //col:3508
IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE =                       0x2000000000000000 //col:3510
IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF =                          0x4000000000000000 //col:3512
IA32_PERF_GLOBAL_INUSE =                                       0x00000392 //col:3519
IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE =               0xFFFFFFFF //col:3523
IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE =                0x700000000 //col:3525
IA32_PERF_GLOBAL_INUSE_PMI_IN_USE =                            0x8000000000000000 //col:3528
IA32_PEBS_ENABLE =                                             0x000003F1 //col:3534
IA32_PEBS_ENABLE_ENABLE_PEBS =                                 0x01 //col:3538
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1 =                    0x0E //col:3540
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2 =                    0xF00000000 //col:3543
IA32_MC0_CTL =                                                 0x00000400 //col:3555
IA32_MC1_CTL =                                                 0x00000404 //col:3556
IA32_MC2_CTL =                                                 0x00000408 //col:3557
IA32_MC3_CTL =                                                 0x0000040C //col:3558
IA32_MC4_CTL =                                                 0x00000410 //col:3559
IA32_MC5_CTL =                                                 0x00000414 //col:3560
IA32_MC6_CTL =                                                 0x00000418 //col:3561
IA32_MC7_CTL =                                                 0x0000041C //col:3562
IA32_MC8_CTL =                                                 0x00000420 //col:3563
IA32_MC9_CTL =                                                 0x00000424 //col:3564
IA32_MC10_CTL =                                                0x00000428 //col:3565
IA32_MC11_CTL =                                                0x0000042C //col:3566
IA32_MC12_CTL =                                                0x00000430 //col:3567
IA32_MC13_CTL =                                                0x00000434 //col:3568
IA32_MC14_CTL =                                                0x00000438 //col:3569
IA32_MC15_CTL =                                                0x0000043C //col:3570
IA32_MC16_CTL =                                                0x00000440 //col:3571
IA32_MC17_CTL =                                                0x00000444 //col:3572
IA32_MC18_CTL =                                                0x00000448 //col:3573
IA32_MC19_CTL =                                                0x0000044C //col:3574
IA32_MC20_CTL =                                                0x00000450 //col:3575
IA32_MC21_CTL =                                                0x00000454 //col:3576
IA32_MC22_CTL =                                                0x00000458 //col:3577
IA32_MC23_CTL =                                                0x0000045C //col:3578
IA32_MC24_CTL =                                                0x00000460 //col:3579
IA32_MC25_CTL =                                                0x00000464 //col:3580
IA32_MC26_CTL =                                                0x00000468 //col:3581
IA32_MC27_CTL =                                                0x0000046C //col:3582
IA32_MC28_CTL =                                                0x00000470 //col:3583
IA32_MC0_STATUS =                                              0x00000401 //col:3593
IA32_MC1_STATUS =                                              0x00000405 //col:3594
IA32_MC2_STATUS =                                              0x00000409 //col:3595
IA32_MC3_STATUS =                                              0x0000040D //col:3596
IA32_MC4_STATUS =                                              0x00000411 //col:3597
IA32_MC5_STATUS =                                              0x00000415 //col:3598
IA32_MC6_STATUS =                                              0x00000419 //col:3599
IA32_MC7_STATUS =                                              0x0000041D //col:3600
IA32_MC8_STATUS =                                              0x00000421 //col:3601
IA32_MC9_STATUS =                                              0x00000425 //col:3602
IA32_MC10_STATUS =                                             0x00000429 //col:3603
IA32_MC11_STATUS =                                             0x0000042D //col:3604
IA32_MC12_STATUS =                                             0x00000431 //col:3605
IA32_MC13_STATUS =                                             0x00000435 //col:3606
IA32_MC14_STATUS =                                             0x00000439 //col:3607
IA32_MC15_STATUS =                                             0x0000043D //col:3608
IA32_MC16_STATUS =                                             0x00000441 //col:3609
IA32_MC17_STATUS =                                             0x00000445 //col:3610
IA32_MC18_STATUS =                                             0x00000449 //col:3611
IA32_MC19_STATUS =                                             0x0000044D //col:3612
IA32_MC20_STATUS =                                             0x00000451 //col:3613
IA32_MC21_STATUS =                                             0x00000455 //col:3614
IA32_MC22_STATUS =                                             0x00000459 //col:3615
IA32_MC23_STATUS =                                             0x0000045D //col:3616
IA32_MC24_STATUS =                                             0x00000461 //col:3617
IA32_MC25_STATUS =                                             0x00000465 //col:3618
IA32_MC26_STATUS =                                             0x00000469 //col:3619
IA32_MC27_STATUS =                                             0x0000046D //col:3620
IA32_MC28_STATUS =                                             0x00000471 //col:3621
IA32_MC0_ADDR =                                                0x00000402 //col:3631
IA32_MC1_ADDR =                                                0x00000406 //col:3632
IA32_MC2_ADDR =                                                0x0000040A //col:3633
IA32_MC3_ADDR =                                                0x0000040E //col:3634
IA32_MC4_ADDR =                                                0x00000412 //col:3635
IA32_MC5_ADDR =                                                0x00000416 //col:3636
IA32_MC6_ADDR =                                                0x0000041A //col:3637
IA32_MC7_ADDR =                                                0x0000041E //col:3638
IA32_MC8_ADDR =                                                0x00000422 //col:3639
IA32_MC9_ADDR =                                                0x00000426 //col:3640
IA32_MC10_ADDR =                                               0x0000042A //col:3641
IA32_MC11_ADDR =                                               0x0000042E //col:3642
IA32_MC12_ADDR =                                               0x00000432 //col:3643
IA32_MC13_ADDR =                                               0x00000436 //col:3644
IA32_MC14_ADDR =                                               0x0000043A //col:3645
IA32_MC15_ADDR =                                               0x0000043E //col:3646
IA32_MC16_ADDR =                                               0x00000442 //col:3647
IA32_MC17_ADDR =                                               0x00000446 //col:3648
IA32_MC18_ADDR =                                               0x0000044A //col:3649
IA32_MC19_ADDR =                                               0x0000044E //col:3650
IA32_MC20_ADDR =                                               0x00000452 //col:3651
IA32_MC21_ADDR =                                               0x00000456 //col:3652
IA32_MC22_ADDR =                                               0x0000045A //col:3653
IA32_MC23_ADDR =                                               0x0000045E //col:3654
IA32_MC24_ADDR =                                               0x00000462 //col:3655
IA32_MC25_ADDR =                                               0x00000466 //col:3656
IA32_MC26_ADDR =                                               0x0000046A //col:3657
IA32_MC27_ADDR =                                               0x0000046E //col:3658
IA32_MC28_ADDR =                                               0x00000472 //col:3659
IA32_MC0_MISC =                                                0x00000403 //col:3669
IA32_MC1_MISC =                                                0x00000407 //col:3670
IA32_MC2_MISC =                                                0x0000040B //col:3671
IA32_MC3_MISC =                                                0x0000040F //col:3672
IA32_MC4_MISC =                                                0x00000413 //col:3673
IA32_MC5_MISC =                                                0x00000417 //col:3674
IA32_MC6_MISC =                                                0x0000041B //col:3675
IA32_MC7_MISC =                                                0x0000041F //col:3676
IA32_MC8_MISC =                                                0x00000423 //col:3677
IA32_MC9_MISC =                                                0x00000427 //col:3678
IA32_MC10_MISC =                                               0x0000042B //col:3679
IA32_MC11_MISC =                                               0x0000042F //col:3680
IA32_MC12_MISC =                                               0x00000433 //col:3681
IA32_MC13_MISC =                                               0x00000437 //col:3682
IA32_MC14_MISC =                                               0x0000043B //col:3683
IA32_MC15_MISC =                                               0x0000043F //col:3684
IA32_MC16_MISC =                                               0x00000443 //col:3685
IA32_MC17_MISC =                                               0x00000447 //col:3686
IA32_MC18_MISC =                                               0x0000044B //col:3687
IA32_MC19_MISC =                                               0x0000044F //col:3688
IA32_MC20_MISC =                                               0x00000453 //col:3689
IA32_MC21_MISC =                                               0x00000457 //col:3690
IA32_MC22_MISC =                                               0x0000045B //col:3691
IA32_MC23_MISC =                                               0x0000045F //col:3692
IA32_MC24_MISC =                                               0x00000463 //col:3693
IA32_MC25_MISC =                                               0x00000467 //col:3694
IA32_MC26_MISC =                                               0x0000046B //col:3695
IA32_MC27_MISC =                                               0x0000046F //col:3696
IA32_MC28_MISC =                                               0x00000473 //col:3697
IA32_VMX_BASIC =                                               0x00000480 //col:3702
IA32_VMX_BASIC_VMCS_REVISION_ID =                              0x7FFFFFFF //col:3706
IA32_VMX_BASIC_MUST_BE_ZERO =                                  0x80000000 //col:3708
IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES =                            0x1FFF00000000 //col:3710
IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH =                   0x1000000000000 //col:3713
IA32_VMX_BASIC_DUAL_MONITOR =                                  0x2000000000000 //col:3715
IA32_VMX_BASIC_MEMORY_TYPE =                                   0x3C000000000000 //col:3717
IA32_VMX_BASIC_INS_OUTS_VMEXIT_INFORMATION =                   0x40000000000000 //col:3719
IA32_VMX_BASIC_TRUE_CONTROLS =                                 0x80000000000000 //col:3721
IA32_VMX_PINBASED_CTLS =                                       0x00000481 //col:3728
IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING =            0x01 //col:3732
IA32_VMX_PINBASED_CTLS_NMI_EXITING =                           0x08 //col:3735
IA32_VMX_PINBASED_CTLS_VIRTUAL_NMIS =                          0x20 //col:3738
IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER =         0x40 //col:3740
IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS =             0x80 //col:3742
IA32_VMX_PROCBASED_CTLS =                                      0x00000482 //col:3749
IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING =             0x04 //col:3754
IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING =                   0x08 //col:3756
IA32_VMX_PROCBASED_CTLS_HLT_EXITING =                          0x80 //col:3759
IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING =                       0x200 //col:3762
IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING =                        0x400 //col:3764
IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING =                        0x800 //col:3766
IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING =                        0x1000 //col:3768
IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING =                     0x8000 //col:3771
IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING =                    0x10000 //col:3773
IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS =           0x20000 //col:3775
IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING =                     0x80000 //col:3778
IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING =                    0x100000 //col:3780
IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW =                       0x200000 //col:3782
IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING =                   0x400000 //col:3784
IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING =                       0x800000 //col:3786
IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING =             0x1000000 //col:3788
IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS =                       0x2000000 //col:3790
IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG =                    0x8000000 //col:3793
IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS =                      0x10000000 //col:3795
IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING =                      0x20000000 //col:3797
IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING =                        0x40000000 //col:3799
IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS =          0x80000000 //col:3801
IA32_VMX_EXIT_CTLS =                                           0x00000483 //col:3808
IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS =                       0x04 //col:3813
IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE =                   0x200 //col:3816
IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL =                0x1000 //col:3819
IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT =             0x8000 //col:3822
IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT =                             0x40000 //col:3825
IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT =                             0x80000 //col:3827
IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER =                            0x100000 //col:3829
IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER =                            0x200000 //col:3831
IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE =           0x400000 //col:3833
IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS =                        0x800000 //col:3835
IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT =                       0x1000000 //col:3837
IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL =                       0x2000000 //col:3839
IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL =                        0x4000000 //col:3841
IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE =                       0x10000000 //col:3844
IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS =                            0x20000000 //col:3846
IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS =               0x80000000 //col:3849
IA32_VMX_ENTRY_CTLS =                                          0x00000484 //col:3856
IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS =                      0x04 //col:3861
IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST =                         0x200 //col:3864
IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM =                             0x400 //col:3866
IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT =        0x800 //col:3868
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL =               0x2000 //col:3871
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT =                            0x4000 //col:3873
IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER =                           0x8000 //col:3875
IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS =                        0x10000 //col:3877
IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT =                      0x20000 //col:3879
IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL =                       0x40000 //col:3881
IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE =                           0x100000 //col:3884
IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL =                        0x200000 //col:3886
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS =                           0x400000 //col:3888
IA32_VMX_MISC =                                                0x00000485 //col:3895
IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP =              0x1F //col:3899
IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT =                       0x20 //col:3901
IA32_VMX_MISC_ACTIVITY_STATES =                                0x1C0 //col:3903
IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX =                      0x4000 //col:3906
IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM =          0x8000 //col:3908
IA32_VMX_MISC_CR3_TARGET_COUNT =                               0x1FF0000 //col:3910
IA32_VMX_MISC_MAX_NUMBER_OF_MSR =                              0xE000000 //col:3912
IA32_VMX_MISC_SMM_MONITOR_CTL_B2 =                             0x10000000 //col:3914
IA32_VMX_MISC_VMWRITE_VMEXIT_INFO =                            0x20000000 //col:3916
IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION =      0x40000000 //col:3918
IA32_VMX_MISC_MSEG_ID =                                        0xFFFFFFFF00000000 //col:3921
IA32_VMX_CR0_FIXED0 =                                          0x00000486 //col:3927
IA32_VMX_CR0_FIXED1 =                                          0x00000487 //col:3928
IA32_VMX_CR4_FIXED0 =                                          0x00000488 //col:3929
IA32_VMX_CR4_FIXED1 =                                          0x00000489 //col:3930
IA32_VMX_VMCS_ENUM =                                           0x0000048A //col:3931
IA32_VMX_VMCS_ENUM_ACCESS_TYPE =                               0x01 //col:3935
IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE =                       0x3FE //col:3937
IA32_VMX_VMCS_ENUM_FIELD_TYPE =                                0xC00 //col:3939
IA32_VMX_VMCS_ENUM_FIELD_WIDTH =                               0x6000 //col:3942
IA32_VMX_PROCBASED_CTLS2 =                                     0x0000048B //col:3949
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES =            0x01 //col:3953
IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT =                          0x02 //col:3955
IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING =            0x04 //col:3957
IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP =                       0x08 //col:3959
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE =              0x10 //col:3961
IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID =                         0x20 //col:3963
IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING =                      0x40 //col:3965
IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST =                  0x80 //col:3967
IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION =        0x100 //col:3969
IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY =          0x200 //col:3971
IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING =                  0x400 //col:3973
IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING =                      0x800 //col:3975
IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID =                      0x1000 //col:3977
IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS =                 0x2000 //col:3979
IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING =                      0x4000 //col:3981
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING =                0x8000 //col:3983
IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING =                      0x10000 //col:3985
IA32_VMX_PROCBASED_CTLS2_ENABLE_PML =                          0x20000 //col:3987
IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION =                       0x40000 //col:3989
IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT =                 0x80000 //col:3991
IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES =                       0x100000 //col:3993
IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT =  0x400000 //col:3996
IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT =  0x800000 //col:3998
IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES =    0x1000000 //col:4000
IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING =                     0x2000000 //col:4002
IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE =              0x4000000 //col:4004
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING =                0x10000000 //col:4007
IA32_VMX_EPT_VPID_CAP =                                        0x0000048C //col:4014
IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES =                     0x01 //col:4018
IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4 =                     0x40 //col:4021
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE =                0x100 //col:4024
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK =                 0x4000 //col:4027
IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES =                          0x10000 //col:4030
IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES =                        0x20000 //col:4032
IA32_VMX_EPT_VPID_CAP_INVEPT =                                 0x100000 //col:4035
IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS =           0x200000 //col:4037
IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION = 0x400000 //col:4039
IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK =                0x800000 //col:4041
IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT =                  0x2000000 //col:4044
IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS =                    0x4000000 //col:4046
IA32_VMX_EPT_VPID_CAP_INVVPID =                                0x100000000 //col:4049
IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS =             0x10000000000 //col:4052
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT =                 0x20000000000 //col:4054
IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS =                   0x40000000000 //col:4056
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS =  0x80000000000 //col:4058
IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE =                   0x3F000000000000 //col:4061
IA32_VMX_TRUE_PINBASED_CTLS =                                  0x0000048D //col:4073
IA32_VMX_TRUE_PROCBASED_CTLS =                                 0x0000048E //col:4074
IA32_VMX_TRUE_EXIT_CTLS =                                      0x0000048F //col:4075
IA32_VMX_TRUE_ENTRY_CTLS =                                     0x00000490 //col:4076
IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS =                        0xFFFFFFFF //col:4080
IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS =                        0xFFFFFFFF00000000 //col:4082
IA32_VMX_VMFUNC =                                              0x00000491 //col:4092
IA32_VMX_VMFUNC_EPTP_SWITCHING =                               0x01 //col:4096
IA32_VMX_PROCBASED_CTLS3 =                                     0x00000492 //col:4103
IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING =                   0x01 //col:4107
IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT =                         0x02 //col:4109
IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE =                    0x04 //col:4111
IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING =                        0x08 //col:4113
IA32_VMX_EXIT_CTLS2 =                                          0x00000493 //col:4120
IA32_VMX_EXIT_CTLS2_RESERVED =                                 0xFFFFFFFFFFFFFFFF //col:4124
IA32_A_PMC0 =                                                  0x000004C1 //col:4135
IA32_A_PMC1 =                                                  0x000004C2 //col:4136
IA32_A_PMC2 =                                                  0x000004C3 //col:4137
IA32_A_PMC3 =                                                  0x000004C4 //col:4138
IA32_A_PMC4 =                                                  0x000004C5 //col:4139
IA32_A_PMC5 =                                                  0x000004C6 //col:4140
IA32_A_PMC6 =                                                  0x000004C7 //col:4141
IA32_A_PMC7 =                                                  0x000004C8 //col:4142
IA32_MCG_EXT_CTL =                                             0x000004D0 //col:4147
IA32_MCG_EXT_CTL_LMCE_EN =                                     0x01 //col:4151
IA32_SGX_SVN_STATUS =                                          0x00000500 //col:4158
IA32_SGX_SVN_STATUS_LOCK =                                     0x01 //col:4162
IA32_SGX_SVN_STATUS_SGX_SVN_SINIT =                            0xFF0000 //col:4165
IA32_RTIT_OUTPUT_BASE =                                        0x00000560 //col:4172
IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS =                  0xFFFFFFFFFF80 //col:4177
IA32_RTIT_OUTPUT_MASK_PTRS =                                   0x00000561 //col:4184
IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK =                        0x7F //col:4188
IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET =              0xFFFFFF80 //col:4190
IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET =                     0xFFFFFFFF00000000 //col:4192
IA32_RTIT_CTL =                                                0x00000570 //col:4198
IA32_RTIT_CTL_TRACE_EN =                                       0x01 //col:4202
IA32_RTIT_CTL_CYC_EN =                                         0x02 //col:4204
IA32_RTIT_CTL_OS =                                             0x04 //col:4206
IA32_RTIT_CTL_USER =                                           0x08 //col:4208
IA32_RTIT_CTL_PWR_EVT_EN =                                     0x10 //col:4210
IA32_RTIT_CTL_FUP_ON_PTW =                                     0x20 //col:4212
IA32_RTIT_CTL_FABRIC_EN =                                      0x40 //col:4214
IA32_RTIT_CTL_CR3_FILTER =                                     0x80 //col:4216
IA32_RTIT_CTL_TOPA =                                           0x100 //col:4218
IA32_RTIT_CTL_MTC_EN =                                         0x200 //col:4220
IA32_RTIT_CTL_TSC_EN =                                         0x400 //col:4222
IA32_RTIT_CTL_DIS_RETC =                                       0x800 //col:4224
IA32_RTIT_CTL_PTW_EN =                                         0x1000 //col:4226
IA32_RTIT_CTL_BRANCH_EN =                                      0x2000 //col:4228
IA32_RTIT_CTL_MTC_FREQ =                                       0x3C000 //col:4230
IA32_RTIT_CTL_CYC_THRESH =                                     0x780000 //col:4233
IA32_RTIT_CTL_PSB_FREQ =                                       0xF000000 //col:4236
IA32_RTIT_CTL_ADDR0_CFG =                                      0xF00000000 //col:4239
IA32_RTIT_CTL_ADDR1_CFG =                                      0xF000000000 //col:4241
IA32_RTIT_CTL_ADDR2_CFG =                                      0xF0000000000 //col:4243
IA32_RTIT_CTL_ADDR3_CFG =                                      0xF00000000000 //col:4245
IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE =                       0x100000000000000 //col:4248
IA32_RTIT_STATUS =                                             0x00000571 //col:4255
IA32_RTIT_STATUS_FILTER_EN =                                   0x01 //col:4259
IA32_RTIT_STATUS_CONTEX_EN =                                   0x02 //col:4261
IA32_RTIT_STATUS_TRIGGER_EN =                                  0x04 //col:4263
IA32_RTIT_STATUS_ERROR =                                       0x10 //col:4266
IA32_RTIT_STATUS_STOPPED =                                     0x20 //col:4268
IA32_RTIT_STATUS_PEND_PSB =                                    0x40 //col:4270
IA32_RTIT_STATUS_PEND_TOPA_PMI =                               0x80 //col:4272
IA32_RTIT_STATUS_PACKET_BYTE_CNT =                             0x1FFFF00000000 //col:4275
IA32_RTIT_CR3_MATCH =                                          0x00000572 //col:4282
IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH =                       0xFFFFFFFFFFFFFFE0 //col:4287
IA32_RTIT_ADDR0_A =                                            0x00000580 //col:4303
IA32_RTIT_ADDR1_A =                                            0x00000582 //col:4304
IA32_RTIT_ADDR2_A =                                            0x00000584 //col:4305
IA32_RTIT_ADDR3_A =                                            0x00000586 //col:4306
IA32_RTIT_ADDR0_B =                                            0x00000581 //col:4316
IA32_RTIT_ADDR1_B =                                            0x00000583 //col:4317
IA32_RTIT_ADDR2_B =                                            0x00000585 //col:4318
IA32_RTIT_ADDR3_B =                                            0x00000587 //col:4319
IA32_RTIT_ADDR_VIRTUAL_ADDRESS =                               0xFFFFFFFFFFFF //col:4327
IA32_RTIT_ADDR_SIGN_EXT_VA =                                   0xFFFF000000000000 //col:4329
IA32_DS_AREA =                                                 0x00000600 //col:4339
IA32_U_CET =                                                   0x000006A0 //col:4340
IA32_U_CET_SH_STK_EN =                                         0x01 //col:4344
IA32_U_CET_WR_SHSTK_EN =                                       0x02 //col:4346
IA32_U_CET_ENDBR_EN =                                          0x04 //col:4348
IA32_U_CET_LEG_IW_EN =                                         0x08 //col:4350
IA32_U_CET_NO_TRACK_EN =                                       0x10 //col:4352
IA32_U_CET_SUPPRESS_DIS =                                      0x20 //col:4354
IA32_U_CET_SUPPRESS =                                          0x400 //col:4357
IA32_U_CET_TRACKER =                                           0x800 //col:4359
IA32_U_CET_EB_LEG_BITMAP_BASE =                                0xFFFFFFFFFFFFF000 //col:4361
IA32_S_CET =                                                   0x000006A2 //col:4367
IA32_S_CET_SH_STK_EN =                                         0x01 //col:4371
IA32_S_CET_WR_SHSTK_EN =                                       0x02 //col:4373
IA32_S_CET_ENDBR_EN =                                          0x04 //col:4375
IA32_S_CET_LEG_IW_EN =                                         0x08 //col:4377
IA32_S_CET_NO_TRACK_EN =                                       0x10 //col:4379
IA32_S_CET_SUPPRESS_DIS =                                      0x20 //col:4381
IA32_S_CET_SUPPRESS =                                          0x400 //col:4384
IA32_S_CET_TRACKER =                                           0x800 //col:4386
IA32_S_CET_EB_LEG_BITMAP_BASE =                                0xFFFFFFFFFFFFF000 //col:4388
IA32_PL0_SSP =                                                 0x000006A4 //col:4394
IA32_PL1_SSP =                                                 0x000006A5 //col:4395
IA32_PL2_SSP =                                                 0x000006A6 //col:4396
IA32_PL3_SSP =                                                 0x000006A7 //col:4397
IA32_INTERRUPT_SSP_TABLE_ADDR =                                0x000006A8 //col:4398
IA32_TSC_DEADLINE =                                            0x000006E0 //col:4399
IA32_PM_ENABLE =                                               0x00000770 //col:4400
IA32_PM_ENABLE_HWP_ENABLE =                                    0x01 //col:4404
IA32_HWP_CAPABILITIES =                                        0x00000771 //col:4411
IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE =                    0xFF //col:4415
IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE =                 0xFF00 //col:4417
IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE =             0xFF0000 //col:4419
IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE =                     0xFF000000 //col:4421
IA32_HWP_REQUEST_PKG =                                         0x00000772 //col:4428
IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE =                     0xFF //col:4432
IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE =                     0xFF00 //col:4434
IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE =                     0xFF0000 //col:4436
IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE =           0xFF000000 //col:4438
IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW =                         0x3FF00000000 //col:4440
IA32_HWP_INTERRUPT =                                           0x00000773 //col:4447
IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE =          0x01 //col:4451
IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM =                      0x02 //col:4453
IA32_HWP_REQUEST =                                             0x00000774 //col:4460
IA32_HWP_REQUEST_MINIMUM_PERFORMANCE =                         0xFF //col:4464
IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE =                         0xFF00 //col:4466
IA32_HWP_REQUEST_DESIRED_PERFORMANCE =                         0xFF0000 //col:4468
IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE =               0xFF000000 //col:4470
IA32_HWP_REQUEST_ACTIVITY_WINDOW =                             0x3FF00000000 //col:4472
IA32_HWP_REQUEST_PACKAGE_CONTROL =                             0x40000000000 //col:4474
IA32_HWP_STATUS =                                              0x00000777 //col:4481
IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE =                0x01 //col:4485
IA32_HWP_STATUS_EXCURSION_TO_MINIMUM =                         0x04 //col:4488
IA32_X2APIC_APICID =                                           0x00000802 //col:4495
IA32_X2APIC_VERSION =                                          0x00000803 //col:4496
IA32_X2APIC_TPR =                                              0x00000808 //col:4497
IA32_X2APIC_PPR =                                              0x0000080A //col:4498
IA32_X2APIC_EOI =                                              0x0000080B //col:4499
IA32_X2APIC_LDR =                                              0x0000080D //col:4500
IA32_X2APIC_SIVR =                                             0x0000080F //col:4501
IA32_X2APIC_ISR0 =                                             0x00000810 //col:4507
IA32_X2APIC_ISR1 =                                             0x00000811 //col:4508
IA32_X2APIC_ISR2 =                                             0x00000812 //col:4509
IA32_X2APIC_ISR3 =                                             0x00000813 //col:4510
IA32_X2APIC_ISR4 =                                             0x00000814 //col:4511
IA32_X2APIC_ISR5 =                                             0x00000815 //col:4512
IA32_X2APIC_ISR6 =                                             0x00000816 //col:4513
IA32_X2APIC_ISR7 =                                             0x00000817 //col:4514
IA32_X2APIC_TMR0 =                                             0x00000818 //col:4524
IA32_X2APIC_TMR1 =                                             0x00000819 //col:4525
IA32_X2APIC_TMR2 =                                             0x0000081A //col:4526
IA32_X2APIC_TMR3 =                                             0x0000081B //col:4527
IA32_X2APIC_TMR4 =                                             0x0000081C //col:4528
IA32_X2APIC_TMR5 =                                             0x0000081D //col:4529
IA32_X2APIC_TMR6 =                                             0x0000081E //col:4530
IA32_X2APIC_TMR7 =                                             0x0000081F //col:4531
IA32_X2APIC_IRR0 =                                             0x00000820 //col:4541
IA32_X2APIC_IRR1 =                                             0x00000821 //col:4542
IA32_X2APIC_IRR2 =                                             0x00000822 //col:4543
IA32_X2APIC_IRR3 =                                             0x00000823 //col:4544
IA32_X2APIC_IRR4 =                                             0x00000824 //col:4545
IA32_X2APIC_IRR5 =                                             0x00000825 //col:4546
IA32_X2APIC_IRR6 =                                             0x00000826 //col:4547
IA32_X2APIC_IRR7 =                                             0x00000827 //col:4548
IA32_X2APIC_ESR =                                              0x00000828 //col:4553
IA32_X2APIC_LVT_CMCI =                                         0x0000082F //col:4554
IA32_X2APIC_ICR =                                              0x00000830 //col:4555
IA32_X2APIC_LVT_TIMER =                                        0x00000832 //col:4556
IA32_X2APIC_LVT_THERMAL =                                      0x00000833 //col:4557
IA32_X2APIC_LVT_PMI =                                          0x00000834 //col:4558
IA32_X2APIC_LVT_LINT0 =                                        0x00000835 //col:4559
IA32_X2APIC_LVT_LINT1 =                                        0x00000836 //col:4560
IA32_X2APIC_LVT_ERROR =                                        0x00000837 //col:4561
IA32_X2APIC_INIT_COUNT =                                       0x00000838 //col:4562
IA32_X2APIC_CUR_COUNT =                                        0x00000839 //col:4563
IA32_X2APIC_DIV_CONF =                                         0x0000083E //col:4564
IA32_X2APIC_SELF_IPI =                                         0x0000083F //col:4565
IA32_DEBUG_INTERFACE =                                         0x00000C80 //col:4566
IA32_DEBUG_INTERFACE_ENABLE =                                  0x01 //col:4570
IA32_DEBUG_INTERFACE_LOCK =                                    0x40000000 //col:4573
IA32_DEBUG_INTERFACE_DEBUG_OCCURRED =                          0x80000000 //col:4575
IA32_L3_QOS_CFG =                                              0x00000C81 //col:4582
IA32_L3_QOS_CFG_ENABLE =                                       0x01 //col:4586
IA32_L2_QOS_CFG =                                              0x00000C82 //col:4593
IA32_L2_QOS_CFG_ENABLE =                                       0x01 //col:4597
IA32_QM_EVTSEL =                                               0x00000C8D //col:4604
IA32_QM_EVTSEL_EVENT_ID =                                      0xFF //col:4608
IA32_QM_EVTSEL_RESOURCE_MONITORING_ID =                        0xFFFFFFFF00000000 //col:4611
IA32_QM_CTR =                                                  0x00000C8E //col:4617
IA32_QM_CTR_RESOURCE_MONITORED_DATA =                          0x3FFFFFFFFFFFFFFF //col:4621
IA32_QM_CTR_UNAVAILABLE =                                      0x4000000000000000 //col:4623
IA32_QM_CTR_ERROR =                                            0x8000000000000000 //col:4625
IA32_PQR_ASSOC =                                               0x00000C8F //col:4631
IA32_PQR_ASSOC_RESOURCE_MONITORING_ID =                        0xFFFFFFFF //col:4635
IA32_PQR_ASSOC_COS =                                           0xFFFFFFFF00000000 //col:4637
IA32_BNDCFGS =                                                 0x00000D90 //col:4643
IA32_BNDCFGS_ENABLE =                                          0x01 //col:4647
IA32_BNDCFGS_BND_PRESERVE =                                    0x02 //col:4649
IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS =                    0xFFFFFFFFFFFFF000 //col:4652
IA32_XSS =                                                     0x00000DA0 //col:4658
IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE =                    0x100 //col:4663
IA32_PKG_HDC_CTL =                                             0x00000DB0 //col:4670
IA32_PKG_HDC_CTL_HDC_PKG_ENABLE =                              0x01 //col:4674
IA32_PM_CTL1 =                                                 0x00000DB1 //col:4681
IA32_PM_CTL1_HDC_ALLOW_BLOCK =                                 0x01 //col:4685
IA32_THREAD_STALL =                                            0x00000DB2 //col:4692
IA32_EFER =                                                    0xC0000080 //col:4697
IA32_EFER_SYSCALL_ENABLE =                                     0x01 //col:4701
IA32_EFER_IA32E_MODE_ENABLE =                                  0x100 //col:4704
IA32_EFER_IA32E_MODE_ACTIVE =                                  0x400 //col:4707
IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE =                         0x800 //col:4709
IA32_STAR =                                                    0xC0000081 //col:4716
IA32_LSTAR =                                                   0xC0000082 //col:4717
IA32_CSTAR =                                                   0xC0000083 //col:4718
IA32_FMASK =                                                   0xC0000084 //col:4719
IA32_FS_BASE =                                                 0xC0000100 //col:4720
IA32_GS_BASE =                                                 0xC0000101 //col:4721
IA32_KERNEL_GS_BASE =                                          0xC0000102 //col:4722
IA32_TSC_AUX =                                                 0xC0000103 //col:4723
IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE =                         0xFFFFFFFF //col:4727
PDE_4MB_32_PRESENT =                                           0x01 //col:4751
PDE_4MB_32_WRITE =                                             0x02 //col:4753
PDE_4MB_32_SUPERVISOR =                                        0x04 //col:4755
PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH =                          0x08 //col:4757
PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE =                          0x10 //col:4759
PDE_4MB_32_ACCESSED =                                          0x20 //col:4761
PDE_4MB_32_DIRTY =                                             0x40 //col:4763
PDE_4MB_32_LARGE_PAGE =                                        0x80 //col:4765
PDE_4MB_32_GLOBAL =                                            0x100 //col:4767
PDE_4MB_32_IGNORED_1 =                                         0xE00 //col:4769
PDE_4MB_32_PAT =                                               0x1000 //col:4771
PDE_4MB_32_PAGE_FRAME_NUMBER_LOW =                             0x1FE000 //col:4773
PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH =                            0xFFC00000 //col:4776
PDE_32_PRESENT =                                               0x01 //col:4785
PDE_32_WRITE =                                                 0x02 //col:4787
PDE_32_SUPERVISOR =                                            0x04 //col:4789
PDE_32_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:4791
PDE_32_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:4793
PDE_32_ACCESSED =                                              0x20 //col:4795
PDE_32_IGNORED_1 =                                             0x40 //col:4797
PDE_32_LARGE_PAGE =                                            0x80 //col:4799
PDE_32_IGNORED_2 =                                             0xF00 //col:4801
PDE_32_PAGE_FRAME_NUMBER =                                     0xFFFFF000 //col:4803
PTE_32_PRESENT =                                               0x01 //col:4812
PTE_32_WRITE =                                                 0x02 //col:4814
PTE_32_SUPERVISOR =                                            0x04 //col:4816
PTE_32_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:4818
PTE_32_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:4820
PTE_32_ACCESSED =                                              0x20 //col:4822
PTE_32_DIRTY =                                                 0x40 //col:4824
PTE_32_PAT =                                                   0x80 //col:4826
PTE_32_GLOBAL =                                                0x100 //col:4828
PTE_32_IGNORED_1 =                                             0xE00 //col:4830
PTE_32_PAGE_FRAME_NUMBER =                                     0xFFFFF000 //col:4832
PT_ENTRY_32_PRESENT =                                          0x01 //col:4841
PT_ENTRY_32_WRITE =                                            0x02 //col:4843
PT_ENTRY_32_SUPERVISOR =                                       0x04 //col:4845
PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH =                         0x08 //col:4847
PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE =                         0x10 //col:4849
PT_ENTRY_32_ACCESSED =                                         0x20 //col:4851
PT_ENTRY_32_DIRTY =                                            0x40 //col:4853
PT_ENTRY_32_LARGE_PAGE =                                       0x80 //col:4855
PT_ENTRY_32_GLOBAL =                                           0x100 //col:4857
PT_ENTRY_32_IGNORED_1 =                                        0xE00 //col:4859
PT_ENTRY_32_PAGE_FRAME_NUMBER =                                0xFFFFF000 //col:4861
PDE_ENTRY_COUNT_32 =                                           0x00000400 //col:4872
PTE_ENTRY_COUNT_32 =                                           0x00000400 //col:4873
PML4E_64_PRESENT =                                             0x01 //col:4890
PML4E_64_WRITE =                                               0x02 //col:4892
PML4E_64_SUPERVISOR =                                          0x04 //col:4894
PML4E_64_PAGE_LEVEL_WRITE_THROUGH =                            0x08 //col:4896
PML4E_64_PAGE_LEVEL_CACHE_DISABLE =                            0x10 //col:4898
PML4E_64_ACCESSED =                                            0x20 //col:4900
PML4E_64_MUST_BE_ZERO =                                        0x80 //col:4903
PML4E_64_IGNORED_1 =                                           0x700 //col:4905
PML4E_64_RESTART =                                             0x800 //col:4907
PML4E_64_PAGE_FRAME_NUMBER =                                   0xFFFFFFFFF000 //col:4909
PML4E_64_IGNORED_2 =                                           0x7FF0000000000000 //col:4912
PML4E_64_EXECUTE_DISABLE =                                     0x8000000000000000 //col:4914
PDPTE_1GB_64_PRESENT =                                         0x01 //col:4923
PDPTE_1GB_64_WRITE =                                           0x02 //col:4925
PDPTE_1GB_64_SUPERVISOR =                                      0x04 //col:4927
PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH =                        0x08 //col:4929
PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE =                        0x10 //col:4931
PDPTE_1GB_64_ACCESSED =                                        0x20 //col:4933
PDPTE_1GB_64_DIRTY =                                           0x40 //col:4935
PDPTE_1GB_64_LARGE_PAGE =                                      0x80 //col:4937
PDPTE_1GB_64_GLOBAL =                                          0x100 //col:4939
PDPTE_1GB_64_IGNORED_1 =                                       0x600 //col:4941
PDPTE_1GB_64_RESTART =                                         0x800 //col:4943
PDPTE_1GB_64_PAT =                                             0x1000 //col:4945
PDPTE_1GB_64_PAGE_FRAME_NUMBER =                               0xFFFFC0000000 //col:4948
PDPTE_1GB_64_IGNORED_2 =                                       0x7F0000000000000 //col:4951
PDPTE_1GB_64_PROTECTION_KEY =                                  0x7800000000000000 //col:4953
PDPTE_1GB_64_EXECUTE_DISABLE =                                 0x8000000000000000 //col:4955
PDPTE_64_PRESENT =                                             0x01 //col:4964
PDPTE_64_WRITE =                                               0x02 //col:4966
PDPTE_64_SUPERVISOR =                                          0x04 //col:4968
PDPTE_64_PAGE_LEVEL_WRITE_THROUGH =                            0x08 //col:4970
PDPTE_64_PAGE_LEVEL_CACHE_DISABLE =                            0x10 //col:4972
PDPTE_64_ACCESSED =                                            0x20 //col:4974
PDPTE_64_LARGE_PAGE =                                          0x80 //col:4977
PDPTE_64_IGNORED_1 =                                           0x700 //col:4979
PDPTE_64_RESTART =                                             0x800 //col:4981
PDPTE_64_PAGE_FRAME_NUMBER =                                   0xFFFFFFFFF000 //col:4983
PDPTE_64_IGNORED_2 =                                           0x7FF0000000000000 //col:4986
PDPTE_64_EXECUTE_DISABLE =                                     0x8000000000000000 //col:4988
PDE_2MB_64_PRESENT =                                           0x01 //col:4997
PDE_2MB_64_WRITE =                                             0x02 //col:4999
PDE_2MB_64_SUPERVISOR =                                        0x04 //col:5001
PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH =                          0x08 //col:5003
PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE =                          0x10 //col:5005
PDE_2MB_64_ACCESSED =                                          0x20 //col:5007
PDE_2MB_64_DIRTY =                                             0x40 //col:5009
PDE_2MB_64_LARGE_PAGE =                                        0x80 //col:5011
PDE_2MB_64_GLOBAL =                                            0x100 //col:5013
PDE_2MB_64_IGNORED_1 =                                         0x600 //col:5015
PDE_2MB_64_RESTART =                                           0x800 //col:5017
PDE_2MB_64_PAT =                                               0x1000 //col:5019
PDE_2MB_64_PAGE_FRAME_NUMBER =                                 0xFFFFFFE00000 //col:5022
PDE_2MB_64_IGNORED_2 =                                         0x7F0000000000000 //col:5025
PDE_2MB_64_PROTECTION_KEY =                                    0x7800000000000000 //col:5027
PDE_2MB_64_EXECUTE_DISABLE =                                   0x8000000000000000 //col:5029
PDE_64_PRESENT =                                               0x01 //col:5038
PDE_64_WRITE =                                                 0x02 //col:5040
PDE_64_SUPERVISOR =                                            0x04 //col:5042
PDE_64_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:5044
PDE_64_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:5046
PDE_64_ACCESSED =                                              0x20 //col:5048
PDE_64_LARGE_PAGE =                                            0x80 //col:5051
PDE_64_IGNORED_1 =                                             0x700 //col:5053
PDE_64_RESTART =                                               0x800 //col:5055
PDE_64_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:5057
PDE_64_IGNORED_2 =                                             0x7FF0000000000000 //col:5060
PDE_64_EXECUTE_DISABLE =                                       0x8000000000000000 //col:5062
PTE_64_PRESENT =                                               0x01 //col:5071
PTE_64_WRITE =                                                 0x02 //col:5073
PTE_64_SUPERVISOR =                                            0x04 //col:5075
PTE_64_PAGE_LEVEL_WRITE_THROUGH =                              0x08 //col:5077
PTE_64_PAGE_LEVEL_CACHE_DISABLE =                              0x10 //col:5079
PTE_64_ACCESSED =                                              0x20 //col:5081
PTE_64_DIRTY =                                                 0x40 //col:5083
PTE_64_PAT =                                                   0x80 //col:5085
PTE_64_GLOBAL =                                                0x100 //col:5087
PTE_64_IGNORED_1 =                                             0x600 //col:5089
PTE_64_RESTART =                                               0x800 //col:5091
PTE_64_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:5093
PTE_64_IGNORED_2 =                                             0x7F0000000000000 //col:5096
PTE_64_PROTECTION_KEY =                                        0x7800000000000000 //col:5098
PTE_64_EXECUTE_DISABLE =                                       0x8000000000000000 //col:5100
PT_ENTRY_64_PRESENT =                                          0x01 //col:5109
PT_ENTRY_64_WRITE =                                            0x02 //col:5111
PT_ENTRY_64_SUPERVISOR =                                       0x04 //col:5113
PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH =                         0x08 //col:5115
PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE =                         0x10 //col:5117
PT_ENTRY_64_ACCESSED =                                         0x20 //col:5119
PT_ENTRY_64_DIRTY =                                            0x40 //col:5121
PT_ENTRY_64_LARGE_PAGE =                                       0x80 //col:5123
PT_ENTRY_64_GLOBAL =                                           0x100 //col:5125
PT_ENTRY_64_IGNORED_1 =                                        0x600 //col:5127
PT_ENTRY_64_RESTART =                                          0x800 //col:5129
PT_ENTRY_64_PAGE_FRAME_NUMBER =                                0xFFFFFFFFF000 //col:5131
PT_ENTRY_64_IGNORED_2 =                                        0x7F0000000000000 //col:5134
PT_ENTRY_64_PROTECTION_KEY =                                   0x7800000000000000 //col:5136
PT_ENTRY_64_EXECUTE_DISABLE =                                  0x8000000000000000 //col:5138
PML4_ENTRY_COUNT_64 =                                          0x00000200 //col:5149
PDPTE_ENTRY_COUNT_64 =                                         0x00000200 //col:5150
PDE_ENTRY_COUNT_64 =                                           0x00000200 //col:5151
PTE_ENTRY_COUNT_64 =                                           0x00000200 //col:5152
INVPCID_INDIVIDUAL_ADDRESS =                                   0x00000000 //col:5165
INVPCID_SINGLE_CONTEXT =                                       0x00000001 //col:5166
INVPCID_ALL_CONTEXT_WITH_GLOBALS =                             0x00000002 //col:5167
INVPCID_ALL_CONTEXT =                                          0x00000003 //col:5168
INVPCID_DESCRIPTOR_PCID =                                      0xFFF //col:5176
INVPCID_DESCRIPTOR_RESERVED1 =                                 0xFFFFFFFFFFFFF000 //col:5178
INVPCID_DESCRIPTOR_LINEAR_ADDRESS =                            0xFFFFFFFFFFFFFFFF0000000000000000 //col:5180
SEGMENT_ACCESS_RIGHTS_TYPE =                                   0xF00 //col:5209
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE =                        0x1000 //col:5211
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL =             0x6000 //col:5213
SEGMENT_ACCESS_RIGHTS_PRESENT =                                0x8000 //col:5215
SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT =                          0x100000 //col:5218
SEGMENT_ACCESS_RIGHTS_LONG_MODE =                              0x200000 //col:5220
SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG =                            0x400000 //col:5222
SEGMENT_ACCESS_RIGHTS_GRANULARITY =                            0x800000 //col:5224
SEGMENT__BASE_ADDRESS_MIDDLE =                                 0xFF //col:5237
SEGMENT__TYPE =                                                0xF00 //col:5239
SEGMENT__DESCRIPTOR_TYPE =                                     0x1000 //col:5241
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL =                          0x6000 //col:5243
SEGMENT__PRESENT =                                             0x8000 //col:5245
SEGMENT__SEGMENT_LIMIT_HIGH =                                  0xF0000 //col:5247
SEGMENT__AVAILABLE_BIT =                                       0x100000 //col:5249
SEGMENT__LONG_MODE =                                           0x200000 //col:5251
SEGMENT__DEFAULT_BIG =                                         0x400000 //col:5253
SEGMENT__GRANULARITY =                                         0x800000 //col:5255
SEGMENT__BASE_ADDRESS_HIGH =                                   0xFF000000 //col:5257
SEGMENT__BASE_ADDRESS_MIDDLE =                                 0xFF //col:5271
SEGMENT__TYPE =                                                0xF00 //col:5273
SEGMENT__DESCRIPTOR_TYPE =                                     0x1000 //col:5275
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL =                          0x6000 //col:5277
SEGMENT__PRESENT =                                             0x8000 //col:5279
SEGMENT__SEGMENT_LIMIT_HIGH =                                  0xF0000 //col:5281
SEGMENT__AVAILABLE_BIT =                                       0x100000 //col:5283
SEGMENT__LONG_MODE =                                           0x200000 //col:5285
SEGMENT__DEFAULT_BIG =                                         0x400000 //col:5287
SEGMENT__GRANULARITY =                                         0x800000 //col:5289
SEGMENT__BASE_ADDRESS_HIGH =                                   0xFF000000 //col:5291
SEGMENT__INTERRUPT_STACK_TABLE =                               0x07 //col:5307
SEGMENT__MUST_BE_ZERO_0 =                                      0xF8 //col:5309
SEGMENT__TYPE =                                                0xF00 //col:5311
SEGMENT__MUST_BE_ZERO_1 =                                      0x1000 //col:5313
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL =                          0x6000 //col:5315
SEGMENT__PRESENT =                                             0x8000 //col:5317
SEGMENT__OFFSET_MIDDLE =                                       0xFFFF0000 //col:5319
SEGMENT_DESCRIPTOR_TYPE_SYSTEM =                               0x00000000 //col:5329
SEGMENT_DESCRIPTOR_TYPE_CODE_OR_DATA =                         0x00000001 //col:5330
SEGMENT_DESCRIPTOR_TYPE_DATA_R =                               0x00000000 //col:5336
SEGMENT_DESCRIPTOR_TYPE_DATA_RA =                              0x00000001 //col:5337
SEGMENT_DESCRIPTOR_TYPE_DATA_RW =                              0x00000002 //col:5338
SEGMENT_DESCRIPTOR_TYPE_DATA_RWA =                             0x00000003 //col:5339
SEGMENT_DESCRIPTOR_TYPE_DATA_RE =                              0x00000004 //col:5340
SEGMENT_DESCRIPTOR_TYPE_DATA_REA =                             0x00000005 //col:5341
SEGMENT_DESCRIPTOR_TYPE_DATA_RWE =                             0x00000006 //col:5342
SEGMENT_DESCRIPTOR_TYPE_DATA_RWEA =                            0x00000007 //col:5343
SEGMENT_DESCRIPTOR_TYPE_CODE_E =                               0x00000008 //col:5344
SEGMENT_DESCRIPTOR_TYPE_CODE_EA =                              0x00000009 //col:5345
SEGMENT_DESCRIPTOR_TYPE_CODE_ER =                              0x0000000A //col:5346
SEGMENT_DESCRIPTOR_TYPE_CODE_ERA =                             0x0000000B //col:5347
SEGMENT_DESCRIPTOR_TYPE_CODE_EC =                              0x0000000C //col:5348
SEGMENT_DESCRIPTOR_TYPE_CODE_ECA =                             0x0000000D //col:5349
SEGMENT_DESCRIPTOR_TYPE_CODE_ERC =                             0x0000000E //col:5350
SEGMENT_DESCRIPTOR_TYPE_CODE_ERCA =                            0x0000000F //col:5351
SEGMENT_DESCRIPTOR_TYPE_RESERVED_1 =                           0x00000000 //col:5361
SEGMENT_DESCRIPTOR_TYPE_TSS_16_AVAILABLE =                     0x00000001 //col:5362
SEGMENT_DESCRIPTOR_TYPE_LDT =                                  0x00000002 //col:5363
SEGMENT_DESCRIPTOR_TYPE_TSS_16_BUSY =                          0x00000003 //col:5364
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE_16 =                         0x00000004 //col:5365
SEGMENT_DESCRIPTOR_TYPE_TASK_GATE =                            0x00000005 //col:5366
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE_16 =                    0x00000006 //col:5367
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE_16 =                         0x00000007 //col:5368
SEGMENT_DESCRIPTOR_TYPE_RESERVED_2 =                           0x00000008 //col:5369
SEGMENT_DESCRIPTOR_TYPE_TSS_AVAILABLE =                        0x00000009 //col:5370
SEGMENT_DESCRIPTOR_TYPE_RESERVED_3 =                           0x0000000A //col:5371
SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY =                             0x0000000B //col:5372
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE =                            0x0000000C //col:5373
SEGMENT_DESCRIPTOR_TYPE_RESERVED_4 =                           0x0000000D //col:5374
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE =                       0x0000000E //col:5375
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE =                            0x0000000F //col:5376
SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL =                     0x03 //col:5384
SEGMENT_SELECTOR_TABLE_INDICATOR =                             0x04 //col:5386
SEGMENT_SELECTOR_INDEX =                                       0xFFF8 //col:5388
VMX_EXIT_REASON_XCPT_OR_NMI =                                  0x00000000 //col:5431
VMX_EXIT_REASON_EXT_INT =                                      0x00000001 //col:5432
VMX_EXIT_REASON_TRIPLE_FAULT =                                 0x00000002 //col:5433
VMX_EXIT_REASON_INIT_SIGNAL =                                  0x00000003 //col:5434
VMX_EXIT_REASON_SIPI =                                         0x00000004 //col:5435
VMX_EXIT_REASON_IO_SMI =                                       0x00000005 //col:5436
VMX_EXIT_REASON_SMI =                                          0x00000006 //col:5437
VMX_EXIT_REASON_INT_WINDOW =                                   0x00000007 //col:5438
VMX_EXIT_REASON_NMI_WINDOW =                                   0x00000008 //col:5439
VMX_EXIT_REASON_TASK_SWITCH =                                  0x00000009 //col:5440
VMX_EXIT_REASON_CPUID =                                        0x0000000A //col:5441
VMX_EXIT_REASON_GETSEC =                                       0x0000000B //col:5442
VMX_EXIT_REASON_HLT =                                          0x0000000C //col:5443
VMX_EXIT_REASON_INVD =                                         0x0000000D //col:5444
VMX_EXIT_REASON_INVLPG =                                       0x0000000E //col:5445
VMX_EXIT_REASON_RDPMC =                                        0x0000000F //col:5446
VMX_EXIT_REASON_RDTSC =                                        0x00000010 //col:5447
VMX_EXIT_REASON_RSM =                                          0x00000011 //col:5448
VMX_EXIT_REASON_VMCALL =                                       0x00000012 //col:5449
VMX_EXIT_REASON_VMCLEAR =                                      0x00000013 //col:5450
VMX_EXIT_REASON_VMLAUNCH =                                     0x00000014 //col:5451
VMX_EXIT_REASON_VMPTRLD =                                      0x00000015 //col:5452
VMX_EXIT_REASON_VMPTRST =                                      0x00000016 //col:5453
VMX_EXIT_REASON_VMREAD =                                       0x00000017 //col:5454
VMX_EXIT_REASON_VMRESUME =                                     0x00000018 //col:5455
VMX_EXIT_REASON_VMWRITE =                                      0x00000019 //col:5456
VMX_EXIT_REASON_VMXOFF =                                       0x0000001A //col:5457
VMX_EXIT_REASON_VMXON =                                        0x0000001B //col:5458
VMX_EXIT_REASON_MOV_CRX =                                      0x0000001C //col:5459
VMX_EXIT_REASON_MOV_DRX =                                      0x0000001D //col:5460
VMX_EXIT_REASON_IO_INSTR =                                     0x0000001E //col:5461
VMX_EXIT_REASON_RDMSR =                                        0x0000001F //col:5462
VMX_EXIT_REASON_WRMSR =                                        0x00000020 //col:5463
VMX_EXIT_REASON_ERR_INVALID_GUEST_STATE =                      0x00000021 //col:5464
VMX_EXIT_REASON_ERR_MSR_LOAD =                                 0x00000022 //col:5465
VMX_EXIT_REASON_MWAIT =                                        0x00000024 //col:5466
VMX_EXIT_REASON_MTF =                                          0x00000025 //col:5467
VMX_EXIT_REASON_MONITOR =                                      0x00000027 //col:5468
VMX_EXIT_REASON_PAUSE =                                        0x00000028 //col:5469
VMX_EXIT_REASON_ERR_MACHINE_CHECK =                            0x00000029 //col:5470
VMX_EXIT_REASON_TPR_BELOW_THRESHOLD =                          0x0000002B //col:5471
VMX_EXIT_REASON_APIC_ACCESS =                                  0x0000002C //col:5472
VMX_EXIT_REASON_VIRTUALIZED_EOI =                              0x0000002D //col:5473
VMX_EXIT_REASON_XDTR_ACCESS =                                  0x0000002E //col:5474
VMX_EXIT_REASON_TR_ACCESS =                                    0x0000002F //col:5475
VMX_EXIT_REASON_EPT_VIOLATION =                                0x00000030 //col:5476
VMX_EXIT_REASON_EPT_MISCONFIG =                                0x00000031 //col:5477
VMX_EXIT_REASON_INVEPT =                                       0x00000032 //col:5478
VMX_EXIT_REASON_RDTSCP =                                       0x00000033 //col:5479
VMX_EXIT_REASON_PREEMPT_TIMER =                                0x00000034 //col:5480
VMX_EXIT_REASON_INVVPID =                                      0x00000035 //col:5481
VMX_EXIT_REASON_WBINVD =                                       0x00000036 //col:5482
VMX_EXIT_REASON_XSETBV =                                       0x00000037 //col:5483
VMX_EXIT_REASON_APIC_WRITE =                                   0x00000038 //col:5484
VMX_EXIT_REASON_RDRAND =                                       0x00000039 //col:5485
VMX_EXIT_REASON_INVPCID =                                      0x0000003A //col:5486
VMX_EXIT_REASON_VMFUNC =                                       0x0000003B //col:5487
VMX_EXIT_REASON_ENCLS =                                        0x0000003C //col:5488
VMX_EXIT_REASON_RDSEED =                                       0x0000003D //col:5489
VMX_EXIT_REASON_PML_FULL =                                     0x0000003E //col:5490
VMX_EXIT_REASON_XSAVES =                                       0x0000003F //col:5491
VMX_EXIT_REASON_XRSTORS =                                      0x00000040 //col:5492
VMX_ERROR_VMCALL =                                             0x00000001 //col:5502
VMX_ERROR_VMCLEAR_INVALID_PHYS_ADDR =                          0x00000002 //col:5503
VMX_ERROR_VMCLEAR_INVALID_VMXON_PTR =                          0x00000003 //col:5504
VMX_ERROR_VMLAUCH_NON_CLEAR_VMCS =                             0x00000004 //col:5505
VMX_ERROR_VMRESUME_NON_LAUNCHED_VMCS =                         0x00000005 //col:5506
VMX_ERROR_VMRESUME_CORRUPTED_VMCS =                            0x00000006 //col:5507
VMX_ERROR_VMENTRY_INVALID_CONTROL_FIELDS =                     0x00000007 //col:5508
VMX_ERROR_VMENTRY_INVALID_HOST_STATE =                         0x00000008 //col:5509
VMX_ERROR_VMPTRLD_INVALID_PHYS_ADDR =                          0x00000009 //col:5510
VMX_ERROR_VMPTRLD_VMXON_PTR =                                  0x0000000A //col:5511
VMX_ERROR_VMPTRLD_WRONG_VMCS_REVISION =                        0x0000000B //col:5512
VMX_ERROR_VMREAD_VMWRITE_INVALID_COMPONENT =                   0x0000000C //col:5513
VMX_ERROR_VMWRITE_READONLY_COMPONENT =                         0x0000000D //col:5514
VMX_ERROR_VMXON_IN_VMX_ROOT_OP =                               0x0000000F //col:5515
VMX_ERROR_VMENTRY_INVALID_VMCS_EXEC_PTR =                      0x00000010 //col:5516
VMX_ERROR_VMENTRY_NON_LAUNCHED_EXEC_VMCS =                     0x00000011 //col:5517
VMX_ERROR_VMENTRY_EXEC_VMCS_PTR =                              0x00000012 //col:5518
VMX_ERROR_VMCALL_NON_CLEAR_VMCS =                              0x00000013 //col:5519
VMX_ERROR_VMCALL_INVALID_VMEXIT_FIELDS =                       0x00000014 //col:5520
VMX_ERROR_VMCALL_INVALID_MSEG_REVISION =                       0x00000016 //col:5521
VMX_ERROR_VMXOFF_DUAL_MONITOR =                                0x00000017 //col:5522
VMX_ERROR_VMCALL_INVALID_SMM_MONITOR =                         0x00000018 //col:5523
VMX_ERROR_VMENTRY_INVALID_VM_EXEC_CTRL =                       0x00000019 //col:5524
VMX_ERROR_VMENTRY_MOV_SS =                                     0x0000001A //col:5525
VMX_ERROR_INVEPTVPID_INVALID_OPERAND =                         0x0000001C //col:5526
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION =  0x0F //col:5557
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED = 0x2000 //col:5560
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION =    0x4000 //col:5562
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR =                  0xFFFF //col:5572
VMX_EXIT_QUALIFICATION_TASK_SWITCH_TYPE =                      0xC0000000 //col:5575
VMX_EXIT_QUALIFICATION_TYPE_CALL =                             0x00000000 //col:5576
VMX_EXIT_QUALIFICATION_TYPE_IRET =                             0x00000001 //col:5577
VMX_EXIT_QUALIFICATION_TYPE_JMP =                              0x00000002 //col:5578
VMX_EXIT_QUALIFICATION_TYPE_IDT =                              0x00000003 //col:5579
VMX_EXIT_QUALIFICATION_CR_ACCESS_CR_NUMBER =                   0x0F //col:5589
VMX_EXIT_QUALIFICATION_REGISTER_CR0 =                          0x00000000 //col:5590
VMX_EXIT_QUALIFICATION_REGISTER_CR2 =                          0x00000002 //col:5591
VMX_EXIT_QUALIFICATION_REGISTER_CR3 =                          0x00000003 //col:5592
VMX_EXIT_QUALIFICATION_REGISTER_CR4 =                          0x00000004 //col:5593
VMX_EXIT_QUALIFICATION_REGISTER_CR8 =                          0x00000008 //col:5594
VMX_EXIT_QUALIFICATION_CR_ACCESS_ACCESS_TYPE =                 0x30 //col:5596
VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR =                      0x00000000 //col:5597
VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR =                    0x00000001 //col:5598
VMX_EXIT_QUALIFICATION_ACCESS_CLTS =                           0x00000002 //col:5599
VMX_EXIT_QUALIFICATION_ACCESS_LMSW =                           0x00000003 //col:5600
VMX_EXIT_QUALIFICATION_CR_ACCESS_LMSW_OPERAND_TYPE =           0x40 //col:5602
VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER =                      0x00000000 //col:5603
VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY =                        0x00000001 //col:5604
VMX_EXIT_QUALIFICATION_CR_ACCESS_GP_REGISTER =                 0xF00 //col:5607
VMX_EXIT_QUALIFICATION_GENREG_RAX =                            0x00000000 //col:5608
VMX_EXIT_QUALIFICATION_GENREG_RCX =                            0x00000001 //col:5609
VMX_EXIT_QUALIFICATION_GENREG_RDX =                            0x00000002 //col:5610
VMX_EXIT_QUALIFICATION_GENREG_RBX =                            0x00000003 //col:5611
VMX_EXIT_QUALIFICATION_GENREG_RSP =                            0x00000004 //col:5612
VMX_EXIT_QUALIFICATION_GENREG_RBP =                            0x00000005 //col:5613
VMX_EXIT_QUALIFICATION_GENREG_RSI =                            0x00000006 //col:5614
VMX_EXIT_QUALIFICATION_GENREG_RDI =                            0x00000007 //col:5615
VMX_EXIT_QUALIFICATION_GENREG_R8 =                             0x00000008 //col:5616
VMX_EXIT_QUALIFICATION_GENREG_R9 =                             0x00000009 //col:5617
VMX_EXIT_QUALIFICATION_GENREG_R10 =                            0x0000000A //col:5618
VMX_EXIT_QUALIFICATION_GENREG_R11 =                            0x0000000B //col:5619
VMX_EXIT_QUALIFICATION_GENREG_R12 =                            0x0000000C //col:5620
VMX_EXIT_QUALIFICATION_GENREG_R13 =                            0x0000000D //col:5621
VMX_EXIT_QUALIFICATION_GENREG_R14 =                            0x0000000E //col:5622
VMX_EXIT_QUALIFICATION_GENREG_R15 =                            0x0000000F //col:5623
VMX_EXIT_QUALIFICATION_CR_ACCESS_LMSW_SOURCE_DATA =            0xFFFF0000 //col:5626
VMX_EXIT_QUALIFICATION_DR_ACCESS_DR_NUMBER =                   0x07 //col:5636
VMX_EXIT_QUALIFICATION_REGISTER_DR0 =                          0x00000000 //col:5637
VMX_EXIT_QUALIFICATION_REGISTER_DR1 =                          0x00000001 //col:5638
VMX_EXIT_QUALIFICATION_REGISTER_DR2 =                          0x00000002 //col:5639
VMX_EXIT_QUALIFICATION_REGISTER_DR3 =                          0x00000003 //col:5640
VMX_EXIT_QUALIFICATION_REGISTER_DR6 =                          0x00000006 //col:5641
VMX_EXIT_QUALIFICATION_REGISTER_DR7 =                          0x00000007 //col:5642
VMX_EXIT_QUALIFICATION_DR_ACCESS_DIRECTION_OF_ACCESS =         0x10 //col:5645
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR =                   0x00000000 //col:5646
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR =                 0x00000001 //col:5647
VMX_EXIT_QUALIFICATION_DR_ACCESS_GP_REGISTER =                 0xF00 //col:5650
VMX_EXIT_QUALIFICATION_IO_INST_SIZE_OF_ACCESS =                0x07 //col:5660
VMX_EXIT_QUALIFICATION_WIDTH_1B =                              0x00000000 //col:5661
VMX_EXIT_QUALIFICATION_WIDTH_2B =                              0x00000001 //col:5662
VMX_EXIT_QUALIFICATION_WIDTH_4B =                              0x00000003 //col:5663
VMX_EXIT_QUALIFICATION_IO_INST_DIRECTION_OF_ACCESS =           0x08 //col:5665
VMX_EXIT_QUALIFICATION_DIRECTION_OUT =                         0x00000000 //col:5666
VMX_EXIT_QUALIFICATION_DIRECTION_IN =                          0x00000001 //col:5667
VMX_EXIT_QUALIFICATION_IO_INST_STRING_INSTRUCTION =            0x10 //col:5669
VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING =                  0x00000000 //col:5670
VMX_EXIT_QUALIFICATION_IS_STRING_STRING =                      0x00000001 //col:5671
VMX_EXIT_QUALIFICATION_IO_INST_REP_PREFIXED =                  0x20 //col:5673
VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP =                        0x00000000 //col:5674
VMX_EXIT_QUALIFICATION_IS_REP_REP =                            0x00000001 //col:5675
VMX_EXIT_QUALIFICATION_IO_INST_OPERAND_ENCODING =              0x40 //col:5677
VMX_EXIT_QUALIFICATION_ENCODING_DX =                           0x00000000 //col:5678
VMX_EXIT_QUALIFICATION_ENCODING_IMM =                          0x00000001 //col:5679
VMX_EXIT_QUALIFICATION_IO_INST_PORT_NUMBER =                   0xFFFF0000 //col:5682
VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET =               0xFFF //col:5692
VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE =               0xF000 //col:5694
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_READ =                      0x00000000 //col:5695
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_WRITE =                     0x00000001 //col:5696
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_INSTR_FETCH =               0x00000002 //col:5697
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_EVENT_DELIVERY =            0x00000003 //col:5698
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_EVENT_DELIVERY =          0x0000000A //col:5699
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_INSTR =                   0x0000000F //col:5700
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_DATA_READ =               0x01 //col:5710
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_DATA_WRITE =              0x02 //col:5712
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_INSTRUCTION_FETCH =       0x04 //col:5714
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_PRESENT =           0x08 //col:5716
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_WRITE =             0x10 //col:5718
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_EXECUTE =           0x20 //col:5720
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ENTRY_EXECUTE_FOR_USER_MODE = 0x40 //col:5722
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS = 0x80 //col:5724
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_TRANSLATED_ACCESS =   0x100 //col:5726
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS = 0x200 //col:5728
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE =  0x400 //col:5730
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE =    0x800 //col:5732
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING =          0x1000 //col:5734
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS =     0x2000 //col:5736
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK = 0x4000 //col:5738
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION = 0x8000 //col:5740
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION = 0x10000 //col:5742
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE =            0x380 //col:5762
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER =        0x38000 //col:5765
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING =               0x03 //col:5775
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE =          0x380 //col:5778
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER =      0x38000 //col:5781
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GP_REGISTER =           0x3C0000 //col:5783
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GP_REGISTER_INVALID =   0x400000 //col:5785
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER =         0x7800000 //col:5787
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID = 0x8000000 //col:5789
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2 =            0xF0000000 //col:5791
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING =         0x03 //col:5801
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE =    0x380 //col:5804
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE =    0x800 //col:5807
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER = 0x38000 //col:5810
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GP_REGISTER =     0x3C0000 //col:5812
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GP_REGISTER_INVALID = 0x400000 //col:5814
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER =   0x7800000 //col:5816
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID = 0x8000000 //col:5818
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_IDENTITY = 0x30000000 //col:5820
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING =           0x03 //col:5830
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1 =             0x78 //col:5833
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE =      0x380 //col:5835
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER =   0x400 //col:5837
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER =  0x38000 //col:5840
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GP_REGISTER =       0x3C0000 //col:5842
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GP_REGISTER_INVALID = 0x400000 //col:5844
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER =     0x7800000 //col:5846
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID = 0x8000000 //col:5848
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_IDENTITY = 0x30000000 //col:5850
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER = 0x78 //col:5861
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE =       0x1800 //col:5864
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING =           0x03 //col:5874
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE =      0x380 //col:5877
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER =  0x38000 //col:5880
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GP_REGISTER =       0x3C0000 //col:5882
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GP_REGISTER_INVALID = 0x400000 //col:5884
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER =     0x7800000 //col:5886
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID = 0x8000000 //col:5888
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING =           0x03 //col:5898
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1 =        0x78 //col:5901
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE =      0x380 //col:5903
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER =   0x400 //col:5905
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER =  0x38000 //col:5908
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GP_REGISTER =       0x3C0000 //col:5910
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GP_REGISTER_INVALID = 0x400000 //col:5912
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER =     0x7800000 //col:5914
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID = 0x8000000 //col:5916
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2 =        0xF0000000 //col:5918
VMX_SEGMENT_ACCESS_RIGHTS_TYPE =                               0x0F //col:5932
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE =                    0x10 //col:5934
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL =         0x60 //col:5936
VMX_SEGMENT_ACCESS_RIGHTS_PRESENT =                            0x80 //col:5938
VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT =                      0x1000 //col:5941
VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE =                          0x2000 //col:5943
VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG =                        0x4000 //col:5945
VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY =                        0x8000 //col:5947
VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE =                           0x10000 //col:5949
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI =                   0x01 //col:5959
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS =                0x02 //col:5961
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI =                   0x04 //col:5963
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI =                   0x08 //col:5965
VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION =              0x10 //col:5967
VMX_ACTIVE =                                                   0x00000000 //col:5974
VMX_HLT =                                                      0x00000001 //col:5975
VMX_SHUTDOWN =                                                 0x00000002 //col:5976
VMX_WAIT_FOR_SIPI =                                            0x00000003 //col:5977
VMX_PENDING_DEBUG_EXCEPTIONS_B0 =                              0x01 //col:5985
VMX_PENDING_DEBUG_EXCEPTIONS_B1 =                              0x02 //col:5987
VMX_PENDING_DEBUG_EXCEPTIONS_B2 =                              0x04 //col:5989
VMX_PENDING_DEBUG_EXCEPTIONS_B3 =                              0x08 //col:5991
VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT =              0x1000 //col:5994
VMX_PENDING_DEBUG_EXCEPTIONS_BS =                              0x4000 //col:5997
VMX_PENDING_DEBUG_EXCEPTIONS_RTM =                             0x10000 //col:6000
VMX_VMEXIT_REASON_BASIC_EXIT_REASON =                          0xFFFF //col:6014
VMX_VMEXIT_REASON_ALWAYS0 =                                    0x10000 //col:6016
VMX_VMEXIT_REASON_RESERVED1 =                                  0x7FE0000 //col:6018
VMX_VMEXIT_REASON_ENCLAVE_MODE =                               0x8000000 //col:6020
VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT =                        0x10000000 //col:6022
VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT =                      0x20000000 //col:6024
VMX_VMEXIT_REASON_RESERVED2 =                                  0x40000000 //col:6026
VMX_VMEXIT_REASON_VM_ENTRY_FAILURE =                           0x80000000 //col:6028
IO_BITMAP_A_MIN =                                              0x00000000 //col:6035
IO_BITMAP_A_MAX =                                              0x00007FFF //col:6036
IO_BITMAP_B_MIN =                                              0x00008000 //col:6037
IO_BITMAP_B_MAX =                                              0x0000FFFF //col:6038
MSR_ID_LOW_MIN =                                               0x00000000 //col:6044
MSR_ID_LOW_MAX =                                               0x00001FFF //col:6045
MSR_ID_HIGH_MIN =                                              0xC0000000 //col:6046
MSR_ID_HIGH_MAX =                                              0xC0001FFF //col:6047
EPTP_MEMORY_TYPE =                                             0x07 //col:6062
EPTP_PAGE_WALK_LENGTH =                                        0x38 //col:6064
EPT_PAGE_WALK_LENGTH_4 =                                       0x00000003 //col:6065
EPTP_ENABLE_ACCESS_AND_DIRTY_FLAGS =                           0x40 //col:6067
EPTP_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES =                    0x80 //col:6069
EPTP_PAGE_FRAME_NUMBER =                                       0xFFFFFFFFF000 //col:6072
EPML4E_READ_ACCESS =                                           0x01 //col:6082
EPML4E_WRITE_ACCESS =                                          0x02 //col:6084
EPML4E_EXECUTE_ACCESS =                                        0x04 //col:6086
EPML4E_ACCESSED =                                              0x100 //col:6089
EPML4E_USER_MODE_EXECUTE =                                     0x400 //col:6092
EPML4E_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:6095
EPDPTE_1GB_READ_ACCESS =                                       0x01 //col:6105
EPDPTE_1GB_WRITE_ACCESS =                                      0x02 //col:6107
EPDPTE_1GB_EXECUTE_ACCESS =                                    0x04 //col:6109
EPDPTE_1GB_MEMORY_TYPE =                                       0x38 //col:6111
EPDPTE_1GB_IGNORE_PAT =                                        0x40 //col:6113
EPDPTE_1GB_LARGE_PAGE =                                        0x80 //col:6115
EPDPTE_1GB_ACCESSED =                                          0x100 //col:6117
EPDPTE_1GB_DIRTY =                                             0x200 //col:6119
EPDPTE_1GB_USER_MODE_EXECUTE =                                 0x400 //col:6121
EPDPTE_1GB_PAGE_FRAME_NUMBER =                                 0xFFFFC0000000 //col:6124
EPDPTE_1GB_VERIFY_GUEST_PAGING =                               0x200000000000000 //col:6127
EPDPTE_1GB_PAGING_WRITE_ACCESS =                               0x400000000000000 //col:6129
EPDPTE_1GB_SUPERVISOR_SHADOW_STACK =                           0x1000000000000000 //col:6132
EPDPTE_1GB_SUPPRESS_VE =                                       0x8000000000000000 //col:6135
EPDPTE_READ_ACCESS =                                           0x01 //col:6144
EPDPTE_WRITE_ACCESS =                                          0x02 //col:6146
EPDPTE_EXECUTE_ACCESS =                                        0x04 //col:6148
EPDPTE_ACCESSED =                                              0x100 //col:6151
EPDPTE_USER_MODE_EXECUTE =                                     0x400 //col:6154
EPDPTE_PAGE_FRAME_NUMBER =                                     0xFFFFFFFFF000 //col:6157
EPDE_2MB_READ_ACCESS =                                         0x01 //col:6167
EPDE_2MB_WRITE_ACCESS =                                        0x02 //col:6169
EPDE_2MB_EXECUTE_ACCESS =                                      0x04 //col:6171
EPDE_2MB_MEMORY_TYPE =                                         0x38 //col:6173
EPDE_2MB_IGNORE_PAT =                                          0x40 //col:6175
EPDE_2MB_LARGE_PAGE =                                          0x80 //col:6177
EPDE_2MB_ACCESSED =                                            0x100 //col:6179
EPDE_2MB_DIRTY =                                               0x200 //col:6181
EPDE_2MB_USER_MODE_EXECUTE =                                   0x400 //col:6183
EPDE_2MB_PAGE_FRAME_NUMBER =                                   0xFFFFFFE00000 //col:6186
EPDE_2MB_VERIFY_GUEST_PAGING =                                 0x200000000000000 //col:6189
EPDE_2MB_PAGING_WRITE_ACCESS =                                 0x400000000000000 //col:6191
EPDE_2MB_SUPERVISOR_SHADOW_STACK =                             0x1000000000000000 //col:6194
EPDE_2MB_SUPPRESS_VE =                                         0x8000000000000000 //col:6197
EPDE_READ_ACCESS =                                             0x01 //col:6206
EPDE_WRITE_ACCESS =                                            0x02 //col:6208
EPDE_EXECUTE_ACCESS =                                          0x04 //col:6210
EPDE_ACCESSED =                                                0x100 //col:6213
EPDE_USER_MODE_EXECUTE =                                       0x400 //col:6216
EPDE_PAGE_FRAME_NUMBER =                                       0xFFFFFFFFF000 //col:6219
EPTE_READ_ACCESS =                                             0x01 //col:6229
EPTE_WRITE_ACCESS =                                            0x02 //col:6231
EPTE_EXECUTE_ACCESS =                                          0x04 //col:6233
EPTE_MEMORY_TYPE =                                             0x38 //col:6235
EPTE_IGNORE_PAT =                                              0x40 //col:6237
EPTE_ACCESSED =                                                0x100 //col:6240
EPTE_DIRTY =                                                   0x200 //col:6242
EPTE_USER_MODE_EXECUTE =                                       0x400 //col:6244
EPTE_PAGE_FRAME_NUMBER =                                       0xFFFFFFFFF000 //col:6247
EPTE_VERIFY_GUEST_PAGING =                                     0x200000000000000 //col:6250
EPTE_PAGING_WRITE_ACCESS =                                     0x400000000000000 //col:6252
EPTE_SUPERVISOR_SHADOW_STACK =                                 0x1000000000000000 //col:6255
EPTE_SUB_PAGE_WRITE_PERMISSIONS =                              0x2000000000000000 //col:6257
EPTE_SUPPRESS_VE =                                             0x8000000000000000 //col:6260
EPT_ENTRY_READ_ACCESS =                                        0x01 //col:6269
EPT_ENTRY_WRITE_ACCESS =                                       0x02 //col:6271
EPT_ENTRY_EXECUTE_ACCESS =                                     0x04 //col:6273
EPT_ENTRY_MEMORY_TYPE =                                        0x38 //col:6275
EPT_ENTRY_IGNORE_PAT =                                         0x40 //col:6277
EPT_ENTRY_LARGE_PAGE =                                         0x80 //col:6279
EPT_ENTRY_ACCESSED =                                           0x100 //col:6281
EPT_ENTRY_DIRTY =                                              0x200 //col:6283
EPT_ENTRY_USER_MODE_EXECUTE =                                  0x400 //col:6285
EPT_ENTRY_PAGE_FRAME_NUMBER =                                  0xFFFFFFFFF000 //col:6288
EPT_ENTRY_SUPPRESS_VE =                                        0x8000000000000000 //col:6291
EPT_LEVEL_PML4E =                                              0x00000003 //col:6302
EPT_LEVEL_PDPTE =                                              0x00000002 //col:6303
EPT_LEVEL_PDE =                                                0x00000001 //col:6304
EPT_LEVEL_PTE =                                                0x00000000 //col:6305
EPML4_ENTRY_COUNT =                                            0x00000200 //col:6315
EPDPTE_ENTRY_COUNT =                                           0x00000200 //col:6316
EPDE_ENTRY_COUNT =                                             0x00000200 //col:6317
EPTE_ENTRY_COUNT =                                             0x00000200 //col:6318
INVEPT_SINGLE_CONTEXT =                                        0x00000001 //col:6327
INVEPT_ALL_CONTEXT =                                           0x00000002 //col:6328
INVVPID_INDIVIDUAL_ADDRESS =                                   0x00000000 //col:6333
INVVPID_SINGLE_CONTEXT =                                       0x00000001 //col:6334
INVVPID_ALL_CONTEXT =                                          0x00000002 //col:6335
INVVPID_SINGLE_CONTEXT_RETAINING_GLOBALS =                     0x00000003 //col:6336
HLATP_PAGE_LEVEL_WRITE_THROUGH =                               0x08 //col:6357
HLATP_PAGE_LEVEL_CACHE_DISABLE =                               0x10 //col:6359
HLATP_PAGE_FRAME_NUMBER =                                      0xFFFFFFFFF000 //col:6362
VMCS_COMPONENT_ENCODING_ACCESS_TYPE =                          0x01 //col:6396
VMCS_COMPONENT_ENCODING_INDEX =                                0x3FE //col:6398
VMCS_COMPONENT_ENCODING_TYPE =                                 0xC00 //col:6400
VMCS_COMPONENT_ENCODING_MUST_BE_ZERO =                         0x1000 //col:6402
VMCS_COMPONENT_ENCODING_WIDTH =                                0x6000 //col:6404
VMCS_CTRL_VPID =                                               0x00000000 //col:6421
VMCS_CTRL_POSTED_INTR_NOTIFY_VECTOR =                          0x00000002 //col:6422
VMCS_CTRL_EPTP_INDEX =                                         0x00000004 //col:6423
VMCS_CTRL_HLAT_PREFIX_SIZE =                                   0x00000006 //col:6424
VMCS_GUEST_ES_SEL =                                            0x00000800 //col:6434
VMCS_GUEST_CS_SEL =                                            0x00000802 //col:6435
VMCS_GUEST_SS_SEL =                                            0x00000804 //col:6436
VMCS_GUEST_DS_SEL =                                            0x00000806 //col:6437
VMCS_GUEST_FS_SEL =                                            0x00000808 //col:6438
VMCS_GUEST_GS_SEL =                                            0x0000080A //col:6439
VMCS_GUEST_LDTR_SEL =                                          0x0000080C //col:6440
VMCS_GUEST_TR_SEL =                                            0x0000080E //col:6441
VMCS_GUEST_INTR_STATUS =                                       0x00000810 //col:6442
VMCS_GUEST_PML_INDEX =                                         0x00000812 //col:6443
VMCS_HOST_ES_SEL =                                             0x00000C00 //col:6453
VMCS_HOST_CS_SEL =                                             0x00000C02 //col:6454
VMCS_HOST_SS_SEL =                                             0x00000C04 //col:6455
VMCS_HOST_DS_SEL =                                             0x00000C06 //col:6456
VMCS_HOST_FS_SEL =                                             0x00000C08 //col:6457
VMCS_HOST_GS_SEL =                                             0x00000C0A //col:6458
VMCS_HOST_TR_SEL =                                             0x00000C0C //col:6459
VMCS_CTRL_IO_BITMAP_A =                                        0x00002000 //col:6478
VMCS_CTRL_IO_BITMAP_B =                                        0x00002002 //col:6479
VMCS_CTRL_MSR_BITMAP =                                         0x00002004 //col:6480
VMCS_CTRL_VMEXIT_MSR_STORE =                                   0x00002006 //col:6481
VMCS_CTRL_VMEXIT_MSR_LOAD =                                    0x00002008 //col:6482
VMCS_CTRL_VMENTRY_MSR_LOAD =                                   0x0000200A //col:6483
VMCS_CTRL_EXEC_VMCS_PTR =                                      0x0000200C //col:6484
VMCS_CTRL_PML_ADDR =                                           0x0000200E //col:6485
VMCS_CTRL_TSC_OFFSET =                                         0x00002010 //col:6486
VMCS_CTRL_VAPIC_PAGEADDR =                                     0x00002012 //col:6487
VMCS_CTRL_APIC_ACCESSADDR =                                    0x00002014 //col:6488
VMCS_CTRL_POSTED_INTR_DESC =                                   0x00002016 //col:6489
VMCS_CTRL_VMFUNC_CTRLS =                                       0x00002018 //col:6490
VMCS_CTRL_EPTP =                                               0x0000201A //col:6491
VMCS_CTRL_EOI_BITMAP_0 =                                       0x0000201C //col:6492
VMCS_CTRL_EOI_BITMAP_1 =                                       0x0000201E //col:6493
VMCS_CTRL_EOI_BITMAP_2 =                                       0x00002020 //col:6494
VMCS_CTRL_EOI_BITMAP_3 =                                       0x00002022 //col:6495
VMCS_CTRL_EPTP_LIST =                                          0x00002024 //col:6496
VMCS_CTRL_VMREAD_BITMAP =                                      0x00002026 //col:6497
VMCS_CTRL_VMWRITE_BITMAP =                                     0x00002028 //col:6498
VMCS_CTRL_VIRTXCPT_INFO_ADDR =                                 0x0000202A //col:6499
VMCS_CTRL_XSS_EXITING_BITMAP =                                 0x0000202C //col:6500
VMCS_CTRL_ENCLS_EXITING_BITMAP =                               0x0000202E //col:6501
VMCS_CTRL_SPP_TABLE_POINTER =                                  0x00002030 //col:6502
VMCS_CTRL_TSC_MULTIPLIER =                                     0x00002032 //col:6503
VMCS_CTRL_PROC_EXEC3 =                                         0x00002034 //col:6504
VMCS_CTRL_ENCLV_EXITING_BITMAP =                               0x00002036 //col:6505
VMCS_CTRL_HLATP =                                              0x00002040 //col:6506
VMCS_CTRL_SECONDARY_EXIT =                                     0x00002044 //col:6507
VMCS_GUEST_PHYS_ADDR =                                         0x00002400 //col:6517
VMCS_GUEST_VMCS_LINK_PTR =                                     0x00002800 //col:6527
VMCS_GUEST_DEBUGCTL =                                          0x00002802 //col:6528
VMCS_GUEST_PAT =                                               0x00002804 //col:6529
VMCS_GUEST_EFER =                                              0x00002806 //col:6530
VMCS_GUEST_PERF_GLOBAL_CTRL =                                  0x00002808 //col:6531
VMCS_GUEST_PDPTE0 =                                            0x0000280A //col:6532
VMCS_GUEST_PDPTE1 =                                            0x0000280C //col:6533
VMCS_GUEST_PDPTE2 =                                            0x0000280E //col:6534
VMCS_GUEST_PDPTE3 =                                            0x00002810 //col:6535
VMCS_GUEST_BNDCFGS =                                           0x00002812 //col:6536
VMCS_GUEST_RTIT_CTL =                                          0x00002814 //col:6537
VMCS_GUEST_LBR_CTL =                                           0x00002816 //col:6538
VMCS_GUEST_PKRS =                                              0x00002818 //col:6539
VMCS_HOST_PAT =                                                0x00002C00 //col:6549
VMCS_HOST_EFER =                                               0x00002C02 //col:6550
VMCS_HOST_PERF_GLOBAL_CTRL =                                   0x00002C04 //col:6551
VMCS_HOST_PKRS =                                               0x00002C06 //col:6552
VMCS_CTRL_PIN_EXEC =                                           0x00004000 //col:6571
VMCS_CTRL_PROC_EXEC =                                          0x00004002 //col:6572
VMCS_CTRL_EXCEPTION_BITMAP =                                   0x00004004 //col:6573
VMCS_CTRL_PAGEFAULT_ERROR_MASK =                               0x00004006 //col:6574
VMCS_CTRL_PAGEFAULT_ERROR_MATCH =                              0x00004008 //col:6575
VMCS_CTRL_CR3_TARGET_COUNT =                                   0x0000400A //col:6576
VMCS_CTRL_PRIMARY_EXIT =                                       0x0000400C //col:6577
VMCS_CTRL_EXIT_MSR_STORE_COUNT =                               0x0000400E //col:6578
VMCS_CTRL_EXIT_MSR_LOAD_COUNT =                                0x00004010 //col:6579
VMCS_CTRL_ENTRY =                                              0x00004012 //col:6580
VMCS_CTRL_ENTRY_MSR_LOAD_COUNT =                               0x00004014 //col:6581
VMCS_CTRL_ENTRY_INTERRUPTION_INFO =                            0x00004016 //col:6582
VMCS_CTRL_ENTRY_EXCEPTION_ERRCODE =                            0x00004018 //col:6583
VMCS_CTRL_ENTRY_INSTR_LENGTH =                                 0x0000401A //col:6584
VMCS_CTRL_TPR_THRESHOLD =                                      0x0000401C //col:6585
VMCS_CTRL_PROC_EXEC2 =                                         0x0000401E //col:6586
VMCS_CTRL_PLE_GAP =                                            0x00004020 //col:6587
VMCS_CTRL_PLE_WINDOW =                                         0x00004022 //col:6588
VMCS_VM_INSTR_ERROR =                                          0x00004400 //col:6598
VMCS_EXIT_REASON =                                             0x00004402 //col:6599
VMCS_EXIT_INTERRUPTION_INFO =                                  0x00004404 //col:6600
VMCS_EXIT_INTERRUPTION_ERROR_CODE =                            0x00004406 //col:6601
VMCS_IDT_VECTORING_INFO =                                      0x00004408 //col:6602
VMCS_IDT_VECTORING_ERROR_CODE =                                0x0000440A //col:6603
VMCS_EXIT_INSTR_LENGTH =                                       0x0000440C //col:6604
VMCS_EXIT_INSTR_INFO =                                         0x0000440E //col:6605
VMCS_GUEST_ES_LIMIT =                                          0x00004800 //col:6615
VMCS_GUEST_CS_LIMIT =                                          0x00004802 //col:6616
VMCS_GUEST_SS_LIMIT =                                          0x00004804 //col:6617
VMCS_GUEST_DS_LIMIT =                                          0x00004806 //col:6618
VMCS_GUEST_FS_LIMIT =                                          0x00004808 //col:6619
VMCS_GUEST_GS_LIMIT =                                          0x0000480A //col:6620
VMCS_GUEST_LDTR_LIMIT =                                        0x0000480C //col:6621
VMCS_GUEST_TR_LIMIT =                                          0x0000480E //col:6622
VMCS_GUEST_GDTR_LIMIT =                                        0x00004810 //col:6623
VMCS_GUEST_IDTR_LIMIT =                                        0x00004812 //col:6624
VMCS_GUEST_ES_ACCESS_RIGHTS =                                  0x00004814 //col:6625
VMCS_GUEST_CS_ACCESS_RIGHTS =                                  0x00004816 //col:6626
VMCS_GUEST_SS_ACCESS_RIGHTS =                                  0x00004818 //col:6627
VMCS_GUEST_DS_ACCESS_RIGHTS =                                  0x0000481A //col:6628
VMCS_GUEST_FS_ACCESS_RIGHTS =                                  0x0000481C //col:6629
VMCS_GUEST_GS_ACCESS_RIGHTS =                                  0x0000481E //col:6630
VMCS_GUEST_LDTR_ACCESS_RIGHTS =                                0x00004820 //col:6631
VMCS_GUEST_TR_ACCESS_RIGHTS =                                  0x00004822 //col:6632
VMCS_GUEST_INTERRUPTIBILITY_STATE =                            0x00004824 //col:6633
VMCS_GUEST_ACTIVITY_STATE =                                    0x00004826 //col:6634
VMCS_GUEST_SMBASE =                                            0x00004828 //col:6635
VMCS_GUEST_SYSENTER_CS =                                       0x0000482A //col:6636
VMCS_GUEST_PREEMPT_TIMER_VALUE =                               0x0000482E //col:6637
VMCS_HOST_SYSENTER_CS =                                        0x00004C00 //col:6647
VMCS_CTRL_CR0_MASK =                                           0x00006000 //col:6666
VMCS_CTRL_CR4_MASK =                                           0x00006002 //col:6667
VMCS_CTRL_CR0_READ_SHADOW =                                    0x00006004 //col:6668
VMCS_CTRL_CR4_READ_SHADOW =                                    0x00006006 //col:6669
VMCS_CTRL_CR3_TARGET_VAL0 =                                    0x00006008 //col:6670
VMCS_CTRL_CR3_TARGET_VAL1 =                                    0x0000600A //col:6671
VMCS_CTRL_CR3_TARGET_VAL2 =                                    0x0000600C //col:6672
VMCS_CTRL_CR3_TARGET_VAL3 =                                    0x0000600E //col:6673
VMCS_EXIT_QUALIFICATION =                                      0x00006400 //col:6683
VMCS_IO_RCX =                                                  0x00006402 //col:6684
VMCS_IO_RSI =                                                  0x00006404 //col:6685
VMCS_IO_RDI =                                                  0x00006406 //col:6686
VMCS_IO_RIP =                                                  0x00006408 //col:6687
VMCS_EXIT_GUEST_LINEAR_ADDR =                                  0x0000640A //col:6688
VMCS_GUEST_CR0 =                                               0x00006800 //col:6698
VMCS_GUEST_CR3 =                                               0x00006802 //col:6699
VMCS_GUEST_CR4 =                                               0x00006804 //col:6700
VMCS_GUEST_ES_BASE =                                           0x00006806 //col:6701
VMCS_GUEST_CS_BASE =                                           0x00006808 //col:6702
VMCS_GUEST_SS_BASE =                                           0x0000680A //col:6703
VMCS_GUEST_DS_BASE =                                           0x0000680C //col:6704
VMCS_GUEST_FS_BASE =                                           0x0000680E //col:6705
VMCS_GUEST_GS_BASE =                                           0x00006810 //col:6706
VMCS_GUEST_LDTR_BASE =                                         0x00006812 //col:6707
VMCS_GUEST_TR_BASE =                                           0x00006814 //col:6708
VMCS_GUEST_GDTR_BASE =                                         0x00006816 //col:6709
VMCS_GUEST_IDTR_BASE =                                         0x00006818 //col:6710
VMCS_GUEST_DR7 =                                               0x0000681A //col:6711
VMCS_GUEST_RSP =                                               0x0000681C //col:6712
VMCS_GUEST_RIP =                                               0x0000681E //col:6713
VMCS_GUEST_RFLAGS =                                            0x00006820 //col:6714
VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS =                          0x00006822 //col:6715
VMCS_GUEST_SYSENTER_ESP =                                      0x00006824 //col:6716
VMCS_GUEST_SYSENTER_EIP =                                      0x00006826 //col:6717
VMCS_GUEST_S_CET =                                             0x00006C28 //col:6718
VMCS_GUEST_SSP =                                               0x00006C2A //col:6719
VMCS_GUEST_INTERRUPT_SSP_TABLE_ADDR =                          0x00006C2C //col:6720
VMCS_HOST_CR0 =                                                0x00006C00 //col:6730
VMCS_HOST_CR3 =                                                0x00006C02 //col:6731
VMCS_HOST_CR4 =                                                0x00006C04 //col:6732
VMCS_HOST_FS_BASE =                                            0x00006C06 //col:6733
VMCS_HOST_GS_BASE =                                            0x00006C08 //col:6734
VMCS_HOST_TR_BASE =                                            0x00006C0A //col:6735
VMCS_HOST_GDTR_BASE =                                          0x00006C0C //col:6736
VMCS_HOST_IDTR_BASE =                                          0x00006C0E //col:6737
VMCS_HOST_SYSENTER_ESP =                                       0x00006C10 //col:6738
VMCS_HOST_SYSENTER_EIP =                                       0x00006C12 //col:6739
VMCS_HOST_RSP =                                                0x00006C14 //col:6740
VMCS_HOST_RIP =                                                0x00006C16 //col:6741
VMCS_HOST_S_CET =                                              0x00006C18 //col:6742
VMCS_HOST_SSP =                                                0x00006C1A //col:6743
VMCS_HOST_INTERRUPT_SSP_TABLE_ADDR =                           0x00006C1C //col:6744
EXTERNAL_INTERRUPT =                                           0x00000000 //col:6757
NON_MASKABLE_INTERRUPT =                                       0x00000002 //col:6758
HARDWARE_EXCEPTION =                                           0x00000003 //col:6759
SOFTWARE_INTERRUPT =                                           0x00000004 //col:6760
PRIVILEGED_SOFTWARE_EXCEPTION =                                0x00000005 //col:6761
SOFTWARE_EXCEPTION =                                           0x00000006 //col:6762
OTHER_EVENT =                                                  0x00000007 //col:6763
VMENTRY_INTERRUPT_INFO_VECTOR =                                0xFF //col:6771
VMENTRY_INTERRUPT_INFO_INTERRUPTION_TYPE =                     0x700 //col:6773
VMENTRY_INTERRUPT_INFO_DELIVER_ERROR_CODE =                    0x800 //col:6775
VMENTRY_INTERRUPT_INFO_VALID =                                 0x80000000 //col:6778
VMEXIT_INTERRUPT_INFO_VECTOR =                                 0xFF //col:6787
VMEXIT_INTERRUPT_INFO_INTERRUPTION_TYPE =                      0x700 //col:6789
VMEXIT_INTERRUPT_INFO_ERROR_CODE_VALID =                       0x800 //col:6791
VMEXIT_INTERRUPT_INFO_NMI_UNBLOCKING =                         0x1000 //col:6793
VMEXIT_INTERRUPT_INFO_VALID =                                  0x80000000 //col:6796
APIC_BASE =                                                    0xFEE00000 //col:6811
APIC_ID =                                                      0x00000020 //col:6812
APIC_VERSION =                                                 0x00000030 //col:6813
APIC_TPR =                                                     0x00000080 //col:6814
APIC_APR =                                                     0x00000090 //col:6815
APIC_PPR =                                                     0x000000A0 //col:6816
APIC_EOI =                                                     0x000000B0 //col:6817
APIC_REMOTE_READ =                                             0x000000C0 //col:6818
APIC_LOGICAL_DESTINATION =                                     0x000000D0 //col:6819
APIC_DESTINATION_FORMAT =                                      0x000000E0 //col:6820
APIC_SIV =                                                     0x000000F0 //col:6821
APIC_ISR_31_0 =                                                0x00000100 //col:6822
APIC_ISR_63_32 =                                               0x00000110 //col:6823
APIC_ISR_95_64 =                                               0x00000120 //col:6824
APIC_ISR_127_96 =                                              0x00000130 //col:6825
APIC_ISR_159_128 =                                             0x00000140 //col:6826
APIC_ISR_191_160 =                                             0x00000150 //col:6827
APIC_ISR_223_192 =                                             0x00000160 //col:6828
APIC_ISR_255_224 =                                             0x00000170 //col:6829
APIC_TMR_31_0 =                                                0x00000180 //col:6830
APIC_TMR_63_32 =                                               0x00000190 //col:6831
APIC_TMR_95_64 =                                               0x000001A0 //col:6832
APIC_TMR_127_96 =                                              0x000001B0 //col:6833
APIC_TMR_159_128 =                                             0x000001C0 //col:6834
APIC_TMR_191_160 =                                             0x000001D0 //col:6835
APIC_TMR_223_192 =                                             0x000001E0 //col:6836
APIC_TMR_255_224 =                                             0x000001F0 //col:6837
APIC_IRR_31_0 =                                                0x00000200 //col:6838
APIC_IRR_63_32 =                                               0x00000210 //col:6839
APIC_IRR_95_64 =                                               0x00000220 //col:6840
APIC_IRR_127_96 =                                              0x00000230 //col:6841
APIC_IRR_159_128 =                                             0x00000240 //col:6842
APIC_IRR_191_160 =                                             0x00000250 //col:6843
APIC_IRR_223_192 =                                             0x00000260 //col:6844
APIC_IRR_255_224 =                                             0x00000270 //col:6845
APIC_ERROR_STATUS =                                            0x00000280 //col:6846
APIC_CMCI =                                                    0x000002F0 //col:6847
APIC_ICR_0_31 =                                                0x00000300 //col:6848
APIC_ICR_32_63 =                                               0x00000310 //col:6849
APIC_LVT_TIMER =                                               0x00000320 //col:6850
APIC_LVT_THERMAL_SENSOR =                                      0x00000330 //col:6851
APIC_LVT_PERFORMANCE_MONITORING_COUNTERS =                     0x00000340 //col:6852
APIC_LVT_LINT0 =                                               0x00000350 //col:6853
APIC_LVT_LINT1 =                                               0x00000360 //col:6854
APIC_LVT_ERROR =                                               0x00000370 //col:6855
APIC_INITIAL_COUNT =                                           0x00000380 //col:6856
APIC_CURRENT_COUNT =                                           0x00000390 //col:6857
APIC_DIVIDE_CONFIGURATION =                                    0x000003E0 //col:6858
EFL_CARRY_FLAG =                                               0x01 //col:6866
EFL_READ_AS_1 =                                                0x02 //col:6868
EFL_PARITY_FLAG =                                              0x04 //col:6870
EFL_AUXILIARY_CARRY_FLAG =                                     0x10 //col:6873
EFL_ZERO_FLAG =                                                0x40 //col:6876
EFL_SIGN_FLAG =                                                0x80 //col:6878
EFL_TRAP_FLAG =                                                0x100 //col:6880
EFL_INTERRUPT_ENABLE_FLAG =                                    0x200 //col:6882
EFL_DIRECTION_FLAG =                                           0x400 //col:6884
EFL_OVERFLOW_FLAG =                                            0x800 //col:6886
EFL_IO_PRIVILEGE_LEVEL =                                       0x3000 //col:6888
EFL_NESTED_TASK_FLAG =                                         0x4000 //col:6890
EFL_RESUME_FLAG =                                              0x10000 //col:6893
EFL_VIRTUAL_8086_MODE_FLAG =                                   0x20000 //col:6895
EFL_ALIGNMENT_CHECK_FLAG =                                     0x40000 //col:6897
EFL_VIRTUAL_INTERRUPT_FLAG =                                   0x80000 //col:6899
EFL_VIRTUAL_INTERRUPT_PENDING_FLAG =                           0x100000 //col:6901
EFL_IDENTIFICATION_FLAG =                                      0x200000 //col:6903
RFL_CARRY_FLAG =                                               0x01 //col:6913
RFL_READ_AS_1 =                                                0x02 //col:6915
RFL_PARITY_FLAG =                                              0x04 //col:6917
RFL_AUXILIARY_CARRY_FLAG =                                     0x10 //col:6920
RFL_ZERO_FLAG =                                                0x40 //col:6923
RFL_SIGN_FLAG =                                                0x80 //col:6925
RFL_TRAP_FLAG =                                                0x100 //col:6927
RFL_INTERRUPT_ENABLE_FLAG =                                    0x200 //col:6929
RFL_DIRECTION_FLAG =                                           0x400 //col:6931
RFL_OVERFLOW_FLAG =                                            0x800 //col:6933
RFL_IO_PRIVILEGE_LEVEL =                                       0x3000 //col:6935
RFL_NESTED_TASK_FLAG =                                         0x4000 //col:6937
RFL_RESUME_FLAG =                                              0x10000 //col:6940
RFL_VIRTUAL_8086_MODE_FLAG =                                   0x20000 //col:6942
RFL_ALIGNMENT_CHECK_FLAG =                                     0x40000 //col:6944
RFL_VIRTUAL_INTERRUPT_FLAG =                                   0x80000 //col:6946
RFL_VIRTUAL_INTERRUPT_PENDING_FLAG =                           0x100000 //col:6948
RFL_IDENTIFICATION_FLAG =                                      0x200000 //col:6950
CONTROL_PROTECTION_EXCEPTION_CPEC =                            0x7FFF //col:6965
CONTROL_PROTECTION_EXCEPTION_ENCL =                            0x8000 //col:6967
DIVIDE_ERROR =                                                 0x00000000 //col:6974
DEBUG =                                                        0x00000001 //col:6975
NMI =                                                          0x00000002 //col:6976
BREAKPOINT =                                                   0x00000003 //col:6977
OVERFLOW =                                                     0x00000004 //col:6978
BOUND_RANGE_EXCEEDED =                                         0x00000005 //col:6979
INVALID_OPCODE =                                               0x00000006 //col:6980
DEVICE_NOT_AVAILABLE =                                         0x00000007 //col:6981
DOUBLE_FAULT =                                                 0x00000008 //col:6982
COPROCESSOR_SEGMENT_OVERRUN =                                  0x00000009 //col:6983
INVALID_TSS =                                                  0x0000000A //col:6984
SEGMENT_NOT_PRESENT =                                          0x0000000B //col:6985
STACK_SEGMENT_FAULT =                                          0x0000000C //col:6986
GENERAL_PROTECTION =                                           0x0000000D //col:6987
PAGE_FAULT =                                                   0x0000000E //col:6988
X87_FLOATING_POINT_ERROR =                                     0x00000010 //col:6989
ALIGNMENT_CHECK =                                              0x00000011 //col:6990
MACHINE_CHECK =                                                0x00000012 //col:6991
SIMD_FLOATING_POINT_ERROR =                                    0x00000013 //col:6992
VIRTUALIZATION_EXCEPTION =                                     0x00000014 //col:6993
CONTROL_PROTECTION =                                           0x00000015 //col:6994
EXCEPTION_ERROR_CODE_EXTERNAL_EVENT =                          0x01 //col:7002
EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION =                     0x02 //col:7004
EXCEPTION_ERROR_CODE_GDT_LDT =                                 0x04 //col:7006
EXCEPTION_ERROR_CODE_INDEX =                                   0xFFF8 //col:7008
PAGE_FAULT_EXCEPTION_PRESENT =                                 0x01 //col:7018
PAGE_FAULT_EXCEPTION_WRITE =                                   0x02 //col:7020
PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS =                        0x04 //col:7022
PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION =                  0x08 //col:7024
PAGE_FAULT_EXCEPTION_EXECUTE =                                 0x10 //col:7026
PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION =                0x20 //col:7028
PAGE_FAULT_EXCEPTION_SHADOW_STACK =                            0x40 //col:7030
PAGE_FAULT_EXCEPTION_HLAT =                                    0x80 //col:7032
PAGE_FAULT_EXCEPTION_SGX =                                     0x8000 //col:7035
MEMORY_TYPE_UC =                                               0x00000000 //col:7051
MEMORY_TYPE_WC =                                               0x00000001 //col:7052
MEMORY_TYPE_WT =                                               0x00000004 //col:7053
MEMORY_TYPE_WP =                                               0x00000005 //col:7054
MEMORY_TYPE_WB =                                               0x00000006 //col:7055
MEMORY_TYPE_UC_MINUS =                                         0x00000007 //col:7056
MEMORY_TYPE_INVALID =                                          0x000000FF //col:7057
VTD_Lower64_PRESENT =                                          0x01 //col:7071
VTD_Lower64_CONTEXT_TABLE_POINTER =                            0xFFFFFFFFFFFFF000 //col:7074
VTD_Upper64_RESERVED =                                         0xFFFFFFFFFFFFFFFF //col:7083
VTD_Lower64_PRESENT =                                          0x01 //col:7095
VTD_Lower64_FAULT_PROCESSING_DISABLE =                         0x02 //col:7097
VTD_Lower64_TRANSLATION_TYPE =                                 0x0C //col:7099
VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER =            0xFFFFFFFFFFFFF000 //col:7102
VTD_Upper64_ADDRESS_WIDTH =                                    0x07 //col:7111
VTD_Upper64_IGNORED =                                          0x78 //col:7113
VTD_Upper64_DOMAIN_IDENTIFIER =                                0x3FF00 //col:7116
VTD_ROOT_ENTRY_COUNT =                                         0x00000100 //col:7130
VTD_CONTEXT_ENTRY_COUNT =                                      0x00000100 //col:7131
VTD_VERSION =                                                  0x00000000 //col:7136
VTD_VERSION_MINOR =                                            0x0F //col:7140
VTD_VERSION_MAJOR =                                            0xF0 //col:7142
VTD_CAPABILITY =                                               0x00000008 //col:7149
VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED =                   0x07 //col:7153
VTD_CAPABILITY_ADVANCED_FAULT_LOGGING =                        0x08 //col:7155
VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING =                0x10 //col:7157
VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION =                   0x20 //col:7159
VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION =                  0x40 //col:7161
VTD_CAPABILITY_CACHING_MODE =                                  0x80 //col:7163
VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS =       0x1F00 //col:7165
VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH =                   0x3F0000 //col:7168
VTD_CAPABILITY_ZERO_LENGTH_READ =                              0x400000 //col:7170
VTD_CAPABILITY_DEPRECATED =                                    0x800000 //col:7172
VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET =               0x3FF000000 //col:7174
VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT =               0x3C00000000 //col:7176
VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION =                   0x8000000000 //col:7179
VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS =           0xFF0000000000 //col:7181
VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE =                    0x3F000000000000 //col:7183
VTD_CAPABILITY_WRITE_DRAINING =                                0x40000000000000 //col:7185
VTD_CAPABILITY_READ_DRAINING =                                 0x80000000000000 //col:7187
VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT =               0x100000000000000 //col:7189
VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT =                     0x800000000000000 //col:7192
VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT =             0x1000000000000000 //col:7194
VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT = 0x4000000000000000 //col:7197
VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT =       0x8000000000000000 //col:7199
VTD_EXTENDED_CAPABILITY =                                      0x00000010 //col:7205
VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY =                  0x01 //col:7209
VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT =          0x02 //col:7211
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT =                   0x04 //col:7213
VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT =          0x08 //col:7215
VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE =              0x10 //col:7217
VTD_EXTENDED_CAPABILITY_DEPRECATED1 =                          0x20 //col:7219
VTD_EXTENDED_CAPABILITY_PASS_THROUGH =                         0x40 //col:7221
VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL =                        0x80 //col:7223
VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET =                0x3FF00 //col:7225
VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE =            0xF00000 //col:7228
VTD_EXTENDED_CAPABILITY_DEPRECATED2 =                          0x1000000 //col:7230
VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT =                  0x2000000 //col:7232
VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT =           0x4000000 //col:7234
VTD_EXTENDED_CAPABILITY_DEPRECATED3 =                          0x10000000 //col:7237
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT =                 0x20000000 //col:7239
VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT =              0x40000000 //col:7241
VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT =                0x200000000 //col:7244
VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT =       0x400000000 //col:7246
VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED =                 0xF800000000 //col:7248
VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT =     0x10000000000 //col:7250
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE =     0x20000000000 //col:7252
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT =           0x40000000000 //col:7254
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT =    0x80000000000 //col:7256
VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT =              0x100000000000 //col:7258
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT =  0x200000000000 //col:7260
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT =     0x400000000000 //col:7262
VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT =      0x800000000000 //col:7264
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY =    0x1000000000000 //col:7266
VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT =                    0x2000000000000 //col:7268
VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT =               0x10000000000000 //col:7271
VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT =                     0x20000000000000 //col:7273
VTD_GLOBAL_COMMAND =                                           0x00000018 //col:7280
VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT =            0x800000 //col:7285
VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER =         0x1000000 //col:7287
VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE =                0x2000000 //col:7289
VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE =                0x4000000 //col:7291
VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH =                        0x8000000 //col:7293
VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING =             0x10000000 //col:7295
VTD_GLOBAL_COMMAND_SET_FAULT_LOG =                             0x20000000 //col:7297
VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER =                    0x40000000 //col:7299
VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE =                        0x80000000 //col:7301
VTD_GLOBAL_STATUS =                                            0x0000001C //col:7307
VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS =      0x800000 //col:7312
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS =   0x1000000 //col:7314
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS =          0x2000000 //col:7316
VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS =          0x4000000 //col:7318
VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS =                  0x8000000 //col:7320
VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS =              0x10000000 //col:7322
VTD_GLOBAL_STATUS_FAULT_LOG_STATUS =                           0x20000000 //col:7324
VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS =                  0x40000000 //col:7326
VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS =                  0x80000000 //col:7328
VTD_ROOT_TABLE_ADDRESS =                                       0x00000020 //col:7334
VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE =                0xC00 //col:7339
VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS =                    0xFFFFFFFFFFFFF000 //col:7341
VTD_CONTEXT_COMMAND =                                          0x00000028 //col:7347
VTD_CONTEXT_COMMAND_DOMAIN_ID =                                0xFFFF //col:7351
VTD_CONTEXT_COMMAND_SOURCE_ID =                                0xFFFF0000 //col:7353
VTD_CONTEXT_COMMAND_FUNCTION_MASK =                            0x300000000 //col:7355
VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY =  0x1800000000000000 //col:7358
VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY = 0x6000000000000000 //col:7360
VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE =                 0x8000000000000000 //col:7362
VTD_INVALIDATE_ADDRESS =                                       0x00000000 //col:7368
VTD_INVALIDATE_ADDRESS_ADDRESS_MASK =                          0x3F //col:7372
VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT =                     0x40 //col:7374
VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS =                          0xFFFFFFFFFFFFF000 //col:7377
VTD_IOTLB_INVALIDATE =                                         0x00000008 //col:7383
VTD_IOTLB_INVALIDATE_DOMAIN_ID =                               0xFFFF00000000 //col:7388
VTD_IOTLB_INVALIDATE_DRAIN_WRITES =                            0x1000000000000 //col:7390
VTD_IOTLB_INVALIDATE_DRAIN_READS =                             0x2000000000000 //col:7392
VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY =   0x600000000000000 //col:7395
VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY =  0x3000000000000000 //col:7398
VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB =                        0x8000000000000000 //col:7401
XCR0_X87 =                                                     0x01 //col:7414
XCR0_SSE =                                                     0x02 //col:7416
XCR0_AVX =                                                     0x04 //col:7418
XCR0_BNDREG =                                                  0x08 //col:7420
XCR0_BNDCSR =                                                  0x10 //col:7422
XCR0_OPMASK =                                                  0x20 //col:7424
XCR0_ZMM_HI256 =                                               0x40 //col:7426
XCR0_ZMM_HI16 =                                                0x80 //col:7428
XCR0_PKRU =                                                    0x200 //col:7431
)

type (
Ia32DefinesOnly interface{
#if defined()(ok bool)//col:50
#pragma pack()(ok bool)//col:5195
#pragma pack()(ok bool)//col:5202
#pragma pack()(ok bool)//col:5229
#pragma pack()(ok bool)//col:5415
#pragma pack()(ok bool)//col:5543
}


















































)

func NewIa32DefinesOnly() { return & ia32DefinesOnly{} }

func (i *ia32DefinesOnly)#if defined()(ok bool){//col:50



































return true
}



















































func (i *ia32DefinesOnly)#pragma pack()(ok bool){//col:5195





return true
}



















































func (i *ia32DefinesOnly)#pragma pack()(ok bool){//col:5202






return true
}



















































func (i *ia32DefinesOnly)#pragma pack()(ok bool){//col:5229

























return true
}



















































func (i *ia32DefinesOnly)#pragma pack()(ok bool){//col:5415


















return true
}



















































func (i *ia32DefinesOnly)#pragma pack()(ok bool){//col:5543
































































































return true
}





















































