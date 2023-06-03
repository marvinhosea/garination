package config

type Casdoor struct {
	Endpoint         string
	ClientId         string
	ClientSecret     string
	OrganisationName string
	ApplicationName  string
	CertificateX509  string
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type App struct {
	Port string
}

type Config struct {
	Casdoor  Casdoor
	Postgres Postgres
	App      App
}
