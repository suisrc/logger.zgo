package logger

import (
	"context"
	"strconv"
)

//==================================================

// ContextTrace ...
type ContextTrace interface {
	GetTraceID() string
	GetTraceCIP() string
	GetTraceUID() string
}

// GetTraceID 从上下文中获取跟踪ID
func GetTraceID(ctx context.Context) string {
	if v, ok := ctx.(ContextTrace); ok {
		return v.GetTraceID()
	}
	return "main-" + strconv.Itoa(pid) // 系统上下文
}

// GetTraceCIP 从上下文中获取用户ID
func GetTraceCIP(ctx context.Context) string {
	if v, ok := ctx.(ContextTrace); ok {
		return v.GetTraceCIP()
	}
	return "0.0.0.0" // 未知IP
}

// GetTraceUID 从上下文中获取用户ID
func GetTraceUID(ctx context.Context) string {
	if v, ok := ctx.(ContextTrace); ok {
		return v.GetTraceUID()
	}
	return "none"
}
