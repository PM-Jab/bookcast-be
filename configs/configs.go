package configs

type Configs struct {
	App        Gin
	PostgreSQL PostgreSQL
}

type Gin struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}
