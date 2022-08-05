package sdk

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"strings"
	"testing"
)

/*
uname -a
sudo pacman  -S linux-headers
sudo pacman  -S linux515-headers
yay -S vmware-workstation
sudo modprobe -a vmw_vmci vmmon
sudo systemctl enable vmware-networks.service
sudo systemctl start  vmware-networks.service
sudo pacman -Sy open-vm-tools
*/
func TestName(t *testing.T) {
	run, err := cmd.Run("uname -a")
	if !mycheck.Error(err) {
		return
	}
	split := strings.Split(run, " ")
	s := split[2]
	before, _, found := strings.Cut(s, "-")
	if !found {
		return
	}
	index := strings.LastIndex(before, ".")
	before = before[:index]
	before = strings.ReplaceAll(before, ".", "")
	run, err = cmd.Run(`sudo pacman  -Sy linux` + before + `-headers`)
	if !mycheck.Error(err) {
		return
	}
	println(run)
}

/*
https://github.com/vlang/v.git
https://github.com/goplus/c2go.git
https://gitlab.com/cznic/builder
https://gitlab.com/cznic/ccgo
https://gitlab.com/cznic/cc
https://gitlab.com/cznic/memory
https://gitlab.com/cznic/sqlite
https://gitlab.com/cznic/ql
https://gitlab.com/cznic/libc
https://gitlab.com/cznic/gc
https://gitlab.com/cznic/parser
git clone --recursive -b dev https://github.com/HyperDbg/HyperDbg.git
https://github.com/vlang/tccbin/tree/thirdparty-windows-amd64

sudo pacman -S github-cli
gh auth login
sudo gh extension install github/gh-net

sudo pacman -S openssh-server





*/
