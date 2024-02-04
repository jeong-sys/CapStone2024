module capstone.com/module/tcp

go 1.21.6

require capstone.com/module/asset v0.0.0

require (
	capstone.com/module/db v0.0.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
)

replace capstone.com/module/asset => ../asset

replace capstone.com/module/db => ../db
