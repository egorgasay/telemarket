package config

import (
	"errors"
	"flag"
)

// Flag struct for parsing from env and cmd args.
type Flag struct {
	Key *string
}

var (
	f Flag

	// ErrKeyNotSet error when the key is not set.
	ErrKeyNotSet = errors.New("key not set")
)

func init() {
	f.Key = flag.String("key", "", "-key=KEY")
}

// Config contains all the settings for configuring the application.
type Config struct {
	Key string
}

// New initializing the config for the application.
func New() (*Config, error) {
	flag.Parse()

	if *f.Key == "" {
		return nil, ErrKeyNotSet
	}

	return &Config{
		Key: *f.Key,
	}, nil
}
