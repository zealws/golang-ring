package ring

import (
	"sync"
	"testing"
)

func TestSetsSize(t *testing.T) {
	r := &Ring{}
	r.SetCapacity(10)
	if r.Capacity() != 10 {
		t.Fatal("Size of ring was not 10", r.Capacity())
	}
}

func TestSavesSomeData(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 7; i++ {
		r.Enqueue(i)
	}
	for i := 0; i < 7; i++ {
		x := r.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

func TestReusesBuffer(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 7; i++ {
		r.Enqueue(i)
	}
	for i := 0; i < 7; i++ {
		r.Dequeue()
	}
	for i := 7; i < 14; i++ {
		r.Enqueue(i)
	}
	for i := 7; i < 14; i++ {
		x := r.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

func TestOverflowsBuffer(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 20; i++ {
		r.Enqueue(i)
	}
	for i := 10; i < 20; i++ {
		x := r.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

func TestPartiallyOverflows(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 15; i++ {
		r.Enqueue(i)
	}
	for i := 5; i < 15; i++ {
		x := r.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
	}
}

func TestPeeks(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 10; i++ {
		r.Enqueue(i)
	}
	for i := 0; i < 10; i++ {
		r.Peek()
		r.Peek()
		x1 := r.Peek()
		x := r.Dequeue()
		if x != i {
			t.Fatal("Unexpected response", x, "wanted", i)
		}
		if x1 != x {
			t.Fatal("Unexpected response", x1, "wanted", x)
		}
	}
}

func TestConstructsArr(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	v := r.Values()
	if len(v) != 0 {
		t.Fatal("Unexpected values", v, "wanted len of", 0)
	}
	for i := 1; i < 21; i++ {
		r.Enqueue(i)
		l := i
		if l > 10 {
			l = 10
		}
		v = r.Values()
		if len(v) != l {
			t.Fatal("Unexpected values", v, "wanted len of", l, "index", i)
		}
	}
}

func TestContentSize(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)

	for i := 1; i < 10; i++ {
		r.Enqueue(i)
		s := r.ContentSize()
		if s != i {
			t.Fatal("Unexpected content size", s, "wanted", i)
		}
	}

	for i := 1; i < 5; i++ {
		r.Enqueue(i)
		s := r.ContentSize()
		if s != 10 {
			t.Fatal("Unexpected content size", s, "wanted 10")
		}
	}

	for i := 9; i > 0; i-- {
		r.Dequeue()
		s := r.ContentSize()
		if s != i {
			t.Fatal("Unexpected content size", s, "wanted", i)
		}
	}
}

func TestConcurrency(t *testing.T) {
	wg := sync.WaitGroup{}

	r := Ring{}
	r.SetCapacity(128)

	for i := 0; i < 2048; i++ {
		wg.Add(1)
		go func(i int) {
			for x := 0; x < 100; x++ {
				r.Enqueue(x)
			}

			if i%10 == 0 {
				r.ContentSize()
			}

			if (i+1)%10 == 0 {
				r.Capacity()
			}

			if (i+2)%10 == 0 {
				r.Peek()
			}

			if (i+3)%10 == 0 {
				r.Values()
			}

			if i%10 == 0 {
				r.SetCapacity(r.Capacity() + 1)
			}

			for x := 0; x < 125; x++ {
				r.Dequeue()
			}

			wg.Done()
		}(i)

	}

	wg.Wait()
}
