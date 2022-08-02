package script_engine_wrapper


type ALLOCATED_MEMORY_FOR_SCRIPT_ENGINE_CASTING struct{
* int8 //col:3
* int8 //col:4
* int8 //col:5
* int8 //col:6
* int8 //col:7
* int8 //col:8
}



type     typedef struct _UNICODE_STRING struct{
Length uint16 //col:12
MaximumLength uint16 //col:13
Buffer PWSTR //col:14
}



type     typedef struct _STUPID_STRUCT1 struct{
Flag32 uint32 //col:18
Flag64 uint64 //col:19
Context PVOID //col:20
StringValue PUNICODE_STRING //col:21
}



type     typedef struct _STUPID_STRUCT2 struct{
Sina32 uint32 //col:25
Sina64 uint64 //col:26
AghaaSina PVOID //col:27
UnicodeStr PUNICODE_STRING //col:28
StupidStruct1 PSTUPID_STRUCT1 //col:29
}



type     typedef struct _TEST_STRUCT struct{
Var1 uint64 //col:68
Var2 uint64 //col:69
Var3 uint64 //col:70
Var4 uint64 //col:71
}




type (
ScriptEngineWrapper interface{
ScriptEngineConvertNameToAddressWrapper()(ok bool)//col:125
AllocateStructForCasting()(ok bool)//col:140
ScriptEngineWrapperGetSize()(ok bool)//col:148
}

scriptEngineWrapper struct{}
)

func NewScriptEngineWrapper()ScriptEngineWrapper{ return & scriptEngineWrapper{} }

func (s *scriptEngineWrapper)ScriptEngineConvertNameToAddressWrapper()(ok bool){//col:125





























































































































return true
}


func (s *scriptEngineWrapper)AllocateStructForCasting()(ok bool){//col:140















return true
}


func (s *scriptEngineWrapper)ScriptEngineWrapperGetSize()(ok bool){//col:148








return true
}




