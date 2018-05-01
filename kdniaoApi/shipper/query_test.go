package shipper

import (
	"testing"

    "fmt"
)

func TestDoQuery(t *testing.T) {
	query, err := NewQuery("1338036", "d113aa98-addf-4188-91e7-d3deb5dc64dd", "540302693641")

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
