package packet

import (
    "net/http"
    "github.com/we-express/kdniaoApi"
    "encoding/json"
    "crypto/md5"
    "encoding/hex"
    "encoding/base64"
    "github.com/we-express/config"
    "fmt"
)

func NewQuery(id, appKey, order, shipper string) (kdniaoApi.Query, error) {
    request := NewRequest(order, shipper)

    requestData, err := request.RequestData()

    query := kdniaoApi.Query{}

    if err != nil {
        return query, fmt.Errorf("encode query failed")
    }

    query.RequestData = requestData
    query.EBusinessID = id
    query.RequestType = request.requestType
    query.DataType = request.dataType
    query.DataSign = sign(request, appKey)

    return query, nil
}

func DoQuery(query kdniaoApi.Query) (Status, error) {
    resp, err := http.PostForm(config.QueryUrl, query.Encode())

    status := Status{}

    if err != nil {
        return status, fmt.Errorf("query packet failed")
    }

    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&status)

    return status, nil
}

func sign(r Request, appKey string) string {
    sum := md5.Sum([]byte(r.jsonString + appKey))

    hexBytes := make([]byte, hex.EncodedLen(len(sum)))

    hex.Encode(hexBytes, sum[:])

    return base64.URLEncoding.EncodeToString(hexBytes)
}
