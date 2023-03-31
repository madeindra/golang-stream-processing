package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/madeindra/golang-stream-processing/internal/model"
	"github.com/madeindra/golang-stream-processing/pkg/reader"
)

func main() {
	// get current directory
	_, dir, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed reading directory")
	}

	// map as temporary database
	stockData := make(map[string]model.StockDetail)

	// sync for goroutines
	var wg sync.WaitGroup
	var mtx sync.Mutex

	// read all files and pass the result to the function
	if err := reader.ReadAllFiles(filepath.Join(dir, "../../../data"), passLine(stockData, &wg, &mtx)); err != nil {
		fmt.Println("error: ", err)
	}

	wg.Wait()
	fmt.Println(stockData)
}

func passLine(stockData map[string]model.StockDetail, wg *sync.WaitGroup, mtx *sync.Mutex) func(lines string) error {
	// create error channel
	chErr := make(chan error)

	return func(lines string) error {
		wg.Add(1)

		// pass to goroutines
		go process(lines, stockData, chErr, wg, mtx)
		if err := <-chErr; err != nil {
			return err
		}

		return nil
	}
}

func process(lines string, stockData map[string]model.StockDetail, chErr chan<- error, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	// parse json
	data := &model.Data{}
	if err := json.Unmarshal([]byte(lines), data); err != nil {
		chErr <- err
		return
	}

	// retrive value
	stockDetail := stockData[data.Name]

	if stockDetail.Name == "" {
		stockDetail.Name = data.Name
	}

	stockDetail.Quantity += data.Quantity
	stockDetail.AveragePrice = float64(stockDetail.Quantity) / float64(data.Price)

	// modify map
	mtx.Lock()
	stockData[data.Name] = stockDetail
	mtx.Unlock()

	chErr <- nil
}
