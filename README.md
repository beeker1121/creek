# Creek

A simple log rotator for Go on Linux platforms.

## Usage

Sample usage:

```go
package main

import (
	"log"

	"creek"
)

func main() {
	// Create a new logger.
	logger := log.New(creek.New("/var/log/your_app/http.log", 1), "Logged: ", log.Lshortfile|log.LstdFlags)

	// Print to the log.
	logger.Println("Testing the log file")
}
```