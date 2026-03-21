# pdbfetch

A cross platform standalone cli application written in Go that fetches PDB symbol files directly from Microsoft's symbol servers by parsing an associated PE formatted executable.

## Usage

```
pdbfetch pefile [directory]
```
#### Windows
```
pdbfetch ExamplePE.exe "C:\symbols"
```
#### Unix
```
pdbfetch ExamplePE.exe "/usr/share/symbols"
```

## How

By parsing a given PE file's `IMAGE_DIRECTORY_ENTRY_DEBUG` directory for PDB symbol information of type `IMAGE_DEBUG_TYPE_CODEVIEW`, and compiling a valid symbol server URL.
The PDB `Name`, `Signtaure`, and `Age` fields are used to compile the symbol download URL.

The `Name` is the name (or path?) of the PDB that is emitted by the linker.

The `Signtaure` is a Windows formatted 128-bit GUID that is emitted by the linker.

The  `Age` is the version of the same PDB (or simply the number of times the PDB file has been written by the linker).

Compiling the link is as simple as:

`https://msdl.microsoft.com/download/symbols/{PDB_NAME}/{PDB_GUID}+{PDB_AGE}/{PDB_NAME}`

Let's take for example `netbios.sys`. The following PDB symbol information is extracted from the PE:
```
PDB Name: netbios.pdb
PDB GUID: 109096de-2cf2-44bd-8d03-0e5473c5253a
PDB Age:  1
```
The resulting valid symbol link for this is:

http://msdl.microsoft.com/download/symbols/netbios.pdb/109096DE2CF244BD8D030E5473C5253A1/netbios.pdb  

## References

This would not have been so simple without the help of `PDB-Downloader`:

https://github.com/rajkumar-rangaraj/PDB-Downloader

Special thanks to Sam O (and Ero Carrera) for a great start to a PE file parser in golang:

https://github.com/soluwalana/pefile-go

LLVM PDB Documentation:

https://llvm.org/docs/PDB/PdbStream.html#matching-a-pdb-to-its-executable

https://github.com/llvm/llvm-project/blob/master/llvm/include/llvm/Object/CVDebugRecord.h

## License

This project is licensed under the Apache 2.0 license.

Additional licenses pertaining to this project can be found under the LICENSES sub-folder.
