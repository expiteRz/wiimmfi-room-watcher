package web

import (
	"fmt"
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
			return fmt.Sprint(address, ":", port), fmt.Errorf("%s is already in use", fmt.Sprint(address, ":", port))
		}
		return fmt.Sprint(address, ":", port), nil
	}
}
