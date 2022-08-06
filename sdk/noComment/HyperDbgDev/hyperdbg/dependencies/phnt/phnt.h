/*
 * This file is part of the Process Hacker project -
 * https://processhacker.sourceforge.io/
 *
 * You can redistribute this file and/or modify it under the terms of the
 * Attribution 4.0 International (CC BY 4.0) license.
 *
 * You must give appropriate credit, provide a link to the license, and
 * indicate if changes were made. You may do so in any reasonable manner, but
 * not in any way that suggests the licensor endorses you or your use.
 */
#ifndef _PHNT_H
#define _PHNT_H
#define PHNT_MODE_KERNEL 0
#define PHNT_MODE_USER 1
#define PHNT_WIN2K 50
#define PHNT_WINXP 51
#define PHNT_WS03 52
#define PHNT_VISTA 60
#define PHNT_WIN7 61
#define PHNT_WIN8 62
#define PHNT_WINBLUE 63
#define PHNT_THRESHOLD 100
#define PHNT_THRESHOLD2 101
#define PHNT_REDSTONE 102
#define PHNT_REDSTONE2 103
#define PHNT_REDSTONE3 104
#define PHNT_REDSTONE4 105
#define PHNT_REDSTONE5 106
#define PHNT_19H1 107
#define PHNT_19H2 108
#define PHNT_20H1 109
#define PHNT_20H2 110
#define PHNT_21H1 111
#define PHNT_21H2 112
#define PHNT_WIN11 113
#ifndef PHNT_MODE
#define PHNT_MODE PHNT_MODE_USER
#endif
#ifndef PHNT_VERSION
#define PHNT_VERSION PHNT_WIN7
#endif
#ifdef __cplusplus
extern "C" {
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#include <ntkeapi.h>
#include <ntnls.h>
#include <phnt_ntdef.h>
#endif
#include <ntexapi.h>
#include <ntldr.h>
#include <ntbcd.h>
#include <ntmmapi.h>
#include <ntobapi.h>
#include <ntpsapi.h>
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#include <cfg.h>
#include <ntdbg.h>
#include <ntioapi.h>
#include <ntlpcapi.h>
#include <ntpfapi.h>
#include <ntpnpapi.h>
#include <ntpoapi.h>
#include <ntregapi.h>
#include <ntrtl.h>
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#include <ntseapi.h>
#include <nttmapi.h>
#include <nttp.h>
#include <ntxcapi.h>
#include <ntwow64.h>
#include <ntlsa.h>
#include <ntsam.h>
#include <ntmisc.h>
#include <ntzwapi.h>
#endif
#ifdef __cplusplus
}

#endif
#endif
