package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\subprocesstag.h.back

const(
_SUBPROCESSTAG_H =  //col:1
)

const(
    eTagInfoLevelNameFromTag  =  1   //col:3
    eTagInfoLevelNamesReferencingModule  = 2  //col:4
    eTagInfoLevelNameTagMapping  = 3  //col:5
    eTagInfoLevelMax = 4  //col:6
)


const(
    eTagTypeService  =  1  //col:10
    eTagTypeMax = 2  //col:11
)



type TAG_INFO_NAME_FROM_TAG_IN_PARAMS struct{
dwPid ULONG
dwTag ULONG
}


type TAG_INFO_NAME_FROM_TAG_OUT_PARAMS struct{
eTagType ULONG
pszName PWSTR
}


type TAG_INFO_NAME_FROM_TAG struct{
InParams TAG_INFO_NAME_FROM_TAG_IN_PARAMS
OutParams TAG_INFO_NAME_FROM_TAG_OUT_PARAMS
}


type TAG_INFO_NAMES_REFERENCING_MODULE_IN_PARAMS struct{
dwPid ULONG
pszModule PWSTR
}


type TAG_INFO_NAMES_REFERENCING_MODULE_OUT_PARAMS struct{
eTagType ULONG
pmszNames PWSTR
}


type TAG_INFO_NAMES_REFERENCING_MODULE struct{
InParams TAG_INFO_NAMES_REFERENCING_MODULE_IN_PARAMS
OutParams TAG_INFO_NAMES_REFERENCING_MODULE_OUT_PARAMS
}


type TAG_INFO_NAME_TAG_MAPPING_IN_PARAMS struct{
dwPid ULONG
}


type TAG_INFO_NAME_TAG_MAPPING_ELEMENT struct{
eTagType ULONG
dwTag ULONG
pszName PWSTR
pszGroupName PWSTR
}


type TAG_INFO_NAME_TAG_MAPPING_OUT_PARAMS struct{
cElements ULONG
pNameTagMappingElements PTAG_INFO_NAME_TAG_MAPPING_ELEMENT
}


type TAG_INFO_NAME_TAG_MAPPING struct{
InParams TAG_INFO_NAME_TAG_MAPPING_IN_PARAMS
pOutParams PTAG_INFO_NAME_TAG_MAPPING_OUT_PARAMS
}




