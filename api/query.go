package api

import (
    "net/http"
    "fmt"
    "encoding/json"
)

const queryUrl = "http://q.kdpt.net/api?id=%s&nu=%s&com=auto&show=json"

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
