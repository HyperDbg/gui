package out
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\ia32-doc\out\ia32_defines_only.h.back

type   struct{
max_cpuid_input_value uint32_t //col:7
ebx_value_genu uint32_t //col:8
ecx_value_ntel uint32_t //col:9
edx_value_inei uint32_t //col:10
}


type   struct{
stepping_id uint32_t //col:21
model uint32_t //col:22
family_id uint32_t //col:23
processor_type uint32_t //col:24
reserved_1 uint32_t //col:25
extended_model_id uint32_t //col:26
extended_family_id uint32_t //col:27
reserved_2 uint32_t //col:28
}


type   struct{
cache_type_field uint32_t //col:118
cache_level uint32_t //col:119
self_initializing_cache_level uint32_t //col:120
fully_associative_cache uint32_t //col:121
reserved_1 uint32_t //col:122
max_addressable_ids_for_logical_processors_sharing_this_cache: uint32_t //col:123
max_addressable_ids_for_processor_cores_in_physical_package: uint32_t //col:124
}


type   struct{
smallest_monitor_line_size uint32_t //col:150
reserved_1 uint32_t //col:151
}


type   struct{
temperature_sensor_supported uint32_t //col:209
intel_turbo_boost_technology_available uint32_t //col:210
apic_timer_always_running uint32_t //col:211
reserved_1 uint32_t //col:212
power_limit_notification uint32_t //col:213
clock_modulation_duty uint32_t //col:214
package_thermal_management uint32_t //col:215
hwp_base_registers uint32_t //col:216
hwp_notification uint32_t //col:217
hwp_activity_window uint32_t //col:218
hwp_energy_performance_preference uint32_t //col:219
hwp_package_level_request uint32_t //col:220
reserved_2 uint32_t //col:221
hdc uint32_t //col:222
intel_turbo_boost_max_technology_3_available uint32_t //col:223
hwp_capabilities uint32_t //col:224
hwp_peci_override uint32_t //col:225
flexible_hwp uint32_t //col:226
fast_access_mode_for_hwp_request_msr uint32_t //col:227
reserved_3 uint32_t //col:228
ignoring_idle_logical_processor_hwp_request uint32_t //col:229
reserved_4 uint32_t //col:230
intel_thread_director uint32_t //col:231
reserved_5 uint32_t //col:232
}


type   struct{
number_of_sub_leaves uint32_t //col:241
}


type   struct{
ia32_platform_dca_cap uint32_t //col:346
}


type   struct{
version_id_of_architectural_performance_monitoring uint32_t //col:375
number_of_performance_monitoring_counter_per_logical_processor: uint32_t //col:376
bit_width_of_performance_monitoring_counter uint32_t //col:377
ebx_bit_vector_length uint32_t //col:378
}


type   struct{
x2apic_id_to_unique_topology_id_shift uint32_t //col:413
reserved_1 uint32_t //col:414
}


type   struct{
x87_state uint32_t //col:451
sse_state uint32_t //col:452
avx_state uint32_t //col:453
mpx_state uint32_t //col:454
avx_512_state uint32_t //col:455
used_for_ia32_xss_1 uint32_t //col:456
pkru_state uint32_t //col:457
reserved_1 uint32_t //col:458
used_for_ia32_xss_2 uint32_t //col:459
reserved_2 uint32_t //col:460
}


type   struct{
reserved_1 uint32_t //col:481
supports_xsavec_and_compacted_xrstor uint32_t //col:482
supports_xgetbv_with_ecx_1 uint32_t //col:483
supports_xsave_xrstor_and_ia32_xss uint32_t //col:484
reserved_2 uint32_t //col:485
}


type   struct{
ia32_platform_dca_cap uint32_t //col:517
}


type   struct{
reserved uint32_t //col:545
}


type   struct{
reserved uint32_t //col:573
}


type   struct{
ia32_platform_dca_cap uint32_t //col:602
}


type   struct{
length_of_capacity_bit_mask uint32_t //col:633
reserved_1 uint32_t //col:634
}


type   struct{
length_of_capacity_bit_mask uint32_t //col:663
reserved_1 uint32_t //col:664
}


type   struct{
max_mba_throttling_value uint32_t //col:691
reserved_1 uint32_t //col:692
}


type   struct{
sgx1 uint32_t //col:725
sgx2 uint32_t //col:726
reserved_1 uint32_t //col:727
sgx_enclv_advanced uint32_t //col:728
sgx_encls_advanced uint32_t //col:729
reserved_2 uint32_t //col:730
}


type   struct{
valid_secs_attributes_0 uint32_t //col:753
}


type   struct{
sub_leaf_type uint32_t //col:780
reserved_1 uint32_t //col:781
}


type   struct{
sub_leaf_type uint32_t //col:808
reserved_1 uint32_t //col:809
epc_base_physical_address_1 uint32_t //col:810
}


type   struct{
max_sub_leaf uint32_t //col:838
}


type   struct{
number_of_configurable_address_ranges_for_filtering uint32_t //col:880
reserved_1 uint32_t //col:881
bitmap_of_supported_mtc_period_encodings uint32_t //col:882
}


type   struct{
denominator uint32_t //col:907
}


type   struct{
procesor_base_frequency_mhz uint32_t //col:934
reserved_1 uint32_t //col:935
}


type   struct{
max_soc_id_index uint32_t //col:962
}


type   struct{
soc_vendor_brand_string uint32_t //col:990
}


type   struct{
reserved uint32_t //col:1016
}


type   struct{
max_sub_leaf uint32_t //col:1042
}


type   struct{
reserved uint32_t //col:1080
}


type   struct{
max_extended_functions uint32_t //col:1118
}


type   struct{
reserved uint32_t //col:1144
}


type   struct{
processor_brand_string_1 uint32_t //col:1184
}


type   struct{
processor_brand_string_5 uint32_t //col:1210
}


type   struct{
processor_brand_string_9 uint32_t //col:1236
}


type   struct{
reserved uint32_t //col:1262
}


type   struct{
reserved uint32_t //col:1288
}


type   struct{
reserved uint32_t //col:1317
}


type   struct{
number_of_physical_address_bits uint32_t //col:1347
number_of_linear_address_bits uint32_t //col:1348
reserved_1 uint32_t //col:1349
}


type   struct{
thread_adjust uint64_t //col:1369
}


type   struct{
mseg_header_revision uint32_t //col:1379
monitor_features uint32_t //col:1380
gdtr_limit uint32_t //col:1381
gdtr_base_offset uint32_t //col:1382
cs_selector uint32_t //col:1383
eip_offset uint32_t //col:1384
esp_offset uint32_t //col:1385
cr3_offset uint32_t //col:1386
}


type   struct{
c0_mcnt uint64_t //col:1382
}


type   struct{
c0_acnt uint64_t //col:1385
}


type   struct{
stall_cycle_cnt uint64_t //col:1388
}


type   struct{
limit uint16_t //col:1392
base_address uint32_t //col:1393
}


type   struct{
limit uint16_t //col:1396
base_address uint64_t //col:1397
}


type   struct{
segment_limit_low uint16_t //col:1413
base_address_low uint16_t //col:1414
base_address_middle uint32_t //col:1417
type uint32_t //col:1418
descriptor_type uint32_t //col:1419
descriptor_privilege_level uint32_t //col:1420
present uint32_t //col:1421
segment_limit_high uint32_t //col:1422
available_bit uint32_t //col:1423
long_mode uint32_t //col:1424
default_big uint32_t //col:1425
granularity uint32_t //col:1426
base_address_high uint32_t //col:1427
}


type   struct{
segment_limit_low uint16_t //col:1433
base_address_low uint16_t //col:1434
base_address_middle uint32_t //col:1437
type uint32_t //col:1438
descriptor_type uint32_t //col:1439
descriptor_privilege_level uint32_t //col:1440
present uint32_t //col:1441
segment_limit_high uint32_t //col:1442
available_bit uint32_t //col:1443
long_mode uint32_t //col:1444
default_big uint32_t //col:1445
granularity uint32_t //col:1446
base_address_high uint32_t //col:1447
}


type   struct{
offset_low uint16_t //col:1451
segment_selector uint16_t //col:1452
interrupt_stack_table uint32_t //col:1455
must_be_zero_0 uint32_t //col:1456
type uint32_t //col:1457
must_be_zero_1 uint32_t //col:1458
descriptor_privilege_level uint32_t //col:1459
present uint32_t //col:1460
offset_middle uint32_t //col:1461
}


type   struct{
reserved_0 uint32_t //col:1473
rsp0 uint64_t //col:1474
rsp1 uint64_t //col:1475
rsp2 uint64_t //col:1476
reserved_1 uint64_t //col:1477
ist1 uint64_t //col:1478
ist2 uint64_t //col:1479
ist3 uint64_t //col:1480
ist4 uint64_t //col:1481
ist5 uint64_t //col:1482
ist6 uint64_t //col:1483
ist7 uint64_t //col:1484
reserved_2 uint64_t //col:1485
reserved_3 uint16_t //col:1486
io_map_base uint16_t //col:1487
}


type   struct{
reason uint32_t //col:1481
exception_mask uint32_t //col:1482
exit uint64_t //col:1483
guest_linear_address uint64_t //col:1484
guest_physical_address uint64_t //col:1485
current_eptp_index uint16_t //col:1486
}


type   struct{
io_a[4096] uint8_t //col:1485
io_b[4096] uint8_t //col:1486
}


type   struct{
rdmsr_low[1024] uint8_t //col:1491
rdmsr_high[1024] uint8_t //col:1492
wrmsr_low[1024] uint8_t //col:1493
wrmsr_high[1024] uint8_t //col:1494
}


type   struct{
ept_pointer uint64_t //col:1495
reserved uint64_t //col:1496
}


type   struct{
vpid uint16_t //col:1501
reserved1 uint16_t //col:1502
reserved2 uint32_t //col:1503
linear_address uint64_t //col:1504
}


type   struct{
revision_id uint32_t //col:1507
shadow_vmcs_indicator uint32_t //col:1508
}


type   struct{
revision_id uint32_t //col:1515
must_be_zero uint32_t //col:1516
}


type   struct{
present uint64_t //col:1525
reserved_1 uint64_t //col:1526
context_table_pointer uint64_t //col:1527
}


type   struct{
present uint64_t //col:1543
fault_processing_disable uint64_t //col:1544
translation_type uint64_t //col:1545
reserved_1 uint64_t //col:1546
second_level_page_translation_pointer uint64_t //col:1547
}




