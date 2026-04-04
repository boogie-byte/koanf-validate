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

	"github.com/stretchr/testify/require"

	"github.com/boogie-byte/koanf-validate"
)

func TestCollectKVs(t *testing.T) {
	t.Run("NoWildcard", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.bar": "val1",
		})

		kvs := validate.CollectKVs(k, []string{"foo", "bar"})

		require.Equal(t, []validate.KV{
			{Key: "foo.bar", Value: "val1"},
		}, kvs)
	})

	t.Run("NoWildcard_Missing_Key", func(t *testing.T) {
		k := newKoanf(t, map[string]any{})

		kvs := validate.CollectKVs(k, []string{"foo", "bar"})

		require.Equal(t, []validate.KV{
			{Key: "foo.bar", Value: nil},
		}, kvs)
	})

	t.Run("SingleWildcard", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.a.bar": "val1",
			"foo.b.bar": "val2",
			"foo.c.bar": "val3",
		})

		kvs := validate.CollectKVs(k, []string{"foo", "*", "bar"})

		require.ElementsMatch(t, []validate.KV{
			{Key: "foo.a.bar", Value: "val1"},
			{Key: "foo.b.bar", Value: "val2"},
			{Key: "foo.c.bar", Value: "val3"},
		}, kvs)
	})

	t.Run("MultipleWildcards", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.a.bar.x.baz": "val1",
			"foo.a.bar.y.baz": "val2",
			"foo.b.bar.x.baz": "val3",
		})

		kvs := validate.CollectKVs(k, []string{"foo", "*", "bar", "*", "baz"})

		require.ElementsMatch(t, []validate.KV{
			{Key: "foo.a.bar.x.baz", Value: "val1"},
			{Key: "foo.a.bar.y.baz", Value: "val2"},
			{Key: "foo.b.bar.x.baz", Value: "val3"},
		}, kvs)
	})

	t.Run("SingleWildcard_NoMatches", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"other.key": "val1",
		})

		kvs := validate.CollectKVs(k, []string{"foo", "*", "bar"})

		require.Empty(t, kvs)
	})

	t.Run("WildcardAtEnd_Panics", func(t *testing.T) {
		k := newKoanf(t, map[string]any{
			"foo.a.bar": "val1",
		})

		require.Panics(t, func() {
			validate.CollectKVs(k, []string{"foo", "*"})
		})
	})
}
