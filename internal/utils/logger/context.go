package logger

import (
	"context"

	"go.uber.org/zap"
)

// contextKey Объявляем кастомный тип для ключа контекста во избежание коллизий
type contextKey int

const (
	// loggerContextKey Устанавливаем единый ключ для логгера,через который и будем его доставать
	loggerContextKey contextKey = 0
)

// ToContext Передает в родительский контекст логгер.
func ToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// FromContext достает логгер из контекста. Если в контексте логгер не
// обнаруживается - возвращает глобальный логгер.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := getLogger(ctx)

	return l
}

// WithName создает именованный логгер из уже имеющегося в контексте.
// Дочерние логгеры наследуют имя.
func WithName(ctx context.Context, name string) context.Context {
	log := FromContext(ctx).Named(name)
	return ToContext(ctx, log)
}

// WithKV создает логгер из уже имеющегося в контексте и устанавливает метаданные.
// Принимает ключ и значение, которые будут наследоваться дочерними логгерами.
func WithKV(ctx context.Context, key string, value any) context.Context {
	log := FromContext(ctx).With(key, value)
	return ToContext(ctx, log)
}

// WithFields создает логгер из уже имеющегося в контексте и устанавливает метаданные,
// используя типизированные поля.
func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	log := FromContext(ctx).
		Desugar().
		With(fields...).
		Sugar()
	return ToContext(ctx, log)
}

// AddKV добавляет в контекст пары ключей и значений, которые
// используются при вызове функций логгера в качестве полей.
func AddKV(ctx context.Context, kvs ...any) context.Context {
	if len(kvs) == 0 {
		return ctx
	}

	kvsFromContext := getKvsFromContext(ctx)
	additionalFields := sweetenFields(ctx, kvs)
	merged := mergeFields(kvsFromContext, additionalFields)

	return context.WithValue(ctx, logFieldKey, merged)
}

// Получение логгера либо из глобального инстанса, либо из контекста, вне зависимости от того,
// подходящий ли сейчас вызван уровень логгирования
func getLogger(ctx context.Context) *zap.SugaredLogger {
	logger := globalLogger
	if contextLogger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		logger = contextLogger
	}
	return logger
}

type logFieldsKeyType string

var logFieldKey = logFieldsKeyType("logger-fields")

// Получение актуальных пар ключ-значение для логгера из контекста
func getKvsFromContext(ctx context.Context) []any {
	if kvs, ok := ctx.Value(logFieldKey).([]any); ok {
		// Пользователь ранее добавил ключи в контекст
		// с помощью вызова logger.AddKV().
		kvsCopy := make([]any, len(kvs))
		copy(kvsCopy, kvs)
		return kvsCopy
	}
	return nil
}
