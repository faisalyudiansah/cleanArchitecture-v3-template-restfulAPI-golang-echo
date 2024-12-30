package scheduler_test

import (
	"fmt"
	"server/config"
	"server/internal/infrastructure/logger"
	"server/internal/infrastructure/scheduler"
	"sync"
	"testing"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var (
	lo         *logger.Logger
	slo        *scheduler.SchedulerLogger
	loggerOnce sync.Once
	sloOnce    sync.Once
	sOpts      *scheduler.SchedulerOption
	sOptsOnce  sync.Once
)

func getLogger(t testing.TB) *logger.Logger {
	var err error

	loggerOnce.Do(func() {
		lo, err = config.NewLogger(&config.LoggerConfig{
			Logger:   logrus.New(),
			Output:   "./../../../public/logs",
			FileName: "scheduler",
		})
	})

	require.NoError(t, err)

	return lo
}

func GetSchedulerLogger(t testing.TB) *scheduler.SchedulerLogger {
	lo := getLogger(t)

	sloOnce.Do(func() {
		slo = &scheduler.SchedulerLogger{
			Logger: lo,
		}
	})

	return slo
}

func GetSchedulerOption(t testing.TB) *scheduler.SchedulerOption {
	logger := GetSchedulerLogger(t)
	sOptsOnce.Do(func() {
		sOpts = &scheduler.SchedulerOption{
			Timezone: "Asia/Jakarta",
			Logger:   logger,
		}
	})
	return sOpts
}

func GetScheduler(t testing.TB) gocron.Scheduler {
	opts := GetSchedulerOption(t)
	s, err := scheduler.NewScheduler(opts)
	require.NoError(t, err)
	return s
}

// TestScheduler ensure scheduler running as expected.
func TestSchedulerDurationJob(t *testing.T) {
	opts := GetSchedulerOption(t)
	s := GetScheduler(t)

	// defer func() {
	// 	opts.Logger.Info("shutting down the duration job scheduler")
	// 	_ = s.Shutdown()
	// }()

	jTotal := 3
	jCount := 1

	job, err := s.NewJob(
		gocron.DurationJob(time.Duration(1)*time.Minute),
		gocron.NewTask(func() {
			opts.Logger.Info(fmt.Sprintf("[%d/%d] duration job scheduler is running", jCount, jTotal))
			jCount++
		}),
		gocron.WithName("test_scheduler_duration"),
	)
	require.NoError(t, err)

	require.NotEmpty(t, job.ID())

	s.Start()

	time.Sleep(time.Duration(jTotal) * time.Second)
}

func TestSchedulerCronTab(t *testing.T) {
	opts := GetSchedulerOption(t)
	s := GetScheduler(t)

	defer func() {
		opts.Logger.Info("shutting down the crontab scheduler")
		_ = s.Shutdown()
	}()

	jTotal := 2
	jCount := 1

	job, err := s.NewJob(
		gocron.CronJob(
			"* * * * *",
			true),
		gocron.NewTask(func() {
			opts.Logger.Info(fmt.Sprintf("[%d/%d] crontab scheduler is running", jCount, jTotal))
			jCount++
		}),
		gocron.WithName("test_scheduler_crontab"),
	)

	require.NoError(t, err)

	s.Start()

	require.NotEmpty(t, job.ID())

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
	}

	time.Sleep(time.Duration(jTotal) * time.Minute)
}
