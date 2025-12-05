package data_parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// GetConfigFile returns the path to the config file.
func GetConfigFile() string {
	homeDir, err := os.UserHomeDir()
	if err!= nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".data-parser.json")
}

// LoadConfig loads the config from a file.
func LoadConfig() (*Config, error) {
	configFile := GetConfigFile()
	data, err := ioutil.ReadFile(configFile)
	if err!= nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	return &config, err
}

// SaveConfig saves the config to a file.
func SaveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err!= nil {
		return err
	}
	return ioutil.WriteFile(GetConfigFile(), data, 0644)
}

// Config is the config for the data parser.
type Config struct {
	// File path to the data file.
	DataFilePath string `json:"data_file_path"`
	// The format of the data file.
	DataFileFormat string `json:"data_file_format"`
	// The format of the data.
	DataFormat string `json:"data_format"`
	// The column names.
	ColumnNames []string `json:"column_names"`
	// The column types.
	ColumnTypes []string `json:"column_types"`
	// The column lengths.
	ColumnLengths []int `json:"column_lengths"`
	// The date format.
	DateFormats []string `json:"date_formats"`
}

// ValidateConfig validates the config.
func ValidateConfig(config *Config) error {
	if config.DataFilePath == "" {
		return fmt.Errorf("data file path is required")
	}
	if config.DataFileFormat == "" {
		return fmt.Errorf("data file format is required")
	}
	if config.DataFormat == "" {
		return fmt.Errorf("data format is required")
	}
	if len(config.ColumnNames) == 0 {
		return fmt.Errorf("column names are required")
	}
	if len(config.ColumnTypes)!= len(config.ColumnNames) {
		return fmt.Errorf("column types must match column names")
	}
	if len(config.ColumnLengths)!= len(config.ColumnNames) {
		return fmt.Errorf("column lengths must match column names")
	}
	for _, dateFormat := range config.DateFormats {
		_, err := time.Parse(dateFormat, "2022-01-01")
		if err!= nil {
			return err
		}
	}
	return nil
}

// ParseCSV parses a CSV file.
func ParseCSV(filePath string) ([]map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err!= nil {
		return nil, err
	}
	defer file.Close()
	return parseCSV(file)
}

// parseCSV parses a CSV file.
func parseCSV(file *os.File) ([]map[string]interface{}, error) {
	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ","))
	}
	if err := scanner.Err(); err!= nil {
		return nil, err
	}
	columnNames := lines[0]
	data := make([]map[string]interface{}, len(lines)-1)
	for i, line := range lines[1:] {
		row := make(map[string]interface{})
		for j, value := range line {
			switch columnTypes[j] {
			case "string":
				row[columnNames[j]] = value
			case "int":
				row[columnNames[j]], _ = strconv.Atoi(value)
			case "float":
				row[columnNames[j]], _ = strconv.ParseFloat(value, 64)
			case "time":
				row[columnNames[j]], _ = time.Parse(dateFormats[j], value)
			default:
				return nil, fmt.Errorf("unknown column type: %s", columnTypes[j])
			}
		}
		data[i] = row
	}
	return data, nil
}

// parseCSVFile parses a CSV file.
func parseCSVFile(filePath string) ([]map[string]interface{}, error) {
	return ParseCSV(filePath)
}

func getColumns() []string {
	return []string{"column1", "column2", "column3"}
}

func getTypes() []string {
	return []string{"string", "int", "time"}
}

func getLengths() []int {
	return []int{10, 10, 10}
}

func getDateFormats() []string {
	return []string{"2006-01-02"}
}

func getColumnType(column string) string {
	switch column {
	case "column1":
		return "string"
	case "column2":
		return "int"
	case "column3":
		return "time"
	default:
		return ""
	}
}

func getColumnTypeValue(column string) interface{} {
	switch getColumnType(column) {
	case "string":
		return "value"
	case "int":
		return 10
	case "time":
		return time.Now()
	default:
		return ""
	}
}

func getColumnTypeString(column string) string {
	return getColumnType(column)
}

func getColumnTypeValueString(column string) string {
	return fmt.Sprintf("%v", getColumnTypeValue(column))
}