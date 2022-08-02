package examples


type typedef struct TestStruct_ struct{
u32 ZyanU32 //col:3
u64 ZyanU64 //col:4
f float //col:5
}




type (
Vector interface{
static_void_InitTestdata()(ok bool)//col:5
static_ZyanStatus_PerformBasicTests()(ok bool)//col:40
static_ZyanStatus_PerformBinarySearchTest()(ok bool)//col:100
static_ZyanStatus_AllocatorReallocate()(ok bool)//col:116
static_ZyanStatus_AllocatorDeallocate()(ok bool)//col:168
}

vector struct{}
)

func NewVector()Vector{ return & vector{} }

func (v *vector)static_void_InitTestdata()(ok bool){//col:5





return true
}


func (v *vector)static_ZyanStatus_PerformBasicTests()(ok bool){//col:40



































return true
}


func (v *vector)static_ZyanStatus_PerformBinarySearchTest()(ok bool){//col:100




























































return true
}


func (v *vector)static_ZyanStatus_AllocatorReallocate()(ok bool){//col:116
















return true
}


func (v *vector)static_ZyanStatus_AllocatorDeallocate()(ok bool){//col:168




















































return true
}




