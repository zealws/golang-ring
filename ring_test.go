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

func TestHollowBufferLength(t *testing.T) {
	r := Ring{}
	if r.Length() != 0 {
		t.Fatal("Length expected to be 0, was ", r.Length())
	}
}

func TestEmptyBufferLength(t *testing.T) {
	r := Ring{}
	r.SetCapacity(1)
	if r.Length() != 0 {
		t.Fatal("Length expected to be 0, was ", r.Length())
	}
}

func TestOneElementBufferLength(t *testing.T) {
	r := Ring{}
	r.SetCapacity(2)
	r.Enqueue(5)
	if r.Length() != 1 {
		t.Fatal("Length expected to be 1, was ", r.Length())
	}
}

func TestFullBufferLength(t *testing.T) {
	r := Ring{}
	r.SetCapacity(2)
	r.Enqueue(5)
	r.Enqueue(6)
	if r.Length() != 2 {
		t.Fatal("Length expected to be 2, was ", r.Length())
	}
}

func TestRolledOverBufferLength(t *testing.T) {
	r := Ring{}
	r.SetCapacity(2)
	r.Enqueue(5)
	r.Enqueue(6)
	r.Enqueue(7)
	if r.Length() != 2 {
		t.Fatal("Length expected to be 2, was ", r.Length())
	}
}

func TestRollOverAndFillLength(t *testing.T) {
	r := Ring{}
	r.SetCapacity(2)
	r.Enqueue(5)
	r.Enqueue(6)
	r.Enqueue(7)
	r.Enqueue(8)
	if r.Length() != 2 {
		t.Fatal("Length expected to be 2, was ", r.Length())
	}
}

func TestDequeReducesLength(t *testing.T) {
	r := Ring{}
	r.SetCapacity(3)
	r.Enqueue(5)
	r.Enqueue(6)
	r.Enqueue(7)
	r.Enqueue(8)
	r.Dequeue()
	if r.Length() != 2 {
		t.Fatal("Length expected to be 2, was ", r.Length())
	}
}
