package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBReconstructorBase.h.back

type (
PdbReconstructorBase interface{
		OnEnumType()(ok bool)//col:124
		OnEnumTypeBegin()(ok bool)//col:240
		OnEnumTypeEnd()(ok bool)//col:349
		OnEnumField()(ok bool)//col:451
		OnUdt()(ok bool)//col:546
		OnUdtBegin()(ok bool)//col:633
		OnUdtEnd()(ok bool)//col:713
		OnUdtFieldBegin()(ok bool)//col:786
		OnUdtFieldEnd()(ok bool)//col:852
		OnUdtField()(ok bool)//col:911
		OnAnonymousUdtBegin()(ok bool)//col:962
		OnAnonymousUdtEnd()(ok bool)//col:1005
		OnUdtFieldBitFieldBegin()(ok bool)//col:1038
		OnUdtFieldBitFieldEnd()(ok bool)//col:1063
		OnPaddingMember()(ok bool)//col:1080
		OnPaddingBitFieldField()(ok bool)//col:1087
}
)

func NewPdbReconstructorBase() { return & pdbReconstructorBase{} }

func (p *pdbReconstructorBase)		OnEnumType()(ok bool){//col:124
/*		OnEnumType(
			const SYMBOL* Symbol
			)
		{
			return false;
		}
		virtual
		void
		OnEnumTypeBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnEnumTypeEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		bool
		OnUdt(
			const SYMBOL* Symbol
			)
		{
			return false;
		}
		virtual
		void
		OnUdtBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnEnumTypeBegin()(ok bool){//col:240
/*		OnEnumTypeBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnEnumTypeEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		bool
		OnUdt(
			const SYMBOL* Symbol
			)
		{
			return false;
		}
		virtual
		void
		OnUdtBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnEnumTypeEnd()(ok bool){//col:349
/*		OnEnumTypeEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		bool
		OnUdt(
			const SYMBOL* Symbol
			)
		{
			return false;
		}
		virtual
		void
		OnUdtBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnEnumField()(ok bool){//col:451
/*		OnEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		bool
		OnUdt(
			const SYMBOL* Symbol
			)
		{
			return false;
		}
		virtual
		void
		OnUdtBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdt()(ok bool){//col:546
/*		OnUdt(
			const SYMBOL* Symbol
			)
		{
			return false;
		}
		virtual
		void
		OnUdtBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtBegin()(ok bool){//col:633
/*		OnUdtBegin(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtEnd()(ok bool){//col:713
/*		OnUdtEnd(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtFieldBegin()(ok bool){//col:786
/*		OnUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtFieldEnd()(ok bool){//col:852
/*		OnUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtField()(ok bool){//col:911
/*		OnUdtField(
			const SYMBOL_UDT_FIELD* UdtField,
			UdtFieldDefinitionBase* MemberDefinition
			)
		{
		}
		virtual
		void
		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnAnonymousUdtBegin()(ok bool){//col:962
/*		OnAnonymousUdtBegin(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField
			)
		{
		}
		virtual
		void
		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnAnonymousUdtEnd()(ok bool){//col:1005
/*		OnAnonymousUdtEnd(
			UdtKind Kind,
			const SYMBOL_UDT_FIELD* FirstUdtField,
			const SYMBOL_UDT_FIELD* LastUdtField,
			DWORD Size
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtFieldBitFieldBegin()(ok bool){//col:1038
/*		OnUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnUdtFieldBitFieldEnd()(ok bool){//col:1063
/*		OnUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField
			)
		{
		}
		virtual
		void
		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnPaddingMember()(ok bool){//col:1080
/*		OnPaddingMember(
			const SYMBOL_UDT_FIELD* UdtField,
			BasicType PaddingBasicType,
			DWORD PaddingBasicTypeSize,
			DWORD PaddingSize
			)
		{
		}
		virtual
		void
		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}

func (p *pdbReconstructorBase)		OnPaddingBitFieldField()(ok bool){//col:1087
/*		OnPaddingBitFieldField(
			const SYMBOL_UDT_FIELD* UdtField,
			const SYMBOL_UDT_FIELD* PreviousUdtField
			)
		{
		}
};*/
return true
}



