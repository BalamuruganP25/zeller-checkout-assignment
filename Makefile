.PHONY: run test build clean

run:
	go run main.go

test:
	go test ./...

build:
	go build -o zeller-checkout-assignment main.go

clean:
	rm -f zeller-checkout-assignment