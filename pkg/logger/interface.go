package logger

type ILogger interface {
	Info(args ...any)

	Infow(msg string, args ...any)

	Infof(msg string, args ...any)

	Debug(args ...any)

	Debugw(msg string, args ...any)

	Debugf(msg string, args ...any)

	Error(args ...any)

	Errorw(msg string, args ...any)

	Errorf(msg string, args ...any)

	Warn(args ...any)

	Warnw(msg string, args ...any)

	Warnf(msg string, args ...any)
}
