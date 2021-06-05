package config

import "fmt"

type Config struct {
	Server ServerConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	HttpPort string
}

type MySQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
}

// mysql -uroot -hvtol-duc.tokyo
func MySQLConfigDefault() MySQLConfig {
	return MySQLConfig{
		Host:     "vtol-duc.tokyo",
		Port:     "3306",
		Username: "root",
		Password: "Minhduc1992",
		Database: "ducdb",
	}
}

func (c MySQLConfig) ToString() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database)
	return dsn
}
