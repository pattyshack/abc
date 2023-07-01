package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Template interface {
	io.WriterTo

	Name() string
}

type CustomStringer struct {}

func (CustomStringer) String() string { return "CustomStringer" }

func main() {
	flag.Parse()

	names := map[string]struct{}{}
	for _, name := range flag.Args() {
		names[name] = struct{}{}
	}

	forChan := make(chan int, 10)
	forChan <- 1
	forChan <- 3
	forChan <- 5
	forChan <- 7
	forChan <- 9
	close(forChan)

	templates := []Template{
		&TextTemplate{},
		&SubstituteTemplate{},
        &OutputTypesTemplate{CustomStringer{}},
		&ForTemplate{forChan, 0},
        &IfTemplate{},
        &SwitchTemplate{1, "hello"},
        &EmbedTemplate{},
        &ErrorTemplate{},
        &TrimWhitespacesTemplate{},
	}

	for _, template := range templates {
		_, ok := names[template.Name()]
		if ok || len(names) == 0 {
			fmt.Println("Template:", template.Name())
			fmt.Println("-----------------------")
			n, err := template.WriteTo(os.Stdout)
			fmt.Println("-----------------------")
			fmt.Println(n, " bytes written")
			fmt.Println("error:", err)
			fmt.Println("=======================")
		}
	}
}
