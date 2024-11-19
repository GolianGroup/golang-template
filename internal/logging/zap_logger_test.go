package logging

import (
	"bytes"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewZapLogger(t *testing.T) {
	logger := zap.NewExample()

	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want *ZapLogger
	}{
		{
			name: "Valid logger",
			args: args{
				logger: logger,
			},
			want: &ZapLogger{
				Logger: logger,
			},
		},
		{
			name: "Nil logger",
			args: args{
				logger: nil,
			},
			want: &ZapLogger{
				Logger: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewZapLogger(tt.args.logger)
			if (got == nil) != (tt.want == nil) || (got != nil && tt.want != nil && got.Logger != tt.want.Logger) {
				t.Errorf("NewZapLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZapLoggerInfo(t *testing.T) {
	// create custome example logger
	buffer := &bytes.Buffer{}
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.AddSync(buffer),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	zapLogger := NewZapLogger(logger)

	tests := []struct {
		name   string
		logger *ZapLogger
		msg    string
		fields []zap.Field
		want   string
	}{
		{
			name:   "Simple Info log",
			logger: zapLogger,
			msg:    "Info log message",
			fields: []zap.Field{zap.String("key", "value")},
			want:   `{"L":"INFO","T":"2024-11-19T13:11:53.385+0330","M":"Info log message","key":"value"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			tt.logger.Info(tt.msg, tt.fields...)
			// if got := buffer.String(); !containsJSON(got, tt.want) {
			// 	t.Errorf("Info() = %s, want contains %s", got, tt.want)
			// }
		})
	}
}

// func TestZapLogger_Error(t *testing.T) {
// 	type args struct {
// 		msg    string
// 		fields []zap.Field
// 	}
// 	tests := []struct {
// 		name string
// 		z    *ZapLogger
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.z.Error(tt.args.msg, tt.args.fields...)
// 		})
// 	}
// }

//	func TestZapLogger_Fatal(t *testing.T) {
//		type args struct {
//			msg    string
//			fields []zap.Field
//		}
//		tests := []struct {
//			name string
//			z    *ZapLogger
//			args args
//		}{
//			// TODO: Add test cases.
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				tt.z.Fatal(tt.args.msg, tt.args.fields...)
//			})
//		}
//	}
func containsJSON(actual, expected string) bool {
	return bytes.Contains([]byte(actual), []byte(expected))
}
