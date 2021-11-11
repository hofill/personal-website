package utils

import (
	"bufio"
	"errors"
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

func GetWriteUpMetaDataAndMD(filePath string) (map[string]interface{}, string, error) {
	metaData := make(map[string]interface{})
	var mdLines string
	lines, err := getLinesFromFile(filePath)
	if err != nil {
		return nil, "", err
	}
	var delimiter = "---"
	var counterDelimiter = 0
	for _, line := range lines {
		if line == delimiter && counterDelimiter < 2 {
			counterDelimiter++
			continue
		}
		if counterDelimiter < 2 {
			key, value, err := parseWriteUpMetaData(line)
			if err != nil {
				return nil, "", err
			}
			metaData[key] = value
			continue
		}
		// Load up MD when finished loading metadata
		mdLines += line + "\n"
	}

	return metaData, mdLines, nil
}

func parseWriteUpMetaData(line string) (string, interface{}, error) {
	pair := strings.Split(line, ":")
	if len(pair) != 2 {
		return "", nil, errors.New("corrupted metadata file")
	}

	key, value := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])

	var newValue interface{}
	var err error

	if strings.HasPrefix(key, "&") {
		newValue, err = strconv.ParseInt(value, 10, 0)
		key = key[1:]
	}
	if strings.HasPrefix(key, "@") {
		newValue, err = strconv.ParseBool(value)
		key = key[1:]
	}
	if err != nil || newValue == nil {
		return key, value, nil
	}
	return key, newValue, nil
}

func GetWriteUp(filePath string) ([]string, error) {
	lines, err := getLinesFromFile(filePath)
	if err != nil {
		return nil, err
	}
	var delimiter = "---"
	var counterDelimiter = 0
	var writeUp []string
	for _, line := range lines {
		if line == delimiter {
			counterDelimiter += 1
		}
		if counterDelimiter < 2 {
			continue
		}

		writeUp = append(writeUp, line)

	}
	if counterDelimiter < 2 {
		return nil, errors.New("corrupted writeup")
	}
	return writeUp, nil
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
