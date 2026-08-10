/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
 * author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
 * license:  GLPv3
 */
#include "pluginmain.h"
#include "internal.h"

#define plugin_name "StaticAnalysis"
#define plugin_version 001

int pluginHandle;
HWND hwndDlg;
int hMenu;

DLL_EXPORT bool pluginit(PLUG_INITSTRUCT* initStruct)
{
    initStruct->pluginVersion=plugin_version;
    initStruct->sdkVersion=PLUG_SDKVERSION;
    strcpy(initStruct->pluginName, plugin_name);
    pluginHandle=initStruct->pluginHandle;
    internalInit(initStruct);
    return true;
}

DLL_EXPORT bool plugstop()
{
    internalStop();
    return true;
}

DLL_EXPORT void plugsetup(PLUG_SETUPSTRUCT* setupStruct)
{
    hwndDlg=setupStruct->hwndDlg;
    hMenu=setupStruct->hMenu;
    internalSetup();
}

extern "C" DLL_EXPORT BOOL APIENTRY DllMain(HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpvReserved)
{
    return TRUE;
}
