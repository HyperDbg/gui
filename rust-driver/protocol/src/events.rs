use alloc::string::String;
use alloc::vec::Vec;
use serde::{Deserialize, Serialize};
use crate::types::*;

#[derive(Clone, Copy, Debug, Serialize, Deserialize, PartialEq, Eq)]
#[repr(u32)]
pub enum EventType {
    Breakpoint = 1,
    Exception = 2,
    MemoryAccess = 3,
    SyscallEntry = 4,
    SyscallExit = 5,
    ProcessCreate = 6,
    ProcessExit = 7,
    ThreadCreate = 8,
    ThreadExit = 9,
    ModuleLoad = 10,
    ModuleUnload = 11,
    DebugPrint = 12,
    VmxExit = 13,
    Trap = 14,
    HiddenHookExec = 15,
    HiddenHookRead = 16,
    HiddenHookWrite = 17,
    Cpuid = 18,
    Tsc = 19,
    Pmc = 20,
    Interrupt = 21,
    ExceptionBitmap = 22,
    CrAccess = 23,
    DrAccess = 24,
    IoPort = 25,
    MsrRead = 26,
    MsrWrite = 27,
    EptViolation = 28,
    Vmcalled = 29,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct EventHeader {
    pub event_type: EventType,
    pub process_id: ProcessId,
    pub thread_id: ThreadId,
    pub core_id: u32,
    pub timestamp: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct BreakpointEvent {
    pub header: EventHeader,
    pub address: Address,
    pub breakpoint_id: u64,
    pub registers: RegisterState,
}

#[derive(Clone, Copy, Debug, Serialize, Deserialize, PartialEq, Eq)]
#[repr(u32)]
pub enum ExceptionCode {
    DivideError = 0,
    Debug = 1,
    Nmi = 2,
    Breakpoint = 3,
    Overflow = 4,
    Bound = 5,
    InvalidOpcode = 6,
    NoMath = 7,
    DoubleFault = 8,
    CoprocessorSegment = 9,
    InvalidTss = 10,
    SegmentNotPresent = 11,
    StackSegment = 12,
    GeneralProtection = 13,
    PageFault = 14,
    FloatingPoint = 16,
    AlignmentCheck = 17,
    MachineCheck = 18,
    SimdFloatingPoint = 19,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ExceptionEvent {
    pub header: EventHeader,
    pub exception_code: ExceptionCode,
    pub address: Address,
    pub error_code: u64,
    pub registers: RegisterState,
}

#[derive(Clone, Copy, Debug, Serialize, Deserialize, PartialEq, Eq)]
#[repr(u32)]
pub enum MemoryAccessType {
    Read = 0,
    Write = 1,
    Execute = 2,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct MemoryAccessEvent {
    pub header: EventHeader,
    pub virtual_address: Address,
    pub physical_address: PhysicalAddress,
    pub access_type: MemoryAccessType,
    pub size: u32,
    pub value: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct SyscallEvent {
    pub header: EventHeader,
    pub syscall_number: u32,
    pub arguments: [u64; 8],
    pub return_value: u64,
    pub is_entry: bool,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ProcessEvent {
    pub header: EventHeader,
    pub process_info: ProcessInfo,
    pub parent_process_id: ProcessId,
    pub is_create: bool,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ThreadEvent {
    pub header: EventHeader,
    pub thread_info: ThreadInfo,
    pub is_create: bool,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct ModuleEvent {
    pub header: EventHeader,
    pub module_info: ModuleInfo,
    pub is_load: bool,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct DebugPrintEvent {
    pub header: EventHeader,
    pub message: String,
    pub level: LogLevel,
}

#[derive(Clone, Copy, Debug, Serialize, Deserialize, PartialEq, Eq)]
#[repr(u32)]
pub enum VmxExitReason {
    ExceptionNmi = 0,
    ExternalInterrupt = 1,
    TripleFault = 2,
    InitSignal = 3,
    StartupIpi = 4,
    IoSmi = 5,
    OtherSmi = 6,
    InterruptWindow = 7,
    NmiWindow = 8,
    TaskSwitch = 9,
    Cpuid = 10,
    Getsec = 11,
    Hlt = 12,
    Invd = 13,
    Invlpg = 14,
    Rdpmc = 15,
    Rdtsc = 16,
    Rsm = 17,
    Vmcall = 18,
    Vmclear = 19,
    Vmlaunch = 20,
    Vmptrld = 21,
    Vmptrst = 22,
    Vmread = 23,
    Vmresume = 24,
    Vmwrite = 25,
    Vmxoff = 26,
    Vmxon = 27,
    CrAccess = 28,
    MovDr = 29,
    IoInstruction = 30,
    Rdmsr = 31,
    Wrmsr = 32,
    EntryFailGuestState = 33,
    EntryFailMsrLoad = 34,
    Mwait = 36,
    Mtf = 37,
    Monitor = 39,
    Pause = 40,
    EntryFailMachineCheck = 41,
    TprBelowThreshold = 43,
    ApicAccess = 44,
    VirtualizedEoi = 45,
    GdtrIdtrAccess = 46,
    LdtrTrAccess = 47,
    EptViolation = 48,
    EptMisconfig = 49,
    Invept = 50,
    Rdtscp = 51,
    VmxPreemptTimer = 52,
    Invvpid = 53,
    Wbinvd = 54,
    Xsetbv = 55,
    ApicWrite = 56,
    Rdrand = 57,
    Invpcid = 58,
    VmfFunc = 59,
    Encls = 60,
    Rdseed = 61,
    PmlFull = 62,
    Xsaves = 63,
    Xrstors = 64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct VmxExitEvent {
    pub header: EventHeader,
    pub exit_reason: VmxExitReason,
    pub exit_qualification: u64,
    pub guest_rip: Address,
    pub guest_rsp: Address,
    pub instruction_length: u32,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct TrapEvent {
    pub header: EventHeader,
    pub trap_number: u32,
    pub error_code: u64,
    pub address: Address,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct HiddenHookEvent {
    pub header: EventHeader,
    pub hook_address: Address,
    pub hook_type: MemoryAccessType,
    pub data: Vec<u8>,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct CpuidEvent {
    pub header: EventHeader,
    pub leaf: u32,
    pub sub_leaf: u32,
    pub eax: u32,
    pub ebx: u32,
    pub ecx: u32,
    pub edx: u32,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct TscEvent {
    pub header: EventHeader,
    pub tsc_value: u64,
    pub rdtsc_exit: bool,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct CrAccessEvent {
    pub header: EventHeader,
    pub cr_number: u32,
    pub is_write: bool,
    pub old_value: u64,
    pub new_value: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct DrAccessEvent {
    pub header: EventHeader,
    pub dr_number: u32,
    pub is_write: bool,
    pub value: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct IoPortEvent {
    pub header: EventHeader,
    pub port: u16,
    pub size: u32,
    pub is_write: bool,
    pub value: u32,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct MsrEvent {
    pub header: EventHeader,
    pub msr: u32,
    pub is_write: bool,
    pub value: u64,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct EptViolationEvent {
    pub header: EventHeader,
    pub guest_physical_address: PhysicalAddress,
    pub guest_virtual_address: Address,
    pub read: bool,
    pub write: bool,
    pub execute: bool,
    pub readable: bool,
    pub writable: bool,
    pub executable: bool,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum DebuggerEvent {
    Breakpoint(BreakpointEvent),
    Exception(ExceptionEvent),
    MemoryAccess(MemoryAccessEvent),
    Syscall(SyscallEvent),
    ProcessCreate(ProcessEvent),
    ProcessExit(ProcessEvent),
    ThreadCreate(ThreadEvent),
    ThreadExit(ThreadEvent),
    ModuleLoad(ModuleEvent),
    ModuleUnload(ModuleEvent),
    DebugPrint(DebugPrintEvent),
    VmxExit(VmxExitEvent),
    Trap(TrapEvent),
    HiddenHook(HiddenHookEvent),
    Cpuid(CpuidEvent),
    Tsc(TscEvent),
    CrAccess(CrAccessEvent),
    DrAccess(DrAccessEvent),
    IoPort(IoPortEvent),
    Msr(MsrEvent),
    EptViolation(EptViolationEvent),
}
