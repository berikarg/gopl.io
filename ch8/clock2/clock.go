// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "port")

func handleConn(c net.Conn, p string) {
	defer c.Close()
	for {
		switch p {
		case "8010":
			tz, _ := time.LoadLocation("US/Eastern")
			_, err := io.WriteString(c, time.Now().In(tz).Format("15:04:05\n"))
			if err != nil {
				return // e.g., client disconnected
			}
			time.Sleep(1 * time.Second)
		case "8020":
			tz, _ := time.LoadLocation("Asia/Tokyo")
			_, err := io.WriteString(c, time.Now().In(tz).Format("15:04:05\n"))
			if err != nil {
				return // e.g., client disconnected
			}
			time.Sleep(1 * time.Second)
		case "8030":
			tz, _ := time.LoadLocation("Europe/London")
			_, err := io.WriteString(c, time.Now().In(tz).Format("15:04:05\n"))
			if err != nil {
				return // e.g., client disconnected
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, *port) // handle connections concurrently
	}
	//!-
}
