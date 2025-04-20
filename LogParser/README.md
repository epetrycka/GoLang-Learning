## Zadanie: Log Parser CLI z własnym interfejsem
Zrobisz idiomatyczną, testowalną aplikację do filtrowania i analizy logów.
W pełni modularną, z interfejsami i testami. ✨

### Struktura projektu
 main.go – wejście CLI

 parser.go – logika parsowania logów

 filter.go – logika filtrowania + Twój własny interfejs

 types.go – definicje struktur LogEntry, LogFilter

 parser_test.go, filter_test.go – testy jednostkowe

 go.mod – moduł Go

### To Do
 Wypisz statystyki (count per service, count per level)

 Zapisz dane w formacie json

### Testy
 Przetestuj FilterLogs() z różnymi filtrami

 Przetestuj parser logów na poprawnych i błędnych liniach JSON

 Użyj strings.NewReader i bytes.Buffer do testowania io.Reader/io.Writer

## Plus
 Zaimplementuj interfejs MultiFilter – który przyjmuje wiele filtrów naraz

 Obsłuż stdin jako źródło danych (os.Stdin)

## Wywołania CLI:
```
go run main.go --file aws_style_logs.jsonl --level ERROR --service EC2

go build -o logparser
sudo mv logparser /usr/local/bin/
```