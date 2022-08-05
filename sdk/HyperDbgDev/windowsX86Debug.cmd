cd .
xmake f -c -p windows -a x86 -m debug
xmake -r
xmake package -vD
xmake project -k cmakelists
xmake project -k vsxmake -m "debug,release"
xmake project -k compile_commands
pause