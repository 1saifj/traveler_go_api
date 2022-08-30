package config

import "fmt"

type AppConfig struct {
	Info  AppInfo        `yaml:"app_info"`
	Hosts []Host         `yaml:"hosts"`
	DB    DatabaseConfig `yaml:"database"`
}

type AppInfo struct {
	AppName    string `yaml:"app_name"`
	AppVersion string `yaml:"app_version"`
}

type Host struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (host1 Host) GetHost() string {
	return fmt.Sprintf("%s:%d", host1.Host, host1.Port)
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Timezone string `yaml:"timezone"`
	SslMode  bool   `yaml:"ssl_mode"`
}
