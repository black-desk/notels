package message_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/black-desk/notels/pkg/jsonrpc2/message"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("message.ID", func() {
	Context("Given a", func() {
		Context("json number (for example:", func() {
			raws := []string{
				`100`,
				`100.10001`,
			}
			for i := range raws {
				buf := &bytes.Buffer{}
				json.Compact(buf, []byte(raws[i]))
				Context(fmt.Sprintf("%s)", buf.String()), func() {
					i := i
					var id ID
					var err error

					BeforeEach(func() {
						err = json.Unmarshal([]byte(raws[i]), &id)
					})

					It("should success to unmarshal", func() {
						Expect(err).To(BeNil())
					})
					It("should be number", func() {
						Expect(id.IsNum()).To(BeTrue())
					})
					It("should not be null", func() {
						Expect(id.IsNull()).To(BeFalse())
					})
					It("should not be string", func() {
						Expect(id.IsStr()).To(BeFalse())
					})
				})
			}
		})
		Context("json string (for example:", func() {
			raws := []string{
				`"null"`,
				`"aaa"`,
			}
			for i := range raws {
				buf := &bytes.Buffer{}
				json.Compact(buf, []byte(raws[i]))
				Context(fmt.Sprintf("%s)", buf.String()), func() {
					i := i
					var id ID
					var err error

					BeforeEach(func() {
						err = json.Unmarshal([]byte(raws[i]), &id)
					})

					It("should success to unmarshal", func() {
						Expect(err).To(BeNil())
					})
					It("should be string", func() {
						Expect(id.IsStr()).To(BeTrue())
					})
					It("should not be null", func() {
						Expect(id.IsNull()).To(BeFalse())
					})
					It("should not be number", func() {
						Expect(id.IsNum()).To(BeFalse())
					})
				})
			}
		})
	})
})
