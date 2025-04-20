package Logic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

type LogEntry map[string]interface{}

var Attributes = map[string]string{
	"logLevel" : "",
	"service" : "",
}

func Parse(c *cli.Context) error {
	var logEntryList []LogEntry
	name := c.String("file")
	
	file, err := os.Open(name)
	if err != nil{
		return fmt.Errorf("error occurred during opening a file %s: %w", name, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		var logEntry LogEntry
		if err := json.Unmarshal([]byte(line), &logEntry); err != nil{
			log.Print("error during parsing json log: ", err)
			continue
		}
		logEntryList = append(logEntryList, logEntry)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error during scanning occurred: %w", err)
	}

	for attribute, _ := range Attributes{
		value := c.String(attribute)
		if value != ""{
			Attributes[attribute] = value
		}
	}

	SaveFilterLogs(logEntryList)

	return nil
}

func SaveFilterLogs(entries []LogEntry) error{
	var filterEntries []LogEntry
	filterEntries, err := FilterLogs(entries)

	if err != nil{
		return fmt.Errorf("error occurred during filtering logs: %w", err)
	}

	output_file := "Data/filtered_logs.jsonl" 
	file, err := os.OpenFile(output_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error occurred during opening output file:", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, entry := range filterEntries {
		jsonBytes, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("error marshalling JSON: %w", err)
		}

		_, err = writer.WriteString(string(jsonBytes) + "\n")
		if err != nil {
			return fmt.Errorf("error writing line: %w", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("error occurred during flush:", err)
		return err
	}
	fmt.Println("Data written successfully to file")

	return nil
}