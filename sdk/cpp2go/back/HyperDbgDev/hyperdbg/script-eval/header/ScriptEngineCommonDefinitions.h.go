package header
//back\HyperDbgDev\hyperdbg\script-eval\header\ScriptEngineCommonDefinitions.h.back

const(
SCRIPT_ENGINE_COMMON_DEFINITIONS_H =  //col:3
SYMBOL_GLOBAL_ID_TYPE = 0 //col:27
SYMBOL_LOCAL_ID_TYPE = 1 //col:28
SYMBOL_NUM_TYPE = 2 //col:29
SYMBOL_REGISTER_TYPE = 3 //col:30
SYMBOL_PSEUDO_REG_TYPE = 4 //col:31
SYMBOL_SEMANTIC_RULE_TYPE = 5 //col:32
SYMBOL_TEMP_TYPE = 6 //col:33
SYMBOL_STRING_TYPE = 7 //col:34
SYMBOL_VARIABLE_COUNT_TYPE = 8 //col:35
SYMBOL_MEM_VALID_CHECK_MASK = (1 << 31) //col:36
SYMBOL_INVALID = 9 //col:37
INVALID = -999 //col:38
LALR_ACCEPT = 999 //col:39
FUNC_INC = 0 //col:43
FUNC_DEC = 1 //col:44
FUNC_REFERENCE = 2 //col:45
FUNC_DEREFERENCE = 3 //col:46
FUNC_OR = 4 //col:47
FUNC_XOR = 5 //col:48
FUNC_AND = 6 //col:49
FUNC_ASR = 7 //col:50
FUNC_ASL = 8 //col:51
FUNC_ADD = 9 //col:52
FUNC_SUB = 10 //col:53
FUNC_MUL = 11 //col:54
FUNC_DIV = 12 //col:55
FUNC_MOD = 13 //col:56
FUNC_GT = 14 //col:57
FUNC_LT = 15 //col:58
FUNC_EGT = 16 //col:59
FUNC_ELT = 17 //col:60
FUNC_EQUAL = 18 //col:61
FUNC_NEQ = 19 //col:62
FUNC_START_OF_IF = 20 //col:63
FUNC_JMP = 21 //col:64
FUNC_JZ = 22 //col:65
FUNC_JNZ = 23 //col:66
FUNC_JMP_TO_END_AND_JZCOMPLETED = 24 //col:67
FUNC_END_OF_IF = 25 //col:68
FUNC_START_OF_WHILE = 26 //col:69
FUNC_END_OF_WHILE = 27 //col:70
FUNC_VARGSTART = 28 //col:71
FUNC_MOV = 29 //col:72
FUNC_START_OF_DO_WHILE = 30 //col:73
FUNC_ = 31 //col:74
FUNC_START_OF_DO_WHILE_COMMANDS = 32 //col:75
FUNC_END_OF_DO_WHILE = 33 //col:76
FUNC_START_OF_FOR = 34 //col:77
FUNC_FOR_INC_DEC = 35 //col:78
FUNC_START_OF_FOR_OMMANDS = 36 //col:79
FUNC_END_OF_IF = 37 //col:80
FUNC_IGNORE_LVALUE = 38 //col:81
FUNC_PRINT = 39 //col:82
FUNC_FORMATS = 40 //col:83
FUNC_EVENT_ENABLE = 41 //col:84
FUNC_EVENT_DISABLE = 42 //col:85
FUNC_TEST_STATEMENT = 43 //col:86
FUNC_SPINLOCK_LOCK = 44 //col:87
FUNC_SPINLOCK_UNLOCK = 45 //col:88
FUNC_PRINTF = 46 //col:89
FUNC_PAUSE = 47 //col:90
FUNC_FLUSH = 48 //col:91
FUNC_EVENT_IGNORE = 49 //col:92
FUNC_SPINLOCK_LOCK_CUSTOM_WAIT = 50 //col:93
FUNC_POI = 51 //col:94
FUNC_DB = 52 //col:95
FUNC_DD = 53 //col:96
FUNC_DW = 54 //col:97
FUNC_DQ = 55 //col:98
FUNC_NEG = 56 //col:99
FUNC_HI = 57 //col:100
FUNC_LOW = 58 //col:101
FUNC_NOT = 59 //col:102
FUNC_CHECK_ADDRESS = 60 //col:103
FUNC_STRLEN = 61 //col:104
FUNC_WCSLEN = 62 //col:105
FUNC_INTERLOCKED_INCREMENT = 63 //col:106
FUNC_INTERLOCKED_DECREMENT = 64 //col:107
FUNC_REFERENCE = 65 //col:108
FUNC_PHYSICAL_TO_VIRTUAL = 66 //col:109
FUNC_VIRTUAL_TO_PHYSICAL = 67 //col:110
FUNC_ED = 68 //col:111
FUNC_EB = 69 //col:112
FUNC_EQ = 70 //col:113
FUNC_INTERLOCKED_EXCHANGE = 71 //col:114
FUNC_INTERLOCKED_EXCHANGE_ADD = 72 //col:115
FUNC_INTERLOCKED_COMPARE_EXCHANGE = 73 //col:116
FUNC_POI = 74 //col:117
FUNC_DB = 75 //col:118
FUNC_DD = 76 //col:119
FUNC_DW = 77 //col:120
FUNC_DQ = 78 //col:121
FUNC_NEG = 79 //col:122
FUNC_HI = 80 //col:123
FUNC_LOW = 81 //col:124
FUNC_NOT = 82 //col:125
FUNC_CHECK_ADDRESS = 83 //col:126
FUNC_STRLEN = 84 //col:127
FUNC_WCSLEN = 85 //col:128
FUNC_INTERLOCKED_INCREMENT = 86 //col:129
FUNC_INTERLOCKED_DECREMENT = 87 //col:130
FUNC_REFERENCE = 88 //col:131
FUNC_PHYSICAL_TO_VIRTUAL = 89 //col:132
FUNC_VIRTUAL_TO_PHYSICAL = 90 //col:133
FUNC_ED = 91 //col:134
FUNC_EB = 92 //col:135
FUNC_EQ = 93 //col:136
FUNC_INTERLOCKED_EXCHANGE = 94 //col:137
FUNC_INTERLOCKED_EXCHANGE_ADD = 95 //col:138
FUNC_INTERLOCKED_COMPARE_EXCHANGE = 96 //col:139
PSEUDO_REGISTER_PID = 0 //col:282
PSEUDO_REGISTER_TID = 1 //col:283
PSEUDO_REGISTER_PNAME = 2 //col:284
PSEUDO_REGISTER_CORE = 3 //col:285
PSEUDO_REGISTER_PROC = 4 //col:286
PSEUDO_REGISTER_THREAD = 5 //col:287
PSEUDO_REGISTER_PEB = 6 //col:288
PSEUDO_REGISTER_TEB = 7 //col:289
PSEUDO_REGISTER_IP = 8 //col:290
PSEUDO_REGISTER_BUFFER = 9 //col:291
PSEUDO_REGISTER_CONTEXT = 10 //col:292
PSEUDO_REGISTER_EVENT_TAG = 11 //col:293
PSEUDO_REGISTER_EVENT_ID = 12 //col:294
)

type 	REGISTER_RAX = 0 uint32
const(
	REGISTER_RAX  typedef enum REGS_ENUM { =  0  //col:141
	REGISTER_EAX  typedef enum REGS_ENUM { =  1  //col:142
	REGISTER_AX  typedef enum REGS_ENUM { =  2  //col:143
	REGISTER_AH  typedef enum REGS_ENUM { =  3  //col:144
	REGISTER_AL  typedef enum REGS_ENUM { =  4  //col:145
	REGISTER_RCX  typedef enum REGS_ENUM { =  5  //col:146
	REGISTER_ECX  typedef enum REGS_ENUM { =  6  //col:147
	REGISTER_CX  typedef enum REGS_ENUM { =  7  //col:148
	REGISTER_CH  typedef enum REGS_ENUM { =  8  //col:149
	REGISTER_CL  typedef enum REGS_ENUM { =  9  //col:150
	REGISTER_RDX  typedef enum REGS_ENUM { =  10  //col:151
	REGISTER_EDX  typedef enum REGS_ENUM { =  11  //col:152
	REGISTER_DX  typedef enum REGS_ENUM { =  12  //col:153
	REGISTER_DH  typedef enum REGS_ENUM { =  13  //col:154
	REGISTER_DL  typedef enum REGS_ENUM { =  14  //col:155
	REGISTER_RBX  typedef enum REGS_ENUM { =  15  //col:156
	REGISTER_EBX  typedef enum REGS_ENUM { =  16  //col:157
	REGISTER_BX  typedef enum REGS_ENUM { =  17  //col:158
	REGISTER_BH  typedef enum REGS_ENUM { =  18  //col:159
	REGISTER_BL  typedef enum REGS_ENUM { =  19  //col:160
	REGISTER_RSP  typedef enum REGS_ENUM { =  20  //col:161
	REGISTER_ESP  typedef enum REGS_ENUM { =  21  //col:162
	REGISTER_SP  typedef enum REGS_ENUM { =  22  //col:163
	REGISTER_SPL  typedef enum REGS_ENUM { =  23  //col:164
	REGISTER_RBP  typedef enum REGS_ENUM { =  24  //col:165
	REGISTER_EBP  typedef enum REGS_ENUM { =  25  //col:166
	REGISTER_BP  typedef enum REGS_ENUM { =  26  //col:167
	REGISTER_BPL  typedef enum REGS_ENUM { =  27  //col:168
	REGISTER_RSI  typedef enum REGS_ENUM { =  28  //col:169
	REGISTER_ESI  typedef enum REGS_ENUM { =  29  //col:170
	REGISTER_SI  typedef enum REGS_ENUM { =  30  //col:171
	REGISTER_SIL  typedef enum REGS_ENUM { =  31  //col:172
	REGISTER_RDI  typedef enum REGS_ENUM { =  32  //col:173
	REGISTER_EDI  typedef enum REGS_ENUM { =  33  //col:174
	REGISTER_DI  typedef enum REGS_ENUM { =  34  //col:175
	REGISTER_DIL  typedef enum REGS_ENUM { =  35  //col:176
	REGISTER_R8  typedef enum REGS_ENUM { =  36  //col:177
	REGISTER_R8D  typedef enum REGS_ENUM { =  37  //col:178
	REGISTER_R8W  typedef enum REGS_ENUM { =  38  //col:179
	REGISTER_R8H  typedef enum REGS_ENUM { =  39  //col:180
	REGISTER_R8L  typedef enum REGS_ENUM { =  40  //col:181
	REGISTER_R9  typedef enum REGS_ENUM { =  41  //col:182
	REGISTER_R9D  typedef enum REGS_ENUM { =  42  //col:183
	REGISTER_R9W  typedef enum REGS_ENUM { =  43  //col:184
	REGISTER_R9H  typedef enum REGS_ENUM { =  44  //col:185
	REGISTER_R9L  typedef enum REGS_ENUM { =  45  //col:186
	REGISTER_R10  typedef enum REGS_ENUM { =  46  //col:187
	REGISTER_R10D  typedef enum REGS_ENUM { =  47  //col:188
	REGISTER_R10W  typedef enum REGS_ENUM { =  48  //col:189
	REGISTER_R10H  typedef enum REGS_ENUM { =  49  //col:190
	REGISTER_R10L  typedef enum REGS_ENUM { =  50  //col:191
	REGISTER_R11  typedef enum REGS_ENUM { =  51  //col:192
	REGISTER_R11D  typedef enum REGS_ENUM { =  52  //col:193
	REGISTER_R11W  typedef enum REGS_ENUM { =  53  //col:194
	REGISTER_R11H  typedef enum REGS_ENUM { =  54  //col:195
	REGISTER_R11L  typedef enum REGS_ENUM { =  55  //col:196
	REGISTER_R12  typedef enum REGS_ENUM { =  56  //col:197
	REGISTER_R12D  typedef enum REGS_ENUM { =  57  //col:198
	REGISTER_R12W  typedef enum REGS_ENUM { =  58  //col:199
	REGISTER_R12H  typedef enum REGS_ENUM { =  59  //col:200
	REGISTER_R12L  typedef enum REGS_ENUM { =  60  //col:201
	REGISTER_R13  typedef enum REGS_ENUM { =  61  //col:202
	REGISTER_R13D  typedef enum REGS_ENUM { =  62  //col:203
	REGISTER_R13W  typedef enum REGS_ENUM { =  63  //col:204
	REGISTER_R13H  typedef enum REGS_ENUM { =  64  //col:205
	REGISTER_R13L  typedef enum REGS_ENUM { =  65  //col:206
	REGISTER_R14  typedef enum REGS_ENUM { =  66  //col:207
	REGISTER_R14D  typedef enum REGS_ENUM { =  67  //col:208
	REGISTER_R14W  typedef enum REGS_ENUM { =  68  //col:209
	REGISTER_R14H  typedef enum REGS_ENUM { =  69  //col:210
	REGISTER_R14L  typedef enum REGS_ENUM { =  70  //col:211
	REGISTER_R15  typedef enum REGS_ENUM { =  71  //col:212
	REGISTER_R15D  typedef enum REGS_ENUM { =  72  //col:213
	REGISTER_R15W  typedef enum REGS_ENUM { =  73  //col:214
	REGISTER_R15H  typedef enum REGS_ENUM { =  74  //col:215
	REGISTER_R15L  typedef enum REGS_ENUM { =  75  //col:216
	REGISTER_DS  typedef enum REGS_ENUM { =  76  //col:217
	REGISTER_ES  typedef enum REGS_ENUM { =  77  //col:218
	REGISTER_FS  typedef enum REGS_ENUM { =  78  //col:219
	REGISTER_GS  typedef enum REGS_ENUM { =  79  //col:220
	REGISTER_CS  typedef enum REGS_ENUM { =  80  //col:221
	REGISTER_SS  typedef enum REGS_ENUM { =  81  //col:222
	REGISTER_RFLAGS  typedef enum REGS_ENUM { =  82  //col:223
	REGISTER_EFLAGS  typedef enum REGS_ENUM { =  83  //col:224
	REGISTER_FLAGS  typedef enum REGS_ENUM { =  84  //col:225
	REGISTER_CF  typedef enum REGS_ENUM { =  85  //col:226
	REGISTER_PF  typedef enum REGS_ENUM { =  86  //col:227
	REGISTER_AF  typedef enum REGS_ENUM { =  87  //col:228
	REGISTER_ZF  typedef enum REGS_ENUM { =  88  //col:229
	REGISTER_SF  typedef enum REGS_ENUM { =  89  //col:230
	REGISTER_TF  typedef enum REGS_ENUM { =  90  //col:231
	REGISTER_IF  typedef enum REGS_ENUM { =  91  //col:232
	REGISTER_DF  typedef enum REGS_ENUM { =  92  //col:233
	REGISTER_OF  typedef enum REGS_ENUM { =  93  //col:234
	REGISTER_IOPL  typedef enum REGS_ENUM { =  94  //col:235
	REGISTER_NT  typedef enum REGS_ENUM { =  95  //col:236
	REGISTER_RF  typedef enum REGS_ENUM { =  96  //col:237
	REGISTER_VM  typedef enum REGS_ENUM { =  97  //col:238
	REGISTER_AC  typedef enum REGS_ENUM { =  98  //col:239
	REGISTER_VIF  typedef enum REGS_ENUM { =  99  //col:240
	REGISTER_VIP  typedef enum REGS_ENUM { =  100  //col:241
	REGISTER_ID  typedef enum REGS_ENUM { =  101  //col:242
	REGISTER_RIP  typedef enum REGS_ENUM { =  102  //col:243
	REGISTER_EIP  typedef enum REGS_ENUM { =  103  //col:244
	REGISTER_IP  typedef enum REGS_ENUM { =  104  //col:245
	REGISTER_IDTR  typedef enum REGS_ENUM { =  105  //col:246
	REGISTER_LDTR  typedef enum REGS_ENUM { =  106  //col:247
	REGISTER_GDTR  typedef enum REGS_ENUM { =  107  //col:248
	REGISTER_TR  typedef enum REGS_ENUM { =  108  //col:249
	REGISTER_CR0  typedef enum REGS_ENUM { =  109  //col:250
	REGISTER_CR2  typedef enum REGS_ENUM { =  110  //col:251
	REGISTER_CR3  typedef enum REGS_ENUM { =  111  //col:252
	REGISTER_CR4  typedef enum REGS_ENUM { =  112  //col:253
	REGISTER_CR8  typedef enum REGS_ENUM { =  113  //col:254
	REGISTER_DR0  typedef enum REGS_ENUM { =  114  //col:255
	REGISTER_DR1  typedef enum REGS_ENUM { =  115  //col:256
	REGISTER_DR2  typedef enum REGS_ENUM { =  116  //col:257
	REGISTER_DR3  typedef enum REGS_ENUM { =  117  //col:258
	REGISTER_DR6  typedef enum REGS_ENUM { =  118  //col:259
	REGISTER_DR7  typedef enum REGS_ENUM { =  119  //col:260
)



type (
ScriptEngineCommonDefinitions interface{
#define SYMBOL_MEM_VALID_CHECK_MASK ()(ok bool)//col:262
}
)

func NewScriptEngineCommonDefinitions() { return & scriptEngineCommonDefinitions{} }

func (s *scriptEngineCommonDefinitions)#define SYMBOL_MEM_VALID_CHECK_MASK ()(ok bool){//col:262
/*#define SYMBOL_MEM_VALID_CHECK_MASK (1 << 31)
#define SYMBOL_INVALID 9
#define INVALID -999
#define LALR_ACCEPT 999
#define FUNC_INC 0
#define FUNC_DEC 1
#define FUNC_REFERENCE 2
#define FUNC_DEREFERENCE 3
#define FUNC_OR 4
#define FUNC_XOR 5
#define FUNC_AND 6
#define FUNC_ASR 7
#define FUNC_ASL 8
#define FUNC_ADD 9
#define FUNC_SUB 10
#define FUNC_MUL 11
#define FUNC_DIV 12
#define FUNC_MOD 13
#define FUNC_GT 14
#define FUNC_LT 15
#define FUNC_EGT 16
#define FUNC_ELT 17
#define FUNC_EQUAL 18
#define FUNC_NEQ 19
#define FUNC_START_OF_IF 20
#define FUNC_JMP 21
#define FUNC_JZ 22
#define FUNC_JNZ 23
#define FUNC_JMP_TO_END_AND_JZCOMPLETED 24
#define FUNC_END_OF_IF 25
#define FUNC_START_OF_WHILE 26
#define FUNC_END_OF_WHILE 27
#define FUNC_VARGSTART 28
#define FUNC_MOV 29
#define FUNC_START_OF_DO_WHILE 30
#define FUNC_ 31
#define FUNC_START_OF_DO_WHILE_COMMANDS 32
#define FUNC_END_OF_DO_WHILE 33
#define FUNC_START_OF_FOR 34
#define FUNC_FOR_INC_DEC 35
#define FUNC_START_OF_FOR_OMMANDS 36
#define FUNC_END_OF_IF 37
#define FUNC_IGNORE_LVALUE 38
#define FUNC_PRINT 39
#define FUNC_FORMATS 40
#define FUNC_EVENT_ENABLE 41
#define FUNC_EVENT_DISABLE 42
#define FUNC_TEST_STATEMENT 43
#define FUNC_SPINLOCK_LOCK 44
#define FUNC_SPINLOCK_UNLOCK 45
#define FUNC_PRINTF 46
#define FUNC_PAUSE 47
#define FUNC_FLUSH 48
#define FUNC_EVENT_IGNORE 49
#define FUNC_SPINLOCK_LOCK_CUSTOM_WAIT 50
#define FUNC_POI 51
#define FUNC_DB 52
#define FUNC_DD 53
#define FUNC_DW 54
#define FUNC_DQ 55
#define FUNC_NEG 56
#define FUNC_HI 57
#define FUNC_LOW 58
#define FUNC_NOT 59
#define FUNC_CHECK_ADDRESS 60
#define FUNC_STRLEN 61
#define FUNC_WCSLEN 62
#define FUNC_INTERLOCKED_INCREMENT 63
#define FUNC_INTERLOCKED_DECREMENT 64
#define FUNC_REFERENCE 65
#define FUNC_PHYSICAL_TO_VIRTUAL 66
#define FUNC_VIRTUAL_TO_PHYSICAL 67
#define FUNC_ED 68
#define FUNC_EB 69
#define FUNC_EQ 70
#define FUNC_INTERLOCKED_EXCHANGE 71
#define FUNC_INTERLOCKED_EXCHANGE_ADD 72
#define FUNC_INTERLOCKED_COMPARE_EXCHANGE 73
#define FUNC_POI 74
#define FUNC_DB 75
#define FUNC_DD 76
#define FUNC_DW 77
#define FUNC_DQ 78
#define FUNC_NEG 79
#define FUNC_HI 80
#define FUNC_LOW 81
#define FUNC_NOT 82
#define FUNC_CHECK_ADDRESS 83
#define FUNC_STRLEN 84
#define FUNC_WCSLEN 85
#define FUNC_INTERLOCKED_INCREMENT 86
#define FUNC_INTERLOCKED_DECREMENT 87
#define FUNC_REFERENCE 88
#define FUNC_PHYSICAL_TO_VIRTUAL 89
#define FUNC_VIRTUAL_TO_PHYSICAL 90
#define FUNC_ED 91
#define FUNC_EB 92
#define FUNC_EQ 93
#define FUNC_INTERLOCKED_EXCHANGE 94
#define FUNC_INTERLOCKED_EXCHANGE_ADD 95
#define FUNC_INTERLOCKED_COMPARE_EXCHANGE 96
typedef enum REGS_ENUM {
	REGISTER_RAX = 0,
	REGISTER_EAX = 1,
	REGISTER_AX = 2,
	REGISTER_AH = 3,
	REGISTER_AL = 4,
	REGISTER_RCX = 5,
	REGISTER_ECX = 6,
	REGISTER_CX = 7,
	REGISTER_CH = 8,
	REGISTER_CL = 9,
	REGISTER_RDX = 10,
	REGISTER_EDX = 11,
	REGISTER_DX = 12,
	REGISTER_DH = 13,
	REGISTER_DL = 14,
	REGISTER_RBX = 15,
	REGISTER_EBX = 16,
	REGISTER_BX = 17,
	REGISTER_BH = 18,
	REGISTER_BL = 19,
	REGISTER_RSP = 20,
	REGISTER_ESP = 21,
	REGISTER_SP = 22,
	REGISTER_SPL = 23,
	REGISTER_RBP = 24,
	REGISTER_EBP = 25,
	REGISTER_BP = 26,
	REGISTER_BPL = 27,
	REGISTER_RSI = 28,
	REGISTER_ESI = 29,
	REGISTER_SI = 30,
	REGISTER_SIL = 31,
	REGISTER_RDI = 32,
	REGISTER_EDI = 33,
	REGISTER_DI = 34,
	REGISTER_DIL = 35,
	REGISTER_R8 = 36,
	REGISTER_R8D = 37,
	REGISTER_R8W = 38,
	REGISTER_R8H = 39,
	REGISTER_R8L = 40,
	REGISTER_R9 = 41,
	REGISTER_R9D = 42,
	REGISTER_R9W = 43,
	REGISTER_R9H = 44,
	REGISTER_R9L = 45,
	REGISTER_R10 = 46,
	REGISTER_R10D = 47,
	REGISTER_R10W = 48,
	REGISTER_R10H = 49,
	REGISTER_R10L = 50,
	REGISTER_R11 = 51,
	REGISTER_R11D = 52,
	REGISTER_R11W = 53,
	REGISTER_R11H = 54,
	REGISTER_R11L = 55,
	REGISTER_R12 = 56,
	REGISTER_R12D = 57,
	REGISTER_R12W = 58,
	REGISTER_R12H = 59,
	REGISTER_R12L = 60,
	REGISTER_R13 = 61,
	REGISTER_R13D = 62,
	REGISTER_R13W = 63,
	REGISTER_R13H = 64,
	REGISTER_R13L = 65,
	REGISTER_R14 = 66,
	REGISTER_R14D = 67,
	REGISTER_R14W = 68,
	REGISTER_R14H = 69,
	REGISTER_R14L = 70,
	REGISTER_R15 = 71,
	REGISTER_R15D = 72,
	REGISTER_R15W = 73,
	REGISTER_R15H = 74,
	REGISTER_R15L = 75,
	REGISTER_DS = 76,
	REGISTER_ES = 77,
	REGISTER_FS = 78,
	REGISTER_GS = 79,
	REGISTER_CS = 80,
	REGISTER_SS = 81,
	REGISTER_RFLAGS = 82,
	REGISTER_EFLAGS = 83,
	REGISTER_FLAGS = 84,
	REGISTER_CF = 85,
	REGISTER_PF = 86,
	REGISTER_AF = 87,
	REGISTER_ZF = 88,
	REGISTER_SF = 89,
	REGISTER_TF = 90,
	REGISTER_IF = 91,
	REGISTER_DF = 92,
	REGISTER_OF = 93,
	REGISTER_IOPL = 94,
	REGISTER_NT = 95,
	REGISTER_RF = 96,
	REGISTER_VM = 97,
	REGISTER_AC = 98,
	REGISTER_VIF = 99,
	REGISTER_VIP = 100,
	REGISTER_ID = 101,
	REGISTER_RIP = 102,
	REGISTER_EIP = 103,
	REGISTER_IP = 104,
	REGISTER_IDTR = 105,
	REGISTER_LDTR = 106,
	REGISTER_GDTR = 107,
	REGISTER_TR = 108,
	REGISTER_CR0 = 109,
	REGISTER_CR2 = 110,
	REGISTER_CR3 = 111,
	REGISTER_CR4 = 112,
	REGISTER_CR8 = 113,
	REGISTER_DR0 = 114,
	REGISTER_DR1 = 115,
	REGISTER_DR2 = 116,
	REGISTER_DR3 = 117,
	REGISTER_DR6 = 118,
	REGISTER_DR7 = 119
} REGS_ENUM;*/
return true
}


