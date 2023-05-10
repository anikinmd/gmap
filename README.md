[![codecov](https://codecov.io/gh/anikinmd/gmap/branch/main/graph/badge.svg?token=J5Eotq61Tz)](https://codecov.io/gh/anikinmd/gmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/anikinmd/gmap)](https://goreportcard.com/report/github.com/anikinmd/gmap)
# gmap

Go concurrent generic map

gmap is simple generic based thread safe comparable to any map implementation.

Use map with any type without type casting.

Usage:
```
// key can be any comparable type
type keyStruct struct {
	str string
	i   int
}

// value can be any type
type valStruct struct {
	b []byte
}

func main() {
	m := gmap.NewGMap[keyStruct, valStruct]()
	m.Set(keyStruct{"a", 1}, valStruct{[]byte{0x1}})
	// Don't need type casting
	val, err := m.Get(keyStruct{"a", 1})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
```
