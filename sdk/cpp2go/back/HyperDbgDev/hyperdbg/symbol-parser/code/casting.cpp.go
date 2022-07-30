package code
//back\HyperDbgDev\hyperdbg\symbol-parser\code\casting.cpp.back

type (
Casting interface{
      EXPRSSION-> cast < STRING0 > ()(ok bool)//col:29
 * FiledOfStructName for future ()(ok bool)//col:234
 * it returns FALSE ()(ok bool)//col:283
main()(ok bool)//col:415
}
)

func NewCasting() { return & casting{} }

func (c *casting)      EXPRSSION-> cast < STRING0 > ()(ok bool){//col:29
/*      EXPRSSION-> cast < STRING0 > ( EXPRESSION ) STRING_SEQUENCE
      STRING_SEQUENCE->.STRING[1 - N] STRING_SEQUENCE
      STRING_SEQUENCE->'->' STRING STRING_SEQUENCE
      STRING_SEQUENCE->eps
typedef struct _UNICODE_STRING
{
} UNICODE_STRING, *PUNICODE_STRING;*/
return true
}

func (c *casting) * FiledOfStructName for future ()(ok bool){//col:234
/* * FiledOfStructName for future (next '->' or '.' )
 * structure
 *
 * was any error it returns FALSE (in that case script engine should throw an
 * error to the user) otherwise return TRUE which means the script engine can
 * safely use the details
BOOLEAN
SymCastingQueryForFiledsAndTypes(_In_ const char * StructName,
                                 _In_ const char * FiledOfStructName,
                                 _Out_ PBOOLEAN    IsStructNamePointerOrNot,
                                 _Out_ PBOOLEAN    IsFiledOfStructNamePointerOrNot,
                                 _Out_ char **     NewStructOrTypeName,
                                 _Out_ UINT32 * OffsetOfFieldFromTop,
                                 _Out_ UINT32 * SizeOfField)
{
    BOOLEAN IsPointer                 = FALSE;
    BOOLEAN IsTheStructItselfAPointer = FALSE;
    UINT32  TempOffsetOfFieldFromTop  = 0;
    UINT32  TempSizeOfField           = 0;
    if (strcmp(StructName, "STUPID_STRUCT1") == 0 ||
        strcmp(StructName, "PSTUPID_STRUCT1") == 0)
    {
        if (strcmp(StructName, "PSTUPID_STRUCT1") == 0)
        {
            IsTheStructItselfAPointer = TRUE;
        }
        if (strcmp(FiledOfStructName, "Flag32") == 0)
        {
            IsPointer                = FALSE;
            TempOffsetOfFieldFromTop = 0x0;
            TempSizeOfField          = 0x4;
            strcpy(*NewStructOrTypeName, "UINT32");
        }
        else if (strcmp(FiledOfStructName, "Flag64") == 0)
        {
            IsPointer                = FALSE;
            TempOffsetOfFieldFromTop = 0x4;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "UINT64");
        }
        else if (strcmp(FiledOfStructName, "Context") == 0)
        {
            IsPointer                = TRUE;
            TempOffsetOfFieldFromTop = 0xc;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "PVOID");
        }
        else if (strcmp(FiledOfStructName, "StringValue") == 0)
        {
            IsPointer                = TRUE;
            TempOffsetOfFieldFromTop = 0x14;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "PUNICODE_STRING");
        }
        else
        {
            return FALSE;
        }
    }
    else if (strcmp(StructName, "STUPID_STRUCT2") == 0 ||
             strcmp(StructName, "PSTUPID_STRUCT2") == 0)
    {
        if (strcmp(StructName, "PSTUPID_STRUCT2") == 0)
        {
            IsTheStructItselfAPointer = TRUE;
        }
        if (strcmp(FiledOfStructName, "Sina32") == 0)
        {
            IsPointer                = FALSE;
            TempOffsetOfFieldFromTop = 0x0;
            TempSizeOfField          = 0x4;
            strcpy(*NewStructOrTypeName, "UINT32");
        }
        else if (strcmp(FiledOfStructName, "Sina64") == 0)
        {
            IsPointer                = FALSE;
            TempOffsetOfFieldFromTop = 0x4;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "UINT64");
        }
        else if (strcmp(FiledOfStructName, "AghaaSina") == 0)
        {
            IsPointer                = TRUE;
            TempOffsetOfFieldFromTop = 0xc;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "PVOID");
        }
        else if (strcmp(FiledOfStructName, "UnicodeStr") == 0)
        {
            IsPointer                = TRUE;
            TempOffsetOfFieldFromTop = 0x14;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "PUNICODE_STRING");
        }
        else if (strcmp(FiledOfStructName, "StupidStruct1") == 0)
        {
            IsPointer                = TRUE;
            TempOffsetOfFieldFromTop = 0x1c;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "PSTUPID_STRUCT1");
        }
        else
        {
            return FALSE;
        }
    }
    else if (strcmp(StructName, "UNICODE_STRING") == 0 ||
             strcmp(StructName, "PUNICODE_STRING") == 0)
    {
        if (strcmp(StructName, "PUNICODE_STRING") == 0)
        {
            IsTheStructItselfAPointer = TRUE;
        }
        if (strcmp(FiledOfStructName, "Length") == 0)
        {
            IsPointer                = FALSE;
            TempOffsetOfFieldFromTop = 0x0;
            TempSizeOfField          = 0x2;
            strcpy(*NewStructOrTypeName, "USHORT");
        }
        else if (strcmp(FiledOfStructName, "MaximumLength") == 0)
        {
            IsPointer                = FALSE;
            TempOffsetOfFieldFromTop = 0x2;
            TempSizeOfField          = 0x2;
            strcpy(*NewStructOrTypeName, "USHORT");
        }
        else if (strcmp(FiledOfStructName, "Buffer") == 0)
        {
            IsPointer                = TRUE;
            TempOffsetOfFieldFromTop = 0x4;
            TempSizeOfField          = 0x8;
            strcpy(*NewStructOrTypeName, "PWSTR");
        }
        else
        {
            return FALSE;
        }
    }
    else
    {
        return FALSE;
    }
    *OffsetOfFieldFromTop            = TempOffsetOfFieldFromTop;
    *SizeOfField                     = TempSizeOfField;
    *IsFiledOfStructNamePointerOrNot = IsPointer;
    *IsStructNamePointerOrNot        = IsTheStructItselfAPointer;
    return TRUE;
}*/
return true
}

func (c *casting) * it returns FALSE ()(ok bool){//col:283
/* * it returns FALSE (in that case script engine should throw an error to the
 * user) otherwise return TRUE which means the script engine can safely use the
 * details
BOOLEAN
SymQuerySizeof(_In_ const char * StructNameOrTypeName, _Out_ UINT32 * SizeOfField)
{
    if (strcmp(StructNameOrTypeName, "STUPID_STRUCT1") == 0)
    {
        *SizeOfField = sizeof(STUPID_STRUCT1);
    }
    else if (strcmp(StructNameOrTypeName, "STUPID_STRUCT2") == 0)
    {
        *SizeOfField = sizeof(STUPID_STRUCT2);
    }
    else if (strcmp(StructNameOrTypeName, "UNICODE_STRING") == 0)
    {
        *SizeOfField = sizeof(UNICODE_STRING);
    }
    else if (strcmp(StructNameOrTypeName, "PSTUPID_STRUCT1") == 0)
    {
        *SizeOfField = sizeof(PSTUPID_STRUCT1);
    }
    else if (strcmp(StructNameOrTypeName, "PSTUPID_STRUCT2") == 0)
    {
        *SizeOfField = sizeof(PSTUPID_STRUCT2);
    }
    else if (strcmp(StructNameOrTypeName, "PUNICODE_STRING") == 0)
    {
        *SizeOfField = sizeof(PUNICODE_STRING);
    }
    else
    {
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (c *casting)main()(ok bool){//col:415
/*main()
{
    UINT32 SizeofResult = 0;
    if (SymQuerySizeof("PUNICODE_STRING", &SizeofResult))
    {
        printf("result of sizeof(PUNICODE_STRING) is : 0x%x\n", SizeofResult);
    }
    if (SymQuerySizeof("SOME_UNKNOWN_STRUCT", &SizeofResult))
    {
        printf("result of sizeof(SOME_UNKNOWN_STRUCT) is : 0x%x\n", SizeofResult);
    }
    else
    {
        printf("SOME_UNKNOWN_STRUCT is not found\n");
    }
    printf("\n\n\n");
        0x00000260`d065ee30 
           +0x000 Sina32           : 0x32
           +0x008 Sina64           : 0x64
           +0x010 AghaaSina        : 0x00000000`00000055 Void
           +0x018 UnicodeStr       : 0x00000260`d065ec70 AllocateStructForCasting::__l2::_UNICODE_STRING
              +0x000 Length           : 0x40
              +0x002 MaximumLength    : 0x40
              +0x008 Buffer           : 0x00000260`d065ecf0  "Goodbye I'm at stupid struct 2!"
           +0x020 StupidStruct1    : 0x00000260`d065eda0 AllocateStructForCasting::__l2::_STUPID_STRUCT1
              +0x000 Flag32           : 0x3232
              +0x008 Flag64           : 0x6464
              +0x010 Context          : 0x00000000`00000085 Void
              +0x018 StringValue      : 0x00000260`d065eb50 AllocateStructForCasting::__l2::_UNICODE_STRING
                 +0x000 Length           : 0x3c
                 +0x002 MaximumLength    : 0x3c
                 +0x008 Buffer           : 0x00000260`d065ebd0  "Hi come from stupid struct 1!"
        my_var =
    BOOLEAN IsStructPointerOrNot = FALSE;
    BOOLEAN IsFieldPointerOrNot  = FALSE;
    UINT32  OffsetOfFieldFromTop = NULL;
    UINT32  SizeOfField          = NULL;
    CHAR *  NewStructOrTypeName  = (CHAR *)malloc(100);
    if (SymCastingQueryForFiledsAndTypes(
            "STUPID_STRUCT2",
            "UnicodeStr",
            &IsStructPointerOrNot,
            &IsFieldPointerOrNot,
            &NewStructOrTypeName,
            &OffsetOfFieldFromTop,
            &SizeOfField))
    {
        printf("is the structure itself pointer or not: %s\n",
               IsStructPointerOrNot ? "yes" : "no");
        printf("is the field of structure itself pointer or not: %s\n",
               IsFieldPointerOrNot ? "yes" : "no");
        printf("offset of field from top : %x\n", OffsetOfFieldFromTop);
        printf("size of field : %x\n", SizeOfField);
    }
    free(NewStructOrTypeName);
}*/
return true
}



