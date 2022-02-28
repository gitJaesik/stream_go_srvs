package logger

import "fmt"

var (
	envs []string
)

type StreamGoLibLoggerConfig struct {
	values map[string]string
	Err    error
}

func init() {
	fmt.Println("logger.go init start")
	InitializeZap()
	fmt.Println("logger.go init end")
}
