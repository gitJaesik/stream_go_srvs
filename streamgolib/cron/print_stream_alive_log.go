package cron

import (
	"github.com/gitJaesik/stream_go_srvs/streamgolib"
	"github.com/gitJaesik/stream_go_srvs/streamgolib/logger"
	"github.com/go-co-op/gocron"
)

// type cronFunc func(ctx context.Context) (interface{}, error)

func PrintStreamAliveLog(s *gocron.Scheduler) {
	_, err := s.Do(func() {
		logger.Logger.Infow("PrintStreamAliveLog", streamgolib.ProductName(), "still alive")
	})

	if err != nil {
		logger.Logger.Error("PrintStreamAliveLog", "cron builder", "error")
	}

	// scheduler starts running jobs and current thread continues to execute
	s.StartAsync()
}
