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
	"slices"
	"strings"

	"github.com/knadh/koanf/v2"
)

// KV holds a resolved configuration key and its value.
type KV struct {
	Key   string
	Value any
}

// CollectKVs resolves a path into key-value pairs, recursively expanding any "*" wildcards.
func CollectKVs(k *koanf.Koanf, path []string) []KV {
	delim := k.Delim()
	wildcardIdx := slices.Index(path, "*")

	if wildcardIdx == -1 {
		key := strings.Join(path, delim)

		return []KV{
			{
				Key:   key,
				Value: k.Get(key),
			},
		}
	}

	if wildcardIdx == len(path)-1 {
		panic("selector cannot end with wildcard")
	}

	prefix := strings.Join(path[:wildcardIdx], delim)
	mapKeys := k.MapKeys(prefix)

	var kvs []KV
	for _, mapKey := range mapKeys {
		expandedPath := slices.Clone(path)
		expandedPath[wildcardIdx] = mapKey
		subKVs := CollectKVs(k, expandedPath)
		kvs = append(kvs, subKVs...)
	}

	return kvs
}
