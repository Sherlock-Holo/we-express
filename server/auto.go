package server

import (
    "errors"

    "github.com/Sherlock-Holo/we-express/db"
)

func autoUpdate(db *db.ExpressDB, apiID string) error {
    rows, err := db.ListExpress()

    if err != nil {
        return err
    }

    var (
        check = make(chan bool)
        count int
    )

    for rows.Next() {
        var (
            order string
            com   string
        )

        err = rows.Scan(&order, &com)
        if err != nil {
            break
        }

        count++
        go func(order, com string) {
            _, err = db.Update(order, com, apiID)
            if err != nil {
                check <- false
            } else {
                check <- true
            }
        }(order, com)
    }

    success := true

    for ; count > 0; count-- {
        if !(<-check) {
            success = false
        }
    }

    if !success {
        return errors.New("update data error")
    } else {
        return nil
    }
}
