MAKEFLAGS = -s

parser.go: grammar.y
	goyacc -o grammar.go -v grammar.output -p Lr grammar.y

clean:
	rm -f grammar.output grammar.go
