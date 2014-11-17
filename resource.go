package amon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Resource interface {
	Patterns() []*regexp.Regexp
	Load(*map[string]string)
	BuildEvents(e *Event) *[]Event
}

func (s Status) GetResource() ([]byte, error) {
	url := s.Url
	res, err := http.Get(url)
	if err != nil {
		log.Printf("unable do reach the url %v", err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("error reading resource %v body - %v", url, err)
	}

	return data, nil
}

func PostEvents(url string, events *[]Event) {
	timestamp := time.Now().Format("2006.01.02")
	for _, value := range *events {
		var postData []byte
		w := bytes.NewBuffer(postData)
		json.NewEncoder(w).Encode(value)
		resp, _ := http.Post(url+"/logstash-"+timestamp+"/amon/", "application/json", w)
		resp.Body.Close()
	}
}
