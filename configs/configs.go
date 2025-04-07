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
	User     string
	Pass     string
	Database string
	SSLMode  string
}
