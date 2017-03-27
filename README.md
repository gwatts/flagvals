# Go Flag Value Types

[![GoDoc](https://godoc.org/github.com/gwatts/flagvals?status.svg)](https://godoc.org/github.com/gwatts/flagvals)
[![Build Status](https://travis-ci.org/gwatts/flagvals.svg?branch=master)](https://travis-ci.org/gwatts/flagvals)

Some handy Go types that can be used as values with the [Go flag package](https://golang.org/pkg/flag/)
as they implement the [flag.Value](https://golang.org/pkg/flag/#Value) interface.

# Range restricted integer values

RangeInt accepts an int64 value within a defined minimum and/or maximum range.

Values passed to the flag outside that range return an error.

```golang
// Specify a flag that accepts value that are at least 1 or higher
// with a default of 10
maxSize := flagvals.GTEInt(10, 1)
flag.Var(maxSize, "max-size", "")

// Specify a flag that accepts a value between 1 and 10 inclusive
// with a default of 4
concurrency := flagvals.BetweenInt(5, 1, 10)
flag.Var(concurrency, "concurrency", "")

flag.Parse()

fmt.Println("max size:", maxSize.Value)
fmt.Println("concurrency:", concurrency.Value)
```

# String choices

OneOfString only accepts a string value that matches one of the choices
configured for the flag.

```golang
// Define a flag that accepts either "one" or "two"
system := NewOneOfString("production", "testing")

flagset.Var(system, "test-flag", "usage here")

flag.Parse()

flagset.Parse([]string{"-test-flag", "production"})
fmt.Println("set value:", system.Value)
```
