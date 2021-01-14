src := $(shell git rev-parse --short HEAD)
date := $(shell date +%F_%T)
rev-full := $(src) $(date)

all:
	go build -ldflags "-X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun