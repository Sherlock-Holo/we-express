package shipper

import (
    "encoding/json"
    "fmt"
)

type Request struct {
    LogisticCode string

    requestType string
    dataType    string

    jsonString string
}

func NewRequest(order string) Request {
    return Request{
        LogisticCode: order,
        requestType:  "2002",
        dataType:     "2",
    }
}

func (r *Request) RequestData() (string, error) {
    bytes, err := json.Marshal(r)

    if err != nil {
        return "", fmt.Errorf("encode query failed")
    }

    r.jsonString = string(bytes)

    return r.jsonString, nil
}
