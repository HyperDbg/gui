x64dbg-StaticAnalysis
=====================
This is a plugin for the project x64_dbg. It detect win32api calls and add useful comments at found parameters for easier analysis. It supports **x86 and x64**.
The plugin delivers a comprehensive list (3304 descriptions of functions) of function prototypes. *These were crafted by hand and RegEx, so please respect the credits and license.

### x86 and x64
It fully supports the x64 calling convention:
![default x86 calling](https://raw.githubusercontent.com/x64dbg/x64dbg-StaticAnalysis/master/x64.PNG)

It even detects different argument order. Notice that the prototype is

   int GetModuleFileNameW(DWORD nBufferSize,LPTSTR lpBuffer,HMODULE hModule)

but the plugin recognize the order of the arguments
![default x86 calling](https://raw.githubusercontent.com/x64dbg/x64dbg-StaticAnalysis/master/x64_2.PNG)
Hence it knows, that in this case the argument order changed.


### advanced parameter detection (MinGW)

![default x86 calling](https://raw.githubusercontent.com/x64dbg/x64dbg-StaticAnalysis/master/analysis2.PNG)
Due some stack emulation it is possible to analyse MingGW parameter passing to the stack
![MingGW parameters](https://raw.githubusercontent.com/x64dbg/x64dbg-StaticAnalysis/master/analysis.PNG)

## Download
See [release page](https://github.com/x64dbg/x64dbg-StaticAnalysis/release) for lastest compiled plugin. The source of the latest release is in the branch "stable".

This master branch may contains experimental code and the latest commits.

Make sure that you installed the "Microsoft Visual C++ 2013 Redistributatble Package".
(http://www.microsoft.com/en-us/download/details.aspx?id=40784)


## implemented Features
- auto comments on parameters for api-function
- real stack emulator for MinGW-arguments detection
- real register emulator for x64 calling convention
- supports x86 and x64 targets

## Roadmap


- function-finder (function body, xrefs, function returns)
- case-switch detection
- "MB_OK"-like symbols
- loop-detection
- flow-heuristic against obfuscation instead of linear scanning

## License
GLPv3
