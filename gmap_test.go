package gmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type keyVal struct {
	key   string
	value string
}

func fillMap(m *GMap[string, string], kv []keyVal) {
	for _, k := range kv {
		m.data[k.key] = k.value
	}
}

func Test_GMap_DeleteKey(t *testing.T) {
	type testCase struct {
		name    string
		filling []keyVal
		key     string
		err     error
	}
	tests := []testCase{
		{
			name: "Delete existed key",
			filling: []keyVal{
				{"a", "a"},
			},
			key: "a",
			err: nil,
		},
		{
			name: "Delete nonexistent key",
			filling: []keyVal{
				{"a", "a"},
			},
			key: "b",
			err: ErrorUnknownKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewGMap[string, string]()
			fillMap(m, tt.filling)
			err := m.Delete(tt.key)
			if tt.err == nil {
				assert.NoError(t, err)
				_, ok := m.data[tt.key]
				assert.False(t, ok)
			} else {
				assert.ErrorIs(t, ErrorUnknownKey, err)
			}
		})
	}
}

func Test_GMap_Get(t *testing.T) {
	type testCase struct {
		name          string
		filling       []keyVal
		key           string
		want          string
		assertionFunc assert.ErrorAssertionFunc
	}
	tests := []testCase{
		{
			name: "get existed key",
			filling: []keyVal{
				{"a", "a"},
			},
			key:           "a",
			want:          "a",
			assertionFunc: assert.NoError,
		},
		{
			name: "Get nonexistent key",
			filling: []keyVal{
				{"a", "a"},
			},
			key:           "b",
			assertionFunc: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewGMap[string, string]()
			fillMap(c, tt.filling)
			val, err := c.Get(tt.key)
			tt.assertionFunc(t, err)
			assert.Equal(t, tt.want, val)
		})
	}
}

func Test_GMap_Set(t *testing.T) {
	type testCase struct {
		name    string
		filling []keyVal
		set     keyVal
	}
	tests := []testCase{
		{
			name:    "Set not existed key",
			filling: []keyVal{{"a", "a"}},
			set:     keyVal{"b", "b"},
		},
		{
			name:    "Set existed key",
			filling: []keyVal{{"a", "a"}},
			set:     keyVal{"a", "b"},
		},
	}

	for _, tt := range tests {
		m := NewGMap[string, string]()
		fillMap(m, tt.filling)
		t.Run(tt.name, func(t *testing.T) {
			m.Set(tt.set.key, tt.set.value)
			assert.Equal(t, tt.set.value, m.data[tt.set.key])
		})
	}
}

func Test_GMap_GetKeys(t *testing.T) {
	type testCase struct {
		name    string
		filling []keyVal
		want    []string
	}
	tests := []testCase{
		{
			name: "Empty cache",
			want: []string{},
		},
		{
			name:    "Not empty cache",
			filling: []keyVal{{"a", "a"}, {"b", "b"}},
			want:    []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewGMap[string, string]()
			fillMap(m, tt.filling)
			assert.ElementsMatch(t, tt.want, m.GetKeys())
		})
	}
}

func Test_GMap_CheckKeyExists(t *testing.T) {
	type testCase struct {
		name    string
		filling []keyVal
		key     string
		want    bool
	}
	tests := []testCase{
		{
			name:    "check existed key",
			filling: []keyVal{{"a", "a"}},
			key:     "a",
			want:    true,
		},
		{
			name:    "check nonexistent key",
			filling: []keyVal{{"a", "a"}},
			key:     "b",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewGMap[string, string]()
			fillMap(m, tt.filling)
			val := m.CheckKeyExists(tt.key)
			assert.Equal(t, tt.want, val)
		})
	}
}
