package sdk

type (
	Objects interface {
		ObjectShowProcessesOrThreadList() (ok bool)
		ObjectShowProcessesOrThreadDetails() (ok bool)
	}
	objects struct{}
)

func (o *objects) ObjectShowProcessesOrThreadList() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *objects) ObjectShowProcessesOrThreadDetails() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func newObjects() Objects {
	return &objects{}
}
