package workerpool_test

import (
	"sync/atomic"
	"testing"

	"github.com/tadvi/workerpool"
)

func TestZero(t *testing.T) {
	wk := workerpool.Run(0, func() {
		t.Errorf("got executed, want no execution")
	})
	wk.Wait()
}

func TestOne(t *testing.T) {
	var executed bool

	wk := workerpool.Run(1, func() {
		executed = true
	})
	wk.Wait()

	if !executed {
		t.Errorf("did not get executed, want one execution")
	}
}

func TestMany(t *testing.T) {

	for i := int32(2); i < 10; i++ {
		var count int32

		wk := workerpool.Run(int(i), func() {
			atomic.AddInt32(&count, 1)
		})
		wk.Wait()

		if count != i {
			t.Errorf("got executed %d times, want %d executions", count, i)
		}

	}
}
