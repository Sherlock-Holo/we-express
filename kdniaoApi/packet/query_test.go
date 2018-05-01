package packet

import (
    "testing"
    "fmt"
)

func TestDoQuery(t *testing.T) {
    query, err := NewQuery("1338036", "d113aa98-addf-4188-91e7-d3deb5dc64dd", "540302693641", "ZTO")

    if err != nil {
        t.Error(err)
    }

    result, err := DoQuery(query)

    if err != nil {
        t.Error(err)
    }

    fmt.Println(result)
}
