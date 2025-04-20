## Log Parser CLI

A modular and idiomatic Go CLI tool for parsing and filtering JSONL logs.

## Features

- Filter logs by attributes like `level`, `service`, etc.
- Supports custom `LogFilter` and `MultiFilter` interfaces
- Counts unique values per field (e.g., `EC2: 123`)
- Saves filtered logs as JSON
- Handles stdin input and error cases idiomatically

## Usage

```bash
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