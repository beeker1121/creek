# Creek

A simple log rotator for Go on Linux platforms.

## Usage

Sample usage:

```go
package main

import (
	"fmt"
	"log"
	"creek"
)

func main() {
	fmt.Println("Starting test of creek logger...")

	logger := log.New(&creek.Logger{Filename: "/var/log/maildblog/http.log", MaxSize: 1}, "Logger: ", log.Lshortfile|log.LstdFlags)

	logger.Println("Testing the log file")
	logger.Println("Testing the log file again")

	fmt.Println("Log finished.")
}
```