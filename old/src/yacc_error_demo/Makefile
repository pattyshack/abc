MAKEFLAGS = -s

parser.go: parser.y
	goyacc -o parser.go -v parser.output -p demo parser.y

clean:
	rm -f parser.output parser.go
