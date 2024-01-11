package config

type Config struct {
	Mysql    MySQLConfig    `toml:"mysql" mapstructure:"mysql"`
	Webauthn WebauthnConfig `toml:"webauthn" mapstructure:"webauthn"`
}

type MySQLConfig struct {
	DSN string `toml:"dsn" env:"DSN" mapstructure:"dsn"`
}

type WebauthnConfig struct {
	RPID          string `toml:"rp_id" env:"RPID" mapstructure:"rp_id"`
	RPDisplayName string `toml:"rp_display_name" env:"RPDisplayName" mapstructure:"rp_display_name"`
	RPOrigins     string `toml:"rp_origins" env:"RPOrigins" mapstructure:"rp_origins"`
}
