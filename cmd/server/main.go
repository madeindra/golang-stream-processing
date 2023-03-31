package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/madeindra/golang-stream-processing/internal/model"
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
		fmt.Println("error: ", err)
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

func process(lines string, chErr chan<- error) {
	// parse json
	data := &model.Data{}
	if err := json.Unmarshal([]byte(lines), data); err != nil {
		chErr <- err
		return
	}

	// print parsed line of data
	fmt.Println(data)
	chErr <- nil
}
