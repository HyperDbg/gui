package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBReconstructorBase.h.back

type (
PdbReconstructorBase interface{
		OnEnumType()(ok bool)//col:216
}
)

func NewPdbReconstructorBase() { return & pdbReconstructorBase{} }

func (p *pdbReconstructorBase)		OnEnumType()(ok bool){//col:216
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



