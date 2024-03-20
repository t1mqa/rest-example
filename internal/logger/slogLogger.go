package logger

import (
	"errors"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

// SlogLogger представляет реализацию интерфейса Logger
type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(envType string) *SlogLogger {
	l := &SlogLogger{}
	l.MustSetupLogger(envType)
	return l
}

func (l *SlogLogger) MustSetupLogger(envType string) {
	var handler slog.Handler

	switch envType {
	case envProd:
		// В случае продакшна - храним в JSON файле
		_, err := os.Stat("logs")
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				panic("logs folder doesnt exists")
			}
			panic(err)
		}

		file, err := os.OpenFile("logs/logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		handler = slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo})

	case envLocal:
		// В случае локала - принтимся в консоль
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})

	default:
		panic("something wrong with env config settings")
	}

	l.logger = slog.New(handler)
}

func (l *SlogLogger) Info(msg string, keysAndValues ...any) {
	l.logger.Info(msg, keysAndValues...)
	// слог вполне себе проверит, сколько их там, но выведет сообщение с BADKEY
	// поэтому хочу оставить проверку во всех случаях

	l.checkUnevenKeyError(len(keysAndValues))
}

func (l *SlogLogger) Warn(msg string, keysAndValues ...any) {
	l.logger.Warn(msg, keysAndValues...)

	l.checkUnevenKeyError(len(keysAndValues))
}

func (l *SlogLogger) Error(msg string, keysAndValues ...any) {
	l.logger.Error(msg, keysAndValues...)

	l.checkUnevenKeyError(len(keysAndValues))
}

// Fatal паникует после Error лога
func (l *SlogLogger) Fatal(msg string, keysAndValues ...any) {
	l.logger.Error(msg, keysAndValues...)

	l.checkUnevenKeyError(len(keysAndValues))

	panic(msg)
}

func (l *SlogLogger) checkUnevenKeyError(keysLength int) {
	if keysLength%2 != 0 {
		l.logger.Error("logging call with uneven key/value pairs")
	}
}
