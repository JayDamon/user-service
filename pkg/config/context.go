package config

import "database/sql"

type Context struct {
	DB     *sql.DB
	Config *Config
}
