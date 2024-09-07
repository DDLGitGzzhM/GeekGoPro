package context

import (
	"context"
	"testing"
	"time"
)

func Test_ContextCancel(t *testing.T) {
	ctx := context.Background()
	subCtx1 := context.WithValue(ctx, "DoSelect", "success")

	cancelCtx, cancel := context.WithCancel(subCtx1)
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	_ = <-cancelCtx.Done()
	t.Log(cancelCtx.Err())
	t.Log(cancelCtx.Value("DoSelect"))
}

func Test_ContextWithDeadLine(t *testing.T) {
	ctx := context.Background()
	cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	defer cancel()
	<-cancelCtx.Done()
	t.Log(cancelCtx.Deadline())
	t.Log(cancelCtx.Err())
}

func Test_ContextWithTimeOut(t *testing.T) {
	ctx := context.Background()
	cancelCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	<-cancelCtx.Done()
	t.Log(cancelCtx.Deadline())
	t.Log(cancelCtx.Err())
}
