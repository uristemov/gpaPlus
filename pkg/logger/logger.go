package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func New() *zap.SugaredLogger {
	opt := zap.AddStacktrace(zap.FatalLevel)
	logger, err := zap.NewProduction(opt)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to product logger error: %v", err))
	}
	l := logger.Sugar()
	//l = l.With(zap.String("app", "repeatPRO"))
	return l
}
