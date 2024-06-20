clang++ -nobuiltininc -Xclang -ast-dump=json -fsyntax-only test.hpp  > ast_output.json 2>&1
clang++ -fsyntax-only -nobuiltininc -emit-llvm -Xclang -fdump-record-layouts -Xclang -fdump-record-layouts-complete test.hpp  > ast_record_layouts.log 2>&1
clang++ -E -dM test.hpp > macros.txt
pause
