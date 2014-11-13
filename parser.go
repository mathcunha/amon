package amon

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
)

func loadStatus(data []byte, status *[]Status) error {
	return json.NewDecoder(bytes.NewBuffer(data)).Decode(status)
}

func FindPattern(r Resource, body []byte) *map[string]string {
	mp := map[string]string{}

	for _, re := range r.Patterns() {
		findPattern(re, body, &mp)
	}

	return &mp
}

func findPattern(re *regexp.Regexp, body []byte, mp *map[string]string) {
	n1 := re.SubexpNames()
	r2 := re.FindSubmatch(body)
	mapa := *mp
	if len(r2) > 0 {
		for i, n := range r2 {
			mapa[n1[i]] = string(n)
		}
	}
}

func ParseFloat(value string) float32 {
	f_value, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Printf("error parsing (%v) to float32 (%v)\n", value, err)
	} else {
		return -1
	}
	return float32(f_value)
}

func ParseInt(value string) int32 {
	i_value, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Printf("error parsing (%v) to int32 (%v)\n", value, err)
	} else {
		return -1
	}
	return int32(i_value)
}
