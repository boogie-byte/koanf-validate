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

// Predicate is a function that validates a configuration value, returning an error on failure.
type Predicate func(val any) error

// Required returns a Predicate that fails if the value is nil.
func Required() Predicate {
	return func(val any) error {
		if val == nil {
			return ErrMissingRequiredField
		}

		return nil
	}
}
