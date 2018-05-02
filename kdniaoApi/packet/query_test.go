package packet

import (
    "testing"
    "fmt"
    "github.com/we-express/config"
)

func TestDoQuery(t *testing.T) {
    conf, _ := config.Parse("/home/sherlock/go/src/github.com/we-express/config/config.toml")
    query, err := NewQuery(conf.EBusinessID, conf.AppKey, "540302693641", "ZTO")

    if err != nil {
        t.Error(err)
    }

    result, err := DoQuery(query)

    if err != nil {
        t.Error(err)
    }

    fmt.Println(result)
}
