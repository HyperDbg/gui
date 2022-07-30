package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\src\MetaInfo.c.back

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
	/*const char* ZydisCategoryGetString(ZydisInstructionCategory category)
	  {
	      if ((ZyanUSize)category >= ZYAN_ARRAY_LENGTH(STR_INSTRUCTIONCATEGORY))
	      {
	          return ZYAN_NULL;
	      }
	      return STR_INSTRUCTIONCATEGORY[category];
	  }*/
	return true
}

func (m *metaInfo) const_charPtr_ZydisISASetGetString() (ok bool) { //col:16
	/*const char* ZydisISASetGetString(ZydisISASet isa_set)
	  {
	      if ((ZyanUSize)isa_set >= ZYAN_ARRAY_LENGTH(STR_ISASET))
	      {
	          return ZYAN_NULL;
	      }
	      return STR_ISASET[isa_set];
	  }*/
	return true
}

func (m *metaInfo) const_charPtr_ZydisISAExtGetString() (ok bool) { //col:24
	/*const char* ZydisISAExtGetString(ZydisISAExt isa_ext)
	  {
	      if ((ZyanUSize)isa_ext >= ZYAN_ARRAY_LENGTH(STR_ISAEXT))
	      {
	          return ZYAN_NULL;
	      }
	      return STR_ISAEXT[isa_ext];
	  }*/
	return true
}
