package driver-loader
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\driver-loader\install.cpp.back

type (
Install interface{
InstallDriver()(ok bool)//col:52
ManageDriver()(ok bool)//col:99
RemoveDriver()(ok bool)//col:124
StartDriver()(ok bool)//col:174
StopDriver()(ok bool)//col:200
SetupDriverName()(ok bool)//col:230
}
)

func NewInstall() { return & install{} }

func (i *install)InstallDriver()(ok bool){//col:52
/*InstallDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName, LPCTSTR ServiceExe);
BOOLEAN
RemoveDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName);
BOOLEAN
StartDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName);
BOOLEAN
StopDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName);
BOOLEAN
InstallDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName, LPCTSTR ServiceExe)
{
    SC_HANDLE SchService;
    DWORD     LastError;
    SchService =
        CreateService(SchSCManager,          
                      DriverName,            
                      DriverName,            
                      SERVICE_ALL_ACCESS,    
                      SERVICE_KERNEL_DRIVER, 
                      SERVICE_DEMAND_START,  
                      SERVICE_ERROR_NORMAL,  
                      ServiceExe,            
                      NULL,                  
                      NULL,                  
                      NULL,                  
                      NULL,                  
                      NULL                   
        );
    if (SchService == NULL)
    {
        LastError = GetLastError();
        if (LastError == ERROR_SERVICE_EXISTS)
        {
            return TRUE;
        }
        else if (LastError == ERROR_SERVICE_MARKED_FOR_DELETE)
        {
            ShowMessages("err, previous instance of the service is not fully deleted. Try "
                         "again...\n");
            return FALSE;
        }
        else
        {
            ShowMessages("err, CreateService failed (%x)\n", LastError);
            return FALSE;
        }
    }
    if (SchService)
    {
        CloseServiceHandle(SchService);
    }
    return TRUE;
}*/
return true
}

func (i *install)ManageDriver()(ok bool){//col:99
/*ManageDriver(LPCTSTR DriverName, LPCTSTR ServiceName, UINT16 Function)
{
    SC_HANDLE schSCManager;
    BOOLEAN   Res = TRUE;
    if (!DriverName || !ServiceName)
    {
        ShowMessages("invalid Driver or Service provided to ManageDriver() \n");
        return FALSE;
    }
    schSCManager = OpenSCManager(NULL,                 
                                 NULL,                 
                                 SC_MANAGER_ALL_ACCESS 
    );
    if (!schSCManager)
    {
        ShowMessages("err, OpenSCManager failed (%x)\n", GetLastError());
        return FALSE;
    }
    switch (Function)
    {
    case DRIVER_FUNC_INSTALL:
        if (InstallDriver(schSCManager, DriverName, ServiceName))
        {
            Res = StartDriver(schSCManager, DriverName);
        }
        else
        {
            Res = FALSE;
        }
        break;
    case DRIVER_FUNC_STOP:
        Res = StopDriver(schSCManager, DriverName);
        break;
    case DRIVER_FUNC_REMOVE:
        Res = RemoveDriver(schSCManager, DriverName);
        break;
    default:
        ShowMessages("unknown ManageDriver() function. \n");
        Res = FALSE;
        break;
    }
    if (schSCManager)
    {
        CloseServiceHandle(schSCManager);
    }
    return Res;
}*/
return true
}

func (i *install)RemoveDriver()(ok bool){//col:124
/*RemoveDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName)
{
    SC_HANDLE SchService;
    BOOLEAN   Res;
    SchService = OpenService(SchSCManager, DriverName, SERVICE_ALL_ACCESS);
    if (SchService == NULL)
    {
        ShowMessages("err, OpenService failed (%x)\n", GetLastError());
        return FALSE;
    }
    if (DeleteService(SchService))
    {
        Res = TRUE;
    }
    else
    {
        ShowMessages("err, DeleteService failed (%x)\n", GetLastError());
        Res = FALSE;
    }
    if (SchService)
    {
        CloseServiceHandle(SchService);
    }
    return Res;
}*/
return true
}

func (i *install)StartDriver()(ok bool){//col:174
/*StartDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName)
{
    SC_HANDLE      SchService;
    DWORD          LastError;
    SERVICE_STATUS serviceStatus;
    UINT64         Status = TRUE;
    SchService = OpenService(SchSCManager, DriverName, SERVICE_ALL_ACCESS);
    if (SchService == NULL)
    {
        ShowMessages("err, OpenService failed (%x)\n", GetLastError());
        return FALSE;
    }
    if (!StartService(SchService, 
                      0,          
                      NULL        
                      ))
    {
        LastError = GetLastError();
        if (LastError == ERROR_SERVICE_ALREADY_RUNNING)
        {
        }
        else if (LastError == ERROR_PATH_NOT_FOUND)
        {
            ShowMessages("err, path to the driver not found, or the access to the driver file is limited\n");
            ShowMessages("most of the time, it's because anti-virus software is not finished scanning the drivers, so, if you try to load the driver again (re-enter the previous command), the problem will be solved\n");
            Status = FALSE;
        }
        else if (LastError == ERROR_INVALID_IMAGE_HASH)
        {
            ShowMessages(
                "err, failed loading driver\n"
                "it's because either the driver signature enforcement is enabled or HVCI prevents the driver from loading\n"
                "you should disable the driver signature enforcement by attaching WinDbg or from the boot menu\n"
                "if the driver signature enforcement is disabled, HVCI might prevent the driver from loading\n"
                "HyperDbg is not compatible with Virtualization Based Security (VBS)\n"
                "please follow the instructions from: https:
            Status = FALSE;
        }
        else
        {
            ShowMessages("err, StartService failure (%x)\n", LastError);
            Status = FALSE;
        }
    }
    if (SchService)
    {
        CloseServiceHandle(SchService);
    }
    return Status;
}*/
return true
}

func (i *install)StopDriver()(ok bool){//col:200
/*StopDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName)
{
    BOOLEAN        Res = TRUE;
    SC_HANDLE      SchService;
    SERVICE_STATUS serviceStatus;
    SchService = OpenService(SchSCManager, DriverName, SERVICE_ALL_ACCESS);
    if (SchService == NULL)
    {
        ShowMessages("err, OpenService failed (%x)\n", GetLastError());
        return FALSE;
    }
    if (ControlService(SchService, SERVICE_CONTROL_STOP, &serviceStatus))
    {
        Res = TRUE;
    }
    else
    {
        ShowMessages("err, ControlService failed (%x)\n", GetLastError());
        Res = FALSE;
    }
    if (SchService)
    {
        CloseServiceHandle(SchService);
    }
    return Res;
}*/
return true
}

func (i *install)SetupDriverName()(ok bool){//col:230
/*SetupDriverName(_Inout_updates_bytes_all_(BufferLength) PCHAR DriverLocation,
                ULONG                                         BufferLength)
{
    HANDLE  FileHandle;
    DWORD   DriverLocLen = 0;
    HMODULE ProcHandle   = GetModuleHandle(NULL);
    char *  Pos;
    GetModuleFileName(ProcHandle, DriverLocation, BufferLength);
    Pos = strrchr(DriverLocation, '\\');
    if (Pos != NULL)
    {
        *Pos = '\0';
    }
    if (FAILED(
            StringCbCat(DriverLocation, BufferLength, "\\" VMM_DRIVER_NAME ".sys")))
    {
        return FALSE;
    }
    if ((FileHandle = CreateFile(DriverLocation, GENERIC_READ, 0, NULL, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL)) ==
        INVALID_HANDLE_VALUE)
    {
        ShowMessages("%s.sys is not loaded.\n", VMM_DRIVER_NAME);
        return FALSE;
    }
    if (FileHandle)
    {
        CloseHandle(FileHandle);
    }
    return TRUE;
}*/
return true
}



