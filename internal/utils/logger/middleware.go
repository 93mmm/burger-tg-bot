package logger

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PanicRecoverUnaryInterceptor Мидлварь для восстановления от паник и логирования их ошибок
func PanicRecoverUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any,
	err error) {
	defer func() {
		if r := recover(); r != nil {
			ErrorKV(ctx, fmt.Sprintf("произошла паника во время запроса: %+#v", r),
				"call_method", info.FullMethod,
				"stacktrace", string(debug.Stack()),
			)
			err = status.Errorf(codes.Internal, "восстановлено от паники: %v", r)
		}
	}()
	return handler(ctx, req)
}

// LoggerUnaryInterceptor Мидлварь для IN/OUT логирования
func LoggerUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	ctx = WithKV(ctx, "request_id", uuid.New().String())
	InfoKV(ctx, "gRPC запрос",
		"call_method", info.FullMethod,
		"request_body", req,
	)

	response, err := handler(ctx, req)
	if err != nil {
		ErrorKV(ctx, "произошла ошибка во время обработки gRPC запроса",
			"call_method", info.FullMethod,
			"response_body", response,
			"error", err,
		)
		return response, err
	}
	InfoKV(ctx, "gRPC ответ",
		"call_method", info.FullMethod,
		"response_body", response,
	)
	return response, err
}
