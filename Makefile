build:
	#windows 32 & 64bit
	GOOS=windows GOARCH=amd64 go build -o bin/fabula64-windows.exe main.go
	GOOS=windows GOARCH=386 go build -o bin/fabula32-windows.exe main.go

	#macos 64bit
	GOOS=darwin GOARCH=amd64 go build -o bin/fabula64-macos main.go

	#linux 64bit
	GOOS=linux GOARCH=amd64 go build -o bin/fabula64-linux main.go

run:
	go run main.go -username=$(USERNAME)

clean:
	rm -rf bin