package server

import (
    "net/http"
    "io"
    "github.com/Sherlock-Holo/we-express/config"
    "github.com/Sherlock-Holo/we-express/api"
    "encoding/json"
    "strconv"
    "log"
    "fmt"
)

var conf config.Config

func query(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    order := query.Get("order")
    com := query.Get("com")

    if order == "" {
        w.WriteHeader(http.StatusBadRequest)
        io.WriteString(w, "order error")
        log.Println("error order")
        return
    }

    response, err := api.Query(order, conf.ID, com)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        log.Println(err)
        return
    }

    resp := Response{}

    resp.Records = make([]Record, 0)

    switch response.Status {

    // 0: 暂无结果
    // 2: 接口出现异常
    case "0", "2":
        resp.Status = false

        bytes, err := json.Marshal(resp)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }

        w.Header().Set("Content-type", "application/json")
        w.Write(bytes)

    case "1":
        resp.Status = true
        i, err := strconv.Atoi(response.State)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }

        resp.Order = order

        resp.StateInfo = api.Statuses[response.Status]
        resp.State = i

        for _, data := range response.Data {
            record := Record{}

            jTime := newJTime(data.Time)
            record.Time = jTime
            record.Info = data.Context

            resp.Records = append(resp.Records, record)
        }

        com, ok := api.ComCode[response.Com]

        if !ok {
            resp.Com = response.Com
        } else {
            resp.Com = com
        }

        bytes, err := json.Marshal(resp)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println(err)
            return
        }

        w.Header().Set("Content-type", "application/json")
        w.Write(bytes)

    default:
        w.WriteHeader(http.StatusInternalServerError)
        log.Println(err)
        return
    }
}

func Start(configFile, addr string, port uint) {
    var err error
    conf, err = config.Parse(configFile)

    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/express", query)

    address := fmt.Sprintf("%s:%d", addr, port)

    log.Printf("listen on %s\n", address)

    log.Fatal(http.ListenAndServe(address, nil))
}
