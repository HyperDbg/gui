#pragma once
#define CONTAINING_RECORD(address, type, field)                                \
  ((type *)((char *)(address) - (unsigned long)(&((type *)0)->field)))
void InitializeListHead(PLIST_ENTRY ListHead);
void InsertHeadList(PLIST_ENTRY ListHead, PLIST_ENTRY Entry);
BOOLEAN
RemoveEntryList(PLIST_ENTRY Entry);
