package config

import (
	"errors"
	"flag"
	"os"
)

// Flag struct for parsing from env and cmd args.
type Flag struct {
	Key         *string
	PathToItems *string
}

var (
	f Flag

	// ErrKeyNotSet error when the key is not set.
	ErrKeyNotSet = errors.New("key not set")
)

func init() {
	f.Key = flag.String("key", "", "-key=KEY")
	f.PathToItems = flag.String("items", "items.json", "-config=path/to/items.json")
}

// Config contains all the settings for configuring the application.
type Config struct {
	Key         string
	PathToItems string
}

// New initializing the config for the application.
func New() (*Config, error) {
	flag.Parse()

	if key, ok := os.LookupEnv("TELEGRAM_BOT_KEY"); ok {
		*f.Key = key
	}

	if *f.Key == "" {
		return nil, ErrKeyNotSet
	}

	return &Config{
		Key:         *f.Key,
		PathToItems: *f.PathToItems,
	}, nil
}
