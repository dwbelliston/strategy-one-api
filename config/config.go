package config

type Config struct {
	DB  *DBConfig
	SQL map[string]string
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Endpoint string
	Port     string
	Name     string
	Charset  string
}

// GetConfig for app
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgresql",
			Username: "awsuser",
			Password: "Password1!",
			Endpoint: "strategy-one-db.cluster-ciwkcai1iw95.us-east-1.rds.amazonaws.com",
			Port:     "5432",
			Name:     "dev",
			Charset:  "utf8",
		},
		SQL: map[string]string{
			"GET_ALL_SHAPES": "SELECT shape_id, title, sides, created_at FROM shapes",
		},
	}
}
