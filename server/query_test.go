package server

import (
    "testing"

    "github.com/we-express/config"
    "fmt"
    "encoding/json"
)

func TestQueryExpress(t *testing.T) {
    conf, err := config.Parse("/home/sherlock/go/src/github.com/we-express/config/config.toml")

    if err != nil {
        t.Error(err)
        return
    }

    resp, err := QueryExpress(conf.EBusinessID, conf.AppKey, "540302693641")

    if err != nil {
        t.Error(err)
        return
    }

    bytes, err := json.MarshalIndent(resp, "", "    ")

    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(string(bytes))
}
