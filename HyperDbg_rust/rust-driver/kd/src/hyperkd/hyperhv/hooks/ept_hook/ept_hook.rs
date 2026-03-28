use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use spin::Mutex;

use crate::hyperkd::hyperhv::vmm::ept::*;
use crate::hyperkd::hyperhv::memory::MemoryManager;
use crate::hyperkd::hyperhv::state::EPT_PML1_ENTRY;

pub type EptPtEntry = EPT_PML1_ENTRY;
pub type EptPml1Entry = EPT_PML1_ENTRY;

pub const PAGE_SIZE: usize = 0x1000;
pub const SIZE_2_MB: u64 = 0x200000;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EptHookError {
    AllocationFailed,
    InvalidAddress,
    NotSupported,
    AlreadyHooked,
    NotHooked,
    PageSplitFailed,
    EptViolation,
}

#[derive(Debug, Clone, Copy)]
pub struct EptHooksContext {
    pub virtual_address: u64,
    pub physical_address: u64,
    pub process_id: u32,
    pub thread_id: u32,
    pub is_hidden_breakpoint: bool,
    pub is_execution_hook: bool,
}

impl Default for EptHooksContext {
    fn default() -> Self {
        Self {
            virtual_address: 0,
            physical_address: 0,
            process_id: 0,
            thread_id: 0,
            is_hidden_breakpoint: false,
            is_execution_hook: false,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum MtfTrapType {
    None = 0,
    EptHookRestore = 1,
    HiddenBreakpoint = 2,
    ModeBasedExec = 3,
}

#[derive(Debug, Clone)]
pub struct MtfTrapState {
    pub trap_type: MtfTrapType,
    pub hooked_entry: Option<usize>,
    pub enabled: bool,
}

impl Default for MtfTrapState {
    fn default() -> Self {
        Self {
            trap_type: MtfTrapType::None,
            hooked_entry: None,
            enabled: false,
        }
    }
}

#[repr(C, align(4096))]
pub struct EptHookedPageDetail {
    pub fake_page_contents: [u8; PAGE_SIZE],
    pub virtual_address: u64,
    pub physical_base_address: u64,
    pub physical_base_address_of_fake_page_contents: u64,
    pub original_entry: EptPtEntry,
    pub changed_entry: EptPtEntry,
    pub trampoline: u64,
    pub is_execution_hook: bool,
    pub is_hidden_breakpoint: bool,
    pub last_context_state: EptHooksContext,
    pub process_id: u32,
    pub hook_offset: u32,
    pub hook_size: u32,
    pub enabled: bool,
}

impl EptHookedPageDetail {
    pub fn new() -> Self {
        Self {
            fake_page_contents: [0u8; PAGE_SIZE],
            virtual_address: 0,
            physical_base_address: 0,
            physical_base_address_of_fake_page_contents: 0,
            original_entry: EptPml1Entry::new(),
            changed_entry: EptPml1Entry::new(),
            trampoline: 0,
            is_execution_hook: false,
            is_hidden_breakpoint: false,
            last_context_state: EptHooksContext::default(),
            process_id: 0,
            hook_offset: 0,
            hook_size: 0,
            enabled: false,
        }
    }

    pub fn set_fake_page_content(&mut self, offset: usize, data: &[u8]) {
        if offset + data.len() <= PAGE_SIZE {
            self.fake_page_contents[offset..offset + data.len()].copy_from_slice(data);
        }
    }

    pub fn get_fake_page_content(&self, offset: usize, size: usize) -> Option<&[u8]> {
        if offset + size <= PAGE_SIZE {
            Some(&self.fake_page_contents[offset..offset + size])
        } else {
            None
        }
    }
}

pub struct EptHookContext {
    pub hooked_pages: BTreeMap<u64, Arc<Mutex<EptHookedPageDetail>>>,
    pub hook_count: usize,
}

impl EptHookContext {
    pub const fn new() -> Self {
        Self {
            hooked_pages: BTreeMap::new(),
            hook_count: 0,
        }
    }

    pub fn add_hook(&mut self, physical_address: u64, detail: EptHookedPageDetail) -> Result<(), EptHookError> {
        if self.hooked_pages.contains_key(&physical_address) {
            return Err(EptHookError::AlreadyHooked);
        }

        self.hooked_pages.insert(physical_address, Arc::new(Mutex::new(detail)));
        self.hook_count += 1;
        Ok(())
    }

    pub fn remove_hook(&mut self, physical_address: u64) -> Result<(), EptHookError> {
        if self.hooked_pages.remove(&physical_address).is_none() {
            return Err(EptHookError::NotHooked);
        }
        self.hook_count -= 1;
        Ok(())
    }

    pub fn get_hook(&self, physical_address: u64) -> Option<Arc<Mutex<EptHookedPageDetail>>> {
        self.hooked_pages.get(&physical_address).cloned()
    }

    pub fn get_hook_count(&self) -> usize {
        self.hook_count
    }

    pub fn is_hooked(&self, physical_address: u64) -> bool {
        self.hooked_pages.contains_key(&physical_address)
    }

    pub fn clear_all_hooks(&mut self) {
        self.hooked_pages.clear();
        self.hook_count = 0;
    }
}

pub static EPT_HOOK_CONTEXT: Mutex<EptHookContext> = Mutex::new(EptHookContext::new());

pub struct EptHookManager;

impl EptHookManager {
    pub unsafe fn perform_page_hook_monitor_and_inline_hook(
        virtual_address: u64,
        process_id: u32,
        hook_handler: Option<u64>,
        is_hidden_breakpoint: bool,
    ) -> Result<(), EptHookError> {
        let physical_address = MemoryManager::virtual_to_physical(virtual_address);
        if physical_address == 0 {
            return Err(EptHookError::InvalidAddress);
        }

        let page_offset = (virtual_address & 0xFFF) as usize;
        let page_aligned_pa = physical_address & !0xFFF;

        {
            let context = EPT_HOOK_CONTEXT.lock();
            if context.is_hooked(page_aligned_pa) {
                return Err(EptHookError::AlreadyHooked);
            }
        }

        let mut hooked_detail = EptHookedPageDetail::new();
        hooked_detail.virtual_address = virtual_address;
        hooked_detail.physical_base_address = page_aligned_pa;
        hooked_detail.process_id = process_id;
        hooked_detail.hook_offset = page_offset as u32;

        if let Ok(original_page) = MemoryManager::map_physical_memory(page_aligned_pa, PAGE_SIZE, 0) {
            core::ptr::copy_nonoverlapping(
                original_page as *const u8,
                hooked_detail.fake_page_contents.as_mut_ptr(),
                PAGE_SIZE,
            );
            MemoryManager::unmap_physical_memory(original_page, PAGE_SIZE);
        } else {
            return Err(EptHookError::AllocationFailed);
        }

        if is_hidden_breakpoint {
            hooked_detail.fake_page_contents[page_offset] = 0xCC;
            hooked_detail.is_hidden_breakpoint = true;
        } else if let Some(handler) = hook_handler {
            let jump_code = Self::build_absolute_jump(handler);
            if page_offset + jump_code.len() <= PAGE_SIZE {
                hooked_detail.fake_page_contents[page_offset..page_offset + jump_code.len()]
                    .copy_from_slice(&jump_code);
                hooked_detail.hook_size = jump_code.len() as u32;
            }
            hooked_detail.is_execution_hook = true;
        }

        let fake_page_va = MemoryManager::allocate_contiguous(PAGE_SIZE)
            .map_err(|_| EptHookError::AllocationFailed)?;

        core::ptr::copy_nonoverlapping(
            hooked_detail.fake_page_contents.as_ptr(),
            fake_page_va as *mut u8,
            PAGE_SIZE,
        );

        hooked_detail.physical_base_address_of_fake_page_contents = 
            MemoryManager::virtual_to_physical(fake_page_va);

        hooked_detail.original_entry = Self::get_current_ept_entry(page_aligned_pa);
        
        let mut changed_entry = EptPml1Entry::new();
        changed_entry.set_read_access(true);
        changed_entry.set_write_access(true);
        changed_entry.set_execute_access(false);
        changed_entry.set_memory_type(hooked_detail.original_entry.memory_type());
        changed_entry.set_ignore_pat(hooked_detail.original_entry.ignore_pat());
        changed_entry.set_page_frame_number(hooked_detail.physical_base_address_of_fake_page_contents / PAGE_SIZE as u64);
        hooked_detail.changed_entry = changed_entry;

        hooked_detail.enabled = true;

        let changed_entry = hooked_detail.changed_entry;

        let mut context = EPT_HOOK_CONTEXT.lock();
        context.add_hook(page_aligned_pa, hooked_detail)?;

        Self::set_ept_entry_and_invalidate_tlb(
            page_aligned_pa,
            &changed_entry,
        );

        Ok(())
    }

    pub unsafe fn remove_hook(physical_address: u64) -> Result<(), EptHookError> {
        let page_aligned_pa = physical_address & !0xFFF;

        let mut context = EPT_HOOK_CONTEXT.lock();
        let hook = context.get_hook(page_aligned_pa)
            .ok_or(EptHookError::NotHooked)?;

        let hook_detail = hook.lock();
        Self::set_ept_entry_and_invalidate_tlb(
            page_aligned_pa,
            &hook_detail.original_entry,
        );

        if hook_detail.physical_base_address_of_fake_page_contents != 0 {
            if let Some(fake_page_va) = MemoryManager::physical_to_virtual(
                hook_detail.physical_base_address_of_fake_page_contents
            ) {
                MemoryManager::free_contiguous(fake_page_va);
            }
        }

        context.remove_hook(page_aligned_pa)
    }

    pub fn get_hooked_page_detail_by_physical_address(physical_address: u64) -> Option<Arc<Mutex<EptHookedPageDetail>>> {
        let page_aligned_pa = physical_address & !0xFFF;
        let context = EPT_HOOK_CONTEXT.lock();
        context.get_hook(page_aligned_pa)
    }

    pub unsafe fn handle_executeViolation(
        physical_address: u64,
        mtf_state: &mut MtfTrapState,
    ) -> Result<(), EptHookError> {
        let page_aligned_pa = physical_address & !0xFFF;
        
        let hook = Self::get_hooked_page_detail_by_physical_address(page_aligned_pa)
            .ok_or(EptHookError::NotHooked)?;

        let hook_detail = hook.lock();

        Self::set_ept_entry_and_invalidate_tlb(
            hook_detail.physical_base_address_of_fake_page_contents,
            &hook_detail.changed_entry,
        );

        mtf_state.trap_type = if hook_detail.is_hidden_breakpoint {
            MtfTrapType::HiddenBreakpoint
        } else {
            MtfTrapType::EptHookRestore
        };
        mtf_state.enabled = true;

        Ok(())
    }

    pub unsafe fn handle_read_or_write_violation(
        physical_address: u64,
        mtf_state: &mut MtfTrapState,
    ) -> Result<(), EptHookError> {
        let page_aligned_pa = physical_address & !0xFFF;

        let hook = Self::get_hooked_page_detail_by_physical_address(page_aligned_pa)
            .ok_or(EptHookError::NotHooked)?;

        let hook_detail = hook.lock();

        Self::set_ept_entry_and_invalidate_tlb(
            hook_detail.physical_base_address,
            &hook_detail.original_entry,
        );

        mtf_state.trap_type = MtfTrapType::EptHookRestore;
        mtf_state.enabled = true;

        Ok(())
    }

    fn build_absolute_jump(target: u64) -> [u8; 14] {
        let mut code = [0u8; 14];
        code[0] = 0xFF;
        code[1] = 0x25;
        code[2] = 0x00;
        code[3] = 0x00;
        code[4] = 0x00;
        code[5] = 0x00;
        code[6..14].copy_from_slice(&target.to_le_bytes());
        code
    }

    unsafe fn get_current_ept_entry(physical_address: u64) -> EptPml1Entry {
        let mut entry = EptPml1Entry::new();
        entry.set_read_access(true);
        entry.set_write_access(true);
        entry.set_execute_access(true);
        entry.set_memory_type(6);
        entry.set_ignore_pat(false);
        entry.set_page_frame_number(physical_address / PAGE_SIZE as u64);
        entry
    }

    pub unsafe fn set_ept_entry_and_invalidate_tlb(
        _physical_address: u64,
        _entry: &EptPml1Entry,
    ) {
        // TODO: 实现实际的 EPT 条目设置和 TLB 失效
        // 需要访问当前处理器的 EPT 页表并更新相应的 PML1 条目
        // 然后调用 invept_single_context
    }
}

pub unsafe fn ept_hook_add_hidden_breakpoint(
    virtual_address: u64,
    process_id: u32,
) -> Result<(), EptHookError> {
    EptHookManager::perform_page_hook_monitor_and_inline_hook(
        virtual_address,
        process_id,
        None,
        true,
    )
}

pub unsafe fn ept_hook_add_execution_hook(
    virtual_address: u64,
    process_id: u32,
    hook_handler: u64,
) -> Result<(), EptHookError> {
    EptHookManager::perform_page_hook_monitor_and_inline_hook(
        virtual_address,
        process_id,
        Some(hook_handler),
        false,
    )
}

pub unsafe fn ept_hook_remove(virtual_address: u64) -> Result<(), EptHookError> {
    let physical_address = MemoryManager::virtual_to_physical(virtual_address);
    if physical_address == 0 {
        return Err(EptHookError::InvalidAddress);
    }
    EptHookManager::remove_hook(physical_address)
}

pub fn ept_hook_get_count() -> usize {
    let context = EPT_HOOK_CONTEXT.lock();
    context.get_hook_count()
}

pub fn ept_hook_is_address_hooked(virtual_address: u64) -> bool {
    let physical_address = unsafe { MemoryManager::virtual_to_physical(virtual_address) };
    if physical_address == 0 {
        return false;
    }
    let page_aligned_pa = physical_address & !0xFFF;
    let context = EPT_HOOK_CONTEXT.lock();
    context.is_hooked(page_aligned_pa)
}

pub fn ept_hook_get_detail(physical_address: u64) -> Option<Arc<Mutex<EptHookedPageDetail>>> {
    EptHookManager::get_hooked_page_detail_by_physical_address(physical_address)
}

pub unsafe fn ept_hook_handle_violation(
    physical_address: u64,
    is_read: bool,
    is_write: bool,
    is_execute: bool,
    mtf_state: &mut MtfTrapState,
) -> bool {
    if is_execute {
        EptHookManager::handle_executeViolation(physical_address, mtf_state).is_ok()
    } else if is_read || is_write {
        EptHookManager::handle_read_or_write_violation(physical_address, mtf_state).is_ok()
    } else {
        false
    }
}
