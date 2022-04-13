.PHONY:all
all:vet test

.PHONY:vet
vet:
	go vet --all .

.PHONY:test
test:
	go test -v -race .
