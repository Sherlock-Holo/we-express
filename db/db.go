package db

// database format
// express_id(varchar 20), express_com(10), express_record(text), update_time(int)

import (
    "errors"
    "database/sql"
    "sync"
    "time"
    "fmt"
    "github.com/Sherlock-Holo/we-express/api"
    "github.com/Sherlock-Holo/we-express/server"
    "strconv"
    "encoding/json"
)

var (
    Timeout = errors.New("database data timeout")
    //NotInDB = errors.New("express not in database")
)

type ExpressDB struct {
    db     *sql.DB
    rwlock sync.RWMutex
}

func (db *ExpressDB) Update(order, com, apiID string, exist bool) (string, error) {
    apiResp, err := api.Query(order, apiID, com)

    db.rwlock.Lock()
    defer db.rwlock.Unlock()

    if err != nil {
        return "", err
    }

    serverResp := server.Response{}

    i, err := strconv.Atoi(apiResp.State)

    if err != nil {
        return "", err
    }

    serverResp.State = i
    serverResp.StateInfo = api.States[apiResp.State]

    switch apiResp.Status {

    // 0: 暂无结果
    // 2: 接口出现异常
    case "0", "2":
        serverResp.Status = false

    case "1":
        serverResp.Status = true
        serverResp.Order = order

        for _, data := range apiResp.Data {
            record := server.Record{}

            jTime := server.NewJTime(data.Time)
            record.Time = jTime
            record.Info = data.Context

            serverResp.Records = append(serverResp.Records, record)
        }

        com, ok := api.ComCode[apiResp.Com]
        if !ok {
            serverResp.Com = apiResp.Com
        } else {
            serverResp.Com = com
        }
    }

    jsonBytes, err := json.Marshal(serverResp)

    if err != nil {
        return "", err
    }

    jsonString := string(jsonBytes)

    unixTime := time.Now().UTC().Unix()

    var stmt *sql.Stmt

    if exist {
        stmt, err = db.db.Prepare("update express set update_time=? express_state=? express_record=?")
    } else {
        stmt, err = db.db.Prepare("insert into express (update_time, express_state, express_record), values (?, ?, ?)")
    }

    if err != nil {
        return "", err
    }

    _, err = stmt.Exec(unixTime, serverResp.State, jsonString)

    if err != nil {
        return "", err
    }

    return jsonString, nil
}

func (db *ExpressDB) Query(order, com string) (string, error) {
    now := time.Now()

    var (
        row *sql.Row
        err error
    )

    db.rwlock.RLock()

    if com == "" || com == "auto" {
        row = db.db.QueryRow("SELECT express_record, update_time FROM express WHERE express_code=?", order)

    } else {
        row = db.db.QueryRow("SELECT express_record, update_time FROM express WHERE express_code=? AND express_com=?", order, com)
    }

    if err != nil {
        db.rwlock.RUnlock()
        return "", err
    }

    var (
        expressRecord string
        updateTime    int64
    )

    err = row.Scan(&expressRecord, &updateTime)
    db.rwlock.RUnlock()

    if err == sql.ErrNoRows {
        return "", sql.ErrNoRows
    }

    if err != nil {
        return "", err
    }

    update := time.Unix(updateTime, 0)

    if update.Sub(now) > 10*time.Minute {
        return "", Timeout
    }

    return expressRecord, nil
}

func Connect(user, password string) (*ExpressDB, error) {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/expresss", user, password))

    if err != nil {
        return nil, err
    }

    return &ExpressDB{
        db: db,
    }, nil
}
