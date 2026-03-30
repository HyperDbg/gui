//! Go script execution engine for HyperDbg hooks
//!
//! This module provides the ability to parse and execute Go code snippets
//! within the kernel driver using WDK bindings.
//!
//! # Architecture
//!
//! ```text
//! Go Code → Parser → AST Nodes → Analyzer → Hook Operations → Executor → WDK API
//! ```
//!
//! # Hook Installation Flow
//!
//! ```text
//! 1. Go code sent via HTTP → install_hook_script()
//! 2. Parse Go code → HookOperations
//! 3. Lookup API address from hook_db
//! 4. Install EPT/Inline hook via hooks.rs
//! 5. Store operations in SCRIPT_HOOK_DATABASE
//! 6. On hook trigger → execute_script_hook() → WDK bindings
//! ```

mod parser;
mod analyzer;
mod generator;
mod executor;

pub use parser::{GoParser, GoNode, parse_go_code};
pub use analyzer::{AstAnalyzer, HookOperation, analyze_go_code};
pub use generator::{RustGenerator, generate_rust_code, generate_hook_function};
pub use executor::ScriptExecutor;

use alloc::string::String;
use alloc::vec::Vec;
use alloc::collections::BTreeMap;
use spin::Mutex;
use crate::common::types_gen::HookFilter;
use crate::framework::log::LogLevel;
use crate::hyperkd::hyperhv::hooks::hooks::{
    HookError, HOOK_CONTEXT,
    install_ept_hook, install_inline_hook,
    set_ept_hook2_detour, EptHook2DetourEntry,
};
use crate::hyperkd::hyperhv::hooks::hook_db::{
    EPT_HOOK_DATABASE, EptHookDb,
    INLINE_HOOK_DATABASE, InlineHookDb,
};

pub const HOOK_TYPE_EPT: u32 = 0;
pub const HOOK_TYPE_INLINE: u32 = 1;

#[derive(Clone)]
pub struct ScriptHookEntry {
    pub api_name: alloc::string::String,
    pub address: u64,
    pub hook_type: u32,
    pub operations: Vec<HookOperation>,
    pub filter: Option<HookFilter>,
    pub enabled: bool,
}

pub static SCRIPT_HOOK_DATABASE: Mutex<BTreeMap<u64, ScriptHookEntry>> = Mutex::new(BTreeMap::new());

pub fn find_api_address(api_name: &str) -> Option<u64> {
    for entry in EPT_HOOK_DATABASE.iter() {
        if entry.name == api_name {
            return Some(entry.address);
        }
    }
    for entry in INLINE_HOOK_DATABASE.iter() {
        if entry.name == api_name {
            return Some(entry.address);
        }
    }
    None
}

pub fn find_api_address_dynamic(api_name: &str) -> Option<u64> {
    use crate::ntapi::exported::MmGetSystemRoutineAddress;
    use crate::ntapi::{UNICODE_STRING, PUNICODE_STRING};
    
    let mut name_buffer = [0u16; 256];
    let bytes = api_name.as_bytes();
    for (i, &b) in bytes.iter().enumerate() {
        if i >= 255 { break; }
        name_buffer[i] = b as u16;
    }
    name_buffer[bytes.len()] = 0;
    
    let mut unicode_name = UNICODE_STRING {
        Length: (bytes.len() * 2) as u16,
        MaximumLength: 256 * 2,
        Buffer: name_buffer.as_mut_ptr(),
    };
    
    unsafe {
        let addr = MmGetSystemRoutineAddress(&mut unicode_name as *mut _ as *mut _);
        if addr.is_null() {
            None
        } else {
            Some(addr as u64)
        }
    }
}

pub struct GoScriptEngine {
    parser: GoParser,
    analyzer: AstAnalyzer,
    generator: RustGenerator,
    executor: ScriptExecutor,
}

impl GoScriptEngine {
    pub fn new() -> Self {
        Self {
            parser: GoParser::new(),
            analyzer: AstAnalyzer::new(),
            generator: RustGenerator::new(),
            executor: ScriptExecutor::new(),
        }
    }

    pub fn compile(&mut self, go_code: &str) -> Option<Vec<HookOperation>> {
        let nodes = parse_go_code(go_code)?;
        let operations = analyze_go_code(&nodes);
        Some(operations)
    }

    pub fn generate_code(&mut self, operations: &[HookOperation]) -> String {
        generate_rust_code(operations)
    }

    pub unsafe fn execute(&mut self, operations: &[HookOperation], regs: &mut crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext) {
        self.executor.load(operations.to_vec());
        self.executor.execute(regs);
    }
}

impl Default for GoScriptEngine {
    fn default() -> Self {
        Self::new()
    }
}

pub fn compile_and_execute(go_code: &str, regs: &mut crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext) -> bool {
    let mut engine = GoScriptEngine::new();

    match engine.compile(go_code) {
        Some(operations) => {
            unsafe { engine.execute(&operations, regs); }
            true
        }
        None => false,
    }
}

pub fn install_hook_script(
    api_name: &str,
    hook_type: u32,
    filter: Option<&HookFilter>,
    code: &str,
) -> Result<(), &'static str> {
    crate::log!(LogLevel::Info, "Installing hook script for {}", api_name);
    
    let mut engine = GoScriptEngine::new();
    let operations = engine.compile(code).ok_or("Failed to parse Go code")?;
    
    let generated_code = engine.generate_code(&operations);
    crate::log!(LogLevel::Info, "Generated Rust code:\n{}", generated_code);
    
    let address = find_api_address(api_name)
        .or_else(|| find_api_address_dynamic(api_name))
        .ok_or("API not found in hook database")?;
    
    crate::log!(LogLevel::Info, "Found {} at address 0x{:X}", api_name, address);
    
    if address == 0 {
        return Err("API address is null, cannot hook");
    }
    
    let hook_handler = script_hook_dispatcher as u64;
    
    match hook_type {
        HOOK_TYPE_EPT => {
            unsafe {
                install_ept_hook(address, &[]).map_err(|e| {
                    crate::log!(LogLevel::Error, "EPT hook failed: {:?}", e);
                    "EPT hook installation failed"
                })?;
            }
            crate::log!(LogLevel::Info, "EPT hook installed at 0x{:X}", address);
        }
        HOOK_TYPE_INLINE => {
            unsafe {
                install_inline_hook(address, hook_handler).map_err(|e| {
                    crate::log!(LogLevel::Error, "Inline hook failed: {:?}", e);
                    "Inline hook installation failed"
                })?;
            }
            crate::log!(LogLevel::Info, "Inline hook installed at 0x{:X}", address);
        }
        _ => return Err("Invalid hook type"),
    }
    
    let entry = ScriptHookEntry {
        api_name: alloc::string::String::from(api_name),
        address,
        hook_type,
        operations,
        filter: filter.cloned(),
        enabled: true,
    };
    
    {
        let mut db = SCRIPT_HOOK_DATABASE.lock();
        db.insert(address, entry);
    }
    
    crate::log!(LogLevel::Info, "Hook script stored in database for {}", api_name);
    
    Ok(())
}

pub fn uninstall_hook_script(api_name: &str) -> Result<(), &'static str> {
    let address = {
        let db = SCRIPT_HOOK_DATABASE.lock();
        let entry = db.values().find(|e| e.api_name == api_name)
            .ok_or("Hook not found")?;
        entry.address
    };
    
    {
        let mut db = SCRIPT_HOOK_DATABASE.lock();
        db.remove(&address);
    }
    
    crate::log!(LogLevel::Info, "Hook script removed for {}", api_name);
    Ok(())
}

pub fn get_script_hook(address: u64) -> Option<ScriptHookEntry> {
    let db = SCRIPT_HOOK_DATABASE.lock();
    db.get(&address).cloned()
}

pub fn execute_script_hook(address: u64, regs: &mut crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext) -> bool {
    let operations = {
        let db = SCRIPT_HOOK_DATABASE.lock();
        match db.get(&address) {
            Some(entry) if entry.enabled => entry.operations.clone(),
            _ => return false,
        }
    };
    
    let mut executor = ScriptExecutor::new();
    unsafe {
        executor.load(operations);
        executor.execute(regs);
    }
    
    true
}

pub extern "C" fn script_hook_dispatcher(regs: *mut crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext) {
    if regs.is_null() {
        return;
    }
    
    unsafe {
        let regs_ref = &mut *regs;
        let called_address = regs_ref.rip;
        
        execute_script_hook(called_address, regs_ref);
    }
}

pub fn get_hook_count() -> usize {
    let db = SCRIPT_HOOK_DATABASE.lock();
    db.len()
}

pub fn list_hooked_apis() -> Vec<alloc::string::String> {
    let db = SCRIPT_HOOK_DATABASE.lock();
    db.values().map(|e| e.api_name.clone()).collect()
}

pub fn enable_hook(address: u64, enabled: bool) -> Result<(), &'static str> {
    let mut db = SCRIPT_HOOK_DATABASE.lock();
    let entry = db.get_mut(&address).ok_or("Hook not found")?;
    entry.enabled = enabled;
    Ok(())
}
