package config

type ConsulConfig struct {
	Host string
	Port int
	Key  string
}

type ServerConfig struct {
	Name        string       `mapstructure:"name" json:"name"`
	Host        string       `mapstructure:"host" json:"host"`
	Port        int          `mapstructure:"port" json:"port"`
	OtelInfo    OtelConfig   `mapstructure:"otel" json:"otel"`
	UserSrvInfo RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
