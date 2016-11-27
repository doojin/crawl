package scheduler

import (
	"github.com/doojin/crawl/queue"
	"github.com/jasonlvhit/gocron"
)

// Scheduler schedules job executions
type Scheduler struct {
	s *gocron.Scheduler
}

// New returns new instance of scheduler
func New() *Scheduler {
	return &Scheduler{
		s: new(gocron.Scheduler),
	}
}

// RunEverySecond runs job every second
func (scheduler *Scheduler) RunEverySecond(job queue.Job) {
	scheduler.schedule(scheduler.s.Every(1).Seconds(), job)
}

// RunEveryMinute runs job every minute
func (scheduler *Scheduler) RunEveryMinute(job queue.Job) {
	scheduler.schedule(scheduler.s.Every(1).Minutes(), job)
}

// RunEveryHour runs job every hour
func (scheduler *Scheduler) RunEveryHour(job queue.Job) {
	scheduler.schedule(scheduler.s.Every(1).Hours(), job)
}

// RunEveryNSeconds runs job every N seconds
func (scheduler *Scheduler) RunEveryNSeconds(job queue.Job, n uint64) {
	scheduler.schedule(scheduler.s.Every(n).Seconds(), job)
}

// RunEveryNMinutes runs job every N minutes
func (scheduler *Scheduler) RunEveryNMinutes(job queue.Job, n uint64) {
	scheduler.schedule(scheduler.s.Every(n).Minutes(), job)
}

// RunEveryNHours runs job every N hours
func (scheduler *Scheduler) RunEveryNHours(job queue.Job, n uint64) {
	scheduler.schedule(scheduler.s.Every(n).Hours(), job)
}

// Start starts configured scheduler
func (scheduler *Scheduler) Start() {
	go func() {
		<-scheduler.s.Start()
	}()
}

func (scheduler *Scheduler) schedule(cronJob *gocron.Job, executableJob queue.Job) {
	cronJob.Do(executableJob.Execute)
}
