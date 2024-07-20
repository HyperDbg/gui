#include <stdio.h>
#include <windows.h>


#define COMMUNICATION_BUFFER_SIZE 256
#define TCP_END_OF_BUFFER_CHARS_COUNT 4

typedef int (*SendMessageWithParamCallback)(const char *Text);

PVOID g_MessageHandler = NULL;

__declspec(dllexport) VOID SetTextMessageCallback(PVOID Handler) {
    g_MessageHandler = Handler;
    printf("call SetTextMessageCallback\n");
}


__declspec(dllexport) VOID ShowMessages(const char *Fmt, ...) {
    char TempMessage[COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT] = {0};
    memset(TempMessage, 0, COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT);
    strcpy(TempMessage, "log callback buf test");
    printf("TempMessage %s\n", TempMessage);
    ((SendMessageWithParamCallback) g_MessageHandler)(TempMessage);
}
