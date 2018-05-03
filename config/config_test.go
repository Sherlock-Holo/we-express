package config

import (
	"testing"
    "fmt"
)

func TestParse(t *testing.T) {
	conf, err := Parse("/home/sherlock/go/src/github.com/we-express/.idea/config.toml")

    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(conf)
}
