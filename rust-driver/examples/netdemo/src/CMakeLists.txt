cmake_minimum_required(VERSION 3.28)
project(wsksample C)

set(CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}")
find_package(WDK REQUIRED)

set(WDK_NTDDI_VERSION "0x0A000008")

set(CMAKE_BUILD_TYPE Debug CACHE STRING "Build type" FORCE)

wdk_add_driver(wsksample
    wsksmple.c
)

target_compile_definitions(wsksample PRIVATE
    NTDDI_VERSION=0x0A000008
)

target_compile_options(wsksample PRIVATE /Od /Zi /RTC1)
target_link_options(wsksample PRIVATE /DEBUG)

target_link_libraries(wsksample WDK::NETIO)
