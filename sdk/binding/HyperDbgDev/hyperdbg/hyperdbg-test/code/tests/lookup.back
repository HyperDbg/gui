










#include "pch.h"







DWORD WINAPI
TestListenerForResultsThread(void * data)
{
    HANDLE OutputPipeHandle;
    UINT32 ReadBytes;
    char * BufferToRead;
    char   TestPipe[] = "\\\\.\\Pipe\\HyperDbgOutputTest";

    
    
    
    OutputPipeHandle = NamedPipeServerCreatePipe(TestPipe,
                                                 TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE,
                                                 TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!OutputPipeHandle)
    {
        
        
        
        return FALSE;
    }

    BufferToRead = (char *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);

    
    
    
    if (!NamedPipeServerWaitForClientConntection(OutputPipeHandle))
    {
        
        
        
        free(BufferToRead);
        return FALSE;
    }

ReadAgain:
    ReadBytes =
        NamedPipeServerReadClientMessage(OutputPipeHandle, BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);

    if (!ReadBytes)
    {
        
        
        

        free(BufferToRead);
        return FALSE;
    }

    
    
    
    printf("Test from debugger : %s\n", BufferToRead);

    
    
    
    goto ReadAgain;

    return TRUE;
}










VOID
TestCreateLookupTable(HANDLE PipeHandle, PVOID KernelInformation, UINT32 KernelInformationSize)
{
    BOOLEAN                                           SentMessageResult;
    PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION        KernelInfoArray;
    char                                              SuccessMessage[]    = "success";
    char                                              KernelTestMessage[] = "perform-kernel-test";
    vector<DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION> LookupTable;
    vector<string>                                    TestCases;
    int                                               IndexOfTestCasesVector = 0;
    DWORD                                             ThreadId;

    printf("start testing event commands...\n");

    KernelInfoArray = (PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION)KernelInformation;

    
    
    
    for (size_t i = 0; i < KernelInformationSize / sizeof(DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION); i++)
    {
        LookupTable.push_back(KernelInfoArray[i]);
    }

    
    
    

    

    
    
    
    HANDLE Thread = CreateThread(NULL, 0, TestListenerForResultsThread, NULL, 0, &ThreadId);

    
    
    
    if (!TestCase(TestCases))
    {
        
        
        
        return;
    }

    
    
    
    for (auto & CurrentCase : TestCases)
    {
        
        
        
        
        smatch Match;
        regex  r("\\[(.*?)\\]");
        string Subject = std::move(CurrentCase);

        int i = 1;
        while (regex_search(Subject, Match, r))
        {
            for (auto & item : LookupTable)
            {
                string Temp = ConvertToString(item.Tag);
                Temp        = "[" + Temp + "]";

                if (!Temp.compare(Match.str(0)))
                {
                    StringReplace(TestCases[IndexOfTestCasesVector], Temp, Uint64ToString(item.Value));
                }
            }

            
            
            
            Subject = Match.suffix().str();
            i++;
        }
        IndexOfTestCasesVector++;
    }

    _getch();

    
    
    
    for (auto NewCase : TestCases)
    {
        printf("new cases : %s\n", NewCase.c_str());

        string OutputCommand = "cmd:" + NewCase;

        SentMessageResult = NamedPipeClientSendMessage(PipeHandle, (char *)OutputCommand.c_str(), OutputCommand.length() + 1);
        if (!SentMessageResult)
        {
            
            
            
            return;
        }
    }

    _getch();

    
    
    
    SentMessageResult = NamedPipeClientSendMessage(PipeHandle, (char *)KernelTestMessage, sizeof(KernelTestMessage));

    if (!SentMessageResult)
    {
        
        
        
        return;
    }

    _getch();

    
    
    
    SentMessageResult = NamedPipeClientSendMessage(PipeHandle, (char *)SuccessMessage, sizeof(SuccessMessage));

    if (!SentMessageResult)
    {
        
        
        
        return;
    }
}
