//go:build windows

package hook

import (
	"fmt"

	"github.com/ddkwork/golibrary/std/mylog"
	wapi "github.com/iamacarpet/go-win64api"
)

func GetServices() {
	svc := mylog.Check2(wapi.GetServices())

	for _, v := range svc {
		fmt.Printf("%-50s - %-75s - Status: %-20s - Accept Stop: %-5t, Running Pid: %d\r\n", v.SCName, v.DisplayName, v.StatusText, v.AcceptStop, v.RunningPid)
	}
}

func StartService() {
	mylog.Check(wapi.StartService("service_name"))
}

func InstalledSoftwareList() {
	sw := mylog.Check2(wapi.InstalledSoftwareList())

	for _, s := range sw {
		fmt.Printf("%-100s - %s - %s\r\n", s.Name(), s.Architecture(), s.Version())
	}
}

func ProcessList() {
	pr := mylog.Check2(wapi.ProcessList())

	for _, p := range pr {
		fmt.Printf("%8d - %-30s - %-30s - %s\r\n", p.Pid, p.Username, p.Executable, p.Fullpath)
	}
}
