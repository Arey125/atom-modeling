.PHONY: all run

all:
	go build -o ./bin/model ./cmd/*.go

run: all
	./bin/model
