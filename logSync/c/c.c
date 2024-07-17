#include <stdio.h>
#include <stdlib.h>
#include "../mydll/main.h"

int main() {
    const char* message = "Hello, World!";
    WriteStringToFile(message);
    return 0;
}
