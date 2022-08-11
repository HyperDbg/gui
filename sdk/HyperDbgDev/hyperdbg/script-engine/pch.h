/**
 * @file pch.h
 * @author M.H. Gholamrezaei (mh@hyperdbg.org)
 *
 * @details Pre-compiled headers
 * @version 0.1
 * @date 2020-10-22
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//
// Exclude rarely-used stuff from Windows headers
//
#define WIN32_LEAN_AND_MEAN

//
// Windows Header Files
//
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>

#include "..\script-eval\header\ScriptEngineCommonDefinitions.h"
#include "SDK/Imports/HyperDbgSymImports.h"
#include "common.h"
#include "globals.h"
#include "parse-table.h"
#include "scanner.h"
#include "script-engine.h"
