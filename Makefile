start:
	go run cmd/api_server/main.go 

build_server:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/prometheus_bridge_linux_amd64 cmd/api_server/main.go
	upx --ultra-brute build/prometheus_bridge_linux_amd64
