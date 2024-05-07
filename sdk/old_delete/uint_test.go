package sdk

import (
	"strings"
	"testing"

	"github.com/ddkwork/cpp2go"
)

func TestCpp2Go(t *testing.T) {
	cpp2go.New().RemoveComment("HyperDbgDev/hyperdbg")
	return
	//o := cpp2go.New()
	//object := o.Src("./HyperDbgDev/hyperdbg").
	//	Dst("binding").
	//	ExpandPath("miscellaneous\\constants", ".txt").
	//	Back()
	//if object == nil {
	//	return
	//}
	//o.Convert()
	//o.String()
}

/*
mkdir cc
cd cc
git clone --recursive https://github.com/vlang/v.git
git clone --recursive https://github.com/goplus/c2go.git
git clone --recursive https://gitlab.com/cznic/builder
git clone --recursive https://gitlab.com/cznic/ccgo
git clone --recursive https://gitlab.com/cznic/cc
git clone --recursive https://gitlab.com/cznic/memory
git clone --recursive https://gitlab.com/cznic/sqlite
git clone --recursive https://gitlab.com/cznic/ql
git clone --recursive https://gitlab.com/cznic/libc
git clone --recursive https://gitlab.com/cznic/gc
git clone --recursive https://gitlab.com/cznic/parser
git clone --recursive -b dev https://github.com/HyperDbg/HyperDbg.git
git clone --recursive https://github.com/vlang/tccbin/tree/thirdparty-windows-amd64
*/

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
	run := stream.RunCommand("uname -a")
	split := strings.Split(run.Result, " ")
	s := split[2]
	before, _, found := strings.Cut(s, "-")
	if !found {
		return
	}
	index := strings.LastIndex(before, ".")
	before = before[:index]
	before = strings.ReplaceAll(before, ".", "")
	stream.RunCommand(`sudo pacman  -Sy linux` + before + `-headers`)
}
