module module/asset

go 1.21.6

require (
	github.com/go-sql-driver/mysql v1.7.1
	module/db v0.0.0-00010101000000-000000000000
)

replace module/db => ../db
