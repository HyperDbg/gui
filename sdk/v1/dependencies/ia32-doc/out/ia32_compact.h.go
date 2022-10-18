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



type   struct{
max_cpuid_input_value uint32_t //col:7
ebx_value_genu uint32_t //col:8
ecx_value_ntel uint32_t //col:9
edx_value_inei uint32_t //col:10
}


type   struct{
stepping_id uint32_t //col:20
model uint32_t //col:21
family_id uint32_t //col:22
processor_type uint32_t //col:23
reserved_1 uint32_t //col:24
extended_model_id uint32_t //col:25
extended_family_id uint32_t //col:26
}


type   struct{
cache_type_field uint32_t //col:116
cache_level uint32_t //col:117
self_initializing_cache_level uint32_t //col:118
fully_associative_cache uint32_t //col:119
reserved_1 uint32_t //col:120
max_addressable_ids_for_logical_processors_sharing_this_cache: uint32_t //col:121
max_addressable_ids_for_processor_cores_in_physical_package: uint32_t //col:122
}


type   struct{
smallest_monitor_line_size uint32_t //col:146
}


type   struct{
temperature_sensor_supported uint32_t //col:202
intel_turbo_boost_technology_available uint32_t //col:203
apic_timer_always_running uint32_t //col:204
reserved_1 uint32_t //col:205
power_limit_notification uint32_t //col:206
clock_modulation_duty uint32_t //col:207
package_thermal_management uint32_t //col:208
hwp_base_registers uint32_t //col:209
hwp_notification uint32_t //col:210
hwp_activity_window uint32_t //col:211
hwp_energy_performance_preference uint32_t //col:212
hwp_package_level_request uint32_t //col:213
reserved_2 uint32_t //col:214
hdc uint32_t //col:215
intel_turbo_boost_max_technology_3_available uint32_t //col:216
hwp_capabilities uint32_t //col:217
hwp_peci_override uint32_t //col:218
flexible_hwp uint32_t //col:219
fast_access_mode_for_hwp_request_msr uint32_t //col:220
reserved_3 uint32_t //col:221
ignoring_idle_logical_processor_hwp_request uint32_t //col:222
reserved_4 uint32_t //col:223
intel_thread_director uint32_t //col:224
}


type   struct{
number_of_sub_leaves uint32_t //col:232
}


type   struct{
ia32_platform_dca_cap uint32_t //col:337
}


type   struct{
version_id_of_architectural_performance_monitoring uint32_t //col:366
number_of_performance_monitoring_counter_per_logical_processor: uint32_t //col:367
bit_width_of_performance_monitoring_counter uint32_t //col:368
ebx_bit_vector_length uint32_t //col:369
}


type   struct{
x2apic_id_to_unique_topology_id_shift uint32_t //col:401
}


type   struct{
x87_state uint32_t //col:436
sse_state uint32_t //col:437
avx_state uint32_t //col:438
mpx_state uint32_t //col:439
avx_512_state uint32_t //col:440
used_for_ia32_xss_1 uint32_t //col:441
pkru_state uint32_t //col:442
reserved_1 uint32_t //col:443
used_for_ia32_xss_2 uint32_t //col:444
}


type   struct{
reserved_1 uint32_t //col:465
supports_xsavec_and_compacted_xrstor uint32_t //col:466
supports_xgetbv_with_ecx_1 uint32_t //col:467
supports_xsave_xrstor_and_ia32_xss uint32_t //col:468
}


type   struct{
ia32_platform_dca_cap uint32_t //col:500
}


type   struct{
reserved uint32_t //col:527
}


type   struct{
reserved uint32_t //col:554
}


type   struct{
ia32_platform_dca_cap uint32_t //col:582
}


type   struct{
length_of_capacity_bit_mask uint32_t //col:611
}


type   struct{
length_of_capacity_bit_mask uint32_t //col:638
}


type   struct{
max_mba_throttling_value uint32_t //col:664
}


type   struct{
sgx1 uint32_t //col:695
sgx2 uint32_t //col:696
reserved_1 uint32_t //col:697
sgx_enclv_advanced uint32_t //col:698
sgx_encls_advanced uint32_t //col:699
}


type   struct{
valid_secs_attributes_0 uint32_t //col:722
}


type   struct{
sub_leaf_type uint32_t //col:748
}


type   struct{
sub_leaf_type uint32_t //col:776
reserved_1 uint32_t //col:777
epc_base_physical_address_1 uint32_t //col:778
}


type   struct{
max_sub_leaf uint32_t //col:804
}


type   struct{
number_of_configurable_address_ranges_for_filtering uint32_t //col:845
reserved_1 uint32_t //col:846
bitmap_of_supported_mtc_period_encodings uint32_t //col:847
}


type   struct{
denominator uint32_t //col:872
}


type   struct{
procesor_base_frequency_mhz uint32_t //col:898
}


type   struct{
max_soc_id_index uint32_t //col:924
}


type   struct{
soc_vendor_brand_string uint32_t //col:951
}


type   struct{
reserved uint32_t //col:977
}


type   struct{
max_sub_leaf uint32_t //col:1003
}


type   struct{
reserved uint32_t //col:1040
}


type   struct{
max_extended_functions uint32_t //col:1077
}


type   struct{
reserved uint32_t //col:1103
}


type   struct{
processor_brand_string_1 uint32_t //col:1141
}


type   struct{
processor_brand_string_5 uint32_t //col:1167
}


type   struct{
processor_brand_string_9 uint32_t //col:1193
}


type   struct{
reserved uint32_t //col:1219
}


type   struct{
reserved uint32_t //col:1245
}


type   struct{
reserved uint32_t //col:1274
}


type   struct{
number_of_physical_address_bits uint32_t //col:1302
number_of_linear_address_bits uint32_t //col:1303
}


type   struct{
thread_adjust uint64_t //col:1324
}


type   struct{
mseg_header_revision uint32_t //col:1334
monitor_features uint32_t //col:1335
gdtr_limit uint32_t //col:1336
gdtr_base_offset uint32_t //col:1337
cs_selector uint32_t //col:1338
eip_offset uint32_t //col:1339
esp_offset uint32_t //col:1340
cr3_offset uint32_t //col:1341
}


type   struct{
c0_mcnt uint64_t //col:1337
}


type   struct{
c0_acnt uint64_t //col:1340
}


type   struct{
stall_cycle_cnt uint64_t //col:1343
}


type   struct{
limit uint16_t //col:1347
base_address uint32_t //col:1348
}


type   struct{
limit uint16_t //col:1351
base_address uint64_t //col:1352
}


type   struct{
segment_limit_low uint16_t //col:1368
base_address_low uint16_t //col:1369
base_address_middle uint32_t //col:1372
type uint32_t //col:1373
descriptor_type uint32_t //col:1374
descriptor_privilege_level uint32_t //col:1375
present uint32_t //col:1376
segment_limit_high uint32_t //col:1377
available_bit uint32_t //col:1378
long_mode uint32_t //col:1379
default_big uint32_t //col:1380
granularity uint32_t //col:1381
base_address_high uint32_t //col:1382
}


type   struct{
segment_limit_low uint16_t //col:1388
base_address_low uint16_t //col:1389
base_address_middle uint32_t //col:1392
type uint32_t //col:1393
descriptor_type uint32_t //col:1394
descriptor_privilege_level uint32_t //col:1395
present uint32_t //col:1396
segment_limit_high uint32_t //col:1397
available_bit uint32_t //col:1398
long_mode uint32_t //col:1399
default_big uint32_t //col:1400
granularity uint32_t //col:1401
base_address_high uint32_t //col:1402
}


type   struct{
offset_low uint16_t //col:1406
segment_selector uint16_t //col:1407
interrupt_stack_table uint32_t //col:1410
must_be_zero_0 uint32_t //col:1411
type uint32_t //col:1412
must_be_zero_1 uint32_t //col:1413
descriptor_privilege_level uint32_t //col:1414
present uint32_t //col:1415
offset_middle uint32_t //col:1416
}


type   struct{
reserved_0 uint32_t //col:1428
rsp0 uint64_t //col:1429
rsp1 uint64_t //col:1430
rsp2 uint64_t //col:1431
reserved_1 uint64_t //col:1432
ist1 uint64_t //col:1433
ist2 uint64_t //col:1434
ist3 uint64_t //col:1435
ist4 uint64_t //col:1436
ist5 uint64_t //col:1437
ist6 uint64_t //col:1438
ist7 uint64_t //col:1439
reserved_2 uint64_t //col:1440
reserved_3 uint16_t //col:1441
io_map_base uint16_t //col:1442
}


type   struct{
reason uint32_t //col:1436
exception_mask uint32_t //col:1437
exit uint64_t //col:1438
guest_linear_address uint64_t //col:1439
guest_physical_address uint64_t //col:1440
current_eptp_index uint16_t //col:1441
}


type   struct{
io_a[4096] uint8_t //col:1440
io_b[4096] uint8_t //col:1441
}


type   struct{
rdmsr_low[1024] uint8_t //col:1446
rdmsr_high[1024] uint8_t //col:1447
wrmsr_low[1024] uint8_t //col:1448
wrmsr_high[1024] uint8_t //col:1449
}


type   struct{
ept_pointer uint64_t //col:1450
reserved uint64_t //col:1451
}


type   struct{
vpid uint16_t //col:1456
reserved1 uint16_t //col:1457
reserved2 uint32_t //col:1458
linear_address uint64_t //col:1459
}


type   struct{
revision_id uint32_t //col:1462
shadow_vmcs_indicator uint32_t //col:1463
}


type   struct{
revision_id uint32_t //col:1470
must_be_zero uint32_t //col:1471
}


type   struct{
present uint64_t //col:1480
reserved_1 uint64_t //col:1481
context_table_pointer uint64_t //col:1482
}


type   struct{
present uint64_t //col:1498
fault_processing_disable uint64_t //col:1499
translation_type uint64_t //col:1500
reserved_1 uint64_t //col:1501
second_level_page_translation_pointer uint64_t //col:1502
}




