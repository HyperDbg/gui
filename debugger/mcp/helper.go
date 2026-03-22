// Code generated from interfaces.go. DO NOT EDIT.

package mcp

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/std/mylog"
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
	bytes := mylog.Check2(json.Marshal(val))

	var result T
	json.Unmarshal(bytes, &result)
	return result
}

// mustConvert 使用 JSON 序列化/反序列化将 any 转换为目标类型
func mustConvert[T any](v any) T {
	if v == nil {
		var zero T
		return zero
	}
	bytes := mylog.Check2(json.Marshal(v))

	var result T
	mylog.Check(json.Unmarshal(bytes, &result))
	return result
}
