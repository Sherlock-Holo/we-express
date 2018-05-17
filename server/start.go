package server

import (
    "fmt"
    "log"
    "net/http"

    "github.com/Sherlock-Holo/we-express/config"
    "github.com/Sherlock-Holo/we-express/db"
)

func Start(configFile, addr string, port uint) {
    var err error
    conf, err = config.Parse(configFile)

    if err != nil {
        log.Fatal(err)
    }

    expressDB, err = db.Connect(conf.DbUser, conf.DbPassword)

    if err != nil {
        log.Fatal(err)
    }

    if err := autoUpdate(expressDB, conf.ID); err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/express", query)

    address := fmt.Sprintf("%s:%d", addr, port)

    log.Printf("listen on %s\n", address)

    log.Fatal(http.ListenAndServe(address, nil))
}
