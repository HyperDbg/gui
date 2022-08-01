package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\UdtFieldDefinitionBase.h.back

type (
UdtFieldDefinitionBase interface{
		VisitBaseType()(ok bool)//col:88
		VisitPointerTypeBegin()(ok bool)//col:169
		VisitPointerTypeEnd()(ok bool)//col:243
		VisitArrayTypeBegin()(ok bool)//col:310
		VisitArrayTypeEnd()(ok bool)//col:370
		VisitFunctionTypeBegin()(ok bool)//col:423
		VisitFunctionTypeEnd()(ok bool)//col:469
		VisitFunctionArgTypeBegin()(ok bool)//col:508
		VisitFunctionArgTypeEnd()(ok bool)//col:540
		SetMemberName()(ok bool)//col:565
		GetPrintableDefinition()(ok bool)//col:583
			return_std::string()(ok bool)//col:599
		SetSettings()(ok bool)//col:611
		GetSettings()(ok bool)//col:616
}
)

func NewUdtFieldDefinitionBase() { return & udtFieldDefinitionBase{} }

func (u *udtFieldDefinitionBase)		VisitBaseType()(ok bool){//col:88
/*		VisitBaseType(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitPointerTypeBegin()(ok bool){//col:169
/*		VisitPointerTypeBegin(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitPointerTypeEnd()(ok bool){//col:243
/*		VisitPointerTypeEnd(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitArrayTypeBegin()(ok bool){//col:310
/*		VisitArrayTypeBegin(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitArrayTypeEnd()(ok bool){//col:370
/*		VisitArrayTypeEnd(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitFunctionTypeBegin()(ok bool){//col:423
/*		VisitFunctionTypeBegin(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitFunctionTypeEnd()(ok bool){//col:469
/*		VisitFunctionTypeEnd(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitFunctionArgTypeBegin()(ok bool){//col:508
/*		VisitFunctionArgTypeBegin(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		VisitFunctionArgTypeEnd()(ok bool){//col:540
/*		VisitFunctionArgTypeEnd(
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
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		SetMemberName()(ok bool){//col:565
/*		SetMemberName(
			const CHAR* MemberName
		)
		{
		}
		virtual
		std::string
		GetPrintableDefinition() const
		{
			return std::string();
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		GetPrintableDefinition()(ok bool){//col:583
/*		GetPrintableDefinition() const
		{
			return std::string();
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)			return_std::string()(ok bool){//col:599
/*			return std::string();
		}
		virtual
		void
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		SetSettings()(ok bool){//col:611
/*		SetSettings(
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
};*/
return true
}

func (u *udtFieldDefinitionBase)		GetSettings()(ok bool){//col:616
/*		GetSettings()
		{
			return nullptr;
		}
};*/
return true
}



