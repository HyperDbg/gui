package phnt

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

type TAG_INFO_NAME_FROM_TAG_IN_PARAMS struct {
	dwPid uint32 //col:3
	dwTag uint32 //col:4
}

type TAG_INFO_NAME_FROM_TAG_OUT_PARAMS struct {
	eTagType uint32 //col:8
	pszName  PWSTR  //col:9
}

type TAG_INFO_NAME_FROM_TAG struct {
	InParams  TAG_INFO_NAME_FROM_TAG_IN_PARAMS  //col:13
	OutParams TAG_INFO_NAME_FROM_TAG_OUT_PARAMS //col:14
}

type TAG_INFO_NAMES_REFERENCING_MODULE_IN_PARAMS struct {
	dwPid     uint32 //col:18
	pszModule PWSTR  //col:19
}

type TAG_INFO_NAMES_REFERENCING_MODULE_OUT_PARAMS struct {
	eTagType  uint32 //col:23
	pmszNames PWSTR  //col:24
}

type TAG_INFO_NAMES_REFERENCING_MODULE struct {
	InParams  TAG_INFO_NAMES_REFERENCING_MODULE_IN_PARAMS  //col:28
	OutParams TAG_INFO_NAMES_REFERENCING_MODULE_OUT_PARAMS //col:29
}

type TAG_INFO_NAME_TAG_MAPPING_IN_PARAMS struct {
	dwPid uint32 //col:33
}

type TAG_INFO_NAME_TAG_MAPPING_ELEMENT struct {
	eTagType     uint32 //col:37
	dwTag        uint32 //col:38
	pszName      PWSTR  //col:39
	pszGroupName PWSTR  //col:40
}

type TAG_INFO_NAME_TAG_MAPPING_OUT_PARAMS struct {
	cElements               uint32                             //col:44
	pNameTagMappingElements PTAG_INFO_NAME_TAG_MAPPING_ELEMENT //col:45
}

type TAG_INFO_NAME_TAG_MAPPING struct {
	InParams   TAG_INFO_NAME_TAG_MAPPING_IN_PARAMS   //col:49
	pOutParams PTAG_INFO_NAME_TAG_MAPPING_OUT_PARAMS //col:50
}
