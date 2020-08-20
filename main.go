package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gwuah/pasta/lib"
)

func main() {

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

	var countStream = make(chan int)

	for _, file := range components {
		go lib.ProcessFile(file, countStream)
	}

	var count = 0
	for i := 0; i < len(components); i++ {
		count += <-countStream
	}

	fmt.Println(count)

}
