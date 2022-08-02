package out
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\ia32-doc\out\ia32_defines_only.h.back

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




