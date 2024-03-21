package db

import (
	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"
)

func GetConnector(DB_name string) *sql.DB {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "12345678",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               DB_name,
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}