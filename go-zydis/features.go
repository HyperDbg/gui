// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

/*
#cgo CFLAGS: -I./lib/include
#include <Zydis/Zydis.h>

int zVerMajor(ZyanU64 v) { return ZYDIS_VERSION_MAJOR(v); }
int zVerMinor(ZyanU64 v) { return ZYDIS_VERSION_MINOR(v); }
int zVerBuild(ZyanU64 v) { return ZYDIS_VERSION_PATCH(v); }
int zVerPatch(ZyanU64 v) { return ZYDIS_VERSION_BUILD(v); }
*/
import "C"
import (
	"fmt"
)

// Version returns the zydis library version number.
func Version() (major, minor, patch, build int) {
	v := C.ZydisGetVersion()
	return int(C.zVerMajor(v)),
		int(C.zVerMinor(v)),
		int(C.zVerBuild(v)),
		int(C.zVerPatch(v))
}

// Feature is an enum of zydis library features.
type Feature int

// Feature enum values.
const (
	FeatureDecoder Feature = iota
	FeatureFormatter
	FeatureAVX512
	FeatureKNC
)

// IsFeatureEnabled returns true when the given feature is included in
// the zydis library.
func IsFeatureEnabled(f Feature) (bool, error) {
	ret := C.ZydisIsFeatureEnabled(C.ZydisFeature(f))
	if zyanFailure(ret) {
		return false, fmt.Errorf("zydis: error checking for feature: 0x%x", ret)
	}
	return ret == C.ZYAN_STATUS_TRUE, nil
}
