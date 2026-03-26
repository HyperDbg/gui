// Code generated from Debugger. DO NOT EDIT.

package mcp

import (
	"context"
	"fmt"

	"hyperdbg-communicator/debugger"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CommunicatorMCPServer struct {
	impl debugger.Debugger
}

func NewCommunicatorMCPServer(impl debugger.Debugger) *CommunicatorMCPServer {
	return &CommunicatorMCPServer{impl: impl}
}

func (s *CommunicatorMCPServer) RegisterTools(server *mcp.Server) {
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Start",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"port": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleStart)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Stop",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleStop)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "WaitForDriver",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"timeout": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleWaitForDriver)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Initialize",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleInitialize)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Terminate",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleTerminate)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "AttachProcess",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"processID": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleAttachProcess)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "DetachProcess",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleDetachProcess)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "SetBreakpoint",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"address": map[string]any{
					"type": "number",
				},
				
				"bpType": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleSetBreakpoint)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "RemoveBreakpoint",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"breakpointID": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleRemoveBreakpoint)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Continue",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleContinue)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Pause",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handlePause)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "StepInto",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleStepInto)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "StepOver",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleStepOver)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "StepOut",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleStepOut)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ReadMemory",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"address": map[string]any{
					"type": "number",
				},
				
				"size": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleReadMemory)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "WriteMemory",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"address": map[string]any{
					"type": "number",
				},
				
				"data": map[string]any{
					"type": "array",
				},
				
			},
		},
	}, s.handleWriteMemory)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ReadRegisters",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleReadRegisters)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "WriteRegisters",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"regs": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleWriteRegisters)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetState",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleGetState)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "WaitForEvent",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"timeout": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleWaitForEvent)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ExecuteScript",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"script": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleExecuteScript)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ExecuteScriptWithContext",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"ctx": map[string]any{
					"type": "string",
				},
				
				"script": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleExecuteScriptWithContext)
	
}



type CommunicatorMCPServerStartParams struct {
	
	port int `json:"port"`
	
}


func (s *CommunicatorMCPServer) handleStart(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerStartParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Start(
		
		
		args.port,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerStopParams struct {}


func (s *CommunicatorMCPServer) handleStop(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerStopParams) (*mcp.CallToolResult, any, error) {
	
	
	s.impl.Stop(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerWaitForDriverParams struct {
	
	timeout any `json:"timeout"`
	
}


func (s *CommunicatorMCPServer) handleWaitForDriver(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerWaitForDriverParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.WaitForDriver(
		
		
		args.timeout.(debugger.time.Duration),
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerInitializeParams struct {}


func (s *CommunicatorMCPServer) handleInitialize(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerInitializeParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Initialize(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerTerminateParams struct {}


func (s *CommunicatorMCPServer) handleTerminate(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerTerminateParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Terminate(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerAttachProcessParams struct {
	
	processID uint32 `json:"processID"`
	
}


func (s *CommunicatorMCPServer) handleAttachProcess(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerAttachProcessParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.AttachProcess(
		
		
		args.processID,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerDetachProcessParams struct {}


func (s *CommunicatorMCPServer) handleDetachProcess(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerDetachProcessParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.DetachProcess(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerSetBreakpointParams struct {
	
	address uint64 `json:"address"`
	
	bpType any `json:"bpType"`
	
}


func (s *CommunicatorMCPServer) handleSetBreakpoint(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerSetBreakpointParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.SetBreakpoint(
		
		
		args.address,
		
		args.bpType.(debugger.BreakpointType),
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerRemoveBreakpointParams struct {
	
	breakpointID uint64 `json:"breakpointID"`
	
}


func (s *CommunicatorMCPServer) handleRemoveBreakpoint(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerRemoveBreakpointParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.RemoveBreakpoint(
		
		
		args.breakpointID,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerContinueParams struct {}


func (s *CommunicatorMCPServer) handleContinue(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerContinueParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Continue(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerPauseParams struct {}


func (s *CommunicatorMCPServer) handlePause(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerPauseParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Pause(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerStepIntoParams struct {}


func (s *CommunicatorMCPServer) handleStepInto(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerStepIntoParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.StepInto(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerStepOverParams struct {}


func (s *CommunicatorMCPServer) handleStepOver(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerStepOverParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.StepOver(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerStepOutParams struct {}


func (s *CommunicatorMCPServer) handleStepOut(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerStepOutParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.StepOut(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerReadMemoryParams struct {
	
	address uint64 `json:"address"`
	
	size uint32 `json:"size"`
	
}


func (s *CommunicatorMCPServer) handleReadMemory(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerReadMemoryParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.ReadMemory(
		
		
		args.address,
		
		args.size,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type CommunicatorMCPServerWriteMemoryParams struct {
	
	address uint64 `json:"address"`
	
	data []byte `json:"data"`
	
}


func (s *CommunicatorMCPServer) handleWriteMemory(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerWriteMemoryParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.WriteMemory(
		
		
		args.address,
		
		args.data,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerReadRegistersParams struct {}


func (s *CommunicatorMCPServer) handleReadRegisters(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerReadRegistersParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.ReadRegisters(
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type CommunicatorMCPServerWriteRegistersParams struct {
	
	regs any `json:"regs"`
	
}


func (s *CommunicatorMCPServer) handleWriteRegisters(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerWriteRegistersParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.WriteRegisters(
		
		
		args.regs.(*debugger.RegisterState),
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type CommunicatorMCPServerGetStateParams struct {}


func (s *CommunicatorMCPServer) handleGetState(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerGetStateParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.GetState(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type CommunicatorMCPServerWaitForEventParams struct {
	
	timeout any `json:"timeout"`
	
}


func (s *CommunicatorMCPServer) handleWaitForEvent(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerWaitForEventParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.WaitForEvent(
		
		
		args.timeout.(debugger.time.Duration),
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type CommunicatorMCPServerExecuteScriptParams struct {
	
	script string `json:"script"`
	
}


func (s *CommunicatorMCPServer) handleExecuteScript(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerExecuteScriptParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.ExecuteScript(
		
		
		args.script,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type CommunicatorMCPServerExecuteScriptWithContextParams struct {
	
	ctx any `json:"ctx"`
	
	script string `json:"script"`
	
}


func (s *CommunicatorMCPServer) handleExecuteScriptWithContext(ctx context.Context, req *mcp.CallToolRequest, args CommunicatorMCPServerExecuteScriptWithContextParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.ExecuteScriptWithContext(
		
		
		args.ctx.(debugger.context.Context),
		
		args.script,
		
		
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}

