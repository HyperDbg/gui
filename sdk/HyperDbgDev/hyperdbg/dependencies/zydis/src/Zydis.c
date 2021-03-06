#include <Zydis/Zydis.h>
ZyanU64
ZydisGetVersion(void) {
    return ZYDIS_VERSION;
}

ZyanStatus
ZydisIsFeatureEnabled(ZydisFeature feature) {
    switch (feature) {
    case ZYDIS_FEATURE_DECODER:
#ifndef ZYDIS_DISABLE_DECODER
        return ZYAN_STATUS_TRUE;
#else
        return ZYAN_STATUS_FALSE;
#endif
    case ZYDIS_FEATURE_FORMATTER:
#ifndef ZYDIS_DISABLE_FORMATTER
        return ZYAN_STATUS_TRUE;
#else
        return ZYAN_STATUS_FALSE;
#endif
    case ZYDIS_FEATURE_AVX512:
#ifndef ZYDIS_DISABLE_AVX512
        return ZYAN_STATUS_TRUE;
#else
        return ZYAN_STATUS_FALSE;
#endif
    case ZYDIS_FEATURE_KNC:
#ifndef ZYDIS_DISABLE_KNC
        return ZYAN_STATUS_TRUE;
#else
        return ZYAN_STATUS_FALSE;
#endif
    default:
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
}
