package nano

import (
	"math/rand"
	"testing"

	"github.com/LucaHhx/nano/session"
)

func TestChannel_Add(t *testing.T) {
	c := NewGroup("test_add")

	var paraCount = 100
	w := make(chan bool, paraCount)
	for i := 0; i < paraCount; i++ {
		go func(id int) {
			s := session.New(nil)
			err := s.Bind(int64(id + 1))
			if err != nil {
				return
			}
			err = c.Add(s)
			if err != nil {
				return
			}
			w <- true
		}(i)
	}

	for i := 0; i < paraCount; i++ {
		<-w
	}

	if c.Count() != paraCount {
		t.Fatalf("count expect: %d, got: %d", paraCount, c.Count())
	}

	n := rand.Int63n(int64(paraCount)) + 1
	if !c.Contains(n) {
		t.Fail()
	}

	// leave
	err := c.LeaveAll()
	if err != nil {
		return
	}
	if c.Count() != 0 {
		t.Fail()
	}
}
