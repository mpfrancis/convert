package convert

import (
	"testing"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		input    any
		expected int
	}{
		{input: 42, expected: 42},
		{input: int8(42), expected: 42},
		{input: int16(42), expected: 42},
		{input: int32(42), expected: 42},
		{input: int64(42), expected: 42},
		{input: uint(42), expected: 42},
		{input: uint8(42), expected: 42},
		{input: uint16(42), expected: 42},
		{input: uint32(42), expected: 42},
		{input: uint64(42), expected: 42},
		{input: float32(42.0), expected: 42},
		{input: float64(42.0), expected: 42},
		{input: true, expected: 1},
		{input: false, expected: 0},
		{input: "42", expected: 42},
		{input: "invalid", expected: 0},
	}

	for _, test := range tests {
		result := ToInt(test.input)
		if result != test.expected {
			t.Errorf("ToInt(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestToFloat(t *testing.T) {
	tests := []struct {
		input    any
		expected float64
	}{
		{input: 42, expected: 42.0},
		{input: int8(42), expected: 42.0},
		{input: int16(42), expected: 42.0},
		{input: int32(42), expected: 42.0},
		{input: int64(42), expected: 42.0},
		{input: uint(42), expected: 42.0},
		{input: uint8(42), expected: 42.0},
		{input: uint16(42), expected: 42.0},
		{input: uint32(42), expected: 42.0},
		{input: uint64(42), expected: 42.0},
		{input: float32(42.0), expected: 42.0},
		{input: float64(42.0), expected: 42.0},
		{input: true, expected: 1.0},
		{input: false, expected: 0.0},
		{input: "42.0", expected: 42.0},
		{input: "invalid", expected: 0.0},
	}

	for _, test := range tests {
		result := ToFloat(test.input)
		if result != test.expected {
			t.Errorf("ToFloat(%v) = %f; want %f", test.input, result, test.expected)
		}
	}
}

func TestPtr(t *testing.T) {
	value := 42
	ptr := Ptr(value)
	if *ptr != value {
		t.Errorf("Ptr(%d) = %d; want %d", value, *ptr, value)
	}
}

func TestPtrD(t *testing.T) {
	value := 42
	ptr := &value
	result := PtrD(ptr)
	if result != value {
		t.Errorf("PtrD(%d) = %d; want %d", *ptr, result, value)
	}

	var nilPtr *int
	result = PtrD(nilPtr)
	if result != 0 {
		t.Errorf("PtrD(nil) = %d; want %d", result, 0)
	}
}
