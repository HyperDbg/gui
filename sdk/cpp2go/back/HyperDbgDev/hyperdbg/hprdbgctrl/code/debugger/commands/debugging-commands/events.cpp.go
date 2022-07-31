package debugging-commands
type (
Events interface{
CommandEventsHelp()(ok bool)//col:52
CommandEvents()(ok bool)//col:148
CommandEventQueryEventState()(ok bool)//col:193
CommandEventsShowEvents()(ok bool)//col:250
CommandEventDisableEvent()(ok bool)//col:300
CommandEventEnableEvent()(ok bool)//col:350
CommandEventClearEvent()(ok bool)//col:436
CommandEventsClearAllEventsAndResetTags()(ok bool)//col:463
CommandEventsHandleModifiedEvent()(ok bool)//col:603
CommandEventsModifyAndQueryEvents()(ok bool)//col:707
}

)
func NewEvents() { return & events{} }
func (e *events)CommandEventsHelp()(ok bool){//col:52
return true
}

func (e *events)CommandEvents()(ok bool){//col:148
return true
}

func (e *events)CommandEventQueryEventState()(ok bool){//col:193
return true
}

func (e *events)CommandEventsShowEvents()(ok bool){//col:250
return true
}

func (e *events)CommandEventDisableEvent()(ok bool){//col:300
return true
}

func (e *events)CommandEventEnableEvent()(ok bool){//col:350
return true
}

func (e *events)CommandEventClearEvent()(ok bool){//col:436
return true
}

func (e *events)CommandEventsClearAllEventsAndResetTags()(ok bool){//col:463
return true
}

func (e *events)CommandEventsHandleModifiedEvent()(ok bool){//col:603
return true
}

func (e *events)CommandEventsModifyAndQueryEvents()(ok bool){//col:707
return true
}

