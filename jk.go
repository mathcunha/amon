package amon

import (
	"log"
	"regexp"
)

type Jk struct {
	HostName string
}

func (j Jk) Patterns() []*regexp.Regexp {
	log.Println("jk patterns")
	return []*regexp.Regexp{regexp.MustCompile("<tr><td>(?P<worker>.+)</td><td>.+</td><td>(?P<host_name>.+)</td><td>(?P<ip>[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}):[0-9]+</td>"), regexp.MustCompile("</td><td>(?P<wname>.+)</td><td>(?P<act>.+)</td><td>(?P<state>.+)</td><td>[0-9]+</td><td>[0-9]+</td><td>[0-9]+</td><td>[0-9]+</td><td>(?P<request_num>[0-9]+) .+</td><td>(?P<session_num>[0-9]+) .+</td><td>(?P<failed_request>[0-9]+)</td><td>(?P<client_error>[0-9]+)</td><td>(?P<reply_error>[0-9]+)</td><td>.+</td><td>.+</td><td>(?P<busy>[0-9]+)</td><td>(?P<max_busy>[0-9]+)</td><td>(?P<backend_con>[0-9]+)</td><td>.+</td><td>.+</td><td>.+</td><td>.+</td><td>.+</td><td>(?P<last_error>.+)</td></tr>")}
}

func (j *Jk) Load(mp *map[string]string) {
	mapa := *mp
	log.Printf("mapa - %v \n", mapa)
}
