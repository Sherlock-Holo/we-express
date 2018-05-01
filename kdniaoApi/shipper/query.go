package shipper

import (
    "github.com/we-express/kdniaoApi"
    "crypto/md5"
    "encoding/hex"
    "encoding/base64"
    "net/http"
    "encoding/json"
    "github.com/we-express/config"
)

func NewQuery(id, appKey, order string) (kdniaoApi.Query, error) {
    request := NewRequest(order)

    requestData, err := request.RequestData()

    query := kdniaoApi.Query{}

    if err != nil {
        return query, err
    }

    query.RequestData = requestData
    query.EBusinessID = id
    query.RequestType = request.requestType
    query.DataType = request.dataType
    query.DataSign = sign(request, appKey)

    return query, nil
}

func DoQuery(query kdniaoApi.Query) (Result, error) {
    resp, err := http.PostForm(config.QueryUrl, query.Encode())

    status := Result{}

    if err != nil {
        return status, err
    }

    decoder := json.NewDecoder(resp.Body)

    err = decoder.Decode(&status)

    if err != nil {
        return Result{}, err
    }

    return status, nil
}

func sign(r Request, appKey string) string {
    sum := md5.Sum([]byte(r.jsonString + appKey))

    hexBytes := make([]byte, hex.EncodedLen(len(sum)))

    hex.Encode(hexBytes, sum[:])

    return base64.URLEncoding.EncodeToString(hexBytes)
}
