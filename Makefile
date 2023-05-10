bin: bin/newl_darwin_amd64 bin/newl_linux_amd64 bin/newl_windows_amd64.exe

bin/newl_darwin_amd64:
	@mkdir -p bin
	@echo "Compiling newl..."
	GOOS=darwin GOARCH=amd64 go build -o $@ main.go

bin/newl_linux_amd64:
	@mkdir -p bin
	@echo "Compiling newl..."
	GOOS=linux GOARCH=amd64 go build -o $@ main.go

bin/newl_windows_amd64.exe:
	@mkdir -p bin
	@echo "Compiling newl..."
	GOOS=windows GOARCH=amd64 go build -o $@ main.go