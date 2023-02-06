package adaptor

import (
	. "github.com/black-desk/lib/go/errwrap"
	"github.com/black-desk/notels/pkg/jsonrpc2"
)

type Adaptor struct {
	impl    any
	methods map[string]jsonrpc2.Method
}

func New(opts ...func(*Adaptor) *Adaptor) (adaptor *Adaptor, err error) {
	result := &Adaptor{
		impl: nil,
	}

	for i := range opts {
		result = opts[i](result)
	}

	verify := struct {
		Impl any `validate:"required"`
	}{
		Impl: adaptor.impl,
	}

	err = validator().Struct(verify)
	if err != nil {
		err = Trace(err)
		return
	}

	err = result.setupMethods()
	if err != nil {
		err = Trace(err)
		return
	}

	adaptor = result

	return
}

func (a *Adaptor) setupMethods() (err error) {
	return
}

func (a *Adaptor) Method(name string) jsonrpc2.Method {
	return a.methods[name]
}

var _ jsonrpc2.Adaptor = &Adaptor{}
