package logger

import (
	"context"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func MessageProducer() grpc_zap.MessageProducer {
	const ReqKey = "request"
	logger := globalLogger.Desugar()

	return func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
		errStatus, ok := status.FromError(err)
		if ok {
			msg = errStatus.Message()
		}
		fields := []zap.Field{
			zap.Uint32("code", uint32(code)),
			methodFieldFromContext(ctx),
			RequestIdField(ctx),
			zap.Any(ReqKey, ctx.Value(ReqKey)),
			zap.Error(err),
			duration,
		}

		logLevel := globalLogger.Level()

		switch {
		case level == zapcore.DebugLevel && logLevel < zapcore.InfoLevel:
			logger.Debug(msg, fields...)
		case level == zapcore.InfoLevel && logLevel <= zapcore.InfoLevel:
			logger.Info(msg, fields...)
		case level == zapcore.WarnLevel && logLevel <= zapcore.WarnLevel:
			logger.Warn(msg, fields...)
		case level == zapcore.DPanicLevel && logLevel <= zapcore.ErrorLevel:
			fallthrough
		case level == zapcore.PanicLevel && logLevel <= zapcore.ErrorLevel:
			fallthrough
		case level == zapcore.FatalLevel && logLevel <= zapcore.ErrorLevel:
			fallthrough
		case level == zapcore.ErrorLevel && logLevel <= zapcore.ErrorLevel:
			logger.Error(msg, fields...)
		}
	}
}

func methodFieldFromContext(ctx context.Context) zap.Field {
	method, _ := grpc.Method(ctx)
	return zap.String("method", method)
}

func RequestIdField(ctx context.Context) zap.Field {
	var requestId string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		requestIds := md.Get("x-request-id")
		if len(requestIds) > 0 {
			requestId = requestIds[len(requestIds)-1]
		}
	}
	return zap.String("request-id", requestId)
}
