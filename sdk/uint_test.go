package sdk

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/Comment"
	"github.com/ddkwork/librarygo/src/myc2go"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	cc "modernc.org/cc/v5"
	"strings"
	"testing"
)

func TestName2(t *testing.T) {
	parse, err := cc.Parse(&cc.Config{
		ABI:                 nil,
		CC:                  "",
		FS:                  nil,
		HostIncludePaths:    nil,
		HostSysIncludePaths: nil,
		IncludePaths:        nil,
		PragmaHandler:       nil,
		Predefined:          "",
		SysIncludePaths:     nil,
	},
		[]cc.Source{
			cc.Source{
				Name:  "",
				Value: nil,
				FS:    nil,
			},
		})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(exampleAST(9, "\U00100001 'a' ( 'b' , 'c' )"))
}

func TestCpp2Go(t *testing.T) {
	Comment.New().Delete("HyperDbgDev/hyperdbg")
	return
	o := myc2go.NewObj()
	object := o.Src("./HyperDbgDev/hyperdbg").
		Dst("binding").
		ExpandPath("miscellaneous\\constants", ".txt").
		Back()
	if object == nil {
		return
	}
	o.Convert()
	//o.Format()
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
