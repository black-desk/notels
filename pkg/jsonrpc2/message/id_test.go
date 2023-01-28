package message_test

import (
	"encoding/json"

	. "github.com/black-desk/lib/go/ginkgo-helper"
	. "github.com/black-desk/notels/pkg/jsonrpc2/message"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("message.ID", func() {
	ContextTable("Given a json number (for example %v)",
		func(raw string) {
			var id ID
			var err error

			BeforeEach(func() {
				err = json.Unmarshal([]byte(raw), &id)
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
		},
		ContextTableEntry(`100`),
		ContextTableEntry(`100.10001`),
	)
	ContextTable("Given a json string (for example %v)",
		func(raw string) {
			var id ID
			var err error

			BeforeEach(func() {
				err = json.Unmarshal([]byte(raw), &id)
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
		},
		ContextTableEntry(`"null"`),
		ContextTableEntry(`"aaa"`),
	)
})
