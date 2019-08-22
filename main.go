package main

import (
	"github.com/trntv/bricks/blueprint"
	"github.com/trntv/bricks/generator"
	"os"
	"plugin"
)

func main() {
	pf := os.Args[1]
	p, err := plugin.Open(pf)
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("NewGenerator")
	if err != nil {
		panic(err)
	}

	outDir := "/Users/terentev/Code/brick/out"
	modName := "testpkg"

	os.Mkdir(outDir, 0755)

	generator.GenModules(outDir, modName)

	bp := &blueprint.DefaultBlueprint{}
	gen := f.(func(string) blueprint.Generator)("/Users/terentev/Code/brick/test")
	gen.Generate(outDir, bp)

	generator.InstallModules(outDir)
}
