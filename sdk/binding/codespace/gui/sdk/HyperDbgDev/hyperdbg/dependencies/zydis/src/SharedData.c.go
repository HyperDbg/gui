package src



type (
	SharedData interface {
		void_ZydisGetInstructionDefinition() (ok bool) //col:31
	}
	sharedData struct{}
)

func NewSharedData() SharedData { return &sharedData{} }

func (s *sharedData) void_ZydisGetInstructionDefinition() (ok bool) { //col:31









































	return true
}


