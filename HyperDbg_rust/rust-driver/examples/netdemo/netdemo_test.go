package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestJSON(t *testing.T) {
	tests := []struct {
		name     string
		req      Request
		expected string
	}{
		{
			name:     "initialize command",
			req:      Request{Action: "initialize"},
			expected: `{"action":"initialize"}`,
		},
		{
			name:     "read_memory command",
			req:      Request{Action: "read_memory", Address: new(uint64(0x1000)), Size: new(uint32(64))},
			expected: `{"action":"read_memory","address":4096,"size":64}`,
		},
		{
			name:     "write_memory command",
			req:      Request{Action: "write_memory", Address: new(uint64(0x2000)), Data: []byte{0xDE, 0xAD, 0xBE, 0xEF}},
			expected: `{"action":"write_memory","address":8192,"data":[222,173,190,239]}`,
		},
		{
			name:     "attach_process command",
			req:      Request{Action: "attach_process", ProcessID: new(uint32(1234))},
			expected: `{"action":"attach_process","process_id":1234}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.req)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}

			var decoded Request
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}
			if decoded.Action != tt.req.Action {
				t.Errorf("Action mismatch: expected %s, got %s", tt.req.Action, decoded.Action)
			}
		})
	}
}

func TestResponseJSON(t *testing.T) {
	tests := []struct {
		name     string
		resp     Response
		expected string
	}{
		{
			name:     "success response",
			resp:     Response{Success: true, Message: "pong", Data: nil},
			expected: `{"success":true,"message":"pong"}`,
		},
		{
			name:     "error response",
			resp:     Response{Success: false, Message: "invalid request", Data: nil},
			expected: `{"success":false,"message":"invalid request"}`,
		},
		{
			name:     "response with data",
			resp:     Response{Success: true, Message: "data retrieved", Data: []byte{0x01, 0x02, 0x03}},
			expected: `{"success":true,"message":"data retrieved","data":[1,2,3]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.resp)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}

func TestHTTPRequestFormat(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		body   []byte
	}{
		{
			name:   "POST /api with initialize",
			method: "POST",
			path:   "/api",
			body:   []byte(`{"action":"initialize"}`),
		},
		{
			name:   "POST /api with read_memory",
			method: "POST",
			path:   "/api",
			body:   []byte(`{"action":"read_memory","address":4096,"size":64}`),
		},
		{
			name:   "POST /api with write_memory",
			method: "POST",
			path:   "/api",
			body:   []byte(`{"action":"write_memory","address":8192,"data":[222,173,190,239]}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, bytes.NewBuffer(tt.body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Host", "127.0.0.1")

			if req.Method != tt.method {
				t.Errorf("Method mismatch: expected %s, got %s", tt.method, req.Method)
			}
			if req.URL.Path != tt.path {
				t.Errorf("Path mismatch: expected %s, got %s", tt.path, req.URL.Path)
			}
			if req.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Content-Type mismatch: expected application/json, got %s", req.Header.Get("Content-Type"))
			}
			if req.Header.Get("Host") != "127.0.0.1" {
				t.Errorf("Host mismatch: expected 127.0.0.1, got %s", req.Header.Get("Host"))
			}
		})
	}
}

func TestHTTPResponseParsing(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := Response{Success: true, Message: "test response", Data: "test data"}
		json.NewEncoder(w).Encode(resp)
	}

	req := httptest.NewRequest("POST", "/api", bytes.NewBuffer([]byte(`{"action":"ping"}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status mismatch: expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}
	if result.Message != "test response" {
		t.Errorf("Message mismatch: expected 'test response', got '%s'", result.Message)
	}
}

func TestAllEndpoints(t *testing.T) {
	endpoints := []struct {
		method string
		path   string
		body   Request
	}{
		{"POST", "/api", Request{Action: "initialize"}},
		{"POST", "/api", Request{Action: "terminate"}},
		{"POST", "/api", Request{Action: "attach_process", ProcessID: new(uint32(1234))}},
		{"POST", "/api", Request{Action: "read_memory", Address: new(uint64(0x1000)), Size: new(uint32(64))}},
		{"POST", "/api", Request{Action: "write_memory", Address: new(uint64(0x2000)), Data: []byte{0xDE, 0xAD, 0xBE, 0xEF}}},
	}

	for _, ep := range endpoints {
		t.Run(fmt.Sprintf("%s %s", ep.method, ep.path), func(t *testing.T) {
			var body []byte
			if ep.method == "POST" {
				var err error
				body, err = json.Marshal(ep.body)
				if err != nil {
					t.Fatalf("Failed to marshal body: %v", err)
				}
			}

			req := httptest.NewRequest(ep.method, ep.path, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Host", "127.0.0.1")

			if req.Method != ep.method {
				t.Errorf("Method mismatch")
			}
			if req.URL.Path != ep.path {
				t.Errorf("Path mismatch")
			}
		})
	}
}

func TestHTTPHeaders(t *testing.T) {
	req := httptest.NewRequest("POST", "/api", bytes.NewBuffer([]byte(`{"action":"ping"}`)))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "127.0.0.1")
	req.Header.Set("Content-Length", "19")

	if req.Header.Get("Content-Type") != "application/json" {
		t.Error("Content-Type header mismatch")
	}
	if req.Header.Get("Host") != "127.0.0.1" {
		t.Error("Host header mismatch")
	}
	if req.Header.Get("Content-Length") != "19" {
		t.Error("Content-Length header mismatch")
	}
}

func TestInvalidJSONHandling(t *testing.T) {
	invalidJSON := []byte(`{"action": invalid}`)
	req := httptest.NewRequest("POST", "/api", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	var reqBody Request
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&reqBody)
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestEmptyBodyHandling(t *testing.T) {
	req := httptest.NewRequest("POST", "/api", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	var reqBody Request
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&reqBody)
	if err == nil {
		t.Error("Expected error for empty body")
	}
}

func TestOptionalFields(t *testing.T) {
	tests := []struct {
		name string
		req  Request
	}{
		{
			name: "all fields",
			req: Request{
				Action:       "read_memory",
				ProcessID:    new(uint32(1234)),
				Address:      new(uint64(0x1000)),
				BreakpointID: new(uint64(1)),
				Size:         new(uint32(64)),
				Type:         new(int32(1)),
				Data:         []byte{0x01, 0x02, 0x03},
			},
		},
		{
			name: "only action",
			req:  Request{Action: "ping"},
		},
		{
			name: "action with address",
			req:  Request{Action: "read_memory", Address: new(uint64(0x1000))},
		},
		{
			name: "action with process_id",
			req:  Request{Action: "attach_process", ProcessID: new(uint32(1234))},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.req)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			var decoded Request
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if decoded.Action != tt.req.Action {
				t.Errorf("Action mismatch: expected %s, got %s", tt.req.Action, decoded.Action)
			}
		})
	}
}
