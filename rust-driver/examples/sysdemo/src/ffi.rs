use wdk_sys::{PIO_STACK_LOCATION, PIRP};

#[allow(non_snake_case)]
pub unsafe fn IoGetCurrentIrpStackLocation(p_irp: PIRP) -> PIO_STACK_LOCATION {
    unsafe {
        assert!((*p_irp).CurrentLocation <= (*p_irp).StackCount + 1);
        (*p_irp)
            .Tail
            .Overlay
            .__bindgen_anon_2
            .__bindgen_anon_1
            .CurrentStackLocation
    }
}
