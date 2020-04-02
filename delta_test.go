package fdelta

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelta(t *testing.T) {
	readFile := func(path string) []byte {
		data, err := ioutil.ReadFile(
			filepath.FromSlash(path),
		)
		assert.Nil(t, err)
		return data
	}

	tests := []struct {
		origin        []byte
		target        []byte
		expectedDelta []byte
	}{
		{
			origin:        readFile("data/1/origin"),
			target:        readFile("data/1/target"),
			expectedDelta: readFile("data/1/delta"),
		},
		{
			origin:        readFile("data/2/origin"),
			target:        readFile("data/2/target"),
			expectedDelta: readFile("data/2/delta"),
		},
		{
			origin:        readFile("data/3/origin"),
			target:        readFile("data/3/target"),
			expectedDelta: readFile("data/3/delta"),
		},
		{
			origin:        readFile("data/4/origin"),
			target:        readFile("data/4/target"),
			expectedDelta: readFile("data/4/delta"),
		},
		{
			origin:        readFile("data/5/origin"),
			target:        readFile("data/5/target"),
			expectedDelta: readFile("data/5/delta"),
		},
	}

	for _, test := range tests {
		delta := Create(test.origin, test.target)
		assert.Equal(t, test.expectedDelta, delta)
	}
}

func BenchmarkCreateDelta(b *testing.B) {
	readFile := func(path string) []byte {
		data, err := ioutil.ReadFile(
			filepath.FromSlash(path),
		)
		assert.Nil(b, err)
		return data
	}

	tests := []struct {
		origin []byte
		target []byte
	}{
		{
			origin: readFile("data/1/origin"),
			target: readFile("data/1/target"),
		},
		{
			origin: readFile("data/2/origin"),
			target: readFile("data/2/target"),
		},
		{
			origin: readFile("data/3/origin"),
			target: readFile("data/3/target"),
		},
		{
			origin: readFile("data/4/origin"),
			target: readFile("data/4/target"),
		},
		{
			origin: readFile("data/5/origin"),
			target: readFile("data/5/target"),
		},
	}

	for i, test := range tests {
		b.Run(fmt.Sprintf("CreateDelta%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Create(test.origin, test.target)
			}
		})
	}
}

func BenchmarkApplyDelta(b *testing.B) {
	readFile := func(path string) []byte {
		data, err := ioutil.ReadFile(
			filepath.FromSlash(path),
		)
		assert.Nil(b, err)
		return data
	}

	tests := []struct {
		origin []byte
		delta  []byte
	}{
		{
			origin: readFile("data/1/origin"),
			delta:  readFile("data/1/delta"),
		},
		{
			origin: readFile("data/2/origin"),
			delta:  readFile("data/2/delta"),
		},
		{
			origin: readFile("data/3/origin"),
			delta:  readFile("data/3/delta"),
		},
		{
			origin: readFile("data/4/origin"),
			delta:  readFile("data/4/delta"),
		},
		{
			origin: readFile("data/5/origin"),
			delta:  readFile("data/5/delta"),
		},
	}

	for i, test := range tests {
		b.Run(fmt.Sprintf("ApplyDelta%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Apply(test.origin, test.delta) //nolint: errcheck
			}
		})
	}
}
