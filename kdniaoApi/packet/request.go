package packet

import "encoding/json"

type Request struct {
    OrderCode    string
    ShipperCode  string
    LogisticCode string

    requestType string
    dataType    string

    jsonString string
}

func NewRequest(order, shipper string) Request {
    return Request{
        LogisticCode: order,
        ShipperCode:  shipper,
        requestType:  "1002",
        dataType:     "2",
    }
}

func (r *Request) RequestData() (string, error) {
    bytes, err := json.Marshal(*r)

    if err != nil {
        return "", err
    }

    r.jsonString = string(bytes)

    return r.jsonString, nil
}
