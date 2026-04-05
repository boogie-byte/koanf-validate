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
	"fmt"
	"reflect"
)

// concreteValueOf unwraps pointers and interfaces and returns the underlying concrete reflect.Value.
func concreteValueOf(val any) reflect.Value {
	rv := reflect.ValueOf(val)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	return rv
}

// lenOf returns the length of a string, slice, or map, dereferencing pointers and interfaces first.
// It returns 0 for nil values and -1 with an error for unsupported types.
func lenOf(val any) (int, error) {
	if val == nil {
		return 0, nil
	}

	rv := concreteValueOf(val)

	switch rv.Kind() {
	case reflect.String, reflect.Slice, reflect.Map:
		return rv.Len(), nil
	default:
		return -1, fmt.Errorf("expected string, slice, or map, got %T", val)
	}
}
