package edsger

import (
	"math"
	"reflect"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Number interface {
	~float32 | ~float64 | Integer
}

func Signed[T Number]() bool {
	var v T
	switch reflect.TypeOf(v).Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return false
	default:
		panic("Unknown type")
	}
}

// Returns true if the type is signed
func SignedInt[T Integer]() bool {
	var zero T
	return zero-1 < 0
}

func MaxInt[T Integer]() T {
	var v T
	bits := reflect.TypeOf(v).Size() * 8
	if SignedInt[T]() {
		return 1<<(bits-1) - 1
	} else {
		return (1 << bits) - 1
	}
}

func MaxValue[N Number]() N {
	var v N
	switch any(v).(type) {
	case float32, float64:
		return N(math.Inf(1))
	case int:
		return N(MaxInt[int]())
	case int8:
		return N(MaxInt[int8]())
	case int16:
		return N(MaxInt[int16]())
	case int32:
		return N(MaxInt[int32]())
	case int64:
		return N(MaxInt[int64]())
	case uint:
		return N(MaxInt[uint]())
	case uint8:
		return N(MaxInt[uint8]())
	case uint16:
		return N(MaxInt[uint16]())
	case uint32:
		return N(MaxInt[uint32]())
	case uint64:
		return N(MaxInt[uint64]())
	}
	panic("Unknown type!")
}
