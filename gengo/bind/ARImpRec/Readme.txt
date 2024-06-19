ARImpRec.DLL - 1.8.1 ---- Updated 2024 February, the 17th

+ Fixed bug when rebuilding export table in minimized files
+ Completed recovering of PE header directories data in minimized files

ARImpRec.DLL - 1.8.0 Beta ---- Updated 2014 February, the 13th

+ Fixed bug related to PE header when rebuilding in Windows7
+ Fixed bug about renaming dumped file 

ARImpRec.DLL - 1.7.9 ---- Updated 2014 June, the 16th

+ Fixed bugs in GetProcOrdinal, GetProcName and GetProcAndName functions

ARImpRec.DLL - 1.7.8 ---- Updated 2013 June, the 3rd

+ Improved performance of the search engine for imports
+ Updated the overlay generation engine
+ Fixed some bugs when analyzing code sections for imports
+ Fixed a bug in import rebuilder when forwarding
+ Updated TryGetImportedFunctionName function to TryGetImportedFunction, see below...

ARImpRec.DLL - 1.7.7 ---- Updated 2011 September, the 5th

+ Added an exported function: TryGetImportedFunctionName
+ Improved code for retrieving imports


ARImpRec.DLL - 1.7.6 ---- Updated 2011 June, the 30th

+ Improved sorting of imported functions and dlls by name
+ Fixed an issue with rebuilding by relocations when module is loaded at a different virtual address than imagebase
+ Fixed bugs when rebuilding without modifying any code


ARImpRec.DLL - 1.7.5 ---- Updated 2010 November, the 3rd

+ Fixed rebuilding of sections for minimized MinGW compiled executables


ARImpRec.DLL - 1.7.4 ---- Updated 2010 October, the 20th

+ Added new parameter to choose saving OEP to file
+ Added an exported function for rebuilding imports without modifying code
+ Improved detection of possible code sections
+ Removed the backup of rebuilt file
+ Fixed a bug when minimizing size for dlls


ARImpRec.DLL - 1.7.3 ---- Updated 2010 September, the 17th

+ Improved relocations analysis and rebuilding performance
+ Improved extraction of files of pdata section
+ Improved rebuilding of overlay for big ones
+ Fixed a bug for import table bigger than the section where it should lay
+ Fixed a bug when realigning sections
+ Fixed a bug when rebuilding resources


ARImpRec.DLL - 1.7.2 ---- Updated 2010 August, the 12th

+ Fixed an error when rebuilding overlay


ARImpRec.DLL - 1.7.1 ---- Updated 2010 July, the 18th

+ Fixed several bugs for relocations, imports retrieving and minimize size option
+ Added support for newest Armadillo versions


ARImpRec.DLL - 1.7.0 Beta ---- Updated 2010 April, the 25th

+ New algo for rebuilding overlay
+ Decryption of some files within pdata
+ Fixed several bugs


ARImpRec.DLL - 1.6.5 ---- Updated 2009 December, the 22th

+ Fixed bug when rebuilding imports by relocations


ARImpRec.DLL - 1.6.4 ---- Updated 2009 December, the 9th

+ An important bug fixed about IAT Size when rebuilding imports by relocations
+ Fixed a bug when rebuilding resources


ARImpRec.DLL - 1.6.3 ---- Updated 2009 December, the 6th

+ Improved detection of overlay offset
+ Fixed rebuilding resources from other section than .rsrc
+ Fixed a bug when rebuilding imports by relocations
+ Added a new type of compiled to rebuild properly all sections
+ Added more possibilities to rebuild overlay 


Previous versions
-----------------

ARImpRec.DLL - 1.6.2 ---- Updated 2009 October, the 17th

+ Improved detection of pdata section
+ Added some checks about PE header of dump
+ Fixed a bug when getting forwarded functions related to Wsock32.dll/ws2_32.dll
+ Fixed a bug that was destroying export table


ARImpRec.DLL - 1.6.1 ---- Updated 2009 October, the 11th

+ Fixed a bug related to pointers for PE header names
+ Improved the rebuilding of sections and relocations for Delphi and Borland C++ targets


ARImpRec.DLL - 1.6.0 ---- Updated 2009 October, the 11th

+ Recoded the Rebuild imports procedure to cover different compiled executables


ARImpRec.DLL - 1.5.2 ---- Updated 2009 October, the 8th

+ Fixed a bug when rebuilding original IT for VC++ targets


ARImpRec.DLL - 1.5.1 ---- Updated 2009 September, the 18th

+ Added support for UPX targets using overlay


ARImpRec.DLL - 1.5.0 ---- Updated 2009 September, the 1st

+ New approach to get overlay offsets


ARImpRec.DLL - 1.4.6 ---- Updated 2009 July, the 27th

+ Fixed a bug when rebuilding imports by using relocations
+ Added overlay detection for newest version of Armadillo


ARImpRec.DLL - 1.4.5 ---- Updated 2009 July, the 11th

+ Fixed a couple of bugs when searching for any possible overlay


ARImpRec.DLL - 1.4.4 ---- Updated 2009 June, the 20th

+ Improved code when rebuilding imports using relocations data


ARImpRec.DLL - 1.4.3 ---- Updated 2009 March, the 13th

+ Fixed bug when rebuilding imports using relocations data


ARImpRec.DLL - 1.4.2 ---- Updated 2009 February, the 26th

+ Fixed some bugs when rebuilding Visual Basic targets
+ Fixed a bug when rebuilding imports using relocations data


ARImpRec.DLL - 1.4.1 ---- Updated 2009 February, the 24th

+ Added analysis of imports using relocations data
+ Fixed some bugs when rebuilding imports


ARImpRec.DLL - 1.4.0 ---- Updated 2009 February, the 17th

+ Added support for zlib packed overlays
+ Improved rebuilding of imports, now based on relocations data, if they exist
+ Added rebuilding of VC++ 3.0 targets
+ Fixed rebuilding of Export Table


ARImpRec.DLL - 1.3.0 ---- Updated 2009 January, the 19th

+ Improved the speed of processing imports, changed the way of accessing the data and the algorithms.
+ Improved the rebuilding of section names for Armadillo 6 when using MinimizeSection.
+ Fixed some bugs for overlay targets.


ARImpRec.DLL - 1.2.2 ---- Updated 2008 November, the 29th

+ Fixed detection of code section of Armadillo for Minimized targets.


ARImpRec.DLL - 1.2.1 ---- Updated 2008 September, the 6th

+ Sorted imports
+ Fixed bug for UPX targets in the Armadillo 6 code.


ARImpRec.DLL - 1.1 ---- Released 2008 March, the 23rd
-----------------------------------------------------------------------
ARTeam Import Reconstructor
-----------------------------------------------------------------------

Abstract
-------

This tool has been coded for completing the tasks of the Armageddon tool by condzero.

The main feature is that it ignores all thunks not valid found between valid ones, and then it rearranges the imports found, rebuilding for every module an only array of thunks.

This is supposed to be in beta version, since it has been tested in a few targets.

Tested on the following compiled:
C++
Delphi
Visual Basic


How to use
----------
Two ways of proceeding:

Either:
call SearchAndRebuildImportsn@24

Or using a function from ARMinSiz.dll for repairing sections, as revirgin:

call RebuildSectionsFromArmadillo@16 function
if return is zero then
 call SearchAndRebuildImportsNoNewSection@24

New
---
It detects, when possible, the exact offset of the import table of virgin.exe. Then , the import table could be rebuilt in the same place where it was created in the first compilation by the developer.
This involves getting uniquely the functions used by the code section. So, it filters the imports found, choosing only the needed by the code section.
In this way the tool is providing a complete unwrapped target, free of protector code.
The import table gets rebuilt inside the existing sections.
These should have been fixed previously by the ARMinSiz.dll module.
It fixes relocation data in PE header, whatever the function called is.

Fixed
-----
Forwarding issues.

Known issues
------------
When using the SearchAndRebuildImportsNoNewSection@24 it takes considerably more time than using the SearchAndRebuildImports@24 function, because of the process of filtering all needed apis among all the ones found.

To be done
----------
Rebuilding import table by ordinal.

//////////////////////////////////////////////////////////////////////////////////////////////
EXPORTED FUNCTIONS
//////////////////////////////////////////////////////////////////////////////////////////////
function GetProcName(IRHModule : cardinal; 
                     IRAddress :Cardinal; 
                     var IRHInt : Cardinal; 
                     IRProcName : PChar): Integer;stdcall;

exports GetProcName name 'GetProcName@16';


It works in a similar way as the well known GetProcAddress API, but in this case you are getting the name of the function.

Input parameters:
 IRHModule - Handle of a module previously loaded in memory.
 IRAddress - Virtual Address of the function from what you want to get the Function Name.

Output parameters:
 IRProcName - Pointer to a string containing the name of a function.
 IRHInt - Address of buffer contaning the Index of the name of funtion in the array AddressOfNames of export table in the module owner of the function 
 Return of function is the length of the Name of the API.


#############################################################################################
function GetProcOrdinal(IRHModule : cardinal; 
                        IRAddress :Cardinal) : Cardinal;stdcall;

exports GetProcOrdinal name 'GetProcOrdinal@8';

It works in a similar way as the well known GetProcAddress API, but in this case you are getting the ordinal.

Input parameters:
 IRHModule - Handle of a module previously loaded in memory.
 IRAddress - Virtual Address of the function from what you want to get the Function Name.

Output parameters:
 Return of function is the ordinal thunk value. It comes in this format: 8000XXXX being XXXX the Ordinal of the function


#############################################################################################
function GetProcNameAndOrdinal(IRHModule : cardinal; 
                               IRAddress :Cardinal; 
                               var IROrdinal : Cardinal; 
                               var IRHint: Cardinal;
                               IRProcName : PChar): Integer;stdcall;

exports GetProcNameAndOrdinal name 'GetProcNameAndOrdinal@20';

It works in a similar way as the well known GetProcAddress API, but in this case you are getting the name of the function and the ordinal.

Input parameters:
 IRHModule - Handle of a module previously loaded in memory.
 IRAddress - Virtual Address of the function from what you want to get the Function Name.

Ouptut parameters:
 IROrdinal returns the ordinal.
 IRHInt - Address of buffer contaning the Index of the name of funtion in the array AddressOfNames of export table in the module owner of the function. 
 IRProcName - Pointer to a string containing the name of a function.
 Return of function is the length of the Name of the API.


#############################################################################################
function TryGetImportedFunction(IRProcessId : Cardinal;
                                IRVAddress : Cardinal;
                                IROrdinal : Cardinal;
                                var IRHint : Cardinal;                                
                                IRFunctionName : PChar;
                                IRModule : PChar): Integer;stdcall;


exports TryGetImportedFunction name 'TryGetImportedFunction@24';

It checks if a given virtual address is the handle of any imported function of the process.

Input parameters:
 IRProcessId - Identifier of a process loaded in memory.
 IRVAddress - Virtual Address what you want to get the function name for.
 IRHint - If it enters with the value 0xDEADC0DE, it check again in memory the list of modules loaded. Otherwise, it doesn't check if there is any new loaded module by the process but only the first execution of this function.

Ouptut parameters:
 IROrdinal - The ordinal of the function in the exported table of its module, always will be returned this value if it exists.
 IRHint - The hint value that shows the index of the function name in the AddressOfNamesOrdinals in the export table of the module owner of the function
 IRFunctionName - Name of imported function, if return is zero.
 IRModule - Name of module that exports the function provided.

 Return - 0 If all has gone right. 
   A function name and its hint (if exists) will be provided in the IRFunctionNameand IRHint fields respectively, its ordinal in IROrdinal and its module in IRModule . If no function name exists for that function, it will be returned as am empty string, and IRHInt will get the value 0xFFFF.
 Return - 1 There is not memory allocated for IRFunctionName field
 Return - 2 There is not memory allocated for IRModule field
 Return - 3 The process with that PiD (IRProcessId) cannot be opened
 Return - 4 Error when retrieving loaded modules information
 Return - 5 There is no module owner of that virtual address passed in IRVAddress
 Return - 6 There is no imported function for that virtual address


#############################################################################################
function SearchAndRebuildImports (IRProcessId : Cardinal;
                                  IRNameOfDumped : PChar;
                                  IROEP : Cardinal; // In VAddress
                                  IRSaveOEPToFile : Cardinal; // 1 - Yes, 0 - No
                                  var IRIATRVA : Cardinal;
                                  var IRIATSize : Cardinal;
                                  var IRWarning : PChar) : Integer;stdcall;

exports SearchAndRebuildImports name 'SearchAndRebuildImports@28';

It rebuilds imports in a file previously dumped. IAT gets rebuilt in the same place where it has been found, functions handles keep their exact offset in the file than in memory had, and Import Table is built in a new section, pasted at the end of the file.
The PE header is fixed for some needed data.

Input parameters:
 IRProcessId - Identifier of a process loaded in memory.
 IRNameOfDumped - Pointer to a string containing the complete name of a file: directory + name of file. This is the dumped from which the rebuilt is going to be generated.
 IROEP - Virtual Address of the Original Entry Point. Note that it cannot come as RVA.
 IRSaveOEPToFile - Set it to 1 if you want the OEP to be the rebuilt EP, otherwise set it to 0.
 IRIATRVA - Start address of IAT, as RVA. Zero indicates a search in the entire file.
 IRIATRVA - Size of IAT, as RVA. Zero indicates a search in the entire file.

Ouptut parameters:
 IRIATRVA - Start address of IAT, as RVA.
 IRIATRVA - Size of IAT, as RVA.
 IRWarning - Message of error, if any has been produced. When it happens, it comes with
 Return - 0 if all has gone right. If an error has been produced, it comes with a number associated to the message returned by IRWarning.


#############################################################################################
function SearchAndRebuildImportsIATOptimized (IRProcessId : Cardinal;
                                              IRNameOfDumped : PChar;
                                              IROEP : Cardinal; // In VAddress
		                              IRSaveOEPToFile : Cardinal; // 1 - Yes, 0 - No
                                              var IRIATRVA : Cardinal;
                                              var IRIATSize : Cardinal;
                                              var IRWarning : PChar) : Integer;stdcall;

exports SearchAndRebuildImports name 'SearchAndRebuildImportsIATOptimized@28';

It rebuilds imports in a file previously dumped. IAT gets rebuilt in the same place where it has been found, minimizing the IAT size by deleting bad handles and redundant zero dwords, and Import Table is built in a new section, pasted at the end of the file.
The PE header is fixed for some needed data.

Input parameters:
 IRProcessId - Identifier of a process loaded in memory.
 IRNameOfDumped - Pointer to a string containing the complete name of a file: directory + name of file. This is the dumped from which the rebuilt is going to be generated.
 IROEP - Virtual Address of the Original Entry Point. Note that it cannot come as RVA.
 IRSaveOEPToFile - Set it to 1 if you want the OEP to be the rebuilt EP, otherwise set it to 0.
 IRIATRVA - Start address of IAT, as RVA. Zero indicates a search in the entire file.
 IRIATRVA - Size of IAT, as RVA. Zero indicates a search in the entire file.

Ouptut parameters:
 IRIATRVA - Start address of IAT, as RVA.
 IRIATRVA - Size of IAT, as RVA.
 IRWarning - Message of error, if any has been produced. When it happens, it comes with
 Return - 0 if all has gone right. If an error has been produced, it comes with a number associated to the message returned by IRWarning.


#############################################################################################
function SearchAndRebuildImportsNoNewSection (IRProcessId : Cardinal;
                                              IRNameOfDumped : PChar;
                                              IROEP : Cardinal; // In VAddress
		                              IRSaveOEPToFile : Cardinal; // 1 - Yes, 0 - No
                                              var IRIATRVA : Cardinal;
                                              var IRIATSize : Cardinal;
                                              var IRWarning : PChar) : Integer;stdcall;

exports SearchAndRebuildImportsNoNewSection name 'SearchAndRebuildImportsNoNewSection@28';

It rebuilds imports in a file previously dumped. IAT gets rebuilt in the place where it was in the virgin.exe, and Import Table is built in the same place where it was in virgin.exe, too. So, no new section is appended.
This has been coded in that way to let the flash and macromedia overlays get in their exact offset. This wouldn't be possible keeping all sections added by Armadillo.

The PE header is fixed for some needed data.

Input parameters:
 IRProcessId - Identifier of a process loaded in memory.
 IRNameOfDumped - Pointer to a string containing the complete name of a file: directory + name of file. This is the dumped from which the rebuilt is going to be generated.
 IROEP - Virtual Address of the Original Entry Point. Note that it cannot come as RVA.
 IRSaveOEPToFile - Set it to 1 if you want the OEP to be the rebuilt EP, otherwise set it to 0.
 IRIATRVA - Start address of IAT, as RVA. Zero indicates a search in the entire file.
 IRIATRVA - Size of IAT, as RVA. Zero indicates a search in the entire file.

Ouptut parameters:
 IRIATRVA - Start address of IAT, as RVA.
 IRIATRVA - Size of IAT, as RVA.
 IRWarning - Message of error, if any has been produced. When it happens, it comes with
 Return - 0 if all has gone right. If an error has been produced, it comes with a number associated to the message returned by IRWarning.


#############################################################################################
function RebuildSectionsFromArmadillo  (MSNameOfProtected : PChar;
                                        MSNameOfDumped : PChar;
                                        MSWarning : PChar) : Integer;stdcall;

exports RebuildSectionsFromArmadillo name 'RebuildSectionsFromArmadillo@12';


You should invoke it in your code as 'RebuildSectionsFromArmadillo@12'.

Input parameters:
 MSNameOfProtected - Address of buffer contaning the path of your Armadillo protected target.
 MSNameOfDumped - Address of buffer contaning the path of your Armadillo dumped target to be processed by the tool.
 MSWarning - Pointer to an allocated buffer of Char, it is recommended a minimum size of 255 bytes.

Output parameters:
 MSWarning - returns the error message, if there is any, or a success message.
 Return of function is an error code, different for every message, 0 in case of success or -1 if there is a not valid allocated buffer as input parameter.

##############################################################################################
function GetNameFileOptimized (MSFileNameOrig : PChar; MSFileNameOptimized: PChar) : Integer;stdcall;

exports GetNameFileOptimized name 'GetNameFileOptimized@8';


You should invoke it in your code as 'GetNameFileOptimized@8'.

Input parameters:
 MSFileNameOrig - Address of buffer contaning the path of your Armadillo dumped target to be processed by the tool.
 MSFileNameOptimized - Pointer to an allocated buffer of Char, it is recommended a minimum size of 255 bytes.

Output parameters:
 MSFileNameOptimized is a pointer to the name of the target created by the RebuildSectionsFromArmadillo function.
 Return of function is the length of the string returned by MSFileNameOptimizedname, or - 1 if an error has happened.

##############################################################################################
function UnpackPdataSection(MSNameOfProtected : PChar;
                            MSNameOfDumped : PChar;
                            MSWarning : PChar): Integer;stdcall;


exports UnpackPdataSection name 'UnpackPdataSection@12';

You should invoke it in your code as 'UnpackPdataSection@12'.

Input parameters:
 MSNameOfProtected - Address of buffer contaning the path of your Armadillo protected target.
 MSNameOfDumped - Address of buffer contaning the path of your Armadillo dumped target to be processed by the tool.

Output parameters:
 MSWarning - returns the error message, if there is any, or a success message with the name of unpacked file.
 Return of function is an error code, different for every message, 0 in case of success or -1 if there is a not valid allocated buffer as input parameter.


You should invoke any of the previous functions from your code as is indicated in exports line.

For instance, you would code this:
call GetProcName@8 


-----------------------------------------------------------------------------------------
- History -
ARTeam Import Reconstructor 1.0 - Released 2008/02/26

For support in rebuilding DLL's, now OEP should come in Virtual Address (not RVA as before).

Release 2008/01/30
This tool has been coded for completing the tasks of the Armageddon tool by condzero.

The main feature is that it ignores all thunks not valid found between valid ones, and then it rearranges the imports found, rebuilding for every module an only array of thunks.

This is supposed to be in beta version, since it has been tested in a few targets.


I hope all of you could use it for a width range of uses and protections.

I know, speed could be improved, just let me play a little with the tool to find the way.

If you find useful, please report all bugs and desiderable improvement.


This tool is devoted to all my friends of ARTeam.

Coded in Delphi 7 Enterprise.

Nacho_dj / ARteam