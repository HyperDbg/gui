package kdserial
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\common.c.back

type (
Common interface{
DllInitialize()(ok bool)//col:5
DllUnload()(ok bool)//col:9
}
)

func NewCommon() { return & common{} }

func (c *common)DllInitialize()(ok bool){//col:5
/*DllInitialize(
    _In_ PUNICODE_STRING RegistryPath)
{
    return STATUS_SUCCESS;
}*/
return true
}

func (c *common)DllUnload()(ok bool){//col:9
/*DllUnload(void)
{
    return STATUS_SUCCESS;
}*/
return true
}



