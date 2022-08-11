package communication

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\communication\SerialConnection.c.back

type (
	SerialConnection interface {
		SerialConnectionTest() (ok bool)       //col:33
		SerialConnectionRecvBuffer() (ok bool) //col:56
		SerialConnectionSend() (ok bool)       //col:116
		SerialConnectionCheckPort() (ok bool)  //col:125
	}
	serialConnection struct{}
)

func NewSerialConnection() SerialConnection { return &serialConnection{} }

func (s *serialConnection) SerialConnectionTest() (ok bool) { //col:33
	/*
	   SerialConnectionTest()

	   	{
	   	    for (size_t i = 0; i < 100; i++)
	   	    {
	   	        KdHyperDbgTest((UINT16)i);

	   SerialConnectionSendEndOfBuffer()

	   	{
	   	    KdHyperDbgSendByte(SERIAL_END_OF_BUFFER_CHAR_1, TRUE);
	   	    KdHyperDbgSendByte(SERIAL_END_OF_BUFFER_CHAR_2, TRUE);
	   	    KdHyperDbgSendByte(SERIAL_END_OF_BUFFER_CHAR_3, TRUE);
	   	    KdHyperDbgSendByte(SERIAL_END_OF_BUFFER_CHAR_4, TRUE);

	   SerialConnectionCheckForTheEndOfTheBuffer(PUINT32 CurrentLoopIndex, BYTE * Buffer)

	   	{
	   	    UINT32 ActualBufferLength;
	   	    ActualBufferLength = *CurrentLoopIndex;
	   	    if (*CurrentLoopIndex <= 3)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (Buffer[ActualBufferLength] == SERIAL_END_OF_BUFFER_CHAR_4 &&
	   	        Buffer[ActualBufferLength - 1] == SERIAL_END_OF_BUFFER_CHAR_3 &&
	   	        Buffer[ActualBufferLength - 2] == SERIAL_END_OF_BUFFER_CHAR_2 &&
	   	        Buffer[ActualBufferLength - 3] == SERIAL_END_OF_BUFFER_CHAR_1)
	   	    {
	   	        Buffer[ActualBufferLength - 3] = NULL;
	   	        Buffer[ActualBufferLength - 2] = NULL;
	   	        Buffer[ActualBufferLength - 1] = NULL;
	   	        Buffer[ActualBufferLength]     = NULL;
	   	        *CurrentLoopIndex = ActualBufferLength - 3;
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (s *serialConnection) SerialConnectionRecvBuffer() (ok bool) { //col:56
	/*
	   SerialConnectionRecvBuffer(CHAR *   BufferToSave,

	   	UINT32 * LengthReceived)

	   	{
	   	    UINT32 Loop = 0;
	   	    while (TRUE)
	   	    {
	   	        UCHAR RecvChar = NULL;
	   	        if (!KdHyperDbgRecvByte(&RecvChar))
	   	        {
	   	            continue;
	   	        }
	   	        if (!(MaxSerialPacketSize > Loop))
	   	        {
	   	            LogError("Err, a buffer received in debuggee which exceeds the buffer limitation");
	   	        if (SerialConnectionCheckForTheEndOfTheBuffer(&Loop, (BYTE *)BufferToSave))
	   	        {
	   	            break;
	   	        }
	   	        Loop++;
	   	    }
	   	    *LengthReceived = Loop;
	   	    return TRUE;
	   	}
	*/
	return true
}

func (s *serialConnection) SerialConnectionSend() (ok bool) { //col:116
	/*
	   SerialConnectionSend(CHAR * Buffer, UINT32 Length)

	   	{
	   	    if (Length + SERIAL_END_OF_BUFFER_CHARS_COUNT > MaxSerialPacketSize)
	   	    {
	   	        LogError("Err, buffer is above the maximum buffer size that can be sent to debuggee (%d > %d)",
	   	                 Length + SERIAL_END_OF_BUFFER_CHARS_COUNT,
	   	                 MaxSerialPacketSize);
	   	    for (size_t i = 0; i < Length; i++)
	   	    {
	   	        KdHyperDbgSendByte(Buffer[i], TRUE);
	   	    SerialConnectionSendEndOfBuffer();

	   SerialConnectionSendTwoBuffers(CHAR * Buffer1, UINT32 Length1, CHAR * Buffer2, UINT32 Length2)

	   	{
	   	    if ((Length1 + Length2 + SERIAL_END_OF_BUFFER_CHARS_COUNT) > MaxSerialPacketSize)
	   	    {
	   	        LogError("Err, buffer is above the maximum buffer size that can be sent to debuggee (%d > %d)",
	   	                 Length1 + Length2 + SERIAL_END_OF_BUFFER_CHARS_COUNT,
	   	                 MaxSerialPacketSize);
	   	    for (size_t i = 0; i < Length1; i++)
	   	    {
	   	        KdHyperDbgSendByte(Buffer1[i], TRUE);
	   	    for (size_t i = 0; i < Length2; i++)
	   	    {
	   	        KdHyperDbgSendByte(Buffer2[i], TRUE);
	   	    SerialConnectionSendEndOfBuffer();

	   SerialConnectionSendThreeBuffers(CHAR * Buffer1,

	   	UINT32 Length1,
	   	CHAR * Buffer2,
	   	UINT32 Length2,
	   	CHAR * Buffer3,
	   	UINT32 Length3)

	   	{
	   	    if ((Length1 + Length2 + Length3 + SERIAL_END_OF_BUFFER_CHARS_COUNT) > MaxSerialPacketSize)
	   	    {
	   	        LogError("Err, buffer is above the maximum buffer size that can be sent to debuggee (%d > %d)",
	   	                 Length1 + Length2 + Length3 + SERIAL_END_OF_BUFFER_CHARS_COUNT,
	   	                 MaxSerialPacketSize);
	   	    for (size_t i = 0; i < Length1; i++)
	   	    {
	   	        KdHyperDbgSendByte(Buffer1[i], TRUE);
	   	    for (size_t i = 0; i < Length2; i++)
	   	    {
	   	        KdHyperDbgSendByte(Buffer2[i], TRUE);
	   	    for (size_t i = 0; i < Length3; i++)
	   	    {
	   	        KdHyperDbgSendByte(Buffer3[i], TRUE);
	   	    SerialConnectionSendEndOfBuffer();

	   SerialConnectionCheckBaudrate(DWORD Baudrate)

	   	{
	   	    if (Baudrate == CBR_110 || Baudrate == CBR_300 || Baudrate == CBR_600 ||
	   	        Baudrate == CBR_1200 || Baudrate == CBR_2400 || Baudrate == CBR_4800 ||
	   	        Baudrate == CBR_9600 || Baudrate == CBR_14400 || Baudrate == CBR_19200 ||
	   	        Baudrate == CBR_38400 || Baudrate == CBR_56000 || Baudrate == CBR_57600 ||
	   	        Baudrate == CBR_115200 || Baudrate == CBR_128000 ||
	   	        Baudrate == CBR_256000)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (s *serialConnection) SerialConnectionCheckPort() (ok bool) { //col:125
	/*
	   SerialConnectionCheckPort(UINT32 SerialPort)

	   	{
	   	    if (SerialPort == COM1_PORT || SerialPort == COM2_PORT || SerialPort == COM3_PORT ||
	   	        SerialPort == COM4_PORT)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

