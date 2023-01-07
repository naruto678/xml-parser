build: 
	@go build -o ./bin/parser 
run: build
	@./bin/parser 
test: 
	@go test -v . 

