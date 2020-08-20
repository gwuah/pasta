package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gwuah/pasta/lib"
)

type FileConfig struct {
	files   []string
	stream  chan int
	counter int
}

func processFiles(container *FileConfig) func() {
	return func() {
		for _, file := range container.files {
			go lib.GetLocStats(file, container.stream)
		}
	}
}

func aggregateDataFromStream(config *FileConfig, wg *sync.WaitGroup) {
	for i := 0; i < len(config.files); i++ {
		config.counter += <-config.stream
	}
	wg.Done()
}

func main() {

	var components []string
	var tests []string
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
			tests = append(tests, path)
		} else if strings.HasSuffix(path, ".spec.tsx.snap") {
			snapshots = append(snapshots, path)
		} else if strings.HasSuffix(path, ".tsx") || strings.HasSuffix(path, ".ts") {
			components = append(components, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	var start = time.Now()
	var wg sync.WaitGroup

	configs := []*FileConfig{
		{files: components, stream: make(chan int)},
		{files: tests, stream: make(chan int)},
		{files: snapshots, stream: make(chan int)},
	}

	for _, config := range configs {
		go processFiles(config)()
	}

	for _, config := range configs {
		wg.Add(1)
		go aggregateDataFromStream(config, &wg)
	}

	wg.Wait()

	var elapsed = time.Since(start)
	var totalNumberOfFiles = len(components) + len(tests) + len(snapshots)

	fmt.Println("\nPasta la vista 🍝")

	fmt.Printf("Processed %d files under %s\n\n", totalNumberOfFiles, elapsed)
	fmt.Printf("Components :=> %d lines of code.\n", configs[0].counter)
	fmt.Printf("Tests :=> %d lines of code.\n", configs[1].counter)
	fmt.Printf("Snapshots :=> %d lines of code.\n", configs[2].counter)

}
