package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	scanner := bufio.NewScanner(file)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines, nil
}

func GetWriteUpMetaData(filePath string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	lines, err := getLinesFromFile(filePath)
	if err != nil {
		return nil, err
	}
	var delimiter = "---"
	var counterDelimiter = 0
	for _, line := range lines {
		if line == delimiter {
			counterDelimiter++
			continue
		}
		if counterDelimiter == 2 {
			break
		}

		pair := strings.Split(line, ":")
		if len(pair) != 2 {
			return nil, errors.New("corrupted metadata file")
		}

		key, value := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])

		var newValue interface{}

		if strings.HasPrefix(key, "&") {
			newValue, err = strconv.ParseInt(value, 10, 0)
			key = key[1:]
		}
		if strings.HasPrefix(key, "@") {
			newValue, err = strconv.ParseBool(value)
			key = key[1:]
		}
		fmt.Println(newValue, value, err)
		if err != nil || newValue == nil {
			result[key] = value
			continue
		}
		result[key] = newValue
	}

	return result, nil
}

func GetEventMetaData(filePath string) (map[string]string, error) {
	result := make(map[string]string)
	lines, err := getLinesFromFile(filePath)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			return nil, errors.New("corrupted metadata file")
		}
		key, value := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])
		result[key] = value
	}

	return result, nil
}
