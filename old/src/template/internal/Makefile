MAKEFLAGS = -s

parser.go: grammar.lr
	go run ./../../lr/bin/generate-lr-parser -o grammar.go grammar.lr

clean:
	rm -f grammar.go
