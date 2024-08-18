package outputter

import (
	"bufio"
	"io"
	"log/slog"
	"os"
)

type Outputter struct {
	StdOut io.Writer
	StdErr *slog.Logger
}

type cfg struct {
	stdOut  io.Writer
	stdErr  io.Writer
	handler slog.Handler
}

type Option func(*cfg)

func WithStdOut(o io.Writer) Option {
	return func(c *cfg) {
		c.stdOut = o
	}
}

func WithBufferedStdOut() Option {
	return func(c *cfg) {
		c.stdOut = bufio.NewWriter(c.stdOut)
	}
}

func WithStdErr(o io.Writer) Option {
	return func(c *cfg) {
		c.stdErr = o
	}
}

func WithTextHandler() Option {
	return func(c *cfg) {
		c.handler = slog.NewTextHandler(c.stdErr, nil)
	}
}

func WithJSONLogging() Option {
	return func(c *cfg) {
		c.handler = slog.NewJSONHandler(c.stdErr, nil)
	}
}

func New(opts ...Option) *Outputter {
	config := &cfg{
		stdOut:  os.Stdout,
		stdErr:  os.Stderr,
		handler: slog.Default().Handler(),
	}
	for _, o := range opts {
		o(config)
	}

	return &Outputter{
		StdOut: config.stdOut,
		StdErr: slog.New(config.handler),
	}
}

func (o *Outputter) Close() error {
	if f, ok := o.StdOut.(*bufio.Writer); ok {
		if err := f.Flush(); err != nil {
			return err
		}
	}
	if c, ok := o.StdOut.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
