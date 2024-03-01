package config

import "time"

type Config struct {
	HttpServer HttpServer `mapstructure:"HttpServer"`
	Database   Database   `mapstructure:"Database"`
	Redis      Redis      `mapstructure:"Redis"`
	Auth       Auth       `mapstructure:"Auth"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type HttpServer struct {
	Port            int           `mapstructure:"port"`
	Timeout         time.Duration `mapstructure:"timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type Redis struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	ExpirationTime time.Duration `mapstructure:"expiration_time"`
}

//type JWTToken struct {
//	TimeToLive time.Duration `mapstructure:"time_to_live"`
//}

type Auth struct {
	Access  Access  `mapstructure:"Access"`
	Refresh Refresh `mapstructure:"Refresh"`
}

type Access struct {
	//PasswordSecretKey string        `mapstructure:"PasswordSecretKey"`
	//JwtSecretKey      string        `mapstructure:"JwtSecretKey"`
	TimeToLive time.Duration `mapstructure:"TimeToLive"`
}

type Refresh struct {
	//PasswordSecretKey string        `mapstructure:"PasswordSecretKey"`
	//JwtSecretKey      string        `mapstructure:"JwtSecretKey"`
	TimeToLive time.Duration `mapstructure:"TimeToLive"`
}
