module github.com/hungknow/hktrading_server

go 1.22

toolchain go1.22.1

require (
	github.com/go-chi/chi/v5 v5.0.8
	github.com/phuslu/log v1.0.89
	hungknow.com/blockchain v0.0.0-00010101000000-000000000000
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/parsers/dotenv v0.1.0 // indirect
	github.com/knadh/koanf/parsers/yaml v0.1.0 // indirect
	github.com/knadh/koanf/providers/file v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace hungknow.com/blockchain => ../..
