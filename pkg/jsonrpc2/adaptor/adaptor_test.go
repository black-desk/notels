package adaptor_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/black-desk/notels/pkg/jsonrpc2/adaptor"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

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

type TestStructWithTag struct {
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
	return "", errors.New("Method4")
}

var _ = Describe("JSON-RPC 2.0 Adaptor", func() {
	Context("Given a struct with methods as member", func() {
                TestStruct
	})
})
