module capstone.com/module/asset

go 1.21.6

require (
	capstone.com/module/db v0.0.0
	capstone.com/module/tcp v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.7.1
)

replace capstone.com/module/db => ../db

replace capstone.com/module/tcp => ../tcp
