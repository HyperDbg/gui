package extension

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
)

type CommandEptHook struct {
	dbg    *debugger.HyperDbg
	packet *debugger.Packet
}

func NewCommandEptHook(dbg *debugger.HyperDbg, packet *debugger.Packet) *CommandEptHook {
	return &CommandEptHook{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandEptHook) Help() {
	slog.Info("!epthook : set EPT hook.")
	slog.Info("syntax : \t!epthook [address]")
	slog.Info("  address - virtual address to hook (hex or decimal)")
}

func (c *CommandEptHook) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	address, err := c.parseAddress(tokens[1].Value)
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}

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

	err := c.packet.SetEptHook(address)
	if err != nil {
		return err
	}

	slog.Info("Setting EPT hook", "address", fmt.Sprintf("0x%x", address))
	slog.Info("EPT hook set successfully")
	return nil
}

type CommandMonitor struct {
	dbg    *debugger.HyperDbg
	packet *debugger.Packet
}

func NewCommandMonitor(dbg *debugger.HyperDbg, packet *debugger.Packet) *CommandMonitor {
	return &CommandMonitor{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandMonitor) Help() {
	slog.Info("!monitor : monitor memory access.")
	slog.Info("syntax : \t!monitor [address] [size] [read] [write]")
	slog.Info("  address - virtual address to monitor (hex or decimal)")
	slog.Info("  size    - size in bytes (hex or decimal)")
	slog.Info("  read    - monitor read access (true/false)")
	slog.Info("  write   - monitor write access (true/false)")
}

func (c *CommandMonitor) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 3 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	address, err := c.parseAddress(tokens[1].Value)
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}

	size, err := strconv.ParseUint(tokens[2].Value, 0, 32)
	if err != nil {
		return fmt.Errorf("invalid size: %w", err)
	}

	read := false
	write := false

	if len(tokens) > 3 {
		read, err = strconv.ParseBool(tokens[3].Value)
		if err != nil {
			return fmt.Errorf("invalid read flag: %w", err)
		}
	}

	if len(tokens) > 4 {
		write, err = strconv.ParseBool(tokens[4].Value)
		if err != nil {
			return fmt.Errorf("invalid write flag: %w", err)
		}
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

	err := c.packet.MonitorMemory(address, size, read, write)
	if err != nil {
		return err
	}

	slog.Info("Monitoring memory", "address", fmt.Sprintf("0x%x", address),
		"size", fmt.Sprintf("0x%x", size), "read", read, "write", write)
	slog.Info("Memory monitor set successfully")
	return nil
}

type CommandPte struct {
	dbg    *debugger.HyperDbg
	packet *debugger.Packet
}

func NewCommandPte(dbg *debugger.HyperDbg, packet *debugger.Packet) *CommandPte {
	return &CommandPte{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandPte) Help() {
	slog.Info("!pte : query page table entry.")
	slog.Info("syntax : \t!pte [virtual_address] [pid]")
	slog.Info("  virtual_address - virtual address to query (hex or decimal)")
	slog.Info("  pid            - process id (optional, default: current process)")
}

func (c *CommandPte) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	virtualAddress, err := c.parseAddress(tokens[1].Value)
	if err != nil {
		return fmt.Errorf("invalid virtual address: %w", err)
	}

	processId := uint32(0)
	if len(tokens) > 2 {
		pid, err := strconv.ParseUint(tokens[2].Value, 0, 32)
		if err != nil {
			return fmt.Errorf("invalid process id: %w", err)
		}
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

	_, err := c.packet.QueryPte(virtualAddress, processId)
	if err != nil {
		return err
	}

	slog.Info("Querying PTE", "virtual_address", fmt.Sprintf("0x%x", virtualAddress), "pid", processId)
	slog.Info("PTE query completed")
	return nil
}

type CommandVa2Pa struct {
	dbg    *debugger.HyperDbg
	packet *debugger.Packet
}

func NewCommandVa2Pa(dbg *debugger.HyperDbg, packet *debugger.Packet) *CommandVa2Pa {
	return &CommandVa2Pa{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandVa2Pa) Help() {
	slog.Info("!va2pa : convert virtual address to physical address.")
	slog.Info("syntax : \t!va2pa [virtual_address] [pid]")
	slog.Info("  virtual_address - virtual address to convert (hex or decimal)")
	slog.Info("  pid            - process id (optional, default: current process)")
}

func (c *CommandVa2Pa) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	virtualAddress, err := c.parseAddress(tokens[1].Value)
	if err != nil {
		return fmt.Errorf("invalid virtual address: %w", err)
	}

	processId := uint32(0)
	if len(tokens) > 2 {
		pid, err := strconv.ParseUint(tokens[2].Value, 0, 32)
		if err != nil {
			return fmt.Errorf("invalid process id: %w", err)
		}
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

	physicalAddress, err := c.packet.Va2Pa(virtualAddress, processId)
	if err != nil {
		return err
	}

	slog.Info("Virtual Address -> Physical Address", "virtual", fmt.Sprintf("0x%x", virtualAddress),
		"physical", fmt.Sprintf("0x%x", physicalAddress), "pid", processId)
	return nil
}

type CommandPa2Va struct {
	dbg    *debugger.HyperDbg
	packet *debugger.Packet
}

func NewCommandPa2Va(dbg *debugger.HyperDbg, packet *debugger.Packet) *CommandPa2Va {
	return &CommandPa2Va{
		dbg:    dbg,
		packet: packet,
	}
}

func (c *CommandPa2Va) Help() {
	slog.Info("!pa2va : convert physical address to virtual address.")
	slog.Info("syntax : \t!pa2va [physical_address] [pid]")
	slog.Info("  physical_address - physical address to convert (hex or decimal)")
	slog.Info("  pid             - process id (optional, default: current process)")
}

func (c *CommandPa2Va) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	physicalAddress, err := c.parseAddress(tokens[1].Value)
	if err != nil {
		return fmt.Errorf("invalid physical address: %w", err)
	}

	processId := uint32(0)
	if len(tokens) > 2 {
		pid, err := strconv.ParseUint(tokens[2].Value, 0, 32)
		if err != nil {
			return fmt.Errorf("invalid process id: %w", err)
		}
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

	virtualAddress, err := c.packet.Pa2Va(physicalAddress, processId)
	if err != nil {
		return err
	}

	slog.Info("Physical Address -> Virtual Address", "physical", fmt.Sprintf("0x%x", physicalAddress),
		"virtual", fmt.Sprintf("0x%x", virtualAddress), "pid", processId)
	return nil
}
