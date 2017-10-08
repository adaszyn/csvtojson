//Package provides functions for converting
//.csv files to json format
package csvtojson

import (
	"io"
	"strings"
	"io/ioutil"
	"log"
	"encoding/json"
)

func ConvertCsvToJson(reader io.Reader, separator string) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal("Cannot read from reader")
	}
	array := CsvStringToArray(string(data), separator)
	arrayOfMaps := ArrayToArrayOfMaps(array)
	jsonString, err := json.Marshal(arrayOfMaps)
	return string(jsonString), err
}

func CsvStringToArray (csv string, separator string) [][] string {
	lines := strings.Split(csv, "\n")
	numOfLines := len(lines)

	array := make([][]string, numOfLines)
	for i, line := range lines {
		cells := strings.Split(strings.TrimSpace(line), separator)
		array[i] = cells
	}
	return array
}

func ArrayToArrayOfMaps(array [][] string) [] map[string] string {
	result := make([] map[string] string, len(array) - 1)
	headers := array[0]
	var i int
	var val [] string
	for i, val = range array[1:] {
		rowMap := make(map[string] string)
		for j, header := range headers {
			rowMap[header] = val[j]
		}
		result[i] = rowMap
	}
	return result
}