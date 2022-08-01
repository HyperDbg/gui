package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBSymbolSorterAlphabetical.h.back

type (
PdbSymbolSorterAlphabetical interface{
		GetSortedSymbols()(ok bool)//col:110
			if_()(ok bool)//col:218
				std::sort()(ok bool)//col:324
					m_SortedSymbols.begin()(ok bool)//col:429
					m_SortedSymbols.end()(ok bool)//col:533
					[]()(ok bool)//col:636
						return_strcmp()(ok bool)//col:738
		GetImageArchitecture()(ok bool)//col:833
		Clear()(ok bool)//col:923
			m_VisitedUdts.clear()(ok bool)//col:1010
			m_SortedSymbols.clear()(ok bool)//col:1096
		VisitEnumType()(ok bool)//col:1178
			if_()(ok bool)//col:1256
			AddSymbol()(ok bool)//col:1333
		VisitPointerType()(ok bool)//col:1407
			if_()(ok bool)//col:1477
				switch_()(ok bool)//col:1545
						assert()(ok bool)//col:1604
		VisitUdt()(ok bool)//col:1657
			if_()(ok bool)//col:1706
			PDBSymbolVisitorBase::VisitUdt()(ok bool)//col:1754
			AddSymbol()(ok bool)//col:1801
		VisitUdtField()(ok bool)//col:1845
			Visit()(ok bool)//col:1885
		HasBeenVisited()(ok bool)//col:1921
			if_()(ok bool)//col:1951
				if_()(ok bool)//col:1975
					Key_+=_std::to_string()(ok bool)//col:1997
		AddSymbol()(ok bool)//col:2012
			if_()(ok bool)//col:2023
}
)

func NewPdbSymbolSorterAlphabetical() { return & pdbSymbolSorterAlphabetical{} }

func (p *pdbSymbolSorterAlphabetical)		GetSortedSymbols()(ok bool){//col:110
/*		GetSortedSymbols() override
		{
			if (m_Dirty)
			{
				std::sort(
					m_SortedSymbols.begin(),
					m_SortedSymbols.end(),
					[](const SYMBOL* lhs, const SYMBOL* rhs) {
						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			if_()(ok bool){//col:218
/*			if (m_Dirty)
			{
				std::sort(
					m_SortedSymbols.begin(),
					m_SortedSymbols.end(),
					[](const SYMBOL* lhs, const SYMBOL* rhs) {
						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)				std::sort()(ok bool){//col:324
/*				std::sort(
					m_SortedSymbols.begin(),
					m_SortedSymbols.end(),
					[](const SYMBOL* lhs, const SYMBOL* rhs) {
						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)					m_SortedSymbols.begin()(ok bool){//col:429
/*					m_SortedSymbols.begin(),
					m_SortedSymbols.end(),
					[](const SYMBOL* lhs, const SYMBOL* rhs) {
						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)					m_SortedSymbols.end()(ok bool){//col:533
/*					m_SortedSymbols.end(),
					[](const SYMBOL* lhs, const SYMBOL* rhs) {
						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)					[]()(ok bool){//col:636
/*					[](const SYMBOL* lhs, const SYMBOL* rhs) {
						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)						return_strcmp()(ok bool){//col:738
/*						return strcmp(lhs->Name, rhs->Name) < 0;
					});
				m_Dirty = false;
			}
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		GetImageArchitecture()(ok bool){//col:833
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		Clear()(ok bool){//col:923
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			m_VisitedUdts.clear()(ok bool){//col:1010
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			m_SortedSymbols.clear()(ok bool){//col:1096
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		VisitEnumType()(ok bool){//col:1178
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			if_()(ok bool){//col:1256
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			AddSymbol()(ok bool){//col:1333
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		VisitPointerType()(ok bool){//col:1407
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			if_()(ok bool){//col:1477
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)				switch_()(ok bool){//col:1545
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)						assert()(ok bool){//col:1604
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		VisitUdt()(ok bool){//col:1657
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			if_()(ok bool){//col:1706
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			PDBSymbolVisitorBase::VisitUdt()(ok bool){//col:1754
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			AddSymbol()(ok bool){//col:1801
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		VisitUdtField()(ok bool){//col:1845
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			Visit()(ok bool){//col:1885
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		HasBeenVisited()(ok bool){//col:1921
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			if_()(ok bool){//col:1951
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)				if_()(ok bool){//col:1975
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)					Key_+=_std::to_string()(ok bool){//col:1997
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
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)		AddSymbol()(ok bool){//col:2012
/*		AddSymbol(
			const SYMBOL* Symbol
			)
		{
			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}

func (p *pdbSymbolSorterAlphabetical)			if_()(ok bool){//col:2023
/*			if (std::find(m_SortedSymbols.begin(), m_SortedSymbols.end(), Symbol) == m_SortedSymbols.end())
			{
				m_SortedSymbols.push_back(Symbol);
				m_Dirty = true;
			}
		}
		ImageArchitecture m_Architecture = ImageArchitecture::None;
		std::map<std::string, const SYMBOL*> m_VisitedUdts;
		std::vector<const SYMBOL*> m_SortedSymbols;
		bool m_Dirty = true;
};*/
return true
}



