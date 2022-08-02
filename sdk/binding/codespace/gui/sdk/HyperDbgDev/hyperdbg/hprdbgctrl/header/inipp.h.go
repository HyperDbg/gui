package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\inipp.h.back

type (
	Inipp interface {
		ltrim() (ok bool)   //col:34
		extract() (ok bool) //col:39
	}
	inipp struct{}
)

func NewInipp() Inipp { return &inipp{} }

func (i *inipp) ltrim() (ok bool) { //col:34
	/*
	   ltrim(std::basic_string<CharT> & s, const std::locale & loc)

	   	{
	   	    s.erase(s.begin(),
	   	            std::find_if(s.begin(), s.end(), [&loc](CharT ch) { return !std::isspace(ch, loc); }));

	   rtrim(std::basic_string<CharT> & s, const std::locale & loc)

	   	{
	   	    s.erase(std::find_if(s.rbegin(), s.rend(), [&loc](CharT ch) { return !std::isspace(ch, loc); }).base(),
	   	            s.end());

	   rtrim2(std::basic_string<CharT> & s, UnaryPredicate pred)

	   	{
	   	    s.erase(std::find_if(s.begin(), s.end(), pred), s.end());

	   replace(std::basic_string<CharT> & str, const std::basic_string<CharT> & from, const std::basic_string<CharT> & to)

	   	{
	   	    auto   changed   = false;
	   	    size_t start_pos = 0;
	   	    while ((start_pos = str.find(from, start_pos)) != std::basic_string<CharT>::npos)
	   	    {
	   	        str.replace(start_pos, from.length(), to);
	   	        start_pos += to.length();

	   extract(const std::basic_string<CharT> & value, T & dst)

	   	{
	   	    CharT                           c;
	   	    std::basic_istringstream<CharT> is {value};
	   	    T                               result;
	   	    if ((is >> std::boolalpha >> result) && !(is >> c))
	   	    {
	   	        dst = result;
	   	        return true;
	   	    }
	   	    else
	   	    {
	   	        return false;
	   	    }
	   	}
	*/
	return true
}

func (i *inipp) extract() (ok bool) { //col:39
	/*
	   extract(const std::basic_string<CharT> & value, std::basic_string<CharT> & dst)

	   	{
	   	    dst = value;
	   	    return true;
	   	}
	*/
	return true
}

