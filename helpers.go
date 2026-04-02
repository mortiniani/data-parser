package data_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Helper function to check if a file exists.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Helper function to read file content to string.
func readFileContent(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

// Helper function to convert string to int.
func stringToInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty string cannot be converted to integer")
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error converting string to integer: %w", err)
	}
	return i, nil
}

// Helper function to convert string to float.
func stringToFloat(s string) (float64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty string cannot be converted to float")
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting string to float: %w", err)
	}
	return f, nil
}

// Helper function to convert string to bool.
func stringToBool(s string) (bool, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return false, fmt.Errorf("empty string cannot be converted to bool")
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("error converting string to bool: %w", err)
	}
	return b, nil
}

// Helper function to check if a string is empty.
func isEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Helper function to convert a struct to JSON string.
func structToJSONString(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling struct to JSON: %w", err)
	}
	return string(jsonData), nil
}

// Helper function to check if a slice contains a specific element.
func stringSliceContains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

// Helper function to get the type of a variable as a string.
func getType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t == nil {
		return "nil"
	}
	return t.String()
}