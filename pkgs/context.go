package pkgs

import (
	"context"
	"fmt"
	"time"
)

// SleuthContext 调用链路
func SleuthContext() context.Context {
	return context.WithValue(context.Background(), `sleuth_code`, fmt.Sprintf("ctx-%d", time.Now().Unix()))
}

// SleuthCtx ctx
type SleuthCtx struct {
	SleuthCode string `json:"sleuth_code"`
}

// GetSleuthCtx 获取调用链
func GetSleuthCtx(ctx context.Context) SleuthCtx {
	return SleuthCtx{SleuthCode: fmt.Sprintf("%v", ctx.Value("sleuth_code"))}
}
