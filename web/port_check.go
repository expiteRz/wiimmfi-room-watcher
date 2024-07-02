package web

import (
	"fmt"
	"log"
	"net"
	"time"
)

func portCheck(address string, port int) (string, error) {
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprint(address, ":", port), time.Second)
		if err != nil {
			return fmt.Sprint(address, ":", port), nil
		}
		if conn != nil {
			newPort := port + 1
			log.Println(fmt.Sprint(address, ":", port), "is already in use.", fmt.Sprint(address, ":", newPort), "will use instead.")
			port = newPort
			continue
		}
		return fmt.Sprint(address, ":", port), nil
	}
}
