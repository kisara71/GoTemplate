package kbase

func ToPtr[T any](src T) *T {
	return &src
}

func Deref[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

func IsZero[T comparable](val T) bool {
	var temp T
	return temp == val
}

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
