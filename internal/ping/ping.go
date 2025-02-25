package ping

import (
	"fmt"
	"net"
	"time"
)

// belirli ip ye ping atma
func Ping(ip string) (string, error) {
	conn, err := net.DialTimeout("ip4:icmp", ip, time.Second*3)
	if err != nil {
		return "", fmt.Errorf("ping atarken hata: %v", err)
	}
	defer conn.Close()

	return fmt.Sprintf("Ping başarili %s IP'sine bağlanildi.", ip), nil
}
