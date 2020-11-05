package main

import (
	"encoding/json"
	"io/ioutil"
)

func readFile(filePath string) ([]byte, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Errorf("can`t read file %v", err)
		return nil, err
	}
	return file, nil
}

func parseInitValues(values []byte) ([]int, error) {
	var valuesInt []int
	err := json.Unmarshal(values, &valuesInt)
	if err != nil {
		logger.Errorf("can`t parse json %v", err)
		return nil, err
	}
	return valuesInt, nil
}

func getInitValues(filePath string) []int {
	fileByte, err := readFile(filePath)
	if err != nil {
		logger.Errorf("read file problem")
	}
	intArr, err := parseInitValues(fileByte)
	if err != nil {
		logger.Errorf("parse file problem")
	}
	return intArr
}
