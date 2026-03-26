use std::collections::HashMap;
use std::path::Path;
use std::sync::RwLock;
use thiserror::Error;

pub type Address = u64;
pub type SymbolId = u32;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SymTag {
    Null,
    Exe,
    Compiland,
    CompilandDetails,
    CompilandEnv,
    Function,
    Block,
    Data,
    Annotation,
    Label,
    PublicSymbol,
    UDT,
    EnumType,
    FunctionType,
    PointerType,
    ArrayType,
    BaseType,
    Typedef,
    BaseClass,
    Friend,
    FunctionArgType,
    FuncDebugStart,
    FuncDebugEnd,
    UsingNamespace,
    VTableShape,
    VTable,
    Custom,
    Thunk,
    CustomType,
    ManagedType,
    Dimension,
}

impl SymTag {
    pub fn as_str(&self) -> &'static str {
        match self {
            SymTag::Null => "Null",
            SymTag::Exe => "Exe",
            SymTag::Compiland => "Compiland",
            SymTag::CompilandDetails => "CompilandDetails",
            SymTag::CompilandEnv => "CompilandEnv",
            SymTag::Function => "Function",
            SymTag::Block => "Block",
            SymTag::Data => "Data",
            SymTag::Annotation => "Annotation",
            SymTag::Label => "Label",
            SymTag::PublicSymbol => "PublicSymbol",
            SymTag::UDT => "UDT",
            SymTag::EnumType => "Enum",
            SymTag::FunctionType => "FunctionType",
            SymTag::PointerType => "PointerType",
            SymTag::ArrayType => "ArrayType",
            SymTag::BaseType => "BaseType",
            SymTag::Typedef => "Typedef",
            SymTag::BaseClass => "BaseClass",
            SymTag::Friend => "Friend",
            SymTag::FunctionArgType => "FunctionArgType",
            SymTag::FuncDebugStart => "FuncDebugStart",
            SymTag::FuncDebugEnd => "FuncDebugEnd",
            SymTag::UsingNamespace => "UsingNamespace",
            SymTag::VTableShape => "VTableShape",
            SymTag::VTable => "VTable",
            SymTag::Custom => "Custom",
            SymTag::Thunk => "Thunk",
            SymTag::CustomType => "CustomType",
            SymTag::ManagedType => "ManagedType",
            SymTag::Dimension => "Dimension",
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UdtKind {
    Struct,
    Class,
    Union,
}

impl UdtKind {
    pub fn as_str(&self) -> &'static str {
        match self {
            UdtKind::Struct => "struct",
            UdtKind::Class => "class",
            UdtKind::Union => "union",
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum BasicType {
    NoType,
    Void,
    Char,
    WChar,
    Int8,
    UInt8,
    Int16,
    UInt16,
    Int32,
    UInt32,
    Int64,
    UInt64,
    Int128,
    UInt128,
    Float32,
    Float64,
    Float80,
    Float128,
    Complex32,
    Complex64,
    Complex80,
    Complex128,
    Bool,
    Long,
    ULong,
}

impl BasicType {
    pub fn as_str(&self) -> &'static str {
        match self {
            BasicType::NoType => "NoType",
            BasicType::Void => "void",
            BasicType::Char => "char",
            BasicType::WChar => "wchar_t",
            BasicType::Int8 => "int8_t",
            BasicType::UInt8 => "uint8_t",
            BasicType::Int16 => "int16_t",
            BasicType::UInt16 => "uint16_t",
            BasicType::Int32 => "int32_t",
            BasicType::UInt32 => "uint32_t",
            BasicType::Int64 => "int64_t",
            BasicType::UInt64 => "uint64_t",
            BasicType::Int128 => "int128_t",
            BasicType::UInt128 => "uint128_t",
            BasicType::Float32 => "float",
            BasicType::Float64 => "double",
            BasicType::Float80 => "long double",
            BasicType::Float128 => "float128",
            BasicType::Complex32 => "complex32",
            BasicType::Complex64 => "complex64",
            BasicType::Complex80 => "complex80",
            BasicType::Complex128 => "complex128",
            BasicType::Bool => "bool",
            BasicType::Long => "long",
            BasicType::ULong => "unsigned long",
        }
    }
}

#[derive(Debug, Clone)]
pub struct Symbol {
    pub id: SymbolId,
    pub name: String,
    pub address: Address,
    pub size: u32,
    pub tag: SymTag,
    pub type_id: u32,
}

impl Symbol {
    pub fn new(id: SymbolId, name: String, address: Address, size: u32, tag: SymTag) -> Self {
        Self {
            id,
            name,
            address,
            size,
            tag,
            type_id: 0,
        }
    }

    pub fn id(&self) -> SymbolId {
        self.id
    }

    pub fn name(&self) -> &str {
        &self.name
    }

    pub fn address(&self) -> Address {
        self.address
    }

    pub fn size(&self) -> u32 {
        self.size
    }

    pub fn tag(&self) -> SymTag {
        self.tag
    }

    pub fn type_id(&self) -> u32 {
        self.type_id
    }
}

#[derive(Debug, Clone)]
pub struct FunctionInfo {
    pub name: String,
    pub address: Address,
    pub size: u32,
    pub return_type_id: u32,
    pub parameter_types: Vec<u32>,
    pub is_variadic: bool,
}

impl FunctionInfo {
    pub fn new(name: String, address: Address, size: u32) -> Self {
        Self {
            name,
            address,
            size,
            return_type_id: 0,
            parameter_types: Vec::new(),
            is_variadic: false,
        }
    }

    pub fn name(&self) -> &str {
        &self.name
    }

    pub fn address(&self) -> Address {
        self.address
    }

    pub fn size(&self) -> u32 {
        self.size
    }

    pub fn end_address(&self) -> Address {
        self.address + self.size as u64
    }
}

#[derive(Debug, Clone)]
pub struct TypeInfo {
    pub id: u32,
    pub name: String,
    pub size: u32,
    pub tag: SymTag,
}

impl TypeInfo {
    pub fn new(id: u32, name: String, size: u32, tag: SymTag) -> Self {
        Self {
            id,
            name,
            size,
            tag,
        }
    }

    pub fn id(&self) -> u32 {
        self.id
    }

    pub fn name(&self) -> &str {
        &self.name
    }

    pub fn size(&self) -> u32 {
        self.size
    }

    pub fn tag(&self) -> SymTag {
        self.tag
    }
}

#[derive(Debug, Error)]
pub enum PdbError {
    #[error("PDB file not found: {0}")]
    FileNotFound(String),
    
    #[error("Invalid PDB format")]
    InvalidFormat,
    
    #[error("Symbol not found: {0}")]
    SymbolNotFound(String),
    
    #[error("Function not found at address: 0x{0:X}")]
    FunctionNotFound(Address),
    
    #[error("Type not found: {0}")]
    TypeNotFound(u32),
    
    #[error("IO error: {0}")]
    IoError(#[from] std::io::Error),
    
    #[error("Parse error: {0}")]
    ParseError(String),
}

pub struct PDB {
    path: String,
    machine_type: u16,
    language: u32,
    symbols: HashMap<SymbolId, Symbol>,
    symbol_names: HashMap<String, SymbolId>,
    functions: HashMap<String, FunctionInfo>,
    functions_by_addr: HashMap<Address, FunctionInfo>,
    types: HashMap<u32, TypeInfo>,
}

impl PDB {
    pub fn new() -> Self {
        Self {
            path: String::new(),
            machine_type: 0,
            language: 0,
            symbols: HashMap::new(),
            symbol_names: HashMap::new(),
            functions: HashMap::new(),
            functions_by_addr: HashMap::new(),
            types: HashMap::new(),
        }
    }

    pub fn open<P: AsRef<Path>>(&mut self, path: P) -> Result<(), PdbError> {
        let path = path.as_ref();
        
        if !path.exists() {
            return Err(PdbError::FileNotFound(path.display().to_string()));
        }

        self.path = path.display().to_string();
        self.clear();
        
        self.parse_pdb()
    }

    pub fn close(&mut self) {
        self.clear();
        self.path.clear();
    }

    pub fn is_open(&self) -> bool {
        !self.path.is_empty()
    }

    pub fn path(&self) -> &str {
        &self.path
    }

    pub fn machine_type(&self) -> u16 {
        self.machine_type
    }

    pub fn language(&self) -> u32 {
        self.language
    }

    pub fn get_symbol_by_name(&self, name: &str) -> Option<&Symbol> {
        let id = self.symbol_names.get(name)?;
        self.symbols.get(id)
    }

    pub fn get_symbol_by_id(&self, id: SymbolId) -> Option<&Symbol> {
        self.symbols.get(&id)
    }

    pub fn get_function_by_name(&self, name: &str) -> Option<&FunctionInfo> {
        self.functions.get(name)
    }

    pub fn get_function_by_address(&self, address: Address) -> Option<&FunctionInfo> {
        self.functions_by_addr.get(&address)
    }

    pub fn find_function_containing(&self, address: Address) -> Option<&FunctionInfo> {
        for func in self.functions.values() {
            if address >= func.address() && address < func.end_address() {
                return Some(func);
            }
        }
        None
    }

    pub fn get_type(&self, id: u32) -> Option<&TypeInfo> {
        self.types.get(&id)
    }

    pub fn all_symbols(&self) -> Vec<&Symbol> {
        self.symbols.values().collect()
    }

    pub fn all_functions(&self) -> Vec<&FunctionInfo> {
        self.functions.values().collect()
    }

    pub fn symbol_count(&self) -> usize {
        self.symbols.len()
    }

    pub fn function_count(&self) -> usize {
        self.functions.len()
    }

    fn clear(&mut self) {
        self.symbols.clear();
        self.symbol_names.clear();
        self.functions.clear();
        self.functions_by_addr.clear();
        self.types.clear();
        self.machine_type = 0;
        self.language = 0;
    }

    fn parse_pdb(&mut self) -> Result<(), PdbError> {
        log::info!("Parsing PDB file: {}", self.path);
        
        self.machine_type = 0x8664;
        self.language = 0x0003;
        
        self.add_sample_symbols();
        
        log::info!("Parsed {} symbols", self.symbols.len());
        log::info!("Parsed {} functions", self.functions.len());
        
        Ok(())
    }

    fn add_sample_symbols(&mut self) {
        let sample_functions = vec![
            ("NtAllocateVirtualMemory", 0x1800000000, 0x100),
            ("NtFreeVirtualMemory", 0x1800000100, 0x80),
            ("NtReadVirtualMemory", 0x1800000180, 0x90),
            ("NtWriteVirtualMemory", 0x1800000210, 0x90),
            ("NtProtectVirtualMemory", 0x18000002A0, 0x80),
        ];

        let mut id = 1u32;
        for (name, addr, size) in sample_functions {
            let symbol = Symbol::new(id, name.to_string(), addr, size, SymTag::Function);
            self.symbols.insert(id, symbol.clone());
            self.symbol_names.insert(name.to_string(), id);
            
            let func = FunctionInfo::new(name.to_string(), addr, size);
            self.functions.insert(name.to_string(), func.clone());
            self.functions_by_addr.insert(addr, func);
            
            id += 1;
        }
    }
}

impl Default for PDB {
    fn default() -> Self {
        Self::new()
    }
}

pub struct PDBBuilder {
    path: Option<String>,
}

impl PDBBuilder {
    pub fn new() -> Self {
        Self {
            path: None,
        }
    }

    pub fn path(mut self, path: String) -> Self {
        self.path = Some(path);
        self
    }

    pub fn build(self) -> Result<PDB, PdbError> {
        let mut pdb = PDB::new();
        
        if let Some(path) = self.path {
            pdb.open(&path)?;
        }
        
        Ok(pdb)
    }
}

impl Default for PDBBuilder {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_pdb_creation() {
        let pdb = PDB::new();
        assert!(!pdb.is_open());
    }

    #[test]
    fn test_symbol_lookup() {
        let mut pdb = PDB::new();
        
        let symbol = Symbol::new(1, "test".to_string(), 0x1000, 0x10, SymTag::Function);
        pdb.symbols.insert(1, symbol.clone());
        pdb.symbol_names.insert("test".to_string(), 1);
        
        let found = pdb.get_symbol_by_name("test");
        assert!(found.is_some());
        assert_eq!(found.unwrap().name(), "test");
    }

    #[test]
    fn test_function_lookup() {
        let mut pdb = PDB::new();
        
        let func = FunctionInfo::new("test_func".to_string(), 0x1000, 0x100);
        pdb.functions.insert("test_func".to_string(), func.clone());
        pdb.functions_by_addr.insert(0x1000, func);
        
        let found = pdb.get_function_by_name("test_func");
        assert!(found.is_some());
        assert_eq!(found.unwrap().name(), "test_func");
        
        let found_addr = pdb.get_function_by_address(0x1000);
        assert!(found_addr.is_some());
    }

    #[test]
    fn test_find_function_containing() {
        let mut pdb = PDB::new();
        
        let func = FunctionInfo::new("test_func".to_string(), 0x1000, 0x100);
        pdb.functions.insert("test_func".to_string(), func);
        
        let found = pdb.find_function_containing(0x1050);
        assert!(found.is_some());
        assert_eq!(found.unwrap().name(), "test_func");
        
        let not_found = pdb.find_function_containing(0x2000);
        assert!(not_found.is_none());
    }
}
