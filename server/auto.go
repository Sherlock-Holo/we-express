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

/*func autoUpdate(db *db.ExpressDB, apiID string) error {
    rows, err := db.ListExpress()

    if err != nil {
        return err
    }

    for rows.Next() {
        var (
            order string
            com   string
        )

        err = rows.Scan(&order, &com)
        if err != nil {
            break
        }

        // count++
        _, err = db.Update(order, com, apiID)
        if err != nil {
            return err
        }
    }

    return nil
}*/
