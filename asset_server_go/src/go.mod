module main

go 1.21.6

replace module/asset => ./asset

require module/asset v0.0.0-00010101000000-000000000000

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	module/db v0.0.0-00010101000000-000000000000 // indirect
)

replace module/db => ./db
