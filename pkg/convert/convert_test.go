package convert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapToValueSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: []int{1},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "duplicates",
			input:    map[string]int{"a": 1, "b": 2, "c": 1},
			expected: []int{1, 2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapToValueSlice(test.input)

			require.Equal(t, len(test.input), len(test.expected))

			// Since map iteration order is nondeterministic, we need to check if all expected values are preset with multiplicity.
			resultMap := make(map[int]int)
			for _, v := range result {
				resultMap[v] += 1
			}

			for _, expected := range test.expected {
				resultMap[expected] -= 1
			}

			for _, v := range resultMap {
				require.Equal(t, 0, v, "expected multiplicity for %d does not match", v)
			}
		})
	}
}

func TestMapToKeySlice(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    map[string]int{"a": 1},
			expected: []string{"a"},
		},
		{
			name:     "multiple elements",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapToKeySlice(test.input)

			require.Equal(t, len(test.input), len(test.expected))

			// Since map iteration order is nondeterministic, we need to check if all expected keys are present.
			resultMap := make(map[string]bool)
			for _, v := range result {
				resultMap[v] = true
			}

			for _, expected := range test.expected {
				require.True(t, resultMap[expected])
			}
		})
	}
}

func TestSliceToSet(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected map[int]bool
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: map[int]bool{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: map[int]bool{1: true},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3},
			expected: map[int]bool{1: true, 2: true, 3: true},
		},
		{
			name:     "duplicates",
			input:    []int{1, 2, 1, 3, 2},
			expected: map[int]bool{1: true, 2: true, 3: true},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SliceToSet(test.input)

			require.Equal(t, len(test.expected), len(result))

			for key, value := range test.expected {
				require.Equal(t, value, result[key])
			}
		})
	}
}
