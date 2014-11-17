package amon

import (
	"regexp"
	"strings"
)

type Extended struct {
	CPULoad     float32
	ReqPerSec   float32
	BytesPerSec float32
	BytesPerReq float32
	BusyWorkers int32
	IdleWorkers int32
	ScoreBoard  string
	HostName    string
	Waiting     int
	Starting    int
	Reading     int
	Replying    int
	Keepalive   int
	Dns         int
	Closing     int
	Logging     int
	Grace       int
	Idle        int
	Available   int
}

func (e *Extended) BuildEvents(event *Event) *[]Event {
	event.Field = e
	events := []Event{*event}
	return &events
}

func (e Extended) Patterns() []*regexp.Regexp {
	return []*regexp.Regexp{regexp.MustCompile("<h1>.*for (?P<host_name>.*)</h1>"), regexp.MustCompile("<dt>.*- (?P<cpu_load>[0-9.]+)% CPU load</dt>"), regexp.MustCompile("<dt>(?P<req_sec>[0-9.]+) requests/sec - (?P<kB_sec>[0-9.]+) kB/second - (?P<B_request>[0-9.]+) .+request</dt>"), regexp.MustCompile("<dt>(?P<requests>[0-9.]+) .+, (?P<idle_workers>[0-9.]+) .+</dt>"), regexp.MustCompile("<pre>(?P<score_board>(.|\n)*)</pre>")}
}

func (e *Extended) Load(mp *map[string]string) {
	mapa := *mp
	e.CPULoad = ParseFloat(mapa["cpu_load"])
	e.ReqPerSec = ParseFloat(mapa["req_sec"])
	e.BytesPerSec = ParseFloat(mapa["kB_sec"])
	e.BytesPerReq = ParseFloat(mapa["B_request"])
	e.BusyWorkers = ParseInt(mapa["requests"])
	e.BusyWorkers = ParseInt(mapa["idle_workers"])
	e.ScoreBoard = mapa["score_board"]
	e.HostName = mapa["host_name"]
	e.Waiting = strings.Count(e.ScoreBoard, "_")
	e.Starting = strings.Count(e.ScoreBoard, "S")
	e.Reading = strings.Count(e.ScoreBoard, "R")
	e.Replying = strings.Count(e.ScoreBoard, "W")
	e.Keepalive = strings.Count(e.ScoreBoard, "K")
	e.Dns = strings.Count(e.ScoreBoard, "D")
	e.Closing = strings.Count(e.ScoreBoard, "C")
	e.Logging = strings.Count(e.ScoreBoard, "L")
	e.Grace = strings.Count(e.ScoreBoard, "G")
	e.Idle = strings.Count(e.ScoreBoard, "I")
	e.Available = strings.Count(e.ScoreBoard, ".")
}
