package maps

// エントリーのスライスをマップに変換する。
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	m := make(map[K]V, len(entries))
	for _, e := range entries {
		m[e.Key] = e.Value
	}
	return m
}
