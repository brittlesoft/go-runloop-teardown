package main

import (
	"context"

	"github.com/brittlesoft/go-runloop-teardown/internal/producing"
	"github.com/brittlesoft/go-runloop-teardown/internal/recording"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	recorder := recording.NewRecorder()
	producer := producing.NewProducer(recorder)

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error { return recorder.Run(egCtx) })
	eg.Go(func() error { return producer.RunCtxSelect(egCtx) })

	eg.Wait()
}
