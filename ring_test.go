package ring

import (
	// "fmt"
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

func TestCount(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	if r.Count() != 0 {
		t.Fatal("Count on empty ring", r.Count(), "wanted count of", 0)
	}
	for i := 1; i < 11; i++ {
		r.Enqueue(i)
	}
	if r.Count() != 10 {
		t.Fatal("Count on ring with 10 enqueues", r.Count(), "wanted count of", 10)
	}
	for i := 1; i < 11; i++ {
		_ = r.Dequeue()
	}
	if r.Count() != 0 {
		t.Fatal("Count on ring with 10 dequeues", r.Count(), "wanted count of", 0)
	}
}

func TestCountOverrun(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)

	for i := 1; i < 16; i++ {
		r.Enqueue(i)
	}
	if r.Count() != r.Capacity() {
		t.Fatal("Count on ring with 15 enqueues", r.Count(), "wanted count of", r.Capacity())
	}

	for i := 1; i < 21; i++ {
		r.Enqueue(i)
	}
	if r.Count() != r.Capacity() {
		t.Fatal("Count on ring with 35 enqueues", r.Count(), "wanted count of", r.Capacity())
	}

	for i := 1; i < 11; i++ {
		_ = r.Dequeue()
	}
	if r.Count() != 0 {
		t.Fatal("Count on ring with 10 dequeues", r.Count(), "wanted count of", 0)
	}

}

func TestCountOnEmpty(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)

	_ = r.Dequeue()
	if r.Count() != 0 {
		t.Fatal("Count on empty ring after 1 dequeues", r.Count(), "wanted count of", 0)
	}
}

func TestCountArr(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)

	for i := 1; i < 6; i++ {
		r.Enqueue(i)
	}
	slice := r.Values()
	if r.Count() != len(slice) {
		t.Fatal("Length of return Values() of ring", len(slice), "wanted count of", r.Count())
	}

	for i := 1; i < 11; i++ {
		r.Enqueue(i)
	}
	slice = r.Values()
	if r.Count() != len(slice) {
		t.Fatal("Length of return Values() of ring", len(slice), "wanted count of", r.Count())
	}
	
	for i := 11; i < 21; i++ {
		r.Enqueue(i)
	}
	slice = r.Values()
	if r.Count() != len(slice) {
		t.Fatal("Length of return Values() of ring", len(slice), "wanted count of", r.Count())
	}
}

