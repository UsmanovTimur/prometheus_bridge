package logging

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

var Logger = &logrus.Logger{}

// Инициация логирования проекта
func Init() {
	l := logrus.New()
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN != "" {
		hook, err := logrus_sentry.NewSentryHook(sentryDSN,
			[]logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
			})
		if err == nil {
			hook.StacktraceConfiguration.Enable = true
			l.Hooks.Add(hook)
		}
	}
	//Настройка логов
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	l.SetOutput(os.Stdout)
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	l.SetOutput(os.Stdout)
	Logger = l
}
