package Slice

import "errors"

//	insert value
func Insert[T any](slice []T, elem T, idx int) ([]T, error) {
	l, c := len(slice), cap(slice)
	if idx < 0 || idx > l {
		return nil, errors.New("incorrect index")
	}
	if l == c {
		var res []T
		if c > 256 {
			res = make([]T, l+1, int(float64(c)*1.25))
		} else {
			res = make([]T, l+1, c*2)
		}
		copy(res[:idx], slice[:idx])
		res[idx] = elem
		copy(res[idx+1:], slice[idx:])
		return res, nil
	}

	slice = slice[:l+1]
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = elem
	return slice, nil
}

//	delete value
func Delete[T any](slice []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(slice) {
		return nil, errors.New("index out of range")
	}

	copy(slice[idx:], slice[idx+1:])
	slice = slice[:len(slice)-1]
	return slice, nil
}

//	convert value type from Src to Dst
func Map[Src any, Dst any](beg int, count int, src []Src, m func(src Src) Dst) ([]Dst, error) {
	if beg < 0 || beg > len(src) || beg+count > len(src) {
		return nil, errors.New("invalidate beg")
	}

	dst := make([]Dst, count)
	for i := 0; i < count; i++ {
		dst[i] = m(src[beg+i])
	}
	return dst, nil
}

//	return a slice that conforms to the logic function
func Filter[T any](slice []T, pred func(T) bool) []T {
	var res []T
	for _, val := range slice {
		if pred(val) {
			res = append(res, val)
		}
	}
	return res
}

//	return an accumulated value in the type of R
func Reduce[T any, R any](slice []T, init R, f func(T, R) R) R {
	acc := init
	for _, val := range slice {
		acc = f(val, acc)
	}
	return acc
}
