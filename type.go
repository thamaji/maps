package maps

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type float interface {
	~float32 | ~float64
}

type complex interface {
	~complex64 | ~complex128
}

type ordered interface {
	integer | float | string
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
