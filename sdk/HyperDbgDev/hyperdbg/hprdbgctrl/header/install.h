#pragma once
#ifdef HPRDBGCTRL_EXPORTS
#    define HPRDBGCTRL_API __declspec(dllexport)
#else
#    define HPRDBGCTRL_API __declspec(dllimport)
#endif
class HPRDBGCTRL_API Chprdbgctrl {
public:
    Chprdbgctrl(void);
};
extern HPRDBGCTRL_API int nhprdbgctrl;
HPRDBGCTRL_API int
fnhprdbgctrl(void);
#define DRIVER_FUNC_INSTALL 0x01
#define DRIVER_FUNC_STOP    0x02
#define DRIVER_FUNC_REMOVE  0x03
BOOLEAN
ManageDriver(_In_ LPCTSTR DriverName, _In_ LPCTSTR ServiceName, _In_ UINT16 Function);
BOOLEAN
SetupDriverName(_Inout_updates_bytes_all_(BufferLength) PCHAR DriverLocation,
                _In_ ULONG                                    BufferLength);
