package adaptor_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/black-desk/lib/go/ginkgo-helper"
	. "github.com/black-desk/notels/pkg/jsonrpc2/adaptor"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var err4 = errors.New("Method4")

func TestAdaptor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adaptor Suite")
}

type TestStruct struct {
	Method1 func() string
	Method2 func(string) string
	Method3 func() (string, error)
	Method4 func() (string, error)
}

type TestTaggedStruct struct {
	Method1 func() string          `jsonrpc2:"test/method1"`
	Method2 func(string) string    `jsonrpc2:"test/method2"`
	Method3 func() (string, error) `jsonrpc2:"test/method3"`
	Method4 func() (string, error) `jsonrpc2:"test/method4"`
}

type TestInterface interface {
	Method1() string
	Method2(string) string
	Method3() (string, error)
	Method4() (string, error)
}

var _ TestInterface = &TestInterfaceImpl{}

type TestInterfaceImpl struct{}

func (*TestInterfaceImpl) Method1() string {
	return "method1"
}

func (*TestInterfaceImpl) Method2(p string) string {
	return "method2 param=" + p
}

func (*TestInterfaceImpl) Method3() (string, error) {
	return "method3", nil
}

func (*TestInterfaceImpl) Method4() (string, error) {
	return "", err4
}

type TestMethodMap MethodMap

var _ = Describe("JSON-RPC 2.0 Adaptor", func() {
	Context("Given a map from string to reflect.Method", func() {
		var methodMap MethodMap
		BeforeEach(func() {
			methodMap = MethodMap{
				"method1": reflect.ValueOf(
					func() string { return "method1" },
				),
				"method2": reflect.ValueOf(
					func(p string) string {
						return "method2 param=" + p
					},
				),
				"method3": reflect.ValueOf(
					func() (string, error) {
						return "method3", nil
					},
				),
				"method4": reflect.ValueOf(
					func() (string, error) {
						return "", err4
					},
				),
			}
		})
		type option func(*Adaptor) *Adaptor
		commonCases := []*ContextTableEntryT{
			ContextTableEntry("method1", []reflect.Value{}, []any{"method1"}),
			ContextTableEntry("method2", []reflect.Value{reflect.ValueOf("param")}, []any{"method2 param=param"}),
			ContextTableEntry("method3", []reflect.Value{}, []any{"method3", nil}),
			ContextTableEntry("method4", []reflect.Value{}, []any{"", err4}),
		}
		ContextTable("then create an Adaptor with %s",
			func(opt option, cases []*ContextTableEntryT) {
				var adaptor *Adaptor
				var err error
				BeforeEach(func() {
					adaptor, err = New(opt)
				})
				It("should success", func() {
					Expect(adaptor).ToNot(BeNil())
					Expect(err).To(BeNil())
				})
				ContextTable("when we call Adaptor.Method(\"%s\")",
					func(method string, para []reflect.Value, ret []any) {
						var m reflect.Value
						var rets []reflect.Value
						BeforeEach(func() {
							m = reflect.ValueOf(adaptor.Method("method"))
							rets = m.Call(para)
						})
						It("should return as expected", func() {
							for i := range rets {
								if rets[i].Type() == reflect.TypeOf(ret[i]) {
									Expect(ret[i]).To(Equal(rets[i]))
								}
							}
						})
					},
					cases...,
				)
			},
			ContextTableEntry(WithMethodMap(methodMap), commonCases).WithFmt("method map"),
			ContextTableEntry(WithImpl(&TestInterfaceImpl{}), commonCases).WithFmt("struct"),
			ContextTableEntry(WithImpl(TestInterface(&TestInterfaceImpl{})), commonCases).WithFmt("interface"),
			ContextTableEntry(WithTaggedStruct(TestTaggedStruct{}), []*ContextTableEntryT{
				ContextTableEntry("test/method1", []reflect.Value{}, []any{"method1"}),
				ContextTableEntry("test/method2", []reflect.Value{reflect.ValueOf("param")}, []any{"method2 param=param"}),
				ContextTableEntry("test/method3", []reflect.Value{}, []any{"method3", nil}),
				ContextTableEntry("test/method4", []reflect.Value{}, []any{"", err4}),
			}).WithFmt("tagged struct"),
		)
	})
})
