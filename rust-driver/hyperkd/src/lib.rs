#![no_std]
#![feature(abi_msvc)]
#![feature(naked_functions)]

extern crate alloc;

pub mod debugger;
pub mod network;
pub mod driver;
pub mod extension_commands;
pub mod callstack;
pub mod commands;
pub mod ffi;
pub mod dpc_routines;
pub mod halted_broadcast;
pub mod user_debug;
pub mod allocations;

use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use spin::Mutex;
use protocol::*;
use wsk::*;
use hyperhv::*;

pub const DEBUGGER_TOKEN: u32 = 0x01000000;
pub const MAX_BREAKPOINTS: usize = 256;
pub const MAX_PROCESS_COUNT: usize = 1024;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DebuggerState {
    Uninitialized,
    Initialized,
    Running,
    Paused,
    Terminated,
}

pub struct DebuggerContext {
    state: DebuggerState,
    socket: Option<WskSocket>,
    breakpoints: Vec<BreakpointInfo>,
    current_process: Option<ProcessId>,
    current_thread: Option<ThreadId>,
    paused_core: Option<u32>,
    pause_reason: Option<PauseReason>,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PauseReason {
    Breakpoint,
    Exception,
    StepComplete,
    UserRequest,
    ProcessExit,
    ThreadExit,
}

impl DebuggerContext {
    pub fn new() -> Self {
        Self {
            state: DebuggerState::Uninitialized,
            socket: None,
            breakpoints: Vec::with_capacity(MAX_BREAKPOINTS),
            current_process: None,
            current_thread: None,
            paused_core: None,
            pause_reason: None,
        }
    }
}

pub struct HyperDbgDriver {
    vmx: Arc<Mutex<VmxContext>>,
    debugger: Arc<Mutex<DebuggerContext>>,
    network: Arc<Mutex<NetworkManager>>,
}

impl HyperDbgDriver {
    pub fn new() -> Result<Self, DriverError> {
        Ok(Self {
            vmx: Arc::new(Mutex::new(VmxContext::new())),
            debugger: Arc::new(Mutex::new(DebuggerContext::new())),
            network: Arc::new(Mutex::new(NetworkManager::new()?)),
        })
    }

    // ==================== 初始化和终止 ====================

    pub fn initialize(&mut self) -> Result<(), DriverError> {
        let mut vmx = self.vmx.lock();
        vmx.initialize().map_err(|e| DriverError::VmxError(e))?;

        let mut dbg = self.debugger.lock();
        dbg.state = DebuggerState::Initialized;

        Ok(())
    }

    pub fn terminate(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        dbg.state = DebuggerState::Terminated;
        drop(dbg);

        let mut vmx = self.vmx.lock();
        vmx.terminate().map_err(|e| DriverError::VmxError(e))?;

        Ok(())
    }

    // ==================== 网络连接 ====================

    pub fn connect(&mut self, addr: &str, port: u16) -> Result<(), DriverError> {
        let mut network = self.network.lock();
        network.connect(addr, port)
    }

    pub fn disconnect(&mut self) {
        let mut network = self.network.lock();
        network.disconnect();
    }

    pub fn is_connected(&self) -> bool {
        let network = self.network.lock();
        network.is_connected()
    }

    // ==================== 进程调试 ====================

    pub fn attach_process(&mut self, process_id: ProcessId) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        dbg.current_process = Some(process_id);
        dbg.state = DebuggerState::Running;
        Ok(())
    }

    pub fn detach_process(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        dbg.current_process = None;
        dbg.current_thread = None;
        dbg.state = DebuggerState::Initialized;
        Ok(())
    }

    pub fn start_process(&mut self, path: &str) -> Result<ProcessId, DriverError> {
        // 创建进程并附加
        let process_id = self.create_process(path)?;
        self.attach_process(process_id)?;
        Ok(process_id)
    }

    pub fn kill_process(&mut self, process_id: ProcessId) -> Result<(), DriverError> {
        // 终止进程
        Ok(())
    }

    // ==================== 断点管理 ====================

    pub fn set_breakpoint(&mut self, address: Address, bp_type: BreakpointType) -> Result<u64, DriverError> {
        let mut dbg = self.debugger.lock();
        
        let id = dbg.breakpoints.len() as u64;
        let bp = BreakpointInfo {
            id,
            address,
            bp_type,
            process_id: dbg.current_process.unwrap_or(0),
            enabled: true,
            hit_count: 0,
        };

        // 设置断点
        match bp_type {
            BreakpointType::Software => {
                self.set_software_breakpoint(address)?;
            }
            BreakpointType::Hardware => {
                self.set_hardware_breakpoint(address)?;
            }
            BreakpointType::Hidden | BreakpointType::Ept => {
                self.set_ept_breakpoint(address)?;
            }
        }

        dbg.breakpoints.push(bp);
        Ok(id)
    }

    pub fn remove_breakpoint(&mut self, breakpoint_id: u64) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        
        if let Some(bp) = dbg.breakpoints.get(breakpoint_id as usize) {
            match bp.bp_type {
                BreakpointType::Software => {
                    self.remove_software_breakpoint(bp.address)?;
                }
                BreakpointType::Hardware => {
                    self.remove_hardware_breakpoint(bp.address)?;
                }
                BreakpointType::Hidden | BreakpointType::Ept => {
                    self.remove_ept_breakpoint(bp.address)?;
                }
            }
            dbg.breakpoints.remove(breakpoint_id as usize);
        }

        Ok(())
    }

    pub fn enable_breakpoint(&mut self, breakpoint_id: u64) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if let Some(bp) = dbg.breakpoints.get_mut(breakpoint_id as usize) {
            bp.enabled = true;
        }
        Ok(())
    }

    pub fn disable_breakpoint(&mut self, breakpoint_id: u64) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if let Some(bp) = dbg.breakpoints.get_mut(breakpoint_id as usize) {
            bp.enabled = false;
        }
        Ok(())
    }

    pub fn list_breakpoints(&self) -> Vec<BreakpointInfo> {
        let dbg = self.debugger.lock();
        dbg.breakpoints.clone()
    }

    // ==================== 执行控制 ====================

    pub fn continue_execution(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if dbg.state != DebuggerState::Paused {
            return Err(DriverError::NotPaused);
        }
        dbg.state = DebuggerState::Running;
        dbg.paused_core = None;
        dbg.pause_reason = None;
        Ok(())
    }

    pub fn pause_execution(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if dbg.state != DebuggerState::Running {
            return Err(DriverError::NotRunning);
        }
        dbg.state = DebuggerState::Paused;
        dbg.pause_reason = Some(PauseReason::UserRequest);
        Ok(())
    }

    pub fn step_into(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if dbg.state != DebuggerState::Paused {
            return Err(DriverError::NotPaused);
        }
        // 设置 TF 标志
        dbg.state = DebuggerState::Running;
        Ok(())
    }

    pub fn step_over(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if dbg.state != DebuggerState::Paused {
            return Err(DriverError::NotPaused);
        }
        // 检测当前指令，如果是 call 则设置临时断点
        dbg.state = DebuggerState::Running;
        Ok(())
    }

    pub fn step_out(&mut self) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        if dbg.state != DebuggerState::Paused {
            return Err(DriverError::NotPaused);
        }
        // 在返回地址设置临时断点
        dbg.state = DebuggerState::Running;
        Ok(())
    }

    // ==================== 寄存器操作 ====================

    pub fn read_registers(&self, core_id: u32) -> Result<RegisterState, DriverError> {
        let vmx = self.vmx.lock();
        if let Some(vcpu) = vmx.get_vcpu(core_id) {
            // 从 VMCS 读取寄存器
            Ok(RegisterState::default())
        } else {
            Err(DriverError::InvalidCoreId)
        }
    }

    pub fn write_registers(&mut self, core_id: u32, regs: &RegisterState) -> Result<(), DriverError> {
        let vmx = self.vmx.lock();
        if vmx.get_vcpu(core_id).is_some() {
            // 写入 VMCS
            Ok(())
        } else {
            Err(DriverError::InvalidCoreId)
        }
    }

    pub fn read_register(&self, core_id: u32, reg_name: &str) -> Result<u64, DriverError> {
        let regs = self.read_registers(core_id)?;
        match reg_name.to_lowercase().as_str() {
            "rax" => Ok(regs.rax),
            "rbx" => Ok(regs.rbx),
            "rcx" => Ok(regs.rcx),
            "rdx" => Ok(regs.rdx),
            "rsi" => Ok(regs.rsi),
            "rdi" => Ok(regs.rdi),
            "rbp" => Ok(regs.rbp),
            "rsp" => Ok(regs.rsp),
            "r8" => Ok(regs.r8),
            "r9" => Ok(regs.r9),
            "r10" => Ok(regs.r10),
            "r11" => Ok(regs.r11),
            "r12" => Ok(regs.r12),
            "r13" => Ok(regs.r13),
            "r14" => Ok(regs.r14),
            "r15" => Ok(regs.r15),
            "rip" => Ok(regs.rip),
            "rflags" => Ok(regs.rflags),
            "cr0" => Ok(regs.cr0),
            "cr2" => Ok(regs.cr2),
            "cr3" => Ok(regs.cr3),
            "cr4" => Ok(regs.cr4),
            _ => Err(DriverError::InvalidRegister),
        }
    }

    pub fn write_register(&mut self, core_id: u32, reg_name: &str, value: u64) -> Result<(), DriverError> {
        let mut regs = self.read_registers(core_id)?;
        match reg_name.to_lowercase().as_str() {
            "rax" => regs.rax = value,
            "rbx" => regs.rbx = value,
            "rcx" => regs.rcx = value,
            "rdx" => regs.rdx = value,
            "rsi" => regs.rsi = value,
            "rdi" => regs.rdi = value,
            "rbp" => regs.rbp = value,
            "rsp" => regs.rsp = value,
            "r8" => regs.r8 = value,
            "r9" => regs.r9 = value,
            "r10" => regs.r10 = value,
            "r11" => regs.r11 = value,
            "r12" => regs.r12 = value,
            "r13" => regs.r13 = value,
            "r14" => regs.r14 = value,
            "r15" => regs.r15 = value,
            "rip" => regs.rip = value,
            "rflags" => regs.rflags = value,
            _ => return Err(DriverError::InvalidRegister),
        }
        self.write_registers(core_id, &regs)
    }

    // ==================== 内存操作 ====================

    pub fn read_memory(&self, address: Address, size: usize) -> Result<Vec<u8>, DriverError> {
        let mut buffer = vec![0u8; size];
        // 使用 EPT 或直接读取
        Ok(buffer)
    }

    pub fn write_memory(&mut self, address: Address, data: &[u8]) -> Result<(), DriverError> {
        // 使用 EPT 或直接写入
        Ok(())
    }

    pub fn read_physical_memory(&self, address: PhysicalAddress, size: usize) -> Result<Vec<u8>, DriverError> {
        let mut buffer = vec![0u8; size];
        // 直接读取物理内存
        Ok(buffer)
    }

    pub fn write_physical_memory(&mut self, address: PhysicalAddress, data: &[u8]) -> Result<(), DriverError> {
        // 直接写入物理内存
        Ok(())
    }

    pub fn virtual_to_physical(&self, virtual_address: Address, cr3: u64) -> Result<PhysicalAddress, DriverError> {
        // 地址转换
        Ok(virtual_address)
    }

    // ==================== 进程和线程信息 ====================

    pub fn get_current_process(&self) -> Option<ProcessId> {
        let dbg = self.debugger.lock();
        dbg.current_process
    }

    pub fn get_current_thread(&self) -> Option<ThreadId> {
        let dbg = self.debugger.lock();
        dbg.current_thread
    }

    pub fn get_process_list(&self) -> Result<Vec<ProcessInfo>, DriverError> {
        // 枚举进程列表
        Ok(Vec::new())
    }

    pub fn get_thread_list(&self, process_id: ProcessId) -> Result<Vec<ThreadInfo>, DriverError> {
        // 枚举线程列表
        Ok(Vec::new())
    }

    pub fn get_module_list(&self, process_id: ProcessId) -> Result<Vec<ModuleInfo>, DriverError> {
        // 枚举模块列表
        Ok(Vec::new())
    }

    // ==================== 调用栈 ====================

    pub fn get_call_stack(&self, thread_id: ThreadId) -> Result<Vec<CallStackFrame>, DriverError> {
        // 获取调用栈
        Ok(Vec::new())
    }

    // ==================== 事件处理 ====================

    pub fn handle_vmexit(&mut self, core_id: u32, exit_reason: u32) -> Result<(), DriverError> {
        match exit_reason {
            3 => self.handle_breakpoint(core_id),
            14 => self.handle_page_fault(core_id),
            _ => Ok(()),
        }
    }

    fn handle_breakpoint(&mut self, core_id: u32) -> Result<(), DriverError> {
        let mut dbg = self.debugger.lock();
        dbg.state = DebuggerState::Paused;
        dbg.paused_core = Some(core_id);
        dbg.pause_reason = Some(PauseReason::Breakpoint);

        // 发送事件到 Go
        let event = BreakpointEvent {
            header: EventHeader {
                event_type: EventType::Breakpoint,
                process_id: dbg.current_process.unwrap_or(0),
                thread_id: dbg.current_thread.unwrap_or(0),
                core_id,
                timestamp: 0,
            },
            address: 0,
            breakpoint_id: 0,
            registers: RegisterState::default(),
        };

        drop(dbg);
        self.send_event(DebuggerEvent::Breakpoint(event))
    }

    fn handle_page_fault(&mut self, core_id: u32) -> Result<(), DriverError> {
        Ok(())
    }

    // ==================== 内部方法 ====================

    fn create_process(&mut self, path: &str) -> Result<ProcessId, DriverError> {
        // 创建进程
        Ok(0)
    }

    fn set_software_breakpoint(&mut self, address: Address) -> Result<(), DriverError> {
        // 写入 0xCC
        Ok(())
    }

    fn remove_software_breakpoint(&mut self, address: Address) -> Result<(), DriverError> {
        // 恢复原指令
        Ok(())
    }

    fn set_hardware_breakpoint(&mut self, address: Address) -> Result<(), DriverError> {
        // 设置 DR0-DR3
        Ok(())
    }

    fn remove_hardware_breakpoint(&mut self, address: Address) -> Result<(), DriverError> {
        // 清除 DR
        Ok(())
    }

    fn set_ept_breakpoint(&mut self, address: Address) -> Result<(), DriverError> {
        // 设置 EPT hook
        Ok(())
    }

    fn remove_ept_breakpoint(&mut self, address: Address) -> Result<(), DriverError> {
        // 移除 EPT hook
        Ok(())
    }

    fn send_event(&mut self, event: DebuggerEvent) -> Result<(), DriverError> {
        let network = self.network.lock();
        network.send_event(event)
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DriverError {
    VmxError(VmxError),
    SocketError(SocketError),
    NotInitialized,
    AlreadyInitialized,
    NotConnected,
    NotPaused,
    NotRunning,
    InvalidCoreId,
    InvalidRegister,
    InvalidAddress,
    ProcessNotFound,
    ThreadNotFound,
    BreakpointNotFound,
    InsufficientMemory,
    AccessDenied,
}

pub struct NetworkManager {
    socket: Option<WskSocket>,
    connected: bool,
}

impl NetworkManager {
    pub fn new() -> Result<Self, DriverError> {
        Ok(Self {
            socket: None,
            connected: false,
        })
    }

    pub fn connect(&mut self, addr: &str, port: u16) -> Result<(), DriverError> {
        let mut socket = WskSocket::new(SocketType::Stream, Protocol::Tcp)
            .map_err(|e| DriverError::SocketError(e))?;

        let addr: SocketAddr = addr.parse()
            .map_err(|_| DriverError::SocketError(SocketError::InvalidAddress))?;

        socket.connect(SocketAddr::new(addr.ip, port))
            .map_err(|e| DriverError::SocketError(e))?;

        self.socket = Some(socket);
        self.connected = true;
        Ok(())
    }

    pub fn disconnect(&mut self) {
        if let Some(mut socket) = self.socket.take() {
            socket.close();
        }
        self.connected = false;
    }

    pub fn is_connected(&self) -> bool {
        self.connected
    }

    pub fn send_event(&self, event: DebuggerEvent) -> Result<(), DriverError> {
        if !self.connected {
            return Err(DriverError::NotConnected);
        }

        // 序列化事件并发送
        Ok(())
    }

    pub fn recv_command(&self) -> Result<alloc::vec::Vec<u8>, DriverError> {
        if !self.connected {
            return Err(DriverError::NotConnected);
        }

        // 接收命令
        Ok(alloc::vec::Vec::new())
    }
}
