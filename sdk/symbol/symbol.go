package symbol

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object   struct{}
	Ntoskrnl interface {
	}
)

func New() Interface { return &object{} }
