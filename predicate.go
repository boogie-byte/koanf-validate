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

import "fmt"

// Predicate is a function that validates a configuration value, returning an error on failure.
type Predicate func(val any) error

// Required is a Predicate that fails if the value is nil.
func Required(val any) error {
	if val == nil {
		return ErrMissingRequiredField
	}

	return nil
}

// MinLen returns a Predicate that fails if the length of a string, slice, or map is less than n.
// It panics if n is negative.
func MinLen(n int) Predicate {
	if n < 0 {
		panic("MinLen: n must not be negative")
	}

	return func(val any) error {
		l, err := lenOf(val)
		if err != nil {
			return err
		}

		if l < n {
			return fmt.Errorf("length %d is less than minimum %d", l, n)
		}

		return nil
	}
}

// MaxLen returns a Predicate that fails if the length of a string, slice, or map exceeds n.
// It panics if n is negative.
func MaxLen(n int) Predicate {
	if n < 0 {
		panic("MaxLen: n must not be negative")
	}

	return func(val any) error {
		l, err := lenOf(val)
		if err != nil {
			return err
		}

		if l > n {
			return fmt.Errorf("length %d exceeds maximum %d", l, n)
		}

		return nil
	}
}
