package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/madeindra/golang-stream-processing/pkg/reader"
)

func main() {
	// get current directory
	_, dir, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed reading directory")
	}

	// read all files and pass the result to the function
	if err := reader.ReadAllFiles(filepath.Join(dir, "../../../data"), passLine); err != nil {
		panic(err)
	}
}

func passLine(lines string) error {
	// create error channel
	chErr := make(chan error)
	defer close(chErr)

	// pass to goroutines
	go process(lines, chErr)
	if err := <-chErr; err != nil {
		return err
	}

	return nil
}

func process(lines string, err chan<- error) {
	// print content of lines
	fmt.Println(lines)
	err <- nil
}
