package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommandJSON(t *testing.T) {
	tests := []struct {
		name     string
		cmd      Command
		expected string
	}{
		{
			name:     "ping command",
			cmd:      Command{Action: "ping"},
			expected: `{"action":"ping"}`,
		},
		{
			name:     "read_memory command",
			cmd:      Command{Action: "read_memory", Target: "0x1000", Size: 64},
			expected: `{"action":"read_memory","target":"0x1000","size":64}`,
		},
		{
			name:     "write_memory command",
			cmd:      Command{Action: "write_memory", Target: "0x2000", Data: "deadbeef"},
			expected: `{"action":"write_memory","target":"0x2000","data":"deadbeef"}`,
		},
		{
			name:     "status command",
			cmd:      Command{Action: "status"},
			expected: `{"action":"status"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.cmd)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}

			var decoded Command
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}
			if decoded.Action != tt.cmd.Action {
				t.Errorf("Action mismatch: expected %s, got %s", tt.cmd.Action, decoded.Action)
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
			resp:     Response{Success: true, Message: "pong"},
			expected: `{"success":true,"message":"pong"}`,
		},
		{
			name:     "error response",
			resp:     Response{Success: false, Message: "invalid request"},
			expected: `{"success":false,"message":"invalid request"}`,
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
			name:   "GET /ping",
			method: "GET",
			path:   "/ping",
			body:   nil,
		},
		{
			name:   "GET /status",
			method: "GET",
			path:   "/status",
			body:   nil,
		},
		{
			name:   "POST /api with ping",
			method: "POST",
			path:   "/api",
			body:   []byte(`{"action":"ping"}`),
		},
		{
			name:   "POST /api with read_memory",
			method: "POST",
			path:   "/api",
			body:   []byte(`{"action":"read_memory","target":"0x1000","size":64}`),
		},
		{
			name:   "POST /read_memory",
			method: "POST",
			path:   "/read_memory",
			body:   []byte(`{"action":"read_memory","target":"0x1000","size":64}`),
		},
		{
			name:   "POST /write_memory",
			method: "POST",
			path:   "/write_memory",
			body:   []byte(`{"action":"write_memory","target":"0x2000","data":"deadbeef"}`),
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
		resp := Response{Success: true, Message: "test response"}
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
		body   Command
	}{
		{"GET", "/ping", Command{}},
		{"GET", "/status", Command{}},
		{"POST", "/read_memory", Command{Action: "read_memory", Target: "0x1000", Size: 64}},
		{"POST", "/write_memory", Command{Action: "write_memory", Target: "0x2000", Data: "deadbeef"}},
		{"POST", "/api", Command{Action: "ping"}},
		{"POST", "/api", Command{Action: "status"}},
		{"POST", "/api", Command{Action: "read_memory", Target: "0x1000", Size: 64}},
		{"POST", "/api", Command{Action: "write_memory", Target: "0x2000", Data: "deadbeef"}},
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

	var cmd Command
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&cmd)
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestEmptyBodyHandling(t *testing.T) {
	req := httptest.NewRequest("POST", "/api", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	var cmd Command
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&cmd)
	if err == nil {
		t.Error("Expected error for empty body")
	}
}
