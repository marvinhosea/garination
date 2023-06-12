package config

type AppDb struct {
	Host string
	Port string
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
	Port     string
	Password string
}

type App struct {
	Port        string
	MetricsPort string
}

type Config struct {
	AppDb   AppDb
	Redis   Redis
	Casdoor Casdoor
	App     App
}
