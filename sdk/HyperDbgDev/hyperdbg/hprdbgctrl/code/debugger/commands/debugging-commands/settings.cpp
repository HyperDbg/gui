#include "pch.h"
extern BOOLEAN g_AutoUnpause;
extern BOOLEAN g_AutoFlush;
extern BOOLEAN g_AddressConversion;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern UINT32  g_DisassemblerSyntax;
VOID
CommandSettingsHelp() {
    ShowMessages(
        "settings : queries, sets, or changes a value for a sepcial settings option.\n\n");
    ShowMessages("syntax : \tsettings [OptionName (string)]\n");
    ShowMessages("syntax : \tsettings [OptionName (string)] [Value (hex)]\n");
    ShowMessages("syntax : \tsettings [OptionName (string)] [Value (string)]\n");
    ShowMessages("syntax : \tsettings [OptionName (string)] [on|off]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : settings autounpause\n");
    ShowMessages("\t\te.g : settings autounpause on\n");
    ShowMessages("\t\te.g : settings autounpause off\n");
    ShowMessages("\t\te.g : settings addressconversion on\n");
    ShowMessages("\t\te.g : settings addressconversion off\n");
    ShowMessages("\t\te.g : settings autoflush on\n");
    ShowMessages("\t\te.g : settings autoflush off\n");
    ShowMessages("\t\te.g : settings syntax intel\n");
    ShowMessages("\t\te.g : settings syntax att\n");
    ShowMessages("\t\te.g : settings syntax masm\n");
}

VOID
CommandSettingsAddressConversion(vector<string> SplittedCommand) {
    if (SplittedCommand.size() == 2) {
        if (g_AddressConversion) {
            ShowMessages("address conversion is enabled\n");
        } else {
            ShowMessages("address conversion is disabled\n");
        }
    } else if (SplittedCommand.size() == 3) {
        if (!SplittedCommand.at(2).compare("on")) {
            g_AddressConversion = TRUE;
            ShowMessages("set address conversion to enabled\n");
        } else if (!SplittedCommand.at(2).compare("off")) {
            g_AddressConversion = FALSE;
            ShowMessages("set address conversion to disabled\n");
        } else {
            ShowMessages("incorrect use of 'settings', please use 'help settings' "
                         "for more details\n");
            return;
        }
    } else {
        ShowMessages("incorrect use of 'settings', please use 'help settings' "
                     "for more details\n");
        return;
    }
}

VOID
CommandSettingsAutoFlush(vector<string> SplittedCommand) {
    if (SplittedCommand.size() == 2) {
        if (g_AutoFlush) {
            ShowMessages("auto-flush is enabled\n");
        } else {
            ShowMessages("auto-flush is disabled\n");
        }
    } else if (SplittedCommand.size() == 3) {
        if (!SplittedCommand.at(2).compare("on")) {
            g_AutoFlush = TRUE;
            ShowMessages("set auto-flush to enabled\n");
        } else if (!SplittedCommand.at(2).compare("off")) {
            g_AutoFlush = FALSE;
            ShowMessages("set auto-flush to disabled\n");
        } else {
            ShowMessages("incorrect use of 'settings', please use 'help settings' "
                         "for more details\n");
            return;
        }
    } else {
        ShowMessages("incorrect use of 'settings', please use 'help settings' "
                     "for more details\n");
        return;
    }
}

VOID
CommandSettingsAutoUpause(vector<string> SplittedCommand) {
    if (SplittedCommand.size() == 2) {
        if (g_AutoUnpause) {
            ShowMessages("auto-unpause is enabled\n");
        } else {
            ShowMessages("auto-unpause is disabled\n");
        }
    } else if (SplittedCommand.size() == 3) {
        if (!SplittedCommand.at(2).compare("on")) {
            g_AutoUnpause = TRUE;
            ShowMessages("set auto-unpause to enabled\n");
        } else if (!SplittedCommand.at(2).compare("off")) {
            g_AutoUnpause = FALSE;
            ShowMessages("set auto-unpause to disabled\n");
        } else {
            ShowMessages("incorrect use of 'settings', please use 'help settings' "
                         "for more details\n");
            return;
        }
    } else {
        ShowMessages("incorrect use of 'settings', please use 'help settings' "
                     "for more details\n");
        return;
    }
}

VOID
CommandSettingsSyntax(vector<string> SplittedCommand) {
    if (SplittedCommand.size() == 2) {
        if (g_DisassemblerSyntax == 1) {
            ShowMessages("disassembler syntax is : intel\n");
        } else if (g_DisassemblerSyntax == 2) {
            ShowMessages("disassembler syntax is : at&t\n");
        } else if (g_DisassemblerSyntax == 3) {
            ShowMessages("disassembler syntax is : masm\n");
        } else {
            ShowMessages("unknown syntax\n");
        }
    } else if (SplittedCommand.size() == 3) {
        if (!SplittedCommand.at(2).compare("intel")) {
            g_DisassemblerSyntax = 1;
            ShowMessages("set syntax to intel\n");
        } else if (!SplittedCommand.at(2).compare("att") ||
                   !SplittedCommand.at(2).compare("at&t")) {
            g_DisassemblerSyntax = 2;
            ShowMessages("set syntax to at&t\n");
        } else if (!SplittedCommand.at(2).compare("masm")) {
            g_DisassemblerSyntax = 3;
            ShowMessages("set syntax to masm\n");
        } else {
            ShowMessages("incorrect use of 'settings', please use 'help settings' "
                         "for more details\n");
            return;
        }
    } else {
        ShowMessages("incorrect use of 'settings', please use 'help settings' "
                     "for more details\n");
        return;
    }
}

VOID
CommandSettings(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() <= 1) {
        ShowMessages("incorrect use of 'settings'\n\n");
        CommandSettingsHelp();
        return;
    }
    if (!SplittedCommand.at(1).compare("autounpause")) {
        CommandSettingsAutoUpause(SplittedCommand);
    } else if (!SplittedCommand.at(1).compare("syntax")) {
        if (g_IsConnectedToRemoteDebuggee) {
            RemoteConnectionSendCommand(Command.c_str(), Command.length() + 1);
        } else {
            CommandSettingsSyntax(SplittedCommand);
        }
    } else if (!SplittedCommand.at(1).compare("autoflush")) {
        if (g_IsConnectedToRemoteDebuggee) {
            RemoteConnectionSendCommand(Command.c_str(), Command.length() + 1);
        } else {
            CommandSettingsAutoFlush(SplittedCommand);
        }
    } else if (!SplittedCommand.at(1).compare("addressconversion")) {
        if (g_IsConnectedToRemoteDebuggee) {
            RemoteConnectionSendCommand(Command.c_str(), Command.length() + 1);
        } else {
            CommandSettingsAddressConversion(SplittedCommand);
        }
    } else {
        ShowMessages("incorrect use of 'settings', please use 'help settings' "
                     "for more details\n");
        return;
    }
}
