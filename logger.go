package logger

import (
	"context"
	"fmt"
	"os"
)

// Debugf 写入调试日志
func Debugf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Debugf(format, args...)
}

// Infof 写入消息日志
func Infof(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Infof(format, args...)
}

// Printf 写入消息日志
func Printf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Printf(format, args...)
}

// Warnf 写入警告日志
func Warnf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Warnf(format, args...)
}

// Errorf 写入错误日志
func Errorf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Errorf(format, args...)
}

// Fatalf 写入重大错误日志
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Fatalf(format, args...)
}

// ErrorStack 输出错误栈
func ErrorStack(ctx context.Context, err error) {
	StartTrace(ctx).WithField(StackKey, fmt.Sprintf("%+v", err)).Errorf(err.Error())
}

//=================================================================分割线
//=================================================================分割线
//=================================================================分割线

// 定义键名
const (
	TraceIDKey   = "trace_id"
	UserIDKey    = "user_id"
	RoleIDKey    = "role_id"
	VersionKey   = "version"
	HostnameKey  = "hostname"
	ClientIPKey  = "clientip"
	StackKey     = "stack"
	LogVerKey    = "@version"
	NamespaceKey = "@namespace"
)

var (
	version     string
	pid         = os.Getpid()
	hostname, _ = os.Hostname()
)

// SetVersion 设定版本
func SetVersion(v string) {
	version = v
}

// @see github.com/suisrc/zgo/helper/helper.go

// 定义上下文中的键
const (
	prefix      = "zgo"
	KeyUserInfo = prefix + ":user-info"
	KeyTraceID  = prefix + ":tract-id"
)

// StartTrace 开始一个追踪单元
func StartTrace(ctx context.Context) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := map[string]interface{}{
		LogVerKey:    logversion,
		NamespaceKey: namespace,
		HostnameKey:  hostname,
		VersionKey:   version,
	}
	if v := GetTraceID(ctx); v != "" {
		fields[TraceIDKey] = v
	}
	if v := GetTraceUID(ctx); v != "" {
		fields[UserIDKey] = v
	}
	if v := GetTraceCIP(ctx); v != "" {
		fields[ClientIPKey] = v
	}

	return newEntryWithFields(fields)
}
