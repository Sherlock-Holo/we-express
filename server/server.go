package server

import (
    "net/http"
    "github.com/we-express/config"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "io"
)

var conf config.Config

func queryHandle(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    order := query.Get("order")
    //deviceID := query.Get("deviceID")

    response, err := QueryExpress(conf.EBusinessID, conf.AppKey, order)

    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        io.WriteString(w, err.Error())
        return
    }

    bytes, err := json.Marshal(response)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(bytes)
}

func Start(addr string, port int, configFile string) {
    var err error
    conf, err = config.Parse(configFile)

    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/express", queryHandle)
    address := fmt.Sprintf("%s:%d", addr, port)
    fmt.Fprintf(os.Stderr, "bind address: %s\n", address)
    log.Fatal(http.ListenAndServe(address, nil))
}
