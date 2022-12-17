package adaptor

type option func(*Adaptor) *Adaptor

func WithImpl(impl any) option {
	return func(a *Adaptor) *Adaptor {
		a.impl = impl
		return a
	}
}
