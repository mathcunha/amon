package scheduler

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

type Task interface {
	Run()
	Interval() string
}

func ScheduleOne(t Task) (*time.Ticker, chan bool) {
	log.Printf("tasks %v", t)
	duration := getInterval(t.Interval())

	if duration > 0 {
		ticker := time.NewTicker(duration)
		control := make(chan bool)
		go schedule(ticker, control, t)
		return ticker, control
	}
	return nil, nil
}

func Schedule(tasks []Task) []*time.Ticker {
	length := len(tasks)
	tickers := make([]*time.Ticker, length, length)
	controls := make([]chan bool, length, length)

	for i, t := range tasks {
		ticker, control := ScheduleOne(t)
		tickers[i] = ticker
		controls[i] = control
	}
	return tickers
}

func schedule(t *time.Ticker, q chan bool, task Task) {
	task.Run()
	for {
		select {
		case <-t.C:
			task.Run()
		case <-q:
			t.Stop()
			return
		}
	}
}

func getInterval(interval string) time.Duration {
	log.Printf("interval %v", interval)
	nPattern := "^[0-9]*"
	dPattern := "[hms]$"

	if matched, _ := regexp.MatchString(nPattern+dPattern, interval); matched {
		re := regexp.MustCompile(nPattern)
		num, _ := strconv.Atoi(re.FindString(interval))

		re = regexp.MustCompile(dPattern)
		duration := re.FindString(interval)

		log.Printf("Num = %v - Duration = %v", num, duration)

		switch {
		case "h" == duration:
			return time.Duration(num) * time.Hour
		case "m" == duration:
			return time.Duration(num) * time.Minute
		case "s" == duration:
			return time.Duration(num) * time.Second
		}
	}

	return -1
}
