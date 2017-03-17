package flagvals

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	_ flag.Value = new(RangeInt)
)

// RangeInt holds an int64 value between a defined min/max range.
// It implements the flag.Value interface.
type RangeInt struct {
	Value int64 // The value as passed to Set, else the initial value.
	IsSet bool  // true if Set has been called successfully.

	min *int64
	max *int64
}

// Set implements the flag.Value interface.
func (ri *RangeInt) Set(v string) error {
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}
	if ri.min != nil && val < *ri.min {
		return fmt.Errorf("value must be >= %d", *ri.min)
	}
	if ri.max != nil && val > *ri.max {
		return fmt.Errorf("value must be <= %d", *ri.max)
	}
	ri.Value = val
	ri.IsSet = true
	return nil
}

// String returns the current value of the flag.
func (ri *RangeInt) String() string {
	return strconv.FormatInt(ri.Value, 10)
}

// Min returns the minimum value required for the RangeInt
// It returns isDefined=false if no minimum required value has
// been defined for the range.
func (ri *RangeInt) Min() (min int64, isDefined bool) {
	if ri.min != nil {
		return *ri.min, true
	}
	return 0, false
}

// Min returns the maximum value required for the RangeInt.
// It returns isDefined=false if no maximum required value has been defined
// for the range.
func (ri *RangeInt) Max() (max int64, isDefined bool) {
	if ri.max != nil {
		return *ri.max, true
	}
	return 0, false
}

// LTEInt defines a RangeInt that accepts values less than, or  equal to the
// supplied value.
func LTEInt(initialValue, max int64) *RangeInt {
	return &RangeInt{max: &max, Value: initialValue}
}

// GETInt defines a RangeInt that accepts values greater than, or equal to
// the supplied value.
func GTEInt(initialValue, min int64) *RangeInt {
	return &RangeInt{min: &min, Value: initialValue}
}

// BetweenInt defines a  RangeInt that accepts values between the supplied
// min and max values.  The values are inclusive, so BetweenInt(1,3)
// would accept values of 1, 2 or 3.
func BetweenInt(initialValue, min, max int64) *RangeInt {
	return &RangeInt{min: &min, max: &max, Value: initialValue}
}
