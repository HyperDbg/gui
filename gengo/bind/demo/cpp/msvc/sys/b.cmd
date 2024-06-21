v -cc msvc -kernel -o driver.sys driver.v                        \
-DUNICODE -D_UNICODE                                             \
-DWINVER=0x0600 -D_WIN32_WINNT=0x0600 -DNTDDI_VERSION=0x06000000 \
-I"C:\WinDDK\7600.16385.1\inc\ddk"                               \
-Wl,/entry:DriverEntry -Wl,/subsystem:native -Wl,/nodefaultlib   \
-lntoskrnl.lib -lhal.lib                                         \
pause