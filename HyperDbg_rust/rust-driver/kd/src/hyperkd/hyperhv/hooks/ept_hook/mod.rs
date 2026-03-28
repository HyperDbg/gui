pub mod ept_hook;
pub mod exec_trap;

pub use ept_hook::{ept_hook_add_hidden_breakpoint, ept_hook_remove, EptHookError};
pub use exec_trap::{
    exec_trap_initialize, exec_trap_uninitialize, exec_trap_is_initialized,
    exec_trap_handle_ept_violation_vmexit, exec_trap_handle_move_to_adjusted_trap_state,
    exec_trap_restore_to_normal_eptp, DebuggerEventModeType,
};
