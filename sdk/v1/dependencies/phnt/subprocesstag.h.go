package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\subprocesstag.h.back

const (
	eTagInfoLevelNameFromTag            = 1 //col:3
	eTagInfoLevelNamesReferencingModule = 2 //col:4
	eTagInfoLevelNameTagMapping         = 3 //col:5
	eTagInfoLevelMax                    = 4 //col:6
)

const (
	eTagTypeService = 1 //col:10
	eTagTypeMax     = 2 //col:11
)

type _TAG_INFO_NAME_FROM_TAG_IN_PARAMS struct {
	dwPid uint32 //col:7
	dwTag uint32 //col:8
}

type _TAG_INFO_NAME_FROM_TAG_OUT_PARAMS struct {
	eTagType uint32  //col:12
	pszName  *uint32 //col:13
}

type _TAG_INFO_NAME_FROM_TAG struct {
	InParams  TAG_INFO_NAME_FROM_TAG_IN_PARAMS  //col:17
	OutParams TAG_INFO_NAME_FROM_TAG_OUT_PARAMS //col:18
}

type _TAG_INFO_NAMES_REFERENCING_MODULE_IN_PARAMS struct {
	dwPid     uint32  //col:22
	pszModule *uint32 //col:23
}

type _TAG_INFO_NAMES_REFERENCING_MODULE_OUT_PARAMS struct {
	eTagType  uint32  //col:27
	pmszNames *uint32 //col:28
}

type _TAG_INFO_NAMES_REFERENCING_MODULE struct {
	InParams  TAG_INFO_NAMES_REFERENCING_MODULE_IN_PARAMS  //col:32
	OutParams TAG_INFO_NAMES_REFERENCING_MODULE_OUT_PARAMS //col:33
}

type _TAG_INFO_NAME_TAG_MAPPING_IN_PARAMS struct {
	dwPid uint32 //col:36
}

type _TAG_INFO_NAME_TAG_MAPPING_ELEMENT struct {
	eTagType     uint32  //col:43
	dwTag        uint32  //col:44
	pszName      *uint32 //col:45
	pszGroupName *uint32 //col:46
}

type _TAG_INFO_NAME_TAG_MAPPING_OUT_PARAMS struct {
	cElements               uint32                             //col:48
	pNameTagMappingElements PTAG_INFO_NAME_TAG_MAPPING_ELEMENT //col:49
}

type _TAG_INFO_NAME_TAG_MAPPING struct {
	InParams   TAG_INFO_NAME_TAG_MAPPING_IN_PARAMS   //col:53
	pOutParams PTAG_INFO_NAME_TAG_MAPPING_OUT_PARAMS //col:54
}

