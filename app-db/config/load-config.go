package config

import (
	"fmt"
	"os"
)

func LoadConfig() (Config, error) {
	var config Config

	// Load Casdoor configuration
	casdoorEndpoint, found := os.LookupEnv("CASDOOR_ENDPOINT")
	if !found {
		return config, fmt.Errorf("CASDOOR_ENDPOINT environment variable not set")
	}
	config.Casdoor.Endpoint = casdoorEndpoint

	config.Casdoor.ClientId, found = os.LookupEnv("CASDOOR_CLIENT_ID")
	if !found {
		return config, fmt.Errorf("CASDOOR_CLIENT_ID environment variable not set")
	}

	config.Casdoor.ClientSecret, found = os.LookupEnv("CASDOOR_CLIENT_SECRET")
	if !found {
		return config, fmt.Errorf("CASDOOR_CLIENT_SECRET environment variable not set")
	}

	config.Casdoor.OrganisationName, found = os.LookupEnv("CASDOOR_ORGANISATION_NAME")
	if !found {
		return config, fmt.Errorf("CASDOOR_ORGANISATION_NAME environment variable not set")
	}

	config.Casdoor.ApplicationName, found = os.LookupEnv("CASDOOR_APPLICATION_NAME")
	if !found {
		return config, fmt.Errorf("CASDOOR_APPLICATION_NAME environment variable not set")
	}

	config.Casdoor.CertificateX509, found = os.LookupEnv("CASDOOR_CERTIFICATE_X509")
	if !found {
		return config, fmt.Errorf("CASDOOR_CERTIFICATE_X509 environment variable not set")
	}

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

	return config, nil
}
