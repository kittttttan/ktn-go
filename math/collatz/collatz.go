package collatz

import (
	"errors"
	"math"
)

// 与えられた n についてコラッツの数列を計算し個数 cnt と最大値 max を返す
func Count(n uint64) (cnt uint64, max uint64, err error) {
	if n < 1 {
		return cnt, max, nil
	}

	var limit uint64 = math.MaxUint64/3 - 1

	max = n
	cnt = 1
	for n > 1 {
		if cnt == math.MaxUint64 {
			return 0, 0, errors.New("overflow cnt")
		}

		cnt += 1
		if n%2 == 0 {
			n /= 2
		} else {
			if n > limit {
				return 0, 0, errors.New("overflow max")
			}

			n = 3*n + 1
			if n > max {
				max = n
			}
		}
	}

	return cnt, max, nil
}

// 与えられた n についてコラッツの数列を計算するイテレーター
func Iter(n uint64) func(func(uint64) bool) {
	return func(yield func(uint64) bool) {
		var limit uint64 = math.MaxUint64/3 - 1
		if !yield(n) {
			return
		}

		for n > 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				if n > limit {
					return
				}

				n = 3*n + 1
			}
			if !yield(n) {
				return
			}
		}
	}
}
