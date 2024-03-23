module main

go 1.21.6

replace capstone.com/module/db => ./db

replace capstone.com/module/handler => ./handler

replace capstone.com/module/models => ./models

require (
	capstone.com/module/handler v0.0.0-00010101000000-000000000000
	github.com/labstack/echo v3.3.10+incompatible
)

require (
	capstone.com/module/db v0.0.0-00010101000000-000000000000 // indirect
	capstone.com/module/models v0.0.0-00010101000000-000000000000 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
