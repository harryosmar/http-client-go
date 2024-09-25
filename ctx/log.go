package ctx

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type (
	LoggerCtxKey struct {
	}

	LoggerCtx interface {
		GetLoggerFromContext(ctx context.Context) *log.Entry
		ContextWithLogger(ctx context.Context, logEntry *log.Entry) context.Context
	}

	loggerCtx struct {
	}
)

func NewLoggerCtx() *loggerCtx {
	return &loggerCtx{}
}

func (l loggerCtx) GetLoggerFromContext(ctx context.Context) *log.Entry {
	entry := ctx.Value(LoggerCtxKey{})
	if logEntry, ok := entry.(*log.Entry); ok {
		return logEntry
	}

	return log.WithContext(ctx)
}

func (l loggerCtx) ContextWithLogger(ctx context.Context, logEntry *log.Entry) context.Context {
	return context.WithValue(ctx, LoggerCtxKey{}, logEntry)
}
