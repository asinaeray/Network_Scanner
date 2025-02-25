package ping

import (
	"testing"
)

// Ping fonksiyonunun test edilmesi
func TestPing(t *testing.T) {
	ip := "192.168.1.1"
	result, err := Ping(ip)
	if err != nil {
		t.Errorf("Beklenen sonuç, fakat hata alındı: %v", err)
	}
	if result == "" {
		t.Errorf("Ping sonucu boş döndü")
	}
}
