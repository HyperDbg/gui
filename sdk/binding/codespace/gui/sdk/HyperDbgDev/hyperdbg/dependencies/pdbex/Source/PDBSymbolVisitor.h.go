package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBSymbolVisitor.h.back

type (
PdbSymbolVisitor interface{
		PDBSymbolVisitor()(ok bool)//col:213
		Run()(ok bool)//col:422
		Visit()(ok bool)//col:626
		VisitBaseType()(ok bool)//col:826
		VisitEnumType()(ok bool)//col:1022
		VisitTypedefType()(ok bool)//col:1214
		VisitPointerType()(ok bool)//col:1402
		VisitArrayType()(ok bool)//col:1586
		VisitFunctionType()(ok bool)//col:1766
		VisitFunctionArgType()(ok bool)//col:1942
		VisitUdt()(ok bool)//col:2114
		VisitOtherType()(ok bool)//col:2282
		VisitEnumField()(ok bool)//col:2446
		VisitUdtField()(ok bool)//col:2606
		VisitUdtFieldEnd()(ok bool)//col:2762
		VisitUdtFieldBitFieldEnd()(ok bool)//col:2914
			AnonymousUdt()(ok bool)//col:3060
			BitFieldRange()(ok bool)//col:3182
				:_FirstUdtFieldBitField()(ok bool)//col:3303
				,_LastUdtFieldBitField()(ok bool)//col:3423
			Clear()(ok bool)//col:3539
			HasValue()(ok bool)//col:3649
			UdtFieldContext()(ok bool)//col:3751
				if_()(ok bool)//col:3840
					NextUdtField___=_GetNextUdtFieldWithRespectToBitFields()(ok bool)//col:3927
			IsFirst()(ok bool)//col:4010
			IsLast()(ok bool)//col:4088
			GetNext()(ok bool)//col:4161
				if_()(ok bool)//col:4229
}
)

func NewPdbSymbolVisitor() { return & pdbSymbolVisitor{} }

func (p *pdbSymbolVisitor)		PDBSymbolVisitor()(ok bool){//col:213
/*		PDBSymbolVisitor(
			PDBReconstructorBase* ReconstructVisitor,
			void* MemberDefinitionSettings = nullptr
			);
		void
		Run(
			const SYMBOL* Symbol
			);
	protected:
		void
		Visit(
			const SYMBOL* Symbol
			) override;
		void
		VisitBaseType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override;
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			) override;
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		Run()(ok bool){//col:422
/*		Run(
			const SYMBOL* Symbol
			);
	protected:
		void
		Visit(
			const SYMBOL* Symbol
			) override;
		void
		VisitBaseType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override;
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			) override;
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		Visit()(ok bool){//col:626
/*		Visit(
			const SYMBOL* Symbol
			) override;
		void
		VisitBaseType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override;
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			) override;
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitBaseType()(ok bool){//col:826
/*		VisitBaseType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override;
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			) override;
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitEnumType()(ok bool){//col:1022
/*		VisitEnumType(
			const SYMBOL* Symbol
			) override;
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			) override;
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitTypedefType()(ok bool){//col:1214
/*		VisitTypedefType(
			const SYMBOL* Symbol
			) override;
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitPointerType()(ok bool){//col:1402
/*		VisitPointerType(
			const SYMBOL* Symbol
			) override;
		void
		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitArrayType()(ok bool){//col:1586
/*		VisitArrayType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitFunctionType()(ok bool){//col:1766
/*		VisitFunctionType(
			const SYMBOL* Symbol
			) override;
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitFunctionArgType()(ok bool){//col:1942
/*		VisitFunctionArgType(
			const SYMBOL* Symbol
			) override;
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitUdt()(ok bool){//col:2114
/*		VisitUdt(
			const SYMBOL* Symbol
			) override;
		void
		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitOtherType()(ok bool){//col:2282
/*		VisitOtherType(
			const SYMBOL* Symbol
			) override;
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitEnumField()(ok bool){//col:2446
/*		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			) override;
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitUdtField()(ok bool){//col:2606
/*		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitUdtFieldEnd()(ok bool){//col:2762
/*		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)		VisitUdtFieldBitFieldEnd()(ok bool){//col:2914
/*		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			) override;
	private:
		struct AnonymousUdt
		{
			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			AnonymousUdt()(ok bool){//col:3060
/*			AnonymousUdt(
				UdtKind Kind,
				const SYMBOL_UDT_FIELD* FirstUdtField,
				const SYMBOL_UDT_FIELD* LastUdtField,
				DWORD Size = 0,
				DWORD MemberCount = 0
				)
			{
				this->Kind          = Kind;
				this->FirstUdtField = FirstUdtField;
				this->LastUdtField  = LastUdtField;
				this->Size          = Size;
				this->MemberCount   = MemberCount;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* LastUdtField;
			DWORD Size;
			DWORD MemberCount;
			UdtKind Kind;
		};
		struct BitFieldRange
		{
			const SYMBOL_UDT_FIELD* FirstUdtFieldBitField;
			const SYMBOL_UDT_FIELD* LastUdtFieldBitField;
			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			BitFieldRange()(ok bool){//col:3182
/*			BitFieldRange()
				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)				:_FirstUdtFieldBitField()(ok bool){//col:3303
/*				: FirstUdtFieldBitField(nullptr)
				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)				,_LastUdtFieldBitField()(ok bool){//col:3423
/*				, LastUdtFieldBitField(nullptr)
			{
			}
			void
			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			Clear()(ok bool){//col:3539
/*			Clear()
			{
				FirstUdtFieldBitField = nullptr;
				LastUdtFieldBitField = nullptr;
			}
			bool
			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			HasValue()(ok bool){//col:3649
/*			HasValue() const
			{
				return
				       LastUdtFieldBitField  != nullptr;
			}
		};
		struct UdtFieldContext
		{
			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			UdtFieldContext()(ok bool){//col:3751
/*			UdtFieldContext(
				const SYMBOL_UDT_FIELD* UdtField,
				BOOL RespectBitFields = TRUE
				)
			{
				SYMBOL_UDT* ParentUdt = &UdtField->Parent->u.Udt;
				DWORD UdtFieldCount = ParentUdt->FieldCount;
				FirstUdtField    = &ParentUdt->Fields[0];
				EndOfUdtField    = &ParentUdt->Fields[UdtFieldCount];
				PreviousUdtField = &UdtField[-1];
				CurrentUdtField  = &UdtField[ 0];
				NextUdtField     = &UdtField[ 1];
				this->RespectBitFields = RespectBitFields;
				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)				if_()(ok bool){//col:3840
/*				if (RespectBitFields)
				{
					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)					NextUdtField___=_GetNextUdtFieldWithRespectToBitFields()(ok bool){//col:3927
/*					NextUdtField   = GetNextUdtFieldWithRespectToBitFields(UdtField);
				}
			}
			bool
			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			IsFirst()(ok bool){//col:4010
/*			IsFirst() const
			{
				return PreviousUdtField < FirstUdtField;
			}
			bool
			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			IsLast()(ok bool){//col:4088
/*			IsLast() const
			{
				return NextUdtField == EndOfUdtField;
			}
			bool
			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)			GetNext()(ok bool){//col:4161
/*			GetNext()
			{
				PreviousUdtField = CurrentUdtField;
				CurrentUdtField  = NextUdtField;
				NextUdtField     = &CurrentUdtField[1];
				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}

func (p *pdbSymbolVisitor)				if_()(ok bool){//col:4229
/*				if (RespectBitFields && IsLast() == false)
				{
					NextUdtField = GetNextUdtFieldWithRespectToBitFields(CurrentUdtField);
				}
				return IsLast() == false;
			}
			const SYMBOL_UDT_FIELD* FirstUdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			const SYMBOL_UDT_FIELD* PreviousUdtField;
			const SYMBOL_UDT_FIELD* CurrentUdtField;
			const SYMBOL_UDT_FIELD* NextUdtField;
			BOOL RespectBitFields;
		};
		using AnonymousUdtStack = std::stack<std::shared_ptr<AnonymousUdt>>;
		using ContextStack      = std::stack<std::shared_ptr<UdtFieldDefinitionBase>>;
	private:
		void
		CheckForDataFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForBitFieldFieldPadding(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousUnion(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForAnonymousStruct(
			const SYMBOL_UDT_FIELD* UdtField
			);
		void
		CheckForEndOfAnonymousUdt(
			const SYMBOL_UDT_FIELD* UdtField
			);
		std::shared_ptr<UdtFieldDefinitionBase>
		MemberDefinitionFactory();
		void
		PushAnonymousUdt(
			std::shared_ptr<AnonymousUdt> Item
			);
		void
		PopAnonymousUdt();
	private:
		static
		const SYMBOL_UDT_FIELD*
		GetNextUdtFieldWithRespectToBitFields(
			const SYMBOL_UDT_FIELD* UdtField
			);
		static
		bool
		Is64BitBasicType(
			const SYMBOL* Symbol
			);
	private:
		DWORD m_SizeOfPreviousUdtField = 0;
		const SYMBOL_UDT_FIELD* m_PreviousUdtField = nullptr;
		const SYMBOL_UDT_FIELD* m_PreviousBitFieldField = nullptr;
		AnonymousUdtStack m_AnonymousUdtStack;
		AnonymousUdtStack m_AnonymousUnionStack;
		AnonymousUdtStack m_AnonymousStructStack;
		BitFieldRange m_CurrentBitField;
		ContextStack m_MemberContextStack;
		PDBReconstructorBase* m_ReconstructVisitor;
		void* m_MemberDefinitionSettings;
};*/
return true
}



