package examples


const(
COLOR_DEFAULT =       ZYAN_VT100SGR_FG_DEFAULT //col:59
COLOR_ERROR =         ZYAN_VT100SGR_FG_BRIGHT_RED //col:60
COLOR_VALUE_R =       ZYAN_VT100SGR_FG_BRIGHT_RED //col:61
COLOR_VALUE_G =       ZYAN_VT100SGR_FG_BRIGHT_GREEN //col:62
COLOR_VALUE_B =       ZYAN_VT100SGR_FG_CYAN //col:63
CVT100_OUT(sequence) = (g_vt100_stdout ? (sequence) : "") //col:82
CVT100_ERR(sequence) = (g_vt100_stderr ? (sequence) : "") //col:90
)

type (
ZydisPerfTest interface{
#if defined()(ok bool)//col:119
static double GetCounter()(ok bool)//col:127
#elif defined()(ok bool)//col:138
static double GetCounter()(ok bool)//col:151
#elif defined()(ok bool)//col:161
static double GetCounter()(ok bool)//col:171
static void AdjustProcessAndThreadPriority()(ok bool)//col:224
static ZyanU64 ProcessBuffer()(ok bool)//col:295
static void TestPerformance()(ok bool)//col:363
static void GenerateTestData()(ok bool)//col:454
int main()(ok bool)//col:586
}






)

func NewZydisPerfTest() { return & zydisPerfTest{} }

func (z *zydisPerfTest)#if defined()(ok bool){//col:119







































return true
}







func (z *zydisPerfTest)static double GetCounter()(ok bool){//col:127






return true
}







func (z *zydisPerfTest)#elif defined()(ok bool){//col:138







return true
}







func (z *zydisPerfTest)static double GetCounter()(ok bool){//col:151









return true
}







func (z *zydisPerfTest)#elif defined()(ok bool){//col:161






return true
}







func (z *zydisPerfTest)static double GetCounter()(ok bool){//col:171







return true
}







func (z *zydisPerfTest)static void AdjustProcessAndThreadPriority()(ok bool){//col:224






































return true
}







func (z *zydisPerfTest)static ZyanU64 ProcessBuffer()(ok bool){//col:295























































return true
}







func (z *zydisPerfTest)static void TestPerformance()(ok bool){//col:363



















































return true
}







func (z *zydisPerfTest)static void GenerateTestData()(ok bool){//col:454






















































































return true
}







func (z *zydisPerfTest)int main()(ok bool){//col:586













































































































return true
}









