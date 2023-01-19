package log

import (
	"context"
	"net/http"

	"github.com/coinbase-samples/ib-api-go/model"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
)

func MakeContextLogger(l *logrus.Entry) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := uuid.New()
			ctx := ctxlogrus.ToContext(
				context.WithValue(r.Context(), model.RequestCtxKey, requestId.String()),
				l.WithField("requestId", requestId.String()),
			)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ToContext(ctx context.Context, entry *Entry) context.Context {
	return context.WithValue(ctx, LogCtxKey, entry)
}

func Extract(ctx context.Context) *Entry {
	l, ok := ctx.Value(LogCtxKey).(*Entry)
	if !ok || l == nil {
		return NewEntry()
	}
	return l
}

func TracefCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Tracef(format, args...)
}

func TraceCtx(ctx context.Context, args ...interface{}) {
	Extract(ctx).Trace(args...)
}

func DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Debugf(format, args...)
}

func DebugCtx(ctx context.Context, args ...interface{}) {
	Extract(ctx).Debug(args...)
}

func InfofCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Infof(format, args...)
}

func InfoCtx(ctx context.Context, args ...interface{}) {
	Extract(ctx).Info(args...)
}

func WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Warnf(format, args...)
}

func WarnCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Warn(args...)
}

func FatalCtx(ctx context.Context, args ...interface{}) {
	Extract(ctx).Fatal(args...)
}

func FatalfCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Fatalf(format, args...)
}

func PanicfCtx(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Panicf(format, args...)
}

func PanicCtx(ctx context.Context, args ...interface{}) {
	Extract(ctx).Panic(args...)
}
