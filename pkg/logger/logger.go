package logger

import (
	"os"

	internal "github.com/Qwepo/InCryipt/internal"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

func NewLogger(conf *internal.Config) *zerolog.Logger {
	level := setLevel(conf)
	writer := logWriter(conf)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(writer).Level(zerolog.Level(level)).With().Timestamp().Logger()

	return &logger

}

func logWriter(conf *internal.Config) zerolog.LevelWriter {
	l := lumberjack.Logger{
		Filename:   conf.Logger.Filename,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	writer := zerolog.MultiLevelWriter(os.Stdout, &l)
	return writer

}

func setLevel(conf *internal.Config) int8 {
	switch conf.Logger.Level {
	case "debug":
		return 0
	case "info":
		return 1
	case "warn":
		return 2
	case "error":
		return 3
	case "fatal":
		return 4
	case "panic":
		return 5
	case "nolevel":
		return 6
	case "disable":
		return 6
	default:
		return 3
	}

}
