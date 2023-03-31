package reader

import (
	"bufio"
	"os"
	"path/filepath"
)

// ReadAllFiles reads all lines and run a function to each line
func ReadAllFiles(dir string, fn func(string) error) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// loop over the files
	for _, file := range files {
		// open the file
		fileReader, err := os.Open(filepath.Clean(filepath.Join(dir, file.Name())))
		if err != nil {
			return err
		}
		defer closeReader(fileReader)

		// read each line of the file
		scanner := bufio.NewScanner(fileReader)
		scanner.Split(bufio.ScanLines)

		// loop over the lines of the file
		for scanner.Scan() {
			if err := fn(scanner.Text()); err != nil {
				return err
			}
		}
	}

	return nil
}

// closeReader close the file reader and panic to prevent leak
func closeReader(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}
