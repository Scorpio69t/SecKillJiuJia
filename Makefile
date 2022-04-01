.PHONY: all build run gotool clean help

BINARY="./build/seckill-jiujia"
BUILDDIR="./build/"

all: check linux-build


linux-build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}
	cp -rf config.toml ${BUILDDIR}

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe
	cp -rf config.toml ${BUILDDIR} 

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}
	cp -rf config.toml ${BUILDDIR} 

run:
	@go run main.go

check:
	go fmt ./
	go vet ./

clean:
	@rm -rf ${BUILDDIR}/*

help:
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make run - 直接运行 main.go"
	@echo "make clean - 移除二进制文件"
	@echo "make check - 运行 Go 工具 'fmt' and 'vet'"
