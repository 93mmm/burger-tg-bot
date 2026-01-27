package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

// Инитим логгер при первом импорте пакета
func init() {
	SetLogger(NewLogger(defaultLevel))
}

// NewLogger Создание нового логгера, пишущего только в stdout
func NewLogger(level zapcore.LevelEnabler, options ...zap.Option) *zap.SugaredLogger {
	return newLogger(level, os.Stdout, options...)
}

// NewLogger Конструктор нового присахаренного zap логгера с дополнительными полями и отконфигурированного под наши нужды.
// Так же принимает дополнительные опции вида zap.Option
// При ините пакета (то бишь первом импорте) или при пустом переданном аргументе writer создастся глобальный логгер только для stdout.
func newLogger(level zapcore.LevelEnabler, writer io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if writer == nil {
		writer = os.Stdout
	}
	loggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(writer),
		level,
	)

	return zap.New(loggerCore, options...).Sugar()
}

// SetLogger устанавливает глобальный логгер
func SetLogger(logger *zap.SugaredLogger) {
	globalLogger = logger
}

// Debug Лог с уровнем Debug
func Debug(ctx context.Context, args ...any) {
	logger := getLogger(ctx)
	logger.Debug(args...)
}

// Debugf Лог с уровнем Debug и форматированием сообщения
func Debugf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Debugf(format, args...)
}

// DebugKV Лог с уровнем Debug и дополнительными ключами и значениями в сообщении
func DebugKV(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Debugw(message, keyValues...)
}

func Info(ctx context.Context, args ...any) {
	logger := getLogger(ctx)

	logger.Info(args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Infof(format, args...)
}

func InfoKV(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Infow(message, keyValues...)
}

func Warn(ctx context.Context, args ...any) {
	logger := getLogger(ctx)
	logger.Warn(args...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Warnf(format, args...)
}

func WarnKV(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Warnw(message, keyValues...)
}

func Error(ctx context.Context, args ...any) {
	logger := getLogger(ctx)
	logger.Error(args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Errorf(format, args...)
}

func ErrorKV(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Errorw(message, keyValues...)
}

func Fatal(ctx context.Context, args ...any) {
	logger := getLogger(ctx)
	logger.Fatal(args...)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Fatalf(format, args...)
}

func FatalKV(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Fatalw(message, keyValues...)
}

func Panic(ctx context.Context, args ...any) {
	logger := getLogger(ctx)
	logger.Panic(args...)
}

func Panicf(ctx context.Context, format string, args ...any) {
	logger := getLogger(ctx)
	logger.Panicf(format, args...)
}

func PanicKV(ctx context.Context, message string, keyValues ...any) {
	logger := getLogger(ctx)
	logger.Panicw(message, keyValues...)
}
