#include <windows.h>
#include <stdio.h>

int hyperdbg_show_messages(const char *Text) {
    printf("Test in the handler | ");
    printf("%s\n", Text);
    return 0;
}

typedef int (*Pfnhyperdbg_u_set_text_message_callback)(const char *);

typedef void  (*Pfnhyperdbg_u_interpreter)(const char *);

int main() {
    auto hModule = LoadLibraryW(L"libhyperdbg.dll");
    if (hModule == NULL) {
        printf("Failed to load libhyperdbg.dll\n");
        return 1;
    }
    auto hyperdbg_u_set_text_message_callback = (Pfnhyperdbg_u_set_text_message_callback) GetProcAddress(hModule,
                                                                                                         "hyperdbg_u_set_text_message_callback");
    auto hyperdbg_u_interpreter = (Pfnhyperdbg_u_interpreter) GetProcAddress(hModule, "hyperdbg_u_interpreter");
    if (hyperdbg_u_set_text_message_callback == NULL || hyperdbg_u_interpreter == NULL) {
        printf("Failed to get function address\n");
        FreeLibrary(hModule);
        return 1;
    }
    hyperdbg_u_set_text_message_callback(reinterpret_cast<const char *>(hyperdbg_show_messages));
    hyperdbg_u_interpreter("help !monitor");
    FreeLibrary(hModule);
    return 0;
}
