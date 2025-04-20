package Logic

type LogFilter interface {
    Match(entry LogEntry) bool
}

type LevelFilter struct {
	logLevel string
}

func (f *LevelFilter) Match(entry LogEntry) bool {
	val, ok := entry["logLevel"].(string)
	return ok && val == f.logLevel
}

type ServiceFilter struct {
	Service string
}

func (f *ServiceFilter) Match(entry LogEntry) bool {
	val, ok := entry["service"].(string)
	return ok && val == f.Service
}

func FilterLogs(entries []LogEntry, filter LogFilter) ([]LogEntry, error){
	if filter == nil {
		return entries, nil
	}
	var filtered []LogEntry
	for _, entry := range entries {
		if filter.Match(entry) {
			filtered = append(filtered, entry)
		}
	}
	return filtered, nil
}