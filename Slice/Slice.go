package Slice

import "errors"

func Insert[T any](slice []T, elem T, idx int) ([]T, error) {
	if idx < 0 || idx > len(slice) {
		return nil, errors.New("incorrect index")
	}

	if len(slice) == cap(slice) {
		var res []T
		if cap(slice) > 256 {
			res = make([]T, len(slice)+1, int(float64(cap(slice))*1.25))
		} else {
			res = make([]T, len(slice)+1, cap(slice)*2)
		}
		copy(res[:idx], slice[:idx])
		res[idx] = elem
		copy(res[idx+1:], slice[idx:])
		return res, nil
	}

	slice = slice[:len(slice)+1]
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = elem
	return slice, nil
}

func Delete[T any](slice []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(slice) {
		return nil, errors.New("index out of range")
	}

	copy(slice[idx:], slice[idx+1:])
	slice = slice[:len(slice)-1]
	return slice, nil
}
