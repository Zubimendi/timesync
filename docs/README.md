# TimeSync

TimeSync is a Go-based CLI tool designed to simplify scheduling across different time zones. Ideal for remote teams!

## Features
- Schedule meetings in different time zones.
- Manage team availability.
- Export schedules in various formats.

## Usage
Run the following command to get started:

```bash
timesync help



---

### `Makefile`
Provides common commands for building and running the project.

```makefile
.PHONY: build test run clean

build:
	go build -o bin/timesync ./cmd/timesync

test:
	go test ./...

run:
	go run ./cmd/timesync

clean:
	rm -rf bin/
