package Source

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\main.cpp.back

type (
	Main interface {
		int_main_impl() (ok bool) //col:5
		int_main() (ok bool)      //col:9
	}
	main struct{}
)

func NewMain() Main { return &main{} }

func (m *main) int_main_impl() (ok bool) { //col:5
	/*
	   int main_impl(int argc, char** argv)

	   	{
	   		PDBExtractor Instance;
	   		return Instance.Run(argc, argv);
	   	}
	*/
	return true
}

func (m *main) int_main() (ok bool) { //col:9
	/*
	   int main(int argc, char** argv)

	   	{
	   		return main_impl(argc, argv);
	   	}
	*/
	return true
}

