package frame

import (
	"fmt"
	"runtime"

	"github.com/gotomicro/ego/core/util/xcolor"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	DefaultLogger, _ = NewDebugLogger().Build(zap.AddCallerSkip(1))
	grpcLogger, _    = NewDebugLogger().Build(zap.AddCallerSkip(5))
)

func NewDebugLogger() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    defaultDebugConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func defaultDebugConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:    "ts",
		LevelKey:   "lv",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		//StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    debugEncodeLevel,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		//EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
}

// debugEncodeLevel ...
func debugEncodeLevel(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var colorize = xcolor.Red
	switch lv {
	case zapcore.DebugLevel:
		colorize = xcolor.Blue
	case zapcore.InfoLevel:
		colorize = xcolor.Green
	case zapcore.WarnLevel:
		colorize = xcolor.Yellow
	case zapcore.ErrorLevel, zap.PanicLevel, zap.DPanicLevel, zap.FatalLevel:
		colorize = xcolor.Red
	default:
	}
	enc.AppendString(colorize(lv.CapitalString()))
}

func LoggerPanic(msg string, fields ...zap.Field) {
	enc := zapcore.NewMapObjectEncoder()
	for _, field := range fields {
		field.AddTo(enc)
	}
	fmt.Printf("%s: \n    %s: %s\n", xcolor.Red("panic"), xcolor.Red("msg"), msg)
	if _, file, line, ok := runtime.Caller(3); ok {
		fmt.Printf("    %s: %s:%d\n", xcolor.Red("loc"), file, line)
	}
	for key, val := range enc.Fields {
		fmt.Printf("    %s: %s\n", xcolor.Red(key), fmt.Sprintf("%+v", val))
	}
	msg = fmt.Sprintf("%-32s", msg)
	panic(msg)
}

func LoggerError(msg string, fields ...zap.Field) {
	normalizeMessage(msg)
	DefaultLogger.Error(msg, fields...)
}

func normalizeMessage(msg string) string {
	return fmt.Sprintf("%-32s", msg)
}
