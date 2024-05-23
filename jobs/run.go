package jobs

import (
	"context"
	"web/jobs/checker"
)

func RunJob(ctx context.Context) {
	go checker.CheckerJob(ctx)
}
