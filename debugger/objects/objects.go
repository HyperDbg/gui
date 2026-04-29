package objects

import "fmt"

type ObjectType uint32

const (
	ObjectTypeProcess ObjectType = iota + 1
	ObjectTypeThread
)

type ObjectInfo struct {
	Type     ObjectType
	Id       uint32
	Name     string
	ParentId uint32
	Address  uint64
	Cr3      uint64
	Is64Bit  bool
}

func GetObjectList(objectType ObjectType) ([]ObjectInfo, error) {
	return nil, fmt.Errorf("not implemented: requires IOCTL")
}

func GetObjectById(id uint32, objectType ObjectType) (*ObjectInfo, error) {
	return nil, fmt.Errorf("not implemented: requires IOCTL")
}

func GetObjectByName(name string, objectType ObjectType) (*ObjectInfo, error) {
	return nil, fmt.Errorf("not implemented: requires IOCTL")
}
