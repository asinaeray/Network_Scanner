package scanner

import (
	"fmt"
	"network-scanner/internal/ping"
)

func ScanNetwork(startIP, endIP string) {
	for ip := startIP; ip != endIP; ip = incrementIP(ip) {
		result, err := ping.Ping(ip)
		if err != nil {
			fmt.Printf("Hata: %v\n", err)
		} else {
			fmt.Println(result)
		}
	}
}

// İp arttırma
func incrementIP(ip string) string {
	return ip
}
