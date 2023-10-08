package main

import (
	"fmt"
	"os"
	"math/rand"
)

// SaveData into file at location path
func SaveData(path string, data []byte) error {
	// store data in temporary file
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp) // clean
		return err
	}

	// flush data immediately to disk to ensure
	// time of persistance is predictable
	if err = fp.Sync(); err != nil {
		os.Remove(tmp)
		return err
	}

	// rename is atomic, move data to destination
	return os.Rename(tmp, path)
}

// Append-Only Logs
func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
}

func LogAppend(fp *os.File, line string) error {
	buffer := []byte(line)
	buffer = append(buffer, '\n')
	_, err := fp.Write(buffer)
	if err != nil {
		return err
	}
	return fp.Sync()
}
