package amon

import (
	"github.com/mathcunha/amon/scheduler"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Config struct {
	Elasticsearch string    `json:"elasticsearch"`
	Status        []*Status `json:"status"`
}

type Status struct {
	Stype     string `json:"type"`
	Url       string
	Intervalo string `json:"Interval"`
	Config    *Config
}

type Event struct {
	Source    string      `json:"@source"`
	Tags      []string    `json:"@tags"`
	Timestamp string      `json:"@timestamp"`
	Field     interface{} `json:"@fields"`
}

func (s *Status) Run() {
	body, _ := s.GetResource()

	var res Resource

	switch s.Stype {
	case "extended":
		res = new(Extended)
	case "jk":
		res = new(Jk)
	}

	LoadAttributes(res, body)
	event := new(Event)
	event.Source = s.Url
	event.Tags = []string{s.Stype}
	event.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z0700")
	PostEvents(s.Config.Elasticsearch, res.BuildEvents(event))
}

func (t Status) Interval() string {
	return t.Intervalo
}

func Monitor(filepath string) (sync.WaitGroup, error) {
	body, err := ioutil.ReadFile(filepath)
	var wg sync.WaitGroup

	if err != nil {
		log.Printf("error loading config file (%v)", err)
		return wg, err
	}
	var c Config
	err = loadStatus(body, &c)
	routines := scheduler.Schedule(tasks(&c))
	wg.Add(len(routines))

	return wg, err
}

func tasks(c *Config) []scheduler.Task {
	s := c.Status
	vals := make([]scheduler.Task, len(s))
	for i := 0; i < len(s); i++ {
		vals[i] = s[i]
		s[i].Config = c
	}
	return vals
}
