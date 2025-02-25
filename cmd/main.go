package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"network-scanner/internal/scanner"
)

type Config struct {
	StartIP string `json:"startIP"`
	EndIP   string `json:"endIP"`
}

func main() {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("Yapilandirma dosyasi okunamadi:", err)
		return
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("Yapilandirma dosyasi hatali:", err)
		return
	}

	// ağ taraması başlatılması
	fmt.Println("Ağ taramasi başlatiliyor...")
	scanner.ScanNetwork(config.StartIP, config.EndIP)
}
