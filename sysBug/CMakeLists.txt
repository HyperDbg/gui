cmake_minimum_required(VERSION 3.15)
project(HyperDbg C CXX)

list(APPEND CMAKE_MODULE_PATH "$ENV{path}")
find_package(WDK REQUIRED)

set(SourceFiles
    "driver.c"
    "driver.h"
    "HyperDbgHyperLogIntrinsics.h"
)
include_directories(
    "."
)
wdk_add_driver(sys
    KMDF 1.15
    ${SourceFiles}
)
