# Go Logging System

A simple and extensible logging system implemented in Go. This system supports log levels, multiple handlers (for writing logs to stdout or files), formatters for custom log outputs, and filters for selective logging.

## Features

- **Log Levels**: `DEBUG`, `INFO`, `ERROR`, `WARNING`, `CRITICAL`
- **Handlers**:
  - `StreamHandler`: Writes logs to `stdout`.
  - `FileHandler`: Writes logs to a specified log file.
- **Formatters**: Customizable log formatting.
- **Filters**: Apply custom filtering on log records before they are written.
- **Thread Safety**: Can be extended to support concurrent logging with mutex locking.

## Usage

Below is an example of how to use the Go logging system. This script sets up a logger, adds handlers, configures a formatter, and logs messages at different levels.

### Full Example Script

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a logger instance
	logger := GetLogger("test")
	logger.SetLevel(INFO)

	// Create a StreamHandler that writes to stdout
	std := GetStreamHandler()
	std.SetLogLevel(DEBUG)
	logger.AddHandlers(std)

	// Create a FileHandler that writes to a file
	fileHandler := GetFileHandler("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileHandler.SetLogLevel(DEBUG)
	logger.AddHandlers(&fileHandler)  // Pass a pointer to FileHandler

	// Set a custom formatter for the logs
	formatter := &StdFormatter{}
	std.SetFormatter(formatter)

	// Log messages at different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Critical("This is a critical message")

	// Close the file handler to release resources
	fileHandler.Close()
}
```

### 1. **Create a Logger**

To create a logger and set its logging level, use the `GetLogger` function and `SetLevel` method:

```go
logger := GetLogger("test")
logger.SetLevel(INFO)
```

### 2. **Add Handlers**

You can use different handlers to output logs either to the console or to a log file.

#### StreamHandler (stdout):

This handler writes logs to the standard output (`stdout`):

```go
std := GetStreamHandler()
std.SetLogLevel(DEBUG)
logger.AddHandlers(std)
```

#### FileHandler (Log file):

This handler writes logs to a specified log file. It requires you to provide the file path and permission settings:

```go
fileHandler := GetFileHandler("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
fileHandler.SetLogLevel(DEBUG)
logger.AddHandlers(&fileHandler)  // Pass a pointer to FileHandler
```

### 3. **Set a Formatter**

The `Formatter` defines how log records are formatted. The default formatter is `StdFormatter`, but you can create and use your own formatter if needed.

```go
formatter := &StdFormatter{}
std.SetFormatter(formatter)
```

### 4. **Log Messages**

Once the logger is set up, you can start logging messages at different levels. Example log levels include `DEBUG`, `INFO`, `WARNING`, `ERROR`, and `CRITICAL`:

```go
logger.Debug("This is a debug message")
logger.Info("This is an info message")
logger.Warning("This is a warning message")
logger.Critical("This is a critical message")
```

### 5. **Close Handlers**

To prevent resource leakage, make sure to close any file handlers when you're done logging:

```go
fileHandler.Close()
```