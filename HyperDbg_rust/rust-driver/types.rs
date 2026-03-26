use alloc::string::String;
use serde::{Deserialize, Serialize};

pub type ProcessId = u32;
pub type ThreadId = u32;
pub type Address = u64;
pub type PhysicalAddress = u64;

#[derive(Clone, Copy, Debug, Serialize, Deserialize, PartialEq, Eq)]
pub enum DebugState {
    Running = 0,
    Paused = 1,
    Stepping = 2,
    Terminated = 3,
}

#[derive(Clone, Copy, Debug, Serialize, Deserialize, PartialEq, Eq)]
pub enum BreakpointType {
    Software = 0,
    Hardware = 1,
    Hidden = 2,
    Ept = 3,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct RegisterState {
    pub rax: u64,
    pub rbx: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub rbp: u64,
    pub rsp: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
    pub rip: u64,
    pub rflags: u64,
    pub cr0: u64,
    pub cr2: u64,
    pub cr3: u64,
    pub cr4: u64,
    pub cr8: u64,
    pub dr0: u64,
    pub dr1: u64,
    pub dr2: u64,
    pub dr3: u64,
    pub dr6: u64,
    pub dr7: u64,
    pub gdtr: u64,
    pub gs_base: u64,
    pub fs_base: u64,
}

impl Default for RegisterState {
    fn default() -> Self {
        Self {
            rax: 0, rbx: 0, rcx: 0, rdx: 0,
            rsi: 0, rdi: 0, rbp: 0, rsp: 0,
            r8: 0, r9: 0, r10: 0, r11: 0,
            r12: 0, r13: 0, r14: 0, r15: 0,
            rip: 0, rflags: 0,
            cr0: 0, cr2: 0, cr3: 0, cr4: 0, cr8: 0,
            dr0: 0, dr1: 0, dr2: 0, dr3: 0, dr6: 0, dr7: 0,
            gdtr: 0, gs_base: 0, fs_base: 0,
        }
    }
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ProcessInfo {
    pub process_id: ProcessId,
    pub image_name: String,
    pub base_address: Address,
    pub size: u64,
    pub cr3: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ThreadInfo {
    pub thread_id: ThreadId,
    pub process_id: ProcessId,
    pub start_address: Address,
    pub teb: Address,
    pub state: DebugState,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct MemoryRegion {
    pub base_address: Address,
    pub size: u64,
    pub protection: u32,
    pub state: u32,
    pub type_: u32,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ModuleInfo {
    pub base_address: Address,
    pub size: u64,
    pub name: String,
    pub path: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct BreakpointInfo {
    pub id: u64,
    pub address: Address,
    pub bp_type: BreakpointType,
    pub process_id: ProcessId,
    pub enabled: bool,
    pub hit_count: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct CallStackFrame {
    pub instruction_pointer: Address,
    pub return_address: Address,
    pub stack_pointer: Address,
    pub frame_pointer: Address,
    pub module_name: String,
    pub symbol_name: Option<String>,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct SymbolInfo {
    pub name: String,
    pub address: Address,
    pub size: u64,
    pub module: String,
}

#[derive(Clone, Copy, Debug, Serialize, Deserialize)]
pub enum LogLevel {
    Trace = 0,
    Debug = 1,
    Info = 2,
    Warn = 3,
    Error = 4,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct VmxCapabilities {
    pub vmx_supported: bool,
    pub ept_supported: bool,
    pub vpid_supported: bool,
    pub msr_bitmap_supported: bool,
    pub io_bitmap_supported: bool,
    pub max_physical_address_width: u8,
    pub processor_count: u32,
}
