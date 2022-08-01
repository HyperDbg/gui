package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBSymbolSorter.h.back

type (
PdbSymbolSorter interface{
		GetSortedSymbols()(ok bool)//col:98
		GetImageArchitecture()(ok bool)//col:191
		Clear()(ok bool)//col:279
			m_VisitedUdts.clear()(ok bool)//col:364
			m_SortedSymbols.clear()(ok bool)//col:448
		VisitEnumType()(ok bool)//col:528
			if_()(ok bool)//col:604
			AddSymbol()(ok bool)//col:679
		VisitPointerType()(ok bool)//col:751
			if_()(ok bool)//col:819
				switch_()(ok bool)//col:885
						assert()(ok bool)//col:942
		VisitUdt()(ok bool)//col:993
			if_()(ok bool)//col:1040
			PDBSymbolVisitorBase::VisitUdt()(ok bool)//col:1086
			AddSymbol()(ok bool)//col:1131
		VisitUdtField()(ok bool)//col:1173
			Visit()(ok bool)//col:1211
		HasBeenVisited()(ok bool)//col:1245
			if_()(ok bool)//col:1273
				if_()(ok bool)//col:1295
					Key_+=_std::to_string()(ok bool)//col:1315
		AddSymbol()(ok bool)//col:1328
			if_()(ok bool)//col:1337
}
pdbSymbolSorter struct{}
)

func NewPdbSymbolSorter()PdbSymbolSorter{ return & pdbSymbolSorter{} }

func (p *pdbSymbolSorter)		GetSortedSymbols()(ok bool){//col:98
/*		GetSortedSymbols() override
		{
			return m_SortedSymbols;
		}
		ImageArchitecture
		GetImageArchitecture() const override
		{
			return m_Architecture;
		}
		void
		Clear() override
		{
			ImageArchitecture m_Architecture = ImageArchitecture::None;
			m_VisitedUdts.clear();
			m_SortedSymbols.clear();
		}
	protected:
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		GetImageArchitecture()(ok bool){//col:191
/*		GetImageArchitecture() const override
		{
			return m_Architecture;
		}
		void
		Clear() override
		{
			ImageArchitecture m_Architecture = ImageArchitecture::None;
			m_VisitedUdts.clear();
			m_SortedSymbols.clear();
		}
	protected:
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		Clear()(ok bool){//col:279
/*		Clear() override
		{
			ImageArchitecture m_Architecture = ImageArchitecture::None;
			m_VisitedUdts.clear();
			m_SortedSymbols.clear();
		}
	protected:
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			m_VisitedUdts.clear()(ok bool){//col:364
/*			m_VisitedUdts.clear();
			m_SortedSymbols.clear();
		}
	protected:
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			m_SortedSymbols.clear()(ok bool){//col:448
/*			m_SortedSymbols.clear();
		}
	protected:
		void
		VisitEnumType(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		VisitEnumType()(ok bool){//col:528
/*		VisitEnumType(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			if_()(ok bool){//col:604
/*			if (HasBeenVisited(Symbol)) return;
			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			AddSymbol()(ok bool){//col:679
/*			AddSymbol(Symbol);
		}
		void
		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		VisitPointerType()(ok bool){//col:751
/*		VisitPointerType(
			const SYMBOL* Symbol
			) override
		{
			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			if_()(ok bool){//col:819
/*			if (m_Architecture == ImageArchitecture::None)
			{
				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)				switch_()(ok bool){//col:885
/*				switch (Symbol->Size)
				{
					case 4:
						m_Architecture = ImageArchitecture::x86;
						break;
					case 8:
						m_Architecture = ImageArchitecture::x64;
						break;
					default:
						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)						assert()(ok bool){//col:942
/*						assert(0);
						break;
				}
			}
		}
		void
		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		VisitUdt()(ok bool){//col:993
/*		VisitUdt(
			const SYMBOL* Symbol
			) override
		{
			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			if_()(ok bool){//col:1040
/*			if (HasBeenVisited(Symbol)) return;
			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			PDBSymbolVisitorBase::VisitUdt()(ok bool){//col:1086
/*			PDBSymbolVisitorBase::VisitUdt(Symbol);
			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			AddSymbol()(ok bool){//col:1131
/*			AddSymbol(Symbol);
		}
		void
		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		VisitUdtField()(ok bool){//col:1173
/*		VisitUdtField(
			const SYMBOL_UDT_FIELD* UdtField
			) override
		{
			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			Visit()(ok bool){//col:1211
/*			Visit(UdtField->Type);
		}
	private:
		bool
		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		HasBeenVisited()(ok bool){//col:1245
/*		HasBeenVisited(
			const SYMBOL* Symbol
			)
		{
			static DWORD UnnamedCounter = 0;
			std::string Key = Symbol->Name;
			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			if_()(ok bool){//col:1273
/*			if (m_VisitedUdts.find(Key) != m_VisitedUdts.end())
			{
				return true;
			}
			else
			{
				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)				if_()(ok bool){//col:1295
/*				if (PDB::IsUnnamedSymbol(Symbol))
				{
					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)					Key_+=_std::to_string()(ok bool){//col:1315
/*					Key += std::to_string(++UnnamedCounter);
				}
				m_VisitedUdts[Key] = Symbol;
				return false;
			}
		}
		void
		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)		AddSymbol()(ok bool){//col:1328
/*		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}

func (p *pdbSymbolSorter)			if_()(ok bool){//col:1337
/*			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
};*/
return true
}



