package logger

import (
	"context"
	"errors"
	"fmt"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.uber.org/zap"
)

// trace, request, exception
type Logger struct {
	Log       *zap.SugaredLogger
	Telemetry appinsights.TelemetryClient
	ReqID     string
}

func NewLogger(client appinsights.TelemetryClient) (*Logger, error) {
	z, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &Logger{
		Log:       z.WithOptions(zap.AddCallerSkip(1)).Sugar(),
		Telemetry: client,
	}, nil
}

func Must(l *Logger, err error) *Logger {
	if err != nil {
		panic(err)
	}
	return l
}

func argsToString(args ...any) string {
	var msg = ""
	for _, v := range args {
		msg = msg + fmt.Sprintf("%v ", v)
	}
	return msg
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	if reqID := ctx.Value("reqID"); reqID != nil {
		return &Logger{
			l.Log.With(zap.String("reqID", reqID.(string))),
			l.Telemetry,
			reqID.(string),
		}
	}
	return l
}

func (l *Logger) Info(args ...any) {
	l.Log.Info(args)

	l.Telemetry.TrackTrace(argsToString(args), contracts.Information)
}

func (l *Logger) InfoCtx(ctx context.Context, args ...any) {
	l = l.WithContext(ctx)
	l.Log.Info(args)

	telemetry := appinsights.NewTraceTelemetry(argsToString(args), contracts.Information)

	telemetry.Properties["reqID"] = l.ReqID
	l.Telemetry.Track(telemetry)
}

func (l *Logger) Infow(msg string, args ...any) {
	l.Log.Info(msg, args)

	msg = msg + " " + argsToString(args)
	l.Telemetry.TrackTrace(msg, contracts.Information)
}

func (l *Logger) Infof(msg string, args ...any) {
	l.Log.Infof(msg, args)

	msg = fmt.Sprintf(msg, args)
	l.Telemetry.TrackTrace(msg, contracts.Information)
}

func (l *Logger) DebugCtx(ctx context.Context, args ...any) {
	l.WithContext(ctx).Debug(args)
}

func (l *Logger) Debug(args ...any) {
	l.Log.Debug(args)
}

func (l *Logger) Debugw(msg string, args ...any) {
	l.Log.Debug(msg, args)
}

func (l *Logger) Debugf(msg string, args ...any) {
	l.Log.Debugf(msg, args)
}

func (l *Logger) Error(args ...any) {
	l.Log.Error(args)

	msg := argsToString(args)

	exception := appinsights.NewExceptionTelemetry(errors.New(msg))
	exception.Frames = exception.Frames[1:]
	l.Telemetry.Track(exception)
}

func (l *Logger) Errorw(msg string, args ...any) {
	l.Log.Error(msg, args)

	msg = msg + " " + argsToString(args)
	exception := appinsights.NewExceptionTelemetry(errors.New(msg))
	exception.Frames = exception.Frames[1:]
	l.Telemetry.Track(exception)
}

func (l *Logger) Errorf(msg string, args ...any) {
	l.Log.Errorf(msg, args)

	msg = fmt.Sprintf(msg, args)
	exception := appinsights.NewExceptionTelemetry(errors.New(msg))
	exception.Frames = exception.Frames[1:]
	l.Telemetry.Track(exception)
}

func (l *Logger) Warn(args ...any) {
	l.Log.Warn(args)
}

func (l *Logger) Warnw(msg string, args ...any) {
	l.Log.Warn(msg, args)
}

func (l *Logger) Warnf(msg string, args ...any) {
	l.Log.Warnf(msg, args)
}
