package ports

import (
	"fmt"
	"net"
)

func ScanPort(ip string, port int) (string, error) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Sprintf("Port %d kapalı", port), err
	}
	defer conn.Close()
	return fmt.Sprintf("Port %d açık", port), nil
}
