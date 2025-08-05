package server

import (
	"bufio"
	"exception-reporter-agent/handler"
	"log"
	"net"
)

func ListenAndServe(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		var raw = scanner.Bytes()
		go handler.HandleException(raw)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Read error: %v", err)
	}
}
