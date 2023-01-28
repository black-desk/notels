package message_test

import (
	"encoding/json"

	. "github.com/black-desk/lib/go/ginkgo-helper"
	. "github.com/black-desk/notels/pkg/jsonrpc2/message"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSON-RPC 2.0 Request message unmarshal", func() {
	ContextTable("from valid raw json (for example %v)",
		func(raw string) {
			var req Request
			It("should success", func() {
				err := json.Unmarshal([]byte(raw), &req)
				Expect(err).To(BeNil())
			})
		},
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":{},
                        "id":"a"
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":[],
                        "id":null
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name"
                }`),
	)

	ContextTable("from invalid raw json (for example %v)",
		func(raw string) {
			var req Request
			It("should failed", func() {
				err := json.Unmarshal([]byte(raw), &req)
				Expect(err).NotTo(BeNil())
			})
		},
		ContextTableEntry(`{
                        "jsonrpc":"1.0",
                        "method":"some function name",
                        "params":{},
                        "id":"a"
                }`),
		ContextTableEntry(`{
                        "method":"some function name",
                        "params":{},
                        "id":"a"
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":"a",
                        "id":"a"
                }`),
	)
})

var _ = Describe("Request.IsNotification: Given a Request", func() {
	ContextTable("without id (for example %v)",
		func(raw string) {
			var req Request

			BeforeEach(func() {
				err := json.Unmarshal([]byte(raw), &req)
				Expect(err).To(BeNil())
			})

			It("should return true", func() {
				Expect(req.IsNotification()).To(BeTrue())
			})
		},
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":{}
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":[]
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name"
                }`),
	)

	ContextTable("with id (for example %v)",
		func(raw string) {
			var req Request

			BeforeEach(func() {
				err := json.Unmarshal([]byte(raw), &req)
				Expect(err).To(BeNil())
			})

			It("should return true", func() {
				Expect(req.IsNotification()).To(BeFalse())
			})
		},
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":{},
                        "id":1
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "params":[],
                        "id":"1"
                }`),
		ContextTableEntry(`{
                        "jsonrpc":"2.0",
                        "method":"some function name",
                        "id":1.0001
                }`),
	)
})
