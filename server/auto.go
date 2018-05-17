package server

import (
    "github.com/Sherlock-Holo/we-express/db"
    "errors"
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
            order        string
            com          string
            recordString string
            updateTime   int64
        )

        err = rows.Scan(&order, &com, &recordString, &updateTime)
        if err != nil {
            break
        }

        count++
        go func() {
            _, err = db.Update(order, com, apiID, true)
            if err != nil {
                check <- false
            } else {
                check <- true
            }
        }()
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
