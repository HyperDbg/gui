#![no_std]
#![feature(asm_const)]
#![feature(naked_functions)]
#![feature(const_maybe_uninit_zeroed)]

extern crate alloc;

pub mod vmm;
pub mod memory;
pub mod hooks;
pub mod globals;
pub mod processor;
pub mod callbacks;
pub mod scheduler;
pub mod broadcast;
pub mod vmcall;
pub mod events;
pub mod interrupts;
pub mod compatibility;
pub mod switch_layout;
pub mod apic;
pub mod asm;
pub mod debugger_asm;
pub mod dirty_logging;
pub mod ept_vpid;
pub mod halted_core;
pub mod hyper_evade;
pub mod idt;
pub mod kernel_tests;
pub mod lbr;
pub mod loader;
pub mod meta_dispatch;
pub mod pci;
pub mod process;
pub mod serial_connection;
pub mod smi;
pub mod termination;
pub mod thread;
pub mod tracing;
pub mod transparency;
pub mod communication;
