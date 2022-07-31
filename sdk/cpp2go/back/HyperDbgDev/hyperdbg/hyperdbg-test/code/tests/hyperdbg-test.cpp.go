package tests
//back\HyperDbgDev\hyperdbg\hyperdbg-test\code\tests\hyperdbg-test.cpp.back

type (
HyperdbgTest interface{
main()(ok bool)//col:157
}
)

func NewHyperdbgTest() { return & hyperdbgTest{} }

func (h *hyperdbgTest)main()(ok bool){//col:157
/*main(int argc, char * argv[])
{
    HANDLE  PipeHandle;
    BOOLEAN SentMessageResult;
    UINT32  ReadBytes;
    char *  Buffer;
    if (argc != 2)
    {
        printf("you should not test functionalities directly, instead use 'test' "
               "command from HyperDbg...\n");
        return 1;
    }
    Buffer = (char *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    strcpy_s(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE, "Hey there, Are you HyperDbg?");
    if (!strcmp(argv[1], "im-hyperdbg"))
    {
        PipeHandle = NamedPipeClientCreatePipe("\\\\.\\Pipe\\HyperDbgTests");
        if (!PipeHandle)
        {
            free(Buffer);
            return 1;
        }
        SentMessageResult =
            NamedPipeClientSendMessage(PipeHandle, Buffer, strlen(Buffer) + 1);
        if (!SentMessageResult)
        {
            free(Buffer);
            return 1;
        }
        ReadBytes = NamedPipeClientReadMessage(PipeHandle, Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
        if (!ReadBytes)
        {
            free(Buffer);
            return 1;
        }
        if (strcmp(Buffer,
                   "Hello, Dear Test Process... Yes, I'm HyperDbg Debugger :)") ==
            0)
        {
            RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
            strcpy_s(
                Buffer,
                TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE,
                "Wow! I miss you... Would you plz send me the kernel information?");
            SentMessageResult =
                NamedPipeClientSendMessage(PipeHandle, Buffer, strlen(Buffer) + 1);
            if (!SentMessageResult)
            {
                free(Buffer);
                return 1;
            }
            RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
            ReadBytes = NamedPipeClientReadMessage(PipeHandle, Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
            if (!ReadBytes)
            {
                free(Buffer);
                return 1;
            }
            TestCreateLookupTable(PipeHandle, (PVOID)Buffer, ReadBytes);
            NamedPipeClientClosePipe(PipeHandle);
            exit(0);
        }
    }
    else
    {
        printf("you should not test functionalities directly, instead use 'test' "
               "command from HyperDbg...\n");
        free(Buffer);
        return 1;
    }
    free(Buffer);
    return 0;
}*/
return true
}



