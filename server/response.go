package server

import "time"

// State and StateInfo:
// "0": "物流单暂无结果",
// "1": "查询成功",
// "2": "接口出现异常",
type Response struct {
    Order     string
    Com       string
    Status    bool
    State     int
    StateInfo string
    Records   []Record
}

type Record struct {
    Time JTime
    Info string
}

type JTime time.Time

func (jt *JTime) MarshalJSON() ([]byte, error) {
    t := time.Time(*jt)
    return []byte(t.Format(`"2006-01-02 15:04:05"`)), nil
}

func newJTime(t string) JTime {
    jTime, _ := time.Parse("2006-01-02 15:04:05", t)
    return JTime(jTime)
}
