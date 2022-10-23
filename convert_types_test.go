package dd

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPtrAndVal(t *testing.T) {
	assert := assert.New(t)

	stringVal := "a"
	stringPtr := Ptr(stringVal)
	assert.Equal(stringVal, *stringPtr)
	assert.Equal(stringVal, Val(stringPtr))

	intVal := 1
	intPtr := Ptr(intVal)
	assert.Equal(intVal, *intPtr)
	assert.Equal(intVal, Val(intPtr))
}

func TestNilPtrVal(t *testing.T) {
	assert := assert.New(t)

	var nilStringPtr *string
	assert.Equal("", Val(nilStringPtr))
	stringDefaultVal := "a"
	assert.Equal(stringDefaultVal, ValD(nilStringPtr, stringDefaultVal))

	var nilIntPtr *int
	assert.Equal(0, Val(nilIntPtr))
	intDefaultVal := 1
	assert.Equal(intDefaultVal, ValD(nilIntPtr, intDefaultVal))
}

var testCasesStringSlice = [][]string{
	{"a", "b", "c", "d", "e"},
	{"a", "b", "", "", "e"},
}

func TestStringSlice(t *testing.T) {
	for idx, in := range testCasesStringSlice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesStringValueSlice = [][]*string{
	{Ptr("a"), Ptr("b"), nil, Ptr("c")},
}

func TestStringValueSlice(t *testing.T) {
	for idx, in := range testCasesStringValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != "" {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != "" {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *in[i], *out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesStringMap = []map[string]string{
	{"a": "1", "b": "2", "c": "3"},
}

func TestStringMap(t *testing.T) {
	for idx, in := range testCasesStringMap {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesBoolSlice = [][]bool{
	{true, true, false, false},
}

func TestBoolSlice(t *testing.T) {
	for idx, in := range testCasesBoolSlice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesBoolValueSlice = [][]*bool{
	{Ptr(true), Ptr(true), Ptr(false), Ptr(false)},
}

func TestBoolValueSlice(t *testing.T) {
	for idx, in := range testCasesBoolValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesBoolMap = []map[string]bool{
	{"a": true, "b": false, "c": true},
}

func TestBoolMap(t *testing.T) {
	for idx, in := range testCasesBoolMap {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUintSlice = [][]uint{
	{1, 2, 3, 4},
}

func TestUintSlice(t *testing.T) {
	for idx, in := range testCasesUintSlice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUintValueSlice = [][]*uint{
	{Ptr(uint(1)), Ptr(uint(2)), Ptr(uint(3)), Ptr(uint(4))},
}

func TestUintValueSlice(t *testing.T) {
	for idx, in := range testCasesUintValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUintMap = []map[string]uint{
	{"a": 3, "b": 2, "c": 1},
}

func TestUintMap(t *testing.T) {
	for idx, in := range testCasesUintMap {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesIntSlice = [][]int{
	{1, 2, 3, 4},
}

func TestIntSlice(t *testing.T) {
	for idx, in := range testCasesIntSlice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesIntValueSlice = [][]*int{
	{Ptr(1), Ptr(2), Ptr(3), Ptr(4)},
}

func TestIntValueSlice(t *testing.T) {
	for idx, in := range testCasesIntValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesIntMap = []map[string]int{
	{"a": 3, "b": 2, "c": 1},
}

func TestIntMap(t *testing.T) {
	for idx, in := range testCasesIntMap {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt8Slice = [][]int8{
	{1, 2, 3, 4},
}

func TestInt8Slice(t *testing.T) {
	for idx, in := range testCasesInt8Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt8ValueSlice = [][]*int8{
	{Ptr(int8(1)), Ptr(int8(2)), Ptr(int8(3)), Ptr(int8(4))},
}

func TestInt8ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt8ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt8Map = []map[string]int8{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt8Map(t *testing.T) {
	for idx, in := range testCasesInt8Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt16Slice = [][]int16{
	{1, 2, 3, 4},
}

func TestInt16Slice(t *testing.T) {
	for idx, in := range testCasesInt16Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt16ValueSlice = [][]*int16{
	{Ptr(int16(1)), Ptr(int16(2)), Ptr(int16(3)), Ptr(int16(4))},
}

func TestInt16ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt16ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt16Map = []map[string]int16{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt16Map(t *testing.T) {
	for idx, in := range testCasesInt16Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt32Slice = [][]int32{
	{1, 2, 3, 4},
}

func TestInt32Slice(t *testing.T) {
	for idx, in := range testCasesInt32Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt32ValueSlice = [][]*int32{
	{Ptr(int32(1)), Ptr(int32(2)), Ptr(int32(3)), Ptr(int32(4))},
}

func TestInt32ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt32ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt32Map = []map[string]int32{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt32Map(t *testing.T) {
	for idx, in := range testCasesInt32Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt64Slice = [][]int64{
	{1, 2, 3, 4},
}

func TestInt64Slice(t *testing.T) {
	for idx, in := range testCasesInt64Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt64ValueSlice = [][]*int64{
	{Ptr(int64(1)), Ptr(int64(2)), Ptr(int64(3)), Ptr(int64(4))},
}

func TestInt64ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt64ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt64Map = []map[string]int64{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt64Map(t *testing.T) {
	for idx, in := range testCasesInt64Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint8Slice = [][]uint8{
	{1, 2, 3, 4},
}

func TestUint8Slice(t *testing.T) {
	for idx, in := range testCasesUint8Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint8ValueSlice = [][]*uint8{
	{Ptr(uint8(1)), Ptr(uint8(2)), Ptr(uint8(3)), Ptr(uint8(4))},
}

func TestUint8ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint8ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint8Map = []map[string]uint8{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint8Map(t *testing.T) {
	for idx, in := range testCasesUint8Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint16Slice = [][]uint16{
	{1, 2, 3, 4},
}

func TestUint16Slice(t *testing.T) {
	for idx, in := range testCasesUint16Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint16ValueSlice = [][]*uint16{
	{Ptr(uint16(1)), Ptr(uint16(2)), Ptr(uint16(3)), Ptr(uint16(4))},
}

func TestUint16ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint16ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint16Map = []map[string]uint16{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint16Map(t *testing.T) {
	for idx, in := range testCasesUint16Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint32Slice = [][]uint32{
	{1, 2, 3, 4},
}

func TestUint32Slice(t *testing.T) {
	for idx, in := range testCasesUint32Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint32ValueSlice = [][]*uint32{
	{Ptr(uint32(1)), Ptr(uint32(2)), Ptr(uint32(3)), Ptr(uint32(4))},
}

func TestUint32ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint32ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint32Map = []map[string]uint32{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint32Map(t *testing.T) {
	for idx, in := range testCasesUint32Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint64Slice = [][]uint64{
	{1, 2, 3, 4},
}

func TestUint64Slice(t *testing.T) {
	for idx, in := range testCasesUint64Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint64ValueSlice = [][]*uint64{
	{Ptr(uint64(1)), Ptr(uint64(2)), Ptr(uint64(3)), Ptr(uint64(4))},
}

func TestUint64ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint64ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint64Map = []map[string]uint64{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint64Map(t *testing.T) {
	for idx, in := range testCasesUint64Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat32Slice = [][]float32{
	{1, 2, 3, 4},
}

func TestFloat32Slice(t *testing.T) {
	for idx, in := range testCasesFloat32Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat32ValueSlice = [][]*float32{
	{Ptr(float32(1)), Ptr(float32(2)), Ptr(float32(3)), Ptr(float32(4))},
}

func TestFloat32ValueSlice(t *testing.T) {
	for idx, in := range testCasesFloat32ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesFloat32Map = []map[string]float32{
	{"a": 3, "b": 2, "c": 1},
}

func TestFloat32Map(t *testing.T) {
	for idx, in := range testCasesFloat32Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat64Slice = [][]float64{
	{1, 2, 3, 4},
}

func TestFloat64Slice(t *testing.T) {
	for idx, in := range testCasesFloat64Slice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat64ValueSlice = [][]*float64{
	{Ptr(float64(1)), Ptr(float64(2)), Ptr(float64(3)), Ptr(float64(4))},
}

func TestFloat64ValueSlice(t *testing.T) {
	for idx, in := range testCasesFloat64ValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), *(out2[i]); e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesFloat64Map = []map[string]float64{
	{"a": 3, "b": 2, "c": 1},
}

func TestFloat64Map(t *testing.T) {
	for idx, in := range testCasesFloat64Map {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesTimeSlice = [][]time.Time{
	{time.Now(), time.Now().AddDate(100, 0, 0)},
}

func TestTimeSlice(t *testing.T) {
	for idx, in := range testCasesTimeSlice {
		if in == nil {
			continue
		}
		out := PtrSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesTimeValueSlice = [][]*time.Time{}

func TestTimeValueSlice(t *testing.T) {
	for idx, in := range testCasesTimeValueSlice {
		if in == nil {
			continue
		}
		out := ValSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if !out[i].IsZero() {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := PtrSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if !(out2[i]).IsZero() {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesTimeMap = []map[string]time.Time{
	{"a": time.Now().AddDate(-100, 0, 0), "b": time.Now()},
}

func TestTimeMap(t *testing.T) {
	for idx, in := range testCasesTimeMap {
		if in == nil {
			continue
		}
		out := PtrMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := ValMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

type TimeValueTestCase struct {
	in        int64
	outSecs   time.Time
	outMillis time.Time
}

var testCasesTimeValue = []TimeValueTestCase{
	{
		in:        int64(1501558289000),
		outSecs:   time.Unix(1501558289, 0),
		outMillis: time.Unix(1501558289, 0),
	},
	{
		in:        int64(1501558289001),
		outSecs:   time.Unix(1501558289, 0),
		outMillis: time.Unix(1501558289, 1*1000000),
	},
}

func TestSecondsTimeValue(t *testing.T) {
	for idx, testCase := range testCasesTimeValue {
		out := SecondsTimeVal(&testCase.in)
		if e, a := testCase.outSecs, out; e != a {
			t.Errorf("Unexpected value for time value at %d", idx)
		}
	}
}

func TestMillisecondsTimeValue(t *testing.T) {
	for idx, testCase := range testCasesTimeValue {
		out := MillisecondsTimeVal(&testCase.in)
		if e, a := testCase.outMillis, out; e != a {
			t.Errorf("Unexpected value for time value at %d", idx)
		}
	}
}
