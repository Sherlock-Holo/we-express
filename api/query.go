package api

import (
    "net/http"
    "fmt"
    "encoding/json"
)

const queryUrl = "http://q.kdpt.net/api"

var ComCode = map[string]string{
    "shunfeng":  "顺丰",
    "huitong":   "百世汇通",
    "zhongtong": "中通快递",
    "shentong":  "申通快递",
    "yuantong":  "圆通速递",
    "yunda":     "韵达",
    "chinapost": "邮政",
    "ems":       "EMS",
    "tiantian":  "天天",
    "jingdong":  "京东",
    "yousu":     "优速",
    "debang":    "德邦",
    "kjkd":      "快捷",
    "zjs":       "宅急送",
    "tnt":       "TNT",
    "ups":       "UPS",
    "dhl":       "DHL",
    "fedex":     "FEDEX联邦",
}

func Query(order, id, com string) (Response, error) {
    request, err := http.NewRequest(http.MethodGet, queryUrl, nil)

    if err != nil {
        return Response{}, fmt.Errorf("query failed")
    }

    query := request.URL.Query()

    query.Add("show", "json")
    query.Add("id", id)
    query.Add("nu", order)

    if com == "" {
        query.Add("com", "auto")
    } else {
        query.Add("com", com)
    }

    request.URL.RawQuery = query.Encode()

    resp, err := http.DefaultClient.Do(request)

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
