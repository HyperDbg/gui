pub use crate::hyperkd::hyperhv::hooks::hooks::{
    SyscallHookManager,
    SyscallHookEntry,
    SYSCALL_HOOK_MANAGER,
    install_global_syscall_hook,
    uninstall_global_syscall_hook,
    add_syscall_filter,
    remove_syscall_filter,
    is_syscall_filtered,
    get_original_lstar,
};
