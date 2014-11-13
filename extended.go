package amon

import (
	"regexp"
)

func Patterns() []*regexp.Regexp {
	patterns := []*regexp.Regexp{regexp.MustCompile("^.*- (?P<cpu_load>[0-9.]+)% CPU load</dt>"), regexp.MustCompile("<dt>(?P<req_sec>[0-9.]+) requests/sec - (?P<kB_sec>[0-9.]+) kB/second - (?P<B_request>[0-9.]+) B/request</dt>"), regexp.MustCompile("<dt>(?P<requests>[0-9.]+) .*, (?P<idle_workers>[0-9.]+) .*</dt>")}
	return patterns
}
