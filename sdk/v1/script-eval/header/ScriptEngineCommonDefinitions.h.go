package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-eval\header\ScriptEngineCommonDefinitions.h.back

const (
	REGISTER_RAX    = 0   //col:2
	REGISTER_EAX    = 1   //col:3
	REGISTER_AX     = 2   //col:4
	REGISTER_AH     = 3   //col:5
	REGISTER_AL     = 4   //col:6
	REGISTER_RCX    = 5   //col:7
	REGISTER_ECX    = 6   //col:8
	REGISTER_CX     = 7   //col:9
	REGISTER_CH     = 8   //col:10
	REGISTER_CL     = 9   //col:11
	REGISTER_RDX    = 10  //col:12
	REGISTER_EDX    = 11  //col:13
	REGISTER_DX     = 12  //col:14
	REGISTER_DH     = 13  //col:15
	REGISTER_DL     = 14  //col:16
	REGISTER_RBX    = 15  //col:17
	REGISTER_EBX    = 16  //col:18
	REGISTER_BX     = 17  //col:19
	REGISTER_BH     = 18  //col:20
	REGISTER_BL     = 19  //col:21
	REGISTER_RSP    = 20  //col:22
	REGISTER_ESP    = 21  //col:23
	REGISTER_SP     = 22  //col:24
	REGISTER_SPL    = 23  //col:25
	REGISTER_RBP    = 24  //col:26
	REGISTER_EBP    = 25  //col:27
	REGISTER_BP     = 26  //col:28
	REGISTER_BPL    = 27  //col:29
	REGISTER_RSI    = 28  //col:30
	REGISTER_ESI    = 29  //col:31
	REGISTER_SI     = 30  //col:32
	REGISTER_SIL    = 31  //col:33
	REGISTER_RDI    = 32  //col:34
	REGISTER_EDI    = 33  //col:35
	REGISTER_DI     = 34  //col:36
	REGISTER_DIL    = 35  //col:37
	REGISTER_R8     = 36  //col:38
	REGISTER_R8D    = 37  //col:39
	REGISTER_R8W    = 38  //col:40
	REGISTER_R8H    = 39  //col:41
	REGISTER_R8L    = 40  //col:42
	REGISTER_R9     = 41  //col:43
	REGISTER_R9D    = 42  //col:44
	REGISTER_R9W    = 43  //col:45
	REGISTER_R9H    = 44  //col:46
	REGISTER_R9L    = 45  //col:47
	REGISTER_R10    = 46  //col:48
	REGISTER_R10D   = 47  //col:49
	REGISTER_R10W   = 48  //col:50
	REGISTER_R10H   = 49  //col:51
	REGISTER_R10L   = 50  //col:52
	REGISTER_R11    = 51  //col:53
	REGISTER_R11D   = 52  //col:54
	REGISTER_R11W   = 53  //col:55
	REGISTER_R11H   = 54  //col:56
	REGISTER_R11L   = 55  //col:57
	REGISTER_R12    = 56  //col:58
	REGISTER_R12D   = 57  //col:59
	REGISTER_R12W   = 58  //col:60
	REGISTER_R12H   = 59  //col:61
	REGISTER_R12L   = 60  //col:62
	REGISTER_R13    = 61  //col:63
	REGISTER_R13D   = 62  //col:64
	REGISTER_R13W   = 63  //col:65
	REGISTER_R13H   = 64  //col:66
	REGISTER_R13L   = 65  //col:67
	REGISTER_R14    = 66  //col:68
	REGISTER_R14D   = 67  //col:69
	REGISTER_R14W   = 68  //col:70
	REGISTER_R14H   = 69  //col:71
	REGISTER_R14L   = 70  //col:72
	REGISTER_R15    = 71  //col:73
	REGISTER_R15D   = 72  //col:74
	REGISTER_R15W   = 73  //col:75
	REGISTER_R15H   = 74  //col:76
	REGISTER_R15L   = 75  //col:77
	REGISTER_DS     = 76  //col:78
	REGISTER_ES     = 77  //col:79
	REGISTER_FS     = 78  //col:80
	REGISTER_GS     = 79  //col:81
	REGISTER_CS     = 80  //col:82
	REGISTER_SS     = 81  //col:83
	REGISTER_RFLAGS = 82  //col:84
	REGISTER_EFLAGS = 83  //col:85
	REGISTER_FLAGS  = 84  //col:86
	REGISTER_CF     = 85  //col:87
	REGISTER_PF     = 86  //col:88
	REGISTER_AF     = 87  //col:89
	REGISTER_ZF     = 88  //col:90
	REGISTER_SF     = 89  //col:91
	REGISTER_TF     = 90  //col:92
	REGISTER_IF     = 91  //col:93
	REGISTER_DF     = 92  //col:94
	REGISTER_OF     = 93  //col:95
	REGISTER_IOPL   = 94  //col:96
	REGISTER_NT     = 95  //col:97
	REGISTER_RF     = 96  //col:98
	REGISTER_VM     = 97  //col:99
	REGISTER_AC     = 98  //col:100
	REGISTER_VIF    = 99  //col:101
	REGISTER_VIP    = 100 //col:102
	REGISTER_ID     = 101 //col:103
	REGISTER_RIP    = 102 //col:104
	REGISTER_EIP    = 103 //col:105
	REGISTER_IP     = 104 //col:106
	REGISTER_IDTR   = 105 //col:107
	REGISTER_LDTR   = 106 //col:108
	REGISTER_GDTR   = 107 //col:109
	REGISTER_TR     = 108 //col:110
	REGISTER_CR0    = 109 //col:111
	REGISTER_CR2    = 110 //col:112
	REGISTER_CR3    = 111 //col:113
	REGISTER_CR4    = 112 //col:114
	REGISTER_CR8    = 113 //col:115
	REGISTER_DR0    = 114 //col:116
	REGISTER_DR1    = 115 //col:117
	REGISTER_DR2    = 116 //col:118
	REGISTER_DR3    = 117 //col:119
	REGISTER_DR6    = 118 //col:120
	REGISTER_DR7    = 119 //col:121
)

type SYMBOL struct {
	Type  uint64 //col:5
	Value uint64 //col:6
}

type SYMBOL_BUFFER struct {
	Head    PSYMBOL //col:11
	Pointer uint32  //col:12
	Size    uint32  //col:13
	Message *int8   //col:14
}

type SYMBOL_MAP struct {
	Name *int8  //col:17
	Type uint64 //col:18
}

type ACTION_BUFFER struct {
	Tag                       uint64 //col:22
	CurrentAction             uint64 //col:23
	ImmediatelySendTheResults int8   //col:24
	Context                   uint64 //col:25
}

