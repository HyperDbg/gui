










#include "pch.h"




extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;






VOID
CommandSymHelp()
{
    ShowMessages(".sym : perfroms the symbol actions.\n\n");

    ShowMessages("syntax : \t.sym [table]\n");
    ShowMessages("syntax : \t.sym [reload] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \t.sym [download]\n");
    ShowMessages("syntax : \t.sym [load]\n");
    ShowMessages("syntax : \t.sym [unload]\n");
    ShowMessages("syntax : \t.sym [add] [base Address (hex)] [path Path (string)]\n");

    ShowMessages("\n");
    ShowMessages("\t\te.g : .sym table\n");
    ShowMessages("\t\te.g : .sym reload\n");
    ShowMessages("\t\te.g : .sym load\n");
    ShowMessages("\t\te.g : .sym download\n");
    ShowMessages("\t\te.g : .sym add base fffff8077356000 path c:\\symbols\\my_dll.pdb\n");
    ShowMessages("\t\te.g : .sym unload\n");
}








VOID
CommandSym(vector<string> SplittedCommand, string Command)
{
    UINT64 BaseAddress   = NULL;
    UINT32 UserProcessId = NULL;

    if (SplittedCommand.size() == 1)
    {
        ShowMessages("incorrect use of '.sym'\n\n");
        CommandSymHelp();
        return;
    }

    if (!SplittedCommand.at(1).compare("table"))
    {
        
        
        
        if (SplittedCommand.size() != 2)
        {
            ShowMessages("incorrect use of '.sym'\n\n");
            CommandSymHelp();
            return;
        }

        
        
        
        SymbolBuildAndShowSymbolTable();
    }
    else if (!SplittedCommand.at(1).compare("load") || !SplittedCommand.at(1).compare("download"))
    {
        
        
        
        if (SplittedCommand.size() != 2)
        {
            ShowMessages("incorrect use of '.sym'\n\n");
            CommandSymHelp();
            return;
        }

        
        
        
        if (!SplittedCommand.at(1).compare("load"))
        {
            SymbolLoadOrDownloadSymbols(FALSE, FALSE);
        }
        else if (!SplittedCommand.at(1).compare("download"))
        {
            SymbolLoadOrDownloadSymbols(TRUE, FALSE);
        }
    }
    else if (!SplittedCommand.at(1).compare("reload"))
    {
        
        
        
        if (SplittedCommand.size() != 2 && SplittedCommand.size() != 4)
        {
            ShowMessages("incorrect use of '.sym'\n\n");
            CommandSymHelp();
            return;
        }

        
        
        
        if (SplittedCommand.size() == 4)
        {
            if (!SplittedCommand.at(2).compare("pid"))
            {
                if (!ConvertStringToUInt32(SplittedCommand.at(3), &UserProcessId))
                {
                    
                    
                    
                    ShowMessages("err, couldn't resolve error at '%s'\n\n",
                                 SplittedCommand.at(3).c_str());
                    CommandSymHelp();
                    return;
                }
            }
            else
            {
                ShowMessages("incorrect use of '.sym'\n\n");
                CommandSymHelp();
                return;
            }
        }

        
        
        
        if (g_IsSerialConnectedToRemoteDebuggee)
        {
            
            
            
            SymbolReloadSymbolTableInDebuggerMode(UserProcessId);
        }
        else
        {
            
            
            
            if (UserProcessId == NULL)
            {
                
                
                
                
                
                
                if (g_ActiveProcessDebuggingState.IsActive)
                {
                    UserProcessId = g_ActiveProcessDebuggingState.ProcessId;
                }
                else
                {
                    UserProcessId = GetCurrentProcessId();
                }
            }

            
            
            
            if (SymbolLocalReload(UserProcessId))
            {
                ShowMessages("symbol table updated successfully\n");
            }
        }
    }
    else if (!SplittedCommand.at(1).compare("unload"))
    {
        
        
        
        if (SplittedCommand.size() != 2)
        {
            ShowMessages("incorrect use of '.sym'\n\n");
            CommandSymHelp();
            return;
        }

        
        
        
        
        ScriptEngineUnloadAllSymbolsWrapper();

        
        
        
        
    }
    else if (!SplittedCommand.at(1).compare("add"))
    {
        
        
        
        if (SplittedCommand.size() < 6)
        {
            ShowMessages("incorrect use of '.sym'\n\n");
            CommandSymHelp();
            return;
        }

        if (!SplittedCommand.at(2).compare("base"))
        {
            string Delimiter = "";
            string PathToPdb = "";
            if (!ConvertStringToUInt64(SplittedCommand.at(3), &BaseAddress))
            {
                ShowMessages("please add a valid hex address to be used as the base address\n\n");
                CommandSymHelp();
                return;
            }

            
            
            
            if (SplittedCommand.at(4).compare("path"))
            {
                ShowMessages("incorrect use of '.sym'\n\n");
                CommandSymHelp();
                return;
            }

            
            
            
            Delimiter = "path ";
            PathToPdb = Command.substr(Command.find(Delimiter) + 5, Command.size());

            
            
            
            if (!IsFileExistA(PathToPdb.c_str()))
            {
                ShowMessages("pdb file not found\n");
                return;
            }

            ShowMessages("loading module symbol at '%s'\n", PathToPdb.c_str());

            
            
            
            
            ScriptEngineLoadFileSymbolWrapper(BaseAddress, PathToPdb.c_str());
        }
        else
        {
            ShowMessages("incorrect use of '.sym'\n\n");
            CommandSymHelp();
            return;
        }
    }
    else
    {
        ShowMessages("unknown parameter at '%s'\n\n", SplittedCommand.at(1).c_str());
        CommandSymHelp();
        return;
    }
}
