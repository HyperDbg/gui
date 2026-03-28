use crate::hyperkd::Vcpu;

pub struct MtfState {
    pub enabled: bool,
    pub callback: Option<unsafe fn(u32)>,
}

impl MtfState {
    pub const fn new() -> Self {
        Self {
            enabled: false,
            callback: None,
        }
    }
}

pub fn mtf_handle_vmexit(_core_id: u32, _mtf_state: &mut MtfState) -> Result<(), &'static str> {
    Ok(())
}

pub fn hv_set_monitor_trap_flag(_enable: bool) {
}
