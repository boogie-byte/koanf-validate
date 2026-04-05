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

package validate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/boogie-byte/koanf-validate"
)

func TestPredicates(t *testing.T) {
	testCases := []struct {
		name      string
		predicate validate.Predicate
		value     any
		wantErr   bool
	}{
		// Required
		{
			name:      "Required_Present",
			predicate: validate.Required,
			value:     "foo",
			wantErr:   false,
		},
		{
			name:      "Required_Missing",
			predicate: validate.Required,
			value:     nil,
			wantErr:   true,
		},

		// MinLen
		{
			name:      "MinLen_Pass",
			predicate: validate.MinLen(3),
			value:     "foobar",
			wantErr:   false,
		},
		{
			name:      "MinLen_Exact",
			predicate: validate.MinLen(3),
			value:     "foo",
			wantErr:   false,
		},
		{
			name:      "MinLen_Too_Short",
			predicate: validate.MinLen(5),
			value:     "foo",
			wantErr:   true,
		},

		// OneOf
		{
			name:      "OneOf_Match",
			predicate: validate.OneOf("foo", "bar", "baz"),
			value:     "bar",
			wantErr:   false,
		},
		{
			name:      "OneOf_No_Match",
			predicate: validate.OneOf("foo", "bar", "baz"),
			value:     "qux",
			wantErr:   true,
		},
		{
			name:      "OneOf_Type_Mismatch",
			predicate: validate.OneOf("foo", "bar"),
			value:     42,
			wantErr:   true,
		},
		{
			name:      "OneOf_Nil",
			predicate: validate.OneOf("foo", "bar"),
			value:     nil,
			wantErr:   true,
		},

		// MaxLen
		{
			name:      "MaxLen_Pass",
			predicate: validate.MaxLen(5),
			value:     "foo",
			wantErr:   false,
		},
		{
			name:      "MaxLen_Exact",
			predicate: validate.MaxLen(3),
			value:     "foo",
			wantErr:   false,
		},
		{
			name:      "MaxLen_Too_Long",
			predicate: validate.MaxLen(2),
			value:     "foo",
			wantErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.predicate(tc.value)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestOneOf_Empty_Panics(t *testing.T) {
	require.Panics(t, func() {
		validate.OneOf[string]()
	})
}

func TestMinLen_Negative_Panics(t *testing.T) {
	require.Panics(t, func() {
		validate.MinLen(-1)
	})
}

func TestMaxLen_Negative_Panics(t *testing.T) {
	require.Panics(t, func() {
		validate.MaxLen(-1)
	})
}
