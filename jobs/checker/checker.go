package checker

import (
	"context"
	"time"

	"web/logger"
)

func CheckerJob(ctx context.Context) {
	// init checker setting
	logger.Info("CheckerJob running now.")

	for {
		select {
		case <-ctx.Done():
			logger.Infof("checker job got exit sign. return now")
			return

		case <-time.After(time.Second * 1):
			// 5 second buffer between range
			checker()
		}
	}
}

func checker() {
	logger.Info("checker running")

}
