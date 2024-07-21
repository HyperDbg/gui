#define VOID void
#define PVOID void*


__declspec(dllexport) VOID SetTextMessageCallback(PVOID Handler);
__declspec(dllexport) VOID ShowMessages(const char *Fmt, ...) ;
