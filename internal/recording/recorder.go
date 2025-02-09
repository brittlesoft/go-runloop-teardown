package recording

import (
	"context"
	"log/slog"
)

type Recorder struct {
	inputCh chan struct{}
}

func NewRecorder() *Recorder {
	return &Recorder{inputCh: make(chan struct{})}
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
