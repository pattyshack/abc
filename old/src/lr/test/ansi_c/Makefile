MAKEFLAGS = -s

all: lr yacc

lr: ansi-c.lr
	go run ./../../bin/generate-lr-parser -o ansi-c.go ansi-c.lr

yacc: ansi-c.y
	goyacc -o ansi-c.yacc.go -v ansi-c.yacc.output -p YC ansi-c.y

clean:
	rm -f ansi-c.go ansi-c.yacc.*
