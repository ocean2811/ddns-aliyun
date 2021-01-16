src := $(shell git rev-parse --short HEAD)
date := $(shell date +%F_%T)
rev-full := $(src) $(date)

all:
	go build -ldflags "-X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun.exe

386:
	GOOS=linux GOARCH=386 go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

mipsle:
	GOOS=linux GOARCH=mipsle go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

mipsle-sf:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

mips64le:
	GOOS=linux GOARCH=mips64le go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

mips:
	GOOS=linux GOARCH=mips go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

mips-sf:
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun

mips64:
	GOOS=linux GOARCH=mips64 go build -ldflags "-s -w -X 'github.com/ocean2811/ddns-aliyun/cmd.version=$(rev-full)'" -o ddns-aliyun
