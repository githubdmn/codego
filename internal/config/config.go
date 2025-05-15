package config

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: ":8080",
		},
		DB: DBConfig{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "postgres",
			DBName:   "restapi",
			SSLMode:  "disable",
		},
	}, nil
}
