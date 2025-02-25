package ports

import (
	"fmt"
	"testing"
)

func TestScanPort(t *testing.T) {
	ip := "192.168.1.1"
	port := 80

	result, err := ScanPort(ip, port)
	if err != nil {
		t.Errorf("Port tarama hatasi: %v", err)
	}

	expected := fmt.Sprintf("Port %d açik", port)
	if result != expected {
		t.Errorf("Beklenen %v, ancak %v alindi", expected, result)
	}
}
