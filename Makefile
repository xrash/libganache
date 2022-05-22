.PHONY : test

test :
	go test -race -cover ./...

.PHONY : check
check :
	staticcheck ./...; go vet ./...
