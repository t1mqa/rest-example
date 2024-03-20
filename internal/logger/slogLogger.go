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

func NewSlogLogger() *SlogLogger {
	// При создании логгера мы не назначаем ему конкретный обработчик.
	// В дальнейшем обязательно вызвать MustSetupLogger
	return &SlogLogger{}
}

func (l *SlogLogger) MustSetupLogger(env string) {
	var handler slog.Handler

	switch env {
	case envProd:
		// В случае продакшна - храним в JSON файле

		_, err := os.Stat("logs")
		if errors.Is(err, os.ErrNotExist) {
			panic("logs folder doesnt exists")
		} else if err != nil {
			panic(err)
		}

		file, err := os.OpenFile("logs/logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if errors.Is(err, os.ErrNotExist) {
			panic("logs.json file doesnt exists")
		} else if err != nil {
			panic(err)
		}

		handler = slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo})

	case envLocal:
		// В случае локала - принтимся в консоль

		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	default:
		//

		panic("something wrong with env config settings")
	}

	l.logger = slog.New(handler)
}

func (l *SlogLogger) Info(msg string, keysAndValues ...any) {
	if len(keysAndValues)%2 == 0 {
		l.logger.Info(msg, keysAndValues...)
	} else {
		l.logger.Error("logging call with uneven key/value pairs")
	}
}

func (l *SlogLogger) Warn(msg string, keysAndValues ...any) {
	if len(keysAndValues)%2 == 0 {
		l.logger.Warn(msg, keysAndValues...)
	} else {
		l.logger.Error("logging call with uneven key/value pairs")
	}
}

func (l *SlogLogger) Error(msg string, keysAndValues ...any) {
	if len(keysAndValues)%2 == 0 {
		l.logger.Error(msg, keysAndValues...)
	} else {
		l.logger.Error("logging call with uneven key/value pairs")
	}
}

// Fatal паникует после Error
func (l *SlogLogger) Fatal(msg string, keysAndValues ...any) {
	if len(keysAndValues)%2 == 0 {
		l.logger.Error(msg, keysAndValues...)
		panic(msg)
	} else {
		l.logger.Error("logging call with uneven key/value pairs")
		panic(msg) // anyway runtime drop
	}
}
