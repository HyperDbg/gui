#include "pch.h"
VOID CommandSympathHelp() {
  ShowMessages(".sympath : shows and sets the symbol server and path.\n\n");
  ShowMessages("syntax : \t.sympath\n");
  ShowMessages("syntax : \t.sympath [SymServer (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .sympath\n");
  ShowMessages(
      "\t\te.g : .sympath "
      "SRV*c:\\Symbols*https:msdl.microsoft.com/download/symbols \n");
}

VOID CommandSympath(vector<string> SplittedCommand, string Command) {
  string SymbolServer = "";
  string Token;
  if (SplittedCommand.size() == 1) {
    if (!CommandSettingsGetValueFromConfigFile("SymbolServer", SymbolServer)) {
      ShowMessages(
          "symbol server is not configured, please use '.help .sympath'\n");
    } else {
      ShowMessages("current symbol server is : %s\n", SymbolServer.c_str());
    }
  } else {
    Trim(Command);
    Command.erase(0, 8);
    Trim(Command);
    char Delimiter = '*';
    if (Command.find(Delimiter) != std::string::npos) {
      Token = Command.substr(0, Command.find(Delimiter));
      transform(Token.begin(), Token.end(), Token.begin(), ::tolower);
      if (!Token.compare("srv")) {
        CommandSettingsSetValueFromConfigFile("SymbolServer", Command);
        //
        ShowMessages("symbol server/path is configured successfully\n");
        ShowMessages("use '.sym load', '.sym reload', or '.sym download' to "
                     "load pdb files\n");
      } else {
        ShowMessages("symbol path is invalid\n\n");
        CommandSympathHelp();
        return;
      }
    } else {
      ShowMessages("symbol path is invalid\n\n");
      CommandSympathHelp();
      return;
    }
  }
}

