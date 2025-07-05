package clang

import "testing"

func Test_switchStruct(t *testing.T) {
	println(switchStruct(`typedef struct SYMBOL
		  {
		      long long unsigned Type;
		      long long unsigned Len;
		      long long unsigned VariableType;
		      long long unsigned Value;
		      
		  } SYMBOL, *PSYMBOL;`))
}
