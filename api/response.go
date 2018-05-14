package api

type State struct {
    Time    string `json:"time"`
    Context string `json:"context"`
}

type Response struct {
    Com    string  `json:"com"`
    Status string  `json:"status"`
    State  string  `json:"state"`
    Data   []State `json:"data"`
}

var States = map[string]string{
    "0": "在途，即货物处于运输过程中",
    "1": "揽件，货物已由快递公司揽收并且产生了第一条跟踪信息",
    "2": "疑难，货物寄送过程出了问题",
    "3": "签收，收件人已签收",
    "4": "退签，即货物由于用户拒签、超区等原因退回，而且发件人已经签收",
    "5": "派件，即快递正在进行同城派件",
    "6": "退回，货物正处于退回发件人的途中",
}

/*var Statuses = map[string]string{
    "0": "物流单暂无结果",
    "1": "查询成功",
    "2": "接口出现异常",
}*/
