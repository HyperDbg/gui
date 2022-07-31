package phnt


const(
_NTOBAPI_H =  //col:13
OBJECT_TYPE_CREATE = 0x0001 //col:16
OBJECT_TYPE_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0x1) //col:17
DIRECTORY_QUERY = 0x0001 //col:21
DIRECTORY_TRAVERSE = 0x0002 //col:22
DIRECTORY_CREATE_OBJECT = 0x0004 //col:23
DIRECTORY_CREATE_SUBDIRECTORY = 0x0008 //col:24
DIRECTORY_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0xf) //col:25
SYMBOLIC_LINK_QUERY = 0x0001 //col:29
SYMBOLIC_LINK_SET = 0x0002 //col:30
SYMBOLIC_LINK_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0x1) //col:31
SYMBOLIC_LINK_ALL_ACCESS_EX = (STANDARD_RIGHTS_REQUIRED | 0xFFFF) //col:32
OBJ_PROTECT_CLOSE = 0x00000001 //col:36
OBJ_INHERIT = 0x00000002 //col:39
OBJ_AUDIT_OBJECT_CLOSE = 0x00000004 //col:42
ObjectBasicInformation = 0 //col:58
ObjectNameInformation = 1 //col:59
ObjectTypesInformation = 3 //col:60
ObjectHandleFlagInformation = 4 //col:61
ObjectSessionInformation = 5 //col:62
ObjectSessionObjectInformation = 6 //col:63
DUPLICATE_CLOSE_SOURCE = 0x00000001 //col:151
DUPLICATE_SAME_ACCESS = 0x00000002 //col:152
DUPLICATE_SAME_ATTRIBUTES = 0x00000004 //col:153
OBJECT_BOUNDARY_DESCRIPTOR_VERSION = 1 //col:343
)

const(
    ObjectBasicInformation // q: OBJECT_BASIC_INFORMATION = 1  //col:48
    ObjectNameInformation // q: OBJECT_NAME_INFORMATION = 2  //col:49
    ObjectTypeInformation // q: OBJECT_TYPE_INFORMATION = 3  //col:50
    ObjectTypesInformation // q: OBJECT_TYPES_INFORMATION = 4  //col:51
    ObjectHandleFlagInformation // qs: OBJECT_HANDLE_FLAG_INFORMATION = 5  //col:52
    ObjectSessionInformation // s: void // change object session // (requires SeTcbPrivilege) = 6  //col:53
    ObjectSessionObjectInformation // s: void // change object session // (requires SeTcbPrivilege) = 7  //col:54
    MaxObjectInfoClass = 8  //col:55
)


const(
    OBNS_Invalid = 1  //col:329
    OBNS_Name = 2  //col:330
    OBNS_SID = 3  //col:331
    OBNS_IL = 4  //col:332
)


const(
    SymbolicLinkGlobalInformation  =  1 // s: ULONG  //col:427
    SymbolicLinkAccessMask // s: ACCESS_MASK = 2  //col:428
    MaxnSymbolicLinkInfoClass = 3  //col:429
)



type (
Ntobapi interface{
#if ()(ok bool)//col:56
#if ()(ok bool)//col:85
#if ()(ok bool)//col:303
NtQueryDirectoryObject()(ok bool)//col:333
NtCreatePrivateNamespace()(ok bool)//col:430
}

)

func NewNtobapi() { return & ntobapi{} }

func (n *ntobapi)#if ()(ok bool){//col:56































return true
}


func (n *ntobapi)#if ()(ok bool){//col:85





return true
}


func (n *ntobapi)#if ()(ok bool){//col:303


























































































































































return true
}


func (n *ntobapi)NtQueryDirectoryObject()(ok bool){//col:333



















return true
}


func (n *ntobapi)NtCreatePrivateNamespace()(ok bool){//col:430





















































return true
}




