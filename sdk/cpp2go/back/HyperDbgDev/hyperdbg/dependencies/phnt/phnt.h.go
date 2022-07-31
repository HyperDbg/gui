package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\phnt.h.back

const(
_PHNT_H =  //col:13
PHNT_MODE_KERNEL = 0 //col:34
PHNT_MODE_USER = 1 //col:35
PHNT_WIN2K = 50 //col:38
PHNT_WINXP = 51 //col:39
PHNT_WS03 = 52 //col:40
PHNT_VISTA = 60 //col:41
PHNT_WIN7 = 61 //col:42
PHNT_WIN8 = 62 //col:43
PHNT_WINBLUE = 63 //col:44
PHNT_THRESHOLD = 100 //col:45
PHNT_THRESHOLD2 = 101 //col:46
PHNT_REDSTONE = 102 //col:47
PHNT_REDSTONE2 = 103 //col:48
PHNT_REDSTONE3 = 104 //col:49
PHNT_REDSTONE4 = 105 //col:50
PHNT_REDSTONE5 = 106 //col:51
PHNT_19H1 = 107 //col:52
PHNT_19H2 = 108 //col:53
PHNT_20H1 = 109 //col:54
PHNT_20H2 = 110 //col:55
PHNT_21H1 = 111 //col:56
PHNT_21H2 = 112 //col:57
PHNT_WIN11 = 113 //col:58
PHNT_MODE = PHNT_MODE_USER //col:61
PHNT_VERSION = PHNT_WIN7 //col:65
// = PHNT_NO_INLINE_INIT_STRING //col:70
)

type (
Phnt interface{
 * Attribution 4.0 International ()(ok bool)//col:121
}
)

func NewPhnt() { return & phnt{} }

func (p *phnt) * Attribution 4.0 International ()(ok bool){//col:121
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
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
#include <phnt_ntdef.h>
#include <ntnls.h>
#include <ntkeapi.h>
#endif
#include <ntldr.h>
#include <ntexapi.h>
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
}*/
return true
}


