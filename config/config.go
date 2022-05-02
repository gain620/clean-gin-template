package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		MyLog    `yaml:"my_log"`
		PG       `yaml:"postgres"`
		RMQ      `yaml:"rabbitmq"`
		Database `yaml:"database"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// MyLog TODO : Fix MyLog reading problem
	MyLog struct {
		Level      string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
		RollbarEnv string `yaml:"rollbar_env"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	RMQ struct {
		ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
		ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
		URL            string `env-required:"true"                            env:"RMQ_URL"`
	}

	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		PoolMax  int    `yaml:"pool_max"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		Name     string `yaml:"name"`
		Location string `yaml:"location"`
	}
)

// NewConfig returns app config
func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to read config file, %v", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into config struct, %v", err)
	}

	return cfg, nil
}
