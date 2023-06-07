package config

import (
	"fmt"
	"os"
)

func LoadConfig() (Config, error) {
	var config Config
	var found bool

	// Load Postgres configuration
	config.Postgres.Host, found = os.LookupEnv("POSTGRES_HOST")
	if !found {
		return config, fmt.Errorf("POSTGRES_HOST environment variable not set")
	}

	config.Postgres.Port, found = os.LookupEnv("POSTGRES_PORT")
	if !found {
		return config, fmt.Errorf("POSTGRES_PORT environment variable not set")
	}

	config.Postgres.User, found = os.LookupEnv("POSTGRES_USER")
	if !found {
		return config, fmt.Errorf("POSTGRES_USER environment variable not set")
	}

	config.Postgres.Password, found = os.LookupEnv("POSTGRES_PASSWORD")
	if !found {
		return config, fmt.Errorf("POSTGRES_PASSWORD environment variable not set")
	}

	config.Postgres.Database, found = os.LookupEnv("POSTGRES_DATABASE")
	if !found {
		return config, fmt.Errorf("POSTGRES_DATABASE environment variable not set")
	}

	// Load App configuration
	config.App.Port, found = os.LookupEnv("APP_PORT")
	if !found {
		return config, fmt.Errorf("APP_PORT environment variable not set")
	}

	config.App.MetricsPort, found = os.LookupEnv("APP_METRICS_PORT")
	if !found {
		return config, fmt.Errorf("APP_METRICS_PORT environment variable not set")
	}

	return config, nil
}
