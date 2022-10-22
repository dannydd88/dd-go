package dd

import "time"

// Ptr
// Generic type version returns a pointer to the value passed in
func Ptr[T any](v T) *T {
	return &v
}

// Val
// Generic type version returns the value of the pointer passed in or
// default value of this type if pointer is nil
func Val[T any](v *T) T {
	if v != nil {
		return *v
	}
	var ret T
	return ret
}

// ValD
// Generic type version returns the value of the pointer passed in or
// return |defaultVal| if pointer is nil
func ValD[T any](v *T, defaultVal T) T {
	if v != nil {
		return *v
	}
	return defaultVal
}

// PtrSlice
// Generic type version converts a slice of T values into a slice of
// T pointers
func PtrSlice[T any](src []T) []*T {
	dst := make([]*T, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// ValSlice
// Generic type version converts a slice of T pointers into a slice of
// T values
func ValSlice[T any](src []*T) []T {
	dst := make([]T, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// PtrMap
// Generic type version converts a string map of T values into a string
// map of T pointers
func PtrMap[T any](src map[string]T) map[string]*T {
	dst := make(map[string]*T)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// ValMap
// Generic type version converts a string map of T pointers into a string
// map of T values
func ValMap[T any](src map[string]*T) map[string]T {
	dst := make(map[string]T)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// SecondsTimeVal converts an int64 pointer to a time.Time value
// representing seconds since Epoch or time.Time{} if the pointer is nil.
func SecondsTimeVal(v *int64) time.Time {
	if v != nil {
		return time.Unix((*v / 1000), 0)
	}
	return time.Time{}
}

// MillisecondsTimeVal converts an int64 pointer to a time.Time value
// representing milliseconds sinch Epoch or time.Time{} if the pointer is nil.
func MillisecondsTimeVal(v *int64) time.Time {
	if v != nil {
		return time.Unix(0, (*v * 1000000))
	}
	return time.Time{}
}

// TimeUnixMilli returns a Unix timestamp in milliseconds from "January 1, 1970 UTC".
// The result is undefined if the Unix time cannot be represented by an int64.
// Which includes calling TimeUnixMilli on a zero Time is undefined.
//
// This utility is useful for service API's such as CloudWatch Logs which require
// their unix time values to be in milliseconds.
//
// See Go stdlib https://golang.org/pkg/time/#Time.UnixNano for more information.
func TimeUnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond/time.Nanosecond)
}
