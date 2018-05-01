package packet

type Status struct {
    LogisticCode string
    ShipperCode  string
    Traces       []State
}

type State struct {
    AcceptStation string
    AcceptTime    string
}
