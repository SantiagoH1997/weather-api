package logger

import (
	"go.uber.org/zap"
)

// Log is the logger for the entire app
var Log *zap.SugaredLogger

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	Log = l.Sugar()
}
