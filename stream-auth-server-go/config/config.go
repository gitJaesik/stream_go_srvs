package config

import (
	"os"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var (
	envs []string
)

type SGLConfig struct {
	values map[string]string
	Err    error
}

var Set = wire.NewSet(
	InitializeViperConfig,
)

var SglConfig SGLConfig

// InitializeViperConfig
func InitializeViperConfig() *SGLConfig {

	os.Setenv("STREAMAUTHGO_AUTH_SERVER_GRPC_PORT", "8280")
	os.Setenv("STREAMAUTHGO_AUTH_SERVER_GW_PORT", "8281")
	os.Setenv("STREAMAUTHGO_CRON_SECONDS", "30")

	defaultEnvs := map[string]string{
		"CRON_SECONDS":          "30",
		"AUTH_SERVER_GRPC_PORT": "8280",
		"AUTH_SERVER_GW_PORT":   "8281",
	}

	SglConfig.values = map[string]string{}
	envs = []string{
		"CRON_SECONDS",
		"AUTH_SERVER_GRPC_PORT",
		"AUTH_SERVER_GW_PORT",
	}

	viper.SetEnvPrefix("STREAMAUTHGO")

	for _, key := range envs {
		err := viper.BindEnv(key)

		if err != nil {
			panic(err)
		}

		// SglConfig.values[key] = viper.Get(key).(string)
		viperEnvVal := viper.Get(key)
		if viperEnvVal != nil {
			SglConfig.values[key] = viperEnvVal.(string)
		} else {
			SglConfig.values[key] = defaultEnvs[key]
		}
	}

	return &SglConfig
}

// GetValues ...
func (s *SGLConfig) GetValues() map[string]string {
	return s.values
}

// GetCronSeconds ...
func (s *SGLConfig) GetCronSeconds() string {
	return s.values["CRON_SECONDS"]
}

// GetAuthServerGrpcPort ...
func (s *SGLConfig) GetAuthServerGrpcPort() string {
	return s.values["AUTH_SERVER_GRPC_PORT"]
}

// GetAuthServerGwPort ...
func (s *SGLConfig) GetAuthServerGwPort() string {
	return s.values["AUTH_SERVER_GW_PORT"]
}

////

// GetValues ...
func GetValues() map[string]string {
	return SglConfig.values
}

// GetCronSeconds ...
func GetCronSeconds() string {
	return SglConfig.values["CRON_SECONDS"]
}

// GetAuthServerGrpcPort ...
func GetAuthServerGrpcPort() string {
	return SglConfig.values["AUTH_SERVER_GRPC_PORT"]
}

// GetAuthServerGwPort ...
func GetAuthServerGwPort() string {
	return SglConfig.values["AUTH_SERVER_GW_PORT"]
}
