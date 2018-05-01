package main

import (
	"testing"

	"github.com/we-express/config"
	"fmt"
)

func TestQueryExpress(t *testing.T) {
	conf, err := config.Parse("/home/sherlock/go/src/github.com/we-express/config/config.json")

	if err != nil {
		t.Error(err)
		return
	}

	status, err := QueryExpress(conf.EBusinessID, conf.AppKey, "540302693641")

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(status)
}
