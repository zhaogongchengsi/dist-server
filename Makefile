# 编译 Windows 可执行文件
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/preview.exe server.go

# 编译 macOS 可执行文件
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o preview server.go

# 打包 Windows 和 macOS 可执行文件
build-all: build-windows build-macos