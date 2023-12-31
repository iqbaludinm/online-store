.PHONY : format install build

#Website/RESTful API
run:
	@echo "Running server..."
	go run ./main.go

init:
	@echo "Initializing dependencies..."
	go mod init
	go mod tidy
	
install:
	@echo "Downloading dependencies..."
	go mod download

build:
	@echo "building binary..."
	go build ./main.go

start:
	@echo "Starting server..."
	./main

swag:
	@echo "Generating swagger docs..."
	swag init -g ./main.go

swag-debug:
	@echo "Generating swagger docs..."
	swag init --parseDependency --parseInternal --parseDepth 2 -g /main.go -d ./

clean:
	@echo "Cleaning..."
	rm -rf ./main.exe

# live reload using nodemon: npm -g i nodemon
run-nodemon:
	@echo "Running server with nodemon..."
	nodemon --exec go run ./main.go