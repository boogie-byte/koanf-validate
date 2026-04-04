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

func TestValidationError(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		err := validate.NewValidationError("foo.bar", validate.ErrMissingRequiredField)

		require.Equal(t, "foo.bar: missing required field", err.Error())
	})

	t.Run("FieldName", func(t *testing.T) {
		err := validate.NewValidationError("foo.bar", validate.ErrMissingRequiredField)

		require.Equal(t, "foo.bar", err.FieldName())
	})

	t.Run("Unwrap", func(t *testing.T) {
		err := validate.NewValidationError("foo.bar", validate.ErrMissingRequiredField)

		require.ErrorIs(t, err, validate.ErrMissingRequiredField)
	})

	t.Run("As", func(t *testing.T) {
		err := validate.NewValidationError("foo.bar", validate.ErrMissingRequiredField)
		wrapped := errors.Join(err)

		var ve validate.ValidationError
		require.ErrorAs(t, wrapped, &ve)
		require.Equal(t, "foo.bar", ve.FieldName())
	})
}
