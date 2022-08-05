package sdk

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"testing"
)

/*
uname -a
sudo pacman  -S linux-headers
yay -S vmware-workstation
sudo modprobe -a vmw_vmci vmmon
sudo systemctl enable vmware-networks.service
sudo systemctl start  vmware-networks.service
sudo pacman -Sy open-vm-tools
*/
func TestName(t *testing.T) {
	fmt.Println(cmd.Run("uname -a"))
}
