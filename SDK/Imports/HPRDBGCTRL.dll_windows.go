package Imports

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"syscall"
)

type (
	Api struct{ procMap map[nameKind]*syscall.Proc }
)

var api = newApi()

func Call(p *syscall.Proc, a ...uintptr) (value uintptr) {
	value, statusCode, err := p.Call(a...)
	if !mylog.Error(err) {
		return
	}
	if statusCode == 0 {
		mylog.Error("statusCode == 0") //?
		return
	}
	return value
}

func newApi() *Api {
	p := &Api{
		procMap: make(map[nameKind]*syscall.Proc, HyperDbgLoadVmm.Len()),
	}
	for _, kind := range HyperDbgLoadVmm.Names() {
		p.SetProc(kind)
	}
	return p
}
func (a *Api) Proc(name nameKind) *syscall.Proc { return a.procMap[name] }
func (a *Api) SetProc(name nameKind)            { a.procMap[name] = dll.MustFindProc(name.String()) }

var (
	//go:embed HPRDBGCTRL.dll
	dllData []byte
	dll     *syscall.DLL
)

func init() {
	dir, err := os.UserCacheDir()
	if !mylog.Error(err) {
		return
	}
	dir = filepath.Join(dir, "hyperdbgDll", "dll_cache")
	if !mylog.Error(os.MkdirAll(dir, 0755)) {
		return
	}
	if !mylog.Error(windows.SetDllDirectory(dir)) {
		return
	}
	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("hyperdbgDll-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	_, err = os.Stat(filePath)
	if !mylog.Error(err) {
		if err == os.ErrNotExist {
			if !mylog.Error(os.WriteFile(filePath, dllData, 0644)) {
				return
			}
		}
		return
	}
	dll = syscall.MustLoadDLL(dllName)
}
