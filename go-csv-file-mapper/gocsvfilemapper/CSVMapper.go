package gocsvfilemapper

/**
 * GO CSV File Mapper
 * 1. Read CSV Into Memory
 * 2. Convert to 2D Array
 * 3. Filter via the 2D array
 */

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CSVMapper struct {
	path         string
	data         [][]string
	Headers      []string
	Len          int
	filteredData [][]string
}

func NewCSVFilter(path string) (*CSVMapper, error) {

	if filepath.Ext(path) != ".csv" {
		return nil, errors.New("Extension is not CSV")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("Filepath invalid")
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to open file")
	}

	defer file.Close()

	var headers []string
	var data [][]string

	reader := csv.NewReader(file)
	//
	// if err != nil {
	// 	fmt.Println("Reader failed")
	// 	fmt.Println(err)
	// }
	//
	// length := len(records)
	//
	h := []string{}

	hFilled := false
	count := 0
	for {
		row, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}

		if !hFilled {
			for _, val := range row {
				h = append(h, val)
			}
			hFilled = true
			headers = h
			continue
		}
		var rowString []string
		for _, element := range row {

			rowString = append(rowString, element)
		}

		// fmt.Println(row)
		data = append(data, rowString)
		count++
	}
	fmt.Printf("\nNumber of rows: %d", len(data))

	// Only check if the path exists and if the extension is CSV
	return &CSVMapper{
		path:         path,
		Headers:      headers,
		Len:          count,
		data:         data,
		filteredData: data,
	}, nil
}

func (m *CSVMapper) Filter(index int, filter string) {
	var filtered [][]string

	fmt.Printf("\nFiltering for :'%s' at %d\n", filter, index)

	for i, _ := range m.filteredData {

		// fmt.Println(m.filteredData[i][index], filter)
		if strings.Contains(m.filteredData[i][index], filter) {

			fmt.Println(m.filteredData[i][index], filter)
			filtered = append(filtered, m.filteredData[i])
			fmt.Println(m.filteredData[i])
		}

	}

	fmt.Printf("\nFiltered Rows Count: %d\n", len(filtered))
	m.filteredData = filtered
}

func (m *CSVMapper) reset() {
	m.filteredData = m.data
}

func (m *CSVMapper) WriteToCSV(path string) error {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("Unable to create file")
		fmt.Println(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(m.Headers); err != nil {
		fmt.Println("Unable to write headers")
		fmt.Println(err)

		return fmt.Errorf("%v", err)
	}

	for _, row := range m.filteredData {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
