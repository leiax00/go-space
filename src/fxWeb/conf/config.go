package conf

type Config struct {
	ServerName string `json:"servername"`
	Version    string `json:"version"`
	Mode       string `json:"mode"`
}

func NewConfig() *Config {
	return &Config{
		ServerName: "DemoApp",
		Version:    "v1.0",
		Mode:       "develop",
	}
}
