package Headers
//back\HyperDbgDev\hyperdbg\include\SDK\Headers\Datatypes.h.back

const(
SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED = sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED) //col:65
)

type (
Datatypes interface{
typedef int ()(ok bool)//col:44
    sizeof()(ok bool)//col:76
static_assert()(ok bool)//col:141
}
)

func NewDatatypes() { return & datatypes{} }

func (d *datatypes)typedef int ()(ok bool){//col:44
/*typedef int (*Callback)(const char * Text);
 *
typedef struct _DEBUGGEE_USER_INPUT_PACKET
{
    UINT32  CommandLen;
    BOOLEAN IgnoreFinishedSignal;
    UINT32  Result;
} DEBUGGEE_USER_INPUT_PACKET, *PDEBUGGEE_USER_INPUT_PACKET;*/
return true
}

func (d *datatypes)    sizeof()(ok bool){//col:76
/*    sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED)
 *
typedef struct _DEBUGGER_PAUSE_PACKET_RECEIVED
{
} DEBUGGER_PAUSE_PACKET_RECEIVED, *PDEBUGGER_PAUSE_PACKET_RECEIVED;*/
return true
}

func (d *datatypes)static_assert()(ok bool){//col:141
/*static_assert(sizeof(DEBUGGEE_UD_PAUSED_PACKET) < PacketChunkSize,
              "err (static_assert), size of PacketChunkSize should be bigger than DEBUGGEE_UD_PAUSED_PACKET");
 *
typedef struct _DEBUGGEE_MESSAGE_PACKET
{
    UINT32 OperationCode;
    CHAR   Message[PacketChunkSize];
} DEBUGGEE_MESSAGE_PACKET, *PDEBUGGEE_MESSAGE_PACKET;*/
return true
}



