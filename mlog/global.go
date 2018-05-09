// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package mlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *Logger

func InitGlobalLogger(logger *Logger) {
	globalLogger = logger
	Debug = globalLogger.Debug
	Info = globalLogger.Info
	Warn = globalLogger.Warn
	Error = globalLogger.Error
	Critical = globalLogger.Critical
}

func RedirectStdLog(logger *Logger) {
	zap.RedirectStdLogAt(logger.zap.With(zap.String("source", "stdlog")), zapcore.ErrorLevel)
}

type LogFunc func(string, ...Field)

var Debug LogFunc
var Info LogFunc
var Warn LogFunc
var Error LogFunc
var Critical LogFunc
