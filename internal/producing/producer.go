package producing

import (
	"context"
	"log/slog"

	"github.com/brittlesoft/go-runloop-teardown/internal/recording"
)

type Producer struct {
	recorder *recording.Recorder
}

func NewProducer(rec *recording.Recorder) *Producer {
	return &Producer{recorder: rec}
}

func (p *Producer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			slog.Info("producer done")
			return ctx.Err()
		default:
			p.recorder.Submit(struct{}{})
		}
	}
}

func (p *Producer) RunCtx(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			slog.Info("producer done")
			return ctx.Err()
		default:
			p.recorder.SubmitCtx(ctx, struct{}{})
		}
	}
}

func (p *Producer) RunCtxSelect(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			slog.Info("producer done")
			return ctx.Err()
		default:
			p.recorder.SubmitCtxSelect(ctx, struct{}{})
		}
	}
}
