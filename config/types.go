package config

type Config struct {
	Mysql    MySQLConfig    `toml:"mysql" mapstructure:"mysql"`
	Webauthn WebauthnConfig `toml:"webauthn" mapstructure:"webauthn"`
}

type MySQLConfig struct {
	DSN string `toml:"dsn" env:"DSN" mapstructure:"dsn"`
}

func NewMySQLConfig() MySQLConfig {
	return MySQLConfig{
		DSN: "root:123456@tcp(127.0.0.1:3306)/passkey",
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

func NewConfig() *Config {
	return &Config{
		Mysql:    NewMySQLConfig(),
		Webauthn: NewDefaultWebauthnConfig(),
	}
}
