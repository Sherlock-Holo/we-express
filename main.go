package main

import (
    "flag"
    "github.com/we-express/server"
    "os"
)

var (
    addr = flag.String("addr", "", "listen addr")
    port = flag.Int("port", 80, "listen port")
    conf = flag.String("conf", "", "config file")
)

func main() {
    flag.Parse()

    if flag.NFlag() == 0 {
        flag.Usage()
        os.Exit(2)
    }

    server.Start(*addr, *port, *conf)
}
