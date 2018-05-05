package api

import (
	"testing"
	"github.com/Sherlock-Holo/we-express/config"
    "fmt"
)

func TestQuery(t *testing.T) {
	conf, err := config.Parse("/home/sherlock/go/src/github.com/Sherlock-Holo/we-express/.idea/config.toml")

    if err != nil {
        t.Error(err)
        return
    }

    response, err := Query("619787212452", conf.ID, "shufeng")

    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(response)
}
