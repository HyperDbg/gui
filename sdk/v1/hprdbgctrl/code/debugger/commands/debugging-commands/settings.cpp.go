package debugging_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\settings.cpp.back

type (
	Settings interface {
		CommandSettingsHelp() (ok bool) //col:43
	}
	settings struct{}
)

func NewSettings() Settings { return &settings{} }

func (s *settings) CommandSettingsHelp() (ok bool) { //col:43
	/*
	   CommandSettingsHelp()

	   	{
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

	   CommandSettingsGetValueFromConfigFile(std::string OptionName, std::string & OptionValue)

	   	{
	   	    inipp::Ini<char> Ini;
	   	    WCHAR            ConfigPath[260] = {0};
	   	    std::string      OptionValueFromFile;
	   	    GetConfigFilePath(ConfigPath);
	   	    if (!IsFileExistW(ConfigPath))
	   	    {
	   	        return FALSE;
	   	    }
	   	    ifstream Is(ConfigPath);
	   	    Ini.parse(Is);
	   	    inipp::get_value(Ini.sections["DEFAULT"], OptionName, OptionValueFromFile);
	   	    Is.close();
	   	    if (!OptionValueFromFile.empty())
	   	    {
	   	        OptionValue = OptionValueFromFile;
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        return FALSE;
	   	    }
	   	}
	*/
	return true
}

