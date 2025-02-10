package config

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func Load() *Config {
	return &Config{
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "minimoodle",
		DBHost:     "localhost",
		DBPort:     "5432",
	}
}
