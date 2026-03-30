use core::arch::naked_asm;

#[allow(static_mut_refs)]
#[repr(C)]
pub struct TrapFrame {
    pub rax: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rbx: u64,
    pub rbp: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
    pub vector: u64,
    pub error: u64,
    pub rip: u64,
    pub cs: u64,
    pub rflags: u64,
    pub rsp: u64,
    pub ss: u64,
}

extern "C" {
    fn IdtEmulationhandleHostInterrupt(frame: *mut TrapFrame);
}

macro_rules! define_isr_no_error {
    ($name:ident, $vector:expr) => {
        #[unsafe(naked)]
        pub unsafe extern "C" fn $name() {
            naked_asm!(
                "push 0",
                "push {0}",
                "push rax",
                "push rcx",
                "push rdx",
                "push rbx",
                "push rbp",
                "push rsi",
                "push rdi",
                "push r8",
                "push r9",
                "push r10",
                "push r11",
                "push r12",
                "push r13",
                "push r14",
                "push r15",
                "mov rcx, rsp",
                "call {1}",
                "pop r15",
                "pop r14",
                "pop r13",
                "pop r12",
                "pop r11",
                "pop r10",
                "pop r9",
                "pop r8",
                "pop rdi",
                "pop rsi",
                "pop rbp",
                "pop rbx",
                "pop rdx",
                "pop rcx",
                "pop rax",
                "add rsp, 16",
                "iretq",
                const $vector,
                sym IdtEmulationhandleHostInterrupt,
            );
        }
    };
}

macro_rules! define_isr_with_error {
    ($name:ident, $vector:expr) => {
        #[unsafe(naked)]
        pub unsafe extern "C" fn $name() {
            naked_asm!(
                "push {0}",
                "push rax",
                "push rcx",
                "push rdx",
                "push rbx",
                "push rbp",
                "push rsi",
                "push rdi",
                "push r8",
                "push r9",
                "push r10",
                "push r11",
                "push r12",
                "push r13",
                "push r14",
                "push r15",
                "mov rcx, rsp",
                "call {1}",
                "pop r15",
                "pop r14",
                "pop r13",
                "pop r12",
                "pop r11",
                "pop r10",
                "pop r9",
                "pop r8",
                "pop rdi",
                "pop rsi",
                "pop rbp",
                "pop rbx",
                "pop rdx",
                "pop rcx",
                "pop rax",
                "add rsp, 16",
                "iretq",
                const $vector,
                sym IdtEmulationhandleHostInterrupt,
            );
        }
    };
}

define_isr_no_error!(interrupt_handler_0, 0);
define_isr_no_error!(interrupt_handler_1, 1);
define_isr_no_error!(interrupt_handler_2, 2);
define_isr_no_error!(interrupt_handler_3, 3);
define_isr_no_error!(interrupt_handler_4, 4);
define_isr_no_error!(interrupt_handler_5, 5);
define_isr_no_error!(interrupt_handler_6, 6);
define_isr_no_error!(interrupt_handler_7, 7);
define_isr_with_error!(interrupt_handler_8, 8);
define_isr_no_error!(interrupt_handler_9, 9);
define_isr_with_error!(interrupt_handler_10, 10);
define_isr_with_error!(interrupt_handler_11, 11);
define_isr_with_error!(interrupt_handler_12, 12);
define_isr_with_error!(interrupt_handler_13, 13);
define_isr_with_error!(interrupt_handler_14, 14);
define_isr_no_error!(interrupt_handler_16, 16);
define_isr_with_error!(interrupt_handler_17, 17);
define_isr_no_error!(interrupt_handler_18, 18);
define_isr_no_error!(interrupt_handler_19, 19);
define_isr_no_error!(interrupt_handler_20, 20);
define_isr_with_error!(interrupt_handler_30, 30);

pub const INTERRUPT_HANDLER_COUNT: usize = 256;

#[repr(C)]
#[derive(Clone, Copy)]
pub struct InterruptDescriptor {
    offset_low: u16,
    selector: u16,
    ist: u8,
    type_attr: u8,
    offset_mid: u16,
    offset_high: u32,
    reserved: u32,
}

impl InterruptDescriptor {
    pub fn new(offset: u64, selector: u16, type_attr: u8, ist: u8) -> Self {
        Self {
            offset_low: (offset & 0xFFFF) as u16,
            selector,
            ist,
            type_attr,
            offset_mid: ((offset >> 16) & 0xFFFF) as u16,
            offset_high: ((offset >> 32) & 0xFFFFFFFF) as u32,
            reserved: 0,
        }
    }

    pub fn set_offset(&mut self, offset: u64) {
        self.offset_low = (offset & 0xFFFF) as u16;
        self.offset_mid = ((offset >> 16) & 0xFFFF) as u16;
        self.offset_high = ((offset >> 32) & 0xFFFFFFFF) as u32;
    }
}

#[repr(C)]
pub struct IdtRegister {
    pub limit: u16,
    pub base: u64,
}

pub struct InterruptHandlerTable {
    handlers: [Option<unsafe extern "C" fn()>; INTERRUPT_HANDLER_COUNT],
}

impl InterruptHandlerTable {
    pub const fn new() -> Self {
        Self {
            handlers: [None; INTERRUPT_HANDLER_COUNT],
        }
    }

    pub fn set_handler(&mut self, vector: usize, handler: unsafe extern "C" fn()) {
        if vector < INTERRUPT_HANDLER_COUNT {
            self.handlers[vector] = Some(handler);
        }
    }

    pub fn get_handler(&self, vector: usize) -> Option<unsafe extern "C" fn()> {
        if vector < INTERRUPT_HANDLER_COUNT {
            self.handlers[vector]
        } else {
            None
        }
    }

    pub fn create_idt(&self) -> [InterruptDescriptor; INTERRUPT_HANDLER_COUNT] {
        let mut idt = [InterruptDescriptor::new(0, 0, 0, 0); INTERRUPT_HANDLER_COUNT];
        
        for i in 0..INTERRUPT_HANDLER_COUNT {
            if let Some(handler) = self.handlers[i] {
                let handler_addr = handler as u64;
                idt[i] = InterruptDescriptor::new(handler_addr, 8, 0x8E, 0);
            }
        }
        
        idt
    }
}

pub static mut INTERRUPT_HANDLERS: InterruptHandlerTable = InterruptHandlerTable::new();

pub unsafe fn initialize_interrupt_handlers() {
    INTERRUPT_HANDLERS.set_handler(0, interrupt_handler_0);
    INTERRUPT_HANDLERS.set_handler(1, interrupt_handler_1);
    INTERRUPT_HANDLERS.set_handler(2, interrupt_handler_2);
    INTERRUPT_HANDLERS.set_handler(3, interrupt_handler_3);
    INTERRUPT_HANDLERS.set_handler(4, interrupt_handler_4);
    INTERRUPT_HANDLERS.set_handler(5, interrupt_handler_5);
    INTERRUPT_HANDLERS.set_handler(6, interrupt_handler_6);
    INTERRUPT_HANDLERS.set_handler(7, interrupt_handler_7);
    INTERRUPT_HANDLERS.set_handler(8, interrupt_handler_8);
    INTERRUPT_HANDLERS.set_handler(9, interrupt_handler_9);
    INTERRUPT_HANDLERS.set_handler(10, interrupt_handler_10);
    INTERRUPT_HANDLERS.set_handler(11, interrupt_handler_11);
    INTERRUPT_HANDLERS.set_handler(12, interrupt_handler_12);
    INTERRUPT_HANDLERS.set_handler(13, interrupt_handler_13);
    INTERRUPT_HANDLERS.set_handler(14, interrupt_handler_14);
    INTERRUPT_HANDLERS.set_handler(16, interrupt_handler_16);
    INTERRUPT_HANDLERS.set_handler(17, interrupt_handler_17);
    INTERRUPT_HANDLERS.set_handler(18, interrupt_handler_18);
    INTERRUPT_HANDLERS.set_handler(19, interrupt_handler_19);
    INTERRUPT_HANDLERS.set_handler(20, interrupt_handler_20);
    INTERRUPT_HANDLERS.set_handler(30, interrupt_handler_30);
}

pub unsafe fn load_idt(idt: &[InterruptDescriptor; INTERRUPT_HANDLER_COUNT]) {
    let idtr = IdtRegister {
        limit: (core::mem::size_of::<[InterruptDescriptor; INTERRUPT_HANDLER_COUNT]>() - 1) as u16,
        base: idt.as_ptr() as u64,
    };
    
    core::arch::asm!("lidt [{}]", in(reg) &idtr, options(nostack, readonly));
}

pub unsafe fn save_idt() -> IdtRegister {
    let mut idtr: IdtRegister = core::mem::zeroed();
    core::arch::asm!("sidt [{}]", in(reg) &mut idtr, options(nostack));
    idtr
}

pub unsafe fn restore_idt(idtr: &IdtRegister) {
    core::arch::asm!("lidt [{}]", in(reg) idtr, options(nostack, readonly));
}

pub fn create_idt_entry(offset: u64, selector: u16, type_attr: u8, ist: u8) -> InterruptDescriptor {
    InterruptDescriptor::new(offset, selector, type_attr, ist)
}

pub const IDT_TYPE_INTERRUPT_GATE: u8 = 0x8E;
pub const IDT_TYPE_TRAP_GATE: u8 = 0x8F;
pub const IDT_TYPE_TASK_GATE: u8 = 0x85;

pub fn create_interrupt_gate(offset: u64, selector: u16, ist: u8) -> InterruptDescriptor {
    create_idt_entry(offset, selector, IDT_TYPE_INTERRUPT_GATE, ist)
}

pub fn create_trap_gate(offset: u64, selector: u16, ist: u8) -> InterruptDescriptor {
    create_idt_entry(offset, selector, IDT_TYPE_TRAP_GATE, ist)
}
