package header
//back\HyperDbgDev\hyperdbg\hyperdbg-test\header\routines.h.back

type (
Routines interface{
TestCase()(ok bool)//col:26
}
)

func NewRoutines() { return & routines{} }

func (r *routines)TestCase()(ok bool){//col:26
/*TestCase(std::vector<std::string> & TestCase);
extern "C" {
extern void inline AsmTest();
}*/
return true
}



