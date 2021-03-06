#include "pch.h"
VOID
XApicIcrWrite(UINT32 Low, UINT32 High) {
    *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset + 0x10) = High;
    *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset)        = Low;
}

VOID
X2ApicIcrWrite(UINT32 Low, UINT32 High) {
    __writemsr(X2_MSR_BASE + TO_X2(ICROffset), ((UINT64)High << 32) | Low);
}

VOID
ApicTriggerGenericNmi() {
    if (g_IsX2Apic) {
        X2ApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
    } else {
        XApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
    }
}

BOOLEAN
ApicInitialize() {
    UINT64           ApicBaseMSR;
    PHYSICAL_ADDRESS PaApicBase;
    ApicBaseMSR = __readmsr(0x1B);
    if (!(ApicBaseMSR & (1 << 11)))
        return FALSE;
    if (ApicBaseMSR & (1 << 10)) {
        g_IsX2Apic = TRUE;
        return FALSE;
    } else {
        PaApicBase.QuadPart = ApicBaseMSR & 0xFFFFFF000;
        g_ApicBase          = MmMapIoSpace(PaApicBase, 0x1000, MmNonCached);
        if (!g_ApicBase)
            return FALSE;
        g_IsX2Apic = FALSE;
    }
    return TRUE;
}

VOID
ApicUninitialize() {
    if (g_ApicBase)
        MmUnmapIoSpace(g_ApicBase, 0x1000);
}

VOID
ApicSelfIpi(UINT32 Vector) {
    if (g_IsX2Apic) {
        X2ApicIcrWrite(APIC_DEST_SELF | APIC_DEST_PHYSICAL | APIC_DM_FIXED | Vector, 0);
    } else {
        XApicIcrWrite(APIC_DEST_SELF | APIC_DEST_PHYSICAL | APIC_DM_FIXED | Vector, 0);
    }
}
