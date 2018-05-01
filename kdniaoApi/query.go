package kdniaoApi

import (
    "net/url"
)

type Query struct {
    RequestData string
    EBusinessID string
    RequestType string
    DataSign    string
    DataType    string
}

func (q *Query) Encode() url.Values {
    return url.Values{
        "RequestData": []string{q.RequestData},
        "EBusinessID": []string{q.EBusinessID},
        "RequestType": []string{q.RequestType},
        "DataSign":    []string{q.DataSign},
        "DataType":    []string{q.DataType},
    }
}
