package logger

type ILogger interface {
	Info(format string, v ...any)
	Error(format string, v ...any)
	Warning(format string, v ...any)
	Fatal(format string, v ...any)
}
