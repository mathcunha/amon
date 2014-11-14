package amon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func loadStatus(data []byte, status *[]Status) error {
	return json.NewDecoder(bytes.NewBuffer(data)).Decode(status)
}

func LoadAttributes(r Resource, body []byte) {
	mp := map[string]string{}

	for _, re := range r.Patterns() {
		findPattern(re, body, &mp)
	}

	r.Load(&mp)
}

func findPattern(re *regexp.Regexp, body []byte, mp *map[string]string) {
	n1 := re.SubexpNames()
	r3 := re.FindAllSubmatch(body, -1)
	mapa := *mp
	if len(r3) > 0 {
		for j, r2 := range r3 {
			if len(r2) > 0 {
				for i, n := range r2 {
					mapa[keyName(n1[i], j)] = string(n)
				}
			}
		}
	}
}

func keyName(key string, i int) string {
	if len(key) != 0 && i != 0 {
		return fmt.Sprintf("%v_%v", key, i)
	}
	return key
}

func ParseFloat(value string) float32 {
	f_value, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Printf("error parsing (%v) to float32 (%v)\n", value, err)
		return -1
	}
	return float32(f_value)
}

func ParseInt(value string) int32 {
	i_value, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Printf("error parsing (%v) to int32 (%v)\n", value, err)
		return -1
	}
	return int32(i_value)
}
