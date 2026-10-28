package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	host := "scanme.nmap.org"

	for port := 20; port <= 25; port++ {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err != nil {
			fmt.Println("Porta", port, "fechada ou filtrada")
			continue
		}
		conn.Close()
		fmt.Println("Porta", port, "aberta")
	}
}
