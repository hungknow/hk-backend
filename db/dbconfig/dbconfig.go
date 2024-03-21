package dbconfig

type DBConfig struct {
	Host         string `koanf:"host"`
	Port         int    `koanf:"port"`
	Username     string `koanf:"username"`
	Password     string `koanf:"password"`
	DatabaseName string `koanf:"database_name"`
}
