package enum

import (
	"strconv"
)

type Enum[T int | int64] struct {
	key   T
	value string
}

func EnumFromKey[T int | int64](key T) Enum[T] {
	return Enum[T]{key: key}
}

func (e Enum[T]) Key() T {
	return e.key
}

func (e Enum[T]) Value() string {
	return e.value
}

func (e Enum[T]) String() string {
	return e.value
}

func (e *Enum[T]) SetKey(key T) {
	e.key = key
}

func (e *Enum[T]) SetValue(value string) {
	e.value = value
}

func (e Enum[T]) Values() []Enum[T] {
	return nil
}

type IEnum[T int | int64] interface {
	Values() []Enum[T]
	Key() T
	Value() string
}

// UnmarshalJSON implement json Unmarshal interface
func UnmarshalJSON[T int | int64](b []byte) (Enum[T], error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	str := string(b)
	var x Enum[T]
	values := x.Values()
	var found bool
	if i, err := strconv.ParseInt(str, 10, 64); err != nil {
		for _, val := range values {
			if val.Value() == str {
				x = val
				found = true
				break
			}
		}
	} else {
		key := T(i)
		for _, val := range values {
			if val.Key() == key {
				x = val
				found = true
				break
			}
		}
	}
	if !found {
		(&x).SetKey(-1)
		(&x).SetValue(str)
	}
	return x, nil
}
