package Logic

import "fmt"

// type LogFilter interface {
//     Match(entry LogEntry) bool
// }

// type LevelFilter struct {
// 	logLevel string
// }

// func (f *LevelFilter) Match(entry LogEntry) bool {
// 	val, ok := entry["logLevel"].(string)
// 	return ok && val == f.logLevel
// }

// type ServiceFilter struct {
// 	Service string
// }

// func (f *ServiceFilter) Match(entry LogEntry) bool {
// 	val, ok := entry["service"].(string)
// 	return ok && val == f.Service
// }

func (entry LogEntry) IfMatching (attribute string, value string) bool {
	val, ok := entry[attribute].(string)
	return ok && val == value
} 

func FilterLogs(entries []LogEntry) ([]LogEntry, error){
	if len(Attributes) == 0 {
		return entries, nil
	}

	counts := make(map[string]map[string]int)
	var filtered []LogEntry
	var matching bool = true
	for _, entry := range entries {
		matching = true
		for attribute, value := range Attributes{
			if value == ""{
				continue
			}
			if !entry.IfMatching(attribute, value)  {
				matching = false
			}
			if uniqueValue, ok := entry[attribute]; ok {
				if attrValue, ok := uniqueValue.(string); ok {
					if _, exists := counts[attribute]; !exists {
						counts[attribute] = make(map[string]int)
					}
					counts[attribute][attrValue]++
				}
			}
		}
		if matching {
			filtered = append(filtered, entry)
		}
	}

	for attribute, valueMap := range counts {
		fmt.Println("Attribute:", attribute)
		for value, count := range valueMap {
			fmt.Printf("  %s: %d\n", value, count)
		}
	}
	
	return filtered, nil
}