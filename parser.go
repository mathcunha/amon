package amon

import (
	"bytes"
	"encoding/json"
)

func loadStatus(data []byte, status *[]Status) error {
	return json.NewDecoder(bytes.NewBuffer(data)).Decode(status)
}
