package main

import (
	"fmt"
	"log"
	"os"
	"path"

	pforest "github.com/p-nordmann/prefix-forest"
	flag "github.com/spf13/pflag"
)

var maxDepth int
var dirPath string

func init() {
	flag.IntVar(&maxDepth, "max-depth", 4, "maximum depth (max 'n' in 'n'-grams)")
	flag.StringVar(&dirPath, "path", "data", "path to corpus directory")
}

func main() {
	flag.Parse()

	infos, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	fmt.Printf("Found %d files.\n", len(infos))

	f1 := pforest.New(maxDepth)
	f2 := pforest.New(maxDepth)
	for k, info := range infos {
		if info.IsDir() {
			continue
		}
		filePath := path.Join(dirPath, info.Name())
		b, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			continue
		}
		text := string(b)
		if k%2 == 0 {
			f1.Learn(text)
		} else {
			f2.Learn(text)
		}
	}
	forest := pforest.Intersection(f1, f2)
	fmt.Println(forest.Generate(1000))
}
