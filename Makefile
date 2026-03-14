COP=cover.out

.PHONY: vet lint test test-cover clean

vet:
	@- go vet ./...

lint: vet
	@- golangci-lint run

test: vet
	@- go test -race -count 1 -v -cover ./... -coverpkg ./...

test-cover: vet
	@- go test -v -coverprofile="$(COP)" -cover ./... -coverpkg ./... -covermode=count
	@- go tool cover -func="$(COP)" -o="$(COP)"

bench:
	@- go test -v -count 1 -bench=. -benchmem -timeout 15m

clean:
	@- rm -f "$(COP)"
