package script-engine-wrapper
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\script-engine-wrapper\script-engine.cpp.back

type (
ScriptEngine interface{
ScriptEngineEvalSingleExpression()(ok bool)//col:43
}
)

func NewScriptEngine() { return & scriptEngine{} }

func (s *scriptEngine)ScriptEngineEvalSingleExpression()(ok bool){//col:43
/*ScriptEngineEvalSingleExpression(string Expr, PBOOLEAN HasError)
{
    PVOID  CodeBuffer;
    UINT64 BufferAddress;
    UINT32 BufferLength;
    UINT32 Pointer;
    UINT64 Result = NULL;
    Expr.insert(0, "formats(");
    Expr.append(");");
    CodeBuffer = ScriptEngineParseWrapper((char *)Expr.c_str(), FALSE);
    if (CodeBuffer == NULL)
    {
        *HasError = TRUE;
        return NULL;
    }
    BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
    BufferLength  = ScriptEngineWrapperGetSize(CodeBuffer);
    Pointer       = ScriptEngineWrapperGetPointer(CodeBuffer);
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        KdSendScriptPacketToDebuggee(BufferAddress, BufferLength, Pointer, TRUE);
        if (g_ErrorStateOfResultOfEvaluatedExpression == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
        {
            Result                                    = g_ResultOfEvaluatedExpression;
            g_ErrorStateOfResultOfEvaluatedExpression = NULL;
            g_ResultOfEvaluatedExpression             = NULL;
            *HasError                                 = FALSE;
        }
        else
        {
            g_ErrorStateOfResultOfEvaluatedExpression = NULL;
            g_ResultOfEvaluatedExpression             = NULL;
            *HasError = TRUE;
            Result    = NULL;
        }
    }
    else
    {
        Result = ScriptEngineEvalUInt64StyleExpressionWrapper(Expr, HasError);
    }
    ScriptEngineWrapperRemoveSymbolBuffer(CodeBuffer);
    return Result;
}*/
return true
}



