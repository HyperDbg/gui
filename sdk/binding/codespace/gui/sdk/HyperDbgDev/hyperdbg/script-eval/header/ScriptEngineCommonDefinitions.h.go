package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-eval\header\ScriptEngineCommonDefinitions.h.back

const (
	SCRIPT_ENGINE_COMMON_DEFINITIONS_H = //col:1
	SYMBOL_GLOBAL_ID_TYPE = 0 //col:2
SYMBOL_LOCAL_ID_TYPE = 1      //col:3
SYMBOL_NUM_TYPE = 2           //col:4
SYMBOL_REGISTER_TYPE = 3   //col:5
SYMBOL_PSEUDO_REG_TYPE = 4 //col:6
SYMBOL_SEMANTIC_RULE_TYPE = 5 //col:7
SYMBOL_TEMP_TYPE = 6          //col:8
SYMBOL_STRING_TYPE = 7        //col:9
SYMBOL_VARIABLE_COUNT_TYPE = 8 //col:10
SYMBOL_MEM_VALID_CHECK_MASK = (1 << 31
) //col:11
SYMBOL_INVALID = 9 //col:12
INVALID = -999     //col:13
LALR_ACCEPT = 999 //col:14
FUNC_INC = 0      //col:15
FUNC_DEC = 1      //col:16
FUNC_REFERENCE = 2   //col:17
FUNC_DEREFERENCE = 3 //col:18
FUNC_OR = 4  //col:19
FUNC_XOR = 5 //col:20
FUNC_AND = 6 //col:21
FUNC_ASR = 7 //col:22
FUNC_ASL = 8 //col:23
FUNC_ADD = 9  //col:24
FUNC_SUB = 10 //col:25
FUNC_MUL = 11 //col:26
FUNC_DIV = 12 //col:27
FUNC_MOD = 13 //col:28
FUNC_GT = 14  //col:29
FUNC_LT = 15  //col:30
FUNC_EGT = 16 //col:31
FUNC_ELT = 17   //col:32
FUNC_EQUAL = 18 //col:33
FUNC_NEQ = 19         //col:34
FUNC_START_OF_IF = 20 //col:35
FUNC_JMP = 21         //col:36
FUNC_JZ = 22  //col:37
FUNC_JNZ = 23 //col:38
FUNC_JMP_TO_END_AND_JZCOMPLETED = 24 //col:39
FUNC_END_OF_IF = 25                  //col:40
FUNC_START_OF_WHILE = 26             //col:41
FUNC_END_OF_WHILE = 27 //col:42
FUNC_VARGSTART = 28    //col:43
FUNC_MOV = 29               //col:44
FUNC_START_OF_DO_WHILE = 30 //col:45
FUNC_ = 31                  //col:46
FUNC_START_OF_DO_WHILE_COMMANDS = 32 //col:47
FUNC_END_OF_DO_WHILE = 33            //col:48
FUNC_START_OF_FOR = 34         //col:49
FUNC_FOR_INC_DEC = 35          //col:50
FUNC_START_OF_FOR_OMMANDS = 36 //col:51
FUNC_END_OF_IF = 37     //col:52
FUNC_IGNORE_LVALUE = 38 //col:53
FUNC_PRINT = 39        //col:54
FUNC_FORMATS = 40      //col:55
FUNC_EVENT_ENABLE = 41 //col:56
FUNC_EVENT_DISABLE = 42  //col:57
FUNC_TEST_STATEMENT = 43 //col:58
FUNC_SPINLOCK_LOCK = 44   //col:59
FUNC_SPINLOCK_UNLOCK = 45 //col:60
FUNC_PRINTF = 46          //col:61
FUNC_PAUSE = 47 //col:62
FUNC_FLUSH = 48 //col:63
FUNC_EVENT_IGNORE = 49              //col:64
FUNC_SPINLOCK_LOCK_CUSTOM_WAIT = 50 //col:65
FUNC_POI = 51                       //col:66
FUNC_DB = 52 //col:67
FUNC_DD = 53 //col:68
FUNC_DW = 54  //col:69
FUNC_DQ = 55  //col:70
FUNC_NEG = 56 //col:71
FUNC_HI = 57  //col:72
FUNC_LOW = 58 //col:73
FUNC_NOT = 59           //col:74
FUNC_CHECK_ADDRESS = 60 //col:75
FUNC_STRLEN = 61        //col:76
FUNC_WCSLEN = 62                //col:77
FUNC_INTERLOCKED_INCREMENT = 63 //col:78
FUNC_INTERLOCKED_DECREMENT = 64 //col:79
FUNC_REFERENCE = 65             //col:80
FUNC_PHYSICAL_TO_VIRTUAL = 66   //col:81
FUNC_VIRTUAL_TO_PHYSICAL = 67 //col:82
FUNC_ED = 68                  //col:83
FUNC_EB = 69                   //col:84
FUNC_EQ = 70                   //col:85
FUNC_INTERLOCKED_EXCHANGE = 71 //col:86
FUNC_INTERLOCKED_EXCHANGE_ADD = 72     //col:87
FUNC_INTERLOCKED_COMPARE_EXCHANGE = 73 //col:88
FUNC_POI = 74 //col:89
FUNC_DB = 75  //col:90
FUNC_DD = 76  //col:91
FUNC_DW = 77 //col:92
FUNC_DQ = 78 //col:93
FUNC_NEG = 79 //col:94
FUNC_HI = 80  //col:95
FUNC_LOW = 81 //col:96
FUNC_NOT = 82           //col:97
FUNC_CHECK_ADDRESS = 83 //col:98
FUNC_STRLEN = 84                //col:99
FUNC_WCSLEN = 85                //col:100
FUNC_INTERLOCKED_INCREMENT = 86 //col:101
FUNC_INTERLOCKED_DECREMENT = 87 //col:102
FUNC_REFERENCE = 88             //col:103
FUNC_PHYSICAL_TO_VIRTUAL = 89 //col:104
FUNC_VIRTUAL_TO_PHYSICAL = 90 //col:105
FUNC_ED = 91                  //col:106
FUNC_EB = 92 //col:107
FUNC_EQ = 93 //col:108
FUNC_INTERLOCKED_EXCHANGE = 94         //col:109
FUNC_INTERLOCKED_EXCHANGE_ADD = 95     //col:110
FUNC_INTERLOCKED_COMPARE_EXCHANGE = 96 //col:111
PSEUDO_REGISTER_PID = 0 //col:112
PSEUDO_REGISTER_TID = 1 //col:113
PSEUDO_REGISTER_PNAME = 2 //col:114
PSEUDO_REGISTER_CORE = 3  //col:115
PSEUDO_REGISTER_PROC = 4  //col:116
PSEUDO_REGISTER_THREAD = 5 //col:117
PSEUDO_REGISTER_PEB = 6    //col:118
PSEUDO_REGISTER_TEB = 7    //col:119
PSEUDO_REGISTER_IP = 8     //col:120
PSEUDO_REGISTER_BUFFER = 9 //col:121
PSEUDO_REGISTER_CONTEXT = 10   //col:122
PSEUDO_REGISTER_EVENT_TAG = 11 //col:123
PSEUDO_REGISTER_EVENT_ID = 12 //col:124
)

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

type typedef struct SYMBOL { struct{
long long //col:2
long long //col:3
}


	type typedef struct SYMBOL_BUFFER {
	struct{
	Head PSYMBOL //col:6
	int unsigned //col:7
	int unsigned //col:8
	Message char* //col:9
	}


	type typedef struct SYMBOL_MAP struct{
	Name char* //col:13
	long long  //col:14
	}


	type typedef struct ACTION_BUFFER {
	struct{
	long long                      //col:17
	long long                      //col:18
	ImmediatelySendTheResults int8 //col:19
	long long //col:20
	}
