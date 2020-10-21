package config

type Config struct {
	MySQLDNS string
}

func NewDefaultConfig() *Config {
	return &Config{
		MySQLDNS: "default:secret@tcp(127.0.0.1:3306)/dogfood?charset=utf8mb4&parseTime=True&loc=Local",
	}
}
