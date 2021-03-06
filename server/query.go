package server

import (
    "compress/gzip"
    "database/sql"
    "io"
    "log"
    "net/http"

    "github.com/Sherlock-Holo/we-express/config"
    "github.com/Sherlock-Holo/we-express/db"
)

var (
    conf      config.Config
    expressDB *db.ExpressDB
)

func query(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    order := query.Get("order")
    com := query.Get("com")
    force := query.Get("force")

    if order == "" {
        w.WriteHeader(http.StatusBadRequest)
        io.WriteString(w, "order error")
        log.Println("error order")
        return
    }

    var (
        jsonString string
        err        error
    )

    if force != "" {
        log.Printf("force refresh, order %s, com %s", order, com)
        jsonString, err = expressDB.Update(order, com, conf.ID)
    } else {
        jsonString, err = expressDB.Query(order, com)
    }

    switch err {
    case db.Timeout, sql.ErrNoRows:
        log.Println(err, "refresh from api")

        jsonString, err = expressDB.Update(order, com, conf.ID)

        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

    case nil:

    default:
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Add("Content-Encoding", "gzip")
    w.Header().Set("Content-type", "application/json")

    gzw := gzip.NewWriter(w)

    gzw.Write([]byte(jsonString))
    gzw.Close()
}
