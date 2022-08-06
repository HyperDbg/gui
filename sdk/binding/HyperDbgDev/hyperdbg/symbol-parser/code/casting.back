










#include "pch.h"











typedef struct _UNICODE_STRING
{
    USHORT Length;        
    USHORT MaximumLength; 
    PWSTR  Buffer;        
} UNICODE_STRING, *PUNICODE_STRING;

typedef struct _STUPID_STRUCT1
{
    UINT32          Flag32;      
    UINT64          Flag64;      
    PVOID           Context;     
    PUNICODE_STRING StringValue; 
} STUPID_STRUCT1, *PSTUPID_STRUCT1;

typedef struct _STUPID_STRUCT2
{
    UINT32          Sina32;        
    UINT64          Sina64;        
    PVOID           AghaaSina;     
    PUNICODE_STRING UnicodeStr;    
    PSTUPID_STRUCT1 StupidStruct1; 

} STUPID_STRUCT2, *PSTUPID_STRUCT2;























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
}












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
}

int
main()
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
}
