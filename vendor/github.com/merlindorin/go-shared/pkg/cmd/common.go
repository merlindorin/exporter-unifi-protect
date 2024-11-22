package cmd

import (
	"fmt"

	"golang.org/x/text/message"

	"github.com/merlindorin/go-shared/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Commons defines the common flags and embedded commands for printing version
// and licence information, utilized by the command-line interface.
type Commons struct {
	Development bool   `short:"D" env:"DEBUG,DEV,DEVELOPMENT" help:"Set to true to enable development mode with debug-level logging."`
	Level       string `short:"l" env:"LOG_LEVEL" help:"Specify the logging level, options are: debug, info, warn, error, fatal." default:"info"`
	Lang        string `env:"LANG" help:"Specify the print lang for tailored message." default:"en"`

	Version Version `cmd:"" help:"Display version information."`
	Licence Licence `cmd:"" help:"Show the application's licence."`
}

// Logger initializes a new zap.Logger based on the Development and Level fields in the commons struct.
// It returns the configured logger or an error if the logging level is invalid or the logger cannot be created.
func (c *Commons) Logger() (*zap.Logger, error) {
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse logger level \"%s\": %w", c.Level, err)
	}

	if c.Development {
		level = zapcore.DebugLevel
	}

	return logger.New(level, c.Development)
}

// MustLogger will panic if a logger can't be provided.
func (c *Commons) MustLogger() *zap.Logger {
	l, err := c.Logger()
	if err != nil {
		panic(err)
	}

	return l
}

// Printer returns a new message.Printer configured for the specified language.
// This printer can be used to format and print localized messages within
// the command-line interface, ensuring that output is tailored to the user's
// language preference.
func (c *Commons) Printer() Printer {
	return message.NewPrinter(message.MatchLanguage(c.Lang, "en"))
}
