package transparency

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\transparency\Transparency.c.back

type (
	Transparency interface {
		TransparentGetRand() (ok bool) //col:14
		TransparentLog() (ok bool)     //col:25
		TransparentSqrt() (ok bool)    //col:45
	}
	transparency struct{}
)

func NewTransparency() Transparency { return &transparency{} }

func (t *transparency) TransparentGetRand() (ok bool) { //col:14
	/*
	   TransparentGetRand()

	   	{
	   	    UINT64 Tsc;
	   	    UINT32 Rand;
	   	    Tsc  = __rdtsc();

	   TransparentPow(int x, int p)

	   	{
	   	    int Res = 1;
	   	    for (int i = 0; i < p; i++)
	   	    {
	   	        Res = Res * x;
	   	    }
	   	    return Res;
	   	}
	*/
	return true
}

func (t *transparency) TransparentLog() (ok bool) { //col:25
	/*
	   TransparentLog(int x)

	   	{
	   	    int n     = x;
	   	    int Digit = 0;
	   	    while (n >= 100)
	   	    {
	   	        n = n / 10;
	   	        Digit++;
	   	    }
	   	    return TransparentTableLog[n] / 100 + (Digit * 23) / 10;
	   	}
	*/
	return true
}

func (t *transparency) TransparentSqrt() (ok bool) { //col:45
	/*
	   TransparentSqrt(int x)

	   	{
	   	    int Res = 0;
	   	    int Bit;
	   	    Bit = 1 << 30;
	   	    while (Bit > x)
	   	        Bit >>= 2;
	   	    while (Bit != 0)
	   	    {
	   	        if (x >= Res + Bit)
	   	        {
	   	            x -= Res + Bit;
	   	            Res = (Res >> 1) + Bit;
	   	        }
	   	        else
	   	            Res >>= 1;
	   	        Bit >>= 2;
	   	    }
	   	    return Res;
	   	}
	*/
	return true
}

