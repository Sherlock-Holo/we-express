package server

import (
    "net/http"
    "io"
    "github.com/we-express/config"
    "github.com/we-express/api"
    "encoding/json"
    "strconv"
    "log"
    "fmt"
)

var conf config.Config

func query(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    order := query.Get("order")

    if order == "" {
        w.WriteHeader(http.StatusBadRequest)
        io.WriteString(w, "error order")
        return
    }

    response, err := api.Query(order, conf.ID)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    resp := Response{}

    resp.Records = make([]Record, 0)

    switch response.Status {
    case "0", "2":
        resp.Status = false

        bytes, err := json.Marshal(resp)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.Write(bytes)

    case "1":
        resp.Status = true
        i, err := strconv.Atoi(response.State)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        resp.StateInfo = api.Statuses[response.Status]
        resp.State = i

        for _, data := range response.Data {
            record := Record{}

            jTime := newJTime(data.Time)
            record.Time = jTime
            record.Info = data.Context

            resp.Records = append(resp.Records, record)
        }

        bytes, err := json.Marshal(resp)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.Write(bytes)

    default:
        w.WriteHeader(http.StatusInternalServerError)
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

    log.Fatal(http.ListenAndServe(address, nil))
}