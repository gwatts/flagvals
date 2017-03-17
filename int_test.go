package flagvals

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type rangeIntTest struct {
	val       int
	shouldErr bool
}

var gteTests = []rangeIntTest{
	{0, true},
	{1, true},
	{2, false},
	{3, false},
}

func TestGTEInt(t *testing.T) {
	for _, test := range gteTests {
		rv := GTEInt(-1, 2)
		testIntFlag(t, rv, "TestGTEInt", test, 2, 0)
	}
}

var lteTests = []rangeIntTest{
	{0, false},
	{1, false},
	{2, false},
	{3, true},
}

func TestLTEInt(t *testing.T) {
	for _, test := range lteTests {
		rv := LTEInt(-1, 2)
		testIntFlag(t, rv, "TestLTEInt", test, 0, 2)
	}
}

var betweenIntTests = []struct {
	val       int
	shouldErr bool
}{
	{0, true},
	{1, false},
	{2, false},
	{3, false},
	{4, true},
}

func TestBetweenInt(t *testing.T) {
	for _, test := range betweenIntTests {
		rv := BetweenInt(-1, 1, 3)
		testIntFlag(t, rv, "TestBetweenInt", test, 1, 3)
	}
}

func testIntFlag(t *testing.T, rv *RangeInt, testName string, test rangeIntTest, withMin, withMax int64) {
	assert := assert.New(t)
	assert.Equal("-1", rv.String(), "input=%d", test.val)
	min, minIsDef := rv.Min()
	max, maxIsdef := rv.Max()

	if withMin != 0 {
		assert.True(minIsDef, "%s input=%d", testName, test.val)
		assert.EqualValues(withMin, min, "%s input=%d", testName, test.val)
	} else {
		assert.False(minIsDef, "%s input=%d", testName, test.val)
		assert.EqualValues(0, min, "%s input=%d", testName, test.val)
	}

	if withMax != 0 {
		assert.True(maxIsdef, "%s input=%d", testName, test.val)
		assert.EqualValues(withMax, max, "%s input=%d", testName, test.val)
	} else {
		assert.False(maxIsdef, "%s input=%d", testName, test.val)
		assert.EqualValues(0, max, "%s input=%d", testName, test.val)
	}

	err := rv.Set(strconv.Itoa(test.val))
	if test.shouldErr {
		assert.Error(err, "%s input=%d", testName, test.val)
		assert.False(rv.IsSet, "%s input=%d", testName, test.val)
		assert.EqualValues(-1, rv.Value, "%s input=%d", testName, test.val)

	} else {
		assert.NoError(err, "%s input=%d", testName, test.val)
		assert.True(rv.IsSet, "%s input=%d", testName, test.val)
		assert.EqualValues(test.val, rv.Value, "%s input=%d", testName, test.val)
		assert.Equal(strconv.Itoa(test.val), rv.String(), "%s input=%d", testName, test.val)
	}
}

func TestInvalidInt(t *testing.T) {
	assert := assert.New(t)
	rv := LTEInt(-1, 10)
	err := rv.Set("5")
	assert.NoError(err)
	err = rv.Set("invalid")
	assert.Error(err)
}
