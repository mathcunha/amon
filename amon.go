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

func (t Status) Evaluate() {
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
	scheduler.Schedule(s)
	return err
}
