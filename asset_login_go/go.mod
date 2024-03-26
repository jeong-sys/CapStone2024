module main

go 1.21.6

replace capstone.com/module/db => ./db

replace capstone.com/module/handler => ./handler

replace capstone.com/module/models => ./models

replace capstone.com/module/hashing => ./hashing

require (
	capstone.com/module/handler v0.0.0-00010101000000-000000000000
	github.com/labstack/echo v3.3.10+incompatible
)

require (
	capstone.com/module/db v0.0.0-00010101000000-000000000000 // indirect
	capstone.com/module/hashing v0.0.0-00010101000000-000000000000 // indirect
	capstone.com/module/models v0.0.0-00010101000000-000000000000 // indirect
	dmitri.shuralyov.com/gpu/mtl v0.0.0-20221208032759-85de2813cf6b // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	gioui.org v0.5.0 // indirect
	gioui.org/cpu v0.0.0-20220412190645-f1e9e8c3b1f7 // indirect
	gioui.org/shader v1.0.8 // indirect
	github.com/BurntSushi/xgb v0.0.0-20210121224620-deaf085860bc // indirect
	github.com/aarzilli/gdlv v1.12.0 // indirect
	github.com/aarzilli/nucular v0.0.0-20240117103348-47eb8d7bfc14 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cilium/ebpf v0.13.2 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/derekparker/trie v0.0.0-20230829180723-39f4de51ef7d // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-delve/delve v1.22.1 // indirect
	github.com/go-delve/liner v1.2.3-0.20231231155935-4726ab1d7f62 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20240307211618-a69d953ea142 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-sql-driver/mysql v1.8.0 // indirect
	github.com/go-text/typesetting v0.1.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/google/go-dap v0.12.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/labstack/echo/v4 v4.11.4 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.starlark.net v0.0.0-20240314022150-ee8ed142361c // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20240318143956-a85f2c67cd81 // indirect
	golang.org/x/exp/shiny v0.0.0-20240318143956-a85f2c67cd81 // indirect
	golang.org/x/image v0.15.0 // indirect
	golang.org/x/mobile v0.0.0-20240320162201-c76e57eead38 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
