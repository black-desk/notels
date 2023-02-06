package adaptor

import (
	"context"
	"reflect"

	"github.com/black-desk/notels/pkg/jsonrpc2"
	"github.com/black-desk/notels/pkg/jsonrpc2/message"
)

type option func(*Adaptor) *Adaptor

type MethodMap map[string]reflect.Value

func WithImpl(impl any) option {
	return WithMethodMap(methodMapFromImplement(impl))
}

func methodMapFromImplement(impl any) (ret MethodMap) {
	v := reflect.ValueOf(impl).Elem()
	t := v.Type()
	numberOfMethods := t.NumMethod()
	for i := 0; i < numberOfMethods; i++ {
		ret[t.Method(i).Name] = v.Method(i)
	}
	return
}

func WithTaggedStruct(taggedStruct any) option {
	return WithMethodMap(methodMapFromTaggedStruct(taggedStruct))
}
func methodMapFromTaggedStruct(taggedStruct any) (ret MethodMap) {
	v := reflect.ValueOf(taggedStruct).Elem()
	t := v.Type()
	numberOfMethods := t.NumField()
	for i := 0; i < numberOfMethods; i++ {
		jsonName, ok := t.Field(i).Tag.Lookup("jsonrpc2")
		if !ok {
			continue
		}
		ret[jsonName] = v.Field(i)
	}
	return
}

func WithMethodMap(methodMap MethodMap) option {
	return func(a *Adaptor) *Adaptor {
		a.methods = make(map[string]jsonrpc2.Method)
		for k := range methodMap {
			a.methods[k] = methodWrapper(methodMap[k])
		}
		return a
	}
}

func methodWrapper(fn reflect.Value) jsonrpc2.Method {
	return func(ctx context.Context, request *message.Request) (response *message.Response, err error) {
		ins := []reflect.Value{}
		{
			paramStart := 0
			if fn.Type().NumIn() > 0 && typeImplementsContext(fn.Type().In(0)) {
				ins = append(ins, reflect.ValueOf(ctx))
				paramStart += 1
			}

			if fn.Type().In(paramStart) {
			}
		}
		return
	}
}

func typeImplementsContext(t reflect.Type) bool {
	return t.Implements(reflect.TypeOf((*context.Context)(nil)).Elem())
}
