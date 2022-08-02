package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\IoHandler.c.back

type (
	IoHandler interface {
		IoHandleIoVmExits() (ok bool) //col:73
	}
	ioHandler struct{}
)

func NewIoHandler() IoHandler { return &ioHandler{} }

func (i *ioHandler) IoHandleIoVmExits() (ok bool) { //col:73
	/*
	   IoHandleIoVmExits(PGUEST_REGS GuestRegs, VMX_EXIT_QUALIFICATION_IO_INSTRUCTION IoQualification, RFLAGS Flags)

	   	{
	   	    UINT16 Port  = 0;
	   	    UINT32 Count = 0;
	   	    UINT32 Size  = 0;
	   	    UINT64 GpReg = 0;
	   	    union
	   	    {
	   	        unsigned char *  AsBytePtr;
	   	        unsigned short * AsWordPtr;
	   	        unsigned long *  AsDwordPtr;
	   	        void * AsPtr;
	   	        UINT64 AsUint64;
	   	    } PortValue;
	   	    if (IoQualification.StringInstruction)
	   	    {
	   	        PortValue.AsPtr = IoQualification.DirectionOfAccess == AccessIn ? GuestRegs->rdi : GuestRegs->rsi;
	   	    }
	   	    else
	   	    {
	   	        PortValue.AsPtr = &GuestRegs->rax;
	   	    }
	   	    Port = IoQualification.PortNumber;
	   	    Count = IoQualification.RepPrefixed ? (GuestRegs->rcx & 0xffffffff) : 1;
	   	    Size = IoQualification.SizeOfAccess + 1;
	   	    switch (IoQualification.DirectionOfAccess)
	   	    {
	   	    case AccessIn:
	   	        if (IoQualification.StringInstruction)
	   	        {
	   	            switch (Size)
	   	            {
	   	            case 1:
	   	                IoInByteString(Port, PortValue.AsBytePtr, Count);
	   	                IoInWordString(Port, PortValue.AsWordPtr, Count);
	   	                IoInDwordString(Port, PortValue.AsDwordPtr, Count);
	   	            switch (Size)
	   	            {
	   	            case 1:
	   	                *PortValue.AsBytePtr = IoInByte(Port);
	   	                *PortValue.AsWordPtr = IoInWord(Port);
	   	                *PortValue.AsDwordPtr = IoInDword(Port);
	   	        if (IoQualification.StringInstruction)
	   	        {
	   	            switch (Size)
	   	            {
	   	            case 1:
	   	                IoOutByteString(Port, PortValue.AsBytePtr, Count);
	   	                IoOutWordString(Port, PortValue.AsWordPtr, Count);
	   	                IoOutDwordString(Port, PortValue.AsDwordPtr, Count);
	   	            switch (Size)
	   	            {
	   	            case 1:
	   	                IoOutByte(Port, *PortValue.AsBytePtr);
	   	                IoOutWord(Port, *PortValue.AsWordPtr);
	   	                IoOutDword(Port, *PortValue.AsDwordPtr);
	   	    if (IoQualification.StringInstruction)
	   	    {
	   	        UINT64 GpReg = IoQualification.DirectionOfAccess == AccessIn ? GuestRegs->rdi : GuestRegs->rsi;
	   	        if (Flags.DirectionFlag)
	   	        {
	   	            GpReg -= Count * Size;
	   	        }
	   	        else
	   	        {
	   	            GpReg += Count * Size;
	   	        }
	   	        if (IoQualification.RepPrefixed)
	   	        {
	   	            GuestRegs->rcx = 0;
	   	        }
	   	    }
	   	}
	*/
	return true
}

