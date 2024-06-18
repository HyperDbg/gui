cmake_minimum_required(VERSION 3.10.2)

set(CMAKE_CXX_STANDARD 20)

project(zydis)

include_directories(.)
file(GLOB SRC_LIST ./*.cpp ./*.def ./*.c)
add_compile_definitions(
        ZYDIS_STATIC_BUILD
        ZYDIS_SHOULD_EXPORT
        DZYAN_NO_LIBC
)
add_library(zydis SHARED ${SRC_LIST})