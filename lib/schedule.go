package lib

import (
	"time"

	"github.com/robfig/cron"
)

var Scheduler = cron.NewWithLocation(time.Local)

func InitSchedule() {
	Scheduler.Start()
}
