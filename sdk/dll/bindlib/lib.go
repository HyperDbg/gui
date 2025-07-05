package bindlib

import (
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/ddkwork/golibrary/std/mylog"
)

type LoadedLibrary interface {
	Lookup(name string) uintptr
}

// func LoadLibrary(name string) (LoadedLibrary, error)
// func FindLibrary(name string) (LoadedLibrary, error)
// func LoadLibraryEmbed(data []byte) (LoadedLibrary, error)

type Library struct {
	Name  string
	cache LoadedLibrary
	done  atomic.Bool
	mu    sync.Mutex
}

func NewLibrary(name string) *Library {
	return &Library{
		Name: name,
	}
}

func (l *Library) Assign(lib LoadedLibrary) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.cache = lib
	l.done.Store(true)
}

func (l *Library) LoadFrom(path string) (LoadedLibrary, error) {
	if l.done.Load() {
		return l.cache, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.done.Load() {
		return l.cache, nil
	}

	loaded := mylog.Check2(LoadLibrary(path))

	l.cache = loaded
	l.done.Store(true)
	return loaded, nil
}

func (l *Library) LoadEmbed(data []byte) (LoadedLibrary, error) {
	loaded := mylog.Check2(LoadLibraryEmbed(data))
	l.cache = loaded
	l.done.Store(true)
	return loaded, nil
}

func (l *Library) Get() (LoadedLibrary, error) {
	return l.LoadFrom(l.Name)
}

func (l *Library) Import(name string) Proc {
	return Proc{
		library: l,
		name:    name,
	}
}

func (l *Library) ImportNow(name string) PreloadProc {
	i := l.Import(name)
	return PreloadProc(i.Addr())
}

var getTmpDir = sync.OnceValue(func() string {
	if cache, e := os.UserCacheDir(); e == nil {
		return cache + string(os.PathSeparator)
	}
	if exec, e := os.Executable(); e == nil {
		return filepath.Dir(exec) + string(os.PathSeparator)
	}
	return os.TempDir() + string(os.PathSeparator)
})
