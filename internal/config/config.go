package config

import (
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3"
	"os"
)

type Config struct {
	listen       string
	dbConnString string
}

func (c *Config) DbConnString() string {
	return c.dbConnString
}

func (c *Config) Listen() string {
	return c.listen
}

func Load() (*Config, error) {
	fs := flag.NewFlagSet("cute-chat-backend", flag.ExitOnError)
	cfg := new(Config)
	fs.StringVar(&cfg.listen, "listen", "localhost:46473", "listen listen")
	//refresh    = fs.Duration("refresh", 15*time.Second, "refresh interval")
	//debug      = fs.Bool("debug", false, "log debug information")
	fs.String("config", "", "config file (optional)")

	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("CCB_"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)
	if err != nil {
		return nil, fmt.Errorf("new config: parse os.Args: %w", err)
	}
	return cfg, nil
}
