package logging

import "go.uber.org/zap"

var logger, _ = zap.NewProduction()
var sugar = logger.Sugar()

func Error(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}
