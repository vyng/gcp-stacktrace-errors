# stacktrace-gcp
Go Errors With Stacktraces for Google Cloud Error Reporting

# Documentation
* [API Reference](https://godoc.org/github.com/vyng/gcp-stacktrace-errors)

# Example
```go
	errorReporter, err := errorreporting.NewClient(ctx, project, config)
	if err != nil {
		panic(err)
	}
	
	// Set up loggerLevels
	
	reporter = &LogReporter{
		errorReporter: errorReporter,
		loggerLevels: loggerLevels,
	}
	
	reporter.errorReporter.Report(WrapError(errors.New('this is an error'), 1))
````