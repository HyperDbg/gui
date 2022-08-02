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
TypedefStruct typedef struct //col:7
MaxCpuidInputValue uint32 //col:9
EbxValueGenu uint32 //col:10
EcxValueNtel uint32 //col:11
EdxValueInei uint32 //col:12
}


type  struct{
TypedefStruct typedef struct //col:22
Union union //col:24
Struct struct //col:26
SteppingId uint32 //col:28
Model uint32 //col:29
FamilyId uint32 //col:30
ProcessorType uint32 //col:31
Reserved1 uint32 //col:32
ExtendedModelId uint32 //col:33
ExtendedFamilyId uint32 //col:34
Reserved2 uint32 //col:35
}


type  struct{
TypedefStruct typedef struct //col:128
Union union //col:130
Struct struct //col:132
CacheTypeField uint32 //col:134
CacheLevel uint32 //col:135
SelfInitializingCacheLevel uint32 //col:136
FullyAssociativeCache uint32 //col:137
Reserved1 uint32 //col:138
MaxAddressableIdsForLogicalProcessorsSharingThisCache uint32 //col:139
MaxAddressableIdsForProcessorCoresInPhysicalPackage uint32 //col:140
}


type  struct{
TypedefStruct typedef struct //col:169
Union union //col:171
Struct struct //col:173
SmallestMonitorLineSize uint32 //col:175
Reserved1 uint32 //col:176
}


type  struct{
TypedefStruct typedef struct //col:237
Union union //col:239
Struct struct //col:241
TemperatureSensorSupported uint32 //col:243
IntelTurboBoostTechnologyAvailable uint32 //col:244
ApicTimerAlwaysRunning uint32 //col:245
Reserved1 uint32 //col:246
PowerLimitNotification uint32 //col:247
ClockModulationDuty uint32 //col:248
PackageThermalManagement uint32 //col:249
HwpBaseRegisters uint32 //col:250
HwpNotification uint32 //col:251
HwpActivityWindow uint32 //col:252
HwpEnergyPerformancePreference uint32 //col:253
HwpPackageLevelRequest uint32 //col:254
Reserved2 uint32 //col:255
Hdc uint32 //col:256
IntelTurboBoostMaxTechnology3Available uint32 //col:257
HwpCapabilities uint32 //col:258
HwpPeciOverride uint32 //col:259
FlexibleHwp uint32 //col:260
FastAccessModeForHwpRequestMsr uint32 //col:261
Reserved3 uint32 //col:262
IgnoringIdleLogicalProcessorHwpRequest uint32 //col:263
Reserved4 uint32 //col:264
IntelThreadDirector uint32 //col:265
Reserved5 uint32 //col:266
}


type  struct{
TypedefStruct typedef struct //col:278
Union union //col:280
Struct struct //col:282
NumberOfSubLeaves uint32 //col:284
}


type  struct{
TypedefStruct typedef struct //col:392
Union union //col:394
Struct struct //col:396
Ia32PlatformDcaCap uint32 //col:398
}


type  struct{
TypedefStruct typedef struct //col:430
Union union //col:432
Struct struct //col:434
VersionIdOfArchitecturalPerformanceMonitoring uint32 //col:436
NumberOfPerformanceMonitoringCounterPerLogicalProcessor: uint32 //col:437
BitWidthOfPerformanceMonitoringCounter uint32 //col:438
EbxBitVectorLength uint32 //col:439
}


type  struct{
TypedefStruct typedef struct //col:477
Union union //col:479
Struct struct //col:481
X2ApicIdToUniqueTopologyIdShift uint32 //col:483
Reserved1 uint32 //col:484
}


type  struct{
TypedefStruct typedef struct //col:524
Union union //col:526
Struct struct //col:528
X87State uint32 //col:530
SseState uint32 //col:531
AvxState uint32 //col:532
MpxState uint32 //col:533
Avx512State uint32 //col:534
UsedForIa32Xss1 uint32 //col:535
PkruState uint32 //col:536
Reserved1 uint32 //col:537
UsedForIa32Xss2 uint32 //col:538
Reserved2 uint32 //col:539
}


type  struct{
TypedefStruct typedef struct //col:563
Union union //col:565
Struct struct //col:567
Reserved1 uint32 //col:569
SupportsXsavecAndCompactedXrstor uint32 //col:570
SupportsXgetbvWithEcx1 uint32 //col:571
SupportsXsaveXrstorAndIa32Xss uint32 //col:572
Reserved2 uint32 //col:573
}


type  struct{
TypedefStruct typedef struct //col:608
Union union //col:610
Struct struct //col:612
Ia32PlatformDcaCap uint32 //col:614
}


type  struct{
TypedefStruct typedef struct //col:645
Union union //col:647
Struct struct //col:649
Reserved uint32 //col:651
}


type  struct{
TypedefStruct typedef struct //col:682
Union union //col:684
Struct struct //col:686
Reserved uint32 //col:688
}


type  struct{
TypedefStruct typedef struct //col:720
Union union //col:722
Struct struct //col:724
Ia32PlatformDcaCap uint32 //col:726
}


type  struct{
TypedefStruct typedef struct //col:760
Union union //col:762
Struct struct //col:764
LengthOfCapacityBitMask uint32 //col:766
Reserved1 uint32 //col:767
}


type  struct{
TypedefStruct typedef struct //col:799
Union union //col:801
Struct struct //col:803
LengthOfCapacityBitMask uint32 //col:805
Reserved1 uint32 //col:806
}


type  struct{
TypedefStruct typedef struct //col:836
Union union //col:838
Struct struct //col:840
MaxMbaThrottlingValue uint32 //col:842
Reserved1 uint32 //col:843
}


type  struct{
TypedefStruct typedef struct //col:879
Union union //col:881
Struct struct //col:883
Sgx1 uint32 //col:885
Sgx2 uint32 //col:886
Reserved1 uint32 //col:887
SgxEnclvAdvanced uint32 //col:888
SgxEnclsAdvanced uint32 //col:889
Reserved2 uint32 //col:890
}


type  struct{
TypedefStruct typedef struct //col:916
Union union //col:918
Struct struct //col:920
ValidSecsAttributes0 uint32 //col:922
}


type  struct{
TypedefStruct typedef struct //col:952
Union union //col:954
Struct struct //col:956
SubLeafType uint32 //col:958
Reserved1 uint32 //col:959
}


type  struct{
TypedefStruct typedef struct //col:989
Union union //col:991
Struct struct //col:993
SubLeafType uint32 //col:995
Reserved1 uint32 //col:996
EpcBasePhysicalAddress1 uint32 //col:997
}


type  struct{
TypedefStruct typedef struct //col:1028
Union union //col:1030
Struct struct //col:1032
MaxSubLeaf uint32 //col:1034
}


type  struct{
TypedefStruct typedef struct //col:1079
Union union //col:1081
Struct struct //col:1083
NumberOfConfigurableAddressRangesForFiltering uint32 //col:1085
Reserved1 uint32 //col:1086
BitmapOfSupportedMtcPeriodEncodings uint32 //col:1087
}


type  struct{
TypedefStruct typedef struct //col:1115
Union union //col:1117
Struct struct //col:1119
Denominator uint32 //col:1121
}


type  struct{
TypedefStruct typedef struct //col:1151
Union union //col:1153
Struct struct //col:1155
ProcesorBaseFrequencyMhz uint32 //col:1157
Reserved1 uint32 //col:1158
}


type  struct{
TypedefStruct typedef struct //col:1188
Union union //col:1190
Struct struct //col:1192
MaxSocIdIndex uint32 //col:1194
}


type  struct{
TypedefStruct typedef struct //col:1225
Union union //col:1227
Struct struct //col:1229
SocVendorBrandString uint32 //col:1231
}


type  struct{
TypedefStruct typedef struct //col:1260
Union union //col:1262
Struct struct //col:1264
Reserved uint32 //col:1266
}


type  struct{
TypedefStruct typedef struct //col:1295
Union union //col:1297
Struct struct //col:1299
MaxSubLeaf uint32 //col:1301
}


type  struct{
TypedefStruct typedef struct //col:1342
Union union //col:1344
Struct struct //col:1346
Reserved uint32 //col:1348
}


type  struct{
TypedefStruct typedef struct //col:1389
Union union //col:1391
Struct struct //col:1393
MaxExtendedFunctions uint32 //col:1395
}


type  struct{
TypedefStruct typedef struct //col:1424
Union union //col:1426
Struct struct //col:1428
Reserved uint32 //col:1430
}


type  struct{
TypedefStruct typedef struct //col:1473
Union union //col:1475
Struct struct //col:1477
ProcessorBrandString1 uint32 //col:1479
}


type  struct{
TypedefStruct typedef struct //col:1508
Union union //col:1510
Struct struct //col:1512
ProcessorBrandString5 uint32 //col:1514
}


type  struct{
TypedefStruct typedef struct //col:1543
Union union //col:1545
Struct struct //col:1547
ProcessorBrandString9 uint32 //col:1549
}


type  struct{
TypedefStruct typedef struct //col:1578
Union union //col:1580
Struct struct //col:1582
Reserved uint32 //col:1584
}


type  struct{
TypedefStruct typedef struct //col:1613
Union union //col:1615
Struct struct //col:1617
Reserved uint32 //col:1619
}


type  struct{
TypedefStruct typedef struct //col:1651
Union union //col:1653
Struct struct //col:1655
Reserved uint32 //col:1657
}


type  struct{
TypedefStruct typedef struct //col:1690
Union union //col:1692
Struct struct //col:1694
NumberOfPhysicalAddressBits uint32 //col:1696
NumberOfLinearAddressBits uint32 //col:1697
Reserved1 uint32 //col:1698
}


type  struct{
TypedefStruct typedef struct //col:1721
ThreadAdjust uint64 //col:1723
}


type  struct{
TypedefStruct typedef struct //col:1732
MsegHeaderRevision uint32 //col:1734
MonitorFeatures uint32 //col:1735
GdtrLimit uint32 //col:1736
GdtrBaseOffset uint32 //col:1737
CsSelector uint32 //col:1738
EipOffset uint32 //col:1739
EspOffset uint32 //col:1740
Cr3Offset uint32 //col:1741
}


type  struct{
TypedefStruct typedef struct //col:1736
C0Mcnt uint64 //col:1738
}


type  struct{
TypedefStruct typedef struct //col:1740
C0Acnt uint64 //col:1742
}


type  struct{
TypedefStruct typedef struct //col:1744
StallCycleCount uint64 //col:1746
}


type  struct{
TypedefStruct typedef struct //col:1749
Limit uint16 //col:1751
BaseAddress uint32 //col:1752
}


type  struct{
TypedefStruct typedef struct //col:1754
Limit uint16 //col:1756
BaseAddress uint64 //col:1757
}


type  struct{
TypedefStruct typedef struct //col:1774
SegmentLimitLow uint16 //col:1776
BaseAddressLow uint16 //col:1777
Union union //col:1778
Struct struct //col:1780
BaseAddressMiddle uint32 //col:1782
Type uint32 //col:1783
DescriptorType uint32 //col:1784
DescriptorPrivilegeLevel uint32 //col:1785
Present uint32 //col:1786
SegmentLimitHigh uint32 //col:1787
System uint32 //col:1788
LongMode uint32 //col:1789
DefaultBig uint32 //col:1790
Granularity uint32 //col:1791
BaseAddressHigh uint32 //col:1792
}


type  struct{
TypedefStruct typedef struct //col:1797
SegmentLimitLow uint16 //col:1799
BaseAddressLow uint16 //col:1800
Union union //col:1801
Struct struct //col:1803
BaseAddressMiddle uint32 //col:1805
Type uint32 //col:1806
DescriptorType uint32 //col:1807
DescriptorPrivilegeLevel uint32 //col:1808
Present uint32 //col:1809
SegmentLimitHigh uint32 //col:1810
System uint32 //col:1811
LongMode uint32 //col:1812
DefaultBig uint32 //col:1813
Granularity uint32 //col:1814
BaseAddressHigh uint32 //col:1815
}


type  struct{
TypedefStruct typedef struct //col:1818
OffsetLow uint16 //col:1820
SegmentSelector uint16 //col:1821
Union union //col:1822
Struct struct //col:1824
InterruptStackTable uint32 //col:1826
MustBeZero0 uint32 //col:1827
Type uint32 //col:1828
MustBeZero1 uint32 //col:1829
DescriptorPrivilegeLevel uint32 //col:1830
Present uint32 //col:1831
OffsetMiddle uint32 //col:1832
}


type  struct{
TypedefStruct typedef struct //col:1841
Reserved0 uint32 //col:1843
Rsp0 uint64 //col:1844
Rsp1 uint64 //col:1845
Rsp2 uint64 //col:1846
Reserved1 uint64 //col:1847
Ist1 uint64 //col:1848
Ist2 uint64 //col:1849
Ist3 uint64 //col:1850
Ist4 uint64 //col:1851
Ist5 uint64 //col:1852
Ist6 uint64 //col:1853
Ist7 uint64 //col:1854
Reserved2 uint64 //col:1855
Reserved3 uint16 //col:1856
IoMapBase uint16 //col:1857
}


type  struct{
TypedefStruct typedef struct //col:1850
Reason uint32 //col:1852
ExceptionMask uint32 //col:1853
Exit uint64 //col:1854
GuestLinearAddress uint64 //col:1855
GuestPhysicalAddress uint64 //col:1856
CurrentEptpIndex uint16 //col:1857
}


type  struct{
TypedefStruct typedef struct //col:1855
IoA[4096] UINT8 //col:1857
IoB[4096] UINT8 //col:1858
}


type  struct{
TypedefStruct typedef struct //col:1862
RdmsrLow[1024] UINT8 //col:1864
RdmsrHigh[1024] UINT8 //col:1865
WrmsrLow[1024] UINT8 //col:1866
WrmsrHigh[1024] UINT8 //col:1867
}


type  struct{
TypedefStruct typedef struct //col:1867
EptPointer uint64 //col:1869
Reserved uint64 //col:1870
}


type  struct{
TypedefStruct typedef struct //col:1874
Vpid uint16 //col:1876
Reserved1 uint16 //col:1877
Reserved2 uint32 //col:1878
LinearAddress uint64 //col:1879
}


type  struct{
TypedefStruct typedef struct //col:1881
Struct struct //col:1883
RevisionId uint32 //col:1885
ShadowVmcsIndicator uint32 //col:1886
}


type  struct{
TypedefStruct typedef struct //col:1891
Struct struct //col:1893
RevisionId uint32 //col:1895
MustBeZero uint32 //col:1896
}


type  struct{
TypedefStruct typedef struct //col:1903
Union union //col:1905
Struct struct //col:1907
Present uint64 //col:1909
Reserved1 uint64 //col:1910
ContextTablePointer uint64 //col:1911
}


type  struct{
TypedefStruct typedef struct //col:1926
Union union //col:1928
Struct struct //col:1930
Present uint64 //col:1932
FaultProcessingDisable uint64 //col:1933
TranslationType uint64 //col:1934
Reserved1 uint64 //col:1935
SecondLevelPageTranslationPointer uint64 //col:1936
}




