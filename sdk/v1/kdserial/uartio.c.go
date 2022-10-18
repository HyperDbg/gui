package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\uartio.c.back

type (
	Uartio interface {
		WriteRegisterWithIndex8() (ok bool) //col:185
	}
	uartio struct{}
)

func NewUartio() Uartio { return &uartio{} }

func (u *uartio) WriteRegisterWithIndex8() (ok bool) { //col:185
	/*
	   WriteRegisterWithIndex8(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PUCHAR Pointer;
	   	    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_REGISTER_UCHAR(Pointer, Value);

	   ReadRegisterWithIndex8(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PUCHAR Pointer;
	   	    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
	   	    return READ_REGISTER_UCHAR(Pointer);

	   WriteRegisterWithIndex16(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PUSHORT Pointer;
	   	    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_REGISTER_USHORT(Pointer, Value);

	   ReadRegisterWithIndex16(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PUSHORT Pointer;
	   	    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
	   	    return (UCHAR)READ_REGISTER_USHORT(Pointer);

	   WriteRegisterWithIndex32(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PULONG Pointer;
	   	    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_REGISTER_ULONG(Pointer, Value);

	   ReadRegisterWithIndex32(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PULONG Pointer;
	   	    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
	   	    return (UCHAR)READ_REGISTER_ULONG(Pointer);

	   WriteRegisterWithIndex64(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PULONG64 Pointer;
	   	    Pointer = (PULONG64)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_REGISTER_ULONG64(Pointer, Value);

	   ReadRegisterWithIndex64(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PULONG64 Pointer;
	   	    Pointer = (PULONG64)(Port->Address + Index * Port->ByteWidth);
	   	    return (UCHAR)READ_REGISTER_ULONG64(Pointer);

	   WritePortWithIndex8(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PUCHAR Pointer;
	   	    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_PORT_UCHAR(Pointer, Value);

	   ReadPortWithIndex8(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PUCHAR Pointer;
	   	    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
	   	    return (UCHAR)READ_PORT_UCHAR(Pointer);

	   WritePortWithIndex16(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PUSHORT Pointer;
	   	    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_PORT_USHORT(Pointer, Value);

	   ReadPortWithIndex16(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PUSHORT Pointer;
	   	    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
	   	    return (UCHAR)READ_PORT_USHORT(Pointer);

	   WritePortWithIndex32(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index,
	   	const UCHAR  Value)

	   	{
	   	    PULONG Pointer;
	   	    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
	   	    WRITE_PORT_ULONG(Pointer, Value);

	   ReadPortWithIndex32(

	   	_In_ PCPPORT Port,
	   	const UCHAR  Index)

	   	{
	   	    PULONG Pointer;
	   	    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
	   	    return (UCHAR)READ_PORT_ULONG(Pointer);

	   UartpSetAccess(

	   	_Inout_ PCPPORT Port,
	   	const BOOLEAN   MemoryMapped,
	   	const UCHAR     AccessSize,
	   	const UCHAR     BitWidth)

	   	{
	   	    UCHAR                             MinRegisterWidth;
	   	    BOOLEAN                           PowerOfTwo;
	   	    UART_HARDWARE_READ_INDEXED_UCHAR  ReadFunction;
	   	    UART_HARDWARE_WRITE_INDEXED_UCHAR WriteFunction;
	   	    MinRegisterWidth = 8;
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        switch ((ACPI_GENERIC_ACCESS_SIZE)AccessSize)
	   	        {

	   #if defined(_X86_) || defined(_AMD64_)

	   	case AcpiGenericAccessSizeLegacy:
	   	    __fallthrough;
	   	case AcpiGenericAccessSizeByte:
	   	    WriteFunction = WritePortWithIndex8;
	   	    ReadFunction  = ReadPortWithIndex8;
	   	    break;
	   	case AcpiGenericAccessSizeWord:
	   	    WriteFunction    = WritePortWithIndex16;
	   	    ReadFunction     = ReadPortWithIndex16;
	   	    MinRegisterWidth = 16;
	   	    break;
	   	case AcpiGenericAccessSizeDWord:
	   	    WriteFunction    = WritePortWithIndex32;
	   	    ReadFunction     = ReadPortWithIndex32;
	   	    MinRegisterWidth = 32;
	   	    break;

	   #endif

	   	    default:
	   	        return FALSE;
	   	    }
	   	}
	   	else
	   	{
	   	    switch ((ACPI_GENERIC_ACCESS_SIZE)AccessSize)
	   	    {
	   	    case AcpiGenericAccessSizeLegacy:
	   	        __fallthrough;
	   	    case AcpiGenericAccessSizeByte:
	   	        WriteFunction = WriteRegisterWithIndex8;
	   	        ReadFunction  = ReadRegisterWithIndex8;
	   	        break;
	   	    case AcpiGenericAccessSizeWord:
	   	        WriteFunction    = WriteRegisterWithIndex16;
	   	        ReadFunction     = ReadRegisterWithIndex16;
	   	        MinRegisterWidth = 16;
	   	        break;
	   	    case AcpiGenericAccessSizeDWord:
	   	        WriteFunction    = WriteRegisterWithIndex32;
	   	        ReadFunction     = ReadRegisterWithIndex32;
	   	        MinRegisterWidth = 32;
	   	        break;

	   #if defined(_WIN64)

	   	case AcpiGenericAccessSizeQWord:
	   	    WriteFunction    = WriteRegisterWithIndex64;
	   	    ReadFunction     = ReadRegisterWithIndex64;
	   	    MinRegisterWidth = 64;
	   	    break;

	   #endif

	   	        default:
	   	            return FALSE;
	   	        }
	   	    }
	   	    PowerOfTwo = ((BitWidth & (BitWidth - 1)) == 0);
	   	    if ((PowerOfTwo == FALSE) ||
	   	        (BitWidth < MinRegisterWidth) ||
	   	        (BitWidth > MAX_REGISTER_WIDTH))
	   	    {
	   	        return FALSE;
	   	    }
	   	    Port->ByteWidth = BitWidth / 8;
	   	    Port->Write     = WriteFunction;
	   	    Port->Read      = ReadFunction;
	   	    return TRUE;
	   	}
	*/
	return true
}

