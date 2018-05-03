package api

import (
    "net/http"
    "fmt"
    "encoding/json"
)

const queryUrl = "http://q.kdpt.net/api?id=%s&nu=%s&com=auto&show=json&format=kuaidi100"

func Query(order, id string) (Response, error) {
    resp, err := http.Get(fmt.Sprintf(queryUrl, id, order))

    if err != nil {
        return Response{}, fmt.Errorf("query failed")
    }

    decoder := json.NewDecoder(resp.Body)

    response := Response{}

    err = decoder.Decode(&response)

    if err != nil {
        return Response{}, fmt.Errorf("data format error")
    }

    return response, nil
}
