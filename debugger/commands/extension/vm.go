package extension

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
)

type CommandEptHook struct {
	dbg    debugger.Packeter
	packet *debugger.Packet
}

func NewCommandEptHook(dbg debugger.Packeter, packet *debugger.Packet) *CommandEptHook {
	return &CommandEptHook{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandEptHook) Help() {
	mylog.Info("!epthook : set EPT hook.")
	mylog.Info("syntax : \t!epthook [address]")
	mylog.Info("  address - virtual address to hook (hex or decimal)")
}

func (c *CommandEptHook) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	address := mylog.Check2(c.parseAddress(tokens[1].Value))

	return c.SetEptHook(address)
}

func (c *CommandEptHook) parseAddress(addr string) (uint64, error) {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X") {
		return strconv.ParseUint(addr[2:], 16, 64)
	}
	return strconv.ParseUint(addr, 0, 64)
}

func (c *CommandEptHook) SetEptHook(address uint64) error {
	if c.packet == nil || !c.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	c.packet.EPTHook(address, 1, debugger.EPTHookTypeExecute)

	mylog.Info(fmt.Sprintf("0x%x", address))
	mylog.Info("EPT hook set successfully")
	return nil
}

type CommandMonitor struct {
	dbg    debugger.Packeter
	packet *debugger.Packet
}

func NewCommandMonitor(dbg debugger.Packeter, packet *debugger.Packet) *CommandMonitor {
	return &CommandMonitor{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandMonitor) Help() {
	mylog.Info("!monitor : monitor memory access.")
	mylog.Info("syntax : \t!monitor [address] [size] [read] [write]")
	mylog.Info("  address - virtual address to monitor (hex or decimal)")
	mylog.Info("  size    - size in bytes (hex or decimal)")
	mylog.Info("  read    - monitor read access (true/false)")
	mylog.Info("  write   - monitor write access (true/false)")
}

func (c *CommandMonitor) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 3 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	address := mylog.Check2(c.parseAddress(tokens[1].Value))

	size := mylog.Check2(strconv.ParseUint(tokens[2].Value, 0, 32))

	read := false
	write := false

	if len(tokens) > 3 {
		read = mylog.Check2(strconv.ParseBool(tokens[3].Value))
	}

	if len(tokens) > 4 {
		write = mylog.Check2(strconv.ParseBool(tokens[4].Value))
	}

	return c.MonitorMemory(address, uint32(size), read, write)
}

func (c *CommandMonitor) parseAddress(addr string) (uint64, error) {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X") {
		return strconv.ParseUint(addr[2:], 16, 64)
	}
	return strconv.ParseUint(addr, 0, 64)
}

func (c *CommandMonitor) MonitorMemory(address uint64, size uint32, read bool, write bool) error {
	if c.packet == nil || !c.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	var monitorType debugger.MonitorType
	if read && write {
		monitorType = debugger.MonitorTypeReadWrite
	} else if read {
		monitorType = debugger.MonitorTypeRead
	} else if write {
		monitorType = debugger.MonitorTypeWrite
	} else {
		return fmt.Errorf("must specify at least read or write")
	}

	c.packet.MonitorMemory(address, size, monitorType)

	mylog.Info(fmt.Sprintf("0x%x", address), fmt.Sprintf("0x%x", size), read, write)
	mylog.Info("Memory monitor set successfully")
	return nil
}

type CommandPte struct {
	dbg    debugger.Packeter
	packet *debugger.Packet
}

func NewCommandPte(dbg debugger.Packeter, packet *debugger.Packet) *CommandPte {
	return &CommandPte{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandPte) Help() {
	mylog.Info("!pte : query page table entry.")
	mylog.Info("syntax : \t!pte [virtual_address] [pid]")
	mylog.Info("  virtual_address - virtual address to query (hex or decimal)")
	mylog.Info("  pid            - process id (optional, default: current process)")
}

func (c *CommandPte) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	virtualAddress := mylog.Check2(c.parseAddress(tokens[1].Value))

	processId := uint32(0)
	if len(tokens) > 2 {
		pid := mylog.Check2(strconv.ParseUint(tokens[2].Value, 0, 32))

		processId = uint32(pid)
	}

	return c.QueryPte(virtualAddress, processId)
}

func (c *CommandPte) parseAddress(addr string) (uint64, error) {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X") {
		return strconv.ParseUint(addr[2:], 16, 64)
	}
	return strconv.ParseUint(addr, 0, 64)
}

func (c *CommandPte) QueryPte(virtualAddress uint64, processId uint32) error {
	if c.packet == nil || !c.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	mylog.Info(fmt.Sprintf("0x%x", virtualAddress), processId)
	mylog.Info("PTE query completed (not implemented)")
	return nil
}

type CommandVa2Pa struct {
	dbg    debugger.Packeter
	packet *debugger.Packet
}

func NewCommandVa2Pa(dbg debugger.Packeter, packet *debugger.Packet) *CommandVa2Pa {
	return &CommandVa2Pa{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandVa2Pa) Help() {
	mylog.Info("!va2pa : convert virtual address to physical address.")
	mylog.Info("syntax : \t!va2pa [virtual_address] [pid]")
	mylog.Info("  virtual_address - virtual address to convert (hex or decimal)")
	mylog.Info("  pid            - process id (optional, default: current process)")
}

func (c *CommandVa2Pa) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	virtualAddress := mylog.Check2(c.parseAddress(tokens[1].Value))

	processId := uint32(0)
	if len(tokens) > 2 {
		pid := mylog.Check2(strconv.ParseUint(tokens[2].Value, 0, 32))

		processId = uint32(pid)
	}

	return c.Va2Pa(virtualAddress, processId)
}

func (c *CommandVa2Pa) parseAddress(addr string) (uint64, error) {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X") {
		return strconv.ParseUint(addr[2:], 16, 64)
	}
	return strconv.ParseUint(addr, 0, 64)
}

func (c *CommandVa2Pa) Va2Pa(virtualAddress uint64, processId uint32) error {
	if c.packet == nil || !c.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	physicalAddress := c.packet.VA2PA(virtualAddress, processId)

	mylog.Info(fmt.Sprintf("0x%x", virtualAddress), fmt.Sprintf("0x%x", physicalAddress), processId)
	return nil
}

type CommandPa2Va struct {
	dbg    debugger.Packeter
	packet *debugger.Packet
}

func NewCommandPa2Va(dbg debugger.Packeter, packet *debugger.Packet) *CommandPa2Va {
	return &CommandPa2Va{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandPa2Va) Help() {
	mylog.Info("!pa2va : convert physical address to virtual address.")
	mylog.Info("syntax : \t!pa2va [physical_address] [pid]")
	mylog.Info("  physical_address - physical address to convert (hex or decimal)")
	mylog.Info("  pid             - process id (optional, default: current process)")
}

func (c *CommandPa2Va) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	physicalAddress := mylog.Check2(c.parseAddress(tokens[1].Value))

	processId := uint32(0)
	if len(tokens) > 2 {
		pid := mylog.Check2(strconv.ParseUint(tokens[2].Value, 0, 32))

		processId = uint32(pid)
	}

	return c.Pa2Va(physicalAddress, processId)
}

func (c *CommandPa2Va) parseAddress(addr string) (uint64, error) {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X") {
		return strconv.ParseUint(addr[2:], 16, 64)
	}
	return strconv.ParseUint(addr, 0, 64)
}

func (c *CommandPa2Va) Pa2Va(physicalAddress uint64, processId uint32) error {
	if c.packet == nil || !c.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	virtualAddress := c.packet.PA2VA(physicalAddress, processId)

	mylog.Info(fmt.Sprintf("0x%x", physicalAddress), fmt.Sprintf("0x%x", virtualAddress), processId)
	return nil
}
