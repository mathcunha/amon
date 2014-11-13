package amon

import (
	"regexp"
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
}
