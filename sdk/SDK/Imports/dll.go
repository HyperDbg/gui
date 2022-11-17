package Imports

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	//go:embed HPRDBGCTRL.dll
	dllData                         []byte
	grBackendRenderTargetNewGLProc  *syscall.Proc
	grBackendRenderTargetDeleteProc *syscall.Proc
)

func init() {
	dir, err := os.UserCacheDir()
	jot.FatalIfErr(err)
	dir = filepath.Join(dir, "unison", "dll_cache")
	jot.FatalIfErr(os.MkdirAll(dir, 0755))
	windows.SetDllDirectory(dir)
	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("skia-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	if !fs.FileExists(filePath) {
		jot.FatalIfErr(os.WriteFile(filePath, dllData, 0644))
	}
	skia := syscall.MustLoadDLL(dllName)
	grBackendRenderTargetNewGLProc = skia.MustFindProc("gr_backendrendertarget_new_gl")
	grBackendRenderTargetDeleteProc = skia.MustFindProc("gr_backendrendertarget_delete")
}

type (
	BackendRenderTarget  unsafe.Pointer
	DirectContext        unsafe.Pointer
	GLInterface          unsafe.Pointer
	Canvas               unsafe.Pointer
	ColorFilter          unsafe.Pointer
	ColorSpace           unsafe.Pointer
	Data                 unsafe.Pointer
	Document             unsafe.Pointer
	DynamicMemoryWStream unsafe.Pointer
	FileWStream          unsafe.Pointer
	Font                 unsafe.Pointer
	FontMgr              unsafe.Pointer
	FontStyle            unsafe.Pointer
	FontStyleSet         unsafe.Pointer
	Image                unsafe.Pointer
	ImageFilter          unsafe.Pointer
	MaskFilter           unsafe.Pointer
	OpBuilder            unsafe.Pointer
	Paint                unsafe.Pointer
	Path                 unsafe.Pointer
	PathEffect           unsafe.Pointer
	SamplingOptions      uintptr
	Shader               unsafe.Pointer
	String               unsafe.Pointer
	Surface              unsafe.Pointer
	SurfaceProps         unsafe.Pointer
	TextBlob             unsafe.Pointer
	TextBlobBuilder      unsafe.Pointer
	TypeFace             unsafe.Pointer
	WStream              unsafe.Pointer
)

func BackendRenderTargetDelete(backend BackendRenderTarget) {
	grBackendRenderTargetDeleteProc.Call(uintptr(backend))
}

//func ContextMakeGL(gl GLInterface) DirectContext {
//	r1, _, _ := grContextMakeGLProc.Call(uintptr(gl))
//	return DirectContext(r1)
//}
