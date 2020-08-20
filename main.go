package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("[PASTA]")

	// find all files in repo
	// find all test files (.spec.tsx)
	// filter out actual typescript files (.tsx)
	// filter all snapshot files (.spec.tsx.snap)

	var files []string
	pathToCodebase := "/Users/user/Desktop/work/petra/web/src"

	err := filepath.Walk(pathToCodebase, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(files)

}
