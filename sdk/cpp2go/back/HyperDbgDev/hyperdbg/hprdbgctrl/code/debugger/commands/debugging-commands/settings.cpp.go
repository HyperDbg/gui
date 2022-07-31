package debugging-commands
type (
Settings interface{
CommandSettingsHelp()(ok bool)//col:50
CommandSettingsGetValueFromConfigFile()(ok bool)//col:104
CommandSettingsSetValueFromConfigFile()(ok bool)//col:153
CommandSettingsLoadDefaultValuesFromConfigFile()(ok bool)//col:264
CommandSettingsAddressConversion()(ok bool)//col:328
CommandSettingsAutoFlush()(ok bool)//col:392
CommandSettingsAutoUpause()(ok bool)//col:455
CommandSettingsSyntax()(ok bool)//col:534
CommandSettings()(ok bool)//col:626
}

)
func NewSettings() { return & settings{} }
func (s *settings)CommandSettingsHelp()(ok bool){//col:50
return true
}

func (s *settings)CommandSettingsGetValueFromConfigFile()(ok bool){//col:104
return true
}

func (s *settings)CommandSettingsSetValueFromConfigFile()(ok bool){//col:153
return true
}

func (s *settings)CommandSettingsLoadDefaultValuesFromConfigFile()(ok bool){//col:264
return true
}

func (s *settings)CommandSettingsAddressConversion()(ok bool){//col:328
return true
}

func (s *settings)CommandSettingsAutoFlush()(ok bool){//col:392
return true
}

func (s *settings)CommandSettingsAutoUpause()(ok bool){//col:455
return true
}

func (s *settings)CommandSettingsSyntax()(ok bool){//col:534
return true
}

func (s *settings)CommandSettings()(ok bool){//col:626
return true
}

