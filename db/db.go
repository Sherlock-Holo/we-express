package db

// database format
// express_id(varchar 20), express_com(10), express_record(text), update_time(int)

import (
    "errors"
    "database/sql"
    "sync"
    "time"
    "fmt"
    "log"
    "github.com/Sherlock-Holo/we-express/api"
    "strconv"
    "encoding/json"

    _ "github.com/Go-SQL-Driver/MySQL"
)

var (
    Timeout = errors.New("database data timeout")

    updatePrepared *sql.Stmt
    insertPrepared *sql.Stmt
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

    serverResp := Response{}

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
            record := Record{}

            jTime := NewJTime(data.Time)
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

    if exist {
        _, err = updatePrepared.Exec(unixTime, jsonString, apiResp.Com, order)
    } else {
        _, err = insertPrepared.Exec(unixTime, jsonString, apiResp.Com, order)
    }

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
        row = db.db.QueryRow("SELECT express_record, update_time FROM express WHERE express_id=?", order)

    } else {
        row = db.db.QueryRow("SELECT express_record, update_time FROM express WHERE express_id=?", order)
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

    if now.Sub(update) > 10*time.Minute {
        return "", Timeout
    }

    return expressRecord, nil
}

func (db *ExpressDB) ListExpress() (*sql.Rows, error) {
    return db.db.Query("select * from express")
}

func (db *ExpressDB) Check(order string) bool {
    db.rwlock.RLock()
    defer db.rwlock.RUnlock()

    row := db.db.QueryRow("SELECT express_id FROM express WHERE express_id=?", order)

    err := row.Scan(new(string))

    if err == sql.ErrNoRows {
        return false
    } else {
        return true
    }
}

func Connect(user, password string) (*ExpressDB, error) {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/express", user, password))

    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    updatePrepared, err = db.Prepare("update express set update_time=?, express_record=?, express_com=? where express_id=?")
    if err != nil {
        return nil, err
    }

    insertPrepared, err = db.Prepare("insert into express (update_time, express_record, express_com, express_id) values (?, ?, ?, ?)")
    if err != nil {
        return nil, err
    }

    return &ExpressDB{
        db: db,
    }, nil
}
