#include <inttypes.h>
#include <stdio.h>
#include <time.h>
#include <Zycore/Allocator.h>
#include <Zycore/Defines.h>
#include <Zycore/LibC.h>
#include <Zycore/Types.h>
#include <Zycore/Vector.h>
typedef struct TestStruct_ {
    ZyanU32 u32;
    ZyanU64 u64;
    float   f;
} TestStruct;
static void
InitTestdata(TestStruct * data, ZyanU32 n) {
    ZYAN_ASSERT(data);
    data->u32 = n;
    data->u64 = n;
    data->f   = (float)n;
}

static ZyanStatus
PerformBasicTests(ZyanVector * vector) {
    ZYAN_ASSERT(vector);
    static TestStruct         e_v;
    static const TestStruct * e_p;
    for (ZyanU32 i = 0; i < 20; ++i) {
        InitTestdata(&e_v, i);
        ZYAN_CHECK(ZyanVectorPushBack(vector, &e_v));
    }
    ZYAN_CHECK(ZyanVectorDeleteRange(vector, 5, 5));
    InitTestdata(&e_v, 12345678);
    ZYAN_CHECK(ZyanVectorInsert(vector, 5, &e_v));
    InitTestdata(&e_v, 87654321);
    ZYAN_CHECK(ZyanVectorSet(vector, 10, &e_v));
    ZyanUSize value;
    ZYAN_CHECK(ZyanVectorGetSize(vector, &value));
    puts("ELEMENTS");
    for (ZyanUSize i = 0; i < value; ++i) {
        ZYAN_CHECK(ZyanVectorGetPointer(vector, i, (const void **)&e_p));
        printf("  Element #%02" PRIuPTR ": %08" PRIu64 "\n", i, e_p->u64);
    }
    puts("INFO");
    printf("  Size       : %08" PRIuPTR "\n", value);
    ZYAN_CHECK(ZyanVectorGetCapacity(vector, &value));
    printf("  Capacity   : %08" PRIuPTR "\n\n", value);
    return ZYAN_STATUS_SUCCESS;
}

static ZyanI32
TestDataComparison(const TestStruct * left, const TestStruct * right) {
    ZYAN_ASSERT(left);
    ZYAN_ASSERT(right);
    if (left->u32 < right->u32) {
        return -1;
    }
    if (left->u32 > right->u32) {
        return 1;
    }
    return 0;
}

static ZyanStatus
PerformBinarySearchTest(ZyanVector * vector) {
    ZYAN_ASSERT(vector);
    static TestStruct         e_v;
    static const TestStruct * e_p;
    ZyanUSize                 value;
    ZYAN_CHECK(ZyanVectorGetCapacity(vector, &value));
    for (ZyanUSize i = 0; i < value; ++i) {
        const ZyanU32 n = rand() % 100;
        InitTestdata(&e_v, n);
        ZyanUSize found_index;
        ZYAN_CHECK(ZyanVectorBinarySearch(vector, &e_v, &found_index, (ZyanComparison)&TestDataComparison));
        ZYAN_CHECK(ZyanVectorInsert(vector, found_index, &e_v));
    }
    ZYAN_CHECK(ZyanVectorGetSize(vector, &value));
    puts("ELEMENTS");
    for (ZyanUSize i = 0; i < value; ++i) {
        ZYAN_CHECK(ZyanVectorGetPointer(vector, i, (const void **)&e_p));
        printf("  Element #%02" PRIuPTR ": %08" PRIu32 "\n", i, e_p->u32);
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
TestDynamic(void) {
    ZyanVector vector;
    ZYAN_CHECK(ZyanVectorInit(&vector, sizeof(TestStruct), 10, ZYAN_NULL));
    ZYAN_CHECK(PerformBasicTests(&vector));
    ZYAN_CHECK(ZyanVectorClear(&vector));
    ZYAN_CHECK(ZyanVectorReserve(&vector, 20));
    ZYAN_CHECK(PerformBinarySearchTest(&vector));
    return ZyanVectorDestroy(&vector);
}

static ZyanStatus
TestStatic(void) {
    static TestStruct buffer[20];
    ZyanVector        vector;
    ZYAN_CHECK(ZyanVectorInitCustomBuffer(&vector, sizeof(TestStruct), buffer, ZYAN_ARRAY_LENGTH(buffer), ZYAN_NULL));
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&vector, &size));
    for (ZyanUSize i = 0; i < size; ++i) {
        static TestStruct * element;
        ZYAN_CHECK(ZyanVectorGetPointer(&vector, i, (const void **)&element));
        if (element->u64 != buffer[i].u64) {
            return ZYAN_STATUS_INVALID_OPERATION;
        }
    }
    ZYAN_CHECK(PerformBasicTests(&vector));
    ZYAN_CHECK(ZyanVectorClear(&vector));
    ZYAN_CHECK(PerformBinarySearchTest(&vector));
    return ZyanVectorDestroy(&vector);
}

static ZyanStatus
AllocatorAllocate(ZyanAllocator * allocator, void ** p, ZyanUSize element_size, ZyanUSize n) {
    ZYAN_ASSERT(allocator);
    ZYAN_ASSERT(p);
    ZYAN_ASSERT(element_size);
    ZYAN_ASSERT(n);
    ZYAN_UNUSED(allocator);
    *p = ZYAN_MALLOC(element_size * n);
    if (!*p) {
        return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
AllocatorReallocate(ZyanAllocator * allocator, void ** p, ZyanUSize element_size, ZyanUSize n) {
    ZYAN_ASSERT(allocator);
    ZYAN_ASSERT(p);
    ZYAN_ASSERT(element_size);
    ZYAN_ASSERT(n);
    ZYAN_UNUSED(allocator);
    void * const x = ZYAN_REALLOC(*p, element_size * n);
    if (!x) {
        return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
    }
    *p = x;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
AllocatorDeallocate(ZyanAllocator * allocator, void * p, ZyanUSize element_size, ZyanUSize n) {
    ZYAN_ASSERT(allocator);
    ZYAN_ASSERT(p);
    ZYAN_ASSERT(element_size);
    ZYAN_ASSERT(n);
    ZYAN_UNUSED(allocator);
    ZYAN_UNUSED(element_size);
    ZYAN_UNUSED(n);
    ZYAN_FREE(p);
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
TestAllocator(void) {
    ZyanAllocator allocator;
    ZYAN_CHECK(ZyanAllocatorInit(&allocator, &AllocatorAllocate, &AllocatorReallocate, &AllocatorDeallocate));
    ZyanVector vector;
    ZYAN_CHECK(ZyanVectorInitEx(&vector, sizeof(TestStruct), 5, ZYAN_NULL, &allocator, 10.0f, 0.0f));
    static TestStruct e_v;
    for (ZyanU32 i = 0; i < 10; ++i) {
        InitTestdata(&e_v, i);
        ZYAN_CHECK(ZyanVectorPushBack(&vector, &e_v));
    }
    ZyanUSize value;
    ZYAN_CHECK(ZyanVectorGetCapacity(&vector, &value));
    if (value != 60) // (5 + 1) * 10.0f
    {
        return ZYAN_STATUS_INVALID_OPERATION;
    }
    ZYAN_CHECK(ZyanVectorClear(&vector));
    puts("INFO");
    ZYAN_CHECK(ZyanVectorGetSize(&vector, &value));
    printf("  Size       : %08" PRIuPTR "\n", value);
    ZYAN_CHECK(ZyanVectorGetCapacity(&vector, &value));
    printf("  Capacity   : %08" PRIuPTR "\n\n", value);
    return ZyanVectorDestroy(&vector);
}

int
main() {
    time_t t;
    srand((unsigned)time(&t));
    if (!ZYAN_SUCCESS(TestDynamic())) {
        return EXIT_FAILURE;
    }
    if (!ZYAN_SUCCESS(TestStatic())) {
        return EXIT_FAILURE;
    }
    if (!ZYAN_SUCCESS(TestAllocator())) {
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}
