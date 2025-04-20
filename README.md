## GoLang Learning Projects
This repository contains small Go programs created as part of my learning journey. Each stage focuses on a different aspect of the Go language, from fundamentals to advanced DevOps and cloud integrations.

### Stage 1: Go Language Fundamentals
Learned: Syntax, variables, functions, structs, slices/maps, project organization
Project: ToDo CLI App — manage tasks via console with add/remove/display features.

### Stage 2: Idiomatic Go & Practical Use
To learn: Error handling, interfaces, testing, Go modules
Project: File utility library with ReadLines and WriteLines, tested with go test.

### Stage 3: Concurrency in Go
To learn: Goroutines, channels, mutexes, sync.WaitGroup
Project: Web crawler — fetch multiple URLs in parallel and gather their content.

### Stage 4: Building Web Apps & Microservices
To learn: HTTP server, JSON, routing, database interaction
Project: notes-service REST API — create, list, and delete notes using SQLite/PostgreSQL.

### Stage 5: DevOps Tools & CLI in Go
Learned: Cross-compilation (GOOS, GOARCH), logging (logrus, zap), CLI design using flags and cobra, working with JSON output
Project: health-checker CLI tool

Checks the HTTP status and response time of multiple URLs

Supports a --json flag for JSON-formatted output

Built using the cobra library for clean CLI structure

Compiles into a standalone binary for various platforms.

### Stage 6: Cloud & Middleware Integration
To learn: AWS SDK, RabbitMQ/Kafka, Prometheus, serialization
Project: Microservice that subscribes to messages from RabbitMQ/Kafka, stores them, and exposes a /messages endpoint.

Add /metrics endpoint for Prometheus

### Stage 7: Final Project
To learn: Full-stack integration of previous topics
Project: A complete CLI or REST-based app using goroutines, API integration, file/stream handling, logging, monitoring, and Dockerized deployment.