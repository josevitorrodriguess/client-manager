package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT     = "LOG_OUTPUT"
	LOG_LEVEL      = "LOG_LEVEL"
	SERVICE_NAME   = "SERVICE_NAME"
	DEFAULT_FIELDS = []zap.Field{
		zap.String("service", getServiceName()),
	}

	// Cores ANSI
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func init() {
	output := getOutputLogs()
	isStdout := output == "stdout"

	var encoder zapcore.Encoder
	if isStdout {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	} else {
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(getLevelLogs()),
	)

	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Info(message string, tags ...zap.Field) {
	tags = append(DEFAULT_FIELDS, tags...)
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(DEFAULT_FIELDS, tags...)
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	log.Sync()
}

func Warn(message string, tags ...zap.Field) {
	tags = append(DEFAULT_FIELDS, tags...)
	log.Warn(message, tags...)
	log.Sync()
}

func Debug(message string, tags ...zap.Field) {
	tags = append(DEFAULT_FIELDS, tags...)
	log.Debug(message, tags...)
	log.Sync()
}

func WithRequestID(requestID string) []zap.Field {
	return []zap.Field{
		zap.String("request_id", requestID),
	}
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.InfoLevel
	}
}

func getServiceName() string {
	serviceName := strings.TrimSpace(os.Getenv(SERVICE_NAME))
	if serviceName == "" {
		return "client-manager"
	}
	return serviceName
}
