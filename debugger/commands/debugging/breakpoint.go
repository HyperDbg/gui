package debugging

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
)

type Breakpoint struct {
	Address    uint64
	ProcessId  uint32
	CoreNumber uint32
	Type       uint32
	Tag        uint32
	Enabled    bool
}

type CommandBp struct {
	dbg         *debugger.HyperDbg
	breakpoints map[uint32]*Breakpoint
	nextTag     uint32
	packet      *debugger.Packet
}

func NewCommandBp(dbg *debugger.HyperDbg, packet *debugger.Packet) *CommandBp {
	return &CommandBp{
		dbg:         dbg,
		breakpoints: make(map[uint32]*Breakpoint),
		nextTag:     1,
		packet:      packet,
	}
}

func (c *CommandBp) Help() {
	slog.Info("bp : set breakpoint.")
	slog.Info("syntax : \tbp [address] [pid] [core] [type]")
	slog.Info("  address - breakpoint address (hex or decimal)")
	slog.Info("  pid     - process id (optional, default: 0)")
	slog.Info("  core    - core number (optional, default: 0)")
	slog.Info("  type    - breakpoint type (optional, default: 0)")
	slog.Info("Breakpoint types:")
	slog.Info("  0 - Software breakpoint")
	slog.Info("  1 - Hardware breakpoint (execute)")
	slog.Info("  2 - Hardware breakpoint (write)")
	slog.Info("  3 - Hardware breakpoint (read/write)")
}

func (c *CommandBp) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	address, err := c.parseAddress(tokens[1].Value)
	if err != nil {
		return fmt.Errorf("invalid address: %w", err)
	}

	processId := uint32(0)
	coreNumber := uint32(0)
	bpType := uint32(0)

	if len(tokens) > 2 {
		pid, err := strconv.ParseUint(tokens[2].Value, 0, 32)
		if err != nil {
			return fmt.Errorf("invalid process id: %w", err)
		}
		processId = uint32(pid)
	}

	if len(tokens) > 3 {
		core, err := strconv.ParseUint(tokens[3].Value, 0, 32)
		if err != nil {
			return fmt.Errorf("invalid core number: %w", err)
		}
		coreNumber = uint32(core)
	}

	if len(tokens) > 4 {
		typ, err := strconv.ParseUint(tokens[4].Value, 0, 32)
		if err != nil {
			return fmt.Errorf("invalid breakpoint type: %w", err)
		}
		bpType = uint32(typ)
	}

	return c.SetBreakpoint(address, processId, coreNumber, bpType)
}

func (c *CommandBp) parseAddress(addr string) (uint64, error) {
	addr = strings.TrimSpace(addr)
	if strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X") {
		return strconv.ParseUint(addr[2:], 16, 64)
	}
	return strconv.ParseUint(addr, 0, 64)
}

func (c *CommandBp) SetBreakpoint(address uint64, processId uint32, coreNumber uint32, bpType uint32) error {
	if c.packet == nil || !c.packet.IsConnected() {
		return fmt.Errorf("driver not available")
	}

	tag := c.nextTag
	c.nextTag++

	bp := &Breakpoint{
		Address:    address,
		ProcessId:  processId,
		CoreNumber: coreNumber,
		Type:       bpType,
		Tag:        tag,
		Enabled:    true,
	}

	err := c.packet.SetBreakpoint(processId, address, coreNumber, bpType, tag)
	if err != nil {
		return fmt.Errorf("failed to set breakpoint: %w", err)
	}

	c.breakpoints[tag] = bp
	slog.Info("Breakpoint set", "address", fmt.Sprintf("0x%x", address), "tag", tag)
	return nil
}

func (c *CommandBp) GetBreakpoint(tag uint32) (*Breakpoint, bool) {
	bp, exists := c.breakpoints[tag]
	return bp, exists
}

func (c *CommandBp) GetAllBreakpoints() []*Breakpoint {
	var bps []*Breakpoint
	for _, bp := range c.breakpoints {
		bps = append(bps, bp)
	}
	return bps
}

type CommandBc struct {
	dbg         *debugger.HyperDbg
	breakpoints map[uint32]*Breakpoint
	packet      *debugger.Packet
}

func NewCommandBc(dbg *debugger.HyperDbg, breakpoints map[uint32]*Breakpoint, packet *debugger.Packet) *CommandBc {
	return &CommandBc{
		dbg:         dbg,
		breakpoints: breakpoints,
		packet:      packet,
	}
}

func (c *CommandBc) Help() {
	slog.Info("bc : clear breakpoint.")
	slog.Info("syntax : \tbc [tag]")
	slog.Info("  tag - breakpoint tag to clear (use '*' to clear all)")
}

func (c *CommandBc) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	if tokens[1].Value == "*" {
		return c.ClearAllBreakpoints()
	}

	tag, err := strconv.ParseUint(tokens[1].Value, 0, 32)
	if err != nil {
		return fmt.Errorf("invalid tag: %w", err)
	}

	return c.ClearBreakpoint(uint32(tag))
}

func (c *CommandBc) ClearBreakpoint(tag uint32) error {
	if _, exists := c.breakpoints[tag]; !exists {
		return fmt.Errorf("breakpoint with tag %d not found", tag)
	}

	if c.packet != nil && c.packet.IsConnected() {
		err := c.packet.RemoveBreakpoint(tag)
		if err != nil {
			return fmt.Errorf("failed to clear breakpoint from driver: %w", err)
		}
	}

	delete(c.breakpoints, tag)
	slog.Info("Breakpoint cleared", "tag", tag)
	return nil
}

func (c *CommandBc) ClearAllBreakpoints() error {
	count := len(c.breakpoints)

	if c.packet != nil && c.packet.IsConnected() {
		for tag := range c.breakpoints {
			err := c.packet.RemoveBreakpoint(tag)
			if err != nil {
				slog.Warn("Failed to clear breakpoint", "tag", tag, "error", err)
			}
		}
	}

	for tag := range c.breakpoints {
		delete(c.breakpoints, tag)
	}
	slog.Info("Cleared breakpoints", "count", count)
	return nil
}

type CommandBl struct {
	dbg         *debugger.HyperDbg
	breakpoints map[uint32]*Breakpoint
}

func NewCommandBl(dbg *debugger.HyperDbg, breakpoints map[uint32]*Breakpoint) *CommandBl {
	return &CommandBl{
		dbg:         dbg,
		breakpoints: breakpoints,
	}
}

func (c *CommandBl) Help() {
	slog.Info("bl : list breakpoints.")
	slog.Info("syntax : \tbl")
}

func (c *CommandBl) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.ListBreakpoints()
}

func (c *CommandBl) ListBreakpoints() error {
	if len(c.breakpoints) == 0 {
		slog.Info("No breakpoints set")
		return nil
	}

	slog.Info("Breakpoints:")
	for tag, bp := range c.breakpoints {
		status := "enabled"
		if !bp.Enabled {
			status = "disabled"
		}
		slog.Info("Breakpoint", "tag", tag, "address", fmt.Sprintf("0x%x", bp.Address),
			"pid", bp.ProcessId, "core", bp.CoreNumber, "type", bp.Type, "status", status)
	}

	return nil
}
