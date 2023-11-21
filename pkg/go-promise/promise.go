package gopromise

// Promise provides an awaitable go-routine.
type Promise[T any] interface {
	// Await returns the result or the error of the go-routine.
	Await() (T, error)
}

type promise[T any] struct {
	val  T
	err  error
	task func() (T, error)
	ch   chan bool
}

func (p *promise[T]) Await() (T, error) {
	sta := <-p.ch
	if !sta {
		var none T
		return none, p.err
	}
	return p.val, nil
}

// All takes an array of promises and returns their result.
func All[T any](p []Promise[T]) Promise[[]T] {
	return New(func() ([]T, error) {
		var err error
		ret := make([]T, len(p))

		for i, pr := range p {
			if err != nil {
				pr.Await()
				continue
			}

			val, er := pr.Await()

			if er != nil {
				err = er
			} else {
				ret[i] = val
			}

		}
		if err != nil {
			return nil, err
		}
		return ret, nil
	})
}

// Result contains the promise result.
type Result[T any] struct {
	Val T
	Err error
}

// AllSettled takes an array of promises and returns their result individually.
func AllSettled[T any](p []Promise[T]) Promise[[]*Result[T]] {
	return New(func() ([]*Result[T], error) {
		ret := make([]*Result[T], len(p))

		for i, pr := range p {
			val, err := pr.Await()
			ret[i] = &Result[T]{Val: val, Err: err}
		}

		return ret, nil
	})
}

// New takes a function and returns it in a Promise form.
func New[T any](f func() (T, error)) Promise[T] {
	p := &promise[T]{ch: make(chan bool), task: f}

	go func(pr *promise[T]) {
		pr.val, pr.err = pr.task()
		pr.ch <- pr.err == nil
	}(p)

	return p
}
