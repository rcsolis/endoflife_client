include .env
export $(shell sed 's/=.*//' .env)

dev:
	@echo "--> Variables"
	@echo $(BINARY_FILENAME)
	@echo $(BINARY_NAME)
	@echo "-->Run dev mode"
	go run ./cmd/$(BINARY_FILENAME).go

test:
	echo "-->Run test"
	go test -v ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

clean:
	@echo "-->Clean"
	go clean
	rm -rf test.db
	rm -rf bin

dep: clean
	@echo "-->Download dependencies"
	go mod download
	go mod verify
	go mod tidy

build: dep
	@echo "==>Building binary"
	@echo "--> Variables"
	@echo $(BINARY_FILENAME)
	@echo $(BINARY_NAME)
	@echo "-->Build command"
	go build -o bin/${BINARY_NAME} -v ./cmd/$(BINARY_FILENAME).go

run: build
	@echo "--> Variables"
	@echo $(BINARY_FILENAME)
	@echo $(BINARY_NAME)
	@echo "==>Run binary"
	./bin/$(BINARY_NAME)