package main

import (
	"kelp/log"
	"testing"
)

func TestProbe(t *testing.T) {

	log.AddLogger(
		"probe_test.log",
		"/opt/repo/go/probe",
		2,
		10,
		5, 0,
	)

	probe("http://probe.mapleque.com")
	format("/opt/repo/go/probe/probe_test.log", "/opt/repo/go/probe/probe_test.json")
}
