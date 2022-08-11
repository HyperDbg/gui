#include <Zycore/Allocator.h>
#include <Zycore/Defines.h>
#include <Zycore/LibC.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#include <stdio.h>
static ZyanStatus PerformBasicTests(ZyanString *string) {
  ZYAN_ASSERT(string);
  ZYAN_UNUSED(string);
  return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus TestDynamic(void) {
  PerformBasicTests(ZYAN_NULL);
  return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus TestStatic(void) {
  PerformBasicTests(ZYAN_NULL);
  return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus TestAllocator(void) { return ZYAN_STATUS_SUCCESS; }
int main() {
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

