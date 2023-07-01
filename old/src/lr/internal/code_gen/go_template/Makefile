MAKEFLAGS = -s

TEMPLATES = $(wildcard *.template)
GENERATED = $(TEMPLATES:.template=.go)

all: touch $(GENERATED)

%.go: %.template
	echo Compiling $<
	go run ./../../../../template/bin/compile-template -o $@ $<

touch:
	touch *.template

clean:
	rm -f $(GENERATED)
