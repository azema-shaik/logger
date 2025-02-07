# Logger

A simple and flexible logging system for Go.

## Features

- Multiple log levels: DEBUG, INFO, WARNING, ERROR, CRITICAL
- Customizable formatters
- Multiple handlers: StreamHandler, FileHandler
- Filter support for log records
- Logger hierarchy with propagation


## Basic Usage

```go
package main

import (
    "fmt"
    "os"

    logging "github.com/azema-shaik/logger/logger"
)

func main() {
    logger := logging.GetLogger("example")
    logger.SetLevel(logging.DEBUG)

    // StreamHandler
    streamHandler, _ := logging.GetStreamHandler()
    streamHandler.SetLogLevel(logging.DEBUG)
    logger.AddHandler(streamHandler)

    // FileHandler
    fileHandler, _ := logging.GetFileHandler("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    fileHandler.SetLogLevel(logging.DEBUG)
    logger.AddHandler(fileHandler)

    logger.Debug("This is a debug message")
    logger.Info("This is an info message")
    logger.Warning("This is a warning message")
    logger.Error("This is an error message")
    logger.Critical("This is a critical message")

    logging.Close()
}
```

## Custom Formatter
The following format strings are available for customizing log messages:
| Format String    | Description |
|-----------------|-------------|
| `%(asctime)s`   | The time when the log record was created. |
| `%(levelname)s` | The log level name (e.g., DEBUG, INFO). |
| `%(levelno)d`   | The log level number. |
| `%(funcName)s`  | The name of the function from which the log message originated. |
| `%(lineno)d`    | The line number in the source code where the log message originated. |
| `%(name)s`      | The name of the logger. |
| `%(Lfilename)s` | The full path of the source file where the log message originated. |
| `%(Sfilename)s` | The short name of the source file where the log message originated. |
| `%(msg)s`       | The log message. |

```go
package main

import (
    "os"

    logging "github.com/azema-shaik/logger/logger"
)

func main() {
    logger := logging.GetLogger("example")
    logger.SetLevel(logging.DEBUG)

    // Custom Formatter
    formatter := &logging.StdFormatter{}
    formatter.SetFormatter("[%(asctime)s] [%(levelname)s] %(message)s", "2006-01-02 15:04:05")

    // StreamHandler with custom formatter
    streamHandler, _ := logging.GetStreamHandler()
    streamHandler.SetLogLevel(logging.DEBUG)
    streamHandler.SetFormatter(formatter)
    logger.AddHandler(streamHandler)

    logger.Debug("This is a debug message with custom formatter")
    logger.Info("This is an info message with custom formatter")

    logging.Close()
}
```

## Using Filters
```go
package main

import (
    "os"

    logging "github.com/azema-shaik/logger/logger"
)

type CustomFilter struct{}

func (f *CustomFilter) Filter(record logging.LogRecord) bool {
    // Only log messages that contain the word "important"
    return strings.Contains(record.Message, "important")
}

func main() {
    logger := logging.GetLogger("example")
    logger.SetLevel(logging.DEBUG)

    // StreamHandler
    streamHandler, _ := logging.GetStreamHandler()
    streamHandler.SetLogLevel(logging.DEBUG)
    logger.AddHandler(streamHandler)

    // Add custom filter
    filter := &CustomFilter{}
    streamHandler.AddFilter(filter)

    logger.Debug("This is an important debug message")
    logger.Info("This is an unimportant info message")

    logging.Close()
}
```


## Using with Goroutines

To use the logging system with goroutines, you can create a logger and share it among multiple goroutines. Here is an example of how to do this:

```go
package main

import (
    "fmt"
    "os"
    "sync"

    logging "github.com/azema-shaik/logger/logger"
)

func main() {
    logger := logging.GetLogger("example")
    logger.SetLevel(logging.DEBUG)

    // StreamHandler
    streamHandler, _ := logging.GetStreamHandler()
    streamHandler.SetLogLevel(logging.DEBUG)
    logger.AddHandler(streamHandler)

    // FileHandler
    fileHandler, _ := logging.GetFileHandler("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    fileHandler.SetLogLevel(logging.DEBUG)
    logger.AddHandler(fileHandler)

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            logger.Debug(fmt.Sprintf("Goroutine %d: This is a debug message", id))
            logger.Info(fmt.Sprintf("Goroutine %d: This is an info message", id))
            logger.Warning(fmt.Sprintf("Goroutine %d: This is a warning message", id))
            logger.Error(fmt.Sprintf("Goroutine %d: This is an error message", id))
            logger.Critical(fmt.Sprintf("Goroutine %d: This is a critical message", id))
        }(i)
    }

    wg.Wait()
    logging.Close()
}
```