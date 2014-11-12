package amon

import (
	"github.com/mathcunha/amon/scheduler"
	"io/ioutil"
	"log"
)

type Status struct {
	Stype     string `json:"type"`
	Url       string
	Intervalo string `json:"Interval"`
}

func (t Status) Run() {
	log.Printf("TODO")
}

func (t Status) Interval() string {
	return t.Intervalo
}

func Monitor(filepath string) error {
	body, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Printf("error loading config file (%v)", err)
		return err
	}
	var s []Status
	err = loadStatus(body, &s)
	scheduler.Schedule(tasks(s))
	return err
}

func tasks(s []Status) []scheduler.Task {
	vals := make([]scheduler.Task, len(s))
	for i, v := range s {
		vals[i] = v
	}
	return vals
}
