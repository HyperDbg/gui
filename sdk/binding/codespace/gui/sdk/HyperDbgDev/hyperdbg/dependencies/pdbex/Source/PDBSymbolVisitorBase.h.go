package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBSymbolVisitorBase.h.back

type (
PdbSymbolVisitorBase interface{
		~PDBSymbolVisitorBase()(ok bool)//col:192
		Visit()(ok bool)//col:381
			switch_()(ok bool)//col:566
					VisitBaseType()(ok bool)//col:748
					VisitEnumType()(ok bool)//col:927
					VisitTypedefType()(ok bool)//col:1103
					VisitPointerType()(ok bool)//col:1276
					VisitArrayType()(ok bool)//col:1446
					VisitFunctionType()(ok bool)//col:1613
					VisitFunctionArgType()(ok bool)//col:1777
					VisitUdt()(ok bool)//col:1938
					VisitOtherType()(ok bool)//col:2096
		VisitBaseType()(ok bool)//col:2247
		VisitEnumType()(ok bool)//col:2391
			for_()(ok bool)//col:2531
				VisitEnumField()(ok bool)//col:2669
		VisitTypedefType()(ok bool)//col:2802
			Visit()(ok bool)//col:2931
		VisitPointerType()(ok bool)//col:3056
			Visit()(ok bool)//col:3177
		VisitArrayType()(ok bool)//col:3294
			Visit()(ok bool)//col:3407
		VisitFunctionType()(ok bool)//col:3516
			for_()(ok bool)//col:3621
				Visit()(ok bool)//col:3724
		VisitFunctionArgType()(ok bool)//col:3822
			Visit()(ok bool)//col:3916
		VisitUdt()(ok bool)//col:4006
			if_()(ok bool)//col:4090
				if_()(ok bool)//col:4166
					VisitUdtFieldBegin()(ok bool)//col:4240
					VisitUdtField()(ok bool)//col:4313
					VisitUdtFieldEnd()(ok bool)//col:4385
					VisitUdtFieldBitFieldBegin()(ok bool)//col:4453
						VisitUdtFieldBitField()(ok bool)//col:4518
					}_while_()(ok bool)//col:4582
					VisitUdtFieldBitFieldEnd()(ok bool)//col:4644
			}_while_()(ok bool)//col:4704
		VisitOtherType()(ok bool)//col:4760
		VisitEnumField()(ok bool)//col:4809
		VisitUdtFieldBegin()(ok bool)//col:4851
		VisitUdtFieldEnd()(ok bool)//col:4886
		VisitUdtField()(ok bool)//col:4914
		VisitUdtFieldBitFieldBegin()(ok bool)//col:4935
		VisitUdtFieldBitFieldEnd()(ok bool)//col:4949
		VisitUdtFieldBitField()(ok bool)//col:4956
}
)

func NewPdbSymbolVisitorBase() { return & pdbSymbolVisitorBase{} }

func (p *pdbSymbolVisitorBase)		~PDBSymbolVisitorBase()(ok bool){//col:192
/*		~PDBSymbolVisitorBase() = default;
		virtual
		void
		Visit(
			const SYMBOL* Symbol
			)
		{
			switch (Symbol->Tag)
			{
				case SymTagBaseType:
					VisitBaseType(Symbol);
					break;
				case SymTagEnum:
					VisitEnumType(Symbol);
					break;
				case SymTagTypedef:
					VisitTypedefType(Symbol);
					break;
				case SymTagPointerType:
					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		Visit()(ok bool){//col:381
/*		Visit(
			const SYMBOL* Symbol
			)
		{
			switch (Symbol->Tag)
			{
				case SymTagBaseType:
					VisitBaseType(Symbol);
					break;
				case SymTagEnum:
					VisitEnumType(Symbol);
					break;
				case SymTagTypedef:
					VisitTypedefType(Symbol);
					break;
				case SymTagPointerType:
					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			switch_()(ok bool){//col:566
/*			switch (Symbol->Tag)
			{
				case SymTagBaseType:
					VisitBaseType(Symbol);
					break;
				case SymTagEnum:
					VisitEnumType(Symbol);
					break;
				case SymTagTypedef:
					VisitTypedefType(Symbol);
					break;
				case SymTagPointerType:
					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitBaseType()(ok bool){//col:748
/*					VisitBaseType(Symbol);
					break;
				case SymTagEnum:
					VisitEnumType(Symbol);
					break;
				case SymTagTypedef:
					VisitTypedefType(Symbol);
					break;
				case SymTagPointerType:
					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitEnumType()(ok bool){//col:927
/*					VisitEnumType(Symbol);
					break;
				case SymTagTypedef:
					VisitTypedefType(Symbol);
					break;
				case SymTagPointerType:
					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitTypedefType()(ok bool){//col:1103
/*					VisitTypedefType(Symbol);
					break;
				case SymTagPointerType:
					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitPointerType()(ok bool){//col:1276
/*					VisitPointerType(Symbol);
					break;
				case SymTagArrayType:
					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitArrayType()(ok bool){//col:1446
/*					VisitArrayType(Symbol);
					break;
				case SymTagFunctionType:
					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitFunctionType()(ok bool){//col:1613
/*					VisitFunctionType(Symbol);
					break;
				case SymTagFunctionArgType:
					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitFunctionArgType()(ok bool){//col:1777
/*					VisitFunctionArgType(Symbol);
					break;
				case SymTagUDT:
					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitUdt()(ok bool){//col:1938
/*					VisitUdt(Symbol);
					break;
				default:
					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitOtherType()(ok bool){//col:2096
/*					VisitOtherType(Symbol);
					break;
			}
		}
	protected:
		virtual
		void
		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitBaseType()(ok bool){//col:2247
/*		VisitBaseType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitEnumType()(ok bool){//col:2391
/*		VisitEnumType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			for_()(ok bool){//col:2531
/*			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)				VisitEnumField()(ok bool){//col:2669
/*				VisitEnumField(&Symbol->u.Enum.Fields[i]);
			}
		}
		virtual
		void
		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitTypedefType()(ok bool){//col:2802
/*		VisitTypedefType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			Visit()(ok bool){//col:2931
/*			Visit(Symbol->u.Typedef.Type);
		}
		virtual
		void
		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitPointerType()(ok bool){//col:3056
/*		VisitPointerType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			Visit()(ok bool){//col:3177
/*			Visit(Symbol->u.Pointer.Type);
		}
		virtual
		void
		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitArrayType()(ok bool){//col:3294
/*		VisitArrayType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			Visit()(ok bool){//col:3407
/*			Visit(Symbol->u.Array.ElementType);
		}
		virtual
		void
		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitFunctionType()(ok bool){//col:3516
/*		VisitFunctionType(
			const SYMBOL* Symbol
			)
		{
			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			for_()(ok bool){//col:3621
/*			for (DWORD i = 0; i < Symbol->u.Function.ArgumentCount; i++)
			{
				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)				Visit()(ok bool){//col:3724
/*				Visit(Symbol->u.Function.Arguments[i]);
			}
		}
		virtual
		void
		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitFunctionArgType()(ok bool){//col:3822
/*		VisitFunctionArgType(
			const SYMBOL* Symbol
			)
		{
			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			Visit()(ok bool){//col:3916
/*			Visit(Symbol->u.FunctionArg.Type);
		}
		virtual
		void
		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdt()(ok bool){//col:4006
/*		VisitUdt(
			const SYMBOL* Symbol
			)
		{
			const SYMBOL_UDT_FIELD* UdtField;
			const SYMBOL_UDT_FIELD* EndOfUdtField;
			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			if_()(ok bool){//col:4090
/*			if (Symbol->u.Udt.FieldCount == 0)
			{
				return;
			}
			UdtField = Symbol->u.Udt.Fields;
			EndOfUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
			do
			{
				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)				if_()(ok bool){//col:4166
/*				if (UdtField->Bits == 0)
				{
					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitUdtFieldBegin()(ok bool){//col:4240
/*					VisitUdtFieldBegin(UdtField);
					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitUdtField()(ok bool){//col:4313
/*					VisitUdtField(UdtField);
					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitUdtFieldEnd()(ok bool){//col:4385
/*					VisitUdtFieldEnd(UdtField);
				}
				else
				{
					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitUdtFieldBitFieldBegin()(ok bool){//col:4453
/*					VisitUdtFieldBitFieldBegin(UdtField);
					do
					{
						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)						VisitUdtFieldBitField()(ok bool){//col:4518
/*						VisitUdtFieldBitField(UdtField);
					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					}_while_()(ok bool){//col:4582
/*					} while (++UdtField < EndOfUdtField &&
					           UdtField->BitPosition != 0);
					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)					VisitUdtFieldBitFieldEnd()(ok bool){//col:4644
/*					VisitUdtFieldBitFieldEnd(--UdtField);
				}
			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)			}_while_()(ok bool){//col:4704
/*			} while (++UdtField < EndOfUdtField);
		}
		virtual
		void
		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitOtherType()(ok bool){//col:4760
/*		VisitOtherType(
			const SYMBOL* Symbol
			)
		{
		}
		virtual
		void
		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitEnumField()(ok bool){//col:4809
/*		VisitEnumField(
			const SYMBOL_ENUM_FIELD* EnumField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdtFieldBegin()(ok bool){//col:4851
/*		VisitUdtFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdtFieldEnd()(ok bool){//col:4886
/*		VisitUdtFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdtField()(ok bool){//col:4914
/*		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdtFieldBitFieldBegin()(ok bool){//col:4935
/*		VisitUdtFieldBitFieldBegin(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdtFieldBitFieldEnd()(ok bool){//col:4949
/*		VisitUdtFieldBitFieldEnd(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
		}
		virtual
		void
		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}

func (p *pdbSymbolVisitorBase)		VisitUdtFieldBitField()(ok bool){//col:4956
/*		VisitUdtFieldBitField(
			const SYMBOL_UDT_FIELD* UdtField
			)
		{
			VisitUdtField(UdtField);
		}
};*/
return true
}



