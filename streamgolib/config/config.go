package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	envs []string
)

type SGLConfig struct {
	values map[string]string
	Err    error
}

var SglConfig SGLConfig

func init() {
	fmt.Println("config.go init start")
	InitializeViperConfig()
	fmt.Println("config.go init end")
}

// InitializeViperConfig
func InitializeViperConfig() *SGLConfig {

	/*
		os.Setenv("KEY", "VALUE")
	*/

	defaultEnvs := map[string]string{
		// "MONGODB_URI": "mongodb://mongodb:27017",
		"MONGODB_URI":                  "mongodb://0.0.0.0:27017",
		"MONGODB_DATABASE":             "stream",
		"MONGODB_NOTE_COLLECTION":      "note",
		"MONGODB_PLAYER_COLLECTION":    "player",
		"MONGODB_VOICENOTE_COLLECTION": "voicenote",
		"MONGODB_USER_COLLECTION":      "user",
		"CRON_SECONDS":                 "30",
		"GOOGLE_STT_IP":                "127.0.0.1",
		"GOOGLE_STT_PORT":              "9999",
		"FILE_SERVER_HTTP_HTTPS":       "http://",
		"FILE_SERVER_IP":               "127.0.0.1",
		"FILE_SERVER_GRPC_PORT":        "8180",
		"AUTH_SERVER_IP":               "127.0.0.1",
		"AUTH_SERVER_GRPC_PORT":        "8280",
		"AUTH_SERVER_GW_PORT":          "8281",
		"CKEY":                         "CIPHERKEY01234567890123456789012",
		"CIVKEY":                       "CIPHERIVKEY01234",
		"LOG_LEVEL":                    "InfoLevel",
		"LOG_USE_CONSOLELOG":           "y",
		"LOG_USE_FILELOG":              "n",
		"LOG_FILENAME":                 "log/app-%Y-%m-%d-%H.log",
		"LOG_FILENAME_LINK":            "log/app.log",
		"JWT_TTL":                      "10800",
		"CERT_FILE":                    "x509/server_cert.pem",
		"KEY_FILE":                     "x509/server_key.pem",
		"ACCESS_PRIVATE_PEM":           "access_refresh/access-private.pem",
		"ACCESS_PUBLIC_PEM":            "access_refresh/access-public.pem",
		"REFRESH_PRIVATE_PEM":          "access_refresh/refresh-private.pem",
		"REFRESH_PUBLIC_PEM":           "access_refresh/refresh-public.pem",
	}

	SglConfig.values = map[string]string{}
	envs = []string{
		"MONGODB_URI",
		"MONGODB_DATABASE",
		"MONGODB_VOICENOTE_COLLECTION",
		"MONGODB_NOTE_COLLECTION",
		"MONGODB_PLAYER_COLLECTION",
		"MONGODB_USER_COLLECTION",
		"CRON_SECONDS",
		"GOOGLE_STT_IP",
		"GOOGLE_STT_PORT",
		"FILE_SERVER_HTTP_HTTPS",
		"FILE_SERVER_IP",
		"FILE_SERVER_GRPC_PORT",
		"AUTH_SERVER_IP",
		"AUTH_SERVER_GRPC_PORT",
		"AUTH_SERVER_GW_PORT",
		"CKEY",
		"CIVKEY",
		"LOG_LEVEL",
		"LOG_USE_CONSOLELOG",
		"LOG_USE_FILELOG",
		"LOG_FILENAME",
		"LOG_FILENAME_LINK",
		"JWT_TTL",
		"CERT_FILE",
		"KEY_FILE",
		"ACCESS_PRIVATE_PEM",
		"ACCESS_PUBLIC_PEM",
		"REFRESH_PRIVATE_PEM",
		"REFRESH_PUBLIC_PEM",
	}

	viper.SetEnvPrefix("STREAMGOLIB")

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

// GetMongoUri ...
func (s *SGLConfig) GetMongoUri() string {
	return s.values["MONGODB_URI"]
}

// GetMongoDatabase ...
func (s *SGLConfig) GetMongoDatabase() string {
	return s.values["MONGODB_DATABASE"]
}

// GetMongoVoicenoteCollection ...
func (s *SGLConfig) GetMongoVoicenoteCollection() string {
	return s.values["MONGODB_VOICENOTE_COLLECTION"]
}

// GetMongoNoteCollection ...
func (s *SGLConfig) GetMongoNoteCollection() string {
	return s.values["MONGODB_NOTE_COLLECTION"]
}

// GetMongoPlayerCollection ...
func (s *SGLConfig) GetMongoPlayerCollection() string {
	return s.values["MONGODB_PLAYER_COLLECTION"]
}

// GetMongoUserCollection ...
func (s *SGLConfig) GetMongoUserCollection() string {
	return s.values["MONGODB_USER_COLLECTION"]
}

// GetCronSeconds ...
func (s *SGLConfig) GetCronSeconds() string {
	return s.values["CRON_SECONDS"]
}

// GetGoogleSttIp ...
func (s *SGLConfig) GetGoogleSttIp() string {
	return s.values["GOOGLE_STT_IP"]
}

// GetGoogleSttPort ...
func (s *SGLConfig) GetGoogleSttPort() string {
	return s.values["GOOGLE_STT_PORT"]
}

// GetFileServerHttpHttps ...
func (s *SGLConfig) GetFileServerHttpHttps() string {
	return s.values["FILE_SERVER_HTTP_HTTPS"]
}

// GetFileServerIp ...
func (s *SGLConfig) GetFileServerIp() string {
	return s.values["FILE_SERVER_IP"]
}

// GetFileServerGrpcPort ...
func (s *SGLConfig) GetFileServerGrpcPort() string {
	return s.values["FILE_SERVER_GRPC_PORT"]
}

// GetAuthServerGrpcIp ...
func (s *SGLConfig) GetAuthServerIp() string {
	return s.values["AUTH_SERVER_IP"]
}

// GetAuthServerGrpcPort ...
func (s *SGLConfig) GetAuthServerGrpcPort() string {
	return s.values["AUTH_SERVER_GRPC_PORT"]
}

// GetAuthServerGwPort ...
func (s *SGLConfig) GetAuthServerGwPort() string {
	return s.values["AUTH_SERVER_GW_PORT"]
}

// GetCKey ...
func (s *SGLConfig) GetCKey() string {
	return s.values["CKEY"]
}

// GetCIVKey ...
func (s *SGLConfig) GetCIVKey() string {
	return s.values["CIVKEY"]
}

// GetLogLevel ...
func (s *SGLConfig) GetLogLevel() string {
	return s.values["LOG_LEVEL"]
}

// GetLogUseConsolelog ...
func (s *SGLConfig) IsUseConsolelog() string {
	return s.values["LOG_USE_CONSOLELOG"]
}

// GetLogUseFilelog ...
func (s *SGLConfig) IsUseFilelog() string {
	return s.values["LOG_USE_FILELOG"]
}

// GetLogFilename ...
func (s *SGLConfig) GetLogFilename() string {
	return s.values["LOG_FILENAME"]
}

// GetLogFilenameLink ...
func (s *SGLConfig) GetLogFilenameLink() string {
	return s.values["LOG_FILENAME_LINK"]
}

// GetJwtTtl ...
func (s *SGLConfig) GetJwtTtl() string {
	return s.values["JWT_TTL"]
}

// GetCertFile ...
func (s *SGLConfig) GetCertFile() string {
	return s.values["CERT_FILE"]
}

// GetKeyFile ...
func (s *SGLConfig) GetKeyFile() string {
	return s.values["KEY_FILE"]
}

// GetAccessPrivatePem ...
func (s *SGLConfig) GetAccessPrivatePem() string {
	return s.values["ACCESS_PRIVATE_PEM"]
}

// GetAccessPublicPem ...
func (s *SGLConfig) GetAccessPublicPem() string {
	return s.values["ACCESS_PUBLIC_PEM"]
}

// GetRefreshPrivatePem ...
func (s *SGLConfig) GetRefreshPrivatePem() string {
	return s.values["REFRESH_PRIVATE_PEM"]
}

// GetRefreshPublicPem ...
func (s *SGLConfig) GetRefreshPublicPem() string {
	return s.values["REFRESH_PUBLIC_PEM"]
}

/////

// GetValues ...
func GetValues() map[string]string {
	return SglConfig.values
}

// GetMongoUri ...
func GetMongoUri() string {
	return SglConfig.values["MONGODB_URI"]
}

// GetMongoDatabase ...
func GetMongoDatabase() string {
	return SglConfig.values["MONGODB_DATABASE"]
}

// GetMongoVoicenoteCollection ...
func GetMongoVoicenoteCollection() string {
	return SglConfig.values["MONGODB_VOICENOTE_COLLECTION"]
}

// GetMongoNoteCollection ...
func GetMongoNoteCollection() string {
	return SglConfig.values["MONGODB_NOTE_COLLECTION"]
}

// GetMongoPlayerCollection ...
func GetMongoPlayerCollection() string {
	return SglConfig.values["MONGODB_PLAYER_COLLECTION"]
}

// GetMongoUserCollection ...
func GetMongoUserCollection() string {
	return SglConfig.values["MONGODB_USER_COLLECTION"]
}

// GetCronSeconds ...
func GetCronSeconds() string {
	return SglConfig.values["CRON_SECONDS"]
}

// GetGoogleSttIp ...
func GetGoogleSttIp() string {
	return SglConfig.values["GOOGLE_STT_IP"]
}

// GetGoogleSttPort ...
func GetGoogleSttPort() string {
	return SglConfig.values["GOOGLE_STT_PORT"]
}

// GetFileServerHttpHttps ...
func GetFileServerHttpHttps() string {
	return SglConfig.values["FILE_SERVER_HTTP_HTTPS"]
}

// GetFileServerIp ...
func GetFileServerIp() string {
	return SglConfig.values["FILE_SERVER_IP"]
}

// GetFileServerGrpcPort ...
func GetFileServerGrpcPort() string {
	return SglConfig.values["FILE_SERVER_GRPC_PORT"]
}

// GetAuthServerGrpcIp ...
func GetAuthServerIp() string {
	return SglConfig.values["AUTH_SERVER_IP"]
}

// GetAuthServerGrpcPort ...
func GetAuthServerGrpcPort() string {
	return SglConfig.values["AUTH_SERVER_GRPC_PORT"]
}

// GetAuthServerGwPort ...
func GetAuthServerGwPort() string {
	return SglConfig.values["AUTH_SERVER_GW_PORT"]
}

// GetCKey ...
func GetCKey() string {
	return SglConfig.values["CKEY"]
}

// GetCIVKey ...
func GetCIVKey() string {
	return SglConfig.values["CIVKEY"]
}

// GetLogLevel ...
func GetLogLevel() string {
	return SglConfig.values["LOG_LEVEL"]
}

// GetLogUseConsolelog ...
func IsUseConsolelog() string {
	return SglConfig.values["LOG_USE_CONSOLELOG"]
}

// GetLogUseFilelog ...
func IsUseFilelog() string {
	return SglConfig.values["LOG_USE_FILELOG"]
}

// GetLogFilename ...
func GetLogFilename() string {
	return SglConfig.values["LOG_FILENAME"]
}

// GetLogFilenameLink ...
func GetLogFilenameLink() string {
	return SglConfig.values["LOG_FILENAME_LINK"]
}

// GetJwtTtl ...
func GetJwtTtl() string {
	return SglConfig.values["JWT_TTL"]
}

// GetCertFile ...
func GetCertFile() string {
	return SglConfig.values["CERT_FILE"]
}

// GetKeyFile ...
func GetKeyFile() string {
	return SglConfig.values["KEY_FILE"]
}

// GetAccessPrivatePem ...
func GetAccessPrivatePem() string {
	return SglConfig.values["ACCESS_PRIVATE_PEM"]
}

// GetAccessPublicPem ...
func GetAccessPublicPem() string {
	return SglConfig.values["ACCESS_PUBLIC_PEM"]
}

// GetRefreshPrivatePem ...
func GetRefreshPrivatePem() string {
	return SglConfig.values["REFRESH_PRIVATE_PEM"]
}

// GetRefreshPublicPem ...
func GetRefreshPublicPem() string {
	return SglConfig.values["REFRESH_PUBLIC_PEM"]
}
