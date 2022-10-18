package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\LengthDisassemblerEngine.h.back

type (
	LengthDisassemblerEngine interface {
		findByte() (ok bool)   //col:11
		parseModRM() (ok bool) //col:29
	}
	lengthDisassemblerEngine struct{}
)

func NewLengthDisassemblerEngine() LengthDisassemblerEngine { return &lengthDisassemblerEngine{} }

func (l *lengthDisassemblerEngine) findByte() (ok bool) { //col:11
	/*
	   findByte(const UINT8 * arr, const size_t N, const UINT8 x)

	   	{
	   	    for (size_t i = 0; i < N; i++)
	   	    {
	   	        if (arr[i] == x)
	   	        {
	   	            return TRUE;
	   	        }
	   	    };
	   	    return FALSE;
	   	}
	*/
	return true
}

func (l *lengthDisassemblerEngine) parseModRM() (ok bool) { //col:29
	/*
	   parseModRM(UINT8 ** b, const BOOLEAN addressPrefix)

	   	{
	   	    UINT8 modrm = *++*b;
	   	    if (!addressPrefix || **b >= 0x40)
	   	    {
	   	        BOOLEAN hasSIB = FALSE;
	   	        if (**b < 0xC0 && (**b & 0b111) == 0b100 && !addressPrefix)
	   	            hasSIB = TRUE, (*b)++;
	   	        if (modrm >= 0x40 && modrm <= 0x7F)
	   	            (*b)++;
	   	        else if ((modrm <= 0x3F && (modrm & 0b111) == 0b101) || (modrm >= 0x80 && modrm <= 0xBF))
	   	            *b += (addressPrefix) ? 2 : 4;
	   	        else if (hasSIB && (**b & 0b111) == 0b101)
	   	            *b += (modrm & 0b01000000) ? 1 : 4;
	   	    }
	   	    else if (addressPrefix && modrm == 0x26)
	   	        *b += 2;
	   	};
	*/
	return true
}

