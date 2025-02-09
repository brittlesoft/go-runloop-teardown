package recording

import (
	"context"
	"fmt"
	"log/slog"
)

type Recorder struct {
	inputCh chan struct{}
}

func NewRecorder() *Recorder {
	return &Recorder{inputCh: make(chan struct{})}
}

func (r *Recorder) SubmitCtxSelect(ctx context.Context, data struct{}) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("context error: %w", ctx.Err())
	case r.inputCh <- data:
		// all good
	}
	return nil
}

func (r *Recorder) SubmitCtx(ctx context.Context, data struct{}) error {
	if ctx.Err() != nil {
		return fmt.Errorf("context error: %w", ctx.Err())
	}
	r.inputCh <- data // FIXME: this will block forever if runloop isn't started or if stopped
	return nil
}

func (r *Recorder) Submit(data struct{}) error {
	r.inputCh <- data // FIXME: this will block forever if runloop isn't started or if stopped
	return nil
}

func (r *Recorder) Run(ctx context.Context) error {
	n := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-r.inputCh:
			n++
			slog.Debug("got data", "n", n)
		}
	}
}
