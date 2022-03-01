package cron

import (
	"time"

	"github.com/gitJaesik/stream_go_srvs/streamgolib/config"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"github.com/go-co-op/gocron"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	InitializeCron,
)

func InitializeCron(sglConfig *config.SGLConfig) *gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)
	logger.Logger.Infow("cron", "sglConfig.GetCronSeconds()", sglConfig.GetCronSeconds())
	s.Every(sglConfig.GetCronSeconds()).Second()

	return s
}
