package amon

import (
	"log"
	"regexp"
)

type Jk struct {
	HostName string
	Workers  []Worker
}

type Worker struct {
	Name           string
	HostName       string
	Status         string
	State          string
	ResquestNumber int32
	SessionNumber  int32
	FailedRequest  int32
	ClientError    int32
	ReplyError     int32
	Busy           int32
	MaxBusy        int32
	BackendCon     int32
	LastError      string
}

func (j Jk) Patterns() []*regexp.Regexp {
	log.Println("jk patterns")
	return []*regexp.Regexp{regexp.MustCompile("<tr><td>(?P<worker>.+)</td><td>.+</td><td>(?P<host_name>.+)</td><td>(?P<ip>[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}):[0-9]+</td>"), regexp.MustCompile("</td><td>(?P<wname>.+)</td><td>(?P<act>.+)</td><td>(?P<state>.+)</td><td>[0-9]+</td><td>[0-9]+</td><td>[0-9]+</td><td>[0-9]+</td><td>(?P<request_num>[0-9]+) .+</td><td>(?P<session_num>[0-9]+) .+</td><td>(?P<failed_request>[0-9]+)</td><td>(?P<client_error>[0-9]+)</td><td>(?P<reply_error>[0-9]+)</td><td>.+</td><td>.+</td><td>(?P<busy>[0-9]+)</td><td>(?P<max_busy>[0-9]+)</td><td>(?P<backend_con>[0-9]+)</td><td>.+</td><td>.+</td><td>.+</td><td>.+</td><td>.+</td><td>(?P<last_error>.+)</td></tr>")}
}

func (j *Jk) Load(mp *map[string]string) {
	keys := []string{"worker", "wname", "host_name", "act", "state", "request_num", "session_num", "failed_request", "client_error", "reply_error", "busy", "max_busy", "backend_con", "last_error"}
	mapa := *mp
	i := 0
	for i = 0; len(mapa[keyName(keys[0], i)]) != 0; i++ {
		if mapa[keyName(keys[0], i)] == mapa[keyName(keys[1], i)] {
			worker := Worker{mapa[keyName(keys[1], i)], mapa[keyName(keys[2], i)], mapa[keyName(keys[3], i)], mapa[keyName(keys[4], i)], ParseInt(mapa[keyName(keys[5], i)]), ParseInt(mapa[keyName(keys[6], i)]), ParseInt(mapa[keyName(keys[7], i)]), ParseInt(mapa[keyName(keys[8], i)]), ParseInt(mapa[keyName(keys[9], i)]), ParseInt(mapa[keyName(keys[10], i)]), ParseInt(mapa[keyName(keys[11], i)]), ParseInt(mapa[keyName(keys[12], i)]), mapa[keyName(keys[13], i)]}
			j.Workers = append(j.Workers, worker)
		} else {
			log.Printf("error worker mismatch - {%v} %v == %v \n", i, mapa[keyName(keys[0], i)], mapa[keyName(keys[1], i)])
		}
	}
	log.Printf("%v workers \n", i)
}

func (jk *Jk) BuildEvents(event *Event) *[]Event {
	events := make([]Event, len(jk.Workers))
	for i, worker := range jk.Workers {
		events[i] = Event{event.Source, []string{"jkstatus"}, event.Timestamp, worker}
	}
	return &events
}
