package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"kelp/log"
	"kelp/util"
)

func main() {
	log.AddLogger(
		"probe.log",
		"/opt/repo/go/probe",
		2,
		10000000,
		5, 0,
	)
	ticker := time.NewTicker(time.Minute * 1)
	for _ = range ticker.C {
		probe("http://probe.mapleque.com")
		format("/opt/repo/go/probe/probe.log", "opt/repo/go/probe/probe.json")
	}
}

func probe(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return
	}
	html, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Error(err)
		return
	}
	if string(html) == "This is a Probe" {
		log.Info(int(time.Now().Sub(start)) / int(time.Millisecond))
	} else {
		log.Error(html)
	}
}

func format(filein, fileout string) {
	data := []string{}
	for _, line := range util.ReadFile(filein) {
		arr := strings.Split(line, " ")
		if len(arr) != 4 || arr[2] != "[INFO]" {
			data = append(data, `"`+arr[0]+" "+arr[1]+`":"0ms"`)
		} else {
			data = append(data, `"`+arr[0]+" "+arr[1]+`":"`+arr[3]+`"`)
		}
	}
	body := "{" + strings.Join(data, ",") + "}"
	json, err := os.OpenFile(
		fileout,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err)
	}
	defer json.Close()
	json.WriteString(body)
}
