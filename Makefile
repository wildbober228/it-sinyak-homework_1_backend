.PHONY: run check fix

run:
	go run .

check:
	@unformatted="$$(gofmt -l .)"; if [ -n "$$unformatted" ]; then echo "gofmt: $$unformatted"; exit 1; fi
	go vet ./...
	go test ./...

fix:
	gofmt -w .
	go vet ./...
	go test ./...
