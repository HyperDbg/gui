package transparency


const(
MSR_IA32_TIME_STAMP_COUNTER = 0x10 //col:35
RAND_MAX = 0x7fff //col:41
)

type (
Transparency interface{
TransparentModeStart()(ok bool)//col:59
}













































)

func NewTransparency() { return & transparency{} }

func (t *transparency)TransparentModeStart()(ok bool){//col:59














return true
}
















































