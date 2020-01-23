all: restore-deps test vet

test:
	go test -v ./tabular/*
vet: 
	go vet ./tabular/*
clean:
	rm *.test
restore-deps:
	go mod tidy
