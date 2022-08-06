typedef enum ZydisInstructionCategory_
{
    ZYDIS_CATEGORY_INVALID,
    ZYDIS_CATEGORY_ADOX_ADCX,
    ZYDIS_CATEGORY_AES,
    ZYDIS_CATEGORY_AMD3DNOW,
    ZYDIS_CATEGORY_AVX,
    ZYDIS_CATEGORY_AVX2,
    ZYDIS_CATEGORY_AVX2GATHER,
    ZYDIS_CATEGORY_AVX512,
    ZYDIS_CATEGORY_AVX512_4FMAPS,
    ZYDIS_CATEGORY_AVX512_4VNNIW,
    ZYDIS_CATEGORY_AVX512_BITALG,
    ZYDIS_CATEGORY_AVX512_VBMI,
    ZYDIS_CATEGORY_AVX512_VP2INTERSECT,
    ZYDIS_CATEGORY_BINARY,
    ZYDIS_CATEGORY_BITBYTE,
    ZYDIS_CATEGORY_BLEND,
    ZYDIS_CATEGORY_BMI1,
    ZYDIS_CATEGORY_BMI2,
    ZYDIS_CATEGORY_BROADCAST,
    ZYDIS_CATEGORY_CALL,
    ZYDIS_CATEGORY_CET,
    ZYDIS_CATEGORY_CLDEMOTE,
    ZYDIS_CATEGORY_CLFLUSHOPT,
    ZYDIS_CATEGORY_CLWB,
    ZYDIS_CATEGORY_CLZERO,
    ZYDIS_CATEGORY_CMOV,
    ZYDIS_CATEGORY_COMPRESS,
    ZYDIS_CATEGORY_COND_BR,
    ZYDIS_CATEGORY_CONFLICT,
    ZYDIS_CATEGORY_CONVERT,
    ZYDIS_CATEGORY_DATAXFER,
    ZYDIS_CATEGORY_DECIMAL,
    ZYDIS_CATEGORY_ENQCMD,
    ZYDIS_CATEGORY_EXPAND,
    ZYDIS_CATEGORY_FCMOV,
    ZYDIS_CATEGORY_FLAGOP,
    ZYDIS_CATEGORY_FMA4,
    ZYDIS_CATEGORY_GATHER,
    ZYDIS_CATEGORY_GFNI,
    ZYDIS_CATEGORY_IFMA,
    ZYDIS_CATEGORY_INTERRUPT,
    ZYDIS_CATEGORY_IO,
    ZYDIS_CATEGORY_IOSTRINGOP,
    ZYDIS_CATEGORY_KMASK,
    ZYDIS_CATEGORY_KNC,
    ZYDIS_CATEGORY_KNCMASK,
    ZYDIS_CATEGORY_KNCSCALAR,
    ZYDIS_CATEGORY_LOGICAL,
    ZYDIS_CATEGORY_LOGICAL_FP,
    ZYDIS_CATEGORY_LZCNT,
    ZYDIS_CATEGORY_MISC,
    ZYDIS_CATEGORY_MMX,
    ZYDIS_CATEGORY_MOVDIR,
    ZYDIS_CATEGORY_MPX,
    ZYDIS_CATEGORY_NOP,
    ZYDIS_CATEGORY_PADLOCK,
    ZYDIS_CATEGORY_PCLMULQDQ,
    ZYDIS_CATEGORY_PCONFIG,
    ZYDIS_CATEGORY_PKU,
    ZYDIS_CATEGORY_POP,
    ZYDIS_CATEGORY_PREFETCH,
    ZYDIS_CATEGORY_PREFETCHWT1,
    ZYDIS_CATEGORY_PT,
    ZYDIS_CATEGORY_PUSH,
    ZYDIS_CATEGORY_RDPID,
    ZYDIS_CATEGORY_RDPRU,
    ZYDIS_CATEGORY_RDRAND,
    ZYDIS_CATEGORY_RDSEED,
    ZYDIS_CATEGORY_RDWRFSGS,
    ZYDIS_CATEGORY_RET,
    ZYDIS_CATEGORY_ROTATE,
    ZYDIS_CATEGORY_SCATTER,
    ZYDIS_CATEGORY_SEGOP,
    ZYDIS_CATEGORY_SEMAPHORE,
    ZYDIS_CATEGORY_SETCC,
    ZYDIS_CATEGORY_SGX,
    ZYDIS_CATEGORY_SHA,
    ZYDIS_CATEGORY_SHIFT,
    ZYDIS_CATEGORY_SMAP,
    ZYDIS_CATEGORY_SSE,
    ZYDIS_CATEGORY_STRINGOP,
    ZYDIS_CATEGORY_STTNI,
    ZYDIS_CATEGORY_SYSCALL,
    ZYDIS_CATEGORY_SYSRET,
    ZYDIS_CATEGORY_SYSTEM,
    ZYDIS_CATEGORY_TBM,
    ZYDIS_CATEGORY_UFMA,
    ZYDIS_CATEGORY_UNCOND_BR,
    ZYDIS_CATEGORY_VAES,
    ZYDIS_CATEGORY_VBMI2,
    ZYDIS_CATEGORY_VFMA,
    ZYDIS_CATEGORY_VPCLMULQDQ,
    ZYDIS_CATEGORY_VTX,
    ZYDIS_CATEGORY_WAITPKG,
    ZYDIS_CATEGORY_WIDENOP,
    ZYDIS_CATEGORY_X87_ALU,
    ZYDIS_CATEGORY_XOP,
    ZYDIS_CATEGORY_XSAVE,
    ZYDIS_CATEGORY_XSAVEOPT,
    ZYDIS_CATEGORY_MAX_VALUE = ZYDIS_CATEGORY_XSAVEOPT,
    ZYDIS_CATEGORY_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CATEGORY_MAX_VALUE)
} ZydisInstructionCategory;
