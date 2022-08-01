package platform

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\platform\MetaMacros.h.back

const (
	MetaScopedExpr(
Pre_, Post_, ScopedExpr_
) = do { Pre_; ScopedExpr_; Post_; } while (FALSE) //col:1
LIST_FOR_EACH(_head, = _struct_type, _var)               _LIST_FOR_EACH(_head, _struct_type, Link, _var) //col:8
LIST_FOR_NEXT(_start, = _head, _type, _var)              _LIST_FOR_NEXT(_start, _head, _type, Link, _var) //col:9
LIST_FOR_EACH_LINK(_head, = _struct_type, _member, _var) _LIST_FOR_EACH(_head, _struct_type, _member, _var) //col:10
_NEXT(_var, = _member)              _var->_member.Flink //col:11
_NEXT_ENTRY(_var, = _member, _type) CONTAINING_RECORD(_NEXT(_var, _member), _type, _member) //col:12
PREPROC_CONCAT(a, = b)     PREPROC_CONCAT_1(a, b) //col:13
PREPROC_CONCAT_1(a, = b)   PREPROC_CONCAT_2(~, a##b) //col:14
PREPROC_CONCAT_2(p, = res) res //col:15
UNIQUE_NAME(base) = PREPROC_CONCAT(base, __LINE__) //col:16
_LIST_FOR_EACH(_head, _type, _member, _var) = for (_type * _var = CONTAINING_RECORD(_head.Flink, _type, _member), *UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type); &_var->_member != &_head; _var = UNIQUE_NAME(_n), UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type)) //col:17
_LIST_FOR_NEXT(_start, _head, _type, _member, _var) = for (_type * _var = _NEXT_ENTRY(_start, _member, _type), *UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type); &_var->_member != &_head; _var = UNIQUE_NAME(_n), UNIQUE_NAME(_n) = _NEXT_ENTRY(_var, _member, _type)) //col:21
)
