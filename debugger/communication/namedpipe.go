package communication

import (
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

const (
	PipeAccessDuplex       = 0x00000003
	PipeTypeMessage        = 0x00000004
	PipeReadmodeMessage    = 0x00000002
	PipeWait               = 0x00000000
	PipeUnlimitedInstances = 255
)

func NamedPipeServerCreatePipe(pipeName string, outputBufferSize, inputBufferSize uint32) (windows.Handle, error) {
	namePtr, err := windows.UTF16PtrFromString(pipeName)
	if err != nil {
		return windows.InvalidHandle, err
	}

	handle, err := windows.CreateNamedPipe(
		namePtr,
		PipeAccessDuplex,
		PipeTypeMessage|PipeReadmodeMessage|PipeWait,
		PipeUnlimitedInstances,
		outputBufferSize,
		inputBufferSize,
		0,
		nil,
	)
	if err != nil {
		mylog.Warning("error occurred while creating the pipe", "error", err)
		return windows.InvalidHandle, err
	}
	return handle, nil
}

func NamedPipeServerWaitForClientConnection(pipeHandle windows.Handle) error {
	err := windows.ConnectNamedPipe(pipeHandle, nil)
	if err != nil {
		windows.CloseHandle(pipeHandle)
		return err
	}
	return nil
}

func NamedPipeServerReadClientMessage(pipeHandle windows.Handle, buffer []byte) (uint32, error) {
	var bytesRead uint32
	err := windows.ReadFile(pipeHandle, buffer, &bytesRead, nil)
	return bytesRead, err
}

func NamedPipeServerSendMessageToClient(pipeHandle windows.Handle, message []byte) (uint32, error) {
	var bytesWritten uint32
	err := windows.WriteFile(pipeHandle, message, &bytesWritten, nil)
	return bytesWritten, err
}

func NamedPipeClientConnectToServer(pipeName string) (windows.Handle, error) {
	namePtr, err := windows.UTF16PtrFromString(pipeName)
	if err != nil {
		return windows.InvalidHandle, err
	}

	handle, err := windows.CreateFile(
		namePtr,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		0,
		nil,
		windows.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		return windows.InvalidHandle, err
	}
	return handle, nil
}

func NamedPipeClientSendMessageToServer(clientHandle windows.Handle, message []byte) (uint32, error) {
	var bytesWritten uint32
	err := windows.WriteFile(clientHandle, message, &bytesWritten, nil)
	return bytesWritten, err
}

func NamedPipeClientReadServerMessage(clientHandle windows.Handle, buffer []byte) (uint32, error) {
	var bytesRead uint32
	err := windows.ReadFile(clientHandle, buffer, &bytesRead, nil)
	return bytesRead, err
}

func NamedPipeCloseHandle(handle windows.Handle) {
	if handle != windows.InvalidHandle && handle != 0 {
		windows.CloseHandle(handle)
	}
}
