package atomicCounter

import "testing"

func TestIncr(t *testing.T) {
	counter := &Counter{}
	for i := 0; i < 5; i++ {
		counter.Incr()
	}
	if counter.Get() != 5 {
		t.Errorf("Counter is in unexpected state ", counter.Get())
	}

}

func TestDecr(t *testing.T) {
	counter := &Counter{}
	for i := 0; i < 5; i++ {
		counter.Incr()
	}
	for i := 0; i < 4; i++ {
		counter.Decr()
	}
	if counter.Get() != 1 {
		t.Errorf("Counter is in unexpected state ", counter.Get())
	}

}

func TestIncrGoRoutine(t *testing.T) {

}

func TestDecrGoRoutine(t *testing.T) {

}
