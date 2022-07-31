package debugging-commands


type (
Settings interface{
CommandSettingsHelp()(ok bool)//col:50
CommandSettingsGetValueFromConfigFile()(ok bool)//col:105
CommandSettingsSetValueFromConfigFile()(ok bool)//col:155
CommandSettingsLoadDefaultValuesFromConfigFile()(ok bool)//col:267
CommandSettingsAddressConversion()(ok bool)//col:332
CommandSettingsAutoFlush()(ok bool)//col:397
CommandSettingsAutoUpause()(ok bool)//col:461
CommandSettingsSyntax()(ok bool)//col:541
CommandSettings()(ok bool)//col:634
}

)

func NewSettings() { return & settings{} }

func (s *settings)CommandSettingsHelp()(ok bool){//col:50




















return true
}


func (s *settings)CommandSettingsGetValueFromConfigFile()(ok bool){//col:105
























return true
}


func (s *settings)CommandSettingsSetValueFromConfigFile()(ok bool){//col:155














return true
}


func (s *settings)CommandSettingsLoadDefaultValuesFromConfigFile()(ok bool){//col:267





































































return true
}


func (s *settings)CommandSettingsAddressConversion()(ok bool){//col:332









































return true
}


func (s *settings)CommandSettingsAutoFlush()(ok bool){//col:397









































return true
}


func (s *settings)CommandSettingsAutoUpause()(ok bool){//col:461









































return true
}


func (s *settings)CommandSettingsSyntax()(ok bool){//col:541
























































return true
}


func (s *settings)CommandSettings()(ok bool){//col:634




















































return true
}




