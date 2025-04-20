package Logic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"log"
	"strconv"
	"github.com/urfave/cli/v2"
)

type LogEntry map[string]interface{}

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

	level := c.String("level")
	service := c.String("service")
	if level != ""{
		var filter LogFilter = &LevelFilter{logLevel: level}
		SaveFilterLogs(logEntryList, filter)
	} else if service != ""{
		var filter LogFilter = &ServiceFilter{Service: service}
		SaveFilterLogs(logEntryList, filter)
	} else {
		var filter LogFilter = nil
		SaveFilterLogs(logEntryList, filter)
	}
	
	return nil
}

func SaveFilterLogs(entries []LogEntry, filter LogFilter) error{
	var filterEntries []LogEntry
	filterEntries, err := FilterLogs(entries, filter)

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

	for key, value := range filterEntries {
		_, err := writer.WriteString("LogEntry: " + strconv.Itoa(key) + "\n")
		if err != nil {
			fmt.Println("error occurred during saving first line:", err)
			return err
		}

		for subKey, subValue := range value {
			_, err := writer.WriteString(fmt.Sprintf("  %s: %v\n", subKey, subValue))
			if err != nil {
				fmt.Println("error occurred during saving description:", err)
				return err
			}
		}
		_, err = writer.WriteString("---\n")
		if err != nil {
			fmt.Println("error occurred during saving delimiter:", err)
			return err
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