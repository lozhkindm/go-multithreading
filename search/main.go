package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches []string
	wg      = sync.WaitGroup{}
	lock    = sync.Mutex{}
)

func main() {
	wg.Add(1)
	search("~/", "boid.go")
	wg.Wait()
	for _, file := range matches {
		log.Println("Found a file:", file)
	}
}

func search(root, filename string) {
	log.Printf("Searching for %s in %s", filename, root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		path := filepath.Join(root, file.Name())
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, path)
			lock.Unlock()
		} else if file.IsDir() {
			wg.Add(1)
			go search(path, filename)
		}
	}
	wg.Done()
}
