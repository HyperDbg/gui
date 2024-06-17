#pragma once
#ifdef HPRDBGCTRL_EXPORTS
#    define HPRDBGCTRL_API __declspec(dllexport)
#else
#    define HPRDBGCTRL_API __declspec(dllimport)
#endif
#define DRIVER_FUNC_INSTALL 0x01
#define DRIVER_FUNC_STOP    0x02
#define DRIVER_FUNC_REMOVE  0x03
BOOLEAN ManageDriver(_In_ LPCTSTR DriverName, _In_ LPCTSTR ServiceName, _In_ UINT16 Function);
BOOLEAN SetupDriverName(const CHAR *                                  DriverName,
                _Inout_updates_bytes_all_(BufferLength) PCHAR DriverLocation,
                ULONG                                         BufferLength);
