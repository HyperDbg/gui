package Headers
//back\HyperDbgDev\hyperdbg\include\SDK\Headers\Symbols.h.back

type (
Symbols interface{
typedef VOID ()(ok bool)//col:76
static_assert()(ok bool)//col:97
}
)

func NewSymbols() { return & symbols{} }

func (s *symbols)typedef VOID ()(ok bool){//col:76
/*typedef VOID (*SymbolMapCallback)(UINT64 Address, char * ModuleName, char * ObjectName, unsigned int ObjectSize);
 * symbol table entry
 *
typedef struct _DEBUGGER_UPDATE_SYMBOL_TABLE
{
    UINT32               TotalSymbols;
    UINT32               CurrentSymbolIndex;
    MODULE_SYMBOL_DETAIL SymbolDetailPacket;
} DEBUGGER_UPDATE_SYMBOL_TABLE, *PDEBUGGER_UPDATE_SYMBOL_TABLE;*/
return true
}

func (s *symbols)static_assert()(ok bool){//col:97
/*static_assert(sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE) < PacketChunkSize,
              "err (static_assert), size of PacketChunkSize should be bigger than DEBUGGER_UPDATE_SYMBOL_TABLE (MODULE_SYMBOL_DETAIL)");
==============================================================================================
 *
typedef struct _DEBUGGEE_SYMBOL_UPDATE_RESULT
{
} DEBUGGEE_SYMBOL_UPDATE_RESULT, *PDEBUGGEE_SYMBOL_UPDATE_RESULT;*/
return true
}



