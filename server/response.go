package server

import "time"

type Response struct {
    Order   string
    Shipper string
    States  []State
}

type State struct {
    Date jTime
    Info string
}

type jTime time.Time

func (jt jTime) MarshalJSON() ([]byte, error) {
    t := time.Time(jt)

    return []byte(t.Format(`"2006-01-02 15:04:05"`)), nil
}
