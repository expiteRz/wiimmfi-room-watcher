package utils

import (
	"log/slog"
	"os"
	"runtime"
	"strings"
)

var logger *slog.Logger

func init() {
	hndr := slog.NewTextHandler(os.Stdout, nil)
	logger = slog.New(hndr)
}

func LogDebug(v ...any) {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	logger.Debug("["+funcName[lastSlash+1:lastDot]+"]", v)
}

func LogInfo(v ...any) {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	logger.Info("["+funcName[lastSlash+1:lastDot]+"]", v)
}

func LogFatal(v ...any) {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	logger.Error("["+funcName[lastSlash+1:lastDot]+"]", v)
	os.Exit(-1)
}

func LogInfoPrefix(prefix string, v ...any) {
	logger.Info(prefix, v)
}
