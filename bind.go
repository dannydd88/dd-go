package dd

// An implement of Closure(https://en.wikipedia.org/wiki/Closure_(computer_programming))
type Closure func() any

// Implement |Task| interface
func (f *Closure) Run() any {
	return (*f)()
}

// === Bind
// Helper function to construct a closure using function currying

func Bind0[R any](f func() R) *Closure {
	c := Closure(func() any {
		return f()
	})
	return &c
}

func Bind1[T any, R any](f func(T) R, p T) *Closure {
	c := Closure(func() any {
		return f(p)
	})
	return &c
}

func Bind2[T1, T2, R any](f func(T1, T2) R, p1 T1, p2 T2) *Closure {
	c := Closure(func() any {
		return f(p1, p2)
	})
	return &c
}

func Bind3[T1, T2, T3, R any](
	f func(T1, T2, T3) R,
	p1 T1, p2 T2, p3 T3) *Closure {
	c := Closure(func() any {
		return f(p1, p2, p3)
	})
	return &c
}

func Bind4[T1, T2, T3, T4, R any](
	f func(T1, T2, T3, T4) R,
	p1 T1, p2 T2, p3 T3, p4 T4) *Closure {
	c := Closure(func() any {
		return f(p1, p2, p3, p4)
	})
	return &c
}

func Bind5[T1, T2, T3, T4, T5, R any](
	f func(T1, T2, T3, T4, T5) R,
	p1 T1, p2 T2, p3 T3, p4 T4, p5 T5) *Closure {
	c := Closure(func() any {
		return f(p1, p2, p3, p4, p5)
	})
	return &c
}

func Bind6[T1, T2, T3, T4, T5, T6, R any](
	f func(T1, T2, T3, T4, T5, T6) R,
	p1 T1, p2 T2, p3 T3, p4 T4, p5 T5, p6 T6) *Closure {
	c := Closure(func() any {
		return f(p1, p2, p3, p4, p5, p6)
	})
	return &c
}

func Bind7[T1, T2, T3, T4, T5, T6, T7, R any](
	f func(T1, T2, T3, T4, T5, T6, T7) R,
	p1 T1, p2 T2, p3 T3, p4 T4, p5 T5, p6 T6, p7 T7) *Closure {
	c := Closure(func() any {
		return f(p1, p2, p3, p4, p5, p6, p7)
	})
	return &c
}

func Bind8[T1, T2, T3, T4, T5, T6, T7, T8, R any](
	f func(T1, T2, T3, T4, T5, T6, T7, T8) R,
	p1 T1, p2 T2, p3 T3, p4 T4, p5 T5, p6 T6, p7 T7, p8 T8) *Closure {
	c := Closure(func() any {
		return f(p1, p2, p3, p4, p5, p6, p7, p8)
	})
	return &c
}

// === Bind
