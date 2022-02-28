package logger

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func getZapLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "DebugLevel":
		return zapcore.DebugLevel
	case "InfoLevel":
		return zapcore.InfoLevel
	case "WarnLevel":
		return zapcore.WarnLevel
	case "ErrorLevel":
		return zapcore.ErrorLevel
	case "DPanicLevel":
		return zapcore.DPanicLevel
	case "PanicLevel":
		return zapcore.PanicLevel
	default:
		return zapcore.DebugLevel
	}
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		panic(err)
	}

	return encCfg
}

func getRotateFile() *rotatelogs.RotateLogs {
	logf, err := rotatelogs.New(
		config.SglConfig.GetLogFilename(),
		rotatelogs.WithLinkName(config.SglConfig.GetLogFilenameLink()),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic("failed to create rotatelogs: %s")
	}
	return logf
}

// InitializeZap ...
func InitializeZap() *zap.SugaredLogger {
	cores := []zapcore.Core{}

	if config.SglConfig.IsUseConsolelog() == "y" {
		consoleDebuggingWriteSyncer := zapcore.Lock(os.Stdout)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
		core := zapcore.NewCore(consoleEncoder, consoleDebuggingWriteSyncer, getZapLogLevel(config.SglConfig.GetLogLevel()))
		cores = append(cores, core)
	}

	// if config.SglConfig.IsUseConsolelog() == "y" {
	// 	consoleErrorsWriteSyncer := zapcore.Lock(os.Stderr)
	// 	consoleEncoder := zapcore.NewConsoleEncoder(getEncoderConfig())
	// 	core := zapcore.NewCore(consoleEncoder, consoleErrorsWriteSyncer, getZapLogLevel(config.SglConfig.GetLogLevel()))
	// 	cores = append(cores, core)
	// }

	if config.SglConfig.IsUseFilelog() == "y" {
		fileWriteSyncer := zapcore.AddSync(getRotateFile())
		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		core := zapcore.NewCore(fileEncoder, fileWriteSyncer, getZapLogLevel(config.SglConfig.GetLogLevel()))
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(
		cores...,
	)

	zapOpts := []zap.Option{zap.AddCaller()}
	Logger = zap.New(combinedCore, zapOpts...).Sugar()

	return Logger
}
