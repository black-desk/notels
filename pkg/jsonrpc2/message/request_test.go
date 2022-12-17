package message_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/black-desk/notels/pkg/jsonrpc2/message"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSON-RPC 2.0 Request message unmarshal", func() {
	Context("from valid raw json (for example:", func() {
		raws := []string{
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":{},
                                "id":"a"
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":[],
                                "id":null
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name"
                        }`,
		}
		for i := range raws {
			buf := &bytes.Buffer{}
			json.Compact(buf, []byte(raws[i]))
			Context(fmt.Sprintf("%s)", buf.String()), func() {
				i := i
				var req Request
				It("should success", func() {
					err := json.Unmarshal([]byte(raws[i]), &req)
					Expect(err).To(BeNil())
				})
			})
		}
	})
	Context("from invalid raw json (for example:", func() {
		raws := []string{
			`{
                                "jsonrpc":"1.0",
                                "method":"some function name",
                                "params":{},
                                "id":"a"
                        }`,
			`{
                                "method":"some function name",
                                "params":{},
                                "id":"a"
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":"a",
                                "id":"a"
                        }`,
		}
		for i := range raws {
			buf := &bytes.Buffer{}
			json.Compact(buf, []byte(raws[i]))
			Context(fmt.Sprintf("%s)", buf.String()), func() {
				i := i
				var req Request
				It("should failed", func() {
					err := json.Unmarshal([]byte(raws[i]), &req)
					Expect(err).NotTo(BeNil())
				})
			})
		}
	})
})

var _ = Describe("Request.IsNotification: Given a Request", func() {
	Context("without id (for example:", func() {
		raws := []string{
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":{}
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":[]
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name"
                        }`,
		}
		for i := range raws {
			buf := &bytes.Buffer{}
			json.Compact(buf, []byte(raws[i]))
			Context(fmt.Sprintf("%s)", buf.String()), func() {
				i := i
				var req Request

				BeforeEach(func() {
					err := json.Unmarshal([]byte(raws[i]), &req)
					Expect(err).To(BeNil())
				})

				It("should return true", func() {
					Expect(req.IsNotification()).To(BeTrue())
				})
			})
		}
	})
	Context("with id (for example:", func() {
		raws := []string{
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":{},
                                "id":1
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "params":[],
                                "id":"1"
                        }`,
			`{
                                "jsonrpc":"2.0",
                                "method":"some function name",
                                "id":1.0001
                        }`,
		}
		for i := range raws {
			buf := &bytes.Buffer{}
			json.Compact(buf, []byte(raws[i]))
			Context(fmt.Sprintf("%s)", buf.String()), func() {
				i := i
				var req Request

				BeforeEach(func() {
					err := json.Unmarshal([]byte(raws[i]), &req)
					Expect(err).To(BeNil())
				})

				It("should return true", func() {
					Expect(req.IsNotification()).To(BeFalse())
				})
			})
		}
	})
})
