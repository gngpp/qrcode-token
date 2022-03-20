#!/bin/sh
export GOPROXY=direct

sudo apt-get update
sudo apt-get install gcc-mingw-w64-i686 gcc-multilib -y

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o linux_amd64_qrcode-token ./main.go
tar -czvf linux_amd64_qrcode-token.tar.gz ./linux_amd64_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o linux_386_qrcode-token ./main.go
tar -czvf linux_386_qrcode-token.tar.gz ./linux_386_qrcode-token

CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o freebsd_386_qrcode-token ./main.go
tar -czvf freebsd_386_qrcode-token.tar.gz ./freebsd_386_qrcode-token

CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o freebsd_amd64_qrcode-token ./main.go
tar -czvf freebsd_amd64_qrcode-token.tar.gz ./freebsd_amd64_qrcode-token

CGO_ENABLED=0 GOOS=freebsd GOARCH=arm go build -ldflags="-s -w" -o freebsd_arm64_qrcode-token ./main.go
tar -czvf freebsd_arm64_qrcode-token.tar.gz ./freebsd_arm64_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o linux_armv7_qrcode-token ./main.go
tar -czvf linux_armv7_qrcode-token.tar.gz ./linux_armv7_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o linux_armv6_qrcode-token ./main.go
tar -czvf linux_armv6_qrcode-token.tar.gz ./linux_armv6_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o linux_armv5_qrcode-token ./main.go
tar -czvf linux_armv5_qrcode-token.tar.gz ./linux_armv5_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o linux_arm64_qrcode-token ./main.go
tar -czvf linux_arm64_qrcode-token.tar.gz ./linux_arm64_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags="-s -w" -o linux_mips64_qrcode-token ./main.go
tar -czvf linux_mips64_qrcode-token.tar.gz ./linux_mips64_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags="-s -w" -o linux_mips64le_qrcode-token ./main.go
tar -czvf linux_mips64le_qrcode-token.tar.gz ./linux_mips64le_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags="-s -w" -o linux_mipsle_qrcode-token ./main.go
tar -czvf linux_mipsle_qrcode-token.tar.gz ./linux_mipsle_qrcode-token

CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -ldflags="-s -w" -o linux_mips_qrcode-token ./main.go
tar -czvf linux_mips_qrcode-token.tar.gz ./linux_mips_qrcode-token

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o windows_386_qrcode-token.exe ./main.go
tar -czvf windows_386_qrcode-token.tar.gz ./windows_386_qrcode-token.exe

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o windows_amd64_qrcode-token.exe ./main.go
tar -czvf windows_amd64_qrcode-token.tar.gz ./windows_amd64_qrcode-token.exe

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o darwin_amd64_qrcode-token ./main.go
tar -czvf darwin_amd64_qrcode-token.tar.gz ./darwin_amd64_qrcode-token

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o darwin_arm64_qrcode-token ./main.go
tar -czvf darwin_arm64_qrcode-token.tar.gz ./darwin_arm64_qrcode-token

