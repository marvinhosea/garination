package config

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
	Postgres Postgres
	App      App
}
