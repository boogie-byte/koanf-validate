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

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/require"

	"github.com/boogie-byte/koanf-validate"
)

func TestValidate(t *testing.T) {
	t.Run("All_Pass", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.bar": "val1",
			"foo.baz": "val2",
		})

		errs := validate.Validate(k,
			validate.Rule("foo.bar", validate.Required),
			validate.Rule("foo.baz", validate.Required),
		)

		require.Empty(t, errs)
	})

	t.Run("Some_Fail", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.bar": "val1",
		})

		errs := validate.Validate(k,
			validate.Rule("foo.bar", validate.Required),
			validate.Rule("foo.baz", validate.Required),
		)

		require.Len(t, errs, 1)

		var ve validate.ValidationError
		require.ErrorAs(t, errs[0], &ve)
		require.Equal(t, "foo.baz", ve.FieldName())
	})

	t.Run("Multiple_Fail", func(t *testing.T) {
		k := newKoanf(t, map[string]any{})

		errs := validate.Validate(k,
			validate.Rule("foo.bar", validate.Required),
			validate.Rule("foo.baz", validate.Required),
		)

		require.Len(t, errs, 2)

		var ve0 validate.ValidationError
		require.ErrorAs(t, errs[0], &ve0)
		require.Equal(t, "foo.bar", ve0.FieldName())

		var ve1 validate.ValidationError
		require.ErrorAs(t, errs[1], &ve1)
		require.Equal(t, "foo.baz", ve1.FieldName())
	})

	t.Run("No_Rules", func(t *testing.T) {
		k := newKoanf(t, map[string]any{})

		errs := validate.Validate(k)

		require.Empty(t, errs)
	})
}

func newKoanf(t *testing.T, data map[string]any) *koanf.Koanf {
	t.Helper()

	const delim = "."

	k := koanf.New(delim)
	err := k.Load(confmap.Provider(data, delim), nil)
	require.NoError(t, err)

	return k
}
