package logger

import (
	"errors"
	"fmt"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.uber.org/zap"
)

// trace, request, exception
type Logger struct {
	log            *zap.SugaredLogger
	insightsClient appinsights.TelemetryClient
}

func NewLogger(client appinsights.TelemetryClient) (*Logger, error) {
	z, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &Logger{
		log:            z.WithOptions(zap.AddCallerSkip(1)).Sugar(),
		insightsClient: client,
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

func (l *Logger) Info(args ...any) {
	l.log.Info(args)

	l.insightsClient.TrackTrace(argsToString(args), contracts.Information)
}

func (l *Logger) Infow(msg string, args ...any) {
	l.log.Info(msg, args)

	msg = msg + " " + argsToString(args)
	l.insightsClient.TrackTrace(msg, contracts.Information)
}

func (l *Logger) Infof(msg string, args ...any) {
	l.log.Infof(msg, args)

	msg = fmt.Sprintf(msg, args)
	l.insightsClient.TrackTrace(msg, contracts.Information)
}

func (l *Logger) Debug(args ...any) {
	l.log.Debug(args)
}

func (l *Logger) Debugw(msg string, args ...any) {
	l.log.Debug(msg, args)
}

func (l *Logger) Debugf(msg string, args ...any) {
	l.log.Debugf(msg, args)
}

func (l *Logger) Error(args ...any) {
	l.log.Error(args)

	msg := argsToString(args)

	exception := appinsights.NewExceptionTelemetry(errors.New(msg))
	exception.Frames = exception.Frames[1:]
	l.insightsClient.Track(exception)
}

func (l *Logger) Errorw(msg string, args ...any) {
	l.log.Error(msg, args)

	msg = msg + " " + argsToString(args)
	exception := appinsights.NewExceptionTelemetry(errors.New(msg))
	exception.Frames = exception.Frames[1:]
	l.insightsClient.Track(exception)
}

func (l *Logger) Errorf(msg string, args ...any) {
	l.log.Errorf(msg, args)

	msg = fmt.Sprintf(msg, args)
	exception := appinsights.NewExceptionTelemetry(errors.New(msg))
	exception.Frames = exception.Frames[1:]
	l.insightsClient.Track(exception)
}

func (l *Logger) Warn(args ...any) {
	l.log.Warn(args)
}

func (l *Logger) Warnw(msg string, args ...any) {
	l.log.Warn(msg, args)
}

func (l *Logger) Warnf(msg string, args ...any) {
	l.log.Warnf(msg, args)
}
