package out
const(
CR0_PROTECTION_ENABLE_BIT =                                    0 //col:73
CR0_PROTECTION_ENABLE_FLAG =                                   0x01 //col:74
CR0_PROTECTION_ENABLE_MASK =                                   0x01 //col:75
CR0_PROTECTION_ENABLE(_) =                                     (((_) >> 0) & 0x01) //col:76
CR0_MONITOR_COPROCESSOR_BIT =                                  1 //col:86
CR0_MONITOR_COPROCESSOR_FLAG =                                 0x02 //col:87
CR0_MONITOR_COPROCESSOR_MASK =                                 0x01 //col:88
CR0_MONITOR_COPROCESSOR(_) =                                   (((_) >> 1) & 0x01) //col:89
CR0_EMULATE_FPU_BIT =                                          2 //col:108
CR0_EMULATE_FPU_FLAG =                                         0x04 //col:109
CR0_EMULATE_FPU_MASK =                                         0x01 //col:110
CR0_EMULATE_FPU(_) =                                           (((_) >> 2) & 0x01) //col:111
CR0_TASK_SWITCHED_BIT =                                        3 //col:135
CR0_TASK_SWITCHED_FLAG =                                       0x08 //col:136
CR0_TASK_SWITCHED_MASK =                                       0x01 //col:137
CR0_TASK_SWITCHED(_) =                                         (((_) >> 3) & 0x01) //col:138
CR0_EXTENSION_TYPE_BIT =                                       4 //col:148
CR0_EXTENSION_TYPE_FLAG =                                      0x10 //col:149
CR0_EXTENSION_TYPE_MASK =                                      0x01 //col:150
CR0_EXTENSION_TYPE(_) =                                        (((_) >> 4) & 0x01) //col:151
CR0_NUMERIC_ERROR_BIT =                                        5 //col:171
CR0_NUMERIC_ERROR_FLAG =                                       0x20 //col:172
CR0_NUMERIC_ERROR_MASK =                                       0x01 //col:173
CR0_NUMERIC_ERROR(_) =                                         (((_) >> 5) & 0x01) //col:174
CR0_WRITE_PROTECT_BIT =                                        16 //col:188
CR0_WRITE_PROTECT_FLAG =                                       0x10000 //col:189
CR0_WRITE_PROTECT_MASK =                                       0x01 //col:190
CR0_WRITE_PROTECT(_) =                                         (((_) >> 16) & 0x01) //col:191
CR0_ALIGNMENT_MASK_BIT =                                       18 //col:202
CR0_ALIGNMENT_MASK_FLAG =                                      0x40000 //col:203
CR0_ALIGNMENT_MASK_MASK =                                      0x01 //col:204
CR0_ALIGNMENT_MASK(_) =                                        (((_) >> 18) & 0x01) //col:205
CR0_NOT_WRITE_THROUGH_BIT =                                    29 //col:215
CR0_NOT_WRITE_THROUGH_FLAG =                                   0x20000000 //col:216
CR0_NOT_WRITE_THROUGH_MASK =                                   0x01 //col:217
CR0_NOT_WRITE_THROUGH(_) =                                     (((_) >> 29) & 0x01) //col:218
CR0_CACHE_DISABLE_BIT =                                        30 //col:232
CR0_CACHE_DISABLE_FLAG =                                       0x40000000 //col:233
CR0_CACHE_DISABLE_MASK =                                       0x01 //col:234
CR0_CACHE_DISABLE(_) =                                         (((_) >> 30) & 0x01) //col:235
CR0_PAGING_ENABLE_BIT =                                        31 //col:248
CR0_PAGING_ENABLE_FLAG =                                       0x80000000 //col:249
CR0_PAGING_ENABLE_MASK =                                       0x01 //col:250
CR0_PAGING_ENABLE(_) =                                         (((_) >> 31) & 0x01) //col:251
CR3_PAGE_LEVEL_WRITE_THROUGH_BIT =                             3 //col:273
CR3_PAGE_LEVEL_WRITE_THROUGH_FLAG =                            0x08 //col:274
CR3_PAGE_LEVEL_WRITE_THROUGH_MASK =                            0x01 //col:275
CR3_PAGE_LEVEL_WRITE_THROUGH(_) =                              (((_) >> 3) & 0x01) //col:276
CR3_PAGE_LEVEL_CACHE_DISABLE_BIT =                             4 //col:287
CR3_PAGE_LEVEL_CACHE_DISABLE_FLAG =                            0x10 //col:288
CR3_PAGE_LEVEL_CACHE_DISABLE_MASK =                            0x01 //col:289
CR3_PAGE_LEVEL_CACHE_DISABLE(_) =                              (((_) >> 4) & 0x01) //col:290
CR3_ADDRESS_OF_PAGE_DIRECTORY_BIT =                            12 //col:303
CR3_ADDRESS_OF_PAGE_DIRECTORY_FLAG =                           0xFFFFFFFFF000 //col:304
CR3_ADDRESS_OF_PAGE_DIRECTORY_MASK =                           0xFFFFFFFFF //col:305
CR3_ADDRESS_OF_PAGE_DIRECTORY(_) =                             (((_) >> 12) & 0xFFFFFFFFF) //col:306
CR4_VIRTUAL_MODE_EXTENSIONS_BIT =                              0 //col:330
CR4_VIRTUAL_MODE_EXTENSIONS_FLAG =                             0x01 //col:331
CR4_VIRTUAL_MODE_EXTENSIONS_MASK =                             0x01 //col:332
CR4_VIRTUAL_MODE_EXTENSIONS(_) =                               (((_) >> 0) & 0x01) //col:333
CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_BIT =                    1 //col:344
CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_FLAG =                   0x02 //col:345
CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_MASK =                   0x01 //col:346
CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS(_) =                     (((_) >> 1) & 0x01) //col:347
CR4_TIMESTAMP_DISABLE_BIT =                                    2 //col:357
CR4_TIMESTAMP_DISABLE_FLAG =                                   0x04 //col:358
CR4_TIMESTAMP_DISABLE_MASK =                                   0x01 //col:359
CR4_TIMESTAMP_DISABLE(_) =                                     (((_) >> 2) & 0x01) //col:360
CR4_DEBUGGING_EXTENSIONS_BIT =                                 3 //col:372
CR4_DEBUGGING_EXTENSIONS_FLAG =                                0x08 //col:373
CR4_DEBUGGING_EXTENSIONS_MASK =                                0x01 //col:374
CR4_DEBUGGING_EXTENSIONS(_) =                                  (((_) >> 3) & 0x01) //col:375
CR4_PAGE_SIZE_EXTENSIONS_BIT =                                 4 //col:385
CR4_PAGE_SIZE_EXTENSIONS_FLAG =                                0x10 //col:386
CR4_PAGE_SIZE_EXTENSIONS_MASK =                                0x01 //col:387
CR4_PAGE_SIZE_EXTENSIONS(_) =                                  (((_) >> 4) & 0x01) //col:388
CR4_PHYSICAL_ADDRESS_EXTENSION_BIT =                           5 //col:399
CR4_PHYSICAL_ADDRESS_EXTENSION_FLAG =                          0x20 //col:400
CR4_PHYSICAL_ADDRESS_EXTENSION_MASK =                          0x01 //col:401
CR4_PHYSICAL_ADDRESS_EXTENSION(_) =                            (((_) >> 5) & 0x01) //col:402
CR4_MACHINE_CHECK_ENABLE_BIT =                                 6 //col:412
CR4_MACHINE_CHECK_ENABLE_FLAG =                                0x40 //col:413
CR4_MACHINE_CHECK_ENABLE_MASK =                                0x01 //col:414
CR4_MACHINE_CHECK_ENABLE(_) =                                  (((_) >> 6) & 0x01) //col:415
CR4_PAGE_GLOBAL_ENABLE_BIT =                                   7 //col:430
CR4_PAGE_GLOBAL_ENABLE_FLAG =                                  0x80 //col:431
CR4_PAGE_GLOBAL_ENABLE_MASK =                                  0x01 //col:432
CR4_PAGE_GLOBAL_ENABLE(_) =                                    (((_) >> 7) & 0x01) //col:433
CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_BIT =                8 //col:442
CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_FLAG =               0x100 //col:443
CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_MASK =               0x01 //col:444
CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE(_) =                 (((_) >> 8) & 0x01) //col:445
CR4_OS_FXSAVE_FXRSTOR_SUPPORT_BIT =                            9 //col:468
CR4_OS_FXSAVE_FXRSTOR_SUPPORT_FLAG =                           0x200 //col:469
CR4_OS_FXSAVE_FXRSTOR_SUPPORT_MASK =                           0x01 //col:470
CR4_OS_FXSAVE_FXRSTOR_SUPPORT(_) =                             (((_) >> 9) & 0x01) //col:471
CR4_OS_XMM_EXCEPTION_SUPPORT_BIT =                             10 //col:484
CR4_OS_XMM_EXCEPTION_SUPPORT_FLAG =                            0x400 //col:485
CR4_OS_XMM_EXCEPTION_SUPPORT_MASK =                            0x01 //col:486
CR4_OS_XMM_EXCEPTION_SUPPORT(_) =                              (((_) >> 10) & 0x01) //col:487
CR4_USERMODE_INSTRUCTION_PREVENTION_BIT =                      11 //col:496
CR4_USERMODE_INSTRUCTION_PREVENTION_FLAG =                     0x800 //col:497
CR4_USERMODE_INSTRUCTION_PREVENTION_MASK =                     0x01 //col:498
CR4_USERMODE_INSTRUCTION_PREVENTION(_) =                       (((_) >> 11) & 0x01) //col:499
CR4_LINEAR_ADDRESSES_57_BIT_BIT =                              12 //col:511
CR4_LINEAR_ADDRESSES_57_BIT_FLAG =                             0x1000 //col:512
CR4_LINEAR_ADDRESSES_57_BIT_MASK =                             0x01 //col:513
CR4_LINEAR_ADDRESSES_57_BIT(_) =                               (((_) >> 12) & 0x01) //col:514
CR4_VMX_ENABLE_BIT =                                           13 //col:524
CR4_VMX_ENABLE_FLAG =                                          0x2000 //col:525
CR4_VMX_ENABLE_MASK =                                          0x01 //col:526
CR4_VMX_ENABLE(_) =                                            (((_) >> 13) & 0x01) //col:527
CR4_SMX_ENABLE_BIT =                                           14 //col:537
CR4_SMX_ENABLE_FLAG =                                          0x4000 //col:538
CR4_SMX_ENABLE_MASK =                                          0x01 //col:539
CR4_SMX_ENABLE(_) =                                            (((_) >> 14) & 0x01) //col:540
CR4_FSGSBASE_ENABLE_BIT =                                      16 //col:549
CR4_FSGSBASE_ENABLE_FLAG =                                     0x10000 //col:550
CR4_FSGSBASE_ENABLE_MASK =                                     0x01 //col:551
CR4_FSGSBASE_ENABLE(_) =                                       (((_) >> 16) & 0x01) //col:552
CR4_PCID_ENABLE_BIT =                                          17 //col:562
CR4_PCID_ENABLE_FLAG =                                         0x20000 //col:563
CR4_PCID_ENABLE_MASK =                                         0x01 //col:564
CR4_PCID_ENABLE(_) =                                           (((_) >> 17) & 0x01) //col:565
CR4_OS_XSAVE_BIT =                                             18 //col:581
CR4_OS_XSAVE_FLAG =                                            0x40000 //col:582
CR4_OS_XSAVE_MASK =                                            0x01 //col:583
CR4_OS_XSAVE(_) =                                              (((_) >> 18) & 0x01) //col:584
CR4_KEY_LOCKER_ENABLE_BIT =                                    19 //col:595
CR4_KEY_LOCKER_ENABLE_FLAG =                                   0x80000 //col:596
CR4_KEY_LOCKER_ENABLE_MASK =                                   0x01 //col:597
CR4_KEY_LOCKER_ENABLE(_) =                                     (((_) >> 19) & 0x01) //col:598
CR4_SMEP_ENABLE_BIT =                                          20 //col:608
CR4_SMEP_ENABLE_FLAG =                                         0x100000 //col:609
CR4_SMEP_ENABLE_MASK =                                         0x01 //col:610
CR4_SMEP_ENABLE(_) =                                           (((_) >> 20) & 0x01) //col:611
CR4_SMAP_ENABLE_BIT =                                          21 //col:621
CR4_SMAP_ENABLE_FLAG =                                         0x200000 //col:622
CR4_SMAP_ENABLE_MASK =                                         0x01 //col:623
CR4_SMAP_ENABLE(_) =                                           (((_) >> 21) & 0x01) //col:624
CR4_PROTECTION_KEY_ENABLE_BIT =                                22 //col:634
CR4_PROTECTION_KEY_ENABLE_FLAG =                               0x400000 //col:635
CR4_PROTECTION_KEY_ENABLE_MASK =                               0x01 //col:636
CR4_PROTECTION_KEY_ENABLE(_) =                                 (((_) >> 22) & 0x01) //col:637
CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_BIT =                      23 //col:648
CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_FLAG =                     0x800000 //col:649
CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_MASK =                     0x01 //col:650
CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE(_) =                       (((_) >> 23) & 0x01) //col:651
CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_BIT =            24 //col:661
CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_FLAG =           0x1000000 //col:662
CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_MASK =           0x01 //col:663
CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE(_) =             (((_) >> 24) & 0x01) //col:664
CR8_TASK_PRIORITY_LEVEL_BIT =                                  0 //col:683
CR8_TASK_PRIORITY_LEVEL_FLAG =                                 0x0F //col:684
CR8_TASK_PRIORITY_LEVEL_MASK =                                 0x0F //col:685
CR8_TASK_PRIORITY_LEVEL(_) =                                   (((_) >> 0) & 0x0F) //col:686
CR8_RESERVED_BIT =                                             4 //col:694
CR8_RESERVED_FLAG =                                            0xFFFFFFFFFFFFFFF0 //col:695
CR8_RESERVED_MASK =                                            0xFFFFFFFFFFFFFFF //col:696
CR8_RESERVED(_) =                                              (((_) >> 4) & 0xFFFFFFFFFFFFFFF) //col:697
DR6_BREAKPOINT_CONDITION_BIT =                                 0 //col:742
DR6_BREAKPOINT_CONDITION_FLAG =                                0x0F //col:743
DR6_BREAKPOINT_CONDITION_MASK =                                0x0F //col:744
DR6_BREAKPOINT_CONDITION(_) =                                  (((_) >> 0) & 0x0F) //col:745
DR6_DEBUG_REGISTER_ACCESS_DETECTED_BIT =                       13 //col:757
DR6_DEBUG_REGISTER_ACCESS_DETECTED_FLAG =                      0x2000 //col:758
DR6_DEBUG_REGISTER_ACCESS_DETECTED_MASK =                      0x01 //col:759
DR6_DEBUG_REGISTER_ACCESS_DETECTED(_) =                        (((_) >> 13) & 0x01) //col:760
DR6_SINGLE_INSTRUCTION_BIT =                                   14 //col:770
DR6_SINGLE_INSTRUCTION_FLAG =                                  0x4000 //col:771
DR6_SINGLE_INSTRUCTION_MASK =                                  0x01 //col:772
DR6_SINGLE_INSTRUCTION(_) =                                    (((_) >> 14) & 0x01) //col:773
DR6_TASK_SWITCH_BIT =                                          15 //col:783
DR6_TASK_SWITCH_FLAG =                                         0x8000 //col:784
DR6_TASK_SWITCH_MASK =                                         0x01 //col:785
DR6_TASK_SWITCH(_) =                                           (((_) >> 15) & 0x01) //col:786
DR6_RESTRICTED_TRANSACTIONAL_MEMORY_BIT =                      16 //col:799
DR6_RESTRICTED_TRANSACTIONAL_MEMORY_FLAG =                     0x10000 //col:800
DR6_RESTRICTED_TRANSACTIONAL_MEMORY_MASK =                     0x01 //col:801
DR6_RESTRICTED_TRANSACTIONAL_MEMORY(_) =                       (((_) >> 16) & 0x01) //col:802
DR7_LOCAL_BREAKPOINT_0_BIT =                                   0 //col:821
DR7_LOCAL_BREAKPOINT_0_FLAG =                                  0x01 //col:822
DR7_LOCAL_BREAKPOINT_0_MASK =                                  0x01 //col:823
DR7_LOCAL_BREAKPOINT_0(_) =                                    (((_) >> 0) & 0x01) //col:824
DR7_GLOBAL_BREAKPOINT_0_BIT =                                  1 //col:834
DR7_GLOBAL_BREAKPOINT_0_FLAG =                                 0x02 //col:835
DR7_GLOBAL_BREAKPOINT_0_MASK =                                 0x01 //col:836
DR7_GLOBAL_BREAKPOINT_0(_) =                                   (((_) >> 1) & 0x01) //col:837
DR7_LOCAL_BREAKPOINT_1_BIT =                                   2 //col:839
DR7_LOCAL_BREAKPOINT_1_FLAG =                                  0x04 //col:840
DR7_LOCAL_BREAKPOINT_1_MASK =                                  0x01 //col:841
DR7_LOCAL_BREAKPOINT_1(_) =                                    (((_) >> 2) & 0x01) //col:842
DR7_GLOBAL_BREAKPOINT_1_BIT =                                  3 //col:844
DR7_GLOBAL_BREAKPOINT_1_FLAG =                                 0x08 //col:845
DR7_GLOBAL_BREAKPOINT_1_MASK =                                 0x01 //col:846
DR7_GLOBAL_BREAKPOINT_1(_) =                                   (((_) >> 3) & 0x01) //col:847
DR7_LOCAL_BREAKPOINT_2_BIT =                                   4 //col:849
DR7_LOCAL_BREAKPOINT_2_FLAG =                                  0x10 //col:850
DR7_LOCAL_BREAKPOINT_2_MASK =                                  0x01 //col:851
DR7_LOCAL_BREAKPOINT_2(_) =                                    (((_) >> 4) & 0x01) //col:852
DR7_GLOBAL_BREAKPOINT_2_BIT =                                  5 //col:854
DR7_GLOBAL_BREAKPOINT_2_FLAG =                                 0x20 //col:855
DR7_GLOBAL_BREAKPOINT_2_MASK =                                 0x01 //col:856
DR7_GLOBAL_BREAKPOINT_2(_) =                                   (((_) >> 5) & 0x01) //col:857
DR7_LOCAL_BREAKPOINT_3_BIT =                                   6 //col:859
DR7_LOCAL_BREAKPOINT_3_FLAG =                                  0x40 //col:860
DR7_LOCAL_BREAKPOINT_3_MASK =                                  0x01 //col:861
DR7_LOCAL_BREAKPOINT_3(_) =                                    (((_) >> 6) & 0x01) //col:862
DR7_GLOBAL_BREAKPOINT_3_BIT =                                  7 //col:864
DR7_GLOBAL_BREAKPOINT_3_FLAG =                                 0x80 //col:865
DR7_GLOBAL_BREAKPOINT_3_MASK =                                 0x01 //col:866
DR7_GLOBAL_BREAKPOINT_3(_) =                                   (((_) >> 7) & 0x01) //col:867
DR7_LOCAL_EXACT_BREAKPOINT_BIT =                               8 //col:878
DR7_LOCAL_EXACT_BREAKPOINT_FLAG =                              0x100 //col:879
DR7_LOCAL_EXACT_BREAKPOINT_MASK =                              0x01 //col:880
DR7_LOCAL_EXACT_BREAKPOINT(_) =                                (((_) >> 8) & 0x01) //col:881
DR7_GLOBAL_EXACT_BREAKPOINT_BIT =                              9 //col:883
DR7_GLOBAL_EXACT_BREAKPOINT_FLAG =                             0x200 //col:884
DR7_GLOBAL_EXACT_BREAKPOINT_MASK =                             0x01 //col:885
DR7_GLOBAL_EXACT_BREAKPOINT(_) =                               (((_) >> 9) & 0x01) //col:886
DR7_RESTRICTED_TRANSACTIONAL_MEMORY_BIT =                      11 //col:898
DR7_RESTRICTED_TRANSACTIONAL_MEMORY_FLAG =                     0x800 //col:899
DR7_RESTRICTED_TRANSACTIONAL_MEMORY_MASK =                     0x01 //col:900
DR7_RESTRICTED_TRANSACTIONAL_MEMORY(_) =                       (((_) >> 11) & 0x01) //col:901
DR7_GENERAL_DETECT_BIT =                                       13 //col:916
DR7_GENERAL_DETECT_FLAG =                                      0x2000 //col:917
DR7_GENERAL_DETECT_MASK =                                      0x01 //col:918
DR7_GENERAL_DETECT(_) =                                        (((_) >> 13) & 0x01) //col:919
DR7_READ_WRITE_0_BIT =                                         16 //col:940
DR7_READ_WRITE_0_FLAG =                                        0x30000 //col:941
DR7_READ_WRITE_0_MASK =                                        0x03 //col:942
DR7_READ_WRITE_0(_) =                                          (((_) >> 16) & 0x03) //col:943
DR7_LENGTH_0_BIT =                                             18 //col:960
DR7_LENGTH_0_FLAG =                                            0xC0000 //col:961
DR7_LENGTH_0_MASK =                                            0x03 //col:962
DR7_LENGTH_0(_) =                                              (((_) >> 18) & 0x03) //col:963
DR7_READ_WRITE_1_BIT =                                         20 //col:965
DR7_READ_WRITE_1_FLAG =                                        0x300000 //col:966
DR7_READ_WRITE_1_MASK =                                        0x03 //col:967
DR7_READ_WRITE_1(_) =                                          (((_) >> 20) & 0x03) //col:968
DR7_LENGTH_1_BIT =                                             22 //col:970
DR7_LENGTH_1_FLAG =                                            0xC00000 //col:971
DR7_LENGTH_1_MASK =                                            0x03 //col:972
DR7_LENGTH_1(_) =                                              (((_) >> 22) & 0x03) //col:973
DR7_READ_WRITE_2_BIT =                                         24 //col:975
DR7_READ_WRITE_2_FLAG =                                        0x3000000 //col:976
DR7_READ_WRITE_2_MASK =                                        0x03 //col:977
DR7_READ_WRITE_2(_) =                                          (((_) >> 24) & 0x03) //col:978
DR7_LENGTH_2_BIT =                                             26 //col:980
DR7_LENGTH_2_FLAG =                                            0xC000000 //col:981
DR7_LENGTH_2_MASK =                                            0x03 //col:982
DR7_LENGTH_2(_) =                                              (((_) >> 26) & 0x03) //col:983
DR7_READ_WRITE_3_BIT =                                         28 //col:985
DR7_READ_WRITE_3_FLAG =                                        0x30000000 //col:986
DR7_READ_WRITE_3_MASK =                                        0x03 //col:987
DR7_READ_WRITE_3(_) =                                          (((_) >> 28) & 0x03) //col:988
DR7_LENGTH_3_BIT =                                             30 //col:990
DR7_LENGTH_3_FLAG =                                            0xC0000000 //col:991
DR7_LENGTH_3_MASK =                                            0x03 //col:992
DR7_LENGTH_3(_) =                                              (((_) >> 30) & 0x03) //col:993
CPUID_SIGNATURE =                                              0x00000000 //col:1022
CPUID_VERSION_INFORMATION =                                    0x00000001 //col:1063
CPUID_VERSION_INFORMATION_STEPPING_ID_BIT =                    0 //col:1074
CPUID_VERSION_INFORMATION_STEPPING_ID_FLAG =                   0x0F //col:1075
CPUID_VERSION_INFORMATION_STEPPING_ID_MASK =                   0x0F //col:1076
CPUID_VERSION_INFORMATION_STEPPING_ID(_) =                     (((_) >> 0) & 0x0F) //col:1077
CPUID_VERSION_INFORMATION_MODEL_BIT =                          4 //col:1079
CPUID_VERSION_INFORMATION_MODEL_FLAG =                         0xF0 //col:1080
CPUID_VERSION_INFORMATION_MODEL_MASK =                         0x0F //col:1081
CPUID_VERSION_INFORMATION_MODEL(_) =                           (((_) >> 4) & 0x0F) //col:1082
CPUID_VERSION_INFORMATION_FAMILY_ID_BIT =                      8 //col:1084
CPUID_VERSION_INFORMATION_FAMILY_ID_FLAG =                     0xF00 //col:1085
CPUID_VERSION_INFORMATION_FAMILY_ID_MASK =                     0x0F //col:1086
CPUID_VERSION_INFORMATION_FAMILY_ID(_) =                       (((_) >> 8) & 0x0F) //col:1087
CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_BIT =                 12 //col:1096
CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_FLAG =                0x3000 //col:1097
CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_MASK =                0x03 //col:1098
CPUID_VERSION_INFORMATION_PROCESSOR_TYPE(_) =                  (((_) >> 12) & 0x03) //col:1099
CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_BIT =              16 //col:1106
CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_FLAG =             0xF0000 //col:1107
CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_MASK =             0x0F //col:1108
CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID(_) =               (((_) >> 16) & 0x0F) //col:1109
CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_BIT =             20 //col:1115
CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_FLAG =            0xFF00000 //col:1116
CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_MASK =            0xFF //col:1117
CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID(_) =              (((_) >> 20) & 0xFF) //col:1118
CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_BIT =                 0 //col:1137
CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_FLAG =                0xFF //col:1138
CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_MASK =                0xFF //col:1139
CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX(_) =                  (((_) >> 0) & 0xFF) //col:1140
CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_BIT =           8 //col:1149
CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_FLAG =          0xFF00 //col:1150
CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_MASK =          0xFF //col:1151
CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE(_) =            (((_) >> 8) & 0xFF) //col:1152
CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_BIT =         16 //col:1162
CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_FLAG =        0xFF0000 //col:1163
CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_MASK =        0xFF //col:1164
CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS(_) =          (((_) >> 16) & 0xFF) //col:1165
CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_BIT =             24 //col:1172
CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_FLAG =            0xFF000000 //col:1173
CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_MASK =            0xFF //col:1174
CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID(_) =              (((_) >> 24) & 0xFF) //col:1175
CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_BIT = 0 //col:1194
CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_FLAG = 0x01 //col:1195
CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_MASK = 0x01 //col:1196
CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3(_) = (((_) >> 0) & 0x01) //col:1197
CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_BIT =      1 //col:1205
CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_FLAG =     0x02 //col:1206
CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_MASK =     0x01 //col:1207
CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION(_) =       (((_) >> 1) & 0x01) //col:1208
CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_BIT =       2 //col:1216
CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_FLAG =      0x04 //col:1217
CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_MASK =      0x01 //col:1218
CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT(_) =        (((_) >> 2) & 0x01) //col:1219
CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_BIT =  3 //col:1227
CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_FLAG = 0x08 //col:1228
CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_MASK = 0x01 //col:1229
CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION(_) =   (((_) >> 3) & 0x01) //col:1230
CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_BIT =  4 //col:1239
CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_FLAG = 0x10 //col:1240
CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_MASK = 0x01 //col:1241
CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE(_) =   (((_) >> 4) & 0x01) //col:1242
CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_BIT = 5 //col:1250
CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_FLAG = 0x20 //col:1251
CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_MASK = 0x01 //col:1252
CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS(_) =  (((_) >> 5) & 0x01) //col:1253
CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_BIT =      6 //col:1263
CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_FLAG =     0x40 //col:1264
CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_MASK =     0x01 //col:1265
CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS(_) =       (((_) >> 6) & 0x01) //col:1266
CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_BIT = 7 //col:1274
CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_FLAG = 0x80 //col:1275
CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_MASK = 0x01 //col:1276
CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY(_) = (((_) >> 7) & 0x01) //col:1277
CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_BIT =          8 //col:1285
CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_FLAG =         0x100 //col:1286
CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_MASK =         0x01 //col:1287
CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2(_) =           (((_) >> 8) & 0x01) //col:1288
CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_BIT = 9 //col:1297
CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_FLAG = 0x200 //col:1298
CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_MASK = 0x01 //col:1299
CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3(_) = (((_) >> 9) & 0x01) //col:1300
CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_BIT =              10 //col:1310
CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_FLAG =             0x400 //col:1311
CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_MASK =             0x01 //col:1312
CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID(_) =               (((_) >> 10) & 0x01) //col:1313
CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_BIT =              11 //col:1321
CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_FLAG =             0x800 //col:1322
CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_MASK =             0x01 //col:1323
CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG(_) =               (((_) >> 11) & 0x01) //col:1324
CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_BIT =             12 //col:1332
CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_FLAG =            0x1000 //col:1333
CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_MASK =            0x01 //col:1334
CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS(_) =              (((_) >> 12) & 0x01) //col:1335
CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_BIT =     13 //col:1343
CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_FLAG =    0x2000 //col:1344
CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_MASK =    0x01 //col:1345
CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION(_) =      (((_) >> 13) & 0x01) //col:1346
CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_BIT =        14 //col:1354
CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_FLAG =       0x4000 //col:1355
CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_MASK =       0x01 //col:1356
CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL(_) =         (((_) >> 14) & 0x01) //col:1357
CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_BIT = 15 //col:1366
CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_FLAG = 0x8000 //col:1367
CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_MASK = 0x01 //col:1368
CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY(_) = (((_) >> 15) & 0x01) //col:1369
CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_BIT = 17 //col:1378
CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_FLAG = 0x20000 //col:1379
CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_MASK = 0x01 //col:1380
CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS(_) = (((_) >> 17) & 0x01) //col:1381
CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_BIT =        18 //col:1390
CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_FLAG =       0x40000 //col:1391
CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_MASK =       0x01 //col:1392
CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS(_) =         (((_) >> 18) & 0x01) //col:1393
CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_BIT =              19 //col:1401
CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_FLAG =             0x80000 //col:1402
CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_MASK =             0x01 //col:1403
CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT(_) =               (((_) >> 19) & 0x01) //col:1404
CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_BIT =              20 //col:1412
CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_FLAG =             0x100000 //col:1413
CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_MASK =             0x01 //col:1414
CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT(_) =               (((_) >> 20) & 0x01) //col:1415
CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_BIT =             21 //col:1423
CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_FLAG =            0x200000 //col:1424
CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_MASK =            0x01 //col:1425
CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT(_) =              (((_) >> 21) & 0x01) //col:1426
CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_BIT =          22 //col:1434
CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_FLAG =         0x400000 //col:1435
CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_MASK =         0x01 //col:1436
CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION(_) =           (((_) >> 22) & 0x01) //col:1437
CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_BIT =         23 //col:1445
CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_FLAG =        0x800000 //col:1446
CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_MASK =        0x01 //col:1447
CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION(_) =          (((_) >> 23) & 0x01) //col:1448
CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_BIT =               24 //col:1457
CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_FLAG =              0x1000000 //col:1458
CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_MASK =              0x01 //col:1459
CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE(_) =                (((_) >> 24) & 0x01) //col:1460
CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_BIT = 25 //col:1468
CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_FLAG = 0x2000000 //col:1469
CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_MASK = 0x01 //col:1470
CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS(_) = (((_) >> 25) & 0x01) //col:1471
CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_BIT =   26 //col:1480
CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_FLAG =  0x4000000 //col:1481
CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_MASK =  0x01 //col:1482
CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION(_) =    (((_) >> 26) & 0x01) //col:1483
CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_BIT =                   27 //col:1492
CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_FLAG =                  0x8000000 //col:1493
CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_MASK =                  0x01 //col:1494
CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE(_) =                    (((_) >> 27) & 0x01) //col:1495
CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_BIT =                28 //col:1503
CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_FLAG =               0x10000000 //col:1504
CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_MASK =               0x01 //col:1505
CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT(_) =                 (((_) >> 28) & 0x01) //col:1506
CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_BIT = 29 //col:1514
CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_FLAG = 0x20000000 //col:1515
CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_MASK = 0x01 //col:1516
CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS(_) = (((_) >> 29) & 0x01) //col:1517
CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_BIT =         30 //col:1525
CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_FLAG =        0x40000000 //col:1526
CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_MASK =        0x01 //col:1527
CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION(_) =          (((_) >> 30) & 0x01) //col:1528
CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_BIT = 0 //col:1548
CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_FLAG = 0x01 //col:1549
CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_MASK = 0x01 //col:1550
CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP(_) = (((_) >> 0) & 0x01) //col:1551
CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_BIT = 1 //col:1561
CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_FLAG = 0x02 //col:1562
CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_MASK = 0x01 //col:1563
CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS(_) = (((_) >> 1) & 0x01) //col:1564
CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_BIT =       2 //col:1573
CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_FLAG =      0x04 //col:1574
CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_MASK =      0x01 //col:1575
CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS(_) =        (((_) >> 2) & 0x01) //col:1576
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_BIT =        3 //col:1585
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_FLAG =       0x08 //col:1586
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_MASK =       0x01 //col:1587
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION(_) =         (((_) >> 3) & 0x01) //col:1588
CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_BIT =          4 //col:1596
CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_FLAG =         0x10 //col:1597
CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_MASK =         0x01 //col:1598
CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER(_) =           (((_) >> 4) & 0x01) //col:1599
CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_BIT =   5 //col:1607
CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_FLAG =  0x20 //col:1608
CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_MASK =  0x01 //col:1609
CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS(_) =    (((_) >> 5) & 0x01) //col:1610
CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_BIT = 6 //col:1619
CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_FLAG = 0x40 //col:1620
CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_MASK = 0x01 //col:1621
CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION(_) =  (((_) >> 6) & 0x01) //col:1622
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_BIT =    7 //col:1633
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_FLAG =   0x80 //col:1634
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_MASK =   0x01 //col:1635
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION(_) =     (((_) >> 7) & 0x01) //col:1636
CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_BIT =                  8 //col:1644
CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_FLAG =                 0x100 //col:1645
CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_MASK =                 0x01 //col:1646
CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B(_) =                   (((_) >> 8) & 0x01) //col:1647
CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_BIT =               9 //col:1657
CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_FLAG =              0x200 //col:1658
CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_MASK =              0x01 //col:1659
CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP(_) =                (((_) >> 9) & 0x01) //col:1660
CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_BIT = 11 //col:1669
CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_FLAG = 0x800 //col:1670
CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_MASK = 0x01 //col:1671
CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS(_) = (((_) >> 11) & 0x01) //col:1672
CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_BIT = 12 //col:1681
CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_FLAG = 0x1000 //col:1682
CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_MASK = 0x01 //col:1683
CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS(_) = (((_) >> 12) & 0x01) //col:1684
CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_BIT =            13 //col:1693
CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_FLAG =           0x2000 //col:1694
CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_MASK =           0x01 //col:1695
CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT(_) =             (((_) >> 13) & 0x01) //col:1696
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_BIT = 14 //col:1705
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_FLAG = 0x4000 //col:1706
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_MASK = 0x01 //col:1707
CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE(_) =  (((_) >> 14) & 0x01) //col:1708
CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_BIT = 15 //col:1717
CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_FLAG = 0x8000 //col:1718
CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_MASK = 0x01 //col:1719
CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS(_) = (((_) >> 15) & 0x01) //col:1720
CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_BIT =       16 //col:1729
CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_FLAG =      0x10000 //col:1730
CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_MASK =      0x01 //col:1731
CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE(_) =        (((_) >> 16) & 0x01) //col:1732
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_BIT =  17 //col:1742
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_FLAG = 0x20000 //col:1743
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_MASK = 0x01 //col:1744
CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT(_) =   (((_) >> 17) & 0x01) //col:1745
CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_BIT =    18 //col:1753
CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_FLAG =   0x40000 //col:1754
CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_MASK =   0x01 //col:1755
CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER(_) =     (((_) >> 18) & 0x01) //col:1756
CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_BIT =                    19 //col:1764
CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_FLAG =                   0x80000 //col:1765
CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_MASK =                   0x01 //col:1766
CPUID_FEATURE_INFORMATION_EDX_CLFLUSH(_) =                     (((_) >> 19) & 0x01) //col:1767
CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_BIT =                21 //col:1779
CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_FLAG =               0x200000 //col:1780
CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_MASK =               0x01 //col:1781
CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE(_) =                 (((_) >> 21) & 0x01) //col:1782
CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_BIT = 22 //col:1791
CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_FLAG = 0x400000 //col:1792
CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_MASK = 0x01 //col:1793
CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI(_) = (((_) >> 22) & 0x01) //col:1794
CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_BIT =                23 //col:1802
CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_FLAG =               0x800000 //col:1803
CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_MASK =               0x01 //col:1804
CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT(_) =                 (((_) >> 23) & 0x01) //col:1805
CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_BIT = 24 //col:1815
CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_FLAG = 0x1000000 //col:1816
CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_MASK = 0x01 //col:1817
CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS(_) = (((_) >> 24) & 0x01) //col:1818
CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_BIT =                25 //col:1826
CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_FLAG =               0x2000000 //col:1827
CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_MASK =               0x01 //col:1828
CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT(_) =                 (((_) >> 25) & 0x01) //col:1829
CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_BIT =               26 //col:1837
CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_FLAG =              0x4000000 //col:1838
CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_MASK =              0x01 //col:1839
CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT(_) =                (((_) >> 26) & 0x01) //col:1840
CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_BIT =                 27 //col:1849
CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_FLAG =                0x8000000 //col:1850
CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_MASK =                0x01 //col:1851
CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP(_) =                  (((_) >> 27) & 0x01) //col:1852
CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_BIT = 28 //col:1862
CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_FLAG = 0x10000000 //col:1863
CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_MASK = 0x01 //col:1864
CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY(_) =  (((_) >> 28) & 0x01) //col:1865
CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_BIT =            29 //col:1873
CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_FLAG =           0x20000000 //col:1874
CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_MASK =           0x01 //col:1875
CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR(_) =             (((_) >> 29) & 0x01) //col:1876
CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_BIT =       31 //col:1887
CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_FLAG =      0x80000000 //col:1888
CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_MASK =      0x01 //col:1889
CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE(_) =        (((_) >> 31) & 0x01) //col:1890
CPUID_CACHE_PARAMETERS =                                       0x00000004 //col:1917
CPUID_EAX_CACHE_TYPE_FIELD_BIT =                               0 //col:1932
CPUID_EAX_CACHE_TYPE_FIELD_FLAG =                              0x1F //col:1933
CPUID_EAX_CACHE_TYPE_FIELD_MASK =                              0x1F //col:1934
CPUID_EAX_CACHE_TYPE_FIELD(_) =                                (((_) >> 0) & 0x1F) //col:1935
CPUID_EAX_CACHE_LEVEL_BIT =                                    5 //col:1941
CPUID_EAX_CACHE_LEVEL_FLAG =                                   0xE0 //col:1942
CPUID_EAX_CACHE_LEVEL_MASK =                                   0x07 //col:1943
CPUID_EAX_CACHE_LEVEL(_) =                                     (((_) >> 5) & 0x07) //col:1944
CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_BIT =                  8 //col:1950
CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_FLAG =                 0x100 //col:1951
CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_MASK =                 0x01 //col:1952
CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL(_) =                   (((_) >> 8) & 0x01) //col:1953
CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_BIT =                        9 //col:1959
CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_FLAG =                       0x200 //col:1960
CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_MASK =                       0x01 //col:1961
CPUID_EAX_FULLY_ASSOCIATIVE_CACHE(_) =                         (((_) >> 9) & 0x01) //col:1962
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_BIT = 14 //col:1973
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_FLAG = 0x3FFC000 //col:1974
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_MASK = 0xFFF //col:1975
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE(_) = (((_) >> 14) & 0xFFF) //col:1976
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_BIT = 26 //col:1987
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_FLAG = 0xFC000000 //col:1988
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_MASK = 0x3F //col:1989
CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE(_) = (((_) >> 26) & 0x3F) //col:1990
CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_BIT =                     0 //col:2006
CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_FLAG =                    0xFFF //col:2007
CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_MASK =                    0xFFF //col:2008
CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE(_) =                      (((_) >> 0) & 0xFFF) //col:2009
CPUID_EBX_PHYSICAL_LINE_PARTITIONS_BIT =                       12 //col:2017
CPUID_EBX_PHYSICAL_LINE_PARTITIONS_FLAG =                      0x3FF000 //col:2018
CPUID_EBX_PHYSICAL_LINE_PARTITIONS_MASK =                      0x3FF //col:2019
CPUID_EBX_PHYSICAL_LINE_PARTITIONS(_) =                        (((_) >> 12) & 0x3FF) //col:2020
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_BIT =                          22 //col:2028
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_FLAG =                         0xFFC00000 //col:2029
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_MASK =                         0x3FF //col:2030
CPUID_EBX_WAYS_OF_ASSOCIATIVITY(_) =                           (((_) >> 22) & 0x3FF) //col:2031
CPUID_ECX_NUMBER_OF_SETS_BIT =                                 0 //col:2047
CPUID_ECX_NUMBER_OF_SETS_FLAG =                                0xFFFFFFFF //col:2048
CPUID_ECX_NUMBER_OF_SETS_MASK =                                0xFFFFFFFF //col:2049
CPUID_ECX_NUMBER_OF_SETS(_) =                                  (((_) >> 0) & 0xFFFFFFFF) //col:2050
CPUID_EDX_WRITE_BACK_INVALIDATE_BIT =                          0 //col:2067
CPUID_EDX_WRITE_BACK_INVALIDATE_FLAG =                         0x01 //col:2068
CPUID_EDX_WRITE_BACK_INVALIDATE_MASK =                         0x01 //col:2069
CPUID_EDX_WRITE_BACK_INVALIDATE(_) =                           (((_) >> 0) & 0x01) //col:2070
CPUID_EDX_CACHE_INCLUSIVENESS_BIT =                            1 //col:2079
CPUID_EDX_CACHE_INCLUSIVENESS_FLAG =                           0x02 //col:2080
CPUID_EDX_CACHE_INCLUSIVENESS_MASK =                           0x01 //col:2081
CPUID_EDX_CACHE_INCLUSIVENESS(_) =                             (((_) >> 1) & 0x01) //col:2082
CPUID_EDX_COMPLEX_CACHE_INDEXING_BIT =                         2 //col:2091
CPUID_EDX_COMPLEX_CACHE_INDEXING_FLAG =                        0x04 //col:2092
CPUID_EDX_COMPLEX_CACHE_INDEXING_MASK =                        0x01 //col:2093
CPUID_EDX_COMPLEX_CACHE_INDEXING(_) =                          (((_) >> 2) & 0x01) //col:2094
CPUID_MONITOR_MWAIT =                                          0x00000005 //col:2111
CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_BIT =                     0 //col:2122
CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_FLAG =                    0xFFFF //col:2123
CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_MASK =                    0xFFFF //col:2124
CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE(_) =                      (((_) >> 0) & 0xFFFF) //col:2125
CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_BIT =                      0 //col:2140
CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_FLAG =                     0xFFFF //col:2141
CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_MASK =                     0xFFFF //col:2142
CPUID_EBX_LARGEST_MONITOR_LINE_SIZE(_) =                       (((_) >> 0) & 0xFFFF) //col:2143
CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_BIT =        0 //col:2158
CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_FLAG =       0x01 //col:2159
CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_MASK =       0x01 //col:2160
CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS(_) =         (((_) >> 0) & 0x01) //col:2161
CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_BIT = 1 //col:2167
CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_FLAG = 0x02 //col:2168
CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_MASK = 0x01 //col:2169
CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT(_) = (((_) >> 1) & 0x01) //col:2170
CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_BIT =                      0 //col:2185
CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_FLAG =                     0x0F //col:2186
CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_MASK =                     0x0F //col:2187
CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES(_) =                       (((_) >> 0) & 0x0F) //col:2188
CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_BIT =                      4 //col:2194
CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_FLAG =                     0xF0 //col:2195
CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_MASK =                     0x0F //col:2196
CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES(_) =                       (((_) >> 4) & 0x0F) //col:2197
CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_BIT =                      8 //col:2203
CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_FLAG =                     0xF00 //col:2204
CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_MASK =                     0x0F //col:2205
CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES(_) =                       (((_) >> 8) & 0x0F) //col:2206
CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_BIT =                      12 //col:2212
CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_FLAG =                     0xF000 //col:2213
CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_MASK =                     0x0F //col:2214
CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES(_) =                       (((_) >> 12) & 0x0F) //col:2215
CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_BIT =                      16 //col:2221
CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_FLAG =                     0xF0000 //col:2222
CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_MASK =                     0x0F //col:2223
CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES(_) =                       (((_) >> 16) & 0x0F) //col:2224
CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_BIT =                      20 //col:2230
CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_FLAG =                     0xF00000 //col:2231
CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_MASK =                     0x0F //col:2232
CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES(_) =                       (((_) >> 20) & 0x0F) //col:2233
CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_BIT =                      24 //col:2239
CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_FLAG =                     0xF000000 //col:2240
CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_MASK =                     0x0F //col:2241
CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES(_) =                       (((_) >> 24) & 0x0F) //col:2242
CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_BIT =                      28 //col:2248
CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_FLAG =                     0xF0000000 //col:2249
CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_MASK =                     0x0F //col:2250
CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES(_) =                       (((_) >> 28) & 0x0F) //col:2251
CPUID_THERMAL_AND_POWER_MANAGEMENT =                           0x00000006 //col:2265
CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_BIT =                   0 //col:2276
CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_FLAG =                  0x01 //col:2277
CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_MASK =                  0x01 //col:2278
CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED(_) =                    (((_) >> 0) & 0x01) //col:2279
CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_BIT =         1 //col:2285
CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_FLAG =        0x02 //col:2286
CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_MASK =        0x01 //col:2287
CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE(_) =          (((_) >> 1) & 0x01) //col:2288
CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_BIT =                      2 //col:2294
CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_FLAG =                     0x04 //col:2295
CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_MASK =                     0x01 //col:2296
CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING(_) =                       (((_) >> 2) & 0x01) //col:2297
CPUID_EAX_POWER_LIMIT_NOTIFICATION_BIT =                       4 //col:2304
CPUID_EAX_POWER_LIMIT_NOTIFICATION_FLAG =                      0x10 //col:2305
CPUID_EAX_POWER_LIMIT_NOTIFICATION_MASK =                      0x01 //col:2306
CPUID_EAX_POWER_LIMIT_NOTIFICATION(_) =                        (((_) >> 4) & 0x01) //col:2307
CPUID_EAX_CLOCK_MODULATION_DUTY_BIT =                          5 //col:2313
CPUID_EAX_CLOCK_MODULATION_DUTY_FLAG =                         0x20 //col:2314
CPUID_EAX_CLOCK_MODULATION_DUTY_MASK =                         0x01 //col:2315
CPUID_EAX_CLOCK_MODULATION_DUTY(_) =                           (((_) >> 5) & 0x01) //col:2316
CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_BIT =                     6 //col:2322
CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_FLAG =                    0x40 //col:2323
CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_MASK =                    0x01 //col:2324
CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT(_) =                      (((_) >> 6) & 0x01) //col:2325
CPUID_EAX_HWP_BASE_REGISTERS_BIT =                             7 //col:2332
CPUID_EAX_HWP_BASE_REGISTERS_FLAG =                            0x80 //col:2333
CPUID_EAX_HWP_BASE_REGISTERS_MASK =                            0x01 //col:2334
CPUID_EAX_HWP_BASE_REGISTERS(_) =                              (((_) >> 7) & 0x01) //col:2335
CPUID_EAX_HWP_NOTIFICATION_BIT =                               8 //col:2341
CPUID_EAX_HWP_NOTIFICATION_FLAG =                              0x100 //col:2342
CPUID_EAX_HWP_NOTIFICATION_MASK =                              0x01 //col:2343
CPUID_EAX_HWP_NOTIFICATION(_) =                                (((_) >> 8) & 0x01) //col:2344
CPUID_EAX_HWP_ACTIVITY_WINDOW_BIT =                            9 //col:2350
CPUID_EAX_HWP_ACTIVITY_WINDOW_FLAG =                           0x200 //col:2351
CPUID_EAX_HWP_ACTIVITY_WINDOW_MASK =                           0x01 //col:2352
CPUID_EAX_HWP_ACTIVITY_WINDOW(_) =                             (((_) >> 9) & 0x01) //col:2353
CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_BIT =              10 //col:2359
CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_FLAG =             0x400 //col:2360
CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_MASK =             0x01 //col:2361
CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE(_) =               (((_) >> 10) & 0x01) //col:2362
CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_BIT =                      11 //col:2368
CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_FLAG =                     0x800 //col:2369
CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_MASK =                     0x01 //col:2370
CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST(_) =                       (((_) >> 11) & 0x01) //col:2371
CPUID_EAX_HDC_BIT =                                            13 //col:2378
CPUID_EAX_HDC_FLAG =                                           0x2000 //col:2379
CPUID_EAX_HDC_MASK =                                           0x01 //col:2380
CPUID_EAX_HDC(_) =                                             (((_) >> 13) & 0x01) //col:2381
CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_BIT =   14 //col:2387
CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_FLAG =  0x4000 //col:2388
CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_MASK =  0x01 //col:2389
CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE(_) =    (((_) >> 14) & 0x01) //col:2390
CPUID_EAX_HWP_CAPABILITIES_BIT =                               15 //col:2396
CPUID_EAX_HWP_CAPABILITIES_FLAG =                              0x8000 //col:2397
CPUID_EAX_HWP_CAPABILITIES_MASK =                              0x01 //col:2398
CPUID_EAX_HWP_CAPABILITIES(_) =                                (((_) >> 15) & 0x01) //col:2399
CPUID_EAX_HWP_PECI_OVERRIDE_BIT =                              16 //col:2405
CPUID_EAX_HWP_PECI_OVERRIDE_FLAG =                             0x10000 //col:2406
CPUID_EAX_HWP_PECI_OVERRIDE_MASK =                             0x01 //col:2407
CPUID_EAX_HWP_PECI_OVERRIDE(_) =                               (((_) >> 16) & 0x01) //col:2408
CPUID_EAX_FLEXIBLE_HWP_BIT =                                   17 //col:2414
CPUID_EAX_FLEXIBLE_HWP_FLAG =                                  0x20000 //col:2415
CPUID_EAX_FLEXIBLE_HWP_MASK =                                  0x01 //col:2416
CPUID_EAX_FLEXIBLE_HWP(_) =                                    (((_) >> 17) & 0x01) //col:2417
CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_BIT =           18 //col:2423
CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_FLAG =          0x40000 //col:2424
CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_MASK =          0x01 //col:2425
CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR(_) =            (((_) >> 18) & 0x01) //col:2426
CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_BIT =    20 //col:2433
CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_FLAG =   0x100000 //col:2434
CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_MASK =   0x01 //col:2435
CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST(_) =     (((_) >> 20) & 0x01) //col:2436
CPUID_EAX_INTEL_THREAD_DIRECTOR_BIT =                          23 //col:2444
CPUID_EAX_INTEL_THREAD_DIRECTOR_FLAG =                         0x800000 //col:2445
CPUID_EAX_INTEL_THREAD_DIRECTOR_MASK =                         0x01 //col:2446
CPUID_EAX_INTEL_THREAD_DIRECTOR(_) =                           (((_) >> 23) & 0x01) //col:2447
CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_BIT = 0 //col:2462
CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_FLAG = 0x0F //col:2463
CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_MASK = 0x0F //col:2464
CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR(_) = (((_) >> 0) & 0x0F) //col:2465
CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_BIT =      0 //col:2482
CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_FLAG =     0x01 //col:2483
CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_MASK =     0x01 //col:2484
CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY(_) =       (((_) >> 0) & 0x01) //col:2485
CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_BIT =        3 //col:2493
CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_FLAG =       0x08 //col:2494
CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_MASK =       0x01 //col:2495
CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES(_) =         (((_) >> 3) & 0x01) //col:2496
CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_BIT =             8 //col:2504
CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_FLAG =            0xFF00 //col:2505
CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_MASK =            0xFF //col:2506
CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE(_) =              (((_) >> 8) & 0xFF) //col:2507
CPUID_EDX_RESERVED_BIT =                                       0 //col:2522
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:2523
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:2524
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:2525
CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS =                      0x00000007 //col:2543
CPUID_EAX_NUMBER_OF_SUB_LEAVES_BIT =                           0 //col:2554
CPUID_EAX_NUMBER_OF_SUB_LEAVES_FLAG =                          0xFFFFFFFF //col:2555
CPUID_EAX_NUMBER_OF_SUB_LEAVES_MASK =                          0xFFFFFFFF //col:2556
CPUID_EAX_NUMBER_OF_SUB_LEAVES(_) =                            (((_) >> 0) & 0xFFFFFFFF) //col:2557
CPUID_EBX_FSGSBASE_BIT =                                       0 //col:2571
CPUID_EBX_FSGSBASE_FLAG =                                      0x01 //col:2572
CPUID_EBX_FSGSBASE_MASK =                                      0x01 //col:2573
CPUID_EBX_FSGSBASE(_) =                                        (((_) >> 0) & 0x01) //col:2574
CPUID_EBX_IA32_TSC_ADJUST_MSR_BIT =                            1 //col:2580
CPUID_EBX_IA32_TSC_ADJUST_MSR_FLAG =                           0x02 //col:2581
CPUID_EBX_IA32_TSC_ADJUST_MSR_MASK =                           0x01 //col:2582
CPUID_EBX_IA32_TSC_ADJUST_MSR(_) =                             (((_) >> 1) & 0x01) //col:2583
CPUID_EBX_SGX_BIT =                                            2 //col:2589
CPUID_EBX_SGX_FLAG =                                           0x04 //col:2590
CPUID_EBX_SGX_MASK =                                           0x01 //col:2591
CPUID_EBX_SGX(_) =                                             (((_) >> 2) & 0x01) //col:2592
CPUID_EBX_BMI1_BIT =                                           3 //col:2598
CPUID_EBX_BMI1_FLAG =                                          0x08 //col:2599
CPUID_EBX_BMI1_MASK =                                          0x01 //col:2600
CPUID_EBX_BMI1(_) =                                            (((_) >> 3) & 0x01) //col:2601
CPUID_EBX_HLE_BIT =                                            4 //col:2607
CPUID_EBX_HLE_FLAG =                                           0x10 //col:2608
CPUID_EBX_HLE_MASK =                                           0x01 //col:2609
CPUID_EBX_HLE(_) =                                             (((_) >> 4) & 0x01) //col:2610
CPUID_EBX_AVX2_BIT =                                           5 //col:2616
CPUID_EBX_AVX2_FLAG =                                          0x20 //col:2617
CPUID_EBX_AVX2_MASK =                                          0x01 //col:2618
CPUID_EBX_AVX2(_) =                                            (((_) >> 5) & 0x01) //col:2619
CPUID_EBX_FDP_EXCPTN_ONLY_BIT =                                6 //col:2625
CPUID_EBX_FDP_EXCPTN_ONLY_FLAG =                               0x40 //col:2626
CPUID_EBX_FDP_EXCPTN_ONLY_MASK =                               0x01 //col:2627
CPUID_EBX_FDP_EXCPTN_ONLY(_) =                                 (((_) >> 6) & 0x01) //col:2628
CPUID_EBX_SMEP_BIT =                                           7 //col:2634
CPUID_EBX_SMEP_FLAG =                                          0x80 //col:2635
CPUID_EBX_SMEP_MASK =                                          0x01 //col:2636
CPUID_EBX_SMEP(_) =                                            (((_) >> 7) & 0x01) //col:2637
CPUID_EBX_BMI2_BIT =                                           8 //col:2643
CPUID_EBX_BMI2_FLAG =                                          0x100 //col:2644
CPUID_EBX_BMI2_MASK =                                          0x01 //col:2645
CPUID_EBX_BMI2(_) =                                            (((_) >> 8) & 0x01) //col:2646
CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_BIT =                       9 //col:2652
CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_FLAG =                      0x200 //col:2653
CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_MASK =                      0x01 //col:2654
CPUID_EBX_ENHANCED_REP_MOVSB_STOSB(_) =                        (((_) >> 9) & 0x01) //col:2655
CPUID_EBX_INVPCID_BIT =                                        10 //col:2661
CPUID_EBX_INVPCID_FLAG =                                       0x400 //col:2662
CPUID_EBX_INVPCID_MASK =                                       0x01 //col:2663
CPUID_EBX_INVPCID(_) =                                         (((_) >> 10) & 0x01) //col:2664
CPUID_EBX_RTM_BIT =                                            11 //col:2670
CPUID_EBX_RTM_FLAG =                                           0x800 //col:2671
CPUID_EBX_RTM_MASK =                                           0x01 //col:2672
CPUID_EBX_RTM(_) =                                             (((_) >> 11) & 0x01) //col:2673
CPUID_EBX_RDT_M_BIT =                                          12 //col:2679
CPUID_EBX_RDT_M_FLAG =                                         0x1000 //col:2680
CPUID_EBX_RDT_M_MASK =                                         0x01 //col:2681
CPUID_EBX_RDT_M(_) =                                           (((_) >> 12) & 0x01) //col:2682
CPUID_EBX_DEPRECATES_BIT =                                     13 //col:2688
CPUID_EBX_DEPRECATES_FLAG =                                    0x2000 //col:2689
CPUID_EBX_DEPRECATES_MASK =                                    0x01 //col:2690
CPUID_EBX_DEPRECATES(_) =                                      (((_) >> 13) & 0x01) //col:2691
CPUID_EBX_MPX_BIT =                                            14 //col:2697
CPUID_EBX_MPX_FLAG =                                           0x4000 //col:2698
CPUID_EBX_MPX_MASK =                                           0x01 //col:2699
CPUID_EBX_MPX(_) =                                             (((_) >> 14) & 0x01) //col:2700
CPUID_EBX_RDT_BIT =                                            15 //col:2706
CPUID_EBX_RDT_FLAG =                                           0x8000 //col:2707
CPUID_EBX_RDT_MASK =                                           0x01 //col:2708
CPUID_EBX_RDT(_) =                                             (((_) >> 15) & 0x01) //col:2709
CPUID_EBX_AVX512F_BIT =                                        16 //col:2715
CPUID_EBX_AVX512F_FLAG =                                       0x10000 //col:2716
CPUID_EBX_AVX512F_MASK =                                       0x01 //col:2717
CPUID_EBX_AVX512F(_) =                                         (((_) >> 16) & 0x01) //col:2718
CPUID_EBX_AVX512DQ_BIT =                                       17 //col:2724
CPUID_EBX_AVX512DQ_FLAG =                                      0x20000 //col:2725
CPUID_EBX_AVX512DQ_MASK =                                      0x01 //col:2726
CPUID_EBX_AVX512DQ(_) =                                        (((_) >> 17) & 0x01) //col:2727
CPUID_EBX_RDSEED_BIT =                                         18 //col:2733
CPUID_EBX_RDSEED_FLAG =                                        0x40000 //col:2734
CPUID_EBX_RDSEED_MASK =                                        0x01 //col:2735
CPUID_EBX_RDSEED(_) =                                          (((_) >> 18) & 0x01) //col:2736
CPUID_EBX_ADX_BIT =                                            19 //col:2742
CPUID_EBX_ADX_FLAG =                                           0x80000 //col:2743
CPUID_EBX_ADX_MASK =                                           0x01 //col:2744
CPUID_EBX_ADX(_) =                                             (((_) >> 19) & 0x01) //col:2745
CPUID_EBX_SMAP_BIT =                                           20 //col:2751
CPUID_EBX_SMAP_FLAG =                                          0x100000 //col:2752
CPUID_EBX_SMAP_MASK =                                          0x01 //col:2753
CPUID_EBX_SMAP(_) =                                            (((_) >> 20) & 0x01) //col:2754
CPUID_EBX_AVX512_IFMA_BIT =                                    21 //col:2760
CPUID_EBX_AVX512_IFMA_FLAG =                                   0x200000 //col:2761
CPUID_EBX_AVX512_IFMA_MASK =                                   0x01 //col:2762
CPUID_EBX_AVX512_IFMA(_) =                                     (((_) >> 21) & 0x01) //col:2763
CPUID_EBX_CLFLUSHOPT_BIT =                                     23 //col:2770
CPUID_EBX_CLFLUSHOPT_FLAG =                                    0x800000 //col:2771
CPUID_EBX_CLFLUSHOPT_MASK =                                    0x01 //col:2772
CPUID_EBX_CLFLUSHOPT(_) =                                      (((_) >> 23) & 0x01) //col:2773
CPUID_EBX_CLWB_BIT =                                           24 //col:2779
CPUID_EBX_CLWB_FLAG =                                          0x1000000 //col:2780
CPUID_EBX_CLWB_MASK =                                          0x01 //col:2781
CPUID_EBX_CLWB(_) =                                            (((_) >> 24) & 0x01) //col:2782
CPUID_EBX_INTEL_BIT =                                          25 //col:2788
CPUID_EBX_INTEL_FLAG =                                         0x2000000 //col:2789
CPUID_EBX_INTEL_MASK =                                         0x01 //col:2790
CPUID_EBX_INTEL(_) =                                           (((_) >> 25) & 0x01) //col:2791
CPUID_EBX_AVX512PF_BIT =                                       26 //col:2797
CPUID_EBX_AVX512PF_FLAG =                                      0x4000000 //col:2798
CPUID_EBX_AVX512PF_MASK =                                      0x01 //col:2799
CPUID_EBX_AVX512PF(_) =                                        (((_) >> 26) & 0x01) //col:2800
CPUID_EBX_AVX512ER_BIT =                                       27 //col:2806
CPUID_EBX_AVX512ER_FLAG =                                      0x8000000 //col:2807
CPUID_EBX_AVX512ER_MASK =                                      0x01 //col:2808
CPUID_EBX_AVX512ER(_) =                                        (((_) >> 27) & 0x01) //col:2809
CPUID_EBX_AVX512CD_BIT =                                       28 //col:2815
CPUID_EBX_AVX512CD_FLAG =                                      0x10000000 //col:2816
CPUID_EBX_AVX512CD_MASK =                                      0x01 //col:2817
CPUID_EBX_AVX512CD(_) =                                        (((_) >> 28) & 0x01) //col:2818
CPUID_EBX_SHA_BIT =                                            29 //col:2824
CPUID_EBX_SHA_FLAG =                                           0x20000000 //col:2825
CPUID_EBX_SHA_MASK =                                           0x01 //col:2826
CPUID_EBX_SHA(_) =                                             (((_) >> 29) & 0x01) //col:2827
CPUID_EBX_AVX512BW_BIT =                                       30 //col:2833
CPUID_EBX_AVX512BW_FLAG =                                      0x40000000 //col:2834
CPUID_EBX_AVX512BW_MASK =                                      0x01 //col:2835
CPUID_EBX_AVX512BW(_) =                                        (((_) >> 30) & 0x01) //col:2836
CPUID_EBX_AVX512VL_BIT =                                       31 //col:2842
CPUID_EBX_AVX512VL_FLAG =                                      0x80000000 //col:2843
CPUID_EBX_AVX512VL_MASK =                                      0x01 //col:2844
CPUID_EBX_AVX512VL(_) =                                        (((_) >> 31) & 0x01) //col:2845
CPUID_ECX_PREFETCHWT1_BIT =                                    0 //col:2859
CPUID_ECX_PREFETCHWT1_FLAG =                                   0x01 //col:2860
CPUID_ECX_PREFETCHWT1_MASK =                                   0x01 //col:2861
CPUID_ECX_PREFETCHWT1(_) =                                     (((_) >> 0) & 0x01) //col:2862
CPUID_ECX_AVX512_VBMI_BIT =                                    1 //col:2868
CPUID_ECX_AVX512_VBMI_FLAG =                                   0x02 //col:2869
CPUID_ECX_AVX512_VBMI_MASK =                                   0x01 //col:2870
CPUID_ECX_AVX512_VBMI(_) =                                     (((_) >> 1) & 0x01) //col:2871
CPUID_ECX_UMIP_BIT =                                           2 //col:2877
CPUID_ECX_UMIP_FLAG =                                          0x04 //col:2878
CPUID_ECX_UMIP_MASK =                                          0x01 //col:2879
CPUID_ECX_UMIP(_) =                                            (((_) >> 2) & 0x01) //col:2880
CPUID_ECX_PKU_BIT =                                            3 //col:2886
CPUID_ECX_PKU_FLAG =                                           0x08 //col:2887
CPUID_ECX_PKU_MASK =                                           0x01 //col:2888
CPUID_ECX_PKU(_) =                                             (((_) >> 3) & 0x01) //col:2889
CPUID_ECX_OSPKE_BIT =                                          4 //col:2895
CPUID_ECX_OSPKE_FLAG =                                         0x10 //col:2896
CPUID_ECX_OSPKE_MASK =                                         0x01 //col:2897
CPUID_ECX_OSPKE(_) =                                           (((_) >> 4) & 0x01) //col:2898
CPUID_ECX_WAITPKG_BIT =                                        5 //col:2904
CPUID_ECX_WAITPKG_FLAG =                                       0x20 //col:2905
CPUID_ECX_WAITPKG_MASK =                                       0x01 //col:2906
CPUID_ECX_WAITPKG(_) =                                         (((_) >> 5) & 0x01) //col:2907
CPUID_ECX_AVX512_VBMI2_BIT =                                   6 //col:2913
CPUID_ECX_AVX512_VBMI2_FLAG =                                  0x40 //col:2914
CPUID_ECX_AVX512_VBMI2_MASK =                                  0x01 //col:2915
CPUID_ECX_AVX512_VBMI2(_) =                                    (((_) >> 6) & 0x01) //col:2916
CPUID_ECX_CET_SS_BIT =                                         7 //col:2924
CPUID_ECX_CET_SS_FLAG =                                        0x80 //col:2925
CPUID_ECX_CET_SS_MASK =                                        0x01 //col:2926
CPUID_ECX_CET_SS(_) =                                          (((_) >> 7) & 0x01) //col:2927
CPUID_ECX_GFNI_BIT =                                           8 //col:2933
CPUID_ECX_GFNI_FLAG =                                          0x100 //col:2934
CPUID_ECX_GFNI_MASK =                                          0x01 //col:2935
CPUID_ECX_GFNI(_) =                                            (((_) >> 8) & 0x01) //col:2936
CPUID_ECX_VAES_BIT =                                           9 //col:2942
CPUID_ECX_VAES_FLAG =                                          0x200 //col:2943
CPUID_ECX_VAES_MASK =                                          0x01 //col:2944
CPUID_ECX_VAES(_) =                                            (((_) >> 9) & 0x01) //col:2945
CPUID_ECX_VPCLMULQDQ_BIT =                                     10 //col:2951
CPUID_ECX_VPCLMULQDQ_FLAG =                                    0x400 //col:2952
CPUID_ECX_VPCLMULQDQ_MASK =                                    0x01 //col:2953
CPUID_ECX_VPCLMULQDQ(_) =                                      (((_) >> 10) & 0x01) //col:2954
CPUID_ECX_AVX512_VNNI_BIT =                                    11 //col:2960
CPUID_ECX_AVX512_VNNI_FLAG =                                   0x800 //col:2961
CPUID_ECX_AVX512_VNNI_MASK =                                   0x01 //col:2962
CPUID_ECX_AVX512_VNNI(_) =                                     (((_) >> 11) & 0x01) //col:2963
CPUID_ECX_AVX512_BITALG_BIT =                                  12 //col:2969
CPUID_ECX_AVX512_BITALG_FLAG =                                 0x1000 //col:2970
CPUID_ECX_AVX512_BITALG_MASK =                                 0x01 //col:2971
CPUID_ECX_AVX512_BITALG(_) =                                   (((_) >> 12) & 0x01) //col:2972
CPUID_ECX_TME_EN_BIT =                                         13 //col:2979
CPUID_ECX_TME_EN_FLAG =                                        0x2000 //col:2980
CPUID_ECX_TME_EN_MASK =                                        0x01 //col:2981
CPUID_ECX_TME_EN(_) =                                          (((_) >> 13) & 0x01) //col:2982
CPUID_ECX_AVX512_VPOPCNTDQ_BIT =                               14 //col:2988
CPUID_ECX_AVX512_VPOPCNTDQ_FLAG =                              0x4000 //col:2989
CPUID_ECX_AVX512_VPOPCNTDQ_MASK =                              0x01 //col:2990
CPUID_ECX_AVX512_VPOPCNTDQ(_) =                                (((_) >> 14) & 0x01) //col:2991
CPUID_ECX_LA57_BIT =                                           16 //col:2998
CPUID_ECX_LA57_FLAG =                                          0x10000 //col:2999
CPUID_ECX_LA57_MASK =                                          0x01 //col:3000
CPUID_ECX_LA57(_) =                                            (((_) >> 16) & 0x01) //col:3001
CPUID_ECX_MAWAU_BIT =                                          17 //col:3007
CPUID_ECX_MAWAU_FLAG =                                         0x3E0000 //col:3008
CPUID_ECX_MAWAU_MASK =                                         0x1F //col:3009
CPUID_ECX_MAWAU(_) =                                           (((_) >> 17) & 0x1F) //col:3010
CPUID_ECX_RDPID_BIT =                                          22 //col:3016
CPUID_ECX_RDPID_FLAG =                                         0x400000 //col:3017
CPUID_ECX_RDPID_MASK =                                         0x01 //col:3018
CPUID_ECX_RDPID(_) =                                           (((_) >> 22) & 0x01) //col:3019
CPUID_ECX_KL_BIT =                                             23 //col:3025
CPUID_ECX_KL_FLAG =                                            0x800000 //col:3026
CPUID_ECX_KL_MASK =                                            0x01 //col:3027
CPUID_ECX_KL(_) =                                              (((_) >> 23) & 0x01) //col:3028
CPUID_ECX_CLDEMOTE_BIT =                                       25 //col:3035
CPUID_ECX_CLDEMOTE_FLAG =                                      0x2000000 //col:3036
CPUID_ECX_CLDEMOTE_MASK =                                      0x01 //col:3037
CPUID_ECX_CLDEMOTE(_) =                                        (((_) >> 25) & 0x01) //col:3038
CPUID_ECX_MOVDIRI_BIT =                                        27 //col:3045
CPUID_ECX_MOVDIRI_FLAG =                                       0x8000000 //col:3046
CPUID_ECX_MOVDIRI_MASK =                                       0x01 //col:3047
CPUID_ECX_MOVDIRI(_) =                                         (((_) >> 27) & 0x01) //col:3048
CPUID_ECX_MOVDIR64B_BIT =                                      28 //col:3054
CPUID_ECX_MOVDIR64B_FLAG =                                     0x10000000 //col:3055
CPUID_ECX_MOVDIR64B_MASK =                                     0x01 //col:3056
CPUID_ECX_MOVDIR64B(_) =                                       (((_) >> 28) & 0x01) //col:3057
CPUID_ECX_SGX_LC_BIT =                                         30 //col:3064
CPUID_ECX_SGX_LC_FLAG =                                        0x40000000 //col:3065
CPUID_ECX_SGX_LC_MASK =                                        0x01 //col:3066
CPUID_ECX_SGX_LC(_) =                                          (((_) >> 30) & 0x01) //col:3067
CPUID_ECX_PKS_BIT =                                            31 //col:3073
CPUID_ECX_PKS_FLAG =                                           0x80000000 //col:3074
CPUID_ECX_PKS_MASK =                                           0x01 //col:3075
CPUID_ECX_PKS(_) =                                             (((_) >> 31) & 0x01) //col:3076
CPUID_EDX_AVX512_4VNNIW_BIT =                                  2 //col:3092
CPUID_EDX_AVX512_4VNNIW_FLAG =                                 0x04 //col:3093
CPUID_EDX_AVX512_4VNNIW_MASK =                                 0x01 //col:3094
CPUID_EDX_AVX512_4VNNIW(_) =                                   (((_) >> 2) & 0x01) //col:3095
CPUID_EDX_AVX512_4FMAPS_BIT =                                  3 //col:3101
CPUID_EDX_AVX512_4FMAPS_FLAG =                                 0x08 //col:3102
CPUID_EDX_AVX512_4FMAPS_MASK =                                 0x01 //col:3103
CPUID_EDX_AVX512_4FMAPS(_) =                                   (((_) >> 3) & 0x01) //col:3104
CPUID_EDX_FAST_SHORT_REP_MOV_BIT =                             4 //col:3110
CPUID_EDX_FAST_SHORT_REP_MOV_FLAG =                            0x10 //col:3111
CPUID_EDX_FAST_SHORT_REP_MOV_MASK =                            0x01 //col:3112
CPUID_EDX_FAST_SHORT_REP_MOV(_) =                              (((_) >> 4) & 0x01) //col:3113
CPUID_EDX_AVX512_VP2INTERSECT_BIT =                            8 //col:3120
CPUID_EDX_AVX512_VP2INTERSECT_FLAG =                           0x100 //col:3121
CPUID_EDX_AVX512_VP2INTERSECT_MASK =                           0x01 //col:3122
CPUID_EDX_AVX512_VP2INTERSECT(_) =                             (((_) >> 8) & 0x01) //col:3123
CPUID_EDX_MD_CLEAR_BIT =                                       10 //col:3130
CPUID_EDX_MD_CLEAR_FLAG =                                      0x400 //col:3131
CPUID_EDX_MD_CLEAR_MASK =                                      0x01 //col:3132
CPUID_EDX_MD_CLEAR(_) =                                        (((_) >> 10) & 0x01) //col:3133
CPUID_EDX_SERIALIZE_BIT =                                      14 //col:3140
CPUID_EDX_SERIALIZE_FLAG =                                     0x4000 //col:3141
CPUID_EDX_SERIALIZE_MASK =                                     0x01 //col:3142
CPUID_EDX_SERIALIZE(_) =                                       (((_) >> 14) & 0x01) //col:3143
CPUID_EDX_HYBRID_BIT =                                         15 //col:3149
CPUID_EDX_HYBRID_FLAG =                                        0x8000 //col:3150
CPUID_EDX_HYBRID_MASK =                                        0x01 //col:3151
CPUID_EDX_HYBRID(_) =                                          (((_) >> 15) & 0x01) //col:3152
CPUID_EDX_PCONFIG_BIT =                                        18 //col:3159
CPUID_EDX_PCONFIG_FLAG =                                       0x40000 //col:3160
CPUID_EDX_PCONFIG_MASK =                                       0x01 //col:3161
CPUID_EDX_PCONFIG(_) =                                         (((_) >> 18) & 0x01) //col:3162
CPUID_EDX_CET_IBT_BIT =                                        20 //col:3170
CPUID_EDX_CET_IBT_FLAG =                                       0x100000 //col:3171
CPUID_EDX_CET_IBT_MASK =                                       0x01 //col:3172
CPUID_EDX_CET_IBT(_) =                                         (((_) >> 20) & 0x01) //col:3173
CPUID_EDX_IBRS_IBPB_BIT =                                      26 //col:3182
CPUID_EDX_IBRS_IBPB_FLAG =                                     0x4000000 //col:3183
CPUID_EDX_IBRS_IBPB_MASK =                                     0x01 //col:3184
CPUID_EDX_IBRS_IBPB(_) =                                       (((_) >> 26) & 0x01) //col:3185
CPUID_EDX_STIBP_BIT =                                          27 //col:3192
CPUID_EDX_STIBP_FLAG =                                         0x8000000 //col:3193
CPUID_EDX_STIBP_MASK =                                         0x01 //col:3194
CPUID_EDX_STIBP(_) =                                           (((_) >> 27) & 0x01) //col:3195
CPUID_EDX_L1D_FLUSH_BIT =                                      28 //col:3202
CPUID_EDX_L1D_FLUSH_FLAG =                                     0x10000000 //col:3203
CPUID_EDX_L1D_FLUSH_MASK =                                     0x01 //col:3204
CPUID_EDX_L1D_FLUSH(_) =                                       (((_) >> 28) & 0x01) //col:3205
CPUID_EDX_IA32_ARCH_CAPABILITIES_BIT =                         29 //col:3211
CPUID_EDX_IA32_ARCH_CAPABILITIES_FLAG =                        0x20000000 //col:3212
CPUID_EDX_IA32_ARCH_CAPABILITIES_MASK =                        0x01 //col:3213
CPUID_EDX_IA32_ARCH_CAPABILITIES(_) =                          (((_) >> 29) & 0x01) //col:3214
CPUID_EDX_IA32_CORE_CAPABILITIES_BIT =                         30 //col:3220
CPUID_EDX_IA32_CORE_CAPABILITIES_FLAG =                        0x40000000 //col:3221
CPUID_EDX_IA32_CORE_CAPABILITIES_MASK =                        0x01 //col:3222
CPUID_EDX_IA32_CORE_CAPABILITIES(_) =                          (((_) >> 30) & 0x01) //col:3223
CPUID_EDX_SSBD_BIT =                                           31 //col:3230
CPUID_EDX_SSBD_FLAG =                                          0x80000000 //col:3231
CPUID_EDX_SSBD_MASK =                                          0x01 //col:3232
CPUID_EDX_SSBD(_) =                                            (((_) >> 31) & 0x01) //col:3233
CPUID_DIRECT_CACHE_ACCESS_INFORMATION =                        0x00000009 //col:3247
CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT =                          0 //col:3258
CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG =                         0xFFFFFFFF //col:3259
CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK =                         0xFFFFFFFF //col:3260
CPUID_EAX_IA32_PLATFORM_DCA_CAP(_) =                           (((_) >> 0) & 0xFFFFFFFF) //col:3261
CPUID_EBX_RESERVED_BIT =                                       0 //col:3275
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:3276
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:3277
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:3278
CPUID_ECX_RESERVED_BIT =                                       0 //col:3292
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:3293
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:3294
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:3295
CPUID_EDX_RESERVED_BIT =                                       0 //col:3309
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:3310
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:3311
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:3312
CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING =                   0x0000000A //col:3331
CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_BIT = 0 //col:3342
CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_FLAG = 0xFF //col:3343
CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_MASK = 0xFF //col:3344
CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING(_) = (((_) >> 0) & 0xFF) //col:3345
CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_BIT = 8 //col:3351
CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_FLAG = 0xFF00 //col:3352
CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_MASK = 0xFF //col:3353
CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR(_) = (((_) >> 8) & 0xFF) //col:3354
CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_BIT =    16 //col:3360
CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_FLAG =   0xFF0000 //col:3361
CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_MASK =   0xFF //col:3362
CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER(_) =     (((_) >> 16) & 0xFF) //col:3363
CPUID_EAX_EBX_BIT_VECTOR_LENGTH_BIT =                          24 //col:3369
CPUID_EAX_EBX_BIT_VECTOR_LENGTH_FLAG =                         0xFF000000 //col:3370
CPUID_EAX_EBX_BIT_VECTOR_LENGTH_MASK =                         0xFF //col:3371
CPUID_EAX_EBX_BIT_VECTOR_LENGTH(_) =                           (((_) >> 24) & 0xFF) //col:3372
CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_BIT =                 0 //col:3386
CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_FLAG =                0x01 //col:3387
CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_MASK =                0x01 //col:3388
CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE(_) =                  (((_) >> 0) & 0x01) //col:3389
CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_BIT =        1 //col:3395
CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_FLAG =       0x02 //col:3396
CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_MASK =       0x01 //col:3397
CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE(_) =         (((_) >> 1) & 0x01) //col:3398
CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_BIT =           2 //col:3404
CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_FLAG =          0x04 //col:3405
CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_MASK =          0x01 //col:3406
CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE(_) =            (((_) >> 2) & 0x01) //col:3407
CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_BIT = 3 //col:3413
CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_FLAG = 0x08 //col:3414
CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_MASK = 0x01 //col:3415
CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE(_) =  (((_) >> 3) & 0x01) //col:3416
CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_BIT =    4 //col:3422
CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_FLAG =   0x10 //col:3423
CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_MASK =   0x01 //col:3424
CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE(_) =     (((_) >> 4) & 0x01) //col:3425
CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_BIT = 5 //col:3431
CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_FLAG = 0x20 //col:3432
CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_MASK = 0x01 //col:3433
CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE(_) =  (((_) >> 5) & 0x01) //col:3434
CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_BIT =  6 //col:3440
CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_FLAG = 0x40 //col:3441
CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_MASK = 0x01 //col:3442
CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE(_) =   (((_) >> 6) & 0x01) //col:3443
CPUID_ECX_RESERVED_BIT =                                       0 //col:3458
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:3459
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:3460
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:3461
CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_BIT =  0 //col:3475
CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_FLAG = 0x1F //col:3476
CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_MASK = 0x1F //col:3477
CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS(_) =   (((_) >> 0) & 0x1F) //col:3478
CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_BIT = 5 //col:3484
CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_FLAG = 0x1FE0 //col:3485
CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_MASK = 0xFF //col:3486
CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS(_) = (((_) >> 5) & 0xFF) //col:3487
CPUID_EDX_ANY_THREAD_DEPRECATION_BIT =                         15 //col:3494
CPUID_EDX_ANY_THREAD_DEPRECATION_FLAG =                        0x8000 //col:3495
CPUID_EDX_ANY_THREAD_DEPRECATION_MASK =                        0x01 //col:3496
CPUID_EDX_ANY_THREAD_DEPRECATION(_) =                          (((_) >> 15) & 0x01) //col:3497
CPUID_EXTENDED_TOPOLOGY =                                      0x0000000B //col:3521
CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_BIT =          0 //col:3535
CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_FLAG =         0x1F //col:3536
CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_MASK =         0x1F //col:3537
CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT(_) =           (((_) >> 0) & 0x1F) //col:3538
CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_BIT = 0 //col:3558
CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_FLAG = 0xFFFF //col:3559
CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_MASK = 0xFFFF //col:3560
CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE(_) = (((_) >> 0) & 0xFFFF) //col:3561
CPUID_ECX_LEVEL_NUMBER_BIT =                                   0 //col:3576
CPUID_ECX_LEVEL_NUMBER_FLAG =                                  0xFF //col:3577
CPUID_ECX_LEVEL_NUMBER_MASK =                                  0xFF //col:3578
CPUID_ECX_LEVEL_NUMBER(_) =                                    (((_) >> 0) & 0xFF) //col:3579
CPUID_ECX_LEVEL_TYPE_BIT =                                     8 //col:3592
CPUID_ECX_LEVEL_TYPE_FLAG =                                    0xFF00 //col:3593
CPUID_ECX_LEVEL_TYPE_MASK =                                    0xFF //col:3594
CPUID_ECX_LEVEL_TYPE(_) =                                      (((_) >> 8) & 0xFF) //col:3595
CPUID_EDX_X2APIC_ID_BIT =                                      0 //col:3610
CPUID_EDX_X2APIC_ID_FLAG =                                     0xFFFFFFFF //col:3611
CPUID_EDX_X2APIC_ID_MASK =                                     0xFFFFFFFF //col:3612
CPUID_EDX_X2APIC_ID(_) =                                       (((_) >> 0) & 0xFFFFFFFF) //col:3613
CPUID_EXTENDED_STATE_INFORMATION =                             0x0000000D //col:3636
CPUID_EAX_X87_STATE_BIT =                                      0 //col:3653
CPUID_EAX_X87_STATE_FLAG =                                     0x01 //col:3654
CPUID_EAX_X87_STATE_MASK =                                     0x01 //col:3655
CPUID_EAX_X87_STATE(_) =                                       (((_) >> 0) & 0x01) //col:3656
CPUID_EAX_SSE_STATE_BIT =                                      1 //col:3662
CPUID_EAX_SSE_STATE_FLAG =                                     0x02 //col:3663
CPUID_EAX_SSE_STATE_MASK =                                     0x01 //col:3664
CPUID_EAX_SSE_STATE(_) =                                       (((_) >> 1) & 0x01) //col:3665
CPUID_EAX_AVX_STATE_BIT =                                      2 //col:3671
CPUID_EAX_AVX_STATE_FLAG =                                     0x04 //col:3672
CPUID_EAX_AVX_STATE_MASK =                                     0x01 //col:3673
CPUID_EAX_AVX_STATE(_) =                                       (((_) >> 2) & 0x01) //col:3674
CPUID_EAX_MPX_STATE_BIT =                                      3 //col:3680
CPUID_EAX_MPX_STATE_FLAG =                                     0x18 //col:3681
CPUID_EAX_MPX_STATE_MASK =                                     0x03 //col:3682
CPUID_EAX_MPX_STATE(_) =                                       (((_) >> 3) & 0x03) //col:3683
CPUID_EAX_AVX_512_STATE_BIT =                                  5 //col:3689
CPUID_EAX_AVX_512_STATE_FLAG =                                 0xE0 //col:3690
CPUID_EAX_AVX_512_STATE_MASK =                                 0x07 //col:3691
CPUID_EAX_AVX_512_STATE(_) =                                   (((_) >> 5) & 0x07) //col:3692
CPUID_EAX_USED_FOR_IA32_XSS_1_BIT =                            8 //col:3698
CPUID_EAX_USED_FOR_IA32_XSS_1_FLAG =                           0x100 //col:3699
CPUID_EAX_USED_FOR_IA32_XSS_1_MASK =                           0x01 //col:3700
CPUID_EAX_USED_FOR_IA32_XSS_1(_) =                             (((_) >> 8) & 0x01) //col:3701
CPUID_EAX_PKRU_STATE_BIT =                                     9 //col:3707
CPUID_EAX_PKRU_STATE_FLAG =                                    0x200 //col:3708
CPUID_EAX_PKRU_STATE_MASK =                                    0x01 //col:3709
CPUID_EAX_PKRU_STATE(_) =                                      (((_) >> 9) & 0x01) //col:3710
CPUID_EAX_USED_FOR_IA32_XSS_2_BIT =                            13 //col:3717
CPUID_EAX_USED_FOR_IA32_XSS_2_FLAG =                           0x2000 //col:3718
CPUID_EAX_USED_FOR_IA32_XSS_2_MASK =                           0x01 //col:3719
CPUID_EAX_USED_FOR_IA32_XSS_2(_) =                             (((_) >> 13) & 0x01) //col:3720
CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_BIT =  0 //col:3736
CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_FLAG = 0xFFFFFFFF //col:3737
CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_MASK = 0xFFFFFFFF //col:3738
CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0(_) =   (((_) >> 0) & 0xFFFFFFFF) //col:3739
CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_BIT =             0 //col:3754
CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_FLAG =            0xFFFFFFFF //col:3755
CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_MASK =            0xFFFFFFFF //col:3756
CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA(_) =              (((_) >> 0) & 0xFFFFFFFF) //col:3757
CPUID_EDX_XCR0_SUPPORTED_BITS_BIT =                            0 //col:3771
CPUID_EDX_XCR0_SUPPORTED_BITS_FLAG =                           0xFFFFFFFF //col:3772
CPUID_EDX_XCR0_SUPPORTED_BITS_MASK =                           0xFFFFFFFF //col:3773
CPUID_EDX_XCR0_SUPPORTED_BITS(_) =                             (((_) >> 0) & 0xFFFFFFFF) //col:3774
CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_BIT =           1 //col:3797
CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_FLAG =          0x02 //col:3798
CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_MASK =          0x01 //col:3799
CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR(_) =            (((_) >> 1) & 0x01) //col:3800
CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_BIT =                     2 //col:3806
CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_FLAG =                    0x04 //col:3807
CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_MASK =                    0x01 //col:3808
CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1(_) =                      (((_) >> 2) & 0x01) //col:3809
CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_BIT =             3 //col:3815
CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_FLAG =            0x08 //col:3816
CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_MASK =            0x01 //col:3817
CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS(_) =              (((_) >> 3) & 0x01) //col:3818
CPUID_EBX_SIZE_OF_XSAVE_AREAD_BIT =                            0 //col:3833
CPUID_EBX_SIZE_OF_XSAVE_AREAD_FLAG =                           0xFFFFFFFF //col:3834
CPUID_EBX_SIZE_OF_XSAVE_AREAD_MASK =                           0xFFFFFFFF //col:3835
CPUID_EBX_SIZE_OF_XSAVE_AREAD(_) =                             (((_) >> 0) & 0xFFFFFFFF) //col:3836
CPUID_ECX_USED_FOR_XCR0_1_BIT =                                0 //col:3850
CPUID_ECX_USED_FOR_XCR0_1_FLAG =                               0xFF //col:3851
CPUID_ECX_USED_FOR_XCR0_1_MASK =                               0xFF //col:3852
CPUID_ECX_USED_FOR_XCR0_1(_) =                                 (((_) >> 0) & 0xFF) //col:3853
CPUID_ECX_PT_STATE_BIT =                                       8 //col:3859
CPUID_ECX_PT_STATE_FLAG =                                      0x100 //col:3860
CPUID_ECX_PT_STATE_MASK =                                      0x01 //col:3861
CPUID_ECX_PT_STATE(_) =                                        (((_) >> 8) & 0x01) //col:3862
CPUID_ECX_USED_FOR_XCR0_2_BIT =                                9 //col:3868
CPUID_ECX_USED_FOR_XCR0_2_FLAG =                               0x200 //col:3869
CPUID_ECX_USED_FOR_XCR0_2_MASK =                               0x01 //col:3870
CPUID_ECX_USED_FOR_XCR0_2(_) =                                 (((_) >> 9) & 0x01) //col:3871
CPUID_ECX_CET_USER_STATE_BIT =                                 11 //col:3878
CPUID_ECX_CET_USER_STATE_FLAG =                                0x800 //col:3879
CPUID_ECX_CET_USER_STATE_MASK =                                0x01 //col:3880
CPUID_ECX_CET_USER_STATE(_) =                                  (((_) >> 11) & 0x01) //col:3881
CPUID_ECX_CET_SUPERVISOR_STATE_BIT =                           12 //col:3887
CPUID_ECX_CET_SUPERVISOR_STATE_FLAG =                          0x1000 //col:3888
CPUID_ECX_CET_SUPERVISOR_STATE_MASK =                          0x01 //col:3889
CPUID_ECX_CET_SUPERVISOR_STATE(_) =                            (((_) >> 12) & 0x01) //col:3890
CPUID_ECX_HDC_STATE_BIT =                                      13 //col:3896
CPUID_ECX_HDC_STATE_FLAG =                                     0x2000 //col:3897
CPUID_ECX_HDC_STATE_MASK =                                     0x01 //col:3898
CPUID_ECX_HDC_STATE(_) =                                       (((_) >> 13) & 0x01) //col:3899
CPUID_ECX_LBR_STATE_BIT =                                      15 //col:3906
CPUID_ECX_LBR_STATE_FLAG =                                     0x8000 //col:3907
CPUID_ECX_LBR_STATE_MASK =                                     0x01 //col:3908
CPUID_ECX_LBR_STATE(_) =                                       (((_) >> 15) & 0x01) //col:3909
CPUID_ECX_HWP_STATE_BIT =                                      16 //col:3915
CPUID_ECX_HWP_STATE_FLAG =                                     0x10000 //col:3916
CPUID_ECX_HWP_STATE_MASK =                                     0x01 //col:3917
CPUID_ECX_HWP_STATE(_) =                                       (((_) >> 16) & 0x01) //col:3918
CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_BIT =                  0 //col:3934
CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_FLAG =                 0xFFFFFFFF //col:3935
CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_MASK =                 0xFFFFFFFF //col:3936
CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS(_) =                   (((_) >> 0) & 0xFFFFFFFF) //col:3937
CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT =                          0 //col:3965
CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG =                         0xFFFFFFFF //col:3966
CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK =                         0xFFFFFFFF //col:3967
CPUID_EAX_IA32_PLATFORM_DCA_CAP(_) =                           (((_) >> 0) & 0xFFFFFFFF) //col:3968
CPUID_EBX_RESERVED_BIT =                                       0 //col:3984
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:3985
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:3986
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:3987
CPUID_ECX_ECX_2_BIT =                                          0 //col:4002
CPUID_ECX_ECX_2_FLAG =                                         0x01 //col:4003
CPUID_ECX_ECX_2_MASK =                                         0x01 //col:4004
CPUID_ECX_ECX_2(_) =                                           (((_) >> 0) & 0x01) //col:4005
CPUID_ECX_ECX_1_BIT =                                          1 //col:4013
CPUID_ECX_ECX_1_FLAG =                                         0x02 //col:4014
CPUID_ECX_ECX_1_MASK =                                         0x01 //col:4015
CPUID_ECX_ECX_1(_) =                                           (((_) >> 1) & 0x01) //col:4016
CPUID_EDX_RESERVED_BIT =                                       0 //col:4031
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4032
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:4033
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4034
CPUID_INTEL_RESOURCE_DIRECTOR_TECHNOLOGY_MONITORING_INFORMATION = 0x0000000F //col:4059
CPUID_EAX_RESERVED_BIT =                                       0 //col:4076
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4077
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:4078
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4079
CPUID_EBX_RMID_MAX_RANGE_BIT =                                 0 //col:4093
CPUID_EBX_RMID_MAX_RANGE_FLAG =                                0xFFFFFFFF //col:4094
CPUID_EBX_RMID_MAX_RANGE_MASK =                                0xFFFFFFFF //col:4095
CPUID_EBX_RMID_MAX_RANGE(_) =                                  (((_) >> 0) & 0xFFFFFFFF) //col:4096
CPUID_ECX_RESERVED_BIT =                                       0 //col:4110
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4111
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:4112
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4113
CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_BIT =         1 //col:4129
CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_FLAG =        0x02 //col:4130
CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_MASK =        0x01 //col:4131
CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING(_) =          (((_) >> 1) & 0x01) //col:4132
CPUID_EAX_RESERVED_BIT =                                       0 //col:4156
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4157
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:4158
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4159
CPUID_EBX_CONVERSION_FACTOR_BIT =                              0 //col:4173
CPUID_EBX_CONVERSION_FACTOR_FLAG =                             0xFFFFFFFF //col:4174
CPUID_EBX_CONVERSION_FACTOR_MASK =                             0xFFFFFFFF //col:4175
CPUID_EBX_CONVERSION_FACTOR(_) =                               (((_) >> 0) & 0xFFFFFFFF) //col:4176
CPUID_ECX_RMID_MAX_RANGE_BIT =                                 0 //col:4190
CPUID_ECX_RMID_MAX_RANGE_FLAG =                                0xFFFFFFFF //col:4191
CPUID_ECX_RMID_MAX_RANGE_MASK =                                0xFFFFFFFF //col:4192
CPUID_ECX_RMID_MAX_RANGE(_) =                                  (((_) >> 0) & 0xFFFFFFFF) //col:4193
CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_BIT =               0 //col:4207
CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_FLAG =              0x01 //col:4208
CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_MASK =              0x01 //col:4209
CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING(_) =                (((_) >> 0) & 0x01) //col:4210
CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_BIT =         1 //col:4216
CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_FLAG =        0x02 //col:4217
CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_MASK =        0x01 //col:4218
CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING(_) =          (((_) >> 1) & 0x01) //col:4219
CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_BIT =         2 //col:4225
CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_FLAG =        0x04 //col:4226
CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_MASK =        0x01 //col:4227
CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING(_) =          (((_) >> 2) & 0x01) //col:4228
CPUID_INTEL_RESOURCE_DIRECTOR_TECHNOLOGY_ALLOCATION_INFORMATION = 0x00000010 //col:4254
CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT =                          0 //col:4271
CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG =                         0xFFFFFFFF //col:4272
CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK =                         0xFFFFFFFF //col:4273
CPUID_EAX_IA32_PLATFORM_DCA_CAP(_) =                           (((_) >> 0) & 0xFFFFFFFF) //col:4274
CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_BIT =        1 //col:4290
CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_FLAG =       0x02 //col:4291
CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_MASK =       0x01 //col:4292
CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY(_) =         (((_) >> 1) & 0x01) //col:4293
CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_BIT =        2 //col:4299
CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_FLAG =       0x04 //col:4300
CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_MASK =       0x01 //col:4301
CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY(_) =         (((_) >> 2) & 0x01) //col:4302
CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_BIT =           3 //col:4308
CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_FLAG =          0x08 //col:4309
CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_MASK =          0x01 //col:4310
CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION(_) =            (((_) >> 3) & 0x01) //col:4311
CPUID_ECX_RESERVED_BIT =                                       0 //col:4326
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4327
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:4328
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4329
CPUID_EDX_RESERVED_BIT =                                       0 //col:4343
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4344
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:4345
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4346
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_BIT =                    0 //col:4369
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_FLAG =                   0x1F //col:4370
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_MASK =                   0x1F //col:4371
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK(_) =                     (((_) >> 0) & 0x1F) //col:4372
CPUID_EBX_EBX_0_BIT =                                          0 //col:4387
CPUID_EBX_EBX_0_FLAG =                                         0xFFFFFFFF //col:4388
CPUID_EBX_EBX_0_MASK =                                         0xFFFFFFFF //col:4389
CPUID_EBX_EBX_0(_) =                                           (((_) >> 0) & 0xFFFFFFFF) //col:4390
CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_BIT = 2 //col:4406
CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_FLAG = 0x04 //col:4407
CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_MASK = 0x01 //col:4408
CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED(_) = (((_) >> 2) & 0x01) //col:4409
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT =                   0 //col:4424
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG =                  0xFFFF //col:4425
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK =                  0xFFFF //col:4426
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_) =                    (((_) >> 0) & 0xFFFF) //col:4427
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_BIT =                    0 //col:4451
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_FLAG =                   0x1F //col:4452
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_MASK =                   0x1F //col:4453
CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK(_) =                     (((_) >> 0) & 0x1F) //col:4454
CPUID_EBX_EBX_0_BIT =                                          0 //col:4469
CPUID_EBX_EBX_0_FLAG =                                         0xFFFFFFFF //col:4470
CPUID_EBX_EBX_0_MASK =                                         0xFFFFFFFF //col:4471
CPUID_EBX_EBX_0(_) =                                           (((_) >> 0) & 0xFFFFFFFF) //col:4472
CPUID_ECX_RESERVED_BIT =                                       0 //col:4486
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4487
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:4488
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4489
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT =                   0 //col:4503
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG =                  0xFFFF //col:4504
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK =                  0xFFFF //col:4505
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_) =                    (((_) >> 0) & 0xFFFF) //col:4506
CPUID_EAX_MAX_MBA_THROTTLING_VALUE_BIT =                       0 //col:4530
CPUID_EAX_MAX_MBA_THROTTLING_VALUE_FLAG =                      0xFFF //col:4531
CPUID_EAX_MAX_MBA_THROTTLING_VALUE_MASK =                      0xFFF //col:4532
CPUID_EAX_MAX_MBA_THROTTLING_VALUE(_) =                        (((_) >> 0) & 0xFFF) //col:4533
CPUID_EBX_RESERVED_BIT =                                       0 //col:4548
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4549
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:4550
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4551
CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_BIT =                    2 //col:4567
CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_FLAG =                   0x04 //col:4568
CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_MASK =                   0x01 //col:4569
CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR(_) =                     (((_) >> 2) & 0x01) //col:4570
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT =                   0 //col:4585
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG =                  0xFFFF //col:4586
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK =                  0xFFFF //col:4587
CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_) =                    (((_) >> 0) & 0xFFFF) //col:4588
CPUID_INTEL_SGX =                                              0x00000012 //col:4611
CPUID_EAX_SGX1_BIT =                                           0 //col:4627
CPUID_EAX_SGX1_FLAG =                                          0x01 //col:4628
CPUID_EAX_SGX1_MASK =                                          0x01 //col:4629
CPUID_EAX_SGX1(_) =                                            (((_) >> 0) & 0x01) //col:4630
CPUID_EAX_SGX2_BIT =                                           1 //col:4636
CPUID_EAX_SGX2_FLAG =                                          0x02 //col:4637
CPUID_EAX_SGX2_MASK =                                          0x01 //col:4638
CPUID_EAX_SGX2(_) =                                            (((_) >> 1) & 0x01) //col:4639
CPUID_EAX_SGX_ENCLV_ADVANCED_BIT =                             5 //col:4646
CPUID_EAX_SGX_ENCLV_ADVANCED_FLAG =                            0x20 //col:4647
CPUID_EAX_SGX_ENCLV_ADVANCED_MASK =                            0x01 //col:4648
CPUID_EAX_SGX_ENCLV_ADVANCED(_) =                              (((_) >> 5) & 0x01) //col:4649
CPUID_EAX_SGX_ENCLS_ADVANCED_BIT =                             6 //col:4655
CPUID_EAX_SGX_ENCLS_ADVANCED_FLAG =                            0x40 //col:4656
CPUID_EAX_SGX_ENCLS_ADVANCED_MASK =                            0x01 //col:4657
CPUID_EAX_SGX_ENCLS_ADVANCED(_) =                              (((_) >> 6) & 0x01) //col:4658
CPUID_EBX_MISCSELECT_BIT =                                     0 //col:4673
CPUID_EBX_MISCSELECT_FLAG =                                    0xFFFFFFFF //col:4674
CPUID_EBX_MISCSELECT_MASK =                                    0xFFFFFFFF //col:4675
CPUID_EBX_MISCSELECT(_) =                                      (((_) >> 0) & 0xFFFFFFFF) //col:4676
CPUID_ECX_RESERVED_BIT =                                       0 //col:4690
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:4691
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:4692
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:4693
CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_BIT =                         0 //col:4707
CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_FLAG =                        0xFF //col:4708
CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_MASK =                        0xFF //col:4709
CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64(_) =                          (((_) >> 0) & 0xFF) //col:4710
CPUID_EDX_MAX_ENCLAVE_SIZE_64_BIT =                            8 //col:4716
CPUID_EDX_MAX_ENCLAVE_SIZE_64_FLAG =                           0xFF00 //col:4717
CPUID_EDX_MAX_ENCLAVE_SIZE_64_MASK =                           0xFF //col:4718
CPUID_EDX_MAX_ENCLAVE_SIZE_64(_) =                             (((_) >> 8) & 0xFF) //col:4719
CPUID_EAX_VALID_SECS_ATTRIBUTES_0_BIT =                        0 //col:4743
CPUID_EAX_VALID_SECS_ATTRIBUTES_0_FLAG =                       0xFFFFFFFF //col:4744
CPUID_EAX_VALID_SECS_ATTRIBUTES_0_MASK =                       0xFFFFFFFF //col:4745
CPUID_EAX_VALID_SECS_ATTRIBUTES_0(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:4746
CPUID_EBX_VALID_SECS_ATTRIBUTES_1_BIT =                        0 //col:4760
CPUID_EBX_VALID_SECS_ATTRIBUTES_1_FLAG =                       0xFFFFFFFF //col:4761
CPUID_EBX_VALID_SECS_ATTRIBUTES_1_MASK =                       0xFFFFFFFF //col:4762
CPUID_EBX_VALID_SECS_ATTRIBUTES_1(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:4763
CPUID_ECX_VALID_SECS_ATTRIBUTES_2_BIT =                        0 //col:4777
CPUID_ECX_VALID_SECS_ATTRIBUTES_2_FLAG =                       0xFFFFFFFF //col:4778
CPUID_ECX_VALID_SECS_ATTRIBUTES_2_MASK =                       0xFFFFFFFF //col:4779
CPUID_ECX_VALID_SECS_ATTRIBUTES_2(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:4780
CPUID_EDX_VALID_SECS_ATTRIBUTES_3_BIT =                        0 //col:4794
CPUID_EDX_VALID_SECS_ATTRIBUTES_3_FLAG =                       0xFFFFFFFF //col:4795
CPUID_EDX_VALID_SECS_ATTRIBUTES_3_MASK =                       0xFFFFFFFF //col:4796
CPUID_EDX_VALID_SECS_ATTRIBUTES_3(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:4797
CPUID_EAX_SUB_LEAF_TYPE_BIT =                                  0 //col:4821
CPUID_EAX_SUB_LEAF_TYPE_FLAG =                                 0x0F //col:4822
CPUID_EAX_SUB_LEAF_TYPE_MASK =                                 0x0F //col:4823
CPUID_EAX_SUB_LEAF_TYPE(_) =                                   (((_) >> 0) & 0x0F) //col:4824
CPUID_EBX_ZERO_BIT =                                           0 //col:4839
CPUID_EBX_ZERO_FLAG =                                          0xFFFFFFFF //col:4840
CPUID_EBX_ZERO_MASK =                                          0xFFFFFFFF //col:4841
CPUID_EBX_ZERO(_) =                                            (((_) >> 0) & 0xFFFFFFFF) //col:4842
CPUID_ECX_ZERO_BIT =                                           0 //col:4856
CPUID_ECX_ZERO_FLAG =                                          0xFFFFFFFF //col:4857
CPUID_ECX_ZERO_MASK =                                          0xFFFFFFFF //col:4858
CPUID_ECX_ZERO(_) =                                            (((_) >> 0) & 0xFFFFFFFF) //col:4859
CPUID_EDX_ZERO_BIT =                                           0 //col:4873
CPUID_EDX_ZERO_FLAG =                                          0xFFFFFFFF //col:4874
CPUID_EDX_ZERO_MASK =                                          0xFFFFFFFF //col:4875
CPUID_EDX_ZERO(_) =                                            (((_) >> 0) & 0xFFFFFFFF) //col:4876
CPUID_EAX_SUB_LEAF_TYPE_BIT =                                  0 //col:4901
CPUID_EAX_SUB_LEAF_TYPE_FLAG =                                 0x0F //col:4902
CPUID_EAX_SUB_LEAF_TYPE_MASK =                                 0x0F //col:4903
CPUID_EAX_SUB_LEAF_TYPE(_) =                                   (((_) >> 0) & 0x0F) //col:4904
CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_BIT =                    12 //col:4911
CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_FLAG =                   0xFFFFF000 //col:4912
CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_MASK =                   0xFFFFF //col:4913
CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1(_) =                     (((_) >> 12) & 0xFFFFF) //col:4914
CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_BIT =                    0 //col:4928
CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_FLAG =                   0xFFFFF //col:4929
CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_MASK =                   0xFFFFF //col:4930
CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2(_) =                     (((_) >> 0) & 0xFFFFF) //col:4931
CPUID_ECX_EPC_SECTION_PROPERTY_BIT =                           0 //col:4949
CPUID_ECX_EPC_SECTION_PROPERTY_FLAG =                          0x0F //col:4950
CPUID_ECX_EPC_SECTION_PROPERTY_MASK =                          0x0F //col:4951
CPUID_ECX_EPC_SECTION_PROPERTY(_) =                            (((_) >> 0) & 0x0F) //col:4952
CPUID_ECX_EPC_SIZE_1_BIT =                                     12 //col:4959
CPUID_ECX_EPC_SIZE_1_FLAG =                                    0xFFFFF000 //col:4960
CPUID_ECX_EPC_SIZE_1_MASK =                                    0xFFFFF //col:4961
CPUID_ECX_EPC_SIZE_1(_) =                                      (((_) >> 12) & 0xFFFFF) //col:4962
CPUID_EDX_EPC_SIZE_2_BIT =                                     0 //col:4976
CPUID_EDX_EPC_SIZE_2_FLAG =                                    0xFFFFF //col:4977
CPUID_EDX_EPC_SIZE_2_MASK =                                    0xFFFFF //col:4978
CPUID_EDX_EPC_SIZE_2(_) =                                      (((_) >> 0) & 0xFFFFF) //col:4979
CPUID_INTEL_PROCESSOR_TRACE_INFORMATION =                      0x00000014 //col:5002
CPUID_EAX_MAX_SUB_LEAF_BIT =                                   0 //col:5018
CPUID_EAX_MAX_SUB_LEAF_FLAG =                                  0xFFFFFFFF //col:5019
CPUID_EAX_MAX_SUB_LEAF_MASK =                                  0xFFFFFFFF //col:5020
CPUID_EAX_MAX_SUB_LEAF(_) =                                    (((_) >> 0) & 0xFFFFFFFF) //col:5021
CPUID_EBX_FLAG0_BIT =                                          0 //col:5035
CPUID_EBX_FLAG0_FLAG =                                         0x01 //col:5036
CPUID_EBX_FLAG0_MASK =                                         0x01 //col:5037
CPUID_EBX_FLAG0(_) =                                           (((_) >> 0) & 0x01) //col:5038
CPUID_EBX_FLAG1_BIT =                                          1 //col:5044
CPUID_EBX_FLAG1_FLAG =                                         0x02 //col:5045
CPUID_EBX_FLAG1_MASK =                                         0x01 //col:5046
CPUID_EBX_FLAG1(_) =                                           (((_) >> 1) & 0x01) //col:5047
CPUID_EBX_FLAG2_BIT =                                          2 //col:5054
CPUID_EBX_FLAG2_FLAG =                                         0x04 //col:5055
CPUID_EBX_FLAG2_MASK =                                         0x01 //col:5056
CPUID_EBX_FLAG2(_) =                                           (((_) >> 2) & 0x01) //col:5057
CPUID_EBX_FLAG3_BIT =                                          3 //col:5063
CPUID_EBX_FLAG3_FLAG =                                         0x08 //col:5064
CPUID_EBX_FLAG3_MASK =                                         0x01 //col:5065
CPUID_EBX_FLAG3(_) =                                           (((_) >> 3) & 0x01) //col:5066
CPUID_EBX_FLAG4_BIT =                                          4 //col:5073
CPUID_EBX_FLAG4_FLAG =                                         0x10 //col:5074
CPUID_EBX_FLAG4_MASK =                                         0x01 //col:5075
CPUID_EBX_FLAG4(_) =                                           (((_) >> 4) & 0x01) //col:5076
CPUID_EBX_FLAG5_BIT =                                          5 //col:5083
CPUID_EBX_FLAG5_FLAG =                                         0x20 //col:5084
CPUID_EBX_FLAG5_MASK =                                         0x01 //col:5085
CPUID_EBX_FLAG5(_) =                                           (((_) >> 5) & 0x01) //col:5086
CPUID_EBX_FLAG6_BIT =                                          6 //col:5094
CPUID_EBX_FLAG6_FLAG =                                         0x40 //col:5095
CPUID_EBX_FLAG6_MASK =                                         0x01 //col:5096
CPUID_EBX_FLAG6(_) =                                           (((_) >> 6) & 0x01) //col:5097
CPUID_EBX_FLAG7_BIT =                                          7 //col:5103
CPUID_EBX_FLAG7_FLAG =                                         0x80 //col:5104
CPUID_EBX_FLAG7_MASK =                                         0x01 //col:5105
CPUID_EBX_FLAG7(_) =                                           (((_) >> 7) & 0x01) //col:5106
CPUID_EBX_FLAG8_BIT =                                          8 //col:5112
CPUID_EBX_FLAG8_FLAG =                                         0x100 //col:5113
CPUID_EBX_FLAG8_MASK =                                         0x01 //col:5114
CPUID_EBX_FLAG8(_) =                                           (((_) >> 8) & 0x01) //col:5115
CPUID_ECX_FLAG0_BIT =                                          0 //col:5131
CPUID_ECX_FLAG0_FLAG =                                         0x01 //col:5132
CPUID_ECX_FLAG0_MASK =                                         0x01 //col:5133
CPUID_ECX_FLAG0(_) =                                           (((_) >> 0) & 0x01) //col:5134
CPUID_ECX_FLAG1_BIT =                                          1 //col:5141
CPUID_ECX_FLAG1_FLAG =                                         0x02 //col:5142
CPUID_ECX_FLAG1_MASK =                                         0x01 //col:5143
CPUID_ECX_FLAG1(_) =                                           (((_) >> 1) & 0x01) //col:5144
CPUID_ECX_FLAG2_BIT =                                          2 //col:5150
CPUID_ECX_FLAG2_FLAG =                                         0x04 //col:5151
CPUID_ECX_FLAG2_MASK =                                         0x01 //col:5152
CPUID_ECX_FLAG2(_) =                                           (((_) >> 2) & 0x01) //col:5153
CPUID_ECX_FLAG3_BIT =                                          3 //col:5159
CPUID_ECX_FLAG3_FLAG =                                         0x08 //col:5160
CPUID_ECX_FLAG3_MASK =                                         0x01 //col:5161
CPUID_ECX_FLAG3(_) =                                           (((_) >> 3) & 0x01) //col:5162
CPUID_ECX_FLAG31_BIT =                                         31 //col:5169
CPUID_ECX_FLAG31_FLAG =                                        0x80000000 //col:5170
CPUID_ECX_FLAG31_MASK =                                        0x01 //col:5171
CPUID_ECX_FLAG31(_) =                                          (((_) >> 31) & 0x01) //col:5172
CPUID_EDX_RESERVED_BIT =                                       0 //col:5186
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5187
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:5188
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5189
CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_BIT = 0 //col:5210
CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_FLAG = 0x07 //col:5211
CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_MASK = 0x07 //col:5212
CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING(_) = (((_) >> 0) & 0x07) //col:5213
CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_BIT =       16 //col:5220
CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_FLAG =      0xFFFF0000 //col:5221
CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_MASK =      0xFFFF //col:5222
CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS(_) =        (((_) >> 16) & 0xFFFF) //col:5223
CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_BIT = 0 //col:5237
CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_FLAG = 0xFFFF //col:5238
CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_MASK = 0xFFFF //col:5239
CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS(_) = (((_) >> 0) & 0xFFFF) //col:5240
CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_BIT = 16 //col:5246
CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_FLAG = 0xFFFF0000 //col:5247
CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_MASK = 0xFFFF //col:5248
CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS(_) = (((_) >> 16) & 0xFFFF) //col:5249
CPUID_ECX_RESERVED_BIT =                                       0 //col:5263
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5264
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:5265
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5266
CPUID_EDX_RESERVED_BIT =                                       0 //col:5280
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5281
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:5282
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5283
CPUID_TIME_STAMP_COUNTER_INFORMATION =                         0x00000015 //col:5307
CPUID_EAX_DENOMINATOR_BIT =                                    0 //col:5318
CPUID_EAX_DENOMINATOR_FLAG =                                   0xFFFFFFFF //col:5319
CPUID_EAX_DENOMINATOR_MASK =                                   0xFFFFFFFF //col:5320
CPUID_EAX_DENOMINATOR(_) =                                     (((_) >> 0) & 0xFFFFFFFF) //col:5321
CPUID_EBX_NUMERATOR_BIT =                                      0 //col:5335
CPUID_EBX_NUMERATOR_FLAG =                                     0xFFFFFFFF //col:5336
CPUID_EBX_NUMERATOR_MASK =                                     0xFFFFFFFF //col:5337
CPUID_EBX_NUMERATOR(_) =                                       (((_) >> 0) & 0xFFFFFFFF) //col:5338
CPUID_ECX_NOMINAL_FREQUENCY_BIT =                              0 //col:5352
CPUID_ECX_NOMINAL_FREQUENCY_FLAG =                             0xFFFFFFFF //col:5353
CPUID_ECX_NOMINAL_FREQUENCY_MASK =                             0xFFFFFFFF //col:5354
CPUID_ECX_NOMINAL_FREQUENCY(_) =                               (((_) >> 0) & 0xFFFFFFFF) //col:5355
CPUID_EDX_RESERVED_BIT =                                       0 //col:5369
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5370
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:5371
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5372
CPUID_PROCESSOR_FREQUENCY_INFORMATION =                        0x00000016 //col:5394
CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_BIT =                    0 //col:5405
CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_FLAG =                   0xFFFF //col:5406
CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_MASK =                   0xFFFF //col:5407
CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ(_) =                     (((_) >> 0) & 0xFFFF) //col:5408
CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_BIT =                0 //col:5423
CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_FLAG =               0xFFFF //col:5424
CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_MASK =               0xFFFF //col:5425
CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ(_) =                 (((_) >> 0) & 0xFFFF) //col:5426
CPUID_ECX_BUS_FREQUENCY_MHZ_BIT =                              0 //col:5441
CPUID_ECX_BUS_FREQUENCY_MHZ_FLAG =                             0xFFFF //col:5442
CPUID_ECX_BUS_FREQUENCY_MHZ_MASK =                             0xFFFF //col:5443
CPUID_ECX_BUS_FREQUENCY_MHZ(_) =                               (((_) >> 0) & 0xFFFF) //col:5444
CPUID_EDX_RESERVED_BIT =                                       0 //col:5459
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5460
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:5461
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5462
CPUID_SOC_VENDOR_INFORMATION =                                 0x00000017 //col:5478
CPUID_EAX_MAX_SOC_ID_INDEX_BIT =                               0 //col:5495
CPUID_EAX_MAX_SOC_ID_INDEX_FLAG =                              0xFFFFFFFF //col:5496
CPUID_EAX_MAX_SOC_ID_INDEX_MASK =                              0xFFFFFFFF //col:5497
CPUID_EAX_MAX_SOC_ID_INDEX(_) =                                (((_) >> 0) & 0xFFFFFFFF) //col:5498
CPUID_EBX_SOC_VENDOR_ID_BIT =                                  0 //col:5512
CPUID_EBX_SOC_VENDOR_ID_FLAG =                                 0xFFFF //col:5513
CPUID_EBX_SOC_VENDOR_ID_MASK =                                 0xFFFF //col:5514
CPUID_EBX_SOC_VENDOR_ID(_) =                                   (((_) >> 0) & 0xFFFF) //col:5515
CPUID_EBX_IS_VENDOR_SCHEME_BIT =                               16 //col:5522
CPUID_EBX_IS_VENDOR_SCHEME_FLAG =                              0x10000 //col:5523
CPUID_EBX_IS_VENDOR_SCHEME_MASK =                              0x01 //col:5524
CPUID_EBX_IS_VENDOR_SCHEME(_) =                                (((_) >> 16) & 0x01) //col:5525
CPUID_ECX_PROJECT_ID_BIT =                                     0 //col:5540
CPUID_ECX_PROJECT_ID_FLAG =                                    0xFFFFFFFF //col:5541
CPUID_ECX_PROJECT_ID_MASK =                                    0xFFFFFFFF //col:5542
CPUID_ECX_PROJECT_ID(_) =                                      (((_) >> 0) & 0xFFFFFFFF) //col:5543
CPUID_EDX_STEPPING_ID_BIT =                                    0 //col:5557
CPUID_EDX_STEPPING_ID_FLAG =                                   0xFFFFFFFF //col:5558
CPUID_EDX_STEPPING_ID_MASK =                                   0xFFFFFFFF //col:5559
CPUID_EDX_STEPPING_ID(_) =                                     (((_) >> 0) & 0xFFFFFFFF) //col:5560
CPUID_EAX_SOC_VENDOR_BRAND_STRING_BIT =                        0 //col:5585
CPUID_EAX_SOC_VENDOR_BRAND_STRING_FLAG =                       0xFFFFFFFF //col:5586
CPUID_EAX_SOC_VENDOR_BRAND_STRING_MASK =                       0xFFFFFFFF //col:5587
CPUID_EAX_SOC_VENDOR_BRAND_STRING(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:5588
CPUID_EBX_SOC_VENDOR_BRAND_STRING_BIT =                        0 //col:5602
CPUID_EBX_SOC_VENDOR_BRAND_STRING_FLAG =                       0xFFFFFFFF //col:5603
CPUID_EBX_SOC_VENDOR_BRAND_STRING_MASK =                       0xFFFFFFFF //col:5604
CPUID_EBX_SOC_VENDOR_BRAND_STRING(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:5605
CPUID_ECX_SOC_VENDOR_BRAND_STRING_BIT =                        0 //col:5619
CPUID_ECX_SOC_VENDOR_BRAND_STRING_FLAG =                       0xFFFFFFFF //col:5620
CPUID_ECX_SOC_VENDOR_BRAND_STRING_MASK =                       0xFFFFFFFF //col:5621
CPUID_ECX_SOC_VENDOR_BRAND_STRING(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:5622
CPUID_EDX_SOC_VENDOR_BRAND_STRING_BIT =                        0 //col:5636
CPUID_EDX_SOC_VENDOR_BRAND_STRING_FLAG =                       0xFFFFFFFF //col:5637
CPUID_EDX_SOC_VENDOR_BRAND_STRING_MASK =                       0xFFFFFFFF //col:5638
CPUID_EDX_SOC_VENDOR_BRAND_STRING(_) =                         (((_) >> 0) & 0xFFFFFFFF) //col:5639
CPUID_EAX_RESERVED_BIT =                                       0 //col:5662
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5663
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:5664
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5665
CPUID_EBX_RESERVED_BIT =                                       0 //col:5679
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5680
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:5681
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5682
CPUID_ECX_RESERVED_BIT =                                       0 //col:5696
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5697
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:5698
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5699
CPUID_EDX_RESERVED_BIT =                                       0 //col:5713
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5714
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:5715
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5716
CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS =           0x00000018 //col:5736
CPUID_EAX_MAX_SUB_LEAF_BIT =                                   0 //col:5756
CPUID_EAX_MAX_SUB_LEAF_FLAG =                                  0xFFFFFFFF //col:5757
CPUID_EAX_MAX_SUB_LEAF_MASK =                                  0xFFFFFFFF //col:5758
CPUID_EAX_MAX_SUB_LEAF(_) =                                    (((_) >> 0) & 0xFFFFFFFF) //col:5759
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_BIT =                     0 //col:5773
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_FLAG =                    0x01 //col:5774
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_MASK =                    0x01 //col:5775
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED(_) =                      (((_) >> 0) & 0x01) //col:5776
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_BIT =                     1 //col:5782
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_FLAG =                    0x02 //col:5783
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_MASK =                    0x01 //col:5784
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED(_) =                      (((_) >> 1) & 0x01) //col:5785
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_BIT =                     2 //col:5791
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_FLAG =                    0x04 //col:5792
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_MASK =                    0x01 //col:5793
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED(_) =                      (((_) >> 2) & 0x01) //col:5794
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_BIT =                     3 //col:5800
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_FLAG =                    0x08 //col:5801
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_MASK =                    0x01 //col:5802
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED(_) =                      (((_) >> 3) & 0x01) //col:5803
CPUID_EBX_PARTITIONING_BIT =                                   8 //col:5810
CPUID_EBX_PARTITIONING_FLAG =                                  0x700 //col:5811
CPUID_EBX_PARTITIONING_MASK =                                  0x07 //col:5812
CPUID_EBX_PARTITIONING(_) =                                    (((_) >> 8) & 0x07) //col:5813
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_BIT =                       16 //col:5820
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_FLAG =                      0xFFFF0000 //col:5821
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_MASK =                      0xFFFF //col:5822
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00(_) =                        (((_) >> 16) & 0xFFFF) //col:5823
CPUID_ECX_NUMBER_OF_SETS_BIT =                                 0 //col:5837
CPUID_ECX_NUMBER_OF_SETS_FLAG =                                0xFFFFFFFF //col:5838
CPUID_ECX_NUMBER_OF_SETS_MASK =                                0xFFFFFFFF //col:5839
CPUID_ECX_NUMBER_OF_SETS(_) =                                  (((_) >> 0) & 0xFFFFFFFF) //col:5840
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_BIT =                   0 //col:5863
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_FLAG =                  0x1F //col:5864
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_MASK =                  0x1F //col:5865
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD(_) =                    (((_) >> 0) & 0x1F) //col:5866
CPUID_EDX_TRANSLATION_CACHE_LEVEL_BIT =                        5 //col:5872
CPUID_EDX_TRANSLATION_CACHE_LEVEL_FLAG =                       0xE0 //col:5873
CPUID_EDX_TRANSLATION_CACHE_LEVEL_MASK =                       0x07 //col:5874
CPUID_EDX_TRANSLATION_CACHE_LEVEL(_) =                         (((_) >> 5) & 0x07) //col:5875
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_BIT =                    8 //col:5881
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_FLAG =                   0x100 //col:5882
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_MASK =                   0x01 //col:5883
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE(_) =                     (((_) >> 8) & 0x01) //col:5884
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_BIT =     14 //col:5893
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_FLAG =    0x3FFC000 //col:5894
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_MASK =    0xFFF //col:5895
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS(_) =      (((_) >> 14) & 0xFFF) //col:5896
CPUID_EAX_RESERVED_BIT =                                       0 //col:5924
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:5925
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:5926
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:5927
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_BIT =                     0 //col:5941
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_FLAG =                    0x01 //col:5942
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_MASK =                    0x01 //col:5943
CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED(_) =                      (((_) >> 0) & 0x01) //col:5944
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_BIT =                     1 //col:5950
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_FLAG =                    0x02 //col:5951
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_MASK =                    0x01 //col:5952
CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED(_) =                      (((_) >> 1) & 0x01) //col:5953
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_BIT =                     2 //col:5959
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_FLAG =                    0x04 //col:5960
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_MASK =                    0x01 //col:5961
CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED(_) =                      (((_) >> 2) & 0x01) //col:5962
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_BIT =                     3 //col:5968
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_FLAG =                    0x08 //col:5969
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_MASK =                    0x01 //col:5970
CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED(_) =                      (((_) >> 3) & 0x01) //col:5971
CPUID_EBX_PARTITIONING_BIT =                                   8 //col:5978
CPUID_EBX_PARTITIONING_FLAG =                                  0x700 //col:5979
CPUID_EBX_PARTITIONING_MASK =                                  0x07 //col:5980
CPUID_EBX_PARTITIONING(_) =                                    (((_) >> 8) & 0x07) //col:5981
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_BIT =                       16 //col:5988
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_FLAG =                      0xFFFF0000 //col:5989
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_MASK =                      0xFFFF //col:5990
CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01(_) =                        (((_) >> 16) & 0xFFFF) //col:5991
CPUID_ECX_NUMBER_OF_SETS_BIT =                                 0 //col:6005
CPUID_ECX_NUMBER_OF_SETS_FLAG =                                0xFFFFFFFF //col:6006
CPUID_ECX_NUMBER_OF_SETS_MASK =                                0xFFFFFFFF //col:6007
CPUID_ECX_NUMBER_OF_SETS(_) =                                  (((_) >> 0) & 0xFFFFFFFF) //col:6008
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_BIT =                   0 //col:6031
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_FLAG =                  0x1F //col:6032
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_MASK =                  0x1F //col:6033
CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD(_) =                    (((_) >> 0) & 0x1F) //col:6034
CPUID_EDX_TRANSLATION_CACHE_LEVEL_BIT =                        5 //col:6040
CPUID_EDX_TRANSLATION_CACHE_LEVEL_FLAG =                       0xE0 //col:6041
CPUID_EDX_TRANSLATION_CACHE_LEVEL_MASK =                       0x07 //col:6042
CPUID_EDX_TRANSLATION_CACHE_LEVEL(_) =                         (((_) >> 5) & 0x07) //col:6043
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_BIT =                    8 //col:6049
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_FLAG =                   0x100 //col:6050
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_MASK =                   0x01 //col:6051
CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE(_) =                     (((_) >> 8) & 0x01) //col:6052
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_BIT =     14 //col:6061
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_FLAG =    0x3FFC000 //col:6062
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_MASK =    0xFFF //col:6063
CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS(_) =      (((_) >> 14) & 0xFFF) //col:6064
CPUID_EXTENDED_FUNCTION_INFORMATION =                          0x80000000 //col:6084
CPUID_EAX_MAX_EXTENDED_FUNCTIONS_BIT =                         0 //col:6095
CPUID_EAX_MAX_EXTENDED_FUNCTIONS_FLAG =                        0xFFFFFFFF //col:6096
CPUID_EAX_MAX_EXTENDED_FUNCTIONS_MASK =                        0xFFFFFFFF //col:6097
CPUID_EAX_MAX_EXTENDED_FUNCTIONS(_) =                          (((_) >> 0) & 0xFFFFFFFF) //col:6098
CPUID_EBX_RESERVED_BIT =                                       0 //col:6112
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6113
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:6114
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6115
CPUID_ECX_RESERVED_BIT =                                       0 //col:6129
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6130
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:6131
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6132
CPUID_EDX_RESERVED_BIT =                                       0 //col:6146
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6147
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:6148
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6149
CPUID_EXTENDED_CPU_SIGNATURE =                                 0x80000001 //col:6161
CPUID_EAX_RESERVED_BIT =                                       0 //col:6172
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6173
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:6174
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6175
CPUID_EBX_RESERVED_BIT =                                       0 //col:6189
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6190
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:6191
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6192
CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_BIT =             0 //col:6206
CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_FLAG =            0x01 //col:6207
CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_MASK =            0x01 //col:6208
CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE(_) =              (((_) >> 0) & 0x01) //col:6209
CPUID_ECX_LZCNT_BIT =                                          5 //col:6216
CPUID_ECX_LZCNT_FLAG =                                         0x20 //col:6217
CPUID_ECX_LZCNT_MASK =                                         0x01 //col:6218
CPUID_ECX_LZCNT(_) =                                           (((_) >> 5) & 0x01) //col:6219
CPUID_ECX_PREFETCHW_BIT =                                      8 //col:6226
CPUID_ECX_PREFETCHW_FLAG =                                     0x100 //col:6227
CPUID_ECX_PREFETCHW_MASK =                                     0x01 //col:6228
CPUID_ECX_PREFETCHW(_) =                                       (((_) >> 8) & 0x01) //col:6229
CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_BIT =        11 //col:6246
CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_FLAG =       0x800 //col:6247
CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_MASK =       0x01 //col:6248
CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE(_) =         (((_) >> 11) & 0x01) //col:6249
CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_BIT =                  20 //col:6256
CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_FLAG =                 0x100000 //col:6257
CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_MASK =                 0x01 //col:6258
CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE(_) =                   (((_) >> 20) & 0x01) //col:6259
CPUID_EDX_PAGES_1GB_AVAILABLE_BIT =                            26 //col:6266
CPUID_EDX_PAGES_1GB_AVAILABLE_FLAG =                           0x4000000 //col:6267
CPUID_EDX_PAGES_1GB_AVAILABLE_MASK =                           0x01 //col:6268
CPUID_EDX_PAGES_1GB_AVAILABLE(_) =                             (((_) >> 26) & 0x01) //col:6269
CPUID_EDX_RDTSCP_AVAILABLE_BIT =                               27 //col:6275
CPUID_EDX_RDTSCP_AVAILABLE_FLAG =                              0x8000000 //col:6276
CPUID_EDX_RDTSCP_AVAILABLE_MASK =                              0x01 //col:6277
CPUID_EDX_RDTSCP_AVAILABLE(_) =                                (((_) >> 27) & 0x01) //col:6278
CPUID_EDX_IA64_AVAILABLE_BIT =                                 29 //col:6285
CPUID_EDX_IA64_AVAILABLE_FLAG =                                0x20000000 //col:6286
CPUID_EDX_IA64_AVAILABLE_MASK =                                0x01 //col:6287
CPUID_EDX_IA64_AVAILABLE(_) =                                  (((_) >> 29) & 0x01) //col:6288
CPUID_BRAND_STRING1 =                                          0x80000002 //col:6301
CPUID_BRAND_STRING2 =                                          0x80000003 //col:6306
CPUID_BRAND_STRING3 =                                          0x80000004 //col:6311
CPUID_EAX_PROCESSOR_BRAND_STRING_1_BIT =                       0 //col:6322
CPUID_EAX_PROCESSOR_BRAND_STRING_1_FLAG =                      0xFFFFFFFF //col:6323
CPUID_EAX_PROCESSOR_BRAND_STRING_1_MASK =                      0xFFFFFFFF //col:6324
CPUID_EAX_PROCESSOR_BRAND_STRING_1(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6325
CPUID_EBX_PROCESSOR_BRAND_STRING_2_BIT =                       0 //col:6339
CPUID_EBX_PROCESSOR_BRAND_STRING_2_FLAG =                      0xFFFFFFFF //col:6340
CPUID_EBX_PROCESSOR_BRAND_STRING_2_MASK =                      0xFFFFFFFF //col:6341
CPUID_EBX_PROCESSOR_BRAND_STRING_2(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6342
CPUID_ECX_PROCESSOR_BRAND_STRING_3_BIT =                       0 //col:6356
CPUID_ECX_PROCESSOR_BRAND_STRING_3_FLAG =                      0xFFFFFFFF //col:6357
CPUID_ECX_PROCESSOR_BRAND_STRING_3_MASK =                      0xFFFFFFFF //col:6358
CPUID_ECX_PROCESSOR_BRAND_STRING_3(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6359
CPUID_EDX_PROCESSOR_BRAND_STRING_4_BIT =                       0 //col:6373
CPUID_EDX_PROCESSOR_BRAND_STRING_4_FLAG =                      0xFFFFFFFF //col:6374
CPUID_EDX_PROCESSOR_BRAND_STRING_4_MASK =                      0xFFFFFFFF //col:6375
CPUID_EDX_PROCESSOR_BRAND_STRING_4(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6376
CPUID_EAX_PROCESSOR_BRAND_STRING_5_BIT =                       0 //col:6397
CPUID_EAX_PROCESSOR_BRAND_STRING_5_FLAG =                      0xFFFFFFFF //col:6398
CPUID_EAX_PROCESSOR_BRAND_STRING_5_MASK =                      0xFFFFFFFF //col:6399
CPUID_EAX_PROCESSOR_BRAND_STRING_5(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6400
CPUID_EBX_PROCESSOR_BRAND_STRING_6_BIT =                       0 //col:6414
CPUID_EBX_PROCESSOR_BRAND_STRING_6_FLAG =                      0xFFFFFFFF //col:6415
CPUID_EBX_PROCESSOR_BRAND_STRING_6_MASK =                      0xFFFFFFFF //col:6416
CPUID_EBX_PROCESSOR_BRAND_STRING_6(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6417
CPUID_ECX_PROCESSOR_BRAND_STRING_7_BIT =                       0 //col:6431
CPUID_ECX_PROCESSOR_BRAND_STRING_7_FLAG =                      0xFFFFFFFF //col:6432
CPUID_ECX_PROCESSOR_BRAND_STRING_7_MASK =                      0xFFFFFFFF //col:6433
CPUID_ECX_PROCESSOR_BRAND_STRING_7(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6434
CPUID_EDX_PROCESSOR_BRAND_STRING_8_BIT =                       0 //col:6448
CPUID_EDX_PROCESSOR_BRAND_STRING_8_FLAG =                      0xFFFFFFFF //col:6449
CPUID_EDX_PROCESSOR_BRAND_STRING_8_MASK =                      0xFFFFFFFF //col:6450
CPUID_EDX_PROCESSOR_BRAND_STRING_8(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6451
CPUID_EAX_PROCESSOR_BRAND_STRING_9_BIT =                       0 //col:6472
CPUID_EAX_PROCESSOR_BRAND_STRING_9_FLAG =                      0xFFFFFFFF //col:6473
CPUID_EAX_PROCESSOR_BRAND_STRING_9_MASK =                      0xFFFFFFFF //col:6474
CPUID_EAX_PROCESSOR_BRAND_STRING_9(_) =                        (((_) >> 0) & 0xFFFFFFFF) //col:6475
CPUID_EBX_PROCESSOR_BRAND_STRING_10_BIT =                      0 //col:6489
CPUID_EBX_PROCESSOR_BRAND_STRING_10_FLAG =                     0xFFFFFFFF //col:6490
CPUID_EBX_PROCESSOR_BRAND_STRING_10_MASK =                     0xFFFFFFFF //col:6491
CPUID_EBX_PROCESSOR_BRAND_STRING_10(_) =                       (((_) >> 0) & 0xFFFFFFFF) //col:6492
CPUID_ECX_PROCESSOR_BRAND_STRING_11_BIT =                      0 //col:6506
CPUID_ECX_PROCESSOR_BRAND_STRING_11_FLAG =                     0xFFFFFFFF //col:6507
CPUID_ECX_PROCESSOR_BRAND_STRING_11_MASK =                     0xFFFFFFFF //col:6508
CPUID_ECX_PROCESSOR_BRAND_STRING_11(_) =                       (((_) >> 0) & 0xFFFFFFFF) //col:6509
CPUID_EDX_PROCESSOR_BRAND_STRING_12_BIT =                      0 //col:6523
CPUID_EDX_PROCESSOR_BRAND_STRING_12_FLAG =                     0xFFFFFFFF //col:6524
CPUID_EDX_PROCESSOR_BRAND_STRING_12_MASK =                     0xFFFFFFFF //col:6525
CPUID_EDX_PROCESSOR_BRAND_STRING_12(_) =                       (((_) >> 0) & 0xFFFFFFFF) //col:6526
CPUID_EAX_RESERVED_BIT =                                       0 //col:6547
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6548
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:6549
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6550
CPUID_EBX_RESERVED_BIT =                                       0 //col:6564
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6565
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:6566
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6567
CPUID_ECX_RESERVED_BIT =                                       0 //col:6581
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6582
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:6583
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6584
CPUID_EDX_RESERVED_BIT =                                       0 //col:6598
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6599
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:6600
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6601
CPUID_EXTENDED_CACHE_INFO =                                    0x80000006 //col:6613
CPUID_EAX_RESERVED_BIT =                                       0 //col:6624
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6625
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:6626
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6627
CPUID_EBX_RESERVED_BIT =                                       0 //col:6641
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6642
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:6643
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6644
CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_BIT =                       0 //col:6658
CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_FLAG =                      0xFF //col:6659
CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_MASK =                      0xFF //col:6660
CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES(_) =                        (((_) >> 0) & 0xFF) //col:6661
CPUID_ECX_L2_ASSOCIATIVITY_FIELD_BIT =                         12 //col:6676
CPUID_ECX_L2_ASSOCIATIVITY_FIELD_FLAG =                        0xF000 //col:6677
CPUID_ECX_L2_ASSOCIATIVITY_FIELD_MASK =                        0x0F //col:6678
CPUID_ECX_L2_ASSOCIATIVITY_FIELD(_) =                          (((_) >> 12) & 0x0F) //col:6679
CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_BIT =                         16 //col:6685
CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_FLAG =                        0xFFFF0000 //col:6686
CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_MASK =                        0xFFFF //col:6687
CPUID_ECX_CACHE_SIZE_IN_1K_UNITS(_) =                          (((_) >> 16) & 0xFFFF) //col:6688
CPUID_EDX_RESERVED_BIT =                                       0 //col:6702
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6703
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:6704
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6705
CPUID_EXTENDED_TIME_STAMP_COUNTER =                            0x80000007 //col:6717
CPUID_EAX_RESERVED_BIT =                                       0 //col:6728
CPUID_EAX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6729
CPUID_EAX_RESERVED_MASK =                                      0xFFFFFFFF //col:6730
CPUID_EAX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6731
CPUID_EBX_RESERVED_BIT =                                       0 //col:6745
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6746
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:6747
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6748
CPUID_ECX_RESERVED_BIT =                                       0 //col:6762
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6763
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:6764
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6765
CPUID_EDX_INVARIANT_TSC_AVAILABLE_BIT =                        8 //col:6781
CPUID_EDX_INVARIANT_TSC_AVAILABLE_FLAG =                       0x100 //col:6782
CPUID_EDX_INVARIANT_TSC_AVAILABLE_MASK =                       0x01 //col:6783
CPUID_EDX_INVARIANT_TSC_AVAILABLE(_) =                         (((_) >> 8) & 0x01) //col:6784
CPUID_EXTENDED_VIRTUAL_PHYSICAL_ADDRESS_SIZE =                 0x80000008 //col:6797
CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_BIT =                0 //col:6811
CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_FLAG =               0xFF //col:6812
CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_MASK =               0xFF //col:6813
CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS(_) =                 (((_) >> 0) & 0xFF) //col:6814
CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_BIT =                  8 //col:6820
CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_FLAG =                 0xFF00 //col:6821
CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_MASK =                 0xFF //col:6822
CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS(_) =                   (((_) >> 8) & 0xFF) //col:6823
CPUID_EBX_RESERVED_BIT =                                       0 //col:6838
CPUID_EBX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6839
CPUID_EBX_RESERVED_MASK =                                      0xFFFFFFFF //col:6840
CPUID_EBX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6841
CPUID_ECX_RESERVED_BIT =                                       0 //col:6855
CPUID_ECX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6856
CPUID_ECX_RESERVED_MASK =                                      0xFFFFFFFF //col:6857
CPUID_ECX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6858
CPUID_EDX_RESERVED_BIT =                                       0 //col:6872
CPUID_EDX_RESERVED_FLAG =                                      0xFFFFFFFF //col:6873
CPUID_EDX_RESERVED_MASK =                                      0xFFFFFFFF //col:6874
CPUID_EDX_RESERVED(_) =                                        (((_) >> 0) & 0xFFFFFFFF) //col:6875
IA32_P5_MC_ADDR =                                              0x00000000 //col:6912
IA32_P5_MC_TYPE =                                              0x00000001 //col:6920
IA32_MONITOR_FILTER_LINE_SIZE =                                0x00000006 //col:6932
IA32_TIME_STAMP_COUNTER =                                      0x00000010 //col:6940
IA32_PLATFORM_ID =                                             0x00000017 //col:6948
IA32_PLATFORM_ID_PLATFORM_ID_BIT =                             50 //col:6973
IA32_PLATFORM_ID_PLATFORM_ID_FLAG =                            0x1C000000000000 //col:6974
IA32_PLATFORM_ID_PLATFORM_ID_MASK =                            0x07 //col:6975
IA32_PLATFORM_ID_PLATFORM_ID(_) =                              (((_) >> 50) & 0x07) //col:6976
IA32_APIC_BASE =                                               0x0000001B //col:6991
IA32_APIC_BASE_BSP_FLAG_BIT =                                  8 //col:7002
IA32_APIC_BASE_BSP_FLAG_FLAG =                                 0x100 //col:7003
IA32_APIC_BASE_BSP_FLAG_MASK =                                 0x01 //col:7004
IA32_APIC_BASE_BSP_FLAG(_) =                                   (((_) >> 8) & 0x01) //col:7005
IA32_APIC_BASE_ENABLE_X2APIC_MODE_BIT =                        10 //col:7012
IA32_APIC_BASE_ENABLE_X2APIC_MODE_FLAG =                       0x400 //col:7013
IA32_APIC_BASE_ENABLE_X2APIC_MODE_MASK =                       0x01 //col:7014
IA32_APIC_BASE_ENABLE_X2APIC_MODE(_) =                         (((_) >> 10) & 0x01) //col:7015
IA32_APIC_BASE_APIC_GLOBAL_ENABLE_BIT =                        11 //col:7021
IA32_APIC_BASE_APIC_GLOBAL_ENABLE_FLAG =                       0x800 //col:7022
IA32_APIC_BASE_APIC_GLOBAL_ENABLE_MASK =                       0x01 //col:7023
IA32_APIC_BASE_APIC_GLOBAL_ENABLE(_) =                         (((_) >> 11) & 0x01) //col:7024
IA32_APIC_BASE_APIC_BASE_BIT =                                 12 //col:7030
IA32_APIC_BASE_APIC_BASE_FLAG =                                0xFFFFFFFFF000 //col:7031
IA32_APIC_BASE_APIC_BASE_MASK =                                0xFFFFFFFFF //col:7032
IA32_APIC_BASE_APIC_BASE(_) =                                  (((_) >> 12) & 0xFFFFFFFFF) //col:7033
IA32_FEATURE_CONTROL =                                         0x0000003A //col:7046
IA32_FEATURE_CONTROL_LOCK_BIT_BIT =                            0 //col:7063
IA32_FEATURE_CONTROL_LOCK_BIT_FLAG =                           0x01 //col:7064
IA32_FEATURE_CONTROL_LOCK_BIT_MASK =                           0x01 //col:7065
IA32_FEATURE_CONTROL_LOCK_BIT(_) =                             (((_) >> 0) & 0x01) //col:7066
IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_BIT =               1 //col:7078
IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_FLAG =              0x02 //col:7079
IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_MASK =              0x01 //col:7080
IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX(_) =                (((_) >> 1) & 0x01) //col:7081
IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_BIT =              2 //col:7092
IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_FLAG =             0x04 //col:7093
IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_MASK =             0x01 //col:7094
IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX(_) =               (((_) >> 2) & 0x01) //col:7095
IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_BIT =       8 //col:7107
IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_FLAG =      0x7F00 //col:7108
IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_MASK =      0x7F //col:7109
IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES(_) =        (((_) >> 8) & 0x7F) //col:7110
IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_BIT =                15 //col:7120
IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_FLAG =               0x8000 //col:7121
IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_MASK =               0x01 //col:7122
IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE(_) =                 (((_) >> 15) & 0x01) //col:7123
IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_BIT =           17 //col:7134
IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_FLAG =          0x20000 //col:7135
IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_MASK =          0x01 //col:7136
IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE(_) =            (((_) >> 17) & 0x01) //col:7137
IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_BIT =                   18 //col:7147
IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_FLAG =                  0x40000 //col:7148
IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_MASK =                  0x01 //col:7149
IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE(_) =                    (((_) >> 18) & 0x01) //col:7150
IA32_FEATURE_CONTROL_LMCE_ON_BIT =                             20 //col:7162
IA32_FEATURE_CONTROL_LMCE_ON_FLAG =                            0x100000 //col:7163
IA32_FEATURE_CONTROL_LMCE_ON_MASK =                            0x01 //col:7164
IA32_FEATURE_CONTROL_LMCE_ON(_) =                              (((_) >> 20) & 0x01) //col:7165
IA32_TSC_ADJUST =                                              0x0000003B //col:7178
IA32_SPEC_CTRL =                                               0x00000048 //col:7196
IA32_SPEC_CTRL_IBRS_BIT =                                      0 //col:7207
IA32_SPEC_CTRL_IBRS_FLAG =                                     0x01 //col:7208
IA32_SPEC_CTRL_IBRS_MASK =                                     0x01 //col:7209
IA32_SPEC_CTRL_IBRS(_) =                                       (((_) >> 0) & 0x01) //col:7210
IA32_SPEC_CTRL_STIBP_BIT =                                     1 //col:7219
IA32_SPEC_CTRL_STIBP_FLAG =                                    0x02 //col:7220
IA32_SPEC_CTRL_STIBP_MASK =                                    0x01 //col:7221
IA32_SPEC_CTRL_STIBP(_) =                                      (((_) >> 1) & 0x01) //col:7222
IA32_SPEC_CTRL_SSBD_BIT =                                      2 //col:7231
IA32_SPEC_CTRL_SSBD_FLAG =                                     0x04 //col:7232
IA32_SPEC_CTRL_SSBD_MASK =                                     0x01 //col:7233
IA32_SPEC_CTRL_SSBD(_) =                                       (((_) >> 2) & 0x01) //col:7234
IA32_PRED_CMD =                                                0x00000049 //col:7247
IA32_PRED_CMD_IBPB_BIT =                                       0 //col:7258
IA32_PRED_CMD_IBPB_FLAG =                                      0x01 //col:7259
IA32_PRED_CMD_IBPB_MASK =                                      0x01 //col:7260
IA32_PRED_CMD_IBPB(_) =                                        (((_) >> 0) & 0x01) //col:7261
IA32_BIOS_UPDATE_TRIGGER =                                     0x00000079 //col:7278
IA32_BIOS_UPDATE_SIGNATURE =                                   0x0000008B //col:7288
IA32_BIOS_UPDATE_SIGNATURE_RESERVED_BIT =                      0 //col:7297
IA32_BIOS_UPDATE_SIGNATURE_RESERVED_FLAG =                     0xFFFFFFFF //col:7298
IA32_BIOS_UPDATE_SIGNATURE_RESERVED_MASK =                     0xFFFFFFFF //col:7299
IA32_BIOS_UPDATE_SIGNATURE_RESERVED(_) =                       (((_) >> 0) & 0xFFFFFFFF) //col:7300
IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_BIT =    32 //col:7313
IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_FLAG =   0xFFFFFFFF00000000 //col:7314
IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_MASK =   0xFFFFFFFF //col:7315
IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE(_) =     (((_) >> 32) & 0xFFFFFFFF) //col:7316
IA32_SGXLEPUBKEYHASH0 =                                        0x0000008C //col:7333
IA32_SGXLEPUBKEYHASH1 =                                        0x0000008D //col:7334
IA32_SGXLEPUBKEYHASH2 =                                        0x0000008E //col:7335
IA32_SGXLEPUBKEYHASH3 =                                        0x0000008F //col:7336
IA32_SMM_MONITOR_CTL =                                         0x0000009B //col:7347
IA32_SMM_MONITOR_CTL_VALID_BIT =                               0 //col:7363
IA32_SMM_MONITOR_CTL_VALID_FLAG =                              0x01 //col:7364
IA32_SMM_MONITOR_CTL_VALID_MASK =                              0x01 //col:7365
IA32_SMM_MONITOR_CTL_VALID(_) =                                (((_) >> 0) & 0x01) //col:7366
IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_BIT =            2 //col:7380
IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_FLAG =           0x04 //col:7381
IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_MASK =           0x01 //col:7382
IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF(_) =             (((_) >> 2) & 0x01) //col:7383
IA32_SMM_MONITOR_CTL_MSEG_BASE_BIT =                           12 //col:7394
IA32_SMM_MONITOR_CTL_MSEG_BASE_FLAG =                          0xFFFFF000 //col:7395
IA32_SMM_MONITOR_CTL_MSEG_BASE_MASK =                          0xFFFFF //col:7396
IA32_SMM_MONITOR_CTL_MSEG_BASE(_) =                            (((_) >> 12) & 0xFFFFF) //col:7397
IA32_STM_FEATURES_IA32E =                                      0x00000001 //col:7430
IA32_SMBASE =                                                  0x0000009E //col:7452
IA32_PMC0 =                                                    0x000000C1 //col:7462
IA32_PMC1 =                                                    0x000000C2 //col:7463
IA32_PMC2 =                                                    0x000000C3 //col:7464
IA32_PMC3 =                                                    0x000000C4 //col:7465
IA32_PMC4 =                                                    0x000000C5 //col:7466
IA32_PMC5 =                                                    0x000000C6 //col:7467
IA32_PMC6 =                                                    0x000000C7 //col:7468
IA32_PMC7 =                                                    0x000000C8 //col:7469
IA32_MPERF =                                                   0x000000E7 //col:7480
IA32_APERF =                                                   0x000000E8 //col:7498
IA32_MTRR_CAPABILITIES =                                       0x000000FE //col:7517
IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_BIT =              0 //col:7528
IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_FLAG =             0xFF //col:7529
IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_MASK =             0xFF //col:7530
IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT(_) =               (((_) >> 0) & 0xFF) //col:7531
IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_BIT =             8 //col:7540
IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_FLAG =            0x100 //col:7541
IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_MASK =            0x01 //col:7542
IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED(_) =              (((_) >> 8) & 0x01) //col:7543
IA32_MTRR_CAPABILITIES_WC_SUPPORTED_BIT =                      10 //col:7552
IA32_MTRR_CAPABILITIES_WC_SUPPORTED_FLAG =                     0x400 //col:7553
IA32_MTRR_CAPABILITIES_WC_SUPPORTED_MASK =                     0x01 //col:7554
IA32_MTRR_CAPABILITIES_WC_SUPPORTED(_) =                       (((_) >> 10) & 0x01) //col:7555
IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_BIT =                    11 //col:7564
IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_FLAG =                   0x800 //col:7565
IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_MASK =                   0x01 //col:7566
IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED(_) =                     (((_) >> 11) & 0x01) //col:7567
IA32_ARCH_CAPABILITIES =                                       0x0000010A //col:7580
IA32_ARCH_CAPABILITIES_RDCL_NO_BIT =                           0 //col:7589
IA32_ARCH_CAPABILITIES_RDCL_NO_FLAG =                          0x01 //col:7590
IA32_ARCH_CAPABILITIES_RDCL_NO_MASK =                          0x01 //col:7591
IA32_ARCH_CAPABILITIES_RDCL_NO(_) =                            (((_) >> 0) & 0x01) //col:7592
IA32_ARCH_CAPABILITIES_IBRS_ALL_BIT =                          1 //col:7598
IA32_ARCH_CAPABILITIES_IBRS_ALL_FLAG =                         0x02 //col:7599
IA32_ARCH_CAPABILITIES_IBRS_ALL_MASK =                         0x01 //col:7600
IA32_ARCH_CAPABILITIES_IBRS_ALL(_) =                           (((_) >> 1) & 0x01) //col:7601
IA32_ARCH_CAPABILITIES_RSBA_BIT =                              2 //col:7608
IA32_ARCH_CAPABILITIES_RSBA_FLAG =                             0x04 //col:7609
IA32_ARCH_CAPABILITIES_RSBA_MASK =                             0x01 //col:7610
IA32_ARCH_CAPABILITIES_RSBA(_) =                               (((_) >> 2) & 0x01) //col:7611
IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_BIT =                3 //col:7617
IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_FLAG =               0x08 //col:7618
IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_MASK =               0x01 //col:7619
IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY(_) =                 (((_) >> 3) & 0x01) //col:7620
IA32_ARCH_CAPABILITIES_SSB_NO_BIT =                            4 //col:7626
IA32_ARCH_CAPABILITIES_SSB_NO_FLAG =                           0x10 //col:7627
IA32_ARCH_CAPABILITIES_SSB_NO_MASK =                           0x01 //col:7628
IA32_ARCH_CAPABILITIES_SSB_NO(_) =                             (((_) >> 4) & 0x01) //col:7629
IA32_ARCH_CAPABILITIES_MDS_NO_BIT =                            5 //col:7635
IA32_ARCH_CAPABILITIES_MDS_NO_FLAG =                           0x20 //col:7636
IA32_ARCH_CAPABILITIES_MDS_NO_MASK =                           0x01 //col:7637
IA32_ARCH_CAPABILITIES_MDS_NO(_) =                             (((_) >> 5) & 0x01) //col:7638
IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_BIT =                 6 //col:7645
IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_FLAG =                0x40 //col:7646
IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_MASK =                0x01 //col:7647
IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO(_) =                  (((_) >> 6) & 0x01) //col:7648
IA32_ARCH_CAPABILITIES_TSX_CTRL_BIT =                          7 //col:7654
IA32_ARCH_CAPABILITIES_TSX_CTRL_FLAG =                         0x80 //col:7655
IA32_ARCH_CAPABILITIES_TSX_CTRL_MASK =                         0x01 //col:7656
IA32_ARCH_CAPABILITIES_TSX_CTRL(_) =                           (((_) >> 7) & 0x01) //col:7657
IA32_ARCH_CAPABILITIES_TAA_NO_BIT =                            8 //col:7663
IA32_ARCH_CAPABILITIES_TAA_NO_FLAG =                           0x100 //col:7664
IA32_ARCH_CAPABILITIES_TAA_NO_MASK =                           0x01 //col:7665
IA32_ARCH_CAPABILITIES_TAA_NO(_) =                             (((_) >> 8) & 0x01) //col:7666
IA32_FLUSH_CMD =                                               0x0000010B //col:7679
IA32_FLUSH_CMD_L1D_FLUSH_BIT =                                 0 //col:7690
IA32_FLUSH_CMD_L1D_FLUSH_FLAG =                                0x01 //col:7691
IA32_FLUSH_CMD_L1D_FLUSH_MASK =                                0x01 //col:7692
IA32_FLUSH_CMD_L1D_FLUSH(_) =                                  (((_) >> 0) & 0x01) //col:7693
IA32_TSX_CTRL =                                                0x00000122 //col:7707
IA32_TSX_CTRL_RTM_DISABLE_BIT =                                0 //col:7716
IA32_TSX_CTRL_RTM_DISABLE_FLAG =                               0x01 //col:7717
IA32_TSX_CTRL_RTM_DISABLE_MASK =                               0x01 //col:7718
IA32_TSX_CTRL_RTM_DISABLE(_) =                                 (((_) >> 0) & 0x01) //col:7719
IA32_TSX_CTRL_TSX_CPUID_CLEAR_BIT =                            1 //col:7726
IA32_TSX_CTRL_TSX_CPUID_CLEAR_FLAG =                           0x02 //col:7727
IA32_TSX_CTRL_TSX_CPUID_CLEAR_MASK =                           0x01 //col:7728
IA32_TSX_CTRL_TSX_CPUID_CLEAR(_) =                             (((_) >> 1) & 0x01) //col:7729
IA32_SYSENTER_CS =                                             0x00000174 //col:7746
IA32_SYSENTER_CS_CS_SELECTOR_BIT =                             0 //col:7755
IA32_SYSENTER_CS_CS_SELECTOR_FLAG =                            0xFFFF //col:7756
IA32_SYSENTER_CS_CS_SELECTOR_MASK =                            0xFFFF //col:7757
IA32_SYSENTER_CS_CS_SELECTOR(_) =                              (((_) >> 0) & 0xFFFF) //col:7758
IA32_SYSENTER_CS_NOT_USED_1_BIT =                              16 //col:7766
IA32_SYSENTER_CS_NOT_USED_1_FLAG =                             0xFFFF0000 //col:7767
IA32_SYSENTER_CS_NOT_USED_1_MASK =                             0xFFFF //col:7768
IA32_SYSENTER_CS_NOT_USED_1(_) =                               (((_) >> 16) & 0xFFFF) //col:7769
IA32_SYSENTER_CS_NOT_USED_2_BIT =                              32 //col:7777
IA32_SYSENTER_CS_NOT_USED_2_FLAG =                             0xFFFFFFFF00000000 //col:7778
IA32_SYSENTER_CS_NOT_USED_2_MASK =                             0xFFFFFFFF //col:7779
IA32_SYSENTER_CS_NOT_USED_2(_) =                               (((_) >> 32) & 0xFFFFFFFF) //col:7780
IA32_SYSENTER_ESP =                                            0x00000175 //col:7796
IA32_SYSENTER_EIP =                                            0x00000176 //col:7807
IA32_MCG_CAP =                                                 0x00000179 //col:7814
IA32_MCG_CAP_COUNT_BIT =                                       0 //col:7823
IA32_MCG_CAP_COUNT_FLAG =                                      0xFF //col:7824
IA32_MCG_CAP_COUNT_MASK =                                      0xFF //col:7825
IA32_MCG_CAP_COUNT(_) =                                        (((_) >> 0) & 0xFF) //col:7826
IA32_MCG_CAP_MCG_CTL_P_BIT =                                   8 //col:7832
IA32_MCG_CAP_MCG_CTL_P_FLAG =                                  0x100 //col:7833
IA32_MCG_CAP_MCG_CTL_P_MASK =                                  0x01 //col:7834
IA32_MCG_CAP_MCG_CTL_P(_) =                                    (((_) >> 8) & 0x01) //col:7835
IA32_MCG_CAP_MCG_EXT_P_BIT =                                   9 //col:7841
IA32_MCG_CAP_MCG_EXT_P_FLAG =                                  0x200 //col:7842
IA32_MCG_CAP_MCG_EXT_P_MASK =                                  0x01 //col:7843
IA32_MCG_CAP_MCG_EXT_P(_) =                                    (((_) >> 9) & 0x01) //col:7844
IA32_MCG_CAP_MCP_CMCI_P_BIT =                                  10 //col:7852
IA32_MCG_CAP_MCP_CMCI_P_FLAG =                                 0x400 //col:7853
IA32_MCG_CAP_MCP_CMCI_P_MASK =                                 0x01 //col:7854
IA32_MCG_CAP_MCP_CMCI_P(_) =                                   (((_) >> 10) & 0x01) //col:7855
IA32_MCG_CAP_MCG_TES_P_BIT =                                   11 //col:7861
IA32_MCG_CAP_MCG_TES_P_FLAG =                                  0x800 //col:7862
IA32_MCG_CAP_MCG_TES_P_MASK =                                  0x01 //col:7863
IA32_MCG_CAP_MCG_TES_P(_) =                                    (((_) >> 11) & 0x01) //col:7864
IA32_MCG_CAP_MCG_EXT_CNT_BIT =                                 16 //col:7871
IA32_MCG_CAP_MCG_EXT_CNT_FLAG =                                0xFF0000 //col:7872
IA32_MCG_CAP_MCG_EXT_CNT_MASK =                                0xFF //col:7873
IA32_MCG_CAP_MCG_EXT_CNT(_) =                                  (((_) >> 16) & 0xFF) //col:7874
IA32_MCG_CAP_MCG_SER_P_BIT =                                   24 //col:7880
IA32_MCG_CAP_MCG_SER_P_FLAG =                                  0x1000000 //col:7881
IA32_MCG_CAP_MCG_SER_P_MASK =                                  0x01 //col:7882
IA32_MCG_CAP_MCG_SER_P(_) =                                    (((_) >> 24) & 0x01) //col:7883
IA32_MCG_CAP_MCG_ELOG_P_BIT =                                  26 //col:7894
IA32_MCG_CAP_MCG_ELOG_P_FLAG =                                 0x4000000 //col:7895
IA32_MCG_CAP_MCG_ELOG_P_MASK =                                 0x01 //col:7896
IA32_MCG_CAP_MCG_ELOG_P(_) =                                   (((_) >> 26) & 0x01) //col:7897
IA32_MCG_CAP_MCG_LMCE_P_BIT =                                  27 //col:7906
IA32_MCG_CAP_MCG_LMCE_P_FLAG =                                 0x8000000 //col:7907
IA32_MCG_CAP_MCG_LMCE_P_MASK =                                 0x01 //col:7908
IA32_MCG_CAP_MCG_LMCE_P(_) =                                   (((_) >> 27) & 0x01) //col:7909
IA32_MCG_STATUS =                                              0x0000017A //col:7922
IA32_MCG_STATUS_RIPV_BIT =                                     0 //col:7933
IA32_MCG_STATUS_RIPV_FLAG =                                    0x01 //col:7934
IA32_MCG_STATUS_RIPV_MASK =                                    0x01 //col:7935
IA32_MCG_STATUS_RIPV(_) =                                      (((_) >> 0) & 0x01) //col:7936
IA32_MCG_STATUS_EIPV_BIT =                                     1 //col:7944
IA32_MCG_STATUS_EIPV_FLAG =                                    0x02 //col:7945
IA32_MCG_STATUS_EIPV_MASK =                                    0x01 //col:7946
IA32_MCG_STATUS_EIPV(_) =                                      (((_) >> 1) & 0x01) //col:7947
IA32_MCG_STATUS_MCIP_BIT =                                     2 //col:7955
IA32_MCG_STATUS_MCIP_FLAG =                                    0x04 //col:7956
IA32_MCG_STATUS_MCIP_MASK =                                    0x01 //col:7957
IA32_MCG_STATUS_MCIP(_) =                                      (((_) >> 2) & 0x01) //col:7958
IA32_MCG_STATUS_LMCE_S_BIT =                                   3 //col:7964
IA32_MCG_STATUS_LMCE_S_FLAG =                                  0x08 //col:7965
IA32_MCG_STATUS_LMCE_S_MASK =                                  0x01 //col:7966
IA32_MCG_STATUS_LMCE_S(_) =                                    (((_) >> 3) & 0x01) //col:7967
IA32_MCG_CTL =                                                 0x0000017B //col:7980
IA32_PERFEVTSEL0 =                                             0x00000186 //col:7990
IA32_PERFEVTSEL1 =                                             0x00000187 //col:7991
IA32_PERFEVTSEL2 =                                             0x00000188 //col:7992
IA32_PERFEVTSEL3 =                                             0x00000189 //col:7993
IA32_PERFEVTSEL_EVENT_SELECT_BIT =                             0 //col:8002
IA32_PERFEVTSEL_EVENT_SELECT_FLAG =                            0xFF //col:8003
IA32_PERFEVTSEL_EVENT_SELECT_MASK =                            0xFF //col:8004
IA32_PERFEVTSEL_EVENT_SELECT(_) =                              (((_) >> 0) & 0xFF) //col:8005
IA32_PERFEVTSEL_U_MASK_BIT =                                   8 //col:8011
IA32_PERFEVTSEL_U_MASK_FLAG =                                  0xFF00 //col:8012
IA32_PERFEVTSEL_U_MASK_MASK =                                  0xFF //col:8013
IA32_PERFEVTSEL_U_MASK(_) =                                    (((_) >> 8) & 0xFF) //col:8014
IA32_PERFEVTSEL_USR_BIT =                                      16 //col:8020
IA32_PERFEVTSEL_USR_FLAG =                                     0x10000 //col:8021
IA32_PERFEVTSEL_USR_MASK =                                     0x01 //col:8022
IA32_PERFEVTSEL_USR(_) =                                       (((_) >> 16) & 0x01) //col:8023
IA32_PERFEVTSEL_OS_BIT =                                       17 //col:8029
IA32_PERFEVTSEL_OS_FLAG =                                      0x20000 //col:8030
IA32_PERFEVTSEL_OS_MASK =                                      0x01 //col:8031
IA32_PERFEVTSEL_OS(_) =                                        (((_) >> 17) & 0x01) //col:8032
IA32_PERFEVTSEL_EDGE_BIT =                                     18 //col:8038
IA32_PERFEVTSEL_EDGE_FLAG =                                    0x40000 //col:8039
IA32_PERFEVTSEL_EDGE_MASK =                                    0x01 //col:8040
IA32_PERFEVTSEL_EDGE(_) =                                      (((_) >> 18) & 0x01) //col:8041
IA32_PERFEVTSEL_PC_BIT =                                       19 //col:8047
IA32_PERFEVTSEL_PC_FLAG =                                      0x80000 //col:8048
IA32_PERFEVTSEL_PC_MASK =                                      0x01 //col:8049
IA32_PERFEVTSEL_PC(_) =                                        (((_) >> 19) & 0x01) //col:8050
IA32_PERFEVTSEL_INTR_BIT =                                     20 //col:8056
IA32_PERFEVTSEL_INTR_FLAG =                                    0x100000 //col:8057
IA32_PERFEVTSEL_INTR_MASK =                                    0x01 //col:8058
IA32_PERFEVTSEL_INTR(_) =                                      (((_) >> 20) & 0x01) //col:8059
IA32_PERFEVTSEL_ANY_THREAD_BIT =                               21 //col:8067
IA32_PERFEVTSEL_ANY_THREAD_FLAG =                              0x200000 //col:8068
IA32_PERFEVTSEL_ANY_THREAD_MASK =                              0x01 //col:8069
IA32_PERFEVTSEL_ANY_THREAD(_) =                                (((_) >> 21) & 0x01) //col:8070
IA32_PERFEVTSEL_EN_BIT =                                       22 //col:8076
IA32_PERFEVTSEL_EN_FLAG =                                      0x400000 //col:8077
IA32_PERFEVTSEL_EN_MASK =                                      0x01 //col:8078
IA32_PERFEVTSEL_EN(_) =                                        (((_) >> 22) & 0x01) //col:8079
IA32_PERFEVTSEL_INV_BIT =                                      23 //col:8085
IA32_PERFEVTSEL_INV_FLAG =                                     0x800000 //col:8086
IA32_PERFEVTSEL_INV_MASK =                                     0x01 //col:8087
IA32_PERFEVTSEL_INV(_) =                                       (((_) >> 23) & 0x01) //col:8088
IA32_PERFEVTSEL_CMASK_BIT =                                    24 //col:8095
IA32_PERFEVTSEL_CMASK_FLAG =                                   0xFF000000 //col:8096
IA32_PERFEVTSEL_CMASK_MASK =                                   0xFF //col:8097
IA32_PERFEVTSEL_CMASK(_) =                                     (((_) >> 24) & 0xFF) //col:8098
IA32_PERF_STATUS =                                             0x00000198 //col:8116
IA32_PERF_STATUS_STATE_VALUE_BIT =                             0 //col:8125
IA32_PERF_STATUS_STATE_VALUE_FLAG =                            0xFFFF //col:8126
IA32_PERF_STATUS_STATE_VALUE_MASK =                            0xFFFF //col:8127
IA32_PERF_STATUS_STATE_VALUE(_) =                              (((_) >> 0) & 0xFFFF) //col:8128
IA32_PERF_CTL =                                                0x00000199 //col:8144
IA32_PERF_CTL_TARGET_STATE_VALUE_BIT =                         0 //col:8153
IA32_PERF_CTL_TARGET_STATE_VALUE_FLAG =                        0xFFFF //col:8154
IA32_PERF_CTL_TARGET_STATE_VALUE_MASK =                        0xFFFF //col:8155
IA32_PERF_CTL_TARGET_STATE_VALUE(_) =                          (((_) >> 0) & 0xFFFF) //col:8156
IA32_PERF_CTL_IDA_ENGAGE_BIT =                                 32 //col:8165
IA32_PERF_CTL_IDA_ENGAGE_FLAG =                                0x100000000 //col:8166
IA32_PERF_CTL_IDA_ENGAGE_MASK =                                0x01 //col:8167
IA32_PERF_CTL_IDA_ENGAGE(_) =                                  (((_) >> 32) & 0x01) //col:8168
IA32_CLOCK_MODULATION =                                        0x0000019A //col:8182
IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_BIT = 0 //col:8193
IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_FLAG = 0x01 //col:8194
IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_MASK = 0x01 //col:8195
IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE(_) = (((_) >> 0) & 0x01) //col:8196
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_BIT = 1 //col:8206
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_FLAG = 0x0E //col:8207
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_MASK = 0x07 //col:8208
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE(_) = (((_) >> 1) & 0x07) //col:8209
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_BIT =  4 //col:8219
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_FLAG = 0x10 //col:8220
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_MASK = 0x01 //col:8221
IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE(_) =   (((_) >> 4) & 0x01) //col:8222
IA32_THERM_INTERRUPT =                                         0x0000019B //col:8239
IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_BIT =   0 //col:8250
IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_FLAG =  0x01 //col:8251
IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_MASK =  0x01 //col:8252
IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE(_) =    (((_) >> 0) & 0x01) //col:8253
IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_BIT =    1 //col:8261
IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_FLAG =   0x02 //col:8262
IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_MASK =   0x01 //col:8263
IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE(_) =     (((_) >> 1) & 0x01) //col:8264
IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_BIT =            2 //col:8272
IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_FLAG =           0x04 //col:8273
IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_MASK =           0x01 //col:8274
IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE(_) =             (((_) >> 2) & 0x01) //col:8275
IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_BIT =            3 //col:8283
IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_FLAG =           0x08 //col:8284
IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_MASK =           0x01 //col:8285
IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE(_) =             (((_) >> 3) & 0x01) //col:8286
IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_BIT = 4 //col:8294
IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_FLAG = 0x10 //col:8295
IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_MASK = 0x01 //col:8296
IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE(_) = (((_) >> 4) & 0x01) //col:8297
IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_BIT =                    8 //col:8306
IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_FLAG =                   0x7F00 //col:8307
IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_MASK =                   0x7F //col:8308
IA32_THERM_INTERRUPT_THRESHOLD1_VALUE(_) =                     (((_) >> 8) & 0x7F) //col:8309
IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_BIT =         15 //col:8317
IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_FLAG =        0x8000 //col:8318
IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_MASK =        0x01 //col:8319
IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE(_) =          (((_) >> 15) & 0x01) //col:8320
IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_BIT =                    16 //col:8328
IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_FLAG =                   0x7F0000 //col:8329
IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_MASK =                   0x7F //col:8330
IA32_THERM_INTERRUPT_THRESHOLD2_VALUE(_) =                     (((_) >> 16) & 0x7F) //col:8331
IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_BIT =         23 //col:8339
IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_FLAG =        0x800000 //col:8340
IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_MASK =        0x01 //col:8341
IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE(_) =          (((_) >> 23) & 0x01) //col:8342
IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_BIT =     24 //col:8350
IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_FLAG =    0x1000000 //col:8351
IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_MASK =    0x01 //col:8352
IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE(_) =      (((_) >> 24) & 0x01) //col:8353
IA32_THERM_STATUS =                                            0x0000019C //col:8370
IA32_THERM_STATUS_THERMAL_STATUS_BIT =                         0 //col:8381
IA32_THERM_STATUS_THERMAL_STATUS_FLAG =                        0x01 //col:8382
IA32_THERM_STATUS_THERMAL_STATUS_MASK =                        0x01 //col:8383
IA32_THERM_STATUS_THERMAL_STATUS(_) =                          (((_) >> 0) & 0x01) //col:8384
IA32_THERM_STATUS_THERMAL_STATUS_LOG_BIT =                     1 //col:8392
IA32_THERM_STATUS_THERMAL_STATUS_LOG_FLAG =                    0x02 //col:8393
IA32_THERM_STATUS_THERMAL_STATUS_LOG_MASK =                    0x01 //col:8394
IA32_THERM_STATUS_THERMAL_STATUS_LOG(_) =                      (((_) >> 1) & 0x01) //col:8395
IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_BIT =                  2 //col:8403
IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_FLAG =                 0x04 //col:8404
IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_MASK =                 0x01 //col:8405
IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT(_) =                   (((_) >> 2) & 0x01) //col:8406
IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_BIT =                    3 //col:8414
IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_FLAG =                   0x08 //col:8415
IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_MASK =                   0x01 //col:8416
IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG(_) =                     (((_) >> 3) & 0x01) //col:8417
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_BIT =            4 //col:8425
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_FLAG =           0x10 //col:8426
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_MASK =           0x01 //col:8427
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS(_) =             (((_) >> 4) & 0x01) //col:8428
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_BIT =        5 //col:8436
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_FLAG =       0x20 //col:8437
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_MASK =       0x01 //col:8438
IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG(_) =         (((_) >> 5) & 0x01) //col:8439
IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_BIT =              6 //col:8447
IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_FLAG =             0x40 //col:8448
IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_MASK =             0x01 //col:8449
IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS(_) =               (((_) >> 6) & 0x01) //col:8450
IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_BIT =                 7 //col:8458
IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_FLAG =                0x80 //col:8459
IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_MASK =                0x01 //col:8460
IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG(_) =                  (((_) >> 7) & 0x01) //col:8461
IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_BIT =              8 //col:8469
IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_FLAG =             0x100 //col:8470
IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_MASK =             0x01 //col:8471
IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS(_) =               (((_) >> 8) & 0x01) //col:8472
IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_BIT =                 9 //col:8480
IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_FLAG =                0x200 //col:8481
IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_MASK =                0x01 //col:8482
IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG(_) =                  (((_) >> 9) & 0x01) //col:8483
IA32_THERM_STATUS_POWER_LIMITATION_STATUS_BIT =                10 //col:8491
IA32_THERM_STATUS_POWER_LIMITATION_STATUS_FLAG =               0x400 //col:8492
IA32_THERM_STATUS_POWER_LIMITATION_STATUS_MASK =               0x01 //col:8493
IA32_THERM_STATUS_POWER_LIMITATION_STATUS(_) =                 (((_) >> 10) & 0x01) //col:8494
IA32_THERM_STATUS_POWER_LIMITATION_LOG_BIT =                   11 //col:8502
IA32_THERM_STATUS_POWER_LIMITATION_LOG_FLAG =                  0x800 //col:8503
IA32_THERM_STATUS_POWER_LIMITATION_LOG_MASK =                  0x01 //col:8504
IA32_THERM_STATUS_POWER_LIMITATION_LOG(_) =                    (((_) >> 11) & 0x01) //col:8505
IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_BIT =                   12 //col:8513
IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_FLAG =                  0x1000 //col:8514
IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_MASK =                  0x01 //col:8515
IA32_THERM_STATUS_CURRENT_LIMIT_STATUS(_) =                    (((_) >> 12) & 0x01) //col:8516
IA32_THERM_STATUS_CURRENT_LIMIT_LOG_BIT =                      13 //col:8524
IA32_THERM_STATUS_CURRENT_LIMIT_LOG_FLAG =                     0x2000 //col:8525
IA32_THERM_STATUS_CURRENT_LIMIT_LOG_MASK =                     0x01 //col:8526
IA32_THERM_STATUS_CURRENT_LIMIT_LOG(_) =                       (((_) >> 13) & 0x01) //col:8527
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_BIT =              14 //col:8535
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_FLAG =             0x4000 //col:8536
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_MASK =             0x01 //col:8537
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS(_) =               (((_) >> 14) & 0x01) //col:8538
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_BIT =                 15 //col:8546
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_FLAG =                0x8000 //col:8547
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_MASK =                0x01 //col:8548
IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG(_) =                  (((_) >> 15) & 0x01) //col:8549
IA32_THERM_STATUS_DIGITAL_READOUT_BIT =                        16 //col:8557
IA32_THERM_STATUS_DIGITAL_READOUT_FLAG =                       0x7F0000 //col:8558
IA32_THERM_STATUS_DIGITAL_READOUT_MASK =                       0x7F //col:8559
IA32_THERM_STATUS_DIGITAL_READOUT(_) =                         (((_) >> 16) & 0x7F) //col:8560
IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_BIT =          27 //col:8569
IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_FLAG =         0x78000000 //col:8570
IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_MASK =         0x0F //col:8571
IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS(_) =           (((_) >> 27) & 0x0F) //col:8572
IA32_THERM_STATUS_READING_VALID_BIT =                          31 //col:8580
IA32_THERM_STATUS_READING_VALID_FLAG =                         0x80000000 //col:8581
IA32_THERM_STATUS_READING_VALID_MASK =                         0x01 //col:8582
IA32_THERM_STATUS_READING_VALID(_) =                           (((_) >> 31) & 0x01) //col:8583
IA32_MISC_ENABLE =                                             0x000001A0 //col:8596
IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_BIT =                     0 //col:8610
IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_FLAG =                    0x01 //col:8611
IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_MASK =                    0x01 //col:8612
IA32_MISC_ENABLE_FAST_STRINGS_ENABLE(_) =                      (((_) >> 0) & 0x01) //col:8613
IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_BIT = 3 //col:8628
IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_FLAG = 0x08 //col:8629
IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_MASK = 0x01 //col:8630
IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE(_) = (((_) >> 3) & 0x01) //col:8631
IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_BIT =        7 //col:8643
IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_FLAG =       0x80 //col:8644
IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_MASK =       0x01 //col:8645
IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE(_) =         (((_) >> 7) & 0x01) //col:8646
IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_BIT =        11 //col:8658
IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_FLAG =       0x800 //col:8659
IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_MASK =       0x01 //col:8660
IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE(_) =         (((_) >> 11) & 0x01) //col:8661
IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_BIT = 12 //col:8672
IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_FLAG = 0x1000 //col:8673
IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_MASK = 0x01 //col:8674
IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE(_) = (((_) >> 12) & 0x01) //col:8675
IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_BIT = 16 //col:8687
IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_FLAG = 0x10000 //col:8688
IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_MASK = 0x01 //col:8689
IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE(_) = (((_) >> 16) & 0x01) //col:8690
IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_BIT =                      18 //col:8705
IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_FLAG =                     0x40000 //col:8706
IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_MASK =                     0x01 //col:8707
IA32_MISC_ENABLE_ENABLE_MONITOR_FSM(_) =                       (((_) >> 18) & 0x01) //col:8708
IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_BIT =                      22 //col:8725
IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_FLAG =                     0x400000 //col:8726
IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_MASK =                     0x01 //col:8727
IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL(_) =                       (((_) >> 22) & 0x01) //col:8728
IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_BIT =                    23 //col:8739
IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_FLAG =                   0x800000 //col:8740
IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_MASK =                   0x01 //col:8741
IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE(_) =                     (((_) >> 23) & 0x01) //col:8742
IA32_MISC_ENABLE_XD_BIT_DISABLE_BIT =                          34 //col:8758
IA32_MISC_ENABLE_XD_BIT_DISABLE_FLAG =                         0x400000000 //col:8759
IA32_MISC_ENABLE_XD_BIT_DISABLE_MASK =                         0x01 //col:8760
IA32_MISC_ENABLE_XD_BIT_DISABLE(_) =                           (((_) >> 34) & 0x01) //col:8761
IA32_ENERGY_PERF_BIAS =                                        0x000001B0 //col:8774
IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_BIT =            0 //col:8786
IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_FLAG =           0x0F //col:8787
IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_MASK =           0x0F //col:8788
IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE(_) =             (((_) >> 0) & 0x0F) //col:8789
IA32_PACKAGE_THERM_STATUS =                                    0x000001B1 //col:8805
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_BIT =                 0 //col:8814
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_FLAG =                0x01 //col:8815
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_MASK =                0x01 //col:8816
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS(_) =                  (((_) >> 0) & 0x01) //col:8817
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_BIT =             1 //col:8823
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_FLAG =            0x02 //col:8824
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_MASK =            0x01 //col:8825
IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG(_) =              (((_) >> 1) & 0x01) //col:8826
IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_BIT =                  2 //col:8832
IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_FLAG =                 0x04 //col:8833
IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_MASK =                 0x01 //col:8834
IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT(_) =                   (((_) >> 2) & 0x01) //col:8835
IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_BIT =                    3 //col:8841
IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_FLAG =                   0x08 //col:8842
IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_MASK =                   0x01 //col:8843
IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG(_) =                     (((_) >> 3) & 0x01) //col:8844
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_BIT =    4 //col:8850
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_FLAG =   0x10 //col:8851
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_MASK =   0x01 //col:8852
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS(_) =     (((_) >> 4) & 0x01) //col:8853
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_BIT = 5 //col:8859
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_FLAG = 0x20 //col:8860
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_MASK = 0x01 //col:8861
IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG(_) = (((_) >> 5) & 0x01) //col:8862
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_BIT =      6 //col:8868
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_FLAG =     0x40 //col:8869
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_MASK =     0x01 //col:8870
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS(_) =       (((_) >> 6) & 0x01) //col:8871
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_BIT =         7 //col:8877
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_FLAG =        0x80 //col:8878
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_MASK =        0x01 //col:8879
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG(_) =          (((_) >> 7) & 0x01) //col:8880
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_BIT =      8 //col:8886
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_FLAG =     0x100 //col:8887
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_MASK =     0x01 //col:8888
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS(_) =       (((_) >> 8) & 0x01) //col:8889
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_BIT =         9 //col:8895
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_FLAG =        0x200 //col:8896
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_MASK =        0x01 //col:8897
IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG(_) =          (((_) >> 9) & 0x01) //col:8898
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_BIT =        10 //col:8904
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_FLAG =       0x400 //col:8905
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_MASK =       0x01 //col:8906
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS(_) =         (((_) >> 10) & 0x01) //col:8907
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_BIT =           11 //col:8913
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_FLAG =          0x800 //col:8914
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_MASK =          0x01 //col:8915
IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG(_) =            (((_) >> 11) & 0x01) //col:8916
IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_BIT =                16 //col:8923
IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_FLAG =               0x7F0000 //col:8924
IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_MASK =               0x7F //col:8925
IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT(_) =                 (((_) >> 16) & 0x7F) //col:8926
IA32_PACKAGE_THERM_INTERRUPT =                                 0x000001B2 //col:8943
IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_BIT = 0 //col:8952
IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_FLAG = 0x01 //col:8953
IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_MASK = 0x01 //col:8954
IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE(_) = (((_) >> 0) & 0x01) //col:8955
IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_BIT = 1 //col:8961
IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_FLAG = 0x02 //col:8962
IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_MASK = 0x01 //col:8963
IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE(_) = (((_) >> 1) & 0x01) //col:8964
IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_BIT =    2 //col:8970
IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_FLAG =   0x04 //col:8971
IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_MASK =   0x01 //col:8972
IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE(_) =     (((_) >> 2) & 0x01) //col:8973
IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_BIT =   4 //col:8980
IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_FLAG =  0x10 //col:8981
IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_MASK =  0x01 //col:8982
IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE(_) =    (((_) >> 4) & 0x01) //col:8983
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_BIT =            8 //col:8990
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_FLAG =           0x7F00 //col:8991
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_MASK =           0x7F //col:8992
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE(_) =             (((_) >> 8) & 0x7F) //col:8993
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_BIT = 15 //col:8999
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_FLAG = 0x8000 //col:9000
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_MASK = 0x01 //col:9001
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE(_) =  (((_) >> 15) & 0x01) //col:9002
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_BIT =            16 //col:9008
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_FLAG =           0x7F0000 //col:9009
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_MASK =           0x7F //col:9010
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE(_) =             (((_) >> 16) & 0x7F) //col:9011
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_BIT = 23 //col:9017
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_FLAG = 0x800000 //col:9018
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_MASK = 0x01 //col:9019
IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE(_) =  (((_) >> 23) & 0x01) //col:9020
IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_BIT = 24 //col:9026
IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_FLAG = 0x1000000 //col:9027
IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_MASK = 0x01 //col:9028
IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE(_) = (((_) >> 24) & 0x01) //col:9029
IA32_DEBUGCTL =                                                0x000001D9 //col:9042
IA32_DEBUGCTL_LBR_BIT =                                        0 //col:9054
IA32_DEBUGCTL_LBR_FLAG =                                       0x01 //col:9055
IA32_DEBUGCTL_LBR_MASK =                                       0x01 //col:9056
IA32_DEBUGCTL_LBR(_) =                                         (((_) >> 0) & 0x01) //col:9057
IA32_DEBUGCTL_BTF_BIT =                                        1 //col:9066
IA32_DEBUGCTL_BTF_FLAG =                                       0x02 //col:9067
IA32_DEBUGCTL_BTF_MASK =                                       0x01 //col:9068
IA32_DEBUGCTL_BTF(_) =                                         (((_) >> 1) & 0x01) //col:9069
IA32_DEBUGCTL_TR_BIT =                                         6 //col:9078
IA32_DEBUGCTL_TR_FLAG =                                        0x40 //col:9079
IA32_DEBUGCTL_TR_MASK =                                        0x01 //col:9080
IA32_DEBUGCTL_TR(_) =                                          (((_) >> 6) & 0x01) //col:9081
IA32_DEBUGCTL_BTS_BIT =                                        7 //col:9089
IA32_DEBUGCTL_BTS_FLAG =                                       0x80 //col:9090
IA32_DEBUGCTL_BTS_MASK =                                       0x01 //col:9091
IA32_DEBUGCTL_BTS(_) =                                         (((_) >> 7) & 0x01) //col:9092
IA32_DEBUGCTL_BTINT_BIT =                                      8 //col:9101
IA32_DEBUGCTL_BTINT_FLAG =                                     0x100 //col:9102
IA32_DEBUGCTL_BTINT_MASK =                                     0x01 //col:9103
IA32_DEBUGCTL_BTINT(_) =                                       (((_) >> 8) & 0x01) //col:9104
IA32_DEBUGCTL_BTS_OFF_OS_BIT =                                 9 //col:9112
IA32_DEBUGCTL_BTS_OFF_OS_FLAG =                                0x200 //col:9113
IA32_DEBUGCTL_BTS_OFF_OS_MASK =                                0x01 //col:9114
IA32_DEBUGCTL_BTS_OFF_OS(_) =                                  (((_) >> 9) & 0x01) //col:9115
IA32_DEBUGCTL_BTS_OFF_USR_BIT =                                10 //col:9123
IA32_DEBUGCTL_BTS_OFF_USR_FLAG =                               0x400 //col:9124
IA32_DEBUGCTL_BTS_OFF_USR_MASK =                               0x01 //col:9125
IA32_DEBUGCTL_BTS_OFF_USR(_) =                                 (((_) >> 10) & 0x01) //col:9126
IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_BIT =                         11 //col:9134
IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_FLAG =                        0x800 //col:9135
IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_MASK =                        0x01 //col:9136
IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI(_) =                          (((_) >> 11) & 0x01) //col:9137
IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_BIT =                      12 //col:9145
IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_FLAG =                     0x1000 //col:9146
IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_MASK =                     0x01 //col:9147
IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI(_) =                       (((_) >> 12) & 0x01) //col:9148
IA32_DEBUGCTL_ENABLE_UNCORE_PMI_BIT =                          13 //col:9156
IA32_DEBUGCTL_ENABLE_UNCORE_PMI_FLAG =                         0x2000 //col:9157
IA32_DEBUGCTL_ENABLE_UNCORE_PMI_MASK =                         0x01 //col:9158
IA32_DEBUGCTL_ENABLE_UNCORE_PMI(_) =                           (((_) >> 13) & 0x01) //col:9159
IA32_DEBUGCTL_FREEZE_WHILE_SMM_BIT =                           14 //col:9167
IA32_DEBUGCTL_FREEZE_WHILE_SMM_FLAG =                          0x4000 //col:9168
IA32_DEBUGCTL_FREEZE_WHILE_SMM_MASK =                          0x01 //col:9169
IA32_DEBUGCTL_FREEZE_WHILE_SMM(_) =                            (((_) >> 14) & 0x01) //col:9170
IA32_DEBUGCTL_RTM_DEBUG_BIT =                                  15 //col:9178
IA32_DEBUGCTL_RTM_DEBUG_FLAG =                                 0x8000 //col:9179
IA32_DEBUGCTL_RTM_DEBUG_MASK =                                 0x01 //col:9180
IA32_DEBUGCTL_RTM_DEBUG(_) =                                   (((_) >> 15) & 0x01) //col:9181
IA32_SMRR_PHYSBASE =                                           0x000001F2 //col:9196
IA32_SMRR_PHYSBASE_TYPE_BIT =                                  0 //col:9207
IA32_SMRR_PHYSBASE_TYPE_FLAG =                                 0xFF //col:9208
IA32_SMRR_PHYSBASE_TYPE_MASK =                                 0xFF //col:9209
IA32_SMRR_PHYSBASE_TYPE(_) =                                   (((_) >> 0) & 0xFF) //col:9210
IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_BIT =            12 //col:9217
IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_FLAG =           0xFFFFF000 //col:9218
IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_MASK =           0xFFFFF //col:9219
IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS(_) =             (((_) >> 12) & 0xFFFFF) //col:9220
IA32_SMRR_PHYSMASK =                                           0x000001F3 //col:9235
IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_BIT =                     11 //col:9246
IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_FLAG =                    0x800 //col:9247
IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_MASK =                    0x01 //col:9248
IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK(_) =                      (((_) >> 11) & 0x01) //col:9249
IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_BIT =               12 //col:9255
IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_FLAG =              0xFFFFF000 //col:9256
IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_MASK =              0xFFFFF //col:9257
IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK(_) =                (((_) >> 12) & 0xFFFFF) //col:9258
IA32_PLATFORM_DCA_CAP =                                        0x000001F8 //col:9271
IA32_CPU_DCA_CAP =                                             0x000001F9 //col:9278
IA32_DCA_0_CAP =                                               0x000001FA //col:9285
IA32_DCA_0_CAP_DCA_ACTIVE_BIT =                                0 //col:9294
IA32_DCA_0_CAP_DCA_ACTIVE_FLAG =                               0x01 //col:9295
IA32_DCA_0_CAP_DCA_ACTIVE_MASK =                               0x01 //col:9296
IA32_DCA_0_CAP_DCA_ACTIVE(_) =                                 (((_) >> 0) & 0x01) //col:9297
IA32_DCA_0_CAP_TRANSACTION_BIT =                               1 //col:9303
IA32_DCA_0_CAP_TRANSACTION_FLAG =                              0x06 //col:9304
IA32_DCA_0_CAP_TRANSACTION_MASK =                              0x03 //col:9305
IA32_DCA_0_CAP_TRANSACTION(_) =                                (((_) >> 1) & 0x03) //col:9306
IA32_DCA_0_CAP_DCA_TYPE_BIT =                                  3 //col:9312
IA32_DCA_0_CAP_DCA_TYPE_FLAG =                                 0x78 //col:9313
IA32_DCA_0_CAP_DCA_TYPE_MASK =                                 0x0F //col:9314
IA32_DCA_0_CAP_DCA_TYPE(_) =                                   (((_) >> 3) & 0x0F) //col:9315
IA32_DCA_0_CAP_DCA_QUEUE_SIZE_BIT =                            7 //col:9321
IA32_DCA_0_CAP_DCA_QUEUE_SIZE_FLAG =                           0x780 //col:9322
IA32_DCA_0_CAP_DCA_QUEUE_SIZE_MASK =                           0x0F //col:9323
IA32_DCA_0_CAP_DCA_QUEUE_SIZE(_) =                             (((_) >> 7) & 0x0F) //col:9324
IA32_DCA_0_CAP_DCA_DELAY_BIT =                                 13 //col:9331
IA32_DCA_0_CAP_DCA_DELAY_FLAG =                                0x1E000 //col:9332
IA32_DCA_0_CAP_DCA_DELAY_MASK =                                0x0F //col:9333
IA32_DCA_0_CAP_DCA_DELAY(_) =                                  (((_) >> 13) & 0x0F) //col:9334
IA32_DCA_0_CAP_SW_BLOCK_BIT =                                  24 //col:9341
IA32_DCA_0_CAP_SW_BLOCK_FLAG =                                 0x1000000 //col:9342
IA32_DCA_0_CAP_SW_BLOCK_MASK =                                 0x01 //col:9343
IA32_DCA_0_CAP_SW_BLOCK(_) =                                   (((_) >> 24) & 0x01) //col:9344
IA32_DCA_0_CAP_HW_BLOCK_BIT =                                  26 //col:9351
IA32_DCA_0_CAP_HW_BLOCK_FLAG =                                 0x4000000 //col:9352
IA32_DCA_0_CAP_HW_BLOCK_MASK =                                 0x01 //col:9353
IA32_DCA_0_CAP_HW_BLOCK(_) =                                   (((_) >> 26) & 0x01) //col:9354
IA32_MTRR_PHYSBASE_TYPE_BIT =                                  0 //col:9379
IA32_MTRR_PHYSBASE_TYPE_FLAG =                                 0xFF //col:9380
IA32_MTRR_PHYSBASE_TYPE_MASK =                                 0xFF //col:9381
IA32_MTRR_PHYSBASE_TYPE(_) =                                   (((_) >> 0) & 0xFF) //col:9382
IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_BIT =                     12 //col:9391
IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_FLAG =                    0xFFFFFFFFF000 //col:9392
IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_MASK =                    0xFFFFFFFFF //col:9393
IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER(_) =                      (((_) >> 12) & 0xFFFFFFFFF) //col:9394
IA32_MTRR_PHYSBASE0 =                                          0x00000200 //col:9401
IA32_MTRR_PHYSBASE1 =                                          0x00000202 //col:9402
IA32_MTRR_PHYSBASE2 =                                          0x00000204 //col:9403
IA32_MTRR_PHYSBASE3 =                                          0x00000206 //col:9404
IA32_MTRR_PHYSBASE4 =                                          0x00000208 //col:9405
IA32_MTRR_PHYSBASE5 =                                          0x0000020A //col:9406
IA32_MTRR_PHYSBASE6 =                                          0x0000020C //col:9407
IA32_MTRR_PHYSBASE7 =                                          0x0000020E //col:9408
IA32_MTRR_PHYSBASE8 =                                          0x00000210 //col:9409
IA32_MTRR_PHYSBASE9 =                                          0x00000212 //col:9410
IA32_MTRR_PHYSMASK_VALID_BIT =                                 11 //col:9435
IA32_MTRR_PHYSMASK_VALID_FLAG =                                0x800 //col:9436
IA32_MTRR_PHYSMASK_VALID_MASK =                                0x01 //col:9437
IA32_MTRR_PHYSMASK_VALID(_) =                                  (((_) >> 11) & 0x01) //col:9438
IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_BIT =                     12 //col:9453
IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_FLAG =                    0xFFFFFFFFF000 //col:9454
IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_MASK =                    0xFFFFFFFFF //col:9455
IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER(_) =                      (((_) >> 12) & 0xFFFFFFFFF) //col:9456
IA32_MTRR_PHYSMASK0 =                                          0x00000201 //col:9463
IA32_MTRR_PHYSMASK1 =                                          0x00000203 //col:9464
IA32_MTRR_PHYSMASK2 =                                          0x00000205 //col:9465
IA32_MTRR_PHYSMASK3 =                                          0x00000207 //col:9466
IA32_MTRR_PHYSMASK4 =                                          0x00000209 //col:9467
IA32_MTRR_PHYSMASK5 =                                          0x0000020B //col:9468
IA32_MTRR_PHYSMASK6 =                                          0x0000020D //col:9469
IA32_MTRR_PHYSMASK7 =                                          0x0000020F //col:9470
IA32_MTRR_PHYSMASK8 =                                          0x00000211 //col:9471
IA32_MTRR_PHYSMASK9 =                                          0x00000213 //col:9472
IA32_MTRR_FIX64K_BASE =                                        0x00000000 //col:9494
IA32_MTRR_FIX64K_SIZE =                                        0x00010000 //col:9495
IA32_MTRR_FIX64K_00000 =                                       0x00000250 //col:9496
IA32_MTRR_FIX16K_BASE =                                        0x00080000 //col:9508
IA32_MTRR_FIX16K_SIZE =                                        0x00004000 //col:9509
IA32_MTRR_FIX16K_80000 =                                       0x00000258 //col:9510
IA32_MTRR_FIX16K_A0000 =                                       0x00000259 //col:9511
IA32_MTRR_FIX4K_BASE =                                         0x000C0000 //col:9523
IA32_MTRR_FIX4K_SIZE =                                         0x00001000 //col:9524
IA32_MTRR_FIX4K_C0000 =                                        0x00000268 //col:9525
IA32_MTRR_FIX4K_C8000 =                                        0x00000269 //col:9526
IA32_MTRR_FIX4K_D0000 =                                        0x0000026A //col:9527
IA32_MTRR_FIX4K_D8000 =                                        0x0000026B //col:9528
IA32_MTRR_FIX4K_E0000 =                                        0x0000026C //col:9529
IA32_MTRR_FIX4K_E8000 =                                        0x0000026D //col:9530
IA32_MTRR_FIX4K_F0000 =                                        0x0000026E //col:9531
IA32_MTRR_FIX4K_F8000 =                                        0x0000026F //col:9532
IA32_MTRR_FIX_COUNT =                                          ((1 + 2 + 8) * 8) //col:9540
IA32_MTRR_VARIABLE_COUNT =                                     0x0000000A //col:9547
IA32_MTRR_COUNT =                                              (IA32_MTRR_FIX_COUNT + IA32_MTRR_VARIABLE_COUNT) //col:9552
IA32_PAT =                                                     0x00000277 //col:9563
IA32_PAT_PA0_BIT =                                             0 //col:9572
IA32_PAT_PA0_FLAG =                                            0x07 //col:9573
IA32_PAT_PA0_MASK =                                            0x07 //col:9574
IA32_PAT_PA0(_) =                                              (((_) >> 0) & 0x07) //col:9575
IA32_PAT_PA1_BIT =                                             8 //col:9582
IA32_PAT_PA1_FLAG =                                            0x700 //col:9583
IA32_PAT_PA1_MASK =                                            0x07 //col:9584
IA32_PAT_PA1(_) =                                              (((_) >> 8) & 0x07) //col:9585
IA32_PAT_PA2_BIT =                                             16 //col:9592
IA32_PAT_PA2_FLAG =                                            0x70000 //col:9593
IA32_PAT_PA2_MASK =                                            0x07 //col:9594
IA32_PAT_PA2(_) =                                              (((_) >> 16) & 0x07) //col:9595
IA32_PAT_PA3_BIT =                                             24 //col:9602
IA32_PAT_PA3_FLAG =                                            0x7000000 //col:9603
IA32_PAT_PA3_MASK =                                            0x07 //col:9604
IA32_PAT_PA3(_) =                                              (((_) >> 24) & 0x07) //col:9605
IA32_PAT_PA4_BIT =                                             32 //col:9612
IA32_PAT_PA4_FLAG =                                            0x700000000 //col:9613
IA32_PAT_PA4_MASK =                                            0x07 //col:9614
IA32_PAT_PA4(_) =                                              (((_) >> 32) & 0x07) //col:9615
IA32_PAT_PA5_BIT =                                             40 //col:9622
IA32_PAT_PA5_FLAG =                                            0x70000000000 //col:9623
IA32_PAT_PA5_MASK =                                            0x07 //col:9624
IA32_PAT_PA5(_) =                                              (((_) >> 40) & 0x07) //col:9625
IA32_PAT_PA6_BIT =                                             48 //col:9632
IA32_PAT_PA6_FLAG =                                            0x7000000000000 //col:9633
IA32_PAT_PA6_MASK =                                            0x07 //col:9634
IA32_PAT_PA6(_) =                                              (((_) >> 48) & 0x07) //col:9635
IA32_PAT_PA7_BIT =                                             56 //col:9642
IA32_PAT_PA7_FLAG =                                            0x700000000000000 //col:9643
IA32_PAT_PA7_MASK =                                            0x07 //col:9644
IA32_PAT_PA7(_) =                                              (((_) >> 56) & 0x07) //col:9645
IA32_MC0_CTL2 =                                                0x00000280 //col:9662
IA32_MC1_CTL2 =                                                0x00000281 //col:9663
IA32_MC2_CTL2 =                                                0x00000282 //col:9664
IA32_MC3_CTL2 =                                                0x00000283 //col:9665
IA32_MC4_CTL2 =                                                0x00000284 //col:9666
IA32_MC5_CTL2 =                                                0x00000285 //col:9667
IA32_MC6_CTL2 =                                                0x00000286 //col:9668
IA32_MC7_CTL2 =                                                0x00000287 //col:9669
IA32_MC8_CTL2 =                                                0x00000288 //col:9670
IA32_MC9_CTL2 =                                                0x00000289 //col:9671
IA32_MC10_CTL2 =                                               0x0000028A //col:9672
IA32_MC11_CTL2 =                                               0x0000028B //col:9673
IA32_MC12_CTL2 =                                               0x0000028C //col:9674
IA32_MC13_CTL2 =                                               0x0000028D //col:9675
IA32_MC14_CTL2 =                                               0x0000028E //col:9676
IA32_MC15_CTL2 =                                               0x0000028F //col:9677
IA32_MC16_CTL2 =                                               0x00000290 //col:9678
IA32_MC17_CTL2 =                                               0x00000291 //col:9679
IA32_MC18_CTL2 =                                               0x00000292 //col:9680
IA32_MC19_CTL2 =                                               0x00000293 //col:9681
IA32_MC20_CTL2 =                                               0x00000294 //col:9682
IA32_MC21_CTL2 =                                               0x00000295 //col:9683
IA32_MC22_CTL2 =                                               0x00000296 //col:9684
IA32_MC23_CTL2 =                                               0x00000297 //col:9685
IA32_MC24_CTL2 =                                               0x00000298 //col:9686
IA32_MC25_CTL2 =                                               0x00000299 //col:9687
IA32_MC26_CTL2 =                                               0x0000029A //col:9688
IA32_MC27_CTL2 =                                               0x0000029B //col:9689
IA32_MC28_CTL2 =                                               0x0000029C //col:9690
IA32_MC29_CTL2 =                                               0x0000029D //col:9691
IA32_MC30_CTL2 =                                               0x0000029E //col:9692
IA32_MC31_CTL2 =                                               0x0000029F //col:9693
IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_BIT =             0 //col:9702
IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_FLAG =            0x7FFF //col:9703
IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_MASK =            0x7FFF //col:9704
IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD(_) =              (((_) >> 0) & 0x7FFF) //col:9705
IA32_MC_CTL2_CMCI_EN_BIT =                                     30 //col:9712
IA32_MC_CTL2_CMCI_EN_FLAG =                                    0x40000000 //col:9713
IA32_MC_CTL2_CMCI_EN_MASK =                                    0x01 //col:9714
IA32_MC_CTL2_CMCI_EN(_) =                                      (((_) >> 30) & 0x01) //col:9715
IA32_MTRR_DEF_TYPE =                                           0x000002FF //col:9732
IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_BIT =                   0 //col:9741
IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_FLAG =                  0x07 //col:9742
IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_MASK =                  0x07 //col:9743
IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE(_) =                    (((_) >> 0) & 0x07) //col:9744
IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_BIT =               10 //col:9751
IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_FLAG =              0x400 //col:9752
IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_MASK =              0x01 //col:9753
IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE(_) =                (((_) >> 10) & 0x01) //col:9754
IA32_MTRR_DEF_TYPE_MTRR_ENABLE_BIT =                           11 //col:9760
IA32_MTRR_DEF_TYPE_MTRR_ENABLE_FLAG =                          0x800 //col:9761
IA32_MTRR_DEF_TYPE_MTRR_ENABLE_MASK =                          0x01 //col:9762
IA32_MTRR_DEF_TYPE_MTRR_ENABLE(_) =                            (((_) >> 11) & 0x01) //col:9763
IA32_FIXED_CTR0 =                                              0x00000309 //col:9782
IA32_FIXED_CTR1 =                                              0x0000030A //col:9787
IA32_FIXED_CTR2 =                                              0x0000030B //col:9792
IA32_PERF_CAPABILITIES =                                       0x00000345 //col:9803
IA32_PERF_CAPABILITIES_LBR_FORMAT_BIT =                        0 //col:9812
IA32_PERF_CAPABILITIES_LBR_FORMAT_FLAG =                       0x3F //col:9813
IA32_PERF_CAPABILITIES_LBR_FORMAT_MASK =                       0x3F //col:9814
IA32_PERF_CAPABILITIES_LBR_FORMAT(_) =                         (((_) >> 0) & 0x3F) //col:9815
IA32_PERF_CAPABILITIES_PEBS_TRAP_BIT =                         6 //col:9821
IA32_PERF_CAPABILITIES_PEBS_TRAP_FLAG =                        0x40 //col:9822
IA32_PERF_CAPABILITIES_PEBS_TRAP_MASK =                        0x01 //col:9823
IA32_PERF_CAPABILITIES_PEBS_TRAP(_) =                          (((_) >> 6) & 0x01) //col:9824
IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_BIT =               7 //col:9830
IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_FLAG =              0x80 //col:9831
IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_MASK =              0x01 //col:9832
IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS(_) =                (((_) >> 7) & 0x01) //col:9833
IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_BIT =                8 //col:9839
IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_FLAG =               0xF00 //col:9840
IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_MASK =               0x0F //col:9841
IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT(_) =                 (((_) >> 8) & 0x0F) //col:9842
IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_BIT =     12 //col:9848
IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_FLAG =    0x1000 //col:9849
IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_MASK =    0x01 //col:9850
IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED(_) =      (((_) >> 12) & 0x01) //col:9851
IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_BIT =          13 //col:9857
IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_FLAG =         0x2000 //col:9858
IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_MASK =         0x01 //col:9859
IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE(_) =           (((_) >> 13) & 0x01) //col:9860
IA32_FIXED_CTR_CTRL =                                          0x0000038D //col:9876
IA32_FIXED_CTR_CTRL_EN0_OS_BIT =                               0 //col:9885
IA32_FIXED_CTR_CTRL_EN0_OS_FLAG =                              0x01 //col:9886
IA32_FIXED_CTR_CTRL_EN0_OS_MASK =                              0x01 //col:9887
IA32_FIXED_CTR_CTRL_EN0_OS(_) =                                (((_) >> 0) & 0x01) //col:9888
IA32_FIXED_CTR_CTRL_EN0_USR_BIT =                              1 //col:9894
IA32_FIXED_CTR_CTRL_EN0_USR_FLAG =                             0x02 //col:9895
IA32_FIXED_CTR_CTRL_EN0_USR_MASK =                             0x01 //col:9896
IA32_FIXED_CTR_CTRL_EN0_USR(_) =                               (((_) >> 1) & 0x01) //col:9897
IA32_FIXED_CTR_CTRL_ANY_THREAD0_BIT =                          2 //col:9905
IA32_FIXED_CTR_CTRL_ANY_THREAD0_FLAG =                         0x04 //col:9906
IA32_FIXED_CTR_CTRL_ANY_THREAD0_MASK =                         0x01 //col:9907
IA32_FIXED_CTR_CTRL_ANY_THREAD0(_) =                           (((_) >> 2) & 0x01) //col:9908
IA32_FIXED_CTR_CTRL_EN0_PMI_BIT =                              3 //col:9914
IA32_FIXED_CTR_CTRL_EN0_PMI_FLAG =                             0x08 //col:9915
IA32_FIXED_CTR_CTRL_EN0_PMI_MASK =                             0x01 //col:9916
IA32_FIXED_CTR_CTRL_EN0_PMI(_) =                               (((_) >> 3) & 0x01) //col:9917
IA32_FIXED_CTR_CTRL_EN1_OS_BIT =                               4 //col:9923
IA32_FIXED_CTR_CTRL_EN1_OS_FLAG =                              0x10 //col:9924
IA32_FIXED_CTR_CTRL_EN1_OS_MASK =                              0x01 //col:9925
IA32_FIXED_CTR_CTRL_EN1_OS(_) =                                (((_) >> 4) & 0x01) //col:9926
IA32_FIXED_CTR_CTRL_EN1_USR_BIT =                              5 //col:9932
IA32_FIXED_CTR_CTRL_EN1_USR_FLAG =                             0x20 //col:9933
IA32_FIXED_CTR_CTRL_EN1_USR_MASK =                             0x01 //col:9934
IA32_FIXED_CTR_CTRL_EN1_USR(_) =                               (((_) >> 5) & 0x01) //col:9935
IA32_FIXED_CTR_CTRL_ANY_THREAD1_BIT =                          6 //col:9945
IA32_FIXED_CTR_CTRL_ANY_THREAD1_FLAG =                         0x40 //col:9946
IA32_FIXED_CTR_CTRL_ANY_THREAD1_MASK =                         0x01 //col:9947
IA32_FIXED_CTR_CTRL_ANY_THREAD1(_) =                           (((_) >> 6) & 0x01) //col:9948
IA32_FIXED_CTR_CTRL_EN1_PMI_BIT =                              7 //col:9954
IA32_FIXED_CTR_CTRL_EN1_PMI_FLAG =                             0x80 //col:9955
IA32_FIXED_CTR_CTRL_EN1_PMI_MASK =                             0x01 //col:9956
IA32_FIXED_CTR_CTRL_EN1_PMI(_) =                               (((_) >> 7) & 0x01) //col:9957
IA32_FIXED_CTR_CTRL_EN2_OS_BIT =                               8 //col:9963
IA32_FIXED_CTR_CTRL_EN2_OS_FLAG =                              0x100 //col:9964
IA32_FIXED_CTR_CTRL_EN2_OS_MASK =                              0x01 //col:9965
IA32_FIXED_CTR_CTRL_EN2_OS(_) =                                (((_) >> 8) & 0x01) //col:9966
IA32_FIXED_CTR_CTRL_EN2_USR_BIT =                              9 //col:9972
IA32_FIXED_CTR_CTRL_EN2_USR_FLAG =                             0x200 //col:9973
IA32_FIXED_CTR_CTRL_EN2_USR_MASK =                             0x01 //col:9974
IA32_FIXED_CTR_CTRL_EN2_USR(_) =                               (((_) >> 9) & 0x01) //col:9975
IA32_FIXED_CTR_CTRL_ANY_THREAD2_BIT =                          10 //col:9985
IA32_FIXED_CTR_CTRL_ANY_THREAD2_FLAG =                         0x400 //col:9986
IA32_FIXED_CTR_CTRL_ANY_THREAD2_MASK =                         0x01 //col:9987
IA32_FIXED_CTR_CTRL_ANY_THREAD2(_) =                           (((_) >> 10) & 0x01) //col:9988
IA32_FIXED_CTR_CTRL_EN2_PMI_BIT =                              11 //col:9994
IA32_FIXED_CTR_CTRL_EN2_PMI_FLAG =                             0x800 //col:9995
IA32_FIXED_CTR_CTRL_EN2_PMI_MASK =                             0x01 //col:9996
IA32_FIXED_CTR_CTRL_EN2_PMI(_) =                               (((_) >> 11) & 0x01) //col:9997
IA32_PERF_GLOBAL_STATUS =                                      0x0000038E //col:10010
IA32_PERF_GLOBAL_STATUS_OVF_PMC0_BIT =                         0 //col:10021
IA32_PERF_GLOBAL_STATUS_OVF_PMC0_FLAG =                        0x01 //col:10022
IA32_PERF_GLOBAL_STATUS_OVF_PMC0_MASK =                        0x01 //col:10023
IA32_PERF_GLOBAL_STATUS_OVF_PMC0(_) =                          (((_) >> 0) & 0x01) //col:10024
IA32_PERF_GLOBAL_STATUS_OVF_PMC1_BIT =                         1 //col:10032
IA32_PERF_GLOBAL_STATUS_OVF_PMC1_FLAG =                        0x02 //col:10033
IA32_PERF_GLOBAL_STATUS_OVF_PMC1_MASK =                        0x01 //col:10034
IA32_PERF_GLOBAL_STATUS_OVF_PMC1(_) =                          (((_) >> 1) & 0x01) //col:10035
IA32_PERF_GLOBAL_STATUS_OVF_PMC2_BIT =                         2 //col:10043
IA32_PERF_GLOBAL_STATUS_OVF_PMC2_FLAG =                        0x04 //col:10044
IA32_PERF_GLOBAL_STATUS_OVF_PMC2_MASK =                        0x01 //col:10045
IA32_PERF_GLOBAL_STATUS_OVF_PMC2(_) =                          (((_) >> 2) & 0x01) //col:10046
IA32_PERF_GLOBAL_STATUS_OVF_PMC3_BIT =                         3 //col:10054
IA32_PERF_GLOBAL_STATUS_OVF_PMC3_FLAG =                        0x08 //col:10055
IA32_PERF_GLOBAL_STATUS_OVF_PMC3_MASK =                        0x01 //col:10056
IA32_PERF_GLOBAL_STATUS_OVF_PMC3(_) =                          (((_) >> 3) & 0x01) //col:10057
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_BIT =                    32 //col:10066
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_FLAG =                   0x100000000 //col:10067
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_MASK =                   0x01 //col:10068
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0(_) =                     (((_) >> 32) & 0x01) //col:10069
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_BIT =                    33 //col:10077
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_FLAG =                   0x200000000 //col:10078
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_MASK =                   0x01 //col:10079
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1(_) =                     (((_) >> 33) & 0x01) //col:10080
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_BIT =                    34 //col:10088
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_FLAG =                   0x400000000 //col:10089
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_MASK =                   0x01 //col:10090
IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2(_) =                     (((_) >> 34) & 0x01) //col:10091
IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_BIT =                   55 //col:10100
IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_FLAG =                  0x80000000000000 //col:10101
IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_MASK =                  0x01 //col:10102
IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI(_) =                    (((_) >> 55) & 0x01) //col:10103
IA32_PERF_GLOBAL_STATUS_LBR_FRZ_BIT =                          58 //col:10114
IA32_PERF_GLOBAL_STATUS_LBR_FRZ_FLAG =                         0x400000000000000 //col:10115
IA32_PERF_GLOBAL_STATUS_LBR_FRZ_MASK =                         0x01 //col:10116
IA32_PERF_GLOBAL_STATUS_LBR_FRZ(_) =                           (((_) >> 58) & 0x01) //col:10117
IA32_PERF_GLOBAL_STATUS_CTR_FRZ_BIT =                          59 //col:10127
IA32_PERF_GLOBAL_STATUS_CTR_FRZ_FLAG =                         0x800000000000000 //col:10128
IA32_PERF_GLOBAL_STATUS_CTR_FRZ_MASK =                         0x01 //col:10129
IA32_PERF_GLOBAL_STATUS_CTR_FRZ(_) =                           (((_) >> 59) & 0x01) //col:10130
IA32_PERF_GLOBAL_STATUS_ASCI_BIT =                             60 //col:10139
IA32_PERF_GLOBAL_STATUS_ASCI_FLAG =                            0x1000000000000000 //col:10140
IA32_PERF_GLOBAL_STATUS_ASCI_MASK =                            0x01 //col:10141
IA32_PERF_GLOBAL_STATUS_ASCI(_) =                              (((_) >> 60) & 0x01) //col:10142
IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_BIT =                       61 //col:10150
IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_FLAG =                      0x2000000000000000 //col:10151
IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_MASK =                      0x01 //col:10152
IA32_PERF_GLOBAL_STATUS_OVF_UNCORE(_) =                        (((_) >> 61) & 0x01) //col:10153
IA32_PERF_GLOBAL_STATUS_OVF_BUF_BIT =                          62 //col:10161
IA32_PERF_GLOBAL_STATUS_OVF_BUF_FLAG =                         0x4000000000000000 //col:10162
IA32_PERF_GLOBAL_STATUS_OVF_BUF_MASK =                         0x01 //col:10163
IA32_PERF_GLOBAL_STATUS_OVF_BUF(_) =                           (((_) >> 62) & 0x01) //col:10164
IA32_PERF_GLOBAL_STATUS_COND_CHGD_BIT =                        63 //col:10172
IA32_PERF_GLOBAL_STATUS_COND_CHGD_FLAG =                       0x8000000000000000 //col:10173
IA32_PERF_GLOBAL_STATUS_COND_CHGD_MASK =                       0x01 //col:10174
IA32_PERF_GLOBAL_STATUS_COND_CHGD(_) =                         (((_) >> 63) & 0x01) //col:10175
IA32_PERF_GLOBAL_CTRL =                                        0x0000038F //col:10190
IA32_PERF_GLOBAL_CTRL_EN_PMCN_BIT =                            0 //col:10201
IA32_PERF_GLOBAL_CTRL_EN_PMCN_FLAG =                           0xFFFFFFFF //col:10202
IA32_PERF_GLOBAL_CTRL_EN_PMCN_MASK =                           0xFFFFFFFF //col:10203
IA32_PERF_GLOBAL_CTRL_EN_PMCN(_) =                             (((_) >> 0) & 0xFFFFFFFF) //col:10204
IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_BIT =                      32 //col:10212
IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_FLAG =                     0xFFFFFFFF00000000 //col:10213
IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_MASK =                     0xFFFFFFFF //col:10214
IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN(_) =                       (((_) >> 32) & 0xFFFFFFFF) //col:10215
IA32_PERF_GLOBAL_STATUS_RESET =                                0x00000390 //col:10227
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_BIT =             0 //col:10238
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_FLAG =            0xFFFFFFFF //col:10239
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_MASK =            0xFFFFFFFF //col:10240
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN(_) =              (((_) >> 0) & 0xFFFFFFFF) //col:10241
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_BIT =       32 //col:10250
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_FLAG =      0x700000000 //col:10251
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_MASK =      0x07 //col:10252
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN(_) =        (((_) >> 32) & 0x07) //col:10253
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_BIT =       55 //col:10262
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_FLAG =      0x80000000000000 //col:10263
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_MASK =      0x01 //col:10264
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI(_) =        (((_) >> 55) & 0x01) //col:10265
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_BIT =              58 //col:10274
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_FLAG =             0x400000000000000 //col:10275
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_MASK =             0x01 //col:10276
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ(_) =               (((_) >> 58) & 0x01) //col:10277
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_BIT =              59 //col:10285
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_FLAG =             0x800000000000000 //col:10286
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_MASK =             0x01 //col:10287
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ(_) =               (((_) >> 59) & 0x01) //col:10288
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_BIT =                 60 //col:10296
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_FLAG =                0x1000000000000000 //col:10297
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_MASK =                0x01 //col:10298
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI(_) =                  (((_) >> 60) & 0x01) //col:10299
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_BIT =           61 //col:10307
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_FLAG =          0x2000000000000000 //col:10308
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_MASK =          0x01 //col:10309
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE(_) =            (((_) >> 61) & 0x01) //col:10310
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_BIT =              62 //col:10318
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_FLAG =             0x4000000000000000 //col:10319
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_MASK =             0x01 //col:10320
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF(_) =               (((_) >> 62) & 0x01) //col:10321
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_BIT =            63 //col:10329
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_FLAG =           0x8000000000000000 //col:10330
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_MASK =           0x01 //col:10331
IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD(_) =             (((_) >> 63) & 0x01) //col:10332
IA32_PERF_GLOBAL_STATUS_SET =                                  0x00000391 //col:10344
IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_BIT =                     0 //col:10355
IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_FLAG =                    0xFFFFFFFF //col:10356
IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_MASK =                    0xFFFFFFFF //col:10357
IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN(_) =                      (((_) >> 0) & 0xFFFFFFFF) //col:10358
IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_BIT =               32 //col:10367
IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_FLAG =              0x700000000 //col:10368
IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_MASK =              0x07 //col:10369
IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN(_) =                (((_) >> 32) & 0x07) //col:10370
IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_BIT =               55 //col:10379
IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_FLAG =              0x80000000000000 //col:10380
IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_MASK =              0x01 //col:10381
IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI(_) =                (((_) >> 55) & 0x01) //col:10382
IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_BIT =                      58 //col:10391
IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_FLAG =                     0x400000000000000 //col:10392
IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_MASK =                     0x01 //col:10393
IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ(_) =                       (((_) >> 58) & 0x01) //col:10394
IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_BIT =                      59 //col:10402
IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_FLAG =                     0x800000000000000 //col:10403
IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_MASK =                     0x01 //col:10404
IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ(_) =                       (((_) >> 59) & 0x01) //col:10405
IA32_PERF_GLOBAL_STATUS_SET_ASCI_BIT =                         60 //col:10413
IA32_PERF_GLOBAL_STATUS_SET_ASCI_FLAG =                        0x1000000000000000 //col:10414
IA32_PERF_GLOBAL_STATUS_SET_ASCI_MASK =                        0x01 //col:10415
IA32_PERF_GLOBAL_STATUS_SET_ASCI(_) =                          (((_) >> 60) & 0x01) //col:10416
IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_BIT =                   61 //col:10424
IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_FLAG =                  0x2000000000000000 //col:10425
IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_MASK =                  0x01 //col:10426
IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE(_) =                    (((_) >> 61) & 0x01) //col:10427
IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_BIT =                      62 //col:10435
IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_FLAG =                     0x4000000000000000 //col:10436
IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_MASK =                     0x01 //col:10437
IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF(_) =                       (((_) >> 62) & 0x01) //col:10438
IA32_PERF_GLOBAL_INUSE =                                       0x00000392 //col:10451
IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_BIT =           0 //col:10462
IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_FLAG =          0xFFFFFFFF //col:10463
IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_MASK =          0xFFFFFFFF //col:10464
IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE(_) =            (((_) >> 0) & 0xFFFFFFFF) //col:10465
IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_BIT =            32 //col:10471
IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_FLAG =           0x700000000 //col:10472
IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_MASK =           0x07 //col:10473
IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE(_) =             (((_) >> 32) & 0x07) //col:10474
IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_BIT =                        63 //col:10481
IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_FLAG =                       0x8000000000000000 //col:10482
IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_MASK =                       0x01 //col:10483
IA32_PERF_GLOBAL_INUSE_PMI_IN_USE(_) =                         (((_) >> 63) & 0x01) //col:10484
IA32_PEBS_ENABLE =                                             0x000003F1 //col:10496
IA32_PEBS_ENABLE_ENABLE_PEBS_BIT =                             0 //col:10507
IA32_PEBS_ENABLE_ENABLE_PEBS_FLAG =                            0x01 //col:10508
IA32_PEBS_ENABLE_ENABLE_PEBS_MASK =                            0x01 //col:10509
IA32_PEBS_ENABLE_ENABLE_PEBS(_) =                              (((_) >> 0) & 0x01) //col:10510
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_BIT =                1 //col:10516
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_FLAG =               0x0E //col:10517
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_MASK =               0x07 //col:10518
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1(_) =                 (((_) >> 1) & 0x07) //col:10519
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_BIT =                32 //col:10526
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_FLAG =               0xF00000000 //col:10527
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_MASK =               0x0F //col:10528
IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2(_) =                 (((_) >> 32) & 0x0F) //col:10529
IA32_MC0_CTL =                                                 0x00000400 //col:10545
IA32_MC1_CTL =                                                 0x00000404 //col:10546
IA32_MC2_CTL =                                                 0x00000408 //col:10547
IA32_MC3_CTL =                                                 0x0000040C //col:10548
IA32_MC4_CTL =                                                 0x00000410 //col:10549
IA32_MC5_CTL =                                                 0x00000414 //col:10550
IA32_MC6_CTL =                                                 0x00000418 //col:10551
IA32_MC7_CTL =                                                 0x0000041C //col:10552
IA32_MC8_CTL =                                                 0x00000420 //col:10553
IA32_MC9_CTL =                                                 0x00000424 //col:10554
IA32_MC10_CTL =                                                0x00000428 //col:10555
IA32_MC11_CTL =                                                0x0000042C //col:10556
IA32_MC12_CTL =                                                0x00000430 //col:10557
IA32_MC13_CTL =                                                0x00000434 //col:10558
IA32_MC14_CTL =                                                0x00000438 //col:10559
IA32_MC15_CTL =                                                0x0000043C //col:10560
IA32_MC16_CTL =                                                0x00000440 //col:10561
IA32_MC17_CTL =                                                0x00000444 //col:10562
IA32_MC18_CTL =                                                0x00000448 //col:10563
IA32_MC19_CTL =                                                0x0000044C //col:10564
IA32_MC20_CTL =                                                0x00000450 //col:10565
IA32_MC21_CTL =                                                0x00000454 //col:10566
IA32_MC22_CTL =                                                0x00000458 //col:10567
IA32_MC23_CTL =                                                0x0000045C //col:10568
IA32_MC24_CTL =                                                0x00000460 //col:10569
IA32_MC25_CTL =                                                0x00000464 //col:10570
IA32_MC26_CTL =                                                0x00000468 //col:10571
IA32_MC27_CTL =                                                0x0000046C //col:10572
IA32_MC28_CTL =                                                0x00000470 //col:10573
IA32_MC0_STATUS =                                              0x00000401 //col:10587
IA32_MC1_STATUS =                                              0x00000405 //col:10588
IA32_MC2_STATUS =                                              0x00000409 //col:10589
IA32_MC3_STATUS =                                              0x0000040D //col:10590
IA32_MC4_STATUS =                                              0x00000411 //col:10591
IA32_MC5_STATUS =                                              0x00000415 //col:10592
IA32_MC6_STATUS =                                              0x00000419 //col:10593
IA32_MC7_STATUS =                                              0x0000041D //col:10594
IA32_MC8_STATUS =                                              0x00000421 //col:10595
IA32_MC9_STATUS =                                              0x00000425 //col:10596
IA32_MC10_STATUS =                                             0x00000429 //col:10597
IA32_MC11_STATUS =                                             0x0000042D //col:10598
IA32_MC12_STATUS =                                             0x00000431 //col:10599
IA32_MC13_STATUS =                                             0x00000435 //col:10600
IA32_MC14_STATUS =                                             0x00000439 //col:10601
IA32_MC15_STATUS =                                             0x0000043D //col:10602
IA32_MC16_STATUS =                                             0x00000441 //col:10603
IA32_MC17_STATUS =                                             0x00000445 //col:10604
IA32_MC18_STATUS =                                             0x00000449 //col:10605
IA32_MC19_STATUS =                                             0x0000044D //col:10606
IA32_MC20_STATUS =                                             0x00000451 //col:10607
IA32_MC21_STATUS =                                             0x00000455 //col:10608
IA32_MC22_STATUS =                                             0x00000459 //col:10609
IA32_MC23_STATUS =                                             0x0000045D //col:10610
IA32_MC24_STATUS =                                             0x00000461 //col:10611
IA32_MC25_STATUS =                                             0x00000465 //col:10612
IA32_MC26_STATUS =                                             0x00000469 //col:10613
IA32_MC27_STATUS =                                             0x0000046D //col:10614
IA32_MC28_STATUS =                                             0x00000471 //col:10615
IA32_MC0_ADDR =                                                0x00000402 //col:10629
IA32_MC1_ADDR =                                                0x00000406 //col:10630
IA32_MC2_ADDR =                                                0x0000040A //col:10631
IA32_MC3_ADDR =                                                0x0000040E //col:10632
IA32_MC4_ADDR =                                                0x00000412 //col:10633
IA32_MC5_ADDR =                                                0x00000416 //col:10634
IA32_MC6_ADDR =                                                0x0000041A //col:10635
IA32_MC7_ADDR =                                                0x0000041E //col:10636
IA32_MC8_ADDR =                                                0x00000422 //col:10637
IA32_MC9_ADDR =                                                0x00000426 //col:10638
IA32_MC10_ADDR =                                               0x0000042A //col:10639
IA32_MC11_ADDR =                                               0x0000042E //col:10640
IA32_MC12_ADDR =                                               0x00000432 //col:10641
IA32_MC13_ADDR =                                               0x00000436 //col:10642
IA32_MC14_ADDR =                                               0x0000043A //col:10643
IA32_MC15_ADDR =                                               0x0000043E //col:10644
IA32_MC16_ADDR =                                               0x00000442 //col:10645
IA32_MC17_ADDR =                                               0x00000446 //col:10646
IA32_MC18_ADDR =                                               0x0000044A //col:10647
IA32_MC19_ADDR =                                               0x0000044E //col:10648
IA32_MC20_ADDR =                                               0x00000452 //col:10649
IA32_MC21_ADDR =                                               0x00000456 //col:10650
IA32_MC22_ADDR =                                               0x0000045A //col:10651
IA32_MC23_ADDR =                                               0x0000045E //col:10652
IA32_MC24_ADDR =                                               0x00000462 //col:10653
IA32_MC25_ADDR =                                               0x00000466 //col:10654
IA32_MC26_ADDR =                                               0x0000046A //col:10655
IA32_MC27_ADDR =                                               0x0000046E //col:10656
IA32_MC28_ADDR =                                               0x00000472 //col:10657
IA32_MC0_MISC =                                                0x00000403 //col:10671
IA32_MC1_MISC =                                                0x00000407 //col:10672
IA32_MC2_MISC =                                                0x0000040B //col:10673
IA32_MC3_MISC =                                                0x0000040F //col:10674
IA32_MC4_MISC =                                                0x00000413 //col:10675
IA32_MC5_MISC =                                                0x00000417 //col:10676
IA32_MC6_MISC =                                                0x0000041B //col:10677
IA32_MC7_MISC =                                                0x0000041F //col:10678
IA32_MC8_MISC =                                                0x00000423 //col:10679
IA32_MC9_MISC =                                                0x00000427 //col:10680
IA32_MC10_MISC =                                               0x0000042B //col:10681
IA32_MC11_MISC =                                               0x0000042F //col:10682
IA32_MC12_MISC =                                               0x00000433 //col:10683
IA32_MC13_MISC =                                               0x00000437 //col:10684
IA32_MC14_MISC =                                               0x0000043B //col:10685
IA32_MC15_MISC =                                               0x0000043F //col:10686
IA32_MC16_MISC =                                               0x00000443 //col:10687
IA32_MC17_MISC =                                               0x00000447 //col:10688
IA32_MC18_MISC =                                               0x0000044B //col:10689
IA32_MC19_MISC =                                               0x0000044F //col:10690
IA32_MC20_MISC =                                               0x00000453 //col:10691
IA32_MC21_MISC =                                               0x00000457 //col:10692
IA32_MC22_MISC =                                               0x0000045B //col:10693
IA32_MC23_MISC =                                               0x0000045F //col:10694
IA32_MC24_MISC =                                               0x00000463 //col:10695
IA32_MC25_MISC =                                               0x00000467 //col:10696
IA32_MC26_MISC =                                               0x0000046B //col:10697
IA32_MC27_MISC =                                               0x0000046F //col:10698
IA32_MC28_MISC =                                               0x00000473 //col:10699
IA32_VMX_BASIC =                                               0x00000480 //col:10712
IA32_VMX_BASIC_VMCS_REVISION_ID_BIT =                          0 //col:10724
IA32_VMX_BASIC_VMCS_REVISION_ID_FLAG =                         0x7FFFFFFF //col:10725
IA32_VMX_BASIC_VMCS_REVISION_ID_MASK =                         0x7FFFFFFF //col:10726
IA32_VMX_BASIC_VMCS_REVISION_ID(_) =                           (((_) >> 0) & 0x7FFFFFFF) //col:10727
IA32_VMX_BASIC_MUST_BE_ZERO_BIT =                              31 //col:10733
IA32_VMX_BASIC_MUST_BE_ZERO_FLAG =                             0x80000000 //col:10734
IA32_VMX_BASIC_MUST_BE_ZERO_MASK =                             0x01 //col:10735
IA32_VMX_BASIC_MUST_BE_ZERO(_) =                               (((_) >> 31) & 0x01) //col:10736
IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_BIT =                        32 //col:10745
IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_FLAG =                       0x1FFF00000000 //col:10746
IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_MASK =                       0x1FFF //col:10747
IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES(_) =                         (((_) >> 32) & 0x1FFF) //col:10748
IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_BIT =               48 //col:10762
IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_FLAG =              0x1000000000000 //col:10763
IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_MASK =              0x01 //col:10764
IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH(_) =                (((_) >> 48) & 0x01) //col:10765
IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_BIT =                      49 //col:10777
IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_FLAG =                     0x2000000000000 //col:10778
IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_MASK =                     0x01 //col:10779
IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT(_) =                       (((_) >> 49) & 0x01) //col:10780
IA32_VMX_BASIC_MEMORY_TYPE_BIT =                               50 //col:10793
IA32_VMX_BASIC_MEMORY_TYPE_FLAG =                              0x3C000000000000 //col:10794
IA32_VMX_BASIC_MEMORY_TYPE_MASK =                              0x0F //col:10795
IA32_VMX_BASIC_MEMORY_TYPE(_) =                                (((_) >> 50) & 0x0F) //col:10796
IA32_VMX_BASIC_INS_OUTS_REPORTING_BIT =                        54 //col:10807
IA32_VMX_BASIC_INS_OUTS_REPORTING_FLAG =                       0x40000000000000 //col:10808
IA32_VMX_BASIC_INS_OUTS_REPORTING_MASK =                       0x01 //col:10809
IA32_VMX_BASIC_INS_OUTS_REPORTING(_) =                         (((_) >> 54) & 0x01) //col:10810
IA32_VMX_BASIC_VMX_CONTROLS_BIT =                              55 //col:10827
IA32_VMX_BASIC_VMX_CONTROLS_FLAG =                             0x80000000000000 //col:10828
IA32_VMX_BASIC_VMX_CONTROLS_MASK =                             0x01 //col:10829
IA32_VMX_BASIC_VMX_CONTROLS(_) =                               (((_) >> 55) & 0x01) //col:10830
IA32_VMX_PINBASED_CTLS =                                       0x00000481 //col:10845
IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_BIT =        0 //col:10857
IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_FLAG =       0x01 //col:10858
IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_MASK =       0x01 //col:10859
IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING(_) =         (((_) >> 0) & 0x01) //col:10860
IA32_VMX_PINBASED_CTLS_NMI_EXITING_BIT =                       3 //col:10872
IA32_VMX_PINBASED_CTLS_NMI_EXITING_FLAG =                      0x08 //col:10873
IA32_VMX_PINBASED_CTLS_NMI_EXITING_MASK =                      0x01 //col:10874
IA32_VMX_PINBASED_CTLS_NMI_EXITING(_) =                        (((_) >> 3) & 0x01) //col:10875
IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_BIT =                       5 //col:10887
IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_FLAG =                      0x20 //col:10888
IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_MASK =                      0x01 //col:10889
IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI(_) =                        (((_) >> 5) & 0x01) //col:10890
IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_BIT =     6 //col:10902
IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_FLAG =    0x40 //col:10903
IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_MASK =    0x01 //col:10904
IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER(_) =      (((_) >> 6) & 0x01) //col:10905
IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_BIT =         7 //col:10917
IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_FLAG =        0x80 //col:10918
IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_MASK =        0x01 //col:10919
IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS(_) =          (((_) >> 7) & 0x01) //col:10920
IA32_VMX_PROCBASED_CTLS =                                      0x00000482 //col:10935
IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_BIT =         2 //col:10951
IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_FLAG =        0x04 //col:10952
IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_MASK =        0x01 //col:10953
IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING(_) =          (((_) >> 2) & 0x01) //col:10954
IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_BIT =               3 //col:10966
IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_FLAG =              0x08 //col:10967
IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_MASK =              0x01 //col:10968
IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING(_) =                (((_) >> 3) & 0x01) //col:10969
IA32_VMX_PROCBASED_CTLS_HLT_EXITING_BIT =                      7 //col:10978
IA32_VMX_PROCBASED_CTLS_HLT_EXITING_FLAG =                     0x80 //col:10979
IA32_VMX_PROCBASED_CTLS_HLT_EXITING_MASK =                     0x01 //col:10980
IA32_VMX_PROCBASED_CTLS_HLT_EXITING(_) =                       (((_) >> 7) & 0x01) //col:10981
IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_BIT =                   9 //col:10990
IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_FLAG =                  0x200 //col:10991
IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_MASK =                  0x01 //col:10992
IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING(_) =                    (((_) >> 9) & 0x01) //col:10993
IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_BIT =                    10 //col:11001
IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_FLAG =                   0x400 //col:11002
IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_MASK =                   0x01 //col:11003
IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING(_) =                     (((_) >> 10) & 0x01) //col:11004
IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_BIT =                    11 //col:11012
IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_FLAG =                   0x800 //col:11013
IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_MASK =                   0x01 //col:11014
IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING(_) =                     (((_) >> 11) & 0x01) //col:11015
IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_BIT =                    12 //col:11023
IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_FLAG =                   0x1000 //col:11024
IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_MASK =                   0x01 //col:11025
IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING(_) =                     (((_) >> 12) & 0x01) //col:11026
IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_BIT =                 15 //col:11040
IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_FLAG =                0x8000 //col:11041
IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_MASK =                0x01 //col:11042
IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING(_) =                  (((_) >> 15) & 0x01) //col:11043
IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_BIT =                16 //col:11053
IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_FLAG =               0x10000 //col:11054
IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_MASK =               0x01 //col:11055
IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING(_) =                 (((_) >> 16) & 0x01) //col:11056
IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_BIT =       17 //col:11065
IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_FLAG =      0x20000 //col:11066
IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_MASK =      0x01 //col:11067
IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS(_) =        (((_) >> 17) & 0x01) //col:11068
IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_BIT =                 19 //col:11077
IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_FLAG =                0x80000 //col:11078
IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_MASK =                0x01 //col:11079
IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING(_) =                  (((_) >> 19) & 0x01) //col:11080
IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_BIT =                20 //col:11088
IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_FLAG =               0x100000 //col:11089
IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_MASK =               0x01 //col:11090
IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING(_) =                 (((_) >> 20) & 0x01) //col:11091
IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_BIT =                   21 //col:11101
IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_FLAG =                  0x200000 //col:11102
IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_MASK =                  0x01 //col:11103
IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW(_) =                    (((_) >> 21) & 0x01) //col:11104
IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_BIT =               22 //col:11114
IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_FLAG =              0x400000 //col:11115
IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_MASK =              0x01 //col:11116
IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING(_) =                (((_) >> 22) & 0x01) //col:11117
IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_BIT =                   23 //col:11125
IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_FLAG =                  0x800000 //col:11126
IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_MASK =                  0x01 //col:11127
IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING(_) =                    (((_) >> 23) & 0x01) //col:11128
IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_BIT =         24 //col:11137
IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_FLAG =        0x1000000 //col:11138
IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_MASK =        0x01 //col:11139
IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING(_) =          (((_) >> 24) & 0x01) //col:11140
IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_BIT =                   25 //col:11153
IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_FLAG =                  0x2000000 //col:11154
IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_MASK =                  0x01 //col:11155
IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS(_) =                    (((_) >> 25) & 0x01) //col:11156
IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_BIT =                27 //col:11167
IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_FLAG =               0x8000000 //col:11168
IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_MASK =               0x01 //col:11169
IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG(_) =                 (((_) >> 27) & 0x01) //col:11170
IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_BIT =                  28 //col:11183
IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_FLAG =                 0x10000000 //col:11184
IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_MASK =                 0x01 //col:11185
IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS(_) =                   (((_) >> 28) & 0x01) //col:11186
IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_BIT =                  29 //col:11194
IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_FLAG =                 0x20000000 //col:11195
IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_MASK =                 0x01 //col:11196
IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING(_) =                   (((_) >> 29) & 0x01) //col:11197
IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_BIT =                    30 //col:11205
IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_FLAG =                   0x40000000 //col:11206
IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_MASK =                   0x01 //col:11207
IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING(_) =                     (((_) >> 30) & 0x01) //col:11208
IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_BIT =      31 //col:11217
IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_FLAG =     0x80000000 //col:11218
IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_MASK =     0x01 //col:11219
IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS(_) =       (((_) >> 31) & 0x01) //col:11220
IA32_VMX_EXIT_CTLS =                                           0x00000483 //col:11235
IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_BIT =                   2 //col:11250
IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_FLAG =                  0x04 //col:11251
IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_MASK =                  0x01 //col:11252
IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS(_) =                    (((_) >> 2) & 0x01) //col:11253
IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_BIT =               9 //col:11264
IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_FLAG =              0x200 //col:11265
IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_MASK =              0x01 //col:11266
IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE(_) =                (((_) >> 9) & 0x01) //col:11267
IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_BIT =            12 //col:11276
IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_FLAG =           0x1000 //col:11277
IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_MASK =           0x01 //col:11278
IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL(_) =             (((_) >> 12) & 0x01) //col:11279
IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_BIT =         15 //col:11292
IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_FLAG =        0x8000 //col:11293
IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_MASK =        0x01 //col:11294
IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT(_) =          (((_) >> 15) & 0x01) //col:11295
IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_BIT =                         18 //col:11304
IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_FLAG =                        0x40000 //col:11305
IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_MASK =                        0x01 //col:11306
IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT(_) =                          (((_) >> 18) & 0x01) //col:11307
IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_BIT =                         19 //col:11315
IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_FLAG =                        0x80000 //col:11316
IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_MASK =                        0x01 //col:11317
IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT(_) =                          (((_) >> 19) & 0x01) //col:11318
IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_BIT =                        20 //col:11326
IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_FLAG =                       0x100000 //col:11327
IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_MASK =                       0x01 //col:11328
IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER(_) =                         (((_) >> 20) & 0x01) //col:11329
IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_BIT =                        21 //col:11337
IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_FLAG =                       0x200000 //col:11338
IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_MASK =                       0x01 //col:11339
IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER(_) =                         (((_) >> 21) & 0x01) //col:11340
IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_BIT =       22 //col:11348
IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_FLAG =      0x400000 //col:11349
IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_MASK =      0x01 //col:11350
IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE(_) =        (((_) >> 22) & 0x01) //col:11351
IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_BIT =                    23 //col:11357
IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_FLAG =                   0x800000 //col:11358
IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_MASK =                   0x01 //col:11359
IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS(_) =                     (((_) >> 23) & 0x01) //col:11360
IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_BIT =                   24 //col:11369
IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_FLAG =                  0x1000000 //col:11370
IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_MASK =                  0x01 //col:11371
IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT(_) =                    (((_) >> 24) & 0x01) //col:11372
IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_BIT =                   25 //col:11380
IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_FLAG =                  0x2000000 //col:11381
IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_MASK =                  0x01 //col:11382
IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL(_) =                    (((_) >> 25) & 0x01) //col:11383
IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_BIT =                    26 //col:11389
IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_FLAG =                   0x4000000 //col:11390
IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_MASK =                   0x01 //col:11391
IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL(_) =                     (((_) >> 26) & 0x01) //col:11392
IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_BIT =                   28 //col:11401
IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_FLAG =                  0x10000000 //col:11402
IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_MASK =                  0x01 //col:11403
IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE(_) =                    (((_) >> 28) & 0x01) //col:11404
IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_BIT =                        29 //col:11410
IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_FLAG =                       0x20000000 //col:11411
IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_MASK =                       0x01 //col:11412
IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS(_) =                         (((_) >> 29) & 0x01) //col:11413
IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_BIT =           31 //col:11421
IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_FLAG =          0x80000000 //col:11422
IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_MASK =          0x01 //col:11423
IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS(_) =            (((_) >> 31) & 0x01) //col:11424
IA32_VMX_ENTRY_CTLS =                                          0x00000484 //col:11439
IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_BIT =                  2 //col:11454
IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_FLAG =                 0x04 //col:11455
IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_MASK =                 0x01 //col:11456
IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS(_) =                   (((_) >> 2) & 0x01) //col:11457
IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_BIT =                     9 //col:11468
IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_FLAG =                    0x200 //col:11469
IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_MASK =                    0x01 //col:11470
IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST(_) =                      (((_) >> 9) & 0x01) //col:11471
IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_BIT =                         10 //col:11480
IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_FLAG =                        0x400 //col:11481
IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_MASK =                        0x01 //col:11482
IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM(_) =                          (((_) >> 10) & 0x01) //col:11483
IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_BIT =    11 //col:11494
IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_FLAG =   0x800 //col:11495
IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_MASK =   0x01 //col:11496
IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT(_) =     (((_) >> 11) & 0x01) //col:11497
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_BIT =           13 //col:11506
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_FLAG =          0x2000 //col:11507
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_MASK =          0x01 //col:11508
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL(_) =            (((_) >> 13) & 0x01) //col:11509
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_BIT =                        14 //col:11517
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_FLAG =                       0x4000 //col:11518
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_MASK =                       0x01 //col:11519
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT(_) =                         (((_) >> 14) & 0x01) //col:11520
IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_BIT =                       15 //col:11528
IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_FLAG =                      0x8000 //col:11529
IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_MASK =                      0x01 //col:11530
IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER(_) =                        (((_) >> 15) & 0x01) //col:11531
IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_BIT =                    16 //col:11537
IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_FLAG =                   0x10000 //col:11538
IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_MASK =                   0x01 //col:11539
IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS(_) =                     (((_) >> 16) & 0x01) //col:11540
IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_BIT =                  17 //col:11549
IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_FLAG =                 0x20000 //col:11550
IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_MASK =                 0x01 //col:11551
IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT(_) =                   (((_) >> 17) & 0x01) //col:11552
IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_BIT =                   18 //col:11558
IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_FLAG =                  0x40000 //col:11559
IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_MASK =                  0x01 //col:11560
IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL(_) =                    (((_) >> 18) & 0x01) //col:11561
IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_BIT =                       20 //col:11568
IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_FLAG =                      0x100000 //col:11569
IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_MASK =                      0x01 //col:11570
IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE(_) =                        (((_) >> 20) & 0x01) //col:11571
IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_BIT =                    21 //col:11577
IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_FLAG =                   0x200000 //col:11578
IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_MASK =                   0x01 //col:11579
IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL(_) =                     (((_) >> 21) & 0x01) //col:11580
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_BIT =                       22 //col:11586
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_FLAG =                      0x400000 //col:11587
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_MASK =                      0x01 //col:11588
IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS(_) =                        (((_) >> 22) & 0x01) //col:11589
IA32_VMX_MISC =                                                0x00000485 //col:11604
IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_BIT =          0 //col:11617
IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_FLAG =         0x1F //col:11618
IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_MASK =         0x1F //col:11619
IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP(_) =           (((_) >> 0) & 0x1F) //col:11620
IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_BIT =                   5 //col:11631
IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_FLAG =                  0x20 //col:11632
IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_MASK =                  0x01 //col:11633
IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT(_) =                    (((_) >> 5) & 0x01) //col:11634
IA32_VMX_MISC_ACTIVITY_STATES_BIT =                            6 //col:11647
IA32_VMX_MISC_ACTIVITY_STATES_FLAG =                           0x1C0 //col:11648
IA32_VMX_MISC_ACTIVITY_STATES_MASK =                           0x07 //col:11649
IA32_VMX_MISC_ACTIVITY_STATES(_) =                             (((_) >> 6) & 0x07) //col:11650
IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_BIT =                  14 //col:11663
IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_FLAG =                 0x4000 //col:11664
IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_MASK =                 0x01 //col:11665
IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX(_) =                   (((_) >> 14) & 0x01) //col:11666
IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_BIT =      15 //col:11677
IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_FLAG =     0x8000 //col:11678
IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_MASK =     0x01 //col:11679
IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM(_) =       (((_) >> 15) & 0x01) //col:11680
IA32_VMX_MISC_CR3_TARGET_COUNT_BIT =                           16 //col:11689
IA32_VMX_MISC_CR3_TARGET_COUNT_FLAG =                          0x1FF0000 //col:11690
IA32_VMX_MISC_CR3_TARGET_COUNT_MASK =                          0x1FF //col:11691
IA32_VMX_MISC_CR3_TARGET_COUNT(_) =                            (((_) >> 16) & 0x1FF) //col:11692
IA32_VMX_MISC_MAX_NUMBER_OF_MSR_BIT =                          25 //col:11703
IA32_VMX_MISC_MAX_NUMBER_OF_MSR_FLAG =                         0xE000000 //col:11704
IA32_VMX_MISC_MAX_NUMBER_OF_MSR_MASK =                         0x07 //col:11705
IA32_VMX_MISC_MAX_NUMBER_OF_MSR(_) =                           (((_) >> 25) & 0x07) //col:11706
IA32_VMX_MISC_SMM_MONITOR_CTL_B2_BIT =                         28 //col:11717
IA32_VMX_MISC_SMM_MONITOR_CTL_B2_FLAG =                        0x10000000 //col:11718
IA32_VMX_MISC_SMM_MONITOR_CTL_B2_MASK =                        0x01 //col:11719
IA32_VMX_MISC_SMM_MONITOR_CTL_B2(_) =                          (((_) >> 28) & 0x01) //col:11720
IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_BIT =                        29 //col:11729
IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_FLAG =                       0x20000000 //col:11730
IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_MASK =                       0x01 //col:11731
IA32_VMX_MISC_VMWRITE_VMEXIT_INFO(_) =                         (((_) >> 29) & 0x01) //col:11732
IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_BIT =  30 //col:11739
IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_FLAG = 0x40000000 //col:11740
IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_MASK = 0x01 //col:11741
IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION(_) =   (((_) >> 30) & 0x01) //col:11742
IA32_VMX_MISC_MSEG_ID_BIT =                                    32 //col:11751
IA32_VMX_MISC_MSEG_ID_FLAG =                                   0xFFFFFFFF00000000 //col:11752
IA32_VMX_MISC_MSEG_ID_MASK =                                   0xFFFFFFFF //col:11753
IA32_VMX_MISC_MSEG_ID(_) =                                     (((_) >> 32) & 0xFFFFFFFF) //col:11754
IA32_VMX_CR0_FIXED0 =                                          0x00000486 //col:11768
IA32_VMX_CR0_FIXED1 =                                          0x00000487 //col:11777
IA32_VMX_CR4_FIXED0 =                                          0x00000488 //col:11786
IA32_VMX_CR4_FIXED1 =                                          0x00000489 //col:11795
IA32_VMX_VMCS_ENUM =                                           0x0000048A //col:11804
IA32_VMX_VMCS_ENUM_ACCESS_TYPE_BIT =                           0 //col:11813
IA32_VMX_VMCS_ENUM_ACCESS_TYPE_FLAG =                          0x01 //col:11814
IA32_VMX_VMCS_ENUM_ACCESS_TYPE_MASK =                          0x01 //col:11815
IA32_VMX_VMCS_ENUM_ACCESS_TYPE(_) =                            (((_) >> 0) & 0x01) //col:11816
IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_BIT =                   1 //col:11822
IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_FLAG =                  0x3FE //col:11823
IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_MASK =                  0x1FF //col:11824
IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE(_) =                    (((_) >> 1) & 0x1FF) //col:11825
IA32_VMX_VMCS_ENUM_FIELD_TYPE_BIT =                            10 //col:11831
IA32_VMX_VMCS_ENUM_FIELD_TYPE_FLAG =                           0xC00 //col:11832
IA32_VMX_VMCS_ENUM_FIELD_TYPE_MASK =                           0x03 //col:11833
IA32_VMX_VMCS_ENUM_FIELD_TYPE(_) =                             (((_) >> 10) & 0x03) //col:11834
IA32_VMX_VMCS_ENUM_FIELD_WIDTH_BIT =                           13 //col:11841
IA32_VMX_VMCS_ENUM_FIELD_WIDTH_FLAG =                          0x6000 //col:11842
IA32_VMX_VMCS_ENUM_FIELD_WIDTH_MASK =                          0x03 //col:11843
IA32_VMX_VMCS_ENUM_FIELD_WIDTH(_) =                            (((_) >> 13) & 0x03) //col:11844
IA32_VMX_PROCBASED_CTLS2 =                                     0x0000048B //col:11859
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_BIT =        0 //col:11872
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_FLAG =       0x01 //col:11873
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_MASK =       0x01 //col:11874
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES(_) =         (((_) >> 0) & 0x01) //col:11875
IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_BIT =                      1 //col:11885
IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_FLAG =                     0x02 //col:11886
IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_MASK =                     0x01 //col:11887
IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT(_) =                       (((_) >> 1) & 0x01) //col:11888
IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_BIT =        2 //col:11896
IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_FLAG =       0x04 //col:11897
IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_MASK =       0x01 //col:11898
IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING(_) =         (((_) >> 2) & 0x01) //col:11899
IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_BIT =                   3 //col:11907
IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_FLAG =                  0x08 //col:11908
IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_MASK =                  0x01 //col:11909
IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP(_) =                    (((_) >> 3) & 0x01) //col:11910
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_BIT =          4 //col:11921
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_FLAG =         0x10 //col:11922
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_MASK =         0x01 //col:11923
IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE(_) =           (((_) >> 4) & 0x01) //col:11924
IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_BIT =                     5 //col:11935
IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_FLAG =                    0x20 //col:11936
IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_MASK =                    0x01 //col:11937
IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID(_) =                      (((_) >> 5) & 0x01) //col:11938
IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_BIT =                  6 //col:11946
IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_FLAG =                 0x40 //col:11947
IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_MASK =                 0x01 //col:11948
IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING(_) =                   (((_) >> 6) & 0x01) //col:11949
IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_BIT =              7 //col:11957
IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_FLAG =             0x80 //col:11958
IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_MASK =             0x01 //col:11959
IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST(_) =               (((_) >> 7) & 0x01) //col:11960
IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_BIT =    8 //col:11971
IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_FLAG =   0x100 //col:11972
IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_MASK =   0x01 //col:11973
IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION(_) =     (((_) >> 8) & 0x01) //col:11974
IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_BIT =      9 //col:11983
IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_FLAG =     0x200 //col:11984
IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_MASK =     0x01 //col:11985
IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY(_) =       (((_) >> 9) & 0x01) //col:11986
IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_BIT =              10 //col:11997
IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_FLAG =             0x400 //col:11998
IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_MASK =             0x01 //col:11999
IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING(_) =               (((_) >> 10) & 0x01) //col:12000
IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_BIT =                  11 //col:12008
IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_FLAG =                 0x800 //col:12009
IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_MASK =                 0x01 //col:12010
IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING(_) =                   (((_) >> 11) & 0x01) //col:12011
IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_BIT =                  12 //col:12019
IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_FLAG =                 0x1000 //col:12020
IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_MASK =                 0x01 //col:12021
IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID(_) =                   (((_) >> 12) & 0x01) //col:12022
IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_BIT =             13 //col:12032
IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_FLAG =            0x2000 //col:12033
IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_MASK =            0x01 //col:12034
IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS(_) =              (((_) >> 13) & 0x01) //col:12035
IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_BIT =                  14 //col:12047
IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_FLAG =                 0x4000 //col:12048
IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_MASK =                 0x01 //col:12049
IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING(_) =                   (((_) >> 14) & 0x01) //col:12050
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_BIT =            15 //col:12062
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_FLAG =           0x8000 //col:12063
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_MASK =           0x01 //col:12064
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING(_) =             (((_) >> 15) & 0x01) //col:12065
IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_BIT =                  16 //col:12073
IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_FLAG =                 0x10000 //col:12074
IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_MASK =                 0x01 //col:12075
IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING(_) =                   (((_) >> 16) & 0x01) //col:12076
IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_BIT =                      17 //col:12087
IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG =                     0x20000 //col:12088
IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_MASK =                     0x01 //col:12089
IA32_VMX_PROCBASED_CTLS2_ENABLE_PML(_) =                       (((_) >> 17) & 0x01) //col:12090
IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_BIT =                   18 //col:12100
IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_FLAG =                  0x40000 //col:12101
IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_MASK =                  0x01 //col:12102
IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION(_) =                    (((_) >> 18) & 0x01) //col:12103
IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_BIT =             19 //col:12114
IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_FLAG =            0x80000 //col:12115
IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_MASK =            0x01 //col:12116
IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT(_) =              (((_) >> 19) & 0x01) //col:12117
IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_BIT =                   20 //col:12125
IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_FLAG =                  0x100000 //col:12126
IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_MASK =                  0x01 //col:12127
IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES(_) =                    (((_) >> 20) & 0x01) //col:12128
IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_BIT = 22 //col:12138
IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG = 0x400000 //col:12139
IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_MASK = 0x01 //col:12140
IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT(_) = (((_) >> 22) & 0x01) //col:12141
IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_BIT = 23 //col:12149
IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_FLAG = 0x800000 //col:12150
IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_MASK = 0x01 //col:12151
IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT(_) = (((_) >> 23) & 0x01) //col:12152
IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_BIT = 24 //col:12161
IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_FLAG = 0x1000000 //col:12162
IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_MASK = 0x01 //col:12163
IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES(_) = (((_) >> 24) & 0x01) //col:12164
IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_BIT =                 25 //col:12176
IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_FLAG =                0x2000000 //col:12177
IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_MASK =                0x01 //col:12178
IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING(_) =                  (((_) >> 25) & 0x01) //col:12179
IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_BIT =          26 //col:12187
IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_FLAG =         0x4000000 //col:12188
IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_MASK =         0x01 //col:12189
IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE(_) =           (((_) >> 26) & 0x01) //col:12190
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_BIT =            28 //col:12203
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_FLAG =           0x10000000 //col:12204
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_MASK =           0x01 //col:12205
IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING(_) =             (((_) >> 28) & 0x01) //col:12206
IA32_VMX_EPT_VPID_CAP =                                        0x0000048C //col:12222
IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_BIT =                 0 //col:12233
IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_FLAG =                0x01 //col:12234
IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_MASK =                0x01 //col:12235
IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES(_) =                  (((_) >> 0) & 0x01) //col:12236
IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_BIT =                 6 //col:12243
IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_FLAG =                0x40 //col:12244
IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_MASK =                0x01 //col:12245
IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4(_) =                  (((_) >> 6) & 0x01) //col:12246
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_BIT =            8 //col:12256
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_FLAG =           0x100 //col:12257
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_MASK =           0x01 //col:12258
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE(_) =             (((_) >> 8) & 0x01) //col:12259
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_BIT =             14 //col:12267
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_FLAG =            0x4000 //col:12268
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_MASK =            0x01 //col:12269
IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK(_) =              (((_) >> 14) & 0x01) //col:12270
IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_BIT =                      16 //col:12278
IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_FLAG =                     0x10000 //col:12279
IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_MASK =                     0x01 //col:12280
IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES(_) =                       (((_) >> 16) & 0x01) //col:12281
IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_BIT =                    17 //col:12288
IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_FLAG =                   0x20000 //col:12289
IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_MASK =                   0x01 //col:12290
IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES(_) =                     (((_) >> 17) & 0x01) //col:12291
IA32_VMX_EPT_VPID_CAP_INVEPT_BIT =                             20 //col:12301
IA32_VMX_EPT_VPID_CAP_INVEPT_FLAG =                            0x100000 //col:12302
IA32_VMX_EPT_VPID_CAP_INVEPT_MASK =                            0x01 //col:12303
IA32_VMX_EPT_VPID_CAP_INVEPT(_) =                              (((_) >> 20) & 0x01) //col:12304
IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_BIT =       21 //col:12312
IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_FLAG =      0x200000 //col:12313
IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_MASK =      0x01 //col:12314
IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS(_) =        (((_) >> 21) & 0x01) //col:12315
IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_BIT = 22 //col:12324
IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_FLAG = 0x400000 //col:12325
IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_MASK = 0x01 //col:12326
IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION(_) = (((_) >> 22) & 0x01) //col:12327
IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_BIT =            23 //col:12335
IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_FLAG =           0x800000 //col:12336
IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_MASK =           0x01 //col:12337
IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK(_) =             (((_) >> 23) & 0x01) //col:12338
IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_BIT =              25 //col:12348
IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_FLAG =             0x2000000 //col:12349
IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_MASK =             0x01 //col:12350
IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT(_) =               (((_) >> 25) & 0x01) //col:12351
IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_BIT =                26 //col:12360
IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_FLAG =               0x4000000 //col:12361
IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_MASK =               0x01 //col:12362
IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS(_) =                 (((_) >> 26) & 0x01) //col:12363
IA32_VMX_EPT_VPID_CAP_INVVPID_BIT =                            32 //col:12370
IA32_VMX_EPT_VPID_CAP_INVVPID_FLAG =                           0x100000000 //col:12371
IA32_VMX_EPT_VPID_CAP_INVVPID_MASK =                           0x01 //col:12372
IA32_VMX_EPT_VPID_CAP_INVVPID(_) =                             (((_) >> 32) & 0x01) //col:12373
IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_BIT =         40 //col:12380
IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_FLAG =        0x10000000000 //col:12381
IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_MASK =        0x01 //col:12382
IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS(_) =          (((_) >> 40) & 0x01) //col:12383
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_BIT =             41 //col:12389
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_FLAG =            0x20000000000 //col:12390
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_MASK =            0x01 //col:12391
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT(_) =              (((_) >> 41) & 0x01) //col:12392
IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_BIT =               42 //col:12398
IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_FLAG =              0x40000000000 //col:12399
IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_MASK =              0x01 //col:12400
IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS(_) =                (((_) >> 42) & 0x01) //col:12401
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_BIT = 43 //col:12407
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_FLAG = 0x80000000000 //col:12408
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_MASK = 0x01 //col:12409
IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS(_) = (((_) >> 43) & 0x01) //col:12410
IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_BIT =               48 //col:12420
IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_FLAG =              0x3F000000000000 //col:12421
IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_MASK =              0x3F //col:12422
IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE(_) =                (((_) >> 48) & 0x3F) //col:12423
IA32_VMX_TRUE_PINBASED_CTLS =                                  0x0000048D //col:12448
IA32_VMX_TRUE_PROCBASED_CTLS =                                 0x0000048E //col:12449
IA32_VMX_TRUE_EXIT_CTLS =                                      0x0000048F //col:12450
IA32_VMX_TRUE_ENTRY_CTLS =                                     0x00000490 //col:12451
IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_BIT =                    0 //col:12461
IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_FLAG =                   0xFFFFFFFF //col:12462
IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_MASK =                   0xFFFFFFFF //col:12463
IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS(_) =                     (((_) >> 0) & 0xFFFFFFFF) //col:12464
IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_BIT =                    32 //col:12471
IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_FLAG =                   0xFFFFFFFF00000000 //col:12472
IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_MASK =                   0xFFFFFFFF //col:12473
IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS(_) =                     (((_) >> 32) & 0xFFFFFFFF) //col:12474
IA32_VMX_VMFUNC =                                              0x00000491 //col:12492
IA32_VMX_VMFUNC_EPTP_SWITCHING_BIT =                           0 //col:12503
IA32_VMX_VMFUNC_EPTP_SWITCHING_FLAG =                          0x01 //col:12504
IA32_VMX_VMFUNC_EPTP_SWITCHING_MASK =                          0x01 //col:12505
IA32_VMX_VMFUNC_EPTP_SWITCHING(_) =                            (((_) >> 0) & 0x01) //col:12506
IA32_VMX_PROCBASED_CTLS3 =                                     0x00000492 //col:12521
IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_BIT =               0 //col:12532
IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_FLAG =              0x01 //col:12533
IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_MASK =              0x01 //col:12534
IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING(_) =                (((_) >> 0) & 0x01) //col:12535
IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_BIT =                     1 //col:12545
IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_FLAG =                    0x02 //col:12546
IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_MASK =                    0x01 //col:12547
IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT(_) =                      (((_) >> 1) & 0x01) //col:12548
IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_BIT =                2 //col:12558
IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_FLAG =               0x04 //col:12559
IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_MASK =               0x01 //col:12560
IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE(_) =                 (((_) >> 2) & 0x01) //col:12561
IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_BIT =                    3 //col:12570
IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_FLAG =                   0x08 //col:12571
IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_MASK =                   0x01 //col:12572
IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING(_) =                     (((_) >> 3) & 0x01) //col:12573
IA32_VMX_EXIT_CTLS2 =                                          0x00000493 //col:12588
IA32_VMX_EXIT_CTLS2_RESERVED_BIT =                             0 //col:12594
IA32_VMX_EXIT_CTLS2_RESERVED_FLAG =                            0xFFFFFFFFFFFFFFFF //col:12595
IA32_VMX_EXIT_CTLS2_RESERVED_MASK =                            0xFFFFFFFFFFFFFFFF //col:12596
IA32_VMX_EXIT_CTLS2_RESERVED(_) =                              (((_) >> 0) & 0xFFFFFFFFFFFFFFFF) //col:12597
IA32_A_PMC0 =                                                  0x000004C1 //col:12612
IA32_A_PMC1 =                                                  0x000004C2 //col:12613
IA32_A_PMC2 =                                                  0x000004C3 //col:12614
IA32_A_PMC3 =                                                  0x000004C4 //col:12615
IA32_A_PMC4 =                                                  0x000004C5 //col:12616
IA32_A_PMC5 =                                                  0x000004C6 //col:12617
IA32_A_PMC6 =                                                  0x000004C7 //col:12618
IA32_A_PMC7 =                                                  0x000004C8 //col:12619
IA32_MCG_EXT_CTL =                                             0x000004D0 //col:12631
IA32_MCG_EXT_CTL_LMCE_EN_BIT =                                 0 //col:12637
IA32_MCG_EXT_CTL_LMCE_EN_FLAG =                                0x01 //col:12638
IA32_MCG_EXT_CTL_LMCE_EN_MASK =                                0x01 //col:12639
IA32_MCG_EXT_CTL_LMCE_EN(_) =                                  (((_) >> 0) & 0x01) //col:12640
IA32_SGX_SVN_STATUS =                                          0x00000500 //col:12660
IA32_SGX_SVN_STATUS_LOCK_BIT =                                 0 //col:12675
IA32_SGX_SVN_STATUS_LOCK_FLAG =                                0x01 //col:12676
IA32_SGX_SVN_STATUS_LOCK_MASK =                                0x01 //col:12677
IA32_SGX_SVN_STATUS_LOCK(_) =                                  (((_) >> 0) & 0x01) //col:12678
IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_BIT =                        16 //col:12690
IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_FLAG =                       0xFF0000 //col:12691
IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_MASK =                       0xFF //col:12692
IA32_SGX_SVN_STATUS_SGX_SVN_SINIT(_) =                         (((_) >> 16) & 0xFF) //col:12693
IA32_RTIT_OUTPUT_BASE =                                        0x00000560 //col:12708
IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_BIT =              7 //col:12730
IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_FLAG =             0xFFFFFFFFFF80 //col:12731
IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_MASK =             0x1FFFFFFFFFF //col:12732
IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS(_) =               (((_) >> 7) & 0x1FFFFFFFFFF) //col:12733
IA32_RTIT_OUTPUT_MASK_PTRS =                                   0x00000561 //col:12748
IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_BIT =                    0 //col:12757
IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_FLAG =                   0x7F //col:12758
IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_MASK =                   0x7F //col:12759
IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK(_) =                     (((_) >> 0) & 0x7F) //col:12760
IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_BIT =          7 //col:12779
IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_FLAG =         0xFFFFFF80 //col:12780
IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_MASK =         0x1FFFFFF //col:12781
IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET(_) =           (((_) >> 7) & 0x1FFFFFF) //col:12782
IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_BIT =                 32 //col:12800
IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_FLAG =                0xFFFFFFFF00000000 //col:12801
IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_MASK =                0xFFFFFFFF //col:12802
IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET(_) =                  (((_) >> 32) & 0xFFFFFFFF) //col:12803
IA32_RTIT_CTL =                                                0x00000570 //col:12816
IA32_RTIT_CTL_TRACE_ENABLED_BIT =                              0 //col:12834
IA32_RTIT_CTL_TRACE_ENABLED_FLAG =                             0x01 //col:12835
IA32_RTIT_CTL_TRACE_ENABLED_MASK =                             0x01 //col:12836
IA32_RTIT_CTL_TRACE_ENABLED(_) =                               (((_) >> 0) & 0x01) //col:12837
IA32_RTIT_CTL_CYC_ENABLED_BIT =                                1 //col:12849
IA32_RTIT_CTL_CYC_ENABLED_FLAG =                               0x02 //col:12850
IA32_RTIT_CTL_CYC_ENABLED_MASK =                               0x01 //col:12851
IA32_RTIT_CTL_CYC_ENABLED(_) =                                 (((_) >> 1) & 0x01) //col:12852
IA32_RTIT_CTL_OS_BIT =                                         2 //col:12861
IA32_RTIT_CTL_OS_FLAG =                                        0x04 //col:12862
IA32_RTIT_CTL_OS_MASK =                                        0x01 //col:12863
IA32_RTIT_CTL_OS(_) =                                          (((_) >> 2) & 0x01) //col:12864
IA32_RTIT_CTL_USER_BIT =                                       3 //col:12873
IA32_RTIT_CTL_USER_FLAG =                                      0x08 //col:12874
IA32_RTIT_CTL_USER_MASK =                                      0x01 //col:12875
IA32_RTIT_CTL_USER(_) =                                        (((_) >> 3) & 0x01) //col:12876
IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_BIT =                  4 //col:12887
IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_FLAG =                 0x10 //col:12888
IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_MASK =                 0x01 //col:12889
IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED(_) =                   (((_) >> 4) & 0x01) //col:12890
IA32_RTIT_CTL_FUP_ON_PTW_BIT =                                 5 //col:12899
IA32_RTIT_CTL_FUP_ON_PTW_FLAG =                                0x20 //col:12900
IA32_RTIT_CTL_FUP_ON_PTW_MASK =                                0x01 //col:12901
IA32_RTIT_CTL_FUP_ON_PTW(_) =                                  (((_) >> 5) & 0x01) //col:12902
IA32_RTIT_CTL_FABRIC_ENABLED_BIT =                             6 //col:12913
IA32_RTIT_CTL_FABRIC_ENABLED_FLAG =                            0x40 //col:12914
IA32_RTIT_CTL_FABRIC_ENABLED_MASK =                            0x01 //col:12915
IA32_RTIT_CTL_FABRIC_ENABLED(_) =                              (((_) >> 6) & 0x01) //col:12916
IA32_RTIT_CTL_CR3_FILTER_BIT =                                 7 //col:12925
IA32_RTIT_CTL_CR3_FILTER_FLAG =                                0x80 //col:12926
IA32_RTIT_CTL_CR3_FILTER_MASK =                                0x01 //col:12927
IA32_RTIT_CTL_CR3_FILTER(_) =                                  (((_) >> 7) & 0x01) //col:12928
IA32_RTIT_CTL_TOPA_BIT =                                       8 //col:12944
IA32_RTIT_CTL_TOPA_FLAG =                                      0x100 //col:12945
IA32_RTIT_CTL_TOPA_MASK =                                      0x01 //col:12946
IA32_RTIT_CTL_TOPA(_) =                                        (((_) >> 8) & 0x01) //col:12947
IA32_RTIT_CTL_MTC_ENABLED_BIT =                                9 //col:12959
IA32_RTIT_CTL_MTC_ENABLED_FLAG =                               0x200 //col:12960
IA32_RTIT_CTL_MTC_ENABLED_MASK =                               0x01 //col:12961
IA32_RTIT_CTL_MTC_ENABLED(_) =                                 (((_) >> 9) & 0x01) //col:12962
IA32_RTIT_CTL_TSC_ENABLED_BIT =                                10 //col:12973
IA32_RTIT_CTL_TSC_ENABLED_FLAG =                               0x400 //col:12974
IA32_RTIT_CTL_TSC_ENABLED_MASK =                               0x01 //col:12975
IA32_RTIT_CTL_TSC_ENABLED(_) =                                 (((_) >> 10) & 0x01) //col:12976
IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_BIT =                   11 //col:12987
IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_FLAG =                  0x800 //col:12988
IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_MASK =                  0x01 //col:12989
IA32_RTIT_CTL_RET_COMPRESSION_DISABLED(_) =                    (((_) >> 11) & 0x01) //col:12990
IA32_RTIT_CTL_PTW_ENABLED_BIT =                                12 //col:12999
IA32_RTIT_CTL_PTW_ENABLED_FLAG =                               0x1000 //col:13000
IA32_RTIT_CTL_PTW_ENABLED_MASK =                               0x01 //col:13001
IA32_RTIT_CTL_PTW_ENABLED(_) =                                 (((_) >> 12) & 0x01) //col:13002
IA32_RTIT_CTL_BRANCH_ENABLED_BIT =                             13 //col:13013
IA32_RTIT_CTL_BRANCH_ENABLED_FLAG =                            0x2000 //col:13014
IA32_RTIT_CTL_BRANCH_ENABLED_MASK =                            0x01 //col:13015
IA32_RTIT_CTL_BRANCH_ENABLED(_) =                              (((_) >> 13) & 0x01) //col:13016
IA32_RTIT_CTL_MTC_FREQUENCY_BIT =                              14 //col:13030
IA32_RTIT_CTL_MTC_FREQUENCY_FLAG =                             0x3C000 //col:13031
IA32_RTIT_CTL_MTC_FREQUENCY_MASK =                             0x0F //col:13032
IA32_RTIT_CTL_MTC_FREQUENCY(_) =                               (((_) >> 14) & 0x0F) //col:13033
IA32_RTIT_CTL_CYC_THRESHOLD_BIT =                              19 //col:13050
IA32_RTIT_CTL_CYC_THRESHOLD_FLAG =                             0x780000 //col:13051
IA32_RTIT_CTL_CYC_THRESHOLD_MASK =                             0x0F //col:13052
IA32_RTIT_CTL_CYC_THRESHOLD(_) =                               (((_) >> 19) & 0x0F) //col:13053
IA32_RTIT_CTL_PSB_FREQUENCY_BIT =                              24 //col:13070
IA32_RTIT_CTL_PSB_FREQUENCY_FLAG =                             0xF000000 //col:13071
IA32_RTIT_CTL_PSB_FREQUENCY_MASK =                             0x0F //col:13072
IA32_RTIT_CTL_PSB_FREQUENCY(_) =                               (((_) >> 24) & 0x0F) //col:13073
IA32_RTIT_CTL_ADDR0_CFG_BIT =                                  32 //col:13092
IA32_RTIT_CTL_ADDR0_CFG_FLAG =                                 0xF00000000 //col:13093
IA32_RTIT_CTL_ADDR0_CFG_MASK =                                 0x0F //col:13094
IA32_RTIT_CTL_ADDR0_CFG(_) =                                   (((_) >> 32) & 0x0F) //col:13095
IA32_RTIT_CTL_ADDR1_CFG_BIT =                                  36 //col:13113
IA32_RTIT_CTL_ADDR1_CFG_FLAG =                                 0xF000000000 //col:13114
IA32_RTIT_CTL_ADDR1_CFG_MASK =                                 0x0F //col:13115
IA32_RTIT_CTL_ADDR1_CFG(_) =                                   (((_) >> 36) & 0x0F) //col:13116
IA32_RTIT_CTL_ADDR2_CFG_BIT =                                  40 //col:13134
IA32_RTIT_CTL_ADDR2_CFG_FLAG =                                 0xF0000000000 //col:13135
IA32_RTIT_CTL_ADDR2_CFG_MASK =                                 0x0F //col:13136
IA32_RTIT_CTL_ADDR2_CFG(_) =                                   (((_) >> 40) & 0x0F) //col:13137
IA32_RTIT_CTL_ADDR3_CFG_BIT =                                  44 //col:13155
IA32_RTIT_CTL_ADDR3_CFG_FLAG =                                 0xF00000000000 //col:13156
IA32_RTIT_CTL_ADDR3_CFG_MASK =                                 0x0F //col:13157
IA32_RTIT_CTL_ADDR3_CFG(_) =                                   (((_) >> 44) & 0x0F) //col:13158
IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_BIT =                   56 //col:13171
IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_FLAG =                  0x100000000000000 //col:13172
IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_MASK =                  0x01 //col:13173
IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE(_) =                    (((_) >> 56) & 0x01) //col:13174
IA32_RTIT_STATUS =                                             0x00000571 //col:13187
IA32_RTIT_STATUS_FILTER_ENABLED_BIT =                          0 //col:13202
IA32_RTIT_STATUS_FILTER_ENABLED_FLAG =                         0x01 //col:13203
IA32_RTIT_STATUS_FILTER_ENABLED_MASK =                         0x01 //col:13204
IA32_RTIT_STATUS_FILTER_ENABLED(_) =                           (((_) >> 0) & 0x01) //col:13205
IA32_RTIT_STATUS_CONTEXT_ENABLED_BIT =                         1 //col:13215
IA32_RTIT_STATUS_CONTEXT_ENABLED_FLAG =                        0x02 //col:13216
IA32_RTIT_STATUS_CONTEXT_ENABLED_MASK =                        0x01 //col:13217
IA32_RTIT_STATUS_CONTEXT_ENABLED(_) =                          (((_) >> 1) & 0x01) //col:13218
IA32_RTIT_STATUS_TRIGGER_ENABLED_BIT =                         2 //col:13228
IA32_RTIT_STATUS_TRIGGER_ENABLED_FLAG =                        0x04 //col:13229
IA32_RTIT_STATUS_TRIGGER_ENABLED_MASK =                        0x01 //col:13230
IA32_RTIT_STATUS_TRIGGER_ENABLED(_) =                          (((_) >> 2) & 0x01) //col:13231
IA32_RTIT_STATUS_ERROR_BIT =                                   4 //col:13245
IA32_RTIT_STATUS_ERROR_FLAG =                                  0x10 //col:13246
IA32_RTIT_STATUS_ERROR_MASK =                                  0x01 //col:13247
IA32_RTIT_STATUS_ERROR(_) =                                    (((_) >> 4) & 0x01) //col:13248
IA32_RTIT_STATUS_STOPPED_BIT =                                 5 //col:13261
IA32_RTIT_STATUS_STOPPED_FLAG =                                0x20 //col:13262
IA32_RTIT_STATUS_STOPPED_MASK =                                0x01 //col:13263
IA32_RTIT_STATUS_STOPPED(_) =                                  (((_) >> 5) & 0x01) //col:13264
IA32_RTIT_STATUS_PEND_PSB_BIT =                                6 //col:13277
IA32_RTIT_STATUS_PEND_PSB_FLAG =                               0x40 //col:13278
IA32_RTIT_STATUS_PEND_PSB_MASK =                               0x01 //col:13279
IA32_RTIT_STATUS_PEND_PSB(_) =                                 (((_) >> 6) & 0x01) //col:13280
IA32_RTIT_STATUS_PEND_TOPA_PMI_BIT =                           7 //col:13293
IA32_RTIT_STATUS_PEND_TOPA_PMI_FLAG =                          0x80 //col:13294
IA32_RTIT_STATUS_PEND_TOPA_PMI_MASK =                          0x01 //col:13295
IA32_RTIT_STATUS_PEND_TOPA_PMI(_) =                            (((_) >> 7) & 0x01) //col:13296
IA32_RTIT_STATUS_PACKET_BYTE_COUNT_BIT =                       32 //col:13311
IA32_RTIT_STATUS_PACKET_BYTE_COUNT_FLAG =                      0x1FFFF00000000 //col:13312
IA32_RTIT_STATUS_PACKET_BYTE_COUNT_MASK =                      0x1FFFF //col:13313
IA32_RTIT_STATUS_PACKET_BYTE_COUNT(_) =                        (((_) >> 32) & 0x1FFFF) //col:13314
IA32_RTIT_CR3_MATCH =                                          0x00000572 //col:13332
IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_BIT =                   5 //col:13343
IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_FLAG =                  0xFFFFFFFFFFFFFFE0 //col:13344
IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_MASK =                  0x7FFFFFFFFFFFFFF //col:13345
IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH(_) =                    (((_) >> 5) & 0x7FFFFFFFFFFFFFF) //col:13346
IA32_RTIT_ADDR0_A =                                            0x00000580 //col:13373
IA32_RTIT_ADDR1_A =                                            0x00000582 //col:13374
IA32_RTIT_ADDR2_A =                                            0x00000584 //col:13375
IA32_RTIT_ADDR3_A =                                            0x00000586 //col:13376
IA32_RTIT_ADDR0_B =                                            0x00000581 //col:13390
IA32_RTIT_ADDR1_B =                                            0x00000583 //col:13391
IA32_RTIT_ADDR2_B =                                            0x00000585 //col:13392
IA32_RTIT_ADDR3_B =                                            0x00000587 //col:13393
IA32_RTIT_ADDR_VIRTUAL_ADDRESS_BIT =                           0 //col:13406
IA32_RTIT_ADDR_VIRTUAL_ADDRESS_FLAG =                          0xFFFFFFFFFFFF //col:13407
IA32_RTIT_ADDR_VIRTUAL_ADDRESS_MASK =                          0xFFFFFFFFFFFF //col:13408
IA32_RTIT_ADDR_VIRTUAL_ADDRESS(_) =                            (((_) >> 0) & 0xFFFFFFFFFFFF) //col:13409
IA32_RTIT_ADDR_SIGN_EXT_VA_BIT =                               48 //col:13415
IA32_RTIT_ADDR_SIGN_EXT_VA_FLAG =                              0xFFFF000000000000 //col:13416
IA32_RTIT_ADDR_SIGN_EXT_VA_MASK =                              0xFFFF //col:13417
IA32_RTIT_ADDR_SIGN_EXT_VA(_) =                                (((_) >> 48) & 0xFFFF) //col:13418
IA32_DS_AREA =                                                 0x00000600 //col:13440
IA32_U_CET =                                                   0x000006A0 //col:13448
IA32_U_CET_SH_STK_EN_BIT =                                     0 //col:13459
IA32_U_CET_SH_STK_EN_FLAG =                                    0x01 //col:13460
IA32_U_CET_SH_STK_EN_MASK =                                    0x01 //col:13461
IA32_U_CET_SH_STK_EN(_) =                                      (((_) >> 0) & 0x01) //col:13462
IA32_U_CET_WR_SHSTK_EN_BIT =                                   1 //col:13470
IA32_U_CET_WR_SHSTK_EN_FLAG =                                  0x02 //col:13471
IA32_U_CET_WR_SHSTK_EN_MASK =                                  0x01 //col:13472
IA32_U_CET_WR_SHSTK_EN(_) =                                    (((_) >> 1) & 0x01) //col:13473
IA32_U_CET_ENDBR_EN_BIT =                                      2 //col:13481
IA32_U_CET_ENDBR_EN_FLAG =                                     0x04 //col:13482
IA32_U_CET_ENDBR_EN_MASK =                                     0x01 //col:13483
IA32_U_CET_ENDBR_EN(_) =                                       (((_) >> 2) & 0x01) //col:13484
IA32_U_CET_LEG_IW_EN_BIT =                                     3 //col:13492
IA32_U_CET_LEG_IW_EN_FLAG =                                    0x08 //col:13493
IA32_U_CET_LEG_IW_EN_MASK =                                    0x01 //col:13494
IA32_U_CET_LEG_IW_EN(_) =                                      (((_) >> 3) & 0x01) //col:13495
IA32_U_CET_NO_TRACK_EN_BIT =                                   4 //col:13503
IA32_U_CET_NO_TRACK_EN_FLAG =                                  0x10 //col:13504
IA32_U_CET_NO_TRACK_EN_MASK =                                  0x01 //col:13505
IA32_U_CET_NO_TRACK_EN(_) =                                    (((_) >> 4) & 0x01) //col:13506
IA32_U_CET_SUPPRESS_DIS_BIT =                                  5 //col:13514
IA32_U_CET_SUPPRESS_DIS_FLAG =                                 0x20 //col:13515
IA32_U_CET_SUPPRESS_DIS_MASK =                                 0x01 //col:13516
IA32_U_CET_SUPPRESS_DIS(_) =                                   (((_) >> 5) & 0x01) //col:13517
IA32_U_CET_SUPPRESS_BIT =                                      10 //col:13527
IA32_U_CET_SUPPRESS_FLAG =                                     0x400 //col:13528
IA32_U_CET_SUPPRESS_MASK =                                     0x01 //col:13529
IA32_U_CET_SUPPRESS(_) =                                       (((_) >> 10) & 0x01) //col:13530
IA32_U_CET_TRACKER_BIT =                                       11 //col:13538
IA32_U_CET_TRACKER_FLAG =                                      0x800 //col:13539
IA32_U_CET_TRACKER_MASK =                                      0x01 //col:13540
IA32_U_CET_TRACKER(_) =                                        (((_) >> 11) & 0x01) //col:13541
IA32_U_CET_EB_LEG_BITMAP_BASE_BIT =                            12 //col:13553
IA32_U_CET_EB_LEG_BITMAP_BASE_FLAG =                           0xFFFFFFFFFFFFF000 //col:13554
IA32_U_CET_EB_LEG_BITMAP_BASE_MASK =                           0xFFFFFFFFFFFFF //col:13555
IA32_U_CET_EB_LEG_BITMAP_BASE(_) =                             (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:13556
IA32_S_CET =                                                   0x000006A2 //col:13569
IA32_S_CET_SH_STK_EN_BIT =                                     0 //col:13580
IA32_S_CET_SH_STK_EN_FLAG =                                    0x01 //col:13581
IA32_S_CET_SH_STK_EN_MASK =                                    0x01 //col:13582
IA32_S_CET_SH_STK_EN(_) =                                      (((_) >> 0) & 0x01) //col:13583
IA32_S_CET_WR_SHSTK_EN_BIT =                                   1 //col:13591
IA32_S_CET_WR_SHSTK_EN_FLAG =                                  0x02 //col:13592
IA32_S_CET_WR_SHSTK_EN_MASK =                                  0x01 //col:13593
IA32_S_CET_WR_SHSTK_EN(_) =                                    (((_) >> 1) & 0x01) //col:13594
IA32_S_CET_ENDBR_EN_BIT =                                      2 //col:13602
IA32_S_CET_ENDBR_EN_FLAG =                                     0x04 //col:13603
IA32_S_CET_ENDBR_EN_MASK =                                     0x01 //col:13604
IA32_S_CET_ENDBR_EN(_) =                                       (((_) >> 2) & 0x01) //col:13605
IA32_S_CET_LEG_IW_EN_BIT =                                     3 //col:13613
IA32_S_CET_LEG_IW_EN_FLAG =                                    0x08 //col:13614
IA32_S_CET_LEG_IW_EN_MASK =                                    0x01 //col:13615
IA32_S_CET_LEG_IW_EN(_) =                                      (((_) >> 3) & 0x01) //col:13616
IA32_S_CET_NO_TRACK_EN_BIT =                                   4 //col:13624
IA32_S_CET_NO_TRACK_EN_FLAG =                                  0x10 //col:13625
IA32_S_CET_NO_TRACK_EN_MASK =                                  0x01 //col:13626
IA32_S_CET_NO_TRACK_EN(_) =                                    (((_) >> 4) & 0x01) //col:13627
IA32_S_CET_SUPPRESS_DIS_BIT =                                  5 //col:13635
IA32_S_CET_SUPPRESS_DIS_FLAG =                                 0x20 //col:13636
IA32_S_CET_SUPPRESS_DIS_MASK =                                 0x01 //col:13637
IA32_S_CET_SUPPRESS_DIS(_) =                                   (((_) >> 5) & 0x01) //col:13638
IA32_S_CET_SUPPRESS_BIT =                                      10 //col:13648
IA32_S_CET_SUPPRESS_FLAG =                                     0x400 //col:13649
IA32_S_CET_SUPPRESS_MASK =                                     0x01 //col:13650
IA32_S_CET_SUPPRESS(_) =                                       (((_) >> 10) & 0x01) //col:13651
IA32_S_CET_TRACKER_BIT =                                       11 //col:13659
IA32_S_CET_TRACKER_FLAG =                                      0x800 //col:13660
IA32_S_CET_TRACKER_MASK =                                      0x01 //col:13661
IA32_S_CET_TRACKER(_) =                                        (((_) >> 11) & 0x01) //col:13662
IA32_S_CET_EB_LEG_BITMAP_BASE_BIT =                            12 //col:13674
IA32_S_CET_EB_LEG_BITMAP_BASE_FLAG =                           0xFFFFFFFFFFFFF000 //col:13675
IA32_S_CET_EB_LEG_BITMAP_BASE_MASK =                           0xFFFFFFFFFFFFF //col:13676
IA32_S_CET_EB_LEG_BITMAP_BASE(_) =                             (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:13677
IA32_PL0_SSP =                                                 0x000006A4 //col:13693
IA32_PL1_SSP =                                                 0x000006A5 //col:13704
IA32_PL2_SSP =                                                 0x000006A6 //col:13715
IA32_PL3_SSP =                                                 0x000006A7 //col:13726
IA32_INTERRUPT_SSP_TABLE_ADDR =                                0x000006A8 //col:13736
IA32_TSC_DEADLINE =                                            0x000006E0 //col:13743
IA32_PM_ENABLE =                                               0x00000770 //col:13750
IA32_PM_ENABLE_HWP_ENABLE_BIT =                                0 //col:13762
IA32_PM_ENABLE_HWP_ENABLE_FLAG =                               0x01 //col:13763
IA32_PM_ENABLE_HWP_ENABLE_MASK =                               0x01 //col:13764
IA32_PM_ENABLE_HWP_ENABLE(_) =                                 (((_) >> 0) & 0x01) //col:13765
IA32_HWP_CAPABILITIES =                                        0x00000771 //col:13778
IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_BIT =                0 //col:13790
IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_FLAG =               0xFF //col:13791
IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_MASK =               0xFF //col:13792
IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE(_) =                 (((_) >> 0) & 0xFF) //col:13793
IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_BIT =             8 //col:13802
IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_FLAG =            0xFF00 //col:13803
IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_MASK =            0xFF //col:13804
IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE(_) =              (((_) >> 8) & 0xFF) //col:13805
IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_BIT =         16 //col:13814
IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_FLAG =        0xFF0000 //col:13815
IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_MASK =        0xFF //col:13816
IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE(_) =          (((_) >> 16) & 0xFF) //col:13817
IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_BIT =                 24 //col:13826
IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_FLAG =                0xFF000000 //col:13827
IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_MASK =                0xFF //col:13828
IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE(_) =                  (((_) >> 24) & 0xFF) //col:13829
IA32_HWP_REQUEST_PKG =                                         0x00000772 //col:13842
IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_BIT =                 0 //col:13854
IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_FLAG =                0xFF //col:13855
IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_MASK =                0xFF //col:13856
IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE(_) =                  (((_) >> 0) & 0xFF) //col:13857
IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_BIT =                 8 //col:13866
IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_FLAG =                0xFF00 //col:13867
IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_MASK =                0xFF //col:13868
IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE(_) =                  (((_) >> 8) & 0xFF) //col:13869
IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_BIT =                 16 //col:13878
IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_FLAG =                0xFF0000 //col:13879
IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_MASK =                0xFF //col:13880
IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE(_) =                  (((_) >> 16) & 0xFF) //col:13881
IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_BIT =       24 //col:13890
IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_FLAG =      0xFF000000 //col:13891
IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_MASK =      0xFF //col:13892
IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE(_) =        (((_) >> 24) & 0xFF) //col:13893
IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_BIT =                     32 //col:13902
IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_FLAG =                    0x3FF00000000 //col:13903
IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_MASK =                    0x3FF //col:13904
IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW(_) =                      (((_) >> 32) & 0x3FF) //col:13905
IA32_HWP_INTERRUPT =                                           0x00000773 //col:13918
IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_BIT =      0 //col:13930
IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_FLAG =     0x01 //col:13931
IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_MASK =     0x01 //col:13932
IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE(_) =       (((_) >> 0) & 0x01) //col:13933
IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_BIT =                  1 //col:13942
IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_FLAG =                 0x02 //col:13943
IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_MASK =                 0x01 //col:13944
IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM(_) =                   (((_) >> 1) & 0x01) //col:13945
IA32_HWP_REQUEST =                                             0x00000774 //col:13958
IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_BIT =                     0 //col:13970
IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_FLAG =                    0xFF //col:13971
IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_MASK =                    0xFF //col:13972
IA32_HWP_REQUEST_MINIMUM_PERFORMANCE(_) =                      (((_) >> 0) & 0xFF) //col:13973
IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_BIT =                     8 //col:13982
IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_FLAG =                    0xFF00 //col:13983
IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_MASK =                    0xFF //col:13984
IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE(_) =                      (((_) >> 8) & 0xFF) //col:13985
IA32_HWP_REQUEST_DESIRED_PERFORMANCE_BIT =                     16 //col:13994
IA32_HWP_REQUEST_DESIRED_PERFORMANCE_FLAG =                    0xFF0000 //col:13995
IA32_HWP_REQUEST_DESIRED_PERFORMANCE_MASK =                    0xFF //col:13996
IA32_HWP_REQUEST_DESIRED_PERFORMANCE(_) =                      (((_) >> 16) & 0xFF) //col:13997
IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_BIT =           24 //col:14006
IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_FLAG =          0xFF000000 //col:14007
IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_MASK =          0xFF //col:14008
IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE(_) =            (((_) >> 24) & 0xFF) //col:14009
IA32_HWP_REQUEST_ACTIVITY_WINDOW_BIT =                         32 //col:14018
IA32_HWP_REQUEST_ACTIVITY_WINDOW_FLAG =                        0x3FF00000000 //col:14019
IA32_HWP_REQUEST_ACTIVITY_WINDOW_MASK =                        0x3FF //col:14020
IA32_HWP_REQUEST_ACTIVITY_WINDOW(_) =                          (((_) >> 32) & 0x3FF) //col:14021
IA32_HWP_REQUEST_PACKAGE_CONTROL_BIT =                         42 //col:14030
IA32_HWP_REQUEST_PACKAGE_CONTROL_FLAG =                        0x40000000000 //col:14031
IA32_HWP_REQUEST_PACKAGE_CONTROL_MASK =                        0x01 //col:14032
IA32_HWP_REQUEST_PACKAGE_CONTROL(_) =                          (((_) >> 42) & 0x01) //col:14033
IA32_HWP_STATUS =                                              0x00000777 //col:14046
IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_BIT =            0 //col:14058
IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_FLAG =           0x01 //col:14059
IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_MASK =           0x01 //col:14060
IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE(_) =             (((_) >> 0) & 0x01) //col:14061
IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_BIT =                     2 //col:14071
IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_FLAG =                    0x04 //col:14072
IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_MASK =                    0x01 //col:14073
IA32_HWP_STATUS_EXCURSION_TO_MINIMUM(_) =                      (((_) >> 2) & 0x01) //col:14074
IA32_X2APIC_APICID =                                           0x00000802 //col:14088
IA32_X2APIC_VERSION =                                          0x00000803 //col:14095
IA32_X2APIC_TPR =                                              0x00000808 //col:14102
IA32_X2APIC_PPR =                                              0x0000080A //col:14109
IA32_X2APIC_EOI =                                              0x0000080B //col:14116
IA32_X2APIC_LDR =                                              0x0000080D //col:14123
IA32_X2APIC_SIVR =                                             0x0000080F //col:14130
IA32_X2APIC_ISR0 =                                             0x00000810 //col:14140
IA32_X2APIC_ISR1 =                                             0x00000811 //col:14141
IA32_X2APIC_ISR2 =                                             0x00000812 //col:14142
IA32_X2APIC_ISR3 =                                             0x00000813 //col:14143
IA32_X2APIC_ISR4 =                                             0x00000814 //col:14144
IA32_X2APIC_ISR5 =                                             0x00000815 //col:14145
IA32_X2APIC_ISR6 =                                             0x00000816 //col:14146
IA32_X2APIC_ISR7 =                                             0x00000817 //col:14147
IA32_X2APIC_TMR0 =                                             0x00000818 //col:14161
IA32_X2APIC_TMR1 =                                             0x00000819 //col:14162
IA32_X2APIC_TMR2 =                                             0x0000081A //col:14163
IA32_X2APIC_TMR3 =                                             0x0000081B //col:14164
IA32_X2APIC_TMR4 =                                             0x0000081C //col:14165
IA32_X2APIC_TMR5 =                                             0x0000081D //col:14166
IA32_X2APIC_TMR6 =                                             0x0000081E //col:14167
IA32_X2APIC_TMR7 =                                             0x0000081F //col:14168
IA32_X2APIC_IRR0 =                                             0x00000820 //col:14182
IA32_X2APIC_IRR1 =                                             0x00000821 //col:14183
IA32_X2APIC_IRR2 =                                             0x00000822 //col:14184
IA32_X2APIC_IRR3 =                                             0x00000823 //col:14185
IA32_X2APIC_IRR4 =                                             0x00000824 //col:14186
IA32_X2APIC_IRR5 =                                             0x00000825 //col:14187
IA32_X2APIC_IRR6 =                                             0x00000826 //col:14188
IA32_X2APIC_IRR7 =                                             0x00000827 //col:14189
IA32_X2APIC_ESR =                                              0x00000828 //col:14200
IA32_X2APIC_LVT_CMCI =                                         0x0000082F //col:14207
IA32_X2APIC_ICR =                                              0x00000830 //col:14214
IA32_X2APIC_LVT_TIMER =                                        0x00000832 //col:14221
IA32_X2APIC_LVT_THERMAL =                                      0x00000833 //col:14228
IA32_X2APIC_LVT_PMI =                                          0x00000834 //col:14235
IA32_X2APIC_LVT_LINT0 =                                        0x00000835 //col:14242
IA32_X2APIC_LVT_LINT1 =                                        0x00000836 //col:14249
IA32_X2APIC_LVT_ERROR =                                        0x00000837 //col:14256
IA32_X2APIC_INIT_COUNT =                                       0x00000838 //col:14263
IA32_X2APIC_CUR_COUNT =                                        0x00000839 //col:14270
IA32_X2APIC_DIV_CONF =                                         0x0000083E //col:14277
IA32_X2APIC_SELF_IPI =                                         0x0000083F //col:14284
IA32_DEBUG_INTERFACE =                                         0x00000C80 //col:14291
IA32_DEBUG_INTERFACE_ENABLE_BIT =                              0 //col:14304
IA32_DEBUG_INTERFACE_ENABLE_FLAG =                             0x01 //col:14305
IA32_DEBUG_INTERFACE_ENABLE_MASK =                             0x01 //col:14306
IA32_DEBUG_INTERFACE_ENABLE(_) =                               (((_) >> 0) & 0x01) //col:14307
IA32_DEBUG_INTERFACE_LOCK_BIT =                                30 //col:14319
IA32_DEBUG_INTERFACE_LOCK_FLAG =                               0x40000000 //col:14320
IA32_DEBUG_INTERFACE_LOCK_MASK =                               0x01 //col:14321
IA32_DEBUG_INTERFACE_LOCK(_) =                                 (((_) >> 30) & 0x01) //col:14322
IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_BIT =                      31 //col:14332
IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_FLAG =                     0x80000000 //col:14333
IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_MASK =                     0x01 //col:14334
IA32_DEBUG_INTERFACE_DEBUG_OCCURRED(_) =                       (((_) >> 31) & 0x01) //col:14335
IA32_L3_QOS_CFG =                                              0x00000C81 //col:14348
IA32_L3_QOS_CFG_ENABLE_BIT =                                   0 //col:14359
IA32_L3_QOS_CFG_ENABLE_FLAG =                                  0x01 //col:14360
IA32_L3_QOS_CFG_ENABLE_MASK =                                  0x01 //col:14361
IA32_L3_QOS_CFG_ENABLE(_) =                                    (((_) >> 0) & 0x01) //col:14362
IA32_L2_QOS_CFG =                                              0x00000C82 //col:14375
IA32_L2_QOS_CFG_ENABLE_BIT =                                   0 //col:14386
IA32_L2_QOS_CFG_ENABLE_FLAG =                                  0x01 //col:14387
IA32_L2_QOS_CFG_ENABLE_MASK =                                  0x01 //col:14388
IA32_L2_QOS_CFG_ENABLE(_) =                                    (((_) >> 0) & 0x01) //col:14389
IA32_QM_EVTSEL =                                               0x00000C8D //col:14402
IA32_QM_EVTSEL_EVENT_ID_BIT =                                  0 //col:14413
IA32_QM_EVTSEL_EVENT_ID_FLAG =                                 0xFF //col:14414
IA32_QM_EVTSEL_EVENT_ID_MASK =                                 0xFF //col:14415
IA32_QM_EVTSEL_EVENT_ID(_) =                                   (((_) >> 0) & 0xFF) //col:14416
IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_BIT =                    32 //col:14427
IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_FLAG =                   0xFFFFFFFF00000000 //col:14428
IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_MASK =                   0xFFFFFFFF //col:14429
IA32_QM_EVTSEL_RESOURCE_MONITORING_ID(_) =                     (((_) >> 32) & 0xFFFFFFFF) //col:14430
IA32_QM_CTR =                                                  0x00000C8E //col:14442
IA32_QM_CTR_RESOURCE_MONITORED_DATA_BIT =                      0 //col:14451
IA32_QM_CTR_RESOURCE_MONITORED_DATA_FLAG =                     0x3FFFFFFFFFFFFFFF //col:14452
IA32_QM_CTR_RESOURCE_MONITORED_DATA_MASK =                     0x3FFFFFFFFFFFFFFF //col:14453
IA32_QM_CTR_RESOURCE_MONITORED_DATA(_) =                       (((_) >> 0) & 0x3FFFFFFFFFFFFFFF) //col:14454
IA32_QM_CTR_UNAVAILABLE_BIT =                                  62 //col:14462
IA32_QM_CTR_UNAVAILABLE_FLAG =                                 0x4000000000000000 //col:14463
IA32_QM_CTR_UNAVAILABLE_MASK =                                 0x01 //col:14464
IA32_QM_CTR_UNAVAILABLE(_) =                                   (((_) >> 62) & 0x01) //col:14465
IA32_QM_CTR_ERROR_BIT =                                        63 //col:14473
IA32_QM_CTR_ERROR_FLAG =                                       0x8000000000000000 //col:14474
IA32_QM_CTR_ERROR_MASK =                                       0x01 //col:14475
IA32_QM_CTR_ERROR(_) =                                         (((_) >> 63) & 0x01) //col:14476
IA32_PQR_ASSOC =                                               0x00000C8F //col:14488
IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_BIT =                    0 //col:14501
IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_FLAG =                   0xFFFFFFFF //col:14502
IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_MASK =                   0xFFFFFFFF //col:14503
IA32_PQR_ASSOC_RESOURCE_MONITORING_ID(_) =                     (((_) >> 0) & 0xFFFFFFFF) //col:14504
IA32_PQR_ASSOC_COS_BIT =                                       32 //col:14514
IA32_PQR_ASSOC_COS_FLAG =                                      0xFFFFFFFF00000000 //col:14515
IA32_PQR_ASSOC_COS_MASK =                                      0xFFFFFFFF //col:14516
IA32_PQR_ASSOC_COS(_) =                                        (((_) >> 32) & 0xFFFFFFFF) //col:14517
IA32_BNDCFGS =                                                 0x00000D90 //col:14529
IA32_BNDCFGS_ENABLE_BIT =                                      0 //col:14538
IA32_BNDCFGS_ENABLE_FLAG =                                     0x01 //col:14539
IA32_BNDCFGS_ENABLE_MASK =                                     0x01 //col:14540
IA32_BNDCFGS_ENABLE(_) =                                       (((_) >> 0) & 0x01) //col:14541
IA32_BNDCFGS_BND_PRESERVE_BIT =                                1 //col:14547
IA32_BNDCFGS_BND_PRESERVE_FLAG =                               0x02 //col:14548
IA32_BNDCFGS_BND_PRESERVE_MASK =                               0x01 //col:14549
IA32_BNDCFGS_BND_PRESERVE(_) =                                 (((_) >> 1) & 0x01) //col:14550
IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_BIT =                12 //col:14557
IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_FLAG =               0xFFFFFFFFFFFFF000 //col:14558
IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_MASK =               0xFFFFFFFFFFFFF //col:14559
IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS(_) =                 (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:14560
IA32_XSS =                                                     0x00000DA0 //col:14572
IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_BIT =                8 //col:14583
IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_FLAG =               0x100 //col:14584
IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_MASK =               0x01 //col:14585
IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE(_) =                 (((_) >> 8) & 0x01) //col:14586
IA32_PKG_HDC_CTL =                                             0x00000DB0 //col:14599
IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_BIT =                          0 //col:14613
IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_FLAG =                         0x01 //col:14614
IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_MASK =                         0x01 //col:14615
IA32_PKG_HDC_CTL_HDC_PKG_ENABLE(_) =                           (((_) >> 0) & 0x01) //col:14616
IA32_PM_CTL1 =                                                 0x00000DB1 //col:14629
IA32_PM_CTL1_HDC_ALLOW_BLOCK_BIT =                             0 //col:14643
IA32_PM_CTL1_HDC_ALLOW_BLOCK_FLAG =                            0x01 //col:14644
IA32_PM_CTL1_HDC_ALLOW_BLOCK_MASK =                            0x01 //col:14645
IA32_PM_CTL1_HDC_ALLOW_BLOCK(_) =                              (((_) >> 0) & 0x01) //col:14646
IA32_THREAD_STALL =                                            0x00000DB2 //col:14659
IA32_EFER =                                                    0xC0000080 //col:14679
IA32_EFER_SYSCALL_ENABLE_BIT =                                 0 //col:14690
IA32_EFER_SYSCALL_ENABLE_FLAG =                                0x01 //col:14691
IA32_EFER_SYSCALL_ENABLE_MASK =                                0x01 //col:14692
IA32_EFER_SYSCALL_ENABLE(_) =                                  (((_) >> 0) & 0x01) //col:14693
IA32_EFER_IA32E_MODE_ENABLE_BIT =                              8 //col:14702
IA32_EFER_IA32E_MODE_ENABLE_FLAG =                             0x100 //col:14703
IA32_EFER_IA32E_MODE_ENABLE_MASK =                             0x01 //col:14704
IA32_EFER_IA32E_MODE_ENABLE(_) =                               (((_) >> 8) & 0x01) //col:14705
IA32_EFER_IA32E_MODE_ACTIVE_BIT =                              10 //col:14714
IA32_EFER_IA32E_MODE_ACTIVE_FLAG =                             0x400 //col:14715
IA32_EFER_IA32E_MODE_ACTIVE_MASK =                             0x01 //col:14716
IA32_EFER_IA32E_MODE_ACTIVE(_) =                               (((_) >> 10) & 0x01) //col:14717
IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_BIT =                     11 //col:14723
IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_FLAG =                    0x800 //col:14724
IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_MASK =                    0x01 //col:14725
IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE(_) =                      (((_) >> 11) & 0x01) //col:14726
IA32_STAR =                                                    0xC0000081 //col:14739
IA32_LSTAR =                                                   0xC0000082 //col:14748
IA32_CSTAR =                                                   0xC0000083 //col:14757
IA32_FMASK =                                                   0xC0000084 //col:14764
IA32_FS_BASE =                                                 0xC0000100 //col:14771
IA32_GS_BASE =                                                 0xC0000101 //col:14778
IA32_KERNEL_GS_BASE =                                          0xC0000102 //col:14785
IA32_TSC_AUX =                                                 0xC0000103 //col:14792
IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_BIT =                     0 //col:14801
IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_FLAG =                    0xFFFFFFFF //col:14802
IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_MASK =                    0xFFFFFFFF //col:14803
IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE(_) =                      (((_) >> 0) & 0xFFFFFFFF) //col:14804
PDE_4MB_32_PRESENT_BIT =                                       0 //col:14845
PDE_4MB_32_PRESENT_FLAG =                                      0x01 //col:14846
PDE_4MB_32_PRESENT_MASK =                                      0x01 //col:14847
PDE_4MB_32_PRESENT(_) =                                        (((_) >> 0) & 0x01) //col:14848
PDE_4MB_32_WRITE_BIT =                                         1 //col:14856
PDE_4MB_32_WRITE_FLAG =                                        0x02 //col:14857
PDE_4MB_32_WRITE_MASK =                                        0x01 //col:14858
PDE_4MB_32_WRITE(_) =                                          (((_) >> 1) & 0x01) //col:14859
PDE_4MB_32_SUPERVISOR_BIT =                                    2 //col:14867
PDE_4MB_32_SUPERVISOR_FLAG =                                   0x04 //col:14868
PDE_4MB_32_SUPERVISOR_MASK =                                   0x01 //col:14869
PDE_4MB_32_SUPERVISOR(_) =                                     (((_) >> 2) & 0x01) //col:14870
PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_BIT =                      3 //col:14879
PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_FLAG =                     0x08 //col:14880
PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_MASK =                     0x01 //col:14881
PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH(_) =                       (((_) >> 3) & 0x01) //col:14882
PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_BIT =                      4 //col:14891
PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_FLAG =                     0x10 //col:14892
PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_MASK =                     0x01 //col:14893
PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE(_) =                       (((_) >> 4) & 0x01) //col:14894
PDE_4MB_32_ACCESSED_BIT =                                      5 //col:14902
PDE_4MB_32_ACCESSED_FLAG =                                     0x20 //col:14903
PDE_4MB_32_ACCESSED_MASK =                                     0x01 //col:14904
PDE_4MB_32_ACCESSED(_) =                                       (((_) >> 5) & 0x01) //col:14905
PDE_4MB_32_DIRTY_BIT =                                         6 //col:14913
PDE_4MB_32_DIRTY_FLAG =                                        0x40 //col:14914
PDE_4MB_32_DIRTY_MASK =                                        0x01 //col:14915
PDE_4MB_32_DIRTY(_) =                                          (((_) >> 6) & 0x01) //col:14916
PDE_4MB_32_LARGE_PAGE_BIT =                                    7 //col:14922
PDE_4MB_32_LARGE_PAGE_FLAG =                                   0x80 //col:14923
PDE_4MB_32_LARGE_PAGE_MASK =                                   0x01 //col:14924
PDE_4MB_32_LARGE_PAGE(_) =                                     (((_) >> 7) & 0x01) //col:14925
PDE_4MB_32_GLOBAL_BIT =                                        8 //col:14933
PDE_4MB_32_GLOBAL_FLAG =                                       0x100 //col:14934
PDE_4MB_32_GLOBAL_MASK =                                       0x01 //col:14935
PDE_4MB_32_GLOBAL(_) =                                         (((_) >> 8) & 0x01) //col:14936
PDE_4MB_32_IGNORED_1_BIT =                                     9 //col:14942
PDE_4MB_32_IGNORED_1_FLAG =                                    0xE00 //col:14943
PDE_4MB_32_IGNORED_1_MASK =                                    0x07 //col:14944
PDE_4MB_32_IGNORED_1(_) =                                      (((_) >> 9) & 0x07) //col:14945
PDE_4MB_32_PAT_BIT =                                           12 //col:14953
PDE_4MB_32_PAT_FLAG =                                          0x1000 //col:14954
PDE_4MB_32_PAT_MASK =                                          0x01 //col:14955
PDE_4MB_32_PAT(_) =                                            (((_) >> 12) & 0x01) //col:14956
PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_BIT =                         13 //col:14962
PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_FLAG =                        0x1FE000 //col:14963
PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_MASK =                        0xFF //col:14964
PDE_4MB_32_PAGE_FRAME_NUMBER_LOW(_) =                          (((_) >> 13) & 0xFF) //col:14965
PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_BIT =                        22 //col:14972
PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_FLAG =                       0xFFC00000 //col:14973
PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_MASK =                       0x3FF //col:14974
PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH(_) =                         (((_) >> 22) & 0x3FF) //col:14975
PDE_32_PRESENT_BIT =                                           0 //col:14992
PDE_32_PRESENT_FLAG =                                          0x01 //col:14993
PDE_32_PRESENT_MASK =                                          0x01 //col:14994
PDE_32_PRESENT(_) =                                            (((_) >> 0) & 0x01) //col:14995
PDE_32_WRITE_BIT =                                             1 //col:15003
PDE_32_WRITE_FLAG =                                            0x02 //col:15004
PDE_32_WRITE_MASK =                                            0x01 //col:15005
PDE_32_WRITE(_) =                                              (((_) >> 1) & 0x01) //col:15006
PDE_32_SUPERVISOR_BIT =                                        2 //col:15014
PDE_32_SUPERVISOR_FLAG =                                       0x04 //col:15015
PDE_32_SUPERVISOR_MASK =                                       0x01 //col:15016
PDE_32_SUPERVISOR(_) =                                         (((_) >> 2) & 0x01) //col:15017
PDE_32_PAGE_LEVEL_WRITE_THROUGH_BIT =                          3 //col:15026
PDE_32_PAGE_LEVEL_WRITE_THROUGH_FLAG =                         0x08 //col:15027
PDE_32_PAGE_LEVEL_WRITE_THROUGH_MASK =                         0x01 //col:15028
PDE_32_PAGE_LEVEL_WRITE_THROUGH(_) =                           (((_) >> 3) & 0x01) //col:15029
PDE_32_PAGE_LEVEL_CACHE_DISABLE_BIT =                          4 //col:15038
PDE_32_PAGE_LEVEL_CACHE_DISABLE_FLAG =                         0x10 //col:15039
PDE_32_PAGE_LEVEL_CACHE_DISABLE_MASK =                         0x01 //col:15040
PDE_32_PAGE_LEVEL_CACHE_DISABLE(_) =                           (((_) >> 4) & 0x01) //col:15041
PDE_32_ACCESSED_BIT =                                          5 //col:15049
PDE_32_ACCESSED_FLAG =                                         0x20 //col:15050
PDE_32_ACCESSED_MASK =                                         0x01 //col:15051
PDE_32_ACCESSED(_) =                                           (((_) >> 5) & 0x01) //col:15052
PDE_32_IGNORED_1_BIT =                                         6 //col:15058
PDE_32_IGNORED_1_FLAG =                                        0x40 //col:15059
PDE_32_IGNORED_1_MASK =                                        0x01 //col:15060
PDE_32_IGNORED_1(_) =                                          (((_) >> 6) & 0x01) //col:15061
PDE_32_LARGE_PAGE_BIT =                                        7 //col:15067
PDE_32_LARGE_PAGE_FLAG =                                       0x80 //col:15068
PDE_32_LARGE_PAGE_MASK =                                       0x01 //col:15069
PDE_32_LARGE_PAGE(_) =                                         (((_) >> 7) & 0x01) //col:15070
PDE_32_IGNORED_2_BIT =                                         8 //col:15076
PDE_32_IGNORED_2_FLAG =                                        0xF00 //col:15077
PDE_32_IGNORED_2_MASK =                                        0x0F //col:15078
PDE_32_IGNORED_2(_) =                                          (((_) >> 8) & 0x0F) //col:15079
PDE_32_PAGE_FRAME_NUMBER_BIT =                                 12 //col:15085
PDE_32_PAGE_FRAME_NUMBER_FLAG =                                0xFFFFF000 //col:15086
PDE_32_PAGE_FRAME_NUMBER_MASK =                                0xFFFFF //col:15087
PDE_32_PAGE_FRAME_NUMBER(_) =                                  (((_) >> 12) & 0xFFFFF) //col:15088
PTE_32_PRESENT_BIT =                                           0 //col:15105
PTE_32_PRESENT_FLAG =                                          0x01 //col:15106
PTE_32_PRESENT_MASK =                                          0x01 //col:15107
PTE_32_PRESENT(_) =                                            (((_) >> 0) & 0x01) //col:15108
PTE_32_WRITE_BIT =                                             1 //col:15116
PTE_32_WRITE_FLAG =                                            0x02 //col:15117
PTE_32_WRITE_MASK =                                            0x01 //col:15118
PTE_32_WRITE(_) =                                              (((_) >> 1) & 0x01) //col:15119
PTE_32_SUPERVISOR_BIT =                                        2 //col:15127
PTE_32_SUPERVISOR_FLAG =                                       0x04 //col:15128
PTE_32_SUPERVISOR_MASK =                                       0x01 //col:15129
PTE_32_SUPERVISOR(_) =                                         (((_) >> 2) & 0x01) //col:15130
PTE_32_PAGE_LEVEL_WRITE_THROUGH_BIT =                          3 //col:15139
PTE_32_PAGE_LEVEL_WRITE_THROUGH_FLAG =                         0x08 //col:15140
PTE_32_PAGE_LEVEL_WRITE_THROUGH_MASK =                         0x01 //col:15141
PTE_32_PAGE_LEVEL_WRITE_THROUGH(_) =                           (((_) >> 3) & 0x01) //col:15142
PTE_32_PAGE_LEVEL_CACHE_DISABLE_BIT =                          4 //col:15151
PTE_32_PAGE_LEVEL_CACHE_DISABLE_FLAG =                         0x10 //col:15152
PTE_32_PAGE_LEVEL_CACHE_DISABLE_MASK =                         0x01 //col:15153
PTE_32_PAGE_LEVEL_CACHE_DISABLE(_) =                           (((_) >> 4) & 0x01) //col:15154
PTE_32_ACCESSED_BIT =                                          5 //col:15162
PTE_32_ACCESSED_FLAG =                                         0x20 //col:15163
PTE_32_ACCESSED_MASK =                                         0x01 //col:15164
PTE_32_ACCESSED(_) =                                           (((_) >> 5) & 0x01) //col:15165
PTE_32_DIRTY_BIT =                                             6 //col:15173
PTE_32_DIRTY_FLAG =                                            0x40 //col:15174
PTE_32_DIRTY_MASK =                                            0x01 //col:15175
PTE_32_DIRTY(_) =                                              (((_) >> 6) & 0x01) //col:15176
PTE_32_PAT_BIT =                                               7 //col:15184
PTE_32_PAT_FLAG =                                              0x80 //col:15185
PTE_32_PAT_MASK =                                              0x01 //col:15186
PTE_32_PAT(_) =                                                (((_) >> 7) & 0x01) //col:15187
PTE_32_GLOBAL_BIT =                                            8 //col:15195
PTE_32_GLOBAL_FLAG =                                           0x100 //col:15196
PTE_32_GLOBAL_MASK =                                           0x01 //col:15197
PTE_32_GLOBAL(_) =                                             (((_) >> 8) & 0x01) //col:15198
PTE_32_IGNORED_1_BIT =                                         9 //col:15204
PTE_32_IGNORED_1_FLAG =                                        0xE00 //col:15205
PTE_32_IGNORED_1_MASK =                                        0x07 //col:15206
PTE_32_IGNORED_1(_) =                                          (((_) >> 9) & 0x07) //col:15207
PTE_32_PAGE_FRAME_NUMBER_BIT =                                 12 //col:15213
PTE_32_PAGE_FRAME_NUMBER_FLAG =                                0xFFFFF000 //col:15214
PTE_32_PAGE_FRAME_NUMBER_MASK =                                0xFFFFF //col:15215
PTE_32_PAGE_FRAME_NUMBER(_) =                                  (((_) >> 12) & 0xFFFFF) //col:15216
PT_ENTRY_32_PRESENT_BIT =                                      0 //col:15230
PT_ENTRY_32_PRESENT_FLAG =                                     0x01 //col:15231
PT_ENTRY_32_PRESENT_MASK =                                     0x01 //col:15232
PT_ENTRY_32_PRESENT(_) =                                       (((_) >> 0) & 0x01) //col:15233
PT_ENTRY_32_WRITE_BIT =                                        1 //col:15235
PT_ENTRY_32_WRITE_FLAG =                                       0x02 //col:15236
PT_ENTRY_32_WRITE_MASK =                                       0x01 //col:15237
PT_ENTRY_32_WRITE(_) =                                         (((_) >> 1) & 0x01) //col:15238
PT_ENTRY_32_SUPERVISOR_BIT =                                   2 //col:15240
PT_ENTRY_32_SUPERVISOR_FLAG =                                  0x04 //col:15241
PT_ENTRY_32_SUPERVISOR_MASK =                                  0x01 //col:15242
PT_ENTRY_32_SUPERVISOR(_) =                                    (((_) >> 2) & 0x01) //col:15243
PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_BIT =                     3 //col:15245
PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_FLAG =                    0x08 //col:15246
PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_MASK =                    0x01 //col:15247
PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH(_) =                      (((_) >> 3) & 0x01) //col:15248
PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_BIT =                     4 //col:15250
PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_FLAG =                    0x10 //col:15251
PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_MASK =                    0x01 //col:15252
PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE(_) =                      (((_) >> 4) & 0x01) //col:15253
PT_ENTRY_32_ACCESSED_BIT =                                     5 //col:15255
PT_ENTRY_32_ACCESSED_FLAG =                                    0x20 //col:15256
PT_ENTRY_32_ACCESSED_MASK =                                    0x01 //col:15257
PT_ENTRY_32_ACCESSED(_) =                                      (((_) >> 5) & 0x01) //col:15258
PT_ENTRY_32_DIRTY_BIT =                                        6 //col:15260
PT_ENTRY_32_DIRTY_FLAG =                                       0x40 //col:15261
PT_ENTRY_32_DIRTY_MASK =                                       0x01 //col:15262
PT_ENTRY_32_DIRTY(_) =                                         (((_) >> 6) & 0x01) //col:15263
PT_ENTRY_32_LARGE_PAGE_BIT =                                   7 //col:15265
PT_ENTRY_32_LARGE_PAGE_FLAG =                                  0x80 //col:15266
PT_ENTRY_32_LARGE_PAGE_MASK =                                  0x01 //col:15267
PT_ENTRY_32_LARGE_PAGE(_) =                                    (((_) >> 7) & 0x01) //col:15268
PT_ENTRY_32_GLOBAL_BIT =                                       8 //col:15270
PT_ENTRY_32_GLOBAL_FLAG =                                      0x100 //col:15271
PT_ENTRY_32_GLOBAL_MASK =                                      0x01 //col:15272
PT_ENTRY_32_GLOBAL(_) =                                        (((_) >> 8) & 0x01) //col:15273
PT_ENTRY_32_IGNORED_1_BIT =                                    9 //col:15279
PT_ENTRY_32_IGNORED_1_FLAG =                                   0xE00 //col:15280
PT_ENTRY_32_IGNORED_1_MASK =                                   0x07 //col:15281
PT_ENTRY_32_IGNORED_1(_) =                                     (((_) >> 9) & 0x07) //col:15282
PT_ENTRY_32_PAGE_FRAME_NUMBER_BIT =                            12 //col:15288
PT_ENTRY_32_PAGE_FRAME_NUMBER_FLAG =                           0xFFFFF000 //col:15289
PT_ENTRY_32_PAGE_FRAME_NUMBER_MASK =                           0xFFFFF //col:15290
PT_ENTRY_32_PAGE_FRAME_NUMBER(_) =                             (((_) >> 12) & 0xFFFFF) //col:15291
PDE_ENTRY_COUNT_32 =                                           0x00000400 //col:15304
PTE_ENTRY_COUNT_32 =                                           0x00000400 //col:15305
PML4E_64_PRESENT_BIT =                                         0 //col:15340
PML4E_64_PRESENT_FLAG =                                        0x01 //col:15341
PML4E_64_PRESENT_MASK =                                        0x01 //col:15342
PML4E_64_PRESENT(_) =                                          (((_) >> 0) & 0x01) //col:15343
PML4E_64_WRITE_BIT =                                           1 //col:15351
PML4E_64_WRITE_FLAG =                                          0x02 //col:15352
PML4E_64_WRITE_MASK =                                          0x01 //col:15353
PML4E_64_WRITE(_) =                                            (((_) >> 1) & 0x01) //col:15354
PML4E_64_SUPERVISOR_BIT =                                      2 //col:15362
PML4E_64_SUPERVISOR_FLAG =                                     0x04 //col:15363
PML4E_64_SUPERVISOR_MASK =                                     0x01 //col:15364
PML4E_64_SUPERVISOR(_) =                                       (((_) >> 2) & 0x01) //col:15365
PML4E_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                        3 //col:15374
PML4E_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                       0x08 //col:15375
PML4E_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                       0x01 //col:15376
PML4E_64_PAGE_LEVEL_WRITE_THROUGH(_) =                         (((_) >> 3) & 0x01) //col:15377
PML4E_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                        4 //col:15386
PML4E_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                       0x10 //col:15387
PML4E_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                       0x01 //col:15388
PML4E_64_PAGE_LEVEL_CACHE_DISABLE(_) =                         (((_) >> 4) & 0x01) //col:15389
PML4E_64_ACCESSED_BIT =                                        5 //col:15397
PML4E_64_ACCESSED_FLAG =                                       0x20 //col:15398
PML4E_64_ACCESSED_MASK =                                       0x01 //col:15399
PML4E_64_ACCESSED(_) =                                         (((_) >> 5) & 0x01) //col:15400
PML4E_64_MUST_BE_ZERO_BIT =                                    7 //col:15407
PML4E_64_MUST_BE_ZERO_FLAG =                                   0x80 //col:15408
PML4E_64_MUST_BE_ZERO_MASK =                                   0x01 //col:15409
PML4E_64_MUST_BE_ZERO(_) =                                     (((_) >> 7) & 0x01) //col:15410
PML4E_64_IGNORED_1_BIT =                                       8 //col:15416
PML4E_64_IGNORED_1_FLAG =                                      0x700 //col:15417
PML4E_64_IGNORED_1_MASK =                                      0x07 //col:15418
PML4E_64_IGNORED_1(_) =                                        (((_) >> 8) & 0x07) //col:15419
PML4E_64_RESTART_BIT =                                         11 //col:15428
PML4E_64_RESTART_FLAG =                                        0x800 //col:15429
PML4E_64_RESTART_MASK =                                        0x01 //col:15430
PML4E_64_RESTART(_) =                                          (((_) >> 11) & 0x01) //col:15431
PML4E_64_PAGE_FRAME_NUMBER_BIT =                               12 //col:15437
PML4E_64_PAGE_FRAME_NUMBER_FLAG =                              0xFFFFFFFFF000 //col:15438
PML4E_64_PAGE_FRAME_NUMBER_MASK =                              0xFFFFFFFFF //col:15439
PML4E_64_PAGE_FRAME_NUMBER(_) =                                (((_) >> 12) & 0xFFFFFFFFF) //col:15440
PML4E_64_IGNORED_2_BIT =                                       52 //col:15447
PML4E_64_IGNORED_2_FLAG =                                      0x7FF0000000000000 //col:15448
PML4E_64_IGNORED_2_MASK =                                      0x7FF //col:15449
PML4E_64_IGNORED_2(_) =                                        (((_) >> 52) & 0x7FF) //col:15450
PML4E_64_EXECUTE_DISABLE_BIT =                                 63 //col:15459
PML4E_64_EXECUTE_DISABLE_FLAG =                                0x8000000000000000 //col:15460
PML4E_64_EXECUTE_DISABLE_MASK =                                0x01 //col:15461
PML4E_64_EXECUTE_DISABLE(_) =                                  (((_) >> 63) & 0x01) //col:15462
PDPTE_1GB_64_PRESENT_BIT =                                     0 //col:15479
PDPTE_1GB_64_PRESENT_FLAG =                                    0x01 //col:15480
PDPTE_1GB_64_PRESENT_MASK =                                    0x01 //col:15481
PDPTE_1GB_64_PRESENT(_) =                                      (((_) >> 0) & 0x01) //col:15482
PDPTE_1GB_64_WRITE_BIT =                                       1 //col:15490
PDPTE_1GB_64_WRITE_FLAG =                                      0x02 //col:15491
PDPTE_1GB_64_WRITE_MASK =                                      0x01 //col:15492
PDPTE_1GB_64_WRITE(_) =                                        (((_) >> 1) & 0x01) //col:15493
PDPTE_1GB_64_SUPERVISOR_BIT =                                  2 //col:15501
PDPTE_1GB_64_SUPERVISOR_FLAG =                                 0x04 //col:15502
PDPTE_1GB_64_SUPERVISOR_MASK =                                 0x01 //col:15503
PDPTE_1GB_64_SUPERVISOR(_) =                                   (((_) >> 2) & 0x01) //col:15504
PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                    3 //col:15513
PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                   0x08 //col:15514
PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                   0x01 //col:15515
PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH(_) =                     (((_) >> 3) & 0x01) //col:15516
PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                    4 //col:15525
PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                   0x10 //col:15526
PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                   0x01 //col:15527
PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE(_) =                     (((_) >> 4) & 0x01) //col:15528
PDPTE_1GB_64_ACCESSED_BIT =                                    5 //col:15536
PDPTE_1GB_64_ACCESSED_FLAG =                                   0x20 //col:15537
PDPTE_1GB_64_ACCESSED_MASK =                                   0x01 //col:15538
PDPTE_1GB_64_ACCESSED(_) =                                     (((_) >> 5) & 0x01) //col:15539
PDPTE_1GB_64_DIRTY_BIT =                                       6 //col:15547
PDPTE_1GB_64_DIRTY_FLAG =                                      0x40 //col:15548
PDPTE_1GB_64_DIRTY_MASK =                                      0x01 //col:15549
PDPTE_1GB_64_DIRTY(_) =                                        (((_) >> 6) & 0x01) //col:15550
PDPTE_1GB_64_LARGE_PAGE_BIT =                                  7 //col:15556
PDPTE_1GB_64_LARGE_PAGE_FLAG =                                 0x80 //col:15557
PDPTE_1GB_64_LARGE_PAGE_MASK =                                 0x01 //col:15558
PDPTE_1GB_64_LARGE_PAGE(_) =                                   (((_) >> 7) & 0x01) //col:15559
PDPTE_1GB_64_GLOBAL_BIT =                                      8 //col:15567
PDPTE_1GB_64_GLOBAL_FLAG =                                     0x100 //col:15568
PDPTE_1GB_64_GLOBAL_MASK =                                     0x01 //col:15569
PDPTE_1GB_64_GLOBAL(_) =                                       (((_) >> 8) & 0x01) //col:15570
PDPTE_1GB_64_IGNORED_1_BIT =                                   9 //col:15576
PDPTE_1GB_64_IGNORED_1_FLAG =                                  0x600 //col:15577
PDPTE_1GB_64_IGNORED_1_MASK =                                  0x03 //col:15578
PDPTE_1GB_64_IGNORED_1(_) =                                    (((_) >> 9) & 0x03) //col:15579
PDPTE_1GB_64_RESTART_BIT =                                     11 //col:15588
PDPTE_1GB_64_RESTART_FLAG =                                    0x800 //col:15589
PDPTE_1GB_64_RESTART_MASK =                                    0x01 //col:15590
PDPTE_1GB_64_RESTART(_) =                                      (((_) >> 11) & 0x01) //col:15591
PDPTE_1GB_64_PAT_BIT =                                         12 //col:15600
PDPTE_1GB_64_PAT_FLAG =                                        0x1000 //col:15601
PDPTE_1GB_64_PAT_MASK =                                        0x01 //col:15602
PDPTE_1GB_64_PAT(_) =                                          (((_) >> 12) & 0x01) //col:15603
PDPTE_1GB_64_PAGE_FRAME_NUMBER_BIT =                           30 //col:15610
PDPTE_1GB_64_PAGE_FRAME_NUMBER_FLAG =                          0xFFFFC0000000 //col:15611
PDPTE_1GB_64_PAGE_FRAME_NUMBER_MASK =                          0x3FFFF //col:15612
PDPTE_1GB_64_PAGE_FRAME_NUMBER(_) =                            (((_) >> 30) & 0x3FFFF) //col:15613
PDPTE_1GB_64_IGNORED_2_BIT =                                   52 //col:15620
PDPTE_1GB_64_IGNORED_2_FLAG =                                  0x7F0000000000000 //col:15621
PDPTE_1GB_64_IGNORED_2_MASK =                                  0x7F //col:15622
PDPTE_1GB_64_IGNORED_2(_) =                                    (((_) >> 52) & 0x7F) //col:15623
PDPTE_1GB_64_PROTECTION_KEY_BIT =                              59 //col:15631
PDPTE_1GB_64_PROTECTION_KEY_FLAG =                             0x7800000000000000 //col:15632
PDPTE_1GB_64_PROTECTION_KEY_MASK =                             0x0F //col:15633
PDPTE_1GB_64_PROTECTION_KEY(_) =                               (((_) >> 59) & 0x0F) //col:15634
PDPTE_1GB_64_EXECUTE_DISABLE_BIT =                             63 //col:15643
PDPTE_1GB_64_EXECUTE_DISABLE_FLAG =                            0x8000000000000000 //col:15644
PDPTE_1GB_64_EXECUTE_DISABLE_MASK =                            0x01 //col:15645
PDPTE_1GB_64_EXECUTE_DISABLE(_) =                              (((_) >> 63) & 0x01) //col:15646
PDPTE_64_PRESENT_BIT =                                         0 //col:15663
PDPTE_64_PRESENT_FLAG =                                        0x01 //col:15664
PDPTE_64_PRESENT_MASK =                                        0x01 //col:15665
PDPTE_64_PRESENT(_) =                                          (((_) >> 0) & 0x01) //col:15666
PDPTE_64_WRITE_BIT =                                           1 //col:15674
PDPTE_64_WRITE_FLAG =                                          0x02 //col:15675
PDPTE_64_WRITE_MASK =                                          0x01 //col:15676
PDPTE_64_WRITE(_) =                                            (((_) >> 1) & 0x01) //col:15677
PDPTE_64_SUPERVISOR_BIT =                                      2 //col:15685
PDPTE_64_SUPERVISOR_FLAG =                                     0x04 //col:15686
PDPTE_64_SUPERVISOR_MASK =                                     0x01 //col:15687
PDPTE_64_SUPERVISOR(_) =                                       (((_) >> 2) & 0x01) //col:15688
PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                        3 //col:15697
PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                       0x08 //col:15698
PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                       0x01 //col:15699
PDPTE_64_PAGE_LEVEL_WRITE_THROUGH(_) =                         (((_) >> 3) & 0x01) //col:15700
PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                        4 //col:15709
PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                       0x10 //col:15710
PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                       0x01 //col:15711
PDPTE_64_PAGE_LEVEL_CACHE_DISABLE(_) =                         (((_) >> 4) & 0x01) //col:15712
PDPTE_64_ACCESSED_BIT =                                        5 //col:15720
PDPTE_64_ACCESSED_FLAG =                                       0x20 //col:15721
PDPTE_64_ACCESSED_MASK =                                       0x01 //col:15722
PDPTE_64_ACCESSED(_) =                                         (((_) >> 5) & 0x01) //col:15723
PDPTE_64_LARGE_PAGE_BIT =                                      7 //col:15730
PDPTE_64_LARGE_PAGE_FLAG =                                     0x80 //col:15731
PDPTE_64_LARGE_PAGE_MASK =                                     0x01 //col:15732
PDPTE_64_LARGE_PAGE(_) =                                       (((_) >> 7) & 0x01) //col:15733
PDPTE_64_IGNORED_1_BIT =                                       8 //col:15739
PDPTE_64_IGNORED_1_FLAG =                                      0x700 //col:15740
PDPTE_64_IGNORED_1_MASK =                                      0x07 //col:15741
PDPTE_64_IGNORED_1(_) =                                        (((_) >> 8) & 0x07) //col:15742
PDPTE_64_RESTART_BIT =                                         11 //col:15751
PDPTE_64_RESTART_FLAG =                                        0x800 //col:15752
PDPTE_64_RESTART_MASK =                                        0x01 //col:15753
PDPTE_64_RESTART(_) =                                          (((_) >> 11) & 0x01) //col:15754
PDPTE_64_PAGE_FRAME_NUMBER_BIT =                               12 //col:15760
PDPTE_64_PAGE_FRAME_NUMBER_FLAG =                              0xFFFFFFFFF000 //col:15761
PDPTE_64_PAGE_FRAME_NUMBER_MASK =                              0xFFFFFFFFF //col:15762
PDPTE_64_PAGE_FRAME_NUMBER(_) =                                (((_) >> 12) & 0xFFFFFFFFF) //col:15763
PDPTE_64_IGNORED_2_BIT =                                       52 //col:15770
PDPTE_64_IGNORED_2_FLAG =                                      0x7FF0000000000000 //col:15771
PDPTE_64_IGNORED_2_MASK =                                      0x7FF //col:15772
PDPTE_64_IGNORED_2(_) =                                        (((_) >> 52) & 0x7FF) //col:15773
PDPTE_64_EXECUTE_DISABLE_BIT =                                 63 //col:15782
PDPTE_64_EXECUTE_DISABLE_FLAG =                                0x8000000000000000 //col:15783
PDPTE_64_EXECUTE_DISABLE_MASK =                                0x01 //col:15784
PDPTE_64_EXECUTE_DISABLE(_) =                                  (((_) >> 63) & 0x01) //col:15785
PDE_2MB_64_PRESENT_BIT =                                       0 //col:15802
PDE_2MB_64_PRESENT_FLAG =                                      0x01 //col:15803
PDE_2MB_64_PRESENT_MASK =                                      0x01 //col:15804
PDE_2MB_64_PRESENT(_) =                                        (((_) >> 0) & 0x01) //col:15805
PDE_2MB_64_WRITE_BIT =                                         1 //col:15813
PDE_2MB_64_WRITE_FLAG =                                        0x02 //col:15814
PDE_2MB_64_WRITE_MASK =                                        0x01 //col:15815
PDE_2MB_64_WRITE(_) =                                          (((_) >> 1) & 0x01) //col:15816
PDE_2MB_64_SUPERVISOR_BIT =                                    2 //col:15824
PDE_2MB_64_SUPERVISOR_FLAG =                                   0x04 //col:15825
PDE_2MB_64_SUPERVISOR_MASK =                                   0x01 //col:15826
PDE_2MB_64_SUPERVISOR(_) =                                     (((_) >> 2) & 0x01) //col:15827
PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                      3 //col:15836
PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                     0x08 //col:15837
PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                     0x01 //col:15838
PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH(_) =                       (((_) >> 3) & 0x01) //col:15839
PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                      4 //col:15848
PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                     0x10 //col:15849
PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                     0x01 //col:15850
PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE(_) =                       (((_) >> 4) & 0x01) //col:15851
PDE_2MB_64_ACCESSED_BIT =                                      5 //col:15859
PDE_2MB_64_ACCESSED_FLAG =                                     0x20 //col:15860
PDE_2MB_64_ACCESSED_MASK =                                     0x01 //col:15861
PDE_2MB_64_ACCESSED(_) =                                       (((_) >> 5) & 0x01) //col:15862
PDE_2MB_64_DIRTY_BIT =                                         6 //col:15870
PDE_2MB_64_DIRTY_FLAG =                                        0x40 //col:15871
PDE_2MB_64_DIRTY_MASK =                                        0x01 //col:15872
PDE_2MB_64_DIRTY(_) =                                          (((_) >> 6) & 0x01) //col:15873
PDE_2MB_64_LARGE_PAGE_BIT =                                    7 //col:15879
PDE_2MB_64_LARGE_PAGE_FLAG =                                   0x80 //col:15880
PDE_2MB_64_LARGE_PAGE_MASK =                                   0x01 //col:15881
PDE_2MB_64_LARGE_PAGE(_) =                                     (((_) >> 7) & 0x01) //col:15882
PDE_2MB_64_GLOBAL_BIT =                                        8 //col:15890
PDE_2MB_64_GLOBAL_FLAG =                                       0x100 //col:15891
PDE_2MB_64_GLOBAL_MASK =                                       0x01 //col:15892
PDE_2MB_64_GLOBAL(_) =                                         (((_) >> 8) & 0x01) //col:15893
PDE_2MB_64_IGNORED_1_BIT =                                     9 //col:15899
PDE_2MB_64_IGNORED_1_FLAG =                                    0x600 //col:15900
PDE_2MB_64_IGNORED_1_MASK =                                    0x03 //col:15901
PDE_2MB_64_IGNORED_1(_) =                                      (((_) >> 9) & 0x03) //col:15902
PDE_2MB_64_RESTART_BIT =                                       11 //col:15911
PDE_2MB_64_RESTART_FLAG =                                      0x800 //col:15912
PDE_2MB_64_RESTART_MASK =                                      0x01 //col:15913
PDE_2MB_64_RESTART(_) =                                        (((_) >> 11) & 0x01) //col:15914
PDE_2MB_64_PAT_BIT =                                           12 //col:15923
PDE_2MB_64_PAT_FLAG =                                          0x1000 //col:15924
PDE_2MB_64_PAT_MASK =                                          0x01 //col:15925
PDE_2MB_64_PAT(_) =                                            (((_) >> 12) & 0x01) //col:15926
PDE_2MB_64_PAGE_FRAME_NUMBER_BIT =                             21 //col:15933
PDE_2MB_64_PAGE_FRAME_NUMBER_FLAG =                            0xFFFFFFE00000 //col:15934
PDE_2MB_64_PAGE_FRAME_NUMBER_MASK =                            0x7FFFFFF //col:15935
PDE_2MB_64_PAGE_FRAME_NUMBER(_) =                              (((_) >> 21) & 0x7FFFFFF) //col:15936
PDE_2MB_64_IGNORED_2_BIT =                                     52 //col:15943
PDE_2MB_64_IGNORED_2_FLAG =                                    0x7F0000000000000 //col:15944
PDE_2MB_64_IGNORED_2_MASK =                                    0x7F //col:15945
PDE_2MB_64_IGNORED_2(_) =                                      (((_) >> 52) & 0x7F) //col:15946
PDE_2MB_64_PROTECTION_KEY_BIT =                                59 //col:15954
PDE_2MB_64_PROTECTION_KEY_FLAG =                               0x7800000000000000 //col:15955
PDE_2MB_64_PROTECTION_KEY_MASK =                               0x0F //col:15956
PDE_2MB_64_PROTECTION_KEY(_) =                                 (((_) >> 59) & 0x0F) //col:15957
PDE_2MB_64_EXECUTE_DISABLE_BIT =                               63 //col:15966
PDE_2MB_64_EXECUTE_DISABLE_FLAG =                              0x8000000000000000 //col:15967
PDE_2MB_64_EXECUTE_DISABLE_MASK =                              0x01 //col:15968
PDE_2MB_64_EXECUTE_DISABLE(_) =                                (((_) >> 63) & 0x01) //col:15969
PDE_64_PRESENT_BIT =                                           0 //col:15986
PDE_64_PRESENT_FLAG =                                          0x01 //col:15987
PDE_64_PRESENT_MASK =                                          0x01 //col:15988
PDE_64_PRESENT(_) =                                            (((_) >> 0) & 0x01) //col:15989
PDE_64_WRITE_BIT =                                             1 //col:15997
PDE_64_WRITE_FLAG =                                            0x02 //col:15998
PDE_64_WRITE_MASK =                                            0x01 //col:15999
PDE_64_WRITE(_) =                                              (((_) >> 1) & 0x01) //col:16000
PDE_64_SUPERVISOR_BIT =                                        2 //col:16008
PDE_64_SUPERVISOR_FLAG =                                       0x04 //col:16009
PDE_64_SUPERVISOR_MASK =                                       0x01 //col:16010
PDE_64_SUPERVISOR(_) =                                         (((_) >> 2) & 0x01) //col:16011
PDE_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                          3 //col:16020
PDE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                         0x08 //col:16021
PDE_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                         0x01 //col:16022
PDE_64_PAGE_LEVEL_WRITE_THROUGH(_) =                           (((_) >> 3) & 0x01) //col:16023
PDE_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                          4 //col:16032
PDE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                         0x10 //col:16033
PDE_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                         0x01 //col:16034
PDE_64_PAGE_LEVEL_CACHE_DISABLE(_) =                           (((_) >> 4) & 0x01) //col:16035
PDE_64_ACCESSED_BIT =                                          5 //col:16043
PDE_64_ACCESSED_FLAG =                                         0x20 //col:16044
PDE_64_ACCESSED_MASK =                                         0x01 //col:16045
PDE_64_ACCESSED(_) =                                           (((_) >> 5) & 0x01) //col:16046
PDE_64_LARGE_PAGE_BIT =                                        7 //col:16053
PDE_64_LARGE_PAGE_FLAG =                                       0x80 //col:16054
PDE_64_LARGE_PAGE_MASK =                                       0x01 //col:16055
PDE_64_LARGE_PAGE(_) =                                         (((_) >> 7) & 0x01) //col:16056
PDE_64_IGNORED_1_BIT =                                         8 //col:16062
PDE_64_IGNORED_1_FLAG =                                        0x700 //col:16063
PDE_64_IGNORED_1_MASK =                                        0x07 //col:16064
PDE_64_IGNORED_1(_) =                                          (((_) >> 8) & 0x07) //col:16065
PDE_64_RESTART_BIT =                                           11 //col:16074
PDE_64_RESTART_FLAG =                                          0x800 //col:16075
PDE_64_RESTART_MASK =                                          0x01 //col:16076
PDE_64_RESTART(_) =                                            (((_) >> 11) & 0x01) //col:16077
PDE_64_PAGE_FRAME_NUMBER_BIT =                                 12 //col:16083
PDE_64_PAGE_FRAME_NUMBER_FLAG =                                0xFFFFFFFFF000 //col:16084
PDE_64_PAGE_FRAME_NUMBER_MASK =                                0xFFFFFFFFF //col:16085
PDE_64_PAGE_FRAME_NUMBER(_) =                                  (((_) >> 12) & 0xFFFFFFFFF) //col:16086
PDE_64_IGNORED_2_BIT =                                         52 //col:16093
PDE_64_IGNORED_2_FLAG =                                        0x7FF0000000000000 //col:16094
PDE_64_IGNORED_2_MASK =                                        0x7FF //col:16095
PDE_64_IGNORED_2(_) =                                          (((_) >> 52) & 0x7FF) //col:16096
PDE_64_EXECUTE_DISABLE_BIT =                                   63 //col:16105
PDE_64_EXECUTE_DISABLE_FLAG =                                  0x8000000000000000 //col:16106
PDE_64_EXECUTE_DISABLE_MASK =                                  0x01 //col:16107
PDE_64_EXECUTE_DISABLE(_) =                                    (((_) >> 63) & 0x01) //col:16108
PTE_64_PRESENT_BIT =                                           0 //col:16125
PTE_64_PRESENT_FLAG =                                          0x01 //col:16126
PTE_64_PRESENT_MASK =                                          0x01 //col:16127
PTE_64_PRESENT(_) =                                            (((_) >> 0) & 0x01) //col:16128
PTE_64_WRITE_BIT =                                             1 //col:16136
PTE_64_WRITE_FLAG =                                            0x02 //col:16137
PTE_64_WRITE_MASK =                                            0x01 //col:16138
PTE_64_WRITE(_) =                                              (((_) >> 1) & 0x01) //col:16139
PTE_64_SUPERVISOR_BIT =                                        2 //col:16147
PTE_64_SUPERVISOR_FLAG =                                       0x04 //col:16148
PTE_64_SUPERVISOR_MASK =                                       0x01 //col:16149
PTE_64_SUPERVISOR(_) =                                         (((_) >> 2) & 0x01) //col:16150
PTE_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                          3 //col:16159
PTE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                         0x08 //col:16160
PTE_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                         0x01 //col:16161
PTE_64_PAGE_LEVEL_WRITE_THROUGH(_) =                           (((_) >> 3) & 0x01) //col:16162
PTE_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                          4 //col:16171
PTE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                         0x10 //col:16172
PTE_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                         0x01 //col:16173
PTE_64_PAGE_LEVEL_CACHE_DISABLE(_) =                           (((_) >> 4) & 0x01) //col:16174
PTE_64_ACCESSED_BIT =                                          5 //col:16182
PTE_64_ACCESSED_FLAG =                                         0x20 //col:16183
PTE_64_ACCESSED_MASK =                                         0x01 //col:16184
PTE_64_ACCESSED(_) =                                           (((_) >> 5) & 0x01) //col:16185
PTE_64_DIRTY_BIT =                                             6 //col:16193
PTE_64_DIRTY_FLAG =                                            0x40 //col:16194
PTE_64_DIRTY_MASK =                                            0x01 //col:16195
PTE_64_DIRTY(_) =                                              (((_) >> 6) & 0x01) //col:16196
PTE_64_PAT_BIT =                                               7 //col:16204
PTE_64_PAT_FLAG =                                              0x80 //col:16205
PTE_64_PAT_MASK =                                              0x01 //col:16206
PTE_64_PAT(_) =                                                (((_) >> 7) & 0x01) //col:16207
PTE_64_GLOBAL_BIT =                                            8 //col:16215
PTE_64_GLOBAL_FLAG =                                           0x100 //col:16216
PTE_64_GLOBAL_MASK =                                           0x01 //col:16217
PTE_64_GLOBAL(_) =                                             (((_) >> 8) & 0x01) //col:16218
PTE_64_IGNORED_1_BIT =                                         9 //col:16224
PTE_64_IGNORED_1_FLAG =                                        0x600 //col:16225
PTE_64_IGNORED_1_MASK =                                        0x03 //col:16226
PTE_64_IGNORED_1(_) =                                          (((_) >> 9) & 0x03) //col:16227
PTE_64_RESTART_BIT =                                           11 //col:16236
PTE_64_RESTART_FLAG =                                          0x800 //col:16237
PTE_64_RESTART_MASK =                                          0x01 //col:16238
PTE_64_RESTART(_) =                                            (((_) >> 11) & 0x01) //col:16239
PTE_64_PAGE_FRAME_NUMBER_BIT =                                 12 //col:16245
PTE_64_PAGE_FRAME_NUMBER_FLAG =                                0xFFFFFFFFF000 //col:16246
PTE_64_PAGE_FRAME_NUMBER_MASK =                                0xFFFFFFFFF //col:16247
PTE_64_PAGE_FRAME_NUMBER(_) =                                  (((_) >> 12) & 0xFFFFFFFFF) //col:16248
PTE_64_IGNORED_2_BIT =                                         52 //col:16255
PTE_64_IGNORED_2_FLAG =                                        0x7F0000000000000 //col:16256
PTE_64_IGNORED_2_MASK =                                        0x7F //col:16257
PTE_64_IGNORED_2(_) =                                          (((_) >> 52) & 0x7F) //col:16258
PTE_64_PROTECTION_KEY_BIT =                                    59 //col:16266
PTE_64_PROTECTION_KEY_FLAG =                                   0x7800000000000000 //col:16267
PTE_64_PROTECTION_KEY_MASK =                                   0x0F //col:16268
PTE_64_PROTECTION_KEY(_) =                                     (((_) >> 59) & 0x0F) //col:16269
PTE_64_EXECUTE_DISABLE_BIT =                                   63 //col:16278
PTE_64_EXECUTE_DISABLE_FLAG =                                  0x8000000000000000 //col:16279
PTE_64_EXECUTE_DISABLE_MASK =                                  0x01 //col:16280
PTE_64_EXECUTE_DISABLE(_) =                                    (((_) >> 63) & 0x01) //col:16281
PT_ENTRY_64_PRESENT_BIT =                                      0 //col:16295
PT_ENTRY_64_PRESENT_FLAG =                                     0x01 //col:16296
PT_ENTRY_64_PRESENT_MASK =                                     0x01 //col:16297
PT_ENTRY_64_PRESENT(_) =                                       (((_) >> 0) & 0x01) //col:16298
PT_ENTRY_64_WRITE_BIT =                                        1 //col:16300
PT_ENTRY_64_WRITE_FLAG =                                       0x02 //col:16301
PT_ENTRY_64_WRITE_MASK =                                       0x01 //col:16302
PT_ENTRY_64_WRITE(_) =                                         (((_) >> 1) & 0x01) //col:16303
PT_ENTRY_64_SUPERVISOR_BIT =                                   2 //col:16305
PT_ENTRY_64_SUPERVISOR_FLAG =                                  0x04 //col:16306
PT_ENTRY_64_SUPERVISOR_MASK =                                  0x01 //col:16307
PT_ENTRY_64_SUPERVISOR(_) =                                    (((_) >> 2) & 0x01) //col:16308
PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_BIT =                     3 //col:16310
PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_FLAG =                    0x08 //col:16311
PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_MASK =                    0x01 //col:16312
PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH(_) =                      (((_) >> 3) & 0x01) //col:16313
PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_BIT =                     4 //col:16315
PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_FLAG =                    0x10 //col:16316
PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_MASK =                    0x01 //col:16317
PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE(_) =                      (((_) >> 4) & 0x01) //col:16318
PT_ENTRY_64_ACCESSED_BIT =                                     5 //col:16320
PT_ENTRY_64_ACCESSED_FLAG =                                    0x20 //col:16321
PT_ENTRY_64_ACCESSED_MASK =                                    0x01 //col:16322
PT_ENTRY_64_ACCESSED(_) =                                      (((_) >> 5) & 0x01) //col:16323
PT_ENTRY_64_DIRTY_BIT =                                        6 //col:16325
PT_ENTRY_64_DIRTY_FLAG =                                       0x40 //col:16326
PT_ENTRY_64_DIRTY_MASK =                                       0x01 //col:16327
PT_ENTRY_64_DIRTY(_) =                                         (((_) >> 6) & 0x01) //col:16328
PT_ENTRY_64_LARGE_PAGE_BIT =                                   7 //col:16330
PT_ENTRY_64_LARGE_PAGE_FLAG =                                  0x80 //col:16331
PT_ENTRY_64_LARGE_PAGE_MASK =                                  0x01 //col:16332
PT_ENTRY_64_LARGE_PAGE(_) =                                    (((_) >> 7) & 0x01) //col:16333
PT_ENTRY_64_GLOBAL_BIT =                                       8 //col:16335
PT_ENTRY_64_GLOBAL_FLAG =                                      0x100 //col:16336
PT_ENTRY_64_GLOBAL_MASK =                                      0x01 //col:16337
PT_ENTRY_64_GLOBAL(_) =                                        (((_) >> 8) & 0x01) //col:16338
PT_ENTRY_64_IGNORED_1_BIT =                                    9 //col:16344
PT_ENTRY_64_IGNORED_1_FLAG =                                   0x600 //col:16345
PT_ENTRY_64_IGNORED_1_MASK =                                   0x03 //col:16346
PT_ENTRY_64_IGNORED_1(_) =                                     (((_) >> 9) & 0x03) //col:16347
PT_ENTRY_64_RESTART_BIT =                                      11 //col:16349
PT_ENTRY_64_RESTART_FLAG =                                     0x800 //col:16350
PT_ENTRY_64_RESTART_MASK =                                     0x01 //col:16351
PT_ENTRY_64_RESTART(_) =                                       (((_) >> 11) & 0x01) //col:16352
PT_ENTRY_64_PAGE_FRAME_NUMBER_BIT =                            12 //col:16358
PT_ENTRY_64_PAGE_FRAME_NUMBER_FLAG =                           0xFFFFFFFFF000 //col:16359
PT_ENTRY_64_PAGE_FRAME_NUMBER_MASK =                           0xFFFFFFFFF //col:16360
PT_ENTRY_64_PAGE_FRAME_NUMBER(_) =                             (((_) >> 12) & 0xFFFFFFFFF) //col:16361
PT_ENTRY_64_IGNORED_2_BIT =                                    52 //col:16368
PT_ENTRY_64_IGNORED_2_FLAG =                                   0x7F0000000000000 //col:16369
PT_ENTRY_64_IGNORED_2_MASK =                                   0x7F //col:16370
PT_ENTRY_64_IGNORED_2(_) =                                     (((_) >> 52) & 0x7F) //col:16371
PT_ENTRY_64_PROTECTION_KEY_BIT =                               59 //col:16373
PT_ENTRY_64_PROTECTION_KEY_FLAG =                              0x7800000000000000 //col:16374
PT_ENTRY_64_PROTECTION_KEY_MASK =                              0x0F //col:16375
PT_ENTRY_64_PROTECTION_KEY(_) =                                (((_) >> 59) & 0x0F) //col:16376
PT_ENTRY_64_EXECUTE_DISABLE_BIT =                              63 //col:16378
PT_ENTRY_64_EXECUTE_DISABLE_FLAG =                             0x8000000000000000 //col:16379
PT_ENTRY_64_EXECUTE_DISABLE_MASK =                             0x01 //col:16380
PT_ENTRY_64_EXECUTE_DISABLE(_) =                               (((_) >> 63) & 0x01) //col:16381
PML4E_ENTRY_COUNT_64 =                                         0x00000200 //col:16394
PDPTE_ENTRY_COUNT_64 =                                         0x00000200 //col:16395
PDE_ENTRY_COUNT_64 =                                           0x00000200 //col:16396
PTE_ENTRY_COUNT_64 =                                           0x00000200 //col:16397
INVPCID_DESCRIPTOR_PCID_BIT =                                  0 //col:16445
INVPCID_DESCRIPTOR_PCID_FLAG =                                 0xFFF //col:16446
INVPCID_DESCRIPTOR_PCID_MASK =                                 0xFFF //col:16447
INVPCID_DESCRIPTOR_PCID(_) =                                   (((_) >> 0) & 0xFFF) //col:16448
INVPCID_DESCRIPTOR_RESERVED1_BIT =                             12 //col:16454
INVPCID_DESCRIPTOR_RESERVED1_FLAG =                            0xFFFFFFFFFFFFF000 //col:16455
INVPCID_DESCRIPTOR_RESERVED1_MASK =                            0xFFFFFFFFFFFFF //col:16456
INVPCID_DESCRIPTOR_RESERVED1(_) =                              (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:16457
INVPCID_DESCRIPTOR_LINEAR_ADDRESS_BIT =                        64 //col:16459
INVPCID_DESCRIPTOR_LINEAR_ADDRESS_FLAG =                       0xFFFFFFFFFFFFFFFF0000000000000000 //col:16460
INVPCID_DESCRIPTOR_LINEAR_ADDRESS_MASK =                       0xFFFFFFFFFFFFFFFF //col:16461
INVPCID_DESCRIPTOR_LINEAR_ADDRESS(_) =                         (((_) >> 64) & 0xFFFFFFFFFFFFFFFF) //col:16462
SEGMENT_ACCESS_RIGHTS_TYPE_BIT =                               8 //col:16535
SEGMENT_ACCESS_RIGHTS_TYPE_FLAG =                              0xF00 //col:16536
SEGMENT_ACCESS_RIGHTS_TYPE_MASK =                              0x0F //col:16537
SEGMENT_ACCESS_RIGHTS_TYPE(_) =                                (((_) >> 8) & 0x0F) //col:16538
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_BIT =                    12 //col:16547
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_FLAG =                   0x1000 //col:16548
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_MASK =                   0x01 //col:16549
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE(_) =                     (((_) >> 12) & 0x01) //col:16550
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_BIT =         13 //col:16560
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_FLAG =        0x6000 //col:16561
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_MASK =        0x03 //col:16562
SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL(_) =          (((_) >> 13) & 0x03) //col:16563
SEGMENT_ACCESS_RIGHTS_PRESENT_BIT =                            15 //col:16574
SEGMENT_ACCESS_RIGHTS_PRESENT_FLAG =                           0x8000 //col:16575
SEGMENT_ACCESS_RIGHTS_PRESENT_MASK =                           0x01 //col:16576
SEGMENT_ACCESS_RIGHTS_PRESENT(_) =                             (((_) >> 15) & 0x01) //col:16577
SEGMENT_ACCESS_RIGHTS_SYSTEM_BIT =                             20 //col:16586
SEGMENT_ACCESS_RIGHTS_SYSTEM_FLAG =                            0x100000 //col:16587
SEGMENT_ACCESS_RIGHTS_SYSTEM_MASK =                            0x01 //col:16588
SEGMENT_ACCESS_RIGHTS_SYSTEM(_) =                              (((_) >> 20) & 0x01) //col:16589
SEGMENT_ACCESS_RIGHTS_LONG_MODE_BIT =                          21 //col:16601
SEGMENT_ACCESS_RIGHTS_LONG_MODE_FLAG =                         0x200000 //col:16602
SEGMENT_ACCESS_RIGHTS_LONG_MODE_MASK =                         0x01 //col:16603
SEGMENT_ACCESS_RIGHTS_LONG_MODE(_) =                           (((_) >> 21) & 0x01) //col:16604
SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_BIT =                        22 //col:16626
SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_FLAG =                       0x400000 //col:16627
SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_MASK =                       0x01 //col:16628
SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG(_) =                         (((_) >> 22) & 0x01) //col:16629
SEGMENT_ACCESS_RIGHTS_GRANULARITY_BIT =                        23 //col:16641
SEGMENT_ACCESS_RIGHTS_GRANULARITY_FLAG =                       0x800000 //col:16642
SEGMENT_ACCESS_RIGHTS_GRANULARITY_MASK =                       0x01 //col:16643
SEGMENT_ACCESS_RIGHTS_GRANULARITY(_) =                         (((_) >> 23) & 0x01) //col:16644
SEGMENT__BASE_ADDRESS_MIDDLE_BIT =                             0 //col:16711
SEGMENT__BASE_ADDRESS_MIDDLE_FLAG =                            0xFF //col:16712
SEGMENT__BASE_ADDRESS_MIDDLE_MASK =                            0xFF //col:16713
SEGMENT__BASE_ADDRESS_MIDDLE(_) =                              (((_) >> 0) & 0xFF) //col:16714
SEGMENT__TYPE_BIT =                                            8 //col:16727
SEGMENT__TYPE_FLAG =                                           0xF00 //col:16728
SEGMENT__TYPE_MASK =                                           0x0F //col:16729
SEGMENT__TYPE(_) =                                             (((_) >> 8) & 0x0F) //col:16730
SEGMENT__DESCRIPTOR_TYPE_BIT =                                 12 //col:16739
SEGMENT__DESCRIPTOR_TYPE_FLAG =                                0x1000 //col:16740
SEGMENT__DESCRIPTOR_TYPE_MASK =                                0x01 //col:16741
SEGMENT__DESCRIPTOR_TYPE(_) =                                  (((_) >> 12) & 0x01) //col:16742
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT =                      13 //col:16752
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG =                     0x6000 //col:16753
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK =                     0x03 //col:16754
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_) =                       (((_) >> 13) & 0x03) //col:16755
SEGMENT__PRESENT_BIT =                                         15 //col:16766
SEGMENT__PRESENT_FLAG =                                        0x8000 //col:16767
SEGMENT__PRESENT_MASK =                                        0x01 //col:16768
SEGMENT__PRESENT(_) =                                          (((_) >> 15) & 0x01) //col:16769
SEGMENT__SEGMENT_LIMIT_HIGH_BIT =                              16 //col:16775
SEGMENT__SEGMENT_LIMIT_HIGH_FLAG =                             0xF0000 //col:16776
SEGMENT__SEGMENT_LIMIT_HIGH_MASK =                             0x0F //col:16777
SEGMENT__SEGMENT_LIMIT_HIGH(_) =                               (((_) >> 16) & 0x0F) //col:16778
SEGMENT__SYSTEM_BIT =                                          20 //col:16786
SEGMENT__SYSTEM_FLAG =                                         0x100000 //col:16787
SEGMENT__SYSTEM_MASK =                                         0x01 //col:16788
SEGMENT__SYSTEM(_) =                                           (((_) >> 20) & 0x01) //col:16789
SEGMENT__LONG_MODE_BIT =                                       21 //col:16801
SEGMENT__LONG_MODE_FLAG =                                      0x200000 //col:16802
SEGMENT__LONG_MODE_MASK =                                      0x01 //col:16803
SEGMENT__LONG_MODE(_) =                                        (((_) >> 21) & 0x01) //col:16804
SEGMENT__DEFAULT_BIG_BIT =                                     22 //col:16826
SEGMENT__DEFAULT_BIG_FLAG =                                    0x400000 //col:16827
SEGMENT__DEFAULT_BIG_MASK =                                    0x01 //col:16828
SEGMENT__DEFAULT_BIG(_) =                                      (((_) >> 22) & 0x01) //col:16829
SEGMENT__GRANULARITY_BIT =                                     23 //col:16841
SEGMENT__GRANULARITY_FLAG =                                    0x800000 //col:16842
SEGMENT__GRANULARITY_MASK =                                    0x01 //col:16843
SEGMENT__GRANULARITY(_) =                                      (((_) >> 23) & 0x01) //col:16844
SEGMENT__BASE_ADDRESS_HIGH_BIT =                               24 //col:16850
SEGMENT__BASE_ADDRESS_HIGH_FLAG =                              0xFF000000 //col:16851
SEGMENT__BASE_ADDRESS_HIGH_MASK =                              0xFF //col:16852
SEGMENT__BASE_ADDRESS_HIGH(_) =                                (((_) >> 24) & 0xFF) //col:16853
SEGMENT__BASE_ADDRESS_MIDDLE_BIT =                             0 //col:16913
SEGMENT__BASE_ADDRESS_MIDDLE_FLAG =                            0xFF //col:16914
SEGMENT__BASE_ADDRESS_MIDDLE_MASK =                            0xFF //col:16915
SEGMENT__BASE_ADDRESS_MIDDLE(_) =                              (((_) >> 0) & 0xFF) //col:16916
SEGMENT__TYPE_BIT =                                            8 //col:16929
SEGMENT__TYPE_FLAG =                                           0xF00 //col:16930
SEGMENT__TYPE_MASK =                                           0x0F //col:16931
SEGMENT__TYPE(_) =                                             (((_) >> 8) & 0x0F) //col:16932
SEGMENT__DESCRIPTOR_TYPE_BIT =                                 12 //col:16941
SEGMENT__DESCRIPTOR_TYPE_FLAG =                                0x1000 //col:16942
SEGMENT__DESCRIPTOR_TYPE_MASK =                                0x01 //col:16943
SEGMENT__DESCRIPTOR_TYPE(_) =                                  (((_) >> 12) & 0x01) //col:16944
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT =                      13 //col:16954
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG =                     0x6000 //col:16955
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK =                     0x03 //col:16956
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_) =                       (((_) >> 13) & 0x03) //col:16957
SEGMENT__PRESENT_BIT =                                         15 //col:16968
SEGMENT__PRESENT_FLAG =                                        0x8000 //col:16969
SEGMENT__PRESENT_MASK =                                        0x01 //col:16970
SEGMENT__PRESENT(_) =                                          (((_) >> 15) & 0x01) //col:16971
SEGMENT__SEGMENT_LIMIT_HIGH_BIT =                              16 //col:16977
SEGMENT__SEGMENT_LIMIT_HIGH_FLAG =                             0xF0000 //col:16978
SEGMENT__SEGMENT_LIMIT_HIGH_MASK =                             0x0F //col:16979
SEGMENT__SEGMENT_LIMIT_HIGH(_) =                               (((_) >> 16) & 0x0F) //col:16980
SEGMENT__SYSTEM_BIT =                                          20 //col:16988
SEGMENT__SYSTEM_FLAG =                                         0x100000 //col:16989
SEGMENT__SYSTEM_MASK =                                         0x01 //col:16990
SEGMENT__SYSTEM(_) =                                           (((_) >> 20) & 0x01) //col:16991
SEGMENT__LONG_MODE_BIT =                                       21 //col:17003
SEGMENT__LONG_MODE_FLAG =                                      0x200000 //col:17004
SEGMENT__LONG_MODE_MASK =                                      0x01 //col:17005
SEGMENT__LONG_MODE(_) =                                        (((_) >> 21) & 0x01) //col:17006
SEGMENT__DEFAULT_BIG_BIT =                                     22 //col:17028
SEGMENT__DEFAULT_BIG_FLAG =                                    0x400000 //col:17029
SEGMENT__DEFAULT_BIG_MASK =                                    0x01 //col:17030
SEGMENT__DEFAULT_BIG(_) =                                      (((_) >> 22) & 0x01) //col:17031
SEGMENT__GRANULARITY_BIT =                                     23 //col:17043
SEGMENT__GRANULARITY_FLAG =                                    0x800000 //col:17044
SEGMENT__GRANULARITY_MASK =                                    0x01 //col:17045
SEGMENT__GRANULARITY(_) =                                      (((_) >> 23) & 0x01) //col:17046
SEGMENT__BASE_ADDRESS_HIGH_BIT =                               24 //col:17052
SEGMENT__BASE_ADDRESS_HIGH_FLAG =                              0xFF000000 //col:17053
SEGMENT__BASE_ADDRESS_HIGH_MASK =                              0xFF //col:17054
SEGMENT__BASE_ADDRESS_HIGH(_) =                                (((_) >> 24) & 0xFF) //col:17055
SEGMENT__INTERRUPT_STACK_TABLE_BIT =                           0 //col:17097
SEGMENT__INTERRUPT_STACK_TABLE_FLAG =                          0x07 //col:17098
SEGMENT__INTERRUPT_STACK_TABLE_MASK =                          0x07 //col:17099
SEGMENT__INTERRUPT_STACK_TABLE(_) =                            (((_) >> 0) & 0x07) //col:17100
SEGMENT__MUST_BE_ZERO_0_BIT =                                  3 //col:17106
SEGMENT__MUST_BE_ZERO_0_FLAG =                                 0xF8 //col:17107
SEGMENT__MUST_BE_ZERO_0_MASK =                                 0x1F //col:17108
SEGMENT__MUST_BE_ZERO_0(_) =                                   (((_) >> 3) & 0x1F) //col:17109
SEGMENT__TYPE_BIT =                                            8 //col:17115
SEGMENT__TYPE_FLAG =                                           0xF00 //col:17116
SEGMENT__TYPE_MASK =                                           0x0F //col:17117
SEGMENT__TYPE(_) =                                             (((_) >> 8) & 0x0F) //col:17118
SEGMENT__MUST_BE_ZERO_1_BIT =                                  12 //col:17124
SEGMENT__MUST_BE_ZERO_1_FLAG =                                 0x1000 //col:17125
SEGMENT__MUST_BE_ZERO_1_MASK =                                 0x01 //col:17126
SEGMENT__MUST_BE_ZERO_1(_) =                                   (((_) >> 12) & 0x01) //col:17127
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT =                      13 //col:17133
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG =                     0x6000 //col:17134
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK =                     0x03 //col:17135
SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_) =                       (((_) >> 13) & 0x03) //col:17136
SEGMENT__PRESENT_BIT =                                         15 //col:17142
SEGMENT__PRESENT_FLAG =                                        0x8000 //col:17143
SEGMENT__PRESENT_MASK =                                        0x01 //col:17144
SEGMENT__PRESENT(_) =                                          (((_) >> 15) & 0x01) //col:17145
SEGMENT__OFFSET_MIDDLE_BIT =                                   16 //col:17151
SEGMENT__OFFSET_MIDDLE_FLAG =                                  0xFFFF0000 //col:17152
SEGMENT__OFFSET_MIDDLE_MASK =                                  0xFFFF //col:17153
SEGMENT__OFFSET_MIDDLE(_) =                                    (((_) >> 16) & 0xFFFF) //col:17154
SEGMENT_DESCRIPTOR_TYPE_SYSTEM =                               0x00000000 //col:17168
SEGMENT_DESCRIPTOR_TYPE_CODE_OR_DATA =                         0x00000001 //col:17169
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY =                       0x00000000 //col:17187
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY_ACCESSED =              0x00000001 //col:17192
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE =                      0x00000002 //col:17197
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE_ACCESSED =             0x00000003 //col:17202
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY_EXPAND_DOWN =           0x00000004 //col:17207
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY_EXPAND_DOWN_ACCESSED =  0x00000005 //col:17212
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE_EXPAND_DOWN =          0x00000006 //col:17217
SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE_EXPAND_DOWN_ACCESSED = 0x00000007 //col:17222
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY =                    0x00000008 //col:17227
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY_ACCESSED =           0x00000009 //col:17232
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ =                    0x0000000A //col:17237
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ_ACCESSED =           0x0000000B //col:17242
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY_CONFORMING =         0x0000000C //col:17247
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY_CONFORMING_ACCESSED = 0x0000000D //col:17252
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ_CONFORMING =         0x0000000E //col:17257
SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ_CONFORMING_ACCESSED = 0x0000000F //col:17262
SEGMENT_DESCRIPTOR_TYPE_RESERVED_1 =                           0x00000000 //col:17291
SEGMENT_DESCRIPTOR_TYPE_TSS_16_AVAILABLE =                     0x00000001 //col:17297
SEGMENT_DESCRIPTOR_TYPE_LDT =                                  0x00000002 //col:17303
SEGMENT_DESCRIPTOR_TYPE_TSS_16_BUSY =                          0x00000003 //col:17309
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE_16 =                         0x00000004 //col:17315
SEGMENT_DESCRIPTOR_TYPE_TASK_GATE =                            0x00000005 //col:17321
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE_16 =                    0x00000006 //col:17327
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE_16 =                         0x00000007 //col:17333
SEGMENT_DESCRIPTOR_TYPE_RESERVED_2 =                           0x00000008 //col:17339
SEGMENT_DESCRIPTOR_TYPE_TSS_AVAILABLE =                        0x00000009 //col:17345
SEGMENT_DESCRIPTOR_TYPE_RESERVED_3 =                           0x0000000A //col:17351
SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY =                             0x0000000B //col:17357
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE =                            0x0000000C //col:17363
SEGMENT_DESCRIPTOR_TYPE_RESERVED_4 =                           0x0000000D //col:17369
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE =                       0x0000000E //col:17375
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE =                            0x0000000F //col:17381
SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_BIT =                 0 //col:17403
SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_FLAG =                0x03 //col:17404
SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_MASK =                0x03 //col:17405
SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL(_) =                  (((_) >> 0) & 0x03) //col:17406
SEGMENT_SELECTOR_TABLE_BIT =                                   2 //col:17413
SEGMENT_SELECTOR_TABLE_FLAG =                                  0x04 //col:17414
SEGMENT_SELECTOR_TABLE_MASK =                                  0x01 //col:17415
SEGMENT_SELECTOR_TABLE(_) =                                    (((_) >> 2) & 0x01) //col:17416
SEGMENT_SELECTOR_INDEX_BIT =                                   3 //col:17424
SEGMENT_SELECTOR_INDEX_FLAG =                                  0xFFF8 //col:17425
SEGMENT_SELECTOR_INDEX_MASK =                                  0x1FFF //col:17426
SEGMENT_SELECTOR_INDEX(_) =                                    (((_) >> 3) & 0x1FFF) //col:17427
VMX_EXIT_REASON_EXCEPTION_OR_NMI =                             0x00000000 //col:17548
VMX_EXIT_REASON_EXTERNAL_INTERRUPT =                           0x00000001 //col:17555
VMX_EXIT_REASON_TRIPLE_FAULT =                                 0x00000002 //col:17563
VMX_EXIT_REASON_INIT_SIGNAL =                                  0x00000003 //col:17570
VMX_EXIT_REASON_STARTUP_IPI =                                  0x00000004 //col:17577
VMX_EXIT_REASON_IO_SMI =                                       0x00000005 //col:17586
VMX_EXIT_REASON_SMI =                                          0x00000006 //col:17595
VMX_EXIT_REASON_INTERRUPT_WINDOW =                             0x00000007 //col:17603
VMX_EXIT_REASON_NMI_WINDOW =                                   0x00000008 //col:17611
VMX_EXIT_REASON_TASK_SWITCH =                                  0x00000009 //col:17618
VMX_EXIT_REASON_EXECUTE_CPUID =                                0x0000000A //col:17625
VMX_EXIT_REASON_EXECUTE_GETSEC =                               0x0000000B //col:17632
VMX_EXIT_REASON_EXECUTE_HLT =                                  0x0000000C //col:17639
VMX_EXIT_REASON_EXECUTE_INVD =                                 0x0000000D //col:17646
VMX_EXIT_REASON_EXECUTE_INVLPG =                               0x0000000E //col:17653
VMX_EXIT_REASON_EXECUTE_RDPMC =                                0x0000000F //col:17660
VMX_EXIT_REASON_EXECUTE_RDTSC =                                0x00000010 //col:17667
VMX_EXIT_REASON_EXECUTE_RSM_IN_SMM =                           0x00000011 //col:17674
VMX_EXIT_REASON_EXECUTE_VMCALL =                               0x00000012 //col:17684
VMX_EXIT_REASON_EXECUTE_VMCLEAR =                              0x00000013 //col:17691
VMX_EXIT_REASON_EXECUTE_VMLAUNCH =                             0x00000014 //col:17698
VMX_EXIT_REASON_EXECUTE_VMPTRLD =                              0x00000015 //col:17705
VMX_EXIT_REASON_EXECUTE_VMPTRST =                              0x00000016 //col:17712
VMX_EXIT_REASON_EXECUTE_VMREAD =                               0x00000017 //col:17719
VMX_EXIT_REASON_EXECUTE_VMRESUME =                             0x00000018 //col:17726
VMX_EXIT_REASON_EXECUTE_VMWRITE =                              0x00000019 //col:17733
VMX_EXIT_REASON_EXECUTE_VMXOFF =                               0x0000001A //col:17740
VMX_EXIT_REASON_EXECUTE_VMXON =                                0x0000001B //col:17747
VMX_EXIT_REASON_MOV_CR =                                       0x0000001C //col:17759
VMX_EXIT_REASON_MOV_DR =                                       0x0000001D //col:17766
VMX_EXIT_REASON_EXECUTE_IO_INSTRUCTION =                       0x0000001E //col:17776
VMX_EXIT_REASON_EXECUTE_RDMSR =                                0x0000001F //col:17789
VMX_EXIT_REASON_EXECUTE_WRMSR =                                0x00000020 //col:17802
VMX_EXIT_REASON_ERROR_INVALID_GUEST_STATE =                    0x00000021 //col:17809
VMX_EXIT_REASON_ERROR_MSR_LOAD =                               0x00000022 //col:17816
VMX_EXIT_REASON_EXECUTE_MWAIT =                                0x00000024 //col:17823
VMX_EXIT_REASON_MONITOR_TRAP_FLAG =                            0x00000025 //col:17833
VMX_EXIT_REASON_EXECUTE_MONITOR =                              0x00000027 //col:17840
VMX_EXIT_REASON_EXECUTE_PAUSE =                                0x00000028 //col:17850
VMX_EXIT_REASON_ERROR_MACHINE_CHECK =                          0x00000029 //col:17859
VMX_EXIT_REASON_TPR_BELOW_THRESHOLD =                          0x0000002B //col:17871
VMX_EXIT_REASON_APIC_ACCESS =                                  0x0000002C //col:17881
VMX_EXIT_REASON_VIRTUALIZED_EOI =                              0x0000002D //col:17888
VMX_EXIT_REASON_GDTR_IDTR_ACCESS =                             0x0000002E //col:17896
VMX_EXIT_REASON_LDTR_TR_ACCESS =                               0x0000002F //col:17904
VMX_EXIT_REASON_EPT_VIOLATION =                                0x00000030 //col:17912
VMX_EXIT_REASON_EPT_MISCONFIGURATION =                         0x00000031 //col:17919
VMX_EXIT_REASON_EXECUTE_INVEPT =                               0x00000032 //col:17926
VMX_EXIT_REASON_EXECUTE_RDTSCP =                               0x00000033 //col:17934
VMX_EXIT_REASON_VMX_PREEMPTION_TIMER_EXPIRED =                 0x00000034 //col:17941
VMX_EXIT_REASON_EXECUTE_INVVPID =                              0x00000035 //col:17948
VMX_EXIT_REASON_EXECUTE_WBINVD =                               0x00000036 //col:17955
VMX_EXIT_REASON_EXECUTE_XSETBV =                               0x00000037 //col:17962
VMX_EXIT_REASON_APIC_WRITE =                                   0x00000038 //col:17971
VMX_EXIT_REASON_EXECUTE_RDRAND =                               0x00000039 //col:17978
VMX_EXIT_REASON_EXECUTE_INVPCID =                              0x0000003A //col:17986
VMX_EXIT_REASON_EXECUTE_VMFUNC =                               0x0000003B //col:17994
VMX_EXIT_REASON_EXECUTE_ENCLS =                                0x0000003C //col:18003
VMX_EXIT_REASON_EXECUTE_RDSEED =                               0x0000003D //col:18010
VMX_EXIT_REASON_PAGE_MODIFICATION_LOG_FULL =                   0x0000003E //col:18018
VMX_EXIT_REASON_EXECUTE_XSAVES =                               0x0000003F //col:18026
VMX_EXIT_REASON_EXECUTE_XRSTORS =                              0x00000040 //col:18034
VMX_ERROR_VMCALL_IN_VMX_ROOT_OPERATION =                       0x00000001 //col:18051
VMX_ERROR_VMCLEAR_INVALID_PHYSICAL_ADDRESS =                   0x00000002 //col:18056
VMX_ERROR_VMCLEAR_INVALID_VMXON_POINTER =                      0x00000003 //col:18061
VMX_ERROR_VMLAUCH_NON_CLEAR_VMCS =                             0x00000004 //col:18066
VMX_ERROR_VMRESUME_NON_LAUNCHED_VMCS =                         0x00000005 //col:18071
VMX_ERROR_VMRESUME_AFTER_VMXOFF =                              0x00000006 //col:18076
VMX_ERROR_VMENTRY_INVALID_CONTROL_FIELDS =                     0x00000007 //col:18081
VMX_ERROR_VMENTRY_INVALID_HOST_STATE =                         0x00000008 //col:18086
VMX_ERROR_VMPTRLD_INVALID_PHYSICAL_ADDRESS =                   0x00000009 //col:18091
VMX_ERROR_VMPTRLD_VMXON_POINTER =                              0x0000000A //col:18096
VMX_ERROR_VMPTRLD_INCORRECT_VMCS_REVISION_ID =                 0x0000000B //col:18101
VMX_ERROR_VMREAD_VMWRITE_INVALID_COMPONENT =                   0x0000000C //col:18106
VMX_ERROR_VMWRITE_READONLY_COMPONENT =                         0x0000000D //col:18111
VMX_ERROR_VMXON_IN_VMX_ROOT_OP =                               0x0000000F //col:18116
VMX_ERROR_VMENTRY_INVALID_VMCS_EXECUTIVE_POINTER =             0x00000010 //col:18121
VMX_ERROR_VMENTRY_NON_LAUNCHED_EXECUTIVE_VMCS =                0x00000011 //col:18126
VMX_ERROR_VMENTRY_EXECUTIVE_VMCS_PTR =                         0x00000012 //col:18132
VMX_ERROR_VMCALL_NON_CLEAR_VMCS =                              0x00000013 //col:18137
VMX_ERROR_VMCALL_INVALID_VMEXIT_FIELDS =                       0x00000014 //col:18142
VMX_ERROR_VMCALL_INVALID_MSEG_REVISION_ID =                    0x00000016 //col:18147
VMX_ERROR_VMXOFF_DUAL_MONITOR =                                0x00000017 //col:18152
VMX_ERROR_VMCALL_INVALID_SMM_MONITOR =                         0x00000018 //col:18157
VMX_ERROR_VMENTRY_INVALID_VM_EXECUTION_CONTROL =               0x00000019 //col:18162
VMX_ERROR_VMENTRY_MOV_SS =                                     0x0000001A //col:18167
VMX_ERROR_INVEPT_INVVPID_INVALID_OPERAND =                     0x0000001C //col:18172
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_BIT = 0 //col:18253
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_FLAG = 0x0F //col:18254
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_MASK = 0x0F //col:18255
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION(_) = (((_) >> 0) & 0x0F) //col:18256
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_BIT = 13 //col:18265
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_FLAG = 0x2000 //col:18266
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_MASK = 0x01 //col:18267
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED(_) = (((_) >> 13) & 0x01) //col:18268
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_BIT = 14 //col:18277
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_FLAG = 0x4000 //col:18278
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_MASK = 0x01 //col:18279
VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION(_) = (((_) >> 14) & 0x01) //col:18280
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_BIT =              0 //col:18298
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_FLAG =             0xFFFF //col:18299
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_MASK =             0xFFFF //col:18300
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR(_) =               (((_) >> 0) & 0xFFFF) //col:18301
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_BIT =                30 //col:18308
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_FLAG =               0xC0000000 //col:18309
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_MASK =               0x03 //col:18310
VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE(_) =                 (((_) >> 30) & 0x03) //col:18311
VMX_EXIT_QUALIFICATION_TYPE_CALL_INSTRUCTION =                 0x00000000 //col:18312
VMX_EXIT_QUALIFICATION_TYPE_IRET_INSTRUCTION =                 0x00000001 //col:18313
VMX_EXIT_QUALIFICATION_TYPE_JMP_INSTRUCTION =                  0x00000002 //col:18314
VMX_EXIT_QUALIFICATION_TYPE_TASK_GATE_IN_IDT =                 0x00000003 //col:18315
VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_BIT =           0 //col:18334
VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_FLAG =          0x0F //col:18335
VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_MASK =          0x0F //col:18336
VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER(_) =            (((_) >> 0) & 0x0F) //col:18337
VMX_EXIT_QUALIFICATION_REGISTER_CR0 =                          0x00000000 //col:18338
VMX_EXIT_QUALIFICATION_REGISTER_CR2 =                          0x00000002 //col:18339
VMX_EXIT_QUALIFICATION_REGISTER_CR3 =                          0x00000003 //col:18340
VMX_EXIT_QUALIFICATION_REGISTER_CR4 =                          0x00000004 //col:18341
VMX_EXIT_QUALIFICATION_REGISTER_CR8 =                          0x00000008 //col:18342
VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_BIT =                4 //col:18348
VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_FLAG =               0x30 //col:18349
VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_MASK =               0x03 //col:18350
VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE(_) =                 (((_) >> 4) & 0x03) //col:18351
VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR =                      0x00000000 //col:18352
VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR =                    0x00000001 //col:18353
VMX_EXIT_QUALIFICATION_ACCESS_CLTS =                           0x00000002 //col:18354
VMX_EXIT_QUALIFICATION_ACCESS_LMSW =                           0x00000003 //col:18355
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_BIT =          6 //col:18361
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_FLAG =         0x40 //col:18362
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_MASK =         0x01 //col:18363
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE(_) =           (((_) >> 6) & 0x01) //col:18364
VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER =                      0x00000000 //col:18365
VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY =                        0x00000001 //col:18366
VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER_BIT =   8 //col:18373
VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER_FLAG =  0xF00 //col:18374
VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER_MASK =  0x0F //col:18375
VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER(_) =    (((_) >> 8) & 0x0F) //col:18376
VMX_EXIT_QUALIFICATION_GENREG_RAX =                            0x00000000 //col:18377
VMX_EXIT_QUALIFICATION_GENREG_RCX =                            0x00000001 //col:18378
VMX_EXIT_QUALIFICATION_GENREG_RDX =                            0x00000002 //col:18379
VMX_EXIT_QUALIFICATION_GENREG_RBX =                            0x00000003 //col:18380
VMX_EXIT_QUALIFICATION_GENREG_RSP =                            0x00000004 //col:18381
VMX_EXIT_QUALIFICATION_GENREG_RBP =                            0x00000005 //col:18382
VMX_EXIT_QUALIFICATION_GENREG_RSI =                            0x00000006 //col:18383
VMX_EXIT_QUALIFICATION_GENREG_RDI =                            0x00000007 //col:18384
VMX_EXIT_QUALIFICATION_GENREG_R8 =                             0x00000008 //col:18385
VMX_EXIT_QUALIFICATION_GENREG_R9 =                             0x00000009 //col:18386
VMX_EXIT_QUALIFICATION_GENREG_R10 =                            0x0000000A //col:18387
VMX_EXIT_QUALIFICATION_GENREG_R11 =                            0x0000000B //col:18388
VMX_EXIT_QUALIFICATION_GENREG_R12 =                            0x0000000C //col:18389
VMX_EXIT_QUALIFICATION_GENREG_R13 =                            0x0000000D //col:18390
VMX_EXIT_QUALIFICATION_GENREG_R14 =                            0x0000000E //col:18391
VMX_EXIT_QUALIFICATION_GENREG_R15 =                            0x0000000F //col:18392
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_BIT =           16 //col:18399
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_FLAG =          0xFFFF0000 //col:18400
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_MASK =          0xFFFF //col:18401
VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA(_) =            (((_) >> 16) & 0xFFFF) //col:18402
VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER_BIT =             0 //col:18420
VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER_FLAG =            0x07 //col:18421
VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER_MASK =            0x07 //col:18422
VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER(_) =              (((_) >> 0) & 0x07) //col:18423
VMX_EXIT_QUALIFICATION_REGISTER_DR0 =                          0x00000000 //col:18424
VMX_EXIT_QUALIFICATION_REGISTER_DR1 =                          0x00000001 //col:18425
VMX_EXIT_QUALIFICATION_REGISTER_DR2 =                          0x00000002 //col:18426
VMX_EXIT_QUALIFICATION_REGISTER_DR3 =                          0x00000003 //col:18427
VMX_EXIT_QUALIFICATION_REGISTER_DR6 =                          0x00000006 //col:18428
VMX_EXIT_QUALIFICATION_REGISTER_DR7 =                          0x00000007 //col:18429
VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_BIT =        4 //col:18436
VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_FLAG =       0x10 //col:18437
VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_MASK =       0x01 //col:18438
VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS(_) =         (((_) >> 4) & 0x01) //col:18439
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR =                   0x00000000 //col:18440
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR =                 0x00000001 //col:18441
VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_BIT =   8 //col:18448
VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_FLAG =  0xF00 //col:18449
VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_MASK =  0x0F //col:18450
VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER(_) =    (((_) >> 8) & 0x0F) //col:18451
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_BIT =     0 //col:18469
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_FLAG =    0x07 //col:18470
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_MASK =    0x07 //col:18471
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS(_) =      (((_) >> 0) & 0x07) //col:18472
VMX_EXIT_QUALIFICATION_WIDTH_1_BYTE =                          0x00000000 //col:18473
VMX_EXIT_QUALIFICATION_WIDTH_2_BYTE =                          0x00000001 //col:18474
VMX_EXIT_QUALIFICATION_WIDTH_4_BYTE =                          0x00000003 //col:18475
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_BIT = 3 //col:18481
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_FLAG = 0x08 //col:18482
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_MASK = 0x01 //col:18483
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS(_) = (((_) >> 3) & 0x01) //col:18484
VMX_EXIT_QUALIFICATION_DIRECTION_OUT =                         0x00000000 //col:18485
VMX_EXIT_QUALIFICATION_DIRECTION_IN =                          0x00000001 //col:18486
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_BIT = 4 //col:18492
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_FLAG = 0x10 //col:18493
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_MASK = 0x01 //col:18494
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION(_) =  (((_) >> 4) & 0x01) //col:18495
VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING =                  0x00000000 //col:18496
VMX_EXIT_QUALIFICATION_IS_STRING_STRING =                      0x00000001 //col:18497
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_BIT =       5 //col:18503
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_FLAG =      0x20 //col:18504
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_MASK =      0x01 //col:18505
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED(_) =        (((_) >> 5) & 0x01) //col:18506
VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP =                        0x00000000 //col:18507
VMX_EXIT_QUALIFICATION_IS_REP_REP =                            0x00000001 //col:18508
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_BIT =   6 //col:18514
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_FLAG =  0x40 //col:18515
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_MASK =  0x01 //col:18516
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING(_) =    (((_) >> 6) & 0x01) //col:18517
VMX_EXIT_QUALIFICATION_ENCODING_DX =                           0x00000000 //col:18518
VMX_EXIT_QUALIFICATION_ENCODING_IMMEDIATE =                    0x00000001 //col:18519
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_BIT =        16 //col:18526
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_FLAG =       0xFFFF0000 //col:18527
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_MASK =       0xFFFF //col:18528
VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER(_) =         (((_) >> 16) & 0xFFFF) //col:18529
VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_BIT =           0 //col:18548
VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_FLAG =          0xFFF //col:18549
VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_MASK =          0xFFF //col:18550
VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET(_) =            (((_) >> 0) & 0xFFF) //col:18551
VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE_BIT =           12 //col:18557
VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE_FLAG =          0xF000 //col:18558
VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE_MASK =          0x0F //col:18559
VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE(_) =            (((_) >> 12) & 0x0F) //col:18560
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_READ =                      0x00000000 //col:18564
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_WRITE =                     0x00000001 //col:18569
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_INSTRUCTION_FETCH =         0x00000002 //col:18574
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_EVENT_DELIVERY =            0x00000003 //col:18579
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_EVENT_DELIVERY =          0x0000000A //col:18584
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_INSTRUCTION_FETCH =       0x0000000F //col:18589
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_BIT =         0 //col:18607
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_FLAG =        0x01 //col:18608
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_MASK =        0x01 //col:18609
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS(_) =          (((_) >> 0) & 0x01) //col:18610
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_BIT =        1 //col:18616
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_FLAG =       0x02 //col:18617
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_MASK =       0x01 //col:18618
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS(_) =         (((_) >> 1) & 0x01) //col:18619
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_BIT =      2 //col:18625
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_FLAG =     0x04 //col:18626
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_MASK =     0x01 //col:18627
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS(_) =       (((_) >> 2) & 0x01) //col:18628
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_BIT =        3 //col:18635
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_FLAG =       0x08 //col:18636
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_MASK =       0x01 //col:18637
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE(_) =         (((_) >> 3) & 0x01) //col:18638
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_BIT =       4 //col:18645
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_FLAG =      0x10 //col:18646
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_MASK =      0x01 //col:18647
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE(_) =        (((_) >> 4) & 0x01) //col:18648
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_BIT =      5 //col:18658
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FLAG =     0x20 //col:18659
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_MASK =     0x01 //col:18660
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE(_) =       (((_) >> 5) & 0x01) //col:18661
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_BIT = 6 //col:18670
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_FLAG = 0x40 //col:18671
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_MASK = 0x01 //col:18672
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE(_) = (((_) >> 6) & 0x01) //col:18673
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_BIT = 7 //col:18680
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_FLAG = 0x80 //col:18681
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_MASK = 0x01 //col:18682
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS(_) = (((_) >> 7) & 0x01) //col:18683
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_BIT = 8 //col:18694
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_FLAG = 0x100 //col:18695
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_MASK = 0x01 //col:18696
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION(_) = (((_) >> 8) & 0x01) //col:18697
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_BIT = 9 //col:18707
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_FLAG = 0x200 //col:18708
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_MASK = 0x01 //col:18709
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS(_) = (((_) >> 9) & 0x01) //col:18710
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_BIT = 10 //col:18720
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_FLAG = 0x400 //col:18721
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_MASK = 0x01 //col:18722
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE(_) = (((_) >> 10) & 0x01) //col:18723
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_BIT = 11 //col:18733
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_FLAG = 0x800 //col:18734
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_MASK = 0x01 //col:18735
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE(_) = (((_) >> 11) & 0x01) //col:18736
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_BIT =      12 //col:18742
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_FLAG =     0x1000 //col:18743
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_MASK =     0x01 //col:18744
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING(_) =       (((_) >> 12) & 0x01) //col:18745
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_BIT = 13 //col:18751
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_FLAG = 0x2000 //col:18752
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_MASK = 0x01 //col:18753
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS(_) =  (((_) >> 13) & 0x01) //col:18754
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_BIT = 14 //col:18763
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_FLAG = 0x4000 //col:18764
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_MASK = 0x01 //col:18765
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK(_) = (((_) >> 14) & 0x01) //col:18766
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_BIT = 15 //col:18772
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_FLAG = 0x8000 //col:18773
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_MASK = 0x01 //col:18774
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION(_) = (((_) >> 15) & 0x01) //col:18775
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_BIT = 16 //col:18782
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_FLAG = 0x10000 //col:18783
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_MASK = 0x01 //col:18784
VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION(_) = (((_) >> 16) & 0x01) //col:18785
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_BIT =        7 //col:18823
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_FLAG =       0x380 //col:18824
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_MASK =       0x07 //col:18825
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE(_) =         (((_) >> 7) & 0x07) //col:18826
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_BIT =    15 //col:18841
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_FLAG =   0x38000 //col:18842
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_MASK =   0x07 //col:18843
VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER(_) =     (((_) >> 15) & 0x07) //col:18844
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_BIT =           0 //col:18868
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_FLAG =          0x03 //col:18869
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_MASK =          0x03 //col:18870
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING(_) =            (((_) >> 0) & 0x03) //col:18871
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_BIT =      7 //col:18883
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_FLAG =     0x380 //col:18884
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_MASK =     0x07 //col:18885
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE(_) =       (((_) >> 7) & 0x07) //col:18886
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_BIT =  15 //col:18901
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_FLAG = 0x38000 //col:18902
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_MASK = 0x07 //col:18903
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER(_) =   (((_) >> 15) & 0x07) //col:18904
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_BIT = 18 //col:18910
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_FLAG = 0x3C0000 //col:18911
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_MASK = 0x0F //col:18912
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER(_) = (((_) >> 18) & 0x0F) //col:18913
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_BIT = 22 //col:18919
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_FLAG = 0x400000 //col:18920
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_MASK = 0x01 //col:18921
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID(_) = (((_) >> 22) & 0x01) //col:18922
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_BIT =     23 //col:18929
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_FLAG =    0x7800000 //col:18930
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_MASK =    0x0F //col:18931
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER(_) =      (((_) >> 23) & 0x0F) //col:18932
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_BIT = 27 //col:18938
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_FLAG = 0x8000000 //col:18939
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_MASK = 0x01 //col:18940
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID(_) = (((_) >> 27) & 0x01) //col:18941
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_BIT =        28 //col:18947
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_FLAG =       0xF0000000 //col:18948
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_MASK =       0x0F //col:18949
VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2(_) =         (((_) >> 28) & 0x0F) //col:18950
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_BIT =     0 //col:18974
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_FLAG =    0x03 //col:18975
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_MASK =    0x03 //col:18976
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING(_) =      (((_) >> 0) & 0x03) //col:18977
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_BIT = 7 //col:18989
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_FLAG = 0x380 //col:18990
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_MASK = 0x07 //col:18991
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE(_) = (((_) >> 7) & 0x07) //col:18992
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_BIT = 11 //col:19003
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_FLAG = 0x800 //col:19004
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_MASK = 0x01 //col:19005
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE(_) = (((_) >> 11) & 0x01) //col:19006
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_BIT = 15 //col:19021
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_FLAG = 0x38000 //col:19022
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_MASK = 0x07 //col:19023
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER(_) = (((_) >> 15) & 0x07) //col:19024
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_BIT = 18 //col:19030
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_FLAG = 0x3C0000 //col:19031
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_MASK = 0x0F //col:19032
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER(_) = (((_) >> 18) & 0x0F) //col:19033
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_BIT = 22 //col:19039
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_FLAG = 0x400000 //col:19040
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_MASK = 0x01 //col:19041
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID(_) = (((_) >> 22) & 0x01) //col:19042
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_BIT = 23 //col:19049
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_FLAG = 0x7800000 //col:19050
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_MASK = 0x0F //col:19051
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER(_) = (((_) >> 23) & 0x0F) //col:19052
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_BIT = 27 //col:19058
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_FLAG = 0x8000000 //col:19059
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_MASK = 0x01 //col:19060
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID(_) = (((_) >> 27) & 0x01) //col:19061
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_BIT = 28 //col:19072
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_FLAG = 0x30000000 //col:19073
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_MASK = 0x03 //col:19074
VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION(_) =  (((_) >> 28) & 0x03) //col:19075
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_BIT =       0 //col:19099
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_FLAG =      0x03 //col:19100
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_MASK =      0x03 //col:19101
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING(_) =        (((_) >> 0) & 0x03) //col:19102
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_BIT =         3 //col:19109
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_FLAG =        0x78 //col:19110
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_MASK =        0x0F //col:19111
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1(_) =          (((_) >> 3) & 0x0F) //col:19112
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_BIT =  7 //col:19123
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_FLAG = 0x380 //col:19124
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_MASK = 0x07 //col:19125
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE(_) =   (((_) >> 7) & 0x07) //col:19126
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_BIT = 10 //col:19132
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_FLAG = 0x400 //col:19133
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_MASK = 0x01 //col:19134
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER(_) = (((_) >> 10) & 0x01) //col:19135
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_BIT = 15 //col:19150
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_FLAG = 0x38000 //col:19151
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_MASK = 0x07 //col:19152
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER(_) = (((_) >> 15) & 0x07) //col:19153
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_BIT = 18 //col:19160
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_FLAG = 0x3C0000 //col:19161
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_MASK = 0x0F //col:19162
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER(_) = (((_) >> 18) & 0x0F) //col:19163
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_BIT = 22 //col:19169
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_FLAG = 0x400000 //col:19170
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_MASK = 0x01 //col:19171
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID(_) = (((_) >> 22) & 0x01) //col:19172
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_BIT = 23 //col:19179
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_FLAG = 0x7800000 //col:19180
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_MASK = 0x0F //col:19181
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER(_) =  (((_) >> 23) & 0x0F) //col:19182
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_BIT = 27 //col:19188
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_FLAG = 0x8000000 //col:19189
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_MASK = 0x01 //col:19190
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID(_) = (((_) >> 27) & 0x01) //col:19191
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_BIT =   28 //col:19202
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_FLAG =  0x30000000 //col:19203
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_MASK =  0x03 //col:19204
VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION(_) =    (((_) >> 28) & 0x03) //col:19205
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_BIT = 3 //col:19225
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_FLAG = 0x78 //col:19226
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_MASK = 0x0F //col:19227
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER(_) = (((_) >> 3) & 0x0F) //col:19228
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_BIT =   11 //col:19240
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_FLAG =  0x1800 //col:19241
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_MASK =  0x03 //col:19242
VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE(_) =    (((_) >> 11) & 0x03) //col:19243
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_BIT =       0 //col:19267
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_FLAG =      0x03 //col:19268
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_MASK =      0x03 //col:19269
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING(_) =        (((_) >> 0) & 0x03) //col:19270
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_BIT =  7 //col:19282
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_FLAG = 0x380 //col:19283
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_MASK = 0x07 //col:19284
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE(_) =   (((_) >> 7) & 0x07) //col:19285
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_BIT = 15 //col:19300
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_FLAG = 0x38000 //col:19301
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_MASK = 0x07 //col:19302
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER(_) = (((_) >> 15) & 0x07) //col:19303
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_BIT = 18 //col:19309
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_FLAG = 0x3C0000 //col:19310
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_MASK = 0x0F //col:19311
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER(_) = (((_) >> 18) & 0x0F) //col:19312
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_BIT = 22 //col:19318
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_FLAG = 0x400000 //col:19319
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_MASK = 0x01 //col:19320
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID(_) = (((_) >> 22) & 0x01) //col:19321
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_BIT = 23 //col:19328
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_FLAG = 0x7800000 //col:19329
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_MASK = 0x0F //col:19330
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER(_) =  (((_) >> 23) & 0x0F) //col:19331
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_BIT = 27 //col:19337
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_FLAG = 0x8000000 //col:19338
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_MASK = 0x01 //col:19339
VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID(_) = (((_) >> 27) & 0x01) //col:19340
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_BIT =       0 //col:19365
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_FLAG =      0x03 //col:19366
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_MASK =      0x03 //col:19367
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING(_) =        (((_) >> 0) & 0x03) //col:19368
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_BIT =    3 //col:19375
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_FLAG =   0x78 //col:19376
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_MASK =   0x0F //col:19377
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1(_) =     (((_) >> 3) & 0x0F) //col:19378
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_BIT =  7 //col:19389
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_FLAG = 0x380 //col:19390
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_MASK = 0x07 //col:19391
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE(_) =   (((_) >> 7) & 0x07) //col:19392
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_BIT = 10 //col:19398
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_FLAG = 0x400 //col:19399
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_MASK = 0x01 //col:19400
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER(_) = (((_) >> 10) & 0x01) //col:19401
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_BIT = 15 //col:19416
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_FLAG = 0x38000 //col:19417
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_MASK = 0x07 //col:19418
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER(_) = (((_) >> 15) & 0x07) //col:19419
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_BIT = 18 //col:19426
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_FLAG = 0x3C0000 //col:19427
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_MASK = 0x0F //col:19428
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER(_) = (((_) >> 18) & 0x0F) //col:19429
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_BIT = 22 //col:19435
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_FLAG = 0x400000 //col:19436
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_MASK = 0x01 //col:19437
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID(_) = (((_) >> 22) & 0x01) //col:19438
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_BIT = 23 //col:19445
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_FLAG = 0x7800000 //col:19446
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_MASK = 0x0F //col:19447
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER(_) =  (((_) >> 23) & 0x0F) //col:19448
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_BIT = 27 //col:19454
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_FLAG = 0x8000000 //col:19455
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_MASK = 0x01 //col:19456
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID(_) = (((_) >> 27) & 0x01) //col:19457
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_BIT =    28 //col:19463
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_FLAG =   0xF0000000 //col:19464
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_MASK =   0x0F //col:19465
VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2(_) =     (((_) >> 28) & 0x0F) //col:19466
VMX_SEGMENT_ACCESS_RIGHTS_TYPE_BIT =                           0 //col:19501
VMX_SEGMENT_ACCESS_RIGHTS_TYPE_FLAG =                          0x0F //col:19502
VMX_SEGMENT_ACCESS_RIGHTS_TYPE_MASK =                          0x0F //col:19503
VMX_SEGMENT_ACCESS_RIGHTS_TYPE(_) =                            (((_) >> 0) & 0x0F) //col:19504
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_BIT =                4 //col:19510
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_FLAG =               0x10 //col:19511
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_MASK =               0x01 //col:19512
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE(_) =                 (((_) >> 4) & 0x01) //col:19513
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_BIT =     5 //col:19519
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_FLAG =    0x60 //col:19520
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_MASK =    0x03 //col:19521
VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL(_) =      (((_) >> 5) & 0x03) //col:19522
VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_BIT =                        7 //col:19528
VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_FLAG =                       0x80 //col:19529
VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_MASK =                       0x01 //col:19530
VMX_SEGMENT_ACCESS_RIGHTS_PRESENT(_) =                         (((_) >> 7) & 0x01) //col:19531
VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_BIT =                  12 //col:19538
VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_FLAG =                 0x1000 //col:19539
VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_MASK =                 0x01 //col:19540
VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT(_) =                   (((_) >> 12) & 0x01) //col:19541
VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_BIT =                      13 //col:19547
VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_FLAG =                     0x2000 //col:19548
VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_MASK =                     0x01 //col:19549
VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE(_) =                       (((_) >> 13) & 0x01) //col:19550
VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_BIT =                    14 //col:19556
VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_FLAG =                   0x4000 //col:19557
VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_MASK =                   0x01 //col:19558
VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG(_) =                     (((_) >> 14) & 0x01) //col:19559
VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_BIT =                    15 //col:19565
VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_FLAG =                   0x8000 //col:19566
VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_MASK =                   0x01 //col:19567
VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY(_) =                     (((_) >> 15) & 0x01) //col:19568
VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_BIT =                       16 //col:19574
VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_FLAG =                      0x10000 //col:19575
VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_MASK =                      0x01 //col:19576
VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE(_) =                        (((_) >> 16) & 0x01) //col:19577
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_BIT =               0 //col:19602
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_FLAG =              0x01 //col:19603
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_MASK =              0x01 //col:19604
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI(_) =                (((_) >> 0) & 0x01) //col:19605
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_BIT =            1 //col:19615
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_FLAG =           0x02 //col:19616
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_MASK =           0x01 //col:19617
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS(_) =             (((_) >> 1) & 0x01) //col:19618
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_BIT =               2 //col:19627
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_FLAG =              0x04 //col:19628
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_MASK =              0x01 //col:19629
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI(_) =                (((_) >> 2) & 0x01) //col:19630
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_BIT =               3 //col:19644
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_FLAG =              0x08 //col:19645
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_MASK =              0x01 //col:19646
VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI(_) =                (((_) >> 3) & 0x01) //col:19647
VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_BIT =          4 //col:19653
VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_FLAG =         0x10 //col:19654
VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_MASK =         0x01 //col:19655
VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION(_) =           (((_) >> 4) & 0x01) //col:19656
VMX_PENDING_DEBUG_EXCEPTIONS_B0_BIT =                          0 //col:19701
VMX_PENDING_DEBUG_EXCEPTIONS_B0_FLAG =                         0x01 //col:19702
VMX_PENDING_DEBUG_EXCEPTIONS_B0_MASK =                         0x01 //col:19703
VMX_PENDING_DEBUG_EXCEPTIONS_B0(_) =                           (((_) >> 0) & 0x01) //col:19704
VMX_PENDING_DEBUG_EXCEPTIONS_B1_BIT =                          1 //col:19711
VMX_PENDING_DEBUG_EXCEPTIONS_B1_FLAG =                         0x02 //col:19712
VMX_PENDING_DEBUG_EXCEPTIONS_B1_MASK =                         0x01 //col:19713
VMX_PENDING_DEBUG_EXCEPTIONS_B1(_) =                           (((_) >> 1) & 0x01) //col:19714
VMX_PENDING_DEBUG_EXCEPTIONS_B2_BIT =                          2 //col:19721
VMX_PENDING_DEBUG_EXCEPTIONS_B2_FLAG =                         0x04 //col:19722
VMX_PENDING_DEBUG_EXCEPTIONS_B2_MASK =                         0x01 //col:19723
VMX_PENDING_DEBUG_EXCEPTIONS_B2(_) =                           (((_) >> 2) & 0x01) //col:19724
VMX_PENDING_DEBUG_EXCEPTIONS_B3_BIT =                          3 //col:19731
VMX_PENDING_DEBUG_EXCEPTIONS_B3_FLAG =                         0x08 //col:19732
VMX_PENDING_DEBUG_EXCEPTIONS_B3_MASK =                         0x01 //col:19733
VMX_PENDING_DEBUG_EXCEPTIONS_B3(_) =                           (((_) >> 3) & 0x01) //col:19734
VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_BIT =          12 //col:19741
VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_FLAG =         0x1000 //col:19742
VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_MASK =         0x01 //col:19743
VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT(_) =           (((_) >> 12) & 0x01) //col:19744
VMX_PENDING_DEBUG_EXCEPTIONS_BS_BIT =                          14 //col:19751
VMX_PENDING_DEBUG_EXCEPTIONS_BS_FLAG =                         0x4000 //col:19752
VMX_PENDING_DEBUG_EXCEPTIONS_BS_MASK =                         0x01 //col:19753
VMX_PENDING_DEBUG_EXCEPTIONS_BS(_) =                           (((_) >> 14) & 0x01) //col:19754
VMX_PENDING_DEBUG_EXCEPTIONS_RTM_BIT =                         16 //col:19762
VMX_PENDING_DEBUG_EXCEPTIONS_RTM_FLAG =                        0x10000 //col:19763
VMX_PENDING_DEBUG_EXCEPTIONS_RTM_MASK =                        0x01 //col:19764
VMX_PENDING_DEBUG_EXCEPTIONS_RTM(_) =                          (((_) >> 16) & 0x01) //col:19765
VMX_VMEXIT_REASON_BASIC_EXIT_REASON_BIT =                      0 //col:19792
VMX_VMEXIT_REASON_BASIC_EXIT_REASON_FLAG =                     0xFFFF //col:19793
VMX_VMEXIT_REASON_BASIC_EXIT_REASON_MASK =                     0xFFFF //col:19794
VMX_VMEXIT_REASON_BASIC_EXIT_REASON(_) =                       (((_) >> 0) & 0xFFFF) //col:19795
VMX_VMEXIT_REASON_ALWAYS0_BIT =                                16 //col:19801
VMX_VMEXIT_REASON_ALWAYS0_FLAG =                               0x10000 //col:19802
VMX_VMEXIT_REASON_ALWAYS0_MASK =                               0x01 //col:19803
VMX_VMEXIT_REASON_ALWAYS0(_) =                                 (((_) >> 16) & 0x01) //col:19804
VMX_VMEXIT_REASON_RESERVED1_BIT =                              17 //col:19806
VMX_VMEXIT_REASON_RESERVED1_FLAG =                             0x7FE0000 //col:19807
VMX_VMEXIT_REASON_RESERVED1_MASK =                             0x3FF //col:19808
VMX_VMEXIT_REASON_RESERVED1(_) =                               (((_) >> 17) & 0x3FF) //col:19809
VMX_VMEXIT_REASON_ENCLAVE_MODE_BIT =                           27 //col:19815
VMX_VMEXIT_REASON_ENCLAVE_MODE_FLAG =                          0x8000000 //col:19816
VMX_VMEXIT_REASON_ENCLAVE_MODE_MASK =                          0x01 //col:19817
VMX_VMEXIT_REASON_ENCLAVE_MODE(_) =                            (((_) >> 27) & 0x01) //col:19818
VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_BIT =                    28 //col:19824
VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_FLAG =                   0x10000000 //col:19825
VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_MASK =                   0x01 //col:19826
VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT(_) =                     (((_) >> 28) & 0x01) //col:19827
VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_BIT =                  29 //col:19833
VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_FLAG =                 0x20000000 //col:19834
VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_MASK =                 0x01 //col:19835
VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT(_) =                   (((_) >> 29) & 0x01) //col:19836
VMX_VMEXIT_REASON_RESERVED2_BIT =                              30 //col:19838
VMX_VMEXIT_REASON_RESERVED2_FLAG =                             0x40000000 //col:19839
VMX_VMEXIT_REASON_RESERVED2_MASK =                             0x01 //col:19840
VMX_VMEXIT_REASON_RESERVED2(_) =                               (((_) >> 30) & 0x01) //col:19841
VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_BIT =                       31 //col:19849
VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_FLAG =                      0x80000000 //col:19850
VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_MASK =                      0x01 //col:19851
VMX_VMEXIT_REASON_VM_ENTRY_FAILURE(_) =                        (((_) >> 31) & 0x01) //col:19852
IO_BITMAP_A_MIN =                                              0x00000000 //col:19860
IO_BITMAP_A_MAX =                                              0x00007FFF //col:19861
IO_BITMAP_B_MIN =                                              0x00008000 //col:19862
IO_BITMAP_B_MAX =                                              0x0000FFFF //col:19863
MSR_ID_LOW_MIN =                                               0x00000000 //col:19870
MSR_ID_LOW_MAX =                                               0x00001FFF //col:19871
MSR_ID_HIGH_MIN =                                              0xC0000000 //col:19872
MSR_ID_HIGH_MAX =                                              0xC0001FFF //col:19873
EPT_POINTER_MEMORY_TYPE_BIT =                                  0 //col:19914
EPT_POINTER_MEMORY_TYPE_FLAG =                                 0x07 //col:19915
EPT_POINTER_MEMORY_TYPE_MASK =                                 0x07 //col:19916
EPT_POINTER_MEMORY_TYPE(_) =                                   (((_) >> 0) & 0x07) //col:19917
EPT_POINTER_PAGE_WALK_LENGTH_BIT =                             3 //col:19925
EPT_POINTER_PAGE_WALK_LENGTH_FLAG =                            0x38 //col:19926
EPT_POINTER_PAGE_WALK_LENGTH_MASK =                            0x07 //col:19927
EPT_POINTER_PAGE_WALK_LENGTH(_) =                              (((_) >> 3) & 0x07) //col:19928
EPT_PAGE_WALK_LENGTH_4 =                                       0x00000003 //col:19929
EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_BIT =                6 //col:19937
EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_FLAG =               0x40 //col:19938
EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_MASK =               0x01 //col:19939
EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS(_) =                 (((_) >> 6) & 0x01) //col:19940
EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_BIT =         7 //col:19948
EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_FLAG =        0x80 //col:19949
EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_MASK =        0x01 //col:19950
EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES(_) =          (((_) >> 7) & 0x01) //col:19951
EPT_POINTER_PAGE_FRAME_NUMBER_BIT =                            12 //col:19958
EPT_POINTER_PAGE_FRAME_NUMBER_FLAG =                           0xFFFFFFFFF000 //col:19959
EPT_POINTER_PAGE_FRAME_NUMBER_MASK =                           0xFFFFFFFFF //col:19960
EPT_POINTER_PAGE_FRAME_NUMBER(_) =                             (((_) >> 12) & 0xFFFFFFFFF) //col:19961
EPT_PML4E_READ_ACCESS_BIT =                                    0 //col:19991
EPT_PML4E_READ_ACCESS_FLAG =                                   0x01 //col:19992
EPT_PML4E_READ_ACCESS_MASK =                                   0x01 //col:19993
EPT_PML4E_READ_ACCESS(_) =                                     (((_) >> 0) & 0x01) //col:19994
EPT_PML4E_WRITE_ACCESS_BIT =                                   1 //col:20000
EPT_PML4E_WRITE_ACCESS_FLAG =                                  0x02 //col:20001
EPT_PML4E_WRITE_ACCESS_MASK =                                  0x01 //col:20002
EPT_PML4E_WRITE_ACCESS(_) =                                    (((_) >> 1) & 0x01) //col:20003
EPT_PML4E_EXECUTE_ACCESS_BIT =                                 2 //col:20012
EPT_PML4E_EXECUTE_ACCESS_FLAG =                                0x04 //col:20013
EPT_PML4E_EXECUTE_ACCESS_MASK =                                0x01 //col:20014
EPT_PML4E_EXECUTE_ACCESS(_) =                                  (((_) >> 2) & 0x01) //col:20015
EPT_PML4E_ACCESSED_BIT =                                       8 //col:20025
EPT_PML4E_ACCESSED_FLAG =                                      0x100 //col:20026
EPT_PML4E_ACCESSED_MASK =                                      0x01 //col:20027
EPT_PML4E_ACCESSED(_) =                                        (((_) >> 8) & 0x01) //col:20028
EPT_PML4E_USER_MODE_EXECUTE_BIT =                              10 //col:20037
EPT_PML4E_USER_MODE_EXECUTE_FLAG =                             0x400 //col:20038
EPT_PML4E_USER_MODE_EXECUTE_MASK =                             0x01 //col:20039
EPT_PML4E_USER_MODE_EXECUTE(_) =                               (((_) >> 10) & 0x01) //col:20040
EPT_PML4E_PAGE_FRAME_NUMBER_BIT =                              12 //col:20047
EPT_PML4E_PAGE_FRAME_NUMBER_FLAG =                             0xFFFFFFFFF000 //col:20048
EPT_PML4E_PAGE_FRAME_NUMBER_MASK =                             0xFFFFFFFFF //col:20049
EPT_PML4E_PAGE_FRAME_NUMBER(_) =                               (((_) >> 12) & 0xFFFFFFFFF) //col:20050
EPT_PDPTE_1GB_READ_ACCESS_BIT =                                0 //col:20068
EPT_PDPTE_1GB_READ_ACCESS_FLAG =                               0x01 //col:20069
EPT_PDPTE_1GB_READ_ACCESS_MASK =                               0x01 //col:20070
EPT_PDPTE_1GB_READ_ACCESS(_) =                                 (((_) >> 0) & 0x01) //col:20071
EPT_PDPTE_1GB_WRITE_ACCESS_BIT =                               1 //col:20077
EPT_PDPTE_1GB_WRITE_ACCESS_FLAG =                              0x02 //col:20078
EPT_PDPTE_1GB_WRITE_ACCESS_MASK =                              0x01 //col:20079
EPT_PDPTE_1GB_WRITE_ACCESS(_) =                                (((_) >> 1) & 0x01) //col:20080
EPT_PDPTE_1GB_EXECUTE_ACCESS_BIT =                             2 //col:20089
EPT_PDPTE_1GB_EXECUTE_ACCESS_FLAG =                            0x04 //col:20090
EPT_PDPTE_1GB_EXECUTE_ACCESS_MASK =                            0x01 //col:20091
EPT_PDPTE_1GB_EXECUTE_ACCESS(_) =                              (((_) >> 2) & 0x01) //col:20092
EPT_PDPTE_1GB_MEMORY_TYPE_BIT =                                3 //col:20100
EPT_PDPTE_1GB_MEMORY_TYPE_FLAG =                               0x38 //col:20101
EPT_PDPTE_1GB_MEMORY_TYPE_MASK =                               0x07 //col:20102
EPT_PDPTE_1GB_MEMORY_TYPE(_) =                                 (((_) >> 3) & 0x07) //col:20103
EPT_PDPTE_1GB_IGNORE_PAT_BIT =                                 6 //col:20111
EPT_PDPTE_1GB_IGNORE_PAT_FLAG =                                0x40 //col:20112
EPT_PDPTE_1GB_IGNORE_PAT_MASK =                                0x01 //col:20113
EPT_PDPTE_1GB_IGNORE_PAT(_) =                                  (((_) >> 6) & 0x01) //col:20114
EPT_PDPTE_1GB_LARGE_PAGE_BIT =                                 7 //col:20120
EPT_PDPTE_1GB_LARGE_PAGE_FLAG =                                0x80 //col:20121
EPT_PDPTE_1GB_LARGE_PAGE_MASK =                                0x01 //col:20122
EPT_PDPTE_1GB_LARGE_PAGE(_) =                                  (((_) >> 7) & 0x01) //col:20123
EPT_PDPTE_1GB_ACCESSED_BIT =                                   8 //col:20132
EPT_PDPTE_1GB_ACCESSED_FLAG =                                  0x100 //col:20133
EPT_PDPTE_1GB_ACCESSED_MASK =                                  0x01 //col:20134
EPT_PDPTE_1GB_ACCESSED(_) =                                    (((_) >> 8) & 0x01) //col:20135
EPT_PDPTE_1GB_DIRTY_BIT =                                      9 //col:20144
EPT_PDPTE_1GB_DIRTY_FLAG =                                     0x200 //col:20145
EPT_PDPTE_1GB_DIRTY_MASK =                                     0x01 //col:20146
EPT_PDPTE_1GB_DIRTY(_) =                                       (((_) >> 9) & 0x01) //col:20147
EPT_PDPTE_1GB_USER_MODE_EXECUTE_BIT =                          10 //col:20155
EPT_PDPTE_1GB_USER_MODE_EXECUTE_FLAG =                         0x400 //col:20156
EPT_PDPTE_1GB_USER_MODE_EXECUTE_MASK =                         0x01 //col:20157
EPT_PDPTE_1GB_USER_MODE_EXECUTE(_) =                           (((_) >> 10) & 0x01) //col:20158
EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_BIT =                          30 //col:20165
EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_FLAG =                         0xFFFFC0000000 //col:20166
EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_MASK =                         0x3FFFF //col:20167
EPT_PDPTE_1GB_PAGE_FRAME_NUMBER(_) =                           (((_) >> 30) & 0x3FFFF) //col:20168
EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_BIT =                        57 //col:20179
EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_FLAG =                       0x200000000000000 //col:20180
EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_MASK =                       0x01 //col:20181
EPT_PDPTE_1GB_VERIFY_GUEST_PAGING(_) =                         (((_) >> 57) & 0x01) //col:20182
EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_BIT =                        58 //col:20191
EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_FLAG =                       0x400000000000000 //col:20192
EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_MASK =                       0x01 //col:20193
EPT_PDPTE_1GB_PAGING_WRITE_ACCESS(_) =                         (((_) >> 58) & 0x01) //col:20194
EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_BIT =                    60 //col:20204
EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_FLAG =                   0x1000000000000000 //col:20205
EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_MASK =                   0x01 //col:20206
EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK(_) =                     (((_) >> 60) & 0x01) //col:20207
EPT_PDPTE_1GB_SUPPRESS_VE_BIT =                                63 //col:20218
EPT_PDPTE_1GB_SUPPRESS_VE_FLAG =                               0x8000000000000000 //col:20219
EPT_PDPTE_1GB_SUPPRESS_VE_MASK =                               0x01 //col:20220
EPT_PDPTE_1GB_SUPPRESS_VE(_) =                                 (((_) >> 63) & 0x01) //col:20221
EPT_PDPTE_READ_ACCESS_BIT =                                    0 //col:20238
EPT_PDPTE_READ_ACCESS_FLAG =                                   0x01 //col:20239
EPT_PDPTE_READ_ACCESS_MASK =                                   0x01 //col:20240
EPT_PDPTE_READ_ACCESS(_) =                                     (((_) >> 0) & 0x01) //col:20241
EPT_PDPTE_WRITE_ACCESS_BIT =                                   1 //col:20247
EPT_PDPTE_WRITE_ACCESS_FLAG =                                  0x02 //col:20248
EPT_PDPTE_WRITE_ACCESS_MASK =                                  0x01 //col:20249
EPT_PDPTE_WRITE_ACCESS(_) =                                    (((_) >> 1) & 0x01) //col:20250
EPT_PDPTE_EXECUTE_ACCESS_BIT =                                 2 //col:20259
EPT_PDPTE_EXECUTE_ACCESS_FLAG =                                0x04 //col:20260
EPT_PDPTE_EXECUTE_ACCESS_MASK =                                0x01 //col:20261
EPT_PDPTE_EXECUTE_ACCESS(_) =                                  (((_) >> 2) & 0x01) //col:20262
EPT_PDPTE_ACCESSED_BIT =                                       8 //col:20272
EPT_PDPTE_ACCESSED_FLAG =                                      0x100 //col:20273
EPT_PDPTE_ACCESSED_MASK =                                      0x01 //col:20274
EPT_PDPTE_ACCESSED(_) =                                        (((_) >> 8) & 0x01) //col:20275
EPT_PDPTE_USER_MODE_EXECUTE_BIT =                              10 //col:20284
EPT_PDPTE_USER_MODE_EXECUTE_FLAG =                             0x400 //col:20285
EPT_PDPTE_USER_MODE_EXECUTE_MASK =                             0x01 //col:20286
EPT_PDPTE_USER_MODE_EXECUTE(_) =                               (((_) >> 10) & 0x01) //col:20287
EPT_PDPTE_PAGE_FRAME_NUMBER_BIT =                              12 //col:20294
EPT_PDPTE_PAGE_FRAME_NUMBER_FLAG =                             0xFFFFFFFFF000 //col:20295
EPT_PDPTE_PAGE_FRAME_NUMBER_MASK =                             0xFFFFFFFFF //col:20296
EPT_PDPTE_PAGE_FRAME_NUMBER(_) =                               (((_) >> 12) & 0xFFFFFFFFF) //col:20297
EPT_PDE_2MB_READ_ACCESS_BIT =                                  0 //col:20315
EPT_PDE_2MB_READ_ACCESS_FLAG =                                 0x01 //col:20316
EPT_PDE_2MB_READ_ACCESS_MASK =                                 0x01 //col:20317
EPT_PDE_2MB_READ_ACCESS(_) =                                   (((_) >> 0) & 0x01) //col:20318
EPT_PDE_2MB_WRITE_ACCESS_BIT =                                 1 //col:20324
EPT_PDE_2MB_WRITE_ACCESS_FLAG =                                0x02 //col:20325
EPT_PDE_2MB_WRITE_ACCESS_MASK =                                0x01 //col:20326
EPT_PDE_2MB_WRITE_ACCESS(_) =                                  (((_) >> 1) & 0x01) //col:20327
EPT_PDE_2MB_EXECUTE_ACCESS_BIT =                               2 //col:20336
EPT_PDE_2MB_EXECUTE_ACCESS_FLAG =                              0x04 //col:20337
EPT_PDE_2MB_EXECUTE_ACCESS_MASK =                              0x01 //col:20338
EPT_PDE_2MB_EXECUTE_ACCESS(_) =                                (((_) >> 2) & 0x01) //col:20339
EPT_PDE_2MB_MEMORY_TYPE_BIT =                                  3 //col:20347
EPT_PDE_2MB_MEMORY_TYPE_FLAG =                                 0x38 //col:20348
EPT_PDE_2MB_MEMORY_TYPE_MASK =                                 0x07 //col:20349
EPT_PDE_2MB_MEMORY_TYPE(_) =                                   (((_) >> 3) & 0x07) //col:20350
EPT_PDE_2MB_IGNORE_PAT_BIT =                                   6 //col:20358
EPT_PDE_2MB_IGNORE_PAT_FLAG =                                  0x40 //col:20359
EPT_PDE_2MB_IGNORE_PAT_MASK =                                  0x01 //col:20360
EPT_PDE_2MB_IGNORE_PAT(_) =                                    (((_) >> 6) & 0x01) //col:20361
EPT_PDE_2MB_LARGE_PAGE_BIT =                                   7 //col:20367
EPT_PDE_2MB_LARGE_PAGE_FLAG =                                  0x80 //col:20368
EPT_PDE_2MB_LARGE_PAGE_MASK =                                  0x01 //col:20369
EPT_PDE_2MB_LARGE_PAGE(_) =                                    (((_) >> 7) & 0x01) //col:20370
EPT_PDE_2MB_ACCESSED_BIT =                                     8 //col:20379
EPT_PDE_2MB_ACCESSED_FLAG =                                    0x100 //col:20380
EPT_PDE_2MB_ACCESSED_MASK =                                    0x01 //col:20381
EPT_PDE_2MB_ACCESSED(_) =                                      (((_) >> 8) & 0x01) //col:20382
EPT_PDE_2MB_DIRTY_BIT =                                        9 //col:20391
EPT_PDE_2MB_DIRTY_FLAG =                                       0x200 //col:20392
EPT_PDE_2MB_DIRTY_MASK =                                       0x01 //col:20393
EPT_PDE_2MB_DIRTY(_) =                                         (((_) >> 9) & 0x01) //col:20394
EPT_PDE_2MB_USER_MODE_EXECUTE_BIT =                            10 //col:20402
EPT_PDE_2MB_USER_MODE_EXECUTE_FLAG =                           0x400 //col:20403
EPT_PDE_2MB_USER_MODE_EXECUTE_MASK =                           0x01 //col:20404
EPT_PDE_2MB_USER_MODE_EXECUTE(_) =                             (((_) >> 10) & 0x01) //col:20405
EPT_PDE_2MB_PAGE_FRAME_NUMBER_BIT =                            21 //col:20412
EPT_PDE_2MB_PAGE_FRAME_NUMBER_FLAG =                           0xFFFFFFE00000 //col:20413
EPT_PDE_2MB_PAGE_FRAME_NUMBER_MASK =                           0x7FFFFFF //col:20414
EPT_PDE_2MB_PAGE_FRAME_NUMBER(_) =                             (((_) >> 21) & 0x7FFFFFF) //col:20415
EPT_PDE_2MB_VERIFY_GUEST_PAGING_BIT =                          57 //col:20426
EPT_PDE_2MB_VERIFY_GUEST_PAGING_FLAG =                         0x200000000000000 //col:20427
EPT_PDE_2MB_VERIFY_GUEST_PAGING_MASK =                         0x01 //col:20428
EPT_PDE_2MB_VERIFY_GUEST_PAGING(_) =                           (((_) >> 57) & 0x01) //col:20429
EPT_PDE_2MB_PAGING_WRITE_ACCESS_BIT =                          58 //col:20438
EPT_PDE_2MB_PAGING_WRITE_ACCESS_FLAG =                         0x400000000000000 //col:20439
EPT_PDE_2MB_PAGING_WRITE_ACCESS_MASK =                         0x01 //col:20440
EPT_PDE_2MB_PAGING_WRITE_ACCESS(_) =                           (((_) >> 58) & 0x01) //col:20441
EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_BIT =                      60 //col:20451
EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_FLAG =                     0x1000000000000000 //col:20452
EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_MASK =                     0x01 //col:20453
EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK(_) =                       (((_) >> 60) & 0x01) //col:20454
EPT_PDE_2MB_SUPPRESS_VE_BIT =                                  63 //col:20465
EPT_PDE_2MB_SUPPRESS_VE_FLAG =                                 0x8000000000000000 //col:20466
EPT_PDE_2MB_SUPPRESS_VE_MASK =                                 0x01 //col:20467
EPT_PDE_2MB_SUPPRESS_VE(_) =                                   (((_) >> 63) & 0x01) //col:20468
EPT_PDE_READ_ACCESS_BIT =                                      0 //col:20485
EPT_PDE_READ_ACCESS_FLAG =                                     0x01 //col:20486
EPT_PDE_READ_ACCESS_MASK =                                     0x01 //col:20487
EPT_PDE_READ_ACCESS(_) =                                       (((_) >> 0) & 0x01) //col:20488
EPT_PDE_WRITE_ACCESS_BIT =                                     1 //col:20494
EPT_PDE_WRITE_ACCESS_FLAG =                                    0x02 //col:20495
EPT_PDE_WRITE_ACCESS_MASK =                                    0x01 //col:20496
EPT_PDE_WRITE_ACCESS(_) =                                      (((_) >> 1) & 0x01) //col:20497
EPT_PDE_EXECUTE_ACCESS_BIT =                                   2 //col:20506
EPT_PDE_EXECUTE_ACCESS_FLAG =                                  0x04 //col:20507
EPT_PDE_EXECUTE_ACCESS_MASK =                                  0x01 //col:20508
EPT_PDE_EXECUTE_ACCESS(_) =                                    (((_) >> 2) & 0x01) //col:20509
EPT_PDE_ACCESSED_BIT =                                         8 //col:20519
EPT_PDE_ACCESSED_FLAG =                                        0x100 //col:20520
EPT_PDE_ACCESSED_MASK =                                        0x01 //col:20521
EPT_PDE_ACCESSED(_) =                                          (((_) >> 8) & 0x01) //col:20522
EPT_PDE_USER_MODE_EXECUTE_BIT =                                10 //col:20531
EPT_PDE_USER_MODE_EXECUTE_FLAG =                               0x400 //col:20532
EPT_PDE_USER_MODE_EXECUTE_MASK =                               0x01 //col:20533
EPT_PDE_USER_MODE_EXECUTE(_) =                                 (((_) >> 10) & 0x01) //col:20534
EPT_PDE_PAGE_FRAME_NUMBER_BIT =                                12 //col:20541
EPT_PDE_PAGE_FRAME_NUMBER_FLAG =                               0xFFFFFFFFF000 //col:20542
EPT_PDE_PAGE_FRAME_NUMBER_MASK =                               0xFFFFFFFFF //col:20543
EPT_PDE_PAGE_FRAME_NUMBER(_) =                                 (((_) >> 12) & 0xFFFFFFFFF) //col:20544
EPT_PTE_READ_ACCESS_BIT =                                      0 //col:20562
EPT_PTE_READ_ACCESS_FLAG =                                     0x01 //col:20563
EPT_PTE_READ_ACCESS_MASK =                                     0x01 //col:20564
EPT_PTE_READ_ACCESS(_) =                                       (((_) >> 0) & 0x01) //col:20565
EPT_PTE_WRITE_ACCESS_BIT =                                     1 //col:20571
EPT_PTE_WRITE_ACCESS_FLAG =                                    0x02 //col:20572
EPT_PTE_WRITE_ACCESS_MASK =                                    0x01 //col:20573
EPT_PTE_WRITE_ACCESS(_) =                                      (((_) >> 1) & 0x01) //col:20574
EPT_PTE_EXECUTE_ACCESS_BIT =                                   2 //col:20583
EPT_PTE_EXECUTE_ACCESS_FLAG =                                  0x04 //col:20584
EPT_PTE_EXECUTE_ACCESS_MASK =                                  0x01 //col:20585
EPT_PTE_EXECUTE_ACCESS(_) =                                    (((_) >> 2) & 0x01) //col:20586
EPT_PTE_MEMORY_TYPE_BIT =                                      3 //col:20594
EPT_PTE_MEMORY_TYPE_FLAG =                                     0x38 //col:20595
EPT_PTE_MEMORY_TYPE_MASK =                                     0x07 //col:20596
EPT_PTE_MEMORY_TYPE(_) =                                       (((_) >> 3) & 0x07) //col:20597
EPT_PTE_IGNORE_PAT_BIT =                                       6 //col:20605
EPT_PTE_IGNORE_PAT_FLAG =                                      0x40 //col:20606
EPT_PTE_IGNORE_PAT_MASK =                                      0x01 //col:20607
EPT_PTE_IGNORE_PAT(_) =                                        (((_) >> 6) & 0x01) //col:20608
EPT_PTE_ACCESSED_BIT =                                         8 //col:20618
EPT_PTE_ACCESSED_FLAG =                                        0x100 //col:20619
EPT_PTE_ACCESSED_MASK =                                        0x01 //col:20620
EPT_PTE_ACCESSED(_) =                                          (((_) >> 8) & 0x01) //col:20621
EPT_PTE_DIRTY_BIT =                                            9 //col:20630
EPT_PTE_DIRTY_FLAG =                                           0x200 //col:20631
EPT_PTE_DIRTY_MASK =                                           0x01 //col:20632
EPT_PTE_DIRTY(_) =                                             (((_) >> 9) & 0x01) //col:20633
EPT_PTE_USER_MODE_EXECUTE_BIT =                                10 //col:20641
EPT_PTE_USER_MODE_EXECUTE_FLAG =                               0x400 //col:20642
EPT_PTE_USER_MODE_EXECUTE_MASK =                               0x01 //col:20643
EPT_PTE_USER_MODE_EXECUTE(_) =                                 (((_) >> 10) & 0x01) //col:20644
EPT_PTE_PAGE_FRAME_NUMBER_BIT =                                12 //col:20651
EPT_PTE_PAGE_FRAME_NUMBER_FLAG =                               0xFFFFFFFFF000 //col:20652
EPT_PTE_PAGE_FRAME_NUMBER_MASK =                               0xFFFFFFFFF //col:20653
EPT_PTE_PAGE_FRAME_NUMBER(_) =                                 (((_) >> 12) & 0xFFFFFFFFF) //col:20654
EPT_PTE_VERIFY_GUEST_PAGING_BIT =                              57 //col:20665
EPT_PTE_VERIFY_GUEST_PAGING_FLAG =                             0x200000000000000 //col:20666
EPT_PTE_VERIFY_GUEST_PAGING_MASK =                             0x01 //col:20667
EPT_PTE_VERIFY_GUEST_PAGING(_) =                               (((_) >> 57) & 0x01) //col:20668
EPT_PTE_PAGING_WRITE_ACCESS_BIT =                              58 //col:20677
EPT_PTE_PAGING_WRITE_ACCESS_FLAG =                             0x400000000000000 //col:20678
EPT_PTE_PAGING_WRITE_ACCESS_MASK =                             0x01 //col:20679
EPT_PTE_PAGING_WRITE_ACCESS(_) =                               (((_) >> 58) & 0x01) //col:20680
EPT_PTE_SUPERVISOR_SHADOW_STACK_BIT =                          60 //col:20690
EPT_PTE_SUPERVISOR_SHADOW_STACK_FLAG =                         0x1000000000000000 //col:20691
EPT_PTE_SUPERVISOR_SHADOW_STACK_MASK =                         0x01 //col:20692
EPT_PTE_SUPERVISOR_SHADOW_STACK(_) =                           (((_) >> 60) & 0x01) //col:20693
EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_BIT =                       61 //col:20704
EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_FLAG =                      0x2000000000000000 //col:20705
EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_MASK =                      0x01 //col:20706
EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS(_) =                        (((_) >> 61) & 0x01) //col:20707
EPT_PTE_SUPPRESS_VE_BIT =                                      63 //col:20718
EPT_PTE_SUPPRESS_VE_FLAG =                                     0x8000000000000000 //col:20719
EPT_PTE_SUPPRESS_VE_MASK =                                     0x01 //col:20720
EPT_PTE_SUPPRESS_VE(_) =                                       (((_) >> 63) & 0x01) //col:20721
EPT_ENTRY_READ_ACCESS_BIT =                                    0 //col:20735
EPT_ENTRY_READ_ACCESS_FLAG =                                   0x01 //col:20736
EPT_ENTRY_READ_ACCESS_MASK =                                   0x01 //col:20737
EPT_ENTRY_READ_ACCESS(_) =                                     (((_) >> 0) & 0x01) //col:20738
EPT_ENTRY_WRITE_ACCESS_BIT =                                   1 //col:20740
EPT_ENTRY_WRITE_ACCESS_FLAG =                                  0x02 //col:20741
EPT_ENTRY_WRITE_ACCESS_MASK =                                  0x01 //col:20742
EPT_ENTRY_WRITE_ACCESS(_) =                                    (((_) >> 1) & 0x01) //col:20743
EPT_ENTRY_EXECUTE_ACCESS_BIT =                                 2 //col:20745
EPT_ENTRY_EXECUTE_ACCESS_FLAG =                                0x04 //col:20746
EPT_ENTRY_EXECUTE_ACCESS_MASK =                                0x01 //col:20747
EPT_ENTRY_EXECUTE_ACCESS(_) =                                  (((_) >> 2) & 0x01) //col:20748
EPT_ENTRY_MEMORY_TYPE_BIT =                                    3 //col:20750
EPT_ENTRY_MEMORY_TYPE_FLAG =                                   0x38 //col:20751
EPT_ENTRY_MEMORY_TYPE_MASK =                                   0x07 //col:20752
EPT_ENTRY_MEMORY_TYPE(_) =                                     (((_) >> 3) & 0x07) //col:20753
EPT_ENTRY_IGNORE_PAT_BIT =                                     6 //col:20755
EPT_ENTRY_IGNORE_PAT_FLAG =                                    0x40 //col:20756
EPT_ENTRY_IGNORE_PAT_MASK =                                    0x01 //col:20757
EPT_ENTRY_IGNORE_PAT(_) =                                      (((_) >> 6) & 0x01) //col:20758
EPT_ENTRY_LARGE_PAGE_BIT =                                     7 //col:20760
EPT_ENTRY_LARGE_PAGE_FLAG =                                    0x80 //col:20761
EPT_ENTRY_LARGE_PAGE_MASK =                                    0x01 //col:20762
EPT_ENTRY_LARGE_PAGE(_) =                                      (((_) >> 7) & 0x01) //col:20763
EPT_ENTRY_ACCESSED_BIT =                                       8 //col:20765
EPT_ENTRY_ACCESSED_FLAG =                                      0x100 //col:20766
EPT_ENTRY_ACCESSED_MASK =                                      0x01 //col:20767
EPT_ENTRY_ACCESSED(_) =                                        (((_) >> 8) & 0x01) //col:20768
EPT_ENTRY_DIRTY_BIT =                                          9 //col:20770
EPT_ENTRY_DIRTY_FLAG =                                         0x200 //col:20771
EPT_ENTRY_DIRTY_MASK =                                         0x01 //col:20772
EPT_ENTRY_DIRTY(_) =                                           (((_) >> 9) & 0x01) //col:20773
EPT_ENTRY_USER_MODE_EXECUTE_BIT =                              10 //col:20775
EPT_ENTRY_USER_MODE_EXECUTE_FLAG =                             0x400 //col:20776
EPT_ENTRY_USER_MODE_EXECUTE_MASK =                             0x01 //col:20777
EPT_ENTRY_USER_MODE_EXECUTE(_) =                               (((_) >> 10) & 0x01) //col:20778
EPT_ENTRY_PAGE_FRAME_NUMBER_BIT =                              12 //col:20781
EPT_ENTRY_PAGE_FRAME_NUMBER_FLAG =                             0xFFFFFFFFF000 //col:20782
EPT_ENTRY_PAGE_FRAME_NUMBER_MASK =                             0xFFFFFFFFF //col:20783
EPT_ENTRY_PAGE_FRAME_NUMBER(_) =                               (((_) >> 12) & 0xFFFFFFFFF) //col:20784
EPT_ENTRY_SUPPRESS_VE_BIT =                                    63 //col:20787
EPT_ENTRY_SUPPRESS_VE_FLAG =                                   0x8000000000000000 //col:20788
EPT_ENTRY_SUPPRESS_VE_MASK =                                   0x01 //col:20789
EPT_ENTRY_SUPPRESS_VE(_) =                                     (((_) >> 63) & 0x01) //col:20790
EPT_LEVEL_PML4E =                                              0x00000003 //col:20803
EPT_LEVEL_PDPTE =                                              0x00000002 //col:20804
EPT_LEVEL_PDE =                                                0x00000001 //col:20805
EPT_LEVEL_PTE =                                                0x00000000 //col:20806
EPT_PML4E_ENTRY_COUNT =                                        0x00000200 //col:20818
EPT_PDPTE_ENTRY_COUNT =                                        0x00000200 //col:20819
EPT_PDE_ENTRY_COUNT =                                          0x00000200 //col:20820
EPT_PTE_ENTRY_COUNT =                                          0x00000200 //col:20821
HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_BIT =                    3 //col:20929
HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_FLAG =                   0x08 //col:20930
HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_MASK =                   0x01 //col:20931
HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH(_) =                     (((_) >> 3) & 0x01) //col:20932
HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_BIT =                    4 //col:20939
HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_FLAG =                   0x10 //col:20940
HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_MASK =                   0x01 //col:20941
HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE(_) =                     (((_) >> 4) & 0x01) //col:20942
HLAT_POINTER_PAGE_FRAME_NUMBER_BIT =                           12 //col:20949
HLAT_POINTER_PAGE_FRAME_NUMBER_FLAG =                          0xFFFFFFFFF000 //col:20950
HLAT_POINTER_PAGE_FRAME_NUMBER_MASK =                          0xFFFFFFFFF //col:20951
HLAT_POINTER_PAGE_FRAME_NUMBER(_) =                            (((_) >> 12) & 0xFFFFFFFFF) //col:20952
VMCS_COMPONENT_ENCODING_ACCESS_TYPE_BIT =                      0 //col:21095
VMCS_COMPONENT_ENCODING_ACCESS_TYPE_FLAG =                     0x01 //col:21096
VMCS_COMPONENT_ENCODING_ACCESS_TYPE_MASK =                     0x01 //col:21097
VMCS_COMPONENT_ENCODING_ACCESS_TYPE(_) =                       (((_) >> 0) & 0x01) //col:21098
VMCS_COMPONENT_ENCODING_INDEX_BIT =                            1 //col:21104
VMCS_COMPONENT_ENCODING_INDEX_FLAG =                           0x3FE //col:21105
VMCS_COMPONENT_ENCODING_INDEX_MASK =                           0x1FF //col:21106
VMCS_COMPONENT_ENCODING_INDEX(_) =                             (((_) >> 1) & 0x1FF) //col:21107
VMCS_COMPONENT_ENCODING_TYPE_BIT =                             10 //col:21117
VMCS_COMPONENT_ENCODING_TYPE_FLAG =                            0xC00 //col:21118
VMCS_COMPONENT_ENCODING_TYPE_MASK =                            0x03 //col:21119
VMCS_COMPONENT_ENCODING_TYPE(_) =                              (((_) >> 10) & 0x03) //col:21120
VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_BIT =                     12 //col:21126
VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_FLAG =                    0x1000 //col:21127
VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_MASK =                    0x01 //col:21128
VMCS_COMPONENT_ENCODING_MUST_BE_ZERO(_) =                      (((_) >> 12) & 0x01) //col:21129
VMCS_COMPONENT_ENCODING_WIDTH_BIT =                            13 //col:21139
VMCS_COMPONENT_ENCODING_WIDTH_FLAG =                           0x6000 //col:21140
VMCS_COMPONENT_ENCODING_WIDTH_MASK =                           0x03 //col:21141
VMCS_COMPONENT_ENCODING_WIDTH(_) =                             (((_) >> 13) & 0x03) //col:21142
VMCS_CTRL_VIRTUAL_PROCESSOR_IDENTIFIER =                       0x00000000 //col:21170
VMCS_CTRL_POSTED_INTERRUPT_NOTIFICATION_VECTOR =               0x00000002 //col:21178
VMCS_CTRL_EPTP_INDEX =                                         0x00000004 //col:21186
VMCS_CTRL_HLAT_PREFIX_SIZE =                                   0x00000006 //col:21193
VMCS_GUEST_ES_SELECTOR =                                       0x00000800 //col:21208
VMCS_GUEST_CS_SELECTOR =                                       0x00000802 //col:21213
VMCS_GUEST_SS_SELECTOR =                                       0x00000804 //col:21218
VMCS_GUEST_DS_SELECTOR =                                       0x00000806 //col:21223
VMCS_GUEST_FS_SELECTOR =                                       0x00000808 //col:21228
VMCS_GUEST_GS_SELECTOR =                                       0x0000080A //col:21233
VMCS_GUEST_LDTR_SELECTOR =                                     0x0000080C //col:21238
VMCS_GUEST_TR_SELECTOR =                                       0x0000080E //col:21243
VMCS_GUEST_INTERRUPT_STATUS =                                  0x00000810 //col:21251
VMCS_GUEST_PML_INDEX =                                         0x00000812 //col:21258
VMCS_HOST_ES_SELECTOR =                                        0x00000C00 //col:21273
VMCS_HOST_CS_SELECTOR =                                        0x00000C02 //col:21278
VMCS_HOST_SS_SELECTOR =                                        0x00000C04 //col:21283
VMCS_HOST_DS_SELECTOR =                                        0x00000C06 //col:21288
VMCS_HOST_FS_SELECTOR =                                        0x00000C08 //col:21293
VMCS_HOST_GS_SELECTOR =                                        0x00000C0A //col:21298
VMCS_HOST_TR_SELECTOR =                                        0x00000C0C //col:21303
VMCS_CTRL_IO_BITMAP_A_ADDRESS =                                0x00002000 //col:21331
VMCS_CTRL_IO_BITMAP_B_ADDRESS =                                0x00002002 //col:21336
VMCS_CTRL_MSR_BITMAP_ADDRESS =                                 0x00002004 //col:21341
VMCS_CTRL_VMEXIT_MSR_STORE_ADDRESS =                           0x00002006 //col:21346
VMCS_CTRL_VMEXIT_MSR_LOAD_ADDRESS =                            0x00002008 //col:21351
VMCS_CTRL_VMENTRY_MSR_LOAD_ADDRESS =                           0x0000200A //col:21356
VMCS_CTRL_EXECUTIVE_VMCS_POINTER =                             0x0000200C //col:21361
VMCS_CTRL_PML_ADDRESS =                                        0x0000200E //col:21366
VMCS_CTRL_TSC_OFFSET =                                         0x00002010 //col:21371
VMCS_CTRL_VIRTUAL_APIC_ADDRESS =                               0x00002012 //col:21376
VMCS_CTRL_APIC_ACCESS_ADDRESS =                                0x00002014 //col:21381
VMCS_CTRL_POSTED_INTERRUPT_DESCRIPTOR_ADDRESS =                0x00002016 //col:21386
VMCS_CTRL_VMFUNC_CONTROLS =                                    0x00002018 //col:21391
VMCS_CTRL_EPT_POINTER =                                        0x0000201A //col:21396
VMCS_CTRL_EOI_EXIT_BITMAP_0 =                                  0x0000201C //col:21401
VMCS_CTRL_EOI_EXIT_BITMAP_1 =                                  0x0000201E //col:21406
VMCS_CTRL_EOI_EXIT_BITMAP_2 =                                  0x00002020 //col:21411
VMCS_CTRL_EOI_EXIT_BITMAP_3 =                                  0x00002022 //col:21416
VMCS_CTRL_EPT_POINTER_LIST_ADDRESS =                           0x00002024 //col:21421
VMCS_CTRL_VMREAD_BITMAP_ADDRESS =                              0x00002026 //col:21426
VMCS_CTRL_VMWRITE_BITMAP_ADDRESS =                             0x00002028 //col:21431
VMCS_CTRL_VIRTUALIZATION_EXCEPTION_INFORMATION_ADDRESS =       0x0000202A //col:21436
VMCS_CTRL_XSS_EXITING_BITMAP =                                 0x0000202C //col:21441
VMCS_CTRL_ENCLS_EXITING_BITMAP =                               0x0000202E //col:21446
VMCS_CTRL_SUB_PAGE_PERMISSION_TABLE_POINTER =                  0x00002030 //col:21451
VMCS_CTRL_TSC_MULTIPLIER =                                     0x00002032 //col:21456
VMCS_CTRL_TERTIARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS =     0x00002034 //col:21461
VMCS_CTRL_ENCLV_EXITING_BITMAP =                               0x00002036 //col:21466
VMCS_CTRL_HLAT_POINTER =                                       0x00002040 //col:21471
VMCS_CTRL_SECONDARY_VMEXIT_CONTROLS =                          0x00002044 //col:21476
VMCS_GUEST_PHYSICAL_ADDRESS =                                  0x00002400 //col:21491
VMCS_GUEST_VMCS_LINK_POINTER =                                 0x00002800 //col:21506
VMCS_GUEST_DEBUGCTL =                                          0x00002802 //col:21511
VMCS_GUEST_PAT =                                               0x00002804 //col:21516
VMCS_GUEST_EFER =                                              0x00002806 //col:21521
VMCS_GUEST_PERF_GLOBAL_CTRL =                                  0x00002808 //col:21526
VMCS_GUEST_PDPTE0 =                                            0x0000280A //col:21531
VMCS_GUEST_PDPTE1 =                                            0x0000280C //col:21536
VMCS_GUEST_PDPTE2 =                                            0x0000280E //col:21541
VMCS_GUEST_PDPTE3 =                                            0x00002810 //col:21546
VMCS_GUEST_BNDCFGS =                                           0x00002812 //col:21551
VMCS_GUEST_RTIT_CTL =                                          0x00002814 //col:21556
VMCS_GUEST_LBR_CTL =                                           0x00002816 //col:21561
VMCS_GUEST_PKRS =                                              0x00002818 //col:21566
VMCS_HOST_PAT =                                                0x00002C00 //col:21581
VMCS_HOST_EFER =                                               0x00002C02 //col:21586
VMCS_HOST_PERF_GLOBAL_CTRL =                                   0x00002C04 //col:21591
VMCS_HOST_PKRS =                                               0x00002C06 //col:21596
VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS =                    0x00004000 //col:21624
VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS =              0x00004002 //col:21629
VMCS_CTRL_EXCEPTION_BITMAP =                                   0x00004004 //col:21634
VMCS_CTRL_PAGEFAULT_ERROR_CODE_MASK =                          0x00004006 //col:21639
VMCS_CTRL_PAGEFAULT_ERROR_CODE_MATCH =                         0x00004008 //col:21644
VMCS_CTRL_CR3_TARGET_COUNT =                                   0x0000400A //col:21649
VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS =                            0x0000400C //col:21654
VMCS_CTRL_VMEXIT_MSR_STORE_COUNT =                             0x0000400E //col:21659
VMCS_CTRL_VMEXIT_MSR_LOAD_COUNT =                              0x00004010 //col:21664
VMCS_CTRL_VMENTRY_CONTROLS =                                   0x00004012 //col:21669
VMCS_CTRL_VMENTRY_MSR_LOAD_COUNT =                             0x00004014 //col:21674
VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD =             0x00004016 //col:21679
VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE =                       0x00004018 //col:21684
VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH =                         0x0000401A //col:21689
VMCS_CTRL_TPR_THRESHOLD =                                      0x0000401C //col:21694
VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS =    0x0000401E //col:21699
VMCS_CTRL_PLE_GAP =                                            0x00004020 //col:21704
VMCS_CTRL_PLE_WINDOW =                                         0x00004022 //col:21709
VMCS_VM_INSTRUCTION_ERROR =                                    0x00004400 //col:21724
VMCS_EXIT_REASON =                                             0x00004402 //col:21729
VMCS_VMEXIT_INTERRUPTION_INFORMATION =                         0x00004404 //col:21734
VMCS_VMEXIT_INTERRUPTION_ERROR_CODE =                          0x00004406 //col:21739
VMCS_IDT_VECTORING_INFORMATION =                               0x00004408 //col:21744
VMCS_IDT_VECTORING_ERROR_CODE =                                0x0000440A //col:21749
VMCS_VMEXIT_INSTRUCTION_LENGTH =                               0x0000440C //col:21754
VMCS_VMEXIT_INSTRUCTION_INFO =                                 0x0000440E //col:21759
VMCS_GUEST_ES_LIMIT =                                          0x00004800 //col:21774
VMCS_GUEST_CS_LIMIT =                                          0x00004802 //col:21779
VMCS_GUEST_SS_LIMIT =                                          0x00004804 //col:21784
VMCS_GUEST_DS_LIMIT =                                          0x00004806 //col:21789
VMCS_GUEST_FS_LIMIT =                                          0x00004808 //col:21794
VMCS_GUEST_GS_LIMIT =                                          0x0000480A //col:21799
VMCS_GUEST_LDTR_LIMIT =                                        0x0000480C //col:21804
VMCS_GUEST_TR_LIMIT =                                          0x0000480E //col:21809
VMCS_GUEST_GDTR_LIMIT =                                        0x00004810 //col:21814
VMCS_GUEST_IDTR_LIMIT =                                        0x00004812 //col:21819
VMCS_GUEST_ES_ACCESS_RIGHTS =                                  0x00004814 //col:21824
VMCS_GUEST_CS_ACCESS_RIGHTS =                                  0x00004816 //col:21829
VMCS_GUEST_SS_ACCESS_RIGHTS =                                  0x00004818 //col:21834
VMCS_GUEST_DS_ACCESS_RIGHTS =                                  0x0000481A //col:21839
VMCS_GUEST_FS_ACCESS_RIGHTS =                                  0x0000481C //col:21844
VMCS_GUEST_GS_ACCESS_RIGHTS =                                  0x0000481E //col:21849
VMCS_GUEST_LDTR_ACCESS_RIGHTS =                                0x00004820 //col:21854
VMCS_GUEST_TR_ACCESS_RIGHTS =                                  0x00004822 //col:21859
VMCS_GUEST_INTERRUPTIBILITY_STATE =                            0x00004824 //col:21864
VMCS_GUEST_ACTIVITY_STATE =                                    0x00004826 //col:21869
VMCS_GUEST_SMBASE =                                            0x00004828 //col:21874
VMCS_GUEST_SYSENTER_CS =                                       0x0000482A //col:21879
VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE =                        0x0000482E //col:21884
VMCS_HOST_SYSENTER_CS =                                        0x00004C00 //col:21899
VMCS_CTRL_CR0_GUEST_HOST_MASK =                                0x00006000 //col:21927
VMCS_CTRL_CR4_GUEST_HOST_MASK =                                0x00006002 //col:21932
VMCS_CTRL_CR0_READ_SHADOW =                                    0x00006004 //col:21937
VMCS_CTRL_CR4_READ_SHADOW =                                    0x00006006 //col:21942
VMCS_CTRL_CR3_TARGET_VALUE_0 =                                 0x00006008 //col:21947
VMCS_CTRL_CR3_TARGET_VALUE_1 =                                 0x0000600A //col:21952
VMCS_CTRL_CR3_TARGET_VALUE_2 =                                 0x0000600C //col:21957
VMCS_CTRL_CR3_TARGET_VALUE_3 =                                 0x0000600E //col:21962
VMCS_EXIT_QUALIFICATION =                                      0x00006400 //col:21977
VMCS_IO_RCX =                                                  0x00006402 //col:21982
VMCS_IO_RSI =                                                  0x00006404 //col:21987
VMCS_IO_RDI =                                                  0x00006406 //col:21992
VMCS_IO_RIP =                                                  0x00006408 //col:21997
VMCS_EXIT_GUEST_LINEAR_ADDRESS =                               0x0000640A //col:22002
VMCS_GUEST_CR0 =                                               0x00006800 //col:22017
VMCS_GUEST_CR3 =                                               0x00006802 //col:22022
VMCS_GUEST_CR4 =                                               0x00006804 //col:22027
VMCS_GUEST_ES_BASE =                                           0x00006806 //col:22032
VMCS_GUEST_CS_BASE =                                           0x00006808 //col:22037
VMCS_GUEST_SS_BASE =                                           0x0000680A //col:22042
VMCS_GUEST_DS_BASE =                                           0x0000680C //col:22047
VMCS_GUEST_FS_BASE =                                           0x0000680E //col:22052
VMCS_GUEST_GS_BASE =                                           0x00006810 //col:22057
VMCS_GUEST_LDTR_BASE =                                         0x00006812 //col:22062
VMCS_GUEST_TR_BASE =                                           0x00006814 //col:22067
VMCS_GUEST_GDTR_BASE =                                         0x00006816 //col:22072
VMCS_GUEST_IDTR_BASE =                                         0x00006818 //col:22077
VMCS_GUEST_DR7 =                                               0x0000681A //col:22082
VMCS_GUEST_RSP =                                               0x0000681C //col:22087
VMCS_GUEST_RIP =                                               0x0000681E //col:22092
VMCS_GUEST_RFLAGS =                                            0x00006820 //col:22097
VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS =                          0x00006822 //col:22102
VMCS_GUEST_SYSENTER_ESP =                                      0x00006824 //col:22107
VMCS_GUEST_SYSENTER_EIP =                                      0x00006826 //col:22112
VMCS_GUEST_S_CET =                                             0x00006C28 //col:22117
VMCS_GUEST_SSP =                                               0x00006C2A //col:22122
VMCS_GUEST_INTERRUPT_SSP_TABLE_ADDR =                          0x00006C2C //col:22127
VMCS_HOST_CR0 =                                                0x00006C00 //col:22142
VMCS_HOST_CR3 =                                                0x00006C02 //col:22147
VMCS_HOST_CR4 =                                                0x00006C04 //col:22152
VMCS_HOST_FS_BASE =                                            0x00006C06 //col:22157
VMCS_HOST_GS_BASE =                                            0x00006C08 //col:22162
VMCS_HOST_TR_BASE =                                            0x00006C0A //col:22167
VMCS_HOST_GDTR_BASE =                                          0x00006C0C //col:22172
VMCS_HOST_IDTR_BASE =                                          0x00006C0E //col:22177
VMCS_HOST_SYSENTER_ESP =                                       0x00006C10 //col:22182
VMCS_HOST_SYSENTER_EIP =                                       0x00006C12 //col:22187
VMCS_HOST_RSP =                                                0x00006C14 //col:22192
VMCS_HOST_RIP =                                                0x00006C16 //col:22197
VMCS_HOST_S_CET =                                              0x00006C18 //col:22202
VMCS_HOST_SSP =                                                0x00006C1A //col:22207
VMCS_HOST_INTERRUPT_SSP_TABLE_ADDR =                           0x00006C1C //col:22212
VMENTRY_INTERRUPT_INFORMATION_VECTOR_BIT =                     0 //col:22282
VMENTRY_INTERRUPT_INFORMATION_VECTOR_FLAG =                    0xFF //col:22283
VMENTRY_INTERRUPT_INFORMATION_VECTOR_MASK =                    0xFF //col:22284
VMENTRY_INTERRUPT_INFORMATION_VECTOR(_) =                      (((_) >> 0) & 0xFF) //col:22285
VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_BIT =          8 //col:22293
VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_FLAG =         0x700 //col:22294
VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_MASK =         0x07 //col:22295
VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE(_) =           (((_) >> 8) & 0x07) //col:22296
VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_BIT =         11 //col:22304
VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_FLAG =        0x800 //col:22305
VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_MASK =        0x01 //col:22306
VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE(_) =          (((_) >> 11) & 0x01) //col:22307
VMENTRY_INTERRUPT_INFORMATION_VALID_BIT =                      31 //col:22317
VMENTRY_INTERRUPT_INFORMATION_VALID_FLAG =                     0x80000000 //col:22318
VMENTRY_INTERRUPT_INFORMATION_VALID_MASK =                     0x01 //col:22319
VMENTRY_INTERRUPT_INFORMATION_VALID(_) =                       (((_) >> 31) & 0x01) //col:22320
VMEXIT_INTERRUPT_INFORMATION_VECTOR_BIT =                      0 //col:22340
VMEXIT_INTERRUPT_INFORMATION_VECTOR_FLAG =                     0xFF //col:22341
VMEXIT_INTERRUPT_INFORMATION_VECTOR_MASK =                     0xFF //col:22342
VMEXIT_INTERRUPT_INFORMATION_VECTOR(_) =                       (((_) >> 0) & 0xFF) //col:22343
VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_BIT =           8 //col:22349
VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_FLAG =          0x700 //col:22350
VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_MASK =          0x07 //col:22351
VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE(_) =            (((_) >> 8) & 0x07) //col:22352
VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_BIT =            11 //col:22358
VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_FLAG =           0x800 //col:22359
VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_MASK =           0x01 //col:22360
VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID(_) =             (((_) >> 11) & 0x01) //col:22361
VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_BIT =              12 //col:22367
VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_FLAG =             0x1000 //col:22368
VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_MASK =             0x01 //col:22369
VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING(_) =               (((_) >> 12) & 0x01) //col:22370
VMEXIT_INTERRUPT_INFORMATION_VALID_BIT =                       31 //col:22377
VMEXIT_INTERRUPT_INFORMATION_VALID_FLAG =                      0x80000000 //col:22378
VMEXIT_INTERRUPT_INFORMATION_VALID_MASK =                      0x01 //col:22379
VMEXIT_INTERRUPT_INFORMATION_VALID(_) =                        (((_) >> 31) & 0x01) //col:22380
APIC_BASE_ADDRESS =                                            0xFEE00000 //col:22413
APIC_ID =                                                      0x00000020 //col:22418
APIC_VERSION =                                                 0x00000030 //col:22423
APIC_TASK_PRIORITY =                                           0x00000080 //col:22428
APIC_ARBITRATION_PRIORITY =                                    0x00000090 //col:22433
APIC_PROCESSOR_PRIORITY =                                      0x000000A0 //col:22438
APIC_EOI =                                                     0x000000B0 //col:22443
APIC_REMOTE_READ =                                             0x000000C0 //col:22448
APIC_LOGICAL_DESTINATION =                                     0x000000D0 //col:22453
APIC_DESTINATION_FORMAT =                                      0x000000E0 //col:22460
APIC_SPURIOUS_INTERRUPT_VECTOR =                               0x000000F0 //col:22467
APIC_IN_SERVICE_BITS_31_0 =                                    0x00000100 //col:22472
APIC_IN_SERVICE_BITS_63_32 =                                   0x00000110 //col:22477
APIC_IN_SERVICE_BITS_95_64 =                                   0x00000120 //col:22482
APIC_IN_SERVICE_BITS_127_96 =                                  0x00000130 //col:22487
APIC_IN_SERVICE_BITS_159_128 =                                 0x00000140 //col:22492
APIC_IN_SERVICE_BITS_191_160 =                                 0x00000150 //col:22497
APIC_IN_SERVICE_BITS_223_192 =                                 0x00000160 //col:22502
APIC_IN_SERVICE_BITS_255_224 =                                 0x00000170 //col:22507
APIC_TRIGGER_MODE_BITS_31_0 =                                  0x00000180 //col:22512
APIC_TRIGGER_MODE_BITS_63_32 =                                 0x00000190 //col:22517
APIC_TRIGGER_MODE_BITS_95_64 =                                 0x000001A0 //col:22522
APIC_TRIGGER_MODE_BITS_127_96 =                                0x000001B0 //col:22527
APIC_TRIGGER_MODE_BITS_159_128 =                               0x000001C0 //col:22532
APIC_TRIGGER_MODE_BITS_191_160 =                               0x000001D0 //col:22537
APIC_TRIGGER_MODE_BITS_223_192 =                               0x000001E0 //col:22542
APIC_TRIGGER_MODE_BITS_255_224 =                               0x000001F0 //col:22547
APIC_INTERRUPT_REQUEST_BITS_31_0 =                             0x00000200 //col:22552
APIC_INTERRUPT_REQUEST_BITS_63_32 =                            0x00000210 //col:22557
APIC_INTERRUPT_REQUEST_BITS_95_64 =                            0x00000220 //col:22562
APIC_INTERRUPT_REQUEST_BITS_127_96 =                           0x00000230 //col:22567
APIC_INTERRUPT_REQUEST_BITS_159_128 =                          0x00000240 //col:22572
APIC_INTERRUPT_REQUEST_BITS_191_160 =                          0x00000250 //col:22577
APIC_INTERRUPT_REQUEST_BITS_223_192 =                          0x00000260 //col:22582
APIC_INTERRUPT_REQUEST_BITS_255_224 =                          0x00000270 //col:22587
APIC_ERROR_STATUS =                                            0x00000280 //col:22592
APIC_LVT_CORRECTED_MACHINE_CHECK_INTERRUPT =                   0x000002F0 //col:22597
APIC_INTERRUPT_COMMAND_BITS_0_31 =                             0x00000300 //col:22602
APIC_INTERRUPT_COMMAND_BITS_32_63 =                            0x00000310 //col:22607
APIC_LVT_TIMER =                                               0x00000320 //col:22612
APIC_LVT_THERMAL_SENSOR =                                      0x00000330 //col:22617
APIC_LVT_PERFORMANCE_MONITORING_COUNTERS =                     0x00000340 //col:22622
APIC_LVT_LINT0 =                                               0x00000350 //col:22627
APIC_LVT_LINT1 =                                               0x00000360 //col:22632
APIC_LVT_ERROR =                                               0x00000370 //col:22637
APIC_INITIAL_COUNT =                                           0x00000380 //col:22642
APIC_CURRENT_COUNT =                                           0x00000390 //col:22647
APIC_DIVIDE_CONFIGURATION =                                    0x000003E0 //col:22652
EFLAGS_CARRY_FLAG_BIT =                                        0 //col:22677
EFLAGS_CARRY_FLAG_FLAG =                                       0x01 //col:22678
EFLAGS_CARRY_FLAG_MASK =                                       0x01 //col:22679
EFLAGS_CARRY_FLAG(_) =                                         (((_) >> 0) & 0x01) //col:22680
EFLAGS_READ_AS_1_BIT =                                         1 //col:22686
EFLAGS_READ_AS_1_FLAG =                                        0x02 //col:22687
EFLAGS_READ_AS_1_MASK =                                        0x01 //col:22688
EFLAGS_READ_AS_1(_) =                                          (((_) >> 1) & 0x01) //col:22689
EFLAGS_PARITY_FLAG_BIT =                                       2 //col:22697
EFLAGS_PARITY_FLAG_FLAG =                                      0x04 //col:22698
EFLAGS_PARITY_FLAG_MASK =                                      0x01 //col:22699
EFLAGS_PARITY_FLAG(_) =                                        (((_) >> 2) & 0x01) //col:22700
EFLAGS_AUXILIARY_CARRY_FLAG_BIT =                              4 //col:22710
EFLAGS_AUXILIARY_CARRY_FLAG_FLAG =                             0x10 //col:22711
EFLAGS_AUXILIARY_CARRY_FLAG_MASK =                             0x01 //col:22712
EFLAGS_AUXILIARY_CARRY_FLAG(_) =                               (((_) >> 4) & 0x01) //col:22713
EFLAGS_ZERO_FLAG_BIT =                                         6 //col:22722
EFLAGS_ZERO_FLAG_FLAG =                                        0x40 //col:22723
EFLAGS_ZERO_FLAG_MASK =                                        0x01 //col:22724
EFLAGS_ZERO_FLAG(_) =                                          (((_) >> 6) & 0x01) //col:22725
EFLAGS_SIGN_FLAG_BIT =                                         7 //col:22734
EFLAGS_SIGN_FLAG_FLAG =                                        0x80 //col:22735
EFLAGS_SIGN_FLAG_MASK =                                        0x01 //col:22736
EFLAGS_SIGN_FLAG(_) =                                          (((_) >> 7) & 0x01) //col:22737
EFLAGS_TRAP_FLAG_BIT =                                         8 //col:22745
EFLAGS_TRAP_FLAG_FLAG =                                        0x100 //col:22746
EFLAGS_TRAP_FLAG_MASK =                                        0x01 //col:22747
EFLAGS_TRAP_FLAG(_) =                                          (((_) >> 8) & 0x01) //col:22748
EFLAGS_INTERRUPT_ENABLE_FLAG_BIT =                             9 //col:22757
EFLAGS_INTERRUPT_ENABLE_FLAG_FLAG =                            0x200 //col:22758
EFLAGS_INTERRUPT_ENABLE_FLAG_MASK =                            0x01 //col:22759
EFLAGS_INTERRUPT_ENABLE_FLAG(_) =                              (((_) >> 9) & 0x01) //col:22760
EFLAGS_DIRECTION_FLAG_BIT =                                    10 //col:22770
EFLAGS_DIRECTION_FLAG_FLAG =                                   0x400 //col:22771
EFLAGS_DIRECTION_FLAG_MASK =                                   0x01 //col:22772
EFLAGS_DIRECTION_FLAG(_) =                                     (((_) >> 10) & 0x01) //col:22773
EFLAGS_OVERFLOW_FLAG_BIT =                                     11 //col:22783
EFLAGS_OVERFLOW_FLAG_FLAG =                                    0x800 //col:22784
EFLAGS_OVERFLOW_FLAG_MASK =                                    0x01 //col:22785
EFLAGS_OVERFLOW_FLAG(_) =                                      (((_) >> 11) & 0x01) //col:22786
EFLAGS_IO_PRIVILEGE_LEVEL_BIT =                                12 //col:22796
EFLAGS_IO_PRIVILEGE_LEVEL_FLAG =                               0x3000 //col:22797
EFLAGS_IO_PRIVILEGE_LEVEL_MASK =                               0x03 //col:22798
EFLAGS_IO_PRIVILEGE_LEVEL(_) =                                 (((_) >> 12) & 0x03) //col:22799
EFLAGS_NESTED_TASK_FLAG_BIT =                                  14 //col:22808
EFLAGS_NESTED_TASK_FLAG_FLAG =                                 0x4000 //col:22809
EFLAGS_NESTED_TASK_FLAG_MASK =                                 0x01 //col:22810
EFLAGS_NESTED_TASK_FLAG(_) =                                   (((_) >> 14) & 0x01) //col:22811
EFLAGS_RESUME_FLAG_BIT =                                       16 //col:22820
EFLAGS_RESUME_FLAG_FLAG =                                      0x10000 //col:22821
EFLAGS_RESUME_FLAG_MASK =                                      0x01 //col:22822
EFLAGS_RESUME_FLAG(_) =                                        (((_) >> 16) & 0x01) //col:22823
EFLAGS_VIRTUAL_8086_MODE_FLAG_BIT =                            17 //col:22831
EFLAGS_VIRTUAL_8086_MODE_FLAG_FLAG =                           0x20000 //col:22832
EFLAGS_VIRTUAL_8086_MODE_FLAG_MASK =                           0x01 //col:22833
EFLAGS_VIRTUAL_8086_MODE_FLAG(_) =                             (((_) >> 17) & 0x01) //col:22834
EFLAGS_ALIGNMENT_CHECK_FLAG_BIT =                              18 //col:22846
EFLAGS_ALIGNMENT_CHECK_FLAG_FLAG =                             0x40000 //col:22847
EFLAGS_ALIGNMENT_CHECK_FLAG_MASK =                             0x01 //col:22848
EFLAGS_ALIGNMENT_CHECK_FLAG(_) =                               (((_) >> 18) & 0x01) //col:22849
EFLAGS_VIRTUAL_INTERRUPT_FLAG_BIT =                            19 //col:22858
EFLAGS_VIRTUAL_INTERRUPT_FLAG_FLAG =                           0x80000 //col:22859
EFLAGS_VIRTUAL_INTERRUPT_FLAG_MASK =                           0x01 //col:22860
EFLAGS_VIRTUAL_INTERRUPT_FLAG(_) =                             (((_) >> 19) & 0x01) //col:22861
EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_BIT =                    20 //col:22870
EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_FLAG =                   0x100000 //col:22871
EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_MASK =                   0x01 //col:22872
EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG(_) =                     (((_) >> 20) & 0x01) //col:22873
EFLAGS_IDENTIFICATION_FLAG_BIT =                               21 //col:22881
EFLAGS_IDENTIFICATION_FLAG_FLAG =                              0x200000 //col:22882
EFLAGS_IDENTIFICATION_FLAG_MASK =                              0x01 //col:22883
EFLAGS_IDENTIFICATION_FLAG(_) =                                (((_) >> 21) & 0x01) //col:22884
RFLAGS_CARRY_FLAG_BIT =                                        0 //col:22908
RFLAGS_CARRY_FLAG_FLAG =                                       0x01 //col:22909
RFLAGS_CARRY_FLAG_MASK =                                       0x01 //col:22910
RFLAGS_CARRY_FLAG(_) =                                         (((_) >> 0) & 0x01) //col:22911
RFLAGS_READ_AS_1_BIT =                                         1 //col:22917
RFLAGS_READ_AS_1_FLAG =                                        0x02 //col:22918
RFLAGS_READ_AS_1_MASK =                                        0x01 //col:22919
RFLAGS_READ_AS_1(_) =                                          (((_) >> 1) & 0x01) //col:22920
RFLAGS_PARITY_FLAG_BIT =                                       2 //col:22928
RFLAGS_PARITY_FLAG_FLAG =                                      0x04 //col:22929
RFLAGS_PARITY_FLAG_MASK =                                      0x01 //col:22930
RFLAGS_PARITY_FLAG(_) =                                        (((_) >> 2) & 0x01) //col:22931
RFLAGS_AUXILIARY_CARRY_FLAG_BIT =                              4 //col:22940
RFLAGS_AUXILIARY_CARRY_FLAG_FLAG =                             0x10 //col:22941
RFLAGS_AUXILIARY_CARRY_FLAG_MASK =                             0x01 //col:22942
RFLAGS_AUXILIARY_CARRY_FLAG(_) =                               (((_) >> 4) & 0x01) //col:22943
RFLAGS_ZERO_FLAG_BIT =                                         6 //col:22952
RFLAGS_ZERO_FLAG_FLAG =                                        0x40 //col:22953
RFLAGS_ZERO_FLAG_MASK =                                        0x01 //col:22954
RFLAGS_ZERO_FLAG(_) =                                          (((_) >> 6) & 0x01) //col:22955
RFLAGS_SIGN_FLAG_BIT =                                         7 //col:22963
RFLAGS_SIGN_FLAG_FLAG =                                        0x80 //col:22964
RFLAGS_SIGN_FLAG_MASK =                                        0x01 //col:22965
RFLAGS_SIGN_FLAG(_) =                                          (((_) >> 7) & 0x01) //col:22966
RFLAGS_TRAP_FLAG_BIT =                                         8 //col:22974
RFLAGS_TRAP_FLAG_FLAG =                                        0x100 //col:22975
RFLAGS_TRAP_FLAG_MASK =                                        0x01 //col:22976
RFLAGS_TRAP_FLAG(_) =                                          (((_) >> 8) & 0x01) //col:22977
RFLAGS_INTERRUPT_ENABLE_FLAG_BIT =                             9 //col:22985
RFLAGS_INTERRUPT_ENABLE_FLAG_FLAG =                            0x200 //col:22986
RFLAGS_INTERRUPT_ENABLE_FLAG_MASK =                            0x01 //col:22987
RFLAGS_INTERRUPT_ENABLE_FLAG(_) =                              (((_) >> 9) & 0x01) //col:22988
RFLAGS_DIRECTION_FLAG_BIT =                                    10 //col:22996
RFLAGS_DIRECTION_FLAG_FLAG =                                   0x400 //col:22997
RFLAGS_DIRECTION_FLAG_MASK =                                   0x01 //col:22998
RFLAGS_DIRECTION_FLAG(_) =                                     (((_) >> 10) & 0x01) //col:22999
RFLAGS_OVERFLOW_FLAG_BIT =                                     11 //col:23007
RFLAGS_OVERFLOW_FLAG_FLAG =                                    0x800 //col:23008
RFLAGS_OVERFLOW_FLAG_MASK =                                    0x01 //col:23009
RFLAGS_OVERFLOW_FLAG(_) =                                      (((_) >> 11) & 0x01) //col:23010
RFLAGS_IO_PRIVILEGE_LEVEL_BIT =                                12 //col:23018
RFLAGS_IO_PRIVILEGE_LEVEL_FLAG =                               0x3000 //col:23019
RFLAGS_IO_PRIVILEGE_LEVEL_MASK =                               0x03 //col:23020
RFLAGS_IO_PRIVILEGE_LEVEL(_) =                                 (((_) >> 12) & 0x03) //col:23021
RFLAGS_NESTED_TASK_FLAG_BIT =                                  14 //col:23029
RFLAGS_NESTED_TASK_FLAG_FLAG =                                 0x4000 //col:23030
RFLAGS_NESTED_TASK_FLAG_MASK =                                 0x01 //col:23031
RFLAGS_NESTED_TASK_FLAG(_) =                                   (((_) >> 14) & 0x01) //col:23032
RFLAGS_RESUME_FLAG_BIT =                                       16 //col:23041
RFLAGS_RESUME_FLAG_FLAG =                                      0x10000 //col:23042
RFLAGS_RESUME_FLAG_MASK =                                      0x01 //col:23043
RFLAGS_RESUME_FLAG(_) =                                        (((_) >> 16) & 0x01) //col:23044
RFLAGS_VIRTUAL_8086_MODE_FLAG_BIT =                            17 //col:23052
RFLAGS_VIRTUAL_8086_MODE_FLAG_FLAG =                           0x20000 //col:23053
RFLAGS_VIRTUAL_8086_MODE_FLAG_MASK =                           0x01 //col:23054
RFLAGS_VIRTUAL_8086_MODE_FLAG(_) =                             (((_) >> 17) & 0x01) //col:23055
RFLAGS_ALIGNMENT_CHECK_FLAG_BIT =                              18 //col:23065
RFLAGS_ALIGNMENT_CHECK_FLAG_FLAG =                             0x40000 //col:23066
RFLAGS_ALIGNMENT_CHECK_FLAG_MASK =                             0x01 //col:23067
RFLAGS_ALIGNMENT_CHECK_FLAG(_) =                               (((_) >> 18) & 0x01) //col:23068
RFLAGS_VIRTUAL_INTERRUPT_FLAG_BIT =                            19 //col:23076
RFLAGS_VIRTUAL_INTERRUPT_FLAG_FLAG =                           0x80000 //col:23077
RFLAGS_VIRTUAL_INTERRUPT_FLAG_MASK =                           0x01 //col:23078
RFLAGS_VIRTUAL_INTERRUPT_FLAG(_) =                             (((_) >> 19) & 0x01) //col:23079
RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_BIT =                    20 //col:23087
RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_FLAG =                   0x100000 //col:23088
RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_MASK =                   0x01 //col:23089
RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG(_) =                     (((_) >> 20) & 0x01) //col:23090
RFLAGS_IDENTIFICATION_FLAG_BIT =                               21 //col:23098
RFLAGS_IDENTIFICATION_FLAG_FLAG =                              0x200000 //col:23099
RFLAGS_IDENTIFICATION_FLAG_MASK =                              0x01 //col:23100
RFLAGS_IDENTIFICATION_FLAG(_) =                                (((_) >> 21) & 0x01) //col:23101
CONTROL_PROTECTION_EXCEPTION_CPEC_BIT =                        0 //col:23131
CONTROL_PROTECTION_EXCEPTION_CPEC_FLAG =                       0x7FFF //col:23132
CONTROL_PROTECTION_EXCEPTION_CPEC_MASK =                       0x7FFF //col:23133
CONTROL_PROTECTION_EXCEPTION_CPEC(_) =                         (((_) >> 0) & 0x7FFF) //col:23134
CONTROL_PROTECTION_EXCEPTION_ENCL_BIT =                        15 //col:23140
CONTROL_PROTECTION_EXCEPTION_ENCL_FLAG =                       0x8000 //col:23141
CONTROL_PROTECTION_EXCEPTION_ENCL_MASK =                       0x01 //col:23142
CONTROL_PROTECTION_EXCEPTION_ENCL(_) =                         (((_) >> 15) & 0x01) //col:23143
EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_BIT =                      0 //col:23328
EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_FLAG =                     0x01 //col:23329
EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_MASK =                     0x01 //col:23330
EXCEPTION_ERROR_CODE_EXTERNAL_EVENT(_) =                       (((_) >> 0) & 0x01) //col:23331
EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_BIT =                 1 //col:23338
EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_FLAG =                0x02 //col:23339
EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_MASK =                0x01 //col:23340
EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION(_) =                  (((_) >> 1) & 0x01) //col:23341
EXCEPTION_ERROR_CODE_GDT_LDT_BIT =                             2 //col:23349
EXCEPTION_ERROR_CODE_GDT_LDT_FLAG =                            0x04 //col:23350
EXCEPTION_ERROR_CODE_GDT_LDT_MASK =                            0x01 //col:23351
EXCEPTION_ERROR_CODE_GDT_LDT(_) =                              (((_) >> 2) & 0x01) //col:23352
EXCEPTION_ERROR_CODE_INDEX_BIT =                               3 //col:23363
EXCEPTION_ERROR_CODE_INDEX_FLAG =                              0xFFF8 //col:23364
EXCEPTION_ERROR_CODE_INDEX_MASK =                              0x1FFF //col:23365
EXCEPTION_ERROR_CODE_INDEX(_) =                                (((_) >> 3) & 0x1FFF) //col:23366
PAGE_FAULT_EXCEPTION_PRESENT_BIT =                             0 //col:23387
PAGE_FAULT_EXCEPTION_PRESENT_FLAG =                            0x01 //col:23388
PAGE_FAULT_EXCEPTION_PRESENT_MASK =                            0x01 //col:23389
PAGE_FAULT_EXCEPTION_PRESENT(_) =                              (((_) >> 0) & 0x01) //col:23390
PAGE_FAULT_EXCEPTION_WRITE_BIT =                               1 //col:23397
PAGE_FAULT_EXCEPTION_WRITE_FLAG =                              0x02 //col:23398
PAGE_FAULT_EXCEPTION_WRITE_MASK =                              0x01 //col:23399
PAGE_FAULT_EXCEPTION_WRITE(_) =                                (((_) >> 1) & 0x01) //col:23400
PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_BIT =                    2 //col:23409
PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_FLAG =                   0x04 //col:23410
PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_MASK =                   0x01 //col:23411
PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS(_) =                     (((_) >> 2) & 0x01) //col:23412
PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_BIT =              3 //col:23423
PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_FLAG =             0x08 //col:23424
PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_MASK =             0x01 //col:23425
PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION(_) =               (((_) >> 3) & 0x01) //col:23426
PAGE_FAULT_EXCEPTION_EXECUTE_BIT =                             4 //col:23435
PAGE_FAULT_EXCEPTION_EXECUTE_FLAG =                            0x10 //col:23436
PAGE_FAULT_EXCEPTION_EXECUTE_MASK =                            0x01 //col:23437
PAGE_FAULT_EXCEPTION_EXECUTE(_) =                              (((_) >> 4) & 0x01) //col:23438
PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_BIT =            5 //col:23449
PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_FLAG =           0x20 //col:23450
PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_MASK =           0x01 //col:23451
PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION(_) =             (((_) >> 5) & 0x01) //col:23452
PAGE_FAULT_EXCEPTION_SHADOW_STACK_BIT =                        6 //col:23462
PAGE_FAULT_EXCEPTION_SHADOW_STACK_FLAG =                       0x40 //col:23463
PAGE_FAULT_EXCEPTION_SHADOW_STACK_MASK =                       0x01 //col:23464
PAGE_FAULT_EXCEPTION_SHADOW_STACK(_) =                         (((_) >> 6) & 0x01) //col:23465
PAGE_FAULT_EXCEPTION_HLAT_BIT =                                7 //col:23477
PAGE_FAULT_EXCEPTION_HLAT_FLAG =                               0x80 //col:23478
PAGE_FAULT_EXCEPTION_HLAT_MASK =                               0x01 //col:23479
PAGE_FAULT_EXCEPTION_HLAT(_) =                                 (((_) >> 7) & 0x01) //col:23480
PAGE_FAULT_EXCEPTION_SGX_BIT =                                 15 //col:23489
PAGE_FAULT_EXCEPTION_SGX_FLAG =                                0x8000 //col:23490
PAGE_FAULT_EXCEPTION_SGX_MASK =                                0x01 //col:23491
PAGE_FAULT_EXCEPTION_SGX(_) =                                  (((_) >> 15) & 0x01) //col:23492
MEMORY_TYPE_UNCACHEABLE =                                      0x00000000 //col:23523
MEMORY_TYPE_WRITE_COMBINING =                                  0x00000001 //col:23540
MEMORY_TYPE_WRITE_THROUGH =                                    0x00000004 //col:23552
MEMORY_TYPE_WRITE_PROTECTED =                                  0x00000005 //col:23561
MEMORY_TYPE_WRITE_BACK =                                       0x00000006 //col:23577
MEMORY_TYPE_UNCACHEABLE_MINUS =                                0x00000007 //col:23586
MEMORY_TYPE_INVALID =                                          0x000000FF //col:23587
VTD_Lower64_PRESENT_BIT =                                      0 //col:23615
VTD_Lower64_PRESENT_FLAG =                                     0x01 //col:23616
VTD_Lower64_PRESENT_MASK =                                     0x01 //col:23617
VTD_Lower64_PRESENT(_) =                                       (((_) >> 0) & 0x01) //col:23618
VTD_Lower64_CONTEXT_TABLE_POINTER_BIT =                        12 //col:23626
VTD_Lower64_CONTEXT_TABLE_POINTER_FLAG =                       0xFFFFFFFFFFFFF000 //col:23627
VTD_Lower64_CONTEXT_TABLE_POINTER_MASK =                       0xFFFFFFFFFFFFF //col:23628
VTD_Lower64_CONTEXT_TABLE_POINTER(_) =                         (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:23629
VTD_Upper64_RESERVED_BIT =                                     0 //col:23643
VTD_Upper64_RESERVED_FLAG =                                    0xFFFFFFFFFFFFFFFF //col:23644
VTD_Upper64_RESERVED_MASK =                                    0xFFFFFFFFFFFFFFFF //col:23645
VTD_Upper64_RESERVED(_) =                                      (((_) >> 0) & 0xFFFFFFFFFFFFFFFF) //col:23646
VTD_Lower64_PRESENT_BIT =                                      0 //col:23673
VTD_Lower64_PRESENT_FLAG =                                     0x01 //col:23674
VTD_Lower64_PRESENT_MASK =                                     0x01 //col:23675
VTD_Lower64_PRESENT(_) =                                       (((_) >> 0) & 0x01) //col:23676
VTD_Lower64_FAULT_PROCESSING_DISABLE_BIT =                     1 //col:23685
VTD_Lower64_FAULT_PROCESSING_DISABLE_FLAG =                    0x02 //col:23686
VTD_Lower64_FAULT_PROCESSING_DISABLE_MASK =                    0x01 //col:23687
VTD_Lower64_FAULT_PROCESSING_DISABLE(_) =                      (((_) >> 1) & 0x01) //col:23688
VTD_Lower64_TRANSLATION_TYPE_BIT =                             2 //col:23703
VTD_Lower64_TRANSLATION_TYPE_FLAG =                            0x0C //col:23704
VTD_Lower64_TRANSLATION_TYPE_MASK =                            0x03 //col:23705
VTD_Lower64_TRANSLATION_TYPE(_) =                              (((_) >> 2) & 0x03) //col:23706
VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_BIT =        12 //col:23716
VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_FLAG =       0xFFFFFFFFFFFFF000 //col:23717
VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_MASK =       0xFFFFFFFFFFFFF //col:23718
VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER(_) =         (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:23719
VTD_Upper64_ADDRESS_WIDTH_BIT =                                0 //col:23745
VTD_Upper64_ADDRESS_WIDTH_FLAG =                               0x07 //col:23746
VTD_Upper64_ADDRESS_WIDTH_MASK =                               0x07 //col:23747
VTD_Upper64_ADDRESS_WIDTH(_) =                                 (((_) >> 0) & 0x07) //col:23748
VTD_Upper64_IGNORED_BIT =                                      3 //col:23754
VTD_Upper64_IGNORED_FLAG =                                     0x78 //col:23755
VTD_Upper64_IGNORED_MASK =                                     0x0F //col:23756
VTD_Upper64_IGNORED(_) =                                       (((_) >> 3) & 0x0F) //col:23757
VTD_Upper64_DOMAIN_IDENTIFIER_BIT =                            8 //col:23773
VTD_Upper64_DOMAIN_IDENTIFIER_FLAG =                           0x3FF00 //col:23774
VTD_Upper64_DOMAIN_IDENTIFIER_MASK =                           0x3FF //col:23775
VTD_Upper64_DOMAIN_IDENTIFIER(_) =                             (((_) >> 8) & 0x3FF) //col:23776
VTD_ROOT_ENTRY_COUNT =                                         0x00000100 //col:23792
VTD_CONTEXT_ENTRY_COUNT =                                      0x00000100 //col:23793
VTD_VERSION =                                                  0x00000000 //col:23805
VTD_VERSION_MINOR_BIT =                                        0 //col:23816
VTD_VERSION_MINOR_FLAG =                                       0x0F //col:23817
VTD_VERSION_MINOR_MASK =                                       0x0F //col:23818
VTD_VERSION_MINOR(_) =                                         (((_) >> 0) & 0x0F) //col:23819
VTD_VERSION_MAJOR_BIT =                                        4 //col:23827
VTD_VERSION_MAJOR_FLAG =                                       0xF0 //col:23828
VTD_VERSION_MAJOR_MASK =                                       0x0F //col:23829
VTD_VERSION_MAJOR(_) =                                         (((_) >> 4) & 0x0F) //col:23830
VTD_CAPABILITY =                                               0x00000008 //col:23844
VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_BIT =               0 //col:23863
VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_FLAG =              0x07 //col:23864
VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_MASK =              0x07 //col:23865
VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED(_) =                (((_) >> 0) & 0x07) //col:23866
VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_BIT =                    3 //col:23876
VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_FLAG =                   0x08 //col:23877
VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_MASK =                   0x01 //col:23878
VTD_CAPABILITY_ADVANCED_FAULT_LOGGING(_) =                     (((_) >> 3) & 0x01) //col:23879
VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_BIT =            4 //col:23891
VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_FLAG =           0x10 //col:23892
VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_MASK =           0x01 //col:23893
VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING(_) =             (((_) >> 4) & 0x01) //col:23894
VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_BIT =               5 //col:23904
VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_FLAG =              0x20 //col:23905
VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_MASK =              0x01 //col:23906
VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION(_) =                (((_) >> 5) & 0x01) //col:23907
VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_BIT =              6 //col:23917
VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_FLAG =             0x40 //col:23918
VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_MASK =             0x01 //col:23919
VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION(_) =               (((_) >> 6) & 0x01) //col:23920
VTD_CAPABILITY_CACHING_MODE_BIT =                              7 //col:23934
VTD_CAPABILITY_CACHING_MODE_FLAG =                             0x80 //col:23935
VTD_CAPABILITY_CACHING_MODE_MASK =                             0x01 //col:23936
VTD_CAPABILITY_CACHING_MODE(_) =                               (((_) >> 7) & 0x01) //col:23937
VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_BIT =   8 //col:23956
VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_FLAG =  0x1F00 //col:23957
VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_MASK =  0x1F //col:23958
VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS(_) =    (((_) >> 8) & 0x1F) //col:23959
VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_BIT =               16 //col:23977
VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_FLAG =              0x3F0000 //col:23978
VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_MASK =              0x3F //col:23979
VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH(_) =                (((_) >> 16) & 0x3F) //col:23980
VTD_CAPABILITY_ZERO_LENGTH_READ_BIT =                          22 //col:23992
VTD_CAPABILITY_ZERO_LENGTH_READ_FLAG =                         0x400000 //col:23993
VTD_CAPABILITY_ZERO_LENGTH_READ_MASK =                         0x01 //col:23994
VTD_CAPABILITY_ZERO_LENGTH_READ(_) =                           (((_) >> 22) & 0x01) //col:23995
VTD_CAPABILITY_DEPRECATED_BIT =                                23 //col:24003
VTD_CAPABILITY_DEPRECATED_FLAG =                               0x800000 //col:24004
VTD_CAPABILITY_DEPRECATED_MASK =                               0x01 //col:24005
VTD_CAPABILITY_DEPRECATED(_) =                                 (((_) >> 23) & 0x01) //col:24006
VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_BIT =           24 //col:24016
VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_FLAG =          0x3FF000000 //col:24017
VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_MASK =          0x3FF //col:24018
VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET(_) =            (((_) >> 24) & 0x3FF) //col:24019
VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_BIT =           34 //col:24035
VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_FLAG =          0x3C00000000 //col:24036
VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_MASK =          0x0F //col:24037
VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT(_) =            (((_) >> 34) & 0x0F) //col:24038
VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_BIT =               39 //col:24054
VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_FLAG =              0x8000000000 //col:24055
VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_MASK =              0x01 //col:24056
VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION(_) =                (((_) >> 39) & 0x01) //col:24057
VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_BIT =       40 //col:24068
VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_FLAG =      0xFF0000000000 //col:24069
VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_MASK =      0xFF //col:24070
VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS(_) =        (((_) >> 40) & 0xFF) //col:24071
VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_BIT =                48 //col:24084
VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_FLAG =               0x3F000000000000 //col:24085
VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_MASK =               0x3F //col:24086
VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE(_) =                 (((_) >> 48) & 0x3F) //col:24087
VTD_CAPABILITY_WRITE_DRAINING_BIT =                            54 //col:24100
VTD_CAPABILITY_WRITE_DRAINING_FLAG =                           0x40000000000000 //col:24101
VTD_CAPABILITY_WRITE_DRAINING_MASK =                           0x01 //col:24102
VTD_CAPABILITY_WRITE_DRAINING(_) =                             (((_) >> 54) & 0x01) //col:24103
VTD_CAPABILITY_READ_DRAINING_BIT =                             55 //col:24116
VTD_CAPABILITY_READ_DRAINING_FLAG =                            0x80000000000000 //col:24117
VTD_CAPABILITY_READ_DRAINING_MASK =                            0x01 //col:24118
VTD_CAPABILITY_READ_DRAINING(_) =                              (((_) >> 55) & 0x01) //col:24119
VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_BIT =           56 //col:24128
VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_FLAG =          0x100000000000000 //col:24129
VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_MASK =          0x01 //col:24130
VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT(_) =            (((_) >> 56) & 0x01) //col:24131
VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_BIT =                 59 //col:24144
VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_FLAG =                0x800000000000000 //col:24145
VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_MASK =                0x01 //col:24146
VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT(_) =                  (((_) >> 59) & 0x01) //col:24147
VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_BIT =         60 //col:24158
VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_FLAG =        0x1000000000000000 //col:24159
VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_MASK =        0x01 //col:24160
VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT(_) =          (((_) >> 60) & 0x01) //col:24161
VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_BIT = 62 //col:24172
VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_FLAG = 0x4000000000000000 //col:24173
VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_MASK = 0x01 //col:24174
VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT(_) = (((_) >> 62) & 0x01) //col:24175
VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_BIT =   63 //col:24185
VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_FLAG =  0x8000000000000000 //col:24186
VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_MASK =  0x01 //col:24187
VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT(_) =    (((_) >> 63) & 0x01) //col:24188
VTD_EXTENDED_CAPABILITY =                                      0x00000010 //col:24201
VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_BIT =              0 //col:24218
VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_FLAG =             0x01 //col:24219
VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_MASK =             0x01 //col:24220
VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY(_) =               (((_) >> 0) & 0x01) //col:24221
VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_BIT =      1 //col:24231
VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_FLAG =     0x02 //col:24232
VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_MASK =     0x01 //col:24233
VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT(_) =       (((_) >> 1) & 0x01) //col:24234
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_BIT =               2 //col:24245
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_FLAG =              0x04 //col:24246
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_MASK =              0x01 //col:24247
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT(_) =                (((_) >> 2) & 0x01) //col:24248
VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_BIT =      3 //col:24259
VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_FLAG =     0x08 //col:24260
VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_MASK =     0x01 //col:24261
VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT(_) =       (((_) >> 3) & 0x01) //col:24262
VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_BIT =          4 //col:24273
VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_FLAG =         0x10 //col:24274
VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_MASK =         0x01 //col:24275
VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE(_) =           (((_) >> 4) & 0x01) //col:24276
VTD_EXTENDED_CAPABILITY_DEPRECATED1_BIT =                      5 //col:24284
VTD_EXTENDED_CAPABILITY_DEPRECATED1_FLAG =                     0x20 //col:24285
VTD_EXTENDED_CAPABILITY_DEPRECATED1_MASK =                     0x01 //col:24286
VTD_EXTENDED_CAPABILITY_DEPRECATED1(_) =                       (((_) >> 5) & 0x01) //col:24287
VTD_EXTENDED_CAPABILITY_PASS_THROUGH_BIT =                     6 //col:24297
VTD_EXTENDED_CAPABILITY_PASS_THROUGH_FLAG =                    0x40 //col:24298
VTD_EXTENDED_CAPABILITY_PASS_THROUGH_MASK =                    0x01 //col:24299
VTD_EXTENDED_CAPABILITY_PASS_THROUGH(_) =                      (((_) >> 6) & 0x01) //col:24300
VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_BIT =                    7 //col:24312
VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_FLAG =                   0x80 //col:24313
VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_MASK =                   0x01 //col:24314
VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL(_) =                     (((_) >> 7) & 0x01) //col:24315
VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_BIT =            8 //col:24326
VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_FLAG =           0x3FF00 //col:24327
VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_MASK =           0x3FF //col:24328
VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET(_) =             (((_) >> 8) & 0x3FF) //col:24329
VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_BIT =        20 //col:24340
VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_FLAG =       0xF00000 //col:24341
VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_MASK =       0x0F //col:24342
VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE(_) =         (((_) >> 20) & 0x0F) //col:24343
VTD_EXTENDED_CAPABILITY_DEPRECATED2_BIT =                      24 //col:24353
VTD_EXTENDED_CAPABILITY_DEPRECATED2_FLAG =                     0x1000000 //col:24354
VTD_EXTENDED_CAPABILITY_DEPRECATED2_MASK =                     0x01 //col:24355
VTD_EXTENDED_CAPABILITY_DEPRECATED2(_) =                       (((_) >> 24) & 0x01) //col:24356
VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_BIT =              25 //col:24370
VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_FLAG =             0x2000000 //col:24371
VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_MASK =             0x01 //col:24372
VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT(_) =               (((_) >> 25) & 0x01) //col:24373
VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_BIT =       26 //col:24385
VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_FLAG =      0x4000000 //col:24386
VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_MASK =      0x01 //col:24387
VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT(_) =        (((_) >> 26) & 0x01) //col:24388
VTD_EXTENDED_CAPABILITY_DEPRECATED3_BIT =                      28 //col:24397
VTD_EXTENDED_CAPABILITY_DEPRECATED3_FLAG =                     0x10000000 //col:24398
VTD_EXTENDED_CAPABILITY_DEPRECATED3_MASK =                     0x01 //col:24399
VTD_EXTENDED_CAPABILITY_DEPRECATED3(_) =                       (((_) >> 28) & 0x01) //col:24400
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_BIT =             29 //col:24412
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_FLAG =            0x20000000 //col:24413
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_MASK =            0x01 //col:24414
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT(_) =              (((_) >> 29) & 0x01) //col:24415
VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_BIT =          30 //col:24427
VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_FLAG =         0x40000000 //col:24428
VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_MASK =         0x01 //col:24429
VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT(_) =           (((_) >> 30) & 0x01) //col:24430
VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_BIT =            33 //col:24442
VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_FLAG =           0x200000000 //col:24443
VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_MASK =           0x01 //col:24444
VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT(_) =             (((_) >> 33) & 0x01) //col:24445
VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_BIT =   34 //col:24457
VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_FLAG =  0x400000000 //col:24458
VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_MASK =  0x01 //col:24459
VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT(_) =    (((_) >> 34) & 0x01) //col:24460
VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_BIT =             35 //col:24472
VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_FLAG =            0xF800000000 //col:24473
VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_MASK =            0x1F //col:24474
VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED(_) =              (((_) >> 35) & 0x1F) //col:24475
VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_BIT = 40 //col:24487
VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_FLAG = 0x10000000000 //col:24488
VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_MASK = 0x01 //col:24489
VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT(_) =  (((_) >> 40) & 0x01) //col:24490
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_BIT = 41 //col:24501
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_FLAG = 0x20000000000 //col:24502
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_MASK = 0x01 //col:24503
VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE(_) =  (((_) >> 41) & 0x01) //col:24504
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_BIT =       42 //col:24515
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_FLAG =      0x40000000000 //col:24516
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_MASK =      0x01 //col:24517
VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT(_) =        (((_) >> 42) & 0x01) //col:24518
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_BIT = 43 //col:24531
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_FLAG = 0x80000000000 //col:24532
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_MASK = 0x01 //col:24533
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT(_) = (((_) >> 43) & 0x01) //col:24534
VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_BIT =          44 //col:24547
VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_FLAG =         0x100000000000 //col:24548
VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_MASK =         0x01 //col:24549
VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT(_) =           (((_) >> 44) & 0x01) //col:24550
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_BIT = 45 //col:24562
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_FLAG = 0x200000000000 //col:24563
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_MASK = 0x01 //col:24564
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT(_) = (((_) >> 45) & 0x01) //col:24565
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_BIT = 46 //col:24576
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_FLAG = 0x400000000000 //col:24577
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_MASK = 0x01 //col:24578
VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT(_) =  (((_) >> 46) & 0x01) //col:24579
VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_BIT =  47 //col:24591
VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_FLAG = 0x800000000000 //col:24592
VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_MASK = 0x01 //col:24593
VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT(_) =   (((_) >> 47) & 0x01) //col:24594
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_BIT = 48 //col:24607
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_FLAG = 0x1000000000000 //col:24608
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_MASK = 0x01 //col:24609
VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY(_) = (((_) >> 48) & 0x01) //col:24610
VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_BIT =                49 //col:24621
VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_FLAG =               0x2000000000000 //col:24622
VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_MASK =               0x01 //col:24623
VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT(_) =                 (((_) >> 49) & 0x01) //col:24624
VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_BIT =           52 //col:24635
VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_FLAG =          0x10000000000000 //col:24636
VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_MASK =          0x01 //col:24637
VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT(_) =            (((_) >> 52) & 0x01) //col:24638
VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_BIT =                 53 //col:24650
VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_FLAG =                0x20000000000000 //col:24651
VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_MASK =                0x01 //col:24652
VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT(_) =                  (((_) >> 53) & 0x01) //col:24653
VTD_GLOBAL_COMMAND =                                           0x00000018 //col:24674
VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_BIT =        23 //col:24694
VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_FLAG =       0x800000 //col:24695
VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_MASK =       0x01 //col:24696
VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT(_) =         (((_) >> 23) & 0x01) //col:24697
VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_BIT =     24 //col:24712
VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_FLAG =    0x1000000 //col:24713
VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_MASK =    0x01 //col:24714
VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER(_) =      (((_) >> 24) & 0x01) //col:24715
VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_BIT =            25 //col:24736
VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_FLAG =           0x2000000 //col:24737
VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_MASK =           0x01 //col:24738
VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE(_) =             (((_) >> 25) & 0x01) //col:24739
VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_BIT =            26 //col:24752
VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_FLAG =           0x4000000 //col:24753
VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_MASK =           0x01 //col:24754
VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE(_) =             (((_) >> 26) & 0x01) //col:24755
VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_BIT =                    27 //col:24767
VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_FLAG =                   0x8000000 //col:24768
VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_MASK =                   0x01 //col:24769
VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH(_) =                     (((_) >> 27) & 0x01) //col:24770
VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_BIT =         28 //col:24785
VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_FLAG =        0x10000000 //col:24786
VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_MASK =        0x01 //col:24787
VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING(_) =          (((_) >> 28) & 0x01) //col:24788
VTD_GLOBAL_COMMAND_SET_FAULT_LOG_BIT =                         29 //col:24802
VTD_GLOBAL_COMMAND_SET_FAULT_LOG_FLAG =                        0x20000000 //col:24803
VTD_GLOBAL_COMMAND_SET_FAULT_LOG_MASK =                        0x01 //col:24804
VTD_GLOBAL_COMMAND_SET_FAULT_LOG(_) =                          (((_) >> 29) & 0x01) //col:24805
VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_BIT =                30 //col:24819
VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_FLAG =               0x40000000 //col:24820
VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_MASK =               0x01 //col:24821
VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER(_) =                 (((_) >> 30) & 0x01) //col:24822
VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_BIT =                    31 //col:24842
VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_FLAG =                   0x80000000 //col:24843
VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_MASK =                   0x01 //col:24844
VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE(_) =                     (((_) >> 31) & 0x01) //col:24845
VTD_GLOBAL_STATUS =                                            0x0000001C //col:24858
VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_BIT =  23 //col:24875
VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_FLAG = 0x800000 //col:24876
VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_MASK = 0x01 //col:24877
VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS(_) =   (((_) >> 23) & 0x01) //col:24878
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_BIT = 24 //col:24889
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_FLAG = 0x1000000 //col:24890
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_MASK = 0x01 //col:24891
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS(_) = (((_) >> 24) & 0x01) //col:24892
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_BIT =      25 //col:24902
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_FLAG =     0x2000000 //col:24903
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_MASK =     0x01 //col:24904
VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS(_) =       (((_) >> 25) & 0x01) //col:24905
VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_BIT =      26 //col:24915
VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_FLAG =     0x4000000 //col:24916
VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_MASK =     0x01 //col:24917
VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS(_) =       (((_) >> 26) & 0x01) //col:24918
VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_BIT =              27 //col:24929
VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_FLAG =             0x8000000 //col:24930
VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_MASK =             0x01 //col:24931
VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS(_) =               (((_) >> 27) & 0x01) //col:24932
VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_BIT =          28 //col:24943
VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_FLAG =         0x10000000 //col:24944
VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_MASK =         0x01 //col:24945
VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS(_) =           (((_) >> 28) & 0x01) //col:24946
VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_BIT =                       29 //col:24957
VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_FLAG =                      0x20000000 //col:24958
VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_MASK =                      0x01 //col:24959
VTD_GLOBAL_STATUS_FAULT_LOG_STATUS(_) =                        (((_) >> 29) & 0x01) //col:24960
VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_BIT =              30 //col:24971
VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_FLAG =             0x40000000 //col:24972
VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_MASK =             0x01 //col:24973
VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS(_) =               (((_) >> 30) & 0x01) //col:24974
VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_BIT =              31 //col:24984
VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_FLAG =             0x80000000 //col:24985
VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_MASK =             0x01 //col:24986
VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS(_) =               (((_) >> 31) & 0x01) //col:24987
VTD_ROOT_TABLE_ADDRESS =                                       0x00000020 //col:25002
VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_BIT =            10 //col:25023
VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_FLAG =           0xC00 //col:25024
VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_MASK =           0x03 //col:25025
VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE(_) =             (((_) >> 10) & 0x03) //col:25026
VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_BIT =                12 //col:25036
VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_FLAG =               0xFFFFFFFFFFFFF000 //col:25037
VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_MASK =               0xFFFFFFFFFFFFF //col:25038
VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS(_) =                 (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:25039
VTD_CONTEXT_COMMAND =                                          0x00000028 //col:25053
VTD_CONTEXT_COMMAND_DOMAIN_ID_BIT =                            0 //col:25068
VTD_CONTEXT_COMMAND_DOMAIN_ID_FLAG =                           0xFFFF //col:25069
VTD_CONTEXT_COMMAND_DOMAIN_ID_MASK =                           0xFFFF //col:25070
VTD_CONTEXT_COMMAND_DOMAIN_ID(_) =                             (((_) >> 0) & 0xFFFF) //col:25071
VTD_CONTEXT_COMMAND_SOURCE_ID_BIT =                            16 //col:25082
VTD_CONTEXT_COMMAND_SOURCE_ID_FLAG =                           0xFFFF0000 //col:25083
VTD_CONTEXT_COMMAND_SOURCE_ID_MASK =                           0xFFFF //col:25084
VTD_CONTEXT_COMMAND_SOURCE_ID(_) =                             (((_) >> 16) & 0xFFFF) //col:25085
VTD_CONTEXT_COMMAND_FUNCTION_MASK_BIT =                        32 //col:25103
VTD_CONTEXT_COMMAND_FUNCTION_MASK_FLAG =                       0x300000000 //col:25104
VTD_CONTEXT_COMMAND_FUNCTION_MASK_MASK =                       0x03 //col:25105
VTD_CONTEXT_COMMAND_FUNCTION_MASK(_) =                         (((_) >> 32) & 0x03) //col:25106
VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_BIT = 59 //col:25129
VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_FLAG = 0x1800000000000000 //col:25130
VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_MASK = 0x03 //col:25131
VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY(_) = (((_) >> 59) & 0x03) //col:25132
VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_BIT = 61 //col:25148
VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_FLAG = 0x6000000000000000 //col:25149
VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_MASK = 0x03 //col:25150
VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY(_) = (((_) >> 61) & 0x03) //col:25151
VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_BIT =             63 //col:25172
VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_FLAG =            0x8000000000000000 //col:25173
VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_MASK =            0x01 //col:25174
VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE(_) =              (((_) >> 63) & 0x01) //col:25175
VTD_INVALIDATE_ADDRESS =                                       0x00000000 //col:25190
VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_BIT =                      0 //col:25207
VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_FLAG =                     0x3F //col:25208
VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_MASK =                     0x3F //col:25209
VTD_INVALIDATE_ADDRESS_ADDRESS_MASK(_) =                       (((_) >> 0) & 0x3F) //col:25210
VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_BIT =                 6 //col:25227
VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_FLAG =                0x40 //col:25228
VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_MASK =                0x01 //col:25229
VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT(_) =                  (((_) >> 6) & 0x01) //col:25230
VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_BIT =                      12 //col:25243
VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_FLAG =                     0xFFFFFFFFFFFFF000 //col:25244
VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_MASK =                     0xFFFFFFFFFFFFF //col:25245
VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS(_) =                       (((_) >> 12) & 0xFFFFFFFFFFFFF) //col:25246
VTD_IOTLB_INVALIDATE =                                         0x00000008 //col:25260
VTD_IOTLB_INVALIDATE_DOMAIN_ID_BIT =                           32 //col:25277
VTD_IOTLB_INVALIDATE_DOMAIN_ID_FLAG =                          0xFFFF00000000 //col:25278
VTD_IOTLB_INVALIDATE_DOMAIN_ID_MASK =                          0xFFFF //col:25279
VTD_IOTLB_INVALIDATE_DOMAIN_ID(_) =                            (((_) >> 32) & 0xFFFF) //col:25280
VTD_IOTLB_INVALIDATE_DRAIN_WRITES_BIT =                        48 //col:25291
VTD_IOTLB_INVALIDATE_DRAIN_WRITES_FLAG =                       0x1000000000000 //col:25292
VTD_IOTLB_INVALIDATE_DRAIN_WRITES_MASK =                       0x01 //col:25293
VTD_IOTLB_INVALIDATE_DRAIN_WRITES(_) =                         (((_) >> 48) & 0x01) //col:25294
VTD_IOTLB_INVALIDATE_DRAIN_READS_BIT =                         49 //col:25305
VTD_IOTLB_INVALIDATE_DRAIN_READS_FLAG =                        0x2000000000000 //col:25306
VTD_IOTLB_INVALIDATE_DRAIN_READS_MASK =                        0x01 //col:25307
VTD_IOTLB_INVALIDATE_DRAIN_READS(_) =                          (((_) >> 49) & 0x01) //col:25308
VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_BIT = 57 //col:25333
VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_FLAG = 0x600000000000000 //col:25334
VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_MASK = 0x03 //col:25335
VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY(_) = (((_) >> 57) & 0x03) //col:25336
VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_BIT = 60 //col:25354
VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_FLAG = 0x3000000000000000 //col:25355
VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_MASK = 0x03 //col:25356
VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY(_) = (((_) >> 60) & 0x03) //col:25357
VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_BIT =                    63 //col:25378
VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_FLAG =                   0x8000000000000000 //col:25379
VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_MASK =                   0x01 //col:25380
VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB(_) =                     (((_) >> 63) & 0x01) //col:25381
XCR0_X87_BIT =                                                 0 //col:25399
XCR0_X87_FLAG =                                                0x01 //col:25400
XCR0_X87_MASK =                                                0x01 //col:25401
XCR0_X87(_) =                                                  (((_) >> 0) & 0x01) //col:25402
XCR0_SSE_BIT =                                                 1 //col:25409
XCR0_SSE_FLAG =                                                0x02 //col:25410
XCR0_SSE_MASK =                                                0x01 //col:25411
XCR0_SSE(_) =                                                  (((_) >> 1) & 0x01) //col:25412
XCR0_AVX_BIT =                                                 2 //col:25419
XCR0_AVX_FLAG =                                                0x04 //col:25420
XCR0_AVX_MASK =                                                0x01 //col:25421
XCR0_AVX(_) =                                                  (((_) >> 2) & 0x01) //col:25422
XCR0_BNDREG_BIT =                                              3 //col:25429
XCR0_BNDREG_FLAG =                                             0x08 //col:25430
XCR0_BNDREG_MASK =                                             0x01 //col:25431
XCR0_BNDREG(_) =                                               (((_) >> 3) & 0x01) //col:25432
XCR0_BNDCSR_BIT =                                              4 //col:25439
XCR0_BNDCSR_FLAG =                                             0x10 //col:25440
XCR0_BNDCSR_MASK =                                             0x01 //col:25441
XCR0_BNDCSR(_) =                                               (((_) >> 4) & 0x01) //col:25442
XCR0_OPMASK_BIT =                                              5 //col:25449
XCR0_OPMASK_FLAG =                                             0x20 //col:25450
XCR0_OPMASK_MASK =                                             0x01 //col:25451
XCR0_OPMASK(_) =                                               (((_) >> 5) & 0x01) //col:25452
XCR0_ZMM_HI256_BIT =                                           6 //col:25459
XCR0_ZMM_HI256_FLAG =                                          0x40 //col:25460
XCR0_ZMM_HI256_MASK =                                          0x01 //col:25461
XCR0_ZMM_HI256(_) =                                            (((_) >> 6) & 0x01) //col:25462
XCR0_ZMM_HI16_BIT =                                            7 //col:25469
XCR0_ZMM_HI16_FLAG =                                           0x80 //col:25470
XCR0_ZMM_HI16_MASK =                                           0x01 //col:25471
XCR0_ZMM_HI16(_) =                                             (((_) >> 7) & 0x01) //col:25472
XCR0_PKRU_BIT =                                                9 //col:25479
XCR0_PKRU_FLAG =                                               0x200 //col:25480
XCR0_PKRU_MASK =                                               0x01 //col:25481
XCR0_PKRU(_) =                                                 (((_) >> 9) & 0x01) //col:25482
)
type typedef enum uint32
const(
typedef enum  = 1  //col:16410
return true
}

func (i *ia32)#define CR3_PAGE_LEVEL_WRITE_THROUGH()(ok bool){//col:311
return true
}

func (i *ia32)     * hardware support for a virtual interrupt flag ()(ok bool){//col:669
return true
}

func (i *ia32)#define CR8_TASK_PRIORITY_LEVEL()(ok bool){//col:701
return true
}

func (i *ia32) * other privilege level generates a general-protection exception ()(ok bool){//col:807
return true
}

func (i *ia32)     * [Bit 0] Enables ()(ok bool){//col:998
return true
}

func (i *ia32) * - EBX <- 756e6547h ()(ok bool){//col:1052
return true
}

func (i *ia32)#define CPUID_VERSION_INFORMATION_STEPPING_ID()(ok bool){//col:1896
return true
}

func (i *ia32) * describe a set of deterministic cache parameters ()(ok bool){//col:2101
return true
}

func (i *ia32)       * [Bits 15:0] Smallest monitor-line size in bytes ()(ok bool){//col:2257
return true
}

func (i *ia32)#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED()(ok bool){//col:2531
return true
}

func (i *ia32) * When CPUID executes with EAX set to 07H and the input value of ECX is invalid ()(ok bool){//col:3239
return true
}

func (i *ia32)       * [Bits 31:0] Value of bits [31:0] of IA32_PLATFORM_DCA_CAP MSR ()(ok bool){//col:3318
return true
}

func (i *ia32)#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING()(ok bool){//col:3504
return true
}

func (i *ia32)#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT()(ok bool){//col:3619
return true
}

func (i *ia32) * When CPUID executes with EAX set to 0DH and ECX = n ()(ok bool){//col:3780
return true
}

func (i *ia32)#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR()(ok bool){//col:3943
return true
}

func (i *ia32) *       If ECX contains an invalid sub-leaf index, EAX/EBX/ECX/EDX return 0. Sub-leaf n ()(ok bool){//col:4040
return true
}

func (i *ia32) * resource type if the bit is set. The bit position corresponds to the sub-leaf index ()(ok bool){//col:4139
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:4235
return true
}

func (i *ia32) * corresponds to a specific resource type if the bit is set. The bit position corresponds to the sub-leaf index ()(ok bool){//col:4352
return true
}

func (i *ia32)#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK()(ok bool){//col:4434
return true
}

func (i *ia32)#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK()(ok bool){//col:4513
return true
}

func (i *ia32)#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE()(ok bool){//col:4595
return true
}

func (i *ia32) * When CPUID executes with EAX set to 12H and ECX = n ()(ok bool){//col:4726
return true
}

func (i *ia32)#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0()(ok bool){//col:4803
return true
}

func (i *ia32)#define CPUID_EAX_SUB_LEAF_TYPE()(ok bool){//col:4882
return true
}

func (i *ia32)       * Enclave Page Cache ()(ok bool){//col:4986
return true
}

func (i *ia32) * When CPUID executes with EAX set to 14H and ECX = n ()(ok bool){//col:5195
return true
}

func (i *ia32)#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING()(ok bool){//col:5289
return true
}

func (i *ia32)#define CPUID_EAX_DENOMINATOR()(ok bool){//col:5378
return true
}

func (i *ia32)       * [Bits 15:0] Processor Base Frequency ()(ok bool){//col:5468
return true
}

func (i *ia32)#define CPUID_EAX_MAX_SOC_ID_INDEX()(ok bool){//col:5566
return true
}

func (i *ia32)#define CPUID_EAX_SOC_VENDOR_BRAND_STRING()(ok bool){//col:5645
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:5722
return true
}

func (i *ia32)#define CPUID_EAX_MAX_SUB_LEAF()(ok bool){//col:5903
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:6071
return true
}

func (i *ia32)#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS()(ok bool){//col:6155
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:6295
return true
}

func (i *ia32)#define CPUID_EAX_PROCESSOR_BRAND_STRING_1()(ok bool){//col:6382
return true
}

func (i *ia32)#define CPUID_EAX_PROCESSOR_BRAND_STRING_5()(ok bool){//col:6457
return true
}

func (i *ia32)#define CPUID_EAX_PROCESSOR_BRAND_STRING_9()(ok bool){//col:6532
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:6607
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:6711
return true
}

func (i *ia32)#define CPUID_EAX_RESERVED()(ok bool){//col:6791
return true
}

func (i *ia32)#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS()(ok bool){//col:6881
return true
}

func (i *ia32) *           IA32_P5_MC_()(ok bool){//col:6981
return true
}

func (i *ia32)#define IA32_APIC_BASE_BSP_FLAG()(ok bool){//col:7038
return true
}

func (i *ia32)     * [Bit 0] When set, locks this MSR from being written; writes to this bit will result in GP()(ok bool){//col:7170
return true
}

func (i *ia32)     * [Bit 0] IBRS: Indirect Branch Restricted Speculation ()(ok bool){//col:7239
return true
}

func (i *ia32)     * [Bit 0] IBPB: Indirect Branch Prediction Barrier ()(ok bool){//col:7266
return true
}

func (i *ia32)#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED()(ok bool){//col:7320
return true
}

func (i *ia32) *           IA32_SGXLEPUBKEYHASH[()(ok bool){//col:7402
return true
}

func (i *ia32) *           IA32_PMC()(ok bool){//col:7490
return true
}

func (i *ia32)#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT()(ok bool){//col:7572
return true
}

func (i *ia32)     * [Bit 0] RDCL_NO: The processor is not susceptible to Rogue Data Cache Load ()(ok bool){//col:7671
return true
}

func (i *ia32)#define IA32_FLUSH_CMD_L1D_FLUSH()(ok bool){//col:7698
return true
}

func (i *ia32) *          Available when CPUID.ARCH_CAP()(ok bool){//col:7734
return true
}

func (i *ia32)#define IA32_SYSENTER_CS_CS_SELECTOR()(ok bool){//col:7784
return true
}

func (i *ia32) * The value of this MSR is loaded into RSP ()(ok bool){//col:7914
return true
}

func (i *ia32)#define IA32_MCG_STATUS_RIPV()(ok bool){//col:7972
return true
}

func (i *ia32) *           IA32_PERFEVTSEL()(ok bool){//col:8103
return true
}

func (i *ia32)#define IA32_PERF_STATUS_STATE_VALUE()(ok bool){//col:8133
return true
}

func (i *ia32) * Performance Control. Software makes a request for a new Performance state ()(ok bool){//col:8173
return true
}

func (i *ia32)#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE()(ok bool){//col:8227
return true
}

func (i *ia32)#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE()(ok bool){//col:8358
return true
}

func (i *ia32)#define IA32_THERM_STATUS_THERMAL_STATUS()(ok bool){//col:8588
return true
}

func (i *ia32)     * [Bit 0] When set, the fast-strings feature ()(ok bool){//col:8766
return true
}

func (i *ia32)#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE()(ok bool){//col:8794
return true
}

func (i *ia32)#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS()(ok bool){//col:8931
return true
}

func (i *ia32)#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE()(ok bool){//col:9034
return true
}

func (i *ia32)#define IA32_DEBUGCTL_LBR()(ok bool){//col:9186
return true
}

func (i *ia32)#define IA32_SMRR_PHYSBASE_TYPE()(ok bool){//col:9225
return true
}

func (i *ia32)#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK()(ok bool){//col:9263
return true
}

func (i *ia32)#define IA32_DCA_0_CAP_DCA_ACTIVE()(ok bool){//col:9359
return true
}

func (i *ia32) *           IA32_MTRR_PHYSBASE()(ok bool){//col:9399
return true
}

func (i *ia32) *           IA32_MTRR_PHYSMASK()(ok bool){//col:9461
return true
}

func (i *ia32) *           IA32_MTRR_FIX()(ok bool){//col:9650
return true
}

func (i *ia32) *           IA32_MC()(ok bool){//col:9720
return true
}

func (i *ia32)#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE()(ok bool){//col:9768
return true
}

func (i *ia32) *           IA32_FIXED_CTR()(ok bool){//col:9865
return true
}

func (i *ia32)#define IA32_FIXED_CTR_CTRL_EN0_OS()(ok bool){//col:10002
return true
}

func (i *ia32)#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0()(ok bool){//col:10179
return true
}

func (i *ia32)     * [Bits 31:0] EN_PMC()(ok bool){//col:10219
return true
}

func (i *ia32)     * [Bits 31:0] Set 1 to clear Ovf_PMC()(ok bool){//col:10336
return true
}

func (i *ia32)     * [Bits 31:0] Set 1 to cause Ovf_PMC()(ok bool){//col:10443
return true
}

func (i *ia32)     * [Bits 31:0] IA32_PERFEVTSEL()(ok bool){//col:10488
return true
}

func (i *ia32)#define IA32_PEBS_ENABLE_ENABLE_PEBS()(ok bool){//col:10534
return true
}

func (i *ia32) *           IA32_MC()(ok bool){//col:10835
return true
}

func (i *ia32)     * guest interrupt-descriptor table ()(ok bool){//col:10925
return true
}

func (i *ia32)#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING()(ok bool){//col:11225
return true
}

func (i *ia32)#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS()(ok bool){//col:11429
return true
}

func (i *ia32)#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS()(ok bool){//col:11594
return true
}

func (i *ia32)     * timestamp counter ()(ok bool){//col:11758
return true
}

func (i *ia32)#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE()(ok bool){//col:11849
return true
}

func (i *ia32)#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES()(ok bool){//col:12211
return true
}

func (i *ia32)     * configure EPT paging-structure entries in which bits 1:0 are clear ()(ok bool){//col:12428
return true
}

func (i *ia32) *           IA32_VMX_TRUE_()(ok bool){//col:12478
return true
}

func (i *ia32)#define IA32_VMX_VMFUNC_EPTP_SWITCHING()(ok bool){//col:12511
return true
}

func (i *ia32)#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING()(ok bool){//col:12578
return true
}

func (i *ia32)#define IA32_VMX_EXIT_CTLS2_RESERVED()(ok bool){//col:12601
return true
}

func (i *ia32) *           IA32_A_PMC()(ok bool){//col:12645
return true
}

func (i *ia32)#define IA32_SGX_SVN_STATUS_LOCK()(ok bool){//col:12698
return true
}

func (i *ia32) *          ()(ok bool){//col:12738
return true
}

func (i *ia32) *          ()(ok bool){//col:12807
return true
}

func (i *ia32)     * Note that the processor will clear this bit on \#SMI ()(ok bool){//col:13179
return true
}

func (i *ia32)#define IA32_RTIT_STATUS_FILTER_ENABLED()(ok bool){//col:13319
return true
}

func (i *ia32)#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH()(ok bool){//col:13350
return true
}

func (i *ia32) *           IA32_RTIT_ADDR()(ok bool){//col:13422
return true
}

func (i *ia32) *          CPUID.()(ok bool){//col:13560
return true
}

func (i *ia32) *          CPUID.()(ok bool){//col:13681
return true
}

func (i *ia32) * ()(ok bool){//col:13770
return true
}

func (i *ia32)#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE()(ok bool){//col:13834
return true
}

func (i *ia32)#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE()(ok bool){//col:13910
return true
}

func (i *ia32)#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE()(ok bool){//col:13950
return true
}

func (i *ia32)#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE()(ok bool){//col:14038
return true
}

func (i *ia32)#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE()(ok bool){//col:14079
return true
}

func (i *ia32) *           IA32_X2APIC_ISR()(ok bool){//col:14340
return true
}

func (i *ia32)     * [Bit 0] Set 1 to enable L3 CAT masks and COS to operate in Code and Data Prioritization ()(ok bool){//col:14367
return true
}

func (i *ia32)     * [Bit 0] Set 1 to enable L2 CAT masks and COS to operate in Code and Data Prioritization ()(ok bool){//col:14394
return true
}

func (i *ia32)#define IA32_QM_EVTSEL_EVENT_ID()(ok bool){//col:14434
return true
}

func (i *ia32)#define IA32_QM_CTR_RESOURCE_MONITORED_DATA()(ok bool){//col:14480
return true
}

func (i *ia32)#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID()(ok bool){//col:14521
return true
}

func (i *ia32)#define IA32_BNDCFGS_ENABLE()(ok bool){//col:14564
return true
}

func (i *ia32)#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE()(ok bool){//col:14591
return true
}

func (i *ia32)#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE()(ok bool){//col:14621
return true
}

func (i *ia32)#define IA32_PM_CTL1_HDC_ALLOW_BLOCK()(ok bool){//col:14651
return true
}

func (i *ia32)#define IA32_EFER_SYSCALL_ENABLE()(ok bool){//col:14731
return true
}

func (i *ia32)#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE()(ok bool){//col:14809
return true
}

func (i *ia32)#define PDE_4MB_32_PRESENT()(ok bool){//col:14979
return true
}

func (i *ia32)#define PDE_32_PRESENT()(ok bool){//col:15092
return true
}

func (i *ia32)#define PTE_32_PRESENT()(ok bool){//col:15220
return true
}

func (i *ia32)#define PT_ENTRY_32_PRESENT()(ok bool){//col:15295
return true
}

func (i *ia32) *           64-Bit ()(ok bool){//col:15466
return true
}

func (i *ia32)#define PDPTE_1GB_64_PRESENT()(ok bool){//col:15650
return true
}

func (i *ia32)#define PDPTE_64_PRESENT()(ok bool){//col:15789
return true
}

func (i *ia32)#define PDE_2MB_64_PRESENT()(ok bool){//col:15973
return true
}

func (i *ia32)#define PDE_64_PRESENT()(ok bool){//col:16112
return true
}

func (i *ia32)#define PTE_64_PRESENT()(ok bool){//col:16285
return true
}

func (i *ia32)#define PT_ENTRY_64_PRESENT()(ok bool){//col:16385
return true
}

func (i *ia32)   * descriptor.2 ()(ok bool){//col:16438
return true
}

func (i *ia32)#define INVPCID_DESCRIPTOR_PCID()(ok bool){//col:16466
return true
}

func (i *ia32)#pragma pack()(ok bool){//col:16490
return true
}

func (i *ia32)#pragma pack()(ok bool){//col:16510
return true
}

func (i *ia32)#pragma pack()(ok bool){//col:16649
return true
}

func (i *ia32)   * processor interprets the segment limit in one of two ways, depending on the setting of the G ()(ok bool){//col:16859
return true
}

func (i *ia32)   * processor interprets the segment limit in one of two ways, depending on the setting of the G ()(ok bool){//col:17071
return true
}

func (i *ia32)   * Offset to procedure entry point ()(ok bool){//col:17166
return true
}

func (i *ia32) * When the S ()(ok bool){//col:17431
return true
}

func (i *ia32)#pragma pack()(ok bool){//col:17519
return true
}

func (i *ia32)#pragma pack()(ok bool){//col:18224
return true
}

func (i *ia32)#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION()(ok bool){//col:18285
return true
}

func (i *ia32)     * [Bits 15:0] Selector of task-state segment ()(ok bool){//col:18320
return true
}

func (i *ia32)     * [Bits 3:0] Number of control register ()(ok bool){//col:18407
return true
}

func (i *ia32)#define VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER()(ok bool){//col:18456
return true
}

func (i *ia32)#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS()(ok bool){//col:18534
return true
}

func (i *ia32)#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET()(ok bool){//col:18594
return true
}

func (i *ia32)#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS()(ok bool){//col:18790
return true
}

func (i *ia32)     * 2: 64-bit ()(ok bool){//col:18849
return true
}

func (i *ia32)     * 3: scale by 8 ()(ok bool){//col:18955
return true
}

func (i *ia32)     * 3: scale by 8 ()(ok bool){//col:19080
return true
}

func (i *ia32)     * 3: scale by 8 ()(ok bool){//col:19210
return true
}

func (i *ia32)#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER()(ok bool){//col:19248
return true
}

func (i *ia32)     * 3: scale by 8 ()(ok bool){//col:19345
return true
}

func (i *ia32)     * 3: scale by 8 ()(ok bool){//col:19471
return true
}

func (i *ia32) *        ()(ok bool){//col:19582
return true
}

func (i *ia32)#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI()(ok bool){//col:19661
return true
}

func (i *ia32)   * The logical processor is inactive because it is waiting for a startup-IPI ()(ok bool){//col:19684
return true
}

func (i *ia32)#define VMX_PENDING_DEBUG_EXCEPTIONS_B0()(ok bool){//col:19770
return true
}

func (i *ia32) * Exit reason ()(ok bool){//col:19856
return true
}

func (i *ia32) * The extended page-table mechanism ()(ok bool){//col:19966
return true
}

func (i *ia32) * extended-page-table pointer ()(ok bool){//col:20055
return true
}

func (i *ia32)#define EPT_PDPTE_1GB_READ_ACCESS()(ok bool){//col:20225
return true
}

func (i *ia32)#define EPT_PDPTE_READ_ACCESS()(ok bool){//col:20302
return true
}

func (i *ia32)#define EPT_PDE_2MB_READ_ACCESS()(ok bool){//col:20472
return true
}

func (i *ia32)#define EPT_PDE_READ_ACCESS()(ok bool){//col:20549
return true
}

func (i *ia32)#define EPT_PTE_READ_ACCESS()(ok bool){//col:20725
return true
}

func (i *ia32)#define EPT_ENTRY_READ_ACCESS()(ok bool){//col:20794
return true
}

func (i *ia32)   * all PCIDs. ()(ok bool){//col:20844
return true
}

func (i *ia32)   * and, for combined mappings, all EP4TAs. ()(ok bool){//col:20882
return true
}

func (i *ia32) * The hypervisor-managed linear-address translation pointer ()(ok bool){//col:20957
return true
}

func (i *ia32) * A logical processor uses virtual-machine control data structures ()(ok bool){//col:21027
return true
}

func (i *ia32)     * ()(ok bool){//col:21074
return true
}

func (i *ia32) *           VMCS ()(ok bool){//col:21147
return true
}

func (i *ia32) * Virtual-processor identifier ()(ok bool){//col:22264
return true
}

func (i *ia32)#define VMENTRY_INTERRUPT_INFORMATION_VECTOR()(ok bool){//col:22324
return true
}

func (i *ia32)#define VMEXIT_INTERRUPT_INFORMATION_VECTOR()(ok bool){//col:22384
return true
}

func (i *ia32) *           Advanced Programmable Interrupt Controller ()(ok bool){//col:22889
return true
}

func (i *ia32)#define RFLAGS_CARRY_FLAG()(ok bool){//col:23106
return true
}

func (i *ia32)#define CONTROL_PROTECTION_EXCEPTION_CPEC()(ok bool){//col:23148
return true
}

func (i *ia32) *        Each exception is given a mnemonic that consists of a pound sign ()(ok bool){//col:23309
return true
}

func (i *ia32) *        error code onto the stack of the exception handler ()(ok bool){//col:23371
return true
}

func (i *ia32)#define PAGE_FAULT_EXCEPTION_PRESENT()(ok bool){//col:23497
return true
}

func (i *ia32) * of system memory, it allows the type of caching ()(ok bool){//col:23652
return true
}

func (i *ia32)       * * 0: Indicates the context-entry is not present. All other fields except Fault Processing Disable ()(ok bool){//col:23783
return true
}

func (i *ia32)#define VTD_VERSION_MINOR()(ok bool){//col:23835
return true
}

func (i *ia32)#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED()(ok bool){//col:24192
return true
}

func (i *ia32)     * structures are coherent ()(ok bool){//col:24658
return true
}

func (i *ia32) * 3. if ()(ok bool){//col:24849
return true
}

func (i *ia32)     * [Bit 23] This field indicates the status of Compatibility format interrupts on Intel()(ok bool){//col:24991
return true
}

func (i *ia32) * SRTP field in the Global Command Register ()(ok bool){//col:25043
return true
}

func (i *ia32)     * to this field is within this limit. Hardware ignores ()(ok bool){//col:25179
return true
}

func (i *ia32)#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK()(ok bool){//col:25250
return true
}

func (i *ia32)     * to this field is within this limit. Hardware may ignore and not implement bits 47:()(ok bool){//col:25385
return true
}

func (i *ia32)#define XCR0_X87()(ok bool){//col:25487
return true
}

