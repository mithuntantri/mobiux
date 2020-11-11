package main

import (
	"os"
	"encoding/csv"
)

// ReadDataset accepts a file and returns its content as a multi-dimensional type
// with lines and each column. Only parses to string type.
func ReadDataset(filename string) ([][]string, error) {

    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}