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
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/boogie-byte/koanf-validate"
)

func TestRuleSpec_Apply(t *testing.T) {
	t.Run("Predicate_Pass", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.bar": "val1",
		})

		rule := validate.Rule("foo.bar", validate.Required)
		errs := rule.Apply(k)

		require.Empty(t, errs)
	})

	t.Run("Predicate_Fail", func(t *testing.T) {
		k := newKoanf(t, map[string]any{})

		rule := validate.Rule("foo.bar", validate.Required)
		errs := rule.Apply(k)

		require.Len(t, errs, 1)

		var ve validate.ValidationError
		require.ErrorAs(t, errs[0], &ve)
		require.Equal(t, "foo.bar", ve.FieldName())
		require.ErrorIs(t, ve, validate.ErrMissingRequiredField)
	})

	t.Run("Multiple_Predicates", func(t *testing.T) {
		k := newKoanf(t, map[string]any{})

		errCustom := errors.New("custom error")
		alwaysFail := func(val any) error { return errCustom }

		rule := validate.Rule("foo.bar", validate.Required, alwaysFail)
		errs := rule.Apply(k)

		require.Len(t, errs, 2)
		require.ErrorIs(t, errs[0], validate.ErrMissingRequiredField)
		require.ErrorIs(t, errs[1], errCustom)
	})
}
