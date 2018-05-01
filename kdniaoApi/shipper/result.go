package shipper

import "fmt"

type Result struct {
    EBusinessID  string
    Success      bool
    LogisticCode string
    Shippers     []Shipper
}

type Shipper struct {
    ShipperCode string
    ShipperName string
}

func (r *Result) ShipperCode() (string, error) {
    if !r.Success {
        return "", fmt.Errorf("order: %s can't find shipper\n", r.LogisticCode)
    }

    return r.Shippers[0].ShipperCode, nil
}
