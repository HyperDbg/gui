package debugging

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
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
	dbg         debugger.Packeter
	breakpoints map[uint32]*Breakpoint
	nextTag     uint32
	packet      *debugger.Packet
}

func NewCommandBp(dbg debugger.Packeter, packet *debugger.Packet) *CommandBp {
	return &CommandBp{
		dbg:         dbg,
		breakpoints: make(map[uint32]*Breakpoint),
		nextTag:     1,
		packet:      packet,
	}
}

func (c *CommandBp) Help() {
	mylog.Info("bp : set breakpoint.")
	mylog.Info("syntax : \tbp [address] [pid] [core] [type]")
	mylog.Info("  address - breakpoint address (hex or decimal)")
	mylog.Info("  pid     - process id (optional, default: 0)")
	mylog.Info("  core    - core number (optional, default: 0)")
	mylog.Info("  type    - breakpoint type (optional, default: 0)")
	mylog.Info("Breakpoint types:")
	mylog.Info("  0 - Software breakpoint")
	mylog.Info("  1 - Hardware breakpoint (execute)")
	mylog.Info("  2 - Hardware breakpoint (write)")
	mylog.Info("  3 - Hardware breakpoint (read/write)")
}

func (c *CommandBp) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	address := mylog.Check2(c.parseAddress(tokens[1].Value))

	processId := uint32(0)
	coreNumber := uint32(0)
	bpType := uint32(0)

	if len(tokens) > 2 {
		pid := mylog.Check2(strconv.ParseUint(tokens[2].Value, 0, 32))

		processId = uint32(pid)
	}

	if len(tokens) > 3 {
		core := mylog.Check2(strconv.ParseUint(tokens[3].Value, 0, 32))

		coreNumber = uint32(core)
	}

	if len(tokens) > 4 {
		typ := mylog.Check2(strconv.ParseUint(tokens[4].Value, 0, 32))

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

	switch dbg := c.dbg.(type) {
	case debugger.UserDebugger:
		(dbg.SetBreakpoint(address))
	case debugger.KernelDebugger:
		(dbg.SetBreakpoint(address, processId))
	default:
		return fmt.Errorf("unsupported debugger type")
	}

	c.breakpoints[tag] = bp
	mylog.Info(fmt.Sprintf("0x%x", address), tag)
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
	dbg         debugger.Packeter
	breakpoints map[uint32]*Breakpoint
	packet      *debugger.Packet
}

func NewCommandBc(dbg debugger.Packeter, breakpoints map[uint32]*Breakpoint, packet *debugger.Packet) *CommandBc {
	return &CommandBc{
		dbg:         dbg,
		breakpoints: breakpoints,
		packet:      packet,
	}
}

func (c *CommandBc) Help() {
	mylog.Info("bc : clear breakpoint.")
	mylog.Info("syntax : \tbc [tag]")
	mylog.Info("  tag - breakpoint tag to clear (use '*' to clear all)")
}

func (c *CommandBc) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	if tokens[1].Value == "*" {
		return c.ClearAllBreakpoints()
	}

	tag := mylog.Check2(strconv.ParseUint(tokens[1].Value, 0, 32))

	return c.ClearBreakpoint(uint32(tag))
}

func (c *CommandBc) ClearBreakpoint(tag uint32) error {
	if _, exists := c.breakpoints[tag]; !exists {
		return fmt.Errorf("breakpoint with tag %d not found", tag)
	}

	if c.packet != nil && c.packet.IsConnected() {
		if dbg, ok := c.dbg.(debugger.UserDebugger); ok {
			(dbg.RemoveBreakpoint(tag))
		}
	}

	delete(c.breakpoints, tag)
	mylog.Info(tag)
	return nil
}

func (c *CommandBc) ClearAllBreakpoints() error {
	count := len(c.breakpoints)

	if c.packet != nil && c.packet.IsConnected() {
		if dbg, ok := c.dbg.(debugger.UserDebugger); ok {
			for tag := range c.breakpoints {
				(dbg.RemoveBreakpoint(tag))
			}
		}
	}

	for tag := range c.breakpoints {
		delete(c.breakpoints, tag)
	}
	mylog.Info(count)
	return nil
}

type CommandBl struct {
	dbg         debugger.Packeter
	breakpoints map[uint32]*Breakpoint
}

func NewCommandBl(dbg debugger.Packeter, breakpoints map[uint32]*Breakpoint) *CommandBl {
	return &CommandBl{
		dbg:         dbg,
		breakpoints: breakpoints,
	}
}

func (c *CommandBl) Help() {
	mylog.Info("bl : list breakpoints.")
	mylog.Info("syntax : \tbl")
}

func (c *CommandBl) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.ListBreakpoints()
}

func (c *CommandBl) ListBreakpoints() error {
	if len(c.breakpoints) == 0 {
		mylog.Info("No breakpoints set")
		return nil
	}

	mylog.Info("Breakpoints:")
	for tag, bp := range c.breakpoints {
		status := "enabled"
		if !bp.Enabled {
			status = "disabled"
		}
		mylog.Info(tag, fmt.Sprintf("0x%x", bp.Address), bp.ProcessId, bp.CoreNumber, bp.Type, status)
	}

	return nil
}
