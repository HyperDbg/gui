package ept


const(
MaximumHiddenBreakpointsOnPage = 40 //col:18
PAGE_ATTRIB_READ =  0x2 //col:28
PAGE_ATTRIB_WRITE = 0x4 //col:29
PAGE_ATTRIB_EXEC =  0x8 //col:30
VMM_EPT_PML4E_COUNT = 512 //col:36
VMM_EPT_PML3E_COUNT = 512 //col:42
VMM_EPT_PML2E_COUNT = 512 //col:49
VMM_EPT_PML1E_COUNT = 512 //col:56
SIZE_2_MB = ((SIZE_T)(512 * PAGE_SIZE)) //col:62
ADDRMASK_EPT_PML1_OFFSET(_VAR_) = (_VAR_ & 0xFFFULL) //col:68
ADDRMASK_EPT_PML1_INDEX(_VAR_) = ((_VAR_ & 0x1FF000ULL) >> 12) //col:74
ADDRMASK_EPT_PML2_INDEX(_VAR_) = ((_VAR_ & 0x3FE00000ULL) >> 21) //col:80
ADDRMASK_EPT_PML3_INDEX(_VAR_) = ((_VAR_ & 0x7FC0000000ULL) >> 30) //col:86
ADDRMASK_EPT_PML4_INDEX(_VAR_) = ((_VAR_ & 0xFF8000000000ULL) >> 39) //col:92
EPT_MTRR_RANGE_DESCRIPTOR_MAX = 0x9 //col:165
)

type (
Ept interface{
    DECLSPEC_ALIGN()(ok bool)//col:148
    DECLSPEC_ALIGN()(ok bool)//col:207
    DECLSPEC_ALIGN()(ok bool)//col:296
}













































)

func NewEpt() { return & ept{} }

func (e *ept)    DECLSPEC_ALIGN()(ok bool){//col:148







return true
}














































func (e *ept)    DECLSPEC_ALIGN()(ok bool){//col:207









return true
}














































func (e *ept)    DECLSPEC_ALIGN()(ok bool){//col:296

















return true
}
















































