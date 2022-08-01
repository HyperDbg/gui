package tools


const(
COLOR_DEFAULT =       ZYAN_VT100SGR_FG_DEFAULT //col:46
COLOR_ERROR =         ZYAN_VT100SGR_FG_BRIGHT_RED //col:47
COLOR_HEADER =        ZYAN_VT100SGR_FG_DEFAULT //col:48
COLOR_HEADER_TITLE =  ZYAN_VT100SGR_FG_CYAN //col:49
COLOR_VALUE_LABEL =   ZYAN_VT100SGR_FG_DEFAULT //col:50
COLOR_VALUE_R =       ZYAN_VT100SGR_FG_BRIGHT_RED //col:51
COLOR_VALUE_G =       ZYAN_VT100SGR_FG_BRIGHT_GREEN //col:52
COLOR_VALUE_B =       ZYAN_VT100SGR_FG_CYAN //col:53
CVT100_OUT(sequence) = (g_vt100_stdout ? (sequence) : "") //col:72
CVT100_ERR(sequence) = (g_vt100_stderr ? (sequence) : "") //col:80
PRINT_VALUE_R(name, format, ...) = PrintValueLabel(name); ZYAN_PRINTF("%s" format "%sn", CVT100_OUT(COLOR_VALUE_R), __VA_ARGS__, CVT100_OUT(COLOR_DEFAULT)); //col:187
PRINT_VALUE_G(name, format, ...) = PrintValueLabel(name); ZYAN_PRINTF("%s" format "%sn", CVT100_OUT(COLOR_VALUE_G), __VA_ARGS__, CVT100_OUT(COLOR_DEFAULT)); //col:199
PRINT_VALUE_B(name, format, ...) = PrintValueLabel(name); ZYAN_PRINTF("%s" format "%sn", CVT100_OUT(COLOR_VALUE_B), __VA_ARGS__, CVT100_OUT(COLOR_DEFAULT)); //col:211
)

type (
ZydisInfo interface{
static const char* FormatZyanStatus()(ok bool)//col:146
static void PrintSectionHeader()(ok bool)//col:165
static void PrintValueLabel()(ok bool)//col:177
    PrintValueLabel()(ok bool)//col:359
static void PrintOperands()(ok bool)//col:565
static void PrintFlags()(ok bool)//col:641
static void PrintAVXInfo()(ok bool)//col:740
static void PrintTokenizedInstruction()(ok bool)//col:802
static void PrintDisassembly()(ok bool)//col:855
static void PrintInstruction()(ok bool)//col:1053
int main()(ok bool)//col:1179
}














































)

func NewZydisInfo() { return & zydisInfo{} }

func (z *zydisInfo)static const char* FormatZyanStatus()(ok bool){//col:146













































return true
}















































func (z *zydisInfo)static void PrintSectionHeader()(ok bool){//col:165








return true
}















































func (z *zydisInfo)static void PrintValueLabel()(ok bool){//col:177





return true
}















































func (z *zydisInfo)    PrintValueLabel()(ok bool){//col:359








































































































































return true
}















































func (z *zydisInfo)static void PrintOperands()(ok bool){//col:565





















































































































































































return true
}















































func (z *zydisInfo)static void PrintFlags()(ok bool){//col:641


































































return true
}















































func (z *zydisInfo)static void PrintAVXInfo()(ok bool){//col:740





















































































return true
}















































func (z *zydisInfo)static void PrintTokenizedInstruction()(ok bool){//col:802




















































return true
}















































func (z *zydisInfo)static void PrintDisassembly()(ok bool){//col:855










































return true
}















































func (z *zydisInfo)static void PrintInstruction()(ok bool){//col:1053






















































































































































































return true
}















































func (z *zydisInfo)int main()(ok bool){//col:1179









































































































return true
}

















































