package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type options struct {
	port string
}

var opt options

func init() {
	flag.StringVar(&opt.port, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if opt.port == "" {
		opt.port = "2007"
	}
}

func main() {
	log.Printf("listening on port %v", opt.port)

	port, err := strconv.Atoi(opt.port)
	if err != nil {
		log.Fatal("port must be an integer, was: " + opt.port)
	}
	log.Println(port)
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Fatalf("listen error, err=%s", err)
		return
	}



	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatalf("listen error, err=%s", err)
		return
	}

}
