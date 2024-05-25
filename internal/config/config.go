package config

import (
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3"
	"os"
)

type Config struct {
	host     string
	port     int
	logLevel string
}

func (c *Config) String() string {
	return fmt.Sprintf("Config(host=%v, port=%v, logLevel=%v)", c.host, c.port, c.logLevel)
}

func (c *Config) Host() string {
	return c.host
}

func (c *Config) Port() int {
	return c.port
}

func (c *Config) LogLevel() string {
	return c.logLevel
}

func Load() (*Config, error) {
	fs := flag.NewFlagSet("golang-junior-dev", flag.ExitOnError)
	cfg := new(Config)
	fs.StringVar(&cfg.host, "host", "localhost", "listen host")
	fs.IntVar(&cfg.port, "port", 45385, "listen port")
	fs.StringVar(&cfg.logLevel, "log level", "debug", "log level")
	fs.String("config", "", "config file (optional)")

	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("GJD_"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)
	if err != nil {
		return nil, fmt.Errorf("new config: parse os.Args: %w", err)
	}
	return cfg, nil
}
