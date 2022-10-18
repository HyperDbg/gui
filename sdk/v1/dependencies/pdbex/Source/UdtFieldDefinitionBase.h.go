package Source

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\UdtFieldDefinitionBase.h.back

type (
	UdtFieldDefinitionBase interface {
		VisitBaseType() (ok bool) //col:85
		GetSettings() (ok bool)   //col:90
	}
	udtFieldDefinitionBase struct{}
)

func NewUdtFieldDefinitionBase() UdtFieldDefinitionBase { return &udtFieldDefinitionBase{} }

func (u *udtFieldDefinitionBase) VisitBaseType() (ok bool) { //col:85
	/*
				VisitBaseType(
					const SYMBOL* Symbol
					)
				{
				}
				virtual
				void
				VisitPointerTypeBegin(
					const SYMBOL* Symbol
				)
				{
				}
				virtual
				void
				VisitPointerTypeEnd(
					const SYMBOL* Symbol
					)
				{
				}
				virtual
				void
				VisitArrayTypeBegin(
					const SYMBOL* Symbol
				)
				{
				}
				virtual
				void
				VisitArrayTypeEnd(
					const SYMBOL* Symbol
					)
				{
				}
				virtual
				void
				VisitFunctionTypeBegin(
					const SYMBOL* Symbol
				)
				{
				}
				virtual
				void
				VisitFunctionTypeEnd(
					const SYMBOL* Symbol
					)
				{
				}
				virtual
				void
				VisitFunctionArgTypeBegin(
					const SYMBOL* Symbol
				)
				{
				}
				virtual
				void
				VisitFunctionArgTypeEnd(
					const SYMBOL* Symbol
					)
				{
				}
				virtual
				void
				SetMemberName(
					const CHAR* MemberName
				)
				{
				}
				virtual
				std::string
				GetPrintableDefinition() const
				{
					return std::string();
				SetSettings(
					void* MemberDefinitionSettings
					)
				{
				}
				virtual
				void*
				GetSettings()
				{
					return nullptr;
				}
		};
	*/
	return true
}

func (u *udtFieldDefinitionBase) GetSettings() (ok bool) { //col:90
	/*
				GetSettings()
				{
					return nullptr;
				}
		};
	*/
	return true
}

