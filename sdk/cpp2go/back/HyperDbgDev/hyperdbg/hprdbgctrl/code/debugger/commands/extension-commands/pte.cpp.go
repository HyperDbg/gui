package extension-commands


type (
Pte interface{
CommandPteHelp()(ok bool)//col:37
CommandPteShowResults()(ok bool)//col:82
CommandPte()(ok bool)//col:249
}

)

func NewPte() { return & pte{} }

func (p *pte)CommandPteHelp()(ok bool){//col:37










return true
}


func (p *pte)CommandPteShowResults()(ok bool){//col:82























return true
}


func (p *pte)CommandPte()(ok bool){//col:249


































































































return true
}




