## Log Parser CLI

A modular and idiomatic Go CLI tool for parsing and filtering JSONL logs.

## Features

- Parses JSON logs line by line from a given file (one JSON object per line)
- Works regardless of the structure of each JSON log entry
- Stores all logs as a list of dynamic entries (slice)
- Filters logs by multiple attributes at once (e.g., `--level ERROR --service EC2`)
- Stores filters globally for easy access across the app
- Calculates metrics: counts unique values for each filtered attribute
- Saves filtered logs to a JSON file

## Usage

```
go run main.go --file logs.jsonl --level ERROR --service EC2
```

## TO DO:
- Tests :c

## Notes:

```
go run main.go --file aws_style_logs.jsonl --level ERROR --service EC2

go build -o logparser
sudo mv logparser /usr/local/bin/
```