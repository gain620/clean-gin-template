package config

const (
	EXT_JSON = ".json"
	EXT_YAML = ".yaml"
	EXT_YML  = ".yml"
)

// Config configuration struct
type Config struct {
	Port           int    `json:"port,omitempty"`
	Addr           string `json:"addr,omitempty"`
	Scheme         string `json:"scheme,omitempty"`
	ProductionMode bool   `json:"production_mode,omitempty"`
	BasePath       string `json:"base_path"`
	DebugEnabled   bool   `json:"debug_enabled,omitempty"`
	AppCfgsPath    string `json:"app_configs_path"`
}

// NewConfig creates an instance of Config from command-line args and/or env vars
func NewConfig() *Config {
	cfg := &Config{}

	return cfg
}
