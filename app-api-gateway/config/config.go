package config

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Casdoor struct {
	Endpoint         string
	ClientId         string
	ClientSecret     string
	OrganisationName string
	ApplicationName  string
	CertificateX509  string
}

type Redis struct {
	Host     string
	Port     int
	Password string
}

type Config struct {
	Postgres Postgres
	Redis    Redis
	Casdoor  Casdoor
}
