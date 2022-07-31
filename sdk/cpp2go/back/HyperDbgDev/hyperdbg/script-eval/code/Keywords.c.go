package code
type (
Keywords interface{
ScriptEngineKeywordPoi()(ok bool)//col:51
ScriptEngineKeywordHi()(ok bool)//col:85
ScriptEngineKeywordLow()(ok bool)//col:119
ScriptEngineKeywordDb()(ok bool)//col:153
ScriptEngineKeywordDd()(ok bool)//col:187
ScriptEngineKeywordDw()(ok bool)//col:221
ScriptEngineKeywordDq()(ok bool)//col:255
}

)
func NewKeywords() { return & keywords{} }
func (k *keywords)ScriptEngineKeywordPoi()(ok bool){//col:51
return true
}

func (k *keywords)ScriptEngineKeywordHi()(ok bool){//col:85
return true
}

func (k *keywords)ScriptEngineKeywordLow()(ok bool){//col:119
return true
}

func (k *keywords)ScriptEngineKeywordDb()(ok bool){//col:153
return true
}

func (k *keywords)ScriptEngineKeywordDd()(ok bool){//col:187
return true
}

func (k *keywords)ScriptEngineKeywordDw()(ok bool){//col:221
return true
}

func (k *keywords)ScriptEngineKeywordDq()(ok bool){//col:255
return true
}

