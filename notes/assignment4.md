# Chapter 2 (cont.)
## 2.10. The http.Handler interface
A handler is just an object that implements the `http.Handler` interface:
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
Creating an object just to satisfy a single method is a long-winded way to define a function, so
instead we use the `http.HandlerFunc` method to transform a function into a handler.

This is a simple interface, but it is powerful. Any handler can be chained to add additional 
functionality as long as it satisfies the `http.Handler` interface. In reality, this is how a
servemux and middleware work in Go.

One last thing to note: all requests in Go are served in a separate goroutine. This provides
performance, but requires developers to keep concurrency in mind.

# Chapter 3. Configuration and error handling
## 3.1. Managing configuration settings
Go standard library provides a `flag` package that allows you to parse command-line arguments.
Flags are typed and optional with a default value. The `-help` flag is automatically provided
and will display the usage of specified flags. Configuration values can also be fetched from
environment variables.

## 3.2. Structured logging
Both the `log.Printf()` and `log.Fatal()` functions output log entries using Go’s standard logger.
For human-readable needs that's great, but that will not play nice with log aggregation tools. Go
provides a `log/slog` package for structured logging needs. This supports log levels, timestamps,
key-value pairs, stack traces, and more.

The structured logger is configured on instantiation and `Debug()`, `Info()`, `Warn()`, and 
`Error()` functions can be used to create a log of the corresponding level. Each accepts a 
"variadic" number of arguments, which can be key-value pairs.

Instead of using strings for all key-value pairs, there are typed equivalents provided in the `slog`
package. Use the `slog.String()`, `slog.Int()`, `slog.Bool()`, `slog.Time()` and `slog.Duration()` 
functions to create attributes with a specific type of value.

To make logs output in JSON, use the `slog.NewJSONHandler` in the instantiation of the logger. This
will play nice with log aggregation tools like Datadog.

Configuration of the logger includes minimum log level and outputting log caller location.

## 3.3. Dependency Injection
Dependency injection is a useful pattern for sharing dependencies between different parts of the 
application. Using a struct to hold dependencies, and defining receiver functions for that struct
is a common pattern in Go.

The single "application struct" pattern  won’t work if handlers are spread across multiple packages.
In that case, create a standalone config package that exports an Application struct, then handler 
functions close over this to form a closure.

## 3.4. Centralized error handling
Centralizing error messaging and handling in "helper" functions standardizes error messages and
simplifies handler code.

Another good practice is to include the stack trace in error messages. 

## 3.5. Isolating the application routes
It's a good practice to separate application routing into a separate file. This encapsulates routing
and reduces the responsibility of the main function.
