package out
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\ia32-doc\out\ia32.h.back

const(
typedef enum = 1  //col:1
  InvpcidIndividualAddress                                      =  0x00000000  //col:3
  InvpcidSingleContext                                          =  0x00000001  //col:4
  InvpcidAllContextWithGlobals                                  =  0x00000002  //col:5
  InvpcidAllContext                                             =  0x00000003  //col:6
)


const(
typedef enum = 1  //col:8
  VmxActive                                                     =  0x00000000  //col:10
  VmxHlt                                                        =  0x00000001  //col:11
  VmxShutdown                                                   =  0x00000002  //col:12
  VmxWaitForSipi                                                =  0x00000003  //col:13
)


const(
typedef enum = 1  //col:15
  InveptSingleContext                                           =  0x00000001  //col:17
  InveptAllContext                                              =  0x00000002  //col:18
)


const(
typedef enum = 1  //col:20
  InvvpidIndividualAddress                                      =  0x00000000  //col:22
  InvvpidSingleContext                                          =  0x00000001  //col:23
  InvvpidAllContext                                             =  0x00000002  //col:24
  InvvpidSingleContextRetainingGlobals                          =  0x00000003  //col:25
)


const(
typedef enum = 1  //col:27
  ExternalInterrupt                                             =  0x00000000  //col:29
  NonMaskableInterrupt                                          =  0x00000002  //col:30
  HardwareException                                             =  0x00000003  //col:31
  SoftwareInterrupt                                             =  0x00000004  //col:32
  PrivilegedSoftwareException                                   =  0x00000005  //col:33
  SoftwareException                                             =  0x00000006  //col:34
  OtherEvent                                                    =  0x00000007  //col:35
)


const(
typedef enum = 1  //col:37
  DivideError                                                   =  0x00000000  //col:39
  Debug                                                         =  0x00000001  //col:40
  Nmi                                                           =  0x00000002  //col:41
  Breakpoint                                                    =  0x00000003  //col:42
  Overflow                                                      =  0x00000004  //col:43
  BoundRangeExceeded                                            =  0x00000005  //col:44
  InvalidOpcode                                                 =  0x00000006  //col:45
  DeviceNotAvailable                                            =  0x00000007  //col:46
  DoubleFault                                                   =  0x00000008  //col:47
  CoprocessorSegmentOverrun                                     =  0x00000009  //col:48
  InvalidTss                                                    =  0x0000000A  //col:49
  SegmentNotPresent                                             =  0x0000000B  //col:50
  StackSegmentFault                                             =  0x0000000C  //col:51
  GeneralProtection                                             =  0x0000000D  //col:52
  PageFault                                                     =  0x0000000E  //col:53
  X87FloatingPointError                                         =  0x00000010  //col:54
  AlignmentCheck                                                =  0x00000011  //col:55
  MachineCheck                                                  =  0x00000012  //col:56
  SimdFloatingPointError                                        =  0x00000013  //col:57
  VirtualizationException                                       =  0x00000014  //col:58
  ControlProtection                                             =  0x00000015  //col:59
)



type  struct{
TypedefStruct typedef struct //col:1
MaxCpuidInputValue uint32 //col:3
EbxValueGenu uint32 //col:4
EcxValueNtel uint32 //col:5
EdxValueInei uint32 //col:6
}


type  struct{
TypedefStruct typedef struct //col:8
Union union //col:10
Struct struct //col:12
SteppingId uint32 //col:14
Model uint32 //col:15
FamilyId uint32 //col:16
ProcessorType uint32 //col:17
Reserved1 uint32 //col:18
ExtendedModelId uint32 //col:19
ExtendedFamilyId uint32 //col:20
Reserved2 uint32 //col:21
}


type  struct{
TypedefStruct typedef struct //col:115
Union union //col:117
Struct struct //col:119
CacheTypeField uint32 //col:121
CacheLevel uint32 //col:122
SelfInitializingCacheLevel uint32 //col:123
FullyAssociativeCache uint32 //col:124
Reserved1 uint32 //col:125
MaxAddressableIdsForLogicalProcessorsSharingThisCache uint32 //col:126
MaxAddressableIdsForProcessorCoresInPhysicalPackage uint32 //col:127
}


type  struct{
TypedefStruct typedef struct //col:161
Union union //col:163
Struct struct //col:165
SmallestMonitorLineSize uint32 //col:167
Reserved1 uint32 //col:168
}


type  struct{
TypedefStruct typedef struct //col:207
Union union //col:209
Struct struct //col:211
TemperatureSensorSupported uint32 //col:213
IntelTurboBoostTechnologyAvailable uint32 //col:214
ApicTimerAlwaysRunning uint32 //col:215
Reserved1 uint32 //col:216
PowerLimitNotification uint32 //col:217
ClockModulationDuty uint32 //col:218
PackageThermalManagement uint32 //col:219
HwpBaseRegisters uint32 //col:220
HwpNotification uint32 //col:221
HwpActivityWindow uint32 //col:222
HwpEnergyPerformancePreference uint32 //col:223
HwpPackageLevelRequest uint32 //col:224
Reserved2 uint32 //col:225
Hdc uint32 //col:226
IntelTurboBoostMaxTechnology3Available uint32 //col:227
HwpCapabilities uint32 //col:228
HwpPeciOverride uint32 //col:229
FlexibleHwp uint32 //col:230
FastAccessModeForHwpRequestMsr uint32 //col:231
Reserved3 uint32 //col:232
IgnoringIdleLogicalProcessorHwpRequest uint32 //col:233
Reserved4 uint32 //col:234
IntelThreadDirector uint32 //col:235
Reserved5 uint32 //col:236
}


type  struct{
TypedefStruct typedef struct //col:271
Union union //col:273
Struct struct //col:275
NumberOfSubLeaves uint32 //col:277
}


type  struct{
TypedefStruct typedef struct //col:385
Union union //col:387
Struct struct //col:389
Ia32PlatformDcaCap uint32 //col:391
}


type  struct{
TypedefStruct typedef struct //col:420
Union union //col:422
Struct struct //col:424
VersionIdOfArchitecturalPerformanceMonitoring uint32 //col:426
NumberOfPerformanceMonitoringCounterPerLogicalProcessor: uint32 //col:427
BitWidthOfPerformanceMonitoringCounter uint32 //col:428
EbxBitVectorLength uint32 //col:429
}


type  struct{
TypedefStruct typedef struct //col:469
Union union //col:471
Struct struct //col:473
X2ApicIdToUniqueTopologyIdShift uint32 //col:475
Reserved1 uint32 //col:476
}


type  struct{
TypedefStruct typedef struct //col:508
Union union //col:510
Struct struct //col:512
X87State uint32 //col:514
SseState uint32 //col:515
AvxState uint32 //col:516
MpxState uint32 //col:517
Avx512State uint32 //col:518
UsedForIa32Xss1 uint32 //col:519
PkruState uint32 //col:520
Reserved1 uint32 //col:521
UsedForIa32Xss2 uint32 //col:522
Reserved2 uint32 //col:523
}


type  struct{
TypedefStruct typedef struct //col:552
Union union //col:554
Struct struct //col:556
Reserved1 uint32 //col:558
SupportsXsavecAndCompactedXrstor uint32 //col:559
SupportsXgetbvWithEcx1 uint32 //col:560
SupportsXsaveXrstorAndIa32Xss uint32 //col:561
Reserved2 uint32 //col:562
}


type  struct{
TypedefStruct typedef struct //col:601
Union union //col:603
Struct struct //col:605
Ia32PlatformDcaCap uint32 //col:607
}


type  struct{
TypedefStruct typedef struct //col:638
Union union //col:640
Struct struct //col:642
Reserved uint32 //col:644
}


type  struct{
TypedefStruct typedef struct //col:675
Union union //col:677
Struct struct //col:679
Reserved uint32 //col:681
}


type  struct{
TypedefStruct typedef struct //col:713
Union union //col:715
Struct struct //col:717
Ia32PlatformDcaCap uint32 //col:719
}


type  struct{
TypedefStruct typedef struct //col:752
Union union //col:754
Struct struct //col:756
LengthOfCapacityBitMask uint32 //col:758
Reserved1 uint32 //col:759
}


type  struct{
TypedefStruct typedef struct //col:791
Union union //col:793
Struct struct //col:795
LengthOfCapacityBitMask uint32 //col:797
Reserved1 uint32 //col:798
}


type  struct{
TypedefStruct typedef struct //col:828
Union union //col:830
Struct struct //col:832
MaxMbaThrottlingValue uint32 //col:834
Reserved1 uint32 //col:835
}


type  struct{
TypedefStruct typedef struct //col:867
Union union //col:869
Struct struct //col:871
Sgx1 uint32 //col:873
Sgx2 uint32 //col:874
Reserved1 uint32 //col:875
SgxEnclvAdvanced uint32 //col:876
SgxEnclsAdvanced uint32 //col:877
Reserved2 uint32 //col:878
}


type  struct{
TypedefStruct typedef struct //col:909
Union union //col:911
Struct struct //col:913
ValidSecsAttributes0 uint32 //col:915
}


type  struct{
TypedefStruct typedef struct //col:944
Union union //col:946
Struct struct //col:948
SubLeafType uint32 //col:950
Reserved1 uint32 //col:951
}


type  struct{
TypedefStruct typedef struct //col:980
Union union //col:982
Struct struct //col:984
SubLeafType uint32 //col:986
Reserved1 uint32 //col:987
EpcBasePhysicalAddress1 uint32 //col:988
}


type  struct{
TypedefStruct typedef struct //col:1021
Union union //col:1023
Struct struct //col:1025
MaxSubLeaf uint32 //col:1027
}


type  struct{
TypedefStruct typedef struct //col:1070
Union union //col:1072
Struct struct //col:1074
NumberOfConfigurableAddressRangesForFiltering uint32 //col:1076
Reserved1 uint32 //col:1077
BitmapOfSupportedMtcPeriodEncodings uint32 //col:1078
}


type  struct{
TypedefStruct typedef struct //col:1108
Union union //col:1110
Struct struct //col:1112
Denominator uint32 //col:1114
}


type  struct{
TypedefStruct typedef struct //col:1143
Union union //col:1145
Struct struct //col:1147
ProcesorBaseFrequencyMhz uint32 //col:1149
Reserved1 uint32 //col:1150
}


type  struct{
TypedefStruct typedef struct //col:1181
Union union //col:1183
Struct struct //col:1185
MaxSocIdIndex uint32 //col:1187
}


type  struct{
TypedefStruct typedef struct //col:1218
Union union //col:1220
Struct struct //col:1222
SocVendorBrandString uint32 //col:1224
}


type  struct{
TypedefStruct typedef struct //col:1253
Union union //col:1255
Struct struct //col:1257
Reserved uint32 //col:1259
}


type  struct{
TypedefStruct typedef struct //col:1288
Union union //col:1290
Struct struct //col:1292
MaxSubLeaf uint32 //col:1294
}


type  struct{
TypedefStruct typedef struct //col:1335
Union union //col:1337
Struct struct //col:1339
Reserved uint32 //col:1341
}


type  struct{
TypedefStruct typedef struct //col:1382
Union union //col:1384
Struct struct //col:1386
MaxExtendedFunctions uint32 //col:1388
}


type  struct{
TypedefStruct typedef struct //col:1417
Union union //col:1419
Struct struct //col:1421
Reserved uint32 //col:1423
}


type  struct{
TypedefStruct typedef struct //col:1466
Union union //col:1468
Struct struct //col:1470
ProcessorBrandString1 uint32 //col:1472
}


type  struct{
TypedefStruct typedef struct //col:1501
Union union //col:1503
Struct struct //col:1505
ProcessorBrandString5 uint32 //col:1507
}


type  struct{
TypedefStruct typedef struct //col:1536
Union union //col:1538
Struct struct //col:1540
ProcessorBrandString9 uint32 //col:1542
}


type  struct{
TypedefStruct typedef struct //col:1571
Union union //col:1573
Struct struct //col:1575
Reserved uint32 //col:1577
}


type  struct{
TypedefStruct typedef struct //col:1606
Union union //col:1608
Struct struct //col:1610
Reserved uint32 //col:1612
}


type  struct{
TypedefStruct typedef struct //col:1644
Union union //col:1646
Struct struct //col:1648
Reserved uint32 //col:1650
}


type  struct{
TypedefStruct typedef struct //col:1681
Union union //col:1683
Struct struct //col:1685
NumberOfPhysicalAddressBits uint32 //col:1687
NumberOfLinearAddressBits uint32 //col:1688
Reserved1 uint32 //col:1689
}


type  struct{
TypedefStruct typedef struct //col:1718
ThreadAdjust uint64 //col:1720
}


type  struct{
TypedefStruct typedef struct //col:1722
MsegHeaderRevision uint32 //col:1724
MonitorFeatures uint32 //col:1725
GdtrLimit uint32 //col:1726
GdtrBaseOffset uint32 //col:1727
CsSelector uint32 //col:1728
EipOffset uint32 //col:1729
EspOffset uint32 //col:1730
Cr3Offset uint32 //col:1731
}


type  struct{
TypedefStruct typedef struct //col:1733
C0Mcnt uint64 //col:1735
}


type  struct{
TypedefStruct typedef struct //col:1737
C0Acnt uint64 //col:1739
}


type  struct{
TypedefStruct typedef struct //col:1741
StallCycleCount uint64 //col:1743
}


type  struct{
TypedefStruct typedef struct //col:1745
Limit uint16 //col:1747
BaseAddress uint32 //col:1748
}


type  struct{
TypedefStruct typedef struct //col:1750
Limit uint16 //col:1752
BaseAddress uint64 //col:1753
}


type  struct{
TypedefStruct typedef struct //col:1755
SegmentLimitLow uint16 //col:1757
BaseAddressLow uint16 //col:1758
Union union //col:1759
Struct struct //col:1761
BaseAddressMiddle uint32 //col:1763
Type uint32 //col:1764
DescriptorType uint32 //col:1765
DescriptorPrivilegeLevel uint32 //col:1766
Present uint32 //col:1767
SegmentLimitHigh uint32 //col:1768
System uint32 //col:1769
LongMode uint32 //col:1770
DefaultBig uint32 //col:1771
Granularity uint32 //col:1772
BaseAddressHigh uint32 //col:1773
}


type  struct{
TypedefStruct typedef struct //col:1778
SegmentLimitLow uint16 //col:1780
BaseAddressLow uint16 //col:1781
Union union //col:1782
Struct struct //col:1784
BaseAddressMiddle uint32 //col:1786
Type uint32 //col:1787
DescriptorType uint32 //col:1788
DescriptorPrivilegeLevel uint32 //col:1789
Present uint32 //col:1790
SegmentLimitHigh uint32 //col:1791
System uint32 //col:1792
LongMode uint32 //col:1793
DefaultBig uint32 //col:1794
Granularity uint32 //col:1795
BaseAddressHigh uint32 //col:1796
}


type  struct{
TypedefStruct typedef struct //col:1803
OffsetLow uint16 //col:1805
SegmentSelector uint16 //col:1806
Union union //col:1807
Struct struct //col:1809
InterruptStackTable uint32 //col:1811
MustBeZero0 uint32 //col:1812
Type uint32 //col:1813
MustBeZero1 uint32 //col:1814
DescriptorPrivilegeLevel uint32 //col:1815
Present uint32 //col:1816
OffsetMiddle uint32 //col:1817
}


type  struct{
TypedefStruct typedef struct //col:1824
Reserved0 uint32 //col:1826
Rsp0 uint64 //col:1827
Rsp1 uint64 //col:1828
Rsp2 uint64 //col:1829
Reserved1 uint64 //col:1830
Ist1 uint64 //col:1831
Ist2 uint64 //col:1832
Ist3 uint64 //col:1833
Ist4 uint64 //col:1834
Ist5 uint64 //col:1835
Ist6 uint64 //col:1836
Ist7 uint64 //col:1837
Reserved2 uint64 //col:1838
Reserved3 uint16 //col:1839
IoMapBase uint16 //col:1840
}


type  struct{
TypedefStruct typedef struct //col:1842
Reason uint32 //col:1844
ExceptionMask uint32 //col:1845
Exit uint64 //col:1846
GuestLinearAddress uint64 //col:1847
GuestPhysicalAddress uint64 //col:1848
CurrentEptpIndex uint16 //col:1849
}


type  struct{
TypedefStruct typedef struct //col:1851
IoA[4096] UINT8 //col:1853
IoB[4096] UINT8 //col:1854
}


type  struct{
TypedefStruct typedef struct //col:1856
RdmsrLow[1024] UINT8 //col:1858
RdmsrHigh[1024] UINT8 //col:1859
WrmsrLow[1024] UINT8 //col:1860
WrmsrHigh[1024] UINT8 //col:1861
}


type  struct{
TypedefStruct typedef struct //col:1863
EptPointer uint64 //col:1865
Reserved uint64 //col:1866
}


type  struct{
TypedefStruct typedef struct //col:1868
Vpid uint16 //col:1870
Reserved1 uint16 //col:1871
Reserved2 uint32 //col:1872
LinearAddress uint64 //col:1873
}


type  struct{
TypedefStruct typedef struct //col:1875
Struct struct //col:1877
RevisionId uint32 //col:1879
ShadowVmcsIndicator uint32 //col:1880
}


type  struct{
TypedefStruct typedef struct //col:1885
Struct struct //col:1887
RevisionId uint32 //col:1889
MustBeZero uint32 //col:1890
}


type  struct{
TypedefStruct typedef struct //col:1894
Union union //col:1896
Struct struct //col:1898
Present uint64 //col:1900
Reserved1 uint64 //col:1901
ContextTablePointer uint64 //col:1902
}


type  struct{
TypedefStruct typedef struct //col:1915
Union union //col:1917
Struct struct //col:1919
Present uint64 //col:1921
FaultProcessingDisable uint64 //col:1922
TranslationType uint64 //col:1923
Reserved1 uint64 //col:1924
SecondLevelPageTranslationPointer uint64 //col:1925
}




