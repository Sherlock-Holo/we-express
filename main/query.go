package main

import (
    "github.com/we-express/kdniaoApi/packet"
    "github.com/we-express/kdniaoApi/shipper"
)

func QueryExpress(id, appKey, order string) (packet.Status, error) {
    shipperQuery, err := shipper.NewQuery(id, appKey, order)

    if err != nil {
        return packet.Status{}, err
    }

    shipperResult, err := shipper.DoQuery(shipperQuery)

    if err != nil {
        return packet.Status{}, err
    }

    shipperCode, err := shipperResult.ShipperCode()

    if err != nil {
        return packet.Status{}, err
    }

    packetQuery, err := packet.NewQuery(id, appKey, order, shipperCode)

    if err != nil {
        return packet.Status{}, err
    }

    status, err := packet.DoQuery(packetQuery)

    if err != nil {
        return packet.Status{}, err
    }

    return status, nil
}
