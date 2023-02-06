package message_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/black-desk/notels/pkg/jsonrpc2/message"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSON-RPC 2.0 Response message unmarshal", func() {
	Context("from valid raw json (for example:", func() {
		raws := []string{
			`{
                                "jsonrpc" : "2.0",
                                "result"  : true,
                                "id": "a"
                        }`,
			`{
                                "jsonrpc" : "2.0",
                                "result"  : [],
                                "id": 1.0
                        }`,
			`{
                                "jsonrpc" : "2.0",
                                "error"  : {
                                        "code": 0,
                                        "message": ""
                                },
                                "id": 123
                        }`,
		}
		for i := range raws {
			buf := &bytes.Buffer{}
			json.Compact(buf, []byte(raws[i]))
			Context(fmt.Sprintf("%s)", buf.String()), func() {
				i := i
				var resp Response
				It("should success", func() {
					err := json.Unmarshal([]byte(raws[i]), &resp)
					Expect(err).To(BeNil())
				})
			})
		}
	})
	Context("from invalid raw json (for example:", func() {
		raws := []string{
			`{
                                "jsonrpc" : "2.0",
                                "id": "a"
                        }`,
			`{
                                "jsonrpc" : "2.0",
                                "result"  : [],
                                "error"  : {
                                        "code": 0,
                                        "message": ""
                                },
                                "id" : null
                        }`,
		}
		for i := range raws {
			buf := &bytes.Buffer{}
			json.Compact(buf, []byte(raws[i]))
			Context(fmt.Sprintf("%s)", buf.String()), func() {
				i := i
				var resp Response
				It("should failed", func() {
					err := json.Unmarshal([]byte(raws[i]), &resp)
					Expect(err).NotTo(BeNil())
				})
			})
		}
	})
})
