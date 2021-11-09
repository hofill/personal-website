package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func getLinesFromFilePath(filePath string) ([]string, error) {
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

func GetWriteUpMetaData(filePath string) (map[string]string, error) {
	result := make(map[string]string)
	lines, err := getLinesFromFilePath(filePath)
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
		result[key] = value
	}

	return result, nil
}

func GetEventMetaData(filePath string) (map[string]string, error) {
	result := make(map[string]string)
	lines, err := getLinesFromFilePath(filePath)
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
