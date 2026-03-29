package debugger

import (
	"encoding/json"
	"fmt"
	"time"
)

type HookFilter struct {
	ProcessNames  []string
	ProcessIds    []uint32
	ExcludeSystem bool
}

func (f *HookFilter) Match(processName string, processId uint32) bool {
	if len(f.ProcessNames) > 0 {
		for _, name := range f.ProcessNames {
			if name == processName {
				return true
			}
		}
		return false
	}
	if len(f.ProcessIds) > 0 {
		for _, pid := range f.ProcessIds {
			if pid == processId {
				return true
			}
		}
		return false
	}
	if f.ExcludeSystem && processId <= 4 {
		return false
	}
	return true
}

type HookEvent struct {
	EventId     uint64          `json:"event_id"`
	ApiName     string          `json:"api_name"`
	ProcessId   uint32          `json:"process_id"`
	ThreadId    uint32          `json:"thread_id"`
	ProcessName string          `json:"process_name"`
	Args        json.RawMessage `json:"args"`
	Timestamp   uint64          `json:"timestamp"`
}

type HookEventResponse struct {
	EventId      uint64      `json:"event_id"`
	Action       string      `json:"action"`
	ModifyArgs   bool        `json:"modify_args"`
	ModifiedArgs interface{} `json:"modified_args,omitempty"`
	SkipCall     bool        `json:"skip_call"`
	ReturnValue  interface{} `json:"return_value,omitempty"`
}

type HookHandler func(event *HookEvent) (*HookEventResponse, error)

type HookConfig struct {
	ApiName  string
	HookType HookType
	Filter   *HookFilter
	Handler  HookHandler
}

func (p *Packet) InstallHook(config *HookConfig) error {
	info := GetHookInfo(config.ApiName)
	if info == nil {
		return fmt.Errorf("unknown API: %s", config.ApiName)
	}

	req := map[string]interface{}{
		"action":    "install_hook",
		"api_name":  config.ApiName,
		"hook_type": config.HookType,
	}
	if config.Filter != nil {
		req["filter"] = map[string]interface{}{
			"process_names":  config.Filter.ProcessNames,
			"process_ids":    config.Filter.ProcessIds,
			"exclude_system": config.Filter.ExcludeSystem,
		}
	}

	data, _ := json.Marshal(req)
	resp := SendReceive[HookResponse](p, data)
	if resp == nil || !resp.Success {
		return fmt.Errorf("install hook failed: %s", resp.Message)
	}
	return nil
}

func (p *Packet) WaitForHookEvent(timeout time.Duration) *HookEvent {
	msg := p.WaitForEvent(timeout)
	if msg == nil {
		return nil
	}

	var event HookEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return nil
	}
	return &event
}

func (p *Packet) RespondHookEvent(resp *HookEventResponse) error {
	data, _ := json.Marshal(resp)
	msg := NewMessage(MessageTypeHookResponse, data)
	return p.Send(msg)
}

func Example_NtDeviceIoControlFile_Hook_EventLoop() {
	packet := NewPacket().(*Packet)

	err := packet.InstallHook(&HookConfig{
		ApiName:  "NtDeviceIoControlFile",
		HookType: HookTypeEPT,
		Filter: &HookFilter{
			ProcessNames: []string{"target.exe"},
		},
		Handler: func(event *HookEvent) (*HookEventResponse, error) {
			var args NtDeviceIoControlFileArgs
			if err := json.Unmarshal(event.Args, &args); err != nil {
				return nil, err
			}

			if args.IoControlCode == 0x12345678 {
				fmt.Printf("[%s] PID=%d IoControlCode=0x%X\n",
					event.ProcessName, event.ProcessId, args.IoControlCode)
				fmt.Printf("  InputBuffer: 0x%X (size=%d)\n",
					args.InputBuffer, args.InputBufferLength)
				fmt.Printf("  OutputBuffer: 0x%X (size=%d)\n",
					args.OutputBuffer, args.OutputBufferLength)

				return &HookEventResponse{
					EventId: event.EventId,
					Action:  "continue",
				}, nil
			}

			return &HookEventResponse{
				EventId: event.EventId,
				Action:  "pass",
			}, nil
		},
	})
	if err != nil {
		fmt.Printf("Failed to install hook: %v\n", err)
		return
	}

	fmt.Println("Hook installed, waiting for events...")

	for {
		event := packet.WaitForHookEvent(5 * time.Second)
		if event == nil {
			continue
		}

		info := GetHookInfo(event.ApiName)
		if info == nil {
			continue
		}

		handler := getHandler(event.ApiName)
		if handler == nil {
			packet.RespondHookEvent(&HookEventResponse{
				EventId: event.EventId,
				Action:  "pass",
			})
			continue
		}

		resp, err := handler(event)
		if err != nil {
			fmt.Printf("Handler error: %v\n", err)
			resp = &HookEventResponse{
				EventId: event.EventId,
				Action:  "pass",
			}
		}

		packet.RespondHookEvent(resp)
	}
}

var hookHandlers = make(map[string]HookHandler)

func getHandler(apiName string) HookHandler {
	return hookHandlers[apiName]
}

func registerHandler(apiName string, handler HookHandler) {
	hookHandlers[apiName] = handler
}
