package main

import (
	"flag"
	"fmt"
	"os"
)

var redisAddr = flag.String("redisaddr", "127.0.0.1:6379", "redis backend port")
var redisNetwork = flag.String("redisnet", "tcp", "redis network (tcp, udp, unix, etc)")
var port = flag.Uint("port", 8080, "port")

func main() {
	flag.Parse()
	err := serve(uint16(*port), *redisNetwork, *redisAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}