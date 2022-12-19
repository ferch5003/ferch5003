package main

import (
	"log"

	"github.com/ferch5003/ferch5003/readme-maker/internal"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	deps, err := newDependencies()
	if err != nil {
		return err
	}

	mainTemplate := internal.NewTemplate()
	mainTemplate.AddTemplates(deps.Templates)

	readmeString, err := deps.Storage.Read()
	if err != nil {
		return err
	}

	mainTemplateString, err := mainTemplate.Parse(readmeString)
	if err != nil {
		return err
	}

	return deps.Storage.Write(mainTemplateString)
}
