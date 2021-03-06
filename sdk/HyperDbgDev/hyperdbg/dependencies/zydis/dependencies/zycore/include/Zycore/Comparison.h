#ifndef ZYCORE_COMPARISON_H
#define ZYCORE_COMPARISON_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef ZyanBool (*ZyanEqualityComparison)(const void * left, const void * right);
typedef ZyanI32 (*ZyanComparison)(const void * left, const void * right);
#define ZYAN_DECLARE_EQUALITY_COMPARISON(name, type)       \
    ZyanBool name(const type * left, const type * right) { \
        ZYAN_ASSERT(left);                                 \
        ZYAN_ASSERT(right);                                \
                                                           \
        return (*left == *right) ? ZYAN_TRUE : ZYAN_FALSE; \
    }
#define ZYAN_DECLARE_EQUALITY_COMPARISON_FOR_FIELD(name, type, field_name)       \
    ZyanBool name(const type * left, const type * right) {                       \
        ZYAN_ASSERT(left);                                                       \
        ZYAN_ASSERT(right);                                                      \
                                                                                 \
        return (left->field_name == right->field_name) ? ZYAN_TRUE : ZYAN_FALSE; \
    }
#define ZYAN_DECLARE_COMPARISON(name, type)               \
    ZyanI32 name(const type * left, const type * right) { \
        ZYAN_ASSERT(left);                                \
        ZYAN_ASSERT(right);                               \
                                                          \
        if (*left < *right) {                             \
            return -1;                                    \
        }                                                 \
        if (*left > *right) {                             \
            return 1;                                     \
        }                                                 \
        return 0;                                         \
    }
#define ZYAN_DECLARE_COMPARISON_FOR_FIELD(name, type, field_name) \
    ZyanI32 name(const type * left, const type * right) {         \
        ZYAN_ASSERT(left);                                        \
        ZYAN_ASSERT(right);                                       \
                                                                  \
        if (left->field_name < right->field_name) {               \
            return -1;                                            \
        }                                                         \
        if (left->field_name > right->field_name) {               \
            return 1;                                             \
        }                                                         \
        return 0;                                                 \
    }
ZYAN_INLINE
ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsPointer, void * const)
    ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsBool, ZyanBool)
        ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric8, ZyanU8)
            ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric16, ZyanU16)
                ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric32, ZyanU32)
                    ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric64, ZyanU64)
                        ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanComparePointer, void * const)
                            ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareBool, ZyanBool)
                                ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric8, ZyanU8)
                                    ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric16, ZyanU16)
                                        ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric32, ZyanU32)
                                            ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric64, ZyanU64)
#ifdef __cplusplus
}

#endif
#endif /* ZYCORE_COMPARISON_H */
