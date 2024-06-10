package main

import (
	"fmt"
	"io"
	"os"
)

type CliConfig struct {
	ErrStream, OutStream io.Writer
}

func app(words []string, cfg CliConfig) {
	for _, w := range words {
		if len(w)%2 == 0 {
			fmt.Fprintf(cfg.OutStream, "word %s is even\n", w)
		} else {
			fmt.Fprintf(cfg.ErrStream, "word %s is odd\n", w)
		}
	}
}

func NewCliConfig(opts ...Option) (CliConfig, error) {
	c := CliConfig{
		ErrStream: os.Stderr,
		OutStream: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return CliConfig{}, err
		}
	}
	return c, nil
}

type Option func(*CliConfig) error

func WithErrStream(errStream io.Writer) Option {
	return func(c *CliConfig) error {
		c.ErrStream = errStream
		return nil
	}
}

func WithOutStream(outStream io.Writer) Option {
	return func(c *CliConfig) error {
		c.OutStream = outStream
		return nil
	}
}

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "No words provided.")
		os.Exit(1)
	}

	cfg, err := NewCliConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config: %v\n", err)
		os.Exit(1)
	}

	app(words, cfg)
}
