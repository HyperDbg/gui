package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\UdtFieldDefinition.h.back

type (
UdtFieldDefinition interface{
		VisitBaseType()(ok bool)//col:102
			if_()(ok bool)//col:200
			if_()(ok bool)//col:294
			if_()(ok bool)//col:384
			const_CHARPtr_BasicTypeString_=_PDB::GetBasicTypeString()(ok bool)//col:470
			m_TypePrefix_+=_()(ok bool)//col:555
		VisitPointerTypeEnd()(ok bool)//col:637
			if_()(ok bool)//col:715
			if_()(ok bool)//col:785
			if_()(ok bool)//col:851
		VisitArrayTypeEnd()(ok bool)//col:911
			if_()(ok bool)//col:967
				const_cast<SYMBOLPtr>()(ok bool)//col:1021
				m_TypeSuffix_+=_"["_+_std::to_string()(ok bool)//col:1069
		VisitFunctionTypeEnd()(ok bool)//col:1113
		SetMemberName()(ok bool)//col:1149
			m_MemberName_=_MemberName_?_MemberName_:_std::string()(ok bool)//col:1181
		GetPrintableDefinition()(ok bool)//col:1210
		SetSettings()(ok bool)//col:1234
			if_()(ok bool)//col:1253
			m_Settings_=_static_cast<SettingsPtr>()(ok bool)//col:1268
		GetSettings()(ok bool)//col:1279
}
)

func NewUdtFieldDefinition() { return & udtFieldDefinition{} }

func (u *udtFieldDefinition)		VisitBaseType()(ok bool){//col:102
/*		VisitBaseType(
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:200
/*			if (Symbol->BaseType == btFloat && Symbol->Size == 10)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:294
/*			if (Symbol->IsConst)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:384
/*			if (Symbol->IsVolatile)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			const_CHARPtr_BasicTypeString_=_PDB::GetBasicTypeString()(ok bool){//col:470
/*			const CHAR* BasicTypeString = PDB::GetBasicTypeString(Symbol, m_Settings->UseStdInt);
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			m_TypePrefix_+=_()(ok bool){//col:555
/*			m_TypePrefix += (BasicTypeString != nullptr ? BasicTypeString : "<unknown_type>");
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)		VisitPointerTypeEnd()(ok bool){//col:637
/*		VisitPointerTypeEnd(
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:715
/*			if (Symbol->u.Pointer.IsReference)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:785
/*			if (Symbol->IsConst)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:851
/*			if (Symbol->IsVolatile)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)		VisitArrayTypeEnd()(ok bool){//col:911
/*		VisitArrayTypeEnd(
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:967
/*			if (Symbol->u.Array.ElementCount == 0)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)				const_cast<SYMBOLPtr>()(ok bool){//col:1021
/*				const_cast<SYMBOL*>(Symbol)->Size = 1;
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)				m_TypeSuffix_+=_"["_+_std::to_string()(ok bool){//col:1069
/*				m_TypeSuffix += "[" + std::to_string(Symbol->u.Array.ElementCount) + "]";
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)		VisitFunctionTypeEnd()(ok bool){//col:1113
/*		VisitFunctionTypeEnd(
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)		SetMemberName()(ok bool){//col:1149
/*		SetMemberName(
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			m_MemberName_=_MemberName_?_MemberName_:_std::string()(ok bool){//col:1181
/*			m_MemberName = MemberName ? MemberName : std::string();
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)		GetPrintableDefinition()(ok bool){//col:1210
/*		GetPrintableDefinition() const override
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)		SetSettings()(ok bool){//col:1234
/*		SetSettings(
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			if_()(ok bool){//col:1253
/*			if (MemberDefinitionSettings == nullptr)
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
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}

func (u *udtFieldDefinition)			m_Settings_=_static_cast<SettingsPtr>()(ok bool){//col:1268
/*			m_Settings = static_cast<Settings*>(MemberDefinitionSettings);
		}
		virtual
		void*
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
};*/
return true
}

func (u *udtFieldDefinition)		GetSettings()(ok bool){//col:1279
/*		GetSettings() override
		{
			return &m_Settings;
		}
	private:
		std::string m_TypePrefix; 
		std::string m_MemberName; 
		std::string m_TypeSuffix; 
		std::string m_Comment;
		Settings* m_Settings;
};*/
return true
}



