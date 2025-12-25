//go:build !sqlite3

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var opens = map[string]func(string) gorm.Dialector{
	"tidb":  mysql.Open,
	"mysql": mysql.Open,

	"postgres":   postgres.Open,
	"postgresql": postgres.Open,

	"sqlite":  sqlite.Open,
	"sqlite3": sqlite.Open,

	"sqlserver": sqlserver.Open,
}
