package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\common.c.back

type (
Common interface{
DllInitialize()(ok bool)//col:15
DllUnload()(ok bool)//col:21
}
)

func NewCommon() { return & common{} }

func (c *common)DllInitialize()(ok bool){//col:15
/*DllInitialize(
    _In_ PUNICODE_STRING RegistryPath)
{
    return STATUS_SUCCESS;
}*/
return true
}

func (c *common)DllUnload()(ok bool){//col:21
/*DllUnload(void)
{
    return STATUS_SUCCESS;
}*/
return true
}



