package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brittlesoft/go-runloop-teardown/internal/producing"
	"github.com/brittlesoft/go-runloop-teardown/internal/recording"
	"golang.org/x/sync/errgroup"
)

func TestPatate(t *testing.T) {
	tcs := []string{"run", "runctx", "runctxselect"}
	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			t.Cleanup(cancel)

			recorder := recording.NewRecorder()
			producer := producing.NewProducer(recorder)

			eg, egCtx := errgroup.WithContext(ctx)
			eg.Go(func() error { return recorder.Run(egCtx) })
			eg.Go(func() error {
				switch tc {
				case "run":
					return producer.Run(egCtx)
				case "runctx":
					return producer.RunCtx(egCtx)
				case "runctxselect":
					return producer.RunCtxSelect(egCtx)
				default:
					return fmt.Errorf("invalid test case: %s", tc)
				}
			})

			go func() {
				time.Sleep(10 * time.Millisecond)
				cancel()
			}()

			doneCh := make(chan error)
			go func() {
				doneCh <- eg.Wait()
			}()

			timeout := time.NewTimer(20 * time.Millisecond)
			select {
			case <-doneCh:
			case <-timeout.C:
				t.Log("timed out")
				t.Fail()
			}
		})
	}
}
