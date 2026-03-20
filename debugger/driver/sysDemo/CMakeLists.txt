cmake_minimum_required(VERSION 3.28)
project(sysDemo C)

set_property(GLOBAL PROPERTY USE_FOLDERS ON)

set(CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}")
find_package(WDK REQUIRED)

set(WDK_NTDDI_VERSION "0x0A000005")

set(CMAKE_BUILD_TYPE Debug CACHE STRING "Build type" FORCE)
set(CMAKE_C_FLAGS_DEBUG "/Od /Zi /RTC1")
set(CMAKE_C_FLAGS "/Od /Zi /RTC1")

add_definitions(
    -DNTDDI_VERSION=0x0A000005
)

file(GLOB SourceFiles "*.c")

wdk_add_driver(sysDemo
    KMDF 1.15
    ${SourceFiles}
)

target_compile_options(sysDemo PRIVATE /Od /Zi /RTC1)
target_link_options(sysDemo PRIVATE /DEBUG)
