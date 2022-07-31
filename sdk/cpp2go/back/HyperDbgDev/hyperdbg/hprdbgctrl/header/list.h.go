package header
const(
CONTAINING_RECORD(address, type, field) = ((type *)((char *)(address) - (unsigned long)(&((type *)0)->field))) //col:14
)
