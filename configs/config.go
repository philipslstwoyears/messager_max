package configs

type Config struct {
	BindAddr string          `toml:"bindAddr"`
	LogLevel string          `toml:"logLevel"`
	Postgres *PostgresConfig `toml:"postgres"`
}

type PostgresConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}
