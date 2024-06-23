package web

import (
	"fmt"
	"log"
	"net"
	"time"
)

func portCheck(address string, port int) (int, error) {
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprint(address, ":", port), time.Second)
		if err != nil {
			return port, nil
		}
		if conn != nil {
			newPort := port + 1
			log.Println(address, "is already in use.", fmt.Sprint(address, ":", newPort), "will use instead.")
			port = newPort
			continue
		}
		return port, nil
	}
}
