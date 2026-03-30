pub mod idt;
pub mod interrupts;
pub mod processor;
pub mod smi;

pub use processor::get_current_processor_number;
pub use processor::get_processor_count;
pub use processor::run_on_processor;
