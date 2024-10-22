package settings

type Config struct {
	Server ServerSetting `yaml:"server"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}