package main

import (
	"fmt"
	"path/filepath"

	"astral/internal/fs"
	"astral/internal/parser"
	"astral/internal/transpiler"
	"astral/internal/generator"
)

func main() {
	entries, err := fs.ListItemFiles()
	if err != nil {
		panic(err)
	}

	for _, e := range entries {
		path := filepath.Join("content/items", e.Name())

		sword, err := parser.ParseSword(path)
		if err != nil {
			panic(err)
		}

		code := transpiler.SwordToCS(sword)

		outPath := filepath.Join("output/Items", sword.Name+".cs")
		err = generator.WriteFile(outPath, code)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated:", outPath)
	}

	fmt.Println("Astral Engine v0.0.1 complete!")
}
