REM AST generation from MSVC 
cl /analyze /d1Aprintast *.cpp > %1.ast
pause