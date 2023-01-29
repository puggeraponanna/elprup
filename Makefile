build:
	go build -o tmp/elprup main.go

run: build
	./tmp/elprup