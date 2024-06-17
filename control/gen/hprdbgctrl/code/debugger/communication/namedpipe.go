package control

func NamedPipeServerCreatePipe() syscall.Handle { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:4
}

func NamedPipeServerWaitForClientConntection() bool { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:23
}

func NamedPipeServerReadClientMessage() uint32 { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:33
}

func NamedPipeServerCloseHandle() { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:66
}

func NamedPipeClientCreatePipe() syscall.Handle { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:69
}

func NamedPipeClientCreatePipeOverlappedIo() syscall.Handle { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:88
}

func NamedPipeClientSendMessage() bool { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:111
}

func NamedPipeClientReadMessage() uint32 { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:129
}

func NamedPipeClientClosePipe() { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:144
}

func NamedPipeServerExample() int { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:147
}

func NamedPipeClientExample() int { // hprdbgctrl\code\debugger\communication\namedpipe.cpp:179
}
