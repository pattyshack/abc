MAKEFLAGS = -s

TEMPLATES = $(wildcard *.template)
GENERATED = $(TEMPLATES:.template=.go)

all: clean $(GENERATED)
	go run .

%.go: %.template
	echo Compiling $<
	go run ./../bin/compile-template -o $@ $<

clean:
	rm -f $(GENERATED)
