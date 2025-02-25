package convert

import (
	"fmt"
	"strconv"
)

func ToInt(in any) int {
	switch v := in.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		out, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("invalid Atoi conversion on '%s'\n", v)
		}
		return out
	default:
		panic(fmt.Sprintf("invalid type %T\n", v))
	}
}

func ToFloat(in any) float64 {
	switch v := in.(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		out, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("invalid ParseFloat conversion on '%s'\n", v)
		}
		return out
	default:
		panic(fmt.Sprintf("invalid type %T\n", v))
	}
}

func Ptr[T any](in T) *T {
	return &in
}

func PtrD[T comparable](in *T) T {
	if in == nil {
		var out T
		return out
	}
	return *in
}
