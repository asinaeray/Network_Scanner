# Network Scanner

A simple network scanning tool built with Go to perform IP and port scanning. This project helps in checking the availability of an IP address and scanning specific ports on a target IP.

## Features

- **Ping Scan:** Ping a given IP address to check its availability on the network.
- **Port Scan:** Scan specific ports on an IP address to determine if they are open or closed.
- **Multi-Port Support:** Allows scanning multiple ports at once to find open ports on a given IP.

## Requirements

- Go 1.18+ (Project is compatible with the latest Go version)
- `net` package (Go's standard library)

## Installation

Follow these steps to install and run the project:

1. Clone the repository or download the files:
    ```bash
    git clone https://github.com/your-username/network-scanner.git
    cd network-scanner
    ```

2. Install the required dependencies:
    ```bash
    go mod tidy
    ```

3. Build and run the project:
    ```bash
    go run main.go
    ```

## Usage

To scan an IP or port, use the following commands:

1. **Ping a single IP:**  
    Run the `main.go` file to ping an IP address:
    ```bash
    go run main.go
    ```
    This will check the availability of a given IP address.

2. **Scan a specific port on an IP:**  
    Run this command to scan a specific port on the target IP:
    ```bash
    go run main.go 192.168.1.1 80
    ```
    This will check if port `80` on IP `192.168.1.1` is open.

3. **Scan multiple ports (optional):**  
    You can modify the code to scan a range of ports or multiple ports, depending on your need.

## Testing

To run tests for this project:
```bash
go test ./...
