package config

import (
    "github.com/BurntSushi/toml"
    "os"
    "bytes"
    "mime"
    "path"
    "encoding/json"
)

const QueryUrl = "http://api.kdniao.cc/Ebusiness/EbusinessOrderHandle.aspx"

var ShipperCode = map[string]string{
    "顺丰速运":         "SF",
    "百世快递":         "HTKY",
    "中通快递":         "ZTO",
    "申通快递":         "STO",
    "圆通速递":         "YTO",
    "韵达速递":         "YD",
    "邮政快递":         "YZPY",
    "EMS":          "EMS",
    "天天快递":         "HHTT",
    "京东物流":         "JD",
    "优速快递":         "UC",
    "德邦":           "DBL",
    "快捷快递":         "FAST",
    "宅急送":          "ZJS",
    "TNT快递":        "TNT",
    "UPS":          "UPS",
    "DHL":          "DHL",
    "FEDEX联邦(国内件)": "FEDEX",
    "FEDEX联邦(国际件)": "FEDEX_GJ",
    "八达通":          "BDT",
    "百世快运":         "BTWL",
}

type Config struct {
    EBusinessID string
    AppKey      string
}

func Parse(f string) (Config, error) {
    file, err := os.Open(f)

    if err != nil {
        return Config{}, err
    }

    info, err := file.Stat()

    if err != nil {
        return Config{}, err
    }

    extension := mime.TypeByExtension(path.Ext(info.Name()))

    conf := Config{}

    if extension == "application/json" {
        decoder := json.NewDecoder(file)

        err = decoder.Decode(&conf)

        if err != nil {
            return Config{}, err
        }

        return conf, nil
    }

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
