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

// Package validate provides rule-based validation for koanf configuration.
//
// It allows defining validation rules with selectors that support wildcard
// matching to validate multiple configuration keys at once.
//
//	errs := validate.Validate(k,
//		validate.Rule("db.host", validate.Required()),
//		validate.Rule("services.*.port", validate.Required()),
//	)
package validate
