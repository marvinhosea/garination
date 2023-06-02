package config

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     int
	Password string
}

type Config struct {
	Postgres Postgres
	Redis    Redis
}
