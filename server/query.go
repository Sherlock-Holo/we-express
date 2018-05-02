package server

import (
    "github.com/we-express/kdniaoApi/packet"
    "github.com/we-express/kdniaoApi/shipper"
    "github.com/we-express/config"
    "time"
    "fmt"
)

func QueryExpress(id, appKey, order string) (Response, error) {
    shipperQuery, err := shipper.NewQuery(id, appKey, order)

    if err != nil {
        return Response{}, err
    }

    shipperResult, err := shipper.DoQuery(shipperQuery)

    if err != nil {
        return Response{}, fmt.Errorf("query shipper code failed")
    }

    shipperCode, err := shipperResult.ShipperCode()

    if err != nil {
        return Response{}, err
    }

    packetQuery, err := packet.NewQuery(id, appKey, order, shipperCode)

    if err != nil {
        return Response{}, err
    }

    status, err := packet.DoQuery(packetQuery)

    if err != nil {
        return Response{}, err
    }

    response := Response{}

    response.Order = status.LogisticCode

    shipperName, ok := config.ShipperCode[status.ShipperCode]

    if !ok {
        response.Shipper = "unknown"
    } else {
        response.Shipper = shipperName
    }

    for _, trace := range status.Traces {
        t, err := parseTime(trace.AcceptTime)

        if err != nil {
            return Response{}, err
        }

        state := State{
            Info: trace.AcceptStation,
            Date: t,
        }

        response.States = append(response.States, state)
    }

    return response, nil
}

func parseTime(t string) (jTime, error) {
    parsed, err := time.Parse("2006-01-02 15:04:05", t)

    if err != nil {
        return jTime{}, err
    }

    return jTime(parsed), nil
}
