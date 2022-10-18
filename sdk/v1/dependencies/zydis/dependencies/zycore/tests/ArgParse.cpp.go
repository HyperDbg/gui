package tests

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\tests\ArgParse.cpp.back

type (
	ArgParse interface {
		auto_cvt_string_view() (ok bool) //col:8
	}
	argParse struct{}
)

func NewArgParse() ArgParse { return &argParse{} }

func (a *argParse) auto_cvt_string_view() (ok bool) { //col:8
	/*
	   auto cvt_string_view(const ZyanStringView *sv)

	   	{
	   	    const char* buf;
	   	    if (ZYAN_FAILED(ZyanStringViewGetData(sv, &buf))) throw std::exception{};
	   	    ZyanUSize len;
	   	    if (ZYAN_FAILED(ZyanStringViewGetSize(sv, &len))) throw std::exception{};
	   	    return std::string_view{buf, len};
	   	}
	*/
	return true
}

