package src



type (
	MetaInfo interface {
		const_charPtr_ZydisCategoryGetString() (ok bool) //col:8
		const_charPtr_ZydisISASetGetString() (ok bool)   //col:16
		const_charPtr_ZydisISAExtGetString() (ok bool)   //col:24
	}
	metaInfo struct{}
)

func NewMetaInfo() MetaInfo { return &metaInfo{} }

func (m *metaInfo) const_charPtr_ZydisCategoryGetString() (ok bool) { //col:8











	return true
}


func (m *metaInfo) const_charPtr_ZydisISASetGetString() (ok bool) { //col:16











	return true
}


func (m *metaInfo) const_charPtr_ZydisISAExtGetString() (ok bool) { //col:24











	return true
}


