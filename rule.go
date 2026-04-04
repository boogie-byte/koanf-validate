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
	"strings"

	"github.com/knadh/koanf/v2"
)

// RuleSpec binds a selector to one or more predicates for validation.
type RuleSpec struct {
	selector   string
	predicates []Predicate
}

// Rule creates a RuleSpec for the given selector and predicates.
// The selector is a dot-separated path that may contain "*" wildcards to match multiple keys.
func Rule(selector string, predicates ...Predicate) RuleSpec {
	return RuleSpec{
		selector:   selector,
		predicates: predicates,
	}
}

// Apply runs all predicates against the values matched by the selector and returns any validation errors.
func (r RuleSpec) Apply(k *koanf.Koanf) []error {
	delim := k.Delim()
	path := strings.Split(r.selector, delim)
	kvs := CollectKVs(k, path)

	var errs []error
	for _, kv := range kvs {
		for _, predicate := range r.predicates {
			if err := predicate(kv.Value); err != nil {
				errs = append(errs, NewValidationError(kv.Key, err))
			}
		}
	}

	return errs
}
