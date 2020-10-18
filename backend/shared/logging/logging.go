package logging

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/wcse/monorepo/backend/api/config/base"
)

// NewLogger ...
func NewLogger(msg *base.Logger) (*zap.Logger, error) {
	var c zap.Config
	var opts []zap.Option
	if msg.GetPretty() {
		c = zap.NewDevelopmentConfig()
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	} else {
		c = zap.NewProductionConfig()
	}

	level := zap.NewAtomicLevel()

	levelName := "INFO"
	if msg.Level != base.Logger_UNSPECIFIED {
		levelName = msg.Level.String()
	}

	if err := level.UnmarshalText([]byte(levelName)); err != nil {
		return nil, fmt.Errorf("could not parse log level %s", msg.Level.String())
	}
	c.Level = level

	return c.Build(opts...)
}

// NewTmpLogger ...
func NewTmpLogger() *zap.Logger {
	c := zap.NewProductionConfig()
	c.DisableStacktrace = true
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	return l
}
