//go:build !windows

package bindlib

import (
	"crypto/sha1"
	"encoding/hex"
	"os"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ebitengine/purego"
)

type solib struct {
	Handle uintptr
}

func (w solib) Lookup(name string) uintptr {
	return mylog.Check2(purego.Dlsym(w.Handle, name))
}

func LoadLibrary(name string) (LoadedLibrary, error) {
	h := mylog.Check2(purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_LOCAL))
	// if err == nil {
	return solib{Handle: h}, nil
	// }
	// if !strings.ContainsAny(name, "/\\") {
	// 	h, elocal := purego.Dlopen("./"+name, purego.RTLD_NOW|purego.RTLD_LOCAL)
	// 	if elocal == nil {
	// 		return solib{Handle: h}, nil
	// 	}
	// }
	// return nil, err
}

func FindLibrary(name string) (LoadedLibrary, error) {
	lib := mylog.Check2(LoadLibrary(name))
	// if err == nil {
	return lib, nil
	// }
	// org := err
	// if !strings.HasSuffix(name, ".so") {
	// 	name += ".so"
	// 	if lib = mylog.Check2(LoadLibrary(name)); err == nil {
	// 		return lib, nil
	// 	}
	// }
	// if !strings.HasPrefix(name, "lib") {
	// 	name = "lib" + name
	// 	if lib = mylog.Check2(LoadLibrary(name)); err == nil {
	// 		return lib, nil
	// 	}
	// }
	// return nil, org
}

func LoadLibraryEmbed(data []byte) (LoadedLibrary, error) {
	cache := getTmpDir()
	hash := sha1.Sum(data)
	name := "." + hex.EncodeToString(hash[:4]) + ".gengo.so"
	path := cache + name
	if stat, e := os.Stat(path); e != nil || stat.Size() != int64(len(data)) {
		os.MkdirAll(cache, 0o755)
		mylog.Check(os.WriteFile(path, data, 0o755))
	}
	return LoadLibrary(path)
}
