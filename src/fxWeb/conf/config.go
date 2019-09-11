package conf

var yml *string

type Config struct {
	YmlFile      *string       `json:"ymlFile"`
	ServerConfig *ServerConfig `json:"server"`
}

type ServerConfig struct {
	ServerName string `json:"servername"`
	Version    string `json:"version"`
	Mode       string `json:"mode"`
}

func NewConfig() *Config {
	return &Config{
		YmlFile: yml,
		ServerConfig: &ServerConfig{
			ServerName: "DemoApp",
			Version:    "v1.0",
			Mode:       "develop",
		},
	}
}
