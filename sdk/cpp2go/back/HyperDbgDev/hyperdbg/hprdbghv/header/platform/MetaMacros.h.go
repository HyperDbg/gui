package platform

//back\HyperDbgDev\hyperdbg\hprdbghv\header\platform\MetaMacros.h.back

const (
	MetaScopedExpr(
Pre_, Post_, ScopedExpr_
) = do { Pre_; ScopedExpr_; Post_; } while (FALSE) //col:15
LIST_FOR_EACH(_head, = _struct_type, _var)               _LIST_FOR_EACH(_head, _struct_type, Link, _var) //col:33
LIST_FOR_NEXT(_start, = _head, _type, _var)              _LIST_FOR_NEXT(_start, _head, _type, Link, _var) //col:34
LIST_FOR_EACH_LINK(_head, = _struct_type, _member, _var) _LIST_FOR_EACH(_head, _struct_type, _member, _var) //col:35
_NEXT(_var, = _member)              _var->_member.Flink //col:40
_NEXT_ENTRY(_var, = _member, _type) CONTAINING_RECORD(_NEXT(_var, _member), _type, _member) //col:41
PREPROC_CONCAT(a, = b)     PREPROC_CONCAT_1(a, b) //col:43
PREPROC_CONCAT_1(a, = b)   PREPROC_CONCAT_2(~, a##b) //col:44
PREPROC_CONCAT_2(p, = res) res //col:45
UNIQUE_NAME(base) = PREPROC_CONCAT(base, __LINE__) //col:47
_LIST_FOR_EACH(_head, _type, _member, _var) = for (_type * _var = CONTAINING_RECORD(_head.Flink, _type, _member), *UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type); &_var->_member != &_head; _var = UNIQUE_NAME(_n), UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type)) //col:49
_LIST_FOR_NEXT(_start, _head, _type, _member, _var) = for (_type * _var = _NEXT_ENTRY(_start, _member, _type), *UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type); &_var->_member != &_head; _var = UNIQUE_NAME(_n), UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type)) //col:54
)
