package utils

import (
	"bookcast/configs"
	"fmt"
)

func ConnectionURLBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var url string
	switch stuff {
	case "gin":
		url = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	case "postgresql":
		url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.PostgreSQL.User,
			cfg.PostgreSQL.Pass,
			cfg.PostgreSQL.Host,
			cfg.PostgreSQL.Port,
			cfg.PostgreSQL.Database,
			cfg.PostgreSQL.SSLMode,
		)
	default:
		return "", fmt.Errorf("unknown connection type: %s", stuff)
	}
	return url, nil
}
