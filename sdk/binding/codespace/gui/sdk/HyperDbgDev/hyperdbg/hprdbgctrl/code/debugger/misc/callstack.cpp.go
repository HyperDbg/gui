package misc

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\misc\callstack.cpp.back

type (
	Callstack interface {
		CallstackReturnAddressToCallingAddress() (ok bool) //col:74
	}
	callstack struct{}
)

func NewCallstack() Callstack { return &callstack{} }

func (c *callstack) CallstackReturnAddressToCallingAddress() (ok bool) { //col:74
	/*
	   CallstackReturnAddressToCallingAddress(UCHAR * ReturnAddress, PUINT32 IndexOfCallFromReturnAddress)

	   	{
	   	    if (ReturnAddress[-7] == 0x9A)
	   	    {
	   	        *IndexOfCallFromReturnAddress = 7;
	   	        return TRUE;
	   	    }
	   	    else if (ReturnAddress[-5] == 0xE8)
	   	    {
	   	        *IndexOfCallFromReturnAddress = 5;
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        const unsigned char RmMask = 0xF8;
	   	        if (ReturnAddress[-7] == 0xFF &&
	   	            (ReturnAddress[-6] == 0x94 || ReturnAddress[-6] == 0x9C))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 7;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-6] == 0xFF &&
	   	                 ((ReturnAddress[-5] & RmMask) == 0x90 || (ReturnAddress[-5] & RmMask) == 0x98) &&
	   	                 (ReturnAddress[-5] != 0x94 && ReturnAddress[-5] != 0x9C))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 6;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-6] == 0xFF &&
	   	                 (ReturnAddress[-5] == 0x15 || ReturnAddress[-5] == 0x1D))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 6;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-4] == 0xFF &&
	   	                 (ReturnAddress[-3] == 0x54 || ReturnAddress[-3] == 0x5C))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 4;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-3] == 0xFF &&
	   	                 ((ReturnAddress[-2] & RmMask) == 0x50 || (ReturnAddress[-2] & RmMask) == 0x58) &&
	   	                 (ReturnAddress[-2] != 0x54 && ReturnAddress[-2] != 0x5C))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 3;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-3] == 0xFF &&
	   	                 (ReturnAddress[-2] == 0x14 || ReturnAddress[-2] == 0x1C))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 3;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-2] == 0xFF &&
	   	                 ((ReturnAddress[-1] & RmMask) == 0xD0 || (ReturnAddress[-1] & RmMask) == 0xD8))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 2;
	   	            return TRUE;
	   	        }
	   	        else if (ReturnAddress[-2] == 0xFF &&
	   	                 ((ReturnAddress[-1] & RmMask) == 0x10 || (ReturnAddress[-1] & RmMask) == 0x18) &&
	   	                 (ReturnAddress[-1] != 0x14 && ReturnAddress[-1] != 0x15 &&
	   	                  ReturnAddress[-1] != 0x1C && ReturnAddress[-1] != 0x1D))
	   	        {
	   	            *IndexOfCallFromReturnAddress = 2;
	   	            return TRUE;
	   	        }
	   	        else
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

