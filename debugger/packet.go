package debugger

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

type Packet struct {
	driver *driver.Handle
}

func NewPacket(dh *driver.Handle) *Packet {
	return &Packet{
		driver: dh,
	}
}

func (p *Packet) IsConnected() bool {
	return p.driver != nil && p.driver.IsConnected()
}

func (p *Packet) InitializeVmx() error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 8)
	err := p.driver.SendBuffer(buffer, IoctlInitializeVmx)
	if err != nil {
		return fmt.Errorf("failed to initialize VMX: %v", err)
	}

	return nil
}

func (p *Packet) TerminateVmx() error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	err := p.driver.SendBuffer(nil, IoctlTerminateVmx)
	if err != nil {
		return fmt.Errorf("failed to terminate VMX: %v", err)
	}

	return nil
}

func (p *Packet) ReturnIrpPendingAndDisallowIoctl() error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	err := p.driver.SendBuffer(nil, IoctlReturnIrpPendingAndDisallowIoctl)
	if err != nil {
		return fmt.Errorf("failed to return IRP pending: %v", err)
	}

	return nil
}

func (p *Packet) ReadMemory(pid uint32, address uint64, size uint32) ([]byte, error) {
	if !p.IsConnected() {
		return nil, fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 32)
	binary.LittleEndian.PutUint32(buffer[0:4], pid)
	binary.LittleEndian.PutUint64(buffer[4:12], address)
	binary.LittleEndian.PutUint32(buffer[12:16], size)
	binary.LittleEndian.PutUint32(buffer[16:20], uint32(DEBUGGER_READ_MEMORY_TYPE_READ_FROM_MEMORY))

	err := p.driver.SendBuffer(buffer, IoctlDebuggerReadMemory)
	if err != nil {
		return nil, fmt.Errorf("failed to send read memory request: %v", err)
	}

	response, err := p.driver.ReceiveBuffer()
	if err != nil {
		return nil, fmt.Errorf("failed to receive read memory response: %v", err)
	}

	returnLength := binary.LittleEndian.Uint32(response[0:4])
	if returnLength == 0 {
		return nil, fmt.Errorf("read memory failed: kernel status indicates failure")
	}

	return response[16 : 16+size], nil
}

func (p *Packet) WriteMemory(pid uint32, address uint64, data []byte) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16+len(data))
	binary.LittleEndian.PutUint32(buffer[0:4], pid)
	binary.LittleEndian.PutUint64(buffer[4:12], address)
	binary.LittleEndian.PutUint32(buffer[12:16], uint32(len(data)))
	copy(buffer[16:], data)

	err := p.driver.SendBuffer(buffer, IoctlDebuggerEditMemory)
	if err != nil {
		return fmt.Errorf("failed to send write memory request: %v", err)
	}

	_, err = p.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive write memory response: %v", err)
	}

	return nil
}

func (p *Packet) ReadPageTableEntries(req *DebuggerReadPageTableEntriesDetails) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16)
	binary.LittleEndian.PutUint64(buffer[0:8], req.VirtualAddress)
	binary.LittleEndian.PutUint32(buffer[8:12], req.ProcessId)

	err := p.driver.SendBuffer(buffer, IoctlDebuggerReadPageTable)
	if err != nil {
		return fmt.Errorf("failed to send read page table request: %v", err)
	}

	response, err := p.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive read page table response: %v", err)
	}

	req.Pml4eVirtualAddress = binary.LittleEndian.Uint64(response[0:8])
	req.Pml4eValue = binary.LittleEndian.Uint64(response[8:16])
	req.PdpteVirtualAddress = binary.LittleEndian.Uint64(response[16:24])
	req.PdpteValue = binary.LittleEndian.Uint64(response[24:32])
	req.PdeVirtualAddress = binary.LittleEndian.Uint64(response[32:40])
	req.PdeValue = binary.LittleEndian.Uint64(response[40:48])
	req.PteVirtualAddress = binary.LittleEndian.Uint64(response[48:56])
	req.PteValue = binary.LittleEndian.Uint64(response[56:64])
	req.KernelStatus = binary.LittleEndian.Uint32(response[64:68])

	return nil
}

func (p *Packet) Va2paAndPa2va(req *DebuggerVa2paAndPa2vaCommands) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 24)
	binary.LittleEndian.PutUint64(buffer[0:8], req.VirtualAddress)
	binary.LittleEndian.PutUint64(buffer[8:16], req.PhysicalAddress)
	binary.LittleEndian.PutUint32(buffer[16:20], req.ProcessId)
	binary.LittleEndian.PutUint32(buffer[20:24], func() uint32 {
		if req.IsVirtual2Physical {
			return 1
		}
		return 0
	}())

	err := p.driver.SendBuffer(buffer, IoctlDebuggerVa2paAndPa2va)
	if err != nil {
		return fmt.Errorf("failed to send va2pa/pa2va request: %v", err)
	}

	response, err := p.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive va2pa/pa2va response: %v", err)
	}

	req.VirtualAddress = binary.LittleEndian.Uint64(response[0:8])
	req.PhysicalAddress = binary.LittleEndian.Uint64(response[8:16])
	req.KernelStatus = binary.LittleEndian.Uint32(response[16:20])

	return nil
}

func (p *Packet) ReadMsr(msr uint32) (uint64, error) {
	if !p.IsConnected() {
		return 0, fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer[0:4], msr)

	err := p.driver.SendBuffer(buffer, IoctlDebuggerReadOrWriteMsr)
	if err != nil {
		return 0, fmt.Errorf("failed to send read msr request: %v", err)
	}

	response, err := p.driver.ReceiveBuffer()
	if err != nil {
		return 0, fmt.Errorf("failed to receive read msr response: %v", err)
	}

	value := binary.LittleEndian.Uint64(response[8:16])
	return value, nil
}

func (p *Packet) WriteMsr(msr uint32, value uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 12)
	binary.LittleEndian.PutUint32(buffer[0:4], msr)
	binary.LittleEndian.PutUint64(buffer[4:12], value)

	err := p.driver.SendBuffer(buffer, IoctlDebuggerReadOrWriteMsr)
	if err != nil {
		return fmt.Errorf("failed to send write msr request: %v", err)
	}

	_, err = p.driver.ReceiveBuffer()
	if err != nil {
		return fmt.Errorf("failed to receive write msr response: %v", err)
	}

	return nil
}

func (p *Packet) SetBreakpoint(pid uint32, address uint64, coreNumber uint32, bpType uint32, tag uint32) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 32)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = address
	*(*uint32)(unsafe.Pointer(&buffer[8])) = pid
	*(*uint32)(unsafe.Pointer(&buffer[12])) = coreNumber
	*(*uint32)(unsafe.Pointer(&buffer[16])) = bpType
	*(*uint32)(unsafe.Pointer(&buffer[20])) = tag

	err := p.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send set breakpoint request: %v", err)
	}

	return nil
}

func (p *Packet) RemoveBreakpoint(tag uint32) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = uint64(tag)

	err := p.driver.SendBuffer(buffer, IoctlUnregisterEvent)
	if err != nil {
		return fmt.Errorf("failed to send remove breakpoint request: %v", err)
	}

	return nil
}

func (p *Packet) EnableBreakpoint(tag uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16)
	binary.LittleEndian.PutUint64(buffer[0:8], tag)
	binary.LittleEndian.PutUint32(buffer[8:12], 1)

	err := p.driver.SendBuffer(buffer, IoctlDebuggerModifyEvents)
	if err != nil {
		return fmt.Errorf("failed to send enable breakpoint request: %v", err)
	}

	return nil
}

func (p *Packet) DisableBreakpoint(tag uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16)
	binary.LittleEndian.PutUint64(buffer[0:8], tag)
	binary.LittleEndian.PutUint32(buffer[8:12], 0)

	err := p.driver.SendBuffer(buffer, IoctlDebuggerModifyEvents)
	if err != nil {
		return fmt.Errorf("failed to send disable breakpoint request: %v", err)
	}

	return nil
}

func (p *Packet) ClearAllBreakpoints() error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 4)

	err := p.driver.SendBuffer(buffer, IoctlClearEvents)
	if err != nil {
		return fmt.Errorf("failed to send clear breakpoints request: %v", err)
	}

	return nil
}

func (p *Packet) RegisterEvent(eventTag uint64, eventType uint32, optionalParam1 uint64, optionalParam2 uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 32)
	binary.LittleEndian.PutUint64(buffer[0:8], eventTag)
	binary.LittleEndian.PutUint32(buffer[8:12], eventType)
	binary.LittleEndian.PutUint64(buffer[12:20], optionalParam1)
	binary.LittleEndian.PutUint64(buffer[20:28], optionalParam2)

	err := p.driver.SendBuffer(buffer, IoctlRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to register event: %v", err)
	}

	return nil
}

func (p *Packet) UnregisterEvent(eventTag uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(buffer[0:8], eventTag)

	err := p.driver.SendBuffer(buffer, IoctlUnregisterEvent)
	if err != nil {
		return fmt.Errorf("failed to unregister event: %v", err)
	}

	return nil
}

func (p *Packet) ModifyEvent(eventTag uint64, newEventType uint32) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16)
	binary.LittleEndian.PutUint64(buffer[0:8], eventTag)
	binary.LittleEndian.PutUint32(buffer[8:12], newEventType)

	err := p.driver.SendBuffer(buffer, IoctlModifyEvent)
	if err != nil {
		return fmt.Errorf("failed to modify event: %v", err)
	}

	return nil
}

func (p *Packet) AddActionToEvent(eventTag uint64, actionType uint32, actionParam uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 24)
	binary.LittleEndian.PutUint64(buffer[0:8], eventTag)
	binary.LittleEndian.PutUint32(buffer[8:12], actionType)
	binary.LittleEndian.PutUint64(buffer[12:20], actionParam)

	err := p.driver.SendBuffer(buffer, IoctlAddActionToEvent)
	if err != nil {
		return fmt.Errorf("failed to add action to event: %v", err)
	}

	return nil
}

func (p *Packet) ClearEvents() error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 4)

	err := p.driver.SendBuffer(buffer, IoctlClearEvents)
	if err != nil {
		return fmt.Errorf("failed to clear events: %v", err)
	}

	return nil
}

func (p *Packet) ReceiveEvent(bufferSize uint32) ([]byte, uint32, error) {
	if !p.IsConnected() {
		return nil, 0, fmt.Errorf("driver not available")
	}

	buffer, err := p.driver.ReceiveBuffer()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to receive event: %v", err)
	}

	bytesRead := uint32(len(buffer))
	return buffer, bytesRead, nil
}

func (p *Packet) SetEptHook(address uint64) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 16)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = 1
	*(*uint64)(unsafe.Pointer(&buffer[8])) = address

	err := p.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to set EPT hook: %v", err)
	}

	return nil
}

func (p *Packet) MonitorMemory(address uint64, size uint32, read bool, write bool) error {
	if !p.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 24)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = 2
	*(*uint64)(unsafe.Pointer(&buffer[8])) = address
	*(*uint32)(unsafe.Pointer(&buffer[16])) = size

	var flags uint32
	if read {
		flags |= 1
	}
	if write {
		flags |= 2
	}
	*(*uint32)(unsafe.Pointer(&buffer[20])) = flags

	err := p.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
	if err != nil {
		return fmt.Errorf("failed to monitor memory: %v", err)
	}

	return nil
}

func (p *Packet) QueryPte(virtualAddress uint64, processId uint32) ([]byte, error) {
	if !p.IsConnected() {
		return nil, fmt.Errorf("driver not available")
	}

	buffer := make([]byte, 24)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = 3
	*(*uint64)(unsafe.Pointer(&buffer[8])) = virtualAddress
	*(*uint32)(unsafe.Pointer(&buffer[16])) = processId

	err := p.driver.SendBuffer(buffer, IoctlDebuggerReadPageTable)
	if err != nil {
		return nil, fmt.Errorf("failed to query PTE: %v", err)
	}

	response, err := p.driver.ReceiveBuffer()
	if err != nil {
		return nil, fmt.Errorf("failed to receive PTE response: %v", err)
	}

	return response, nil
}

func (p *Packet) Va2Pa(virtualAddress uint64, processId uint32) (uint64, error) {
	if !p.IsConnected() {
		return 0, fmt.Errorf("driver not available")
	}

	req := &DebuggerVa2paAndPa2vaCommands{
		VirtualAddress:     virtualAddress,
		ProcessId:          processId,
		IsVirtual2Physical: true,
	}

	err := p.Va2paAndPa2va(req)
	if err != nil {
		return 0, err
	}

	return req.PhysicalAddress, nil
}

func (p *Packet) Pa2Va(physicalAddress uint64, processId uint32) (uint64, error) {
	if !p.IsConnected() {
		return 0, fmt.Errorf("driver not available")
	}

	req := &DebuggerVa2paAndPa2vaCommands{
		PhysicalAddress:    physicalAddress,
		ProcessId:          processId,
		IsVirtual2Physical: false,
	}

	err := p.Va2paAndPa2va(req)
	if err != nil {
		return 0, err
	}

	return req.VirtualAddress, nil
}
