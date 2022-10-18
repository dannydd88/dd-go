package dd

// An implement of Closure(https://en.wikipedia.org/wiki/Closure_(computer_programming))
type Closure func() any

// Implement |Task| interface
func (f Closure) Run() any {
	return f()
}

// === Bind
// Helper function to construct a closure using function currying

func Bind0[R any](f func() R) Closure {
	return func() any {
		return f()
	}
}

func Bind1[T any, R any](f func(T) R, p T) Closure {
	return func() any {
		return f(p)
	}
}

func Bind2[T1, T2, R any](f func(T1, T2) R, p1 T1, p2 T2) Closure {
	return func() any {
		return f(p1, p2)
	}
}

func Bind3[T1, T2, T3, R any](f func(T1, T2, T3) R, p1 T1, p2 T2, p3 T3) Closure {
	return func() any {
		return f(p1, p2, p3)
	}
}

func Bind4[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, p1 T1, p2 T2, p3 T3, p4 T4) Closure {
	return func() any {
		return f(p1, p2, p3, p4)
	}
}

func Bind5[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, p1 T1, p2 T2, p3 T3, p4 T4, p5 T5) Closure {
	return func() any {
		return f(p1, p2, p3, p4, p5)
	}
}

func Bind6[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, p1 T1, p2 T2, p3 T3, p4 T4, p5 T5, p6 T6) Closure {
	return func() any {
		return f(p1, p2, p3, p4, p5, p6)
	}
}

func Bind7[T1, T2, T3, T4, T5, T6, T7, R any](f func(T1, T2, T3, T4, T5, T6, T7) R, p1 T1, p2 T2, p3 T3, p4 T4, p5 T5, p6 T6, p7 T7) Closure {
	return func() any {
		return f(p1, p2, p3, p4, p5, p6, p7)
	}
}

func Bind8[T1, T2, T3, T4, T5, T6, T7, T8, R any](f func(T1, T2, T3, T4, T5, T6, T7, T8) R, p1 T1, p2 T2, p3 T3, p4 T4, p5 T5, p6 T6, p7 T7, p8 T8) Closure {
	return func() any {
		return f(p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

// === Bind
