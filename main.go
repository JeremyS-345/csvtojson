package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	fd, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	records, err := csv.NewReader(fd).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	columns := records[0]
	resultSlice := make([]map[string]any, 0)
	for _, record := range records {
		jsonMap := make(map[string]any)
		for i := range record {
			jsonMap[columns[i]] = record[i]
		}
		fmt.Println(jsonMap)
		resultSlice = append(resultSlice, jsonMap)
	}

	jsonBytes, err := json.MarshalIndent(resultSlice,"","	")
	if err != nil {
		log.Fatal(err)
	}

	os.Remove(fileName)
	filenameRunes := []rune(fileName)
	filenameRunes = filenameRunes[:len(filenameRunes)-4]
	os.WriteFile(string(filenameRunes)+".json", jsonBytes, os.ModeDevice)
}
