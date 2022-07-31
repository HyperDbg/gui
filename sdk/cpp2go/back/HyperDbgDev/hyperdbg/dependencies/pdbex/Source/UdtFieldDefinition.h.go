package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\UdtFieldDefinition.h.back

type (
UdtFieldDefinition interface{
		VisitBaseType()(ok bool)//col:173
}
)

func NewUdtFieldDefinition() { return & udtFieldDefinition{} }

func (u *udtFieldDefinition)		VisitBaseType()(ok bool){//col:173
/*		VisitBaseType(
			const SYMBOL* Symbol
			) override
		{
			if (Symbol->BaseType == btFloat && Symbol->Size == 10)
			{
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
		}
		void
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
		}
		void
		SetMemberName(
			const CHAR* MemberName
			) override
		{
			m_MemberName = MemberName ? MemberName : std::string();
		}
		std::string
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
		}
		virtual
		void*
		GetSettings() override
		{
			return &m_Settings;
		}
	private:
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}



