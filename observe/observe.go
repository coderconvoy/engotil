package observe

type Observer interface {
	Notify(interface{})
}

type Observable struct {
	Ref      interface{}
	watchers []Observer
}

type FlagObserver bool

func (p *Observable) SetORef(i interface{}) {
	p.Ref = i
}

func (p *Observable) Observe(w Observer) {
	p.watchers = append(p.watchers, w)
}

func (p *Observable) Unobserve(w Observer) {
	var res []Observer
	for _, v := range p.watchers {
		if v != w {
			res = append(res, v)
		}
	}
	p.watchers = res
}

func (p *Observable) NotifyAll() {
	for _, w := range p.watchers {
		w.Notify(p.Ref)
	}

}

func (f *FlagObserver) Notify(p interface{}) {
	*f = true
}
