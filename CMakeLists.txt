cmake_minimum_required(VERSION 3.28)

include("C:/Program Files/CMake/bin/ewdk.cmake")

project(HyperDbgUnified C CXX ASM_MASM)
set_property(GLOBAL PROPERTY USE_FOLDERS ON)

set(WDK_NTDDI_VERSION "0x0A000008")

# ============================================================
# Common source groups
# ============================================================

file(GLOB ZydisSrc "HyperDbg/hyperdbg/dependencies/zydis/src/*.c")
file(GLOB ZycoreSrc "HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/src/*.c")
list(APPEND ZydisSrc ${ZycoreSrc})

file(GLOB KernelOptimizations "HyperDbg/hyperdbg/include/components/optimizations/code/*.c")
file(GLOB KernelSpinlock    "HyperDbg/hyperdbg/include/components/spinlock/code/*.c")
file(GLOB KernelCallback    "HyperDbg/hyperdbg/include/components/callback/code/*.c")
file(GLOB KernelPlatform    "HyperDbg/hyperdbg/include/platform/kernel/code/*.c")
file(GLOB ScriptEvalSrc     "HyperDbg/hyperdbg/script-eval/code/*.c")

# ============================================================
# KERNEL-MODE TARGETS
# ============================================================

# ---- Platform (shared kernel platform lib) ----
km_lib(Platform ${KernelPlatform}
    WINVER 0x0A00
    INCLUDES
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
        "HyperDbg/hyperdbg/dependencies/zydis/include"
        "HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
        "HyperDbg/hyperdbg/hyperhv"
        "HyperDbg/hyperdbg/hyperhv/header"
        "HyperDbg/hyperdbg/hyperhv/code"
        "HyperDbg/hyperdbg/script-eval"
    DEFINES
        HYPERDBG_KERNEL_MODE HYPERDBG_VMM DHYPERDBG_VMM
        HYPERDBG_HYPEREVADE HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE
)
target_precompile_headers(Platform PRIVATE "HyperDbg/hyperdbg/hyperhv/pch.h")
set_target_properties(Platform PROPERTIES UNITY_BUILD OFF)

#### ---- kernel DLLs ---- ####

# hyperlog
file(GLOB _hyperlog_src "HyperDbg/hyperdbg/hyperlog/code/*.c")
km_lib(Hyperlog ${_hyperlog_src} ${KernelSpinlock}
    INCLUDES
        "HyperDbg/hyperdbg/hyperlog"
        "HyperDbg/hyperdbg/hyperlog/header"
        "HyperDbg/hyperdbg/hyperlog/code"
        "HyperDbg/hyperdbg/include"
    DEFINES
        DHYPERDBG_HYPERLOG HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE
        HYPERDBG_HYPEREVADE HYPERDBG_VMM
)
target_precompile_headers(Hyperlog PRIVATE "HyperDbg/hyperdbg/hyperlog/header/pch.h")

# hyperevade
file(GLOB _hev_src "HyperDbg/hyperdbg/hyperevade/code/*.c")
km_lib(Hyperevade ${_hev_src} ${KernelCallback} ${KernelOptimizations} ${KernelSpinlock}
    INCLUDES
        "HyperDbg/hyperdbg/hyperevade"
        "HyperDbg/hyperdbg/hyperevade/header"
        "HyperDbg/hyperdbg/hyperevade/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
    DEFINES
        DHYPERDBG_HYPEREVADE HYPERDBG_HYPEREVADE HYPERDBG_HYPER_LOG
        HYPERDBG_HYPERTRACE HYPERDBG_VMM
)
target_precompile_headers(Hyperevade PRIVATE "HyperDbg/hyperdbg/hyperevade/header/pch.h")

# hypertrace
file(GLOB_RECURSE _ht_src "HyperDbg/hyperdbg/hypertrace/code/*.c")
km_lib(Hypertrace ${_ht_src} ${KernelCallback}
    INCLUDES
        "HyperDbg/hyperdbg/hypertrace"
        "HyperDbg/hyperdbg/hypertrace/header"
        "HyperDbg/hyperdbg/hypertrace/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
    DEFINES
        DHYPERDBG_HYPERTRACE HYPERDBG_HYPERTRACE HYPERDBG_HYPER_LOG
        HYPERDBG_HYPEREVADE HYPERDBG_VMM
)
target_precompile_headers(Hypertrace PRIVATE "HyperDbg/hyperdbg/hypertrace/header/pch.h")

# hyperperf
file(GLOB _hp_src "HyperDbg/hyperdbg/hyperperf/code/*.c")
km_lib(Hyperperf ${_hp_src} ${KernelCallback}
    INCLUDES
        "HyperDbg/hyperdbg/hyperperf"
        "HyperDbg/hyperdbg/hyperperf/header"
        "HyperDbg/hyperdbg/hyperperf/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
    DEFINES
        DHYPERDBG_HYPERPERF HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE
        HYPERDBG_HYPEREVADE HYPERDBG_VMM
)
target_precompile_headers(Hyperperf PRIVATE "HyperDbg/hyperdbg/hyperperf/header/pch.h")

# kdserial
file(GLOB _kd_src "HyperDbg/hyperdbg/kdserial/*.c")
km_lib(Kdserial ${_kd_src}
    INCLUDES "HyperDbg/hyperdbg/kdserial"
)
set_target_properties(Kdserial PROPERTIES UNITY_BUILD OFF)

# zydis_kernel
km_lib(zydis_kernel ${ZydisSrc}
    INCLUDES
        "HyperDbg/hyperdbg/dependencies/zydis/include"
        "HyperDbg/hyperdbg/dependencies/zydis/src"
        "HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
    DEFINES
        ZYAN_NO_LIBC ZYDIS_NO_LIBC ZYCORE_STATIC_BUILD ZYDIS_STATIC_BUILD
)
set_target_properties(zydis_kernel PROPERTIES UNITY_BUILD OFF)

# hyperhv
file(GLOB_RECURSE _hv_src "HyperDbg/hyperdbg/hyperhv/code/*.c" "HyperDbg/hyperdbg/hyperhv/code/*.asm")
list(FILTER _hv_src EXCLUDE REGEX "UnloadDll")
km_lib(Hyperhv ${_hv_src} ${KernelOptimizations} ${KernelSpinlock}
    WINVER 0x0A00
    INCLUDES
        "HyperDbg/hyperdbg/hyperhv"
        "HyperDbg/hyperdbg/hyperhv/header"
        "HyperDbg/hyperdbg/hyperhv/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
        "HyperDbg/hyperdbg/dependencies/zydis/include"
        "HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
        "HyperDbg/hyperdbg/script-eval"
    DEFINES
        HYPERDBG_VMM DHYPERDBG_VMM HYPERDBG_HYPEREVADE
        HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE
        ZYDIS_STATIC_BUILD ZYCORE_STATIC_BUILD ZYAN_NO_LIBC ZYDIS_NO_LIBC
    LIBS
        zydis_kernel Hyperevade Platform
)
target_precompile_headers(Hyperhv PRIVATE "HyperDbg/hyperdbg/hyperhv/pch.h")

# hyperkd (.sys driver)
file(GLOB_RECURSE _kd_sys_src "HyperDbg/hyperdbg/hyperkd/code/*.c")
file(GLOB _kd_sys_asm "HyperDbg/hyperdbg/hyperkd/code/assembly/*.asm")
km_sys(hyperkd
    ${_kd_sys_src}
    ${_kd_sys_asm}
    ${KernelOptimizations}
    ${KernelSpinlock}
    ${ScriptEvalSrc}
    INCLUDES
        "HyperDbg/hyperdbg/hyperkd"
        "HyperDbg/hyperdbg/hyperkd/header"
        "HyperDbg/hyperdbg/hyperkd/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
        "HyperDbg/hyperdbg/script-eval"
        "HyperDbg/hyperdbg/include/components/optimizations/code"
        "HyperDbg/hyperdbg/include/components/spinlock/code"
    DEFINES
        DHYPERDBG_DEBUGGER WINNT
        HYPERDBG_HYPEREVADE HYPERDBG_HYPER_LOG
        HYPERDBG_HYPERTRACE HYPERDBG_VMM
    LIBS
        Hyperhv Hyperlog Kdserial Hypertrace Platform
)
target_precompile_headers(hyperkd PRIVATE "HyperDbg/hyperdbg/hyperkd/header/pch.h")
target_link_options(hyperkd PRIVATE /FORCE:MULTIPLE)

# ============================================================
# USER-MODE TARGETS
# ============================================================

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_MSVC_RUNTIME_LIBRARY "MultiThreaded")

# ---- Zydis (user-mode static lib) ----
um_lib(Zydis
    SOURCES ${ZydisSrc}
    INCLUDES
        "HyperDbg/hyperdbg/dependencies/zydis/include"
        "HyperDbg/hyperdbg/dependencies/zydis/src"
        "HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
    DEFINES ZYDIS_STATIC_BUILD ZYCORE_STATIC_BUILD
)
set_target_properties(Zydis PROPERTIES UNITY_BUILD OFF)

# ---- Pdbex ----
file(GLOB_RECURSE _pdbex "HyperDbg/hyperdbg/dependencies/pdbex/Source/*.cpp")
list(FILTER _pdbex EXCLUDE REGEX ".*main\\.cpp$")
um_lib(Pdbex
    SOURCES ${_pdbex}
    INCLUDES "HyperDbg/hyperdbg/dependencies/pdbex/Source"
    DEFINES _CRT_SECURE_NO_WARNINGS
)
set_target_properties(Pdbex PROPERTIES UNITY_BUILD OFF)

# ---- SymbolParser ----
file(GLOB _sp_src "HyperDbg/hyperdbg/symbol-parser/code/*.cpp")
um_lib(SymbolParser
    SOURCES ${_sp_src}
    INCLUDES
        "HyperDbg/hyperdbg/symbol-parser"
        "HyperDbg/hyperdbg/symbol-parser/header"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
        "HyperDbg/hyperdbg/dependencies/pdbex/Source"
    DEFINES HYPERDBG_SYMBOL_PARSER _CRT_SECURE_NO_WARNINGS
    LIBS Pdbex
)
target_precompile_headers(SymbolParser PRIVATE "HyperDbg/hyperdbg/symbol-parser/pch.h")

# ---- ScriptEngine ----
file(GLOB _se_src "HyperDbg/hyperdbg/script-engine/code/*.c")
list(APPEND _se_src
    "HyperDbg/hyperdbg/script-engine/code/script_include.c"
    "HyperDbg/hyperdbg/include/platform/user/code/platform-lib-calls.c"
)
um_lib(ScriptEngine
    SOURCES ${_se_src}
    INCLUDES
        "HyperDbg/hyperdbg/script-engine"
        "HyperDbg/hyperdbg/script-engine/header"
        "HyperDbg/hyperdbg/script-engine/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/script-eval"
        "HyperDbg/hyperdbg/include/platform/user/header"
    DEFINES
        SCRIPTENGINE_EXPORTS HYPERDBG_SCRIPT_ENGINE
        HYPERDBG_SYMBOL_PARSER _CRT_SECURE_NO_WARNINGS
    LIBS SymbolParser
)
set_target_properties(ScriptEngine PROPERTIES UNITY_BUILD OFF)
target_precompile_headers(ScriptEngine PRIVATE "HyperDbg/hyperdbg/script-engine/header/pch.h")

# ---- Libhyperdbg ----
# All .c files renamed to .cpp for C++ compilation (PCH requires C++ unity)
file(GLOB_RECURSE _lh_src "HyperDbg/hyperdbg/libhyperdbg/code/*.cpp")
list(FILTER _lh_src EXCLUDE REGEX ".*dllmain\\.cpp$")
list(APPEND _lh_src "HyperDbg/hyperdbg/libhyperdbg/pch.cpp")
file(GLOB _lh_asm "HyperDbg/hyperdbg/libhyperdbg/code/assembly/*.asm")
file(GLOB_RECURSE _up_src "HyperDbg/hyperdbg/include/platform/user/code/*.c")

foreach(_s ${ScriptEvalSrc})
    get_filename_component(_n ${_s} NAME)
    configure_file(${_s} "${CMAKE_CURRENT_BINARY_DIR}/script-eval-cpp/${_n}pp" COPYONLY)
endforeach()
foreach(_s ${_up_src})
    get_filename_component(_n ${_s} NAME)
    configure_file(${_s} "${CMAKE_CURRENT_BINARY_DIR}/user-platform-cpp/${_n}pp" COPYONLY)
endforeach()
file(GLOB _se_cpp "${CMAKE_CURRENT_BINARY_DIR}/script-eval-cpp/*.cpp")
file(GLOB _up_cpp "${CMAKE_CURRENT_BINARY_DIR}/user-platform-cpp/*.cpp")

um_lib(Libhyperdbg
    SOURCES
        ${_lh_src}
        ${_lh_asm}
        ${_se_cpp}
        "HyperDbg/hyperdbg/include/components/pe/code/pe-image-reader.cpp"
        "HyperDbg/hyperdbg/libhyperdbg/ucpuid.cpp"
        ${_up_cpp}
    INCLUDES
        "HyperDbg/hyperdbg/libhyperdbg"
        "HyperDbg/hyperdbg/libhyperdbg/header"
        "HyperDbg/hyperdbg/libhyperdbg/code"
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/include/components/pe/header"
        "HyperDbg/hyperdbg/include/components/pe/code"
        "HyperDbg/hyperdbg/include/platform/user/header"
        "HyperDbg/hyperdbg/include/platform/user/header/windows-only"
        "HyperDbg/hyperdbg/dependencies"
        "HyperDbg/hyperdbg/dependencies/zydis/include"
        "HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
        "HyperDbg/hyperdbg/dependencies/keystone/include"
        "HyperDbg/hyperdbg/dependencies/libipt"
        "HyperDbg/hyperdbg/dependencies/pdbex/Source"
        "HyperDbg/hyperdbg/script-eval"
        "HyperDbg/hyperdbg/script-engine/header"
        "HyperDbg/hyperdbg/symbol-parser/header"
    DEFINES
        HYPERDBG_LIBHYPERDBG HYPERDBG_SCRIPT_ENGINE HYPERDBG_SYMBOL_PARSER
        ZYCORE_STATIC_BUILD ZYDIS_STATIC_BUILD HPRDBGCTRL_EXPORTS
        _WIN32_WINNT=0x0601 WINVER=0x0601 NTDDI_VERSION=0x06010000
        _CRT_SECURE_NO_WARNINGS
    LIBS ScriptEngine SymbolParser Pdbex Zydis
)
set_source_files_properties(${_lh_asm} PROPERTIES LANGUAGE ASM_MASM)
target_precompile_headers(Libhyperdbg PRIVATE "HyperDbg/hyperdbg/libhyperdbg/pch.h")

# ---- hyperdbg-cli ----
um_exe(hyperdbg-cli
    SOURCES "HyperDbg/hyperdbg/hyperdbg-cli/hyperdbg-cli.cpp"
    INCLUDES
        "HyperDbg/hyperdbg/include"
        "HyperDbg/hyperdbg/dependencies"
    DEFINES
        HYPERDBG_LIBHYPERDBG HYPERDBG_SCRIPT_ENGINE HYPERDBG_SYMBOL_PARSER
        HYPERDBG_DEBUG _CRT_SECURE_NO_WARNINGS
    LIBS
        Libhyperdbg
        dbghelp urlmon shell32 ole32 oleaut32
        keystone
        libipt
    COMPILE_OPTIONS /std:c++20
)
target_link_directories(hyperdbg-cli PRIVATE
    $<$<CONFIG:Release>:${CMAKE_CURRENT_SOURCE_DIR}/HyperDbg/hyperdbg/libraries/keystone/release-lib>
    $<$<CONFIG:Debug>:${CMAKE_CURRENT_SOURCE_DIR}/HyperDbg/hyperdbg/libraries/keystone/debug-lib>
    "${CMAKE_CURRENT_SOURCE_DIR}/HyperDbg/hyperdbg/libraries/libipt"
    "${CMAKE_CURRENT_BINARY_DIR}"
)
target_link_options(hyperdbg-cli PRIVATE
    /FORCE:MULTIPLE
)
add_custom_command(TARGET hyperdbg-cli PRE_LINK
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_CURRENT_BINARY_DIR}/Zydis.lib"
        "${CMAKE_CURRENT_BINARY_DIR}/Zycore.lib"
)
add_custom_command(TARGET hyperdbg-cli POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_CURRENT_SOURCE_DIR}/HyperDbg/hyperdbg/libraries/libipt/libipt.dll"
        "$<TARGET_FILE_DIR:hyperdbg-cli>"
    COMMENT "Copying libipt.dll to output directory"
)

 