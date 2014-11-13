package amon

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Resource interface {
	Patterns() []*regexp.Regexp
	Load(*map[string]string)
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
