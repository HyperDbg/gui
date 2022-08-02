package out
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\ia32-doc\out\ia32_compact.h.back

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
}


type typedef struct { struct{
cache_type_field uint32_t //col:106
cache_level uint32_t //col:107
self_initializing_cache_level uint32_t //col:108
fully_associative_cache uint32_t //col:109
reserved_1 uint32_t //col:110
max_addressable_ids_for_logical_processors_sharing_this_cache: uint32_t //col:111
max_addressable_ids_for_processor_cores_in_physical_package: uint32_t //col:112
}


type typedef struct { struct{
smallest_monitor_line_size uint32_t //col:142
}


type typedef struct { struct{
temperature_sensor_supported uint32_t //col:176
intel_turbo_boost_technology_available uint32_t //col:177
apic_timer_always_running uint32_t //col:178
reserved_1 uint32_t //col:179
power_limit_notification uint32_t //col:180
clock_modulation_duty uint32_t //col:181
package_thermal_management uint32_t //col:182
hwp_base_registers uint32_t //col:183
hwp_notification uint32_t //col:184
hwp_activity_window uint32_t //col:185
hwp_energy_performance_preference uint32_t //col:186
hwp_package_level_request uint32_t //col:187
reserved_2 uint32_t //col:188
hdc uint32_t //col:189
intel_turbo_boost_max_technology_3_available uint32_t //col:190
hwp_capabilities uint32_t //col:191
hwp_peci_override uint32_t //col:192
flexible_hwp uint32_t //col:193
fast_access_mode_for_hwp_request_msr uint32_t //col:194
reserved_3 uint32_t //col:195
ignoring_idle_logical_processor_hwp_request uint32_t //col:196
reserved_4 uint32_t //col:197
intel_thread_director uint32_t //col:198
}


type typedef struct { struct{
number_of_sub_leaves uint32_t //col:228
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t //col:333
}


type typedef struct { struct{
version_id_of_architectural_performance_monitoring uint32_t //col:359
number_of_performance_monitoring_counter_per_logical_processor: uint32_t //col:360
bit_width_of_performance_monitoring_counter uint32_t //col:361
ebx_bit_vector_length uint32_t //col:362
}


type typedef struct { struct{
x2apic_id_to_unique_topology_id_shift uint32_t //col:397
}


type typedef struct { struct{
x87_state uint32_t //col:424
sse_state uint32_t //col:425
avx_state uint32_t //col:426
mpx_state uint32_t //col:427
avx_512_state uint32_t //col:428
used_for_ia32_xss_1 uint32_t //col:429
pkru_state uint32_t //col:430
reserved_1 uint32_t //col:431
used_for_ia32_xss_2 uint32_t //col:432
}


type typedef struct { struct{
reserved_1 uint32_t //col:458
supports_xsavec_and_compacted_xrstor uint32_t //col:459
supports_xgetbv_with_ecx_1 uint32_t //col:460
supports_xsave_xrstor_and_ia32_xss uint32_t //col:461
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t //col:496
}


type typedef struct { struct{
reserved uint32_t //col:523
}


type typedef struct { struct{
reserved uint32_t //col:550
}


type typedef struct { struct{
ia32_platform_dca_cap uint32_t //col:578
}


type typedef struct { struct{
length_of_capacity_bit_mask uint32_t //col:607
}


type typedef struct { struct{
length_of_capacity_bit_mask uint32_t //col:634
}


type typedef struct { struct{
max_mba_throttling_value uint32_t //col:660
}


type typedef struct { struct{
sgx1 uint32_t //col:687
sgx2 uint32_t //col:688
reserved_1 uint32_t //col:689
sgx_enclv_advanced uint32_t //col:690
sgx_encls_advanced uint32_t //col:691
}


type typedef struct { struct{
valid_secs_attributes_0 uint32_t //col:718
}


type typedef struct { struct{
sub_leaf_type uint32_t //col:744
}


type typedef struct { struct{
sub_leaf_type uint32_t //col:770
reserved_1 uint32_t //col:771
epc_base_physical_address_1 uint32_t //col:772
}


type typedef struct { struct{
max_sub_leaf uint32_t //col:800
}


type typedef struct { struct{
number_of_configurable_address_ranges_for_filtering uint32_t //col:839
reserved_1 uint32_t //col:840
bitmap_of_supported_mtc_period_encodings uint32_t //col:841
}


type typedef struct { struct{
denominator uint32_t //col:868
}


type typedef struct { struct{
procesor_base_frequency_mhz uint32_t //col:894
}


type typedef struct { struct{
max_soc_id_index uint32_t //col:920
}


type typedef struct { struct{
soc_vendor_brand_string uint32_t //col:947
}


type typedef struct { struct{
reserved uint32_t //col:973
}


type typedef struct { struct{
max_sub_leaf uint32_t //col:999
}


type typedef struct { struct{
reserved uint32_t //col:1036
}


type typedef struct { struct{
max_extended_functions uint32_t //col:1073
}


type typedef struct { struct{
reserved uint32_t //col:1099
}


type typedef struct { struct{
processor_brand_string_1 uint32_t //col:1137
}


type typedef struct { struct{
processor_brand_string_5 uint32_t //col:1163
}


type typedef struct { struct{
processor_brand_string_9 uint32_t //col:1189
}


type typedef struct { struct{
reserved uint32_t //col:1215
}


type typedef struct { struct{
reserved uint32_t //col:1241
}


type typedef struct { struct{
reserved uint32_t //col:1270
}


type typedef struct { struct{
number_of_physical_address_bits uint32_t //col:1297
number_of_linear_address_bits uint32_t //col:1298
}


type typedef struct { struct{
thread_adjust uint64_t //col:1322
}


type typedef struct { struct{
mseg_header_revision uint32_t //col:1325
monitor_features uint32_t //col:1326
gdtr_limit uint32_t //col:1327
gdtr_base_offset uint32_t //col:1328
cs_selector uint32_t //col:1329
eip_offset uint32_t //col:1330
esp_offset uint32_t //col:1331
cr3_offset uint32_t //col:1332
}


type typedef struct { struct{
c0_mcnt uint64_t //col:1335
}


type typedef struct { struct{
c0_acnt uint64_t //col:1338
}


type typedef struct { struct{
stall_cycle_cnt uint64_t //col:1341
}


type typedef struct { struct{
limit uint16_t //col:1344
base_address uint32_t //col:1345
}


type typedef struct { struct{
limit uint16_t //col:1348
base_address uint64_t //col:1349
}


type typedef struct { struct{
segment_limit_low uint16_t //col:1352
base_address_low uint16_t //col:1353
base_address_middle uint32_t //col:1356
type uint32_t //col:1357
descriptor_type uint32_t //col:1358
descriptor_privilege_level uint32_t //col:1359
present uint32_t //col:1360
segment_limit_high uint32_t //col:1361
available_bit uint32_t //col:1362
long_mode uint32_t //col:1363
default_big uint32_t //col:1364
granularity uint32_t //col:1365
base_address_high uint32_t //col:1366
}


type typedef struct { struct{
segment_limit_low uint16_t //col:1372
base_address_low uint16_t //col:1373
base_address_middle uint32_t //col:1376
type uint32_t //col:1377
descriptor_type uint32_t //col:1378
descriptor_privilege_level uint32_t //col:1379
present uint32_t //col:1380
segment_limit_high uint32_t //col:1381
available_bit uint32_t //col:1382
long_mode uint32_t //col:1383
default_big uint32_t //col:1384
granularity uint32_t //col:1385
base_address_high uint32_t //col:1386
}


type typedef struct { struct{
offset_low uint16_t //col:1394
segment_selector uint16_t //col:1395
interrupt_stack_table uint32_t //col:1398
must_be_zero_0 uint32_t //col:1399
type uint32_t //col:1400
must_be_zero_1 uint32_t //col:1401
descriptor_privilege_level uint32_t //col:1402
present uint32_t //col:1403
offset_middle uint32_t //col:1404
}


type typedef struct { struct{
reserved_0 uint32_t //col:1412
rsp0 uint64_t //col:1413
rsp1 uint64_t //col:1414
rsp2 uint64_t //col:1415
reserved_1 uint64_t //col:1416
ist1 uint64_t //col:1417
ist2 uint64_t //col:1418
ist3 uint64_t //col:1419
ist4 uint64_t //col:1420
ist5 uint64_t //col:1421
ist6 uint64_t //col:1422
ist7 uint64_t //col:1423
reserved_2 uint64_t //col:1424
reserved_3 uint16_t //col:1425
io_map_base uint16_t //col:1426
}


type typedef struct { struct{
reason uint32_t //col:1429
exception_mask uint32_t //col:1430
exit uint64_t //col:1431
guest_linear_address uint64_t //col:1432
guest_physical_address uint64_t //col:1433
current_eptp_index uint16_t //col:1434
}


type typedef struct { struct{
io_a[4096] uint8_t //col:1437
io_b[4096] uint8_t //col:1438
}


type typedef struct { struct{
rdmsr_low[1024] uint8_t //col:1441
rdmsr_high[1024] uint8_t //col:1442
wrmsr_low[1024] uint8_t //col:1443
wrmsr_high[1024] uint8_t //col:1444
}


type typedef struct { struct{
ept_pointer uint64_t //col:1447
reserved uint64_t //col:1448
}


type typedef struct { struct{
vpid uint16_t //col:1451
reserved1 uint16_t //col:1452
reserved2 uint32_t //col:1453
linear_address uint64_t //col:1454
}


type typedef struct { struct{
revision_id uint32_t //col:1458
shadow_vmcs_indicator uint32_t //col:1459
}


type typedef struct { struct{
revision_id uint32_t //col:1466
must_be_zero uint32_t //col:1467
}


type typedef struct { struct{
present uint64_t //col:1474
reserved_1 uint64_t //col:1475
context_table_pointer uint64_t //col:1476
}


type typedef struct { struct{
present uint64_t //col:1490
fault_processing_disable uint64_t //col:1491
translation_type uint64_t //col:1492
reserved_1 uint64_t //col:1493
second_level_page_translation_pointer uint64_t //col:1494
}




