package server

import (
    "net/http"
    "io"
    "github.com/Sherlock-Holo/we-express/config"
    "log"
    "github.com/Sherlock-Holo/we-express/db"
    "database/sql"
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
        jsonString, err = expressDB.Update(order, com, conf.ID, expressDB.Check(order))
    } else {
        jsonString, err = expressDB.Query(order, com)
    }

    switch {
    case err == db.Timeout:
        jsonString, err = expressDB.Update(order, com, conf.ID, true)

        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

    case err == sql.ErrNoRows:
        jsonString, err = expressDB.Update(order, com, conf.ID, false)

        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

    case err != nil:
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-type", "application/json")
    io.WriteString(w, jsonString)

}
