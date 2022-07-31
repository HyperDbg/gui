package out


const(
CPUID_SIGNATURE =                                              0x00000000 //col:154
CPUID_VERSION_INFO =                                           0x00000001 //col:162
CPUID_CACHE_PARAMS =                                           0x00000004 //col:268
CPUID_MONITOR_MWAIT =                                          0x00000005 //col:314
CPUID_THERMAL_POWER_MANAGEMENT =                               0x00000006 //col:358
CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS =                      0x00000007 //col:420
CPUID_DIRECT_CACHE_ACCESS_INFO =                               0x00000009 //col:535
CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING =                   0x0000000A //col:571
CPUID_EXTENDED_TOPOLOGY =                                      0x0000000B //col:619
CPUID_EXTENDED_STATE =                                         0x0000000D //col:661
CPUID_INTEL_RDT_MONITORING =                                   0x0000000F //col:797
CPUID_INTEL_RDT_ALLOCATION =                                   0x00000010 //col:880
CPUID_INTEL_SGX =                                              0x00000012 //col:1035
CPUID_INTEL_PROCESSOR_TRACE =                                  0x00000014 //col:1194
CPUID_TIME_STAMP_COUNTER =                                     0x00000015 //col:1285
CPUID_PROCESSOR_FREQUENCY =                                    0x00000016 //col:1321
CPUID_SOC_VENDOR =                                             0x00000017 //col:1362
CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS =           0x00000018 //col:1478
CPUID_EXTENDED_FUNCTION =                                      0x80000000 //col:1575
CPUID_EXTENDED_CPU_SIGNATURE =                                 0x80000001 //col:1611
CPUID_BRAND_STRING1 =                                          0x80000002 //col:1659
CPUID_BRAND_STRING2 =                                          0x80000003 //col:1660
CPUID_BRAND_STRING3 =                                          0x80000004 //col:1661
CPUID_EXTENDED_CACHE_INFO =                                    0x80000006 //col:1802
CPUID_EXTENDED_TIME_STAMP_COUNTER =                            0x80000007 //col:1841
CPUID_EXTENDED_VIRT_PHYS_ADDRESS_SIZE =                        0x80000008 //col:1878
IA32_P5_MC_ADDR =                                              0x00000000 //col:1929
IA32_P5_MC_TYPE =                                              0x00000001 //col:1930
IA32_MONITOR_FILTER_SIZE =                                     0x00000006 //col:1935
IA32_TIME_STAMP_COUNTER =                                      0x00000010 //col:1936
IA32_PLATFORM_ID =                                             0x00000017 //col:1937
IA32_APIC_BASE =                                               0x0000001B //col:1947
IA32_FEATURE_CONTROL =                                         0x0000003A //col:1961
IA32_TSC_ADJUST =                                              0x0000003B //col:1980
IA32_SPEC_CTRL =                                               0x00000048 //col:1985
IA32_PRED_CMD =                                                0x00000049 //col:1996
IA32_BIOS_UPDT_TRIG =                                          0x00000079 //col:2005
IA32_BIOS_SIGN_ID =                                            0x0000008B //col:2006
IA32_SGXLEPUBKEYHASH0 =                                        0x0000008C //col:2021
IA32_SGXLEPUBKEYHASH1 =                                        0x0000008D //col:2022
IA32_SGXLEPUBKEYHASH2 =                                        0x0000008E //col:2023
IA32_SGXLEPUBKEYHASH3 =                                        0x0000008F //col:2024
IA32_SMM_MONITOR_CTL =                                         0x0000009B //col:2029
IA32_STM_FEATURES_IA32E =                                      0x00000001 //col:2046
IA32_SMBASE =                                                  0x0000009E //col:2055
IA32_PMC0 =                                                    0x000000C1 //col:2061
IA32_PMC1 =                                                    0x000000C2 //col:2062
IA32_PMC2 =                                                    0x000000C3 //col:2063
IA32_PMC3 =                                                    0x000000C4 //col:2064
IA32_PMC4 =                                                    0x000000C5 //col:2065
IA32_PMC5 =                                                    0x000000C6 //col:2066
IA32_PMC6 =                                                    0x000000C7 //col:2067
IA32_PMC7 =                                                    0x000000C8 //col:2068
IA32_MPERF =                                                   0x000000E7 //col:2073
IA32_APERF =                                                   0x000000E8 //col:2078
IA32_MTRRCAP =                                                 0x000000FE //col:2083
IA32_ARCH_CAPABILITIES =                                       0x0000010A //col:2096
IA32_FLUSH_CMD =                                               0x0000010B //col:2113
IA32_TSX_CTRL =                                                0x00000122 //col:2122
IA32_SYSENTER_CS =                                             0x00000174 //col:2132
IA32_SYSENTER_ESP =                                            0x00000175 //col:2143
IA32_SYSENTER_EIP =                                            0x00000176 //col:2144
IA32_MCG_CAP =                                                 0x00000179 //col:2145
IA32_MCG_STATUS =                                              0x0000017A //col:2164
IA32_MCG_CTL =                                                 0x0000017B //col:2176
IA32_PERFEVTSEL0 =                                             0x00000186 //col:2182
IA32_PERFEVTSEL1 =                                             0x00000187 //col:2183
IA32_PERFEVTSEL2 =                                             0x00000188 //col:2184
IA32_PERFEVTSEL3 =                                             0x00000189 //col:2185
IA32_PERF_STATUS =                                             0x00000198 //col:2208
IA32_PERF_CTL =                                                0x00000199 //col:2217
IA32_CLOCK_MODULATION =                                        0x0000019A //col:2228
IA32_THERM_INTERRUPT =                                         0x0000019B //col:2239
IA32_THERM_STATUS =                                            0x0000019C //col:2258
IA32_MISC_ENABLE =                                             0x000001A0 //col:2286
IA32_ENERGY_PERF_BIAS =                                        0x000001B0 //col:2311
IA32_PACKAGE_THERM_STATUS =                                    0x000001B1 //col:2320
IA32_PACKAGE_THERM_INTERRUPT =                                 0x000001B2 //col:2342
IA32_DEBUGCTL =                                                0x000001D9 //col:2361
IA32_SMRR_PHYSBASE =                                           0x000001F2 //col:2382
IA32_SMRR_PHYSMASK =                                           0x000001F3 //col:2393
IA32_PLATFORM_DCA_CAP =                                        0x000001F8 //col:2404
IA32_CPU_DCA_CAP =                                             0x000001F9 //col:2405
IA32_DCA_0_CAP =                                               0x000001FA //col:2406
IA32_MTRR_PHYSBASE0 =                                          0x00000200 //col:2439
IA32_MTRR_PHYSBASE1 =                                          0x00000202 //col:2440
IA32_MTRR_PHYSBASE2 =                                          0x00000204 //col:2441
IA32_MTRR_PHYSBASE3 =                                          0x00000206 //col:2442
IA32_MTRR_PHYSBASE4 =                                          0x00000208 //col:2443
IA32_MTRR_PHYSBASE5 =                                          0x0000020A //col:2444
IA32_MTRR_PHYSBASE6 =                                          0x0000020C //col:2445
IA32_MTRR_PHYSBASE7 =                                          0x0000020E //col:2446
IA32_MTRR_PHYSBASE8 =                                          0x00000210 //col:2447
IA32_MTRR_PHYSBASE9 =                                          0x00000212 //col:2448
IA32_MTRR_PHYSMASK0 =                                          0x00000201 //col:2468
IA32_MTRR_PHYSMASK1 =                                          0x00000203 //col:2469
IA32_MTRR_PHYSMASK2 =                                          0x00000205 //col:2470
IA32_MTRR_PHYSMASK3 =                                          0x00000207 //col:2471
IA32_MTRR_PHYSMASK4 =                                          0x00000209 //col:2472
IA32_MTRR_PHYSMASK5 =                                          0x0000020B //col:2473
IA32_MTRR_PHYSMASK6 =                                          0x0000020D //col:2474
IA32_MTRR_PHYSMASK7 =                                          0x0000020F //col:2475
IA32_MTRR_PHYSMASK8 =                                          0x00000211 //col:2476
IA32_MTRR_PHYSMASK9 =                                          0x00000213 //col:2477
IA32_MTRR_FIX64K_BASE =                                        0x00000000 //col:2492
IA32_MTRR_FIX64K_SIZE =                                        0x00010000 //col:2493
IA32_MTRR_FIX64K_00000 =                                       0x00000250 //col:2494
IA32_MTRR_FIX16K_BASE =                                        0x00080000 //col:2504
IA32_MTRR_FIX16K_SIZE =                                        0x00004000 //col:2505
IA32_MTRR_FIX16K_80000 =                                       0x00000258 //col:2506
IA32_MTRR_FIX16K_A0000 =                                       0x00000259 //col:2507
IA32_MTRR_FIX4K_BASE =                                         0x000C0000 //col:2517
IA32_MTRR_FIX4K_SIZE =                                         0x00001000 //col:2518
IA32_MTRR_FIX4K_C0000 =                                        0x00000268 //col:2519
IA32_MTRR_FIX4K_C8000 =                                        0x00000269 //col:2520
IA32_MTRR_FIX4K_D0000 =                                        0x0000026A //col:2521
IA32_MTRR_FIX4K_D8000 =                                        0x0000026B //col:2522
IA32_MTRR_FIX4K_E0000 =                                        0x0000026C //col:2523
IA32_MTRR_FIX4K_E8000 =                                        0x0000026D //col:2524
IA32_MTRR_FIX4K_F0000 =                                        0x0000026E //col:2525
IA32_MTRR_FIX4K_F8000 =                                        0x0000026F //col:2526
IA32_MTRR_FIX_COUNT =                                          ((1 + 2 + 8) * 8) //col:2531
IA32_MTRR_VARIABLE_COUNT =                                     0x0000000A //col:2532
IA32_MTRR_COUNT =                                              (IA32_MTRR_FIX_COUNT + IA32_MTRR_VARIABLE_COUNT) //col:2533
IA32_PAT =                                                     0x00000277 //col:2538
IA32_MC0_CTL2 =                                                0x00000280 //col:2566
IA32_MC1_CTL2 =                                                0x00000281 //col:2567
IA32_MC2_CTL2 =                                                0x00000282 //col:2568
IA32_MC3_CTL2 =                                                0x00000283 //col:2569
IA32_MC4_CTL2 =                                                0x00000284 //col:2570
IA32_MC5_CTL2 =                                                0x00000285 //col:2571
IA32_MC6_CTL2 =                                                0x00000286 //col:2572
IA32_MC7_CTL2 =                                                0x00000287 //col:2573
IA32_MC8_CTL2 =                                                0x00000288 //col:2574
IA32_MC9_CTL2 =                                                0x00000289 //col:2575
IA32_MC10_CTL2 =                                               0x0000028A //col:2576
IA32_MC11_CTL2 =                                               0x0000028B //col:2577
IA32_MC12_CTL2 =                                               0x0000028C //col:2578
IA32_MC13_CTL2 =                                               0x0000028D //col:2579
IA32_MC14_CTL2 =                                               0x0000028E //col:2580
IA32_MC15_CTL2 =                                               0x0000028F //col:2581
IA32_MC16_CTL2 =                                               0x00000290 //col:2582
IA32_MC17_CTL2 =                                               0x00000291 //col:2583
IA32_MC18_CTL2 =                                               0x00000292 //col:2584
IA32_MC19_CTL2 =                                               0x00000293 //col:2585
IA32_MC20_CTL2 =                                               0x00000294 //col:2586
IA32_MC21_CTL2 =                                               0x00000295 //col:2587
IA32_MC22_CTL2 =                                               0x00000296 //col:2588
IA32_MC23_CTL2 =                                               0x00000297 //col:2589
IA32_MC24_CTL2 =                                               0x00000298 //col:2590
IA32_MC25_CTL2 =                                               0x00000299 //col:2591
IA32_MC26_CTL2 =                                               0x0000029A //col:2592
IA32_MC27_CTL2 =                                               0x0000029B //col:2593
IA32_MC28_CTL2 =                                               0x0000029C //col:2594
IA32_MC29_CTL2 =                                               0x0000029D //col:2595
IA32_MC30_CTL2 =                                               0x0000029E //col:2596
IA32_MC31_CTL2 =                                               0x0000029F //col:2597
IA32_MTRR_DEF_TYPE =                                           0x000002FF //col:2612
IA32_FIXED_CTR0 =                                              0x00000309 //col:2629
IA32_FIXED_CTR1 =                                              0x0000030A //col:2630
IA32_FIXED_CTR2 =                                              0x0000030B //col:2631
IA32_PERF_CAPABILITIES =                                       0x00000345 //col:2636
IA32_FIXED_CTR_CTRL =                                          0x0000038D //col:2650
IA32_PERF_GLOBAL_STATUS =                                      0x0000038E //col:2670
IA32_PERF_GLOBAL_CTRL =                                        0x0000038F //col:2695
IA32_PERF_GLOBAL_STATUS_RESET =                                0x00000390 //col:2705
IA32_PERF_GLOBAL_STATUS_SET =                                  0x00000391 //col:2724
IA32_PERF_GLOBAL_INUSE =                                       0x00000392 //col:2742
IA32_PEBS_ENABLE =                                             0x000003F1 //col:2754
IA32_MC0_CTL =                                                 0x00000400 //col:2771
IA32_MC1_CTL =                                                 0x00000404 //col:2772
IA32_MC2_CTL =                                                 0x00000408 //col:2773
IA32_MC3_CTL =                                                 0x0000040C //col:2774
IA32_MC4_CTL =                                                 0x00000410 //col:2775
IA32_MC5_CTL =                                                 0x00000414 //col:2776
IA32_MC6_CTL =                                                 0x00000418 //col:2777
IA32_MC7_CTL =                                                 0x0000041C //col:2778
IA32_MC8_CTL =                                                 0x00000420 //col:2779
IA32_MC9_CTL =                                                 0x00000424 //col:2780
IA32_MC10_CTL =                                                0x00000428 //col:2781
IA32_MC11_CTL =                                                0x0000042C //col:2782
IA32_MC12_CTL =                                                0x00000430 //col:2783
IA32_MC13_CTL =                                                0x00000434 //col:2784
IA32_MC14_CTL =                                                0x00000438 //col:2785
IA32_MC15_CTL =                                                0x0000043C //col:2786
IA32_MC16_CTL =                                                0x00000440 //col:2787
IA32_MC17_CTL =                                                0x00000444 //col:2788
IA32_MC18_CTL =                                                0x00000448 //col:2789
IA32_MC19_CTL =                                                0x0000044C //col:2790
IA32_MC20_CTL =                                                0x00000450 //col:2791
IA32_MC21_CTL =                                                0x00000454 //col:2792
IA32_MC22_CTL =                                                0x00000458 //col:2793
IA32_MC23_CTL =                                                0x0000045C //col:2794
IA32_MC24_CTL =                                                0x00000460 //col:2795
IA32_MC25_CTL =                                                0x00000464 //col:2796
IA32_MC26_CTL =                                                0x00000468 //col:2797
IA32_MC27_CTL =                                                0x0000046C //col:2798
IA32_MC28_CTL =                                                0x00000470 //col:2799
IA32_MC0_STATUS =                                              0x00000401 //col:2809
IA32_MC1_STATUS =                                              0x00000405 //col:2810
IA32_MC2_STATUS =                                              0x00000409 //col:2811
IA32_MC3_STATUS =                                              0x0000040D //col:2812
IA32_MC4_STATUS =                                              0x00000411 //col:2813
IA32_MC5_STATUS =                                              0x00000415 //col:2814
IA32_MC6_STATUS =                                              0x00000419 //col:2815
IA32_MC7_STATUS =                                              0x0000041D //col:2816
IA32_MC8_STATUS =                                              0x00000421 //col:2817
IA32_MC9_STATUS =                                              0x00000425 //col:2818
IA32_MC10_STATUS =                                             0x00000429 //col:2819
IA32_MC11_STATUS =                                             0x0000042D //col:2820
IA32_MC12_STATUS =                                             0x00000431 //col:2821
IA32_MC13_STATUS =                                             0x00000435 //col:2822
IA32_MC14_STATUS =                                             0x00000439 //col:2823
IA32_MC15_STATUS =                                             0x0000043D //col:2824
IA32_MC16_STATUS =                                             0x00000441 //col:2825
IA32_MC17_STATUS =                                             0x00000445 //col:2826
IA32_MC18_STATUS =                                             0x00000449 //col:2827
IA32_MC19_STATUS =                                             0x0000044D //col:2828
IA32_MC20_STATUS =                                             0x00000451 //col:2829
IA32_MC21_STATUS =                                             0x00000455 //col:2830
IA32_MC22_STATUS =                                             0x00000459 //col:2831
IA32_MC23_STATUS =                                             0x0000045D //col:2832
IA32_MC24_STATUS =                                             0x00000461 //col:2833
IA32_MC25_STATUS =                                             0x00000465 //col:2834
IA32_MC26_STATUS =                                             0x00000469 //col:2835
IA32_MC27_STATUS =                                             0x0000046D //col:2836
IA32_MC28_STATUS =                                             0x00000471 //col:2837
IA32_MC0_ADDR =                                                0x00000402 //col:2847
IA32_MC1_ADDR =                                                0x00000406 //col:2848
IA32_MC2_ADDR =                                                0x0000040A //col:2849
IA32_MC3_ADDR =                                                0x0000040E //col:2850
IA32_MC4_ADDR =                                                0x00000412 //col:2851
IA32_MC5_ADDR =                                                0x00000416 //col:2852
IA32_MC6_ADDR =                                                0x0000041A //col:2853
IA32_MC7_ADDR =                                                0x0000041E //col:2854
IA32_MC8_ADDR =                                                0x00000422 //col:2855
IA32_MC9_ADDR =                                                0x00000426 //col:2856
IA32_MC10_ADDR =                                               0x0000042A //col:2857
IA32_MC11_ADDR =                                               0x0000042E //col:2858
IA32_MC12_ADDR =                                               0x00000432 //col:2859
IA32_MC13_ADDR =                                               0x00000436 //col:2860
IA32_MC14_ADDR =                                               0x0000043A //col:2861
IA32_MC15_ADDR =                                               0x0000043E //col:2862
IA32_MC16_ADDR =                                               0x00000442 //col:2863
IA32_MC17_ADDR =                                               0x00000446 //col:2864
IA32_MC18_ADDR =                                               0x0000044A //col:2865
IA32_MC19_ADDR =                                               0x0000044E //col:2866
IA32_MC20_ADDR =                                               0x00000452 //col:2867
IA32_MC21_ADDR =                                               0x00000456 //col:2868
IA32_MC22_ADDR =                                               0x0000045A //col:2869
IA32_MC23_ADDR =                                               0x0000045E //col:2870
IA32_MC24_ADDR =                                               0x00000462 //col:2871
IA32_MC25_ADDR =                                               0x00000466 //col:2872
IA32_MC26_ADDR =                                               0x0000046A //col:2873
IA32_MC27_ADDR =                                               0x0000046E //col:2874
IA32_MC28_ADDR =                                               0x00000472 //col:2875
IA32_MC0_MISC =                                                0x00000403 //col:2885
IA32_MC1_MISC =                                                0x00000407 //col:2886
IA32_MC2_MISC =                                                0x0000040B //col:2887
IA32_MC3_MISC =                                                0x0000040F //col:2888
IA32_MC4_MISC =                                                0x00000413 //col:2889
IA32_MC5_MISC =                                                0x00000417 //col:2890
IA32_MC6_MISC =                                                0x0000041B //col:2891
IA32_MC7_MISC =                                                0x0000041F //col:2892
IA32_MC8_MISC =                                                0x00000423 //col:2893
IA32_MC9_MISC =                                                0x00000427 //col:2894
IA32_MC10_MISC =                                               0x0000042B //col:2895
IA32_MC11_MISC =                                               0x0000042F //col:2896
IA32_MC12_MISC =                                               0x00000433 //col:2897
IA32_MC13_MISC =                                               0x00000437 //col:2898
IA32_MC14_MISC =                                               0x0000043B //col:2899
IA32_MC15_MISC =                                               0x0000043F //col:2900
IA32_MC16_MISC =                                               0x00000443 //col:2901
IA32_MC17_MISC =                                               0x00000447 //col:2902
IA32_MC18_MISC =                                               0x0000044B //col:2903
IA32_MC19_MISC =                                               0x0000044F //col:2904
IA32_MC20_MISC =                                               0x00000453 //col:2905
IA32_MC21_MISC =                                               0x00000457 //col:2906
IA32_MC22_MISC =                                               0x0000045B //col:2907
IA32_MC23_MISC =                                               0x0000045F //col:2908
IA32_MC24_MISC =                                               0x00000463 //col:2909
IA32_MC25_MISC =                                               0x00000467 //col:2910
IA32_MC26_MISC =                                               0x0000046B //col:2911
IA32_MC27_MISC =                                               0x0000046F //col:2912
IA32_MC28_MISC =                                               0x00000473 //col:2913
IA32_VMX_BASIC =                                               0x00000480 //col:2918
IA32_VMX_PINBASED_CTLS =                                       0x00000481 //col:2935
IA32_VMX_PROCBASED_CTLS =                                      0x00000482 //col:2950
IA32_VMX_EXIT_CTLS =                                           0x00000483 //col:2986
IA32_VMX_ENTRY_CTLS =                                          0x00000484 //col:3017
IA32_VMX_MISC =                                                0x00000485 //col:3042
IA32_VMX_CR0_FIXED0 =                                          0x00000486 //col:3063
IA32_VMX_CR0_FIXED1 =                                          0x00000487 //col:3064
IA32_VMX_CR4_FIXED0 =                                          0x00000488 //col:3065
IA32_VMX_CR4_FIXED1 =                                          0x00000489 //col:3066
IA32_VMX_VMCS_ENUM =                                           0x0000048A //col:3067
IA32_VMX_PROCBASED_CTLS2 =                                     0x0000048B //col:3080
IA32_VMX_EPT_VPID_CAP =                                        0x0000048C //col:3117
IA32_VMX_TRUE_PINBASED_CTLS =                                  0x0000048D //col:3157
IA32_VMX_TRUE_PROCBASED_CTLS =                                 0x0000048E //col:3158
IA32_VMX_TRUE_EXIT_CTLS =                                      0x0000048F //col:3159
IA32_VMX_TRUE_ENTRY_CTLS =                                     0x00000490 //col:3160
IA32_VMX_VMFUNC =                                              0x00000491 //col:3174
IA32_VMX_PROCBASED_CTLS3 =                                     0x00000492 //col:3183
IA32_VMX_EXIT_CTLS2 =                                          0x00000493 //col:3195
IA32_A_PMC0 =                                                  0x000004C1 //col:3209
IA32_A_PMC1 =                                                  0x000004C2 //col:3210
IA32_A_PMC2 =                                                  0x000004C3 //col:3211
IA32_A_PMC3 =                                                  0x000004C4 //col:3212
IA32_A_PMC4 =                                                  0x000004C5 //col:3213
IA32_A_PMC5 =                                                  0x000004C6 //col:3214
IA32_A_PMC6 =                                                  0x000004C7 //col:3215
IA32_A_PMC7 =                                                  0x000004C8 //col:3216
IA32_MCG_EXT_CTL =                                             0x000004D0 //col:3221
IA32_SGX_SVN_STATUS =                                          0x00000500 //col:3230
IA32_RTIT_OUTPUT_BASE =                                        0x00000560 //col:3241
IA32_RTIT_OUTPUT_MASK_PTRS =                                   0x00000561 //col:3251
IA32_RTIT_CTL =                                                0x00000570 //col:3262
IA32_RTIT_STATUS =                                             0x00000571 //col:3296
IA32_RTIT_CR3_MATCH =                                          0x00000572 //col:3314
IA32_RTIT_ADDR0_A =                                            0x00000580 //col:3334
IA32_RTIT_ADDR1_A =                                            0x00000582 //col:3335
IA32_RTIT_ADDR2_A =                                            0x00000584 //col:3336
IA32_RTIT_ADDR3_A =                                            0x00000586 //col:3337
IA32_RTIT_ADDR0_B =                                            0x00000581 //col:3347
IA32_RTIT_ADDR1_B =                                            0x00000583 //col:3348
IA32_RTIT_ADDR2_B =                                            0x00000585 //col:3349
IA32_RTIT_ADDR3_B =                                            0x00000587 //col:3350
IA32_DS_AREA =                                                 0x00000600 //col:3368
IA32_U_CET =                                                   0x000006A0 //col:3369
IA32_S_CET =                                                   0x000006A2 //col:3387
IA32_PL0_SSP =                                                 0x000006A4 //col:3405
IA32_PL1_SSP =                                                 0x000006A5 //col:3406
IA32_PL2_SSP =                                                 0x000006A6 //col:3407
IA32_PL3_SSP =                                                 0x000006A7 //col:3408
IA32_INTERRUPT_SSP_TABLE_ADDR =                                0x000006A8 //col:3409
IA32_TSC_DEADLINE =                                            0x000006E0 //col:3410
IA32_PM_ENABLE =                                               0x00000770 //col:3411
IA32_HWP_CAPABILITIES =                                        0x00000771 //col:3420
IA32_HWP_REQUEST_PKG =                                         0x00000772 //col:3432
IA32_HWP_INTERRUPT =                                           0x00000773 //col:3445
IA32_HWP_REQUEST =                                             0x00000774 //col:3455
IA32_HWP_STATUS =                                              0x00000777 //col:3469
IA32_X2APIC_APICID =                                           0x00000802 //col:3480
IA32_X2APIC_VERSION =                                          0x00000803 //col:3481
IA32_X2APIC_TPR =                                              0x00000808 //col:3482
IA32_X2APIC_PPR =                                              0x0000080A //col:3483
IA32_X2APIC_EOI =                                              0x0000080B //col:3484
IA32_X2APIC_LDR =                                              0x0000080D //col:3485
IA32_X2APIC_SIVR =                                             0x0000080F //col:3486
IA32_X2APIC_ISR0 =                                             0x00000810 //col:3492
IA32_X2APIC_ISR1 =                                             0x00000811 //col:3493
IA32_X2APIC_ISR2 =                                             0x00000812 //col:3494
IA32_X2APIC_ISR3 =                                             0x00000813 //col:3495
IA32_X2APIC_ISR4 =                                             0x00000814 //col:3496
IA32_X2APIC_ISR5 =                                             0x00000815 //col:3497
IA32_X2APIC_ISR6 =                                             0x00000816 //col:3498
IA32_X2APIC_ISR7 =                                             0x00000817 //col:3499
IA32_X2APIC_TMR0 =                                             0x00000818 //col:3509
IA32_X2APIC_TMR1 =                                             0x00000819 //col:3510
IA32_X2APIC_TMR2 =                                             0x0000081A //col:3511
IA32_X2APIC_TMR3 =                                             0x0000081B //col:3512
IA32_X2APIC_TMR4 =                                             0x0000081C //col:3513
IA32_X2APIC_TMR5 =                                             0x0000081D //col:3514
IA32_X2APIC_TMR6 =                                             0x0000081E //col:3515
IA32_X2APIC_TMR7 =                                             0x0000081F //col:3516
IA32_X2APIC_IRR0 =                                             0x00000820 //col:3526
IA32_X2APIC_IRR1 =                                             0x00000821 //col:3527
IA32_X2APIC_IRR2 =                                             0x00000822 //col:3528
IA32_X2APIC_IRR3 =                                             0x00000823 //col:3529
IA32_X2APIC_IRR4 =                                             0x00000824 //col:3530
IA32_X2APIC_IRR5 =                                             0x00000825 //col:3531
IA32_X2APIC_IRR6 =                                             0x00000826 //col:3532
IA32_X2APIC_IRR7 =                                             0x00000827 //col:3533
IA32_X2APIC_ESR =                                              0x00000828 //col:3538
IA32_X2APIC_LVT_CMCI =                                         0x0000082F //col:3539
IA32_X2APIC_ICR =                                              0x00000830 //col:3540
IA32_X2APIC_LVT_TIMER =                                        0x00000832 //col:3541
IA32_X2APIC_LVT_THERMAL =                                      0x00000833 //col:3542
IA32_X2APIC_LVT_PMI =                                          0x00000834 //col:3543
IA32_X2APIC_LVT_LINT0 =                                        0x00000835 //col:3544
IA32_X2APIC_LVT_LINT1 =                                        0x00000836 //col:3545
IA32_X2APIC_LVT_ERROR =                                        0x00000837 //col:3546
IA32_X2APIC_INIT_COUNT =                                       0x00000838 //col:3547
IA32_X2APIC_CUR_COUNT =                                        0x00000839 //col:3548
IA32_X2APIC_DIV_CONF =                                         0x0000083E //col:3549
IA32_X2APIC_SELF_IPI =                                         0x0000083F //col:3550
IA32_DEBUG_INTERFACE =                                         0x00000C80 //col:3551
IA32_L3_QOS_CFG =                                              0x00000C81 //col:3563
IA32_L2_QOS_CFG =                                              0x00000C82 //col:3572
IA32_QM_EVTSEL =                                               0x00000C8D //col:3581
IA32_QM_CTR =                                                  0x00000C8E //col:3592
IA32_PQR_ASSOC =                                               0x00000C8F //col:3603
IA32_BNDCFGS =                                                 0x00000D90 //col:3613
IA32_XSS =                                                     0x00000DA0 //col:3625
IA32_PKG_HDC_CTL =                                             0x00000DB0 //col:3635
IA32_PM_CTL1 =                                                 0x00000DB1 //col:3644
IA32_THREAD_STALL =                                            0x00000DB2 //col:3653
IA32_EFER =                                                    0xC0000080 //col:3658
IA32_STAR =                                                    0xC0000081 //col:3672
IA32_LSTAR =                                                   0xC0000082 //col:3673
IA32_CSTAR =                                                   0xC0000083 //col:3674
IA32_FMASK =                                                   0xC0000084 //col:3675
IA32_FS_BASE =                                                 0xC0000100 //col:3676
IA32_GS_BASE =                                                 0xC0000101 //col:3677
IA32_KERNEL_GS_BASE =                                          0xC0000102 //col:3678
IA32_TSC_AUX =                                                 0xC0000103 //col:3679
PDE_ENTRY_COUNT_32 =                                           0x00000400 //col:3781
PTE_ENTRY_COUNT_32 =                                           0x00000400 //col:3782
PML4_ENTRY_COUNT_64 =                                          0x00000200 //col:3960
PDPTE_ENTRY_COUNT_64 =                                         0x00000200 //col:3961
PDE_ENTRY_COUNT_64 =                                           0x00000200 //col:3962
PTE_ENTRY_COUNT_64 =                                           0x00000200 //col:3963
SEGMENT_DESCRIPTOR_TYPE_SYSTEM =                               0x00000000 //col:4098
SEGMENT_DESCRIPTOR_TYPE_CODE_OR_DATA =                         0x00000001 //col:4099
SEGMENT_DESCRIPTOR_TYPE_DATA_R =                               0x00000000 //col:4105
SEGMENT_DESCRIPTOR_TYPE_DATA_RA =                              0x00000001 //col:4106
SEGMENT_DESCRIPTOR_TYPE_DATA_RW =                              0x00000002 //col:4107
SEGMENT_DESCRIPTOR_TYPE_DATA_RWA =                             0x00000003 //col:4108
SEGMENT_DESCRIPTOR_TYPE_DATA_RE =                              0x00000004 //col:4109
SEGMENT_DESCRIPTOR_TYPE_DATA_REA =                             0x00000005 //col:4110
SEGMENT_DESCRIPTOR_TYPE_DATA_RWE =                             0x00000006 //col:4111
SEGMENT_DESCRIPTOR_TYPE_DATA_RWEA =                            0x00000007 //col:4112
SEGMENT_DESCRIPTOR_TYPE_CODE_E =                               0x00000008 //col:4113
SEGMENT_DESCRIPTOR_TYPE_CODE_EA =                              0x00000009 //col:4114
SEGMENT_DESCRIPTOR_TYPE_CODE_ER =                              0x0000000A //col:4115
SEGMENT_DESCRIPTOR_TYPE_CODE_ERA =                             0x0000000B //col:4116
SEGMENT_DESCRIPTOR_TYPE_CODE_EC =                              0x0000000C //col:4117
SEGMENT_DESCRIPTOR_TYPE_CODE_ECA =                             0x0000000D //col:4118
SEGMENT_DESCRIPTOR_TYPE_CODE_ERC =                             0x0000000E //col:4119
SEGMENT_DESCRIPTOR_TYPE_CODE_ERCA =                            0x0000000F //col:4120
SEGMENT_DESCRIPTOR_TYPE_RESERVED_1 =                           0x00000000 //col:4130
SEGMENT_DESCRIPTOR_TYPE_TSS_16_AVAILABLE =                     0x00000001 //col:4131
SEGMENT_DESCRIPTOR_TYPE_LDT =                                  0x00000002 //col:4132
SEGMENT_DESCRIPTOR_TYPE_TSS_16_BUSY =                          0x00000003 //col:4133
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE_16 =                         0x00000004 //col:4134
SEGMENT_DESCRIPTOR_TYPE_TASK_GATE =                            0x00000005 //col:4135
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE_16 =                    0x00000006 //col:4136
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE_16 =                         0x00000007 //col:4137
SEGMENT_DESCRIPTOR_TYPE_RESERVED_2 =                           0x00000008 //col:4138
SEGMENT_DESCRIPTOR_TYPE_TSS_AVAILABLE =                        0x00000009 //col:4139
SEGMENT_DESCRIPTOR_TYPE_RESERVED_3 =                           0x0000000A //col:4140
SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY =                             0x0000000B //col:4141
SEGMENT_DESCRIPTOR_TYPE_CALL_GATE =                            0x0000000C //col:4142
SEGMENT_DESCRIPTOR_TYPE_RESERVED_4 =                           0x0000000D //col:4143
SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE =                       0x0000000E //col:4144
SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE =                            0x0000000F //col:4145
VMX_EXIT_REASON_XCPT_OR_NMI =                                  0x00000000 //col:4197
VMX_EXIT_REASON_EXT_INT =                                      0x00000001 //col:4198
VMX_EXIT_REASON_TRIPLE_FAULT =                                 0x00000002 //col:4199
VMX_EXIT_REASON_INIT_SIGNAL =                                  0x00000003 //col:4200
VMX_EXIT_REASON_SIPI =                                         0x00000004 //col:4201
VMX_EXIT_REASON_IO_SMI =                                       0x00000005 //col:4202
VMX_EXIT_REASON_SMI =                                          0x00000006 //col:4203
VMX_EXIT_REASON_INT_WINDOW =                                   0x00000007 //col:4204
VMX_EXIT_REASON_NMI_WINDOW =                                   0x00000008 //col:4205
VMX_EXIT_REASON_TASK_SWITCH =                                  0x00000009 //col:4206
VMX_EXIT_REASON_CPUID =                                        0x0000000A //col:4207
VMX_EXIT_REASON_GETSEC =                                       0x0000000B //col:4208
VMX_EXIT_REASON_HLT =                                          0x0000000C //col:4209
VMX_EXIT_REASON_INVD =                                         0x0000000D //col:4210
VMX_EXIT_REASON_INVLPG =                                       0x0000000E //col:4211
VMX_EXIT_REASON_RDPMC =                                        0x0000000F //col:4212
VMX_EXIT_REASON_RDTSC =                                        0x00000010 //col:4213
VMX_EXIT_REASON_RSM =                                          0x00000011 //col:4214
VMX_EXIT_REASON_VMCALL =                                       0x00000012 //col:4215
VMX_EXIT_REASON_VMCLEAR =                                      0x00000013 //col:4216
VMX_EXIT_REASON_VMLAUNCH =                                     0x00000014 //col:4217
VMX_EXIT_REASON_VMPTRLD =                                      0x00000015 //col:4218
VMX_EXIT_REASON_VMPTRST =                                      0x00000016 //col:4219
VMX_EXIT_REASON_VMREAD =                                       0x00000017 //col:4220
VMX_EXIT_REASON_VMRESUME =                                     0x00000018 //col:4221
VMX_EXIT_REASON_VMWRITE =                                      0x00000019 //col:4222
VMX_EXIT_REASON_VMXOFF =                                       0x0000001A //col:4223
VMX_EXIT_REASON_VMXON =                                        0x0000001B //col:4224
VMX_EXIT_REASON_MOV_CRX =                                      0x0000001C //col:4225
VMX_EXIT_REASON_MOV_DRX =                                      0x0000001D //col:4226
VMX_EXIT_REASON_IO_INSTR =                                     0x0000001E //col:4227
VMX_EXIT_REASON_RDMSR =                                        0x0000001F //col:4228
VMX_EXIT_REASON_WRMSR =                                        0x00000020 //col:4229
VMX_EXIT_REASON_ERR_INVALID_GUEST_STATE =                      0x00000021 //col:4230
VMX_EXIT_REASON_ERR_MSR_LOAD =                                 0x00000022 //col:4231
VMX_EXIT_REASON_MWAIT =                                        0x00000024 //col:4232
VMX_EXIT_REASON_MTF =                                          0x00000025 //col:4233
VMX_EXIT_REASON_MONITOR =                                      0x00000027 //col:4234
VMX_EXIT_REASON_PAUSE =                                        0x00000028 //col:4235
VMX_EXIT_REASON_ERR_MACHINE_CHECK =                            0x00000029 //col:4236
VMX_EXIT_REASON_TPR_BELOW_THRESHOLD =                          0x0000002B //col:4237
VMX_EXIT_REASON_APIC_ACCESS =                                  0x0000002C //col:4238
VMX_EXIT_REASON_VIRTUALIZED_EOI =                              0x0000002D //col:4239
VMX_EXIT_REASON_XDTR_ACCESS =                                  0x0000002E //col:4240
VMX_EXIT_REASON_TR_ACCESS =                                    0x0000002F //col:4241
VMX_EXIT_REASON_EPT_VIOLATION =                                0x00000030 //col:4242
VMX_EXIT_REASON_EPT_MISCONFIG =                                0x00000031 //col:4243
VMX_EXIT_REASON_INVEPT =                                       0x00000032 //col:4244
VMX_EXIT_REASON_RDTSCP =                                       0x00000033 //col:4245
VMX_EXIT_REASON_PREEMPT_TIMER =                                0x00000034 //col:4246
VMX_EXIT_REASON_INVVPID =                                      0x00000035 //col:4247
VMX_EXIT_REASON_WBINVD =                                       0x00000036 //col:4248
VMX_EXIT_REASON_XSETBV =                                       0x00000037 //col:4249
VMX_EXIT_REASON_APIC_WRITE =                                   0x00000038 //col:4250
VMX_EXIT_REASON_RDRAND =                                       0x00000039 //col:4251
VMX_EXIT_REASON_INVPCID =                                      0x0000003A //col:4252
VMX_EXIT_REASON_VMFUNC =                                       0x0000003B //col:4253
VMX_EXIT_REASON_ENCLS =                                        0x0000003C //col:4254
VMX_EXIT_REASON_RDSEED =                                       0x0000003D //col:4255
VMX_EXIT_REASON_PML_FULL =                                     0x0000003E //col:4256
VMX_EXIT_REASON_XSAVES =                                       0x0000003F //col:4257
VMX_EXIT_REASON_XRSTORS =                                      0x00000040 //col:4258
VMX_ERROR_VMCALL =                                             0x00000001 //col:4268
VMX_ERROR_VMCLEAR_INVALID_PHYS_ADDR =                          0x00000002 //col:4269
VMX_ERROR_VMCLEAR_INVALID_VMXON_PTR =                          0x00000003 //col:4270
VMX_ERROR_VMLAUCH_NON_CLEAR_VMCS =                             0x00000004 //col:4271
VMX_ERROR_VMRESUME_NON_LAUNCHED_VMCS =                         0x00000005 //col:4272
VMX_ERROR_VMRESUME_CORRUPTED_VMCS =                            0x00000006 //col:4273
VMX_ERROR_VMENTRY_INVALID_CONTROL_FIELDS =                     0x00000007 //col:4274
VMX_ERROR_VMENTRY_INVALID_HOST_STATE =                         0x00000008 //col:4275
VMX_ERROR_VMPTRLD_INVALID_PHYS_ADDR =                          0x00000009 //col:4276
VMX_ERROR_VMPTRLD_VMXON_PTR =                                  0x0000000A //col:4277
VMX_ERROR_VMPTRLD_WRONG_VMCS_REVISION =                        0x0000000B //col:4278
VMX_ERROR_VMREAD_VMWRITE_INVALID_COMPONENT =                   0x0000000C //col:4279
VMX_ERROR_VMWRITE_READONLY_COMPONENT =                         0x0000000D //col:4280
VMX_ERROR_VMXON_IN_VMX_ROOT_OP =                               0x0000000F //col:4281
VMX_ERROR_VMENTRY_INVALID_VMCS_EXEC_PTR =                      0x00000010 //col:4282
VMX_ERROR_VMENTRY_NON_LAUNCHED_EXEC_VMCS =                     0x00000011 //col:4283
VMX_ERROR_VMENTRY_EXEC_VMCS_PTR =                              0x00000012 //col:4284
VMX_ERROR_VMCALL_NON_CLEAR_VMCS =                              0x00000013 //col:4285
VMX_ERROR_VMCALL_INVALID_VMEXIT_FIELDS =                       0x00000014 //col:4286
VMX_ERROR_VMCALL_INVALID_MSEG_REVISION =                       0x00000016 //col:4287
VMX_ERROR_VMXOFF_DUAL_MONITOR =                                0x00000017 //col:4288
VMX_ERROR_VMCALL_INVALID_SMM_MONITOR =                         0x00000018 //col:4289
VMX_ERROR_VMENTRY_INVALID_VM_EXEC_CTRL =                       0x00000019 //col:4290
VMX_ERROR_VMENTRY_MOV_SS =                                     0x0000001A //col:4291
VMX_ERROR_INVEPTVPID_INVALID_OPERAND =                         0x0000001C //col:4292
VMX_EXIT_QUALIFICATION_TYPE_CALL =                             0x00000000 //col:4336
VMX_EXIT_QUALIFICATION_TYPE_IRET =                             0x00000001 //col:4337
VMX_EXIT_QUALIFICATION_TYPE_JMP =                              0x00000002 //col:4338
VMX_EXIT_QUALIFICATION_TYPE_IDT =                              0x00000003 //col:4339
VMX_EXIT_QUALIFICATION_REGISTER_CR0 =                          0x00000000 //col:4348
VMX_EXIT_QUALIFICATION_REGISTER_CR2 =                          0x00000002 //col:4349
VMX_EXIT_QUALIFICATION_REGISTER_CR3 =                          0x00000003 //col:4350
VMX_EXIT_QUALIFICATION_REGISTER_CR4 =                          0x00000004 //col:4351
VMX_EXIT_QUALIFICATION_REGISTER_CR8 =                          0x00000008 //col:4352
VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR =                      0x00000000 //col:4354
VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR =                    0x00000001 //col:4355
VMX_EXIT_QUALIFICATION_ACCESS_CLTS =                           0x00000002 //col:4356
VMX_EXIT_QUALIFICATION_ACCESS_LMSW =                           0x00000003 //col:4357
VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER =                      0x00000000 //col:4359
VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY =                        0x00000001 //col:4360
VMX_EXIT_QUALIFICATION_GENREG_RAX =                            0x00000000 //col:4363
VMX_EXIT_QUALIFICATION_GENREG_RCX =                            0x00000001 //col:4364
VMX_EXIT_QUALIFICATION_GENREG_RDX =                            0x00000002 //col:4365
VMX_EXIT_QUALIFICATION_GENREG_RBX =                            0x00000003 //col:4366
VMX_EXIT_QUALIFICATION_GENREG_RSP =                            0x00000004 //col:4367
VMX_EXIT_QUALIFICATION_GENREG_RBP =                            0x00000005 //col:4368
VMX_EXIT_QUALIFICATION_GENREG_RSI =                            0x00000006 //col:4369
VMX_EXIT_QUALIFICATION_GENREG_RDI =                            0x00000007 //col:4370
VMX_EXIT_QUALIFICATION_GENREG_R8 =                             0x00000008 //col:4371
VMX_EXIT_QUALIFICATION_GENREG_R9 =                             0x00000009 //col:4372
VMX_EXIT_QUALIFICATION_GENREG_R10 =                            0x0000000A //col:4373
VMX_EXIT_QUALIFICATION_GENREG_R11 =                            0x0000000B //col:4374
VMX_EXIT_QUALIFICATION_GENREG_R12 =                            0x0000000C //col:4375
VMX_EXIT_QUALIFICATION_GENREG_R13 =                            0x0000000D //col:4376
VMX_EXIT_QUALIFICATION_GENREG_R14 =                            0x0000000E //col:4377
VMX_EXIT_QUALIFICATION_GENREG_R15 =                            0x0000000F //col:4378
VMX_EXIT_QUALIFICATION_REGISTER_DR0 =                          0x00000000 //col:4389
VMX_EXIT_QUALIFICATION_REGISTER_DR1 =                          0x00000001 //col:4390
VMX_EXIT_QUALIFICATION_REGISTER_DR2 =                          0x00000002 //col:4391
VMX_EXIT_QUALIFICATION_REGISTER_DR3 =                          0x00000003 //col:4392
VMX_EXIT_QUALIFICATION_REGISTER_DR6 =                          0x00000006 //col:4393
VMX_EXIT_QUALIFICATION_REGISTER_DR7 =                          0x00000007 //col:4394
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR =                   0x00000000 //col:4397
VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR =                 0x00000001 //col:4398
VMX_EXIT_QUALIFICATION_WIDTH_1B =                              0x00000000 //col:4409
VMX_EXIT_QUALIFICATION_WIDTH_2B =                              0x00000001 //col:4410
VMX_EXIT_QUALIFICATION_WIDTH_4B =                              0x00000003 //col:4411
VMX_EXIT_QUALIFICATION_DIRECTION_OUT =                         0x00000000 //col:4413
VMX_EXIT_QUALIFICATION_DIRECTION_IN =                          0x00000001 //col:4414
VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING =                  0x00000000 //col:4416
VMX_EXIT_QUALIFICATION_IS_STRING_STRING =                      0x00000001 //col:4417
VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP =                        0x00000000 //col:4419
VMX_EXIT_QUALIFICATION_IS_REP_REP =                            0x00000001 //col:4420
VMX_EXIT_QUALIFICATION_ENCODING_DX =                           0x00000000 //col:4422
VMX_EXIT_QUALIFICATION_ENCODING_IMM =                          0x00000001 //col:4423
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_READ =                      0x00000000 //col:4435
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_WRITE =                     0x00000001 //col:4436
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_INSTR_FETCH =               0x00000002 //col:4437
VMX_EXIT_QUALIFICATION_TYPE_LINEAR_EVENT_DELIVERY =            0x00000003 //col:4438
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_EVENT_DELIVERY =          0x0000000A //col:4439
VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_INSTR =                   0x0000000F //col:4440
IO_BITMAP_A_MIN =                                              0x00000000 //col:4668
IO_BITMAP_A_MAX =                                              0x00007FFF //col:4669
IO_BITMAP_B_MIN =                                              0x00008000 //col:4670
IO_BITMAP_B_MAX =                                              0x0000FFFF //col:4671
MSR_ID_LOW_MIN =                                               0x00000000 //col:4677
MSR_ID_LOW_MAX =                                               0x00001FFF //col:4678
MSR_ID_HIGH_MIN =                                              0xC0000000 //col:4679
MSR_ID_HIGH_MAX =                                              0xC0001FFF //col:4680
EPT_PAGE_WALK_LENGTH_4 =                                       0x00000003 //col:4696
EPT_LEVEL_PML4E =                                              0x00000003 //col:4855
EPT_LEVEL_PDPTE =                                              0x00000002 //col:4856
EPT_LEVEL_PDE =                                                0x00000001 //col:4857
EPT_LEVEL_PTE =                                                0x00000000 //col:4858
EPML4_ENTRY_COUNT =                                            0x00000200 //col:4868
EPDPTE_ENTRY_COUNT =                                           0x00000200 //col:4869
EPDE_ENTRY_COUNT =                                             0x00000200 //col:4870
EPTE_ENTRY_COUNT =                                             0x00000200 //col:4871
VMCS_CTRL_VPID =                                               0x00000000 //col:4962
VMCS_CTRL_POSTED_INTR_NOTIFY_VECTOR =                          0x00000002 //col:4963
VMCS_CTRL_EPTP_INDEX =                                         0x00000004 //col:4964
VMCS_CTRL_HLAT_PREFIX_SIZE =                                   0x00000006 //col:4965
VMCS_GUEST_ES_SEL =                                            0x00000800 //col:4975
VMCS_GUEST_CS_SEL =                                            0x00000802 //col:4976
VMCS_GUEST_SS_SEL =                                            0x00000804 //col:4977
VMCS_GUEST_DS_SEL =                                            0x00000806 //col:4978
VMCS_GUEST_FS_SEL =                                            0x00000808 //col:4979
VMCS_GUEST_GS_SEL =                                            0x0000080A //col:4980
VMCS_GUEST_LDTR_SEL =                                          0x0000080C //col:4981
VMCS_GUEST_TR_SEL =                                            0x0000080E //col:4982
VMCS_GUEST_INTR_STATUS =                                       0x00000810 //col:4983
VMCS_GUEST_PML_INDEX =                                         0x00000812 //col:4984
VMCS_HOST_ES_SEL =                                             0x00000C00 //col:4994
VMCS_HOST_CS_SEL =                                             0x00000C02 //col:4995
VMCS_HOST_SS_SEL =                                             0x00000C04 //col:4996
VMCS_HOST_DS_SEL =                                             0x00000C06 //col:4997
VMCS_HOST_FS_SEL =                                             0x00000C08 //col:4998
VMCS_HOST_GS_SEL =                                             0x00000C0A //col:4999
VMCS_HOST_TR_SEL =                                             0x00000C0C //col:5000
VMCS_CTRL_IO_BITMAP_A =                                        0x00002000 //col:5019
VMCS_CTRL_IO_BITMAP_B =                                        0x00002002 //col:5020
VMCS_CTRL_MSR_BITMAP =                                         0x00002004 //col:5021
VMCS_CTRL_VMEXIT_MSR_STORE =                                   0x00002006 //col:5022
VMCS_CTRL_VMEXIT_MSR_LOAD =                                    0x00002008 //col:5023
VMCS_CTRL_VMENTRY_MSR_LOAD =                                   0x0000200A //col:5024
VMCS_CTRL_EXEC_VMCS_PTR =                                      0x0000200C //col:5025
VMCS_CTRL_PML_ADDR =                                           0x0000200E //col:5026
VMCS_CTRL_TSC_OFFSET =                                         0x00002010 //col:5027
VMCS_CTRL_VAPIC_PAGEADDR =                                     0x00002012 //col:5028
VMCS_CTRL_APIC_ACCESSADDR =                                    0x00002014 //col:5029
VMCS_CTRL_POSTED_INTR_DESC =                                   0x00002016 //col:5030
VMCS_CTRL_VMFUNC_CTRLS =                                       0x00002018 //col:5031
VMCS_CTRL_EPTP =                                               0x0000201A //col:5032
VMCS_CTRL_EOI_BITMAP_0 =                                       0x0000201C //col:5033
VMCS_CTRL_EOI_BITMAP_1 =                                       0x0000201E //col:5034
VMCS_CTRL_EOI_BITMAP_2 =                                       0x00002020 //col:5035
VMCS_CTRL_EOI_BITMAP_3 =                                       0x00002022 //col:5036
VMCS_CTRL_EPTP_LIST =                                          0x00002024 //col:5037
VMCS_CTRL_VMREAD_BITMAP =                                      0x00002026 //col:5038
VMCS_CTRL_VMWRITE_BITMAP =                                     0x00002028 //col:5039
VMCS_CTRL_VIRTXCPT_INFO_ADDR =                                 0x0000202A //col:5040
VMCS_CTRL_XSS_EXITING_BITMAP =                                 0x0000202C //col:5041
VMCS_CTRL_ENCLS_EXITING_BITMAP =                               0x0000202E //col:5042
VMCS_CTRL_SPP_TABLE_POINTER =                                  0x00002030 //col:5043
VMCS_CTRL_TSC_MULTIPLIER =                                     0x00002032 //col:5044
VMCS_CTRL_PROC_EXEC3 =                                         0x00002034 //col:5045
VMCS_CTRL_ENCLV_EXITING_BITMAP =                               0x00002036 //col:5046
VMCS_CTRL_HLATP =                                              0x00002040 //col:5047
VMCS_CTRL_SECONDARY_EXIT =                                     0x00002044 //col:5048
VMCS_GUEST_PHYS_ADDR =                                         0x00002400 //col:5058
VMCS_GUEST_VMCS_LINK_PTR =                                     0x00002800 //col:5068
VMCS_GUEST_DEBUGCTL =                                          0x00002802 //col:5069
VMCS_GUEST_PAT =                                               0x00002804 //col:5070
VMCS_GUEST_EFER =                                              0x00002806 //col:5071
VMCS_GUEST_PERF_GLOBAL_CTRL =                                  0x00002808 //col:5072
VMCS_GUEST_PDPTE0 =                                            0x0000280A //col:5073
VMCS_GUEST_PDPTE1 =                                            0x0000280C //col:5074
VMCS_GUEST_PDPTE2 =                                            0x0000280E //col:5075
VMCS_GUEST_PDPTE3 =                                            0x00002810 //col:5076
VMCS_GUEST_BNDCFGS =                                           0x00002812 //col:5077
VMCS_GUEST_RTIT_CTL =                                          0x00002814 //col:5078
VMCS_GUEST_LBR_CTL =                                           0x00002816 //col:5079
VMCS_GUEST_PKRS =                                              0x00002818 //col:5080
VMCS_HOST_PAT =                                                0x00002C00 //col:5090
VMCS_HOST_EFER =                                               0x00002C02 //col:5091
VMCS_HOST_PERF_GLOBAL_CTRL =                                   0x00002C04 //col:5092
VMCS_HOST_PKRS =                                               0x00002C06 //col:5093
VMCS_CTRL_PIN_EXEC =                                           0x00004000 //col:5112
VMCS_CTRL_PROC_EXEC =                                          0x00004002 //col:5113
VMCS_CTRL_EXCEPTION_BITMAP =                                   0x00004004 //col:5114
VMCS_CTRL_PAGEFAULT_ERROR_MASK =                               0x00004006 //col:5115
VMCS_CTRL_PAGEFAULT_ERROR_MATCH =                              0x00004008 //col:5116
VMCS_CTRL_CR3_TARGET_COUNT =                                   0x0000400A //col:5117
VMCS_CTRL_PRIMARY_EXIT =                                       0x0000400C //col:5118
VMCS_CTRL_EXIT_MSR_STORE_COUNT =                               0x0000400E //col:5119
VMCS_CTRL_EXIT_MSR_LOAD_COUNT =                                0x00004010 //col:5120
VMCS_CTRL_ENTRY =                                              0x00004012 //col:5121
VMCS_CTRL_ENTRY_MSR_LOAD_COUNT =                               0x00004014 //col:5122
VMCS_CTRL_ENTRY_INTERRUPTION_INFO =                            0x00004016 //col:5123
VMCS_CTRL_ENTRY_EXCEPTION_ERRCODE =                            0x00004018 //col:5124
VMCS_CTRL_ENTRY_INSTR_LENGTH =                                 0x0000401A //col:5125
VMCS_CTRL_TPR_THRESHOLD =                                      0x0000401C //col:5126
VMCS_CTRL_PROC_EXEC2 =                                         0x0000401E //col:5127
VMCS_CTRL_PLE_GAP =                                            0x00004020 //col:5128
VMCS_CTRL_PLE_WINDOW =                                         0x00004022 //col:5129
VMCS_VM_INSTR_ERROR =                                          0x00004400 //col:5139
VMCS_EXIT_REASON =                                             0x00004402 //col:5140
VMCS_EXIT_INTERRUPTION_INFO =                                  0x00004404 //col:5141
VMCS_EXIT_INTERRUPTION_ERROR_CODE =                            0x00004406 //col:5142
VMCS_IDT_VECTORING_INFO =                                      0x00004408 //col:5143
VMCS_IDT_VECTORING_ERROR_CODE =                                0x0000440A //col:5144
VMCS_EXIT_INSTR_LENGTH =                                       0x0000440C //col:5145
VMCS_EXIT_INSTR_INFO =                                         0x0000440E //col:5146
VMCS_GUEST_ES_LIMIT =                                          0x00004800 //col:5156
VMCS_GUEST_CS_LIMIT =                                          0x00004802 //col:5157
VMCS_GUEST_SS_LIMIT =                                          0x00004804 //col:5158
VMCS_GUEST_DS_LIMIT =                                          0x00004806 //col:5159
VMCS_GUEST_FS_LIMIT =                                          0x00004808 //col:5160
VMCS_GUEST_GS_LIMIT =                                          0x0000480A //col:5161
VMCS_GUEST_LDTR_LIMIT =                                        0x0000480C //col:5162
VMCS_GUEST_TR_LIMIT =                                          0x0000480E //col:5163
VMCS_GUEST_GDTR_LIMIT =                                        0x00004810 //col:5164
VMCS_GUEST_IDTR_LIMIT =                                        0x00004812 //col:5165
VMCS_GUEST_ES_ACCESS_RIGHTS =                                  0x00004814 //col:5166
VMCS_GUEST_CS_ACCESS_RIGHTS =                                  0x00004816 //col:5167
VMCS_GUEST_SS_ACCESS_RIGHTS =                                  0x00004818 //col:5168
VMCS_GUEST_DS_ACCESS_RIGHTS =                                  0x0000481A //col:5169
VMCS_GUEST_FS_ACCESS_RIGHTS =                                  0x0000481C //col:5170
VMCS_GUEST_GS_ACCESS_RIGHTS =                                  0x0000481E //col:5171
VMCS_GUEST_LDTR_ACCESS_RIGHTS =                                0x00004820 //col:5172
VMCS_GUEST_TR_ACCESS_RIGHTS =                                  0x00004822 //col:5173
VMCS_GUEST_INTERRUPTIBILITY_STATE =                            0x00004824 //col:5174
VMCS_GUEST_ACTIVITY_STATE =                                    0x00004826 //col:5175
VMCS_GUEST_SMBASE =                                            0x00004828 //col:5176
VMCS_GUEST_SYSENTER_CS =                                       0x0000482A //col:5177
VMCS_GUEST_PREEMPT_TIMER_VALUE =                               0x0000482E //col:5178
VMCS_HOST_SYSENTER_CS =                                        0x00004C00 //col:5188
VMCS_CTRL_CR0_MASK =                                           0x00006000 //col:5207
VMCS_CTRL_CR4_MASK =                                           0x00006002 //col:5208
VMCS_CTRL_CR0_READ_SHADOW =                                    0x00006004 //col:5209
VMCS_CTRL_CR4_READ_SHADOW =                                    0x00006006 //col:5210
VMCS_CTRL_CR3_TARGET_VAL0 =                                    0x00006008 //col:5211
VMCS_CTRL_CR3_TARGET_VAL1 =                                    0x0000600A //col:5212
VMCS_CTRL_CR3_TARGET_VAL2 =                                    0x0000600C //col:5213
VMCS_CTRL_CR3_TARGET_VAL3 =                                    0x0000600E //col:5214
VMCS_EXIT_QUALIFICATION =                                      0x00006400 //col:5224
VMCS_IO_RCX =                                                  0x00006402 //col:5225
VMCS_IO_RSI =                                                  0x00006404 //col:5226
VMCS_IO_RDI =                                                  0x00006406 //col:5227
VMCS_IO_RIP =                                                  0x00006408 //col:5228
VMCS_EXIT_GUEST_LINEAR_ADDR =                                  0x0000640A //col:5229
VMCS_GUEST_CR0 =                                               0x00006800 //col:5239
VMCS_GUEST_CR3 =                                               0x00006802 //col:5240
VMCS_GUEST_CR4 =                                               0x00006804 //col:5241
VMCS_GUEST_ES_BASE =                                           0x00006806 //col:5242
VMCS_GUEST_CS_BASE =                                           0x00006808 //col:5243
VMCS_GUEST_SS_BASE =                                           0x0000680A //col:5244
VMCS_GUEST_DS_BASE =                                           0x0000680C //col:5245
VMCS_GUEST_FS_BASE =                                           0x0000680E //col:5246
VMCS_GUEST_GS_BASE =                                           0x00006810 //col:5247
VMCS_GUEST_LDTR_BASE =                                         0x00006812 //col:5248
VMCS_GUEST_TR_BASE =                                           0x00006814 //col:5249
VMCS_GUEST_GDTR_BASE =                                         0x00006816 //col:5250
VMCS_GUEST_IDTR_BASE =                                         0x00006818 //col:5251
VMCS_GUEST_DR7 =                                               0x0000681A //col:5252
VMCS_GUEST_RSP =                                               0x0000681C //col:5253
VMCS_GUEST_RIP =                                               0x0000681E //col:5254
VMCS_GUEST_RFLAGS =                                            0x00006820 //col:5255
VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS =                          0x00006822 //col:5256
VMCS_GUEST_SYSENTER_ESP =                                      0x00006824 //col:5257
VMCS_GUEST_SYSENTER_EIP =                                      0x00006826 //col:5258
VMCS_GUEST_S_CET =                                             0x00006C28 //col:5259
VMCS_GUEST_SSP =                                               0x00006C2A //col:5260
VMCS_GUEST_INTERRUPT_SSP_TABLE_ADDR =                          0x00006C2C //col:5261
VMCS_HOST_CR0 =                                                0x00006C00 //col:5271
VMCS_HOST_CR3 =                                                0x00006C02 //col:5272
VMCS_HOST_CR4 =                                                0x00006C04 //col:5273
VMCS_HOST_FS_BASE =                                            0x00006C06 //col:5274
VMCS_HOST_GS_BASE =                                            0x00006C08 //col:5275
VMCS_HOST_TR_BASE =                                            0x00006C0A //col:5276
VMCS_HOST_GDTR_BASE =                                          0x00006C0C //col:5277
VMCS_HOST_IDTR_BASE =                                          0x00006C0E //col:5278
VMCS_HOST_SYSENTER_ESP =                                       0x00006C10 //col:5279
VMCS_HOST_SYSENTER_EIP =                                       0x00006C12 //col:5280
VMCS_HOST_RSP =                                                0x00006C14 //col:5281
VMCS_HOST_RIP =                                                0x00006C16 //col:5282
VMCS_HOST_S_CET =                                              0x00006C18 //col:5283
VMCS_HOST_SSP =                                                0x00006C1A //col:5284
VMCS_HOST_INTERRUPT_SSP_TABLE_ADDR =                           0x00006C1C //col:5285
APIC_BASE =                                                    0xFEE00000 //col:5342
APIC_ID =                                                      0x00000020 //col:5343
APIC_VERSION =                                                 0x00000030 //col:5344
APIC_TPR =                                                     0x00000080 //col:5345
APIC_APR =                                                     0x00000090 //col:5346
APIC_PPR =                                                     0x000000A0 //col:5347
APIC_EOI =                                                     0x000000B0 //col:5348
APIC_REMOTE_READ =                                             0x000000C0 //col:5349
APIC_LOGICAL_DESTINATION =                                     0x000000D0 //col:5350
APIC_DESTINATION_FORMAT =                                      0x000000E0 //col:5351
APIC_SIV =                                                     0x000000F0 //col:5352
APIC_ISR_31_0 =                                                0x00000100 //col:5353
APIC_ISR_63_32 =                                               0x00000110 //col:5354
APIC_ISR_95_64 =                                               0x00000120 //col:5355
APIC_ISR_127_96 =                                              0x00000130 //col:5356
APIC_ISR_159_128 =                                             0x00000140 //col:5357
APIC_ISR_191_160 =                                             0x00000150 //col:5358
APIC_ISR_223_192 =                                             0x00000160 //col:5359
APIC_ISR_255_224 =                                             0x00000170 //col:5360
APIC_TMR_31_0 =                                                0x00000180 //col:5361
APIC_TMR_63_32 =                                               0x00000190 //col:5362
APIC_TMR_95_64 =                                               0x000001A0 //col:5363
APIC_TMR_127_96 =                                              0x000001B0 //col:5364
APIC_TMR_159_128 =                                             0x000001C0 //col:5365
APIC_TMR_191_160 =                                             0x000001D0 //col:5366
APIC_TMR_223_192 =                                             0x000001E0 //col:5367
APIC_TMR_255_224 =                                             0x000001F0 //col:5368
APIC_IRR_31_0 =                                                0x00000200 //col:5369
APIC_IRR_63_32 =                                               0x00000210 //col:5370
APIC_IRR_95_64 =                                               0x00000220 //col:5371
APIC_IRR_127_96 =                                              0x00000230 //col:5372
APIC_IRR_159_128 =                                             0x00000240 //col:5373
APIC_IRR_191_160 =                                             0x00000250 //col:5374
APIC_IRR_223_192 =                                             0x00000260 //col:5375
APIC_IRR_255_224 =                                             0x00000270 //col:5376
APIC_ERROR_STATUS =                                            0x00000280 //col:5377
APIC_CMCI =                                                    0x000002F0 //col:5378
APIC_ICR_0_31 =                                                0x00000300 //col:5379
APIC_ICR_32_63 =                                               0x00000310 //col:5380
APIC_LVT_TIMER =                                               0x00000320 //col:5381
APIC_LVT_THERMAL_SENSOR =                                      0x00000330 //col:5382
APIC_LVT_PERFORMANCE_MONITORING_COUNTERS =                     0x00000340 //col:5383
APIC_LVT_LINT0 =                                               0x00000350 //col:5384
APIC_LVT_LINT1 =                                               0x00000360 //col:5385
APIC_LVT_ERROR =                                               0x00000370 //col:5386
APIC_INITIAL_COUNT =                                           0x00000380 //col:5387
APIC_CURRENT_COUNT =                                           0x00000390 //col:5388
APIC_DIVIDE_CONFIGURATION =                                    0x000003E0 //col:5389
MEMORY_TYPE_UC =                                               0x00000000 //col:5525
MEMORY_TYPE_WC =                                               0x00000001 //col:5526
MEMORY_TYPE_WT =                                               0x00000004 //col:5527
MEMORY_TYPE_WP =                                               0x00000005 //col:5528
MEMORY_TYPE_WB =                                               0x00000006 //col:5529
MEMORY_TYPE_UC_MINUS =                                         0x00000007 //col:5530
MEMORY_TYPE_INVALID =                                          0x000000FF //col:5531
VTD_ROOT_ENTRY_COUNT =                                         0x00000100 //col:5593
VTD_CONTEXT_ENTRY_COUNT =                                      0x00000100 //col:5594
VTD_VERSION =                                                  0x00000000 //col:5599
VTD_CAPABILITY =                                               0x00000008 //col:5609
VTD_EXTENDED_CAPABILITY =                                      0x00000010 //col:5643
VTD_GLOBAL_COMMAND =                                           0x00000018 //col:5686
VTD_GLOBAL_STATUS =                                            0x0000001C //col:5704
VTD_ROOT_TABLE_ADDRESS =                                       0x00000020 //col:5722
VTD_CONTEXT_COMMAND =                                          0x00000028 //col:5733
VTD_INVALIDATE_ADDRESS =                                       0x00000000 //col:5748
VTD_IOTLB_INVALIDATE =                                         0x00000008 //col:5760
)

const(
  invpcid_individual_address                                    =  0x00000000  //col:3977
  invpcid_single_context                                        =  0x00000001  //col:3978
  invpcid_all_context_with_globals                              =  0x00000002  //col:3979
  invpcid_all_context                                           =  0x00000003  //col:3980
)


const(
  vmx_active                                                    =  0x00000000  //col:4625
  vmx_hlt                                                       =  0x00000001  //col:4626
  vmx_shutdown                                                  =  0x00000002  //col:4627
  vmx_wait_for_sipi                                             =  0x00000003  //col:4628
)


const(
  invept_single_context                                         =  0x00000001  //col:4881
  invept_all_context                                            =  0x00000002  //col:4882
)


const(
  invvpid_individual_address                                    =  0x00000000  //col:4886
  invvpid_single_context                                        =  0x00000001  //col:4887
  invvpid_all_context                                           =  0x00000002  //col:4888
  invvpid_single_context_retaining_globals                      =  0x00000003  //col:4889
)


const(
  external_interrupt                                            =  0x00000000  //col:5299
  non_maskable_interrupt                                        =  0x00000002  //col:5300
  hardware_exception                                            =  0x00000003  //col:5301
  software_interrupt                                            =  0x00000004  //col:5302
  privileged_software_exception                                 =  0x00000005  //col:5303
  software_exception                                            =  0x00000006  //col:5304
  other_event                                                   =  0x00000007  //col:5305
)


const(
  divide_error                                                  =  0x00000000  //col:5465
  debug                                                         =  0x00000001  //col:5466
  nmi                                                           =  0x00000002  //col:5467
  breakpoint                                                    =  0x00000003  //col:5468
  overflow                                                      =  0x00000004  //col:5469
  bound_range_exceeded                                          =  0x00000005  //col:5470
  invalid_opcode                                                =  0x00000006  //col:5471
  device_not_available                                          =  0x00000007  //col:5472
  double_fault                                                  =  0x00000008  //col:5473
  coprocessor_segment_overrun                                   =  0x00000009  //col:5474
  invalid_tss                                                   =  0x0000000A  //col:5475
  segment_not_present                                           =  0x0000000B  //col:5476
  stack_segment_fault                                           =  0x0000000C  //col:5477
  general_protection                                            =  0x0000000D  //col:5478
  page_fault                                                    =  0x0000000E  //col:5479
  x87_floating_point_error                                      =  0x00000010  //col:5480
  alignment_check                                               =  0x00000011  //col:5481
  machine_check                                                 =  0x00000012  //col:5482
  simd_floating_point_error                                     =  0x00000013  //col:5483
  virtualization_exception                                      =  0x00000014  //col:5484
  control_protection                                            =  0x00000015  //col:5485
)



type (
Ia32Compact interface{
#if defined()(ok bool)//col:38
#pragma pack()(ok bool)//col:4002
#pragma pack()(ok bool)//col:4009
#pragma pack()(ok bool)//col:4027
#pragma pack()(ok bool)//col:4181
#pragma pack()(ok bool)//col:4309
}











































)

func NewIa32Compact() { return & ia32Compact{} }

func (i *ia32Compact)#if defined()(ok bool){//col:38























return true
}












































func (i *ia32Compact)#pragma pack()(ok bool){//col:4002





return true
}












































func (i *ia32Compact)#pragma pack()(ok bool){//col:4009






return true
}












































func (i *ia32Compact)#pragma pack()(ok bool){//col:4027
















return true
}












































func (i *ia32Compact)#pragma pack()(ok bool){//col:4181


















return true
}












































func (i *ia32Compact)#pragma pack()(ok bool){//col:4309
































































































return true
}














































