










#include "pch.h"

using namespace std;




extern BOOLEAN g_ExecutingScript;






VOID
CommandScriptHelp()
{
    ShowMessages(".script : runs a HyperDbg script.\n\n");

    ShowMessages("syntax : \.script [FilePath (string)] [Args (string)]\n");

    ShowMessages("\n");
    ShowMessages("\t\te.g : .script C:\\scripts\\script.ds\n");
    ShowMessages("\t\te.g : .script C:\\scripts\\script.ds 95 85 @rsp\n");
    ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\"\n");
    ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\" @rax\n");
    ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\" @rax @rcx+55 $pid\n");
    ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\" 12 55 @rip\n");
}






VOID
CommandScriptRunCommand(std::string Input, vector<string> PathAndArgs)
{
    int    CommandExecutionResult = 0;
    char * LineContent            = NULL;
    int    i                      = 0;

    
    
    
    
    
    
    
    for (auto item : PathAndArgs)
    {
        string ToReplace = "$arg" + std::to_string(i);
        i++;

        ReplaceAll(Input, ToReplace, item);
    }

    
    
    
    LineContent = (char *)Input.c_str();

    
    
    
    InterpreterRemoveComments(LineContent);

    if (IsEmptyString(LineContent))
    {
        return;
    }

    
    
    
    HyperDbgShowSignature();

    ShowMessages("%s\n", LineContent);

    CommandExecutionResult = HyperDbgInterpreter(LineContent);

    ShowMessages("\n");

    
    
    
    if (CommandExecutionResult == 1)
    {
        
        
        
        exit(0);
    }
}






VOID
HyperDbgScriptReadFileAndExecuteCommand(std::vector<std::string> & PathAndArgs)
{
    std::string Line;
    BOOLEAN     IsOpened         = FALSE;
    bool        Reset            = false;
    string      CommandToExecute = "";
    string      PathOfScriptFile = "";

    
    
    
    
    PathOfScriptFile = PathAndArgs.at(0);
    ReplaceAll(PathOfScriptFile, "\"", "");

    ifstream File(PathOfScriptFile);

    if (File.is_open())
    {
        IsOpened = TRUE;

        
        
        
        g_ExecutingScript = TRUE;

        
        
        
        Reset = true;

        while (std::getline(File, Line))
        {
            
            
            
            if (HyperDbgCheckMultilineCommand(Line, Reset))
            {
                
                
                
                if (Reset)
                {
                    CommandToExecute.clear();
                }

                
                
                
                Reset = false;

                
                
                
                CommandToExecute += Line + "\n";

                continue;
            }
            else
            {
                
                
                
                Reset = true;

                
                
                
                CommandToExecute += Line;
            }

            
            
            
            CommandScriptRunCommand(CommandToExecute, PathAndArgs);

            
            
            
            CommandToExecute.clear();
        }

        
        
        
        if (!CommandToExecute.empty())
        {
            CommandScriptRunCommand(CommandToExecute, PathAndArgs);

            
            
            
            CommandToExecute.clear();
        }

        
        
        
        g_ExecutingScript = FALSE;

        File.close();
    }

    if (!IsOpened)
    {
        ShowMessages("err, invalid file specified for the script\n");
    }
}








VOID
CommandScript(vector<string> SplittedCommand, string Command)
{
    vector<string> PathAndArgs;

    if (SplittedCommand.size() == 1)
    {
        ShowMessages("please specify a file\n");
        CommandScriptHelp();
        return;
    }

    
    
    
    Trim(Command);

    
    
    
    Command.erase(0, 7);

    
    
    
    Trim(Command);

    
    
    
    SplitPathAndArgs(PathAndArgs, Command);

    











    
    
    
    HyperDbgScriptReadFileAndExecuteCommand(PathAndArgs);
}
