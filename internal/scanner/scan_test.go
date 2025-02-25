package scanner

import (
	"testing"
)

func TestScanNetwork(t *testing.T) {
	startIP := "192.168.1.1"
	endIP := "192.168.1.5"
	ScanNetwork(startIP, endIP)
}
