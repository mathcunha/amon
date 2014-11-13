package amon

import (
	"testing"
)

func TestMonitor(t *testing.T) {
	wg, err := Monitor("/vagrant/config.json")
	if err != nil {
		t.Errorf("config error")
	}
	wg.Wait()
}
