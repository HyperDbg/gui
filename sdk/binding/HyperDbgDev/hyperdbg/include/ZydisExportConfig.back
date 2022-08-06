#define NDEBUG

#ifndef ZYDIS_EXPORT_H
#    define ZYDIS_EXPORT_H

#    ifdef ZYDIS_STATIC_DEFINE
#        define ZYDIS_EXPORT
#        define ZYDIS_NO_EXPORT
#    else
#        ifndef ZYDIS_EXPORT
#            ifdef Zydis_EXPORTS

#                define ZYDIS_EXPORT __declspec(dllexport)
#            else

#                define ZYDIS_EXPORT __declspec(dllimport)
#            endif
#        endif

#        ifndef ZYDIS_NO_EXPORT
#            define ZYDIS_NO_EXPORT
#        endif
#    endif

#    ifndef ZYDIS_DEPRECATED
#        define ZYDIS_DEPRECATED __declspec(deprecated)
#    endif

#    ifndef ZYDIS_DEPRECATED_EXPORT
#        define ZYDIS_DEPRECATED_EXPORT ZYDIS_EXPORT ZYDIS_DEPRECATED
#    endif

#    ifndef ZYDIS_DEPRECATED_NO_EXPORT
#        define ZYDIS_DEPRECATED_NO_EXPORT ZYDIS_NO_EXPORT ZYDIS_DEPRECATED
#    endif

#    if 0 
#        ifndef ZYDIS_NO_DEPRECATED
#            define ZYDIS_NO_DEPRECATED
#        endif
#    endif

#endif 
