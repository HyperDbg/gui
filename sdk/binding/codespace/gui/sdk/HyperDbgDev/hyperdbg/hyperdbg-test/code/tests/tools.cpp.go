package tests

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hyperdbg-test\code\tests\tools.cpp.back

type (
	Tools interface {
		Uint64ToString() (ok bool)  //col:6
		StringReplace() (ok bool)   //col:14
		ConvertToString() (ok bool) //col:19
	}
	tools struct{}
)

func NewTools() Tools { return &tools{} }

func (t *tools) Uint64ToString() (ok bool) { //col:6
	/*Uint64ToString(UINT64 value)
	  {
	      ostringstream Os;
	      Os << setw(16) << setfill('0') << hex << value;
	      return Os.str();
	  }*/
	return true
}

func (t *tools) StringReplace() (ok bool) { //col:14
	/*StringReplace(std::string & str, const std::string & from, const std::string & to)
	  {
	      size_t start_pos = str.find(from);
	      if (start_pos == string::npos)
	          return FALSE;
	      str.replace(start_pos, from.length(), to);
	      return TRUE;
	  }*/
	return true
}

func (t *tools) ConvertToString() (ok bool) { //col:19
	/*ConvertToString(char * Str)
	  {
	      string s(Str);
	      return s;
	  }*/
	return true
}
