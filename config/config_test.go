package config

import (
	"testing"
    "fmt"
)

func TestParse(t *testing.T) {
	conf, err := Parse("/home/sherlock/go/src/github.com/we-express/config/config.json")

    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(conf)
}
