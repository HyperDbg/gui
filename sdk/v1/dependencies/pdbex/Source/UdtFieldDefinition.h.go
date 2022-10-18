package Source

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\UdtFieldDefinition.h.back

type (
	UdtFieldDefinition interface {
		VisitBaseType() (ok bool) //col:95
	}
	udtFieldDefinition struct{}
)

func NewUdtFieldDefinition() UdtFieldDefinition { return &udtFieldDefinition{} }

func (u *udtFieldDefinition) VisitBaseType() (ok bool) { //col:95
	/*
				VisitBaseType(
					const SYMBOL* Symbol
					) override
				{
					if (Symbol->BaseType == btFloat && Symbol->Size == 10)
					{
						m_Comment += "";
					}
					if (Symbol->IsConst)
					{
						m_TypePrefix += "const ";
					}
					if (Symbol->IsVolatile)
					{
						m_TypePrefix += "volatile ";
					}
					const CHAR* BasicTypeString = PDB::GetBasicTypeString(Symbol, m_Settings->UseStdInt);
					m_TypePrefix += (BasicTypeString != nullptr ? BasicTypeString : "<unknown_type>");
				VisitPointerTypeEnd(
					const SYMBOL* Symbol
					) override
				{
					if (Symbol->u.Pointer.IsReference)
					{
						m_TypePrefix += "&";
					}
					else
					{
						m_TypePrefix += "*";
					}
					if (Symbol->IsConst)
					{
						m_TypePrefix += " const";
					}
					if (Symbol->IsVolatile)
					{
						m_TypePrefix += " volatile";
					}
				}
				void
				VisitArrayTypeEnd(
					const SYMBOL* Symbol
					) override
				{
					if (Symbol->u.Array.ElementCount == 0)
					{
						const_cast<SYMBOL*>(Symbol)->Size = 1;
						m_TypePrefix += "*";
						m_Comment += "";
					}
					else
					{
						m_TypeSuffix += "[" + std::to_string(Symbol->u.Array.ElementCount) + "]";
					}
				}
				void
				VisitFunctionTypeEnd(
					const SYMBOL* Symbol
					) override
				{
					m_TypePrefix += "void";
					m_Comment += "";
				}
				void
				SetMemberName(
					const CHAR* MemberName
					) override
				{
					m_MemberName = MemberName ? MemberName : std::string();
				GetPrintableDefinition() const override
				{
					return m_TypePrefix + " " + m_MemberName + m_TypeSuffix + m_Comment;
				}
				void
				SetSettings(
					void* MemberDefinitionSettings
					) override
				{
					static Settings DefaultSettings;
					if (MemberDefinitionSettings == nullptr)
					{
						MemberDefinitionSettings = &DefaultSettings;
					}
					m_Settings = static_cast<Settings*>(MemberDefinitionSettings);
				GetSettings() override
				{
					return &m_Settings;
				}
			private:
				std::string m_TypePrefix;
				std::string m_MemberName;
				std::string m_TypeSuffix;
				std::string m_Comment;
				Settings* m_Settings;
		};
	*/
	return true
}

