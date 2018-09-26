package comparison

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

type compareFunc func(interface{}, interface{}) bool

var priority map[reflect.Kind]int

func init() {
	priority = map[reflect.Kind]int{
		reflect.String:  -1,
		reflect.Bool:    0,
		reflect.Int8:    1,
		reflect.Int16:   2,
		reflect.Int32:   3,
		reflect.Int:     4,
		reflect.Float32: 5,
		reflect.Int64:   6,
		reflect.Float64: 7,
	}
}

func compare(v1 interface{}, v2 interface{}, compareFunc compareFunc) (bool, error) {
	k1 := reflect.TypeOf(v1).Kind()
	k2 := reflect.TypeOf(v2).Kind()
	if k1 == k2 {
		return v1 == v2, nil
	}

	var (
		p1 int
		p2 int
		ok bool
	)
	if p1, ok = priority[k1]; !ok {
		return false, fmt.Errorf("type %v not supported", k1)
	}

	if p2, ok = priority[k2]; !ok {
		return false, fmt.Errorf("type %v not supported", k2)
	}

	if p1 < p2 {
		// transform v1
		var err error
		v1, err = transform(v1, k2)
		if err != nil {
			return false, fmt.Errorf("transform error: %v", err)
		}
	} else {
		// transform v2
		var err error
		v2, err = transform(v2, k1)
		if err != nil {
			return false, fmt.Errorf("transform error: %v", err)
		}
	}

	// maybe transformation upgrade the v1's kind, so we should double check.
	if reflect.TypeOf(v1).Kind() != reflect.TypeOf(v2).Kind() {
		return compare(v1, v2, compareFunc)
	}

	return compareFunc(v1, v2), nil
}

// Equal returns if the given two objects equal to each other, it's not a deep equality,
// some implicit type conversion will be made according to the type priority.
func Equal(v1 interface{}, v2 interface{}) (bool, error) {
	return compare(v1, v2, func(v1 interface{}, v2 interface{}) bool {
		return v1 == v2
	})
}

// GreaterThan returns if the given v1 is greater than v2,
// some implicit type conversion will be made according to the type priority.
func GreaterThan(v1 interface{}, v2 interface{}) (bool, error) {
	return compare(v1, v2, func(v1 interface{}, v2 interface{}) bool {
		k := reflect.TypeOf(v1).Kind()
		switch k {
		case reflect.Float32:
			return v1.(float32) > v2.(float32)
		case reflect.Float64:
			return v1.(float64) > v2.(float64)
		case reflect.String:
			return v1.(string) > v2.(string)
		case reflect.Int:
			return v1.(int) > v2.(int)
		case reflect.Int8:
			return v1.(int8) > v2.(int8)
		case reflect.Int16:
			return v1.(int16) > v2.(int16)
		case reflect.Int32:
			return v1.(int32) > v2.(int32)
		case reflect.Int64:
			return v1.(int64) > v2.(int64)
		}
		return false
	})
}

// LessThan returns if the given v1 is less than v2,
// some implicit type conversion will be made according to the type priority.
func LessThan(v1 interface{}, v2 interface{}) (bool, error) {
	e, err := Equal(v1, v2)
	if err != nil {
		return false, err
	}

	if e {
		return false, nil
	}

	g, err := GreaterThan(v1, v2)
	if err != nil {
		return false, err
	}
	return !g, nil
}

// LessEqual returns if the given v1 is less equal to v2,
// some implicit type conversion will be made according to the type priority.
func LessEqual(v1 interface{}, v2 interface{}) (bool, error) {
	g, err := GreaterThan(v1, v2)
	if err != nil {
		return false, err
	}
	return !g, nil
}

// GreaterEqual returns if the given v1 is greater equal to v2,
// some implicit type conversion will be made according to the type priority.
func GreaterEqual(v1 interface{}, v2 interface{}) (bool, error) {
	e, err := Equal(v1, v2)
	if err != nil {
		return false, err
	}

	if e {
		return true, nil
	}

	g, err := GreaterThan(v1, v2)
	if err != nil {
		return false, err
	}
	return g, nil
}

func transform(v1 interface{}, targetKind reflect.Kind) (interface{}, error) {
	// some special processes when source kind is string
	sourceKind := reflect.TypeOf(v1).Kind()
	switch sourceKind {
	case reflect.String:
		// upgrade the target kind to float64 when source string contains dot
		if strings.Contains(v1.(string), ".") &&
			(targetKind != reflect.Float64 && targetKind != reflect.Float32) {
			targetKind = reflect.Float64
		}
	}

	var v2 interface{}
	switch targetKind {
	case reflect.Int8:
		v2 = cast.ToInt8(v1)
	case reflect.Int16:
		v2 = cast.ToInt16(v1)
	case reflect.Int:
		v2 = cast.ToInt(v1)
	case reflect.Int32:
		v2 = cast.ToInt32(v1)
	case reflect.Int64:
		v2 = cast.ToInt64(v1)
	case reflect.Float32:
		v2 = cast.ToFloat32(v1)
	case reflect.Float64:
		v2 = cast.ToFloat64(v1)
	case reflect.Bool:
		v2 = cast.ToBool(v1)
	case reflect.String:
		v2 = cast.ToString(v1)
	default:
		return nil, fmt.Errorf("transform err: type %v not found", targetKind)

	}
	return v2, nil
}
