all:
	@wire ./internal/setup/server
	@go build -o ./build ./cmd/app/app.go

run: all
	@./build/app

clean:
	rm ./build/app