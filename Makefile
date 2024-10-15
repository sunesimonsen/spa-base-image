default: build

tmp/main: **/*.go
	go build -o tmp/main

build: tmp/main

clean:
	rm -rf tmp
