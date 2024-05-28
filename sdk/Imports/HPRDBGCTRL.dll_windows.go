package Imports

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/windows"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

type (
	Api struct{ procMap map[nameKind]*syscall.Proc }
)

var api = newApi()

func Call(p *syscall.Proc, a ...uintptr) (value uintptr) {
	value, statusCode := mylog.Check3(p.Call(a...))

	if statusCode == 0 {
		mylog.Check("statusCode == 0") //?
		return
	}
	return value
}

func newApi() *Api {
	if dll == nil {
		return nil
		Init()
	}
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

func Init() {
	defer func() {
		if dll == nil {
			// panic("dll not load")
		}
	}()
	dir := mylog.Check2(os.UserCacheDir())

	dir = filepath.Join(dir, "hyperdbgDll", "dll_cache")
	mylog.Check(os.MkdirAll(dir, 0o755))
	mylog.Check(windows.SetDllDirectory(dir))
	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("hyperdbgDll-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	stream.WriteTruncate(filePath, dllData)
	dll = syscall.MustLoadDLL(filePath)
}
