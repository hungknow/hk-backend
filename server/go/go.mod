module github.com/hungknow/hktrading_server

go 1.22

toolchain go1.22.1

require (
	github.com/go-chi/chi/v5 v5.0.8
	hungknow.com/blockchain v0.0.0-00010101000000-000000000000
)

replace hungknow.com/blockchain => ../..
