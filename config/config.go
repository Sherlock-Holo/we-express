package config

import (
    "github.com/BurntSushi/toml"
    "os"
    "bytes"
)

var ShipperCode = map[string]string{
    "SF":       "顺丰速运",
    "HTKY":     "百世快递",
    "ZTO":      "中通快递",
    "STO":      "申通快递",
    "YTO":      "圆通速递",
    "YD":       "韵达速递",
    "YZPY":     "邮政快递",
    "EMS":      "EMS",
    "HHTT":     "天天快递",
    "JD":       "京东物流",
    "UC":       "优速快递",
    "DBL":      "德邦",
    "FAST":     "快捷快递",
    "ZJS":      "宅急送",
    "TNT":      "TNT快递",
    "UPS":      "UPS",
    "DHL":      "DHL",
    "FEDEX":    "FEDEX联邦(国内件)",
    "FEDEX_GJ": "FEDEX联邦(国际件)",
    "BDT":      "八达通",
    "BTWL":     "百世快运",
}

type Config struct {
    ID string `toml:"id"`
}

func Parse(f string) (Config, error) {
    file, err := os.Open(f)

    if err != nil {
        return Config{}, err
    }

    if err != nil {
        return Config{}, err
    }

    conf := Config{}

    buf := bytes.NewBufferString("")

    _, err = buf.ReadFrom(file)

    if err != nil {
        return Config{}, err
    }

    _, err = toml.Decode(buf.String(), &conf)

    if err != nil {
        return Config{}, err
    }

    return conf, nil
}
