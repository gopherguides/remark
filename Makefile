test: peg
	go test -cover -count 1 ./...

peg:
	pigeon internal/parser/remark.peg | goimports > internal/parser/parse.go

deps:
	go get github.com/mna/pigeon
	go get -t ./...
