// Code generated from Debugger. DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ddkwork/HyperDbg_rust/debugger"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type DebuggerMCPServer struct {
	impl debugger.Debugger
}

func NewDebuggerMCPServer(impl debugger.Debugger) *DebuggerMCPServer {
	return &DebuggerMCPServer{impl: impl}
}

func (s *DebuggerMCPServer) RegisterTools(server *mcp.Server) {
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Start",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
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
		Name:        "IsConnected",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleIsConnected)
	
	
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
		Name:        "Status",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleStatus)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Ping",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handlePing)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "LoadVmm",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleLoadVmm)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "UnloadVmm",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleUnloadVmm)
	
	
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
		Name:        "GetProcessList",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleGetProcessList)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetThreadList",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"processID": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleGetThreadList)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetModuleList",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"processID": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleGetModuleList)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "RegisterCallback",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"msgType": map[string]any{
					"type": "string",
				},
				
				"cb": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleRegisterCallback)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "RegisterCpuidCallback",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"cb": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleRegisterCpuidCallback)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetEvent",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleGetEvent)
	
	
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
		Name:        "GetConnectedDrivers",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleGetConnectedDrivers)
	
	
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
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "Disassemble",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"address": map[string]any{
					"type": "number",
				},
				
				"bytes": map[string]any{
					"type": "array",
				},
				
				"maxInstructions": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleDisassemble)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "LoadSymbols",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"pdbPath": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleLoadSymbols)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "UnloadSymbols",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
			},
		},
	}, s.handleUnloadSymbols)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetSymbolByName",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"name": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleGetSymbolByName)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetSymbolByAddress",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"address": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleGetSymbolByAddress)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "GetFunctionByAddress",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"address": map[string]any{
					"type": "number",
				},
				
			},
		},
	}, s.handleGetFunctionByAddress)
	
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "InstallHookScript",
		Description: "",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				
				"script": map[string]any{
					"type": "string",
				},
				
			},
		},
	}, s.handleInstallHookScript)
	
}



type DebuggerMCPServerStartParams struct {}


func (s *DebuggerMCPServer) handleStart(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerStartParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Start(
		
		
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


type DebuggerMCPServerStopParams struct {}


func (s *DebuggerMCPServer) handleStop(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerStopParams) (*mcp.CallToolResult, any, error) {
	
	
	s.impl.Stop(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type DebuggerMCPServerIsConnectedParams struct {}


func (s *DebuggerMCPServer) handleIsConnected(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerIsConnectedParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.IsConnected(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerGetStateParams struct {}


func (s *DebuggerMCPServer) handleGetState(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetStateParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.GetState(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerStatusParams struct {}


func (s *DebuggerMCPServer) handleStatus(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerStatusParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.Status(
		
		
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


type DebuggerMCPServerPingParams struct {}


func (s *DebuggerMCPServer) handlePing(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerPingParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.Ping(
		
		
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


type DebuggerMCPServerLoadVmmParams struct {}


func (s *DebuggerMCPServer) handleLoadVmm(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerLoadVmmParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.LoadVmm(
		
		
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


type DebuggerMCPServerUnloadVmmParams struct {}


func (s *DebuggerMCPServer) handleUnloadVmm(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerUnloadVmmParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.UnloadVmm(
		
		
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


type DebuggerMCPServerAttachProcessParams struct {
	
	processID uint32 `json:"processID"`
	
}


func (s *DebuggerMCPServer) handleAttachProcess(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerAttachProcessParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerDetachProcessParams struct {}


func (s *DebuggerMCPServer) handleDetachProcess(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerDetachProcessParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerSetBreakpointParams struct {
	
	address uint64 `json:"address"`
	
	bpType any `json:"bpType"`
	
}


func (s *DebuggerMCPServer) handleSetBreakpoint(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerSetBreakpointParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerRemoveBreakpointParams struct {
	
	breakpointID uint64 `json:"breakpointID"`
	
}


func (s *DebuggerMCPServer) handleRemoveBreakpoint(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerRemoveBreakpointParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerContinueParams struct {}


func (s *DebuggerMCPServer) handleContinue(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerContinueParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerPauseParams struct {}


func (s *DebuggerMCPServer) handlePause(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerPauseParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerStepIntoParams struct {}


func (s *DebuggerMCPServer) handleStepInto(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerStepIntoParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerStepOverParams struct {}


func (s *DebuggerMCPServer) handleStepOver(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerStepOverParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerStepOutParams struct {}


func (s *DebuggerMCPServer) handleStepOut(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerStepOutParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerReadMemoryParams struct {
	
	address uint64 `json:"address"`
	
	size uint32 `json:"size"`
	
}


func (s *DebuggerMCPServer) handleReadMemory(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerReadMemoryParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerWriteMemoryParams struct {
	
	address uint64 `json:"address"`
	
	data []byte `json:"data"`
	
}


func (s *DebuggerMCPServer) handleWriteMemory(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerWriteMemoryParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerReadRegistersParams struct {}


func (s *DebuggerMCPServer) handleReadRegisters(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerReadRegistersParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerWriteRegistersParams struct {
	
	regs any `json:"regs"`
	
}


func (s *DebuggerMCPServer) handleWriteRegisters(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerWriteRegistersParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerGetProcessListParams struct {}


func (s *DebuggerMCPServer) handleGetProcessList(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetProcessListParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.GetProcessList(
		
		
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


type DebuggerMCPServerGetThreadListParams struct {
	
	processID uint32 `json:"processID"`
	
}


func (s *DebuggerMCPServer) handleGetThreadList(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetThreadListParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.GetThreadList(
		
		
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
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerGetModuleListParams struct {
	
	processID uint32 `json:"processID"`
	
}


func (s *DebuggerMCPServer) handleGetModuleList(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetModuleListParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.GetModuleList(
		
		
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
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerRegisterCallbackParams struct {
	
	msgType any `json:"msgType"`
	
	cb any `json:"cb"`
	
}


func (s *DebuggerMCPServer) handleRegisterCallback(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerRegisterCallbackParams) (*mcp.CallToolResult, any, error) {
	
	
	s.impl.RegisterCallback(
		
		
		args.msgType.(debugger.MessageType),
		
		args.cb.(debugger.EventCallback),
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type DebuggerMCPServerRegisterCpuidCallbackParams struct {
	
	cb any `json:"cb"`
	
}


func (s *DebuggerMCPServer) handleRegisterCpuidCallback(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerRegisterCpuidCallbackParams) (*mcp.CallToolResult, any, error) {
	
	
	s.impl.RegisterCpuidCallback(
		
		
		args.cb.(debugger.EventCallback),
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type DebuggerMCPServerGetEventParams struct {}


func (s *DebuggerMCPServer) handleGetEvent(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetEventParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.GetEvent(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerWaitForEventParams struct {
	
	timeout any `json:"timeout"`
	
}


func (s *DebuggerMCPServer) handleWaitForEvent(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerWaitForEventParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.WaitForEvent(
		
		
		args.timeout.(time.Duration),
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerGetConnectedDriversParams struct {}


func (s *DebuggerMCPServer) handleGetConnectedDrivers(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetConnectedDriversParams) (*mcp.CallToolResult, any, error) {
	
	
	result := s.impl.GetConnectedDrivers(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	
	
}


type DebuggerMCPServerWaitForDriverParams struct {
	
	timeout any `json:"timeout"`
	
}


func (s *DebuggerMCPServer) handleWaitForDriver(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerWaitForDriverParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.WaitForDriver(
		
		
		args.timeout.(time.Duration),
		
		
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


type DebuggerMCPServerExecuteScriptParams struct {
	
	script string `json:"script"`
	
}


func (s *DebuggerMCPServer) handleExecuteScript(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerExecuteScriptParams) (*mcp.CallToolResult, any, error) {
	
	
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


type DebuggerMCPServerExecuteScriptWithContextParams struct {
	
	ctx any `json:"ctx"`
	
	script string `json:"script"`
	
}


func (s *DebuggerMCPServer) handleExecuteScriptWithContext(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerExecuteScriptWithContextParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.ExecuteScriptWithContext(
		
		
		args.ctx.(context.Context),
		
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


type DebuggerMCPServerDisassembleParams struct {
	
	address uint64 `json:"address"`
	
	bytes []byte `json:"bytes"`
	
	maxInstructions uint32 `json:"maxInstructions"`
	
}


func (s *DebuggerMCPServer) handleDisassemble(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerDisassembleParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.Disassemble(
		
		
		args.address,
		
		args.bytes,
		
		args.maxInstructions,
		
		
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


type DebuggerMCPServerLoadSymbolsParams struct {
	
	pdbPath string `json:"pdbPath"`
	
}


func (s *DebuggerMCPServer) handleLoadSymbols(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerLoadSymbolsParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.LoadSymbols(
		
		
		args.pdbPath,
		
		
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


type DebuggerMCPServerUnloadSymbolsParams struct {}


func (s *DebuggerMCPServer) handleUnloadSymbols(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerUnloadSymbolsParams) (*mcp.CallToolResult, any, error) {
	
	
	s.impl.UnloadSymbols(
		
		
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	
	
}


type DebuggerMCPServerGetSymbolByNameParams struct {
	
	name string `json:"name"`
	
}


func (s *DebuggerMCPServer) handleGetSymbolByName(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetSymbolByNameParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.GetSymbolByName(
		
		
		args.name,
		
		
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


type DebuggerMCPServerGetSymbolByAddressParams struct {
	
	address uint64 `json:"address"`
	
}


func (s *DebuggerMCPServer) handleGetSymbolByAddress(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetSymbolByAddressParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.GetSymbolByAddress(
		
		
		args.address,
		
		
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


type DebuggerMCPServerGetFunctionByAddressParams struct {
	
	address uint64 `json:"address"`
	
}


func (s *DebuggerMCPServer) handleGetFunctionByAddress(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerGetFunctionByAddressParams) (*mcp.CallToolResult, any, error) {
	
	
	result, err := s.impl.GetFunctionByAddress(
		
		
		args.address,
		
		
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


type DebuggerMCPServerInstallHookScriptParams struct {
	
	script any `json:"script"`
	
}


func (s *DebuggerMCPServer) handleInstallHookScript(ctx context.Context, req *mcp.CallToolRequest, args DebuggerMCPServerInstallHookScriptParams) (*mcp.CallToolResult, any, error) {
	
	
	err := s.impl.InstallHookScript(
		
		
		args.script.(*debugger.HookScript),
		
		
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


func main() {
	impl := debugger.NewPacket()
	server := NewDebuggerMCPServer(impl)

	mcpServer := mcp.NewServer(&mcp.Implementation{
		Name:    "hyperdbg-mcp",
		Version: "1.0.0",
	}, nil)
	server.RegisterTools(mcpServer)

	ctx := context.Background()
	if err := mcpServer.Run(ctx, &mcp.StdioTransport{}); err != nil {
		panic(err)
	}
}
