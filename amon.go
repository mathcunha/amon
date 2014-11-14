package amon

import (
	"github.com/mathcunha/amon/scheduler"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Status struct {
	Stype     string `json:"type"`
	Url       string
	Intervalo string `json:"Interval"`
}

type Event struct {
	Source    string      `json:"@source"`
	Tags      []string    `json:"@tags"`
	Timestamp string      `json:"@timestamp"`
	Field     interface{} `json:"@fields"`
}

func (s Status) Run() {
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
	PostEvents(res.BuildEvents(event))
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
	var s []Status
	err = loadStatus(body, &s)
	routines := scheduler.Schedule(tasks(s))
	wg.Add(len(routines))

	return wg, err
}

func tasks(s []Status) []scheduler.Task {
	vals := make([]scheduler.Task, len(s))
	for i, v := range s {
		vals[i] = v
	}
	return vals
}
