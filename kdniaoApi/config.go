package kdniaoApi

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
