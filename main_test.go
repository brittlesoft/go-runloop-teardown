package main

import (
	"context"
	"testing"
	"time"

	"github.com/brittlesoft/go-runloop-teardown/internal/producing"
	"github.com/brittlesoft/go-runloop-teardown/internal/recording"
	"golang.org/x/sync/errgroup"
)

func TestPatate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	recorder := recording.NewRecorder()
	producer := producing.NewProducer(recorder)

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error { return recorder.Run(egCtx) })
	eg.Go(func() error { return producer.Run(egCtx) })

	doneCh := make(chan error)
	go func() {
		doneCh <- eg.Wait()
	}()

	timeout := time.NewTimer(time.Second)
	select {
	case <-doneCh:
	case <-timeout.C:
		t.Fatal("timed out")
	}
}
