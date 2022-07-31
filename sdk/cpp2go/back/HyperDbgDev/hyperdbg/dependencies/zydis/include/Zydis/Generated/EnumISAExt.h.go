package Generated


const(
    ZYDIS_ISA_EXT_INVALID = 1  //col:6
    ZYDIS_ISA_EXT_ADOX_ADCX = 2  //col:7
    ZYDIS_ISA_EXT_AES = 3  //col:8
    ZYDIS_ISA_EXT_AMD3DNOW = 4  //col:9
    ZYDIS_ISA_EXT_AVX = 5  //col:10
    ZYDIS_ISA_EXT_AVX2 = 6  //col:11
    ZYDIS_ISA_EXT_AVX2GATHER = 7  //col:12
    ZYDIS_ISA_EXT_AVX512EVEX = 8  //col:13
    ZYDIS_ISA_EXT_AVX512VEX = 9  //col:14
    ZYDIS_ISA_EXT_AVXAES = 10  //col:15
    ZYDIS_ISA_EXT_BASE = 11  //col:16
    ZYDIS_ISA_EXT_BMI1 = 12  //col:17
    ZYDIS_ISA_EXT_BMI2 = 13  //col:18
    ZYDIS_ISA_EXT_CET = 14  //col:19
    ZYDIS_ISA_EXT_CLDEMOTE = 15  //col:20
    ZYDIS_ISA_EXT_CLFLUSHOPT = 16  //col:21
    ZYDIS_ISA_EXT_CLFSH = 17  //col:22
    ZYDIS_ISA_EXT_CLWB = 18  //col:23
    ZYDIS_ISA_EXT_CLZERO = 19  //col:24
    ZYDIS_ISA_EXT_ENQCMD = 20  //col:25
    ZYDIS_ISA_EXT_F16C = 21  //col:26
    ZYDIS_ISA_EXT_FMA = 22  //col:27
    ZYDIS_ISA_EXT_FMA4 = 23  //col:28
    ZYDIS_ISA_EXT_GFNI = 24  //col:29
    ZYDIS_ISA_EXT_INVPCID = 25  //col:30
    ZYDIS_ISA_EXT_KNC = 26  //col:31
    ZYDIS_ISA_EXT_KNCE = 27  //col:32
    ZYDIS_ISA_EXT_KNCV = 28  //col:33
    ZYDIS_ISA_EXT_LONGMODE = 29  //col:34
    ZYDIS_ISA_EXT_LZCNT = 30  //col:35
    ZYDIS_ISA_EXT_MMX = 31  //col:36
    ZYDIS_ISA_EXT_MONITOR = 32  //col:37
    ZYDIS_ISA_EXT_MONITORX = 33  //col:38
    ZYDIS_ISA_EXT_MOVBE = 34  //col:39
    ZYDIS_ISA_EXT_MOVDIR = 35  //col:40
    ZYDIS_ISA_EXT_MPX = 36  //col:41
    ZYDIS_ISA_EXT_PADLOCK = 37  //col:42
    ZYDIS_ISA_EXT_PAUSE = 38  //col:43
    ZYDIS_ISA_EXT_PCLMULQDQ = 39  //col:44
    ZYDIS_ISA_EXT_PCONFIG = 40  //col:45
    ZYDIS_ISA_EXT_PKU = 41  //col:46
    ZYDIS_ISA_EXT_PREFETCHWT1 = 42  //col:47
    ZYDIS_ISA_EXT_PT = 43  //col:48
    ZYDIS_ISA_EXT_RDPID = 44  //col:49
    ZYDIS_ISA_EXT_RDPRU = 45  //col:50
    ZYDIS_ISA_EXT_RDRAND = 46  //col:51
    ZYDIS_ISA_EXT_RDSEED = 47  //col:52
    ZYDIS_ISA_EXT_RDTSCP = 48  //col:53
    ZYDIS_ISA_EXT_RDWRFSGS = 49  //col:54
    ZYDIS_ISA_EXT_RTM = 50  //col:55
    ZYDIS_ISA_EXT_SGX = 51  //col:56
    ZYDIS_ISA_EXT_SGX_ENCLV = 52  //col:57
    ZYDIS_ISA_EXT_SHA = 53  //col:58
    ZYDIS_ISA_EXT_SMAP = 54  //col:59
    ZYDIS_ISA_EXT_SMX = 55  //col:60
    ZYDIS_ISA_EXT_SSE = 56  //col:61
    ZYDIS_ISA_EXT_SSE2 = 57  //col:62
    ZYDIS_ISA_EXT_SSE3 = 58  //col:63
    ZYDIS_ISA_EXT_SSE4 = 59  //col:64
    ZYDIS_ISA_EXT_SSE4A = 60  //col:65
    ZYDIS_ISA_EXT_SSSE3 = 61  //col:66
    ZYDIS_ISA_EXT_SVM = 62  //col:67
    ZYDIS_ISA_EXT_TBM = 63  //col:68
    ZYDIS_ISA_EXT_VAES = 64  //col:69
    ZYDIS_ISA_EXT_VMFUNC = 65  //col:70
    ZYDIS_ISA_EXT_VPCLMULQDQ = 66  //col:71
    ZYDIS_ISA_EXT_VTX = 67  //col:72
    ZYDIS_ISA_EXT_WAITPKG = 68  //col:73
    ZYDIS_ISA_EXT_X87 = 69  //col:74
    ZYDIS_ISA_EXT_XOP = 70  //col:75
    ZYDIS_ISA_EXT_XSAVE = 71  //col:76
    ZYDIS_ISA_EXT_XSAVEC = 72  //col:77
    ZYDIS_ISA_EXT_XSAVEOPT = 73  //col:78
    ZYDIS_ISA_EXT_XSAVES = 74  //col:79
    ZYDIS_ISA_EXT_MAX_VALUE  =  ZYDIS_ISA_EXT_XSAVES  //col:84
    ZYDIS_ISA_EXT_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ISA_EXT_MAX_VALUE)  //col:88
)



type (
EnumIsaExt interface{
    ZYDIS_ISA_EXT_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:89
}









































)

func NewEnumIsaExt() { return & enumIsaExt{} }

func (e *enumIsaExt)    ZYDIS_ISA_EXT_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:89


return true
}












































