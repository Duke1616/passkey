package config

type Config struct {
	Mysql    MySQLConfig    `toml:"mysql" mapstructure:"mysql"`
	Webauthn WebauthnConfig `toml:"webauthn" mapstructure:"webauthn"`
	Redis    RedisConfig    `toml:"redis" mapstructure:"redis"`
}

type MySQLConfig struct {
	HOST     string `toml:"host" env:"HOST" mapstructure:"host"`
	PORT     int32  `toml:"port" env:"PORT" mapstructure:"port"`
	USERNAME string `toml:"username" env:"USERNAME" mapstructure:"username"`
	PASSWORD string `toml:"password" env:"PASSWORD" mapstructure:"password"`
	DATABASE string `toml:"database" env:"DATABASE" mapstructure:"database"`
}

func NewDefaultMySQLConfig() MySQLConfig {
	return MySQLConfig{
		USERNAME: "root",
		PASSWORD: "123456",
		PORT:     3306,
		HOST:     "127.0.0.1",
		DATABASE: "passkey",
	}
}

type WebauthnConfig struct {
	RPID          string `toml:"rp_id" env:"RPID" mapstructure:"rp_id"`
	RPDisplayName string `toml:"rp_display_name" env:"RPDisplayName" mapstructure:"rp_display_name"`
	RPOrigins     string `toml:"rp_origins" env:"RPOrigins" mapstructure:"rp_origins"`
}

func NewDefaultWebauthnConfig() WebauthnConfig {
	return WebauthnConfig{
		RPID:          "localhost",
		RPDisplayName: "WebAuthn Example Application",
		RPOrigins:     "http://localhost:8100",
	}
}

type RedisConfig struct {
	Addr string `toml:"addr" env:"ADDR" mapstructure:"addr"`
}

func NewDefaultRedisConfig() RedisConfig {
	return RedisConfig{
		Addr: "127.0.0.1:6379",
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		Mysql:    NewDefaultMySQLConfig(),
		Webauthn: NewDefaultWebauthnConfig(),
		Redis:    NewDefaultRedisConfig(),
	}
}
