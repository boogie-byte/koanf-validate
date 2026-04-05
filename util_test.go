// Copyright 2026 Sergey Vinogradov
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validate

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcreteValueOf(t *testing.T) {
	ptr := func(v any) *any { return &v }

	testCases := []struct {
		name     string
		value    any
		expected reflect.Kind
	}{
		{
			name:     "String",
			value:    "foo",
			expected: reflect.String,
		},
		{
			name:     "Int",
			value:    42,
			expected: reflect.Int,
		},
		{
			name:     "Slice",
			value:    []any{"a"},
			expected: reflect.Slice,
		},
		{
			name:     "Map",
			value:    map[string]any{"a": 1},
			expected: reflect.Map,
		},
		{
			name:     "Pointer_To_String",
			value:    ptr("foo"),
			expected: reflect.String,
		},
		{
			name:     "Pointer_To_Slice",
			value:    ptr([]any{"a"}),
			expected: reflect.Slice,
		},
		{
			name:     "Pointer_To_Map",
			value:    ptr(map[string]any{"a": 1}),
			expected: reflect.Map,
		},
		{
			name:     "Double_Pointer",
			value:    ptr(ptr("foo")),
			expected: reflect.String,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rv := concreteValueOf(tc.value)
			require.Equal(t, tc.expected, rv.Kind())
		})
	}
}

func TestLenOf(t *testing.T) {
	testCases := []struct {
		name        string
		value       any
		expectedLen int
		wantErr     bool
	}{
		{
			name:        "Nil",
			value:       nil,
			expectedLen: 0,
		},
		{
			name:        "String",
			value:       "foo",
			expectedLen: 3,
		},
		{
			name:        "Empty_String",
			value:       "",
			expectedLen: 0,
		},
		{
			name:        "Slice",
			value:       []any{"a", "b", "c"},
			expectedLen: 3,
		},
		{
			name:        "Map",
			value:       map[string]any{"a": 1, "b": 2},
			expectedLen: 2,
		},
		{
			name:    "Unsupported_Type",
			value:   42,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := lenOf(tc.value)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedLen, l)
			}
		})
	}
}
