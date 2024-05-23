package logger

import (
	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// error logger
var ErrorLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

const ErrorDir = "error/"

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func getWriter(filename string, expireDay int32, format string) io.Writer {
	if format == "" {
		format = "%Y%m%d%H"
	}
	hook, err := rotatelogs.New(
		filename+"."+format,
		rotatelogs.WithMaxAge(time.Duration(expireDay)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

// Setup initialize the log instance
func setup(filePath, fileName, logLevel, runMode string, expireDay int32) {
	fileFullName := filePath + fileName
	errFullName := filePath + ErrorDir + fileName + "_error"
	level := getLoggerLevel(logLevel)
	debugLevel := zap.DebugLevel

	// log slipt setting
	syncWriter := getWriter(fileFullName, expireDay, "")
	syncErrorWriter := getWriter(errFullName, expireDay, "")

	// log encode config
	var encoder zapcore.EncoderConfig = zap.NewProductionEncoderConfig()

	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	//真正的来配置zap
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.AddSync(syncWriter),
			zap.NewAtomicLevelAt(level),
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(debugLevel),
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.AddSync(syncErrorWriter),
			zap.NewAtomicLevelAt(zapcore.WarnLevel),
		),
	)
	var logger *zap.Logger
	additionalFields := zap.Fields(
		zap.Int("pid", os.Getpid()),
		zap.String("process", path.Base(os.Args[0])),
	)
	if runMode == "debug" {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development(), additionalFields)
	} else {
		logger = zap.New(core)
	}
	//logger := zap.New(core, zap.AddCaller(), zap.Development())
	ErrorLogger = logger.Sugar()
}

func Setup2(filePath, fileName, logLevel, runMode string, expireDay int32, format string) {
	fileFullName := filePath + fileName
	level := getLoggerLevel(logLevel)
	debugLevel := zap.DebugLevel

	syncWriter := getWriter(fileFullName, expireDay, format)

	var encoder zapcore.EncoderConfig = zap.NewProductionEncoderConfig()

	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.AddSync(syncWriter),
			zap.NewAtomicLevelAt(level),
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(debugLevel),
		),
	)
	var logger *zap.Logger
	additionalFields := zap.Fields(
		zap.Int("pid", os.Getpid()),
		zap.String("process", path.Base(os.Args[0])),
	)
	if runMode == "debug" {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development(), additionalFields)
	} else {
		logger = zap.New(core)
	}
	//logger := zap.New(core, zap.AddCaller(), zap.Development())
	ErrorLogger = logger.Sugar()
}

func SetupWithNoPrintln(filePath, fileName, logLevel, runMode string, expireDay int32) {
	fileFullName := filePath + fileName
	level := getLoggerLevel(logLevel)

	syncWriter := getWriter(fileFullName, expireDay, "")

	var encoder zapcore.EncoderConfig = zap.NewProductionEncoderConfig()

	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.AddSync(syncWriter),
		zap.NewAtomicLevelAt(level),
	))
	var logger *zap.Logger
	additionalFields := zap.Fields(
		zap.Int("pid", os.Getpid()),
		zap.String("process", path.Base(os.Args[0])),
	)
	if runMode == "debug" {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development(), additionalFields)
	} else {
		logger = zap.New(core)
	}
	//logger := zap.New(core, zap.AddCaller(), zap.Development())
	ErrorLogger = logger.Sugar()
}

func Debug(args ...any) {
	ErrorLogger.Debug(args...)
}

func Debugf(template string, args ...any) {
	ErrorLogger.Debugf(template, args...)
}

func Info(args ...any) {
	ErrorLogger.Info(args...)
}

func Infof(template string, args ...any) {
	ErrorLogger.Infof(template, args...)
}

func Warn(args ...any) {
	ErrorLogger.Warn(args...)
}

func Warnf(template string, args ...any) {
	ErrorLogger.Warnf(template, args...)
}

func Error(args ...any) {
	ErrorLogger.Error(args...)
}

func Errorf(template string, args ...any) {
	ErrorLogger.Errorf(template, args...)
}

// fatal不能使用，进程会直接退出
func Fatal(args ...any) {
	ErrorLogger.Error(args...)
}

func Fatalf(template string, args ...any) {
	ErrorLogger.Errorf(template, args...)
}
