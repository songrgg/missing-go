package comparison

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolean(t *testing.T) {
	// boolean to int
	result, err := Equal(true, 1)
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = Equal(false, 0)
	assert.True(t, result)

	// boolean to string
	result, err = Equal(true, "1")
	assert.True(t, result)

	result, err = Equal(false, "0")
	assert.True(t, result)

	result, err = Equal(true, "true")
	assert.True(t, result)

	result, err = Equal(false, "false")
	assert.True(t, result)

	result, _ = GreaterThan(true, "1")
	assert.False(t, result)

	result, _ = LessThan(true, "1")
	assert.False(t, result)
}

func TestIntToString(t *testing.T) {
	// int64
	result, err := Equal(int64(1000000000000), "1000000000000")
	assert.True(t, result)
	assert.Nil(t, err)

	result, _ = GreaterThan(int64(1000000000000), "11111")
	assert.True(t, result)

	result, _ = LessThan(int64(1000), "11111")
	assert.True(t, result)

	result, _ = LessThan(int64(1000000000000), "11111")
	assert.False(t, result)

	result, _ = LessEqual(int64(100000), "100000")
	assert.True(t, result)

	result, _ = LessEqual(int64(1000000), "100000")
	assert.False(t, result)

	result, _ = GreaterEqual(int64(1000000), "100000")
	assert.True(t, result)

	result, _ = LessEqual(int64(100000), "100000")
	assert.True(t, result)

	result, err = Equal(int64(1000000000000), "2")
	assert.False(t, result)

	// when string is not integer
	result, err = Equal(int64(1000000000000), "sss")
	assert.False(t, result)

	// int8
	result, err = Equal(int8(10), "10")
	assert.True(t, result)

	result, err = Equal(int8(10), "2")
	assert.False(t, result)

	// int16
	result, err = Equal(int16(10), "10")
	assert.True(t, result)

	result, err = Equal(int16(10), "2")
	assert.False(t, result)

	// int32
	result, err = Equal(int32(10), "10")
	assert.True(t, result)

	result, err = Equal(int32(10), "2")
	assert.False(t, result)

	// int
	result, err = Equal(int(10), "10")
	assert.True(t, result)

	result, err = Equal(int(10), "2")
	assert.False(t, result)

	// int to string with points
	result, err = Equal(int(10), "10.0")
	assert.True(t, result)
}

func TestIntToFloat(t *testing.T) {
	// int64
	result, err := Equal(int64(1000000000000), float64(1000000000000))
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = Equal(int64(1000000000000), float64(2))
	assert.False(t, result)

	result, err = Equal(int64(2), float32(2))
	assert.True(t, result)

	// int8
	result, err = Equal(int8(10), float64(10))
	assert.True(t, result)
	result, err = Equal(int8(10), float64(2))
	assert.False(t, result)

	// int16
	result, err = Equal(int16(10), float64(10))
	assert.True(t, result)
	result, err = Equal(int16(10), float64(2))
	assert.False(t, result)

	// int32
	result, err = Equal(int32(10), float64(10))
	assert.True(t, result)
	result, err = Equal(int32(10), float64(2))
	assert.False(t, result)

	// int
	result, err = Equal(int(10), float64(10.0))
	assert.True(t, result)
	result, err = Equal(int(10), float64(2.00))
	assert.False(t, result)
}

func TestFloatToString(t *testing.T) {
	// float64 to string
	result, err := Equal(float64(1.111), "1.111")
	assert.True(t, result)
	assert.Nil(t, err)

	result, err = Equal(float64(1.000), "1.00")
	assert.True(t, result)

	// float32 to string
	result, _ = Equal(float32(1.111), "1.111")
	assert.True(t, result)

	result, err = Equal(float32(1.000), "1.00")
	assert.True(t, result)
}
