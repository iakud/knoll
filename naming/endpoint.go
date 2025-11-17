package naming

type Attributes struct {
	m map[any]any
}

func New(key, value any) *Attributes {
	return &Attributes{m: map[any]any{key: value}}
}

func (a *Attributes) WithValue(key, value any) *Attributes {
	if a == nil {
		return New(key, value)
	}
	n := &Attributes{m: make(map[any]any, len(a.m)+1)}
	for k, v := range a.m {
		n.m[k] = v
	}
	n.m[key] = value
	return n
}

func (a *Attributes) Value(key any) any {
	if a == nil {
		return nil
	}
	return a.m[key]
}

type Endpoint struct {
	Addr       string
	Attributes *Attributes
}
