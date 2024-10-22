package settings

type Config struct {
	Server ServerSetting `yaml:"server"`
	Logger LoggerSetting `yaml:"logger"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}

type LoggerSetting struct {
	Log_level string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups int `mapstructure:"max_backups"`
	Max_age int `mapstructure:"max_age"`
	Max_size int `mapstructure:"max_size"`
	Compress bool `mapstructure:"compress"`
}