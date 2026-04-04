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
	"errors"
)

// ErrMissingRequiredField is returned when a required configuration field is nil.
var ErrMissingRequiredField = errors.New("missing required field")

// ValidationError represents a validation failure for a specific configuration field.
type ValidationError struct {
	fieldName string
	err       error
}

var _ error = (*ValidationError)(nil)

// Error returns the field name and the underlying error message.
func (err ValidationError) Error() string {
	return err.fieldName + ": " + err.err.Error()
}

// Unwrap returns the underlying error for use with errors.Is and errors.As.
func (err ValidationError) Unwrap() error {
	return err.err
}

// FieldName returns the path to a field that failed validation.
func (err ValidationError) FieldName() string {
	return err.fieldName
}

// NewValidationError creates a ValidationError for the given field name and cause.
func NewValidationError(fieldName string, err error) ValidationError {
	return ValidationError{
		fieldName: fieldName,
		err:       err,
	}
}
