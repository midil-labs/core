package logger

import (
    "io"
    "os"
    "runtime/debug"
    "strconv"
    "sync"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/pkgerrors"
    "gopkg.in/natefinch/lumberjack.v2"
    "github.com/midil-labs/core/pkg/config"
    "github.com/newrelic/go-agent/v3/newrelic"
    "github.com/newrelic/go-agent/v3/integrations/logcontext-v2/zerologWriter"
)

var once sync.Once

var log zerolog.Logger

func Get(config *config.LoggingConfig) zerolog.Logger {
    once.Do(func() {
        zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
        zerolog.TimeFieldFormat = time.RFC3339Nano

        logLevel, err := strconv.Atoi(config.Level)
        if err != nil {
            logLevel = int(zerolog.InfoLevel) // default to INFO
        }

        var output io.Writer = zerolog.ConsoleWriter{
            Out:        os.Stdout,
            TimeFormat: time.RFC3339,
        }

        if config.Level != "development" {
            fileLogger := &lumberjack.Logger{
                Filename:   config.OutputPath,
                MaxSize:    5, //
                MaxBackups: 10,
                MaxAge:     14,
                Compress:   true,
            }
            nrWriter := zerologWriter.New(os.Stdout, newRelicApp)
            output = zerolog.MultiLevelWriter(os.Stderr, fileLogger,)
        }

        var gitRevision string

        buildInfo, ok := debug.ReadBuildInfo()
        if ok {
            for _, v := range buildInfo.Settings {
                if v.Key == "vcs.revision" {
                    gitRevision = v.Value
                    break
                }
            }
        }

        log = zerolog.New(output).
            Level(zerolog.Level(logLevel)).
            With().
            Timestamp().
            Str("git_revision", gitRevision).
            Str("go_version", buildInfo.GoVersion).
            Logger()
    })

    return log
}