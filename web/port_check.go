package web

import (
	log2 "app.rz-public.xyz/wiimmfi-room-watcher/utils/log"
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
			newPort := port + 1
			log2.Logger.Info().Msg(fmt.Sprint(address, ":", port) + " is already in use. " + fmt.Sprint(address, ":", newPort) + " will use instead.")
			port = newPort
			continue
		}
		return fmt.Sprint(address, ":", port), nil
	}
}
