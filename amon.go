package amon

import (
	"github.com/mathcunha/amon/scheduler"
	"io/ioutil"
	"log"
	"sync"
)

type Status struct {
	Stype     string `json:"type"`
	Url       string
	Intervalo string `json:"Interval"`
}

func (s Status) Run() {
	body, _ := s.GetResource()

	log.Println(body)
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
