package header
//back\HyperDbgDev\hyperdbg\hprdbgctrl\header\install.h.back

const(
DRIVER_FUNC_INSTALL = 0x01 //col:52
DRIVER_FUNC_STOP =    0x02 //col:53
DRIVER_FUNC_REMOVE =  0x03 //col:54
)

type (
Install interface{
#    define HPRDBGCTRL_API __declspec()(ok bool)//col:40
}
)

func NewInstall() { return & install{} }

func (i *install)#    define HPRDBGCTRL_API __declspec()(ok bool){//col:40
/*#    define HPRDBGCTRL_API __declspec(dllexport)
#else
#    define HPRDBGCTRL_API __declspec(dllimport)
#endif
class HPRDBGCTRL_API Chprdbgctrl
{
public:
    Chprdbgctrl(void);
};*/
return true
}



