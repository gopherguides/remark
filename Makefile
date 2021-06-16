test:
	go test -race -timeout 5s -cover -count 1 ./...

install:
	go install ./cmd/remark
