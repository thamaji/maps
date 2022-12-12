package maps

import "math/rand"

// 指定したキーの値を返す。
func Get[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}

// 指定したキーの値を返す。無い場合はvを返す。
func GetOrElse[K comparable, V any](m map[K]V, k K, v V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return v
}

// マップのキーのスライスを返す。
func Keys[K comparable, V any](m map[K]V) []K {
	dst := make([]K, 0, len(m))
	for k := range m {
		dst = append(dst, k)
	}
	return dst
}

// マップの値のスライスを返す。
func Values[K comparable, V any](m map[K]V) []V {
	dst := make([]V, 0, len(m))
	for _, v := range m {
		dst = append(dst, v)
	}
	return dst
}

// マップのエントリーのスライスを返す。
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	dst := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		dst = append(dst, Entry[K, V]{k, v})
	}
	return dst
}

// 要素をすべて削除する。
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		m[k] = *new(V)
		delete(m, k)
	}
}

// 要素をすべてコピーしたマップを返す。
func Clone[K comparable, V any](m map[K]V) map[K]V {
	dst := make(map[K]V, len(m))
	for k, v := range m {
		dst[k] = v
	}
	return dst
}

// 値を１つランダムに返す。
func Sample[K comparable, V any](m map[K]V, r *rand.Rand) V {
	i, n := 0, r.Intn(len(m))
	var k K
	for k = range m {
		if i == n {
			break
		}
		i++
	}
	return m[k]
}

// ひとつでも一致する値が存在したらtrueを返す。
func Contains[K comparable, V comparable](m map[K]V, v V) bool {
	for _, v1 := range m {
		if v1 == v {
			return true
		}
	}
	return false
}

// ひとつでも条件を満たす要素が存在したらtrueを返す。
func ContainsBy[K comparable, V comparable](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

// 一致する値の数を返す。
func Count[K comparable, V comparable](m map[K]V, v V) int {
	c := 0
	for _, v1 := range m {
		if v1 == v {
			c++
		}
	}
	return c
}

// 条件を満たす要素の数を返す。
func CountBy[K comparable, V comparable](m map[K]V, f func(K, V) bool) int {
	c := 0
	for k, v := range m {
		if f(k, v) {
			c++
		}
	}
	return c
}

// ゼロ値の要素を除いたマップを返す。
func Clean[K comparable, V comparable](m map[K]V) map[K]V {
	zero := *new(V)
	dst := make(map[K]V, len(m))
	for k, v := range m {
		if v != zero {
			dst[k] = v
		}
	}
	return dst
}

// 値を変換したマップを返す。
func Map[K comparable, V1 any, V2 any](m map[K]V1, f func(K, V1) V2) map[K]V2 {
	dst := make(map[K]V2, len(m))
	for k, v := range m {
		dst[k] = f(k, v)
	}
	return dst
}

// 値を順に演算する。
func Reduce[K comparable, V any](m map[K]V, f func(V, K, V) V) V {
	v := *new(V)
	head := true
	for k, v1 := range m {
		if head {
			v = v1
			head = false
			continue
		}
		v = f(v, k, v1)
	}
	return v
}

// 初期値と値を順に演算する。
func Fold[K comparable, V1 any, V2 any](m map[K]V1, v V2, f func(V2, K, V1) V2) V2 {
	for k, v1 := range m {
		v = f(v, k, v1)
	}
	return v
}

// 値と一致する要素を返す。
func Find[K comparable, V comparable](m map[K]V, v V) (Entry[K, V], bool) {
	for k, v1 := range m {
		if v == v1 {
			return Entry[K, V]{k, v1}, true
		}
	}
	return Entry[K, V]{}, false
}

// 条件を満たす要素を返す。
func FindBy[K comparable, V any](m map[K]V, f func(K, V) bool) (Entry[K, V], bool) {
	for k, v := range m {
		if f(k, v) {
			return Entry[K, V]{k, v}, true
		}
	}
	return Entry[K, V]{}, false
}

// 値の一致する要素だけのマップを返す。
func Filter[K comparable, V comparable](m map[K]V, v V) map[K]V {
	dst := make(map[K]V, len(m))
	for k, v1 := range m {
		if v1 == v {
			dst[k] = v1
		}
	}
	return dst
}

// 条件を満たす要素だけのマップを返す。
func FilterBy[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	dst := make(map[K]V, len(m))
	for k, v := range m {
		if f(k, v) {
			dst[k] = v
		}
	}
	return dst
}

// 値の一致しない要素だけのマップを返す。
func FilterNot[K comparable, V comparable](m map[K]V, v V) map[K]V {
	dst := make(map[K]V, len(m))
	for k, v1 := range m {
		if v1 != v {
			dst[k] = v1
		}
	}
	return dst
}

// 条件を満たさない要素だけのマップを返す。
func FilterNotBy[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	dst := make(map[K]V, len(m))
	for k, v := range m {
		if !f(k, v) {
			dst[k] = v
		}
	}
	return dst
}

// 条件を満たす要素を変換したマップを返す。
func Collect[K comparable, V1 any, V2 any](m map[K]V1, f func(K, V1) (V2, bool)) map[K]V2 {
	dst := make(map[K]V2, len(m))
	for k, v := range m {
		if v1, ok := f(k, v); ok {
			dst[k] = v1
		}
	}
	return dst
}

// 値の一致するマップと一致しないマップを返す。
func Partition[K comparable, V comparable](m map[K]V, v V) (map[K]V, map[K]V) {
	dst1 := make(map[K]V, len(m)/2)
	dst2 := make(map[K]V, len(m)/2)
	for k, v1 := range m {
		if v1 == v {
			dst1[k] = v1
		} else {
			dst2[k] = v1
		}
	}
	return dst1, dst2
}

// 条件を満たすマップと満たさないマップを返す。
func PartitionBy[K comparable, V any](m map[K]V, f func(K, V) bool) (map[K]V, map[K]V) {
	dst1 := make(map[K]V, len(m)/2)
	dst2 := make(map[K]V, len(m)/2)
	for k, v := range m {
		if f(k, v) {
			dst1[k] = v
		} else {
			dst2[k] = v
		}
	}
	return dst1, dst2
}

// 値の合計を返す。
func Sum[K comparable, V ordered | complex](m map[K]V) V {
	return Reduce(m, func(sum V, k K, v V) V {
		return sum + v
	})
}

// 値を変換して合計を返す。
func SumBy[K comparable, V1 any, V2 ordered | complex](m map[K]V1, f func(K, V1) V2) V2 {
	return Fold(m, *new(V2), func(sum V2, k K, v V1) V2 {
		return sum + f(k, v)
	})
}

// 最大の値を返す。
func Max[K comparable, V ordered](m map[K]V) V {
	return Reduce(m, func(max V, k K, v V) V {
		if max < v {
			return v
		}
		return max
	})
}

// 値を変換して最大の値を返す。
func MaxBy[K comparable, V1 any, V2 ordered](m map[K]V1, f func(K, V1) V2) V2 {
	head := true
	return Fold(m, *new(V2), func(max V2, k K, v V1) V2 {
		if head {
			head = false
			return f(k, v)
		}
		if v := f(k, v); max < v {
			return v
		}
		return max
	})
}

// 最小の値を返す。
func Min[K comparable, V ordered](m map[K]V) V {
	return Reduce(m, func(min V, k K, v V) V {
		if min > v {
			return v
		}
		return min
	})
}

// 最小の値を返す。
func MinBy[K comparable, V1 any, V2 ordered](m map[K]V1, f func(K, V1) V2) V2 {
	head := true
	return Fold(m, *new(V2), func(min V2, k K, v V1) V2 {
		if head {
			head = false
			return f(k, v)
		}
		if v := f(k, v); min > v {
			return v
		}
		return min
	})
}

// すべての要素に値を代入する。
func Fill[K comparable, V any](m map[K]V, v V) {
	for k := range m {
		m[k] = v
	}
}

// すべての要素にゼロ値を代入する。
func FillZero[K comparable, V any](m map[K]V) {
	zero := *new(V)
	for k := range m {
		m[k] = zero
	}
}

// すべての要素に関数の実行結果を代入する。
func FillBy[K comparable, V any](m map[K]V, f func(K) V) {
	for k := range m {
		m[k] = f(k)
	}
}
