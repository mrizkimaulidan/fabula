build:
	#windows 32 & 64bit
	GOOS=windows GOARCH=amd64 go build -o bin/fabula64-windows.exe *.go
	GOOS=windows GOARCH=386 go build -o bin/fabula32-windows.exe *.go

	#macos 64bit
	GOOS=darwin GOARCH=amd64 go build -o bin/fabula64-macos *.go

	#linux 64bit
	GOOS=linux GOARCH=amd64 go build -o bin/fabula64-linux *.go

story:
	go build && ./fabula -username=$(USERNAME) -option=story

highlight:
	go build && ./fabula -username=$(USERNAME) -option=highlight

clean:
	rm -rf bin
	rm -rf stories
	rm -rf highlights
	rm -f fabula
