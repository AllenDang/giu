package giu

type queue struct {
	items chan []func()
	empty chan bool
}

func newQueue() *queue {
	q := &queue{
		items: make(chan []func(), 1),
		empty: make(chan bool, 1),
	}
	q.empty <- true
	return q
}

func (q *queue) add(f func()) {
	var items []func()
	select {
	case <-q.empty:
	case items = <-q.items:
	}
	q.items <- append(items, f)
}

func (q *queue) get(done <-chan struct{}) (func(), bool) {
	var items []func()
	select {
	case <-done:
		return nil, false
	case items = <-q.items:
	}
	f, n := items[0], copy(items, items[1:])
	if items = items[:n]; len(items) == 0 {
		q.empty <- true
	} else {
		q.items <- items
	}
	return f, true
}

var callQueue = newQueue()

// Run enables mainthread package functionality. To use mainthread package, put your main function
// code into the run function (the argument to Run) and simply call Run from the real main function.
//
// Run returns when run (argument) function finishes.
func Run(run func()) {
	done := make(chan struct{})
	go func() { run(); close(done) }()
	for {
		if f, ok := callQueue.get(done); ok {
			f()
		} else {
			break
		}
	}
}

// CallNonBlock queues function f on the main thread and returns immediately. Does not wait until f
// finishes.
func CallNonBlock(f func()) { callQueue.add(f) }

// Call queues function f on the main thread and blocks until the function f finishes.
func Call(f func()) {
	done := make(chan struct{})
	callQueue.add(func() { f(); done <- struct{}{} })
	<-done
}

// CallErr queues function f on the main thread and returns an error returned by f.
func CallErr(f func() error) error {
	errChan := make(chan error)
	callQueue.add(func() { errChan <- f() })
	return <-errChan
}

// CallVal queues function f on the main thread and returns a value returned by f.
func CallVal(f func() interface{}) interface{} {
	respChan := make(chan interface{})
	callQueue.add(func() { respChan <- f() })
	return <-respChan
}
