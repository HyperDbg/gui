package driver
type (
	Ioctl interface {
		DrvDispatchIoControl() (ok bool) //col:1453
	}
)
func NewIoctl() { return &ioctl{} }
func (i *ioctl) DrvDispatchIoControl() (ok bool) { //col:1453
	return true
}

