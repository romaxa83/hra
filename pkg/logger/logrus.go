package logger

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func NewLoggerConfig(logLevel string, devMode bool, encoder string) *Config {
	return &Config{
		LogLevel: logLevel,
		DevMode:  devMode,
		Encoder:  encoder,
	}
}

// Application logger
type appLogger struct {
	level    string
	devMode  bool
	encoding string
	logger   *logrus.Logger
}

// NewAppLogger App Logger constructor
func NewAppLogger(cfg *Config) *appLogger {
	return &appLogger{
		level:    cfg.LogLevel,
		devMode:  cfg.DevMode,
		encoding: cfg.Encoder,
	}
}

var loggerLevelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"trace": logrus.TraceLevel,
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
}

func (l *appLogger) getLoggerLevel() logrus.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return logrus.DebugLevel
	}

	return level
}

func (l *appLogger) InitLogger() {
	//logLevel := l.getLoggerLevel()

	logger := logrus.New()

	//t := logrus.TextFormatter{}

	logger.SetFormatter(&logrus.JSONFormatter{})

	l.logger = logger
}

func (l *appLogger) Debug(msg ...interface{}) {
	l.logger.Debug(msg...)
}

func (l *appLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *appLogger) Info(msg ...interface{}) {
	l.logger.Info(msg...)
}

func (l *appLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *appLogger) Warn(msg ...interface{}) {
	l.logger.Warn(msg...)
}

func (l *appLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *appLogger) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

func (l *appLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
