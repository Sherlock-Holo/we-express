package shipper

import (
	"testing"

    "fmt"
    "github.com/we-express/config"
    "log"
)

func TestDoQuery(t *testing.T) {
    conf, err := config.Parse("/home/sherlock/go/src/github.com/we-express/config/config.toml")

    if err != nil {
        log.Fatal(err)
    }

	query, err := NewQuery(conf.EBusinessID, conf.AppKey, "540302693641")

    if err != nil {
        t.Error(err)
        return
    }

    status, err := DoQuery(query)

    if err != nil {
        t.Error(err)
        return
    }

    if !status.Success {
        t.Error("not success")
    }

    fmt.Println(status.ShipperCode())
}
