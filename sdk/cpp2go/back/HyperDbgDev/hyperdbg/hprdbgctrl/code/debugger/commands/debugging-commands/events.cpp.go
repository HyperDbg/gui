package debugging-commands


type (
Events interface{
CommandEventsHelp()(ok bool)//col:52
CommandEvents()(ok bool)//col:149
CommandEventQueryEventState()(ok bool)//col:195
CommandEventsShowEvents()(ok bool)//col:253
CommandEventDisableEvent()(ok bool)//col:304
CommandEventEnableEvent()(ok bool)//col:355
CommandEventClearEvent()(ok bool)//col:442
CommandEventsClearAllEventsAndResetTags()(ok bool)//col:470
CommandEventsHandleModifiedEvent()(ok bool)//col:611
CommandEventsModifyAndQueryEvents()(ok bool)//col:716
}

)

func NewEvents() { return & events{} }

func (e *events)CommandEventsHelp()(ok bool){//col:52

















return true
}


func (e *events)CommandEvents()(ok bool){//col:149























































return true
}


func (e *events)CommandEventQueryEventState()(ok bool){//col:195


























return true
}


func (e *events)CommandEventsShowEvents()(ok bool){//col:253































return true
}


func (e *events)CommandEventDisableEvent()(ok bool){//col:304

























return true
}


func (e *events)CommandEventEnableEvent()(ok bool){//col:355

























return true
}


func (e *events)CommandEventClearEvent()(ok bool){//col:442






































return true
}


func (e *events)CommandEventsClearAllEventsAndResetTags()(ok bool){//col:470









return true
}


func (e *events)CommandEventsHandleModifiedEvent()(ok bool){//col:611




































































































return true
}


func (e *events)CommandEventsModifyAndQueryEvents()(ok bool){//col:716










































return true
}




