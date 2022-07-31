package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\main.cpp.back

type (
Main interface{
#pragma comment()(ok bool)//col:9
int main()(ok bool)//col:14
}
)

func NewMain() { return & main{} }

func (m *main)#pragma comment()(ok bool){//col:9
/*#pragma comment(lib, "dbghelp.lib")
int main_impl(int argc, char** argv)
{
	PDBExtractor Instance;
	return Instance.Run(argc, argv);
}*/
return true
}

func (m *main)int main()(ok bool){//col:14
/*int main(int argc, char** argv)
{
	return main_impl(argc, argv);
}*/
return true
}



