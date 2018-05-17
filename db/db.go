package db

// database format
// express_id(varchar 20), express_com(10), express_record(text), update_time(int)

import (
    "database/sql"
    "encoding/json"
    "errors"
    "fmt"
    _ "github.com/Go-SQL-Driver/MySQL"
    "github.com/Sherlock-Holo/we-express/api"
    "log"
    "strconv"
    "time"
)

var (
    Timeout = errors.New("database data timeout")

    quickQueryPrepared     *sql.Stmt
    queryPrepared          *sql.Stmt
    insertOrUpdatePrepared *sql.Stmt
)

type ExpressDB struct {
    db *sql.DB
}

func (db *ExpressDB) Update(order, com, apiID string) (string, error) {
    apiResp, err := api.Query(order, apiID, com)

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

    _, err = insertOrUpdatePrepared.Exec(order, unixTime, jsonString, apiResp.Com)
    if err != nil {
        return "", err
    }

    return jsonString, nil
}

func (db *ExpressDB) Query(order, com string) (string, error) {
    now := time.Now()

    var (
        row           *sql.Row
        err           error
        expressRecord string
        updateTime    int64
    )

    if com == "" || com == "auto" {
        row = quickQueryPrepared.QueryRow(order)
    } else {
        row = queryPrepared.QueryRow(order, com)
    }

    err = row.Scan(&expressRecord, &updateTime)

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
    return db.db.Query("select express_id, express_com from express")
}

func Connect(user, password string) (*ExpressDB, error) {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/express", user, password))

    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    quickQueryPrepared, err = db.Prepare("SELECT express_record, update_time FROM express WHERE express_id=?")
    if err != nil {
        return nil, err
    }

    queryPrepared, err = db.Prepare("SELECT express_record, update_time FROM express WHERE express_id=? AND express_com=?")
    if err != nil {
        return nil, err
    }

    insertOrUpdatePrepared, err = db.Prepare("INSERT INTO express (express_id, update_time, express_record, express_com) values (?, ?, ?, ?) " +
        "on DUPLICATE KEY UPDATE update_time=VALUES(update_time), express_record=VALUES(express_record), express_com=VALUES(express_com)")

    if err != nil {
        return nil, err
    }

    return &ExpressDB{
        db: db,
    }, nil
}
