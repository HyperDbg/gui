use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use spin::Mutex;

use crate::hyperkd::hyperhv::state::*;
use crate::hyperkd::hyperhv::vmm::ept::*;
use crate::hyperkd::hyperhv::hooks::ept_hook::{ept_hook_add_hidden_breakpoint, ept_hook_remove, EptHookError};
use crate::hyperkd::hyperhv::common::msr::{read_msr, write_msr, IA32_LSTAR};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum HookError {
    AlreadyHooked,
    NotHooked,
    InvalidAddress,
    InsufficientMemory,
    EptError(EptHookError),
}

pub struct HookContext {
    pub ept_hooks: BTreeMap<Address, EptHookEntry>,
    pub syscall_hooks: BTreeMap<u32, SyscallHookEntry>,
    pub inline_hooks: BTreeMap<Address, InlineHookEntry>,
    pub ept_hook2_detours: BTreeMap<Address, EptHook2DetourEntry>,
    pub ept_hook2_initialized: bool,
}

pub struct EptHookEntry {
    pub address: Address,
    pub original_bytes: [u8; 16],
    pub hook_bytes: [u8; 16],
    pub hook_size: usize,
    pub shadow_page: PhysicalAddress,
    pub original_page: PhysicalAddress,
    pub process_id: ProcessId,
}

pub struct SyscallHookEntry {
    pub syscall_number: u32,
    pub original_handler: Address,
    pub hook_handler: Address,
    pub enabled: bool,
    pub is_entry: bool,
    pub original_msr_value: u64,
}

pub struct InlineHookEntry {
    pub address: Address,
    pub original_bytes: [u8; 14],
    pub hook_bytes: [u8; 14],
    pub hook_size: usize,
    pub enabled: bool,
    pub trampoline: Address,
    pub trampoline_size: usize,
}

#[derive(Clone, Copy)]
pub struct EptHook2DetourEntry {
    pub hooked_function_address: Address,
    pub return_address: Address,
    pub hook_function: Address,
    pub custom_code: Option<Address>,
    pub custom_code_size: usize,
    pub process_id: ProcessId,
    pub core_id: u32,
    pub enabled: bool,
}

pub struct SyscallHookManager {
    original_lstar: u64,
    hook_installed: bool,
    syscall_table: BTreeMap<u32, SyscallHookEntry>,
}

impl SyscallHookManager {
    pub fn new() -> Self {
        Self {
            original_lstar: 0,
            hook_installed: false,
            syscall_table: BTreeMap::new(),
        }
    }

    pub unsafe fn install_global_hook(&mut self, hook_handler: Address) -> Result<(), HookError> {
        if self.hook_installed {
            return Err(HookError::AlreadyHooked);
        }

        self.original_lstar = read_msr(IA32_LSTAR);

        write_msr(IA32_LSTAR, hook_handler);
        
        self.hook_installed = true;
        Ok(())
    }

    pub unsafe fn uninstall_global_hook(&mut self) -> Result<(), HookError> {
        if !self.hook_installed {
            return Err(HookError::NotHooked);
        }

        write_msr(IA32_LSTAR, self.original_lstar);
        
        self.hook_installed = false;
        Ok(())
    }

    pub fn add_syscall_filter(&mut self, syscall_number: u32, hook_handler: Address, is_entry: bool) -> Result<(), HookError> {
        if self.syscall_table.contains_key(&syscall_number) {
            return Err(HookError::AlreadyHooked);
        }

        let entry = SyscallHookEntry {
            syscall_number,
            original_handler: 0,
            hook_handler,
            enabled: true,
            is_entry,
            original_msr_value: 0,
        };

        self.syscall_table.insert(syscall_number, entry);
        Ok(())
    }

    pub fn remove_syscall_filter(&mut self, syscall_number: u32) -> Result<(), HookError> {
        self.syscall_table.remove(&syscall_number)
            .ok_or(HookError::NotHooked)?;
        Ok(())
    }

    pub fn is_hooked(&self, syscall_number: u32) -> bool {
        self.syscall_table.contains_key(&syscall_number)
    }

    pub fn get_hook(&self, syscall_number: u32) -> Option<&SyscallHookEntry> {
        self.syscall_table.get(&syscall_number)
    }

    pub fn get_original_lstar(&self) -> u64 {
        self.original_lstar
    }

    pub fn is_global_hook_installed(&self) -> bool {
        self.hook_installed
    }
}

pub static SYSCALL_HOOK_MANAGER: Mutex<SyscallHookManager> = Mutex::new(SyscallHookManager {
    original_lstar: 0,
    hook_installed: false,
    syscall_table: BTreeMap::new(),
});

impl HookContext {
    pub fn new() -> Self {
        Self {
            ept_hooks: BTreeMap::new(),
            syscall_hooks: BTreeMap::new(),
            inline_hooks: BTreeMap::new(),
            ept_hook2_detours: BTreeMap::new(),
            ept_hook2_initialized: false,
        }
    }

    pub fn set_ept_hook(
        &mut self,
        address: Address,
        _hook_handler: Address,
        process_id: ProcessId,
    ) -> Result<(), HookError> {
        if self.ept_hooks.contains_key(&address) {
            return Err(HookError::AlreadyHooked);
        }

        unsafe {
            ept_hook_add_hidden_breakpoint(address, process_id)
                .map_err(|e| HookError::EptError(e))?;
        }

        let physical_address = crate::memory::MemoryManager::virtual_to_physical(address);
        let entry = EptHookEntry {
            address,
            original_bytes: [0; 16],
            hook_bytes: [0; 16],
            hook_size: 0,
            shadow_page: 0,
            original_page: physical_address,
            process_id,
        };

        self.ept_hooks.insert(address, entry);
        Ok(())
    }

    pub fn remove_ept_hook(&mut self, address: Address) -> Result<(), HookError> {
        let entry = self.ept_hooks.remove(&address)
            .ok_or(HookError::NotHooked)?;

        unsafe {
            let _ = ept_hook_remove(address);
        }

        let _ = entry;
        Ok(())
    }

    pub fn set_syscall_hook(
        &mut self,
        syscall_number: u32,
        hook_handler: Address,
    ) -> Result<(), HookError> {
        if self.syscall_hooks.contains_key(&syscall_number) {
            return Err(HookError::AlreadyHooked);
        }

        let entry = SyscallHookEntry {
            syscall_number,
            original_handler: 0,
            hook_handler,
            enabled: true,
            is_entry: true,
            original_msr_value: 0,
        };

        self.syscall_hooks.insert(syscall_number, entry);
        Ok(())
    }

    pub fn set_syscall_hook_with_options(
        &mut self,
        syscall_number: u32,
        hook_handler: Address,
        is_entry: bool,
    ) -> Result<(), HookError> {
        if self.syscall_hooks.contains_key(&syscall_number) {
            return Err(HookError::AlreadyHooked);
        }

        let entry = SyscallHookEntry {
            syscall_number,
            original_handler: 0,
            hook_handler,
            enabled: true,
            is_entry,
            original_msr_value: 0,
        };

        self.syscall_hooks.insert(syscall_number, entry);
        Ok(())
    }

    pub fn enable_syscall_hook(&mut self, syscall_number: u32, enabled: bool) -> Result<(), HookError> {
        let entry = self.syscall_hooks.get_mut(&syscall_number)
            .ok_or(HookError::NotHooked)?;
        entry.enabled = enabled;
        Ok(())
    }

    pub fn set_inline_hook(
        &mut self,
        address: Address,
        hook_handler: Address,
    ) -> Result<(), HookError> {
        if self.inline_hooks.contains_key(&address) {
            return Err(HookError::AlreadyHooked);
        }

        if !is_valid_hook_address(address) {
            return Err(HookError::InvalidAddress);
        }

        let jump_code = build_absolute_jump(address, hook_handler);

        let mut original_bytes = [0u8; 14];
        unsafe {
            core::ptr::copy_nonoverlapping(
                address as *const u8,
                original_bytes.as_mut_ptr(),
                14,
            );
        }

        let trampoline = unsafe {
            crate::memory::MemoryManager::allocate_executable_pool(32)
                .map_err(|_| HookError::InsufficientMemory)?
        };

        unsafe {
            core::ptr::copy_nonoverlapping(
                original_bytes.as_ptr(),
                trampoline as *mut u8,
                14,
            );
            
            let jump_back = build_absolute_jump(trampoline + 14, address + 14);
            core::ptr::copy_nonoverlapping(
                jump_back.as_ptr(),
                (trampoline + 14) as *mut u8,
                14,
            );
        }

        let entry = InlineHookEntry {
            address,
            original_bytes,
            hook_bytes: jump_code,
            hook_size: 14,
            enabled: true,
            trampoline,
            trampoline_size: 28,
        };

        self.inline_hooks.insert(address, entry);
        Ok(())
    }

    pub fn set_inline_hook_with_trampoline(
        &mut self,
        address: Address,
        hook_handler: Address,
        hook_size: usize,
    ) -> Result<Address, HookError> {
        if self.inline_hooks.contains_key(&address) {
            return Err(HookError::AlreadyHooked);
        }

        if !is_valid_hook_address(address) {
            return Err(HookError::InvalidAddress);
        }

        let actual_hook_size = if hook_size < 5 { 5 } else { hook_size };
        let trampoline_size = actual_hook_size + 14;

        let trampoline = unsafe {
            crate::memory::MemoryManager::allocate_executable_pool(trampoline_size)
                .map_err(|_| HookError::InsufficientMemory)?
        };

        let mut original_bytes = [0u8; 14];
        unsafe {
            core::ptr::copy_nonoverlapping(
                address as *const u8,
                original_bytes.as_mut_ptr(),
                core::cmp::min(actual_hook_size, 14),
            );
            
            core::ptr::copy_nonoverlapping(
                original_bytes.as_ptr(),
                trampoline as *mut u8,
                actual_hook_size,
            );
            
            let jump_back = build_absolute_jump(trampoline + actual_hook_size as u64, address + actual_hook_size as u64);
            core::ptr::copy_nonoverlapping(
                jump_back.as_ptr(),
                (trampoline + actual_hook_size as u64) as *mut u8,
                14,
            );
        }

        let jump_code = build_absolute_jump(address, hook_handler);

        let entry = InlineHookEntry {
            address,
            original_bytes,
            hook_bytes: jump_code,
            hook_size: actual_hook_size,
            enabled: true,
            trampoline,
            trampoline_size,
        };

        self.inline_hooks.insert(address, entry);
        Ok(trampoline)
    }

    pub fn remove_inline_hook(&mut self, address: Address) -> Result<(), HookError> {
        let entry = self.inline_hooks.remove(&address)
            .ok_or(HookError::NotHooked)?;

        unsafe {
            core::ptr::copy_nonoverlapping(
                entry.original_bytes.as_ptr(),
                address as *mut u8,
                entry.hook_size,
            );

            if entry.trampoline != 0 {
                crate::memory::MemoryManager::free_pool(entry.trampoline);
            }
        }

        Ok(())
    }

    pub fn enable_inline_hook(&mut self, address: Address, enabled: bool) -> Result<(), HookError> {
        let entry = self.inline_hooks.get_mut(&address)
            .ok_or(HookError::NotHooked)?;
        
        if entry.enabled != enabled {
            entry.enabled = enabled;
            
            unsafe {
                if enabled {
                    core::ptr::copy_nonoverlapping(
                        entry.hook_bytes.as_ptr(),
                        address as *mut u8,
                        entry.hook_size,
                    );
                } else {
                    core::ptr::copy_nonoverlapping(
                        entry.original_bytes.as_ptr(),
                        address as *mut u8,
                        entry.hook_size,
                    );
                }
            }
        }
        
        Ok(())
    }

    pub fn get_trampoline(&self, address: Address) -> Option<Address> {
        self.inline_hooks.get(&address).map(|e| e.trampoline)
    }

    pub fn is_ept_hooked(&self, address: Address) -> bool {
        self.ept_hooks.contains_key(&address)
    }

    pub fn is_syscall_hooked(&self, syscall_number: u32) -> bool {
        self.syscall_hooks.contains_key(&syscall_number)
    }

    pub fn is_inline_hooked(&self, address: Address) -> bool {
        self.inline_hooks.contains_key(&address)
    }

    pub fn get_ept_hook(&self, address: Address) -> Option<&EptHookEntry> {
        self.ept_hooks.get(&address)
    }

    pub fn get_syscall_hook(&self, syscall_number: u32) -> Option<&SyscallHookEntry> {
        self.syscall_hooks.get(&syscall_number)
    }

    pub fn get_inline_hook(&self, address: Address) -> Option<&InlineHookEntry> {
        self.inline_hooks.get(&address)
    }
}

pub static HOOK_CONTEXT: Mutex<HookContext> = Mutex::new(HookContext {
    ept_hooks: BTreeMap::new(),
    syscall_hooks: BTreeMap::new(),
    inline_hooks: BTreeMap::new(),
    ept_hook2_detours: BTreeMap::new(),
    ept_hook2_initialized: false,
});

pub unsafe fn install_ept_hook(address: Address, _hook_code: &[u8]) -> Result<(), HookError> {
    let mut _ctx = HOOK_CONTEXT.lock();

    ept_hook_add_hidden_breakpoint(address, 0)
        .map_err(|e| HookError::EptError(e))?;

    Ok(())
}

pub unsafe fn uninstall_ept_hook(address: Address) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.remove_ept_hook(address)
}

pub unsafe fn install_syscall_hook(syscall_number: u32, hook_handler: Address) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.set_syscall_hook(syscall_number, hook_handler)
}

pub unsafe fn uninstall_syscall_hook(syscall_number: u32) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.syscall_hooks.remove(&syscall_number)
        .ok_or(HookError::NotHooked)?;
    Ok(())
}

pub unsafe fn enable_syscall_hook(syscall_number: u32, enabled: bool) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.enable_syscall_hook(syscall_number, enabled)
}

pub unsafe fn install_global_syscall_hook(hook_handler: Address) -> Result<(), HookError> {
    let mut manager = SYSCALL_HOOK_MANAGER.lock();
    manager.install_global_hook(hook_handler)
}

pub unsafe fn uninstall_global_syscall_hook() -> Result<(), HookError> {
    let mut manager = SYSCALL_HOOK_MANAGER.lock();
    manager.uninstall_global_hook()
}

pub unsafe fn add_syscall_filter(syscall_number: u32, hook_handler: Address, is_entry: bool) -> Result<(), HookError> {
    let mut manager = SYSCALL_HOOK_MANAGER.lock();
    manager.add_syscall_filter(syscall_number, hook_handler, is_entry)
}

pub unsafe fn remove_syscall_filter(syscall_number: u32) -> Result<(), HookError> {
    let mut manager = SYSCALL_HOOK_MANAGER.lock();
    manager.remove_syscall_filter(syscall_number)
}

pub fn is_syscall_filtered(syscall_number: u32) -> bool {
    let manager = SYSCALL_HOOK_MANAGER.lock();
    manager.is_hooked(syscall_number)
}

pub fn get_original_lstar() -> u64 {
    let manager = SYSCALL_HOOK_MANAGER.lock();
    manager.get_original_lstar()
}

pub unsafe fn install_inline_hook_with_trampoline(
    address: Address,
    hook_handler: Address,
    hook_size: usize,
) -> Result<Address, HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.set_inline_hook_with_trampoline(address, hook_handler, hook_size)
}

pub unsafe fn enable_inline_hook(address: Address, enabled: bool) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.enable_inline_hook(address, enabled)
}

pub fn get_inline_trampoline(address: Address) -> Option<Address> {
    let ctx = HOOK_CONTEXT.lock();
    ctx.get_trampoline(address)
}

pub unsafe fn install_inline_hook(address: Address, hook_handler: Address) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.set_inline_hook(address, hook_handler)
}

pub unsafe fn uninstall_inline_hook(address: Address) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.remove_inline_hook(address)
}

pub fn build_absolute_jump(from: Address, to: Address) -> [u8; 14] {
    let mut code = [0u8; 14];

    code[0] = 0xFF;
    code[1] = 0x25;
    code[2] = 0x00;
    code[3] = 0x00;
    code[4] = 0x00;
    code[5] = 0x00;

    let target = to.to_le_bytes();
    code[6..14].copy_from_slice(&target);

    code
}

pub fn build_relative_jump(from: Address, to: Address) -> [u8; 5] {
    let mut code = [0u8; 5];

    code[0] = 0xE9;

    let offset = (to as i64 - (from as i64 + 5)) as i32;
    let offset_bytes = offset.to_le_bytes();
    code[1..5].copy_from_slice(&offset_bytes);

    code
}

pub fn build_call(from: Address, to: Address) -> [u8; 5] {
    let mut code = [0u8; 5];

    code[0] = 0xE8;

    let offset = (to as i64 - (from as i64 + 5)) as i32;
    let offset_bytes = offset.to_le_bytes();
    code[1..5].copy_from_slice(&offset_bytes);

    code
}

pub fn build_int3() -> [u8; 1] {
    [0xCC]
}

pub fn build_nop() -> [u8; 1] {
    [0x90]
}

pub fn calculate_relative_offset(from: Address, to: Address, instruction_size: u32) -> i32 {
    (to as i64 - (from as i64 + instruction_size as i64)) as i32
}

pub fn is_valid_hook_address(address: Address) -> bool {
    address != 0 && crate::memory::MemoryManager::is_address_valid(address)
}

pub fn get_hook_count() -> (usize, usize, usize, usize) {
    let ctx = HOOK_CONTEXT.lock();
    (ctx.ept_hooks.len(), ctx.syscall_hooks.len(), ctx.inline_hooks.len(), ctx.ept_hook2_detours.len())
}

pub fn clear_all_hooks() {
    let mut ctx = HOOK_CONTEXT.lock();
    ctx.ept_hooks.clear();
    ctx.syscall_hooks.clear();
    ctx.inline_hooks.clear();
    ctx.ept_hook2_detours.clear();
    ctx.ept_hook2_initialized = false;
}

pub unsafe fn initialize_ept_hook2() -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    if ctx.ept_hook2_initialized {
        return Ok(());
    }
    ctx.ept_hook2_initialized = true;
    Ok(())
}

pub unsafe fn cleanup_ept_hook2() {
    let mut ctx = HOOK_CONTEXT.lock();
    for (address, _detour) in ctx.ept_hook2_detours.iter() {
        let _ = ept_hook_remove(*address);
    }
    ctx.ept_hook2_detours.clear();
    ctx.ept_hook2_initialized = false;
}

pub unsafe fn set_ept_hook2_detour(
    address: Address,
    hook_function: Address,
    custom_code: Option<Address>,
    custom_code_size: usize,
    process_id: ProcessId,
    core_id: u32,
) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    
    if !ctx.ept_hook2_initialized {
        return Err(HookError::InvalidAddress);
    }

    if ctx.ept_hook2_detours.contains_key(&address) {
        return Err(HookError::AlreadyHooked);
    }

    let physical_address = crate::memory::MemoryManager::virtual_to_physical(address);

    ept_hook_add_hidden_breakpoint(physical_address, process_id)
        .map_err(|e| HookError::EptError(e))?;

    let return_address = address + 0x10;

    let detour = EptHook2DetourEntry {
        hooked_function_address: address,
        return_address,
        hook_function,
        custom_code,
        custom_code_size,
        process_id,
        core_id,
        enabled: true,
    };

    ctx.ept_hook2_detours.insert(address, detour);
    Ok(())
}

pub unsafe fn remove_ept_hook2_detour(address: Address) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    
    let detour = ctx.ept_hook2_detours.remove(&address)
        .ok_or(HookError::NotHooked)?;

    let physical_address = crate::memory::MemoryManager::virtual_to_physical(address);
    let _ = ept_hook_remove(physical_address);

    Ok(())
}

pub unsafe fn ept_hook2_general_detour_handler(
    regs: *mut crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext,
    called_from: Address,
) -> Address {
    let ctx = HOOK_CONTEXT.lock();
    
    for (address, detour) in ctx.ept_hook2_detours.iter() {
        if detour.hooked_function_address == called_from && detour.enabled {
            return detour.return_address;
        }
    }

    called_from
}

pub fn is_ept_hook2_initialized() -> bool {
    let ctx = HOOK_CONTEXT.lock();
    ctx.ept_hook2_initialized
}

pub fn get_ept_hook2_count() -> usize {
    let ctx = HOOK_CONTEXT.lock();
    ctx.ept_hook2_detours.len()
}

pub fn find_ept_hook2_detour_by_address(address: Address) -> Option<EptHook2DetourEntry> {
    let ctx = HOOK_CONTEXT.lock();
    ctx.ept_hook2_detours.get(&address).copied()
}

pub fn enable_ept_hook2_detour(address: Address, enabled: bool) -> Result<(), HookError> {
    let mut ctx = HOOK_CONTEXT.lock();
    let detour = ctx.ept_hook2_detours.get_mut(&address)
        .ok_or(HookError::NotHooked)?;
    detour.enabled = enabled;
    Ok(())
}
