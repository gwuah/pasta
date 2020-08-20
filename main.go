package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func processFile(path string) bool {
	return true
}

func main() {
	fmt.Println("[PASTA]")

	// find all files in repo
	// find all test files (.spec.tsx)
	// filter out actual typescript files (.tsx)
	// filter all snapshot files (.spec.tsx.snap)

	var files []string
	var components []string
	var testFiles []string
	var snapshots []string
	var blackList = map[string]bool{
		"/Users/user/Desktop/work/petra/web/src/setupTests.js":      true,
		"/Users/user/Desktop/work/petra/web/src/serviceWorker.ts":   true,
		"/Users/user/Desktop/work/petra/web/src/react-app-env.d.ts": true,
	}

	pathToCodebase := "/Users/user/Desktop/work/petra/web/src"

	err := filepath.Walk(pathToCodebase, func(path string, info os.FileInfo, err error) error {
		if _, ok := blackList[path]; ok {
			return nil
		}

		if strings.HasSuffix(path, ".spec.tsx") {
			testFiles = append(testFiles, path)
		} else if strings.HasSuffix(path, ".spec.tsx.snap") {
			snapshots = append(snapshots, path)
		} else if strings.HasSuffix(path, ".tsx") || strings.HasSuffix(path, ".ts") {
			components = append(components, path)
		} else {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	// fmt.Println(testFiles)
	// fmt.Println(components)
	// fmt.Println(snapshots)

	for _, file := range components {
		fmt.Println(file)
	}

}
