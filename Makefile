test:
	go test --cover ./...

build:
	go build -o ./bin/main .

run:
	./bin/main -n "1,2,3,4,5"

build-windows:
	go build -o ./bin/main.exe .

run-windows: build-windows
	./bin/main.exe -n "1,2,3,4,5"