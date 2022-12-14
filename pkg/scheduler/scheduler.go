package scheduler

import (
	"time"

	"go.uber.org/zap"
)

var logger = zap.Must(zap.NewDevelopment())

func ScheduleTask(task func(), d time.Duration) (stop func()) {
	logger.Info("Scheduling periodic task... :", zap.Duration("delay", d))
	timer := time.AfterFunc(d, task)
	return func() { timer.Stop() }
}

func SchedulePeriodicTask(task func(), d time.Duration) (stop func()) {
	logger.Info("Scheduling periodic task... :", zap.Duration("delay", d))
	ticker := time.NewTicker(d)
	done := make(chan struct{}, 1)
	go func(ticker *time.Ticker, task func(), done <-chan struct{}) {
		for {
			select {
			case <-ticker.C:
				task()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}(ticker, task, done)
	return func() { done <- struct{}{} }
}
