#include "internal.h"
#include "pluginsdk\TitanEngine\TitanEngine.hpp"
#include <windows.h>
#include <stdio.h>
#include <psapi.h>
#include "AnalysisRunner.h"
#include "ApiDB.h"

ApiDB *db;

duint memfindbaseaddr(HANDLE hProcess, duint addr, duint* size)
{
	MEMORY_BASIC_INFORMATION mbi;
	DWORD numBytes;
	duint MyAddress=0, newAddress=0;
	do
	{
		numBytes=VirtualQueryEx(hProcess, (LPCVOID)MyAddress, &mbi, sizeof(mbi));
		newAddress=(duint)mbi.BaseAddress+mbi.RegionSize;
		if((mbi.State==MEM_COMMIT) && (addr<newAddress) && (addr>=MyAddress))
		{
			if(size)
				*size=mbi.RegionSize;
			return (duint)mbi.BaseAddress;
		}
		if(newAddress<=MyAddress)
			numBytes=0;
		else
			MyAddress=newAddress;
	}
	while(numBytes);
	return 0;
}

static bool cbRunAnalysis(){
	duint entry = TE::UE::GetContextData(TE::UE::UE_CIP);
	char mod[MAX_MODULE_SIZE]="";
	if(!DbgGetModuleAt(entry, mod))
	{
		_plugin_logprintf("[StaticAnalysis] no module found at %p... for analysis\n", entry);
		return false;
	}
	duint base=DbgModBaseFromName(mod);
	if(!base)
	{
		_plugin_logputs("[StaticAnalysis] could not get module base...");
		return false;
	}
	duint size = DbgMemGetPageSize(base);
	duint BaseAddr = base + size;
	duint memSize=0;
	HANDLE hProcess = ((PROCESS_INFORMATION*)TE::UE::GetProcessInformation())->hProcess;
	memfindbaseaddr(hProcess,entry,&memSize);


	AnalysisRunner AR(BaseAddr,memSize);
	AR.setFunctionInformation(db);
	AR.start();
	return true;
}


static void cbMenuEntry(CBTYPE cbType, void* callbackInfo)
{
    PLUG_CB_MENUENTRY* info=(PLUG_CB_MENUENTRY*)callbackInfo;
    switch(info->hEntry)
    {
    case MENU_ANALYSIS:
		cbRunAnalysis();
        break;
    }
}


void internalInit(PLUG_INITSTRUCT* initStruct)
{
	_plugin_registercallback(pluginHandle, CB_MENUENTRY, cbMenuEntry);
}

void internalStop()
{
	_plugin_unregistercallback(pluginHandle, CB_MENUENTRY);
}

void internalSetup()
{
	db = new ApiDB();
    _plugin_menuaddentry(hMenu, MENU_ANALYSIS, "&Analyze Code...");
}