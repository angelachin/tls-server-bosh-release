package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"flag"
)

func main() {
	configFilePath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := LoadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("load config file: %s", err)
	}

	cert, err := tls.LoadX509KeyPair(cfg.ServerCertFile, cfg.ServerKeyFile)
	if err != nil {
		log.Fatalf("%s", err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	l, err := tls.Listen("tcp4", fmt.Sprintf(":%d", cfg.ListenPort), &config)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer l.Close()

	helloMessage := "hello over TLS from bosh deployed service"

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("accepting connection: %s", err)
			continue
		}
		go handleConnection(c, helloMessage)
	}
}

func handleConnection(c net.Conn, helloMessage string) error {
	defer c.Close()
	log.Printf("greeting %s...", c.RemoteAddr().String())

	_, err := c.Write([]byte(helloMessage))
	if err != nil {
		log.Printf("error writing: %s", err)
		return err
	}

	log.Printf("  done with %s...", c.RemoteAddr().String())
	return nil
}
