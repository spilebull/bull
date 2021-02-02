/*
	Package logger では
	独自ロガーの定義
	ログ出力の構造化
	ログ出力先の指定
	などを行います
*/
package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Logger はSugaredLoggerよりも高速かつ低アロケーションで動作しますが、構造化スタイルのロギングしかできません（printやprintfスタイルのロギングはできません）
var Logger *zap.Logger

//SugarLogger はLoggerから取得できる簡易なLogger
//構造化ロギングとprintfスタイルのロギングどちらも利用可能ですが、Loggerよりも低速で高アロケーションです
var SugarLogger *zap.SugaredLogger

// InitLogger Logger/SugarLogger初期化
func InitLogger() {
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr))
	encoder := getEncoder()
	var enab zapcore.LevelEnabler
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "Debug":
		enab = zapcore.DebugLevel
	case "Info":
		enab = zapcore.InfoLevel
	case "Warn":
		enab = zapcore.WarnLevel
	case "Error":
		enab = zapcore.ErrorLevel
	case "DPanic":
		enab = zapcore.DPanicLevel
	case "Panic":
		enab = zapcore.PanicLevel
	case "Fatal":
		enab = zapcore.FatalLevel
	default:
		enab = zapcore.DebugLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, enab)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(3))
	SugarLogger = Logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "severity"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig) //Console Encoder-->NewJSONEncoder
}

func Debug(args ...interface{}) {
	SugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	SugarLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	SugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	SugarLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	SugarLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	SugarLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	SugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	SugarLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	SugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	SugarLogger.Fatalf(template, args...)
}

func LoggerSync() error {
	return Logger.Sync()
}
