// Code generated from interfaces.go. DO NOT EDIT.

package mcp

import (
	"encoding/json"
)

func getArg[T any](args map[string]any, key string) T {
	if args == nil {
		var zero T
		return zero
	}
	val, ok := args[key]
	if !ok {
		var zero T
		return zero
	}
	bytes, err := json.Marshal(val)
	if err != nil {
		var zero T
		return zero
	}
	var result T
	if err := json.Unmarshal(bytes, &result); err != nil {
		var zero T
		return zero
	}
	return result
}

func mustConvert[T any](v any) T {
	if v == nil {
		var zero T
		return zero
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		var zero T
		return zero
	}
	var result T
	if err := json.Unmarshal(bytes, &result); err != nil {
		var zero T
		return zero
	}
	return result
}
