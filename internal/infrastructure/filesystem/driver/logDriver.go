package driver

import (
	"fmt"
	"strings"
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sharedInstance *logDriver

type logDriver struct {
	*zap.Logger
}

func SharedLogger() *logDriver {
	if sharedInstance == nil {
		// まだ作っていな場合は、デフォルトのロガーを引き渡す。
		logger, _ := zap.NewDevelopment()
		return &logDriver{Logger: logger}
	}
	return sharedInstance
}

func NewLogDriver(level string, destination string, lifecycle int) *logDriver{
	if sharedInstance == nil {
		loglevel := zapcore.InfoLevel
		// levelcheck
		switch strings.ToLower(level) {
		case "debug":
			loglevel = zapcore.DebugLevel
		case "warn":
			loglevel = zapcore.WarnLevel
		case "error":
			loglevel = zapcore.ErrorLevel
		case "dpanic":
			loglevel = zapcore.DPanicLevel
		case "panic":
			loglevel = zapcore.PanicLevel
		case "fatal":
			loglevel = zapcore.FatalLevel
		}

		// ログ出力
		zaplevel := zap.NewAtomicLevel()
		zaplevel.SetLevel(loglevel) // ここでログレベルを制御する。

		config := zap.NewProductionConfig()
		config.Level = zaplevel

		encodeConfig := zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     JSTTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		
		// ログローテーション指定
		sink := zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   destination,
				MaxSize:	30, // megabytes
				MaxAge:     lifecycle, //days
			},
		)

		logger := zap.New(
			zapcore.NewCore(zapcore.NewJSONEncoder(encodeConfig), sink, config.Level),
			zap.AddCaller(),      // ファイルの位置情報を取得するために必要なおまじない。
			zap.AddCallerSkip(1), // wrapしているため、caller位置がこのモジュールとなってしまうため、一つ上にする。
		)
		zap.ReplaceGlobals(logger)	// グローバルにロガーを引き渡し
		sharedInstance = &logDriver{Logger: logger}
	}
	return sharedInstance
}

func JSTTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    const layout = "2006-01-02 15:04:05.000"
    jst := time.FixedZone("Asia/Tokyo", 9*60*60)
    enc.AppendString(t.In(jst).Format(layout))
}

func createFields(fieldmap map[string]interface{}) []zapcore.Field{
	fields := []zapcore.Field{}
	for key, val := range fieldmap {
		fields = append(fields, zap.String(key, fmt.Sprintf("%v", val)))
	}
	return fields
}

func (d *logDriver) Info(message string, fieldmap map[string]interface{}){
	if d.Logger != nil {
		fields := createFields(fieldmap)
		d.Logger.Info(message, fields...)
	}
}

func (d *logDriver) Error(message string, fieldmap map[string]interface{}){
	if d.Logger != nil {
		fields := createFields(fieldmap)
		d.Logger.Error(message, fields...)
	}
}

func (d *logDriver) Debug(message string, fieldmap map[string]interface{}){
	if d.Logger != nil {
		fields := createFields(fieldmap)
		d.Logger.Debug(message, fields...)
	}
}

func (d *logDriver) Warn(message string, fieldmap map[string]interface{}){
	if d.Logger != nil {
		fields := createFields(fieldmap)
		d.Logger.Warn(message, fields...)
	}
}

func (d *logDriver) Panic(message string, fieldmap map[string]interface{}){
	if d.Logger != nil {
		fields := createFields(fieldmap)
		d.Logger.DPanic(message, fields...)
	}
}

func (d *logDriver) Fatal(message string, fieldmap map[string]interface{}){
	if d.Logger != nil {
		fields := createFields(fieldmap)
		d.Logger.Fatal(message, fields...)
	}
}

func (d *logDriver) Sync(){
	if d.Logger != nil {
		d.Logger.Sync()
	}
}
