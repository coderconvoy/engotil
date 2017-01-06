package observe

import "testing"

type ee struct {
	name string
	Observable
}

type er struct {
	n int
}

func (e *er) Notify(o interface{}) {
	p, ok := o.(*ee)
	if !ok {
		return
	}
	e.n++
	p.name = "poo"
}

func Test_observe(t *testing.T) {

	e := ee{name: "hello"}
	e.SetORef(&e)

	r := er{0}

	f := FlagObserver(false)

	e.Observe(&r)
	e.Observe(&f)
	e.NotifyAll()

	if !f {
		t.Log("f should be true")
		t.Fail()
	}

	if r.n != 1 {
		t.Logf("r should be 1: got %d", r.n)
		t.Fail()
	}

	e.Unobserve(&r)
	e.NotifyAll()
	if r.n != 1 {
		t.Logf("r should still be 1: got %d", r.n)
		t.Fail()
	}

	if e.name != "poo" {
		t.Logf("name not updated: expected \"poo\", got : %s", e.name)
		t.Fail()
	}
}
