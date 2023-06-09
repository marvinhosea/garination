package config

import (
	"fmt"
	"os"
	"strconv"
)

func LoadConfig() (Config, error) {
	var config Config
	var found bool

	// Load AppDb configuration
	config.AppDb.Host, found = os.LookupEnv("APP_DB_HOST")
	if !found {
		return config, fmt.Errorf("APP_DB_HOST environment variable not set")
	}

	portStr, found := os.LookupEnv("APP_DB_PORT")
	if !found {
		return config, fmt.Errorf("APP_DB_PORT environment variable not set")
	}
	_, err := strconv.Atoi(portStr)
	if err != nil {
		return config, fmt.Errorf("failed to parse APP_DB_PORT as integer: %v", err)
	}
	config.AppDb.Port = portStr

	// Load Redis configuration
	config.Redis.Host, found = os.LookupEnv("REDIS_HOST")
	if !found {
		return config, fmt.Errorf("REDIS_HOST environment variable not set")
	}

	portStr, found = os.LookupEnv("REDIS_PORT")
	if !found {
		return config, fmt.Errorf("REDIS_PORT environment variable not set")
	}
	config.Redis.Port = portStr

	config.Redis.Password, found = os.LookupEnv("REDIS_PASSWORD")
	if !found {
		return config, fmt.Errorf("REDIS_PASSWORD environment variable not set")
	}

	// Load Casdoor configuration
	config.Casdoor.Endpoint, found = os.LookupEnv("CASDOOR_ENDPOINT")
	if !found {
		return config, fmt.Errorf("CASDOOR_ENDPOINT environment variable not set")
	}

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

	// Load AppDb configuration
	config.AppDb.Host, found = os.LookupEnv("APP_DB_HOST")
	if !found {
		return config, fmt.Errorf("APP_DB_HOST environment variable not set")
	}

	portStr, found = os.LookupEnv("APP_DB_PORT")
	if !found {
		return config, fmt.Errorf("APP_DB_PORT environment variable not set")
	}
	_, err = strconv.Atoi(portStr)
	if err != nil {
		return config, fmt.Errorf("failed to parse APP_DB_PORT as integer: %v", err)
	}
	config.AppDb.Port = portStr

	return config, nil
}
