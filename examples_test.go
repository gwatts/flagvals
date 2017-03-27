package flagvals

import (
	"flag"
	"fmt"
)

func ExampleRangeInt() {
	// Define a flag that accepts integers between 50 and 150
	// with a default value of 100.
	flagval := BetweenInt(100, 50, 150)

	flagset := flag.NewFlagSet("test-flags", flag.ContinueOnError)
	flagset.Var(flagval, "test-flag", "usage here")

	fmt.Println("default:", flagval.Value)
	flagset.Parse([]string{"-test-flag", "123"})
	fmt.Println("set value:", flagval.Value)

	// Output:
	// default: 100
	// set value: 123
}

func ExampleOneOfString() {
	// Define a flag that accepts either "one" or "two"
	flagval := NewOneOfString("one", "two")

	flagset := flag.NewFlagSet("test-flags", flag.ContinueOnError)
	flagset.Var(flagval, "test-flag", "usage here")

	flagset.Parse([]string{"-test-flag", "one"})
	fmt.Println("set value:", flagval.Value)

	// Output:
	// set value: one
}
