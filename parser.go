package amon

import (
	"bytes"
	"encoding/json"
	"regexp"
)

func loadStatus(data []byte, status *[]Status) error {
	return json.NewDecoder(bytes.NewBuffer(data)).Decode(status)
}

func FindPattern(patterns []*regexp.Regexp, body string) map[string]string {
	mp := map[string]string{}

	for _, re := range patterns {
		findPattern(re, body, &mp)
	}

	return mp
}

func findPattern(re *regexp.Regexp, body string, mp *map[string]string) {
	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch(body, -1)
	mapa := *mp
	if len(r2) > 0 {
		for i, n := range r2[0] {
			mapa[n1[i]] = n
		}
	}
}
