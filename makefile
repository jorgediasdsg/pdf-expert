test:
	go test ./...

test-race:
	go test -race ./...

test-cover:
	go test -cover ./...

test-usecase:
	go test ./internal/app/usecase -v