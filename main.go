package main

import (
    "flag"
    "os"

    "github.com/Sherlock-Holo/we-express/server"
)

var (
    addr = flag.String("addr", "::", "listen addr")
    port = flag.Uint("port", 80, "listen port")
    conf = flag.String("conf", "", "config file")
)

func main() {
    flag.Parse()

    if flag.NFlag() == 0 {
        flag.Usage()
        os.Exit(2)
    }

    server.Start(*conf, *addr, *port)
}
