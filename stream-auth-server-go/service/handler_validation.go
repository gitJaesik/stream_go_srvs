package service

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
)

func ValidationHandler(ctx context.Context, validate func() error) error {
	name := getCallerName(2)
	if validate != nil {
		err := validate()
		if err != nil {
			logger.Logger.Errorf("%s validation error :%v", name, err)
			return fmt.Errorf("VALIDATION_ERROR:%s", err.Error())
		}
	}
	return nil
}

func getCallerName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	name := details.Name()
	lastDotIndex := strings.LastIndex(name, ".")

	if lastDotIndex == -1 {
		return name
	}
	return name[lastDotIndex+1:]
}
