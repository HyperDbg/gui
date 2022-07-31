package code
type (
Scanner interface{
GetToken()(ok bool)//col:730
Scan()(ok bool)//col:797
sgetc()(ok bool)//col:819
IsKeyword()(ok bool)//col:848
IsRegister()(ok bool)//col:862
IsId()(ok bool)//col:875
}

)
func NewScanner() { return & scanner{} }
func (s *scanner)GetToken()(ok bool){//col:730
return true
}

func (s *scanner)Scan()(ok bool){//col:797
return true
}

func (s *scanner)sgetc()(ok bool){//col:819
return true
}

func (s *scanner)IsKeyword()(ok bool){//col:848
return true
}

func (s *scanner)IsRegister()(ok bool){//col:862
return true
}

func (s *scanner)IsId()(ok bool){//col:875
return true
}

