package main

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required: "true"`
	Listen struct {
		Type string `yaml:"type" env-default: "port"`
		BindIP string `yaml:"bind-ip" env-default: "127.0.0.1"`
		Port string `yaml:"port" env-default: "8080"`
	} `yaml: "listen"`
	Storage StorageConfig yaml:"storage"
}

type StorageConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`

}
