// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// ISAExt is an enum of Instruction Set Architecture Extensions.
type ISAExt int

//go:generate stringer -type=ISAExt -linecomment

// IsaExt enum values.
const (
	ISAExtInvalid           ISAExt = iota // INVALID
	IsaExtADOX_ADCX                       // ADOX_ADCX
	IsaExtAES                             // AES
	IsaExtAMD3DNOW                        // AMD3DNOW
	IsaExtAMD3DNOW_PREFETCH               // AMD3DNOW_PREFETCH
	IsaExtAMD_INVLPGB                     // AMD_INVLPGB
	IsaExtAMX_BF16                        // AMX_BF16
	IsaExtAMX_INT8                        // AMX_INT8
	IsaExtAMX_TILE                        // AMX_TILE
	IsaExtAVX                             // AVX
	IsaExtAVX2                            // AVX2
	IsaExtAVX2GATHER                      // AVX2GATHER
	IsaExtAVX512EVEX                      // AVX512EVEX
	IsaExtAVX512VEX                       // AVX512VEX
	IsaExtAVXAES                          // AVXAES
	IsaExtBASE                            // BASE
	IsaExtBMI1                            // BMI1
	IsaExtBMI2                            // BMI2
	IsaExtCET                             // CET
	IsaExtCLDEMOTE                        // CLDEMOTE
	IsaExtCLFLUSHOPT                      // CLFLUSHOPT
	IsaExtCLFSH                           // CLFSH
	IsaExtCLWB                            // CLWB
	IsaExtCLZERO                          // CLZERO
	IsaExtENQCMD                          // ENQCMD
	IsaExtF16C                            // F16C
	IsaExtFMA                             // FMA
	IsaExtFMA4                            // FMA4
	IsaExtGFNI                            // GFNI
	IsaExtINVPCID                         // INVPCID
	IsaExtKNC                             // KNC
	IsaExtKNCE                            // KNCE
	IsaExtKNCV                            // KNCV
	IsaExtLONGMODE                        // LONGMODE
	IsaExtLZCNT                           // LZCNT
	IsaExtMCOMMIT                         // MCOMMIT
	IsaExtMMX                             // MMX
	IsaExtMONITOR                         // MONITOR
	IsaExtMONITORX                        // MONITORX
	IsaExtMOVBE                           // MOVBE
	IsaExtMOVDIR                          // MOVDIR
	IsaExtMPX                             // MPX
	IsaExtPADLOCK                         // PADLOCK
	IsaExtPAUSE                           // PAUSE
	IsaExtPCLMULQDQ                       // PCLMULQDQ
	IsaExtPCONFIG                         // PCONFIG
	IsaExtPKU                             // PKU
	IsaExtPREFETCHWT1                     // PREFETCHWT1
	IsaExtPT                              // PT
	IsaExtRDPID                           // RDPID
	IsaExtRDPRU                           // RDPRU
	IsaExtRDRAND                          // RDRAND
	IsaExtRDSEED                          // RDSEED
	IsaExtRDTSCP                          // RDTSCP
	IsaExtRDWRFSGS                        // RDWRFSGS
	IsaExtRTM                             // RTM
	IsaExtSERIALIZE                       // SERIALIZE
	IsaExtSGX                             // SGX
	IsaExtSGX_ENCLV                       // SGX_ENCLV
	IsaExtSHA                             // SHA
	IsaExtSMAP                            // SMAP
	IsaExtSMX                             // SMX
	IsaExtSNP                             // SNP
	IsaExtSSE                             // SSE
	IsaExtSSE2                            // SSE2
	IsaExtSSE3                            // SSE3
	IsaExtSSE4                            // SSE4
	IsaExtSSE4A                           // SSE4A
	IsaExtSSSE3                           // SSSE3
	IsaExtSVM                             // SVM
	IsaExtTBM                             // TBM
	IsaExtTSX_LDTRK                       // TSX_LDTRK
	IsaExtVAES                            // VAES
	IsaExtVMFUNC                          // VMFUNC
	IsaExtVPCLMULQDQ                      // VPCLMULQDQ
	IsaExtVTX                             // VTX
	IsaExtWAITPKG                         // WAITPKG
	IsaExtX87                             // X87
	IsaExtXOP                             // XOP
	IsaExtXSAVE                           // XSAVE
	IsaExtXSAVEC                          // XSAVEC
	IsaExtXSAVEOPT                        // XSAVEOPT
	IsaExtXSAVES                          // XSAVES
)
